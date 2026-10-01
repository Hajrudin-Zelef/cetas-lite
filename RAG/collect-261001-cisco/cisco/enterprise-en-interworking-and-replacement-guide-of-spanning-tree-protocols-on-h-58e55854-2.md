---
id: collect-261001-cisco/cisco/enterprise-en-interworking-and-replacement-guide-of-spanning-tree-protocols-on-h-58e55854-2
title: "Configure HuaweiA."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/enterprise-en-interworking-and-replacement-guide-of-spanning-tree-protocols-on-h-58e55854.md
source_anchor: ""
source_lines: [29, 76]
sha256: 5330aff5676d11600cacb20bd63ea857e173e0bd38015b0cdbce7e2e2ad9b623
---

# Configure HuaweiA.

To interwork with standard IEEE spanning tree protocols, Cisco develops PVST+ based on PVST. PVST+ provides interoperability with standard spanning tree protocols, which is an improvement made to PVST.
On an access interface, PVST+ sends standard STP BPDUs in its native VLAN. On a trunk interface, PVST+ sends standard STP BPDUs with the destination MAC address of 01-80-C2-00-00-00 only in VLAN 1, and sends Cisco proprietary BPDUs with the destination MAC address of 01-00-0C-CC-CC-CD in other VLANs allowed by the trunk interface.
Huawei S series switches support standard IEEE spanning tree protocols, and can process standard STP BPDUs from Cisco switches. However, Huawei S series switches forward Cisco proprietary BPDUs as multicast packets but not process them.
Rapid PVST+ is an extension of PVST+. Compared with PVST+, Rapid PVST+ uses the Rapid Spanning Tree Protocol (RSTP) mechanism to implement rapid transition.
Cisco MST supports VLAN-instance mapping and defines the region; therefore, it can be considered as a standard MSTP protocol. MST BPDUs use the standard format defined by the IEEE. Huawei and Cisco switches use different keys to generate MSTP digests in BPDUs, so the digests in BPDUs are different. By default, MSTP and Cisco MST can implement only inter-region interworking because Huawei and Cisco switches generate different digests. To enable MSTP and Cisco MST to interwork within an MST region, enable digest snooping on a Huawei S series switch connected to a Cisco switch and the Huawei S series switch's interface connected to the Cisco switch.
Processing mode of Cisco PVST+ BPDUs (the implementation is similar to that of Rapid PVST+ BPDUs)
On a trunk interface:
− In VLAN 1, a PVST+ device sends standard STP BPDUs and untagged PVST BPDUs to negotiate with the remote device.
− In the native VLAN but not VLAN 1, a PVST+ device sends untagged PVST BPDUs to negotiate with the remote device.
− In other VLANs, a PVST+ device sends PVST BPDUs to negotiate with the remote device.
− A PVST+ device sends standard STP BPDUs to negotiate with the remote device in VLAN 1 after the no spanning-tree vlan 1 command is configured globally.
On an access interface:
In all VLANs, a PVST+ device sends standard STP BPDUs to negotiate with the remote device.
By default, VLAN 1 is the native VLAN on a Cisco switch.
Processing of Huawei VBST BPDUs
− In VLAN 1, a VBST-enabled device sends standard STP or RSTP BPDUs and VBST BPDUs to negotiate with the remote device.
− In other VLANs, a VBST-enabled device sends VBST BPDUs to negotiate with the remote device.
l On an access interface:
A VBST-enabled device sends standard STP or RSTP BPDUs to negotiate with the remote device only in the VLAN where the access interface is located.
The Data field of VBST BPDUs and selection of packets of a standard protocol depend on the remote device connected to the Huawei S series switch. By default, standard RSTP BPDUs are used.
Table 1-2 Differences in command formats
| Function | Command on Huawei S Series Switches | Command on Cisco Switches | 
| Configure a spanning tree mode. | stp mode | spanning-tree mode | 
| Configure a path cost algorithm. | stp pathcost-standard | spanning-tree pathcost method | 
| Configure a fast transition mode on an interface. | stp no-agreement-check | - NOTE No such command is available on Cisco switches. Cisco switches support different convergence modes depending on the product model. For details, see Cisco product manuals. | 
| Enable digest snooping. | stp config-digest-snoop | - NOTE No such command is available on Cisco switches. Cisco switches do not support digest snooping. | 
Table 1-3 Differences between path cost algorithms
| Path Cost Algorithm | Command on Huawei S Series Switches |  | Command on Cisco Switches |  | 
|  | Query Command | Configuration Command | Query Command | Configuration Command | 
| IEEE 802.1t | display stp | stp pathcost-standard dot1t | show spanning-tree detail | spanning-tree pathcost method long | 
| IEEE 802.1d-1998 |  | stp pathcost-standard dot1d-1998 |  | spanning-tree pathcost method short | 
Before the 802.1s standard (MSTP) is released, vendors use different formats of digest fields in MSTP BPDUs. When devices from different vendors interwork with each other, negotiation may fail.
When a Huawei S series switch is connected to a Cisco switch, the two devices may fail to communicate because of different keys in BPDUs even though they have the same domain name, revision level, and VLAN mapping table. To solve this problem, enable digest snooping on the interface of the Huawei S series switch connected to the remote device. This function enables the Huawei switch to use the same key as the remote device, so that the Huawei S series switch can negotiate with the remote device.
Table 1-4 Comparison between digest commands
| Function | Command on Huawei S Series Switches | Command on Cisco Switches | 
| Check the digest information. | display stp region-configuration digest | show spanning-tree mst digest | 
| Enable digest snooping. | stp config-digest-snoop | - NOTE No such command is available on Cisco switches. Cisco switches do not support digest snooping. | 
There are three interworking and replacement solutions.
Huawei S series switches transparently transmit PVST BPDUs, and Cisco switches remove loops through negotiation.
Huawei S series switches running VBST interwork with Cisco switches running PVST, PVST+, or Rapid PVST+.
Huawei S series switches running MSTP interwork with Cisco switches running MST.
All models and versions of Huawei S series switches support MSTP, whereas VBST is supported in V200R005 and later versions.
A Huawei S series switch running VBST interworks with a Cisco switch running PVST, PVST+, or Rapid PVST+. They process protocol packets using the same mechanism, identify packets of each other, and use the same multicast MAC address 01-00-0C-CC-CC-CD. That is, their communication is similar to the communication between Huawei S series switches running VBST.
When a Huawei S series switch enabled with MSTP interworks with a Cisco switch running MST, digest snooping needs to be enabled on the Huawei S series switch because their digest formats are different. Their other implementations are the same.
A Cisco switch running PVST+ or Rapid PVST+ sends both PVST BPDUs and STP or RSTP BPDUs to negotiate with the remote device, so a Huawei S series switch running STP or RSTP can interwork with the Cisco switch. STP or RSTP convergence on the Huawei S series switch is based on ports, whereas PVST+ or Rapid PVST+ convergence on the Cisco switch is based on VLANs. The convergence results are as follows:
When a blocked port is located on the Huawei S series switch, data packets of all VLANs including PVST BPDUs of the Cisco switch are discarded on the blocked port. Therefore, the port is blocked in any VLAN.
When a blocked port is located on the Cisco switch, the Cisco switch running PVST+ or Rapid PVST+ only sends standard STP or RSTP BPDUs to negotiate with the remote device in VLAN 1. In this case, the blocked port only blocks packets from VLAN 1. It normally processes and forwards PVST BPDUs of other VLANs, and calculates the spanning tree of the VLAN where the port is located. The Huawei S series switch running STP or RSTP does not process PVST BPDUs, so blocked ports in other VLANs must be located on the Cisco switch.
Huawei S series switches transparently transmit Cisco PVST BPDUs to remove loops between Cisco switches or on themselves.
