---
id: collect-261001-general-networking/general-networking/questions-1541989-connecting-nintendo-switches-over-site-to-site-vpn-a6237bc1
title: "Connecting Nintendo Switches over Site-to-Site VPN"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1541989-connecting-nintendo-switches-over-site-to-site-vpn-a6237bc1.md
source_anchor: ""
source_lines: [1, 11]
sha256: 61884df780c79943f2845fbdf9e1030cd7264ee477e5bc983488a88b99a8d875
---

# Connecting Nintendo Switches over Site-to-Site VPN

*Score : 0 | Source : https://superuser.com/questions/1541989/connecting-nintendo-switches-over-site-to-site-vpn*

I have two sites connected with an OpenVPN tunnel on which I am trying to get two Nintendo Switch devices (one at each site) to see eachother so that they can game as if they're on the same LAN. All TCP / UDP traffic between the two devices is allowed and flowing correctly, so I believe the issue is unrouted multicast traffic. I suspect that I need to somehow create a virtual IP on each network with a 1:1 NAT that forwards traffic to the device on the other network but none of the combinations I've tried have worked so far. One site is running OPNSense on its edge firewall, the other PFSense, fully updated.

---

### Reponse (acceptee) — score 0

For whatever it's worth, the problem I thought I was solving, IGMP/mDNS traffic not crossing the VPN was solved by converting my OpenVPN TUN (Layer 3) connection to a TAP (Layer 2) connection and adding IGMP and mDNS proxy services to both firewalls to push the traffic. I later found out that the switch was not using the LAN at all and was using Nintendo's proprietary "Local Connection" (read Bluetooth) to host rooms for Mario Party so cross site play was not possible. As a plus(?) I can now see all the chromecasts on the other network from here.
