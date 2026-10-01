---
id: collect-261001-rattrapage/rattrapage/php-guide-2
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [207, 440]
sha256: f3fd5c2458323d05e348b852f78070fbe97bc3b9529914a6245c284ecc8018c4
---

# PHP 8 — Le guide complet du sysadmin qui héberge

```php
<?php
// Tout fichier PHP commence par <?php (jamais la balise courte <? seule en prod)

declare(strict_types=1); // voir section 9 — à mettre en tête des fichiers applicatifs

echo "Bonjour Zelef\n";          // affiche + retour ligne
print "Bonjour\n";               // équivalent (retourne 1)

$nom = "Monde";                  // variable : $ + nom
echo "Bonjour $nom\n";           // interpolation dans les doubles quotes
echo 'Bonjour ' . $nom . "\n";   // concaténation avec le point .

// Commentaires : // ligne, # ligne, /* bloc */ , /** docblock */

/*
 * Commentaire sur
 * plusieurs lignes
 */
```

**Règles de nommage :**

| Élément | Convention |
|---|---|
| Variables, fonctions, méthodes | `camelCase` : `$dateDebut`, `calculerTotal()` |
| Classes, interfaces, traits, enums | `PascalCase` : `FactureClient` |
| Constantes | `MAJUSCULES_SNAKE` : `TAUX_TVA` |

Ferme la balise `?>` ? **Non** : dans un fichier 100 % PHP, on **omet** la balise fermante pour éviter tout caractère parasite (espace, BOM) qui casserait les en-têtes HTTP (`headers already sent`).

---

## 8. Les types en PHP 8

PHP est à typage **dynamique** mais supporte un typage **strict et complet** depuis PHP 7/8. En 2026, on type tout.

```php
<?php declare(strict_types=1);

// --- Scalaires ---
$entier   = 42;            // int
$flottant = 19.99;         // float
$texte    = "câble";       // string
$vrai     = true;          // bool

// --- Composés ---
$tableau  = [1, 2, 3];     // array
$objet    = new DateTime();// object

// --- Spéciaux ---
$nulle    = null;          // null
$callable = 'strlen';      // callable (nom de fonction, closure, [obj, méthode])

// --- Types du système de typage (déclarations) ---
function demo(
    int $a,                 // scalaire
    ?string $b,             // nullable (? = null autorisé)
    array $c,               // tableau
    int|float $d,           // union de types (PHP 8.0)
    mixed $e,               // tout type (PHP 8.0)
    callable $f,            // appelable
): int|string {             // type de retour (union)
    return $a;
}

// Types spéciaux de retour :
function neRevientJamais(): never { throw new Exception("stop"); } // PHP 8.1 : ne retourne jamais
```

**Conversions (cast) :**

```php
$x = (int) "42abc";    // 42 (PHP tronque, ne pas abuser)
$y = (float) "19.99";  // 19.99
$z = (bool) "";        // false — "" , "0", 0, 0.0, [], null sont falsy
$s = (string) 42;      // "42"
$a = (array) "x";      // ["x"]
```

⚠️ **Piège classique :** `==` compare avec conversion de type (`"42" == 42` → `true`), `===` compare type + valeur (`"42" === 42` → `false`). **En PHP 8, utilise `===` par défaut.**

---

## 9. `declare(strict_types=1)` — le mode strict

Par défaut, PHP convertit silencieusement les types scalaires à l'appel des fonctions typées (mode **coercitif**) :

```php
<?php
// SANS strict_types :
function addition(int $a, int $b): int { return $a + $b; }
echo addition("2", "3");   // 5 — "2" converti en 2 sans prévenir
```

Avec `declare(strict_types=1);` en **première ligne** du fichier appelant, toute incohérence lève une `TypeError` :

```php
<?php declare(strict_types=1);

function addition(int $a, int $b): int { return $a + $b; }

echo addition("2", "3");
// TypeError: addition(): Argument #1 ($a) must be of type int, string given
```

**Règles :**

- Le `declare` s'applique au **fichier où se trouve l'appel**, pas où la fonction est définie.
- Il doit être la **toute première instruction** du fichier (avant tout code, après `<?php`).
- À mettre systématiquement dans ton code applicatif. Les libs externes gèrent leur propre fichier.

🔒 Le mode strict élimine toute une classe de bugs silencieux (dates, montants, IDs passés en string). **Non négociable dans du code neuf.**

---

## 10. Variables, constantes, portée

```php
<?php declare(strict_types=1);

// --- Variables ---
$compteur = 0;
$compteur += 5;                 // opérateurs combinés : += -= *= /= .= %= **=
$nomComplet = $prenom . ' ' . $nom;

// --- Constantes : define() ou const (préféré dans les classes) ---
define('SEUIL_ALERTE', 80);     // globale, partout
const DOSSIER_LOGS = '/var/log/appli'; // au niveau fichier ou classe

echo SEUIL_ALERTE;

// --- Constantes magiques ---
echo __FILE__;   // chemin du fichier courant
echo __DIR__;    // dossier du fichier courant (préféré à dirname(__FILE__))
echo __LINE__;   // numéro de ligne
echo __FUNCTION__, __CLASS__, __METHOD__;

// --- Portée : les fonctions ne voient PAS les variables globales ---
$taxe = 0.20;
function prixTTC(float $ht): float {
    // echo $taxe; // ERREUR : undefined variable
    global $taxe;              // ❌ à éviter
    return $ht * (1 + $taxe);
}
// ✅ Bien : passer en paramètre
function prixTtcPropre(float $ht, float $taux): float {
    return $ht * (1 + $taux);
}

// --- Variables statiques : conservent leur valeur entre appels ---
function compteurAppels(): int {
    static $n = 0;   // initialisée une seule fois
    return ++$n;
}
echo compteurAppels(); // 1
echo compteurAppels(); // 2

// --- Variables variables (à connaître, à éviter) ---
$cle = 'nom';
$$cle = 'Zelef';   // crée $nom = 'Zelef' — obscur, banni des revues de code
```

---

## 11. Les superglobales

Accessibles partout, sans `global`. **Toutes les données venant de l'utilisateur sont hostiles** : `$_GET`, `$_POST`, `$_COOKIE`, `$_FILES`, `$_REQUEST` = à valider systématiquement (sections 38, 48-50).

```php
<?php
$_GET      // paramètres d'URL : page.php?id=42  → $_GET['id']
$_POST     // corps des formulaires POST
$_COOKIE   // cookies envoyés par le navigateur
$_FILES    // fichiers uploadés (structure spéciale, voir section 43)
$_SERVER   // infos serveur/requête : $_SERVER['REQUEST_METHOD'], ['REMOTE_ADDR'], ['HTTP_HOST']...
$_SESSION  // session (après session_start())
$_ENV      // variables d'environnement
$GLOBALS   // toutes les variables globales

// Exemples sûrs :
$methode = $_SERVER['REQUEST_METHOD'] ?? 'GET';   // ?? = null coalescent
$ip      = $_SERVER['REMOTE_ADDR'] ?? 'inconnue';
$hote    = $_SERVER['HTTP_HOST'] ?? '';

// ⚠️ $_REQUEST fusionne GET+POST+COOKIE : ordre configurable, source ambiguë.
// Préfère $_GET / $_POST explicites. Ne JAMAIS utiliser $_REQUEST en prod.
```

**Opérateur null coalescent `??` et assignation `??=` :**

```php
$page  = $_GET['page'] ?? 1;          // 1 si absent ou null
$tri   = $_GET['tri'] ?? $_POST['tri'] ?? 'date';  // chaînage
$config['debug'] ??= false;           // assigne seulement si null/absent
```

---

## 12. Opérateurs — l'essentiel et les pièges

```php
<?php declare(strict_types=1);

// --- Arithmétiques : + - * / % ** (puissance) ---
$reste = 10 % 3;      // 1
$carre = 4 ** 2;       // 16

// --- Comparaison : === !== == != <=> ??
$a = "42";
var_dump($a == 42);    // true  (conversion)
var_dump($a === 42);   // false (type différent) ✅ préféré
var_dump(3 <=> 5);     // -1 (spaceship : -1, 0 ou 1) — idéal pour usort

// --- Logiques : && || ! (et : and or xor, priorité différente — préfère && ||) ---
if ($connecte && $admin) { /* ... */ }

// --- Ternaire et Elvis ---
$role = $estAdmin ? 'admin' : 'user';
$nom  = $saisie ?: 'anonyme';   // ?: = $saisie si truthy, sinon 'anonyme'

// --- Chaînes ---
echo "Il est " . date('H:i');   // concaténation avec .

// --- Incrémentation ---
$i = 0; $i++; ++$i; $i--; --$i;

// --- Opérateur de contrôle d'erreur @ : À BANNIR ---
// $x = @file_get_contents($url); // ❌ masque les erreurs, ralentit, rend le debug impossible
// ✅ :
$contenu = file_get_contents($url);
if ($contenu === false) { /* gérer l'erreur */ }
```

