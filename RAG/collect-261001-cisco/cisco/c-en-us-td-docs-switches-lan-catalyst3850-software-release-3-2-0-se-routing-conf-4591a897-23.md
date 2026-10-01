---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-23
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [1202, 1266]
sha256: 0fd28d7764202e5b50dce4aa91a64061c5d3a0746985f41c4f4f266f1388c0e3
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

The integrated IS-IS NSF Awareness feature is supported for IPv4G. The feature allows customer premises equipment (CPE) routers that are NSF-aware to help NSF-capable routers perform nonstop forwarding of packets. The local router is not necessarily performing NSF, but its awareness of NSF allows the integrity and accuracy of the routing database and link-state database on the neighboring NSF-capable router to be maintained during the switchover process.
This feature is automatically enabled and requires no configuration. For more information on this feature, see the Integrated IS-IS Nonstop Forwarding (NSF) Awareness Feature Guide.
IS-IS Global Parameters
These are some optional IS-IS global parameters that you can configure:
-  
			 You can force a default route into an IS-IS routing domain by configuring a default route controlled by a route map. You can also specify other filtering options configurable under a route map.
-  
			 You can configure the router to ignore IS-IS LSPs that are received with internal checksum errors or to purge corrupted LSPs, which causes the initiator of the LSP to regenerate it.
-  
			 You can assign passwords to areas and domains.
-  
			 You can create aggregate addresses that are represented in the routing table by a summary address (route-summarization). Routes learned from other routing protocols can also be summarized. The metric used to advertise the summary is the smallest metric of all the specific routes.
-  
			 You can set an overload bit.
-  
			 You can configure the LSP refresh interval and the maximum time that an LSP can remain in the router database without a refresh.
-  
			 You can set the throttling timers for LSP generation, shortest path first computation, and partial route computation.
-  
			 You can configure the Device to generate a log message when an IS-IS adjacency changes state (up or down).
-  
			 If a link in the network has a maximum transmission unit (MTU) size of less than 1500 bytes, you can lower the LSP MTU so that routing will still occur.
-  
			 The partition avoidance router configuration command prevents an area from becoming partitioned when full connectivity is lost among a Level1-2 border router, adjacent Level 1 routers, and end hosts.
IS-IS Interface Parameters
You can optionally configure certain interface-specific IS-IS parameters, independently from other attached routers. However, if you change some values from the defaults, such as multipliers and time intervals, it makes sense to also change them on multiple routers and interfaces. Most of the interface parameters can be configured for level 1, level 2, or both.
These are some interface level parameters you can configure:
- The default metric on the interface, which is used as a value for the IS-IS metric and assigned when there is no quality of service (QoS) routing performed.
- The hello interval (length of time between hello packets sent on the interface) or the default hello packet multiplier used on the interface to determine the hold time sent in IS-IS hello packets. The hold time determines how long a neighbor waits for another hello packet before declaring the neighbor down. This determines how quickly a failed link or neighbor is detected so that routes can be recalculated. Change the hello-multiplier in circumstances where hello packets are lost frequently and IS-IS adjacencies are failing unnecessarily. You can raise the hello multiplier and lower the hello interval correspondingly to make the hello protocol more reliable without increasing the time required to detect a link failure.
- Other time intervals: 
  - Complete sequence number PDU (CSNP) interval. CSNPs are sent by the designated router to maintain database synchronization
  - Retransmission interval. This is the time between retransmission of IS-IS LSPs for point-to-point links.
  - IS-IS LSP retransmission throttle interval. This is the maximum rate (number of milliseconds between packets) at which IS-IS LSPs are re-sent on point-to-point links This interval is different from the retransmission interval, which is the time between successive retransmissions of the same LSP
- Designated router election priority, which allows you to reduce the number of adjacencies required on a multiaccess network, which in turn reduces the amount of routing protocol traffic and the size of the topology database.
- The interface circuit type, which is the type of adjacency desired for neighbors on the specified interface
- Password authentication for the interface
How to Configure ISO CLNS Routing
Default IS-IS Configuration
| Table 12 Default IS-IS 		Configuration |  | 
|---|---|
| Feature | Default Setting | 
|---|---|
| Ignore link-state PDU (LSP) errors | Enabled. | 
| IS-IS type | Conventional IS-IS: the router acts as both a Level 1 (station) and a Level 2 (area) router. Multiarea IS-IS: the first instance of the IS-IS routing process is a Level 1-2 router. Remaining instances are Level 1 routers. | 
| Default-information originate | Disabled. | 
| Log IS-IS adjacency state changes. | Disabled. | 
| LSP generation throttling timers | Maximum interval between two consecutive occurrences: 5 seconds. Initial LSP generation delay: 50 ms. Hold time between the first and second LSP generation: 5000 ms. | 
| LSP maximum lifetime (without a refresh) | 1200 seconds (20 minutes) before t.he LSP packet is deleted. | 
| LSP refresh interval | Send LSP refreshes every 900 seconds (15 minutes). | 
| Maximum LSP packet size | 1497 bytes. | 
| NSF Awareness | Enabled. Allows Layer 3 Devicees to continue forwarding packets from a neighboring NSF-capable router during hardware or software changes. | 
| Partial route computation (PRC) throttling timers | Maximum PRC wait interval: 5 seconds. Initial PRC calculation delay after a topology change: 2000 ms. Hold time between the first and second PRC calculation: 5000 ms. | 
| Partition avoidance | Disabled. | 
| Password | No area or domain password is defined, and authentication is disabled. | 
| Set-overload-bit | Disabled. When enabled, if no arguments are entered, the overload bit is set immediately and remains set until you enter the no set-overload-bit command. | 
| Shortest path first (SPF) throttling timers | Maximum interval between consecutive SFPs: 10 seconds. Initial SFP calculation after a topology change: 5500 ms. Holdtime between the first and second SFP calculation: 5500 ms. | 
| Summary-address | Disabled. | 
Enabling IS-IS Routing
To enable IS-IS, you specify a name and NET for each routing process. You then enable IS-IS routing on the interface and specify the area for each instance of the routing process.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | clns 				  routing Example:  Device(config)# clns routing  | Enables ISO connectionless routing on the switch. | 
| Step 3 | router isis [area 				  tag] Example:  Device(config)# router isis tag1  | Enables the IS-IS routing for the specified routing process and enter IS-IS routing configuration mode. (Optional) Use the area tag argument to identify the area to which the IS-IS router is assigned. You must enter a value if you are configuring multiple IS-IS areas. The first IS-IS instance configured is Level 1-2 by default. Later instances are automatically Level 1. You can change the level of routing by using the is-type global configuration command. | 
| Step 4 | net 			 network-entity-title Example:  Device(config-router)# net 47.0004.004d.0001.0001.0c11.1111.00  | Configures the NETs for the routing process. If you are configuring multiarea IS-IS, specify a NET for each routing process. You can specify a name for a NET and for an address. | 
