---
id: collect-240926-datacamp/datacamp/les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescri-3
title: "les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/clean_en/datacamp/les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20.md
source_anchor: ""
source_lines: [511, 766]
sha256: f6a37b64920b0f3764a9719f2433d80819642012b7dfb1f7c3077e77469a8d85
---

# les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20

`const` applies to variables, while `readonly` applies to object properties.

Here is an example:

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
Here are some best practices for using `readonly` and `const`:

- Use readonly for objects that should not be modified, such as data returned by an API.
- Use const for variables that should not be reassigned.
- Combine destructuring with type annotations to achieve clearer, self-documenting code.

### 27. What is variance (covariance and contravariance) of function types?

Function type variance describes how TypeScript determines whether one function type can safely replace another, based on its parameter types and return types.

Return types: Covariant

A function can return a more specific type than the caller expects. For example:

```
type Animal = { name: string };
type Dog = Animal & { bark: () => void };
let f: () => Animal;
let g: () => Dog = () => ({ name: "Rex", bark() {} });
f = g; // OK: Dog is a subtype of Animal
```
Parameter types: Contravariant

A function can accept more general parameter types than expected. For example:

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
### 28. What are these types in TypeScript?

`this` types allow methods to return the same type as the class, which is useful for fluent APIs:

```
class Builder {
  setName(name: string): this {
    // ...
    return this;
  }
}
new Builder().setName("Test"); // chaining works
```
## TypeScript Object-Oriented Programming Interview Questions

The following questions cover key object-oriented programming (OOP) concepts in TypeScript that are frequently discussed in technical interviews.

### 29. How does TypeScript support object-oriented programming (OOP)?

TypeScript supports object-oriented programming (OOP) principles such as classes, inheritance, and access modifiers.

Classes define reusable templates for objects, and inheritance allows one class to extend another class in order to share properties and behaviors.

Here is an example:

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
Access modifiers include:

- public: accessible everywhere
- protected: accessible within the class and its subclasses
- private: accessible only within the class that declares it

It is recommended to use access modifiers to ensure encapsulation, which prevents unintended modification of the class's internal logic or state.

### 30. What is the difference between abstract classes and interfaces in TypeScript?

Abstract classes and interfaces both define contracts for objects or classes, but they serve different purposes.

Abstract classes can include shared implementation and abstract methods.

Interfaces describe only structure, not implementation.

Here is an example:

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
Best practices include:

- Use abstract classes when you need common logic across multiple related classes.
- Use interfaces for flexible type contracts and decoupled code design.

### 31. What is the difference between the private, protected, and public modifiers?

Here is a comparison of TypeScript access modifiers and what they control:

| **Modifier** | **Description** | 
| public | Accessible everywhere. | 
| protected | Accessible within the class and its subclasses. | 
| private | Accessible only inside the declaring class. | 

Using the appropriate access modifier promotes encapsulation, which ensures that internal state and logic cannot be accessed or modified unexpectedly.

### 32. What is the difference between "implements" and "extends"?

`extends` inherits behavior and properties, while `implements` enforces a contract but does not provide an implementation.

Here is an example:

```
interface Logger { log(msg: string): void; }
class Base { foo() {} }
class MyClass extends Base implements Logger {
  log(msg: string) {}
}
```
### 33. How do mixins work in TypeScript?

Mixins allow you to use multiple reusable class components.

Here is an example:

```
type Constructor<T = {}> = new (...args: any[]) => T;
function Timestamped<TBase extends Constructor>(Base: TBase) {
  return class extends Base {
    timestamp = Date.now();
  };
}
class User {}
const TimestampedUser = Timestamped(User);
```
Useful when you need multiple inheritance patterns.

## Practical TypeScript Interview Questions

These questions focus on concrete TypeScript applications that are frequently discussed in interviews.

### 34. How would you go about migrating a JavaScript project to TypeScript?

Migrating an existing JavaScript codebase to TypeScript should be done gradually to minimize risk and disruption.

The recommended approach is as follows:

1. 
Rename the `.js` files to `.ts`.
2. 
Enable the `allowJs` and `checkJs` settings in the tsconfig.json file.
3. 
Gradually add types, starting with the main functions or modules.
4. 
Use any type temporarily, then replace it with the appropriate types as you refine the code.

Example tsconfig.json file for a gradual migration:

```
{
  "compilerOptions": {
    "allowJs": true,
    "checkJs": true,
    "strict": true
  }
}
```
Gradual migration ensures stability and allows teams to quickly benefit from type checking without a major rewrite.

### 35. How do you use TypeScript with React?

Using TypeScript in React improves type safety for props and state, which helps prevent runtime errors.

Here is an example:

```
interface Props {
  title: string;
}
const Header: React.FC<Props> = ({ title }) => <h1>{title}</h1>;
```
Defining property types allows TypeScript to detect missing or incorrectly typed properties at compile time, which improves reliability and developer productivity.

### 36. How do you use TypeScript with Node.js?

TypeScript improves Node.js applications by adding type safety for API routes, configurations, and middlewares.

Here is an example:

```
import express, { Request, Response } from "express";
const app = express();
app.get("/", (req: Request, res: Response) => res.send("Hello TypeScript"));
```
Typing `Request` and `Response` ensures correct use of parameters and prevents common runtime errors.

TypeScript makes it easier to maintain and safely refactor Node.js APIs.

### 37. How do you handle third-party JavaScript libraries that do not have TypeScript definitions?

Some libraries do not include built-in types. In these cases, you have two main options:

Option 1: Install community-managed types

`npm install --save-dev @types/library-name`
Option 2: Create a custom declaration file.

```
// custom.d.ts
declare module "legacy-library" {
  export function init(): void;
}
```
Here are a few tips:

- 
Always check the `@types` site before writing custom declarations.
- 
Gradually add types for complex libraries to ensure compatibility.

Correct typing of third-party libraries builds confidence and reduces runtime errors.

## TypeScript Configuration and Tooling Interview Questions

