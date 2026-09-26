---
id: etape5-trackb-amd/00-amd/5-epyc-turin-9005-5th-gen-2026-adoption-and-news
title: "5. EPYC TURIN 9005 (5th Gen) — 2026 adoption and news"
domain: step-5-track-b-amd-hardware-instinct-gpus-epyc-ryzen-cpus
role: deep-dive
task: hardware
actors: ["AMD", "Google", "Meta"]
dates: []
keywords: ["amd", "gpu", "memory", "pricing", "research"]
source: docs/RAG/etape5_trackB_amd.md
source_anchor: ""
source_lines: [137, 151]
section: "Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)"
sha256: 23c244685613820189ea749aed27f743a837d63d76179b3dd5064487b7faaa00
---

# 5. EPYC TURIN 9005 (5th Gen) — 2026 adoption and news

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

