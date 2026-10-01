---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046-3
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046.md
source_anchor: ""
source_lines: [214, 301]
sha256: dd143cc23c0d5468f82561a026220c2d78cc3d4889cf34deda389b1a45ebdef8
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046

has the lowest MAC address. However, because of traffic patterns, number of forwarding interfaces, or link types, Switch A
might not be the ideal root device. By increasing the priority (lowering the numerical value) of the ideal device so that
it becomes the root device, you force a spanning-tree recalculation to form a new topology with the ideal device as the root.
When the spanning-tree topology is calculated based on default parameters, the path between source and destination end stations
in a switched network might not be ideal. For instance, connecting higher-speed links to an interface that has a higher number
than the root port can cause a root-port change. The goal is to make the fastest link the root port.
For example, assume that one port on Switch B is a Gigabit Ethernet link and that another port on Switch B (a 10/100 link)
is the root port. Network traffic might be more efficient over the Gigabit Ethernet link. By changing the spanning-tree port
priority on the Gigabit Ethernet port to a higher priority (lower numerical value) than the root port, the Gigabit Ethernet
port becomes the new root port.
Spanning Tree and Redundant Connectivity
You can create a redundant backbone with spanning tree by connecting two switch interfaces to another device or to two different
devices. Spanning tree automatically disables one interface but enables it if the other one fails. If one link is high-speed
and the other is low-speed, the low-speed link is always disabled. If the speeds are the same, the port priority and port
ID are added together, and spanning tree disables the link with the highest value.
Figure 3. Spanning Tree and Redundant Connectivity
You can also create redundant links between devices by using EtherChannel groups.
Spanning-Tree Address Management
IEEE 802.1D specifies 17 multicast addresses, ranging from 0x00180C2000000 to 0x0180C2000010, to be used by different bridge
protocols. These addresses are static addresses that cannot be removed.
Regardless of the spanning-tree state, each device in the stack receives but does not forward packets that are destined for
addresses between 0x0180C2000000 and 0x0180C200000F.
If spanning tree is enabled, the CPU on the switch or on each switch in the stack receives packets that are destined for 0x0180C2000000
and 0x0180C2000010. If spanning tree is disabled, the switch or each switch in the stack forwards those packets as unknown
multicast addresses.
Accelerated Aging to Retain Connectivity
The default for aging dynamic addresses is 5 minutes, the default setting of the mac address-table aging-time global configuration command. However, a spanning-tree reconfiguration can cause many station locations to change. Because
these stations could be unreachable for 5 minutes or more during a reconfiguration, the address-aging time is accelerated
so that station addresses can be dropped from the address table and then relearned. The accelerated aging is the same as the
forward-delay parameter value (spanning-tree vlanvlan-idforward-timeseconds global configuration command) when the spanning tree reconfigures.
Because each VLAN is a separate spanning-tree instance, the switch accelerates aging on a per-VLAN basis. A spanning-tree
reconfiguration on one VLAN can cause the dynamic addresses that are learned on that VLAN to be subject to accelerated aging.
Dynamic addresses on other VLANs can be unaffected and remain subject to the aging interval entered for the switch.
Spanning-Tree Modes and Protocols
The device supports these spanning-tree modes and protocols:
PVST+—This spanning-tree mode is based on the IEEE 802.1D standard and Cisco proprietary extensions. The PVST+ runs on each
VLAN on the device up to the maximum supported, ensuring that each has a loop-free path through the network.
The PVST+ provides Layer 2 load-balancing for the VLAN on which it runs. You can create different logical topologies by using
the VLANs on your network to ensure that all of your links are used but that no one link is oversubscribed. Each instance
of PVST+ on a VLAN has a single root switch. This root switch propagates the spanning-tree information that is associated
with that VLAN to all other devices in the network. Because each device has the same information about the network, this process
ensures that the network topology is maintained.
Rapid PVST+—Rapid PVST+ is the default STP mode on your device. This spanning-tree mode is the same as PVST+ except that is
uses a rapid convergence based on the IEEE 802.1w standard. To provide rapid convergence, the Rapid PVST+ immediately deletes
dynamically learned MAC address entries on a per-port basis upon receiving a topology change. By contrast, PVST+ uses a short
aging time for dynamically learned MAC address entries.
Rapid PVST+ uses the same configuration as PVST+ (except where noted), and the device needs only minimal extra configuration.
The benefit of Rapid PVST+ is that you can migrate a large PVST+ install base to Rapid PVST+ without having to learn the complexities
of the Multiple Spanning Tree Protocol (MSTP) configuration and without having to reprovision your network. In Rapid PVST+
mode, each VLAN runs its own spanning-tree instance up to the maximum supported.
MSTP—This spanning-tree mode is based on the IEEE 802.1s standard. You can map multiple VLANs to the same spanning-tree instance,
which reduces the number of spanning-tree instances that are required to support many VLANs. The MSTP runs on top of the RSTP
(based on IEEE 802.1w), which provides for rapid convergence of the spanning tree by eliminating the forward delay and by
quickly transitioning root ports and designated ports to the forwarding state. In a switch stack, the cross-stack rapid transition (CSRT) feature performs the same function as RSTP. You cannot run MSTP
without RSTP or CSRT.
Supported Spanning-Tree Instances
Starting with Cisco IOS XE Amsterdam 17.2.1 release, in PVST+ or Rapid PVST+ mode, the device or device stack supports up to 300 spanning-tree instances.
On the Cisco Catalyst 9500 Series Switches (C9500-12Q, C9500-16X, C9500-24Q, C9500-40X models), in MSTP mode, the device or
device stack supports up to 65 MST instances. The number of VLANs that can be mapped to a particular MST instance is unlimited.
On the Cisco Catalyst 9500 Series High Performance Switches (C9500-32C, C9500-32QC, C9500-48Y4C, C9500-24Y4C models), in MSTP
mode, the device or device stack supports up to 64 MST instances. The number of VLANs that can be mapped to a particular MST
instance is 1000. In PVST+ or Rapid PVST+ mode, the device or device stack supports up to 1000 spanning-tree instances.
Spanning-Tree Interoperability and Backward Compatibility
In a mixed MSTP and PVST+ network, the common spanning-tree (CST) root must be inside the MST backbone, and a PVST+ device
cannot connect to multiple MST regions.
When a network contains devices running Rapid PVST+ and devices running PVST+, we recommend that the Rapid PVST+ devices and
PVST+ devices be configured for different spanning-tree instances. In the Rapid PVST+ spanning-tree instances, the root switch
must be a Rapid PVST+ device. In the PVST+ instances, the root switch must be a PVST+ device. The PVST+ devices should be
at the edge of the network.
All stack members run the same version of spanning tree (all PVST+, all Rapid PVST+, or all MSTP).
Table 2. PVST+, MSTP, and Rapid-PVST+ Interoperability and Compatibility
PVST+
MSTP
Rapid PVST+
PVST+
Yes
Yes (with restrictions)
Yes (reverts to PVST+)
MSTP
Yes (with restrictions)
Yes
Yes (reverts to PVST+)
Rapid PVST+
Yes (reverts to PVST+)
Yes (reverts to PVST+)
Yes
Spanning Tree Protocols and IEEE 802.1Q Trunks
The IEEE 802.1Q standard for VLAN trunks imposes some limitations on the spanning-tree strategy for a network. The standard
