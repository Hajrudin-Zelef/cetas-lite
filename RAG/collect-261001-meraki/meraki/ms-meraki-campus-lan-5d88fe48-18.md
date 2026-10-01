---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-18
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution", "latency"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [726, 789]
sha256: 2bf61b2df42dd6c670918a9aee60bada0a1bd11b36cb9086a6e1b4b236bdee80
---

# ms-meraki-campus-lan-5d88fe48

  - Normal Areas (LSA types 1,2,3,4 and 5)
  - Stub Areas (LSA types 1,2, and 3)
  - Not-So-Stubby Areas NSSA (LSA types 1,2 and 7)
The OSPF area IDs must be consistent on all OSPF peers
- It is recommended to keep your backbone area manageable in terms of size (e.g. maximum 30 routers) for better performance and convergence
- It is recommended to design your backbone area such that you have clear demarcation from core to access (e.g. backbone area covers core and distribution and access is segregated into multiple Stub/NNSA areas) so basically making your aggregation switches ABRs
- It is recommended to summarize routes where possible for instance at the edge of your backbone area (e.g. Hybrid Campus LAN with Cat9500 Layer 3 Core)
- It is recommended to use route filtering in the backbone area to avoid asymetrical routing (e.g. Hybrid Campus LAN with Cat9500 Core)
- The default cost is 1, but can be increased to give lower priority
- Choose passive on interfaces that do not require forming OSPF peerings
- We recommend leaving the “hello” and “dead” timers to a default of 10s and 40s respectively (If more aggressive timers are required, ensure adequate testing is performed)
The value configured for timers must be identical between all participating OSPF neighbors. If introducing an MS switch to an existing OSPF topology, be sure to reference the existing configuration
- Ensure all areas are directly attached to the backbone Area 0 (Virtual links are not supported)
- Configure a Router ID for ease of management
- Meraki Router Priority is 1 (this cannot be adjusted)
In a hybrid Campus LAN, it is recommended to set priorities on the Catalyst switches. If OSPF peering is happening over LACP channels, it is recommended to set the LACP mode on Catalyst switches to active mode
- Create a Transit VLAN for OSPF peering between access and distribution (or use management VLAN) and set OSPF to passive on all other interfaces (this will reduce load on CPU)
- Configure MD5 authentication for security purposes
Please note that routing protocol redistribution is not supported on MS platforms. As such, redistribution can be implemented on higher layers (e.g. Catalyst distribution or core). Virtual links are not supported on MS platforms
Layer 3 Interfaces (SVIs)
General Guidance
- In order to route traffic between VLANs, routed interfaces must be configured.
- Only VLANs with a routed interface configured will be able to route traffic locally on the switch, and only if clients/devices on the VLAN are configured to use the switch's routed interface IP address as their gateway or next hop.
- The layer 3 interface IP cannot be the same as the switch's management IP
- Multicast can be enabled per SVI if required (Refer to Multicast section)
- The Default gateway is the next hop for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a routed interface. This option is available for the first configured SVI interface and will automatically create a static route (essentials a default route via the configured default gateway)
- OSPF can be enabled per SVI if required (Refer to OSPF section)
- Stay within the limits provided in the below table "Routing Scaling Consideration for MS Platforms"
- Each SVI can be configured per switch/stack
- Each switch can have a single SVI per VLAN
- You can edit or move an existing SVI from one switch/stack to another
- You can also delete an existing SVI but please note that A switch must retain at least one routed interface and the default route
- To delete an existing SVI, please follow these steps in exact order otherwise you will get an error and will not allow the route/interface to be deleted:
    
  - Navigate to Switch > Configure > Routing and DHCP
  - Delete any static routes other than the Default route for the desired switch
  - Delete any layer 3 interfaces other than the one which contains the next hop IP for the default route on the desired switch
  - Delete the last layer 3 interface to disable layer 3 routing
Important Notes
- The management IP is treated entirely different from the layer 3 routed interfaces and must be a different IP address.
- Traffic using the management IP address to communicate with the Cisco Meraki Cloud Controller will not use the layer 3 routing settings, instead using its configured default gateway.
- Therefore, it is important that the IP address, VLAN, and default gateway entered for the management/LAN IP ALWAYS provide connectivity to the internet
- The management interface for a switch (stack) performing L3 routing cannot have a configured gateway of one of its own L3 interfaces
- For switch stacks performing L3 routing, ensure that the management IP subnet does not overlap with the subnet of any of its own configured L3 interfaces (except MS390)
- Overlapping subnets on the management IP and L3 interfaces can result in packet loss when pinging or polling (via SNMP) the management IP of stack members (except MS390)
- 
    MS Switches with Layer 3 enabled will prioritize forwarding traffic over responding to pings
- 
    Because of this, packet loss and/or latency may be observed for pings destined for a Layer 3 interface.
- 
    In such circumstances, it's recommended to ping another device in a given subnet to determine network stability and reachability.
MS390 Specific Guidance
- In order to route traffic between VLANs, routed interfaces must be configured.
- Only VLANs with a routed interface configured will be able to route traffic locally on the switch, and only if clients/devices on the VLAN are configured to use the switch's routed interface IP address as their gateway or next hop.
- Multicast can be enabled per SVI if required (Refer to Multicast section)
- The Default gateway is the next hop for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a routed interface. This option is available for the first configured SVI interface and will automatically create a static route (essentials a default route via the configured default gateway)
- OSPF can be enabled per SVI if required (Refer to OSPF section)
- Stay within the limits provided in the below table "Routing Scaling Consideration for MS Platforms"
- Each SVI can be configured per switch/stack
- Each switch can have a single SVI per VLAN
- You can edit or move an existing SVI from one switch/stack to another
- You can also delete an existing SVI but please note that A switch must retain at least one routed interface and the default route
- To delete an existing SVI, please follow these steps in exact order otherwise you will get an error and will not allow the route/interface to be deleted:
    
