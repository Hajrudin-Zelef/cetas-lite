---
id: collect-261001-general-networking/general-networking/switches-layer-3-switching-layer-3-switch-overview-ed7e193d-1
title: "switches-layer-3-switching-layer-3-switch-overview-ed7e193d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution", "training"]
source: docs/RAG/collect-261001-general-networking/switches-layer-3-switching-layer-3-switch-overview-ed7e193d.md
source_anchor: ""
source_lines: [1, 44]
sha256: e0a9787b68a266660d0bf03cba0edc86771434a4fad15518e8fbac6af4de1e80
---

# switches-layer-3-switching-layer-3-switch-overview-ed7e193d

MS Layer 3 Switching and Routing
Click 日本語 for Japanese
Layer 3 (L3) SVI Routing Overview
A typical network design segments/divides the network based on the group that a device belongs or the function it performs. For example, a Corporate or a Production VLAN only has devices that belong to your own Organization; on the other hand, a Guest-WIFI VLAN only has visitor devices that are connected to a Guest wireless network (Guest SSID).
In such use-cases, VLANs separate devices into different broadcast domains and Layer 3 (L3) subnets. Network switches typically forward data to its destination based on Layer 2 (L2) packet attributes like VLAN-ID and MAC Address. Network switches that perform this forwarding operation are known as L2 switches. Devices within the same VLAN can communicate with each other without the need for routing as they live within the same broadcast domain and L3 subnet.
As a result, devices in separate VLANs require a L3 forwarding device (e.g., a router) to communicate with each other. The L3 device can be external to the switch or it can be a feature in the same switch. Layer 2-only switches require an external L3 routing device to provide communication between VLANs as they don't have L3 routing functionality i.e., they don't forward data to destination based on L3 attributes like destination IP address.
Many Cisco Meraki switches have L3 routing capability within the switch itself. E.g., a switch receives a packet, determines that the packet belongs to another VLAN, and sends the packet to the appropriate port within the destination VLAN. L3 routing is possible due to an internal interface called a "Switched Virtual Interface" or SVI for short.
When L3 routing is enabled on the switch, SVIs per VLAN/subnet are created which will forward data based on the destination IP address. Afterwards, devices living within separate VLANs can talk to one another without the need to be in the same broadcast domain or to have a router acting as a gateway.
Learn more with this free online training course on the Meraki Learning Hub:
Layer 3 (L3) Routed Ports Overview
A routed port is a physical port on a switch or router that is configured to act as a Layer 3 interface. It is used for routing IP packets instead of switching layer 2 frames. Unlike regular switch ports, a routed port is not associated with a specific VLAN and does not participate in Layer 2 operations such as STP. Routed ports are commonly used for connections between switches, routers, firewalls, or other network devices where routing is needed. Routed ports are a great option for connecting switches in core/distribution network for inter-switch routing and are ideal for point-to-point links. Routed ports, being physical interfaces, have a more straightforward and faster process for coming up and down compared to SVIs. SVIs need a Layer 2 port to be active before the SVI itself can be enabled as well as all ports using that same VLAN to be down prior to bringing the SVI down. Routed ports are ideal for point-to-point links, where you don't want to spread a VLAN across multiple devices. They can also be easier to manage when using features like ECMP (Equal Cost Multi-Path). Routed ports allow you to reuse the same VLAN ID on another L3 port or from an L2 VLAN without the risk of L2 traffic in between. Routed ports create separate Layer 3 segments, which means that broadcast traffic is confined to that specific subnet and doesn't affect other VLANs or routed ports.
Layer 3 routing capabilities are available on most Cisco Meraki switches. This allows the switches to route traffic between VLANs in a network without the need for an additional layer 3 device. This feature also allows you to control traffic between VLANs using Access Control Lists (ACLs) while reducing workload on your Router or Firewall. In addition, the Cisco Meraki L3 switch can also offer simple DHCP services through its SVIs or Routed Ports.
This feature is supported in IOS XE 17.18.1 and later on MS390 and Cloud Managed Catalyst switches.
Note: The configuration of routed ports is only available on the New Version of the UI
Supported Models
In order to enable and configure layer 3 routing on MS switches, a layer 3 capable switch is required.
The alert, "This switch is routing for too many hosts. Performance may be affected" will be displayed if the current number of routed clients exceeds the values listed in the table below.
| Model | Layer 3 Interfaces | Routes | Maximum Routable Clients | Features | 
|---|---|---|---|---|
| MS150 | 16 | 16 Static routes | 8192 | SVI Static Routing 17.2.1+ Required* DHCP Relay | 
| MS210 | 16 | 16 static routes | 8192 | SVI Static Routing DHCP Relay | 
| MS225 | 16 | 16 static routes | 8192 | SVI Static Routing DHCP Relay | 
| MS250 | 256 | 1024 1 (256 static routes) | 8192 | SVI Static Routing OSPFv2 DHCP Server + Relay Warm Spare (VRRP) Multicast Routing (PIM-SM) | 
| MS350 | 256 | 16384 1 (256 static routes) | 24K | SVI Static Routing OSPFv2 DHCP Server + Relay Warm Spare (VRRP) Multicast Routing (PIM-SM) | 
| MS350X | 256 | 8192 (256 static routes) | 45K | SVI Static Routing OSPFv2 DHCP Server + Relay Warm Spare (VRRP) Multicast Routing (PIM-SM) | 
| MS355 | 256 | 8192 (256 static routes) | 68K | SVI Static Routing OSPFv2 DHCP Server + Relay Warm Spare (VRRP) Multicast Routing (PIM-SM) | 
| MS390 | 256 | 8192 (256 static routes) | 24K | Routed Port SVI Static Routing OSPFv2 DHCP Server + Relay Multicast Routing (PIM-SM) IPv6 Layer-3 Interfaces & Static Routing 2 BGP3 | 
| C9300(-M) | 256 | 8192 (256 static routes) | 24K | Routed Port SVI Static Routing OSPFv2 DHCP Server + Relay Multicast Routing (PIM-SM) IPv6 Layer-3 Interfaces & Static Routing 2 BGP3 | 
| C9350 | 1.9K | 96K | 64K | Routed Port SVI Static Routing OSPFv2 DHCP Server + Relay Multicast Routing (PIM-SM) IPv6 Layer-3 Interfaces & Static Routing 2 BGP4 | 
| Cloud managed C9200, C9200L(-M), C9500H, C9200CX | 256 | 8192 (256 static routes) | 24K | Routed Port SVI Static Routing OSPFv2 DHCP Server + Relay Multicast Routing (PIM-SM) IPv6 Layer-3 Interfaces & Static Routing 2 BGP3 | 
| MS410 | 256 | 16384 1 (256 static routes) | 24K | SVI Static Routing OSPFv2 DHCP Server + Relay Warm Spare (VRRP) Multicast Routing (PIM-SM) | 
| MS425 | 256 | 8192 (256 static routes) | 212K | SVI Static Routing OSPFv2 DHCP Server + Relay Warm Spare (VRRP) Multicast Routing (PIM-SM) | 
| MS450 | 256 | 8192 (256 static routes) | 68K | SVI Static Routing OSPFv2 DHCP Server + Relay Warm Spare (VRRP) Multicast Routing (PIM-SM) | 
1 To prevent hardware TCAM exhaustion, the following platform limitations are enforced on the number of dynamically (OSPF) learned routes
MS250: 900
MS350, MS410: 15000
If the limit is reached, routes will be rejected indiscriminately and may result in erratic routing behavior. To minimize the impact of this, the default route will not be affected by the limit and will be accepted regardless.
2 Supported only on the specified switches series, on firmware versions CS 15.21.1 and higher.
3 Supported on IOS XE 17.18.1 and higher.
4. C9350 BGP support on IOS XE 26.1.1 and higher
Initializing Layer 3 Routing
You must create Layer 3 Interfaces in order to route traffic between VLANs. These special interfaces are called "Switched Virtual Interface" or SVI for short or routed ports. Only VLANs that have an SVI configured will be able to route traffic on the switch. An SVI is a kind of Layer 3 Routing interface and the term Layer 3 / L3 interface and SVI / SVI interface are used interchangeably. A routed port is a physical port on a switch or router that is configured to act as a Layer 3 interface.
Note: only clients/devices configured to use the correct SVI IP address as their gateway or next-hop will have its packets routed by your switch SVI. If a client/device within a VLAN use another IP address as their gateway, then your switch won't be doing the Layer 3 packet forwarding decision.
SVI MTU size is 1500
