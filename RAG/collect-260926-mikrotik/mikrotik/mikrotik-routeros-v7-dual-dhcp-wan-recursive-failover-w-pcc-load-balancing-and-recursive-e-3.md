---
id: collect-260926-mikrotik/mikrotik/mikrotik-routeros-v7-dual-dhcp-wan-recursive-failover-w-pcc-load-balancing-and-recursive-e-3
title: "MikroTik RouterOS v7 — Dual DHCP WAN Recursive Failover with PCC Load Balancing"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/mikrotik-routeros-v7-dual-dhcp-wan-recursive-failover-w-pcc-load-balancing-and-recursive-ecmp.md
source_anchor: ""
source_lines: [265, 393]
sha256: 1bde4d449e1f4a648360af2a8dbe0501760bdfe04407c3dff33cf4ba32b28fbb
---

# MikroTik RouterOS v7 — Dual DHCP WAN Recursive Failover with PCC Load Balancing

```
/ip route
# -------------------------------------------------------------------
# Main-table recursive probes for router-originated traffic.
# These are updated by the DHCP scripts.
# Probes deliberately avoid Cloudflare and Google DNS so those can be
# used as normal DNS resolvers without also being route-health targets.
# -------------------------------------------------------------------
add disabled=yes dst-address=64.6.64.6/32 gateway=ether1 scope=10 target-scope=10 comment="WAN1_RECURSIVE"
add disabled=yes dst-address=149.112.112.112/32 gateway=ether2 scope=10 target-scope=10 comment="WAN2_RECURSIVE"
# Main-table recursive defaults for router-originated new traffic.
# Equal distances use ECMP. If you prefer WAN1 primary / WAN2 backup
# for main-table traffic, set MAIN_RECURSIVE_WAN2 distance=2.
add check-gateway=ping distance=1 dst-address=0.0.0.0/0 gateway=64.6.64.6 scope=10 target-scope=11 comment="MAIN_RECURSIVE_WAN1"
add check-gateway=ping distance=1 dst-address=0.0.0.0/0 gateway=149.112.112.112 scope=10 target-scope=11 comment="MAIN_RECURSIVE_WAN2"
# -------------------------------------------------------------------
# Recursive host routes for the policy routing tables.
# WAN1_RECURSIVE routes are forced through WAN1's current DHCP gateway.
# WAN2_RECURSIVE routes are forced through WAN2's current DHCP gateway.
#
# Each table's primary and backup intentionally use different DNS providers
# so a single provider outage does not deactivate both routes at once.
# -------------------------------------------------------------------
add disabled=yes dst-address=9.9.9.9/32 gateway=ether1 scope=10 target-scope=10 comment="WAN1_RECURSIVE"
add disabled=yes dst-address=208.67.222.222/32 gateway=ether1 scope=10 target-scope=10 comment="WAN1_RECURSIVE"
add disabled=yes dst-address=64.6.65.6/32 gateway=ether2 scope=10 target-scope=10 comment="WAN2_RECURSIVE"
add disabled=yes dst-address=208.67.220.220/32 gateway=ether2 scope=10 target-scope=10 comment="WAN2_RECURSIVE"
# -------------------------------------------------------------------
# Policy-table defaults.
# Each table prefers its own ISP, then falls back to the other ISP.
# Primary and backup deliberately use different probe providers.
# -------------------------------------------------------------------
add check-gateway=ping distance=1 dst-address=0.0.0.0/0 gateway=9.9.9.9 routing-table=to_ISP1 scope=10 target-scope=11 comment="to_ISP1_primary_via_WAN1"
add check-gateway=ping distance=2 dst-address=0.0.0.0/0 gateway=208.67.220.220 routing-table=to_ISP1 scope=10 target-scope=11 comment="to_ISP1_backup_via_WAN2"
add check-gateway=ping distance=1 dst-address=0.0.0.0/0 gateway=64.6.65.6 routing-table=to_ISP2 scope=10 target-scope=11 comment="to_ISP2_primary_via_WAN2"
add check-gateway=ping distance=2 dst-address=0.0.0.0/0 gateway=208.67.222.222 routing-table=to_ISP2 scope=10 target-scope=11 comment="to_ISP2_backup_via_WAN1"
```

To make the main table active-backup instead of ECMP:

```
/ip route set [find where comment="MAIN_RECURSIVE_WAN1"] distance=1
/ip route set [find where comment="MAIN_RECURSIVE_WAN2"] distance=2
```

If the router has a public management NIC and you need to keep SSH/WinBox reachable through that public interface, consider omitting the `MAIN_RECURSIVE_WAN1` and `MAIN_RECURSIVE_WAN2` defaults during the lab. The policy tables still exercise PCC and failover for LAN traffic, while the `main` table keeps the cloud provider's management default route.

The main-table recursive defaults exist for router-originated traffic — the router's own DNS resolver, NTP, package downloads, DDNS updates, pings — which uses `main` because it does not enter the LAN `prerouting` chain. Equal distances make these defaults ECMP; different distances make them active-backup. Static defaults, DHCP-added defaults, or deliberate output policy routing can fill the same role if ECMP is not the right shape for your design.


### 6.7 DHCP clients with route-update scripts

Add the DHCP clients after creating the route placeholders above.

```
/ip dhcp-client
add interface=ether1 add-default-route=no use-peer-dns=no use-peer-ntp=no comment="WAN1 DHCP - updates WAN1_RECURSIVE routes"
add interface=ether2 add-default-route=no use-peer-dns=no use-peer-ntp=no comment="WAN2 DHCP - updates WAN2_RECURSIVE routes"
```

Set the WAN1 DHCP script:

```
/ip dhcp-client
set [find where interface=ether1] script={
    :if ([:tobool $bound]) do={
        :local gw $"gateway-address";
        /ip/route/set [find where comment="WAN1_RECURSIVE"] gateway=$gw disabled=no;
        :log info ("WAN1 DHCP bound; recursive routes now use " . $gw);
    } else={
        /ip/route/set [find where comment="WAN1_RECURSIVE"] disabled=yes;
        :log warning "WAN1 DHCP unbound; WAN1 recursive routes disabled";
    }
    /ip/firewall/connection/remove [find where connection-mark="ISP1_conn"];
}
```

Set the WAN2 DHCP script:

```
/ip dhcp-client
set [find where interface=ether2] script={
    :if ([:tobool $bound]) do={
        :local gw $"gateway-address";
        /ip/route/set [find where comment="WAN2_RECURSIVE"] gateway=$gw disabled=no;
        :log info ("WAN2 DHCP bound; recursive routes now use " . $gw);
    } else={
        /ip/route/set [find where comment="WAN2_RECURSIVE"] disabled=yes;
        :log warning "WAN2 DHCP unbound; WAN2 recursive routes disabled";
    }
    /ip/firewall/connection/remove [find where connection-mark="ISP2_conn"];
}
```

If your terminal rejects the multiline script form, paste the script body into the DHCP client script field in WinBox/WebFig, or convert it to a quoted one-liner with escaped quotes.

After applying the scripts, renew the leases:

```
/ip dhcp-client renew [find where interface=ether1]
/ip dhcp-client renew [find where interface=ether2]
```

The DHCP route-update script disables placeholder routes during any unbound transition, including the brief unbound state that occurs during initial paste/import. Its first run can therefore leave the placeholders disabled even when the leases bind moments later. Bootstrap once, unconditionally, after first apply:

```
# Replace these gateways with the values shown by /ip dhcp-client print detail.
/ip route set [find where comment="WAN1_RECURSIVE"] gateway=<WAN1_DHCP_GATEWAY> disabled=no
/ip route set [find where comment="WAN2_RECURSIVE"] gateway=<WAN2_DHCP_GATEWAY> disabled=no
```

Verify:

```
/ip dhcp-client print detail
/ip route print detail where comment~"WAN1_RECURSIVE|WAN2_RECURSIVE"
```

Future DHCP bind/unbind events update the routes via the script; this manual bootstrap is needed only once.

Each script clears only its own ISP's connection marks. A WAN1 lease change (new public IP, recovery after failure, etc.) leaves stale `ISP1_conn` flows pinned to the old path or old NAT public IP, so they are flushed and re-classified. Healthy `ISP2_conn` flows on the still-up WAN2 are not disturbed.


### 6.8 Mangle rules for inbound symmetry, PCC, and routing marks

Order matters. Keep the local-traffic accept rule before the PCC rules. `add` appends, so paste this section into a router with no existing prerouting mangle rules — or reorder afterward with `move`.

The chain=output rules below intentionally do **not** filter on `connection-state=new`. They must mark every reply packet whose connection-mark matches; replies on already-established flows are not in state `new`, and adding that condition will break per-WAN reply symmetry.

