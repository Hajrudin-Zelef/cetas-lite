---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-10
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [538, 589]
sha256: 19963cb125f55f6f098cdd7cac1a32ee39a418a9d08fb80d60ccbd7e4b7fff17
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

| Interface parameters | Cost: 1. Retransmit interval: 5 seconds. Transmit delay: 1 second. Priority: 1. Hello interval: 10 seconds. Dead interval: 4 times the hello interval. No authentication. No password specified. MD5 authentication disabled. | 
| Area | Authentication type: 0 (no authentication). Default cost: 1. Range: Disabled. Stub: No stub area defined. NSSA: No NSSA area defined. | 
| Auto cost | 100 Mb/s. | 
| Default-information originate | Disabled. When enabled, the default metric setting is 10, and the external route type default is Type 2. | 
| Default metric | Built-in, automatic metric translation, as appropriate for each routing protocol. | 
| Distance OSPF | dist1 (all routes within an area): 110. dist2 (all routes from one area to another): 110. and dist3 (routes from other routing domains): 110. | 
| OSPF database filter | Disabled. All outgoing link-state advertisements (LSAs) are flooded to the interface. | 
| IP OSPF name lookup | Disabled. | 
| Log adjacency changes | Enabled. | 
| Neighbor | None specified. | 
| Neighbor database filter | Disabled. All outgoing LSAs are flooded to the neighbor. | 
| Network area | Disabled. | 
| Nonstop Forwarding (NSF) awareness | Enabled. Allows Layer 3 switches to continue forwarding packets from a neighboring NSF-capable router during hardware or software changes. | 
| NSF capability | Disabled. | 
| Router ID | No OSPF routing process defined. | 
| Summary address | Disabled. | 
| Timers LSA group pacing | 240 seconds. | 
| Timers shortest path first (spf) | spf delay: 5 seconds.; spf-holdtime: 10 seconds. | 
| Virtual link | No area ID or router ID defined. Hello interval: 10 seconds. Retransmit interval: 5 seconds. Transmit delay: 1 second. Dead interval: 40 seconds. Authentication key: no key predefined. Message-digest key (MD5): no key predefined. | 
| Note | The switch stack supports OSPF NSF-capable routing for IPv4. | 
Configuring Basic OSPF Parameters
To enable OSPF, create an OSPF routing process, specify the range of IP addresses to associate with the routing process, and assign area IDs to be associated with that range. For switches running the IP services image, you can configure either the Cisco OSPFv2 NSF format or the IETF OSPFv2 NSF format.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | router 				  ospf process-id Example:  Device(config)# router ospf 15  | Enables OSPF routing, and enter router configuration mode. The process ID is an internally used identification parameter that is locally assigned and can be any positive integer. Each OSPF routing process has a unique value. | 
| Step 3 | nsf cisco [enforce global] Example:  Device(config)# nsf cisco enforce global  | (Optional) Enables Cisco NSF operations for OSPF. The enforce global keyword cancels NSF restart when non-NSF-aware neighboring networking devices are detected. | 
| Step 4 | nsf ietf [restart-interval  				seconds] Example:  Device(config)# nsf ietf restart-interval 60  | (Optional) Enables IETF NSF operations for OSPF. The restart-interval keyword specifies the length of the graceful restart interval, in seconds. The range is from 1 to 1800. The default is 120. | 
| Step 5 | network address wildcard-mask  			 area area-id Example:  Device(config)# network 10.1.1.1 255.240.0.0 area 20  | Define an interface on which OSPF runs and the area ID for that interface. You can use the wildcard-mask to use a single command to define one or more multiple interfaces to be associated with a specific OSPF area. The area ID can be a decimal value or an IP address. | 
| Step 6 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 7 | show ip 				  protocols Example:  Device# show ip protocols  | Verifies your entries. | 
| Step 8 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
| Note | OSPF for Routed Access supports only one OSPFv2 and one OSPFv3 instance with a maximum number of 200 dynamically learned routes. | 
| Note | Enter the command in Step 3 or Step 4, and go to Step 5. | 
Configuring OSPF Interfaces
You can use the ip ospf interface configuration commands to modify interface-specific OSPF parameters. You are not required to modify any of these parameters, but some interface parameters (hello interval, dead interval, and authentication key) must be consistent across all routers in an attached network. If you modify these parameters, be sure all routers in the network have compatible values.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | interface interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 3 | ip 				  ospf cost Example:  Device(config-if)# ip ospf 8  | (Optional) Explicitly specifies the cost of sending a packet on the interface. | 
| Step 4 | ip ospf 				  retransmit-interval seconds Example:  Device(config-if)# ip ospf transmit-interval 10  | (Optional) Specifies the number of seconds between link state advertisement transmissions. The range is 1 to 65535 seconds. The default is 5 seconds. | 
| Step 5 | ip ospf 				  transmit-delay seconds Example:  Device(config-if)# ip ospf transmit-delay 2  | (Optional) Sets the estimated number of seconds to wait before sending a link state update packet. The range is 1 to 65535 seconds. The default is 1 second. | 
| Step 6 | ip ospf 				  priority number Example:  Device(config-if)# ip ospf priority 5  | (Optional) Sets priority to help find the OSPF designated router for a network. The range is from 0 to 255. The default is 1. | 
| Step 7 | ip ospf 				  hello-interval seconds Example:  Device(config-if)# ip ospf hello-interval 12  | (Optional) Sets the number of seconds between hello packets sent on an OSPF interface. The value must be the same for all nodes on a network. The range is 1 to 65535 seconds. The default is 10 seconds. | 
| Step 8 | ip ospf 				  dead-interval seconds Example:  Device(config-if)# ip ospf dead-interval 8  | (Optional) Sets the number of seconds after the last device hello packet was seen before its neighbors declare the OSPF router to be down. The value must be the same for all nodes on a network. The range is 1 to 65535 seconds. The default is 4 times the hello interval. | 
| Step 9 | ip ospf 				  authentication-key key Example:  Device(config-if)# ip ospf authentication-key password  | (Optional) Assign a password to be used by neighboring OSPF routers. The password can be any string of keyboard-entered characters up to 8 bytes in length. All neighboring routers on the same network must have the same password to exchange OSPF information. | 
| Step 10 | ip ospf message 				  digest-key keyid  			 md5 key Example:  Device(config-if)# ip ospf message digest-key 16 md5 your1pass  | (Optional) Enables MDS authentication. | 
| Step 11 | ip ospf 				  database-filter all out Example:  Device(config-if)# ip ospf database-filter all out  | (Optional) Block flooding of OSPF LSA packets to the interface. By default, OSPF floods new LSAs over all interfaces in the same area, except the interface on which the LSA arrives. | 
| Step 12 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 13 | show ip ospf interface [interface-name] Example:  Device# show ip ospf interface  | Displays OSPF-related interface information. | 
| Step 14 | show ip ospf 				  neighbor detail Example:  Device# show ip ospf neighbor detail  | Displays NSF awareness status of neighbor switch. The output matches one of these examples: | 
