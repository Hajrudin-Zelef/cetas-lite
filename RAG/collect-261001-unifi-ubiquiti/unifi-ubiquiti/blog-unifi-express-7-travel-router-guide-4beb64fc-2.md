---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-unifi-express-7-travel-router-guide-4beb64fc-2
title: "blog-unifi-express-7-travel-router-guide-4beb64fc"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: ["2026-07"]
keywords: ["consumer", "ethernet", "memory", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-unifi-express-7-travel-router-guide-4beb64fc.md
source_anchor: ""
source_lines: [38, 107]
sha256: 91766456b3377574948cab5e895eb89ee29f759507094528c91d91d27d7e26b6
---

# blog-unifi-express-7-travel-router-guide-4beb64fc

The supplier's IT team authorized our client to use a dedicated ethernet wall jack in a conference room, provisioned on an isolated guest VLAN that the supplier controlled. This is a critical point: never connect a router to a host network without explicit written authorization from the site's IT team. Even though the UX7 uses NAT to separate its LAN from the upstream network, traffic still traverses the host infrastructure, and the UX7 may be able to reach upstream private addresses unless either side blocks them.
Our solution, deployed with the supplier's IT approval:
- Configure a UniFi Express 7 as a standalone gateway
- Connect its WAN port to the supplier-provided ethernet jack on their guest VLAN
- Enable Site Magic SD-WAN to create an encrypted tunnel back to the client's headquarters
- Configure UX7 firewall rules to block traffic to the supplier's private upstream subnets
- The team connects to the UX7's WiFi network, accessing headquarters resources through the Site Magic tunnel
The supplier's IT department verified and approved the setup because the UX7 sat behind their guest VLAN, our firewall rules prevented lateral access, and all headquarters-bound traffic was encrypted through the Site Magic tunnel. (By default, Site Magic carries traffic destined for selected remote networks — not all public internet traffic. If full-tunnel routing is required, policy-based routing must be configured separately.) The client's team maintained full access to internal systems, file servers, and applications without exposing credentials over untrusted networks.
We now replicate this configuration for:
- Field service teams servicing client sites
- Temporary project offices at construction sites or event venues
- Remote staff working from coworking spaces where network security is unknown
- Trade show booths requiring secure connectivity for demonstrations
Technical Specifications: What You're Working With
The UX7 is a 443-gram, USB-C powered cloud gateway with one 10 GbE WAN port, one 2.5 GbE LAN port and tri-band WiFi 7. Full specifications follow.
Core Specifications
| Component | Specification | 
|---|---|
| Price | $199 (Ubiquiti Store, July 2026) | 
| Dimensions | 117 × 117 × 42.5 mm (4.6" × 4.6" × 1.7") | 
| Weight | 443 g (1 lb) | 
| WiFi | WiFi 7 (802.11be), 6-stream, tri-band (BE10700 class) | 
| Wireless Speed | Up to 10.7 Gbps aggregated (5.7 Gbps @ 6GHz, 4.3 Gbps @ 5GHz, 688 Mbps @ 2.4GHz) | 
| Coverage | Up to 160 m² / 1,750 sq ft (Ubiquiti vendor estimate — real-world coverage varies with environment) | 
| WAN Port | 1× 10GBASE-T RJ45 (10 Gbps) | 
| LAN Port | 1× 2.5 GbE RJ45 (2.5 Gbps) | 
| Component | Specification | 
|---|---|
| Power | USB-C (5V DC, 5A / 25W adapter included) | 
| Max Power Draw | 22W peak (Ubiquiti spec; we have not published measured idle/load figures) | 
| Processor | Quad-core ARM Cortex-A53 @ 1.5 GHz | 
| Memory | 3 GB RAM | 
| Capacity | 300+ clients, 30+ UniFi devices | 
| VPN | Site-to-site: Site Magic, IPsec, OpenVPN; Client/Server: WireGuard, L2TP; Remote access: Teleport (WiFiman) | 
| Management | UniFi Network application only | 
Power Warning: The 5V/5A Challenge
The UX7 requires 5V at 5A (25W), which is a non-standard USB-C PD profile. Most power banks jump to 9V or 15V for high wattage. If your power bank doesn't support PPS (Programmable Power Supply) or a high-amperage 5V rail, the UX7 may fail to boot or experience power cycling.
For field deployments, either use the included wall adapter or verify your power bank's voltage/amperage profiles before relying on battery operation.
What the 10 GbE WAN Port Does — and Does Not — Mean
The UX7 has a 10 GbE WAN connector, but that does not make it a "10G field office." Its only wired LAN port is 2.5 GbE, IDS/IPS throughput is rated at 2.3 Gbps, VPN performance is not officially published (and is likely well below 2.3 Gbps), and the 10.7 Gbps WiFi figure is an aggregate theoretical total across three radios — not a single-client result. The 10 GbE WAN port provides headroom for fast ISP uplinks and future-proofing, but real-world field throughput is bounded by the 2.5 GbE LAN port, IDS/IPS processing, and VPN overhead.
What Makes It Travel-Friendly
Compact for a Cloud Gateway: At 443 grams and 117 mm square, the UX7 is larger and heavier than consumer travel routers like the GL.iNet Beryl 7 (205 g) or Slate 7 (295 g), but it is compact for a full Cloud Gateway with IDS/IPS and tri-band WiFi 7. It fits in a laptop bag alongside a power bank and cables.
USB-C power: The included adapter is 5V/5A (25W). Battery operation is possible but has not been independently tested — we recommend using the included wall adapter for reliable operation and testing any power bank thoroughly before depending on it in the field. See the power warning below.
Fanless design: Silent operation means you can run the UX7 in conference rooms, hotel rooms, or quiet office spaces without generating noise complaints.
No external antennas: The internal antenna design means nothing to break or lose during transport. The compact square form factor is laptop-bag friendly.
Limitations to Understand
Single ethernet port: The UX7 has one WAN and one LAN port. If you need to connect multiple wired devices, you'll need a separate switch (we address this in the equipment recommendations).
No PoE output: Unlike the Cloud Gateway Fiber or Dream Router 7, the UX7 cannot power devices via Power over Ethernet. All connected equipment needs independent power.
Network app only: The UX7 runs only the UniFi Network application. It cannot host Protect (cameras), Talk (phones), or Access (door locks). This is a gateway-only device, which is appropriate for travel use.
WiFi coverage limits: Ubiquiti rates the UX7 at 160 m² / 1,750 sq ft, but real-world coverage varies with walls, interference and client hardware. It is best suited for single-room deployments or small temporary offices where everyone works within WiFi range.
Complete Equipment Setup: Two Configurations
Two kits cover most field deployments: a basic $210 setup for one to three people, and a $310 professional setup that adds wired ports. Accessory prices are approximate and change often.
Configuration 1: Basic Travel Router (about $210 total)
Scenario: Single user or small team (2-3 people) working from client sites with available ethernet and power outlets.
Equipment:
- UniFi Express 7 - $199
- Included USB-C power adapter
- 6-foot ethernet cable (for WAN connection) - ~$10
About battery operation: The UX7 requires 5V at 5A (25W) — an unusual USB-C profile. Most power banks (including popular 30W models like the Anker 20K) deliver only 5V/3A at the 5V rail and reach higher wattages by stepping up to 9V or 15V, which the UX7 will not accept. Use the included wall adapter unless you have tested and confirmed a specific power bank can supply 5V/5A continuously. We have not yet identified and tested a verified-compatible power bank to recommend.
What you get: Portable gateway with gateway-level security, WiFi 7 connectivity for multiple devices, and wall-powered operation.
Setup time: 10-15 minutes including initial UniFi console configuration.
Configuration 2: Professional Field Setup (about $310 total)
Scenario: Field service teams, project offices, or situations requiring multiple wired connections.
Equipment:
- UniFi Express 7 - $199
- UniFi Flex Mini 2.5G Switch - $49 (5-port 2.5 GbE switch)
- USB-C power cables (2×) - about $24
- Cat6 ethernet cables (3-foot and 6-foot) - ~$15 (Cat6 is sufficient for short 2.5 GbE runs; Cat6A is unnecessary at these lengths)
- Small carrying case - ~$15
What you get: Expanded wired connectivity (4 additional device ports — the fifth port on the Flex Mini serves as the uplink to the UX7), 2.5G throughput to connected devices, and organized transport solution. Both the UX7 and switch require wall power or verified-compatible USB-C power sources.
