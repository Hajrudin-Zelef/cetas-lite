---
id: collect-260926-mikrotik/mikrotik/github-whilcayangyang-mikrotik-homelab-config-a-production-grade-routeros-7-configuration--1
title: "Route to the VPN endpoint itself must go via the real WAN gateway, not the tunnel:"
domain: mikrotik
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/lot-mikrotik/RouterOS/github-whilcayangyang-mikrotik-homelab-config-a-production-grade-routeros-7-configuration-for-a-segm.md
source_anchor: ""
source_lines: [1, 127]
sha256: 4bbc4353b9d43ffb9ee8653c5a33a73a4d7291ef9fb728c06d192994dda57af1
---

# Route to the VPN endpoint itself must go via the real WAN gateway, not the tunnel:

A production-grade RouterOS 7 configuration for a segmented homelab network: VLAN isolation, per-device ProtonVPN policy-based routing (kill-switched), hardened admin access with brute-force lockout, and QoS.

**Disclaimer:** All IP addresses, MAC addresses, hostnames, domains, keys, and credentials in this README are **placeholder examples**. Replace every `<...>` and example value with your own before deploying. Do not reuse the example WireGuard keys or AWS credentials shown here — they are illustrative only.


| Item | Example Value | Notes | 
|---|---|---|
| RouterOS version | `7.21.4` | Stable/long-term channel recommended for production | 
| Model | `C53UiG+5HPaxD2HPaxD` | e.g. MikroTik hAP ax³ or similar dual-band Wi-Fi 6 router | 
| Update channel | `long-term` | Avoid `testing` /`current` on a router you depend on | 

**Note:** Serial number and software ID are omitted intentionally — never publish these, they can be used for warranty/support impersonation.


| VLAN | Name | Subnet (example) | Purpose | 
|---|---|---|---|
| 32 | `vlan32-private` | `172.31.62.0/26` | Trusted personal devices (also doubles as MGMT-reachable) | 
| 33 | `vlan33-server` | `172.31.62.128/26` | Homelab / server tier (K8s, VMs, sandboxes) | 
| 34 | `vlan34-mgmt` | `172.31.62.192/26` | Out-of-band management network | 
| 35 | `vlan35-public` | `172.31.62.64/26` | Guest / untrusted devices, isolated from LAN | 

```
                 ┌─────────────┐
   WAN (ether2) ─┤   MikroTik  │
                 │   Router    │
                 └──────┬──────┘
                        │ bridge1 (VLAN-filtering)
        ┌───────┬───────┼───────┬────────┐
     VLAN32   VLAN33  VLAN34  VLAN35   WireGuard
    (private) (server) (mgmt) (public)  (wg1/wg2 -> ProtonVPN)
```
A single bridge (`bridge1`) with VLAN filtering enabled carries all internal traffic; physical/Wi-Fi ports are assigned a **PVID** to tag untagged traffic into the correct VLAN.

```
/interface bridge
add name=bridge1 vlan-filtering=yes
/interface vlan
add interface=bridge1 name=vlan32-private vlan-id=32
add interface=bridge1 name=vlan33-server  vlan-id=33
add interface=bridge1 name=vlan34-mgmt    vlan-id=34
add interface=bridge1 name=vlan35-public  vlan-id=35
/interface bridge port
add bridge=bridge1 interface=ether1 pvid=33 frame-types=admit-only-untagged-and-priority-tagged
add bridge=bridge1 interface=ether3 pvid=32 frame-types=admit-only-untagged-and-priority-tagged
add bridge=bridge1 interface=ether4 pvid=32 frame-types=admit-only-untagged-and-priority-tagged
add bridge=bridge1 interface=wifi1  pvid=32 frame-types=admit-only-untagged-and-priority-tagged
add bridge=bridge1 interface=wifi2  pvid=35 frame-types=admit-only-untagged-and-priority-tagged
```
**Caution:** `frame-types=admit-only-untagged-and-priority-tagged` silently drops any tagged frame arriving on that port that doesn't match the PVID. This is intentional here (prevents VLAN-hopping from a compromised endpoint) but will break trunking if you later connect a switch expecting tagged traffic on that port.


Interface lists group physical/logical interfaces for use in firewall rules instead of hardcoding interface names everywhere.

```
/interface list
add name=LAN
add name=WAN
add name=MGT
/interface list member
add interface=ether2           list=WAN
add interface=vlan32-private   list=LAN
add interface=vlan33-server    list=LAN
add interface=vlan34-mgmt      list=LAN
add interface=vlan35-public    list=LAN
add interface=wg1              list=WAN
add interface=wg2              list=WAN
add interface=lo               list=LAN
add interface=vlan32-private   list=MGT
add interface=vlan34-mgmt      list=MGT
```
**Note:** WireGuard interfaces (`wg1`, `wg2`) are added to the `WAN` list. This matters for firewall rules like "drop non-DSTNAT WAN input" — traffic arriving from the VPN tunnel is treated the same as internet-facing WAN traffic for forwarding purposes.


Two SSIDs are defined: a WPA3-only network for trusted devices, and a WPA2/WPA3-mixed guest network (for device compatibility) mapped to the isolated public VLAN.

```
/interface wifi security
add name="Home-Trusted" authentication-types=wpa3-psk encryption=ccmp,gcmp \
    management-protection=required management-encryption=gmac-256 wps=disable disabled=no
add name="Home-Guest"   authentication-types=wpa2-psk,wpa3-psk encryption=ccmp,gcmp \
    wps=disable disabled=no
/interface wifi configuration
add name=cfg1 mode=ap ssid="Home-Trusted" security="Home-Trusted" \
    datapath.client-isolation=yes \
    channel.band=5ghz-ax channel.frequency=5785 channel.width=20/40/80mhz \
    channel.skip-dfs-channels=disabled country=<your-country> max-clients=8 disabled=no
add name=cfg2 mode=ap ssid="Home-Guest" security="Home-Guest" \
    datapath.client-isolation=yes \
    channel.band=2ghz-ax channel.frequency=2437,2462 channel.width=20mhz \
    channel.skip-dfs-channels=10min-cac country=<your-country> max-clients=12 disabled=no
/interface wifi
set [ find default-name=wifi1 ] configuration=cfg1 configuration.mode=ap disabled=no
set [ find default-name=wifi2 ] configuration=cfg2 configuration.mode=ap disabled=no
```
**Warning:** Use WPA3-only (`wpa3-psk`) wherever every client supports it. Only fall back to mixed WPA2/WPA3 for a legacy-device SSID, and put that SSID on an isolated/untrusted VLAN (as done here with `vlan35-public`) so a downgrade attack on the weaker network doesn't expose the trusted tier.


**Note:** `management-protection=required` (PMF/802.11w, mandatory under WPA3) and `wps=disable` should be set on every SSID — WPS is a well-known brute-force vector regardless of the underlying authentication type. `datapath.client-isolation=yes` blocks client-to-client traffic within each SSID, which matters most on the guest network but costs nothing on the trusted one either. `cfg2`'s `channel.frequency=2437,2462` gives RouterOS a list of allowed 2.4GHz channels to auto-select from (channels 6 and 11) rather than pinning to one, avoiding manual retuning if a neighbor collides on the same channel.


| Pool | Range (example) | Bound to | 
|---|---|---|
| `vlan32-private-pool` | `172.31.62.2 – 172.31.62.62` | `vlan32-private` | 
| `vlan35-public-pool` | `172.31.62.66 – 172.31.62.126` | `vlan35-public` | 
| `vlan33-server-pool` | `172.31.62.130 – 172.31.62.190` | `vlan33-server` | 

```
/ip pool
add name=vlan32-private-pool ranges=172.31.62.2-172.31.62.62
add name=vlan35-public-pool  ranges=172.31.62.66-172.31.62.126
add name=vlan33-server-pool          ranges=172.31.62.130-172.31.62.190
/ip dhcp-server
add name=dhcp-vlan32-private address-pool=vlan32-private-pool interface=vlan32-private
add name=dhcp-vlan35-public  address-pool=vlan35-public-pool  interface=vlan35-public
add name=dhcp-vlan33-server  address-pool=vlan33-server-pool          interface=vlan33-server
/ip dhcp-server network
add address=172.31.62.0/26   gateway=172.31.62.1   dns-server=10.2.0.1
add address=172.31.62.64/26  gateway=172.31.62.65  dns-server=172.31.62.65
add address=172.31.62.128/26 gateway=172.31.62.129 dns-server=1.1.1.2,1.0.0.2
```
**Note:** Each VLAN's DHCP-issued DNS server differs on purpose:


`vlan32-private` clients are pointed at `10.2.0.1` — the DNS server *inside the WireGuard tunnel* (ProtonVPN-provided DNS), so trusted clients' DNS queries are also tunneled and don't leak.
`vlan35-public` clients resolve via the router itself (`172.31.62.65`), which is redirected to a local resolver (see DNS redirect NAT rule below).
`vlan33-server` clients use Cloudflare's malware-filtering resolver (`1.1.2.2`/`1.0.0.2`) directly, since server-tier traffic is not policy-routed through the VPN.

