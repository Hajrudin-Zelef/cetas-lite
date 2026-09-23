---
id: etape6-phasea-vendors-dc/00-front-matter/supplementary-complementary-research-pass-2026-09-22
title: "Supplementary / Complementary Research Pass — 2026-09-22"
domain: front-matter
role: reference
task: reference
actors: ["AMD", "Broadcom", "Nvidia"]
dates: ["2026-09-22"]
keywords: ["research", "amd", "asic", "cpo", "distribution", "ethernet", "nvidia", "optics", "pricing"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [334, 378]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 4587b6a5ff14f9d7217e1a12eff277055812f6b6870fbb36f136e31ff8626f1c
---

# Supplementary / Complementary Research Pass — 2026-09-22

## Supplementary / Complementary Research Pass — 2026-09-22

**Scope note:** This section is an additive pass. It contains only material gathered after the base report was written and does not duplicate or rewrite §§1–8. Base sections were left untouched. Read the base report first; the provenance legend from the base applies here.

### A. Dell Enterprise SONiC licensing model (structure clarified)

- Dell Enterprise SONiC Distribution is sold as **1-, 3-, or 5-year subscriptions**, offered in two bundles — **Cloud** and **Enterprise** — with **no price difference between the two bundles** [secondary — Datamation review].
- Two pricing tiers exist: **Standard** and **Premium**. Pricing is banded by platform port speed: **1/10 GbE**, **25/100 GbE**, and **400 GbE** [secondary — Datamation].
- Dell does **not publish specific dollar prices** for Enterprise SONiC subscriptions as of 2026-09-22; TrustRadius notes "no pricing plans listed at this time" [secondary — TrustRadius].
- Reference structure from the Broadcom-sold equivalent (Enterprise SONiC Distribution by Broadcom, v4.5 datasheet): per-platform, per-term SmartCare 9×5 subscription part numbers, e.g. **S-SONIC-EB-1G-1/-3/-5** (1G platforms), **S-SONIC-EB-10G-1/-3/-5** (10G platforms), **S-SONIC-EB-100G-1/-3/-5** (25/50/100G platforms), each in 1-, 3-, and 5-year subscription terms [secondary — Broadcom datasheet via direktronik.se]. This illustrates the per-speed, per-term model Dell mirrors; it is not Dell pricing.
- Value-adds Dell cites over community SONiC: scale-out VXLAN EVPN, unified management framework, Ansible automation, silicon telemetry [secondary — Datamation].

### B. Dell Enterprise SONiC 4.6.0 on Broadcom platforms — confirmed (closes base §7 item 14)

- Dell KB **000228560** ("Minimum, Recommended, and Latest Code Versions for Networking Products") lists **Enterprise SONiC Distribution 4.6.0 (released 30-June-2026) as the latest** and **4.5.3 (05-June-2026) as recommended** for the full Broadcom-based lineup: **Z9264F-ON, Z9332F-ON, Z9432F-ON, Z9664F-ON, Z9864F-ON**, S-series (S5212F/S5224F/S5232F/S5248F/S5296F/S5448F-ON, S3248T, S4348F/T), plus N3248/N3248TE and E3248P/E3248PXE families [official — Dell support KB, 2026].
- Dell's Enterprise SONiC 4.6 is officially supported on: **PowerSwitch Z series** (Z9332F-ON, Z9264F-ON, Z9432F-ON, Z9664F-ON, Z9864F-ON), **S series** (S5448F-ON, S4348F/T-ON, S5248F-ON, S5232F-ON, S5224F-ON, S5212F-ON, S3248T), **SN series** (SN5600, SN5610, SN4700, SN2201), **N series** (N3248TE-ON, N3248P), **E series** (E3248PXE, E3248P) [official — Dell Enterprise SONiC spec sheet, © 2026].
- Note: the Dell KB support matrix for Secure Connect Gateway lists older max versions (4.4.2/4.5.0, 4.5.1) for some models — that matrix tracks telemetry-collection compatibility, not the latest shipping OS [official — Dell Secure Connect Gateway support matrix].

### C. Dell PowerSwitch SN6000 series (Spectrum-6) — spec-sheet detail

Dell published the **PowerSwitch SN6000 Series spec sheet** (Spectrum-6-based, 102.4 Tb/s per ASIC) with three models [official — Dell spec sheet, delltechnologies.com]:

| Model | Form factor / optics | Capacity | Notes |
|---|---|---|---|
| **SN6800-LD** | 5RU, liquid-cooled, CPO, MMC-12 800G fibers | 409.6 Tb/s (4 ASICs), 512× 800G ports, up to 2,048 breakouts | "industry's first 4-ASIC co-packaged Ethernet switch for AI scaleout"; UQD8 v2 connectors; 48–60 VDC busbar power |
| **SN6810-LD** | 2RU, liquid-cooled, CPO, UQD8 v2 connectors | 102.4 Tb/s (1 ASIC), silicon-photonics CPO | 48–60 VDC busbar; high-performance Ethernet for AI clusters |
| **SN6600-LD** | 64× OSFP 800 GbE (pluggable) | 102.4 Tb/s class | ⚠️ marked "*Example only. Actual product may vary." in the spec sheet — product details not final |

- Series-wide: **160 MB per-ASIC fully shared packet buffer**, 750,000 shared entries (routes/MAC/ACL), PTP time sync, open NOS ecosystem support, zero-trust security foundation, for leaf/spine/superspine [official].
- Software options on SN6000: **NVIDIA Air** (digital-twin network simulation, "reduces time to first token") and **NVIDIA NetQ** (end-to-end visibility, per-hop RoCE queue telemetry) [official — Dell spec sheet].
- GA date and pricing for the SN6000 series had **not been located** as of 2026-09-22 [gap].

### D. Spectrum-4 lineup clarification (new detail, not a duplicate of base §2.2)

Base §2.2 named SN5600 as "Spectrum-4, 51.2 Tbps, 64× 800G". The official NVIDIA Spectrum SN5600 datasheet adds [official — NVIDIA SN5600 datasheet, via mirror]:

| Model | Ports | CPU | SSD | PSU / fans | Power input |
|---|---|---|---|---|---|
| **SN5610** | 64× OSFP 800GbE + 2× SFP28 25GbE | Octacore AMD | 80 GB NVMe | 4 (2+2), 5 fans N+1 | 200–240 VAC |
| **SN5600** | 64× OSFP 800GbE + 1× SFP28 25GbE | Hexacore x86 | 160 GB SATA-3 | 2 (1+1), 4 fans N+1 | 200–240 VAC |
| **SN5600D** | 64× OSFP 800GbE + 1× SFP28 25GbE | Hexacore x86 | 160 GB SATA-3 | 2 (1+1), 4 fans N+1 | **40–60 VDC busbar** (DC-data-center variant) |

- All three: 51.2 Tb/s, 33.3 Bpps, 8× 100G PAM4 lanes/port, 160 MB packet buffer, 512K-entry scale (MAC/ARP/IPv4/ACL/ECMP/VXLAN shared) [official].
- Sibling **SN5400** (Dell PowerSwitch SN5000 family): 64× QSFP-DD 400GbE + 2× SFP28, 25.6 Tb/s [official — Dell PowerSwitch SN5000 spec sheet via blueally.com].

