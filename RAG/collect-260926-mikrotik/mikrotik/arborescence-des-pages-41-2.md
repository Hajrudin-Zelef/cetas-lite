---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-41-2
title: "arborescence-des-pages-41"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-41.md
source_anchor: ""
source_lines: [217, 293]
sha256: b7627b75ddfeb28ca5ae0dad7905085dd473e9581a48067e219c728b41eec1d9
---

# arborescence-des-pages-41

Bitwise operators are working only on IP, and IPv6 address data types.

| Operator | Description | Example | 
|---|---|---|
| **“~”** | bit inversion | `:put (~0.0.0.0)``:put (~::ffff)` | 
| **“\|”** | bitwise OR. Performs logical OR operation on each pair of corresponding bits. In each pair the result is “1” if one of the bits or both bits is “1”, otherwise the result is “0”. | `:put (192.168.88.0\|0.0.0.255)``:put (2001::1\|::ffff)` | 
| **“^”** | bitwise XOR. The same as OR, but the result in each position is “1” if two bits are not equal, and “0” if the bits are equal. | `:put (1.1.1.1^255.255.0.0)``:put (2001::ffff:1^::ffff:0)` | 
| **“&”** | bitwise AND. In each pair, the result is “1” if the first and second bit is “1”. Otherwise, the result is “0”. | `:put (192.168.88.77&255.255.255.0)``:put (2001::1111&ffff::)` | 
| **“<<”** | left shift by a given amount of bits | `:put (192.168.88.77<<8)` | 
| **“>>”** | right shift by a given amount of bits | `:put (192.168.88.77>>24)` | 

Calculate the subnet address from the given IP and CIDR Netmask using the "&" operator:

Get the last 8 bits from the given IP addresses:

Use the "|" operator and inverted CIDR mask to calculate the broadcast address:

#### Concatenation Operators

| Operator | Description | Example | 
|---|---|---|
| **"."** | concatenates two strings | `:put ("concatenate" . " " . "string");` | 
| **","** | concatenates two arrays or adds an element to the array | `:put ({1;2;3} , 5 );` | 

It is possible to add variable values to strings without a concatenation operator:

By using $[] and $() in the string it is possible to add expressions inside strings:

#### Other Operators

| Operator | Description | Example | 
|---|---|---|
| **“[]”** | command substitution. Can contain only a single command line | `:put [ :len "my test string"; ];` | 
| **“()”** | subexpression or grouping operator | `:put ( "value is " . (4+5));` | 
| **“$”** | substitution operator | `:global a 5; :put $a;` | 
| **“~”** | the binary operator that matches value against POSIX extended regular expression | Print all routes whose gateway ends with 202: `/ip route print where gateway~"^[0-9 \\.]*202\$"` | 
| **“->”** | Get an array element by key | ``` [admin@x86] >:global aaa {a=1;b=2} [admin@x86] > :put ($aaa->"a") 1 [admin@x86] > :put ($aaa->"b") 2 ```  | 

### Variables

The scripting language has two types of variables:

- global - accessible from all scripts created by the current user, defined by global keyword;
- local - accessible only within the current scope, defined by local keyword.

There can be **undefined** variables. When a variable is undefined, the parser will try to look for variables set, for example, by DHCP lease-script or Hotspot on-login

Every variable, except for built-in RouterOS variables, must be declared before usage by local or global keywords. Undefined variables will be marked as undefined and will result in a compilation error. Example:

Correct code:

The exception is when using variables set, for example, by DHCP lease-script

Valid characters in variable names are letters and digits. If the variable name contains any other character, then the variable name should be put in double quotes. Example:

If a variable is initially defined without value then the variable data type is set to *nil*, otherwise, a data type is determined automatically by the scripting engine. Sometimes conversion from one data type to another is required. It can be achieved using data conversion commands. Example:

Variable names are case-sensitive.

Set command without value will un-define the variable (remove from environment, new in v6.2)

Use quotes on the full variable name when the name of the variable contains operators. Example:

#### Reserved variable names

All built-in RouterOS properties are reserved variables. Variables that will be defined the same as the RouterOS built-in properties can cause errors. To avoid such errors, use custom designations.

For example, the following script will not work:

But will work with different defined variables:

### Commands

#### Global commands

Every global command should start with the *":"* token, otherwise, it will be treated as a variable.

