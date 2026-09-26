---
id: etape5-trackd-servers/03-data-center-networking-for-ai-2026/1-5-nvidia-spectrum-x-switches-sn5000-class-sn5600-spectrum-
title: "1.5 NVIDIA Spectrum-X switches (SN5000-class / SN5600 / Spectrum-X800)"
domain: data-center-networking-for-ai-2026
role: deep-dive
task: actor-profile
actors: ["Broadcom", "Nvidia"]
dates: ["2026-07"]
keywords: ["nvidia", "asic", "cpo", "ethernet", "gpu", "hyperscaler", "nvlink", "rubin"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [728, 744]
section: "Data-Center Networking for AI (2026)"
sha256: 1b3a108036b1f5abd3e03879b1fc4f8a19a36798cb34422c988451c31f04919d
---

# 1.5 NVIDIA Spectrum-X switches (SN5000-class / SN5600 / Spectrum-X800)

- Cisco introduced **Silicon One G300** — its own 102.4 Tbps networking silicon — debuting in **liquid-cooled N9000 and 8000 series switches**, aimed at competing with Broadcom and NVIDIA switching chips. [secondary] (ainvest)
- Simultaneously, Cisco ships **N9100 series switches explicitly powered by NVIDIA Spectrum-X Ethernet silicon**, running Cisco NX-OS on NVIDIA hardware. [secondary] (ainvest)
- Cisco FY2026 hyperscaler AI design wins: three new wins in Q4 alone — one **Silicon One P200 scale-across** deployment, one **G200 scale-out** project, and one optical line-system deployment; management flagged line-of-sight to more wins over the next six months across G300, G200, P200 and A100 platforms. [vendor-reported] (infotechlead)
- TrendForce platform table (2025 vintage, cited 2026): Cisco Silicon One G200 series at 51.2 Tbps (2023) alongside Marvell Teralynx 10; CPO prototypes. [secondary]

### 1.5 NVIDIA Spectrum-X switches (SN5000-class / SN5600 / Spectrum-X800)

- **Spectrum-X800 Ethernet switch** (SN5600-class): 800 Gb/s per port, 51.2 Tb/s switching capacity (per temperature2.com / TrendForce table). Based on the **Spectrum-4 switch ASIC**; SN5600 = 64-port 800GbE. [secondary]
- **Spectrum-6** (announced July 2026): 102.4 Tbps Ethernet switch system, 2× previous-generation capacity, integrated into the Rubin-architecture stack (Vera CPU, Rubin GPU, NVLink 6, ConnectX-9 SuperNICs, BlueField-4 DPUs, Spectrum-6). [vendor-reported] (cxotoday citing NVIDIA)
- TrendForce roadmap table: Spectrum-X800 (800G/port, 51.2 Tbps) → Spectrum-X1600 (1.6T/port, 102.4 Tbps, paired with Rubin) → Spectrum-X3200 (3.2T/port, 204.8 Tbps). ConnectX-8 (800G, PCIe 6.0) → ConnectX-9 (1.6T, PCIe 7.0) → ConnectX-10 (3.2T, PCIe 8.0). [secondary]

### 1.6 Celestica's position

- Celestica regained #1 in Ethernet AI backend networks in Q1 2026 per Dell'Oro (above). Celestica is NVIDIA's primary Ethernet-switch manufacturing partner for Spectrum-X systems; its ranking reflects white-box/OEM volume into AI fabrics. [independent] [secondary]

---

