---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-67-2
title: "Fast Path"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-67.md
source_anchor: ""
source_lines: [108, 190]
sha256: 644f1d704159bacda785a7dd0e41d8f0ade96def87b77132f416e4b45897a275
---

# Fast Path

1. A very similar process happens when a packet's destination is a router (routing input): Packet enters prerouting processing:
  1. - check if there is a hotspot and modify the packet for hotspot use;
  2. - process packet through RAW prerouting chain;
  3. - send a packet through connection tracking;
  4. - process packet through Mangle prerouting chain;
  5. - process packet through NATs dst-nat chain;
2. Run packet through routing table to make routing decision;
3. A Packet enters the input process;
  1. - process packet through Mangle input chain;
  2. - process packet through Filter input chain;
  3. - process packet through queue tree (HTB Global);
  4. - process packet through simple queues;
4. Check if there is IPsec and then process through IPsec policies.

### Output

Or when a packet is originated from the router (routing output):

1. The packet is originated from the router itself
  1. the packet goes through the routing table to make a routing decision
2. the packet goes through the routing table to make a routing decision
3. A packet enters the output process
  1. process packet through the Bridge decision;
  2. send the packet through connection tracking;
  3. process packet through the Mangle output chain;
  4. process packet through the Filter output chain;
  5. send the packet to routing adjustment ( policy routing)
4.  The packet enters postrouting process; 
  1. - process packet through Mangle postrouting chain;
  2. - process packet through NATs src-nat chain;
  3. - if there is a hotspot undo any modifications made in hotspot-in;
  4. - process packet through queue tree (HTB Global);
  5. - process packet through simple queues;
5. Check if there is IPsec and then process through IPsec policies;

## Flow of Bridged Packet

Below is discussed a general bridging process in RouterOS. Most of the packets will always follow the same processing path, but in certain configurations (e.g. with enabled VLAN filtering, horizon, STP, DHCP, or IGMP snooping) some packets can be treated differently. Please visit the bridging manual for more specific information.

### Bridge Forward

Bridge forward is a process that takes place when a packet is forwarded from one bridge port to another, essentially connecting multiple devices on the same network. After receiving a packet on the in-interface, the device determines that the in-interface is a bridge port, so it gets passed through the bridging process:

1. A packet goes through the bridge NAT dst-nat chain, where MAC destination and priority can be changed, apart from that, a packet can be simply accepted, dropped, or marked;
2. Checks whether the use-ip-firewall option is enabled in the bridge settings;
3. Run packet through the bridge host table to make a forwarding decision. A packet that ends up being flooded (e.g. broadcast, multicast, unknown unicast traffic), gets multiplied per bridge port and then processed further in the bridge forward chain. When using `vlan-filtering=yes` , packets that are not allowed due to the "/interface bridge vlan" table, will be dropped at this stage.
4. A packet goes through the bridge filter forward chain, where priority can be changed or the packet can be simply accepted, dropped, or marked;
5. Checks whether the use-ip-firewall option is enabled in the bridge settings;
6. A packet goes through the bridge NAT src-nat chain, where MAC source and priority can be changed, apart from that, a packet can be simply accepted, dropped, or marked;
7. Checks whether the use-ip-firewall option is enabled in the bridge settings;

**For RouterOS v6:**

When bridge `vlan-filtering` is enabled, received untagged packets might get encapsulated into the VLAN header before the "DST-NAT" block, which means these packets can be filtered using the `mac-protocol=vlan` and `vlan-encap` settings. Encapsulation can happen if the outgoing interface has `frame-types` set to `admit-all` or `admit-only-untagged-and-priority-tagged`.

Tagged packets might get decapsulated on the "BRIDGING DECISION" block, which means these packets will no longer match the `mac-protocol=vlan` and `vlan-encap` settings. Decapsulation can happen if the packet's VLAN ID matches the outgoing port's untagged VLAN membership.**For RouterOS v7 and newer:**

When bridge `vlan-filtering` is enabled, received untagged packets might get encapsulated into the VLAN header on the "BRIDGING-DECISION" block, which means these packets can be filtered using the `mac-protocol=vlan` and `vlan-encap` settings.  Encapsulation can happen if the outgoing interface has `frame-types` set to `admit-all` or `admit-only-untagged-and-priority-tagged`.

Tagged packets might get decapsulated on the "BRIDGING DECISION" block, which means these packets will no longer match the `mac-protocol=vlan` and `vlan-encap` settings. Decapsulation can happen if the packet's VLAN ID matches the outgoing port's untagged VLAN membership.

### Bridge Input

Bridge input is a process that takes place when a packet is destined for the bridge interface. Most commonly this happens when you need to reach some services that are running on the bridge interface (e.g. a DHCP server) or you need to route traffic to other networks. The very first steps are similar to the bridge forward process - after receiving a packet on the in-interface, the device determines that the in-interface is a bridge port, so it gets passed through the bridging process:

1. A packet goes through the bridge NAT dst-nat chain, where MAC destination and priority can be changed, apart from that, a packet can be simply accepted, dropped, or marked;
2. Checks whether the use-ip-firewall option is enabled in the bridge settings;
3. Run packet through the bridge host table to make a forwarding decision. A packet where the destination MAC address matches the bridge MAC address will be passed to the bridge input chain. A packet that ends up being flooded (e.g. broadcast, multicast, unknown unicast traffic), also reaches the bridge input chain as the bridge interface itself is one of the many destinations;
4. A packet goes through the bridge filter input chain, where priority can be changed or the packet can be simply accepted, dropped, or marked;

### Bridge Output

Bridge output is a process that takes place when a packet should exit the device through one or multiple bridge ports. Most commonly this happens when a bridge interface itself tries to reach a device connected to a certain bridge port (e.g. when a DHCP server running on a bridge interface is responding to a DHCP client). After a packet is processed on other higher-level RouterOS processes and the device finally determines that the output interface is a bridge, the packet gets passed through the bridging process:

1. Run packet through the bridge host table to make a forwarding decision. A packet that ends up being flooded (e.g. broadcast, multicast, unknown unicast traffic), gets multiplied per bridge port and then processed further in the bridge output chain.
2. A packet goes through the bridge filter output chain, where priority can be changed or the packet can be simply accepted, dropped, or marked;
3. A packet goes through the bridge NAT src-nat chain, where MAC source and priority can be changed, apart from that, a packet can be simply accepted, dropped, or marked;
4. Checks whether the use-ip-firewall option is enabled in the bridge settings;

### Forward With Firewall Enabled

In certain network configurations, you might need to enable additional processing on routing chains for bridged traffic, for example, to use simple queues or an IP firewall. This can be done when the use-ip-firewall is enabled under the bridge settings. Note that additional processing will consume more CPU resources to handle these packets. All the steps were already discussed in previous points, below is a recap:

