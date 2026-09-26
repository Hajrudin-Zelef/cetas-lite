---
id: collect-260926-mikrotik/mikrotik/how-to-establish-site-to-site-vpn-with-mikrotik-routers-ea6f3f95
title: "How to establish site to site VPN with Mikrotik routers"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/how-to-establish-site-to-site-vpn-with-mikrotik-routers-ea6f3f95.md
source_anchor: ""
source_lines: [1, 51]
sha256: d2d036ac5a0c943809e7ddd900b5fc0f292b7eddd22adaf1581f43499e994900
---

# How to establish site to site VPN with Mikrotik routers

*Source : https://www.informaticar.net/how-to-establish-site-to-site-vpn-with-mikrotik-routers/*

Quick tutorial: IPSec Site-to-Site VPN tunnel with MikroTik RB RouterOS 6.46.1 on both sides. Pre-shared key (lab only — for production, use certificate-based IPSec tunnels, not PSK). Clean config, no default MikroTik config preloaded; only one NAT rule (srcnat masquerade).

## Lab setup

- **Office 1:** Router 1 WAN `192.168.155.131/24`, gateway `192.168.155.2/24`, LAN `10.50.50.0/24`, PC1 `10.50.50.2/24`
- **Office 2:** Router 25 WAN `192.168.155.130/24`, gateway `192.168.155.2/24`, LAN `192.168.11.0/24`, PC2 `192.168.11.2/24`

Goal: connect Office 1 LAN with Office 2 LAN via IPSec. Steps shown for Office 1, mirrored on Office 2.

## Peers

IP > IPSec > Peers > Plus (+). Name: `Router2`, Address: Office 2 WAN `192.168.155.130`, Exchange Mode: **IKE2**. Rest default.
(Office 2: name `Router1`, address `192.168.155.131`, IKE2.)

## Identities

IP > IPSec > Identities > Plus (+). Peer: `Router2`, Authentication Method: pre shared key, Secret: password (same on both sides; in production use >20 chars with letters, numbers, special characters). Rest default.

## Proposals

IP > IPSec > Proposals > edit `*default`: Auth. Algorithms `sha256`, Encr. Algorithms `aes-256 cbc`, lifetime 30 minutes, PFS Group `modp2048`. Same on both sides.

## Profiles

IP > IPSec > Profiles > edit `*default`: Hash Algorithms `sha256`, Encryption `aes-256`, DH Group `modp2048`, Proposal check `obey`, lifetime `1day`, NAT Traversal checked, DPD Maximum Failure 5. Same on both sides.

## Policies

IP > IPSec > Policies > Plus (+). General tab: Peer `Router2`, check **Tunnel**, Src. Address `10.50.50.0/24` (local LAN), Dst. Address `192.168.11.0/24` (remote LAN). Action tab: defaults.
(Office 2: Peer `Router1`, Src. Address `192.168.11.0/24`, Dst. Address `10.50.50.0/24`.)

When done on both sides: Policy screen shows **PH2 State: established**, Active Peers populated, Installed SAs created.

If the tunnel won't establish: double-check IPSec settings (Peer, Proposals, Secret, Profile, Policy), routing, gateway, NAT and firewall (ports 500, 4500, protocol 50 may need to pass).

## Fixing traffic through the tunnel

Tunnel was established but no traffic passed (ping between PC1 `10.50.50.2` and PC2 `192.168.11.x` failed). Fix — IP > Firewall > NAT tab > Plus (+):

- Chain: `srcnat`, Src. Address: local subnet (`10.50.50.0/24`), Dst. Address: remote subnet (`192.168.11.0/24`)
- Action tab: **Accept**

(Office 2 mirrored: Src. Address `192.168.11.0/24`, Dst. Address `10.50.50.0/24`, action accept.)

**Important:** this NAT accept rule must be the FIRST rule in the NAT tab. Then **reboot both routers** — ping should start working, and traffic is visible via the NAT rule counter.

Working IPSec Site-to-Site tunnel done. Simple lab scenario (VM); in production ensure working public IPs and routes so routers can reach each other.
