---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-18
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [3117, 3366]
sha256: 317aedf9f62828185be90c737fd8f2ae501c8dae92aa176581fa7a166c70ed85
---

# Copyright (c) 2016, Andrea Dainese

EVE-NG Community Cookbook 
Version 1.11 
Page 123 of 165 © EVE-NG LTD 
8.8 Deleting labs 
Refer to section 6.2.2.2 
8.9 Moving labs 
Refer to section 6.2.2.4

EVE-NG Community Cookbook 
Version 1.11 
Page 124 of 165 © EVE-NG LTD 
9  EVE Clouds and Networks 
9.1 Bridge Network  
The EVE Bridge interface act s like an unmanaged Switch. It supports passing along tagged 
dot1q packets. 
Example: We have to connect many nodes in a flat (dot1q) network 
Step 1: Add a Bridge Network onto the topol ogy. There are two ways to do this: Right -clicking 
on the topology area and selecting “Add Network” or in the sidebar  click “Add an Object” and 
then select “Network.” Please refer to sections 7.2.3 and 7.1.1.2 
  
Step 2: Name/prefix can be changed in order to rename your Bridge network. Make sure your 
network type is set to bridge. 
 
Step 3: Connect your nodes using the drag and drop connector. Re fer to sections 8.1.4 and 
7.2.3

EVE-NG Community Cookbook 
Version 1.11 
Page 125 of 165 © EVE-NG LTD 
9.2 Management Cloud0 interface 
EVE management interface is also known  as the Cloud0 
network for labs. The Cloud0 interface is bridged with your 
EVEs first NIC. “Cloud” is used as an alias to pnet. Pnet is 
the bridge interface name inside of EVE. 
Cloud0 is commonly used inside EVE labs to get management access to nodes running inside 
EVE from a host machine external to EVE. 
 IMPORTANT NOTE: For EVE VMs running  on ESXi, make sure your management 
interface bridged with the vSwitch (Port group) has the security settings for 
Promiscuous Mode set to Accept. Any port group or vSwitch used to connect an 
external network to an EVE Cloud network need s to have the Promis cuous mode set 
to “Accept”! 
vSwitch Settings 
 
Portgroup Settings

EVE-NG Community Cookbook 
Version 1.11 
Page 126 of 165 © EVE-NG LTD 
EVE Cloud0 bridging table. 
Lab name 
EVE 
interface 
name (inside) 
Type Notes 
Cloud0 pnet0 Bridged 
Cloud0/pnet0 is bridged with your primary 
EVE ethernet port. It is assigned a 
management I P address used for WEB 
GUI access. The  EVE management  
subnet can be used as a management 
network in labs. 
 Question: How can I obtain my Cloud0 subnet and gateway IP. Many EVE VMs only 
have a DHCP address assigned on the pnet0 interface. 
Answer: SSH to EVE and type the following from the CLI: 
route 
 
Example: We want to use Cloud0 as a management network for an ASAv node in an EVE lab. 
From the above-obtained information, we know that our Cloud management subnet is 
192.168.90.0 with a mask of 255.255.255.0 and the Gateway IP is 192.168.90.1. 
Step 1:  Add A New Network onto the topology. 
There are two ways to do this: Right -clicking on 
topology area and selecting “Network” or in the 
sidebar, “Add an Object” and then select “Network.” 
Step 2:  Name/prefix can b e changed in order to 
rename your Cloud0 network. Make sure your 
network type is set to Management(Cloud0). 
Step 3: Connect your ASAv using the drag and drop 
connector to the Cloud0 network. Refer to sections 
8.1.4 and 7.2.3 
Step 4: Start the node and configure the interface connected to Cloud0 with an IP address from 
the management subnet (192.168.90.0/24 in this example). Make sure you do not assign 
duplicate IPs.

EVE-NG Community Cookbook 
Version 1.11 
Page 127 of 165 © EVE-NG LTD 
 
NOTE: Cloud interfaces can be used to connect multiple nodes to a single cloud instance on 
the topology. 
 
9.3 Other cloud interfaces 
Other cloud interfaces can be used to extend a lab connection inside of EVE or bridged with 
other EVE interfaces to connect external networks or devices. 
EVE Cloud bridging table. 
Lab cloud 
name 
EVE 
interface 
name 
(inside) 
Type 
ESXi VM 
corresponding 
interface 
VMware 
Workstation 
corresponding 
interface 
Bare HW 
Server Notes 
Cloud0 Pnet0 bridged Network 
adapter 1 
Network 
Adapter 
First 
ethernet 
Eth0 
Cloud0/pnet0 is bridged with 
your primary EVE ethernet port. 
It is assigned a management IP 
address used for WEB GUI 
access. The EVE management 
subnet can be used as

EVE-NG Community Cookbook 
Version 1.11 
Page 128 of 165 © EVE-NG LTD 
management network in the 
labs. 
Cloud1 Pnet1 bridged Network 
adapter 2 
Network 
Adapter 2 
Second 
ethernet 
Eth1 
Cloud1 can be bridged with your 
EVE second ethernet port to 
achieve connection to another 
network or device. The IP 
address is not required to be 
configured on it. It will act like a 
pure bridge your external 
connection with EVE lab node. 
Cloud2 Pnet2 bridged Network 
adapter 3 
Network 
Adapter 3 
Third 
ethernet 
Eth2 
Same as Cloud1 
Cloud3 Pnet3 bridged Network 
adapter 4 
Network 
Adapter 4 
Fourth 
ethernet 
Eth3 
Same as Cloud1 
Cloud4-9 Pnet4-9 bridged Network 
adapter 5-10 
Network 
Adapter 5-10  Same as Cloud1 
Example: Cloud7 network is used as an extended connector between nodes: 
Step 1: Add two Cloud7 networks onto the topology. 
 
Step 2: Connect your lab nodes to Cloud7. Your 
configured nodes will work like being connected to 
the same switch (or the same bridge in EVE). 
Even CDP works. It is convenient if it is necessary 
to have connections across the lab and you don’t 
want to have connections going from one end of 
the lab to the other.

EVE-NG Community Cookbook 
Version 1.11 
Page 129 of 165 © EVE-NG LTD 
 
If some of the clouds (e.g. Cloud1 ) are bridged to another ethernet (VMnet) you can connect 
your EVE lab to an external VM or physical device (like e.g. a switch, IP phone or access point). 
 For ESXi make sure that you have set Promiscuous mode security setting s on the 
vSwitch and Port group to Accept. Please refer to section 9.2 
The next sections will explain how you can use Cloud networks in EVE to connect to other 
external (e.g. VMWare) VMs or physical devices. 
9.4 Connecting external VM machines to the EVE Lab 
9.4.1 ESXi VM machines 
External ESXi VM machines can be connected to EVE labs using cloud interfaces.  
 NOTE: A single Cloud interface can be used to connect more than one external VM to 
the EVE lab. 
Example:  Connecting a Web Security Appliance (WSA) to the lab using the Cloud1 interface. 
Step 1: Create a new or use an existing portgroup on your ESXi and assign it to EVE and WSA 
VMs as shown below. Make sure you have set Promiscuous mode on the vSwitch (portgroup 
WSA-MGMT) to Accept. 
 NOTE: VM machines must be in a powered off state to assign network interfaces. 
Portgroup WSA-MGMT (with vSwitch5 as parent) settings:

EVE-NG Community Cookbook 
Version 1.11 
Page 130 of 165 © EVE-NG LTD 
 
Parent vSwitch5 settings: 
 
EVE and WSA VMs settings 
EVE VM, second port is assigned to 
portgroup WSA-MGMT. It is Cloud1 on the 
EVE topology.  
 
Cisco Web security appliance (WSA), 
Management port is assigned in portgroup 
WSA-MGMT.  
 
 
EVE Lab connected to the WSA (Cloud1) 
 NOTE: ESXi WSA VM obtained the IP 192.168.10.3 from the DHCP pool on the lab 
switch. The gateway is 192.168.10.1 
 NOTE: The Firefox Docker node user for management obtained the IP 192.168.10.2 
from the DHCP pool configured on the lab switch.

EVE-NG Community Cookbook 
Version 1.11 
Page 131 of 165 © EVE-NG LTD 
 
 
 
 
 
 
 
 
 
 
 
 
9.4.2 VMWare workstation machines 
External (meaning not running inside EVE) VMWare workstation machines can be connected 
to EVE labs using cloud interfaces. 
 NOTE: A single Cloud interface can be used to connect more than one external VM to 
the EVE lab. 
Example:  Connecting Web security Appliance (WSA) to the lab using Cloud2 interface. 
 NOTE: VMs must be in a powered off state to assign network interfaces. 
Step 1:  Open your VMWare Workstation Virtual Network Editor and configure the VMnet 
interface for the Cloud and WSA VMs. If necessary, add a new VMnet. The example below is 
showing VMnet2 Settings in VMWare workstation. DHCP must be disabled for VMnet2. 
Virtual Network Editor settings:

EVE-NG Community Cookbook 
Version 1.11 
Page 132 of 165 © EVE-NG LTD 
 
