---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-14
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [2353, 2530]
sha256: c7e74d47401cf983c82dd71f9d8857a45d96ac6552d853ccb0467b003fd5b5be
---

# Copyright (c) 2016, Andrea Dainese

EVE-NG Community Cookbook 
Version 1.11 
Page 91 of 165 © EVE-NG LTD 
6.7.3 Lab preview information 
Description, version, UUID etc. 
 
6.7.4 Lab Global Settings 
Lab Global Setting s Page is opened when you click on the Edit button below the Lab 
preview window or from the Topology page Side bar: 
 
This page allows you to fill out important information about the lab. The red numbers in the 
picture correlate with the numbers listed below 
1. Lab name. 
2. Version: Version numbers allow a lab author to assign a value to a unique state of a 
lab. Increase the number to correspond to new developments in the lab. If left unfilled, 
EVE will assign a value of 1 automatically. 
3. Author: You can add a lab author name in this field 
4. Config Script Timeout: It is the value in seconds used for the “Configuration Export” 
and “Boot from exported configs ” operations. Refer to section 10.3 for more 
information.

EVE-NG Community Cookbook 
Version 1.11 
Page 92 of 165 © EVE-NG LTD 
5. Description: In the Description field you can write a short description of the lab. 
6. Tasks: In the Tasks field you can write the task for your lab. 
The Lab details window can be opened from the Topology Canvas page 
sidebar during labbing, to read the Tasks for the lab.

EVE-NG Community Cookbook 
Version 1.11 
Page 93 of 165 © EVE-NG LTD 
7 EVE WEB Topology page 
Once you open a lab, the topology page for that lab will open. 
 
7.1 Side bar functions 
Move your mouse pointer over to the left on top of the minimized sidebar to expand the 
interactive sidebar as shown in below screenshot

EVE-NG Community Cookbook 
Version 1.11 
Page 94 of 165 © EVE-NG LTD 
7.1.1 Add an object 
The “Add an object” menu can be accessed in two different ways, from the sidebar and by right-
clicking on the Topology Page 
  
7.1.1.1 Node object 
The Node object opens the “Add a new node ” window. Only nodes that appear blue in the 
dropdown menu can be added. A grey image name signifies that you have not yet properly 
uploaded an image to the proper folder. A blue image name means that at least one i mage 
exists in the proper folder for this template.  
 
7.1.1.2 Network object 
The Network object opens the “Add a new network” window. This function is used to add any 
kind of network (Cloud, Bridge). For details on these, please refer to section 9

EVE-NG Community Cookbook 
Version 1.11 
Page 95 of 165 © EVE-NG LTD 
 
7.1.1.3 Picture object 
The picture object opens the “Add Picture” window and allows you to upload custom topologies 
in jpg or png format. After uploading, you can edit these pictures and map selected areas to 
nodes from the topology to use your own designs as a lab topology from which you can directly 
connect to the nodes. For details, refer to section 10.2 
 
7.1.1.4 Custom shape object 
The Custom shape object allows you to add shape elements onto the topology; these currently 
include squares, round squares and circles. For details, refer to section 10.1

EVE-NG Community Cookbook 
Version 1.11 
Page 96 of 165 © EVE-NG LTD 
7.1.1.5 Text object 
The Text object allows you to add Text elements onto the topology. For details, refer to section 
10.1.3 
 
7.1.2 Nodes 
The Nodes object in the sidebar opens the “Configured Nodes” window.  
 
In this window,  you can m ake changes for nodes that are on the lab topology. More options 
can be found in the detailed node specific menu, for details refer to section 8.1.2. 
 NOTE: Running nodes are highlighted in Blue, their settings c annot be changed. You 
can only change settings of nodes that are not currently running. 
You can change the following values: 
• Node Name 
• Boot image 
• Number of CPUs for the node 
• Enable or disable CPU Limit (Refer to section 6.4.1) 
• IDLE PC for Dynamips node 
• NVRAM in Kbyte 
• RAM in Mbyte 
• Ethernet quantity. NOTE: The Node must be disconnected from any other nodes to 
make this change. You cannot change the interface quantit y if the node is connected 
to any other node. 
• Serial interface quantity, IOL nodes only. You cannot change Serial interface quantity 
if the node is connected to any other node. 
• Type of Console 
• Node Icon that appears on the Topology 
• Startup configuration to boot from

EVE-NG Community Cookbook 
Version 1.11 
Page 97 of 165 © EVE-NG LTD 
Actions Buttons (Stopped node): 
 
• Start node 
• Stop node 
• Wipe node 
• Export the nodes config 
• Networks 
• Edit node 
• Delete Node 
 
Actions Buttons (Running node): 
 
• Console to the node 
• Stop node 
• Wipe node 
• Export the nodes config 
• Edit node 
• Delete Node 
 
7.1.3 Networks 
 The Networks object in the sidebar will open the “Configured 
Networks” window. 
The “Configured Networks” window will only show networks that were specifically added to the 
topology; it will not show node interconnections. The example below is showing information for 
networks on the Topology. For Cloud networks  and how to connect EVE labs to a network 
external to EVE, please refer to section 9

EVE-NG Community Cookbook 
Version 1.11 
Page 98 of 165 © EVE-NG LTD 
 
 
• Edit Network 
• Delete Network 
7.1.4 Startup-configs 
The Startup-configs object in the sidebar opens the “Startup-configs” 
window. 
This window will show you startup -config for each node and if the node is set to boot from it 
(ON) or not (OFF). 
 
7.1.5 Logical Maps 
NOTE: The Logical Maps object will only appear in the sidebar after 
you have uploaded a custom topology picture to the lab EV E lab 
(Please refer to section 7.1.1.3). The  Pictures object in the sidebar opens the “Picture 
Management” window.  
For details on the Picture / custom topology feature, refer to section 10.2

EVE-NG Community Cookbook 
Version 1.11 
Page 99 of 165 © EVE-NG LTD 
7.1.6 Configured Objects 
The “Configured Objects” window will display a list of all objects that 
are added onto the topology. For details on different objects, refer to 
section 10.1 
NOTE: You will not see any objects in this window if none have been added to the lab yet. 
 
7.1.7 More actions 
The More actions menu in the sidebar has a submenu with the following functions.  
 
7.1.7.1 Start all nodes 
The “Start all nodes” action will start all nodes on your topology, taking 
the (configurable) startup delay of each node into consideration.  
 IMPORTANT. Starting many nodes at once can seriously spike your CPU utilization. 
Please make sure that you are not using the “Start all nodes” option for heavy labs or 
that you have configured a proper de lay between the nodes. For heavy nodes and 
large quantities, it is recommended to start them in smaller groups, wait for them to 
finish booting and then start another small group of nodes. 
7.1.7.2 Stop all nodes 
  Stopping all nodes will power off all nodes on your topology. 
 NOTE: It is recommended to save your (running) config urations on the nodes in your 
lab before you stop the lab if you want to continue where you left off the next time. 
Stopping the nodes will leave the images in a temporary folder and will ta ke up space 
on your drive until they have been wiped.

