---
id: collect-261001-rattrapage/rattrapage/php-guide-7
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["1970-01-01", "2019-03-15", "2026-09-12", "2026-10-01", "2026-10-05", "2026-10-07"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [1442, 1646]
sha256: 0e96d66c02743317f5efb7481361f017696e1106e84d2183c70276778cfe436a
---

# PHP 8 — Le guide complet du sysadmin qui héberge

// --- Création : TOUJOURS avec fuseau explicite ---
$maintenant = new DateTimeImmutable('now', new DateTimeZone('Europe/Paris'));
$intervention = new DateTimeImmutable('2026-10-05 14:30:00', new DateTimeZone('Europe/Paris'));

// --- Immuable = chaque modification retourne un NOUVEL objet ---
$rappel = $intervention->modify('-2 days');  // $intervention INCHANGÉE ✅
// (avec DateTime mutable, modify() change l'objet lui-même → bugs vicieux)

// --- Calculs ---
$duree = new DateInterval('PT2H30M');        // Période : 2h30 (P1D = 1 jour, PT15M = 15 min)
$fin = $intervention->add($duree);
$diff = $intervention->diff($fin);           // DateInterval
echo $diff->format('%h heures %i minutes');

// --- Comparaison directe (objets comparables) ---
if ($rappel < $maintenant) {
    echo "le rappel est passé";
}

// --- Formats ---
echo $intervention->format('Y-m-d H:i:s');  // 2026-10-05 14:30:00 (BDD)
echo $intervention->format('d/m/Y à H\hi'); // 05/10/2026 à 14h30 (affichage FR)
echo $intervention->format('c');            // ISO 8601 : 2026-10-05T14:30:00+02:00

// --- Depuis un format précis (formulaires FR !) ---
$date = DateTimeImmutable::createFromFormat('d/m/Y H:i', '05/10/2026 14:30');
if ($date === false) {
    // format invalide → erreur de validation
}

// --- Timestamp ---
$ts = $intervention->getTimestamp();        // secondes depuis 1970-01-01 UTC
$depuis = new DateTimeImmutable('@' . $ts);  // @ = timestamp (toujours UTC)
```

---

## 36. Dates — pièges et bonnes pratiques

```php
<?php declare(strict_types=1);

// --- 1. Stocker en UTC, afficher en locale ---
// BDD : colonne DATETIME/TIMESTAMP en UTC. Conversion à l'affichage :
$utc = new DateTimeImmutable('2026-10-05 12:30:00', new DateTimeZone('UTC'));
$paris = $utc->setTimezone(new DateTimeZone('Europe/Paris'));
echo $paris->format('d/m/Y H:i'); // 05/10/2026 14:30 (heure d'été)

// --- 2. "now" sans fuseau = fuseau de date.timezone (php.ini) ---
// ⚠️ Si date.timezone n'est pas défini : warning + supposition système. Toujours le définir.

// --- 3. Comparer des dates : objets, pas des strings ---
// "05/10/2026" < "12/09/2026" en string → FAUX résultat. En objets DateTime : correct.

// --- 4. Générer une plage de dates (planning) ---
$debut = new DateTimeImmutable('2026-10-01');
$fin   = new DateTimeImmutable('2026-10-07');
$periode = new DatePeriod($debut, new DateInterval('P1D'), $fin->modify('+1 day'));
foreach ($periode as $jour) {
    echo $jour->format('Y-m-d') . "\n";
}

// --- 5. Âge / ancienneté d'un équipement ---
$installation = new DateTimeImmutable('2019-03-15');
$age = $installation->diff(new DateTimeImmutable('now'));
echo "En service depuis {$age->y} ans et {$age->m} mois";

// --- 6. Tester "est-ce aujourd'hui ?" ---
$aujourdhui = (new DateTimeImmutable('now'))->format('Y-m-d') === $intervention->format('Y-m-d');
```

📋 **Checklist dates :** `DateTimeImmutable` partout (jamais `DateTime` mutable), fuseau explicite, stockage UTC, `date.timezone` défini dans php.ini, `createFromFormat` + test `=== false` pour les saisies.


---

## 37. Formulaires HTML — GET vs POST

```php
<?php // form.php — affiche ET traite (PRG pattern, voir plus bas) ?>
<!DOCTYPE html>
<html lang="fr">
<head><meta charset="UTF-8"><title>Demande d'intervention</title></head>
<body>
<form method="post" action="form.php">
    <!-- Toujours : label + name + type adapté + required côté HTML (1re barrière, pas la seule) -->
    <label for="site">Site :</label>
    <input type="text" id="site" name="site" required maxlength="100">

    <label for="email">E-mail :</label>
    <input type="email" id="email" name="email" required>

    <label for="priorite">Priorité :</label>
    <select id="priorite" name="priorite">
        <option value="basse">Basse</option>
        <option value="normale" selected>Normale</option>
        <option value="haute">Haute</option>
    </select>

    <label for="description">Description :</label>
    <textarea id="description" name="description" rows="5" required></textarea>

    <button type="submit">Envoyer</button>
</form>
</body>
</html>
```

| | GET | POST |
|---|---|---|
| Usage | Recherche, filtres, pagination (actions **sans effet**) | Création/modification/suppression (actions **avec effet**) |
| Données | Dans l'URL (visibles, loggées, limitées ~2 Ko) | Dans le corps (non visibles dans l'URL) |
| Idempotent | Oui (répétable sans risque) | Non → pattern **PRG** |

**Pattern PRG (Post/Redirect/Get) :** après un POST réussi, on **redirige** (302) vers une page GET. Ça évite le double envoi au refresh (F5) :

```php
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    // ... valider, enregistrer ...
    header('Location: confirmation.php?id=' . $id);
    exit; // ⚠️ exit APRÈS header() : sinon le script continue !
}
```

---

## 38. Validation et assainissement — la discipline

**Deux opérations distinctes :**

1. **Valider** = vérifier que la donnée est conforme (refuser si non). *En entrée.*
2. **Assainir/échapper** = neutraliser pour un contexte de sortie (HTML, SQL, shell). *En sortie.*

```php
<?php declare(strict_types=1);

// --- VALIDATION : whitelist, jamais blacklist ---
function validerDemande(array $post): array
{
    $erreurs = [];

    // Texte : longueur + caractères autorisés
    $site = trim($post['site'] ?? '');
    if ($site === '' || mb_strlen($site) > 100) {
        $erreurs['site'] = 'Site requis (max 100 caractères).';
    } elseif (!preg_match('/^[\p{L}\p{N} .\-_]+$/u', $site)) {
        $erreurs['site'] = 'Caractères non autorisés.';
    }

    // E-mail : filtre dédié
    $email = trim($post['email'] ?? '');
    if (!filter_var($email, FILTER_VALIDATE_EMAIL)) {
        $erreurs['email'] = 'E-mail invalide.';
    }

    // Valeur fermée : enum (impossible de tricher)
    $priorite = Priorite::tryFrom($post['priorite'] ?? '');
    if ($priorite === null) {
        $erreurs['priorite'] = 'Priorité invalide.';
    }

    // Entier dans une plage
    $duree = filter_var($post['duree'] ?? null, FILTER_VALIDATE_INT, [
        'options' => ['min_range' => 15, 'max_range' => 480],
    ]);
    if ($duree === false) {
        $erreurs['duree'] = 'Durée entre 15 et 480 minutes.';
    }

    return $erreurs; // vide = tout est valide
}

// --- ASSAINISSEMENT EN SORTIE (contexte HTML) ---
// On stocke la donnée BRUTE (validée) en BDD, on échappe À L'AFFICHAGE :
echo htmlspecialchars($site, ENT_QUOTES | ENT_HTML5, 'UTF-8');
```

🔒 **Règles :** valider **côté serveur** (le JS/HTML `required` se contourne en 2 clics), whitelist > blacklist, stocker brut + échapper à la sortie (jamais l'inverse — sinon double échappement).

---

## 39. filter_var / filter_input — les filtres natifs

```php
<?php declare(strict_types=1);

// --- Validation ---
filter_var('zelef@exemple.fr', FILTER_VALIDATE_EMAIL);   // string ou false
filter_var('10.0.0.11', FILTER_VALIDATE_IP);              // IP v4/v6
filter_var('10.0.0.11', FILTER_VALIDATE_IP, FILTER_FLAG_IPV4 | FILTER_FLAG_NO_PRIV_RANGE); // IPv4 publique
filter_var('https://exemple.fr/x', FILTER_VALIDATE_URL);
filter_var('42', FILTER_VALIDATE_INT, ['options' => ['min_range' => 1, 'max_range' => 100]]);
filter_var('oui', FILTER_VALIDATE_BOOLEAN, FILTER_NULL_ON_FAILURE); // true/false/null

// --- Assainissement (nettoyage) ---
filter_var('<b>test</b>', FILTER_SANITIZE_SPECIAL_CHARS); // &lt;b&gt;test&lt;/b&gt;
$tel = preg_replace('/[^0-9+ ]/', '', $saisieTel);        // garde chiffres/+ /espace

// --- filter_input : lit DIRECTEMENT la superglobale (préféré à $_POST brut) ---
$email = filter_input(INPUT_POST, 'email', FILTER_VALIDATE_EMAIL); // null si absent, false si invalide
$id    = filter_input(INPUT_GET, 'id', FILTER_VALIDATE_INT);

// --- Tester un booléen de case à cocher ---
$urgent = isset($_POST['urgent']); // checkbox : présente = cochée

