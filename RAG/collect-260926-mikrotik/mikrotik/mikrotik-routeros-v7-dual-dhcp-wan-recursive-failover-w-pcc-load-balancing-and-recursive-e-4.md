---
id: collect-260926-mikrotik/mikrotik/mikrotik-routeros-v7-dual-dhcp-wan-recursive-failover-w-pcc-load-balancing-and-recursive-e-4
title: "MikroTik RouterOS v7 — Dual DHCP WAN Recursive Failover with PCC Load Balancing"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/mikrotik-routeros-v7-dual-dhcp-wan-recursive-failover-w-pcc-load-balancing-and-recursive-ecmp.md
source_anchor: ""
source_lines: [394, 545]
sha256: 3d6976d7c3771401711b0344a5badab3f7c8c5f320a7859721a438568fe40325
---

# MikroTik RouterOS v7 — Dual DHCP WAN Recursive Failover with PCC Load Balancing

```
/ip firewall mangle
# Do not PCC traffic whose destination is another local subnet.
add chain=prerouting action=accept in-interface-list=LAN dst-address-list=local comment="skip local destinations before PCC"
# Mark new inbound connections by the WAN where they arrived.
# This helps replies from router services and dst-nat flows return through the same ISP.
add chain=prerouting action=mark-connection connection-state=new connection-mark=no-mark in-interface=ether1 new-connection-mark=ISP1_conn passthrough=yes comment="new inbound via WAN1"
add chain=prerouting action=mark-connection connection-state=new connection-mark=no-mark in-interface=ether2 new-connection-mark=ISP2_conn passthrough=yes comment="new inbound via WAN2"
# PCC for new LAN-originated internet connections.
add chain=prerouting action=mark-connection connection-state=new connection-mark=no-mark in-interface-list=LAN dst-address-list=!local dst-address-type=!local new-connection-mark=ISP1_conn passthrough=yes per-connection-classifier=both-addresses-and-ports:2/0 comment="PCC bucket 0 -> ISP1"
add chain=prerouting action=mark-connection connection-state=new connection-mark=no-mark in-interface-list=LAN dst-address-list=!local dst-address-type=!local new-connection-mark=ISP2_conn passthrough=yes per-connection-classifier=both-addresses-and-ports:2/1 comment="PCC bucket 1 -> ISP2"
# Apply routing marks to LAN packets based on connection mark.
add chain=prerouting action=mark-routing connection-mark=ISP1_conn in-interface-list=LAN new-routing-mark=to_ISP1 passthrough=no comment="route ISP1_conn via to_ISP1"
add chain=prerouting action=mark-routing connection-mark=ISP2_conn in-interface-list=LAN new-routing-mark=to_ISP2 passthrough=no comment="route ISP2_conn via to_ISP2"
# Router-generated replies for marked inbound connections.
# Do NOT add connection-state=new here; replies on established flows must also be marked.
add chain=output action=mark-routing connection-mark=ISP1_conn dst-address-list=!local new-routing-mark=to_ISP1 passthrough=no comment="router output replies via ISP1"
add chain=output action=mark-routing connection-mark=ISP2_conn dst-address-list=!local new-routing-mark=to_ISP2 passthrough=no comment="router output replies via ISP2"
```

The `mark-routing` rules use `passthrough=no` on purpose. Once a packet has its routing mark, no later mangle rule should change it; `passthrough=yes` here would let a subsequent rule overwrite the mark and silently misroute traffic. The earlier `mark-connection` rules use `passthrough=yes` because the next stage (`mark-routing`) needs to read the connection mark.

Notes:

- PCC is per connection, not per packet.
- A single TCP session uses one WAN; it is not split across both links.
- Existing sessions usually break during failover. New sessions should use the active route.
- Some banking, gaming, streaming, and CDN flows may dislike changing public IPs across different connections.


```
/ip firewall nat
add chain=srcnat action=masquerade out-interface-list=WAN ipsec-policy=out,none comment="masquerade both WANs"
```

If you use IPsec, WireGuard, GRE, EOIP, VLANs, or other tunnels, review NAT bypass rules before this masquerade rule.


```
/ip dhcp-client print detail
```

Confirm each WAN has:

- `status=bound` ;
- an address;
- a gateway value, shown as `gateway=` in print output and exposed to scripts as`$"gateway-address"` ;
- `add-default-route=no` .

### 7.2 Check recursive route state

```
/ip route print detail where comment~"WAN|MAIN|to_ISP"
```

Expected behavior:

- `WAN1_RECURSIVE` routes should use WAN1's current DHCP gateway.
- `WAN2_RECURSIVE` routes should use WAN2's current DHCP gateway.
- `to_ISP1_primary_via_WAN1` should be active when WAN1's probe is reachable.
- `to_ISP2_primary_via_WAN2` should be active when WAN2's probe is reachable.
- backup routes should become active when their primary route fails.

### 7.3 Test routing-table-specific pings

```
/tool ping address=9.9.9.9 routing-table=to_ISP1 count=5
/tool ping address=64.6.65.6 routing-table=to_ISP2 count=5
```

`routing-table=` on `/tool ping` requires RouterOS 7.1 or later, but syntax and CLI context can vary between versions and shells. If your build rejects that form, validate with route state plus source-address probes:

```
/ip route print detail where comment~"WAN|to_ISP"
# Replace the source addresses with the DHCP addresses on your WAN interfaces.
/ping address=9.9.9.9 src-address=<WAN1_DHCP_ADDRESS> count=5
/ping address=64.6.65.6 src-address=<WAN2_DHCP_ADDRESS> count=5
```

For PCC itself, prefer real client traffic from the LAN. Router-originated ping tests do not traverse the LAN `prerouting` PCC rules.

### 7.4 Check mangle hit counters

```
/ip firewall mangle print stats all
```

Generate client traffic, then confirm:

- PCC bucket counters increase;
- routing-mark rule counters increase;
- local-destination accept rule increases when clients talk to local subnets.

Example client-side flow generator:

for i in $(seq 1 30); do
  curl -4 -s --max-time 10 https://ifconfig.me
  echo
done | sort | uniq -c

If the client also has a separate public management NIC, make sure the test traffic is sourced through the LAN side of the MikroTik path, for example with `curl --interface <LAN_IP>`.

### 7.5 Check connection marks

```
/ip firewall connection print where connection-mark~"ISP"
```

You should see client flows marked as either `ISP1_conn` or `ISP2_conn`.

Test WAN1 failure:

```
/ip route print detail where routing-table=to_ISP1
```

Then break WAN1 upstream connectivity. Do not rely only on unplugging the Ethernet cable; a real ISP failure may leave the local link up while upstream internet is dead.

Expected result:

- WAN1 recursive probe becomes unreachable.
- `to_ISP1_primary_via_WAN1` becomes inactive.
- `to_ISP1_backup_via_WAN2` becomes the active route.
- New connections marked `ISP1_conn` use WAN2 until WAN1 recovers.

Repeat the test for WAN2.

### 7.7 Validate router-originated DNS if you depend on it

If LAN clients use the MikroTik as their DNS server, also test the router's own upstream DNS path. Client DNS queries enter the router locally, but the router's upstream resolver queries are generated by the router itself and normally use the `main` table.

Use resolver IPs that are not the same as your recursive probe IPs, flush the cache, resolve a new name, and check the DNS connections:

```
/ip dns set allow-remote-requests=yes servers=1.1.1.1,8.8.8.8
/ip dns cache flush
/ip firewall connection remove [find where dst-port=53]
:put [/resolve example.net]
:put [/resolve openai.com]
/ip firewall connection print detail where dst-port=53
/ip route print detail where comment~"MAIN_RECURSIVE"
```

Expected result:

- DNS resolution succeeds.
- DNS connections use the expected WAN source address or expected main-table path.
- If one main-table recursive default is failed, new resolver lookups use the surviving default.

Router-originated traffic uses `main`. If `main` contains a non-recursive default route — common on cloud VMs with a public management NIC, or on routers where DHCP-added WAN defaults still exist — that route can mask the recursive defaults during this test. To validate the real recursive path, manage the router through an out-of-band path (console, LAN-jump, noVNC) and temporarily disable the masking default before repeating the test.


