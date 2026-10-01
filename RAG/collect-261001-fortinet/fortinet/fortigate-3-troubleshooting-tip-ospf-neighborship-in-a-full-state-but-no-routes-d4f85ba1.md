---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-ospf-neighborship-in-a-full-state-but-no-routes-d4f85ba1
title: "fortigate-3-troubleshooting-tip-ospf-neighborship-in-a-full-state-but-no-routes--d4f85ba1"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-ospf-neighborship-in-a-full-state-but-no-routes--d4f85ba1.md
source_anchor: ""
source_lines: [1, 4]
sha256: e22208d7688e32c0b8b04ffe8861c8d013ac8513ddcf9dbf374a95f5f5a58124
---

# fortigate-3-troubleshooting-tip-ospf-neighborship-in-a-full-state-but-no-routes--d4f85ba1

Troubleshooting Tip: OSPF Neighborship in a Full state but no routes are received or advertised to the neighbor
| Description | This article describes two scenarios where the OSPF neighbor is in a Full state, yet no routes are being advertised from the FortiGate. | 
| Scope | FortiGate. | 
| Solution | In the routing table, the neighbor's state is displayed, and if it shows a Full state, it indicates that the neighbors are fully adjacent, allowing routers to use the learned OSPF routes to forward traffic. However, there are instances where no routes are visible even when the neighbor is in a Full state, as demonstrated in the following example. Scenario 1: Single link.  In this situation, it is important to verify the network type configured for the OSPF interface on both sides, ensuring that they match. For instance, if the FortiGate is set to broadcast, the neighboring device should also be set to broadcast.  An example configuration is below:     Scenario 2: OSPF parallel link: FortiGate # get router info ospf  neighbor OSPF process 0, VRF 0: Neighbor ID     Pri   State           Dead Time   Address         Interface 10.1.1.1          1   Full/DR         00:00:38    10.182.0.57     wan1 10.1.1.1          1   Full/DR         00:00:38    10.183.0.57     wan2 In this scenario, two OSPF interfaces are configured to the same OSPF neighbor for redundancy with different network types. By default, network-type BROADCAST has a cost of 1 while network-type POINTOPOINT has a cost of 10, and FortiOS does not keep the higher cost routes in the kernel or routing database. This is intended behavior and is often used to control OSPF routes in cases of parallel OSPF links. See Technical Tip: Controlling OSPF routes in case of parallel (redundant) links: cost on OSPF interfaces. To debug routing communication, use the commands below: diagnose debug disable diagnose debug reset diagnose ip router ospf all enable diagnose ip router ospf level info diagnose debug console timestamp enable diagnose debug enable diagnose debug disable <----- To stop the debug output Related articles: |
