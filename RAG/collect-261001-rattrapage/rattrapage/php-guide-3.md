---
id: collect-261001-rattrapage/rattrapage/php-guide-3
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [441, 679]
sha256: 5da15e5b3b96082fe842de970a171f84190adf4b5c2018c40946dfef091d42a1
---

# PHP 8 — Le guide complet du sysadmin qui héberge

| Priorité à connaître | Exemple correct |
|---|---|
| `&&` prioritaire sur `and` | `$ok = $a && $b;` (pas `$ok = $a and $b;` qui assigne d'abord !) |
| `.` et `+` : même niveau, associatif gauche | `'total: ' . $a + $b` → **piège** : parenthèse ! `('total: ' . ($a + $b))` |
| `??` basse priorité à droite | `$x = $a ?? $b;` OK |


---

## 13. Structures de contrôle : if, switch, match

```php
<?php declare(strict_types=1);

// --- if / elseif / else : accolades OBLIGATOIRES (même pour une ligne) ---
$charge = 82.5;
if ($charge >= 90) {
    $niveau = 'critique';
} elseif ($charge >= 70) {
    $niveau = 'alerte';
} else {
    $niveau = 'ok';
}

// --- switch : utile pour de nombreuses valeurs discrètes ---
$code = 'C6000';
switch ($code) {
    case 'C6000':
    case 'C6020':
        $action = 'vérifier le four';
        break;
    case 'C0640':
        $action = 'vérifier le disque dur';
        break;
    default:
        $action = 'code inconnu';
}
// ⚠️ Ne pas oublier break : sinon ça "tombe" dans le case suivant (fall-through).

// --- match (PHP 8.0) : expression, comparaison STRICTE (===), pas de break ---
$action = match ($code) {
    'C6000', 'C6020' => 'vérifier le four',
    'C0640'          => 'vérifier le disque dur',
    default          => 'code inconnu',
};
// match retourne une valeur et lève UnhandledMatchError si aucun bras ne correspond
// (sans default). Préfère match à switch dans le code neuf.
```

---

## 14. Boucles : for, foreach, while

```php
<?php declare(strict_types=1);

// --- foreach : LA boucle à utiliser sur les tableaux ---
$onduleurs = ['Eaton 9PX', 'APC SRT', 'Vertiv Liebert'];
foreach ($onduleurs as $index => $nom) {
    echo "$index : $nom\n";
}

// Modification par référence (avec & — attention à unset après !)
$puissances = [10, 20, 40];
foreach ($puissances as &$p) {
    $p *= 1000;   // passe en VA
}
unset($p);        // ⚠️ OBLIGATOIRE : sinon $p reste une référence piégeuse

// --- for : quand l'index compte ---
for ($i = 0; $i < 5; $i++) {
    echo "itération $i\n";
}

// --- while / do...while ---
$tentatives = 0;
while ($tentatives < 3) {
    $tentatives++;
    // ...
}

// --- break / continue ---
foreach ($fichiers as $f) {
    if ($f === '.') { continue; }          // passe au suivant
    if ($f === 'stop.txt') { break; }      // sort de la boucle
    if (str_ends_with($f, '.log')) { break 2; } // break 2 = sort de 2 niveaux (boucles imbriquées)
}
```

---

## 15. Les tableaux — création et manipulation

```php
<?php declare(strict_types=1);

// --- Création (syntaxe courte [] — toujours) ---
$vide    = [];
$notes   = [12, 15, 9];
$serveur = [
    'hostname' => 'srv-web-01',
    'ip'       => '10.0.0.11',
    'role'     => 'web',
];

// --- Accès ---
echo $serveur['hostname'];          // srv-web-01
echo $serveur['absent'] ?? 'n/a';   // ?? évite le warning "undefined array key"

// --- Ajout ---
$notes[] = 18;                      // ajoute à la fin
$serveur['os'] = 'Debian 12';       // nouvelle clé

// --- Suppression ---
unset($notes[1]);                   // supprime l'élément (les clés numériques ne sont PAS réindexées)
$notes = array_values($notes);      // réindexe 0..n

// --- Compter / tester ---
count($notes);                      // nombre d'éléments
isset($serveur['ip']);               // true si existe ET non null
array_key_exists('ip', $serveur);    // true même si valeur null
empty($vide);                       // true si vide / falsy
in_array('web', $serveur, true);    // ⚠️ true en 3e arg = comparaison stricte

// --- Déstructuration (list) ---
[$premier, $deuxieme] = $notes;
['hostname' => $h, 'ip' => $ip] = $serveur;  // déstructuration par clés (PHP 7.1+)

// --- Spread operator ---
$tous = [...$notes, ...[20, 21]];    // fusion (PHP 7.4+)
$config = [...$defauts, ...$specifiques]; // les clés string de droite écrasent celles de gauche
```

---

## 16. Fonctions sur tableaux — la boîte à outils

Le sysadmin manipule des listes (fichiers, lignes de log, équipements) : ces fonctions sont quotidiennes.

```php
<?php declare(strict_types=1);

$logs = [
    ['niveau' => 'info',    'msg' => 'démarrage'],
    ['niveau' => 'erreur',  'msg' => 'disque plein'],
    ['niveau' => 'info',    'msg' => 'backup ok'],
    ['niveau' => 'erreur',  'msg' => 'timeout bdd'],
];

// --- array_map : transforme chaque élément ---
$niveaux = array_map(fn(array $l): string => $l['niveau'], $logs);
// ['info', 'erreur', 'info', 'erreur']

// --- array_filter : garde ceux qui satisfont le test ---
$erreurs = array_filter($logs, fn(array $l): bool => $l['niveau'] === 'erreur');
$erreurs = array_values($erreurs); // réindexer après un filter

// --- array_reduce : réduit à une valeur (comptage, somme...) ---
$nbErreurs = array_reduce($logs, fn(int $n, array $l): int => $n + ($l['niveau'] === 'erreur' ? 1 : 0), 0);

// --- array_column : extrait une colonne ---
$messages = array_column($logs, 'msg');
// --- array_column avec clé d'index ---
$parNiveau = array_column($logs, 'msg', 'niveau');

// --- Tri ---
sort($notes);                 // valeurs croissantes (réindexe)
rsort($notes);                // décroissant
asort($notes);                // croissant en conservant les clés
ksort($serveur);              // tri par clé
usort($logs, fn($a, $b) => $a['msg'] <=> $b['msg']); // tri personnalisé (spaceship)

// --- Recherche ---
$cles = array_keys($serveur);            // ['hostname', 'ip', 'role']
$vals = array_values($serveur);
array_search('10.0.0.11', $serveur, true); // retourne la clé ou false (strict !)

// --- Ensembles ---
$fusion   = array_merge($a, $b);   // concatène (clés string de $b écrasent)
$communs  = array_intersect($a, $b);
$diff     = array_diff($a, $b);    // dans $a mais pas dans $b
$uniques  = array_unique($liste);

// --- Découpage ---
$premiers = array_slice($logs, 0, 10);            // 10 premiers
$pages    = array_chunk($logs, 50);               // découpe en paquets de 50 (pagination !)
[$ok, $ko] = [array_filter(...), array_filter(...)]; // séparer en deux groupes

// --- Clés/valeurs ---
$inverse  = array_flip(['a' => 1, 'b' => 2]); // [1 => 'a', 2 => 'b']
$rempli   = array_fill_keys(['lundi', 'mardi'], 0); // initialise un compteur par jour
$combine  = array_combine($cles, $vals);       // construit un tableau clé => valeur
```

📋 **Mémo express :** transformer → `array_map`, filtrer → `array_filter`, agréger → `array_reduce`, colonne → `array_column`, tri custom → `usort` + `<=>`, paginer → `array_chunk`.

---

## 17. Chaînes de caractères

```php
<?php declare(strict_types=1);

// --- Guillemets ---
$nom = 'Zelef';
$a = "Bonjour $nom, il est " . date('H:i');  // doubles quotes : interpolation + séquences \n \t
$b = 'Bonjour $nom';                          // simples quotes : LITTÉRAL (affiche $nom)
$c = 'C:\Windows\System32';                   // simples quotes : \ non interprété (sauf \' et \\)

// --- Heredoc / Nowdoc (textes multilignes : templates, SQL, mails) ---
$mail = <<<TEXTE
Bonjour $nom,

L'intervention sur l'onduleur est planifiée.
Cordialement,
TEXTE;   // heredoc : interpolation active, le marqueur final doit être en début de ligne (ou indenté, PHP 7.3+)

$sql = <<<'SQL'
SELECT * FROM interventions WHERE statut = 'planifiee'
SQL;     // nowdoc (quotes autour du marqueur) : AUCUNE interpolation — idéal pour SQL/regex

// --- Longueur : TOUJOURS mb_strlen en UTF-8 ---
$ville = "Hébergeur";
strlen($ville);      // 10 octets ❌
mb_strlen($ville);   // 9 caractères ✅

// --- Concaténer proprement ---
$ligne = sprintf("Serveur %s : charge %d%%", $hostname, $charge); // formaté
```

---

## 18. Fonctions sur chaînes — référence

```php
<?php declare(strict_types=1);

$s = "  Rapport d'intervention N°42  ";

