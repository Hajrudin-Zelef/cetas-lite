---
id: collect-240926-datacamp/datacamp/les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescri-1
title: "les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20"
domain: datacamp
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: ["inference", "parameters"]
source: docs/RAG/clean_en/datacamp/les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20.md
source_anchor: ""
source_lines: [1, 210]
sha256: 8632868d4a2b5fb8206160735310371dc2b86958e73bdb9451368ad53254f608
---

# les-40-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-sur-typescript-pour-20

<!-- source: https://www.datacamp.com/fr/blog/typescript-interview-questions -->

TypeScript has evolved from a niche JavaScript superset into a core language for modern web development.

Throughout my experience in the field of web applications, I have seen how strong typing and maintainability can significantly reduce bugs and improve collaboration between teams.

TypeScript offers all these features, which explains why it has become an industry standard for modern development.

From Microsoft (its creator) to Google, Airbnb, and Netflix, many organizations rely on TypeScript to power products written in frameworks such as React, Angular, and Node.js.

In my work, I have often observed job postings where proficiency in TypeScript is mentioned as a fundamental requirement, not as an optional skill.

Having used TypeScript myself, I can say without hesitation that understanding its types, interfaces, and advanced features is essential for any developer preparing for technical interviews.

## Basic TypeScript Interview Questions

Let's now look at some common TypeScript job interview questions and their connection to code safety and clarity.

### 1. Why choose TypeScript over JavaScript?

TypeScript adds static typing and compile-time error checking, which helps developers catch bugs early and makes codebases more predictable.

This ensures better maintainability and scalability.

### 2. What problems does TypeScript solve?

TypeScript addresses JavaScript's weaknesses in type safety, scalability, and readability.

Enforcing clear type definitions helps prevent subtle runtime bugs and makes refactoring safer.

### 3. What are the main benefits of using static typing in TypeScript?

Static typing improves code reliability, autocompletion, and developer productivity.

It makes it easier to catch type errors early and allows IDEs to provide more comprehensive tooling support.

### 4. What are type annotations and why are they useful?

Type annotations allow developers to explicitly declare variable types.

This improves code readability and reduces runtime errors.

Here is an example:

```
let username: string = "DataCamper";
let age: number = 25;
let isAdmin: boolean = true;
```
In the code above, the compiler ensures that only the correct type can be assigned to each variable.

For example, if you try to assign a string to age, TypeScript will throw an error, allowing errors to be caught quickly.

Explicit typing is especially important for function parameters, return types, and API contracts, as it communicates expectations to other developers.

### 5. What is type inference in TypeScript?

TypeScript can automatically infer types based on assigned values.

This means it is not always necessary to explicitly declare a type. TypeScript does it for you.

Here is an example:

```
let count = 10;  // inferred as number
count = "hello"; // Error: Type 'string' is not assignable to type 'number'
```
In this example, TypeScript infers that count must be a number based on its initial value, so subsequently assigning a string results in a compilation error.

Best practices for type inference include:

- Use explicit types for exported functions, classes, and interfaces to provide clear contracts.
- Allow inference for local variables when the type is obvious.
- Avoid using it unless absolutely necessary, as it bypasses TypeScript's safety features.

### 6. How does TypeScript handle arrays, tuples, and enums?

TypeScript allows developers to enforce strict types on arrays, tuples, and enums to improve clarity and prevent runtime errors.

Here is an example of a typed array:

```
let scores: number[] = [95, 80, 85];
scores.push(100); // OK
scores.push("A+"); // Error
```
Typed arrays ensure that all elements in the array are of the same type.

Here is an example of a tuple:

`let user: [string, number] = ["Don", 25];`
Tuples define fixed-length arrays with specific types for each position.

They are useful for structured data where the order and type of elements matter.

Here is an example of an enum:

```
enum Status {
  Active,
  Inactive,
  Pending,
}
let currentStatus: Status = Status.Active;
```
Enums provide named constants, improving readability and reducing invalid values.

They are especially useful for status codes, roles, and configuration options.

### 7. What is the difference between the "any", "unknown", and "never" types?

The following table highlights the main differences between the any, unknown, and never types in TypeScript.

| **Type** | **Description** | 
| any | Completely disables type checking, allowing any value. | 
| unknown | A safer alternative to any, it is important to check the type before use. | 
| never | Represents values that never occur (for example, functions that always throw an exception). | 

Here is an example:

```
function fail(): never {
  throw new Error("Something went wrong");
}
```
One tip would be to favor the unknown over any other element in order to preserve type safety while retaining some flexibility.

### 8. What is the difference between null and undefined in TypeScript?

The main differences between null and undefined are as follows:

- 
`undefined` means that a variable has been declared but no value has been assigned to it.
- 
`null` is an explicit value that means "no value."

TypeScript's strictNullChecks mode treats them as distinct types, which helps avoid accidental null-related errors.

Here is an example:

```
let a: string | null = null;
let b: string | undefined = undefined;
```
### 9. What does the strict compilation option do?

The `strict` flag enables all of TypeScript's strict type-checking rules, such as:

- **strictNullChecks:** Prevents the assignment of `null` and `undefined` values to variables unless explicitly allowed, which helps avoid null reference errors at runtime.
- 
**noImplicitAny:** Requires all variables to have explicit or inferred types, thereby preventing TypeScript from defaulting to the `any` type.
- 
**Strict function types:** Applies stricter rules for function type compatibility, detecting incompatibilities between function parameter and return types.
- 
**strictBindCallApply:** Ensures that the `bind`, `call`, and `apply` methods are used with the appropriate argument types for functions.

It ensures safer and more predictable code, and quickly detects subtle bugs.

## TypeScript Type System Interview Questions

Let's now look at a few key concepts of the TypeScript type system that are frequently discussed in technical interviews.

### 10. What is the difference between interfaces and type aliases in TypeScript?

Interfaces and type aliases both describe the shape of objects, but they have slightly different purposes.

Interfaces can be extended using extends, which makes them ideal for hierarchical object models.

Type aliases can represent unions, intersections, primitive types, and more complex composite types.

Here is an example:

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
Please use interfaces when defining object shapes or class contracts.

Use type aliases when you need flexible and composable types, such as unions or intersections.

### 11. What are union and intersection types in TypeScript?

Union types allow a variable to contain multiple possible types, while intersection types combine multiple types into one.

Here is an example of a union:

```
let id: string | number;
id = ‘abc’; // OK
id = 123; // OK
```
Here is an example of an intersection:

```
type Admin = { name: string };
type Permissions = { canEdit: boolean };
type AdminUser = Admin & Permissions;
```
Union types are frequently used in function parameters.

Intersection types are common in complex object compositions. are common in complex object compositions.

### 12. What is type narrowing and how do type guards work?

