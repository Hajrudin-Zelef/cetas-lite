---
id: collect-261001-general-networking/general-networking/etape5-trackb-amd-4
title: "Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Google", "Meta", "Nvidia", "TSMC"]
dates: ["2026-07"]
keywords: ["amd", "gpu", "2nm", "agentic", "chiplet", "datacenter", "helios", "inference", "lpddr5x", "memory", "mi455x", "nvidia"]
source: docs/RAG/collect-261001-general-networking/etape5_trackB_amd.md
source_anchor: ""
source_lines: [137, 172]
sha256: f6e9ee1720c515cf42aa021d3563c5b6c78c4ed6758f3eb9354e6001603880bc
---

# Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)

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

## 6. EPYC VENICE 9006 (6th Gen, Zen 6) — 2026 announcement

### 6.1 Timeline
- **CES 2026 (Jan 6, 2026):** Lisa Su held up a bare (no-IHS) Venice sample in her keynote, revealing TSMC **InFO_oS** advanced chiplet packaging — 2 I/O dies surrounded by 2 rows of 4 CCDs **[independent — https://abit.ee/en/hard/processors/amd-epyc-venice-processor-server-zen-6-ces-2026-chiplet-infoos-tsmc-256-cores-en; https://www.guru3d.com/story/amd-unveils-256core-epyc-venice-for-helios-mi455x-ai-racks-platform/]**.
- **Advancing AI 2026 (Jul 22–23, 2026), San Francisco:** AMD formally unveiled the 6th Gen EPYC "Venice" (EPYC 9006 series); announced as **entering volume production on TSMC 2nm** — the first HPC/server chip on 2nm **[independent — https://www.martincid.com/technology-sv/amd-epyc-venice-256-cores-2nm/; https://finance.biggo.com/news/1fa76c25-95b5-4fd8-a771-5dd63c39e1a2]**. AMD CTO Mark Papermaster confirmed the July 2026 launch window, starting with server/datacenter EPYC **[independent — https://www.club386.com/amd-confirms-zen-6-launch-date/]**.
- EPYC roadmap through 2030 also published: **Venice (Zen 6) 2026 → Florence (Zen 7) 2028 → Ravenna (Zen 8) 2030** **[secondary — https://www.kad8.com/hardware/amd-epyc-roadmap-venice-florence-and-ravenna-through-2030/]**.

### 6.2 Specs and variants
- **Up to 256 cores / 512 threads** per socket: 8 CCDs × 32 Zen 6c cores (density-optimized), or up to **96 standard Zen 6 cores** (192 threads) for single-thread leadership **[official via press — https://finance.biggo.com/news/1fa76c25-95b5-4fd8-a771-5dd63c39e1a2; https://www.martincid.com/technology-sv/amd-epyc-venice-256-cores-2nm/]**. ⚠️ Whether the 32-core CCDs are full Zen 6 or dense Zen 6c was initially ambiguous in CES coverage; the Advancing AI 2026 reveal clarified the Zen 6 / Zen 6c split.
- **3D V-Cache up to 1,152 MB L3** on stacked variants **[official via press — https://finance.biggo.com/news/1fa76c25-95b5-4fd8-a771-5dd63c39e1a2]**.
- New **SP7 socket**; **16 channels DDR5-8000** (~1.6 TB/s aggregate, ~3× Turin's ~614 GB/s); **MRDIMM** support reaching DDR5-12800-equivalent bandwidth **[official via press — https://www.martincid.com/technology-sv/amd-epyc-venice-256-cores-2nm/; https://finance.biggo.com/news/1fa76c25-95b5-4fd8-a771-5dd63c39e1a2]**.
- **PCIe Gen 6 (64 Gbps)**, 5th-gen Infinity Fabric **[official via press — https://www.martincid.com/technology-sv/amd-epyc-venice-256-cores-2nm/]**.
- TSMC 2nm GAA nanosheet: 10–15% higher perf at same power or 25–30% lower power at same perf, +15% density (TSMC figures) **[independent — https://www.martincid.com/technology-sv/amd-epyc-venice-256-cores-2nm/]**.
- Product lines (4): **EPYC 9006 SP7** (flagship 256c "agentic AI" part, 5 GHz boost, 128 PCIe Gen6 lanes); **EPYC 9006X SP7** (up to 96c, 5.15 GHz, 3× L3 per core, HPC/simulation); **EPYC 9006 SP8** (8–128c, edge/power-constrained); **EPYC 9006 LP "Verano"** (LPDDR5X, 24 channels, SOCAMM2 form factor, 112 Gbps CPU-to-GPU bandwidth for rack-scale AI chassis) **[official via press — https://www.martincid.com/technology-sv/amd-epyc-venice-256-cores-2nm/; https://www.tweaktown.com/news/112793/amd-announces-6th-gen-amd-epyc-venice-cpus-up-to-256-cores-and-512-threads-and-built-for-ai/index.html]**.
- AMD vendor claims: **+70% CPU throughput vs Turin Zen 5**; **2.2× AI inference throughput vs NVIDIA's "Vera" platform** at the 256-core tier; Venice **>3× faster than NVIDIA Vera** (per TweakTown headline) **[vendor-reported — https://www.martincid.com/technology-sv/amd-epyc-venice-256-cores-2nm/; https://www.tweaktown.com/news/112793/amd-announces-6th-gen-amd-epyc-venice-cpus-up-to-256-cores-and-512-threads-and-built-for-ai/index.html]**.
- Manufacturing: initial production at **TSMC Taiwan**, later also **TSMC Arizona** **[independent — https://www.club386.com/amd-confirms-zen-6-launch-date/]**.

---

## 7. RYZEN CLIENT AI — 2026 releases

