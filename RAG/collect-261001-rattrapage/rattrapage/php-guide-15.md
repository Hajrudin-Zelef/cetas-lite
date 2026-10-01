---
id: collect-261001-rattrapage/rattrapage/php-guide-15
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [3072, 3275]
sha256: f34206c8f37811955c104113dd6cc804998a4e03ab8052cbc4d28a4e1585125d
---

# PHP 8 — Le guide complet du sysadmin qui héberge

**Ce que ce code fait bien :** CSRF, validation serveur stricte, honeypot, PRG, stockage brut + échappement affichage, en-têtes mail sans injection CRLF, `htmlspecialchars` partout.

---

## 73. Cas pratique n°2 — mini CRUD avec PDO

CRUD complet sur `equipements(id, reference, type, emplacement)` — lister, créer, modifier, supprimer, avec CSRF et PRG.

```php
<?php declare(strict_types=1);
// public/equipements.php — LISTE + SUPPRESSION
session_start();
require __DIR__ . '/../bootstrap.php';

if (empty($_SESSION['csrf_token'])) {
    $_SESSION['csrf_token'] = bin2hex(random_bytes(32));
}

// --- SUPPRESSION (POST + CSRF, jamais GET) ---
if ($_SERVER['REQUEST_METHOD'] === 'POST' && ($_POST['action'] ?? '') === 'supprimer') {
    if (!hash_equals($_SESSION['csrf_token'] ?? '', $_POST['csrf_token'] ?? '')) {
        http_response_code(403);
        exit('Jeton invalide.');
    }
    $id = filter_input(INPUT_POST, 'id', FILTER_VALIDATE_INT);
    if ($id !== false && $id !== null) {
        $stmt = $pdo->prepare('DELETE FROM equipements WHERE id = ?');
        $stmt->execute([$id]);
    }
    header('Location: equipements.php');
    exit;
}

// --- LISTE paginée ---
$page = max(1, (int) ($_GET['page'] ?? 1));
$parPage = 20;
$total = (int) $pdo->query('SELECT COUNT(*) FROM equipements')->fetchColumn();
$pages = (int) ceil($total / $parPage);

$stmt = $pdo->prepare('SELECT id, reference, type, emplacement FROM equipements ORDER BY id DESC LIMIT :lim OFFSET :off');
$stmt->bindValue(':lim', $parPage, PDO::PARAM_INT);
$stmt->bindValue(':off', ($page - 1) * $parPage, PDO::PARAM_INT);
$stmt->execute();
$equipements = $stmt->fetchAll();
?>
<h1>Équipements (<?= $total ?>)</h1>
<p><a href="equipement-form.php">+ Nouvel équipement</a></p>
<table border="1">
    <tr><th>Référence</th><th>Type</th><th>Emplacement</th><th>Actions</th></tr>
    <?php foreach ($equipements as $eq): ?>
    <tr>
        <td><?= htmlspecialchars($eq['reference']) ?></td>
        <td><?= htmlspecialchars($eq['type']) ?></td>
        <td><?= htmlspecialchars($eq['emplacement']) ?></td>
        <td>
            <a href="equipement-form.php?id=<?= (int) $eq['id'] ?>">Modifier</a>
            <form method="post" style="display:inline"
                  onsubmit="return confirm('Supprimer <?= htmlspecialchars($eq['reference']) ?> ?');">
                <input type="hidden" name="csrf_token" value="<?= htmlspecialchars($_SESSION['csrf_token']) ?>">
                <input type="hidden" name="action" value="supprimer">
                <input type="hidden" name="id" value="<?= (int) $eq['id'] ?>">
                <button type="submit">Supprimer</button>
            </form>
        </td>
    </tr>
    <?php endforeach; ?>
</table>
<p>Page <?= $page ?> / <?= max(1, $pages) ?></p>
```

```php
<?php declare(strict_types=1);
// public/equipement-form.php — CRÉATION + MODIFICATION
session_start();
require __DIR__ . '/../bootstrap.php';

if (empty($_SESSION['csrf_token'])) {
    $_SESSION['csrf_token'] = bin2hex(random_bytes(32));
}

$id = filter_input(INPUT_GET, 'id', FILTER_VALIDATE_INT);
$eq = ['reference' => '', 'type' => '', 'emplacement' => ''];
if ($id) {
    $stmt = $pdo->prepare('SELECT reference, type, emplacement FROM equipements WHERE id = ?');
    $stmt->execute([$id]);
    $eq = $stmt->fetch() ?: $eq;
}

$erreurs = [];
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    if (!hash_equals($_SESSION['csrf_token'] ?? '', $_POST['csrf_token'] ?? '')) {
        http_response_code(403);
        exit('Jeton invalide.');
    }

    $eq['reference']   = trim($_POST['reference'] ?? '');
    $eq['type']        = trim($_POST['type'] ?? '');
    $eq['emplacement'] = trim($_POST['emplacement'] ?? '');

    if (!preg_match('/^[A-Z0-9-]{3,20}$/', $eq['reference'])) {
        $erreurs[] = 'Référence : 3-20 caractères (A-Z, 0-9, tiret).';
    }
    $typesAutorises = ['onduleur', 'clim', 'groupe', 'tdbt'];
    if (!in_array($eq['type'], $typesAutorises, true)) {
        $erreurs[] = 'Type invalide.';
    }
    if (mb_strlen($eq['emplacement']) > 100) {
        $erreurs[] = 'Emplacement trop long.';
    }

    if ($erreurs === []) {
        try {
            if ($id) {
                $stmt = $pdo->prepare(
                    'UPDATE equipements SET reference = :r, type = :t, emplacement = :e WHERE id = :id'
                );
                $stmt->execute(['r' => $eq['reference'], 't' => $eq['type'], 'e' => $eq['emplacement'], 'id' => $id]);
            } else {
                $stmt = $pdo->prepare(
                    'INSERT INTO equipements (reference, type, emplacement) VALUES (:r, :t, :e)'
                );
                $stmt->execute(['r' => $eq['reference'], 't' => $eq['type'], 'e' => $eq['emplacement']]);
            }
            header('Location: equipements.php');
            exit;
        } catch (PDOException $e) {
            if ($e->getCode() === '23000') {
                $erreurs[] = 'Cette référence existe déjà.';
            } else {
                throw $e; // remonte au handler global (section 34)
            }
        }
    }
}
?>
<h1><?= $id ? 'Modifier' : 'Nouvel' ?> équipement</h1>
<?php foreach ($erreurs as $err): ?><p style="color:red"><?= htmlspecialchars($err) ?></p><?php endforeach; ?>
<form method="post">
    <input type="hidden" name="csrf_token" value="<?= htmlspecialchars($_SESSION['csrf_token']) ?>">
    <label>Référence : <input name="reference" value="<?= htmlspecialchars($eq['reference']) ?>" required></label><br>
    <label>Type :
        <select name="type">
            <?php foreach (['onduleur', 'clim', 'groupe', 'tdbt'] as $t): ?>
                <option value="<?= $t ?>" <?= $t === $eq['type'] ? 'selected' : '' ?>><?= $t ?></option>
            <?php endforeach; ?>
        </select>
    </label><br>
    <label>Emplacement : <input name="emplacement" value="<?= htmlspecialchars($eq['emplacement']) ?>" maxlength="100"></label><br>
    <button type="submit">Enregistrer</button>
</form>
```

---

## 74. Cas pratique n°3 — page de login sécurisée

```php
<?php declare(strict_types=1);
// public/login.php
session_start();
require __DIR__ . '/../bootstrap.php';

if (empty($_SESSION['csrf_token'])) {
    $_SESSION['csrf_token'] = bin2hex(random_bytes(32));
}

// Déjà connecté ? → espace privé
if (!empty($_SESSION['utilisateur_id'])) {
    header('Location: dashboard.php');
    exit;
}

$erreur = null;
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    if (!hash_equals($_SESSION['csrf_token'] ?? '', $_POST['csrf_token'] ?? '')) {
        http_response_code(403);
        exit('Jeton invalide.');
    }

    $login = trim($_POST['login'] ?? '');
    $mdp   = $_POST['mdp'] ?? '';

    // --- Anti brute-force : compteur par login (table login_tentatives) ---
    $stmt = $pdo->prepare(
        "SELECT COUNT(*) FROM login_tentatives WHERE login = ? AND date_heure > DATE_SUB(NOW(), INTERVAL 15 MINUTE)"
    );
    $stmt->execute([$login]);
    if ((int) $stmt->fetchColumn() >= 5) {
        $erreur = 'Compte temporairement verrouillé (trop de tentatives). Réessayez dans 15 minutes.';
    } else {
        $stmt = $pdo->prepare('SELECT id, hash_mdp, role, actif FROM utilisateurs WHERE login = ?');
        $stmt->execute([$login]);
        $user = $stmt->fetch();

        if ($user && (int) $user['actif'] === 1 && password_verify($mdp, $user['hash_mdp'])) {
            // --- Succès ---
            $pdo->prepare('DELETE FROM login_tentatives WHERE login = ?')->execute([$login]);

            // Re-hachage si nécessaire
            if (password_needs_rehash($user['hash_mdp'], PASSWORD_DEFAULT)) {
                $pdo->prepare('UPDATE utilisateurs SET hash_mdp = ? WHERE id = ?')
                    ->execute([password_hash($mdp, PASSWORD_DEFAULT), $user['id']]);
            }

