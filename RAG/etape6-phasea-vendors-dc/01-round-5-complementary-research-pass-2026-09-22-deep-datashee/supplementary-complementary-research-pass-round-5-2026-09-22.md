---
id: etape6-phasea-vendors-dc/01-round-5-complementary-research-pass-2026-09-22-deep-datashee/supplementary-complementary-research-pass-round-5-2026-09-22
title: "Supplementary / Complementary Research Pass — round 5 (2026-09-22)"
domain: round-5-complementary-research-pass-2026-09-22-deep-datashee
role: deep-dive
task: reference
actors: ["Broadcom", "Nvidia"]
dates: ["2026-01-22", "2026-01-27", "2026-06", "2026-07", "2026-07-10", "2026-09", "2026-09-22"]
keywords: ["research", "asic", "distribution", "latency", "license", "nvidia", "nvlink", "pricing"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [2140, 2185]
section: "Round-5 complementary research pass (2026-09-22) — deep datasheet and radix detail"
sha256: a9394359ba4fe6a8682b6d07c53aae3c2bb9e8a41d52375e0ac2d643a120179f
---

# Supplementary / Complementary Research Pass — round 5 (2026-09-22)

## Supplementary / Complementary Research Pass — round 5 (2026-09-22)

**Scope note:** This round-5 pass adds only material not present in any earlier section (base §§1–8, passes §§A–M, §§N–X, "round 2" §§N–S, "round 3" §§AA–II, "round 3b" §§R3b-A–I, the four "round 4" waves, and "pass #3" §§A–M). It closes five explicit open items from earlier verification logs: Enterprise SONiC H2 2026 release status, Meraki MS210/MS225 street pricing, AOS-CX 10.17 feature delta, Aruba CX 6300/6400 street pricing, and Cisco C9350/C9610 street pricing plus license tiers. Nothing above was modified. The provenance legend from the base applies.

### S1. Enterprise SONiC H2 2026 release status — closes round-4 §OO item 6 (new)

Dell's official code-version bulletin answers the open question directly: as of September 2026 **no Enterprise SONiC 4.6.1 or 4.7.0 has shipped**; 4.6.0 remains the current train [official — Dell support KB 000228560 "Minimum, Recommended, and Latest Code Versions for Networking Products", crawled ~July 2026]:

- **Enterprise SONiC Distribution by Dell Technologies on NVIDIA Spectrum** — products SN2201, SN4700, SN5600, **SN5601**: Latest = **SN-4.6.0**, Release Date **11-June-2026**; Recommended = SN-4.6.0; Minimum = SN-4.6.0 [official]. (The table strings "SN5601" rather than "SN5610" used in the 4.6 spec sheet (§B) — unverified whether this is a distinct variant string or a KB typo; flagged, not corrected.)
- **NVIDIA Switch NVOS Operating System** (Quantum-X800 era) — Q3200-RA, Q3400-RA, Q3401-RD: Latest **25.02.8008** (02-June-2026); Recommended 25.02.7002 (23-Feb-2026); Minimum 25.02.5002/25.02.4014/25.02.6077 by model [official — same KB]. This confirms NVOS 25.02.x as the shipping train for the Quantum-X800 chassis family (§2.3/§R3b-F context).
- **NVIDIA Switch MLNX-OS** — QM9700, QM9701 (Quantum-2/NVLink switch): Latest **3.12.6300** (15-July-2026); Recommended 3.12.2500/3.12.6200 [official — same KB].
- Enterprise SONiC 4.6.0 technical foundation (new detail): Debian 12 "Bookworm" base, Linux kernel 6.1, Broadcom SAI Adapter 15.3.0, Broadcom SDK 6.5.35, FRR 8.2.2, Config DB version_4_6_1 [secondary/independent — stordis.com release breakdown, ~July 2026]. Operational highlights of 4.6.0: route-table overflow detection with automatic recovery and retry; per-prefix route-consistency checker with transient-route exclusion and CLI-triggered recovery — explicitly relevant for large BGP/EVPN fabrics [secondary — stordis.com].
- Automation ecosystem: the official Ansible collection **dellemc.enterprise_sonic** reached **v4.1.0 (2026-01-27)**; v4.0.0 (2026-01-22) added support for Enterprise SONiC 4.5.1 features (AAA accounting options, BAM/MSA interface config, LAG speed/adv_speed) [official — github.com/ansible-collections/dellemc.enterprise_sonic CHANGELOG.rst].
- Reading of the H2-2026 picture: with 4.6.0 dated June 2026 and no 4.6.1/4.7.0 in Dell's recommended-code matrix as of the September 2026 check, the next SONiC release appears to be a **H2-2026 or later event** — the round-4 open item is therefore closed as "not shipped as of 2026-09-22" rather than "shipped and unlocated" [independent assessment].

### S2. Quantum-X800 ship confirmation + XDR software stack 2026 — closes §EE ship-confirmation gap (new)

Quantum-X800 (XDR InfiniBand) is now confirmed **shipping** in 2026:

- Independent status check: **"Status: SHIPPING · Verified 2026-07-10"** [secondary — gpusmith.com Quantum-X800 datasheet page]. Key anchors: Q3400-RA = 4U, **144× 800G XDR (72 twinport OSFP cages), 115.2 Tb/s non-blocking aggregate**, sub-100 ns port-to-port latency, NVIDIA Quantum-3 ASIC; Q3200-RA = 2U, 72× 800G, 57.6 Tb/s; power 2,900 W typical, up to 7,000 W with active cables; both air-cooled, 19-inch; NVIDIA reference design scales a single two-tier fat tree to **10,368 NICs** [secondary — gpusmith.com].
- Ordering part numbers (new): **920-9B36F-00RX-8S0** (Q3400-RA, 4U, 144 ports, 8 PSUs, C2P airflow — Production Release), **920-9B34F-00RX-FS0** (Q3200-RA, 2U, 36 ports over 18 OSFP cages, 4 PSUs — Production Release), **920-9B36P-00RX-8S0** (Q3401-RD, 4U, 144 ports, **48 VDC busbar**, DGX variant — ES phase) [official — networking-docs.nvidia.com "NVIDIA Quantum-X800 (XDR) Clusters", "Tested and Supported Switches" table].
- **2026 Rel2A XDR software stack** (latest official train): ConnectX-8 firmware **40.50.1002**, DOCA-OFED **3.5.0-082000**, MFT **4.37.0-154**, **XDR Switch NVOS 25.03.1010** (superseding the 25.02.8008 in the Dell KB — newer), Quantum-3 firmware **35.2020.1122** (part of NVOS), HPC-X 2.51.0, IButils 2.27.0, OpenSM 5.28.1, **NVIDIA SHARP 3.16.3**, **UFM Enterprise 6.26.1**, **UFM XDR Enterprise Appliance 2.6.1**, **UFM Telemetry 1.24.2** [official — networking-docs.nvidia.com]. 2026 Rel1B train: NVOS 25.02.8008, Quantum-3 FW 35.2016.4060, UFM Enterprise 6.25.1 [official — same].
- ConnectX-8 SuperNICs: C8180 HHHL (800G XDR IB default mode / 2×400GbE, PCIe 6 x16, crypto + secure boot enabled) at **Mass Production**; C8180L partner-cooled variant at Production Release; C8280Z dual-ConnectX-8 mezzanine for GB300 systems at MP [official — same docs]. The 2026 releases deliver as part of the **GB300 release package** (2026 Rel2A) [official].
- In-network computing: Quantum-X800 integrates **SHARP v4** with in-network data aggregation and collective operations (e.g., All-Reduce), claimed up to 9× performance improvement; OSFP interfaces support flexible port breakout (800G → 400G and below) [secondary — ascentoptics.com]. This complements the base §2.3 feature list with deployment-relevant software detail.

### S3. Aruba CX 6300/6400 street pricing — closes round-4 §OO item 7 (new)

The CX 6300 access/aggregation tier was unpriced in earlier passes. September 2026 street anchors [all secondary — reseller listings; hardware only, license/Central subscription separate]:

| SKU | Model | Price (USD) | Source |
|---|---|---|---|
| JL663A | 6300M 48× 1G RJ45 + 4× SFP56 | **$3,432.75** (new, in stock) | buyrouterswitch.com |
| JL668A | 6300F 24× 1G SFP + 4× SFP56 | **$2,785.10** (new, in stock) | buyrouterswitch.com |
| JL667A | 6300F 48× 1G RJ45 + 4× SFP56 | **$4,348.15** (new, in stock) | buyrouterswitch.com |
| JL658A | 6300M 24× 10G SFP+ + 4× SFP56 | **$9,171.80** (new, in stock) | directmacro.com |
| JL660A#OD1 | 6300M 24× 10GBase-T Class 6 PoE + 4× SFP56 | **$3,850.00** (was $4,400) | orangehardwares.com |
| JL660AR | 6300M 24× 10GBase-T Class 6 PoE + 4× SFP56 | **$8,853.31** | orangehardwares.com |
| S0G97A (TAA) | 6300F 48× 1G PoE+ + 4× SFP56 50G | **$9,110.99** (list $11,404.00), 0 units left | avendor.com |
| S0G97A#B2E (TAA) | 6300F 48× 1G PoE+ + 4× SFP56 50G | **$8,376.00**, MSRP **$11,404.00** | SHI (shi.com) |

- Spec anchors confirmed via the same listings: 6300M/6300F L3 stackable, Gen7 ASICs, up to 496 Gbps switching (48-port), 32K MAC table, IEEE 802.1Q/w/s/x, PoE+ budgets per model [secondary]. New strategic detail: the 6300F listing documents **BGP, EVPN, VXLAN and robust security on an access-tier switch** — Aruba's campus-DC OS commonality (AOS-CX everywhere) extends fabric protocols down to the access layer, a positioning point not previously stated [secondary — avendor.com; independent assessment].

### S4. Cisco C9350 street pricing + license tiers — closes round-4 §OO item 5 (new)

Base §R3b-C covered C9350 hardware context; pricing and the unified license structure were open. September 2026 anchors:

