---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-40-1
title: "Login Options"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-40.md
source_anchor: ""
source_lines: [1, 93]
sha256: 8a0510a8bbe988ae9e96a08bd469c9a73892157f593e9f514e41b5e176328b9c
---

# Login Options

The console is used for accessing the MikroTik Router's configuration and management features using text terminals, either remotely using a serial port, telnet, SSH, console screen within WinBox, or directly using monitor and keyboard. The console is also used for writing scripts. This manual describes the general console operation principles. Please consult the Scripting Manual on some advanced console commands and on how to write scripts.

# Login Options

Console login options enable or disable various console features like color, terminal detection, and many other.

Additional login parameters can be appended to the login name after the '+' sign.

```
    login_name ::= user_name [ '+' parameters ]
    parameters ::= parameter [ parameters ]
    parameter ::= [ number ] 'a'..'z'
    number ::= '0'..'9' [ number ]
  
```
If the parameter is not present, then the default value is used. If the number is not present then the implicit value of the parameter is used.

Example: admin+ct80w - will disable console colors, disable auto detection and then set terminal width to 80.

| Param | Default | Implicit | Description | 
|---|---|---|---|
| **"w"** | auto | auto | Set terminal width | 
| **"h"** | auto | auto | Set terminal height | 
| **"c"** | on | off | disable/enable console colors | 
| **"t"** | off | off | Disable auto-detection of terminal capabilities | 
| **"e"** | on | off | Enables "dumb" terminal mode | 

# Banner and Messages

The login process will display the MikroTik banner and short help after validating the user name and password.

After the banner can be printed other important information, like **/system note** set by another admin, the last few critical log messages, demo version upgrade reminder, and default configuration description.

For example, the demo license prompt and last critical messages are printed

# Command Prompt

At the end of the successful login sequence, the login process prints a banner that shows the command prompt, and hands over control to the user.

Default command prompt consists of user name, system identity, and current command path />

For example, change the current path from the root to the interface then go back to the root

Use **up arrow** to recall previous commands (if this is a multiline command, then you can press **F8** in order to expand it) from command history (commands that added sensitive data, like passwords, will not be available in the history), **TAB** key to automatically complete words in the command you are typing, **ENTER** key to execute the command, **Control-C** to interrupt currently running command and return to prompt and **?** to display built-in help, in **RouterOS v7**, **F1** has to be used instead.

The easiest way to log out of the console is to press **Control-D** at the command prompt while the command line is empty (You can cancel the current command and get an empty line with **Control-C**, so **Control-C** followed by **Control-D** will log you out in most cases).

It is possible to write commands that consist of multiple lines. When the entered line is not a complete command and more input is expected, the console shows a continuation prompt that lists all open parentheses, braces, brackets, and quotes, and also trailing backslash if the previous line ended with **backslash**-white-space.

When you are editing such multiple line entries, the prompt shows the number of current lines and total line count instead of the usual username and system name.

line 2 of 3> :put (\

Sometimes commands ask for additional input from the user. For example, the command '`/password`' asks for old and new passwords. In such cases, the prompt shows the name of the requested value, followed by colon and space.

# Hierarchy

The console allows the configuration of the router's settings using text commands. Since there is a lot of available commands, they are split into groups organized in a way of hierarchical menu levels. The name of a menu level reflects the configuration information accessible in the relevant section.

For example, you can issue the `/ip route print` command:

Instead of typing `/ip route` path before each command, the path can be typed only once to move into this particular branch of the menu hierarchy. Thus, the example above could also be executed like this:

Each word in the path can be separated by **space** (as in the example above) or by "/"

Notice that the prompt changes in order to reflect where you are located in the menu hierarchy at the moment. To move to the top level again, type " / "

To move up one command level, type " .. "

You can also use **/** and **..** to execute commands from other menu levels without changing the current level:

# Item Names and Numbers

Many of the command levels operate with arrays of items: interfaces, routes, users, etc. Such arrays are displayed in similarly-looking lists. All items in the list have an item number followed by flags and parameter values.

To change the properties of an item, you have to use the set command and specify the name or number of the item.

## Item Names

Some lists have items with specific names assigned to each of them. Examples are interface or user levels. There you can use item names instead of item numbers.

You do not have to use the print command before accessing items by their names, which, as opposed to numbers, are not assigned by the console internally, but are properties of the items. Thus, they would not change on their own. However, there are all kinds of obscure situations possible when several users are changing the router's configuration at the same time. Generally, item names are more "stable" than the numbers, and also more informative, so you should prefer them to numbers when writing console scripts.

## Item Numbers

Item numbers are assigned by the print command and are not constant - it is possible that two successive print commands will order items differently. But the results of the last print commands are memorized and, thus, once assigned, item numbers can be used even after add, remove and move operations (since version 3, move operation does not renumber items). Item numbers are assigned on a per session basis, they will remain the same until you quit the console or until the next print command is executed. Also, numbers are assigned separately for every item list, so /`ip address print` will not change the numbering of the interface list.

You can specify multiple items as targets to some commands. Almost everywhere, where you can write the number of items, you can also write a list of numbers.

# General Commands

There are some commands that are common to nearly all menu levels, namely: **print, set, remove, add, find, get, export, enable, disable, comment, move.** These commands have similar behavior throughout different menu levels.

