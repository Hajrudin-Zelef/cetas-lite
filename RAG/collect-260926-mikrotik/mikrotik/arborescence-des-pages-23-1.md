---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-23-1
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-23.md
source_anchor: ""
source_lines: [1, 97]
sha256: 40e41e22a1333945802f8820c2acf43b78406aa1520d5552c2bf88fbef2a7fbc
---

# Summary

**Sub-menu:** `/interface vlan`

**Standards:** `IEEE 802.1Q, IEEE 802.1ad`

Virtual Local Area Network (VLAN) is a Layer 2 method that allows multiple Virtual LANs on a single physical interface (ethernet, wireless, etc.), giving the ability to segregate LANs efficiently.

You can use MikroTik RouterOS (as well as Cisco IOS, Linux, and other router systems) to mark these packets as well as to accept and route marked ones.

As VLAN works on OSI Layer 2, it can be used just like any other network interface without any restrictions. VLAN successfully passes through regular Ethernet bridges.

You can also transport VLANs over wireless links and put multiple VLAN interfaces on a single wireless interface. Note that as VLAN is not a full tunnel protocol (i.e., it does not have additional fields to transport MAC addresses of sender and recipient), the same limitation applies to bridging over VLAN as to bridging plain wireless interfaces. In other words, while wireless clients may participate in VLANs put on wireless interfaces, it is not possible to have VLAN put on a wireless interface in station mode bridged with any other interface.

# 802.1Q

The most commonly used protocol for Virtual LANs (VLANs) is IEEE 802.1Q. It is a standardized encapsulation protocol that defines how to insert a four-byte VLAN identifier into the Ethernet header.

Each VLAN is treated as a separate subnet. It means that by default, a host in a specific VLAN cannot communicate with a host that is a member of another VLAN, although they are connected to the same switch. So if you want inter-VLAN communication you need a router. RouterOS supports up to 4094 VLAN interfaces, each with a unique VLAN ID, per interface. VLAN priorities may also be used and manipulated.

When the VLAN extends over more than one switch, the inter-switch link has to become a 'trunk', where packets are tagged to indicate which VLAN they belong to. A trunk carries the traffic of multiple VLANs; it is like a point-to-point link that carries tagged packets between switches or between a switch and router.

The IEEE 802.1Q standard has reserved VLAN IDs with special use cases, the following VLAN IDs should not be used in generic VLAN setups: 0, 1, 4095

# Q-in-Q

Original 802.1Q allows only one VLAN header, Q-in-Q on the other hand allows two or more VLAN headers. In RouterOS, Q-in-Q can be configured by adding one VLAN interface over another. Example:

If any packet is sent over the 'vlan2' interface, two VLAN tags will be added to the Ethernet header - '11' and '12'.

# Properties

**Sub-menu:** `/interface vlan`

| Property | Description | 
|---|---|
| **arp** (*disabled \| enabled \| local-proxy-arp \| proxy-arp \| reply-only* ; Default:**enabled** ) | Address Resolution Protocol setting  | 
| **arp-timeout** (*auto \| integer* ; Default:**auto** ) | How long the ARP record is kept in the ARP table after no packets are received from IP. Value `auto` equals to the value of`arp-timeout` in IP/Settings, default is 30s. | 
| **disabled** (*yes \| no* ; Default:**no** ) | Changes whether the bridge is disabled. | 
| **interface** (*name* ; Default: ) | Name of the interface on top of which VLAN will work. Adding a VLAN interface to a bridge with vlan-filtering enabled will automatically tag the bridge interface as a member port. A dynamic entry with the comment "added by vlan on bridge" will appear under the `/interface/bridge/vlan` menu. | 
| **l3-hw-offloading** (*yes \| no* ; Default:**yes** ) | Enables or disabled L3HW on a per-VLAN interface. This setting is only applicable to devices that support L3HW offloading and is available starting from RouterOS v7.21. More details - Per-VLAN offloading. | 
| **mvrp** (*yes \| no* ; Default:**no** ) | Specifies whether this VLAN should declare its attributes through Multiple VLAN Registration Protocol (MVRP) as an applicant. Its main use case is for VLANs that is created on Ethernet interface (such as a "router on a stick" setup) that is connected to a bridge supporting MVRP. Enabling this option on a VLAN interface that is already part of an MVRP-enabled bridge has no effect, as the bridge manages MVRP in that case. This property only has an effect when `use-service-tag` is disabled. | 
| **mtu** (*integer: 68..65535* ; Default:**1500** ) | Layer3 Maximum transmission unit | 
| **name** (*string* ; Default: ) | Interface name | 
| **use-service-tag** (*yes \| no* ; Default: ) | IEEE 802.1ad compatible Service Tag | 
| **vlan-id** (*integer: 1..4094* ; Default:**1** ) | Virtual LAN identifier or tag that is used to distinguish VLANs. Must be equal for all computers that belong to the same VLAN. | 

MTU should be set to 1500 bytes same as on Ethernet interfaces. But this may not work with some Ethernet cards that do not support receiving/transmitting of full-size Ethernet packets with VLAN header added (1500 bytes data + 4 bytes VLAN header + 14 bytes Ethernet header). In this situation, MTU 1496 can be used, but note that this will cause packet fragmentation if larger packets have to be sent over the interface. At the same time remember that MTU 1496 may cause problems if path MTU discovery is not working properly between source and destination.

# Setup examples

## Video examples

VLANs pt1, VLANs pt2, VLANs pt3

## Layer2 VLAN examples

There are multiple possible configurations that you can use, but each configuration type is designed for a special set of devices since some configuration methods will give you the benefits of the built-in switch chip and gain larger throughput. Check the Basic VLAN switching guide to see which configuration to use for each type of device to gain maximum possible throughput and compatibility, the guide shows how to set up a very basic VLAN trunk/access port configuration.

There are some other ways to set up VLAN tagging or VLAN switching, but the recommended way is to use Bridge VLAN Filtering. Make sure you have not used any known Layer2 misconfigurations.

## Layer3 VLAN examples

### Simple VLAN routing

Let us assume that we have several MikroTik routers connected to a hub. Remember that a hub is an OSI physical layer device (if there is a hub between routers, then from the L3 point of view it is the same as an Ethernet cable connection between them). For simplification assume that all routers are connected to the hub using the ether1 interface and have assigned IP addresses as illustrated in the figure below. Then on each of them the VLAN interface is created.

Configuration for R2 and R4 is shown below:

R2:

R4:

The next step is to assign IP addresses to the VLAN interfaces.

R2:

R4:

At this point, it should be possible to ping router R4 from router R2 and vice versa:

To make sure if the VLAN setup is working properly, try to ping R1 from R2. If pings are timing out then VLANs are successfully isolated.

### InterVLAN routing

If separate VLANs are implemented on a switch, then a router is required to provide communication between VLANs. A switch works at OSI layer 2 so it uses only the Ethernet header to forward and does not check the IP header. For this reason, we must use the router that is working as a gateway for each VLAN. Without a router, a host is unable to communicate outside of its own VLAN. The routing process between VLANs described above is called inter-VLAN communication.

To illustrate inter-VLAN communication, we will create a trunk that will carry traffic from three VLANs (VLAN2 and VLAN3, VLAN4) across a single link between a Mikrotik router and a manageable switch that supports VLAN trunking.

Each VLAN has its own separate subnet (broadcast domain) as we see in the figure above:

- VLAN 2 – 10.10.20.0/24;
- VLAN 3 – 10.10.30.0/24;
- VLAN 4 – 10.10.40.0./24.

VLAN configuration on most switches is straightforward, we need to define which ports are members of the VLANs and define a 'trunk' port that can carry tagged frames between the switch and the router.

