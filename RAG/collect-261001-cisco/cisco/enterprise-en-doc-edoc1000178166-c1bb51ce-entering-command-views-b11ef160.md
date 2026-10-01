---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1000178166-c1bb51ce-entering-command-views-b11ef160
title: "enterprise-en-doc-edoc1000178166-c1bb51ce-entering-command-views-b11ef160"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1000178166-c1bb51ce-entering-command-views-b11ef160.md
source_anchor: ""
source_lines: [1, 37]
sha256: 98c7911c073a0ea862eb94eb8ffda30bc44753df297b89c8535bfd0ba92fca5e
---

# enterprise-en-doc-edoc1000178166-c1bb51ce-entering-command-views-b11ef160

Enterprise
The device has many functions; therefore various configuration commands and query commands are provided to facilitate device management and maintenance. Huawei switch registers commands to different command views based on the functions of the commands so that users can easily use them. To configure a function, enter the corresponding command view and then run corresponding commands.
The device provides various command views. For the methods of entering the command views except the following views, see the S1720, S2700, S5700, and S6720 V200R011C10 Command Reference.
| Name | How To Enter | Function | 
|---|---|---|
| User view | When a user logs in to the device, the user enters the user view and the following prompt is displayed: <HUAWEI> | In the user view, you can view the running status and statistics of the device. | 
| System view | Run the system-view command and press Enter in the user view. The system view is displayed. <HUAWEI> system-view Enter system view, return user view with Ctrl+Z. [HUAWEI] | In the system view, you can set the system parameters of the device, and enter other function views from this view. | 
| Interface view | Run the interface command and specify an interface type and number to enter the interface view. [HUAWEI] interface gigabitethernet X/Y/Z [HUAWEI-GigabitEthernetX/Y/Z]  X/Y/Z indicates the number of an interface that needs to be specified. It is in the format of stack ID/card number/interface sequence number. The interface GigabitEthernet is used as an example. | In the interface view, you can configure interface parameters including physical attributes, link layer protocols, and IP addresses. | 
The command line prompt HUAWEI is the default host name (sysname). The prompt indicates the current view. For example, <> indicates the user view and [] indicates all other views except the user view.
You can enter ! or # followed by a character string in any view. All entered content (including ! and #) is displayed as comments. That is, the corresponding configuration is not generated.
Some commands can be executed in multiple views, but they have different functions after being executed in different views. For example, you can run the lldp enable command in the system view to enable LLDP globally and in the interface view to enable LLDP on an interface.
You can run the quit command to return from the current view to an upper-level view.
[HUAWEI-aaa] quit
[HUAWEI] quit
<HUAWEI>
To return from the AAA view directly to the user view, press Ctrl+Z or run the return command.
[HUAWEI-aaa]           // Enter Ctrl+Z
<HUAWEI> 
[HUAWEI-aaa] return
<HUAWEI> 
Intelligent rollback enables the system to automatically return to the previous view if a command fails to be executed in the current view. The system performs view return attempts until the applicable view of the command is displayed. The system can return to the system view at the maximum extent.
Intelligent rollback cannot be performed in the port group view and VLAN-Range view.
If command matching fails because an ambiguous command is entered in the current view, no intelligent rollback can be performed.
If the intelligent rollback function is enabled, commands may be executed in unexpected views, and services may be interrupted. Before configuring a command, check whether the command to be configured exists in the view. If the command does not exist, run the command in the correct view.
The following provides two application examples for intelligent rollback. The system enters the applicable view of a command after performing one view return attempt in the first example, and performs multiple attempts in the second example.
After entering an OSPF area view, the system allows a user to directly enter another OSPF area view, without the need to manually return to the OSPF view.
<HUAWEI> system-view
[HUAWEI] ospf 100
[HUAWEI-ospf-100] area 1
[HUAWEI-ospf-100-area-0.0.0.1] area 2
[HUAWEI-ospf-100-area-0.0.0.2] 
After entering an OSPF area view, the system allows a user to directly enter an interface view, without the need to manually return to the system view.
<HUAWEI> system-view
[HUAWEI] ospf 100
[HUAWEI-ospf-100] area 1
[HUAWEI-ospf-100-area-0.0.0.1] interface gigabitEthernet 0/0/3
[HUAWEI-GigabitEthernet0/0/3]
