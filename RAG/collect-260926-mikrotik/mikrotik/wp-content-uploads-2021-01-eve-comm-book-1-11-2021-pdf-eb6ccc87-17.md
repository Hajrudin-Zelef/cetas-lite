---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-17
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "license", "licenses"]
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [2891, 3116]
sha256: b30bdaf9508cab5f28f116ab3d52de1274c0c9cd35218f9d55cb1c5b172545b9
---

# Copyright (c) 2016, Andrea Dainese

EVE-NG Community Cookbook 
Version 1.11 
Page 114 of 165 © EVE-NG LTD 
Step 3: Edit “Add a new node” settings. Please refer to the picture and table below.

EVE-NG Community Cookbook 
Version 1.11 
Page 115 of 165 © EVE-NG LTD 
8.1.1.1 Node values Table 
Number Description 
1.  
Template menu. Choose which node template to add 
to the topology 
2.   Chose the number of nodes of this type you want to add 
to the topology 
3.  
Choose your preferred version from preloaded images 
list (if you have more than one image loaded fo r a 
single template). 
4.  
Type your preferred 
node name. If you are 
adding more than one, EVE will automatically append numbers to the nodes 
name.  
Example. We are adding 5 CSR nodes with the name R. On the topology they 
will appear as R1, R2, R3, R4, R5. Later using the Nodes window, you can edit 
the node names per your needs. Refer to section 7.1.2 or edit the node 
individually, refer to section 8.1.2. 
5.  
Node icons can be changed from the 
default per your preference, simply choose 
the preferred icon from the dropdown list. 
Node icons can be changed later per your 
needs. Refer to section 7.1.2 
6.  
The UUID number is 
assigned automatically after 
a node is created. You may 
also set it manually in case you are using a license that is tied to a particular 
UUID.

EVE-NG Community Cookbook 
Version 1.11 
Page 116 of 165 © EVE-NG LTD 
7.  
CPU limit per node. This option is already set 
(checked/unchecked) per EVE recommendations. Refer to 
section 6.4.1 
8.  
Each node template has a pre-set CPU value that aligns 
with vendor requirements. This value can be changed per 
your needs. 
9.  
Each node template has a pre-set RAM value that aligns 
with vendor requirements. This value is displayed in MB 
and may be changed per your needs. 
10.  
The number of ethernets interfaces. 
 NOTE for IOL nodes: 
Ethernet interfaces for IOL nodes are placed into groups of 4. A value of 1 for 
Ethernet means your node will have 4 interfaces. 
The serial interface option is available for IOL nodes only and follows the same 
grouping structure as ethernet interfaces. A value of 1 for Serial means your 
node will have 4 serial 
interfaces. 
 
11.  
Custom MAC address for Qemu nodes only. You can define your own MAC 
address for first interface: 
 
12.  
EVE will pre-set the best recommended QEMU version 
for each node template. This value can be changed per 
your needs. 
13.  
Qemu architecture is pre-set per image vendor 
recommendations. This value can be changed per your 
needs

EVE-NG Community Cookbook 
Version 1.11 
Page 117 of 165 © EVE-NG LTD 
14.  
Type of Qemu NIC is pre-set per image vendor 
recommendations. This value can be changed per your 
needs. 
15.  
Qemu custom options are pre-set 
per image vendor 
recommendations. This value can be changed per your needs 
16.  
Startup configuration: Value can 
be changed to set your node  to 
boot from saved configurations. Refer to section 10.3 for more details. 
17.  
The Delay value is set in seconds and can 
be used to delay a node from booting after 
it is started. Example: if the value is set to 30, the node will wait 30 seconds 
before processing its boot sequense. This feature is useful in conjunction with 
the “Start all nodes” function if your lab requires certain nodes to start up before 
others or to avoid a mass-start of very heavy nodes. 
18.  
Console types for each 
template are pre-set with 
recommended settings.   
The setting can be changes per your needs.  
 NOTE:  The Docker template contains a wide variety of images, 
therefore, please refer to section 14.1.3 for recommended console 
types for each docker image. Windows nodes can use either RDP or 
VNC but RDP needs to be enabled in Windows itself. 
19.  
OPTIONAL: Templates for 
Cisco FirePower, F5, Linux, 
and Citrix have the option to 
manually set the MAC address for the first ethernet interface. This will enable 
the use of licenses that are tied to a particular MAC address. 
MAC Address format must be like: 00:50:0a:00:0b:00 
8.1.2 Edit node 
EVE provides two ways to edit nodes after being added to the topology canvas.

EVE-NG Community Cookbook 
Version 1.11 
Page 118 of 165 © EVE-NG LTD 
 NOTE: A node must be wiped each time an image or startup configuration has been 
changed. 
8.1.2.1 Edit nodes globally 
From the Topology page. Click “Nodes” from the left sidebar to bring up the nodes list. Refer to 
section 7.1.2 for more details. 
 
 
8.1.2.2 Edit node individually. 
Right click on the node and click Edit 
 
 
The “Edit node” window will appear. It is very similar to the window that is displayed when you 
add a new node . To change values for the node, refer to the nodes value table in section 
8.1.1.1.

EVE-NG Community Cookbook 
Version 1.11 
Page 119 of 165 © EVE-NG LTD 
 
 
8.1.3 Wipe Node 
The “Wipe node” function will clear the NVRAM of the node. Each time 
a node setting  is changed ( CPU, RAM, boot image or startup 
configuration) a wipe must be issued on that node. For more information 
refer to section 10.3

EVE-NG Community Cookbook 
Version 1.11 
Page 120 of 165 © EVE-NG LTD 
8.1.4 Interconnecting nodes 
To connect nodes on the lab, use the drag and drop style method 
Connector symbol: Moving the mouse over a node will make an orange male plug 
appear. The male plug is used to connect nodes on the topology , drag and drop 
style. Release the mouse pointer on the second node. 
 
 
 
 
The connection window will appear. Choose the interface you want to use to  interconnect the 
nodes. Click Save when finished. 
 
 
 
 
 
 
8.1.5 Delete connection between nodes 
To delete a connection, right click on it and hit “Delete.”

EVE-NG Community Cookbook 
Version 1.11 
Page 121 of 165 © EVE-NG LTD 
8.1.6 Delete Node 
To delete a node, right click it and hit “Delete.” This is a non -reversable 
function 
NOTE: It is strongly recommended to delete connections from a node 
before deleting the node itself.  
 
8.2 Running labs 
8.2.1 Starting lab 
Nodes inside a lab may be started individually, in groups, or all at once. 
The Start all nodes option will start all nodes on your topology.  
 IMPORTANT. Starting all the nodes at once can  result in major  spikes in CPU 
utilization. Please make sure you are not using the “Start all nodes” option for heavy 
labs. Instead, it is recommended to start nodes in small groups. 
Starting a node or group of nodes: 
Right click on single node or node group and hit “Start.” 
 
 
 
 
Running nodes will turn blue. Refer to section 7.3 for node states 
 
8.3 Saving labs 
To save a running lab, refer to the vendor recommended save commands for each node.   
Example:  
Cisco: “copy run start” 
Juniper “commit”

EVE-NG Community Cookbook 
Version 1.11 
Page 122 of 165 © EVE-NG LTD 
 
Your current work will be saved in the nodes ’ NVRAM and the lab can be stopped safely . 
Starting the lab again will allow you to pick up from where you left off. 
 WARNING: Using the wipe action on a node will clear its  NVRAM. This is similar to 
doing a factory reset on a device. 
The configurations of nodes can be  exported and used as initial or startup configurations for 
your labs. To export configurations and configuration sets for labs refer to section 10.1 
8.4 Stopping labs 
  The Stop all nodes option will stop all nodes on your topology. 
NOTE: It is recommended to save your running configurations before you stop your nodes. 
Stopping a node or group of nodes: 
Right click on single node or node group and hit “Stop.” 
For individual node Stop options refer to section 7.2.5 
8.5 Start saved lab 
Select the lab you want to start and click “Open”. To start Lab refer section 8.2.1 
 
8.6 Importing labs 
Refer to section 6.2.2.6 
8.7 Exporting labs 
Refer to section 6.2.2.5

