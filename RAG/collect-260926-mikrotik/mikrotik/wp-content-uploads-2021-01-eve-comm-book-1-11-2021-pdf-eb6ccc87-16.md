---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-16
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [2733, 2890]
sha256: c8fbe2e8137bd01b52dc557df3fb7d7662462b19eb6519800718c8d883d64d11
---

# Copyright (c) 2016, Andrea Dainese

EVE-NG Community Cookbook 
Version 1.11 
Page 107 of 165 © EVE-NG LTD 
Wipe Selected: The Wipe Selected nodes action will wipe the NVRAM or currently saved 
image of the selected nodes in the current lab. 
Example: You have saved the nodes configuration by sav ing the running configuration to the 
startup configuration. The Wipe comm and will delete the saved NVRAM startup configuration 
and on the next boot it will boot from factory defaults.  
The same applies to images without configurations, e.g. a linux node. If you make modifications 
to the system and afterwards wipe this node, the  next time it will boot from the original base 
image again as the modified image was deleted. 
The Wipe node action is commonly used with initial startup configuration modifications. T he 
Wipe node action does not delete configured startup configurations or sets. Please refer to 
section 10.3 
Console To Selected Nodes: Console To Selected Nodes will open a console to all selected 
running nodes in the current l ab. This includes all different kinds of configured console types 
for lab nodes like VNC, Telnet and RDP  
Export all CFG s: The Export all configurations action wil l export current configs of selected 
nodes to the EVE startup-configs. 
For a full explanation of exporting configurations, please refer to section 10.3 
Set nodes s tartup-cfg to default configset:  Sets nodes to Default startup config, used 
commonly with the wipe nodes function. NOTE: If you have nothing saved in the default config 
set for any node, that node will boot from factory default instead. This is commonly us ed with 
the wipe nodes function so the node will boot from the configured startup -config on next boot 
and not from the startup-config in its NVRAM in case the node was started before already. 
Please refer to section 10.3 
Set nodes startup-cfg to none. Setting selected lab nodes to boot from factory default. Used 
commonly with the wipe nodes function. The example below shows the ste ps to set selected 
nodes to boot from factory default. 
Step 1: Wipe selected nodes 
Step 2: Set nodes startup-cfg to none 
 
Please refer to section 10.3 
Horizontal Align. Aligns the selected nodes in one horizontal line.  
Step 1: Select the nodes you wish to align. 
Step 2: Right click on one of the selected nodes and choose Horizontal align, this will align all 
nodes to the selected node. 
Picture before:

EVE-NG Community Cookbook 
Version 1.11 
Page 108 of 165 © EVE-NG LTD 
 
Picture after: 
 
Vertical Align: Aligns the nodes in one vertical line. 
Step 1: Select the nodes you wish to align. 
Step 2: Right click on one of the selected nodes and choose Vertical align , this will align all 
nodes to the selected node. 
Picture before  Picture after 
  
Circular Align: Aligns the nodes in a circle.  
Step 1: Select the nodes you wish to align. 
Step 2: Right click on one of the selected nodes and choose Circular Align, this will align all 
nodes in a circle, the midpoint of the circle will be at the coordinates the selected node was at 
before. 
Picture Before    Picture After

EVE-NG Community Cookbook 
Version 1.11 
Page 109 of 165 © EVE-NG LTD 
  
Delete nodes startup-config.  
 WARNING, this action will delete the configurations of the selected nodes that are 
saved to your Default config set. Please make sure that is what you want to do before 
you execute this. 
Delete selected: This will delete the selected nodes from your current lab. 
Selected nodes can be moved as a group across the topology. 
Example: You can select nodes and objects to better position them on the Topology.  
 
 
7.3 EVE Lab node states and symbols 
7.3.1 Stopped (non-running) nodes 
Grey colour and a square symbol below a node means that the node is stopped 
and not running. Once you will start it, the node will change to one of the running 
states below.

EVE-NG Community Cookbook 
Version 1.11 
Page 110 of 165 © EVE-NG LTD 
A grey node with an exclamation mark inside a triangle below the node means 
that there was a problem during the boot process, this could be a corrupted boot 
image, insufficient resources or problems with the initial configuration. A node in 
this state cannot be started again. 
Workaround: Right-click on the node and wipe it, the symbol will then change to a grey colour 
with a square symbol below it. Then edit the node and make sure you have configured sufficient 
resources and the correct settings for this node, if it has startup-configs you can check them as 
well. Afterwards start the node again. 
7.3.2 Running nodes 
The blue colour and black Play triangle symbol means that the node is started and 
running, the node is in a working/functional state. 
A running node with a clock symbol below the node means that the node is waiting 
to finish loading from the set exported/startup configuration. Once the configuration 
has been successfully applied, the node symbol will change to a Play triangle 
symbol. If the node has finished booting but the clock symbol does not change to 
the Play triangle symbol, the problem could be in the uploaded startup configuration. For how 
to use exported configurations and boot nodes from them, please refer to section 10.1 
A running node with a turning red gear symbol means that the node is either in the 
process of hibernating the node or it has sent the shutdown signal to the node and 
is waiting for it to turn off. Once this process has successfully finished, the symbol 
will turn into a grey node with a black square symbol below it (stopped state). 
 NOTE: If the node does not support a system shutdown or does not recognize the 
shutdown signal (example: Cisco router), after clicking on Shutdown, the node can 
stay with a turning red gear symbol below it indefinitely.  
Workaround: Use Stop or Stop/PowerOff to stop the node. 
Example nodes where Stop/Shutdown is supported: Microsoft Windows and most Linux nodes 
as well as a lot of appliances based on linux. 
7.3.3 Node connector symbol 
Connector symbol: If you move your mouse pointer on top of a running 
or stopped node, an orange connector symbol appears. It is used to 
connect nodes on the topology in a drag and drop style. Drag the 
symbol from one node and release the mouse pointer on th e second 
node. A new window will appear where you can select the 
interfaces the link should connect to.

EVE-NG Community Cookbook 
Version 1.11 
Page 111 of 165 © EVE-NG LTD 
7.4 Other 
7.4.1 Notifications area 
 The Notification area in the top right is displaying 
informational or error messages.

EVE-NG Community Cookbook 
Version 1.11 
Page 112 of 165 © EVE-NG LTD 
8 Working with EVE labs 
 IMPORTANT NOTE: You must prepare and upload at least a couple of images to start 
building your labs. Refer to section 12 
8.1 Creating a lab 
Step 1: Click Add new lab. For more information on creating new labs, please refer to section 
6.2.2.1 
 
Step 2:  
Fill out the  lab information . Name and Ver sion are required fields. Next hit  Save. Refer to 
section 6.7.4 for more information about the different fields in the Edit lab window. 
 
8.1.1 Adding nodes to the lab 
The new Topology page will open. There are two different ways to add nodes to the topology 
canvas: 
Step 1: Object/Add Node 
Left Side Bar > Add object > node. Refer to 
section 7.1.1.1 for more information. 
Right click on a free a rea of the t opology 
page and click on “Node” to add a new node. 
Refer to section 7.2.1 for more information.

EVE-NG Community Cookbook 
Version 1.11 
Page 113 of 165 © EVE-NG LTD 
  
 
Step 2: The Add new node window will appear. You can scroll down to choose which node you 
wish to add to the lab topology, or you can type the node name to filter through the node list. 
 NOTE: It will only be possible to select and add nodes that have images preloaded in 
EVE. These nodes will be displayed in  a blue font. To prepare images for EVE, refer  
to section

