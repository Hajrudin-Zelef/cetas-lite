---
id: etape5-trackb-amd/00-amd/6-epyc-venice-9006-6th-gen-zen-6-2026-announcement
title: "6. EPYC VENICE 9006 (6th Gen, Zen 6) — 2026 announcement"
domain: step-5-track-b-amd-hardware-instinct-gpus-epyc-ryzen-cpus
role: deep-dive
task: hardware
actors: ["AMD", "Intel", "Microsoft", "Nvidia", "TSMC"]
dates: ["2026-07"]
keywords: ["2nm", "agentic", "amd", "chiplet", "copilot", "datacenter", "gpu", "helios", "inference", "intel", "lpddr5x", "mi455x"]
source: docs/RAG/etape5_trackB_amd.md
source_anchor: ""
source_lines: [152, 180]
section: "Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)"
sha256: 5932d3caa5027a79fa2c2f02a1303aaa6e003d6bb06ee1770436ee2fc46bb13b
---

# 6. EPYC VENICE 9006 (6th Gen, Zen 6) — 2026 announcement

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

### 7.1 Gorgon Point / Ryzen AI 400 series (announced CES 2026, Jan 6, 2026)
- AMD announced the **Ryzen AI 400 series mobile processors** (codename **Gorgon Point**) — a refresh of Strix Point (Ryzen AI 300): same **Zen 5 / Zen 5c CPU + RDNA 3.5 GPU + XDNA 2 NPU** formula, 4nm-class, with higher clocks and an upgraded NPU **[independent — https://www.techpowerup.com/338398/amd-ryzen-ai-9-hx-475-470-gorgon-point-apus-surface-in-shipping-manifests; https://abit.ee/en/hard/processors/amd-ryzen-ai-400-gorgon-point-processor-apu-zen-5-copilot-strix-halo-ces-2026-npu-en; https://wccftech.com/amd-confirms-zen-6-medusa-cpus-2027-next-gen-gaming-gpu-improved-ai-raytracing/]**.
- Flagship **Ryzen AI 9 HX 475**: 12c/24t (4× Zen 5 + 8× Zen 5c), up to **5.2 GHz** (vs 5.1 on Strix Point), 16-CU Radeon 890M, **L3 up to 36 MB** (vs 34), **NPU up to 60 TOPS** (vs 50 on Strix Point; 55+ TOPS in other configs), LPDDR5X-8533 support (vs 8000), 28 W default TDP **[official via press / secondary — https://abit.ee/en/hard/processors/amd-ryzen-ai-400-gorgon-point-processor-apu-zen-5-copilot-strix-halo-ces-2026-npu-en; https://www.tweaktown.com/news/104174/amds-next-gen-gorgon-point-apus-leaked-zen-5-rdna-3-refresh-expected-in-2026/index.html; https://www.techpowerup.com/338398/amd-ryzen-ai-9-hx-475-470-gorgon-point-apus-surface-in-shipping-manifests]**.
- Also surfaced in shipping manifests: **Ryzen AI 9 HX 470** (5.2 GHz), plus new **Ryzen AI 7 450** and **Ryzen 5/3** entry models **[secondary — https://www.techpowerup.com/338398/amd-ryzen-ai-9-hx-475-470-gorgon-point-apus-surface-in-shipping-manifests]**.
- **Desktop + enterprise expansion (MWC 2026, Mar 2):** **Ryzen AI PRO 400 series** (enterprise-hardened), and **Ryzen AI 7 450G / 5 440G desktop APUs** — positioned by one source as the first **Copilot+-certified desktop processors** (Gorgon Point silicon, XDNA 2 NPU ≥40 TOPS) **[secondary — https://markets.financialcontent.com/pentictonherald/article/marketminute-2026-3-6-amd-solidifies-ai-everywhere-strategy-with-massive-ryzen-ai-400-expansion-into-desktop-and-enterprise-markets; https://abit.ee/en/hard/processors/amd-ryzen-ai-400-gorgon-point-processor-apu-zen-5-copilot-strix-halo-ces-2026-npu-en]**. ⚠️ HP, Lenovo, Dell OEM design wins reported; the "first Copilot+ desktop" claim is thinly sourced — treat cautiously.
- AMD's claimed positioning: flagship Ryzen AI 9 HX 470 shows **1.3× multitasking, 1.7× content creation, +10% gaming** vs Intel Core Ultra 9 288V at comparable TDP (28 W vs 30 W); up to 24h laptop battery life **[vendor-reported — https://abit.ee/en/hard/processors/amd-ryzen-ai-400-gorgon-point-processor-apu-zen-5-copilot-strix-halo-ces-2026-npu-en]**.
- **Roadmap context:** Gorgon Point is a Zen 5 refresh; true next-gen **Medusa Point (Zen 6, N3-class)** arrives **2027** with >10× AI performance claims, plus **Medusa Baby** for mainstream (H2 2027); Strix Halo remains in-market through at least end-2027 **[secondary — https://www.tomshardware.com/pc-components/cpus/amd-mobile-cpu-roadmap-leak-claims-zen-6-arrives-in-2027; https://wccftech.com/amd-confirms-zen-6-medusa-cpus-2027-next-gen-gaming-gpu-improved-ai-raytracing/]**.

