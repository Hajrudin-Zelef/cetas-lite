---
id: collect-261001-general-networking/general-networking/etape5-trackb-amd-5
title: "Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "China", "Intel", "Microsoft", "Moonshot", "Nvidia"]
dates: []
keywords: ["amd", "gpu", "benchmark", "copilot", "intel", "kimi", "llama", "llama.cpp", "lpddr5x", "memory", "nvidia", "research"]
source: docs/RAG/collect-261001-general-networking/etape5_trackB_amd.md
source_anchor: ""
source_lines: [173, 193]
sha256: 9bb7f5618b08d42eb9157fba4dd2e36835efb1996869384f1fed3a8d2cfc0d29
---

# Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)

### 7.1 Gorgon Point / Ryzen AI 400 series (announced CES 2026, Jan 6, 2026)
- AMD announced the **Ryzen AI 400 series mobile processors** (codename **Gorgon Point**) — a refresh of Strix Point (Ryzen AI 300): same **Zen 5 / Zen 5c CPU + RDNA 3.5 GPU + XDNA 2 NPU** formula, 4nm-class, with higher clocks and an upgraded NPU **[independent — https://www.techpowerup.com/338398/amd-ryzen-ai-9-hx-475-470-gorgon-point-apus-surface-in-shipping-manifests; https://abit.ee/en/hard/processors/amd-ryzen-ai-400-gorgon-point-processor-apu-zen-5-copilot-strix-halo-ces-2026-npu-en; https://wccftech.com/amd-confirms-zen-6-medusa-cpus-2027-next-gen-gaming-gpu-improved-ai-raytracing/]**.
- Flagship **Ryzen AI 9 HX 475**: 12c/24t (4× Zen 5 + 8× Zen 5c), up to **5.2 GHz** (vs 5.1 on Strix Point), 16-CU Radeon 890M, **L3 up to 36 MB** (vs 34), **NPU up to 60 TOPS** (vs 50 on Strix Point; 55+ TOPS in other configs), LPDDR5X-8533 support (vs 8000), 28 W default TDP **[official via press / secondary — https://abit.ee/en/hard/processors/amd-ryzen-ai-400-gorgon-point-processor-apu-zen-5-copilot-strix-halo-ces-2026-npu-en; https://www.tweaktown.com/news/104174/amds-next-gen-gorgon-point-apus-leaked-zen-5-rdna-3-refresh-expected-in-2026/index.html; https://www.techpowerup.com/338398/amd-ryzen-ai-9-hx-475-470-gorgon-point-apus-surface-in-shipping-manifests]**.
- Also surfaced in shipping manifests: **Ryzen AI 9 HX 470** (5.2 GHz), plus new **Ryzen AI 7 450** and **Ryzen 5/3** entry models **[secondary — https://www.techpowerup.com/338398/amd-ryzen-ai-9-hx-475-470-gorgon-point-apus-surface-in-shipping-manifests]**.
- **Desktop + enterprise expansion (MWC 2026, Mar 2):** **Ryzen AI PRO 400 series** (enterprise-hardened), and **Ryzen AI 7 450G / 5 440G desktop APUs** — positioned by one source as the first **Copilot+-certified desktop processors** (Gorgon Point silicon, XDNA 2 NPU ≥40 TOPS) **[secondary — https://markets.financialcontent.com/pentictonherald/article/marketminute-2026-3-6-amd-solidifies-ai-everywhere-strategy-with-massive-ryzen-ai-400-expansion-into-desktop-and-enterprise-markets; https://abit.ee/en/hard/processors/amd-ryzen-ai-400-gorgon-point-processor-apu-zen-5-copilot-strix-halo-ces-2026-npu-en]**. ⚠️ HP, Lenovo, Dell OEM design wins reported; the "first Copilot+ desktop" claim is thinly sourced — treat cautiously.
- AMD's claimed positioning: flagship Ryzen AI 9 HX 470 shows **1.3× multitasking, 1.7× content creation, +10% gaming** vs Intel Core Ultra 9 288V at comparable TDP (28 W vs 30 W); up to 24h laptop battery life **[vendor-reported — https://abit.ee/en/hard/processors/amd-ryzen-ai-400-gorgon-point-processor-apu-zen-5-copilot-strix-halo-ces-2026-npu-en]**.
- **Roadmap context:** Gorgon Point is a Zen 5 refresh; true next-gen **Medusa Point (Zen 6, N3-class)** arrives **2027** with >10× AI performance claims, plus **Medusa Baby** for mainstream (H2 2027); Strix Halo remains in-market through at least end-2027 **[secondary — https://www.tomshardware.com/pc-components/cpus/amd-mobile-cpu-roadmap-leak-claims-zen-6-arrives-in-2027; https://wccftech.com/amd-confirms-zen-6-medusa-cpus-2027-next-gen-gaming-gpu-improved-ai-raytracing/]**.

### 7.2 Strix Halo (Ryzen AI Max 300) — 2026 adoption and new SKUs
- Strix Halo (Zen 5, up to 16c, 40-CU RDNA 3.5 Radeon 8060S, XDNA 2 50 TOPS, up to 128 GB unified LPDDR5X with up to 96 GB assignable as VRAM) launched Jan 2025; by 2026 it became the **de facto local-LLM workstation chip** **[secondary — https://www.techradar.com/pro/ryzen-ai-max-395-cpu-could-be-amds-sleeper-hit-against-nvidias-ai-dominance-as-nearly-30-strix-halo-mini-ai-workstation-models-hit-the-market-including-some-rather-funky-ones]**.
- **Adoption (2026):** nearly **30 mini AI workstation models** launched by 2026 — Beelink GTR9 Pro, Seaviv AideaStation R1, HP Z2 Mini G1a, GMKtec EVO-X2, Minisforum MS-S1 MAX, Abee AI Station 395 Max; laptops: Asus ROG Flow Z13, Chinese mobile-workstation designs (Sixunited, Linglong, Tianba). Price positioning **$1,800–$2,800** vs $21,000+ for multi-GPU AI servers **[secondary — TechRadar]**. ⚠️ Big-brand availability remains limited; "real-life tests remain scarce" per TechRadar.
- **New SKUs at CES 2026:** **Ryzen AI Max+ 392** and **Ryzen AI Max+ 388** — detuned CPU core counts from the Max+ 395 but **full 40-CU GPU** retained; the 392 surfaced in Geekbench (~2,917 single-core, ~12–15% behind the 395 in multi-core) **[independent — https://hothardware.com/news/amd-ces-2026?ref=thetechstreetnow.com; https://www.notebookcheck.net/New-AMD-Strix-Halo-Ryzen-AI-Max-392-stars-in-early-benchmark-after-CES-2026-debut.1204390.0.html]**.
- **Local-AI positioning:** AMD claims HP Z2 Mini G1a advantages of **1.5×/1.7× tokens-per-second-per-dollar** in LM Studio on GPT-OSS 20B/120B vs NVIDIA DGX Spark **[vendor-reported — https://hothardware.com/news/amd-ces-2026?ref=thetechstreetnow.com]**; community reports ~47–53 tok/s on 120B models vs DGX Spark's 56 tok/s at lower price **[secondary — https://github.com/getnyrex/strix-halo-guide/blob/HEAD/RESEARCH.md]**.
- **Framework Desktop** clusters: 4-node Strix Halo clusters running **1T-parameter Kimi K2.5** via distributed llama.cpp reported by community researchers **[secondary — https://github.com/getnyrex/strix-halo-guide/blob/HEAD/RESEARCH.md]**.
- **Ryzen AI Embedded (Jan 8, 2026):** new embedded portfolio — **P100** (Zen 5, up to 12c, RDNA 3.5, XDNA 2, 50 TOPS, 15–54 W, −40°C to +105°C) for in-vehicle/industrial; **X100** (Strix Halo silicon, up to 16c) for "physical AI"/autonomous systems **[official via press — https://videocardz.com/newz/amd-introduces-ryzen-ai-embedded-p100-strix-point-krackan-and-x100-strix-halo-series]**.
- Rumored roadmap: "Gorgon Halo" (Ryzen AI Max 400, same arch, higher clocks) Q4 2026; "Medusa Halo" with LPDDR6 (~80% more memory bandwidth) **[unverified — https://github.com/getnyrex/strix-halo-guide/blob/HEAD/RESEARCH.md]**.

---

## 8. AMD AI PARTNERSHIPS AND STRATEGIC DEALS — 2026

