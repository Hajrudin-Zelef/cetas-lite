---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6-3
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["preemption"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6.md
source_anchor: ""
source_lines: [41, 94]
sha256: e0c06e98c1fe3c44dcba66f57cd8304160857cf84469f4c6051a82d82c07ce7f
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6

An HSRP IPv6 group has a virtual MAC address that is derived from the HSRP group number and a virtual IPv6 link-local address that is derived, by default, from the HSRP virtual MAC address. The default virtual MAC address for an HSRP IPv6 group is always used to form the virtual IPv6 link-local address, regardless of the actual virtual MAC address used by the group.
Table 17-1 shows the MAC and IP addresses used for IPv6 neighbor discovery packets and HSRP packets.
|  |  |  |  |  | 
|---|---|---|---|---|
| Neighbor solicitation (NS) | Interface MAC address | Interface IPv6 address | — | Interface MAC address | 
| Router solicitation (RS) | Interface MAC address | Interface IPv6 address | — | Interface MAC address | 
| Neighbor advertisement (NA) | Interface MAC address | Interface IPv6 address | Virtual IPv6 address | HSRP virtual MAC address | 
| Route advertisement (RA) | Interface MAC address | Virtual IPv6 address | — | HSRP virtual MAC address | 
| HSRP (inactive) | Interface MAC address | Interface IPv6 address | — | — | 
| HSRP (active) | Virtual MAC address | Interface IPv6 address | — | — | 
HSRP does not add IPv6 link-local addresses to the Unicast Routing Information Base (URIB). Link-local addresses have no secondary virtual IP addresses.
For global unicast addresses, HSRP adds the virtual IPv6 address to the URIB and IPv6.
HSRP Authentication
HSRP message digest 5 (MD5) algorithm authentication protects against HSRP-spoofing software and uses the industry-standard MD5 algorithm for improved reliability and security. HSRP includes the IPv4 or IPv6 address in the authentication TLVs.
HSRP Messages
Routers that are configured with HSRP exchange the following three types of multicast messages:
- Hello—The hello message conveys the HSRP priority and state information of the router to other HSRP routers.
- Coup—When a standby router wants to assume the function of the active router, it sends a coup message.
- Resign—A router that is the active router sends this message when it is about to shut down or when a router that has a higher priority sends a hello or coup message.
HSRP Load Sharing
HSRP allows you to configure multiple groups on an interface. You can configure two overlapping IPv4 HSRP groups to load share traffic from the connected hosts while providing the default router redundancy expected from HSRP. Figure 17-2 shows an example of a load-sharing HSRP IPv4 configuration.
Figure 17-2 shows two routers A and B and two HSRP groups. Router A is the active router for group A but is the standby router for group B. Similarly, router B is the active router for group B and the standby router for group A. If both routers remain active, HSRP load balances the traffic from the hosts across both routers. If either router fails, the remaining router continues to process traffic for both hosts.
Note HSRP for IPv6 load balances by default. If two HSRP IPv6 groups are on the subnet, hosts learn of both groups from their router advertisements and choose to use one so that the load is shared between the advertised routers.
Object Tracking and HSRP
You can use object tracking to modify the priority of an HSRP interface based on the operational state of another interface. Object tracking allows you to route to a standby router if the interface to the main network fails.
Two objects that you can track are the line protocol state of an interface or the reachability of an IP route. If the specified object goes down, Cisco NX-OS reduces the HSRP priority by the configured amount. For more information, see the “Configuring HSRP Object Tracking” section.
vPC and HSRP
HSRP interoperates with virtual port channels (vPCs). vPCs allow links that are physically connected to two different Cisco Nexus 9000 Series devices to appear as a single port channel by a third device. See the Cisco Nexus 9000 Series NX-OS Layer 2 Switching Configuration Guide for more information on vPCs.
vPC forwards traffic through both the active HSRP router and the standby HSRP router. For more information, see the “Configuring the HSRP Priority” section and the “Configuration Examples for HSRP” section.
Note HSRP active can be distributed on both the primary and secondary vPC peers for different SVIs.
vPC Peer Gateway and HSRP
Some third-party devices can ignore the HSRP virtual MAC address and instead use the source MAC address of an HSRP router. In a vPC environment, the packets that use this source MAC address might be sent across the vPC peer link, causing a potential dropped packet. Configure the vPC peer gateway to enable the HSRP routers to directly handle packets sent to the local vPC peer MAC address and the remote vPC peer MAC address, as well as the HSRP virtual MAC address. See the Cisco Nexus 9000 Series NX-OS Layer 2 Switching Configuration Guide for more information on the vPC peer gateway.
BFD
This feature supports bidirectional forwarding detection (BFD). BFD is a detection protocol that provides fast-forwarding and path-failure detection times. BFD provides subsecond failure detection between two adjacent devices and can be less CPU-intensive than protocol hello messages because some of the BFD load can be distributed onto the data plane on supported modules. See the Cisco Nexus 9000 Series NX-OS Interfaces Configuration Guide for more information.
High Availability and Extended Nonstop Forwarding
HSRP supports stateful restarts and stateful switchovers. A stateful restart occurs when the HSRP process fails and is restarted. A stateful switchover occurs when the active supervisor switches to the standby supervisor. Cisco NX-OS applies the run-time configuration after the switchover.
If HSRP hold timers are configured for short time periods, these timers might expire during a controlled switchover. HSRP supports extended nonstop forwarding (NSF) to temporarily extend these HSRP hold timers during a controlled switchover.
With extended NSF configured, HSRP sends hello messages with the extended timers. HSRP peers update their hold timers with these new values. The extended timers prevent unnecessary HSRP state changes during the switchover. After the switchover, HSRP restores the hold timers to their original configured values. If the switchover fails, HSRP restores the hold timers after the extended hold timer values expire.
See the “Configuring Extended Hold Timers for HSRP” section for more information.
Virtualization Support
HSRP supports virtual routing and forwarding (VRF) instances.
Licensing Requirements for HSRP
The following table shows the licensing requirements for this feature:
Prerequisites for HSRP
Guidelines and Limitations for HSRP
HSRP has the following configuration guidelines and limitations:
- You must configure an IP address for the interface that you configure HSRP on and enable that interface before HSRP becomes active.
- You must configure HSRP version 2 when you configure an IPv6 interface for HSRP.
- For IPv4, the virtual IP address must be in the same subnet as the interface IP address.
- We recommend that you do not configure more than one first-hop redundancy protocol on the same interface.
- HSRP version 2 does not interoperate with HSRP version 1. An interface cannot operate both version 1 and version 2 because both versions are mutually exclusive. However, the different versions can be run on different physical interfaces of the same router.
- You cannot change from version 2 to version 1 if you have configured groups above the group number range allowed for version 1 (0 to 255).
- HSRP for IPv4 is supported with BFD. HSRP for IPv6 is not supported with BFD.
- If HSRP IPv4 and IPv6 use the same virtual MAC address on an SVI, the HSRP state should be the same for both HSRP IPv4 and IPv6. The priority and preemption should be configured to result in the same state after failovers.
