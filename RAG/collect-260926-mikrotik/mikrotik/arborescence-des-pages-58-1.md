---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-58-1
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "memory", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-58.md
source_anchor: ""
source_lines: [1, 127]
sha256: c76f42902b5c44b5721edab994f02f59b0a0903f48e6fea8b8fc68eac937ebba
---

# Overview

This article describes a set of commands used for configuration management.

# Configuration Undo and Redo

Any action done in the GUI or any command executed from the CLI is recorded in `/system history`.  You can undo or redo any action by running undo or redo commands from the CLI or by clicking on Undo, and Redo buttons from the GUI. 

A simple example to demonstrate the addition of the firewall rule and how to undo and redo the action:

We have added a firewall rule and in `/system history` we can see all that was done.

Let's undo everything:

As you can see firewall rule disappeared.

Now redo the last change:

System history is capable of showing exact CLI commands that will be executed during "Undo" or "Redo" actions even if we perform the action from GUI, for example, detailed history output after adding TCP accept rule from WinBox:

# Safe Mode

It is sometimes possible to change router configuration in a way that will make the router inaccessible (except from the local console). Usually, this is done by accident, but there is no way to undo the last change when the connection to the router is already cut. Safe mode can be used to minimize such risk.

The **"Safe Mode"** button in the Winbox GUI allows you to enter Safe Mode, while in the CLI, you can access it by either using the keyboard shortcut **F4** or pressing **[CTRL]+[X]**. To exit without saving the made changes in CLI, hit **[CTRL]+[D].**

Message **Safe Mode taken** is displayed and prompt changes to reflect that session is now in safe mode. All configuration changes that are made (also from other login sessions), while the router is in safe mode, are automatically undone if the safe mode session terminates abnormally. You can see all such changes that will be automatically undone and tagged with an **F** flag in the system history:

Now, if the telnet connection (or WinBox terminal) is cut, then after a while (TCP timeout is **9** minutes) all changes that were made while in safe mode will be undone. Exiting session by **[Ctrl]+[D]** also undoes all safe mode changes, while **/quit** does not.

If another user tries to enter safe mode, he's given the following message:

- [u] - undoes all safe mode changes, and puts the current session in safe mode.
- [r] - keeps all current safe mode changes, and puts the current session in a safe mode. The previous owner of safe mode is notified about this:

- [d] - leaves everything as-is.

If too many changes are made while in safe mode, and there's no room in history to hold them all (currently history keeps up to 100 most recent actions), then the session is automatically put out of the safe mode, and no changes are automatically undone. Thus, it is best to change the configuration in small steps, while in safe mode. Pressing **[Ctrl]**+**[X]** twice is an easy way to empty the safe mode action list.

# System Backup and Restore

System backup is the way to completely clone router configuration in binary format.

More information about Backup and Restore is found here.

# Configuration Export and Import

RouterOS allows exporting and importing parts of the configuration in plain text format. This method can be used to copy bits of configuration between different devices, for example, clone the whole firewall from one router to another.

An export command can be executed from each menu (resulting in configuration export only from this specific menu and all its sub-menus) or from the root menu for complete config export and is available for CLI only.

The Export command does not export system user passwords, installed certificates, SSH keys, Dude, or a User-manager database.

Installed certificates, Dude, and User-manager databases must be manually exported and imported into a new device.

System user passwords and user SSH keys can not be exported.

During config import, we suggest using the same RouterOS version used during config export to prevent cases when some of the commands do not exist in one or another RouterOS version.

## Configuration Export

The following command parameters are accepted:

| Property | Description | 
|---|---|
| **compact** | Outputs only the modified configuration. Starting from v6rc1, "export compact" became the default behavior, so "export" and "export compact" now produce identical output. | 
| **file** | Export configuration to a specified file. When the file is not specified export output will be printed to the terminal | 
| **path** | Parameter allows to include or exclude specific configuration menu from router entire configuration | 
| **show-sensitive** | Show sensitive information, like passwords, keys, etc. By default, sensitive information is hidden. List of menus with sensitive parameters | 
| **terse** | With this parameter, the export command will output the configuration as full commands on separate lines, each including the corresponding menu path. | 
| **verbose** | With this parameter, the export command will output whole configuration parameters and items including defaults. | 

For example, export configuration from `/ip address` the menu and save it to a file:

By default, the export command writes only user-edited configuration, RouterOS defaults are omitted.

For example, the IPSec default policy will not be exported, and if we change one property then only our change will be exported:

Example how to import entire /ip firewall configuration, except /ip firewall nat

Note:

The ***** flag, indicates that the entry is system default and cannot be removed manually.

Here is the list of all menus containing default system entries

| Menu | Default Entry | 
|---|---|
| **/interface wireless security-profiles** | default | 
| **/ppp profile** | "default", "default-encryption" | 
| **/ip hotspot profile**  | default | 
| **/ip hotspot user profile**  | default | 
| **/ip ipsec policy**  | default | 
| **/ip ipsec policy group**  | default | 
| **/ip ipsec proposal** | default | 
| **/ip ipsec mode-conf** | read-only | 
| **/ip smb shares**  | pub | 
| **/ip smb users**  | guest | 
| **/ipv6 nd**  | any | 
| **/mpls interface** | all | 
| **/routing bfd interface** | all | 
| **/routing bgp instance** | default | 
| **/routing ospf instance** | default | 
| **/routing ospf area** | backbone | 
| **/routing ospf-v3 instance** | defailt | 
| **/routing ospf-v3 area** | backbone | 
| **/snmp community** | public | 
| **/tool mac-server mac-winbox** | all | 
| **/tool mac-server** | all | 
| **/system logging** | "info", "error", "warning", "critical" | 
| **/system logging action** | "memory", "disk", "echo", "remote" | 
| **/queue type** | "default", "ethernet-default", "wireless-default", "synchronous-default", "hotspot-default", "only-hardware-queue", "multi-queue-ethernet-default", "default-small" | 

If some specific menu will not be able to respond to the export command, starting from the RouterOS v7.11, an error message will be printed out in the export command output after a timeout ("#error exporting "/xxx" (timeout)") and the process will move on to the next menu.

Starting from RouterOS 7.13, you can export parts of a specific menu. For instance, it is possible to export a specific address-list among multiple address-lists on your router.

[admin@MikroTik] > ip firewall address-list export where list=mylist

## Configuration Import

Root menu command import allows running configuration script from the specified file. Script file (with extension ".rsc") can contain any console command including complex scripts.

For example, load saved configuration file

Import command allows to specify the following parameters:

