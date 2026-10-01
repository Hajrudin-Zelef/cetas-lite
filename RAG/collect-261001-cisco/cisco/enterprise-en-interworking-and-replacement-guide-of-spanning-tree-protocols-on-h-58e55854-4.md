---
id: collect-261001-cisco/cisco/enterprise-en-interworking-and-replacement-guide-of-spanning-tree-protocols-on-h-58e55854-4
title: "Configure HuaweiA."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["cost", "parameters"]
source: docs/RAG/collect-261001-cisco/enterprise-en-interworking-and-replacement-guide-of-spanning-tree-protocols-on-h-58e55854.md
source_anchor: ""
source_lines: [168, 287]
sha256: c3e47beb4163aa0a06a76c24ec7408798d5991483290a18a00d660081931332c
---

# Configure HuaweiA.

6. Perform the preceding steps to migrate services on downstream access switches one by one. Figure 1-10 shows the network where migration is completed.
Figure 1-10 Network where migration is completed
a. Run the display stp [ vlan vlan-id ] [ interface interface-type interface-number | slot slot-id ] [ brief ] command to check the spanning tree status and statistics on Huawei S series switches.
PVST on Cisco switches is changed to MST so that Cisco switches can interwork with Huawei S series switches running MSTP.
In Figure 1-11, all switches are Cisco switches. Layer 3 switches establish a virtual switching system (VSS) to implement connectivity. Two aggregation switches establish a port channel in manual mode to provide link redundancy. They are configured with OSPF and establish OSPF relationships with core switches to receive and transmit routes, and are configured with the Hot Standby Router Protocol (HSRP) to implement virtual gateway backup. Switches are enabled with PVST to remove loops.
Cisco switches use the short algorithm to calculate the path cost. The short algorithm corresponds to the dot1d-1998 algorithm on Huawei S series switches. Cisco switches do not support fast transition in enhanced mode, whereas Huawei S series switches use fast transition in enhanced mode by default. You must run the stp no-agreement-check command to configure fast transition in common mode on the interfaces that do not support fast transition in enhanced mode. The format of the digest on a Cisco MST switch is different from that defined by the IEEE. CiscoA is the root bridge and CiscoB is the secondary root bridge in VLAN 10, and GE0/2 on CiscoC is the blocked port. CiscoB is the root bridge and CiscoA is the secondary root bridge in VLAN 20, and GE0/1 on CiscoD is the blocked port.
Figure 1-11 Networking where Cisco switches use MST to replace PVST to interwork with Huawei S series switches running MSTP
4. Configure MST on Cisco switches. Configure MSTP on Huawei S series switches and configure the path cost calculation algorithm and fast transmission mode so that Huawei S series switches can interwork with Cisco MST switches.
a. Configure MST on Cisco switches based on the original network plan.
i. Configure an MST region, create multiple MSTIs, and map VLAN 10 to MSTI 1 and VLAN 20 to MSTI 2.
ii. Configure the root bridge and secondary root bridge of each MSTI in each MST region.
iii. Configure the path cost of a port in each instance so that the port can be blocked.
iv. Enable MST.
b. Configure MSTP on Huawei S series switches based on the original network plan.
iii. Configure the path cost calculation algorithm on Huawei S series switches to be consistent with that on Cisco switches.
iv. Configure digest snooping on interfaces of Huawei S series switches connected to Cisco access switches.
v. Enable MSTP.
2. Power on two Huawei S series switches, and connect links between them and their uplinks. Configure addresses for downlink interfaces of core switches, and configure addresses for uplink interfaces and loopback addresses on Huawei S series switches. Change the spanning tree protocol to MST on Cisco switches and set parameters based on the original network plan. Complete the configuration on Huawei S series switches, and shut down VLANIF 10 and VLANIF 20 on Huawei A and Huawei B.
a. Configure MST on CiscoA and CiscoB to be replaced.
# Configure CiscoA.
CiscoA# configure terminal 
CiscoA(config)# spanning-tree mst configuration 
CiscoA(config)# spanning-tree extend system-id 
CiscoA(config-mst)# instance 1 vlan 10 
CiscoA(config-mst)# instance 2 vlan 20 
CiscoA(config-mst)# spanning-tree mst 1 priority 0 
CiscoA(config-mst)# spanning-tree mst 2 priority 24576 
CiscoA(config-mst)# name BG1 
CiscoA(config-mst)# revision 0 
CiscoA(config-mst)# exit 
CiscoA(config)# spanning-tree mode mst 
CiscoA(config)# end 
# Configure CiscoB.
CiscoB# configure terminal 
CiscoB(config)# spanning-tree mst configuration 
CiscoB(config)# spanning-tree extend system-id 
CiscoB(config-mst)# instance 1 vlan 10 
CiscoB(config-mst)# instance 2 vlan 20 
CiscoB(config-mst)# spanning-tree mst 1 priority 24576 
CiscoB(config-mst)# spanning-tree mst 2 priority 0 
CiscoB(config-mst)# name BG1 
CiscoB(config-mst)# revision 0 
CiscoB(config-mst)# exit 
CiscoB(config)# spanning-tree mode mst 
CiscoB(config)# end 
b. Configure MST on CiscoC and CiscoD (access switches).
# Configure CiscoC.
CiscoC# configure terminal 
CiscoC(config)# spanning-tree mst configuration 
CiscoC(config)# spanning-tree extend system-id 
CiscoC(config-mst)# instance 1 vlan 10 
CiscoC(config-mst)# instance 2 vlan 20 
CiscoC(config-mst)# name BG1 
CiscoC(config-mst)# revision 0 
CiscoC(config-mst)# exit 
CiscoC(config)# spanning-tree mode mst 
CiscoC(config)# interface gigabitethernet 0/2 
CiscoC(config-if)# spanning-tree mst 1 cost 20000 
CiscoC(config-if)# exit 
CiscoC(config)# end 
# Configure CiscoD.
CiscoC# configure terminal 
CiscoC(config)# spanning-tree mst configuration 
CiscoC(config)# spanning-tree extend system-id 
CiscoC(config-mst)# instance 1 vlan 10 
CiscoC(config-mst)# instance 2 vlan 20 
CiscoC(config-mst)# name BG1 
CiscoC(config-mst)# revision 0 
CiscoC(config-mst)# exit 
CiscoC(config)# spanning-tree mode mst 
CiscoC(config)# interface gigabitethernet 0/1 
CiscoC(config-if)# spanning-tree mst 2 cost 20000 
CiscoC(config-if)# exit 
CiscoC(config)# end 
c. Configure MSTP on HuaweiA and HuaweiB.
<HUAWEI> system-view 
[HUAWEI] syaname HuaweiA 
[HuaweiA] stp region-configuration 
[HuaweiA-mst-region] region-name RG1 
[HuaweiA-mst-region] instance 1 vlan 10 
[HuaweiA-mst-region] instance 2 vlan 20 
[HuaweiA-mst-region] active region-configuration 
[HuaweiA-mst-region] quit 
[HuaweiA] stp pathcost-standard dot1d-1998 
[HuaweiA] stp instance 1 root primary 
[HuaweiA] stp instance 2 root secondary 
[HuaweiA] interface gigabitethernet 0/0/1 
[HuaweiA-GigabitEthernet0/0/1] stp no-agreement-check 
[HuaweiA-GigabitEthernet0/0/1] stp config-digest-snoop 
[HuaweiA-GigabitEthernet0/0/1] quit 
[HuaweiA] interface gigabitethernet 0/0/2 
[HuaweiA-GigabitEthernet0/0/2] stp no-agreement-check 
[HuaweiA-GigabitEthernet0/0/2] stp config-digest-snoop 
[HuaweiA-GigabitEthernet0/0/2] quit 
<HUAWEI> system-view 
[HUAWEI] syaname HuaweiB 
[HuaweiB] stp region-configuration 
[HuaweiB-mst-region] region-name RG1 
[HuaweiB-mst-region] instance 1 vlan 10 
[HuaweiB-mst-region] instance 2 vlan 20 
[HuaweiB-mst-region] active region-configuration 
[HuaweiB-mst-region] quit 
[HuaweiB] stp pathcost-standard dot1d-1998 
[HuaweiB] stp instance 1 root secondary 
[HuaweiB] stp instance 2 root primary 
[HuaweiB] interface gigabitethernet 0/0/1 
[HuaweiB-GigabitEthernet0/0/1] stp no-agreement-check 
[HuaweiB-GigabitEthernet0/0/1] stp config-digest-snoop 
[HuaweiB-GigabitEthernet0/0/1] quit 
[HuaweiB] interface gigabitethernet 0/0/2 
[HuaweiB-GigabitEthernet0/0/2] stp no-agreement-check 
[HuaweiB-GigabitEthernet0/0/2] stp config-digest-snoop 
[HuaweiB-GigabitEthernet0/0/2] quit 
3. Migrate services of the backup uplink of CiscoD to HuaweiB and shut down VLANIF 20, as shown in Figure 1-12.
Figure 1-12 Migration process 1
5. Test services on CiscoD. When verifying that services on CiscoD are normal, migrate services on the link between CiscoD and CiscoA to HuaweiA. The migration of the access switch is completed, as shown in Figure 1-13.
Figure 1-13 Migration process 2
6. Perform the preceding steps to migrate services on downstream access switches one by one. Figure 1-14 shows the network where migration is completed.
Figure 1-14 Network where migration is completed
If you have any problems, please post them in our Community. We are happy to solve them for you!
