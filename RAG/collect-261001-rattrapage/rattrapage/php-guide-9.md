---
id: collect-261001-rattrapage/rattrapage/php-guide-9
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [1869, 2081]
sha256: c3dc8d200904dd7795edf99438412a49d155ce24e02902d49f546c71897232e9
---

# PHP 8 — Le guide complet du sysadmin qui héberge

```php
<?php declare(strict_types=1);

function creerPdo(array $config): PDO
{
    $dsn = sprintf(
        'mysql:host=%s;port=%d;dbname=%s;charset=utf8mb4',
        $config['host'], $config['port'], $config['dbname']
    );

    $pdo = new PDO($dsn, $config['user'], $config['pass'], [
        // 🔒 Les 3 options vitales :
        PDO::ATTR_ERRMODE            => PDO::ERRMODE_EXCEPTION, // erreurs = exceptions (pas de die silencieux)
        PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,       // fetch = tableau associatif
        PDO::ATTR_EMULATE_PREPARES   => false,                  // vraies requêtes préparées côté serveur
    ]);

    return $pdo;
}

// --- Config depuis l'environnement (jamais en dur !) ---
$pdo = creerPdo([
    'host'   => $_ENV['DB_HOST'] ?? '127.0.0.1',
    'port'   => (int) ($_ENV['DB_PORT'] ?? 3306),
    'dbname' => $_ENV['DB_NAME'] ?? 'appli',
    'user'   => $_ENV['DB_USER'] ?? 'appli',
    'pass'   => $_ENV['DB_PASS'] ?? '',
]);

// --- PostgreSQL : DSN différent ---
// $dsn = 'pgsql:host=127.0.0.1;port=5432;dbname=appli;';
// --- SQLite (petits outils, tests) ---
// $pdo = new PDO('sqlite:/var/www/appli/data/appli.sqlite');
```

🔒 **Identifiants :** variables d'environnement ou fichier `.env` **hors racine web**, droits `0600`, jamais dans le code, jamais dans git. `charset=utf8mb4` obligatoire (emoji + sécurité).

---

## 45. Requêtes préparées — le seul moyen

```php
<?php declare(strict_types=1);

// --- SELECT avec paramètre nommé ---
$stmt = $pdo->prepare('SELECT id, site, statut FROM interventions WHERE site = :site AND statut = :statut');
$stmt->execute(['site' => $site, 'statut' => 'planifiee']);
$interventions = $stmt->fetchAll();          // tableau de tableaux
$premiere      = $stmt->fetch();             // une ligne ou false

// --- Paramètre positionnel (?) ---
$stmt = $pdo->prepare('SELECT * FROM equipements WHERE id = ?');
$stmt->execute([$id]);

// --- INSERT : lastInsertId ---
$stmt = $pdo->prepare('INSERT INTO interventions (site, description, priorite, cree_le)
                       VALUES (:site, :description, :priorite, NOW())');
$stmt->execute([
    'site'        => $site,
    'description' => $description,
    'priorite'    => $priorite->value, // enum adossée → string
]);
$id = (int) $pdo->lastInsertId();

// --- UPDATE / DELETE : rowCount ---
$stmt = $pdo->prepare('UPDATE interventions SET statut = :statut WHERE id = :id');
$stmt->execute(['statut' => 'cloturee', 'id' => $id]);
echo $stmt->rowCount() . " ligne(s) modifiée(s)";

// --- fetchColumn : une seule valeur ---
$nb = $pdo->query('SELECT COUNT(*) FROM interventions')->fetchColumn();

// --- Itérer sans tout charger en mémoire (gros volumes) ---
$stmt = $pdo->prepare('SELECT * FROM logs ORDER BY id');
$stmt->execute();
while ($ligne = $stmt->fetch()) {           // curseur : ligne par ligne
    traiterLigne($ligne);
}
```

**Pourquoi c'est LE seul moyen :** la requête et les données sont envoyées **séparément** au serveur SQL. Même si `$site` contient `' OR '1'='1`, c'est traité comme une **valeur**, jamais comme du code SQL. Aucun échappement manuel ne vaut ça.

---

## 46. Transactions PDO — tout ou rien

```php
<?php declare(strict_types=1);

// --- Cas : clôturer une intervention = MAJ intervention + INSERT rapport + décrément stock ---
try {
    $pdo->beginTransaction();

    $pdo->prepare('UPDATE interventions SET statut = :s WHERE id = :id')
        ->execute(['s' => 'cloturee', 'id' => $id]);

    $pdo->prepare('INSERT INTO rapports (intervention_id, contenu) VALUES (:id, :contenu)')
        ->execute(['id' => $id, 'contenu' => $contenu]);

    $pdo->prepare('UPDATE stock SET quantite = quantite - :q WHERE piece = :p')
        ->execute(['q' => $quantite, 'p' => $piece]);

    $pdo->commit(); // ✅ tout a réussi : on valide
    echo "Intervention clôturée.";
} catch (Throwable $e) {
    $pdo->rollBack(); // ❌ échec quelque part : on annule TOUT (aucune écriture partielle)
    error_log("Transaction échouée : " . $e->getMessage());
    http_response_code(500);
    echo "Erreur lors de la clôture.";
}

// --- Vérifier qu'on n'est pas déjà en transaction (code imbriqué) ---
if (!$pdo->inTransaction()) {
    $pdo->beginTransaction();
}
```

**Quand :** dès que **2+ écritures** dépendent l'une de l'autre. Sans transaction : panne entre les deux = données incohérentes (intervention clôturée sans rapport).

---

## 47. PDO — erreurs, timeouts, bonnes pratiques

```php
<?php declare(strict_types=1);

// --- DSN avec timeout de connexion (évite les workers FPM bloqués) ---
$dsn = 'mysql:host=127.0.0.1;port=3306;dbname=appli;charset=utf8mb4;connect_timeout=5';

// --- Attraper précisément ---
try {
    $stmt = $pdo->prepare('INSERT INTO equipements (reference) VALUES (:ref)');
    $stmt->execute(['ref' => $ref]);
} catch (PDOException $e) {
    // 23000 = violation de contrainte (doublon UNIQUE, clé étrangère...)
    if ($e->getCode() === '23000') {
        http_response_code(409);
        echo "Cette référence existe déjà.";
    } else {
        error_log("PDO : " . $e->getMessage()); // log complet côté serveur
        http_response_code(500);
        echo "Erreur base de données.";          // message générique côté client 🔒
    }
}

// --- LIKE avec paramètre : le % est dans la VALEUR, pas dans la requête ---
$stmt = $pdo->prepare('SELECT * FROM equipements WHERE reference LIKE :rech');
$stmt->execute(['rech' => '%' . $recherche . '%']);

// --- IN (...) avec un nombre variable de valeurs ---
$ids = [3, 7, 12];
$placeholders = implode(',', array_fill(0, count($ids), '?')); // "?,?,?"
$stmt = $pdo->prepare("SELECT * FROM equipements WHERE id IN ($placeholders)");
$stmt->execute($ids); // ✅ les placeholders sont générés par NOUS, les valeurs liées

// --- LIMIT/OFFSET : caster en int (ou bind en PARAM_INT) ---
$page = max(1, (int) ($_GET['page'] ?? 1));
$parPage = 20;
$stmt = $pdo->prepare('SELECT * FROM interventions ORDER BY id DESC LIMIT :lim OFFSET :off');
$stmt->bindValue(':lim', $parPage, PDO::PARAM_INT);
$stmt->bindValue(':off', ($page - 1) * $parPage, PDO::PARAM_INT);
$stmt->execute();
```

⚠️ **À bannir :** `"SELECT * FROM t WHERE id = $id"` (interpolation), `PDO::query()` avec des variables, `ATTR_EMULATE_PREPARES => true` en prod.

---

## 48. Sécurité : injections SQL — le dossier complet

```php
<?php declare(strict_types=1);

// --- ❌ VULNÉRABLE : concaténation ---
$id = $_GET['id']; // attaquant : 1 OR 1=1
$result = $pdo->query("SELECT * FROM users WHERE id = $id"); // dump complet !

// --- ✅ CORRIGÉ : requête préparée ---
$stmt = $pdo->prepare('SELECT * FROM users WHERE id = ?');
$stmt->execute([(int) $id]);

// --- ❌ VULNÉRABLE : ORDER BY injectable (ne peut pas être un paramètre lié) ---
$tri = $_GET['tri']; // attaquant : (SELECT password FROM users)
$result = $pdo->query("SELECT * FROM logs ORDER BY $tri");

// --- ✅ CORRIGÉ : whitelist stricte ---
$colonnesAutorisees = ['date', 'niveau', 'message'];
$tri = $_GET['tri'] ?? 'date';
if (!in_array($tri, $colonnesAutorisees, true)) {
    $tri = 'date';
}
$sens = ($_GET['sens'] ?? 'asc') === 'desc' ? 'DESC' : 'ASC';
$stmt = $pdo->query("SELECT * FROM logs ORDER BY $tri $sens"); // $tri/$sens = valeurs sûres (nous)

// --- ❌ VULNÉRABLE : nom de table dynamique ---
$table = $_GET['table'];
$pdo->query("SELECT * FROM $table");

// --- ✅ CORRIGÉ : whitelist de tables ---
$tables = ['interventions' => true, 'equipements' => true];
$table = $_GET['table'] ?? 'interventions';
if (!isset($tables[$table])) { $table = 'interventions'; }
```

**Règle universelle :** tout ce qui est **structure SQL** (table, colonne, ASC/DESC) → **whitelist**. Tout ce qui est **donnée** → **paramètre lié**. Aucune exception.

---

## 49. Sécurité : XSS — échapper en sortie

```php
<?php declare(strict_types=1);

