---
id: etape6-phaseb-smb-networking/03-supplementary-complementary-research-pass-september-22-2026/overview
title: "Supplementary / Complementary Research Pass — September 22, 2026"
domain: supplementary-complementary-research-pass-september-22-2026
role: deep-dive
task: reference
actors: []
dates: ["2026-09-22"]
keywords: ["research", "distribution", "license", "memory", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [372, 427]
section: "Supplementary / Complementary Research Pass — September 22, 2026"
sha256: 2537c766bf1d9fe774a11d4607e5ae1f23c6f8d492a44cd90887504e3186e99e
---

# Supplementary / Complementary Research Pass — September 22, 2026

**Scope:** fills gaps identified after the two 2026-09-22 research passes above. New vendors added: **NETGEAR** (AV/IT M4350), **HPE Aruba Instant On** (SMB line). New detail: TP-Link Omada Pro MSRP lineup, TP-Link "Agile" easy-managed series, Zyxel USG FLEX H 2026 street pricing, Zyxel Wi-Fi 7 AP lineup, D-Link DQS-5000 25G/100G DC switch + Nuclias controllers, UniFi Enterprise ECS line confirmations + memory surcharge, UISP Fiber XGS, UniFi G6 ecosystem. No existing section was altered.

Provenance tags: [official] = vendor price list/store/datasheet; [vendor-reported] = vendor claim via trade press; [independent] = third-party test/retailer listing; [secondary] = press/blog/forum; [unverified] = single weak source or inferred.

---

## S1. TP-Link Omada Pro — full lineup and official MSRP (was: price unconfirmed)

The TP-Link Canada **2026 February Distribution Price List** (MSRP, CAD) confirms the Omada Pro S5500 family [official]:
- **S5500-8XF** — 8× 10GE SFP+ L2+ managed; MSRP ~$999.99 CAD (compsource regular) / street **$1,010.23** USD [independent, compsource, listing updated Jun 26, 2026]
- **S5500-8MHP2XF** — 8× 2.5G PoE+ + 2× 10G SFP+, 240W PoE; MSRP **$2,299.99** CAD [official]; street **$1,869.99** USD (compsource, updated Jul 3, 2026) [independent]
- **S5500-24MPP4XF** — 24× 2.5G (16× PoE+ + 8× PoE++), 4× 10G SFP+, 500W PoE; MSRP **$3,999.99** CAD [official]
- **S5500-48GP4XF** — 48× Gigabit PoE+ + 4× 10G SFP+, 500W PoE; MSRP **$4,499.99** CAD [official]; compsource street **$3,469.77** (regular $3,499.99) USD [independent]
- **S5500-24GP4F** — 24× Gigabit PoE+ + 4× 1G SFP (non-XF gigabit variant), 250W PoE; retailer-listed ("Request a Quote") [secondary]
- S5500-48GP4F (48-port gigabit PoE+, 384W PoE budget, 104 Gbps / 77.38 Mpps, 4,096 VLANs) — BE retail listing [secondary]
- All S5500 models: integration with Omada **PRO** SDN controller, static routing, OAM, DDM, sFlow, QinQ, MSTP, 802.1X/RADIUS/TACACS+, LACP, CLI, SNMP, dual image/config, IPv6 [vendor-reported via reseller specs]
- Sources: https://micro-informa.ca/files/Omada_MSRP_2026.pdf ; https://www.compsource.com/buy/S550048GP4XF/Tp-Link-3623/TPLink-Omada-Pro-48Port-PoE-Gigabit-L2-Managed-Switch-with-4-SFP-Slots--48-Ports--Manageable--S550048GP4XF/

**Open gap retained:** Omada Pro controller (hardware/software) pricing still not found publicly; TP-Link positions Omada Pro as the enterprise step-up above standard Omada (features like Anomaly Detection / OUI VLANs have trickled down to standard Omada v6.3, see main §1.1) [unverified].

## S2. TP-Link "Agile" easy-managed series — fully detailed (was: thin sourcing)

"Agile" is TP-Link's 2026 sub-brand for **easy-managed (smart-managed-lite) switches** aimed at surveillance/SMB — positioned between unmanaged and full Omada SDN switches, managed via free **Omada Cloud Essentials** [official]:
- **ES205GP** — 4× GbE PoE+ + 1× GbE RJ45 uplink; 65W PoE budget; desktop/wall [official]
- **ES206GP** — 4× GbE PoE+ + 2× GbE (incl. SFP combo); 65W; desktop/wall [official]
- **ES210GMP** — 8× GbE PoE+ + 1× GbE RJ45 + 1× SFP combo; 123W; desktop/wall, fanless [official]
- **ES220GMP** — 16× GbE PoE+ + 2× RJ45 + 2× SFP; 250W; rackmount [official]
- **ES228GMP** — 24× GbE PoE+ + 2× GbE SFP; 384W; rackmount [official]
- **ES228GP** — 24× GbE PoE+ + 2× GbE SFP; 250W; rackmount [official]
- **ES206XPP-M2** — 5× 2.5G RJ45 (4× PoE++) + 1× 10G SFP+; 120W; fanless [official]
- **ES210XPP-M2** — 8× 2.5G RJ45 (8× PoE++) + 1× 10G RJ45 + 1× 10G SFP+; 200W; fanless [official]
- Features: plug-and-play, ZTP, PoE Auto Recovery, 250 m PoE under Extend Mode (gigabit models), port isolation, remote camera reboot, auto loop detection, cable test, DSCP/port/VLAN QoS [official]
- Portfolio context (ISC West 2026 flyer): Access line split = **Access Max** (L2 with 10G downlink), **Access Pro** (L2 with 2.5G downlink), **Access Plus** (L2 with 1G downlink), **Agile** (easy-managed 1G, surveillance/prosumer) — plus Campus L3 and Aggregation tiers above [official]
- Sources: https://www.tp-link.com/us/solution/poe-switches-for-surveillance/ ; https://static.tp-link.com/document/pdf/en/isc-west-2026/2026_ISC_WEST_Omada_Business_Switches.pdf

## S3. TP-Link Wi-Fi 7 AP — EAP770 BE11000 MSRP (Omada 2026 price list)

- **EAP770** — Omada BE11000 tri-band Wi-Fi 7 ceiling-mount AP: 574 Mbps (2.4 GHz) + 4,320 Mbps (5 GHz) + 5,760 Mbps (6 GHz), 1× 2.5G RJ45, 802.3bt PoE++, MLO, 4K-QAM, 320 MHz, seamless roaming, managed by Omada SDN Controller; MSRP **$309.99** CAD [official, Feb 2026 price list]
- Source: https://micro-informa.ca/files/Omada_MSRP_2026.pdf

## S4. Zyxel USG FLEX H — 2026 street pricing (was: absent)

USG FLEX H series (launched Oct 2023, current 2026 SMB firewall line) — street prices per Zyxel's Apr 15, 2025 press update [secondary]:
- USG FLEX 100H (8× GbE): **$299.99** street; bundled (1-yr Gold Security License) **$399.99**
- USG FLEX 100HP (8× GbE, 1× PoE+ at): **$399.99**; bundled **$499.99**
- USG FLEX 200H (2× 2.5G + 6× GbE): **$399.99**; bundled **$549.99**
- USG FLEX 200HP (2× 2.5G + 6× GbE, 1× PoE+): **$499.99**; bundled **$649.99**
- USG FLEX 500H (2× 2.5G + 2× 2.5G PoE + 8× GbE): **$799.99**; bundled **$1,099.99**
- USG FLEX 700H (2× 2.5G + 2× 10G PoE + 8× GbE + 2× 10G SFP+): **$1,299.99**; bundled **$1,699.99**
- Performance (Zyxel comparison sheet, Rev 2026-05) [official via partner doc]: firewall SPI up to 7 Gbps (700H); IPS up to 4 Gbps; UTM with Gold license up to 3 Gbps; VPN up to 1.2 Gbps; up to 100 IPsec tunnels (200H)
- Management: standalone or Nebula Cloud (status, firmware, license management, Nebula Secure remote GUI, config backup) [vendor-reported]
- Sources: https://businesswire.com/news/home/20250415755419/en/Zyxel-Networks-Upgrades-Leading-Firewall-Family-to-Provide-SMBs-with-Unified-Cloud-and-On-Premises-Security ; https://info.zyxel.com/hubfs/USG%20FLEX%20H%20Comparison_Partners_Rev2026.pdf
- **Flag:** no NEW USG FLEX model announced in 2026 — this is the current lineup, not a 2026 launch.

