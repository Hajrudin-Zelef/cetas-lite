---
id: etape5-trackb-amd/00-amd/3-instinct-mi400-series-roadmap-status-announced-specs-timel
title: "3. INSTINCT MI400 SERIES — roadmap status, announced specs, timeline"
domain: step-5-track-b-amd-hardware-instinct-gpus-epyc-ryzen-cpus
role: deep-dive
task: hardware
actors: ["AMD", "Google", "Meta", "Nvidia", "OpenAI", "TSMC"]
dates: []
keywords: ["mi400", "2nm", "accelerator", "amd", "compute", "ethernet", "fp4", "gpu", "gpus", "hbm4", "helios", "inference"]
source: docs/RAG/etape5_trackB_amd.md
source_anchor: ""
source_lines: [99, 151]
section: "Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)"
sha256: 31af16e2b78f3e98e9e366b855a63f9dbcc8d7777d76881cd5b9085b9c43bcb7
---

# 3. INSTINCT MI400 SERIES — roadmap status, announced specs, timeline

## 3. INSTINCT MI400 SERIES — roadmap status, announced specs, timeline

### 3.1 Naming and positioning
- The **MI400 series** is AMD's next-generation AI/HPC accelerator family (CDNA 5), announced at **CES 2026** (Jan 2026) with flagship **MI455X**, and detailed further at **Advancing AI 2026** (Jul 22–23, 2026) and **Hot Chips 2026** (Aug 2026) **[independent — https://www.theregister.com/special-features/2025/06/12/amd-shines-a-light-on-its-helios-rack-scale-compute-platform/1352958; https://markets.financialcontent.com/bentoncourier/article/tokenring-2026-1-8-amd-ignites-the-yotta-scale-era-unveiling-the-instinct-mi400-and-helios-ai-infrastructure-at-ces-2026; https://www.servethehome.com/amd-helios-mi400-system-architecture-at-hot-chips-2026/]**.
- ⚠️ Naming caveats: sources variously reference "MI450", "MI455X", and "MI430X". The flagship rack part is **MI455X**; **MI430X** is positioned for sovereign AI/HPC offerings (e.g., Cirrascale's sovereign-AI cloud) **[secondary — https://www.businesswire.com/news/home/20260714481734/en/Cirrascale-Advances-Open-AI-Infrastructure-with-AMD-Helios-Rackscale-Solution-and-AMD-Instinct-MI400-Series-GPUs; https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/amd-deep-dive.md]**. A "custom MI450 variant" is the named vehicle for the Meta deal — see §7.
- One source (low-tier aggregator) claimed MI400 is "slated for 2026" at ~10× MI300X performance — treat the 10× figure as **[unverified — https://webpronews.com via search result]**.

### 3.2 Announced specs (MI455X)
- **CDNA 5**, TSMC **2nm (N2)** process, **320 billion transistors** **[official via press — https://markets.financialcontent.com/bentoncourier/article/tokenring-2026-1-8-amd-ignites-the-yotta-scale-era-unveiling-the-instinct-mi400-and-helios-ai-infrastructure-at-ces-2026]**.
- Memory: **432 GB HBM4 per GPU**, **19.6 TB/s** peak bandwidth **[official via press — https://www.theregister.com/special-features/2025/06/12/amd-shines-a-light-on-its-helios-rack-scale-compute-platform/1352958; https://markets.financialcontent.com/bentoncourier/article/tokenring-2026-1-8-amd-ignites-the-yotta-scale-era-unveiling-the-instinct-mi400-and-helios-ai-infrastructure-at-ces-2026]**.
- Compute: **40 petaFLOPS FP4** per GPU (OCP MXFP4) **[independent — https://www.theregister.com/special-features/2025/06/12/amd-shines-a-light-on-its-helios-rack-scale-compute-platform/1352958]**.
- Interconnect: **UALink** open standard for scale-up — AMD claims **3.6 TBps per-accelerator** bandwidth, "just as fast as NVLink" **[independent — The Register]**; AMD opened its Infinity Fabric technology to partners in 2024, foundational to UALink **[independent — same]**.
- **Timeline:** volume shipments **H2 2026**; first 1 GW customer deployments (Meta, OpenAI) scheduled to begin H2 2026 **[secondary — https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/amd-deep-dive.md; https://markets.financialcontent.com/ms.intelvalue/article/marketminute-2026-2-25-amd-secures-historic-6-gigawatt-ai-infrastructure-deal-with-meta-platforms]**.
- Beyond MI400: **MI500 series** announced as the 2027 next generation at AMD's Analyst Day **[secondary — https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/amd-deep-dive.md]**.

### 3.3 Analyst framing
- Coverage frames MI400/Helios as AMD's answer to NVIDIA Vera Rubin NVL144, with the differentiation being **memory capacity leadership (432 GB HBM4 vs competitors)** and an **open-standards stack** (UALink, UEC Ethernet, OCP) vs NVIDIA's closed NVLink ecosystem **[independent — https://www.theregister.com/special-features/2025/06/12/amd-shines-a-light-on-its-helios-rack-scale-compute-platform/1352958; https://www.servethehome.com/amd-helios-mi400-system-architecture-at-hot-chips-2026/]**.

---

## 4. HELIOS RACK PLATFORM — status

### 4.1 Overview
- **Helios** is AMD's rack-scale AI infrastructure blueprint/platform, co-developed with Meta through the **Open Compute Project (OCP)**, unveiled at CES 2026 and detailed at Advancing AI 2025/2026 and Hot Chips 2026 **[independent — https://www.theregister.com/special-features/2025/06/12/amd-shines-a-light-on-its-helios-rack-scale-compute-platform/1352958; https://www.servethehome.com/amd-helios-mi400-system-architecture-at-hot-chips-2026/]**.
- Target workloads: frontier-model training and large-scale inference serving; AMD markets it as "exaflop-class AI performance" in a single rack **[official via press — https://www.businesswire.com/news/home/20260714481734/en/Cirrascale-Advances-Open-AI-Infrastructure-with-AMD-Helios-Rackscale-Solution-and-AMD-Instinct-MI400-Series-GPUs]**.

### 4.2 Configuration (Hot Chips 2026 details)
- **72× MI455X accelerators** per rack, paired with 6th Gen EPYC "Venice" CPUs **[independent — https://www.theregister.com/special-features/2025/06/12/amd-shines-a-light-on-its-helios-rack-scale-compute-platform/1352958]**.
- Rack: **44OU ORW-HPR chassis**, 18 compute trays + 6 switch trays (72-GPU configuration); blind-mate quick-disconnect liquid cooling, redundant PSUs, 50V DC LC busbar; **OCP ORW-aligned 4-GPU compute trays** **[independent — https://www.servethehome.com/amd-helios-mi400-system-architecture-at-hot-chips-2026/]**.
- Per compute tray: **4× MI455X + 1× Venice SP7 CPU host**; each GPU node gets **1.8 TB/s/direction UALoE** (UALink-over-Ethernet) scale-up bandwidth across 12× 3×2 links, **coherent Infinity Fabric 128 GB/s/direction** to the CPU, and up to **3× Pensando Vulcano 800 AI NICs** per EAM via UALink **[independent — ServeTheHome]**.
- Scale-up fabric: two **512-port 200G UALoE switch ASICs** per switch tray, 10.8 TB/s/direction, 72 active links per switch, multi-plane all-to-all GPU connectivity; ~7 kW liquid-cooled tray managed by a Ryzen-based Run BMC **[independent — ServeTheHome]**.
- Aggregate performance: up to **2.9 exaFLOPS peak FP4 (OCP MXFP4)**, **31 TB HBM4 memory** per rack, ~**43 TB/s aggregate scale-out bandwidth** (BusinessWire/Cirrascale figures) / **260 TBps of UALink bandwidth** (The Register) **[vendor-reported via partners — https://www.businesswire.com/news/home/20260714481734/en/Cirrascale-Advances-Open-AI-Infrastructure-with-AMD-Helios-Rackscale-Solution-and-AMD-Instinct-MI400-Series-GPUs; https://www.theregister.com/special-features/2025/06/12/amd-shines-a-light-on-its-helios-rack-scale-compute-platform/1352958]**. ⚠️ Figures vary by source; treat as marketing maxima.
- Networking: **Pensando Vulcano superNICs** — **300 GBps scale-out bandwidth per accelerator** over (presumed) 12× 200 Gbps Ultra Ethernet links; scale-out across racks on **UEC-aligned Ethernet** (no proprietary fabric lock-in) **[independent/secondary — https://www.theregister.com/special-features/2025/06/12/amd-shines-a-light-on-its-helios-rack-scale-compute-platform/1352958; https://www.businesswire.com/news/home/20260714481734/en/Cirrascale-Advances-Open-AI-Infrastructure-with-AMD-Helios-Rackscale-Solution-and-AMD-Instinct-MI400-Series-GPUs]**.
- Software: shared-memory fabric managed by **AMD Instinct Fabric Manager**, UALoE transport over Ethernet with ESUN; defense-in-depth security, virtualization, rack-scale software acceleration; full ROCm stack **[independent — ServeTheHome; https://www.businesswire.com/news/home/20260714481734/en/Cirrascale-Advances-Open-AI-Infrastructure-with-AMD-Helios-Rackscale-Solution-and-AMD-Instinct-MI400-Series-GPUs]**.
- **Status:** system architecture public (Hot Chips 2026); customer deployments (Meta, OpenAI first gigawatts) scheduled from H2 2026; **Cirrascale** announced **Jul 14, 2026** it will offer Helios racks + MI455X (and MI430X for sovereign AI/HPC) in its AI innovation cloud — the first neocloud commitment **[official — https://www.businesswire.com/news/home/20260714481734/en/Cirrascale-Advances-Open-AI-Infrastructure-with-AMD-Helios-Rackscale-Solution-and-AMD-Instinct-MI400-Series-GPUs]**.

---

## 5. EPYC TURIN 9005 (5th Gen) — 2026 adoption and news

### 5.1 Specs (launch context, Oct 2024)
- 5th Gen EPYC "Turin" (EPYC 9005), Zen 5 / Zen 5c, SP5 socket, up to **192 cores / 384 threads** (Zen 5c 9965) or 128 Zen 5 cores (9755), 12-channel DDR5-6400, 128 PCIe 5.0 lanes, up to 5 GHz boost, full 512-bit AVX-512 + VNNI + BF16 **[independent — https://WWW.TECHSPOT.COM/news/105100-amd-launches-epyc-9005-turin-processors-up-192.html; https://convergedigest.com/amd-unveils-5th-gen-epyc-cpu-claims-34-of-server-business/]**.
- 27 SKUs at launch; pricing examples: **EPYC 9965 $14,813**, 9755 $12,984, 9015 (8c) $527 **[independent — https://WWW.TECHSPOT.COM/news/105100-amd-launches-epyc-9005-turin-processors-up-192.html]**.

### 5.2 2026 adoption and pricing trends
- **Google** adopted 5th Gen EPYC (Turin) 9005 for new AI server infrastructure (reported via CPU industry press) **[secondary — https://cpu.itbrief.co/posts/google-adopts-amd-5th-gen-epyc-cpus-for-new-ai-servers/]**.
- **OVHcloud** refreshed its Scale dedicated-server range on EPYC 9005 (up to 192c/384 threads, 384/768 GB RAM configs) — announcement dated **Oct 28, 2025** (context) **[secondary — https://github.com/ovh/infrastructure-roadmap/issues/285]**.
- **Meta AI servers** pair Turin EPYC host CPUs with MI350 racks (per AMD rack materials) **[secondary — https://www.tweaktown.com/news/105766/amd-launches-instinct-mi350-series-ai-chips-185-billion-transistors-288gb-hbm3e-memory/index.html]**.
- The **EPYC 9575F** (frequency-optimized, up to 5 GHz) is AMD's SKU positioned for GPU-powered AI host nodes (28% faster AI processing per AMD's claims) **[vendor-reported — https://convergedigest.com/amd-unveils-5th-gen-epyc-cpu-claims-34-of-server-business/]**.
- ⚠️ No major new Turin SKUs were announced in the Feb–Sep 2026 window found in this research; the 2026 EPYC news is dominated by Venice (see §6). 2026 pricing-trend data for Turin SKUs was not located — flagged as a gap.

---

