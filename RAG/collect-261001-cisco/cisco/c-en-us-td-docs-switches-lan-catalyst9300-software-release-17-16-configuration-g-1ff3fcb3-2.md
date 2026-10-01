---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-1ff3fcb3-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-1ff3fcb3"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-1ff3fcb3.md
source_anchor: ""
source_lines: [107, 255]
sha256: 7b8252cfb1779915be6540b42f148410acafb9fdbfe2f84d7b740c78d24f82d0
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-1ff3fcb3

Step 4
switchport mode trunk
Example:
Device(config-if)# switchport mode trunk
Configures the interface as a trunk port.
Step 5
switchport vlan mappingvlan-id translated-id
Example:
Device(config-if)# switchport vlan mapping 2 102
Enters the VLAN IDs to be mapped:
vlan-id —the customer VLAN ID (C-VLAN) entering the switch from the customer network. The range is from 1 to 4094.
translated-id —the assigned service-provider VLAN ID (S-VLAN). The range is from 1 to 4094.
Step 6
exit
Example:
Device(config-if)# exit
Returns to global configuration mode.
Step 7
spanning-tree bpdufilter enable
Example:
Device(config)# spanning-tree bpdufilter enable
Inserts a BPDU filter for spanning tree.
Note
To process control traffic consistently, either enable Layer 2 protocol tunneling (recommended) or insert a BPDU filter for
spanning tree.
Step 8
end
Example:
Device(config)# end
Returns to privileged EXEC mode.
Step 9
show vlan mapping
Example:
Device# show vlan mapping
Verifies the configuration.
Step 10
copy running-config startup-config
Example:
Device# copy running-config startup-config
(Optional) Saves your entries in the configuration file.
Example
Use no switchport vlan mapping command to remove the VLAN mapping information. Entering no switchport vlan mapping all command deletes all mapping configurations.
This example shows how to map VLAN IDs 2 to 6 in the customer network to VLANs 101 to 105 in the service-provider network
(Figure 3-5). You configure the same VLAN mapping commands for a port in Switch A and Switch B; the traffic on all other VLAN
IDs is forwarded as normal traffic.
In the previous example, at the ingress of the service-provider network, VLAN IDs 2 to 6 in the customer network are mapped
to VLANs 101 to 105, in the service provider network. At the egress of the service provider network, VLANs 101 to 105 in the
service provider network are mapped to VLAN IDs 2 to 6, in the customer network.
Note
Packets with VLAN IDs other than the ones with configured VLAN Mapping are forwarded as normal traffic.
Use show vlan mapping command to view information about configured vlans.
Device> enable
Device# configure terminal
Device(config)# show vlan mapping
Total no of vlan mappings configured: 1
Interface Po5:
VLANs on wire Translated VLAN Operation
------------------------------ --------------- --------------
20 30 1-to-1
Selective Q-in-Q on a Trunk Port
To configure VLAN mapping for selective Q-in-Q on a trunk port, perform this task:
Note
You cannot configure one-to-one mapping and selective Q-in-Q on the same interface.
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configure terminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
interface interface-id
Example:
Device(config)# interface gigabitethernet1/0/1
Enters interface configuration mode for the interface that is connected to the service-provider network. You can enter a physical
interface or an EtherChannel port channel.
vlan-id —the customer VLAN ID (C-VLAN) entering the switch from the customer network. The range is from 1 to 4094. You can
enter a string of VLAN-IDs.
outer-vlan-id —The outer VLAN ID (S-VLAN) of the service provider network. The range is from 1 to 4094.
Use the no form of this command to remove the VLAN mapping configuration. Entering the no switchport vlan mapping all command deletes all mapping configurations.
Specifies that all unmapped packets on the port are forwarded with the specified S-VLAN.
By default, packets that do not match the mapped VLANs, are dropped.
Untagged traffic are forwarded without dropping.
Step 7
exit
Example:
Device(config-if)# exit
Returns to global configuration mode.
Step 8
spanning-tree bpdufilter enable
Example:
Device(config)# spanning-tree bpdufilter enable
Inserts a BPDU filter for spanning tree.
Note
To process control traffic consistently, either enable Layer 2 protocol tunneling (recommended) or insert a BPDU filter for
spanning tree.
Step 9
end
Example:
Device(config)# end
Returns to privileged EXEC mode.
Step 10
show interfaces interface-id vlan mapping
Example:
Device# show interfaces gigabitethernet1/0/1 vlan mapping
Verifies the configuration.
Step 11
copy running-config startup-config
Example:
Device# copy running-config startup-config
(Optional) Saves your entries in the configuration file.
Example
This example shows how to configure selective QinQ mapping on the port so that traffic with a C-VLAN ID of 2 to 5 enters the
switch with an S-VLAN ID of 100. By default, the traffic of any other VLAN ID is dropped.
This example shows how to configure selective QinQ mapping on the port so that traffic with a C-VLAN ID of 2 to 5 enters the
switch with an S-VLAN ID of 100. The traffic of any other VLAN ID is forwarded with the S-VLAN ID of 200.
Device(config)# interface GigabiEthernet0/1
Device(config-if)# switchport vlan mapping 2-5 dot1q-tunnel 100
Device(config-if)# switchport vlan mapping default dot1q-tunnel 200
Device(config-if)# exit
Device# show vlan mapping
Total no of vlan mappings configured: 5
Interface Hu1/0/50:
VLANs on wire Translated VLAN Operation
------------------------------ --------------- --------------
2-5 100 selective QinQ
* 200 default QinQ
Feature History for VLAN Mapping
This table provides release and related information for features explained in this module.
These features are available on all releases subsequent to the one they were introduced in, unless noted otherwise.
Release
Feature
Feature Information
Cisco IOS XE Gibraltar 16.11.1
One-to-One VLAN mapping
One-to-One VLAN mapping allows to map customer VLANs to service-provider VLANs on trunk ports that are connected to a customer
network.
Cisco IOS XE Bengaluru 17.5.1
Selective Q-in-Q
Support for selective Q-in-Q was introduced
Use Cisco Feature Navigator to find information about platform and software image support. To access Cisco Feature Navigator,
go to http://www.cisco.com/go/cfn.
