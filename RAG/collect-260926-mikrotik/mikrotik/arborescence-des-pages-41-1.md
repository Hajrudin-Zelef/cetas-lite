---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-41-1
title: "arborescence-des-pages-41"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-41.md
source_anchor: ""
source_lines: [1, 216]
sha256: 71ea3d93d1c79a801129dc9a28cc229cf1b20e76eda22c47e4957ccd5634463f
---

# arborescence-des-pages-41

## Scripting language manual

This manual provides an introduction to RouterOS's built-in powerful scripting language.

Scripting host provides a way to automate some router maintenance tasks by means of executing user-defined scripts bounded to some event occurrence.

Scripts can be stored in the Script repository or can be written directly to the console. The events used to trigger script execution include, but are not limited to the System Scheduler, the Traffic Monitoring Tool, and the Netwatch Tool generated events.

If you are already familiar with scripting in RouterOS, you might want to see our Tips & Tricks.

### Line structure

The RouterOS script is divided into a number of command lines. Command lines are executed one by one until the end of the script or until a runtime error occurs.

#### Command-line

The RouterOS console uses the following command syntax:

`[prefix] [path] command [uparam] [param=[value]] .. [param=[value]]`

- [prefix] - ":" or "/" character which indicates if a command is ICE or path. It may not be required.
- [path] - relative path to the desired menu level. It may not be required.
- command - one of the commands available at the specified menu level.
- [uparam] - unnamed parameter, must be specified if the command requires it.
- [params] - a sequence of named parameters followed by respective values

The end of the command line is represented by the token *“;”* or *NEWLINE*. Sometimes *“;”* or *NEWLINE* is not required to end the command line.

Single command inside `(), [] or {}` does not require any end-of-command character. The end of the command is determined by the content of the whole script

Each command line inside another command line starts and ends with square brackets "[ ]" (command concatenation).

Notice that the code above contains three command lines:

- :put
- /ip route get
- find gateway=1.1.1.1

Command-line can be constructed from more than one physical line by following line joining rules.

#### Physical Line

A physical line is a sequence of characters terminated by an end-of-line (EOL) sequence. Any of the standard platform line termination sequences can be used:

- **Unix** – ASCII LF;
- **Windows** – ASCII CR LF;
- **mac** – ASCII CR;

Standard C conventions for newline characters can be used ( the \n character).

#### Comments

The following rules apply to a comment:

- A comment starts with a hash character (#) and ends at the end of the physical line.
- RouterOS does not support multiline comments.
- If a **#** character appears inside the string it is not considered a comment.

##### Example

#### Line joining

Two or more physical lines may be joined into logical lines using the backslash character (\).

The following rules apply to using backslash as a line-joining tool:

- A line ending in a backslash cannot carry a comment.
- A backslash does not continue a comment.
- A backslash does not continue a token except for string literals.
- A backslash is illegal elsewhere on a line outside a string literal.

##### Example

#### Whitespace between tokens

Whitespace can be used to separate tokens. Whitespace is necessary between two tokens only if their concatenation could be interpreted as a different token. Example:

Whitespace characters are not allowed

- between '<parameter>='
- between 'from=' 'to=' 'step=' 'in=' 'do=' 'else='

Example:

##### Scopes

Variables can be used only in certain regions of the script called scopes. These regions determine the visibility of the variable. There are two types of scopes - global and local. A variable declared within a block is accessible only within that block and blocks enclosed by it, and only after the point of declaration.

###### Global scope

Global scope or root scope is the default scope of the script. It is created automatically and can not be turned off.

###### Local scope

User can define their own groups to block access to certain variables, these scopes are called local scopes. Each local scope is enclosed in curly braces ("{ }").

In the code above variable, b has local scope and will not be accessible after a closing curly brace.

So for example, the defined local variable will not be visible in the next command line and will generate a syntax error

[admin@MikroTik] > :local myVar a;
[admin@MikroTik] > :put $myVar
syntax error (line 1 column 7)

Note that even variable can be defined as global, it will be available only from its scope unless it is not referenced to be visible outside of the scope.

The code above will output 3, because outside of the scope b is not visible.

The following code will fix the problem and will output 7:

### Keywords

The following words are keywords and cannot be used as variable and function names:

and       or       in

### Delimiters

The following tokens serve as delimiters in the grammar:

```
()  []  {}  :   ;   $   / 
```
### Data types

RouterOS scripting language has the following data types:

| Type | Description | 
|---|---|
| **num (number)** | - 64bit signed integer, possible hexadecimal input; | 
| **bool (boolean)** | - values can bee `true` or`false` ; | 
| **str (string)** | - character sequence; | 
| **ip** | - IP address; | 
| **ip-prefix** | - IP prefix; | 
| **ip6** | - IPv6 address | 
| **ip6-prefix** | - IPv6 prefix | 
| **id (internal ID)** | - hexadecimal value prefixed by '*' sign. Each menu item has an assigned unique number - internal ID; | 
| **time** | - date and time value; | 
| **array** | - sequence of values organized in an array; | 
| **nil** | - default variable type if no value is assigned; | 

#### Constant Escape Sequences

Following escape sequences can be used to define certain special characters within a string:

| **\"** | Insert double quote | 
| **\\** | Insert backslash | 
| **\n** | Insert newline | 
| **\r** | Insert carriage return | 
| **\t** | Insert horizontal tab | 
| **\$** | Output $ character. Otherwise, $ is used to link the variable. | 
| **\?** | ~~Output ? character. Otherwise ? is used to print "help" in the console.~~ Removed since v7.1rc2 | 
| **\_** | - space | 
| **\a** | - BEL (0x07) | 
| **\b** | - backspace (0x08) | 
| **\f** | - form feed (0xFF) | 
| **\v** | Insert vertical tab | 
| **\xx** | A print character from hex value. Hex numbers should use capital letters. | 

##### Example

:put "\48\45\4C\4C\4F\r\nThis\r\nis\r\na\r\ntest";

which will show on the display`HELLO`

This

is

a

test

### Operators

#### Arithmetic Operators

Usual arithmetic operators are supported in the RouterOS scripting language

| Operator | Description | Example | 
|---|---|---|
| **"+"** | binary addition | `:put (3+4);` | 
| **"-"** | binary subtraction | `:put (1-6);` | 
| **"*"** | binary multiplication | `:put (4*5);` | 
| **"/"** | binary division | `:put (10 / 2); :put ((10)/2)` | 
| **"%"** | modulo operation | `:put (5 % 3);` | 
| **"-"** | unary negation | `{ :local a 1; :put (-a); }` | 

**Note:** for the division to work you have to use braces or spaces around the dividend so it is not mistaken as an IP address

#### Relational Operators

| Operator | Description | Example | 
|---|---|---|
| **"<"** | less | `:put (3<4);` | 
| **">"** | greater | `:put (3>4);` | 
| **"="** | equal | `:put (2=2);` | 
| **"<="** | less or equal |  | 
| **">="** | greater or equal |  | 
| **"!="** | not equal |  | 

To negate an expression, you can use "expression=false". To print all interfaces that are not "ethernet", you can use expression negation like this:

Or to do the opposite, you can use "expression=true":

#### Logical Operators

| Operator | Description | Example | 
|---|---|---|
| **“!”** | logical NOT | `:put (!true);` | 
| **“&&”, “and”** | logical AND | `:put (true&&true)` | 
| **“\|\|”, “or”** | logical OR | `:put (true\|\|false);` | 
| **“in”** |  | `:put (1.1.1.1/32 in 1.0.0.0/8);` | 

#### Bitwise Operators

