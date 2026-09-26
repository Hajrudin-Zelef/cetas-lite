---
id: collect-240926-datacamp/datacamp/les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescri-2
title: "les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/clean_en/datacamp/les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20.md
source_anchor: ""
source_lines: [211, 510]
sha256: bc023de037498c67ce6a9c05ffb117f16082c9b6c24fb5b78e3a97d2b63d11a8
---

# les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20

Type narrowing allows TypeScript to infer a more specific type based on the control flow.

Here is an example:

```
function printId(id: string | number) {
  if (typeof id === "string") {
    console.log(id.toUpperCase());
  } else {
    console.log(id.toFixed());
  }
}
```
Type guards ensure the safety of runtime operations by checking the type before performing specific actions.

You can use:

- 
`typeof`: for primitives
- 
`instanceof`: for classes
- 
`Custom type guards`: using the parameter follows the `Type` syntax

Type guards improve type safety and reduce runtime errors.

### 13. What are literal types and discriminated unions in TypeScript?

Literal types restrict a variable to a specific set of predefined values.

Here is an example:

`let direction: "up" | "down" | "left" | "right";`
Discriminated unions combine literal types with object variants for pattern matching.

Here is an example:

```
type Shape =
  | { kind: "circle"; radius: number }
  | { kind: "square"; side: number };
function area(shape: Shape): number {
  if (shape.kind === "circle") return Math.PI * shape.radius ** 2;
  return shape.side ** 2;
}
```
Discriminated unions enable safe branching on object types and simplify logic in complex applications.

### 14. What do keyof, typeof, and in represent in TypeScript?

These operators provide powerful reflection at the type level.

- 
`Keyof`: retrieves the property names of a type
- 
`typeof`: gets the type of a value
- 
`in`: used in mapped types

Here is an example:

```
type Keys = keyof User; // "name" | "age"
const person = { name: "Alice", age: 25 };
type PersonType = typeof person; // { name: string; age: number }
```
### 15. What are index signatures in TypeScript?

Index signatures allow objects with dynamic keys to be used.

Here is an example:

```
interface Errors {
  [key: string]: string;
}
const messages: Errors = {
  email: "Invalid email",
  password: "Required"
};
```
They are used for dictionaries, map-like objects, or dynamic configuration patterns.

### 16. What is structural typing in TypeScript?

Structural typing (or duck typing) means that types are compatible based on their structure, not their name.

Here is an example:

```
interface Point { x: number; y: number; }
let p = { x: 10, y: 20, z: 30 };
let q: Point = p; // OK because structural match
```
This makes TypeScript flexible and allows it to work well with JavaScript patterns.

### 17. What is declaration merging in TypeScript?

Declaration merging occurs when TypeScript combines multiple declarations with the same name.

Here is an example of merging two interfaces:

```
interface User { name: string; }
interface User { age: number; }
const u: User = { name: "Rob", age: 30 }; // merged
```
This is useful for extending third-party types or enriching libraries.

## Advanced TypeScript Interview Questions

Here are advanced TypeScript-related questions and answers for an interview.

### 18. What are generics in TypeScript and why are they useful?

Generics allow you to create reusable, type-safe functions and classes that work with multiple data types.

Here is an example of a generic function:

```
function identity<T>(value: T): T {
  return value;
}
let output = identity<string>("DataCamp");
let outputNumber = identity<number>(42);
```
Here is an example of a generic class:

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
Best practices include:

- 
Please use descriptive names such as - 
Avoid unnecessary complexity.
- 
Please use constraints (extends) to limit acceptable types.

Generics improve code reusability by reducing duplication and enforcing consistent typing.

### 19. What are TypeScript utility types and how are they used?

TypeScript provides built-in utility types to efficiently manipulate and transform existing types.

Here is an example:

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
These are particularly useful when working with API responses, forms, or component props.

One tip would be to use utility types as much as possible to simplify code and ensure type consistency.

### 20. What are conditional and mapped types in TypeScript?

Conditional types allow you to define logic within types based on relationships between them.

Here is an example:

```
type IsString<T> = T extends string ? "yes" : "no";
type Result1 = IsString<string>; // "yes"
type Result2 = IsString<number>; // "no"
```
Mapped types allow you to transform all properties of a type at once.

Here is an example:

```
type ReadonlyTodo = {
  readonly [K in keyof Todo]: Todo[K];
};
type OptionalTodo = {
  [K in keyof Todo]?: Todo[K];
};
```
The main differences are as follows:

- keyof extracts the property names of a type.
- typeof allows you to get the type of a variable or object. Their combination enables dynamic type manipulation and safer refactoring.

### 21. What is the infer keyword used for in TypeScript?

`infer` allows you to extract (infer) a type inside a conditional type.

Here is an example:

```
type ReturnType<T> =
  T extends (...args: any[]) => infer R ? R : never;
type R = ReturnType<() => number>; // number
```
Often used in advanced utilities and generic type transformations.

### 22. What are template literal types?

Template literal types build string types using interpolation.

Here is an example:

```
type Event = ${string}Changed;
let e: Event = "nameChanged"; // OK
```
This is useful for:

- Event naming
- CSS class patterns
- API route patterns

### 23. What do "satisfies" and "as const" mean in TypeScript?

`satisfies` ensures that a value matches a type while preserving its own literal type. 

For example:

```
const config = {
  mode: "dark",
  version: 1
} satisfies { mode: string; version: number };
```
`as const` makes all properties readonly and literal. 

For example:

```
const directions = ["up", "down"] as const;
// type is readonly ["up", "down"]
```
`satisfies` and `as const` are both useful in React, configuration objects, and discriminated unions.

## TypeScript Functions and Objects Interview Questions

In this section, we will look at TypeScript interview questions related to functions and objects.

### 24. How does function overloading work in TypeScript?

Function overloading allows you to define multiple function signatures for different argument types or patterns.

Here is an example:

```
function add(a: number, b: number): number;
function add(a: string, b: string): string;
function add(a: any, b: any): any {
  return a + b;
}
let sumNumbers = add(5, 10); // 15
let sumStrings = add("Hello, ", "World"); // "Hello, World"
```
Best practices include:

- Please declare all possible overloads before the implementation.
- Please use type guards to handle multiple types safely.

### 25. What are optional parameters and rest parameters in TypeScript?

Optional and rest parameters make functions more flexible and expressive.

Here is an example:

```
function greet(name: string, age?: number) {
  console.log(Hello ${name}, age: ${age ?? "unknown"});
}
function sum(...numbers: number[]) {
  return numbers.reduce((a, b) => a + b, 0);
}
```
The main rules for using optional and rest parameters are as follows:

- 
Use `?` for optional parameters.
- 
Use `...` for rest parameters (variadic arguments).

These skills are frequently assessed in interviews to evaluate your ability to handle dynamic function inputs.

### 26. What is the difference between readonly and const in TypeScript?

