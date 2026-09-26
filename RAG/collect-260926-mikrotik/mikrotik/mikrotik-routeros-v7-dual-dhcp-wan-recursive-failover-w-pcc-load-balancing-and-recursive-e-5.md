---
id: collect-260926-mikrotik/mikrotik/mikrotik-routeros-v7-dual-dhcp-wan-recursive-failover-w-pcc-load-balancing-and-recursive-e-5
title: "MikroTik RouterOS v7 — Dual DHCP WAN Recursive Failover with PCC Load Balancing"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["distribution", "ethernet", "throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/mikrotik-routeros-v7-dual-dhcp-wan-recursive-failover-w-pcc-load-balancing-and-recursive-ecmp.md
source_anchor: ""
source_lines: [546, 672]
sha256: 66fa3bea08cd0510f2e3a41279a5f10869b46fc7f1bf35614c8f079d268a480c
---

# MikroTik RouterOS v7 — Dual DHCP WAN Recursive Failover with PCC Load Balancing

| Symptom | Likely cause | What to check | 
| Recursive routes show unreachable | DHCP script did not update gateway, route comments do not match, or probe is blocked | `/ip route print detail where comment~"WAN"` ,`/log print` , DHCP client status | 
| Traffic marked for ISP1 exits ISP2 unexpectedly | FastTrack still active or main table is being used | Disable FastTrack; check mangle counters and route marks | 
| Router services reply through the wrong WAN | Inbound connection was not marked, or output mark rule missing | Check inbound mangle rules and `chain=output` mark-routing rules | 
| LAN clients cannot reach local VLANs/subnets | Local destinations are being PCC-marked | Add every internal subnet to `address-list=local` ; keep local accept rule before PCC | 
| Failover works, but existing downloads or SSH sessions die | Expected behavior; PCC is connection-based and NAT public IP changed | Start new sessions; optionally clear affected connection marks | 
| Failback does not happen after WAN recovery | Probe route still inactive, probe target does not answer, or DHCP route still disabled | Ping probe via intended WAN; check DHCP script logs and route state | 
| Port forwards work on only one WAN | Inbound connection marking or dst-nat rules are incomplete | Mark new inbound connections by WAN and ensure dst-nat rules exist for both WANs | 
| Traceroute looks odd | PCC, ECMP, ICMP behavior, and router-originated traffic can differ from TCP flows | Test with routing-table-specific pings/traceroutes and with real TCP traffic | 
| BGP/OSPF/tunnels behave oddly | Mangle rules can override route decisions for traffic that should use another policy | Add explicit accept/bypass rules before PCC for routing protocols, tunnels, or management subnets | 
| Load balancing seems uneven | PCC balances connection buckets, not bandwidth | Use weighted PCC buckets for asymmetric links; test with many flows | 

Manual cleanup commands:

```
/ip firewall connection remove [find where connection-mark="ISP1_conn"]
/ip firewall connection remove [find where connection-mark="ISP2_conn"]
```


### 9.1 Load-balance only selected clients

Use this when most clients should use WAN1 primary / WAN2 backup, while selected clients are PCC load-balanced.

First, make the main-table recursive defaults active-backup instead of ECMP:

```
/ip route set [find where comment="MAIN_RECURSIVE_WAN1"] distance=1
/ip route set [find where comment="MAIN_RECURSIVE_WAN2"] distance=2
```

Create an address list for clients allowed to use PCC:

```
/ip firewall address-list
add list=MultiWAN-Clients address=192.168.88.50 comment="example load-balanced client"
add list=MultiWAN-Clients address=192.168.88.60 comment="example load-balanced client"
```

Then add `src-address-list=MultiWAN-Clients` to the two PCC mark-connection rules:

```
/ip firewall mangle
set [find where comment="PCC bucket 0 -> ISP1"] src-address-list=MultiWAN-Clients
set [find where comment="PCC bucket 1 -> ISP2"] src-address-list=MultiWAN-Clients
```

Non-listed clients will not receive PCC routing marks and will use the main table.


### 9.2 Weighted PCC for asymmetric WAN speeds

PCC does not aggregate bandwidth, but you can bias new connection distribution.

Example: WAN1 is 250 Mbps, WAN2 is 500 Mbps. Use a 1:2 connection-bucket ratio.

Replace the two PCC rules with three buckets:

```
/ip firewall mangle
remove [find where comment="PCC bucket 0 -> ISP1"]
remove [find where comment="PCC bucket 1 -> ISP2"]
add chain=prerouting action=mark-connection connection-state=new connection-mark=no-mark in-interface-list=LAN dst-address-list=!local dst-address-type=!local new-connection-mark=ISP1_conn passthrough=yes per-connection-classifier=both-addresses-and-ports:3/0 comment="PCC weighted bucket 0 -> ISP1"
add chain=prerouting action=mark-connection connection-state=new connection-mark=no-mark in-interface-list=LAN dst-address-list=!local dst-address-type=!local new-connection-mark=ISP2_conn passthrough=yes per-connection-classifier=both-addresses-and-ports:3/1 comment="PCC weighted bucket 1 -> ISP2"
add chain=prerouting action=mark-connection connection-state=new connection-mark=no-mark in-interface-list=LAN dst-address-list=!local dst-address-type=!local new-connection-mark=ISP2_conn passthrough=yes per-connection-classifier=both-addresses-and-ports:3/2 comment="PCC weighted bucket 2 -> ISP2"
```

For different ratios, increase the denominator and allocate more buckets to the faster WAN.


### 9.3 Static WAN instead of DHCP

For a static Ethernet WAN, do not use the DHCP script for that interface. Set the recursive host routes manually to the ISP next-hop IP.

Example WAN1 static gateway `203.0.113.1`:

```
/ip route set [find where comment="WAN1_RECURSIVE"] gateway=203.0.113.1 disabled=no
```

Do not use `gateway=ether1` on a normal broadcast Ethernet WAN as a long-term substitute for a real next-hop IP. Use the actual gateway IP from the ISP.


### 9.4 PPPoE WAN instead of DHCP

For PPPoE, use the PPPoE interface as the gateway for recursive probe host routes because PPPoE is point-to-point.

Example WAN1 PPPoE interface `pppoe-out1`:

```
/interface list member add interface=pppoe-out1 list=WAN comment="WAN1 PPPoE"
/ip route set [find where comment="WAN1_RECURSIVE"] gateway=pppoe-out1 disabled=no
```

Then remove (not just disable) the WAN1 DHCP client, including its script. The script in §6.7 references `$"gateway-address"` from the DHCP lease; it has no meaning for PPPoE and must not run alongside this manual `set`. If you keep the DHCP client object disabled, also clear its script field to avoid confusion later.


PCC is not bandwidth bonding. It does not split one TCP flow across both WANs. It assigns each new connection to a bucket and keeps that connection sticky to the selected path as long as the connection is tracked.

Practical effects:

- speed tests with one connection may show only one WAN;
- multiple clients or multi-connection tests show balancing better;
- long-lived sessions may break when their selected WAN fails;
- sites that bind login/session state to public IP may behave poorly when different connections leave through different ISPs.

Use probe IPs that:

- are stable;
- answer ICMP consistently;
- are outside your ISP network if you want to detect wider internet reachability;
- are not all in the same provider or DNS service;
- are not the same IPs you plan to use as the router's DNS resolvers;
- are not blocked by the ISP.

If a provider blocks ICMP, choose another target or monitor by another method.

The example uses well-known anycast public DNS IPs from Quad9, OpenDNS, and Verisign/UltraDNS because they are globally reachable and tend to answer ICMP. It deliberately avoids Cloudflare and Google as route probes so those popular resolvers can be used for actual DNS service without also becoming route-health sensors.

The main-table probes (`64.6.64.6` / `149.112.112.112`) and the policy-table probes use different IPs on purpose. Reachability of one set does not guarantee reachability of the other, so it is normal during partial outages to see the router itself reach the internet via the main table while LAN clients on a policy table appear stuck — or the other way around. Compare both sets of probes when triaging.

Keep DNS resolvers and route probes separate where practical. If the router's only DNS servers are also the only recursive probe targets, DNS-service failure and route-health failure become harder to distinguish during troubleshooting.

### FastTrack and performance

Disabling FastTrack/FastPath can reduce maximum throughput on low-end routers because more packets stay on the slow path. On devices such as hEX-class routers, test CPU under real traffic. On more capable routers such as RB5009-class hardware, this design is usually much more comfortable, but throughput still depends on firewall complexity, queueing, packet size, and enabled services.

