---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-68-1
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters", "preemption"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-68.md
source_anchor: ""
source_lines: [1, 109]
sha256: e256f01320b857a1a8980367f3b55193b5197ad630e06d46a91bcb03a229120e
---

# Summary

This chapter describes the Virtual Router Redundancy Protocol (VRRP) support in RouterOS.

Mostly on larger LANs dynamic routing protocols (OSPF or RIP) are used, however, there are a number of factors that may make it undesirable to use dynamic routing protocols. One alternative is to use static routing, but if the statically configured first hop fails, then the host will not be able to communicate with other hosts.

In IPv6 networks, hosts learn about routers by receiving Router Advertisements used by the Neighbor Discovery (ND) protocol. ND already has a built-in mechanism to determine unreachable routers. However, it can take up to 38 seconds to detect an unreachable router. It is possible to change parameters and make detection faster, but it will increase the overhead of ND traffic especially if there are a lot of hosts. VRRP allows the detection of unreachable routers within 3 seconds without additional traffic overhead.

Virtual Router Redundancy Protocol (VRRP) provides a solution by combining a number of routers into a logical group called *Virtual Router* (VR). VRRP implementation in RouterOS is based on VRRPv2 RFC 3768 and VRRPv3 RFC 5798.

It is recommended to use the same version of RouterOS for all devices with the same VRID used to implement VRRP.

According to RFC authentication is deprecated for VRRP v3.

# Protocol Overview

The purpose of the VRRP is to communicate to all VRRP routers associated with the Virtual Router ID and support router redundancy through a prioritized election process among them.

All messaging is done by IPv4 or IPv6 multicast packets using protocol 112 (VRRP). The destination address of an IPv4 packet is *224.0.0.18* and for IPv6 it is *FF02:0:0:0:0:0:0:12*. The source address of the packet is always the primary IP address of an interface from which the packet is being sent. In IPv6 networks, the source address is the link-local address of an interface.

These packets are always sent with TTL=255 and are not forwarded by the router. If for any reason the router receives a packet with lower TTL, a packet is discarded.

Each VR node has a single assigned MAC address. This MAC address is used as a source for all periodic messages sent by Master.

Virtual Router is defined by VRID and mapped set of IPv4 or IPv6 addresses. The master router is said to be the **owner** of mapped IPv4/IPv6 addresses. There are no limits to using the same VRID for IPv4 and IPv6, however, these will be two different Virtual Routers.

Only the Master router is sending periodic Advertisement messages to minimize the traffic. A backup will try to preempt the Master only if it has the higher priority and preemption is not prohibited.

All VRRP routers belonging to the same VR must be configured with the same advertisement interval. If the interval does not match router will discard the received advertisement packet.

## Virtual Router (VR)

A Virtual Router (VR) consists of one Owner router and one or more backup routers belonging to the same network.

VR includes:

- VRID configured on each VRRP router
- the same virtual IP on each router
- Owner and Backup configured on each router. On a given VR there can be only one Owner.

### Virtual MAC address

VRRP automatically assigns MAC address to VRRP interface based on standard MAC prefix for VRRP packets and VRID number. The first five octets are 00:00:5E:00:01 and the last octet is configured VRID. For example, if Virtual Routers VRID is 49, then the virtual MAC address will be *00:00:5E:00:01:31*.

Virtual mac addresses can not be manually set or edited.

## Owner

An Owner router for a VR is the default Master router and operates as the Owner for all subnets included in the VR. Priority on an owner router must be the highest value (255) and virtual IP is the same as real IP (owns the virtual IP address).

RouterOS can not be configured as Owner. The Pure virtual IP configuration is the only valid configuration unless a non-RouterOS device is set as the owner.

## Master

A master router in a VR operates as the physical gateway for the network for which it is configured. The selection of the Master is controlled by priority value. The Master state describes the behavior of the Master router. For example network, **R1** is the Master router. When R1 is no longer available R2 becomes master.

## Backup

VR must contain at least one Backup router. A backup router must be configured with the same virtual IP as the Master for that VR. The default priority for Backup routers is 100. When the current master router is no longer available, a backup router with the highest priority will become a current master. Every time when a router with higher priority becomes available it is switched to master. Sometimes this behavior is not necessary. To override it preemption mode should be disabled.

## Virtual Address

Virtual IP associated with VR must be identical and set on all VR nodes. All virtual and real addresses should be from the same network.

RouterOS can not be configured as Owner. VRRP address and real IP address should not be the same.

If the Master of VR is associated with multiple IP addresses, then Backup routers belonging to the same VR must also be associated with the same set of virtual IP addresses. If the virtual address on the Master is not also on Backup a misconfiguration exists and VRRP advertisement packets will be discarded.

All Virtual Router members can be configured so that virtual IP is not the same as physical IP. Such a virtual address can be called a floating or pure virtual IP address. The advantage of this setup is the flexibility given to the administrator. Since the virtual IP address is not the real address of any one of the participant routers, the administrator can change these physical routers or their addresses without any need to reconfigure the virtual router itself.

In IPv6 networks, the first address is always a link-local address associated with VR. If multiple IPv6 addresses are configured, then they are added to the advertisement packet after the link-local address.

### IPv4 ARP

The Master for a given VR responds to ARP requests with the VR's assigned MAC address. The virtual MAC address is also used as the source MAC address for advertisement packets sent by the Master. To ARP requests for non-virtual IP, addresses router responds with the system MAC address. Backup routers are not responding to ARP requests for Virtual IPs.

### IPv6 ND

As you may know, in IPv6 networks, the Neighbor Discovery protocol is used instead of ARP. When a router becomes the Master, an unsolicited ND Neighbor Advertisement with the Router Flag is sent for each IPv6 address associated with the virtual router.

# VRRP state machine

As you can see from the diagram, each VRRP node can be in one of three states:

- Init state
- Backup state
- Master state

### Init state

The purpose of this state is to wait for a Startup event. When this event is received, the following actions are taken:

- if priority is 255,
- * for IPv4 send advertisement packet and broadcast ARP requests;
- * for IPv6 send an unsolicited ND Neighbor Advertisement for each IPv6 address associated with the virtual router and set target address to link-local address associated with VR;
- * transit to MASTER state;
- else transit to BACKUP state.

### Backup state

When in the backup state,

- in IPv4 networks, a node is not responding to ARP requests and is not forwarding traffic for the IP associated with the VR.
- in IPv6 networks, a node is not responding to ND Neighbor Solicitation messages and is not sending ND Router Advertisement messages for VR-associated IPv6 addresses.

Routers' main task is to receive advertisement packets and check if the master node is available.

The backup router will transmit itself to the master state in two cases:

