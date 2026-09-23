---
id: etape6-phasea-vendors-dc/00-front-matter/supplementary-complementary-research-pass-round-3b-2026-09-2
title: "Supplementary / Complementary Research Pass — round 3b (2026-09-22): Meraki pricing & licensing tiers, Cisco C9350/C9610, IOS XE 17.18, BlueField-4, Quantum-X800 detail, Cumulus 5.18"
domain: front-matter
role: reference
task: pricing
actors: ["China"]
dates: ["2025-02", "2025-06", "2026-09-22"]
keywords: ["pricing", "research", "asic", "latency", "license", "throughput"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [967, 1018]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 1afb62cf385659ea408dee26fb11e3be9041355119fe598704729f62e1ad703b
---

# Supplementary / Complementary Research Pass — round 3b (2026-09-22): Meraki pricing & licensing tiers, Cisco C9350/C9610, IOS XE 17.18, BlueField-4, Quantum-X800 detail, Cumulus 5.18

## Supplementary / Complementary Research Pass — round 3b (2026-09-22): Meraki pricing & licensing tiers, Cisco C9350/C9610, IOS XE 17.18, BlueField-4, Quantum-X800 detail, Cumulus 5.18

**Scope note:** A third additive pass, independently researched the same day. It contains only material not present in the base §§1–8, the first supplementary pass (§§A–M), Pass #2 (§§N–X), or the round-2 pass (second §§N–S). Earlier sections were left untouched. Read the base report and prior passes first; the provenance legend from the base applies here.

### R3b-A. Meraki MS450 / MS355 / MS390 — 2026 datasheet refreshes and street pricing detail

The base report (§4.1–4.4) documented the MS portfolio and 2026 datasheet refreshes; Pass #1 (§H) documented the licensing model. New 2026 pricing evidence:

- **MS450-12-HW** (12× 40G QSFP+ + 2× 100G QSFP28 uplinks, 100G aggregation): **MSRP $20,439.21, street $17,938.91** (hardware-only listing, power cord sold separately) [secondary — 4tekgear.com, 2026]. A second listing showed $20,571.00 (datadirectglobal, quote-based) [secondary].
- **LIC-MS450-12-5YR** (Enterprise license & support, 5 years): **street $3,630.09 (list $5,490.15)**; **LIC-MS450-12-3YR**: **street $2,178.31 (list $3,294.49)** [secondary — hummingbirdnetworks.com, 2026].
- **MS355-48X2-HW**: **$12,866.95** (barcodesinc) [secondary]; **MS355-24X-HW**: ~$1,559–1,641 refurbished [secondary — gotodirect.com]. Note: refurb/secondary-market prices are not Cisco pricing signals.
- **MS355-48X 5-year license (LIC-MS355-48X-5YR)**: **street $2,089.27 (list $4,519.76; MSRP $3,025.00)**; **MS355-48X2 1-year license (LIC-MS355-48X2-1YR)**: MSRP **$1,110.00** [secondary — 4tekgear.com, 2026].
- **LIC-MS-100-L-E** (MS Series Large Essentials + Support, annual): street **$362.99** (backorder, 2026) [secondary — pc-canada.com].
- Datasheet refreshes confirmed current (Jul 21 / Sep 4, 2026); no new MS hardware family in 2026 — status quo stands [official — Meraki documentation].

### R3b-B. Meraki licensing tiers: Essentials vs Advantage (unified Cisco subscription nomenclature)

- Meraki MS switches now use tiered licensing: **Essentials** (core L2/L3 functionality) vs **Advantage** (adds Adaptive Policy, advanced analytics, SD-Branch capabilities), available in 1-, 3-, 5-, 7-, and 10-year terms; tiers vary by model size (MS100/200/300/400) [secondary — stratusinfosystems.com MS series comparison guide, 2026].
- Cisco's unified **Networking Subscription datasheet** (2026, cisco.com) defines the cross-portfolio SKU model that now governs cloud-managed switching [official]:
  - **LIC-CS-AC1-L-E**: Cisco Switching Essentials License, Access Tier 1, Large (48-port)
  - **LIC-CS-AC1-M-E**: Cisco Switching Essentials License, Access Tier 1, Medium (up to 24-port)
  - **LIC-CS-AC1-L-A**: Cisco Switching Advantage License, Access Tier 1, Large (48-port)
  - **LIC-CS-AC1-M-A**: Cisco Switching Advantage License, Access Tier 1, Medium (up to 24-port)
  - **LIC-CS-CO-MO-A**: Cisco Switching Advantage License, Core Modular — **only available with C9610**
- Implication: the Meraki Essentials/Advantage tiering is part of Cisco's unified subscription structure shared with cloud-managed Catalyst — one licensing language across Meraki MS and Catalyst cloud operating mode [official].
- TCO caveat restated: this pass confirms Pass #1's finding — Meraki devices stop working when the license expires; license + hardware are both mandatory [secondary].

### R3b-C. Cisco smart switches C9350 / C9610 — the adjacent Cisco enterprise play (Meraki-convergence context)

Scope note: the C9350/C9610 are Cisco (not Meraki) hardware, but they are directly adjacent to Phase A scope because (a) they are Meraki-dashboard-managed via Cloud Management with IOS XE, and (b) the C9610 is the first modular smart switch in the dashboard — the device that extends cloud management into the modular/core tier covered for Meraki in base §4.2–4.3.

- **Cisco Live 2025** (June 2025): Cisco unveiled **C9350 and C9610 Smart Switches**, powered by **Silicon One**, delivering **up to 51.2 Tbps throughput, sub-5-microsecond latency and quantum-resistant secure networking** for high-stakes AI applications; embedded DPUs carry **Hypershield** (AI-powered security architecture) as a service — a continuation of the February 2025 Nexus 9300 smart-switch launch [vendor-reported — CRN, TechPowerUp].
- **C9350 series** (access smart switches, 2026 datasheet) [official — Cisco C9350 datasheet, 2026]:
  - **C9350-48U**: Silicon One **A100/L** ASIC; 48× 10M/100M/1G downlinks with **60W UPoE per port (2,880W total budget)**; up to 200G modular uplinks (C9350-NM-2C, C9350-NM-8Y).
  - **C9350-48P**: 48× 1G, 30W PoE per port, 1,440W budget.
  - **C9350-48T**: 48× 1G, non-PoE.
  - **C9350-48HXN**: 1.3 Tbps throughput; **90W UPoE+ on all ports (4,320W total budget)**; 36× mGig (10M–5G) + 12× (10M–10G) downlinks; StackWise-1.6T (up to 1.6 Tbps stack bandwidth, 8-switch stack up to 448 multigigabit / 384 UPoE+ ports); up to 200G modular uplinks.
  - **C9350-48TX**: multigigabit access without PoE.
  - Scale/security claims: **4× MAC scale, 4× ACL scale, 6× route scale**, crypto-ready/PQC-resistant (Cisco marketing claims — vendor-reported).
- **C9610 Smart Switch**: 10-slot modular chassis (2 redundant supervisors, 8 line-card slots, 8 PSU slots, 4 fan trays), up to **480 ports on a single page** in the Meraki dashboard; unified licensing via **LIC-CS-CO-MO-A** (Advantage Core Modular) [vendor-reported/official — Meraki Community, Cisco subscription datasheet].
- Cloud-management eligibility (2026): 9300-M, 9300X-M, 9300L-M and MS390 families are all upgradable to **Cloud Management with IOS XE**; the Dec 2025 Cisco Tech Club switching update introduced the C9350 series as new access smart-switch hardware on the same IOS XE cloud track [official — Meraki documentation; Cisco Tech Club webinar].

### R3b-D. IOS XE 17.18 — cloud-management milestones (2026)

Extends base §4.2 (Cloud Monitoring EoS Mar 31, 2026; device-config mode) with the 2026 release train [official — Meraki documentation, 2026]:

- **IOS XE 17.18.1**: brings Cloud Management to additional platforms — **C9200, C9300, C9500** high-performance families; advanced routing protocols, high availability, expanded cloud visibility and troubleshooting tools.
- **IOS XE 17.18.2 (release candidate)**: adds **BGP, VRF, and Cloud EVPN Fabric**, plus **SmartPorts** and **Intelligent Capture (scheduling)** — EVPN fabric creation directly from the Meraki dashboard.
- Unsupported features on 17.15.x cloud configuration (SmartPorts, MAC blocklist, DOM, RSPAN/VLAN SPAN, IPv6 RA/DHCP guard, DAI, etc.) are being addressed in later 17.18 releases; 17.15.x remains the floor for device-config mode (17.15.3+ minimum) [official].
- **Cloud-managed Catalyst is now supported in the Meraki India, China, and Canada regional clouds**, in both cloud-configuration and device-configuration modes [official].
- Downgrade restriction: after upgrading to Cloud Management with IOS XE, downgrading to CS firmware is restricted — factory reset and support assistance may be required [official].

