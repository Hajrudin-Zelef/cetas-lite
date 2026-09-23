---
id: etape6-phaseb-smb-networking/03-supplementary-complementary-research-pass-september-22-2026/s5-zyxel-wi-fi-7-aps-2026-current-lineup-and-prices-was-brie
title: "S5. Zyxel Wi-Fi 7 APs — 2026 current lineup and prices (was: brief)"
domain: supplementary-complementary-research-pass-september-22-2026
role: deep-dive
task: reference
actors: ["EU", "Qualcomm"]
dates: []
keywords: ["consumer", "ethernet", "licenses", "memory"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [428, 482]
section: "Supplementary / Complementary Research Pass — September 22, 2026"
sha256: c84af676f63bfe10b998c22c91153fe5b2a26defed16c4588980c4a341a75673
---

# S5. Zyxel Wi-Fi 7 APs — 2026 current lineup and prices (was: brief)

## S5. Zyxel Wi-Fi 7 APs — 2026 current lineup and prices (was: brief)

Zyxel launched six Wi-Fi 7 APs in Sep 2025; all current in 2026, priced $80–$100 [secondary]:
- **NWA90BE PRO** (BE6500, 4-stream, dual-radio, BandFlex tri-band) — $99.99 street
- **NWA50BE PRO** (BE6500, 4-stream) — $89.99 street
- **NWA90BE** (BE5100) — $89.99 street
- **NWA50BE** (BE5100) — $79.99 street
- NWA55BE / NWA30BE (wall-mount / outdoor) — ~$80–100
- Earlier dual-radio MLO line (high-density): **WBE510D** $199.99, **WBE630S** $299.99 [secondary]
- All: NebulaFlex (cloud or local management), QR-code onboarding, SSID VLAN tagging, Smart Mesh MLO, limited lifetime warranty, Qualcomm chipsets [vendor-reported/secondary]
- Sources: https://www.morningstar.com/news/business-wire/20250929963115/zyxel-networks-expands-portfolio-of-wifi-7-access-points-for-small-businesses ; https://www.pocnetwork.net/technology-news/zyxel-launches-new-affordable-wifi-7-nebulaflex-access-points/

## S6. NETGEAR — AV/IT M4350 switching line (new vendor)

NETGEAR's **M4350** series (AV Line + IT use) is the current mid-enterprise stackable platform — 1G to 100G in ≤40 cm depth, no licenses required for the full feature set [vendor-reported via retailers]:
- **M4350-8X8F** (8× 10G RJ45 + 8× 10G SFP+) — street **$838.99** USD [independent]
- **M4350-24G4XF** (24× GbE PoE+ 648W budget + 4× 10G SFP+) — street **$3,129.80** (regular $3,397.64) [independent]
- **M4350-24F4X** (24× 10G SFP+ + 4× 100G QSFP28) — street **$4,947.54** (regular $5,358.60) [independent]
- **M4350-32F8V** (32× 25G SFP28 + 8× 100G QSFP28) — street **$6,756.63** (regular $7,386.19) [independent]
- **M4350-44M4X4V** (44× 25G + 4× 100G) — street **$8,689.64** [independent]
- Features: Virtual Chassis stacking with NSF hitless failover, IGMP Plus (AV-over-IP out of box), auto-LAG/auto-trunk, AV UI + Engage Controller (also manages M4250/M4300/M4500), SMPTE ST 2110 on select models, Dante/Q-SYS/NDI/AES67/AES70 profiles, CLI/GUI/SNMP/sFlow/RSPAN, ProSAFE Limited Lifetime hardware warranty with next-business-day replacement [vendor-reported via reseller]
- Positioning: Pro-AV-over-IP leader plus edge-to-core for midsize organizations; NETGEAR Engage Controller + ProAVDesign services as differentiators [independent assessment]
- **Flag:** no confirmed NEW NETGEAR switch model in 2026 — M4350/M4250/M4500 are the current lines; 2026 MSRP/list updates not found. NETGEAR consumer side (Orbi) excluded from scope.
- Sources: https://www.lttpartners.com/products/netgear-av-line-m4350-24g4xf-ethernet-switch ; https://www.camcor.com/cat/Netgear-AV-Line-M4350-8X8F-Ethernet-Switch-XSM4316-100NES.html

## S7. HPE Aruba Instant On — SMB line (new vendor)

Aruba Instant On is HPE's cloud-managed SMB line (app/cloud portal, no subscription, limited-lifetime warranty) — current in 2026; no 2026 hardware launch found [secondary/unverified]:
- **Instant On 1830** (entry, L2 smart-managed): 8/24/48-port gigabit, PoE variants (65W/195W/370W), 4× SFP uplinks on 48-port; street example JL811A (8G PoE 65W) [secondary/retailer specs]
- **Instant On 1930** (mid, 8/24/48-port, 4× SFP+ 10G): JL683B (24G 4SFP+ 195W), JL686B (48G 4SFP+ 370W), 176 Gbps fabric on 48-port [secondary/retailer specs]
- **Instant On 1960** (high, stackable up to 4, 1G/10G local+cloud-managed stacking): JL809A 48G 2XGT 2SFP+ 600W [secondary/retailer specs]
- Management: Instant On mobile app / cloud web portal / local web GUI; SNMP manager support [vendor-reported via retailer]
- **Wi-Fi 7 note:** Instant On Wi-Fi 7 APs not confirmed as of Sep 22, 2026 — the Wi-Fi 7 portfolio (730 Series, ultra tri-band, launched Apr 2024) is enterprise Aruba Central only, not Instant On [secondary — flag]
- Post-HPE-Juniper integration: Instant On remains the SMB brand distinct from enterprise CX; convergence so far targets Mist↔Central, not Instant On [unverified assessment]
- Sources: https://www.crn.com/news/networking/new-aruba-instant-on-switch-series-powers-smb-networking-digital-transformation ; http://www.lelong.com.my/hpe-aruba-instant-1960-48g-2xgt-2sfp-600w-switch-jl809a-server-222329420-2026-09-Sale-P.htm

## S8. D-Link additions: DQS-5000 25G/100G DC switch + Nuclias controllers (was: Nuclias roadmap thin)

- **DQS-5000 56ZS** (Feb 2026 new-devices listing) [secondary]: managed L3 DC switch — 48× 10G/25G SFP28 + 8× 40G/100G QSFP28, 4 Tbps switching, 2+1 hot-swap PSUs (1+1 redundancy), 5 fans (N+1), ONIE pre-installed, EVPN/VXLAN/MC-LAG, extended routing (BGP, OSPF, ECMP, IS-IS), up to 512K IPv4 / 256K IPv6 routes, REST API / ZTP / CLI / sFlow [secondary]
- **Nuclias controllers (launched Oct 2025, current 2026)** [vendor-reported via PRNewswire]: **DNH-1000** (entry, up to 500 devices — switches + APs, for SMB/K-12/retail); **DNH-3000** (up to 1,500 devices, multi-port, for campuses/hospitality); **DNC-5000** (software flagship for multi-site control). Centralized management, real-time visibility, tiered admin, automated optimization [vendor-reported]
- **Wi-Fi 7 APs (Japan 2026–27 guide)** [secondary]: DAP-E9560 (802.11be 2×2×2 tri-band, 10G PoE port, 802.3bt), DAP-E3620 (802.11be 2×2, multi-gig PoE), DAP-E3620OU (outdoor) — all Nuclias Connect compatible
- 10G L2 lines in the same guide: DXS-1250 series (10G, PoE bt models), DXS-1100 series (easy smart) [secondary]
- Sources: https://tiny-bar-e01b.ceriseaurora.workers.dev (D-Link "new devices Feb 2026" list) ; https://beta.manilatimes.net/2025/10/16/tmt-newswire/pr-newswire/d-link-unveils-nuclias-network-controllers/2202000

## S9. UniFi Enterprise ECS line — official confirmations + memory surcharge (was: SKU/price unconfirmed)

- **Enterprise Campus Switch Core**: official SKU **ECS-Core**, MSRP **$4,999** + $394 memory surcharge = **$5,393** [official, store.ui.com]. This resolves the §1.1 uncertainty: "32× QSFP28 model" = ECS-Core. SFP28 variant noted at $4,314 [secondary, Notebookcheck]
- **Enterprise Campus Aggregation (ECS-Aggregation)**: 48× 25G SFP28 + 6× 100G QSFP28, 3.6 Tbps / 1.8 Tbps, MC-LAG, L3 Etherlighting — MSRP **$3,999** + $315 surcharge [official/secondary]
- **ECS-48-PoE** (16× 2.5G PoE+++ + 32× 10G PoE+++ + 4× 25G SFP28, 2150W, stacking): **$3,499** + surcharge; "out of stock" at one retailer Sep 2026 [secondary]
- **ECS-24-PoE** (1050W): $3,585 CAD ($3,868 surcharge incl.) [official, ca.store.ui.com]
- **EAV-XG-24-PoE** (Enterprise **Audio/Video**: 24× 10G PoE+++ + 4× 100G QSFP28, PTP timing, SMPTE ST 2110, SDVoE, AES67, OCXO clock, GPS grandmaster input): $5,735 CAD / €3,599 (EU €4,659.60 incl. VAT & surcharge) [official store / secondary]
- **Ubiquiti memory surcharge (effective ~Jul 1, 2026)** [secondary]: applied to switches (ECS-Core $394, ECS-Aggregation $315, USW-Pro-XG-24/48 etc. $17–$46) and other lines, explicitly "due to global memory and storage market conditions" plus tariffs. 5-year coverage add-on for ECS-Core: **$1,099/unit** [official/secondary]
- Accessories [official]: QSFP28 DAC from **$35**; 100G MM MPO-12 transceiver **$69**
- Sources: https://store.ui.com/us/en/category/switching-enterprise/collections/enterprise-campus-aggregation/products/ecs-core ; https://blog.streakwave.com/hubfs/Ubiquiti/2026/Notice-Ubiquiti-Memory-Surcharge-July-2026-with-Tariffs.pdf ; https://ca.store.ui.com/ca/en/collections/unifi-switching-enterprise-power-over-ethernet

