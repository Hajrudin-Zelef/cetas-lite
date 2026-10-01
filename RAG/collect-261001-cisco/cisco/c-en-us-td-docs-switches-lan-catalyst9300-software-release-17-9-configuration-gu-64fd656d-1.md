---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-9-configuration-gu-64fd656d-1
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-9-configuration-gu-64fd656d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-9-configuration-gu-64fd656d.md
source_anchor: ""
source_lines: [1, 46]
sha256: f6550f2c43815e9648391c7daca81032a458d7b191a8ef063ae431e3e8d74c3f
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-9-configuration-gu-64fd656d

Information About Optional Spanning-Tree Features
The following sections provide information about Optional Spanning-Tree features:
PortFast
PortFast immediately brings an interface that is configured as an access or trunk port to the forwarding state from a blocking state, bypassing the listening and learning states.
Interfaces connected to a single workstation or server should not receive bridge protocol data units (BPDUs). An interface with PortFast enabled goes through the normal cycle of spanning-tree status changes when the switch is restarted.
You can enable this feature by enabling it on either the interface or on all nontrunking ports.
Bridge Protocol Data Unit Guard
The Bridge Protocol Data Unit (BPDU) guard feature can be globally enabled on the switch or can be enabled per port, but the feature operates with some differences.
When you enable BPDU guard at the global level on PortFast enabled ports, spanning tree shuts down ports that are in a PortFast operational state if any BPDU is received on them. In a valid configuration, PortFast enabled ports do not receive BPDUs. Receiving a BPDU on a PortFast enabled port means an invalid configuration, such as the connection of an unauthorized device, and the BPDU guard feature puts the port in the error-disabled state. When this happens, the switch shuts down the entire port on which the violation occurred.
When you enable BPDU guard at the interface level on any port without also enabling the PortFast feature, and the port receives a BPDU, it is put in the error-disabled state.
The BPDU guard feature provides a secure response to invalid configurations because you must manually put the interface back in service. Use the BPDU guard feature in a service-provider network to prevent an access port from participating in the spanning tree.
Bridge Protocol Data Unit Filtering
The BPDU filtering feature can be globally enabled on the switch or can be enabled per interface, but the feature operates with some differences.
Enabling BPDU filtering on PortFast enabled interfaces at the global level keeps those interfaces that are in a PortFast operational state from sending or receiving BPDUs. The interfaces still send a few BPDUs at link-up before the switch begins to filter outbound BPDUs. You should globally enable BPDU filtering on a switch so that hosts that are connected to these interfaces do not receive BPDUs. If a BPDU is received on a PortFast enabled interface, the interface loses its PortFast operational status, and BPDU filtering is disabled.
Enabling BPDU filtering on an interface without also enabling the PortFast feature keeps the interface from sending or receiving BPDUs.
| Caution | Enabling BPDU filtering on an interface is the same as disabling spanning tree on it and can result in spanning-tree loops. | 
You can enable the BPDU filtering feature for the entire switch or for an interface.
UplinkFast
If a switch loses connectivity, it begins using the alternate paths when the spanning tree selects a new root port. You can accelerate the choice of a new root port when a link or switch fails or when the spanning tree reconfigures itself by enabling UplinkFast. The root port transitions to the forwarding state immediately without going through the listening and learning states, as it would with the normal spanning-tree procedures.
When the spanning tree reconfigures the new root port, other interfaces flood the network with multicast packets, one for each address that was learned on the interface. You can limit these bursts of multicast traffic by reducing the max-update-rate parameter (the default for this parameter is 150 packets per second). However, if you enter zero, station-learning frames are not generated, so the spanning-tree topology converges more slowly after a loss of connectivity.
| Note | UplinkFast is most useful in wiring-closet switches at the access or edge of the network. It is not appropriate for backbone devices. This feature might not be useful for other types of applications. | 
UplinkFast provides fast convergence after a direct link failure and achieves load-balancing between redundant Layer 2 links using uplink groups. An uplink group is a set of Layer 2 interfaces (per VLAN), only one of which is forwarding at any given time. Specifically, an uplink group consists of the root port (which is forwarding) and a set of blocked ports, except for self-looping ports. The uplink group provides an alternate path in case the currently forwarding link fails.
Cross-Stack UplinkFast
Cross-Stack UplinkFast (CSUF) provides a fast spanning-tree transition (fast convergence in less than 1 second under normal network conditions) across a switch stack. During the fast transition, an alternate redundant link on the switch stack is placed in the forwarding state without causing temporary spanning-tree loops or loss of connectivity to the backbone. With this feature, you can have a redundant and resilient network in some configurations. CSUF is automatically enabled when you enable the UplinkFast feature.
CSUF might not provide a fast transition all the time; in these cases, the normal spanning-tree transition occurs, completing in 30 to 40 seconds. For more information, see Related Topics.
How Cross-Stack UplinkFast Works
Cross-Stack UplinkFast (CSUF) ensures that one link in the stack is elected as the path to the root.
When certain link loss or spanning-tree events occur (described in the following topic), the Fast Uplink Transition Protocol uses the neighbor list to send fast-transition requests to stack members.
The switch sending the fast-transition request needs to do a fast transition to the forwarding state of a port that it has chosen as the root port, and it must obtain an acknowledgment from each stack switch before performing the fast transition.
Each switch in the stack decides if the sending switch is a better choice than itself to be the stack root of this spanning-tree instance by comparing the root, cost, and bridge ID. If the sending switch is the best choice as the stack root, each switch in the stack returns an acknowledgment; otherwise, it sends a fast-transition request. The sending switch then has not received acknowledgments from all stack switches.
When acknowledgments are received from all stack switches, the Fast Uplink Transition Protocol on the sending switch immediately transitions its alternate stack-root port to the forwarding state. If acknowledgments from all stack switches are not obtained by the sending switch, the normal spanning-tree transitions (blocking, listening, learning, and forwarding) take place, and the spanning-tree topology converges at its normal rate (2 * forward-delay time + max-age time).
The Fast Uplink Transition Protocol is implemented on a per-VLAN basis and affects only one spanning-tree instance at a time.
Events That Cause Fast Convergence
Depending on the network event or failure, the CSUF fast convergence might or might not occur.
Fast convergence (less than 1 second under normal network conditions) occurs under these circumstances:
- 
                                    
                                    The stack-root port link fails. If two switches in the stack have alternate paths to the root, only one of the switches performs the fast transition.
- 
                                    
                                    The failed link, which connects the stack root to the spanning-tree root, recovers.
- 
                                    
                                    A network reconfiguration causes a new stack-root switch to be selected.
- 
                                    
