---
id: etape6-phaseb-smb-networking/10-supplementary-complementary-research-pass-round-3-september-/overview
title: "Supplementary / Complementary Research Pass — Round 3 — September 22, 2026"
domain: supplementary-complementary-research-pass-round-3-september-
role: deep-dive
task: reference
actors: ["Samsung", "United States"]
dates: ["2026-03", "2026-04-28", "2026-05-28", "2026-07-27", "2026-09-01", "2026-09-09", "2026-09-22"]
keywords: ["research", "advisory", "asic", "disclosure", "distribution", "memory", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1818, 1872]
section: "Supplementary / Complementary Research Pass — Round 3 — September 22, 2026"
sha256: fa3211d1b7dbdca4d7b6cd1055614edeba9edeb40230590818bcc04bf8f8729f
---

# Supplementary / Complementary Research Pass — Round 3 — September 22, 2026

**Scope:** third complementary pass. Only genuinely new material vs base + S1–S18 + R1–R10. Focus: 2026 launches/pricing not previously captured (NETGEAR ISE 2026 models, Aruba Instant On 1960 12-port multi-gig, HPE CX 6000), UniFi release-train correction, EAP772 street prices (listings created Sep 22, 2026), DXS-3130 street prices, GWN7816 pricing, UISP Fiber XGS 2026 prices, XS1935-12F retail sightings, Omada Pro AP MSRPs, and the 2026 memory-price crisis context ("memflation") behind observed surcharges. No existing section altered.

Provenance tags: [official] / [vendor-reported] / [independent] / [secondary] / [unverified].

---

## T1. NETGEAR — NEW M4350 models at ISE 2026 (Jan 2026) [vendor-reported via trade press]

- Two new M4350 models announced **Jan 27, 2026** ahead of **ISE 2026 (Feb 3–6, Barcelona)**; availability stated as **March 2026** [vendor-reported]:
  - **M4350-16M4V** — ruggedized AV/broadcast model: industry-standard **Neutrik locking connectors** (etherCON for network, opticalCON for fiber, powerCON for power); 16× 2.5G PoE++ ports (up to **1,130W total** PoE budget); 4× 25G SFP28 uplinks; **4 modular interface card slots** for uplink customization [vendor-reported].
  - **M4350-16C** — aggregation/core model: **16× 100G ports**, aimed at large AV-over-IP deployments where high-resolution video streams converge [vendor-reported].
- Portfolio now **18 M4350 models** (1G to 100G), all with high-power PoE++, redundant modular PSUs, SMPTE ST 2110 timing support; TAA-compliant versions for government projects [vendor-reported].
- **NETGEAR Engage Controller v2.4** (Feb 2026): **offline provisioning** — create virtual switches and Wi-Fi 7 APs, design/test networks without physical hardware, export configs, save reusable templates [vendor-reported].
- Firmware sync note (NETGEAR community, Aug 2026-era): **M4250 latest 13.0.4.26**, **M4350 latest 14.0.2.26** — recommended to keep same FW/AV-UI generation across AVoIP installs for profile/PTP-TC parity; Shure profiles qualified on M4250 only at time of posting (awaiting validation on M4300/M4350) [vendor-reported/community].
- Sources: https://www.businesswire.com/news/home/20260127516549/en/ ; https://www.expresscomputer.in/news/netgear-unveils-rugged-av-ready-m4350-switches-with-offline-network-configuration-at-ise-2026/131738/ ; https://community.netgear.com/discussions/en-business-pro-av-over-ip-switches/av-profile-discrepancy-between-m4250-and-m4350/2401926

## T2. Aruba Instant On 1960 — NEW 12-port multi-gigabit model (S0F35A) [independent]

- **NEW SKU S0F35A** listed at PB Tech (Pacific) with **"Date Created: 12:02, 22-09-2026"** — i.e., a brand-new 2026 Instant On 1960 model [independent]:
  - 12-port multi-gigabit smart-managed L2+ stackable: 8× 802.3af/at PoE ports + **4× 2.5G 802.3bt PoE ports (max 480W total)** + **2× 10G RJ45 + 2× SFP+** uplinks; limited lifetime warranty, 90-day 24×7 phone support [independent via retailer].
  - PB Tech price: **NZD $2,876.66** (ex GST) [independent].
- Significance: prior passes found "no new 2026 Instant On hardware" — this listing overturns that for the 1960 line: a compact multi-gig 12-port member joining the existing 24G/48G models (JL807A etc.) [independent assessment].
- Source: https://www.pbtech.com/pacific/product/SWHAUB19605/Aruba-Instant-On-1960-S0F35A-12-Port-Multi-gigabit

## T3. HPE Aruba CX 6000 — NEW L2 branch/retail switch series, NRF 2026 [secondary]

- Announced at **NRF 2026 (Jan 11–13, 2026)** alongside Marvis/Juniper-analytics integration: new **HPE Aruba Networking CX 6000** 8-port L2 switch family targeting retailers/branch/edge [secondary]:
  - 5 fixed 1U models: 24 and 48× 1GbE access ports + **4× 1G SFP uplinks**; up to **104 Gbps non-blocking**, **77.3 Mpps**; HPE Aruba Networking ASIC architecture; programmable CX Operating System [secondary].
  - PoE models: 24-port up to **370W**, 48-port up to **740W**, 802.3at Class 4 (30W/port) [secondary].
- Positioned between Instant On (SMB) and CX 6100/6300 (enterprise campus) for POS, digital signage, IoT endpoints in small stores [secondary].
- Source: https://www.networkworld.com/article/4115610/nrf-2026-hpe-expands-network-server-products-for-retailers.html

## T4. UniFi Network Application release train — correction/addition [official via community trackers]

- Prior pass listed 10.6.106 as "latest tracked (Sep 10, 2026)". Community Docker trackers give a more precise picture [official via ui.com release pages, mirrored on GitHub]:
  - **10.6.101** — **official stable release 2026-09-01**.
  - **10.6.106** — **beta/release candidate track 2026-09-09** (not yet promoted to official stable as of Sep 10 listing).
  - **10.5.67** — official release **2026-07-27** (not just "Aug 2026" per security bulletins).
  - 10.4.57 official 2026-05-28; 10.3.58 official 2026-04-28.
- **SAB-067 (Aug 2026)** — the "21 critical flaws" disclosure maps to Ubiquiti's official **SAB-067** advisory: **three simultaneous CVSS-10.0** flaws across Protect (CVE-2026-77537/77533/77548 → fix 7.2.105), OS Server (CVE-2026-77550/77539/77540 → fix 5.1.37), and Talk (CVE-2026-77554 → fix 5.3.2) [secondary/official-advisory]. Key remediation note from trade press: no single isolation action closes all three attack surfaces — full patching of every listed component required before any can be considered secure [secondary].
- Sources: https://github.com/goofball222/unifi/blob/HEAD/README.md ; https://www.techtimes.com/articles/325720/20260827/ubiquiti-patches-three-simultaneous-maximum-severity-unifi-flaws-cameras-os-voip.htm

## T5. TP-Link Omada Pro — Wi-Fi 6 AP MSRPs + controller street prices [official/independent]

- The TP-Link Canada **Feb 2026 Distribution Price List** (MSRP, CAD) includes the **Omada Pro** AP lineup (new detail vs S1 which had switches only) [official]:
  - **AP9635** — Omada Pro AX1800 ceiling-mount HD AP (1000+ clients, 802.3at, mesh, WIPS/WIDS): **$949.99 CAD** [official].
  - **AP8635-E** — Omada Pro AX1800 indoor/outdoor, IP67, external antennas: **$1,999.99 CAD** [official].
  - **AP8635-I** — Omada Pro AX1800 indoor/outdoor, IP67, internal antennas: **$1,699.99 CAD** [official].
  - **AP7650** — Omada Pro AX3000 wall-plate (2.5G uplink + 3× GbE downlink, 802.3at/af PoE): **$849.99 CAD** [official].
- Compsource US street (Jul 3, 2026 update): S5500-48GP4XF **$3,849.99**; S5500-8MHP2XF **$1,869.99** [independent].
- Hardware controller streets (Sep 2026): **OC300 $160** (Best Buy US, in stock, ships by Sep 25); OC200 KWD 22.183 (Microless Kuwait, 2 left) / QAR 260.95 (Qatar); OC300 IQD 246,945 (elryan.com) [independent].
- Source: https://micro-informa.ca/files/Omada_MSRP_2026.pdf ; https://www.compsource.com/buy/S550048GP4XF/Tp-Link-3623/TPLink-Omada-Pro-48Port-PoE-Gigabit-L2-Managed-Switch-with-4-SFP-Slots--48-Ports--Manageable--S550048GP4XF/buy/SMT577UZKGN14/Samsung-389

