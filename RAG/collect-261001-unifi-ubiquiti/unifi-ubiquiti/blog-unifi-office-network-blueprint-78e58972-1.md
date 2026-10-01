---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-unifi-office-network-blueprint-78e58972-1
title: "blog-unifi-office-network-blueprint-78e58972"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["AWS", "United States"]
dates: ["2026-08-10"]
keywords: ["cost", "disclosure", "ethernet", "license", "pricing", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-unifi-office-network-blueprint-78e58972.md
source_anchor: ""
source_lines: [1, 56]
sha256: e110dfcf23da6584c50af148286671c76821d4858acf4be00fbe65a901141f15
---

# blog-unifi-office-network-blueprint-78e58972

UniFi Network Design Blueprint: Complete 2026 Office Setup
Step-by-step guide to designing a complete UniFi office network for 2026. Covers Dream Machine Pro Max, multi-gigabit switching for WiFi 7, U7 access points, and G6 security cameras.
This blueprint walks through a complete UniFi office network using the current 2026 hardware lineup — from gateway and switch selection through VLAN design, zone-based firewall policies, and post-install validation. It targets offices with 20–80+ employees, wired and wireless workstations, VoIP, security cameras, and guest/IoT segmentation.
If your office has fewer than 15 employees or minimal segmentation needs, this Pro Max 48 PoE stack may be more than you need — see our UniFi Buyer's Guide for a right-sized starting point. For the full UniFi product ecosystem, this guide focuses on the planning and configuration decisions that matter most.
Prices checked August 10, 2026 against the Ubiquiti US store. All prices exclude tax and shipping and are subject to change, including tariff-related surcharges.
Affiliate Disclosure: This article contains affiliate links. If you make a purchase through these links, we may earn a small commission at no extra cost to you. As an Amazon Associate, iFeelTech earns from qualifying purchases.
Why Choose UniFi for Your Network Setup?
UniFi provides a single-vendor ecosystem where routers, switches, access points, and cameras are managed through related but distinct applications — the Network app for switching, routing and wireless, and Protect for cameras. Site Manager provides a centralized entry point across all consoles and sites. This reduces complexity compared to assembling separate vendor dashboards for each function.
Modular Expansion
UniFi's modular ecosystem allows granular scaling. You can daisy-chain additional Switch Pro Max units via SFP+ 10Gb uplinks. Wireless meshing with U7 Series access points is available as a fallback when wiring is not immediately feasible, though wired backhaul is recommended — mesh hops reduce available throughput on the shared radio.
Centralized Visibility
Network devices, APs, and gateways are managed through the Network application; cameras run in the Protect application on the same console. Cloud access via unifi.ui.com and the Site Manager portal enables remote management from anywhere.
Competitive Pricing with Trade-Offs
UniFi hardware delivers strong performance at price points below traditional enterprise vendors. The trade-off is a smaller support organization and tighter ecosystem dependency — cameras, switches, and gateways work best within the UniFi family. Optional UI Care provides expedited replacement, prepaid returns, and five-year protection for eligible devices at an additional per-device cost.
Essential Equipment for a 2026 Network Setup with UniFi
Choosing the right equipment starts with understanding your office layout, employee count, and which services (Wi-Fi, VoIP, cameras, guest access) you need. Pre-wire network cabling before installing devices — Cat6A for new builds, or verify that existing Cat5e/Cat6 supports the multigig speeds your APs and switches require.
Selecting specific products? Our UniFi Buyer's Guide provides recommendations by office size. For switch-specific guidance with PoE sizing, see our Best UniFi Switches guide.
Essential Equipment for 2026
Dream Machine Pro Max – Gateway with 5 Gbps IDS/IPS throughput, dual 3.5″ drive bays for NVR recording, built-in VPN/firewall, and Shadow Mode (VRRP) for high-availability failover (requires a second Pro Max with mirrored WAN/LAN connections).
UniFi Switch Pro Max 48 PoE – 48-port Layer 3 switch: 16 × 2.5GbE (8 PoE+, 8 PoE++) + 32 × 1GbE (24 PoE+, 8 PoE++), four 10G SFP+ uplinks, 720W total PoE budget, and Etherlighting for visual port management. A 2.5GbE PoE port prevents a 1GbE wired bottleneck on APs like the U7 Pro. The U7 Pro XG and XGS have 10GbE RJ45 uplinks — the Pro Max 48's 10G ports are SFP+ (no PoE), so these APs will negotiate at 2.5GbE on the switch's multigig PoE ports. For full 10G operation, you need either a switch with 10GbE RJ45 PoE output, or an SFP+-to-RJ45 transceiver paired with a 10G PoE adapter.
U7 Series Access Points – Wi-Fi 7 access points. The U7 Lite covers 2.4 GHz and 5 GHz only (no 6 GHz radio). The U7 Pro, U7 Pro Max, and U7 Pro XGS add the 6 GHz band. The U7 Pro XG is Ubiquiti's flagship model with a 10GbE RJ45 uplink.
UniFi G6 Series Cameras – 4K cameras with on-device AI. The G6 Bullet, Turret, and Dome include face recognition and license plate detection — all three are IP66 all-weather rated. The G6 Instant is Wi-Fi connected and USB-C powered (not Ethernet PoE), with AI capabilities included.
Cat6A Cables – Recommended for new builds to support 10GBASE-T up to 100 m. Existing Cat5e or Cat6 may support 2.5/5GbE over shorter runs — Cat6A is not a prerequisite for Wi-Fi 7, but it provides the most headroom.
Get Your Custom Equipment List
Not sure exactly what you need? Use our UniFi Network Configurator to build a personalized equipment list based on your office size, employee count, and security requirements. Configurator prices are estimates — verify current pricing at the Ubiquiti US Store before purchasing.
Sizing by Office Size
A single 48-port switch does not automatically serve 80 employees — once you count wired workstations, VoIP phones, APs, cameras, and printers, ports fill up quickly. Here are realistic configurations for three common office sizes:
~25 employees (1 × Pro Max 48 PoE — comfortable fit)
| Qty | Item | Unit Price | Extended | 
|---|---|---|---|
| 1 | Dream Machine Pro Max | $646 | $646 | 
| 1 | Switch Pro Max 48 PoE | $1,401 | $1,401 | 
| 3 | U7 Pro (Wi-Fi) | $203 | $609 | 
| 4 | G6 Bullet (cameras) | $228 | $912 | 
| 1 | 10G SFP+ DAC cable (gateway–switch uplink) | ~$13 | $13 | 
|  | Hardware subtotal |  | $3,581 | 
Port usage: ~15 wired workstations + 6 VoIP phones + 3 APs + 4 cameras + 2 printers = 30 ports of 48 available.
~50 employees (1 × Pro Max 48 PoE + 1 × Pro Max 24 PoE)
| Qty | Item | Unit Price | Extended | 
|---|---|---|---|
| 1 | Dream Machine Pro Max | $646 | $646 | 
| 1 | Switch Pro Max 48 PoE | $1,401 | $1,401 | 
| 1 | Switch Pro Max 24 PoE | $862 | $862 | 
| 5 | U7 Pro (Wi-Fi) | $203 | $1,015 | 
| 6 | G6 Turret (cameras) | $228 | $1,368 | 
| 2 | 10G SFP+ DAC cables | ~$13 | $26 | 
|  | Hardware subtotal |  | $5,318 | 
Port usage: ~30 wired workstations + 15 VoIP phones + 5 APs + 6 cameras + 3 printers = 59 ports across 72 available.
80+ employees — requires additional access switches. Count every wired endpoint, phone, AP, and camera to determine how many switch ports you need, then add 20% headroom. PoE budget must cover all powered devices simultaneously; verify with the PoE budget calculator.
What These Prices Exclude
Surveillance hard drives, UPS, rack/enclosure, patch panels, cabling and labor, and tax/shipping are not included above. Budget approximately $20–$50 per cable run for materials and connectors, $100–$300 for a UPS, and $60–$350 per surveillance HDD depending on capacity. For a complete accessories checklist, see our cabling installation guide.
Network Topology and Design
Layout and Device Placement
Place network devices based on your floor plan and coverage requirements:
- Access Points: To ensure strong Wi-Fi coverage, access points should be placed in central locations, such as open workspaces and meeting rooms. Avoid physical obstructions like thick walls, which can weaken signals.
- Switch: Position your switches in a secure area, such as a network closet or server room, where they are easy to access for maintenance yet protected from interference or accidental damage.
