---
id: collect-261001-rattrapage/rattrapage/php-guide-14
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [2878, 3071]
sha256: c440e4a3270ea07cfd5b130f70b8328aee1313b8d56e17e93caae21de4a1c35d
---

# 3. Tester en STAGING : installer PHP 8.x en parallèle (les paquets cohabitent : php7.4-fpm + php8.4-fpm)
sudo apt install php8.4-fpm php8.4-mysql php8.4-mbstring # + extensions de l'appli
# Dupliquer le pool FPM en 8.4, faire pointer un vhost de test dessus

# 4. Corriger les incompatibilités (section 70), lancer les tests
vendor/bin/phpunit

# 5. BASCULE : changer le fastcgi_pass du vhost prod vers le socket 8.4
#    Garder le pool 7.4 une semaine en cas de rollback

# 6. NETTOYAGE (après validation) : apt remove php7.4*
```

⚠️ **Jamais** de mise à jour PHP un vendredi à 17h. Fenêtre de maintenance annoncée, backup avant, rollback préparé (pool précédent conservé).

---

## 70. Migration 7.4 → 8.x — breaking changes à connaître

```php
<?php
// --- 1. Comparaisons string/nombre : "0 == 'foo'" ---
var_dump(0 == "foo"); // PHP 7 : true (!!) — PHP 8 : false ✅ (comparaison numérique seulement si la string est numérique)

// --- 2. Division par zéro ---
var_dump(1/0);        // PHP 7 : warning + INF — PHP 8 : DivisionByZeroError (exception !)
intdiv(1, 0);        // DivisionByZeroError dans les deux versions

// --- 3. match remplace switch (pas de fall-through, === strict) --- voir section 13

// --- 4. Paramètres nommés : les noms des paramètres des fonctions INTERNES peuvent changer ---
// htmlentities($s, ENT_QUOTES, 'UTF-8') reste valide, mais évite de nommer les args des fonctions natives.

// --- 5. each() supprimé (depuis 8.0), create_function() supprimé, money_format() supprimé ---
// Remplacements : foreach, closures, NumberFormatter (intl)

// --- 6. curly braces $str{0} : syntaxe supprimée → $str[0] ---

// --- 7. Type juggling en arithmétique : "10 euros" + 5 ---
// PHP 7 : 15 (notice) — PHP 8 : TypeError si non numérique ("euros" n'est pas numérique)

// --- 8. str_contains/str_starts_with/str_ends_with (8.0) : remplace strpos(...) !== false ---

// --- 9. get_magic_quotes_gpc() supprimé, magic quotes disparus depuis longtemps ---
// Si du vieux code fait stripslashes($_POST[...]) systématiquement : bug en 8.x → supprimer.

// --- 10. PDO : les erreurs par défaut restent silencieuses si ERRMODE non défini ---
// (pas un breaking change, mais l'occasion de passer en ERRMODE_EXCEPTION — section 44)
```

📋 **Ordre de migration conseillé :** 7.4 → 8.0 (gros changements) → 8.1/8.2/8.3/8.4 (incrémental, enums puis readonly puis...). Tester chaque palier.

---

## 71. Debug — Xdebug et les logs applicatifs

```bash
# Xdebug : UNIQUEMENT en dev/staging, JAMAIS en prod (divise les perfs par 2-3)
sudo apt install php8.4-xdebug
# /etc/php/8.4/fpm/conf.d/20-xdebug.ini — en DEV :
# xdebug.mode=debug
# xdebug.start_with_request=yes
# xdebug.client_host=127.0.0.1
# En prod : xdebug.mode=off (ou paquet désinstallé)
```

```php
<?php declare(strict_types=1);

// --- En prod, le "debug" = logs structurés ---
// Monolog (standard de fait) : composer require monolog/monolog
use Monolog\Logger;
use Monolog\Handler\StreamHandler;

$log = new Logger('appli');
$log->pushHandler(new StreamHandler('/var/log/appli.log', Logger::INFO));

$log->info('Intervention créée', ['id' => $id, 'site' => $site, 'user' => $userId]);
$log->warning('Charge onduleur élevée', ['charge' => 87.5, 'seuil' => 80]);
$log->error('Échec paiement', ['exception' => $e->getMessage()]);
// Contexte en tableau = requêtes grep/jq efficaces sur les logs JSON

// --- Handler JSON pour centralisation (Loki/ELK) ---
use Monolog\Handler\StreamHandler;
use Monolog\Formatter\JsonFormatter;
$h = new StreamHandler('/var/log/appli.json.log', Logger::DEBUG);
$h->setFormatter(new JsonFormatter());

// --- Debug rapide SANS xdebug : backtrace ciblée ---
debug_print_backtrace(DEBUG_BACKTRACE_IGNORE_ARGS); // qui appelle cette fonction ?
```

**En staging avec Xdebug :** step-debug dans VS Code (extension PHP Debug, `launch.json` port 9003), profiler avec `xdebug.mode=profile` + cachegrind pour trouver les fonctions gourmandes.

---

## 72. Cas pratique n°1 — formulaire de contact sécurisé

Le classique, fait **correctement** : validation, CSRF, PRG, anti-spam honeypot, envoi mail sécurisé.

```php
<?php declare(strict_types=1);
// public/contact.php
session_start();
require __DIR__ . '/../bootstrap.php'; // error handlers, config, PDO (sections 34, 44)

/* ---------- CSRF ---------- */
if (empty($_SESSION['csrf_token'])) {
    $_SESSION['csrf_token'] = bin2hex(random_bytes(32));
}

/* ---------- TRAITEMENT ---------- */
$erreurs = [];
$valeurs = ['nom' => '', 'email' => '', 'message' => ''];

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    // 1. CSRF
    if (!hash_equals($_SESSION['csrf_token'] ?? '', $_POST['csrf_token'] ?? '')) {
        http_response_code(403);
        exit('Jeton invalide.');
    }

    // 2. Honeypot anti-bot (champ caché : rempli = robot)
    if (!empty($_POST['site_web'])) {
        // On fait semblant que ça marche (ne pas informer le bot)
        header('Location: contact.php?ok=1');
        exit;
    }

    // 3. Validation
    $valeurs['nom']     = trim($_POST['nom'] ?? '');
    $valeurs['email']    = trim($_POST['email'] ?? '');
    $valeurs['message']  = trim($_POST['message'] ?? '');

    if (mb_strlen($valeurs['nom']) < 2 || mb_strlen($valeurs['nom']) > 80) {
        $erreurs['nom'] = 'Nom requis (2-80 caractères).';
    }
    if (!filter_var($valeurs['email'], FILTER_VALIDATE_EMAIL)) {
        $erreurs['email'] = 'E-mail invalide.';
    }
    if (mb_strlen($valeurs['message']) < 10 || mb_strlen($valeurs['message']) > 5000) {
        $erreurs['message'] = 'Message entre 10 et 5000 caractères.';
    }

    // 4. Enregistrement + envoi
    if ($erreurs === []) {
        $stmt = $pdo->prepare(
            'INSERT INTO contacts (nom, email, message, cree_le) VALUES (:nom, :email, :message, NOW())'
        );
        $stmt->execute([
            'nom'     => $valeurs['nom'],
            'email'   => $valeurs['email'],
            'message' => $valeurs['message'], // stocké BRUT, échappé à l'affichage
        ]);

        // E-mail : en-têtes SANS injection (pas de \r\n utilisateur dans les headers)
        $sujet = 'Nouveau contact : ' . mb_substr(str_replace(["\r", "\n"], '', $valeurs['nom']), 0, 60);
        mail('contact@entreprise.fr', $sujet, $valeurs['message'],
            "From: noreply@entreprise.fr\r\nReply-To: noreply@entreprise.fr\r\nContent-Type: text/plain; charset=UTF-8");
        // Note : Reply-To = adresse saisie UNIQUEMENT après validation FILTER_VALIDATE_EMAIL (fait plus haut)

        header('Location: contact.php?ok=1'); // PRG
        exit;
    }
}
?>
<!DOCTYPE html>
<html lang="fr">
<head><meta charset="UTF-8"><title>Contact</title></head>
<body>
<?php if (isset($_GET['ok'])): ?>
    <p>Message envoyé, merci.</p>
<?php else: ?>
    <form method="post" action="contact.php" novalidate>
        <input type="hidden" name="csrf_token" value="<?= htmlspecialchars($_SESSION['csrf_token']) ?>">
        <!-- honeypot : invisible pour l'humain -->
        <div style="display:none"><input type="text" name="site_web" value="" autocomplete="off"></div>

        <label>Nom : <input type="text" name="nom" value="<?= htmlspecialchars($valeurs['nom']) ?>" required></label>
        <?php if (isset($erreurs['nom'])): ?><span><?= htmlspecialchars($erreurs['nom']) ?></span><?php endif; ?>

        <label>E-mail : <input type="email" name="email" value="<?= htmlspecialchars($valeurs['email']) ?>" required></label>
        <?php if (isset($erreurs['email'])): ?><span><?= htmlspecialchars($erreurs['email']) ?></span><?php endif; ?>

        <label>Message : <textarea name="message" required><?= htmlspecialchars($valeurs['message']) ?></textarea></label>
        <?php if (isset($erreurs['message'])): ?><span><?= htmlspecialchars($erreurs['message']) ?></span><?php endif; ?>

        <button type="submit">Envoyer</button>
    </form>
<?php endif; ?>
</body>
</html>
```

