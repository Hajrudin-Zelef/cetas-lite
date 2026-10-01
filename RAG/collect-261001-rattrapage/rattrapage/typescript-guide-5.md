---
id: collect-261001-rattrapage/rattrapage/typescript-guide-5
title: "TypeScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/typescript_guide.md
source_anchor: ""
source_lines: [922, 1190]
sha256: abaadfa940506489ec85bab91c31cfc9af86c869c2aa2e7a392440fac37b2e65
---

# TypeScript — Guide ultra-complet

function traiterEvenement(e: Evenement): void {
  switch (e.type) {
    case "ping":
      console.log(`ping ${e.hote}`); // e: { type: "ping"; hote: string }
      break;
    case "alerte":
      console.log(`[${e.niveau}] ${e.message}`);
      break;
    case "fin":
      console.log(`code ${e.code}`);
      break;
    default:
      // Exhaustiveness : si on ajoute un variant sans case, ça casse ici
      const impossible: never = e;
      throw new Error(`Événement non géré: ${JSON.stringify(impossible)}`);
  }
}
```

Le `const impossible: never = e` est le **contrôle d'exhaustivité** : si un jour vous ajoutez `{ type: "reboot" }` à l'union sans ajouter de `case`, `e` ne sera plus `never` → **erreur de compilation**. Inestimable pour les handlers d'événements et les reducers.

---

## 27. Classes : les bases

```typescript
class Sonde {
  // Propriétés déclarées (strictPropertyInitialization les exige initialisées)
  nom: string;
  private dernierPing: number | null = null;

  constructor(nom: string) {
    this.nom = nom;
  }

  ping(): string {
    this.dernierPing = Date.now();
    return `pong ${this.nom}`;
  }
}

const s = new Sonde("sonde-dc1");
s.ping();
s.dernierPing; // Erreur : private
```

**Rappel important** : les modificateurs `private`/`public` sont **compile-time uniquement** — au runtime, tout reste accessible en JS. Pour une vraie encapsulation runtime : `#champPrive` (champs privés ECMAScript).

```typescript
class Coffre {
  #secret = "s3cret"; // vraiment privé au runtime
}
```

---

## 28. Constructeurs et parameter properties

Les **parameter properties** déclarent + assignent en une ligne :

```typescript
class Equipement3 {
  constructor(
    public readonly nom: string,   // déclare this.nom (public, readonly)
    private ip: string,            // déclare this.ip (private)
    public enLigne = true,         // avec valeur par défaut
  ) {}
  // Pas de corps nécessaire !
}

const e = new Equipement3("srv-01", "10.0.0.1");
e.nom; // "srv-01"
```

**Surcharge de constructeur** (même syntaxe que les overloads, section 20) :

```typescript
class Plage {
  constructor(debut: string, fin: string);
  constructor(cidr: string);
  constructor(a: string, b?: string) {
    // ...
  }
}
```

> `strictPropertyInitialization` : toute propriété non optionnelle doit être assignée dans le constructeur ou à la déclaration. Échappatoire assumée : `propriete!: string` (definite assignment assertion) quand l'init est faite par un framework.

---

## 29. Modificateurs : public, private, protected, readonly, static

```typescript
class Compteur {
  static instances = 0;          // partagé par la classe (pas l'instance)
  readonly id: string;           // assignable une fois (constructeur)
  private valeur = 0;            // classe uniquement (compile-time)
  protected nom: string;         // classe + sous-classes

  constructor(nom: string) {
    this.nom = nom;
    this.id = `c-${Compteur.instances++}`;
  }

  incrementer(): void { this.valeur++; }
  get lecture(): number { return this.valeur; } // getter
}

class CompteurNomme extends Compteur {
  etiquette(): string { return this.nom; } // OK : protected visible ici
}
```

| Modificateur | Visible depuis | Runtime |
|---|---|---|
| `public` (défaut) | partout | normal |
| `private` | la classe seule | encore accessible en JS ! |
| `protected` | classe + héritiers | idem |
| `#prive` | la classe seule | **vraiment privé** |
| `readonly` | lecture partout, écriture au constructeur | normal |
| `static` | via `Classe.membre` | normal |

---

## 30. Héritage et override

```typescript
class EquipementReseau {
  constructor(public nom: string, public ip: string) {}

  decrire(): string {
    return `${this.nom} (${this.ip})`;
  }
}

class Commutateur extends EquipementReseau {
  constructor(nom: string, ip: string, public nbPorts: number) {
    super(nom, ip); // obligatoire avant d'utiliser this
  }

  override decrire(): string {  // "override" exigé si noImplicitOverride
    return `${super.decrire()} - ${this.nbPorts} ports`;
  }
}
```

**Points d'attention** :
- `super(...)` **avant** tout usage de `this` dans le constructeur dérivé.
- `override` explicite : activez `noImplicitOverride: true` dans le tsconfig pour forcer le mot-clé (évite les fautes de frappe qui créent une méthode au lieu d'en surcharger une).
- Préférez la **composition** à l'héritage quand il n'y a pas de vraie relation « est-un ».

---

## 31. Classes abstraites

Une classe abstraite définit un **squelette** : elle ne s'instancie pas, elle s'étend.

```typescript
abstract class Collecteur {
  abstract source(): string;          // à implémenter obligatoirement
  abstract collecter(): Promise<Record<string, number>>;

  // Méthode concrète partagée (template method pattern)
  async cycle(): Promise<void> {
    const debut = Date.now();
    const mesures = await this.collecter();
    console.log(`[${this.source()}] ${Object.keys(mesures).length} mesures en ${Date.now() - debut}ms`);
  }
}

class CollecteurSnmp extends Collecteur {
  source(): string { return "snmp"; }
  async collecter(): Promise<Record<string, number>> {
    return { charge: 42 }; // ... vraie implémentation
  }
}

// new Collecteur(); // Erreur : classe abstraite
await new CollecteurSnmp().cycle();
```

> Abstrait vs interface : l'abstrait peut porter du **code partagé** et un état ; l'interface ne porte que le contrat. Pour un plugin système (collecteurs, exporters), l'abstrait + template method est idéal.

---

## 32. Implémenter des interfaces avec des classes

```typescript
interface Pingable {
  hote: string;
  ping(): Promise<number>;
}

class ServeurPing implements Pingable {
  constructor(public hote: string) {}
  async ping(): Promise<number> {
    return 12; // ... vraie implémentation
  }
}
```

**`implements` ne vérifie que la forme publique** : il ne contrôle pas les types internes. Et une classe peut implémenter plusieurs interfaces :

```typescript
class MultiSonde implements Pingable, CollecteurDeMetriques { /* ... */ }
```

> Astuce : `implements` + `satisfies`-like via `implements` sur des object literals n'existe pas — pour un objet littéral conforme à une interface, annotez-le (`const x: MonInterface = {...}`) ou utilisez `satisfies` (section 23).

---

## 33. Génériques : les fondamentaux

Les génériques paramètrent un type par **un autre type**. Sans eux, on perd l'information (retour `any`) ou on duplique.

```typescript
// Sans générique : le type de retour est perdu
function premierAny(tab: any[]): any { return tab[0]; }

// Avec générique : le type circule
function premier<T>(tab: T[]): T | undefined {
  return tab[0];
}

const n = premier([1, 2, 3]);       // T = number → number | undefined
const s = premier(["a", "b"]);      // T = string → string | undefined
const srv = premier([{ nom: "x" }]); // { nom: string } | undefined
```

**Génériques sur interfaces, classes, alias** :

```typescript
interface ReponseApi<T> {
  data: T;
  erreur?: string;
  dureeMs: number;
}

class File<T> {
  private items: T[] = [];
  enfiler(x: T): void { this.items.push(x); }
  defiler(): T | undefined { return this.items.shift(); }
}

type Paire<K, V> = [K, V];
```

**Inférence vs explicite** : TypeScript infère `T` depuis les arguments. Précisez-le quand l'inférence se trompe ou pour contraindre l'appel :

```typescript
const x = premier<number>([1, 2, 3]); // explicite, rarement nécessaire
```

---

## 34. Contraintes (extends) et valeurs par défaut

`extends` sur un paramètre générique = **contrainte** : « T doit au moins avoir cette forme ».

```typescript
function nomDe<T extends { nom: string }>(e: T): string {
  return e.nom.toUpperCase();
}
nomDe({ nom: "srv-01", ip: "x" }); // OK
nomDe({ ip: "x" });                // Erreur : pas de 'nom'
```

**Valeur par défaut** :

