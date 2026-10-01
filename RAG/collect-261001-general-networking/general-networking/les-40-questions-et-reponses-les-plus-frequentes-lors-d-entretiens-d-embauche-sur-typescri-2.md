---
id: collect-261001-general-networking/general-networking/les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescri-2
title: "les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20.md
source_anchor: ""
source_lines: [159, 410]
sha256: 75bcc61bb2826fa11dfadcb28ae1661cc6297cc98de4a9f83bc2a851ebe7c708
---

# les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20

Examinons maintenant quelques concepts clés du système de types TypeScript fréquemment abordés lors des entretiens techniques.

### 10. Quelle est la différence entre les interfaces et les alias de type dans TypeScript ?

Les interfaces et les alias de type décrivent tous deux la forme des objets, mais ils ont des objectifs légèrement différents.

Les interfaces peuvent être étendues à l'aide de extends, ce qui les rend idéales pour les modèles d'objets hiérarchiques.

Les alias de type peuvent représenter des unions, des intersections, des types primitifs et des types composites plus complexes.

Voici un exemple :

```
interface Person {
  name: string;
  age: number;
}
type Employee = {
  name: string;
  department: string;
};
```
Veuillez utiliser des interfaces lorsque vous définissez des formes d'objets ou des contrats de classe.

Utilisez des alias de type lorsque vous avez besoin de types flexibles et composables, tels que des unions ou des intersections.

### 11. Que sont les types union et intersection dans TypeScript ?

Les types union permettent à une variable de contenir plusieurs types possibles, tandis que les types intersection combinent plusieurs types en un seul.

Voici un exemple d'union :

```
let id: string | number;
id = ‘abc’; // OK
id = 123; // OK
```
Voici un exemple d'intersection :

```
type Admin = { name: string };
type Permissions = { canEdit: boolean };
type AdminUser = Admin & Permissions;
```
Les types union sont fréquemment utilisés dans les paramètres de fonction.

Les types d'intersection sont fréquents dans les compositions d'objets complexes. sont fréquents dans les compositions d'objets complexes.

### 12. Qu'est-ce que le rétrécissement de type et comment fonctionnent les gardes de type ?

Le rétrécissement de type permet à TypeScript de déduire un type plus spécifique en fonction du flux de contrôle.

Voici un exemple :

```
function printId(id: string | number) {
  if (typeof id === "string") {
    console.log(id.toUpperCase());
  } else {
    console.log(id.toFixed());
  }
}
```
Les protections de type garantissent la sécurité des opérations d'exécution en vérifiant le type avant d'effectuer des actions spécifiques.

Vous pouvez utiliser :

- 
`typeof` : pour les primitives
- 
`instanceof` : pour les cours
- 
`Custom type guards` : l'utilisation du paramètre suit la syntaxe`Type`

Les protections de type améliorent la sécurité des types et réduisent les erreurs d'exécution.

### 13. Que sont les types littéraux et les unions discriminées dans TypeScript ?

Les types littéraux limitent une variable à un ensemble spécifique de valeurs prédéfinies.

Voici un exemple :

`let direction: "up" | "down" | "left" | "right";`
Les unions discriminées combinent des types littéraux avec des variantes d'objets pour la correspondance de motifs.

Voici un exemple :

```
type Shape =
  | { kind: "circle"; radius: number }
  | { kind: "square"; side: number };
function area(shape: Shape): number {
  if (shape.kind === "circle") return Math.PI * shape.radius ** 2;
  return shape.side ** 2;
}
```
Les unions discriminées permettent une ramification sécurisée sur les types d'objets et simplifient la logique dans les applications complexes.

### 14. Que représentent keyof, typeof et in dans TypeScript ?

Ces opérateurs fournissent une réflexion puissante au niveau du type.

- 
`Keyof` : récupère les noms des propriétés d'un type
- 
`typeof` : obtient le type d'une valeur
- 
`in` : utilisé dans les types mappés

Voici un exemple :

```
type Keys = keyof User; // "name" | "age"
const person = { name: "Alice", age: 25 };
type PersonType = typeof person; // { name: string; age: number }
```
### 15. Que sont les signatures d'index dans TypeScript ?

Les signatures d'index permettent d'utiliser des objets avec des clés dynamiques.

Voici un exemple :

```
interface Errors {
  [key: string]: string;
}
const messages: Errors = {
  email: "Invalid email",
  password: "Required"
};
```
Ils sont utilisés pour les dictionnaires, les objets de type carte ou les modèles de configuration dynamique.

### 16. Qu'est-ce que le typage structurel dans TypeScript ?

Le typage structurel (ou duck typing) signifie que les types sont compatibles en fonction de leur structure, et non de leur nom.

Voici un exemple :

```
interface Point { x: number; y: number; }
let p = { x: 10, y: 20, z: 30 };
let q: Point = p; // OK because structural match
```
Cela rend TypeScript flexible et lui permet de bien fonctionner avec les modèles JavaScript.

### 17. Qu'est-ce que la fusion de déclarations dans TypeScript ?

La fusion de déclarations se produit lorsque TypeScript combine plusieurs déclarations portant le même nom.

Voici un exemple de fusion de deux interfaces :

```
interface User { name: string; }
interface User { age: number; }
const u: User = { name: "Rob", age: 30 }; // merged
```
Ceci est utile pour étendre des types tiers ou enrichir des bibliothèques.

## Questions d'entretien avancées sur TypeScript

Voici des questions et réponses avancées relatives à TypeScript pour un entretien.

### 18. Que sont les génériques dans TypeScript et pourquoi sont-ils utiles ?

Les génériques vous permettent de créer des fonctions et des classes réutilisables et sécurisées qui fonctionnent avec plusieurs types de données.

Voici un exemple de fonction générique :

```
function identity<T>(value: T): T {
  return value;
}
let output = identity<string>("DataCamp");
let outputNumber = identity<number>(42);
```
Voici un exemple de classe générique :

```
class Box<T> {
  content: T;
  constructor(content: T) {
    this.content = content;
  }
  getContent(): T {
    return this.content;
  }
}
let stringBox = new Box("Hello");
let numberBox = new Box(123);
```
Les meilleures pratiques comprennent :

- 
Veuillez utiliser des noms descriptifs tels que - 
Évitez toute complexité inutile.
- 
Veuillez utiliser des contraintes (extends) pour limiter les types acceptables.

Les génériques améliorent la réutilisabilité du code en réduisant les doublons et en imposant une typage cohérent.

### 19. Que sont les types utilitaires TypeScript et comment sont-ils utilisés ?

TypeScript fournit des types utilitaires intégrés pour manipuler et transformer efficacement les types existants.

Voici un exemple :

```
interface Todo {
  title: string;
  description: string;
  completed: boolean;
}
type PartialTodo = Partial<Todo>; // all properties optional
type RequiredTodo = Required<Todo>; // all properties required
type TodoPreview = Pick<Todo, "title" | "completed">; // select specific properties
type TodoWithoutDescription = Omit<Todo, "description">; // exclude properties
```
Ces éléments sont particulièrement utiles lorsque l'on travaille avec des réponses API, des formulaires ou des accessoires de composants.

Un conseil serait d'utiliser autant que possible les types utilitaires afin de simplifier le code et de garantir la cohérence des types.

### 20. Que sont les types conditionnels et mappés dans TypeScript ?

Les types conditionnels vous permettent de définir une logique au sein des types en fonction des relations entre eux.

Voici un exemple :

```
type IsString<T> = T extends string ? "yes" : "no";
type Result1 = IsString<string>; // "yes"
type Result2 = IsString<number>; // "no"
```
Les types mappés vous permettent de transformer toutes les propriétés d'un type en une seule fois.

Voici un exemple :

```
type ReadonlyTodo = {
  readonly [K in keyof Todo]: Todo[K];
};
type OptionalTodo = {
  [K in keyof Todo]?: Todo[K];
};
```
Les principales différences sont les suivantes :

- keyof extrait les noms des propriétés d'un type.
- typeof permet d'obtenir le type d'une variable ou d'un objet. Leur combinaison permet une manipulation dynamique des types et une refactorisation plus sécurisée.

### 21. À quoisertle mot-clé inferdans TypeScript ?

