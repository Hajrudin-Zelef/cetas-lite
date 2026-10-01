---
id: collect-261001-rattrapage/rattrapage/php-guide-5
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [900, 1211]
sha256: c1cc16eebead4d48750bdfa8f610108931c83ed04523a2777de7061efb74fdd9
---

# PHP 8 — Le guide complet du sysadmin qui héberge

// --- readonly (PHP 8.1) : propriété assignée UNE fois (au constructeur), puis immuable ---
class ConfigServeur
{
    public function __construct(
        public readonly string $hostname,
        public readonly string $ip,
    ) {}
}

$cfg = new ConfigServeur('srv-web-01', '10.0.0.11');
// $cfg->ip = '10.0.0.99'; // ERREUR : Cannot modify readonly property

// --- Constantes de classe ---
class NiveauAlerte
{
    public const INFO = 'info';
    public const CRITIQUE = 'critique';
}
echo NiveauAlerte::CRITIQUE;

// --- Membres statiques : compteur partagé par toutes les instances ---
class CompteurRequetes
{
    private static int $total = 0;

    public static function incrementer(): void
    {
        self::$total++;
    }

    public static function getTotal(): int
    {
        return self::$total;
    }
}
CompteurRequetes::incrementer();
```

---

## 24. Héritage, classes abstraites, méthodes

```php
<?php declare(strict_types=1);

// --- Classe abstraite : ne s'instancie pas, définit un contrat partiel ---
abstract class Equipement
{
    public function __construct(
        protected string $reference,
        protected string $emplacement,
    ) {}

    abstract public function getType(): string;  // chaque enfant DOIT l'implémenter

    public function etiquette(): string          // méthode concrète partagée
    {
        return "[{$this->getType()}] {$this->reference} @ {$this->emplacement}";
    }

    final public function getReference(): string // final = non surchargeable
    {
        return $this->reference;
    }
}

class Onduleur extends Equipement
{
    public function __construct(string $reference, string $emplacement, private float $puissanceKva)
    {
        parent::__construct($reference, $emplacement); // appel du constructeur parent
    }

    public function getType(): string
    {
        return 'onduleur';
    }

    public function etiquette(): string  // surcharge (override)
    {
        return parent::etiquette() . " ({$this->puissanceKva} kVA)";
    }
}

$ups = new Onduleur('EATON-9PX-20K', 'Local TGBT', 20.0);
echo $ups->etiquette(); // [onduleur] EATON-9PX-20K @ Local TGBT (20 kVA)
// new Equipement(...); // ERREUR : classe abstraite
```

**Règles :** `extends` = **une seule** classe parente. Préfère la **composition** (un objet qui *possède* un autre) à l'héritage profond. Un héritage de plus de 2 niveaux = code smell.


---

## 25. Interfaces — les contrats

```php
<?php declare(strict_types=1);

// --- Une interface = un contrat : "tout objet qui fait X expose ces méthodes" ---
interface Alertable
{
    public function envoyerAlerte(string $message, string $niveau): bool;
}

interface Supervisable
{
    public function getMetriques(): array;
}

// --- Une classe peut implémenter PLUSIEURS interfaces ---
class OnduleurConnecte implements Alertable, Supervisable
{
    public function envoyerAlerte(string $message, string $niveau): bool
    {
        // envoi mail/SNMP...
        return true;
    }

    public function getMetriques(): array
    {
        return ['charge' => 65.0, 'batterie' => 98.0];
    }
}

// --- Typer par l'interface, pas par la classe : découplage ---
function diffuserAlerte(Alertable $equipement, string $msg): void
{
    $equipement->envoyerAlerte($msg, 'critique');
    // Accepte OnduleurConnecte ET tout autre objet Alertable (clim, groupe électrogène...)
}
```

**Quand utiliser quoi :**

| Besoin | Outil |
|---|---|
| Contrat de comportement partagé par des classes sans lien | `interface` |
| Code + état partagés dans une hiérarchie | classe abstraite / héritage |
| Réutiliser du code entre classes sans lien | `trait` (section 26) |

---

## 26. Traits — réutilisation horizontale de code

```php
<?php declare(strict_types=1);

// --- Un trait = un bloc de méthodes réutilisable dans n'importe quelle classe ---
trait Journalisable
{
    private array $journal = [];

    public function journaliser(string $message): void
    {
        $this->journal[] = date('Y-m-d H:i:s') . ' ' . $message;
    }

    public function getJournal(): array
    {
        return $this->journal;
    }
}

trait Horodatable
{
    public function maintenant(): string
    {
        return date('Y-m-d H:i:s');
    }
}

class Intervention
{
    use Journalisable;   // importe les méthodes du trait
    use Horodatable;

    public function cloturer(): void
    {
        $this->journaliser('Intervention clôturée');
    }
}

$inter = new Intervention();
$inter->cloturer();
print_r($inter->getJournal());

// --- Conflit entre traits : résolution explicite ---
trait A { public function qui(): string { return 'A'; } }
trait B { public function qui(): string { return 'B'; } }
class C
{
    use A, B {
        A::qui insteadof B;   // A gagne
        B::qui as quiDeB;      // alias pour garder B accessible
    }
}
```

---

## 27. Enums (PHP 8.1) — fini les constantes magiques

```php
<?php declare(strict_types=1);

// --- Enum pure ---
enum StatutIntervention
{
    case Planifiee;
    case EnCours;
    case Cloturee;
    case Annulee;
}

function changerStatut(StatutIntervention $statut): void
{
    // $statut ne peut être QUE l'une des 4 valeurs — impossible de passer "nimporte quoi"
    echo match ($statut) {
        StatutIntervention::Planifiee => 'à venir',
        StatutIntervention::EnCours   => 'en cours',
        StatutIntervention::Cloturee  => 'terminée',
        StatutIntervention::Annulee   => 'annulée',
    };
}
changerStatut(StatutIntervention::EnCours);

// --- Enum adossée (backed) : valeur scalaire associée — idéale pour la BDD ---
enum NiveauAlarme: string
{
    case Info     = 'info';
    case Warning  = 'warning';
    case Critique = 'critique';
}

$niveau = NiveauAlarme::from('critique');      // lève ValueError si valeur inconnue
$niveau = NiveauAlarme::tryFrom($_POST['n'] ?? ''); // null si inconnue ✅ (validation d'entrée !)
echo $niveau->value;  // "critique"
echo $niveau->name;   // "Critique"

// --- Méthodes dans les enums ---
enum Priorite: int
{
    case Basse = 1;
    case Normale = 2;
    case Haute = 3;

    public function delaiInterventionHeures(): int
    {
        return match ($this) {
            self::Basse   => 72,
            self::Normale  => 24,
            self::Haute   => 4,
        };
    }
}
```

🔒 **Usage sysadmin :** les enums adossées sont parfaites pour valider les entrées utilisateur à valeurs fermées (statuts, rôles, niveaux) — voir section 38.

---

## 28. Namespaces — organiser le code

```php
<?php declare(strict_types=1);

// --- Déclaration : première instruction après declare ---
namespace App\Supervision;

class Sonde
{
    // ...
}

// --- Utilisation depuis un autre fichier ---
namespace App\Web;

use App\Supervision\Sonde;              // import
use App\Supervision\Alerte as AlerteSup; // alias en cas de collision
use function App\outils\formaterDuree;   // import de fonction
use const App\outils\SEUIL;              // import de constante

$sonde = new Sonde();                    // résolu via le use
$autre = new \App\Autre\Truc();          // \ = nom pleinement qualifié (depuis la racine)

// --- Group use ---
use App\Supervision\{Sonde, Alerte, Rapport};
```

**Règles :** un namespace par fichier, `PascalCase` par segment, reflète l'arborescence des dossiers (`App\Supervision\Sonde` → `src/Supervision/Sonde.php`) — c'est la convention **PSR-4** (section 29).

---

## 29. Autoloading PSR-4 — plus jamais de require

```php
<?php
// ❌ AVANT (à bannir) :
// require 'src/Supervision/Sonde.php';
// require 'src/Supervision/Alerte.php';
// ... 50 lignes de require

// ✅ APRÈS : un seul autoloader (généré par Composer, section 30) :
require __DIR__ . '/vendor/autoload.php';

use App\Supervision\Sonde;
$sonde = new Sonde(); // la classe est chargée AUTOMATIQUEMENT à la première utilisation
```

Correspondance PSR-4 (`composer.json`) :

