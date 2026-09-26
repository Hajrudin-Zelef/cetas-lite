---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-19
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["ethernet", "intel"]
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [3367, 3565]
sha256: 8ca54c53c2c174c5bbddb3c4c7c8f16abe56de15996a3c8cadf93735d127d4b1
---

# Copyright (c) 2016, Andrea Dainese

EVE and WSA VMs settings 
EVE VM, the third port (Network adapter 3) is 
assigned to VMnet2. This is Cloud2 inside your 
EVE labs.  
 
Cisco Web Security Appliance (WSA), 
Management port is assigned to VMnet2 
 
EVE Lab connected to the WSA (Cloud2) 
 NOTE: ESXi WSA VM obtained the IP 192.168.10.3 from the DHCP pool on the lab 
switch. The gateway is 192.168.10.1 
 NOTE: The Firefox Docker node user for management obtained the IP 192.168.10.2 
from the DHCP pool configured on the lab switch.

EVE-NG Community Cookbook 
Version 1.11 
Page 133 of 165 © EVE-NG LTD 
 
9.5 Connecting EVE Lab to a physical device 
 IMPORTANT NOTE : To bypass MAC addressing over pnet/cloud interface please 
SSH to your EVE and type: 
for i in /sys/class/net/pnet*/bridge/group_fwd_mask  ; do echo 8 > $i ; done 
 
9.5.1 ESXi EVE 
To connect a physical device (e.g. router, switch) to an EVE lab over a cloud interface, we have 
to bridge the ESXi NICs ethernet port to a VMnet interface. 
 IMPORTANT NOTE: Make sure that you have set Promiscuous mode security settings 
on the vSwitch and Port group to Accept. 
 IMPORTANT NOTE: If you are building trunk between EVE lab node to real Switch, 
please make sure you have set your ESXi vSwitch interface to accept all vlans.  
Reference: https://kb.vmware.com/s/article/1004074 
The Example below is showing ESXi Server settings of the virtual network bridged to the 
physical interface.

EVE-NG Community Cookbook 
Version 1.11 
Page 134 of 165 © EVE-NG LTD 
Logical chain of the networking bridge: 
EVE Lab Cloud0 →Portgroup “Management 90 UD”→vSwitch 1→Physical Adapter eth1 
 
vSwitch1 settings bridged with Server Ethernet port vmnic1 (physical adapter) 
 
Portgroup “Management 90 UD” Settings associated with vSwitch1 
 
EVE VM Settings 
EVE VM Cloud0 is connected to Portgroup “Management 90 UD”

EVE-NG Community Cookbook 
Version 1.11 
Page 135 of 165 © EVE-NG LTD 
 
 
EVE Lab Connected to a physical device 
Physical Topology 
Cisco 887M device port Fastethernet 3 is physically connected to Server port eth1.  
 
EVE Lab Topology 
EVE lab switch port G0/0 is configured as trunk and connected to Cloud0 over bridged chain 
to the physical Cisco 887M Router switchport Fastethernet 3

EVE-NG Community Cookbook 
Version 1.11 
Page 136 of 165 © EVE-NG LTD 
9.5.2 VMWare workstation EVE 
Similar to the ESXi connection, it is recommended to have a second ethernet interface on your 
PC. It can be a USB ethernet extender as well. Not all ethernet adapters fully support a layer2 
connectivity over it. MS Windows OS itself strips off any tags added to the packet. Even if your 
NIC supports 802.1q VLAN tagging, Windows 10 strips these tags off. The example below will 
show a Windows 10 host connected to a physical 3750G-24 switch. The Windows 10 Host has 
an Intel (R) PRO/1000 PT Dual port server adapter and is bridged with VMWare workstation 
(version 14) VMnets. 
Virtual Network Editor Settings, Bridged VMnet interfaces with Real NIC Ports 
 
EVE VM Settings. Network adapter is bridged to VMnet0 (ethernet Intel Pro 1), and Network 
adapter 2 is bridged to VMnet1 (ethernet Intel Pro 2).  
Responding cloud interfaces on EVE VM: 
Cloud0→Network Adapter→VMnet0→IntelPro 
Cloud1→Network Adapter 2→VMnet1→IntelPro#2

EVE-NG Community Cookbook 
Version 1.11 
Page 137 of 165 © EVE-NG LTD 
 
Physical connection scheme and VMware bridging. 
 
EVE Lab scheme.

EVE-NG Community Cookbook 
Version 1.11 
Page 138 of 165 © EVE-NG LTD 
The following solution allows Windows hosts to transmit tagged packets over ethernet. This 
has been used in the example above. 
 Warning. You are making changes to your Windows registry files! This is at your own 
risk. 
https://www.intel.co.uk/content/www/uk/en/support/articles/000005498/network -and-i-
o/ethernet-products.html 
9.5.3 Bare metal server EVE 
A physical server usually has more than one ethernet port, free ports 
can be bridged with EVE clouds and used for external connections. 
EVEs internal interface settings are already bridged in order, pnet0 -9 
are mapped to eth0-9. Refer to the bridging table in section 9.3 
cat /etc/network/interfaces 
Basically, your servers physical port eth0 is bridged to pnet0 which is Cloud0 in your labs, eth1 
is bridged to pnet1 which is Cloud1 in your labs (and so on). Refer to the bridgin g table in 
section 9.3 
The example below shows how to connect a bare -metal EVE server with a physical Cisco 
3750E switch. 
Physical connection topology: 
 
The EVE lab switch’s CDP neighbor is the 3750E switch’s p ort Gig 1/0/25: A trunk has been 
configured between the EVE lab switch and the physical 3750E switch.

EVE-NG Community Cookbook 
Version 1.11 
Page 139 of 165 © EVE-NG LTD 
10  Advanced EVE Lab features 
10.1 Lab design objects 
EVE Community has drawing elements integrates to add drawings and text information to the 
lab topology. Objects can be placed on the topology in two ways. 
Example below, EVE lab with design elements: 
 
Option 1: Side bar -> Add an object 
 
Option 2: Right -click on a fr ee area on the 
topology canvas to add an object. 
 
10.1.1 Custom shape  
There are three custom sha pes that can be added to the topology: square, round square and 
circle (sphere). 
Type: Square, round square or circle

EVE-NG Community Cookbook 
Version 1.11 
Page 140 of 165 © EVE-NG LTD 
Name: This field can be filled with your preferred shape’s name. If the field is left empty, EVE 
will generate a name for the shape. 
Border type: Two options: line or dashed 
Border width : Increase or decrease the width of the 
border. This can be edited later in the “Shape Edit” menu. 
Border colour: Allows you to choose a colour for the 
shape’s border. This can be edited later in the “Shape 
Edit” menu. 
Background colour: Allows you to choose a colour to fill 
your shape with. This can be edited later in the “Shape 
Edit” menu. 
 
Example: Added a circle and square on the topology. Shapes can be moved 
around the topology drag and drop style (click and move with mouse). 
 
10.1.2 Resize square or circle objects 
Move your mouse over the right bottom corner of the object until a 
corner symbol appears. Left click and drag your mouse to change 
object size or style (rectangle, sphere)  
10.1.3 Text object 
It is also possible to add text to your EVE topology.

EVE-NG Community Cookbook 
Version 1.11 
Page 141 of 165 © EVE-NG LTD 
 
Example: text objects added to the topology. 
 
 
 
10.1.4 Add custom picture on the Lab using Text object feature 
Sometimes you may have to add pictures, like logos on your topology. It is possible but you 
need to convert your png or jpg to html format. We have tested this one as the best to achieve 
result. Load your image in the web, and convert to html format. 
https://www.askapache.com/online-tools/base64-image-converter/ 
Step 1: Load your picture jpg or png format and encode it. 
 
Step 2: Scroll down to find HTML format 
 
Step 3: Set your desirable size of picture. 
 
Step 4: Mark and copy all content from HTML window above

EVE-NG Community Cookbook 
Version 1.11 
Page 142 of 165 © EVE-NG LTD 
Step 4: Copy content to EVE text object 
 
Step 5: Move and place your picture to the Lab. 
 
10.1.5 Cloning objects and overlay positions 
Right click on the object y ou want to clone and choose “Duplicate”. You can also change the 
object’s overlay position using the “Send to Back” or “Send to front” options. 
 
 
 
10.1.6 Objects Editing 
Right click the object and choose “Edit” for additional options.

EVE-NG Community Cookbook 
Version 1.11 
Page 143 of 165 © EVE-NG LTD 
 
At the bottom of the “Topology Canvas” page, additional object options will appear  
 
