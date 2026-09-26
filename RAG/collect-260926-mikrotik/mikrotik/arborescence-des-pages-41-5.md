---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-41-5
title: "arborescence-des-pages-41"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-41.md
source_anchor: ""
source_lines: [493, 581]
sha256: 4801f40855c369c0b627524169b14df4f237da96935dc749879a9534e88aa759
---

# arborescence-des-pages-41

- **on event** - scripts are executed automatically on some facility events ( scheduler, netwatch, VRRP)
- **by another script** - running script within the script is allowed
- **manually** - from console executing a run command or in winbox

Only scripts (including schedulers, netwatch, etc) with equal or higher permission rights can execute other scripts.

When executing script from GUI or CLI, user permissions are used. To run a script with script permissions, a script must be executed from CLI with additional "*use-script-permissions*" parameter.

| Property | Description | 
|---|---|
| **comment** (*string* ; Default: ) | Descriptive comment for the script | 
| **dont-require-permissions** (*yes \| no* ; Default:*no* ) | Bypass permissions check when the script is being executed, useful when scripts are being executed from services that have limited permissions, such as Netwatch | 
| **name** (*string* ; Default:*"Script[num]"* ) | name of the script | 
| **policy** (*string* ; Default: ftp,reboot,read,write,policy,test,password,sniff,sensitive,romon) | list of applicable policies:  Read more detailed policy descriptions here | 
| **source** (*string* ;) | Script source code | 

Read-only status properties:

| Property | Description | 
|---|---|
| **last-started** (*date* ) | Date and time when the script was last invoked. | 
| **owner** (*string* ) | The user who created the script | 
| **run-count** (*integer* ) | Counter that counts how many times the script has been executed | 

Menu specific commands

| Command | Description | 
|---|---|
| **run** (*run [id\|name]* ) | Execute the specified script by ID or name using user permissions. | 
| **use-script-permissions** | Additional parameter to execute script using script permissions. | 

### Environment

**Sub-menu level:**

- `/system script environment`
- `/environment`

Contains all user-defined variables and their assigned values.

[admin@MikroTik] > :global example;
[admin@MikroTik] > :set example 123
[admin@MikroTik] > /environment print  
"example"=123


Read-only status properties:

| Property | Description | 
|---|---|
| **name** (*string* ) | Variable name | 
| **user** (*string* ) | The user who defined variable | 
| **value** () | The value assigned to a variable | 

### Job

**Sub-menu level:** `/system script job`

Contains a list of all currently running scripts.

Read-only status properties:

| Property | Description | 
|---|---|
| **owner** (*string* ) | The user who is running the script | 
| **policy** (*array* ) | List of all policies applied to the script | 
| **started** (*date* ) | Local date and time when the script was started | 

## Script permissions

There are four ways a script can be run: using script permissions, user permissions, scheduler permissions or on-event permissions (such as `/system routerboard mode-button` settings).

Depending on how a script is called, it may inherit different permissions or use its own.

When using `/system script run`, a script can inherit the permissions of the caller.

In this example, we use an admin user with full permissions and test running a script both with and without the `use-script-permissions` parameter.

Similarly, there are multiple ways to run a script using the scheduler tool. When using the scheduler, the script can run with the scheduler's permissions.

You can also call the script by name using the scheduler, which works the same way as `/system script run use-script-permissions`.

To demonstrate this, we create three schedulers, each configured to run the script in a different way:

After these schedulers run, the logs show that the two methods using `use-script-permissions` or calling the script by name fail due to insufficient permissions.

In contrast, `run-script-scheduler-perms` executes the script successfully, as it inherits the scheduler's permissions.

A script with higher or more permissions than the user/scheduler cannot be run; use-script-permissions won’t override this.
