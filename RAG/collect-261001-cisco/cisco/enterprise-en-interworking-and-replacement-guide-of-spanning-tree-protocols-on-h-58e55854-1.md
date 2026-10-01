---
id: collect-261001-cisco/cisco/enterprise-en-interworking-and-replacement-guide-of-spanning-tree-protocols-on-h-58e55854-1
title: "Configure HuaweiA."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/enterprise-en-interworking-and-replacement-guide-of-spanning-tree-protocols-on-h-58e55854.md
source_anchor: ""
source_lines: [1, 28]
sha256: 447c4a3e32e75b4bf27459ef15ab7bbb31f4d76f62ca54fb58cf9cccbd40b023
---

# Configure HuaweiA.

Hi, everyone!
Today I’m going to introduce the interworking and replacement guide of spanning-tree protocols on Huawei switches and Cisco switches.
Huawei S series switches support the following spanning tree protocols: Spanning Tree Protocol (STP), Rapid Spanning Tree Protocol (RSTP), Multiple Spanning Tree Protocol (MSTP), and VLAN-Based Spanning Tree (VBST). VBST is supported by V200R005 and later versions.
STP is used on a LAN to prevent loops. Devices running STP discover loops on a network by exchanging information, and block some ports to eliminate loops. Both STP and RSTP (which is an evolution of STP and allows for fast network topology convergence) suffer from a significant limitation: neither can implement VLAN-based load balancing because all VLANs on a LAN use one spanning tree. When a link is blocked, it no longer transmits traffic, which wastes bandwidth and prevents certain VLAN packets from being forwarded. Based on STP and RSTP, MSTP allows fast convergence and provides multiple paths to load balance VLAN traffic. VBST, a Huawei spanning tree protocol, constructs a spanning tree in each VLAN so that traffic from different VLANs is forwarded through different spanning trees. Traffic is therefore load balanced.
Huawei MSTP and Cisco Multiple Spanning Tree (MST) are based on IEEE standards, but their implementations are different. VBST is a Huawei proprietary protocol. The following describes working mechanisms of Huawei MSTP and VBST.
MSTP is a new spanning tree protocol defined in IEEE 802.1s. MSTP uses Multiple Spanning Tree Instances (MSTIs) and Multiple Spanning Tree (MST) regions. An MSTI is a collection of VLANs. Binding multiple VLANs to a single MSTI reduces communication costs and resource usage. The topology of each MSTI is calculated independently, and traffic can be balanced among MSTIs. Multiple VLANs with the same topology can be mapped to a single MSTI. The forwarding state of the VLANs for a port is determined by the port state in the MSTI.
MST region: consists of the configuration name, revision level, configuration identifier format selector, and mapping between VLANs and MSTIs. The configuration name, configuration identifier format selector, and revision level have corresponding fields in a Bridge Protocol Data Unit (BPDU). The mapping between VLANs and MSTIs is the configuration digest in the BPDU, which is a 16-byte signature calculated depending on the mapping. All switches in an MST region must have the same MST region configuration. By default, the configuration name is the first MAC address of a switch, the revision level is 0, the configuration identifier format selector is 0, and all VLANs are mapped to MSTI 0.
MSTI: Each MSTI corresponds to one VLAN or a group of VLANs, whereas each VLAN corresponds only to one MSTI. Each switch can run multiple MSTIs. When the mapping between VLANs and MSTIs is not configured, all VLANs are mapped to MSTI 0.
CIST: The Common Spanning Tree (CST) and the Internal Spanning Tree (IST) construct a CIST. The IST provides connectivity for an MST region. The IST is a segment of the CIST in an MST region and is a special MSTI with an MSTI ID of 0. The CST is a segment of the CIST (each MST region is considered as a single node).
CIST root and regional root: Compared with STP and RSTP, MSTP uses the CIST root and regional root. For all connected switches running STP/RSTP/MSTP, there is only one CIST root. Each MSTI in each MST region has a regional root. There is only one CIST root on the switching network, whereas the number of regional roots in each region depends on the number of MSTIs.
External path cost and internal path cost: Compared with STP and RSTP, MSTP uses the external path cost and internal path cost. The external path cost corresponds to the CIST and is consistent in an MST region. Each MSTI in each MST region has an internal path cost. Different internal path costs correspond to different MSTIs on the same port.
Edge port, master port, and alternate port: Compared with STP and RSTP, MSTP uses the edge port and master port. The regional edge port connects ports in different MST regions, MST regions and regions running STP, or MST regions and regions running RSTP, and is located at the edge of an MST region. Among all edge ports in an MST region, the master port has the smallest cost to the CIST root. The master port is located on the shortest path connecting MST regions to the CIST root. The alternate port is the backup port of the master port. If the master port is blocked, the alternate port becomes the new master port.
Figure 1-1 shows the format of an MST BPDU.
Figure 1-1 Format of an MST BPDU
VBST, a Huawei spanning tree protocol, constructs a spanning tree in each VLAN so that traffic from different VLANs is forwarded through different spanning trees. VBST is equivalent to STP or RSTP running in each VLAN. Spanning trees in different VLANs are independent of each other.
VBST transmits VBST BPDUs in VLANs but not VLAN 1 to determine the network topology. VBST BPDUs are protocol packets of VBST based on STP or RSTP BPDUs. Compared with STP or RSTP BPDUs, a 4-byte 802.1Q tag is added between the source MAC address field and the protocol length field of a VBST BPDU. Figure 1-2 compares encapsulation formats between an STP or RSTP BPDU and a VBST BPDU.
Figure 1-2 Comparisons between encapsulation formats of an STP or RSTP BPDU and a VBST BPDU
Table 1-1 compares VBST, STP, RSTP, and MSTP of Huawei S series switches in convergence speed, traffic forwarding, applicable scenario, and configuration complexity.
Table 1-1 Comparison among four spanning tree protocols
| Spanning Tree Protocol | Similarity | Difference |  |  |  | 
|  |  | Convergence Speed | Traffic Forwarding | Applicable Scenario | Configuration Complexity | 
| STP | Forms a loop-free tree to prevent broadcast storms and implements redundancy. | Slowest | All VLANs share one spanning tree, and traffic from all VLANs is transmitted along the same path. | Service traffic does not need to be differentiated. | Low | 
| RSTP |  | RSTP, MSTP, and VBST provide the same topology convergence speed and offer faster topology convergence than STP. |  |  | Low | 
| MSTP |  |  | MSTP provides mappings between MSTIs and VLANs so that traffic from different VLANs is forwarded through different spanning trees that are independent of each other. | Service traffic needs to be differentiated and load balanced. | Medium | 
| VBST |  |  | A spanning tree is formed in each VLAN, so that traffic from different VLANs is forwarded through different spanning trees that are independent of each other. | l Service traffic needs to be differentiated and load balanced. l VBST can interwork with the Per-VLAN Spanning Tree (PVST), Per-VLAN Spanning Tree Plus (PVST+), and Rapid PVST+. | High | 
Cisco switches support the following spanning tree protocols: Per VLAN Spanning Tree (PVST), Per VLAN Spanning Tree Plus (PVST+), Rapid PVST+, and Multiple Spanning Tree (MST).
Cisco Catalyst series switches of IOS 12.2 and later versions support PVST, PVST+, Rapid PVST+, and MST. Some BPDUs of these spanning tree protocols use Cisco proprietary BPDU formats, which are different from the BPDU format defined by the IEEE. Trunk interfaces on Cisco switches enabled with PVST+ or Rapid PVST+ send Cisco proprietary BPDUs in VLANs but not VLAN 1. The source MAC address of these BPDUs is the MAC address of the trunk interface, and the destination MAC address is Cisco's reserved MAC address 01-00-0C-CC-CC-CD.
PVST can be considered as STP running in each VLAN. Each VLAN has an independent STP status and a spanning tree calculated. Although PVST does not define the instances used in MSTP, PVST can load balance traffic from different VLANs. PVST BPDUs carry VLAN information in data frames. The destination MAC address of PVST BPDUs is 01-00-0C-CC-CC-CD; therefore, PVST cannot interwork with standard IEEE spanning tree protocols.
