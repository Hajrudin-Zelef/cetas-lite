---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-7
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [369, 430]
sha256: b7314d7d44e1807585fab485f816ff2496a17387a5e337b9ed6d767115259e9a
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | ip 				  routing Example:  Device(config)# ip routing  | Enables IP routing. | 
| Step 4 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 5 | show running-config Example:  Device# show running-config    | Verifies your entries. | 
| Step 6 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Example of Enabling IP Routing
This example shows how to enable IP routingusing RIP as the routing protocol :
Device# configure terminal
Enter configuration commands, one per line.  End with CNTL/Z.
Device(config)# ip routing
Device(config)# router rip
Device(config-router)# network 10.0.0.0
Device(config-router)# end
What to Do Next
You can now set up parameters for the selected routing protocols as described in these sections:
Information About RIP
The Routing Information Protocol (RIP) is an interior gateway protocol (IGP) created for use in small, homogeneous networks. It is a distance-vector routing protocol that uses broadcast User Datagram Protocol (UDP) data packets to exchange routing information. The protocol is documented in RFC 1058. You can find detailed information about RIP in IP Routing Fundamentals, published by Cisco Press.
Using RIP, the Device sends routing information updates (advertisements) every 30 seconds. If a router does not receive an update from another router for 180 seconds or more, it marks the routes served by that router as unusable. If there is still no update after 240 seconds, the router removes all routing table entries for the non-updating router.
RIP uses hop counts to rate the value of different routes. The hop count is the number of routers that can be traversed in a route. A directly connected network has a hop count of zero; a network with a hop count of 16 is unreachable. This small range (0 to 15) makes RIP unsuitable for large networks.
If the router has a default network path, RIP advertises a route that links the router to the pseudonetwork 0.0.0.0. The 0.0.0.0 network does not exist; it is treated by RIP as a network to implement the default routing feature. The Device advertises the default network if a default was learned by RIP or if the router has a gateway of last resort and RIP is configured with a default metric. RIP sends updates to the interfaces in specified networks. If an interface’s network is not specified, it is not advertised in any RIP update.
Summary Addresses and Split Horizon
Routers connected to broadcast-type IP networks and using distance-vector routing protocols normally use the split-horizon mechanism to reduce the possibility of routing loops. Split horizon blocks information about routes from being advertised by a router on any interface from which that information originated. This feature usually optimizes communication among multiple routers, especially when links are broken.
How to Configure RIP
Default RIP Configuration
| Table 4 Default RIP 		Configuration |  | 
|---|---|
| Feature | Default Setting | 
|---|---|
| Auto summary | Enabled. | 
| Default-information originate | Disabled. | 
| Default metric | Built-in; automatic metric translations. | 
| IP RIP authentication key-chain | No authentication. Authentication mode: clear text. | 
| IP RIP triggered | Disabled | 
| IP split horizon | Varies with media. | 
| Neighbor | None defined. | 
| Network | None specified. | 
| Offset list | Disabled. | 
| Output delay | 0 milliseconds. | 
| Timers basic |  | 
| Validate-update-source | Enabled. | 
| Version | Receives RIP Version 1 and 2 packets; sends Version 1 packets. | 
Configuring Basic RIP Parameters
To configure RIP, you enable RIP routing for a network and optionally configure other parameters. On the Device, RIP configuration commands are ignored until you configure the network number.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | ip 				  routing Example:  Device(config)# ip routing  | Enables IP routing. (Required only if IP routing is disabled.) | 
| Step 4 | router 				  rip Example:  Device(config)# router rip  | Enables a RIP routing process, and enter router configuration mode. | 
| Step 5 | network  				network number Example:  Device(config)# network 12  | Associates a network with a RIP routing process. You can specify multiple network commands. RIP routing updates are sent and received through interfaces only on these networks. | 
| Step 6 | neighbor  				ip-address Example:  Device(config)# neighbor 10.2.5.1  | (Optional) Defines a neighboring router with which to exchange routing information. This step allows routing updates from RIP (normally a broadcast protocol) to reach nonbroadcast networks. | 
| Step 7 | offset-list [access-list 				  number \|  				name] {in \|  				out}  				offset [type number] Example:  Device(config)# offset-list 103 in 10  | (Optional) Applies an offset list to routing metrics to increase incoming and outgoing metrics to routes learned through RIP. You can limit the offset list with an access list or an interface. | 
| Step 8 | timers 				  basic  				update invalid holddown 				  flush Example:  Device(config)# timers basic 45 360 400 300  | (Optional) Adjusts routing protocol timers. Valid ranges for all timers are 0 to 4294967295 seconds.  | 
| Step 9 | version {1 \|  				2} Example:  Device(config)# version 2  | (Optional) Configures the switch to receive and send only RIP Version 1 or RIP Version 2 packets. By default, the switch receives Version 1 and 2 but sends only Version 1. You can also use the interface commands ip rip {send \| receive} version 1 \| 2 \| 1 2} to control what versions are used for sending and receiving on interfaces. | 
| Step 10 | no auto 				  summary Example:  Device(config)# no auto summary  | (Optional) Disables automatic summarization. By default, the switch summarizes subprefixes when crossing classful network boundaries. Disable summarization (RIP Version 2 only) to advertise subnet and host routing information to classful network boundaries. | 
| Step 11 | no 				  validate-update-source Example:  Device(config)# no validdate-update-source  | (Optional) Disables validation of the source IP address of incoming RIP routing updates. By default, the switch validates the source IP address of incoming RIP routing updates and discards the update if the source address is not valid. Under normal circumstances, disabling this feature is not recommended. However, if you have a router that is off-network and you want to receive its updates, you can use this command. | 
| Step 12 | output-delay  				delay Example:  Device(config)# output-delay 8  | (Optional) Adds interpacket delay for RIP updates sent. By default, packets in a multiple-packet RIP update have no delay added between packets. If you are sending packets to a lower-speed device, you can add an interpacket delay in the range of 8 to 50 milliseconds. | 
| Step 13 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 14 | show ip 				  protocols Example:  Device# show ip protocols  | Verifies your entries. | 
| Step 15 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
| Note | You must configure a network number for the RIP commands to take effect. | 
Configuring RIP Authentication
