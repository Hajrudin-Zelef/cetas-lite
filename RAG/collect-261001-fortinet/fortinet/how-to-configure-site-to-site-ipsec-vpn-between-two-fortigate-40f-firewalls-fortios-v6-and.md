---
id: collect-261001-fortinet/fortinet/how-to-configure-site-to-site-ipsec-vpn-between-two-fortigate-40f-firewalls-fortios-v6-and
title: "how-to-configure-site-to-site-ipsec-vpn-between-two-fortigate-40f-firewalls-fortios-v6-and-v7-gkhan"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/how-to-configure-site-to-site-ipsec-vpn-between-two-fortigate-40f-firewalls-fortios-v6-and-v7-gkhan.md
source_anchor: ""
source_lines: [1, 263]
sha256: fe39fd4f3ce48e46f1614a38225f5a9eb85ec6e5e4aaae6828638a2c680c3b26
---

# how-to-configure-site-to-site-ipsec-vpn-between-two-fortigate-40f-firewalls-fortios-v6-and-v7-gkhan

Connecting branch offices or partners through a secure VPN tunnel is one of the most common FortiGate tasks.

This guide shows how to configure a **manual IPsec site-to-site VPN** between two FortiGate 40F units running different firmware versions:

| Site | Model | Firmware | 
|---|---|---|
| **Site B** | FortiGate-40F | v7.6.3, build 3510 (GA.F) | 
| **Site C** | FortiGate-40F | v6.4.7, build 8726 (GA) | 

## 🧩 Network Plan

| Parameter | Site B | Site C | 
|---|---|---|
| WAN IP | 172.16.10.109 | 172.16.10.111 | 
| LAN Subnet | 192.168.20.0/24 | 192.168.30.0/24 | 
| Tunnel Name | vpn-to-SiteC | SiteB | 
| Pre-Shared Key | Forti@123 | Forti@123 | 

## ⚙️ Step 1 – Create Phase 1 Interface (IKE Negotiation)

### 📍 Site B ( FortiOS v7.x )

```
config vpn ipsec phase1-interface
    edit "vpn-to-siteC"
        set interface "wan"                 # WAN interface carrying IKE
        set ike-version 2                   # IKEv2 (default in v7)
        set peertype any
        set net-device disable
        set proposal aes256-sha256          # Encryption / authentication
        set dhgrp 14                        # Diffie-Hellman group
        set transport auto                  # <— New in FortiOS 7.x, auto-selects UDP 500/4500
        set remote-gw 172.16.10.111         # Peer’s WAN IP
        set psksecret "Forti@123"           # Shared key
    next
end
```
### 📍 Site C ( FortiOS v6.x )

```
config vpn ipsec phase1-interface
    edit "SiteB"
        set interface "wan"
        set ike-version 2
        set peertype any
        set net-device disable
        set proposal aes256-sha256
        set dhgrp 14
        # ❌ 'set transport' not available in v6.x
        set remote-gw 172.16.10.109
        set psksecret "Forti@123"
    next
end
```
📝 **Explanation:**

- *Phase 1* defines how the two peers authenticate and exchange encryption keys.
- `ike-version 2` → modern and more stable.
- `dhgrp` must match on both sides.
- `transport` (v7 only) automatically handles NAT-Traversal; in v6 FortiGate does it by default when needed.

## ⚙️ Step 2 – Create Phase 2 Interface (Traffic Selectors)

### 📍 Site B (v7)

```
config vpn ipsec phase2-interface
    edit "Site_C"
        set phase1name "vpn-to-siteC"
        set proposal aes256-sha256
        set src-subnet 192.168.20.0 255.255.255.0
        set dst-subnet 192.168.30.0 255.255.255.0
    next
end
```
### 📍 Site C (v6)

```
config vpn ipsec phase2-interface
    edit "Site_B"
        set phase1name "SiteB"
        set proposal aes256-sha256
        set src-subnet 192.168.30.0 255.255.255.0
        set dst-subnet 192.168.20.0 255.255.255.0
    next
end
```
📝 **Explanation:**

- *Phase 2* defines which networks are encrypted inside the tunnel.
- Subnets must mirror each other (local ↔ remote).

## ⚙️ Step 3 – Add Firewall Policies

Both sides need **two policies**:

- LAN → VPN
- VPN → LAN

### 📍 Site B (v7)

```
config firewall policy
    edit 0
        set name "LAN-to-SiteC"
        set srcintf "lan"
        set dstintf "vpn-to-siteC"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set nat disable
    next
    edit 0
        set name "SiteC-to-LAN"
        set srcintf "vpn-to-siteC"
        set dstintf "lan"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set nat disable
    next
end
```
### 📍 Site C (v6)

```
config firewall policy
    edit 0
        set name "LAN-to-SiteB"
        set srcintf "lan"
        set dstintf "SiteB"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set nat disable
    next
    edit 0
        set name "SiteB-to-LAN"
        set srcintf "SiteB"
        set dstintf "lan"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set nat disable
    next
end
```
📝 **Explanation:**

- NAT must be **disabled** so that private subnets pass unmodified.
- Policies are stateful—responses automatically return.

## ⚙️ Step 4 – Add Static Routes

Each site needs a route to the opposite LAN through the VPN interface.

### 📍 Site B

```
config router static
    edit 0
        set dst 192.168.30.0 255.255.255.0
        set device "vpn-to-siteC"
    next
end
```
### 📍 Site C

```
config router static
    edit 0
        set dst 192.168.20.0 255.255.255.0
        set device "SiteB"
    next
end
```
## ⚙️ Step 5 – Bring Up and Verify the Tunnel

### Check tunnel summary

```
get vpn ipsec tunnel summary
```
### Detailed info

```
diagnose vpn tunnel list
```
### Force tunnel initiation

```
diagnose vpn tunnel up name vpn-to-siteC
```
### Test reachability

```
execute ping-options source 192.168.20.1
execute ping 192.168.30.1
```
## 🧠 Step 6 – Debug If Tunnel Fails

```
diagnose debug reset
diagnose vpn ike log-filter clear
diagnose vpn ike log-filter dst-addr4 172.16.10.111
diagnose debug application ike -1
diagnose debug enable
```
Look for:

- `no proposal chosen` → mismatch in Phase 1/2 encryption.
- `AUTHENTICATION_FAILED` → wrong pre-shared key.
- `invalid id information` → subnet mismatch.

Stop debug:

```
diagnose debug disable
```
## 🧰 Optional: Stability Enhancements (v7 and v6)

```
config vpn ipsec phase1-interface
    edit "vpn-to-siteC"
        set dpd on-idle
        set keylife 28800
        set keepalive enable
    next
end
```
## 📋 Version Differences Quick Reference

| Command | v7.x | v6.x | Description | 
|---|---|---|---|
| `set transport auto` | ✅ | ❌ | Controls NAT-T / UDP encapsulation (v7+) | 
| `set ike-version 2` | ✅ | ✅ | Enables IKEv2 | 
| `set net-device disable` | ✅ | ✅ | Use interface-mode VPN | 
| `set dpd on-idle` | ✅ | ✅ | Dead Peer Detection | 
| `set psksecret "..."` | ✅ | ✅ | Shared key | 
| `set proposal aes256-sha256` | ✅ | ✅ | Cipher / hash pair | 

## 🧾 Final Verification Checklist

| Test | Command | Expect | 
|---|---|---|
| Tunnel status | `get vpn ipsec tunnel summary` | up/active | 
| Routing table | `get router info routing-table all` | Remote LAN via VPN | 
| Ping test | `execute ping 192.168.30.1` | Success | 
| Log | `diagnose vpn tunnel list` | Established = yes | 

### ✅ Summary

With these CLI commands, any engineer—junior or senior—can configure a secure, working **site-to-site IPsec VPN** between different FortiGate firmware versions.

The key is to **mirror Phase 1/2 parameters**, ensure **no NAT**, and verify routing and policies.
