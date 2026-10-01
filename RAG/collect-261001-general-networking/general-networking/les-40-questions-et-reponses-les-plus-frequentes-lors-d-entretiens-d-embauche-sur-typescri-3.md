---
id: collect-261001-general-networking/general-networking/les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescri-3
title: "les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20.md
source_anchor: ""
source_lines: [411, 644]
sha256: c48d141f2ec1fb2fab7532978d26f1d0c4790572d27d071f311a334d62bd38a2
---

# les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20

`infer` vous permet d'extraire (déduire) un type à l'intérieur d'un type conditionnel.

Voici un exemple :

```
type ReturnType<T> =
  T extends (...args: any[]) => infer R ? R : never;
type R = ReturnType<() => number>; // number
```
Souvent utilisé dans les utilitaires avancés et les transformations de types génériques.

### 22. Que sont les types littéraux de modèle ?

Les types littéraux de modèle construisent des types de chaînes à l'aide de l'interpolation.

Voici un exemple :

```
type Event = ${string}Changed;
let e: Event = "nameChanged"; // OK
```
Ceci est utile pour :

- Dénomination de l'événement
- Modèles de classes CSS
- Modèles de routes API

### 23. Que signifient « satisfies » et « as const » dansTypeScript ?

`satisfies` garantit qu'une valeur correspond à un type tout en conservant son propre type littéral. 

Par exemple :

```
const config = {
  mode: "dark",
  version: 1
} satisfies { mode: string; version: number };
```
`as const` rend toutes les propriétés en lecture seule et littérales. 

Par exemple :

```
const directions = ["up", "down"] as const;
// type is readonly ["up", "down"]
```
`satisfies` et `as const` sont tous deux utiles dans React, les objets de configuration et les unions discriminées.

## Questions d'entretien sur les fonctions et les objets TypeScript

Dans cette section, nous examinerons les questions d'entretien TypeScript liées aux fonctions et aux objets.

### 24. Comment fonctionne la surcharge de fonctions dans TypeScript ?

La surcharge de fonction vous permet de définir plusieurs signatures de fonction pour différents types ou modèles d'arguments.

Voici un exemple :

```
function add(a: number, b: number): number;
function add(a: string, b: string): string;
function add(a: any, b: any): any {
  return a + b;
}
let sumNumbers = add(5, 10); // 15
let sumStrings = add("Hello, ", "World"); // "Hello, World"
```
Les meilleures pratiques comprennent :

- Veuillez déclarer toutes les surcharges possibles avant la mise en œuvre.
- Veuillez utiliser des gardes de type pour gérer plusieurs types en toute sécurité.

### 25. Que sont les paramètres facultatifs et les paramètres de reste dans TypeScript ?

Les paramètres facultatifs et de repos rendent les fonctions plus flexibles et expressives.

Voici un exemple :

```
function greet(name: string, age?: number) {
  console.log(Hello ${name}, age: ${age ?? "unknown"});
}
function sum(...numbers: number[]) {
  return numbers.reduce((a, b) => a + b, 0);
}
```
Les règles principales pour l'utilisation des paramètres facultatifs et résiduels sont les suivantes :

- 
Veuillez utiliser `?` pour les paramètres facultatifs.
- 
Veuillez utiliser `...` pour les paramètres de repos (arguments variadiques).

Ces compétences sont fréquemment évaluées lors d'entretiens afin d'évaluer votre capacité à gérer des entrées de fonctions dynamiques.

### 26. Quelle est la différence entre readonly et const dans TypeScript ?

`const` s'applique aux variables, tandis que `readonly` s'applique aux propriétés des objets.

Voici un exemple :

```
interface User {
  readonly id: number;
  name: string;
}
const user: User = { id: 1, name: "Bob" };
// user.id = 2; // Error: Cannot assign to 'id'
const { id, name } = user;
console.log(User ${name} has ID ${id});
```
Voici quelques bonnes pratiques pour l'utilisation de `readonly` et `const`:

- Veuillez utiliser readonly pour les objets qui ne doivent pas être modifiés, tels que les données renvoyées par une API.
- Veuillez utiliser const pour les variables qui ne doivent pas être réaffectées.
- Combinez la déstructuration avec les annotations de type pour obtenir un code plus clair et auto-documenté.

### 27. Qu'est-ce que la variance (covariance et contravariance) des types de fonctions ?

La variance de type de fonction décrit comment TypeScript détermine si un type de fonction peut remplacer un autre en toute sécurité, en fonction de ses types de paramètres et de ses types de retour.

Types de retour : Covariant

Une fonction peut renvoyer un type plus spécifique que celui attendu par l'appelant. Par exemple :

```
type Animal = { name: string };
type Dog = Animal & { bark: () => void };
let f: () => Animal;
let g: () => Dog = () => ({ name: "Rex", bark() {} });
f = g; // OK: Dog is a subtype of Animal
```
Types de paramètres : Contravariant

Une fonction peut accepter des types de paramètres plus généraux que prévu. Par exemple :

```
type Dog = { name: string; bark: () => void };
type Animal = { name: string };
type HandleDog = (dog: Dog) => void;
type HandleAnimal = (animal: Animal) => void;
let handleDog: HandleDog = d => console.log(d.bark());
let handleAnimal: HandleAnimal = a => console.log(a.name);
// A function that accepts an Animal is more general,
// so it can stand in for one that accepts a Dog
handleDog = handleAnimal; // OK
```
### 28. Quels sont ces types dans TypeScript ?

`this` Les types permettent aux méthodes de renvoyer le même type que la classe, ce qui est utile pour les API fluides :

```
class Builder {
  setName(name: string): this {
    // ...
    return this;
  }
}
new Builder().setName("Test"); // chaining works
```
## Questions d'entretien sur la programmation orientée objet en TypeScript

Les questions suivantes couvrent les concepts clés de la programmation orientée objet (OOP) en TypeScript fréquemment abordés lors des entretiens techniques.

### 29. Comment TypeScript prend-il en charge la programmation orientée objet (POO) ?

TypeScript prend en charge les principes de la programmation orientée objet (POO) tels que les classes, l'héritage et les modificateurs d'accès.

Les classes définissent des modèles réutilisables pour les objets, et l'héritage permet à une classe d'étendre une autre classe afin de partager des propriétés et des comportements.

Voici un exemple :

```
class Animal {
  constructor(public name: string) {}
  move(distance: number) {
    console.log(${this.name} moved ${distance}m.);
  }
}
class Dog extends Animal {
  bark() {
    console.log("Woof!");
  }
}
const dog = new Dog("Buddy");
dog.bark();       // Woof!
dog.move(10);     // Buddy moved 10m.
```
Les modificateurs d'accès comprennent :

- public : accessible partout
- protégé : accessible dans la classe et ses sous-classes
- privé : accessible uniquement au sein de la classe qui le déclare

Il est recommandé d'utiliser des modificateurs d'accès pour garantir l'encapsulation, ce qui empêche toute modification involontaire de la logique ou de l'état interne de la classe.

### 30. Quelle est la différence entre les classes abstraites et les interfaces dans TypeScript ?

Les classes abstraites et les interfaces définissent toutes deux des contrats pour les objets ou les classes, mais elles ont des objectifs différents.

Les classes abstraites peuvent inclure une implémentation partagée et des méthodes abstraites.

Les interfaces décrivent uniquement la structure, et non l'implémentation.

Voici un exemple :

```
abstract class Shape {
  abstract getArea(): number; // must be implemented
}
class Circle extends Shape {
  constructor(public radius: number) {
    super();
  }
  getArea(): number {
    return Math.PI * this.radius ** 2;
  }
}
```
Les meilleures pratiques comprennent :

- Veuillez utiliser les classes abstraites lorsque vous avez besoin d'une logique commune à plusieurs classes apparentées.
- Veuillez utiliser des interfaces pour des contrats de type flexible et une conception de code découplée.

### 31. Quelle est la différence entre les modificateurs privé, protégé et public ?

Voici une comparaison des modificateurs d'accès de TypeScript et de ce qu'ils contrôlent :

| **Modificateur** | **Description** | 
| public | Accessible partout. | 
| protégé | Accessible au sein de la classe et de ses sous-classes. | 
| privé | Accessible uniquement à l'intérieur de la classe déclarante. | 

