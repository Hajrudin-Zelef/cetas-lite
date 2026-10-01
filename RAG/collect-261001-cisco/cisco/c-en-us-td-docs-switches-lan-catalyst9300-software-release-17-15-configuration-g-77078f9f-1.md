---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-15-configuration-g-77078f9f-1
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-15-configuration-g-77078f9f"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "license", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-15-configuration-g-77078f9f.md
source_anchor: ""
source_lines: [1, 67]
sha256: d8df2d8d06b90ba9d2e55c78b90ba19a4cad495fb4e5f9f75ec125520acb0444
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-15-configuration-g-77078f9f

Information About OSPF
OSPF is an Interior Gateway Protocol (IGP) designed expressly for IP networks, supporting IP subnetting and tagging of externally derived routing information. OSPF also allows packet authentication and uses IP multicast when sending and receiving packets. The Cisco implementation supports RFC 1253, OSPF management information base (MIB).
The Cisco implementation conforms to the OSPF Version 2 specifications with these key features:
- 
                              				
                              Definition of stub areas is supported.
- 
                              				
                              Routes learned through any IP routing protocol can be redistributed into another IP routing protocol. At the intradomain level, this means that OSPF can import routes learned through EIGRP and RIP. OSPF routes can also be exported into RIP.
- 
                              				
                              Plain text and MD5 authentication among neighboring routers within an area is supported.
- 
                              				
                              Configurable routing interface parameters include interface output cost, retransmission interval, interface transmit delay, router priority, router dead and hello intervals, and authentication key.
- 
                              				
                              Virtual links are supported.
- 
                              				
                              Not-so-stubby-areas (NSSAs) per RFC 1587are supported.
OSPF typically requires coordination among many internal routers, area border routers (ABRs) connected to multiple areas, and autonomous system boundary routers (ASBRs). The minimum configuration would use all default parameter values, no authentication, and interfaces assigned to areas. If you customize your environment, you must ensure coordinated configuration of all routers.
OSPF for IPv6
The switch supports Open Shortest Path First (OSPF) for IPv6, a link-state protocol for IP.
For configuring OSPF for IPv6, see the Configuring OSPF for IPv6 section.
For more information, see Cisco IOS IPv6 Configuration Library on Cisco.com.
OSPF Nonstop Forwarding
The switch or switch stack supports two levels of nonstop forwarding (NSF):
OSPF NSF Awareness
The Network Advantage license supports OSPF NSF Awareness for IPv4. When the neighboring router is NSF-capable, the Layer 3 device continues to forward packets from the neighboring router during the interval between the primary Route Processor (RP) in a router crashing and the backup RP taking over, or while the primary RP is manually reloaded for a non-disruptive software upgrade.
This feature cannot be disabled.
OSPF NSF Capability
The Network Advantage license supports the OSPFv2 NSF IETF format in addition to the OSPFv2 NSF Cisco format that is supported in earlier releases. For information about this feature, see : NSF—OSPF (RFC 3623 OSPF Graceful Restart).
The Network Advantage license also supports OSPF NSF-capable routing for IPv4 for better convergence and lower traffic loss following a stack's active switch change.
When an active switch change occurs in an OSPF NSF-capable stack, the new active switch must do two things to resynchronize its link-state database with its OSPF neighbors:
- 
                                    				
                                    Release the available OSPF neighbors on the network without resetting the neighbor relationship.
- 
                                    				
                                    Reacquire the contents of the link-state database for the network.
After an active switch change, the new active switch sends an OSPF NSF signal to neighboring NSF-aware devices. A device recognizes this signal to mean that it should not reset the neighbor relationship with the stack. As the NSF-capable active switch receives signals from other routes on the network, it begins to rebuild its neighbor list.
When the neighbor relationships are reestablished, the NSF-capable active switch resynchronizes its database with its NSF-aware neighbors, and routing information is exchanged between the OSPF neighbors. The new active switch uses this routing information to remove stale routes, to update the routing information database (RIB), and to update the forwarding information base (FIB) with the new information. The OSPF protocols then fully converge.
| Note | OSPF NSF requires that all neighbor networking devices be NSF-aware. If an NSF-capable router discovers non-NSF aware neighbors on a network segment, it disables NSF capabilities for that segment. Other network segments where all devices are NSF-aware or NSF-capable continue to provide NSF capabilities. | 
Use the nsf OSPF routing configuration command to enable OSPF NSF routing. Use the show ip ospf privileged EXEC command to verify that it is enabled.
OSPF Area Parameters
You can optionally configure several OSPF area parameters. These parameters include authentication for password-based protection against unauthorized access to an area, stub areas, and not-so-stubby-areas (NSSAs). Stub areas are areas into which information on external routes is not sent. Instead, the area border router (ABR) generates a default external route into the stub area for destinations outside the autonomous system (AS). An NSSA does not flood all LSAs from the core into the area, but can import AS external routes within the area by redistribution.
Route summarization is the consolidation of advertised addresses into a single summary route to be advertised by other areas. If network numbers are contiguous, you can use the area range router configuration command to configure the ABR to advertise a summary route that covers all networks in the range.
Other OSPF Parameters
You can optionally configure other OSPF parameters in router configuration mode.
- 
                                 				
                                 Route summarization: When redistributing routes from other protocols. Each route is advertised individually in an external LSA. To help decrease the size of the OSPF link state database, you can use the summary-address router configuration command to advertise a single router for all the redistributed routes included in a specified network address and mask.
- 
                                 				
                                 Virtual links: In OSPF, all areas must be connected to a backbone area. You can establish a virtual link in case of a backbone-continuity break by configuring two Area Border Routers as endpoints of a virtual link. Configuration information includes the identity of the other virtual endpoint (the other ABR) and the nonbackbone link that the two routers have in common (the transit area). Virtual links cannot be configured through a stub area.
- 
                                 				
                                 Default route: When you specifically configure redistribution of routes into an OSPF routing domain, the route automatically becomes an autonomous system boundary router (ASBR). You can force the ASBR to generate a default route into the OSPF routing domain.
- 
                                 				
                                 Domain Name Server (DNS) names for use in all OSPF show privileged EXEC command displays makes it easier to identify a router than displaying it by router ID or neighbor ID.
- 
                                 				
                                 Default Metrics: OSPF calculates the OSPF metric for an interface according to the bandwidth of the interface. The metric is calculated as ref-bw divided by bandwidth, where ref is 10 by default, and bandwidth (bw ) is specified by the bandwidth interface configuration command. For multiple links with high bandwidth, you can specify a larger number to differentiate the cost on those links.
- 
                                 				
