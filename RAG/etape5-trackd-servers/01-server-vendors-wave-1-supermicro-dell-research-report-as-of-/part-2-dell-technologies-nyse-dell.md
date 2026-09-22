---
id: etape5-trackd-servers/01-server-vendors-wave-1-supermicro-dell-research-report-as-of-/part-2-dell-technologies-nyse-dell
title: "PART 2 — DELL TECHNOLOGIES (NYSE: DELL)"
domain: server-vendors-wave-1-supermicro-dell-research-report-as-of-
role: deep-dive
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: ["2024-05", "2025-07", "2025-11", "2026-06"]
keywords: ["agentic", "amd", "blackwell", "ethernet", "fp4", "fp8", "gpu", "gpus", "inference", "intel", "liquid cooling", "memory"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [254, 304]
section: "Server Vendors — Wave 1: Supermicro + Dell (Research Report, as of 2026-09-22)"
sha256: 865d2bb6d0e2de26414cef9c9b6e2f378d69f01db995bd55b2fb019ad3ca0d45
---

# PART 2 — DELL TECHNOLOGIES (NYSE: DELL)

## PART 2 — DELL TECHNOLOGIES (NYSE: DELL)

### 2.1 PowerEdge AI servers 2026 — portfolio and launch dates

Dell's AI server portfolio in 2026 centers on the **Dell AI Factory with NVIDIA** and
the PowerEdge XE series **[official]**:

- **PowerEdge XE9680** — Dell's "fastest ramping solution ever" (predecessor to the
  2026 XE generation); 8-GPU platform supporting NVIDIA H100/H200/B200-class GPUs
  **[official]**.
- **PowerEdge XE9780 and XE9785** (air-cooled) + **XE9780L and XE9785L** (liquid-cooled),
  announced as the next-generation XE line with NVIDIA: successors to the XE9680,
  supporting **8-way NVIDIA HGX B300** and delivering **up to 4× faster LLM training**
  vs XE9680; configurations support **up to 192 NVIDIA Blackwell Ultra GPUs** with
  direct-to-chip liquid cooling, customizable to **up to 256 Blackwell Ultra GPUs per
  Dell IR7000 rack** **[official (Dell announcement via TechPowerUp techpowerup.com/336994;
  smetechguru.co.za)]**. The announcement (Dell Technologies World 2025 cycle) also
  introduced: **PowerEdge XE9712 with NVIDIA GB300 NVL72** (rack-scale; **50× more AI
  reasoning inference output, 5× throughput improvement**; new Dell PowerCool
  technology) **[official]**; **PowerEdge XE7745 with NVIDIA RTX PRO 6000 Blackwell
  Server Edition** (available July 2025, 4U, up to 8 GPUs, for physical/agentic AI —
  robotics, digital twins, multimodal) **[official]**; planned support for the **NVIDIA
  Vera CPU** and a new Dell PowerEdge XE server for **Vera Rubin** (Dell Integrated
  Rack Scalable Systems) **[official]**; networking: Dell PowerSwitch SN5600/SN2201
  (Spectrum-X Ethernet) and NVIDIA Quantum-X800 InfiniBand (800 Gb/s)
  **[official]**.
- **PowerEdge XE9785 / XE9785L with AMD Instinct MI355X** (announced at **SC25, St.
  Louis — November 2025**, per HotHardware): dual AMD EPYC CPUs + **8× AMD Instinct
  MI355X per node**; liquid-cooled variant is **3U per node** → **128 MI355X GPUs per
  standard 42U rack**; each MI355X has **288 GB HBM3e**, FP6/FP4 support, ~5 FP8
  PFLOPS per GPU (~40 PFLOPS per air-cooled 10U node); AMD Pensando Pollara 400 NICs;
  Dell PowerSwitch AI fabric (new **PowerSwitch Z9964F-ON / Z9964FL-ON**, liquid-cooled
  networking variant); also **PowerEdge R770AP** with Intel Xeon 6 6900-series
  (up to 128 P-cores per socket, dual socket, 12-channel memory), configurable with
  Instinct accelerators **[independent (HotHardware, Nov 2025)]**. ⚠️ Date note: the
  HotHardware article is from SC25 (Nov 2025), not 2026 — it is included because it
  defines Dell's 2026 AMD platform lineup.
- **PowerEdge XE9680L** (liquid-cooled 4U): supports **8× NVIDIA Blackwell GPUs**;
  "highest possible rack-scale density for NVIDIA GPUs in an industry standard x86
  rack," **33% more GPU density per node**, 20% more PCIe Gen5 slots, double
  North/South network expansion capacity; direct liquid cooling; factory-integrated
  rack-scale deployment **[official (Dell AI Factory update via
  digitalisationworld.com, updated ~June 2026)]**. ⚠️ The XE9680L was first announced
  at Dell Tech World 2024 (May 2024); the June-2026 item may be a re-announcement or
  availability update — flagged **[unverified]** which 2026 event it belongs to.
- **Turnkey rack-scale variants** (announced alongside XE9680L): air-cooled design
  supporting **64 GPUs in a single rack**, or liquid-cooled format with **72 NVIDIA
  Blackwell GPUs in a single rack** **[official]**.
- Dell AI Data Platform / Dell NativeEdge / Generative AI Solutions for Digital
  Assistants updates were announced alongside (non-server software) **[official]**.

