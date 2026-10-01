---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-22
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [1151, 1201]
sha256: f4e59b1501592d8b27244b7b7ffe191df1f902e82ce63829905d16dd49d78d6c
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

| show ip bgp community [community-number] [exact] | Displays routes that belong to the specified communities. | 
| show ip bgp community-list community-list-number [exact-match] | Displays routes that are permitted by the community list. | 
| show ip bgp filter-list access-list-number | Displays routes that are matched by the specified AS path access list. | 
| show ip bgp inconsistent-as | Displays the routes with inconsistent originating autonomous systems. | 
| show ip bgp regexp regular-expression | Displays the routes that have an AS path that matches the specified regular expression entered on the command line. | 
| show ip bgp | Displays the contents of the BGP routing table. | 
| show ip bgp neighbors [address] | Displays detailed information on the BGP and TCP connections to individual neighbors. | 
| show ip bgp neighbors [address] [advertised-routes \| dampened-routes \| flap-statistics \| paths regular-expression \| received-routes \| routes] | Displays routes learned from a particular BGP neighbor. | 
| show ip bgp paths | Displays all BGP paths in the database. | 
| show ip bgp peer-group [tag] [summary] | Displays information about BGP peer groups. | 
| show ip bgp summary | Displays the status of all BGP connections. | 
The bgp log-neighbor changes command is enabled by default. It allows to log messages that are generated when a BGP neighbor resets, comes up, or goes down.
Configuration Examples for BGP
Example: Configuring BGP on Routers
Device(config)# router bgp 100
Device(config-router)# neighbor 129.213.1.1 remote-as 200
Device(config)# router bgp 200
Device(config-router)# neighbor 129.213.1.2 remote-as 100
Device(config-router)# neighbor 175.220.1.2 remote-as 200
Device(config)# router bgp 200
Device(config-router)# neighbor 175.220.212.1 remote-as 200
Device(config-router)# neighbor 192.208.10.1 remote-as 300
Device(config)# router bgp 300
Device(config-router)# neighbor 192.208.10.2 remote-as 200
To verify that BGP peers are running, use the show ip bgp neighbors privileged EXEC command. This is the output of this command on Router A:
Device# show ip bgp neighbors
BGP neighbor is 129.213.1.1, remote AS 200, external link
 BGP version 4, remote router ID 175.220.212.1
 BGP state = established, table version = 3, up for 0:10:59
 Last read 0:00:29, hold time is 180, keepalive interval is 60 seconds
 Minimum time between advertisement runs is 30 seconds
 Received 2828 messages, 0 notifications, 0 in queue
 Sent 2826 messages, 0 notifications, 0 in queue
 Connections established 11; dropped 10
Anything other than state = established means that the peers are not running. The remote router ID is the highest IP address on that router (or the highest loopback interface). Each time the table is updated with new information, the table version number increments. A table version number that continually increments means that a route is flapping, causing continual routing updates.
For exterior protocols, a reference to an IP network from the network router configuration command controls only which networks are advertised. This is in contrast to Interior Gateway Protocols (IGPs), such as EIGRP, which also use the network command to specify where to send updates.
For detailed descriptions of BGP configuration, see the “IP Routing Protocols” part of the Cisco IOS IP Configuration Guide, Release 12.4. For details about specific commands, see the Cisco IOS IP Command Reference, Volume 2 of 3: Routing Protocols, Release 12.4.
Information About ISO CLNS Routing
Connectionless Routing
The International Organization for Standardization (ISO) Connectionless Network Service (CLNS) protocol is a standard for the network layer of the Open System Interconnection (OSI) model. Addresses in the ISO network architecture are referred to as network service access point (NSAP) addresses and network entity titles (NETs). Each node in an OSI network has one or more NETs. In addition, each node has many NSAP addresses.
When you enable connectionless routing on the Device by using the clns routing global configuration command, the Device makes only forwarding decisions, with no routing-related functionality. For dynamic routing, you must also enable a routing protocol. The Device supports the Intermediate System-to-Intermediate System (IS-IS) dynamic routing protocol that is based on the OSI routing protocol for ISO CLNS networks.
When dynamically routing, you use IS-IS. This routing protocol supports the concept of areas. Within an area, all routers know how to reach all the system IDs. Between areas, routers know how to reach the proper area. IS-IS supports two levels of routing: station routing (within an area) and area routing (between areas).
The key difference between the ISO IGRP and IS-IS NSAP addressing schemes is in the definition of area addresses. Both use the system ID for Level 1 routing (routing within an area). However, they differ in the way addresses are specified for area routing. An ISO IGRP NSAP address includes three separate fields for routing: the domain, area, and system ID. An IS-IS address includes two fields: a single continuous area field (comprising the domain and area fields) and the system ID.
| Note | For more detailed information about ISO CLNS, see the Cisco IOS Apollo Domain, Banyan VINES, DECnet, ISO CLNS and XNS Configuration Guide, Release 12.4. For complete syntax and usage information for the commands used in this chapter, see the Cisco IOS Apollo Domain, Banyan VINES, DECnet, ISO CLNS and XNS Command Reference, Release 12.4, use the IOS command reference master index, or search online. | 
IS-IS Dynamic Routing
IS-IS is an ISO dynamic routing protocol (described in ISO 105890). Unlike other routing protocols, enabling IS-IS requires that you create an IS-IS routing process and assign it to a specific interface, rather than to a network. You can specify more than one IS-IS routing process per Layer 3 Device or router by using the multiarea IS-IS configuration syntax. You then configure the parameters for each instance of the IS-IS routing process.
Small IS-IS networks are built as a single area that includes all the routers in the network. As the network grows larger, it is usually reorganized into a backbone area made up of the connected set of all Level 2 routers from all areas, which is in turn connected to local areas. Within a local area, routers know how to reach all system IDs. Between areas, routers know how to reach the backbone, and the backbone routers know how to reach other areas.
Routers establish Level 1 adjacencies to perform routing within a local area (station routing). Routers establish Level 2 adjacencies to perform routing between Level 1 areas (area routing).
A single Cisco router can participate in routing in up to 29 areas and can perform Level 2 routing in the backbone. In general, each routing process corresponds to an area. By default, the first instance of the routing process configured performs both Level 1and Level 2 routing. You can configure additional router instances, which are automatically treated as Level 1 areas. You must configure the parameters for each instance of the IS-IS routing process individually.
For IS-IS multiarea routing, you can configure only one process to perform Level 2 routing, although you can define up to 29 Level 1 areas for each Cisco unit. If Level 2 routing is configured on any process, all additional processes are automatically configured as Level 1. You can configure this process to perform Level 1 routing at the same time. If Level 2 routing is not desired for a router instance, remove the Level 2 capability using the is-type global configuration command. Use the is-type command also to configure a different router instance as a Level 2 router.
Nonstop Forwarding Awareness
