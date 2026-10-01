---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac-1
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac.md
source_anchor: ""
source_lines: [1, 54]
sha256: cd049ebf1980b258412f5df53195ecac2ad3282805e29d294e01666e52e9acf2
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac

Protocol-Independent Features
This section describes IP routing protocol-independent features that are available on switches running the Network Essentials feature set .
Distributed Cisco Express Forwarding
Information About Cisco Express Forwarding
Cisco Express Forwarding (CEF) is a Layer 3 IP switching technology used to optimize network performance. CEF implements an advanced IP look-up and forwarding algorithm to deliver maximum Layer 3 switching performance. CEF is less CPU-intensive than fast switching route caching, allowing more CPU processing power to be dedicated to packet forwarding. In a switch stack, the hardware uses distributed CEF (dCEF) in the stack. In dynamic networks, fast switching cache entries are frequently invalidated because of routing changes, which can cause traffic to be process switched using the routing table, instead of fast switched using the route cache. CEF and dCEF use the Forwarding Information Base (FIB) lookup table to perform destination-based switching of IP packets.
The two main components in CEF and dCEF are the distributed FIB and the distributed adjacency tables.
- 
                                 The FIB is similar to a routing table or information base and maintains a mirror image of the forwarding information in the IP routing table. When routing or topology changes occur in the network, the IP routing table is updated, and those changes are reflected in the FIB. The FIB maintains next-hop address information based on the information in the IP routing table. Because the FIB contains all known routes that exist in the routing table, CEF eliminates route cache maintenance, is more efficient for switching traffic, and is not affected by traffic patterns.
- 
                                 Nodes in the network are said to be adjacent if they can reach each other with a single hop across a link layer. CEF uses adjacency tables to prepend Layer 2 addressing information. The adjacency table maintains Layer 2 next-hop addresses for all FIB entries.
Because the switch or switch stack uses Application Specific Integrated Circuits (ASICs) to achieve Gigabit-speed line rate IP traffic, CEF or dCEF forwarding applies only to the software-forwarding path, that is, traffic that is forwarded by the CPU.
How to Configure Cisco Express Forwarding
CEF or distributed CEF is enabled globally by default. If for some reason it is disabled, you can re-enable it by using the ip cef or ip cef distributed global configuration command.
The default configuration is CEF or dCEF enabled on all Layer 3 interfaces. Entering the no ip route-cache cef interface configuration command disables CEF for traffic that is being forwarded by software. This command does not affect the hardware forwarding path. Disabling CEF and using the debug ip packet detail privileged EXEC command can be useful to debug software-forwarded traffic. To enable CEF on an interface for the software-forwarding path, use the ip route-cache cef interface configuration command.
| Caution | Although the no ip route-cache cef interface configuration command to disable CEF on an interface is visible in the CLI, we strongly recommend that you do not disable CEF or dCEF on interfaces except for debugging purposes. | 
To enable CEF or dCEF globally and on an interface for software-forwarded traffic if it has been disabled:
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure terminal Example:  Device# configure terminal   | Enters global configuration mode. | 
| Step 2 | ip cef Example:  Device(config)# ip cef  | Enables CEF operation on a non-stacking switch. Go to Step 4. | 
| Step 3 | ip cef distributed Example:  Device(config)# ip cef distributed  | Enables CEF operation on a active switch. | 
| Step 4 | interface interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 5 | ip route-cache cef Example:  Device(config-if)# ip route-cache cef  | Enables CEF on the interface for software-forwarded traffic. | 
| Step 6 | end Example:  Device(config-if)# end  | Returns to privileged EXEC mode. | 
| Step 7 | show ip cef Example:  Device# show ip cef  | Displays the CEF status on all interfaces. | 
| Step 8 | show cef linecard [detail] Example:  Device# show cef linecard detail  | (Optional) Displays CEF-related interface information on a non-stacking switch. | 
| Step 9 | show cef linecard [slot-number] [detail] Example:  Device# show cef linecard 5 detail  | (Optional) Displays CEF-related interface information on a switch by stack member for all switches in the stack or for the specified switch. (Optional) For slot-number , enter the stack member switch number. | 
| Step 10 | show cef interface [interface-id] Example:  Device# show cef interface gigabitethernet 1/0/1  | Displays detailed CEF information for all interfaces or the specified interface. | 
| Step 11 | show adjacency Example:  Device# show adjacency  | Displays CEF adjacency table information. | 
| Step 12 | copy running-config startup-config Example:  Device# copy running-config startup-config  | (Optional) Saves your entries in the configuration file. | 
Load-Balancing Scheme for CEF Traffic
Restrictions for Configuring a Load-Balancing Scheme for CEF Traffic
- 
                                 
                                 You must globally configure load balancing on device or device stack members in the same way.
- 
                                 
                                 Per-packet load balancing for CEF traffic is not supported.
CEF Load-Balancing Overview
CEF load balancing allows you to optimize resources by distributing traffic over multiple paths. CEF load balancing works based on a combination of source and destination packet information.
You can configure load balancing on a per-destination. Because load-balancing decisions are made on the outbound interface, load balancing must be configured on the outbound interface.
Per-Destination Load Balancing for CEF Traffic
Per-destination load balancing allows the device to use multiple paths to achieve load sharing across multiple source-destination host pairs. Packets for a given source-destination host pair are guaranteed to take the same path, even if multiple paths are available. Traffic streams destined for different pairs tend to take different paths.
Per-destination load balancing is enabled by default when you enable CEF. To use per-destination load balancing, you do not perform any additional tasks once CEF is enabled. Per-destination is the load-balancing method of choice for most situations.
Because per-destination load balancing depends on the statistical distribution of traffic, load sharing becomes more effective as the number of source-destination host pairs increases.
You can use per-destination load balancing to ensure that packets for a given host pair arrive in order. All packets intended for a certain host pair are routed over the same link (or links).
Load-Balancing Algorithms for CEF Traffic
The following load-balancing algorithms are provided for use with CEF traffic. Select a load-balancing algorithm with the ip cef load-sharing algorithm command.
- 
                                 
                                 Original algorithm—The original load-balancing algorithm produces distortions in load sharing across multiple devices because the same algorithm was used on every device. Depending on your network environment, you should select the algorithm.
- 
                                 
