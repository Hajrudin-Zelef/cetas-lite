---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-interfaces-cisco-nexus-90-adc2c4ac
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-interfaces-cisco-nexus-90-adc2c4ac"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-interfaces-cisco-nexus-90-adc2c4ac.md
source_anchor: ""
source_lines: [1, 51]
sha256: 9f4bed6a4c6bb6632d95b3c078d7105fd9ee252bf5cb209b462b635d354f8815
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-interfaces-cisco-nexus-90-adc2c4ac

Bidirectional Forwarding Detection
Bidirectional Forwarding Detection (BFD) is a protocol designed to quickly identify faults in the forwarding path between two devices. BFD simplifies network profiling and planning by offering predictable reconvergence time.
BFD detects forwarding path failures across various media types, encapsulations, topologies, and routing protocols. It provides subsecond failure detection between two adjacent devices, distributing some load onto the data plane on supported modules. BFD can be less CPU-intensive than protocol hello messages.
Asynchronous mode
BFD asynchronous mode is a BFD session mode that:
- 
                                 
                                 involves the exchange of periodic control packets to monitor connectivity,
- 
                                 
                                 establishes and maintains BFD neighbor sessions, and
- 
                                 
                                 negotiates session parameters.
BFD session parameters
The table lists the BFD session parameters and the intervals.
| Table 1. BFD session parameters |  | 
|---|---|
| Session Parameters | Description | 
|---|---|
| Desired minimum transmit interval | The interval at which the device is configured to send BFD hello messages. | 
| Required minimum receive interval | The minimum interval at which the device can accept BFD hello messages from another BFD device. | 
| Detect multiplier | The number of missing BFD hello messages required to detects a fault in the forwarding path. | 
BFD neighbor workflow
The figure details the BFD neighbor sessions establishment between two routers.
The stages that establish a BFD neighbor session are:
- 
                                       
                                       The OSPF process discovers a BFD neighbor.
- 
                                       
                                       The local BFD process gets a request to start a session BFD neighbor session with the OSPF neighbor router.
- 
                                       
                                       The session is established between the BFD neighbor with the OSPF neighbor router.
BFD Detection of Failures
Once a BFD session has been established and timer negotiations are complete, BFD neighbors send BFD control packets that act in the same manner as an IGP hello protocol to detect liveliness, except at a more accelerated rate. BFD detects a failure, but the protocol must take action to bypass a failed peer.
BFD sends a failure detection notice to the BFD-enabled protocols when it detects a failure in the forwarding path. The local device can then initiate the protocol recalculation process and reduce the overall network convergence time.
The following figure shows what happens when a failure occurs in the network (1). The BFD neighbor session with the OSPF neighbor router is torn down (2). BFD notifies the local OSPF process that the BFD neighbor is no longer reachable (3). The local OSPF process tears down the OSPF neighbor relationship (4). If an alternative path is available, the routers immediately start converging on it.
| Note | Note The BFD failure detection occurs in less than a second, which is much faster than OSPF Hello messages could detect the same failure. | 
Distributed Operation
Cisco NX-OS can distribute the BFD operation to compatible modules that support BFD. This process offloads the CPU load for BFD packet processing to the individual modules that connect to the BFD neighbors. All BFD session traffic occurs on the module CPU. The module informs the supervisor when a BFD failure is detected.
BFD Echo Function
Echo packets are defined and processed only by the transmitting system. For IPv4 and IPv6, the echo packets' destination address is that of the transmitting device. It is chosen in such a way as to cause the remote system to forward the packet back to the local system. This bypasses the routing lookup on the remote system and relies on the forwarding information base (FIB) instead. BFD can use the slow timer to slow down the asynchronous session when the echo function is enabled and reduce the number of BFD control packets that are sent between two BFD neighbors. The Echo function tests only the forwarding path of the remote system by having the remote (neighbor) system loop them back, so there is less inter-packet delay variability and faster failure detection times.
Security
Cisco NX-OS uses the packet Time to Live (TTL) value to verify that the BFD packets came from an adjacent BFD peer. For all asynchronous and echo request packets, the BFD neighbor sets the TTL value to 255 and the local BFD process verifies the TTL value as 255 before processing the incoming packet. For the echo response packet, BFD sets the TTL value to 254.
You can configure SHA-1 authentication of BFD packets.
High Availability
BFD supports stateless restarts. After a reboot or supervisor switchover, Cisco NX-OS applies the running configuration and BFD immediately sends control packets to the BFD peers.
Virtualization Support
BFD supports virtual routing and forwarding instances (VRFs). VRFs exist within virtual device contexts (VDCs). By default, Cisco NX-OS places you in the default VDC and default VRF.
