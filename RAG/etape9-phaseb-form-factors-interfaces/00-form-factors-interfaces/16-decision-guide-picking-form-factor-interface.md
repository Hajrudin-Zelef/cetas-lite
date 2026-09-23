---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/16-decision-guide-picking-form-factor-interface
title: "16. Decision guide — picking form factor & interface"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["AMD", "Intel", "Samsung"]
dates: []
keywords: ["benchmark", "compute", "datacenter", "ethernet", "gpu", "intel", "latency", "memory", "nand", "training"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [403, 486]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 5eabe4220f32eb29c785acf2ff76f9043e5768e71e81e3387fc08f81c8e240df
---

# 16. Decision guide — picking form factor & interface

## 16. Decision guide — picking form factor & interface

### 16.1 Workload → form factor

- **AI training storage node (2026)**: 16–36× E3.S Gen5 NVMe, front hot-swap, direct CPU attach or PCIe switch; NVMe/TCP to GPU compute nodes; ZNS/FDP-aware software if available [independent].
- **General virtualization / mixed fleet**: U.3 tri-mode bays (NVMe now, SAS/SATA reuse) + PERC 12 / Smart Array Gen11 [independent].
- **1U edge / telco**: E1.S (12–25 W, dense, hot-swap) [independent].
- **Capacity archive**: E1.L rulers (245 TB class) or 3.5" HDD hybrid with E3.S rear cache tier [independent].
- **Cold bulk**: 3.5" SATA HDDs at ~$25/TB — unbeatable $/TB in 2026 [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).

### 16.2 Interface decision tree

1. Need > 7 GB/s per drive or < 100 µs latency? → NVMe (Gen5 x4) [independent].
2. Need dual-path HA without NVMe multipath maturity? → 24G SAS dual-port [independent].
3. Reusing HDDs + SSDs in one chassis? → tri-mode U.3 backplane + 9600-series controller [independent].
4. Sharing flash across servers? → NVMe/TCP on existing Ethernet; RDMA only if you own the fabric end-to-end [secondary](https://www.eetimes.com/nvme-over-tcp-will-take-time-to-eclipse-rdma/).
5. Power/cooling constrained 1U? → E1.S at 12–25 W, not 25 W U.2 [secondary](https://www.ariat-tech.com/blog/EDSFF-SSD-Guide-E1.S-vs.E1.L-vs.E3.S-vs.E3.L.html).

### 16.3 Homelab / SMB notes

- Used enterprise SATA/SAS SSDs at $78–$127/TB (3.84 TB class) undercut new NVMe by 4–10× — the rational homelab buy in the 2026 NAND price spike [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).
- U.3 tri-mode backplanes in refurbished 16G/Gen11 servers accept cheap used SAS/SATA today and NVMe tomorrow [independent].
- M.2 boot + U.2 data via MCIO breakout cables is the common whitebox pattern; verify bifurcation support before buying [independent].

### 16.4 2026–2028 watch list

- NVMe 2.4 drives with post-quantum crypto and voltage monitoring (spec released Aug 2026; drives TBD) [official](https://www.businesswire.com/news/home/20260804315628/en/NVM-Express-Publishes-Set-of-NVMe-Specifications-Enhancing-Security-Manageability-and-Sustainability-for-AI-Cloud-Enterprise-and-Client-Storage).
- PCIe 6.0 SSDs/hosts (spec 2022, still early in 2026) and CXL 2.0/3.x memory expansion on 17G/Gen12 [secondary].
- PCIe 8.0 pathfinding toward 2028 at 256 GT/s; optical PCIe via the Optical Aware Retimer ECN [secondary](https://convergedigest.com/pci-sig-targets-2028-release-of-pcie-8-0-at-256-gt-s/).
- SAS 48G on the roadmap; >30 TB 2.5" SAS SSDs as the 2027–2028 shelf-replacement milestone [secondary](https://dataintelo.com/report/serial-attached-storage-sas-solid-state-drive-ssd-market)[secondary](https://www.eetimes.com/sas-customers-expect-innovation-backwards-compatibility/).
- EDSFF E3 2T 70 W class drives and x8/x16 NVMe links for Gen6-era SSDs [secondary](https://www.ariat-tech.com/blog/EDSFF-SSD-Guide-E1.S-vs.E1.L-vs.E3.S-vs.E3.L.html).

## 17. Gaps, conflicts, and unverified claims register

1. **Dell R670/R770 E3.S counts**: WWT brief says R670 16× E3.S / R770 44× E3.S; a reseller matrix says "up to 20x EDSFF E3.S" for R670-class and 12× 3.5"/24× 2.5" for R770. Chassis options vary by configuration; prefer Dell's configurator. [conflict — secondary vs secondary]
2. **NVMe 2.2**: no 2.2 Base revision found — the family went 2.1 (2024) → 2.3 (2025) → 2.4 (2026). Do not cite a "NVMe 2.2" [gap closed].
3. **FDP drive availability 2026**: spec (TP4146) and Samsung whitepaper confirmed; no 2026-dated shipping FDP SSD announcements found. [gap]
4. **ZNS deployment scale 2026**: Samsung PM1731a and Samsung/WD MOU confirmed; hyperscale production deployment figures not found. [gap]
5. **PCIe 6.0/7.0 product dates**: 7.0 products "unlikely before 2028"; 6.0 "not yet widespread" — vendor roadmaps vary. [unverified forward-looking]
6. **SAS 48G timeline**: "upcoming" per KIOXIA/EE Times; no ratification date found. [gap]
7. **E1.S/E1.L sustained power**: SNIA gives no separate sustained max below connector limit — platform-dependent. Do not quote a fixed number. [spec gap — official-by-omission]
8. **Supermicro EPYC/Xeon-6 general-purpose EDSFF counts**: not researched; only the Grace Petascale node verified. [gap]
9. **Lenovo ThinkSystem V4 / other OEMs**: not researched. [gap]
10. **SATA Express / U.2 vs "SFF-8639" naming**: U.2 is the marketing name for the SFF-8639 connector usage; some sources conflate them. Kept distinct in this file. [terminology note]
11. **NAND price figures**: Aug 2026 figures from a storage reseller's index (VDURA/TrendForce citations) — directionally useful, not a contract benchmark. [secondary]
12. **Icy Dock datasheet URL**: extremely long CDN URL cited verbatim; link-rot risk high. [unverified long-term]

## 18. Glossary

- **AIC/HHHL**: Add-In Card / Half-Height Half-Length — PCIe card form factor (CEM edge connector).
- **ANA**: Asymmetric Namespace Access — NVMe multipath path-state mechanism.
- **BOSS**: Boot Optimized Storage Solution (Dell) — M.2 RAID boot carrier.
- **CEM**: Card Electromechanical — the PCIe slot/card spec.
- **CMB**: Controller Memory Buffer — host-mappable BAR memory on the NVMe controller.
- **CXL**: Compute Express Link — PCIe-based memory/coherency interconnect.
- **DC-MHS**: Datacenter Modular Hardware System (OCP) — HPM + DC-SCM modular server architecture.
- **DC-SCM**: Datacenter Secure Control Module — management/security module in DC-MHS.
- **EDSFF**: Enterprise and Datacenter SSD Form Factor (SNIA) — E1.S/E1.L/E3.S/E3.L.
- **FDP**: Flexible Data Placement (NVMe TP4146) — host-directed placement hints.
- **HHHL**: see AIC.
- **HPM**: Host Processor Module — compute module in DC-MHS.
- **IBPI**: International Blinking Pattern Interpretation (SFF-8489) — LED patterns.
- **KV**: Key Value command set — Store/Retrieve/Delete/Exist/List instead of LBAs.
- **LFF**: Large Form Factor — 3.5" drives.
- **MCIO**: Mini Cool Edge IO (SFF-TA-1016) — x8 Gen5/Gen6 internal cable connector.
- **NVMe-MI**: NVMe Management Interface — out-of-band/in-band management spec.
- **NVMe-oF**: NVMe over Fabrics — NVMe over RDMA/TCP/FC transports.
- **OCP**: Open Compute Project.
- **PMM**: Pluggable Multipurpose Module (SFF-TA-1034) — slot family E3 can share.
- **SFF**: Small Form Factor — 2.5" drives; also the SNIA spec series prefix.
- **SFF-TA-1001**: U.3 — universal x4 link definition for SFF-8639.
- **SFF-TA-1005**: UBM — Universal Backplane Management.
- **SFF-TA-1006/1007/1008/1009**: E1.S / E1.L / E3 mechanicals / pin-and-signal.
- **SFF-TA-1016**: MCIO connector spec.
- **SFF-TA-1034**: host system slot (PMM/DC-MHS family).
- **SGPIO**: Serial GPIO (SFF-8485) — legacy LED/activity serial bus.
- **SES**: SCSI Enclosure Services — SAS enclosure management protocol.
- **SlimSAS**: SFF-8654 — high-density internal connector (24G tri-mode era).
- **SLM**: Subsystem Local Memory command set (new in NVMe 2.1).
- **UBM**: Universal Backplane Management (SFF-TA-1005).
- **VPP**: Virtual Pin Port (Intel) — per-lane NVMe LED/control virtualization.
- **WAF**: Write Amplification Factor — NAND writes ÷ host writes.
- **ZNS**: Zoned Namespaces — sequential-write zones mapped to NAND.

