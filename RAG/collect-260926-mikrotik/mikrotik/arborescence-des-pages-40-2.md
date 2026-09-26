---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-40-2
title: "Login Options"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-40.md
source_anchor: ""
source_lines: [94, 179]
sha256: da1c3892c757c11c7eb0cc83862ecbf188e45edf0302d67c8d6d8da98af390a5
---

# Login Options

| Property | Description | 
|---|---|
| **add** | This command usually has all the same arguments as a set, except the item number argument. It adds a new item with the values you have specified, usually at the end of the item list, in places where the order of items is relevant. There are some required properties that you have to supply, such as the interface for a new address, while other properties are set to defaults unless you explicitly specify them. Common Parameters  Return Values  | 
| **edit** | This command is associated with the set command. It can be used to edit values of properties that contain a large amount of text, such as scripts, but it works with all editable properties. Depending on the capabilities of the terminal, either a full-screen editor or a single line editor is launched to edit the value of the specified property. | 
| **find** | The find command has the same arguments as a set, plus the flag arguments like disabled or active that take values yes or no depending on the value of the respective flag. To see all flags and their names, look at the top of the print command's output. The find command returns internal numbers of all items that have the same values of arguments as specified. | 
| **move** | Changes the order of items in the list. Parameters:  | 
| **print** | Shows all information that's accessible from a particular command level. Thus, `/system clock print` shows the system date and time,`/ip route print` shows all routes etc. If there\'s a list of items in the current level and they are not read-only, i.e. you can change/remove them (example of read-only item list is`/system history` , which shows a history of executed actions), then print command also assigns numbers that are used by all commands that operate with items in this list.Common Parameters:  | 
| **remove** | Removes specified item(-s) from a list. | 
| **set** | Allows you to change values of general parameters or item parameters. The set command has arguments with names corresponding to values you can change. Use ? or double Tab to see a list of all arguments. If there is a list of items in this command level, then the set has one action argument that accepts the number of items (or list of numbers) you wish to set up. This command does not return anything. | 

# Input Modes

It is possible to switch between several input modes:

- **Normal mode** - indicated by normal command prompt.
- **Safe mode** - safe mode is indicated by the word SAFE after the command prompt. In this mode, the configuration is saved to disk only after the safe mode is turned off. Safe mode can be turned on/off with**Ctrl+X or F4.** Read more >>
- **Hot-lock mode** - indicated by additional yellow >. Hot-lock mode autocompletes commands and can be turned on/off with**F7**

# Quick Typing

There are two features in the console that help entering commands much quicker and easier - the [**Tab**] key completions, and abbreviations of command names. Completions work similarly to the bash shell in UNIX. If you press the [**Tab**] key after a part of a word, the console tries to find the command within the current context that begins with this word. If there is only one match, it is automatically appended, followed by a space:

*/inte***[Tab]_** becomes **/interface _**

If there is more than one match, but they all have a common beginning, which is longer than that what you have typed, then the word is completed to this common part, and no space is appended:

*/interface set e***[Tab]_** becomes **/interface set ether_**

If you've typed just the common part, pressing the tab key once has no effect. However, pressing it for the second time shows all possible completions in compact form:

[admin@MikroTik] > interface set e[Tab]_
[admin@MikroTik] > interface set ether[Tab]_
[admin@MikroTik] > interface set ether[Tab]_
ether1 ether5
[admin@MikroTik] > interface set ether_

The **[Tab]** key can be used almost in any context where the console might have a clue about possible values - command names, argument names, arguments that have only several possible values (like names of items in some lists or name of the protocol in firewall and NAT rules). You cannot complete numbers, IP addresses, and similar values.

Another way to press fewer keys while typing is to abbreviate command and argument names. You can type only the beginning of the command name, and, if it is not ambiguous, the console will accept it as a full name. So typing:

[admin@MikroTik] > pi 10.1 c 3 si 100

equals to:

[admin@MikroTik] > ping 10.0.0.1 count 3 size 100

It is possible to complete not only the beginning, but also any distinctive sub-string of a name: if there is no exact match, the console starts looking for words that have string being completed as first letters of a multiple word name, or that simply contain letters of this string in the same order. If a single such word is found, it is completed at the cursor position. For example:

[admin@MikroTik] > interface x[TAB]_
[admin@MikroTik] > interface export _
[admin@MikroTik] > interface mt[TAB]_
[admin@MikroTik] > interface monitor-traffic _

# Console Search

Console search allows performing keyword search through the list of RouterOS menus and the history. The search prompt is accessible with the **[Ctrl+r]** shortcut. 

# Internal Chat System

RouterOS console has a built-in internal chat system. This allows remotely located admins to talk to each other directly in RouterOS CLI. To start the conversation prefix the intended message with the # symbol, anyone who is logged in at the time of sending the message will see it.

# List of Keys

| Key | Description | 
|---|---|
| Control-C | keyboard interrupt | 
| Control-D | log out (if an input line is empty) | 
| Control-K | clear from the cursor to the end of the line | 
| Control-U | clear from the cursor to the beginning of the line | 
| Control-X or F4 | toggle safe mode | 
| F7 | toggle hot-lock mode | 
| Control-R or F3 | toggle console search | 
| F6 | toggle cellar | 
| F1 | show context-sensitive help. | 
| Tab | perform line completion. When pressed a second time, show possible completions. | 
| # | Send a message to an internal chat system | 
| Delete | remove character at the cursor | 
| Control-H or Backspace | removes character before cursor and moves the cursor back one position. | 
| Control-\ | split line at cursor. Insert newline at the cursor position. Display second of the two resulting lines. | 
| Control-B or Left | move cursor backward one character | 
| Control-F or Right | move cursor forward one character | 
| Control-P or Up | go to the previous line. If this is the first line of input then recall previous input from history. | 
| Control-N or Down | go to the next line. If this is the last line of input then recall the next input from the history | 
| Control-A or Home | move the cursor to the beginning of the line. If the cursor is already at the beginning of the line, then go to the beginning of the first line of the current input | 
| Control-E or End | move the cursor to the end of the line. If the cursor is already at the end of the line, then move it to the end of the last line of the current input | 
| Control-L or F5 | reset terminal and repaint screen |
