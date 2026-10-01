---
id: collect-261001-general-networking/general-networking/switches-layer-3-switching-layer-3-switch-overview-ed7e193d-4
title: "switches-layer-3-switching-layer-3-switch-overview-ed7e193d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-general-networking/switches-layer-3-switching-layer-3-switch-overview-ed7e193d.md
source_anchor: ""
source_lines: [146, 189]
sha256: 2e9e2d2ab1aaf4636510d586440832ad4223a7d49167621d4d0ff64e170a0116
---

# switches-layer-3-switching-layer-3-switch-overview-ed7e193d

- Navigate to Switching > Configure > Routing & DHCP.
- Click Add a static route.
- Select the Switch it should be applied to.
- Provide the following information:
    
  - Name: A friendly name/description for the static route.
  - Subnet: The network that this static route is for, in CIDR notation (ex. 10.1.1.0/24 or 2001:db8::/32).
  - Next hop IP: The IP address of the next L3 device along the path to this network. This address must exist in a subnet with a L3 interface. On switches that support IPv6 static routing, an IPv6 global unicast address can be entered as the next hop IP.
- Click Save or Save and add another if additional static routes are needed.
Note: A default route cannot point to an SVI configured on that same switch. Likewise, avoid configuring any static route with a Next-Hop which is the IP address of another SVI in the same switch. Doing this can create a Layer 3 routing loop and consequent network outages.
Editing an Existing Layer 3 SVI, Routed Port, or Static Route
To modify an existing layer 3 interface or static route on a specific switch:
- Navigate to Switching > Configure > Routing & DHCP.
- Click on the desired Interface or Route.
- Make any desired changes.
- Click Save.
Moving a Layer 3 SVI or Routed Port to Another Switch
To move a layer 3 interface from one switch to another:
- Navigate to Switching > Configure > Routing & DHCP.
- Select the layer 3 interfaces that will be moved.
- Click Edit > Move...
- Select destination switch or switch stack, then click Submit.
Deleting a Layer 3 SVI, Routed Port, or Static Route
In order to delete a layer 3 interface or static route:
- Navigate to Switching > Configure > Routing & DHCP.
- Click on the desired Interface or Route.
- Click Delete Interface/Route, then click Confirm delete.
Note: A switch must retain at least one layer 3 interface and the default route. The default route cannot be manually deleted.
Disabling Layer 3 Routing
In order to disable layer 3 routing, any configured static routes and layer 3 interfaces must be deleted in a specific order.
- Navigate to Switching > Configure > Routing & DHCP.
- Delete any static routes other than the Default route for the desired switch.
- Delete any layer 3 interfaces other than the one which contains the next hop IP for the default route on the desired switch.
- Delete the last layer 3 interface to disable layer 3 routing.
Performing these steps out of order will result in an error and will not allow the route/interface to be deleted.
Layer 3 Interface (SVI) Caveats
Switch Management IP and Layer 3 Interfaces (SVIs)
The management IP is treated entirely different from the SVIs and must be a different IP address. It can be placed on a routed or non-routed VLAN (e.g.: a management VLAN independent from client traffic). Traffic from the management IP address to the Cisco Meraki Cloud Controller will not use the layer 3 routing settings; instead, it will be using its configured default gateway. Therefore, it is important that the IP address, VLAN, and default gateway configured in your switch management IP can still provide connectivity to the internet independently from the switch's own L3 routing settings.
The Switch (or Stack) management IP configuration cannot have Gateway address defined as one of its own SVI address when it is performing Layer 3 routing. It will not be able to check in using the assigned management IP when the gateway is pointed to itself. For example, if 192.168.1.1 is one of the L3 interfaces (SVI) on a switch (or stack), you cannot have 192.168.1.1 as the gateway for its management IP (Switching > Switches > LAN IP).
For switch stacks performing L3 routing, ensure that the management IP subnet does not overlap with the subnet of any of it's own configured L3 interfaces. Overlapping subnets on the management IP and L3 interfaces can result in packet loss when pinging or polling (via SNMP) the management IP of stack members.
Note: The overlapping subnet limitation does not apply to Catalyst switches (MS390/C9300-M) running IOS XE firmware version.
Pings Destined for a Layer 3 Interface
MS Switches with Layer 3 enabled will prioritize forwarding traffic over responding to pings. As a result, packet loss and/or latency may be observed for pings destined to an SVI address. Therefore, it's recommended to ping another device in a given subnet to check network stability and reachability.
Note: Meraki MS classic switches (excluding MS390) are unable to ping their own Layer 3 Interface.
