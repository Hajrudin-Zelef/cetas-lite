---
id: collect-261001-rattrapage/rattrapage/php-guide-4
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-10-01"]
keywords: ["datacenter", "decode"]
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [680, 899]
sha256: df93ba6242969f98bd6d3ba0012ebe05319e7c7ca14b17eada4b0ca2e2ad8bf0
---

# PHP 8 — Le guide complet du sysadmin qui héberge

// --- Nettoyage ---
trim($s);                    // supprime espaces début/fin
ltrim($s); rtrim($s, "/");   // un côté ; rtrim($url, "/") pour normaliser les URL

// --- Recherche (strpos retourne la POSITION ou false — comparer avec ===) ---
$pos = strpos($s, "intervention");
if ($pos !== false) { /* trouvé */ }           // ⚠️ !== false, pas juste if ($pos) (position 0 = falsy !)
str_contains($s, "42");                        // PHP 8 : true/false — PRÉFÉRÉ à strpos pour un simple test
str_starts_with($s, "  Rapport");              // PHP 8
str_ends_with($fichier, ".log");               // PHP 8

// --- Extraction / remplacement ---
substr($s, 2, 8);                              // sous-chaîne
str_replace("42", "43", $s);                   // remplace (sensible à la casse)
str_ireplace("rapport", "compte-rendu", $s);   // insensible à la casse
explode(",", "a,b,c");                         // string → tableau
implode(",", ["a", "b", "c"]);                 // tableau → string ("a,b,c")

// --- Casse ---
strtolower($s); strtoupper($s);
ucfirst("bonjour");        // "Bonjour"
mb_strtolower($s, 'UTF-8');// version multibyte (accents)

// --- Padding / répétition ---
str_pad("42", 5, "0", STR_PAD_LEFT);  // "00042" (numéros d'intervention)
str_repeat("-", 40);                   // ligne de séparation

// --- HTML (sécurité, voir section 49) ---
htmlspecialchars($saisie, ENT_QUOTES, 'UTF-8'); // échappe < > & " ' — ANTI-XSS
strip_tags($html);                              // supprime les balises (brutal, à éviter seul)

// --- Comparaison ---
strcmp($a, $b);            // binaire, sensible à la casse (<0, 0, >0)
strcasecmp($a, $b);        // insensible à la casse
$a === $b;                 // ✅ le plus simple pour l'égalité exacte
```

---

## 19. Déclarer des fonctions — typage complet

```php
<?php declare(strict_types=1);

// --- Fonction typée : paramètres + retour ---
function calculerAutonomie(float $capaciteAh, float $tensionV, float $chargeW): float
{
    if ($chargeW <= 0) {
        throw new InvalidArgumentException("La charge doit être positive");
    }
    return ($capaciteAh * $tensionV) / $chargeW; // heures
}

// --- Paramètres optionnels (valeur par défaut) ---
function formaterDuree(int $minutes, string $format = '%dh %02dmin'): string
{
    return sprintf($format, intdiv($minutes, 60), $minutes % 60);
}

// --- Arguments nommés (PHP 8.0) : clarté à l'appel ---
formaterDuree(minutes: 135);                    // "2h 15min"
formaterDuree(minutes: 135, format: '%d heures'); // ordre libre

// --- Retour nullable / void ---
function trouverServeur(string $ip): ?array  // ?array = array|null
{
    // ...
    return null; // ou un tableau
}

function logger(string $message): void  // ne retourne rien (même pas null implicite exploitable)
{
    file_put_contents('/var/log/appli.log', date('c') . " $message\n", FILE_APPEND);
}
```

**Règles d'or :** une fonction = **une** responsabilité, nom **verbe** explicite (`validerEmail`, pas `check`), typer **tous** les paramètres et le retour, **jamais** de `echo` dans une fonction métier (elle retourne, l'appelant affiche).

---

## 20. Closures et arrow functions

```php
<?php declare(strict_types=1);

// --- Closure classique : capture avec use ---
$seuil = 80;
$alerte = function (float $charge) use ($seuil): bool {
    return $charge >= $seuil;
};
var_dump($alerte(85.0)); // true

// --- Arrow function (PHP 7.4) : capture AUTOMATIQUE par valeur, une seule expression ---
$niveaux = ['info', 'erreur', 'info'];
$estErreur = fn(string $n): bool => $n === 'erreur';
$erreurs = array_filter($niveaux, $estErreur);

// Enchaîné avec array_map :
$hostnames = array_map(fn(array $s): string => $s['hostname'], $serveurs);

// --- Callable stocké et invoqué ---
$traitements = [
    'majuscules' => fn(string $s): string => mb_strtoupper($s, 'UTF-8'),
    'trim'       => fn(string $s): string => trim($s),
];
echo $traitements['trim']("  abc  "); // "abc"

// --- Fonction variadique : nombre d'arguments variable ---
function somme(float ...$nombres): float
{
    return array_sum($nombres);
}
echo somme(1.5, 2.5, 3.0); // 7.0
```

---

## 21. Fonctions utilitaires du quotidien sysadmin

```php
<?php declare(strict_types=1);

// --- Fichiers ---
file_get_contents('/etc/hostname');                    // lit tout un fichier
file_put_contents('/tmp/test.txt', "ligne\n", FILE_APPEND | LOCK_EX); // écrit (LOCK_EX = verrou)
file_exists('/var/www/app/.env');                      // existe ?
is_readable($chemin); is_writable($chemin); is_dir($chemin);
unlink($fichier);                                      // supprime
rename($ancien, $nouveau);                             // déplace/renomme
filesize($fichier); filemtime($fichier);               // taille, date modif
$lignes = file('/var/log/syslog', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES); // fichier → tableau de lignes
glob('/var/log/*.log');                                // liste les fichiers par motif

// --- JSON (API, configs) ---
$json = json_encode($data, JSON_PRETTY_PRINT | JSON_UNESCAPED_UNICODE | JSON_THROW_ON_ERROR);
$data = json_decode($json, true, 512, JSON_THROW_ON_ERROR); // true = tableau associatif
// ⚠️ JSON_THROW_ON_ERROR (PHP 7.3+) : lève une exception au lieu de retourner null silencieusement

// --- Divers ---
sleep(2);                    // pause en secondes (scripts CLI)
uniqid('job_', true);         // identifiant unique
random_int(1000, 9999);       // aléatoire CRYPTOGRAPHIQUEMENT sûr (pas rand() !)
var_dump($x);                // debug : type + valeur
print_r($x, true);            // debug : retourne au lieu d'afficher (pour logger)
gettype($x);                 // "integer", "string"...
```

🔒 `random_int()` pour tout ce qui touche à la sécurité (tokens, OTP). `rand()`/`mt_rand()` sont prédictibles : **interdits** pour les tokens (voir CSRF section 50).

---

## 22. POO : les classes — bases

```php
<?php declare(strict_types=1);

// --- Classe simple avec propriétés typées ---
class Onduleur
{
    public string $marque;
    public float $puissanceKva;
    private float $chargePct = 0.0;          // privée : accès via méthodes
    protected string $numeroSerie;           // protégée : classe + enfants

    // --- Constructeur ---
    public function __construct(string $marque, float $puissanceKva, string $numeroSerie)
    {
        $this->marque = $marque;
        $this->puissanceKva = $puissanceKva;
        $this->numeroSerie = $numeroSerie;
    }

    // --- Méthode ---
    public function definirCharge(float $pct): void
    {
        if ($pct < 0 || $pct > 100) {
            throw new InvalidArgumentException("Charge entre 0 et 100");
        }
        $this->chargePct = $pct;
    }

    public function puissanceRestanteKw(): float
    {
        return $this->puissanceKva * 0.9 * (1 - $this->chargePct / 100);
    }

    // --- Getter (lecture contrôlée d'une propriété privée) ---
    public function getCharge(): float
    {
        return $this->chargePct;
    }
}

$ups = new Onduleur('Eaton', 20.0, 'SN-8842');
$ups->definirCharge(65.0);
echo $ups->puissanceRestanteKw(); // 6.3
// echo $ups->chargePct; // ERREUR : propriété privée
```

---

## 23. Constructeurs modernes : promotion et readonly

```php
<?php declare(strict_types=1);

// --- Promotion des propriétés (PHP 8.0) : fini le $this->x = $x répétitif ---
class Intervention
{
    public function __construct(
        public int $id,
        public string $site,
        public DateTimeImmutable $date,
        private string $statut = 'planifiee',  // valeur par défaut possible
    ) {}
}

$inter = new Intervention(42, 'Datacenter A', new DateTimeImmutable('2026-10-01'));
echo $inter->site; // "Datacenter A"

