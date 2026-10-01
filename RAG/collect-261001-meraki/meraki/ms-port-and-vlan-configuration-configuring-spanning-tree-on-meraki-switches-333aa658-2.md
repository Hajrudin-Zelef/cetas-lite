---
id: collect-261001-meraki/meraki/ms-port-and-vlan-configuration-configuring-spanning-tree-on-meraki-switches-333aa658-2
title: "ms-port-and-vlan-configuration-configuring-spanning-tree-on-meraki-switches-333aa658"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/ms-port-and-vlan-configuration-configuring-spanning-tree-on-meraki-switches-333aa658.md
source_anchor: ""
source_lines: [56, 90]
sha256: fd89799524cb55e79fc79a303e1c3ceb1013bb44a420866691ae7890b1947484
---

# ms-port-and-vlan-configuration-configuring-spanning-tree-on-meraki-switches-333aa658

| Loop guard | Loop guard is used to protect a network from unidirectional loops. A unidirectional link failure may stop a port in the blocking state from receiving BPDUs causing it to erroneously transition the forwarding state, creating a loop in the network. If a non-designated port with Loop Guard enabled stops receiving BPDUs, it will transition into a loop-inconsistent blocking state. In this state, the port will still process BPDUs but will not learn MAC addresses or forward traffic, thereby preventing a loop from forming. It is recommended that Loop Guard be enabled on non-designated fiber ports in physically redundant topologies. It is also recommended that Loop Guard be paired with Unidirectional Link Detection (UDLD). For more information on UDLD, check out our Unidirectional Link Detection article. | 
Configuring Portfast Trunk on a Switch Port
Portfast Trunk allows a trunk port to more quickly transition to the STP Forwarding state. To enable this feature, toggle the Portfast Trunk setting to Enabled in the Type configuration area.
Note: This feature is supported starting with IOS XE 26.1.1. To utilize portfast trunk please upgrade to IOS XE 26.1.1 or any later releases.
- 
    Before enabling this feature, make sure that there are no loops in the network between the trunk port and the connected end-device as it may lead to network instability.
- 
    Enabling this feature allows the IOS XE portfast trunk to be enabled on trunk ports, applying the spanning-tree portfast trunk option to the port(s) in question.
Save Changes to a Switch Port’s STP Configuration
Select Update at the bottom of the switch port configuration menu to save your configuration.
For deployments with mixed Meraki and none Meraki switches please refer to the section below for protocol interoperability.
Interoperability Terms and Guidelines
Root bridge
When connecting access devices including the Meraki MS series switches, it is important to first ensure that a root bridge has been properly configured. In a larger switch fabric, this is typically done on the core switch.
Spanning tree will use several values to elect the root bridge. Once elected, this information is displayed per switch in the Meraki dashboard interface under Switch > Monitor > Switches.
Automatic edge port
MS switches will automatically place all access interfaces into EDGE mode. This will cause the interface to immediately transition the port into STP forwarding mode upon linkup. The port still participates in STP. So if the port is to be a part of the loop, the port eventually transitions into STP blocking mode.
Note: MS390 running <=12.28 firmware does not support edge-port (portfast). To utilize portfast please upgrade to 12.28.1 or any later releases.
Protocol interoperability
Several other protocols have been developed including pre-standard and non-standard protocols to either improve upon or facilitate the same functionality of spanning tree. It may be necessary to connect a Meraki MS series switch to an existing infrastructure running one of these protocols.
MSTP (802.1s)
MSTP is an expansion of RSTP and adds a per-VLAN spanning tree instance to make use of better paths on each VLAN. MSTP is fully compatible with RSTP bridges, in that an MSTP BPDU can be interpreted by an RSTP bridge as an RSTP BPDU. Meraki switches do not support MSTP but are compatible and will see all MSTP instances as a single RSTP region.
Note: The MS390/C9K runs single MSTP instance by default (instance 0) and is not configurable for more than 1 instance. This aligns the platform to maintain compatibility with all other MS switches running Rapid Spanning-tree.
PVST/PVST+
This is a Cisco proprietary protocol on Catalyst/Nexus switches that is compatible with spanning tree (802.1D). It is important to note however that because PVST/PVST+ is a multi-VLAN spanning tree protocol, in order for the MS series switches to participate in spanning tree a spanning tree instance must be running on VLAN 1 of all switches and VLAN 1 is allowed on all trunk ports running PVST+ so that BPDUs are seen by the Meraki switches in the topology. Connecting an MS series to an existing switch fabric running PVST+ will force the MS series switch(es) to run in legacy mode (STP) which can increase convergence time. In this configuration, the MS series switches should never be the STP Root Bridge.
Note: When connecting a PVST+ bridge to an MS series switch, make sure both ports are configured as an 802.1Q trunk. Otherwise, the PVST+ bridge will go into a blocking state due to port inconsistency. To avoid any issues with STP, it is recommended to convert the Cisco Catalyst environment to single instance MSTP. This will ensure maximum compatibility in the STP environment.
Rapid-PVST
This is a Cisco proprietary protocol on Catalyst/Nexus switches that is compatible with spanning tree (802.1D) and RSTP (802.1w). It is important to note however that as Rapid-PVST is a multi-VLAN spanning tree protocol, MS series switches can participate in spanning tree only when a spanning tree instance is running on VLAN 1 of all switches. In addition, VLAN 1 must be allowed on all trunk ports running Rapid-PVST, so that BPDUs are seen by the Meraki switches in the topology. In this configuration, the MS series switches should never be the STP Root Bridge.
Note: To avoid unexpected issues where an MS is between two R-PVST switches (such as a Cisco Catalyst Root Bridge < MS > Catalyst switch). It is recommended to convert the topology to a MSTP on the Catalyst switches. This will avoid potential port inconsistency errors and other issues that may cause instability in the STP topology.
To find out the maximum supported stp instances on catalyst switches, use #show spanning-tree instances
C9500-12Q, C9500-24Q, C9500-40X, C9500-16X supports 300
C9500-32C, C9500-32QC, C9500-24Y4C, C9500-48Y4C supports 1000
Performance numbers for all switch models supports 4000 - 4096
Catalyst 9200 Series switches support a maximum of 128 PVST/Rapid PVST+ instances per device
A Cisco Catalyst 9300 switch or switch stack supports up to 300 spanning-tree instances when running PVST+ or Rapid PVST+ (on modern Cisco IOS XE software releases like 17.x). Older software versions or initial releases for the C9300 may have limited support to 128 instances
