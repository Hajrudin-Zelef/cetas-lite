---
id: collect-260926-mikrotik/mikrotik/github-whilcayangyang-mikrotik-homelab-config-a-production-grade-routeros-7-configuration--3
title: "Route to the VPN endpoint itself must go via the real WAN gateway, not the tunnel:"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["sandbox"]
source: docs/RAG/lot-mikrotik/RouterOS/github-whilcayangyang-mikrotik-homelab-config-a-production-grade-routeros-7-configuration-for-a-segm.md
source_anchor: ""
source_lines: [247, 339]
sha256: 35108b8580e3a2bc181ffe1fb81edd76a995c8d6450ec1ba353519c22b37d49e
---

# Route to the VPN endpoint itself must go via the real WAN gateway, not the tunnel:

| # | Action | Match | Purpose | 
|---|---|---|---|
| 1.1 | accept | `connection-state=established,related,untracked` | Fast-path return traffic | 
| 1.2 | drop | `connection-state=invalid` | Reject malformed/invalid packets | 
| 1.3 | accept | `protocol=icmp` | Allow ping for diagnostics | 
| 1.4 | accept | `dst-address=127.0.0.1` | Loopback | 
| 1.5 | jump | → `management` chain | Route admin-access checks | 
| 1.6 | drop | `in-interface-list=!LAN` | Drop anything not from a defined LAN interface | 
| 1.7 | drop | (all) | Explicit final deny | 

```
/ip firewall filter
add chain=input action=accept connection-state=established,related,untracked comment="1.1 ACCEPT: Established/Related/Untracked"
add chain=input action=drop   connection-state=invalid                       comment="1.2 DROP: Invalid Packets"
add chain=input action=accept protocol=icmp                                  comment="1.3 ACCEPT: ICMP (Ping)"
add chain=input action=accept dst-address=127.0.0.1                          comment="1.4 ACCEPT: Loopback"
add chain=input action=jump jump-target=management                          comment="1.5 JUMP: Check Management Access"
add chain=input action=drop   in-interface-list=!LAN                         comment="1.6 DROP: Non-LAN Input"
add chain=input action=drop                                                  comment="1.7 DROP: Everything Else"
```
Implements a 3-stage progressive blacklist for SSH/HTTPS/Winbox login attempts, similar in spirit to fail2ban.

```
/ip firewall address-list
add list=allowed_to_router address=<192.168.88.0/30>   comment="mgmt-p2p-link"
add list=allowed_to_router address=<172.31.62.192/26>  comment="vlan34-mgmt"
add list=allowed_to_router address=<172.31.62.2-172.31.62.5> comment="trusted-devices"
add list=allowed_to_router address=<172.31.62.128/29>  comment="k8s-homelab"
/ip firewall filter
add chain=management action=accept src-address-list=allowed_to_router \
    comment="2.0 ALLOW: Whitelisted Admin IPs Skip Blacklist Check"
add chain=management action=drop src-address-list=admin-blacklist \
    comment="2.0b DROP: Blacklisted Brute-Force Source"
add chain=management action=add-src-to-address-list address-list=admin-blacklist address-list-timeout=1d \
    connection-state=new protocol=tcp dst-port=22,443,8291 src-address-list=admin-stage2 \
    comment="2.0c BLACKLIST: 3rd Failed Attempt in 5 Min"
add chain=management action=add-src-to-address-list address-list=admin-stage2 address-list-timeout=1m \
    connection-state=new protocol=tcp dst-port=22,443,8291 src-address-list=admin-stage1 \
    comment="2.0d STAGE2: 2nd Attempt in 1 Min"
add chain=management action=add-src-to-address-list address-list=admin-stage1 address-list-timeout=1m \
    connection-state=new protocol=tcp dst-port=22,443,8291 \
    comment="2.0e STAGE1: 1st Attempt"
add chain=management action=accept protocol=tcp dst-port=22,80,443,8291 src-address-list=allowed_to_router \
    comment="2.1 ALLOW: Admin Access (TCP)"
add chain=management action=accept protocol=udp dst-port=5678 src-address-list=allowed_to_router \
    comment="2.2 ALLOW: Admin Access (UDP, MNDP/neighbor discovery)"
add chain=management action=accept protocol=udp dst-port=53,67-68,123 src-address-list=lan-network \
    comment="2.3 ALLOW: DNS/DHCP/NTP (UDP)"
add chain=management action=accept protocol=tcp dst-port=53 src-address-list=lan-network \
    comment="2.4 ALLOW: DNS (TCP)"
add chain=management action=return comment="2.5 RETURN: To Input Chain (for Drop)"
```
**How the staged lockout works:** Any *new* connection attempt to an admin port (22/443/8291) from a source not already whitelisted gets added to `admin-stage1` (1 min TTL). A second attempt within that minute promotes the source into `admin-stage2` (also 1 min TTL). A third attempt while still in stage 2 promotes it to `admin-blacklist` for 1 day, at which point rule `2.0b` drops it outright. Legitimate admins should always be in `allowed_to_router` so they skip this entirely (rule `2.0`).


**Warning:** This throttles brute force but is not a substitute for strong authentication. Also enforce: disabled password-only root/admin login where possible, key-based SSH, and a non-default admin username.


```
/ip firewall filter
add chain=forward action=fasttrack-connection connection-mark=no-mark connection-state=established,related \
    comment="3.1 FASTTRACK: Standard Traffic (Excludes VPN)"
add chain=forward action=accept connection-state=established,related,untracked \
    comment="3.2 ACCEPT: Established/Related/Untracked"
add chain=forward action=drop connection-state=invalid \
    comment="3.3 DROP: Invalid Packets"
add chain=forward action=drop connection-nat-state=!dstnat connection-state=new in-interface-list=WAN \
    comment="3.4 DROP: WAN Garbage (Non-DSTNAT)"
add chain=forward action=drop src-address-list=no_forward_ipv4 \
    comment="3.5 DROP: Bad Public IPs (RFC6890/multicast/broadcast)"
add chain=forward action=jump jump-target=vlan-acl \
    comment="3.6 JUMP: Check VLAN ACLs"
add chain=forward action=jump jump-target=port-restriction \
    comment="3.7 JUMP: Check Port Restrictions"
add chain=forward action=drop \
    comment="3.8 DROP: Block Unknown Forwarding"
```
**Note:** Rule `3.1` deliberately excludes marked (VPN-policy-routed) connections from fasttrack — fasttracked packets bypass further firewall/mangle processing, which would break the kill-switch and PBR marking for VPN traffic. Only unmarked ("no-mark") traffic gets the fasttrack performance boost.


**Caution:** `3.5` (`no_forward_ipv4`) should reference an address-list containing RFC 6890 reserved ranges (`0.0.0.0/8`, `169.254.0.0/16`), multicast (`224.0.0.0/4`), and limited broadcast (`255.255.255.255`) — this is MikroTik's default `defconf` bogon list, not a full bogon feed. For internet-facing routers, consider adding a full bogon list.


| Rule | Allow | 
|---|---|
| 4.1 / 4.2 | `vlan32-private` ⇄`vlan33-server` | 
| 4.3 / 4.4 | `vlan33-server` ⇄`vlan34-mgmt` | 
| 4.5 / 4.6 | `vlan34-mgmt` ⇄`allowed_to_router` | 
| 4.7 / 4.7.1 | `vlan35-public` → homelab HTTP/S service only (specific host, ports 80/443) | 
| 4.8 / 4.8.1 | `non-vpn` list → homelab SSH/sandbox service only (specific host, port 22) | 
| 4.9 | Drop everything else between LAN networks | 
| 4.10 | Return to forward chain for internet-egress check | 

