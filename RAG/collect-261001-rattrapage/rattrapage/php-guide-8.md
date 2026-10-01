---
id: collect-261001-rattrapage/rattrapage/php-guide-8
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "decode", "exploit"]
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [1647, 1868]
sha256: 8ca7ba08e316841cfd9c52fd234bc2177ff6419ce24b95d0291b28ee1d25cef7
---

# PHP 8 — Le guide complet du sysadmin qui héberge

// --- Regex utiles ---
preg_match('/^[\p{L}\p{N} .\-_]+$/u', $s);   // texte simple unicode
preg_match('/^[A-Z0-9-]{3,20}$/', $ref);     // référence équipement
```

| Filtre | Usage typique |
|---|---|
| `FILTER_VALIDATE_EMAIL` | e-mails |
| `FILTER_VALIDATE_INT` + `min_range`/`max_range` | IDs, quantités, durées |
| `FILTER_VALIDATE_IP` | adresses IP (whitelist réseau) |
| `FILTER_VALIDATE_URL` | URLs (limité : compléter par une whitelist de schémas `https`) |
| `FILTER_VALIDATE_BOOLEAN` + `FILTER_NULL_ON_FAILURE` | cases à cocher / flags |

---

## 40. Sessions — le mécanisme

```php
<?php declare(strict_types=1);
session_start(); // ⚠️ AVANT tout output HTML (en-têtes HTTP)

// --- Écriture / lecture ---
$_SESSION['utilisateur_id'] = 42;
$_SESSION['role'] = 'technicien';
$_SESSION['derniere_activite'] = time();

echo $_SESSION['role'] ?? 'invité';

// --- Déconnexion PROPRE ---
function deconnexion(): void
{
    $_SESSION = [];                        // vide le tableau
    if (ini_get('session.use_cookies')) {
        $params = session_get_cookie_params();
        setcookie(session_name(), '', time() - 42000, $params['path'], $params['domain'],
            $params['secure'], $params['httponly']); // supprime le cookie
    }
    session_destroy();                     // détruit la session côté serveur
}

// --- Où sont stockées les sessions ? ---
// Par défaut : fichiers dans /var/lib/php/sessions (session.save_path).
// En multi-serveurs / haute dispo : Redis (php8.4-redis) :
// session.save_handler = redis
// session.save_path = "tcp://127.0.0.1:6379"
```

**Cycle de vie :** `session_start()` lit le cookie `PHPSESSID` → charge le fichier correspondant → `$_SESSION` disponible. À la fin du script, PHP **verrouille** le fichier pendant toute la requête : deux requêtes AJAX simultanées du même user se **bloquent** mutuellement.

**Astuce perf :** `session_write_close()` dès qu'on n'écrit plus en session (libère le verrou) :

```php
session_start();
$userId = $_SESSION['utilisateur_id'] ?? null;
session_write_close(); // libère le verrou, lecture seule ensuite
// ... traitement long (appel API, génération PDF) sans bloquer les autres onglets ...
```

---

## 41. Sécurité des sessions — anti-hijacking

```php
<?php declare(strict_types=1);
session_start();

// --- 1. Régénérer l'ID à chaque changement de privilège (anti-fixation) ---
function login(int $userId): void
{
    session_regenerate_id(true); // true = supprime l'ancien fichier de session
    $_SESSION['utilisateur_id'] = $userId;
    $_SESSION['connecte'] = true;
}

// --- 2. Timeout d'inactivité (30 min) ---
$timeout = 1800;
if (isset($_SESSION['derniere_activite']) && time() - $_SESSION['derniere_activite'] > $timeout) {
    deconnexion();
    header('Location: login.php?timeout=1');
    exit;
}
$_SESSION['derniere_activite'] = time();

// --- 3. Lier (souplement) la session au client : user-agent + début d'IP ---
// ⚠️ IP complète = faux positifs (4G, VPN). On lie le /24 ou juste l'UA.
$empreinte = hash('sha256', $_SERVER['HTTP_USER_AGENT'] ?? '');
if (isset($_SESSION['empreinte']) && !hash_equals($_SESSION['empreinte'], $empreinte)) {
    deconnexion(); // user-agent changé en cours de session = suspect
    exit;
}
$_SESSION['empreinte'] ??= $empreinte;

// --- 4. php.ini (rappel section 5) ---
// session.cookie_httponly=1, session.cookie_secure=1, session.cookie_samesite=Lax,
// session.use_strict_mode=1, session.use_only_cookies=1
```

📋 **Checklist session :** `regenerate_id(true)` au login, timeout d'inactivité, empreinte UA, cookies `HttpOnly`+`Secure`+`SameSite`, `deconnexion()` complète, stockage Redis si plusieurs serveurs.

---

## 42. Cookies — posés proprement

```php
<?php declare(strict_types=1);

// --- setcookie() : AVANT tout output, tableau d'options (PHP 7.3+) ---
setcookie('theme', 'sombre', [
    'expires'  => time() + 86400 * 30, // 30 jours
    'path'     => '/',
    'domain'   => '',                   // '' = domaine courant
    'secure'   => true,                 // HTTPS uniquement 🔒
    'httponly' => true,                 // invisible pour JavaScript (anti-vol via XSS) 🔒
    'samesite' => 'Lax',                // 'Strict' = plus sûr, 'Lax' = compromis, 'None' = cross-site (requiert Secure)
]);

// --- Lecture : TOUJOURS considérer comme hostile ---
$theme = $_COOKIE['theme'] ?? 'clair';
if (!in_array($theme, ['clair', 'sombre'], true)) {
    $theme = 'clair'; // whitelist
}

// --- Suppression ---
setcookie('theme', '', ['expires' => time() - 3600, 'path' => '/']);

// --- Cookie signé : détecter la falsification (préférences, panier non critique) ---
define('COOKIE_SECRET', $_ENV['COOKIE_SECRET']); // 32 octets aléatoires, dans .env JAMAIS dans le code

function signerCookie(string $valeur): string
{
    $sig = hash_hmac('sha256', $valeur, COOKIE_SECRET);
    return base64_encode($valeur) . '.' . $sig;
}

function verifierCookie(string $cookie): ?string
{
    $parts = explode('.', $cookie);
    if (count($parts) !== 2) {
        return null;
    }
    [$b64, $sig] = $parts;
    $valeur = base64_decode($b64, true);
    if ($valeur === false) {
        return null;
    }
    // hash_equals : comparaison en temps constant (anti timing attack)
    return hash_equals(hash_hmac('sha256', $valeur, COOKIE_SECRET), $sig) ? $valeur : null;
}
```

🔒 **Jamais** de données sensibles en clair dans un cookie (rôle, user_id non signé). Le cookie de session suffit ; le reste = BDD.

---

## 43. Upload de fichiers — la procédure sécurisée

L'upload est **la faille la plus exploitée** : un `shell.php` uploadé = serveur compromis. Procédure stricte :

```php
<?php declare(strict_types=1);

// Formulaire : <form method="post" enctype="multipart/form-data">
//              <input type="file" name="rapport">

function traiterUpload(array $file): string
{
    // 1. Erreur d'upload ?
    if (($file['error'] ?? UPLOAD_ERR_NO_FILE) !== UPLOAD_ERR_OK) {
        throw new RuntimeException("Échec de l'envoi (code {$file['error']}).");
    }

    // 2. Taille (défense en profondeur : php.ini + ici)
    $maxOctets = 5 * 1024 * 1024; // 5 Mo
    if ($file['size'] > $maxOctets) {
        throw new RuntimeException("Fichier trop volumineux.");
    }

    // 3. Type MIME RÉEL (jamais $file['type'] — fourni par le client !)
    $finfo = new finfo(FILEINFO_MIME_TYPE);
    $mime = $finfo->file($file['tmp_name']);
    $autorises = [
        'application/pdf' => 'pdf',
        'image/jpeg'      => 'jpg',
        'image/png'       => 'png',
    ];
    if (!isset($autorises[$mime])) {
        throw new RuntimeException("Type de fichier non autorisé ($mime).");
    }

    // 4. Nom de fichier : ON LE GÉNÈRE (jamais le nom d'origine)
    $nomFinal = bin2hex(random_bytes(16)) . '.' . $autorises[$mime];

    // 5. Destination HORS de la racine web si possible
    $destination = '/var/www/uploads/' . $nomFinal; // /var/www/uploads PAS servi par nginx

    // 6. Déplacement sécurisé (vérifie que c'est bien un upload HTTP)
    if (!move_uploaded_file($file['tmp_name'], $destination)) {
        throw new RuntimeException("Impossible d'enregistrer le fichier.");
    }
    chmod($destination, 0640);

    return $nomFinal; // à stocker en BDD
}

// Usage :
try {
    $nom = traiterUpload($_FILES['rapport']);
    // INSERT INTO documents (nom_stocke, ...) VALUES (...)
} catch (RuntimeException $e) {
    http_response_code(400);
    echo htmlspecialchars($e->getMessage());
}
```

📋 **Checklist upload :** `enctype` correct, test `error`, taille, MIME réel via `finfo` (whitelist extension→MIME), nom **généré** (pas celui du client), dossier **hors racine web** (ou `deny all` nginx + pas d'exécution PHP), `move_uploaded_file` uniquement, droits `0640`.

---

## 44. PDO — connexion propre

**PDO = la seule API BDD à utiliser** (`mysql_*` supprimé depuis PHP 7, `mysqli` procédural = à oublier).

