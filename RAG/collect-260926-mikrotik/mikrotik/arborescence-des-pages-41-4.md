---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-41-4
title: "arborescence-des-pages-41"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-41.md
source_anchor: ""
source_lines: [340, 492]
sha256: 62c1a1be2690e0647b44ab291a56bf5f1829e93d063492fdb9e68135a27812f5
---

# arborescence-des-pages-41

| **tolf** | `:tolf <var>` | converts line endings to LF | ``` :put [:tolf  "AAA\nBBB\nCCC" ]       AAA    BBB       CCC ```  | 
| **nothing** | `:nothing` | return a value of nothing | ``` :if ([:nothing] = 0) do={:put true} else={:put false} false :if ([:nothing] > 0) do={:put true} else={:put false}  false :if ([:nothing] < 0) do={:put true} else={:put false}  true ```  | 

If a variable type conversion function cannot apply the new format to the provided data, the output will be empty.

For example, if you run the :tonum <var> command on a variable with a non-integer value such as "23.8" or "cow&chicken", the result will be empty, and its data type will be shown as nil.

#### 

Menu specific commands

##### Common commands

The following commands are available from most sub-menus:

| Command | Syntax | Description | 
|---|---|---|
| **add** | `add <param>=<value>..<param>=<value>` | add new item | 
| **remove** | `remove <id>` | remove selected item | 
| **enable** | `enable <id>` | enable selected item | 
| **disable** | `disable <id>` | disable selected item | 
| **set** | `set <id> <param>=<value>..<param>=<value>` | change selected items parameter, more than one parameter can be specified at the time. The parameter can be unset by specifying '!' before the parameter. Example:`/ip firewall filter add chain=blah action=accept protocol=tcp port=123 nth=4,2` print set 0 !port chain=blah2 !nth protocol=udp | 
| **get** | `get <id> <param>=<value>` | get the selected item's parameter value | 
| **print** | `print <param><param>=[<value>]` | print menu items. Output depends on the print parameters specified. The most common print parameters are described here | 
| **export** | `export [file=<value>]` | export configuration from the current menu and its sub-menus (if present). If the file parameter is specified output will be written to the file with the extension '.rsc', otherwise the output will be printed to the console. Exported commands can be imported by import command | 
| **edit** | `edit <id> <param>` | edit selected items property in the built-in text editor | 
| **find** | `find <expression>` | Returns list of internal numbers for items that are matched by given expression. For example:  `:put [/interface find name~"ether"]` | 

##### import

The import command is available from the root menu and is used to import configuration from files created by an export command or written manually by hand.

Starting from 7.16.x version, its possible to catch syntax errors:

New parameter *onerror* can be used:

In addition, the *import* command has new options in *verbose* mode - the *dry-run* parameter is specially designed for debugging and can find multiple errors without changing the configuration.

##### print parameters

Several parameters are available for print command:

| Parameter | Description | Example | 
|---|---|---|
| **append** |  |  | 
| **as-value** | print output as an array of parameters and its values | `:put [/ip address print as-value]` | 
| **brief** | print brief description |  | 
| **detail** | print detailed description, the output is not as readable as brief output but may be useful to view all parameters |  | 
| **count-only** | print only count of menu items |  | 
| **file** | print output to a file |  | 
| **follow** | print all current entries and track new entries until ctrl-c is pressed, very useful when viewing log entries | `/log print follow` | 
| **follow-only** | print and track only new entries until ctrl-c is pressed, very useful when viewing log entries | `/log print follow-only` | 
| **timestamp** | parameter only usable with  **follow** command to highlight the "time" when the change was made | `/ip arp print follow timestamp` | 
| **on-event** | parameter only usable with **follow** command to be able to run scripts when new entries appear | `/log print follow on-event={:log info "Someone logged-in"} where message~"admin logged"` | 
| **from** | print parameters only from specified item | `/user print from=admin` | 
| **interval** | continuously print output in a selected time interval, useful to track down changes where `follow` is not acceptable | `/interface print interval=2` | 
| **terse** | show details in a compact and machine-friendly format |  | 
| **value-list** | show values single per line (good for parsing purposes) |  | 
| **without-paging** | If the output does not fit in the console screen then do not stop, print all information in one piece |  | 
| **where** | expressions followed by where parameters can be used to filter outmatched entries | `/ip route print where interface="ether1"` | 
| **about** | returns entries that have the "about" parameter, such as "managed by CAPsMAN "information or warnings | `/interface wifi print where about` | 

More than one parameter can be specified at a time, for example, `/ip route print count-only interval=1 where interface="ether1"`

### Loops and conditional statements

#### Loops

| Command | Syntax | Description | 
|---|---|---|
| **do..while** | `:do { <commands> } while=( <conditions> ); :while ( <conditions> ) do={ <commands> };` | execute commands until a given condition is met. | 
| **for** | `:for <var> from=<int> to=<int> step=<int> do={ <commands> }` | execute commands over a given number of iterations | 
| **foreach** | `:foreach <var> in=<array> do={ <commands> };` | execute commands for each element in a list | 

#### Conditional statement

| Command | Syntax | Description | 
|---|---|---|
| **if** | `:if (<condition>) do={<commands>} else={<commands>}` | If a given condition is `true` then execute commands in the`do` block, otherwise execute commands in the`else` block if specified. | 

Example:

### Functions

Scripting language does not allow you to create functions directly, however, you could use :parse command as a workaround.

Starting from v6.2 new syntax is added to easier define such functions and even pass parameters. It is also possible to return function value with **:return** command.

See examples below:

Notice that there are two ways how to pass arguments:

- pass arg with a specific name ("a" in our example)
- pass value without arg name, in such case arg "1", "2" .. "n" is used.

**Return example**

You can even clone an existing script from the script environment and use it as a function.

**Avoid using parameters with the same name as global variables.**

For example:

The output will be:

1234
lala
global value 123

**Nested function example**

**Note:** to call another function its name needs to be declared (the same as for variables)

### Catch run-time errors

Starting from v6.2 scripting has the ability to catch run-time errors.

For example, the [code]:reslove[/code] command if failed will throw an error and break the script.

```
[admin@MikroTik] > { :put [:resolve www.example.com]; :put "lala";}
failure: dns name does not exist
```
Now we want to catch this error and proceed with our script:

### Operations with Arrays

**Warning:** Key name in the array contains any character other than a lowercase character, it should be put in quotes

For example:

**Loop through keys and values**

"foreach" command can be used to loop through keys and elements:

If the "foreach" command is used with one argument, then the element value will be returned:

**Note:** If the array element has a key then these elements are sorted in alphabetical order, elements without keys are moved before elements with keys and their order is not changed (see example above).

**Change the value of a single array element**

```
[admin@MikroTik] > :global a {x=1; y=2}
[admin@MikroTik] > :set ($a->"x") 5 
[admin@MikroTik] > :environment print 
a={x=5; y=2}
```
## Script repository

**Sub-menu level:** `/system script`

Contains all user-created scripts. Scripts can be executed in several different ways:

