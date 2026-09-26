---
id: etape6-phasef2-fpga/00-fpga/20-microchip-polarfire-additional-detail
title: "20. Microchip PolarFire — additional detail"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: hardware
actors: ["Anthropic", "TSMC"]
dates: ["2026-09-16"]
keywords: ["agentic", "claude", "compute", "dsp", "energy", "ethernet", "latency", "mcp", "memory", "model context protocol", "pricing", "serdes"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [318, 335]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: 75340a961196be501ef19638e913e23c9044635b2c47d9280bf97a576cd7a198
---

# 20. Microchip PolarFire — additional detail

- Nexus 2 platform (advance datasheet FPGA-DS-02122-0.73, ©2026): **TSMC 16 nm FinFET**, **65K–220K system logic cells**, 120–520 18×18 sysDSP multipliers, **4–12 Mb** EBR, timing closure up to **350 MHz** typical (EBR/DSP/clocks 625 MHz) `[official]` — source: https://latticesemi-prod-latticesemi-live.vercel.app/view_document?document_id=54437 `[official]`
- Nexus 2 interfaces: **2–8× 16G** multiprotocol SerDes, hardened **PCIe Gen1/2/3/4**, DDR4/LPDDR4 to **2400 Mbps**, DDR3L 1866, hardened MIPI D-PHY to 4.5 Gbps/lane and C-PHY to 7.98 Gbps/trio, up to **349** programmable sysI/O `[official]`
- Nexus 2 security: **AES-256-GCM**, ECDSA-521 / RSA4096 auth, anti-tamper, PUF/unique ID, user-data encryption, side-channel resistance, TRNG `[official]`
- Vendor comparisons for Nexus 2 (unnamed competitor baselines — keep `[vendor-reported]`): up to **3× lower power**, up to **10× more energy-efficient** edge sensor monitoring, up to **3.2× faster MIPI**, up to **10× faster configuration**, up to **5× smaller** form factor `[vendor-reported]` — source: https://www.edn.com/lattice-launches-small-size-fpga-platform/ `[secondary]`
- **Certus-N2** (first Nexus 2 family, LN2-CT): 64K–220K logic cells, up to 135K LUTs, up to **520** 18×18 multipliers, up to **12 Mb** embedded memory, **8× 16G** SerDes (128 Gbps total), up to 7 GPLLs, <**20 ms** instant-on via xSPI flash, packages down to **9×9 mm** (0.5/0.8/1.0 mm pitch), commercial + industrial grades, integrated AES-GCM/ECDSA/RSA engines; **sampling** as of Dec 2024 announcements `[vendor-reported]` — sources: https://www.cnx-software.com/2024/12/12/lattice-unveils-nexus-2-small-fpga-platform-lattice-avant-30-and-avant-50-mid-range-devices-updated-lattice-design-software-tools/ and https://www.hackster.io/news/lattice-semiconductor-launches-the-certus-n2-small-fpga-family-built-around-the-nexus-2-platform-f0e437012f1b.amp `[secondary]`
- Certus-N2 example solutions: industrial machine-vision platform (MIPI CSI-2/LVDS/sub-LVDS, custom ISP, 10GbE), automotive zonal controller/gateway, 10GbE frame-grabber with PCIe output and LPDDR4 frame buffer `[vendor-reported]`
- **Mach-N2** (announced **2026-09-16**, six days before observation date): secure-control FPGAs on Nexus 2, **65K–220K** SLCs across five devices, integrated nonvolatile flash, hardware Root of Trust, **CNSA 2.0-compliant post-quantum cryptography** (ML-DSA, LMS, XMSS auth; **ML-KEM** key exchange), crypto agility, 128–256 Mb config flash, **sub-30 ms** instant-on, up to 3 stored recovery images, PCIe 4.0 / 10G Ethernet / 16G SerDes; **available for order now**, samples shipped to compute/communications customers `[vendor-reported]` — source: https://futurumgroup.com/insights/lattice-mach-n2-turns-the-post-quantum-procurement-gate-into-an-fpga-moat/ `[secondary]`
- **Lattice Prompt** (announced 2026-09-16): free AI FPGA design assistant linking agentic IDEs/LLMs (**Claude Code, Cursor, VS Code**) to **Lattice Radiant** via open **Model Context Protocol (MCP**), orchestrating simulation, synthesis, place-and-route, timing analysis, bitstream generation; vendor-reported **10×+ productivity gains** `[vendor-reported]`
- **Avant 30 and Avant 50** (announced Dec 2024): new midrange entries expanding the Avant portfolio for connectivity and functional capacity; evaluation board listed "coming soon" (status per Sep 2026 article); **no public pricing confirmed** `[secondary]` — source: https://www.hackster.io/news/lattice-semiconductor-launches-the-certus-n2-small-fpga-family-built-around-the-nexus-2-platform-f0e437012f1b.amp `[secondary]`
- Lattice software: **Radiant** and **Propel** updated for Nexus 2/Certus-N2/Avant (RISC-V variants, improved debug, power calculation); solution stacks **sensAI** (edge AI), **mVision** (embedded vision), **Automate** (factory automation), **Drive** (automotive) `[vendor-reported]` — source: https://www.eeworldonline.com/pcie-gen-4-integration-added-to-small-form-factor-fpga/ `[secondary]`
- FIPS 140-3 **Level 2** compliance claimed for Nexus 2 AES-GCM/SHA3-512 functions (vendor claim; certification status should be verified in NIST CMVP listings before quoting as certified) `[vendor-reported]`

## 20. Microchip PolarFire — additional detail

- PolarFire SoC product overview: 5-core RISC-V subsystem (4× RV64GC + 1× RV64IMAC monitor), deterministic low-latency fabric, SECDED on memories, single-event-upset (SEU) immunity positioning for space/defense `[official]` — source: https://ww1.microchip.com/downloads/en/DeviceDoc/Polarfire_SOC_Product_Overview.pdf `[official]`
- PolarFire (non-SoC) focus: lowest-power midrange FPGA, **12.7 Gbps** SerDes, PCIe Gen2, DDR3/4, hardened crypto — suited to always-on, thermally constrained edge and comms workloads `[vendor-reported]`
- RT PolarFire (radiation-tolerant) extends the family into space applications — named here for completeness; detailed 2026 status not researched (gap) `[unverified]`

