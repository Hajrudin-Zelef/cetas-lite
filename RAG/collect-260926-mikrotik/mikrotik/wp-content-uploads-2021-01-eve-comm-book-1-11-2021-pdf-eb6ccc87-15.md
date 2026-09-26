---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-15
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [2531, 2732]
sha256: cb8173e71af0265536f49b0ec119ee69a7cc8f2f8b6951a3226f77b5a21f7f1d
---

# Copyright (c) 2016, Andrea Dainese

EVE-NG Community Cookbook 
Version 1.11 
Page 100 of 165 © EVE-NG LTD 
7.1.7.3 Wipe all nodes 
 The “Wipe all nodes” action will wipe the NVRAM or currently saved 
image of all your nodes in the current lab. 
Example: You have saved the nodes configuration by saving the running conf iguration to the 
startup configuration. The Wipe command will delete the saved NVRAM startup configuration 
and on the next boot it will boot from factory defaults.  
The same applies to images without configurations, e.g. a linux node. If you make modifications 
to the system and afterwards wipe this node, the next time it will boot from the original base 
image again as the modified image was deleted. 
The “Wipe node” action is commonly used with initial  startup configuration modifications. The 
Wipe node actio n does not delete configured startup configurations or sets. Please refer to 
section 10.3 
7.1.7.4 Console to All Nodes 
“Console to all nodes ” will open a console to all of your running 
nodes in the current lab. This includes all diffe rent kinds of 
configured console types for lab nodes like VNC, Telnet and RDP. 
7.1.7.5 Export all CFGs 
The “Export all configurations” action will export current configs to 
the EVE startup-configs. 
Export configurations are supported for: 
Cisco Dynamips all nodes 
Cisco IOL (IOS on Linux) 
Cisco ASA 
Cisco ASAv 
Cisco CSR1000v 
Cisco Nexus 9K 
Cisco Nexus Titanium 
Cisco vIOS L3 
Cisco vIOS L2 
Cisco XRv 
Cisco XRv9K 
Juniper VRR 
Juniper VMX 
Juniper vMX-NG 
Juniper vQFX 
Juniper vSRX 
Juniper vSRX-NG 
Mikrotik 
PFsense FW 
Timos Alcatel 
vEOS Arista 
For a full explanation of exporting configurations, please refer to section 10.3 
7.1.7.6 Edit lab 
Opens the Edit lab window. Refer to section: 6.7.4

EVE-NG Community Cookbook 
Version 1.11 
Page 101 of 165 © EVE-NG LTD 
 
7.1.7.7 Set node’s startup-cfg to default configset 
Sets nodes to the d efault startup -config. NOTE: If you have 
nothing saved in the default config set for any node, that node 
will boot from factory default instead. This is commonly used with the wipe nodes function so 
the node will boot from the  configured startup-config on next boot and not from the startup -
config in its NVRAM in case the node was started before already. 
Please refer to section 10.3 
7.1.7.8 Set node’s startup-cfg to none 
Setting all lab nodes to boot from f actory default. Used 
commonly with the wipe nodes function. The example below 
shows the steps to set a lab to boot from factory default. 
Step 1: Wipe all nodes 
Step 2: Set all nodes to startup-cfg none 
 
Please refer to section 10.3 
7.1.7.9 Delete default startup-cfgs 
 
 WARNING: this action will delete all configurations saved to your saved default config 
set. Please make sure that is what you want to do before you execute this.  
7.1.8 Refresh Topology 
Sometimes it is necessary to refresh the topology if many objects are 
added on the topology. 
7.1.9 Lab page zoom/unzoom 
This action is used to zoom or unzoom a large topology in EVE.

EVE-NG Community Cookbook 
Version 1.11 
Page 102 of 165 © EVE-NG LTD 
 
7.1.10 Status 
Opens the EVE Status window.  
Especially useful while working with labs to monitor your EVE ’s resource utiliz ation. It shows 
EVEs CPU, RAM and disk utilization in real time. You can also see the number of running 
nodes per node type. For details on UKSM and CPU Limit, please refer to section 6.4.1 
 
7.1.11 Lab details 
Lab details display information about a lab, its UUID, description and 
lab tasks. To edit the lab description and lab tasks, please refer to 
section 6.7.4 and 7.1.7.6

EVE-NG Community Cookbook 
Version 1.11 
Page 103 of 165 © EVE-NG LTD 
7.1.12 Lock Lab with password 
“Lock Lab” disables some of the functions on the lab topology. If the lab is locked, you cannot 
move any node or object nor edit any node settings. Basically, the whole lab will be in read -
only mode except for the lab settings itself , which you can st ill edit as Administrator from the 
main menu.  
 
Lab is unlocked and all ope rations are 
working 
 
Enter and confirm your lab lock password 
To unlock a Lab, simply press on the red “Unlock Lab” button with an Administrator account. 
 
Lab is locked and all operations are 
restricted 
 
Enter lab unlock password to unlock lab. 
Warning: Please remember your Lab lock password. In case of a lost password, you will not be 
able to recover it. Unlocking a lab / removal of password can be done by EVE-NG support only. 
7.1.13 Dark mode or Light mode 
 
Sets your lab background to the dark mode 
 
Sets your lab background to light mode 
7.1.14 Close lab 
Closes the lab topology. The lab can be closed while the nodes in 
the lab nodes are stopped. 
7.1.15 Logout 
Log out from the EVE WEB GUI session.

EVE-NG Community Cookbook 
Version 1.11 
Page 104 of 165 © EVE-NG LTD 
 
7.2 EVE Lab topology menus 
Right-clicking within the EVE topology can open new menus with various functions and options 
for managing nodes. 
7.2.1 Lab topology menu 
Right-clicking on the (free/unused) canvas of the EVE topology opens a new 
menu. (Add-) Node, Network, Picture, Custom Shape  and Text are the same 
functions referred to in section 7.1.1. 
Auto Align.  This function will help align object s on the topology. The l ab 
creator does not need to worry about small  displacements of object s. Auto  
Align will align all object s to a virtual grid  with a single click and can make 
neatly arranged labs look even neater. 
7.2.2 Connection menu 
Right-clicking on the connection between nodes allows you 
to delete this connection. 
 
7.2.3 Cloud or Bridge network menu 
Right-clicking on a Cloud or Bridge network allows you to edit or delete it. 
  
If you have chosen Edit, the Network edit window will 
open a window where you can change the placement, 
network type or name/prefix. 
 
For details on  how to operate EVE Cloud networks and 
external connections, please refer to section 9

EVE-NG Community Cookbook 
Version 1.11 
Page 105 of 165 © EVE-NG LTD 
 
 
7.2.4 Stopped node menu 
Right-clicking on a stopped node also opens a menu: 
 
Start node: This will start the selected node in this lab 
Wipe node: Wiping a node will erase the NVRAM 
(running config) or the temporary image snapshot 
depending on the type of node. This option is used to 
clean up a node in order to boot it from factory 
defaults or a custom set of configurations. 
Edit node: Opens the Edit node window (picture on 
the right). For details please refer to section 8.1.2 
Delete node. Deletes the node from the lab. It is 
recommended to disconnect (delete connections to it) 
the node before you delete it. 
 
 
 
 
7.2.5 Running node menu 
Right-clicking on a running node also opens a menu:

EVE-NG Community Cookbook 
Version 1.11 
Page 106 of 165 © EVE-NG LTD 
Wipe node: Wiping a node will erase the NVRAM (running config) or the temporary image 
snapshot depending on the type of node. This option is used to clean up a node in order to 
boot it from factory defaults or a custom set of configurations. 
Export CFG: This function is used to export the saved running configuration to the EVE 
startup configuration sets. Reference section 10.3 
Capture. Wireshark capture. Select the interface which you wish 
to capture. Reference section 5.1.2 
 
 
 
7.2.6 Selected nodes menu and features 
It is possible to select many objects or nodes at once in EVE. Using your mouse, you can select 
an area which will cover your nodes and/or you can click on nodes while holding the CTRL key 
on your keyboard. 
 
A right-click on any of the selected nodes opens a group menu: 
 
Start Selected: This will start the selected nodes in this lab. 
Stop Selected: This will stop the selected nodes in this lab

