---
id: etape5-trackd-servers/01-server-vendors-wave-1-supermicro-dell-research-report-as-of-/overview
title: "Server Vendors — Wave 1: Supermicro + Dell (Research Report, as of 2026-09-22)"
domain: server-vendors-wave-1-supermicro-dell-research-report-as-of-
role: deep-dive
task: reference
actors: ["Intel", "Nvidia"]
dates: ["2026-01", "2026-03", "2026-09-22"]
keywords: ["research", "accelerator", "blackwell", "compute", "ethernet", "governance", "gpu", "gpus", "intel", "liquid cooling", "memory", "nvidia"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [28, 84]
section: "Server Vendors — Wave 1: Supermicro + Dell (Research Report, as of 2026-09-22)"
sha256: 2387a87642cea67dde03eb3cc457a583dea0d788bf682d6a9fa0f1e05cd2cf14
---

# Server Vendors — Wave 1: Supermicro + Dell (Research Report, as of 2026-09-22)

> Scope: 2026 product platforms/launches, specs, pricing, customer wins, AI revenue
> figures, governance/accounting updates, controversies. Provenance tag on EVERY claim:
> **[official]** = vendor press release / official IR or corporate page;
> **[vendor-reported]** = figure stated by the vendor in earnings/report;
> **[independent]** = press or tech media with its own reporting (ServeTheHome-style,
> Reuters, Bloomberg, DCD, Tom's Hardware, DIGITIMES…);
> **[secondary]** = aggregators, analyst blogs, market commentary (lower reliability);
> **[unverified]** = cannot be confirmed from the sources collected.
> Dates given only where a source states them; search-crawl dates are NOT treated as
> publication dates. Any figure without an explicit source state is flagged **[unverified]**.

---

## PART 1 — SUPERMICRO (Super Micro Computer, Inc., NASDAQ: SMCI)

### 1.1 2026 AI server platforms

Supermicro's 2026 AI portfolio is built around its Data Center Building Block Solutions
(DCBBS) approach and NVIDIA's Blackwell / Blackwell Ultra platforms **[official]**:

- **NVIDIA HGX B300 8-GPU systems** in two form factors: a compact **2OU (OCP)**
  variant and a **4U** variant, both with Supermicro **DLC-2** direct liquid cooling
  **[official]**. The 2OU OCP-based rack-scale design supports **up to 144 GPUs in a
  single rack** **[official]**. Each HGX B300 system carries **2.1 TB of HBM3e GPU
  memory**; cluster-level compute-fabric throughput is **up to 800 Gb/s** via integrated
  NVIDIA ConnectX-8 SuperNICs with NVIDIA Quantum-X800 InfiniBand or Spectrum-4
  Ethernet **[official]**.
- **NVIDIA HGX B300 NVL16** systems (announced March 2026 as part of the Blackwell Ultra
  wave), designed "for every data center" **[official]**.
- **NVIDIA GB300 NVL72** (Blackwell Ultra, rack-scale Grace-Blackwell) — launched in
  Supermicro's March 2026 Blackwell Ultra portfolio **[official]**.
- **NVIDIA GB200 NVL72** liquid-cooled systems (Blackwell) — launched late 2024, with
  samples shipping to customers in Q4 2024 and full production from end of Q4 2024
  **[official via secondary (Zhitong Finance, Oct 2024)]**; still part of the 2026
  lineup alongside GB300 NVL72 **[official]**.
- **NVIDIA HGX B200** 8-GPU systems (Blackwell, pre-Ultra generation) **[official]**.
- **GB200 NVL4 HPC** rack-scale solutions **[official]**.
- **Super AI Station** — deskside AI supercomputer based on NVIDIA GB300
  **[official]**.
- **NVIDIA RTX PRO 6000 Blackwell Server Edition** systems **[official]**.
- Planned **NVIDIA Vera Rubin NVL144 and Vera Rubin NVL144 CPX** platforms for 2026
  delivery, announced at NVIDIA GTC Washington, D.C. **[official]** (see 1.2).
- **4-GPU platforms**: Supermicro's March 2026 Blackwell Ultra announcement also covers
  smaller-node systems; the HGX B300 NVL16 (8-GPU NVL16-domain) is the mid-scale offer
  **[official]**. Standalone 4-GPU 2026 platforms are marketed via the
  accelerator/GPU-server portfolio **[official]**; no dedicated new 4-GPU Blackwell
  Ultra launch SKU was identified in collected sources — **[unverified]** whether a
  2026 4-GPU flagship exists beyond the RTX PRO 6000 workstation/server line.
- **SuperBlade**: in January 2026 Supermicro launched **SBI-622BA-1NE12-LCC**, a
  high-density liquid-cooled blade server "powered by dual Intel Xeon 6900 series
  processors," described by CEO Charles Liang as "the most core-dense SuperBlade we've
  ever created," with up to **3 TB of memory** support and 93% cabling reduction
  **[secondary (Barchart via greencitylivestock.com, Jan 2 2026)]**; target workloads:
  finance, climate modeling, scientific research **[secondary]**.

