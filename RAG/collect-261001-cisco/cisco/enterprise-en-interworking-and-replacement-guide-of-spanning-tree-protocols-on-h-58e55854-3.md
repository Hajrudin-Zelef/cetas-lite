---
id: collect-261001-cisco/cisco/enterprise-en-interworking-and-replacement-guide-of-spanning-tree-protocols-on-h-58e55854-3
title: "Configure HuaweiA."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["cost", "parameters"]
source: docs/RAG/collect-261001-cisco/enterprise-en-interworking-and-replacement-guide-of-spanning-tree-protocols-on-h-58e55854.md
source_anchor: ""
source_lines: [77, 167]
sha256: 928ae3ca0a003476030ed2d1a83a17ff1f7bfe6b5da0b4e2c66df3ca6f26f29f
---

# Configure HuaweiA.

In Figure 1-3, all switches are Cisco switches. Layer 3 switches establish a virtual switching system (VSS) to implement connectivity. Two aggregation switches establish a port channel in manual mode to provide link redundancy. They are configured with OSPF and establish OSPF relationships with core switches to receive and transmit routes, and are configured with the Hot Standby Router Protocol (HSRP) to implement virtual gateway backup. Switches are enabled with PVST to remove loops.
Huawei S series switches need to replace the two aggregation switches, without changing the network plan.
Figure 1-3 Networking where Huawei S series switches transparently transmit Cisco PVST BPDUs
The configuration roadmap is as follows:
1. Configure OSPF on Huawei S series switches to establish OSPF relationships with core switches to receive and transmit routes.
2. Configure link aggregation in manual mode on Huawei S series switches to implement load balancing.
3. Configure VRRP on Huawei S series switches to interwork with and replace HSRP on Cisco switches. VRRP implements virtual gateway backup.
4. Configure Huawei S series switches to transparently transmit Cisco PVST BPDUs to remove loops between Cisco switches or on themselves.
a. Disable STP on Huawei S series switches.
b. Configure Huawei S series switches to transparently transmit Cisco PVST BPDUs.
5. Configure service forwarding on Huawei S series switches based on the original network plan.
During migration, connect Huawei switches in bypass mode and establish OSPF routes. Migrate services on access switches to Huawei S series switches one by one.
1. Check the configuration of Cisco switches before the replacement.
a. Run the show running-config command to check the spanning tree configuration on Cisco switches.
b. Run the show spanning-tree summary command to check spanning tree parameters and status information on Cisco switches.
Cisco switches use PVST to calculate spanning trees.
2. Power on two Huawei S series switches, and connect links between them and their uplinks. Configure addresses for downlink interfaces of core switches, and configure addresses for uplink interfaces and loopback addresses on Huawei S series switches. Complete the configuration on Huawei S series switches, and shut down VLANIF 10 and VLANIF 20 on HuaweiA and HuaweiB. Retain the configuration of Cisco switches.
a. Disable STP on HuaweiA and HuaweiB.
# Configure HuaweiA.
<HUAWEI> system-view 
[HUAWEI] syaname HuaweiA 
[HuaweiA] stp disable 
# Configure HuaweiB.
<HUAWEI> system-view 
[HUAWEI] syaname HuaweiB 
[HuaweiB] stp disable 
[HuaweiA] interface eth-trunk 1 
[HuaweiA-Eth-Trunk1] l2protocol-tunnel PVST+ enable 
[HuaweiA-Eth-Trunk1] quit 
[HuaweiA] interface gigabitethernet 0/0/1 
[HuaweiA-GigabitEthernet0/0/1] l2protocol-tunnel PVST+ enable 
[HuaweiA-GigabitEthernet0/0/1] quit 
[HuaweiA] interface gigabitethernet 0/0/2 
[HuaweiA-GigabitEthernet0/0/2] l2protocol-tunnel PVST+ enable 
[HuaweiA-GigabitEthernet0/0/2] quit 
[HuaweiB] interface eth-trunk 1 
[HuaweiB-Eth-Trunk1] l2protocol-tunnel PVST+ enable 
[HuaweiB-Eth-Trunk1] quit 
[HuaweiB] interface gigabitethernet 0/0/1 
[HuaweiB-GigabitEthernet0/0/1] l2protocol-tunnel PVST+ enable 
[HuaweiB-GigabitEthernet0/0/1] quit 
[HuaweiB] interface gigabitethernet 0/0/2 
[HuaweiB-GigabitEthernet0/0/2] l2protocol-tunnel PVST+ enable 
[HuaweiB-GigabitEthernet0/0/2] quit 
3. Migrate services of the backup uplink of CiscoD to HuaweiB and shut down VLANIF 20, as shown in Figure 1-4.
Figure 1-4 Migration process 1
4. Disconnect the cable between CiscoA and CiscoD, shut down VLANIF 20 on CiscoA and CiscoB, and enable VLANIF 20 on Huawei S series switches.
5. Test services on CiscoD. When verifying that services on CiscoD are normal, migrate services on the link between CiscoD and CiscoA to Huawei A. The migration of the access switch is completed, as shown in Figure 1-5.
Figure 1-5 Migration process 2
6. Perform the preceding steps to migrate services on downstream access switches one by one. Figure 1-6 shows the network where migration is completed.
Figure 1-6 Network where migration is completed
7. Check the configuration of Huawei S series switches after the replacement.
a. Run the display l2protocol-tunnel group-mac { all | protocol-type | user-defined-protocol protocol-name } command to check whether Huawei S series switches can transparently transmit Cisco PVST BPDUs.
b. Run the show spanning-tree summary command to check spanning tree status information on Cisco switches.
c. Verify services on user-side devices and check whether the replacement is successful.
Huawei S series switches transparently transmit Cisco PVST BPDUs to implement spanning tree negotiation between Cisco switches, so Huawei S series switches broadcast received PVST BPDUs in VLANs. As a result, P2P negotiation between two switches is changed to P2MP negotiation, affecting spanning tree convergence. Solution 1 causes slow spanning tree convergence and easily results in temporary loops.
Huawei S series switches are configured with VBST to interwork with Cisco PVST switches to remove loops.
In Figure 1-7, all switches are Cisco switches. Layer 3 switches establish a virtual switching system (VSS) to implement connectivity. Two aggregation switches establish a port channel in manual mode to provide link redundancy. They are configured with OSPF and establish OSPF relationships with core switches to receive and transmit routes, and are configured with the Hot Standby Router Protocol (HSRP) to implement virtual gateway backup. Switches are enabled with Rapid PVST+ to remove loops.
Cisco switches use the short algorithm to calculate the path cost. The short algorithm corresponds to the dot1d-1998 algorithm on Huawei S series switches. Cisco switches do not support fast transition in enhanced mode, whereas Huawei S series switches use fast transition in enhanced mode by default. You must run the stp no-agreement-check command to configure fast transition in common mode on the interfaces that do not support fast transition in enhanced mode.
Figure 1-7 Networking where Huawei S series switches use VBST to interwork with Cisco PVST switches
4. Configure VBST on Huawei S series switches to interwork with Rapid PVST+ on Cisco switches.
a. Configure Huawei S series switches to work in VBST mode.
b. Configure Huawei S series switches to use the dot1d-1998 algorithm to calculate the path cost.
c. Configure fast transition in common mode on Huawei S series switches.
Cisco switches use Rapid PVST+ to calculate spanning trees and use the short algorithm to calculate the path cost. In addition, they do not support fast transition in enhanced mode.
Configure VBST on Huawei S series switches.
a. Configure HuaweiA and HuaweiB to work in VBST mode.
<HUAWEI> system-view 
[HUAWEI] syaname HuaweiA 
[HuaweiA] stp mode vbst 
<HUAWEI> system-view 
[HUAWEI] syaname HuaweiB 
[HuaweiB] stp mode vbst 
[HuaweiA] stp pathcost-standard dot1d-1998 
[HuaweiB] stp pathcost-standard dot1d-1998 
[HuaweiA] interface gigabitethernet 0/0/1 
[HuaweiA-GigabitEthernet0/0/1] stp no-agreement-check 
[HuaweiA-GigabitEthernet0/0/1] quit 
[HuaweiA] interface gigabitethernet 0/0/2 
[HuaweiA-GigabitEthernet0/0/2] stp no-agreement-check 
[HuaweiA-GigabitEthernet0/0/2] quit 
[HuaweiB] interface gigabitethernet 0/0/1 
[HuaweiB-GigabitEthernet0/0/1] stp no-agreement-check 
[HuaweiB-GigabitEthernet0/0/1] quit 
[HuaweiB] interface gigabitethernet 0/0/2 
[HuaweiB-GigabitEthernet0/0/2] stp no-agreement-check 
[HuaweiB-GigabitEthernet0/0/2] quit 
3. Migrate services of the backup uplink of CiscoD to HuaweiB and shut down VLANIF 20, as shown in Figure 1-8.
Figure 1-8 Migration process 1
5. Test services on CiscoD. When verifying that services on CiscoD are normal, migrate services on the link between CiscoD and CiscoA to Huawei A. The migration of the access switch is completed, as shown in Figure 1-9.
Figure 1-9 Migration process 2
