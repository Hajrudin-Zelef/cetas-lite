---
id: etape6-phasea-vendors-dc/00-front-matter/overview
title: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "Malaysia", "Nvidia", "xAI"]
dates: ["2025-11", "2025-11-17", "2026-05", "2026-06", "2026-06-11", "2026-09-22"]
keywords: ["accelerator", "asic", "distribution", "ethernet", "gpu", "gpus", "neocloud", "nvidia", "pricing", "rack-scale", "research", "revenue"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1, 90]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 9b8ed328f4e30c4aabfd0d76d26b825081a48c12a023a50be362858b24978152
---

# Step 6 — Phase A: Enterprise Data-Center Switching Vendors
## Dell Networking + NVIDIA Networking (Mellanox heritage) + Aruba (HPE) + Cisco Meraki

**Research date:** 2026-09-22
**Coverage window:** February 1 → September 22, 2026 (with necessary 2025 context for platform generations)
**Method:** vendor press releases, spec sheets, documentation; IDC/Dell'Oro market data via trade press; ServeTheHome, Data Center Dynamics, SDxCentral, SiliconANGLE, The Register, Next Platform, TechTarget.
**Provenance legend:** `[official]` vendor documentation/press · `[vendor-reported]` vendor statements relayed by press · `[independent]` IDC/Dell'Oro/market analysts and independent measurements · `[secondary]` trade-press reporting · `[unverified]` single-sourced or unconfirmed claims.

---

## Master timeline (2026)

| Date | Event | Provenance |
|---|---|---|
| Jan 2026 | HPE adds CX 6000 8-port models at NRF 2026 (retail edge) | [official] |
| Feb 2026 | NVIDIA ConnectX-9 SuperNIC GA (firmware v82.48.1000) — 800 Gb/s/port, up to 1.6 Tb/s to Rubin GPUs | [official] |
| Mar 2026 | Cisco ends Cloud Monitoring EoS path for Catalyst 9500; device-config mode in Meraki dashboard GA track (min IOS XE 17.15.3+) | [official] |
| May 2026 | Dell launches PowerRack for networking: 8× PowerSwitch SN 6600 LD (NVIDIA Spectrum-6), >800 Tbps/rack, PowerCool C7000 CDU | [vendor-reported] |
| Jun 2026 | IDC Q1 2026 tracker: **NVIDIA #1 in data-center Ethernet switching by revenue** — $2.1B, 21.5% share, +192.7% YoY | [independent] |
| Jun 11, 2026 | Dell Enterprise SONiC 4.6.0 GA for NVIDIA Spectrum platforms (Debian 12 Bookworm, kernel 6.1, FRR 8.2.2) | [official] |
| Jun 2026 | HPE Discover 2026: Marvis → Aruba Central; Aruba CX switches → HPE Mist; QFX/PTX 12000 AI-fabric hardware | [vendor-reported] |
| Jul 15, 2026 | NVIDIA MLNX-OS 3.12.6300 for QM9700/QM9701 (Quantum-2 NDR) | [official] |
| Jul 21, 2026 | Meraki MS450 datasheet refresh | [official] |
| Aug 31, 2026 | AI-powered support cases launch in Meraki Dashboard (Client VPN + switching) | [official] |
| Sep 4, 2026 | Meraki MS390 datasheet refresh | [official] |
| Sep 22, 2026 | NVIDIA networking-docs "2026 Rel2A" release train: NVOS 25.03.1010 (XDR), SHARP 3.16.3, UFM Enterprise 6.26.1 | [official] |

---

## §1 Dell Networking

### 1.1 AI-fabric flagship: PowerSwitch Z9964 series (2026 headline launch)

- In November 2025, as part of the Dell AI Factory refresh, Dell announced the **PowerSwitch Z9964F-ON and Z9964FL-ON**, powered by **Broadcom Tomahawk-6**, delivering **102.4 Tbps** switching capacity with **64 ports of 1.6 TbE** (and up to 512× 200 GbE via breakout), targeting AI data centers with **more than 100,000 accelerator chips**; air-cooled and direct-liquid-cooled (DLC) variants [official — Dell blog & BusinessWire release 2025-11-17].
- Dell positions the Z9964 as enabling streamlined two-tier fabrics for AI scale versus traditional multi-layer networks [vendor-reported].
- List/street pricing for Z9964 series had **not been published** as of 2026-09-22 [gap].

### 1.2 Current Z-series portfolio (verified via Dell product pages, 2026)

| Model | ASIC | Capacity | Ports | Positioning |
|---|---|---|---|---|
| Z9964F-ON / Z9964FL-ON | Broadcom Tomahawk-6 | 102.4 Tbps | 64× 1.6T (OSFP-XD class), up to 512× 200G breakout | AI fabric flagship, 100k+ accelerators |
| Z9864F-ON | Broadcom Tomahawk-5 | 51.2 Tbps, 20.3 Bpps, 165 MB buffer, <700 ns | 64× 800 GbE OSFP112, 2× SFP+ | 800G AI fabric aggregation; "up to 8K GPU nodes in a single two-tier 400G fabric" |
| Z9664F-ON | Broadcom Tomahawk-4 | 25.6 Tbps, 10.2 Bpps, 114 MB, <850 ns | 64× 400 GbE QSFP56-DD, 2× 10G SFP+ | 400G ToR/MoR, leaf/spine |
| Z9432F-ON | Broadcom Trident4-X11 | 12.8 Tbps, 5.2 Bpps, 132 MB | 32× 400 GbE QSFP56-DD, 2× 10G SFP+ | high-density aggregation |

Source: Dell official product/spec-sheet pages [official]. Z9864F-ON validated in production: neocloud **Hot Aisle** deployed Z9864F-ON + Dell SONiC for GPU cloud (raised MI300X pricing Jul 2026 citing full capacity) [vendor-reported].

### 1.3 S-series and N-series (data-center ToR + campus, 2026 lineup)

Verified from Dell's data-center switches page (2026) [official]:

- **S5448F-ON**: Broadcom Trident4-X9, 16 Tbps, 48× 100G SFP56-DD + 8× 400G QSFP56-DD (ToR for 400G uplinks).
- **S5232F-ON**: Trident3-X7, 6.4 Tbps, 32× 100G QSFP28 — workhorse 100G ToR.
- **S5296F-ON**: Trident3-X7, 6.4 Tbps, 96× 25G SFP28 + 8× 100G QSFP28.
- **S5248F-ON**: Trident3-X5, 4 Tbps, 48× 25G + 4× 100G + 2× 100G QSFP28-DD.
- **S5224F-ON**: Trident3-X5, 2.16 Tbps, 24× 25G + 4× 100G.
- **S4348F-ON / S4148F-ON / S4128F-ON**: 10G ToR family (Trident3-X5/Maverick), 48× 10G SFP+ with 100G uplinks.
- N-series (N1500/N2000/N2200-ON/N3000/N3200-ON/N4000): campus/aggregation with OS10; E3200-ON; MX modular (MX5108n/MX9116n for PowerEdge MX chassis) [official].

### 1.4 Software: Enterprise SONiC + SmartFabric OS10

- **Enterprise SONiC Distribution by Dell Technologies** is Dell's strategic NOS; 2026 current GA on NVIDIA Spectrum platforms (SN2201, SN4700, SN5600, SN5601) is **SN-4.6.0, released 2026-06-11** [official — Dell support KB 000228560].
- **Enterprise SONiC 4.6.0** (community/vendor distributions): Debian 12 Bookworm, Linux kernel 6.1, Broadcom SAI 15.3.0, SDK 6.5.35, FRR 8.2.2; headlined by route-table overflow detection with automatic recovery and route-consistency checker improvements for large BGP/EVPN fabrics [secondary — STORDIS engineering breakdown, 2026-07].
- Ansible support: `dellemc.enterprise_sonic` collection v4.0.0/v4.1.0 (Jan 2026) adding Enterprise SONiC 4.5.1 feature support [official — GitHub changelog].
- **Dell Enterprise SONiC is now certified/compliant with NVIDIA Spectrum-X Ethernet** — announced at SC25 (Nov 2025): enterprises can run the same Spectrum-X architecture as hyperscalers with Dell support; Dell sells NVIDIA Spectrum-X switches as part of the PowerSwitch family with Enterprise SONiC [vendor-reported — theCUBE/SiliconANGLE interview].
- Dell offers OS10 → Enterprise SONiC migration guidance (tech brief H19795.1, **June 2026**) comparing leaf/spine configurations [official — Dell InfoHub].
- SmartFabric Manager: validated blueprints, AI Factory integration, rack-scale OpenManage Enterprise integration; in the May 2026 PowerRack refresh the whole rack (servers, CDUs, PowerShells, switches) is controllable from a single UI with automatic fault isolation [vendor-reported].

### 1.5 Spectrum-6 partnership hardware: PowerSwitch SN 6600 LD

- **May 2026**: Dell PowerRack for networking — >800 Tbps switching capacity powered by **eight new Dell PowerSwitch SN 6600 LD Ethernet switches incorporating the NVIDIA Spectrum-6 ASIC** (introduced with Vera Rubin systems). Paired with PowerCool C7000 in-rack CDU (>220 kW cooling, leak detection, IRC/OpenManage integration) [vendor-reported — TechTarget, 2026-05].
- This confirms Dell as an OEM channel for Spectrum-6 silicon alongside its Broadcom-based Z-series [vendor-reported].

### 1.6 AI cluster wins and market position

- **xAI Colossus**: Dell supplied ~50,000 of the first 100,000 GPUs (Supermicro the other half); phase 2's 100,000 GPUs went to Supermicro per Next Platform analysis [secondary]. A **$5B+ GB200 AI-server deal** with xAI was in advanced talks (Bloomberg, Feb 2025); finalization status in 2026 could not be confirmed [unverified].
- Dell AI-server scale: ~420,000 GPUs sold in FY2026; ~616,000 GPUs total over three years of Dell's AI-server business; 4,000+ enterprise/neocloud/sovereign AI customers; average customer ~142 GPUs [independent — Next Platform analysis, 2026-03]. These server wins carry PowerSwitch AI-fabric attach, though Dell does not disclose networking attach rates [gap].
- Traditional PowerEdge systems grew only 5.8% in FY2026 to $18.41B (44.4% of systems sales) vs AI servers dominating — Dell's networking is increasingly AI-fabric-led [independent — Next Platform].
- Dell networking segment revenue is **not separately disclosed**; market share figures for Dell switches in AI fabrics were not located [gap].

### 1.7 Pricing snapshots (official Dell stores, 2026)

Dell publishes component pricing; switch list prices are quote-based. Observed 2026 store prices [official — dell.com regional stores]:
- Dell Networking SFP+ 10G SR transceiver: ~€96.51 (Ireland); 100G QSFP28→4×SFP28 DAC breakout 3 m: ~€440.87; 25G SFP28 DAC 3 m: ~€94.06 [official].
- 25G SFP28 AOC 2 m: RM 1,781 (Malaysia); QSFP28→4×SFP28 DAC 5 m: HK$7,422 (Hong Kong) [official].
- Full switch pricing is quote-only ("Shop Now" leads to configuration/quote) — no public list prices for Z9864/Z9964 [gap].

---

