---
id: collect-261001-rattrapage/rattrapage/php-guide-16
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "decode"]
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [3276, 3501]
sha256: 0e24e47ee54bc699cce4f7a4aa596bfa81d2b6b48d65fdefd49d7f6ee6a132a7
---

# PHP 8 — Le guide complet du sysadmin qui héberge

            session_regenerate_id(true); // 🔒 anti-fixation
            $_SESSION['utilisateur_id'] = (int) $user['id'];
            $_SESSION['role'] = $user['role'];
            $_SESSION['derniere_activite'] = time();
            $_SESSION['empreinte'] = hash('sha256', $_SERVER['HTTP_USER_AGENT'] ?? '');

            header('Location: dashboard.php');
            exit;
        }

        // --- Échec : on loggue la tentative (sans dire pourquoi ça a échoué) ---
        $pdo->prepare('INSERT INTO login_tentatives (login, ip, date_heure) VALUES (?, ?, NOW())')
            ->execute([$login, $_SERVER['REMOTE_ADDR'] ?? '']);
        // Délai volontaire anti-énumération rapide
        sleep(1);
        $erreur = 'Identifiants invalides.'; // message GÉNÉRIQUE
    }
}
?>
<!DOCTYPE html>
<html lang="fr">
<head><meta charset="UTF-8"><title>Connexion</title></head>
<body>
<h1>Connexion</h1>
<?php if ($erreur): ?><p style="color:red"><?= htmlspecialchars($erreur) ?></p><?php endif; ?>
<form method="post" action="login.php" autocomplete="off">
    <input type="hidden" name="csrf_token" value="<?= htmlspecialchars($_SESSION['csrf_token']) ?>">
    <label>Identifiant : <input type="text" name="login" required autofocus></label><br>
    <label>Mot de passe : <input type="password" name="mdp" required></label><br>
    <button type="submit">Se connecter</button>
</form>
</body>
</html>
```

```php
<?php declare(strict_types=1);
// dashboard.php — page protégée (modèle à inclure en tête de chaque page privée)
session_start();
require __DIR__ . '/../bootstrap.php';

if (empty($_SESSION['utilisateur_id'])) {
    header('Location: login.php');
    exit;
}
// Timeout d'inactivité (section 41)
if (time() - ($_SESSION['derniere_activite'] ?? 0) > 1800) {
    require __DIR__ . '/logout.php'; // qui fait deconnexion() + redirection
}
$_SESSION['derniere_activite'] = time();
```

**Ce que ce login fait bien :** CSRF, `password_verify`, message générique, anti brute-force (5/15 min), `session_regenerate_id(true)`, re-hachage auto, timeout, empreinte UA. **Manque pour aller plus loin :** 2FA TOTP (lib `robthree/twofactorauth`), rate-limit par IP via fail2ban.

---

## 75. Cas pratique n°4 — mini API JSON

```php
<?php declare(strict_types=1);
// public/api/equipements.php — GET /api/equipements.php?type=onduleur
require __DIR__ . '/../../bootstrap.php';

header('Content-Type: application/json; charset=UTF-8');

// --- Auth simple par clé API (header) ---
$cle = $_SERVER['HTTP_X_API_KEY'] ?? '';
$stmt = $pdo->prepare('SELECT id FROM api_cles WHERE cle_hash = ? AND active = 1');
$stmt->execute([hash('sha256', $cle)]); // on stocke le hash, pas la clé en clair
if (!$stmt->fetch()) {
    http_response_code(401);
    echo json_encode(['erreur' => 'Clé API invalide.']);
    exit;
}

// --- Paramètres validés ---
$type = $_GET['type'] ?? null;
$typesAutorises = ['onduleur', 'clim', 'groupe', 'tdbt'];
if ($type !== null && !in_array($type, $typesAutorises, true)) {
    http_response_code(400);
    echo json_encode(['erreur' => 'Type invalide.']);
    exit;
}

$sql = 'SELECT id, reference, type, emplacement FROM equipements';
$params = [];
if ($type !== null) {
    $sql .= ' WHERE type = :type';
    $params['type'] = $type;
}
$stmt = $pdo->prepare($sql);
$stmt->execute($params);

echo json_encode(
    ['donnees' => $stmt->fetchAll()],
    JSON_UNESCAPED_UNICODE | JSON_THROW_ON_ERROR
);
```

Test : `curl -s -H "X-API-Key: <cle>" "https://appli.entreprise.fr/api/equipements.php?type=onduleur" | python3 -m json.tool`


---

## 76. Checklist de mise en production — le rituel

📋 **À dérouler à chaque déploiement / nouvelle appli :**

**Avant**
- [ ] Backup code + BDD à jour et **restauré une fois en test**
- [ ] `composer install --no-dev --optimize-autoloader` (pas `update`)
- [ ] `composer audit` : 0 vulnérabilité critique
- [ ] `php -l` sur tous les fichiers modifiés (ou `composer lint`)
- [ ] Tests automatisés au vert (si existants)
- [ ] `.env` prod présent, `APP_ENV=prod`, `APP_DEBUG=false`
- [ ] Fenêtre de maintenance annoncée si interruption

**Pendant**
- [ ] Déploiement : `git pull` / rsync par utilisateur `deploy` (jamais root sur les fichiers web)
- [ ] `composer dump-autoload --optimize`
- [ ] Migrations BDD jouées (script versionné, jamais à la main en prod)
- [ ] `systemctl reload php8.4-fpm` (vide OPcache)
- [ ] Droits : `chown -R appli:www-data`, `find -type d -exec chmod 750`, fichiers `640`, `public/` seul en lecture web

**Après**
- [ ] Page d'accueil + login + 1 parcours critique testés en HTTPS
- [ ] `curl -I` : en-têtes sécurité présents, pas de `X-Powered-By`
- [ ] Logs : 0 fatal/warning dans les 15 min (`tail -f` pool + nginx)
- [ ] Status FPM : `listen queue = 0`, workers OK
- [ ] Rollback prêt : commit précédent taggé, dump pré-déploiement conservé

---

## 77. Pense-bête de poche — une page

```php
// --- Base ---
$e = fn(string $s): string => htmlspecialchars($s, ENT_QUOTES, 'UTF-8'); // anti-XSS
$id = filter_input(INPUT_GET, 'id', FILTER_VALIDATE_INT);                 // entrée validée

// --- BDD : toujours préparé ---
$stmt = $pdo->prepare('SELECT * FROM t WHERE id = :id');
$stmt->execute(['id' => $id]);
$lignes = $stmt->fetchAll();

// --- CSRF ---
$_SESSION['csrf_token'] ??= bin2hex(random_bytes(32));                    // génération
hash_equals($_SESSION['csrf_token'], $_POST['csrf_token'] ?? '');         // vérif

// --- Mots de passe ---
$hash = password_hash($mdp, PASSWORD_DEFAULT);   // création
password_verify($mdp, $hash);                    // vérif

// --- Dates ---
$d = new DateTimeImmutable('now', new DateTimeZone('Europe/Paris'));
echo $d->format('Y-m-d H:i:s');

// --- Redirection ---
header('Location: /cible.php'); exit;

// --- JSON ---
$json = json_encode($data, JSON_UNESCAPED_UNICODE | JSON_THROW_ON_ERROR);
$data = json_decode($json, true, 512, JSON_THROW_ON_ERROR);

// --- Fichiers ---
file_put_contents($f, $contenu, FILE_APPEND | LOCK_EX);
$lignes = file($f, FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);
```

```bash
# --- Serveur ---
systemctl reload php8.4-fpm          # après modif ini/pool/déploiement
nginx -t && systemctl reload nginx   # après modif vhost
tail -f /var/log/php-fpm-appli-error.log
curl -s "http://127.0.0.1/fpm-status-appli?json" | python3 -m json.tool
php -l fichier.php                   # lint avant déploiement
composer install --no-dev --optimize-autoloader
```

---

## 78. Scripts utiles du sysadmin PHP

```php
<?php declare(strict_types=1);
// /usr/local/bin/health-check.php — cron toutes les 5 min, alerte si KO
$checks = [];

// 1. PHP-FPM répond ?
$ping = @file_get_contents('http://127.0.0.1/fpm-ping-appli');
$checks['fpm'] = ($ping === 'pong');

// 2. BDD joignable ?
try {
    $pdo = new PDO('mysql:host=127.0.0.1;dbname=appli;charset=utf8mb4', 'monitor', $_ENV['DB_MONITOR_PASS'],
        [PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION, PDO::ATTR_TIMEOUT => 3]);
    $pdo->query('SELECT 1');
    $checks['bdd'] = true;
} catch (Throwable) {
    $checks['bdd'] = false;
}

// 3. Espace disque < 85 % ?
$checks['disque'] = disk_free_space('/') / disk_total_space('/') > 0.15;

// 4. OPcache sain (hit rate > 90 %) ?
$st = opcache_get_status(false);
$checks['opcache'] = ($st['opcache_statistics']['opcache_hit_rate'] ?? 0) > 90;

$ko = array_keys(array_filter($checks, fn(bool $v): bool => !$v));
if ($ko !== []) {
    $msg = 'ALERTE appli : ' . implode(', ', $ko) . ' en échec';
    error_log($msg);
    mail('astreinte@entreprise.fr', '[ALERTE] appli', $msg);
    exit(1);
}
echo "OK\n";
```

```bash
#!/bin/bash
# purge-sessions.sh — nettoie les vieux fichiers de session si save_path en fichiers
find /var/lib/php/sessions -type f -cmin +1440 -delete
# (PHP le fait via le cron système ; utile si session.save_path custom par pool)
```

