---
id: collect-261001-cisco/cisco/multiple-privilege-levels-ciscozine
title: "multiple-privilege-levels-ciscozine"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/multiple-privilege-levels-ciscozine.md
source_anchor: ""
source_lines: [1, 85]
sha256: d896318b609c550cadb8ec575fff829b44ab2e8daa358e940f652aeb68667a92
---

# multiple-privilege-levels-ciscozine

Cisco IOS permits to define multiple privilege levels for different accounts. This could be useful when many people work on the same router / switch, but with different roles (operator, tecnhician, network manager) and there is no time to implement an authentication server.

There are 16 different privilege levels that can be used. Level 0 is **user mode**. Level 15 is the **privileged mode**. Level 1 through 14 are available for **customization** and use.

The command used are:

Ciscozine(config)#privilege *mode* level *level command*
Ciscozine(config)#enable secret level *level password*

The privilege command is used to add authorized IOS commands to each customized levels. The enable secret command defines the secret password needed to access this particular privilege level. The options for these commands are:

- *mode* : it selects an exec or configuration option that will be included with this privilege level
- *level* : it defines the privilege level (number between 1 and 14)
- *command* : it is a specific IOS command at the specified*mode* that is included in this privilege level
- *password* : it is the password associated to the*level*

An example of privilege level configuration:

Ciscozine(config)#privilege interface level 2 ip address
Ciscozine(config)#privilege configure level 2 interface
Ciscozine(config)#privilege exec level 2 configure terminal
Ciscozine(config)#privilege exec level 2 show interfaces
Ciscozine(config)#privilege exec level 2 show running-config
Ciscozine(config)#enable secret level 2 c$sc0Z1Ne


This configuration permit:

- Access configuration mode
- Access the interfaces
- Configure ONLY the IP addresses on the interfaces
Ciscozine(config)#interface fastEthernet 0/0 Ciscozine(config-if)#? Interface configuration commands: default Set a command to its defaults exit Exit from interface configuration mode help Description of the interactive help system ip Interface Internet Protocol config commands no Negate a command or set its defaults Ciscozine(config-if)#ip ? Interface IP configuration subcommands: address Set the IP address of an interface Ciscozine(config-if)#
- Display the interfaces
- Displat the running configuration. Even though the level 2 user can execute the show running-config command, only the configure commands that are permitted are actually displayed. In this example the user would see only IP address information from interfaces in the running configuration.
Ciscozine#sh running-config Building configuration... Current configuration : 141 bytes ! boot-start-marker boot-end-marker ! ! ! ! ! ! interface FastEthernet0/0 no ip address ! interface FastEthernet0/1 no ip address ! ! end Ciscozine#


To access a particular privilege level, use the command:

Ciscozine>enable *level*
Password:
Ciscozine#

Where *level* in this case is ‘2’.

Remember that is **not possible** to access a customized privilege level that does **not have a password**.


More info on http://www.cisco.com/…/configuration/guide/scfpass.html

this is a nice article……thanx man

hello, i’m working on helpdesk level 1 & privilege 1, can you give me the supplement command on a router or switch in CLI thanks

Thank you Fab. Was helpful.

i cant get into the interfaces with these commands.plz help me out.

these are my commands

privilege configure level 10 interface

privilege exec level 10 configure

privilege exec level 10 configure terminal

privilege exec level 10 show

privilege exec level 10 show interfaces

privilege exec level 10 show running-config

privilege interface level 10 ip

privilege interface level 10 ip address

privilege interface level 10 ipv6

privilege interface level 10 ipv6 address

Hi bala,

How to config privilege configure level 2 (set vlan name) ?

This could be accomplished on IOS, but there is no equivalent way to do this on CATOS.
