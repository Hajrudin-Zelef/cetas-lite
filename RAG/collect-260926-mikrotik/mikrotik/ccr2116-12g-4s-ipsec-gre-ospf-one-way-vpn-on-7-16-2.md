---
id: collect-260926-mikrotik/mikrotik/ccr2116-12g-4s-ipsec-gre-ospf-one-way-vpn-on-7-16-2
title: "ccr2116-12g-4s-ipsec-gre-ospf-one-way-vpn-on-7-16-2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/lot-mikrotik/forum/ipsec/ccr2116-12g-4s-ipsec-gre-ospf-one-way-vpn-on-7-16-2.md
source_anchor: ""
source_lines: [1, 103]
sha256: dea4f8a522969dd1ca7620773c60e7a04107dbfaf25e2a89ebffe32811f1ba03
---

# ccr2116-12g-4s-ipsec-gre-ospf-one-way-vpn-on-7-16-2

I’m posting my support ticket here in case someone else has the same problem and is googling (or better yet, knows what’s going on)…

Symptom

We have a VPN configuration where we connect to two routers over IPSEC then build a GRE tunnel, then do OSPF over the gre tunnels so that we have redundant vpn uplinks. When running 7.15.3 everything works fine, but then we upgrade to 7.16.2, 7.17.2, 7.18.2, or 7.19rc1 we can send traffic through the tunnel egress from the CCR2116-12G-4S+ but ingress traffic is has a lot of TCP retransmissions, TCP DUP ACK, and generally doesn’t work.

Here is how the config works:

/interface bridge

add frame-types=admit-only-vlan-tagged name=bridge port-cost-mode=short vlan-filtering=yes

add fast-forward=no mtu=1500 name=loopback port-cost-mode=short

/interface pppoe-client

add add-default-route=yes disabled=no interface=ether1 keepalive-timeout=60 name=isp use-peer-dns=yes user=pppoeuser

/interface gre

add allow-fast-path=no clamp-tcp-mss=no !keepalive local-address=.96.237 mtu=1472 name=gre1 remote-address=.96.235

add allow-fast-path=no clamp-tcp-mss=no !keepalive local-address=.96.236 mtu=1472 name=gre2 remote-address=.96.253

/interface vlan

add interface=bridge name=vlan25 vlan-id=25

add interface=bridge name=vlan71 vlan-id=71

/interface bridge port

add bridge=bridge interface=ether2

/interface bridge vlan

add bridge=bridge tagged=ether2,bridge vlan-ids=71

add bridge=bridge tagged=ether2,bridge vlan-ids=25

/ip address

add address=.200.14/30 interface=gre1 network=.200.12

add address=.201.14/30 interface=gre2 network=.201.12

/ip ipsec peer

add address=x.x.x.x/32 name=peer_1 profile=profile

add address=y.y.y.y/32 name=peer_2 profile=profile

/ip ipsec policy

add dst-address=.96.253/32 peer=peer_1 proposal=proposal protocol=gre src-address=.96.236/32 tunnel=yes

add dst-address=.96.235/32 peer=peer_2 proposal=proposal protocol=gre src-address=.96.237/32 tunnel=yes

/routing ospf interface-template

add area=backbone-v2 auth-id=1 auth-key=“” cost=10 dead-interval=4s disabled=no hello-interval=1s interfaces=gre1 networks=.200.12/30 type=ptp

add area=backbone-v2 auth-id=1 auth-key=“” cost=300 dead-interval=4s disabled=no hello-interval=1s interfaces=gre2 networks=.201.12/30 type=ptp

/ip firewall mangle

add action=change-mss chain=forward in-interface-list=vpninterfaces new-mss=1000 passthrough=yes protocol=tcp tcp-flags=syn tcp-mss=1001-65535

add action=change-mss chain=forward new-mss=1000 out-interface-list=vpninterfaces passthrough=yes protocol=tcp tcp-flags=syn tcp-mss=1001-65535

There is obviously more but you get the idea:

IPSEC tunnel to two hosts, GRE Tunnel over to those hosts using loopback addresses on each end, GRE interfaces have a /30 point to point network, OSPF turns on GRE interfaces to pass the routes.

On 7.15.3 this works fine no problems:

% scp test.bin hostinvlan71:

test.bin 100% 58MB 8.4MB/s 00:06

On 7.16.x 7.17.x 7.18.x 7.19rc1 same thing gets to 255KB and stalls:

% scp test.bin hostinvlan71:

test.bin 0% 255KB 1.2KB/s - stalled -^

If I copy a file to the CCR2116 over the VPN performance is fine, if I copy a file to a host on the other side of the CCR2116 then things stall after 255KB.

We are not using hardware routing:

/routing> /interface/ethernet/switch/print

Columns: NAME, TYPE, L3-HW-OFFLOADING, QOS-HW-OFFLOADING

NAME TYPE L3-HW-OFFLOADING QOS-HW-OFFLOADING

0 switch1 Marvell-98DX3255 no no

We don’t have any queues other than default.

“/queue export” returns nothing.

Also note I’m running 1000 MSS to ensure this is not an MTU issue which is confirmed in packet captures.
