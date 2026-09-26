---
id: ai-industry-kb-2026/15-hardware-chips/additional-dated-datapoints
title: "Additional dated datapoints"
domain: hardware-chips
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Anthropic", "Broadcom", "Cerebras", "DeepSeek", "Google", "Groq", "Meta", "Nebius", "Nvidia", "OpenAI", "SpaceX", "United States", "xAI"]
dates: ["2025-10", "2026-03", "2026-03-19", "2026-04-27", "2026-09-16", "2026-09-20"]
keywords: ["acquisition", "amd", "aws", "blackwell", "claude", "compute", "cost", "deepseek", "foundry", "gpu", "gpus", "hbm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7539, 7557]
section: "15. Hardware & Chips"
sha256: ed5b80ed6efe57fd8da89d9b8a3c3c4a25992691dab29fb1abc90482a4ad6f9c
---

# Additional dated datapoints

- **Q1 2026 sourcing:** new H100 wait times 6+ months → **2–3 months**; Blackwell-class lead times moved the opposite way: **3–7 months**; allocation "highly unstable across distributors."
- **September 16, 2026 (Silicon Data):** **B200 residual value at 158% of launch price** — the market values installed Blackwell above new list ~1 year after mass availability; A100/H100 residuals also above straight-line depreciation.
- **September 20, 2026:** structural shortage confirmed across HBM, foundry, and power; Jensen Huang says enterprises and clouds buy GPUs for both training and inference, requiring "sustained, high-volume compute capacity."
- Utilization paradox: field reports of ~9% average GPU utilization (single jobs pinned to whole A100s); the 2026 fix is GPU sharing (MIG partitions, time-slicing, Karpenter autoscaling on spot), not more hardware.
- Demand wall: OpenAI ~26 GW secured pipeline (NVIDIA $100B/10 GW, AMD 6 GW, Broadcom); Stargate $500B/10 GW; xAI Colossus 2 (1M GPUs); NVIDIA FY2026 record **$215.9B** Data Center revenue (Mar 2026).
- Inference price floor verified (Sep 2026): DeepInfra gpt-oss-20b **$0.03/$0.14** per 1M (input/output); gpt-oss-120b $0.04/$0.17; Llama-3.1-8B $0.03/$0.05; SemiAnalysis: ~$0.20/1M for B200-served DeepSeek R1. Inference costs fell **280-fold in two years** (Jan 2026, citing Deloitte/SDxCentral) — Jevons paradox in effect.
- **Export-control hardware events (see §18):** DOJ/Super Micro indictment unsealed March 19, 2026 (Manhattan federal court); NDRC prohibited Meta–Manus April 27, 2026.

### Additional dated datapoints

- **GB200 vs 8×B200 decision rule (2026):** for models under ~100B parameters, 8×B200 HGX nodes are "almost always sufficient and far more cost-effective"; NVL72 earns its keep at 200B+ parameter training needing the 130 TB/s all-to-all domain, or 671B-scale reasoning models that must hold the full model in one rack.
- **AMD rental strategy note:** MI300X's value prop in rental is memory-per-dollar (192 GB vs 80 GB H100 — a 70B FP16 model fits on one card), not kernel throughput; ROCm tuning maturity still trails CUDA for custom-kernel work.
- **Cerebras serving modes:** **native** (model fits on wafer — full speed advantage) vs **Weight Streaming** (larger models stream weights from MemoryX — narrows the advantage); AWS Marketplace integration reached GA (provision via console, bill against EDPs; pay-per-token and dedicated reservations).
- **GroqCloud continuity:** GroqCloud per-token pricing in 2026 ($0.05/$0.08 for 8B-class, $0.59/$0.79 for 70B-class) reflects the **prior LPU generation**; the LPX generation's first production route is enterprise/cloud (Nebius), not the public API.
- **October 2025:** Google–Anthropic up-to-1M-TPU / 1-GW commitment — frames why Rubin faces non-NVIDIA competition at the frontier-training tier. **Reported:** Anthropic leasing SpaceX Colossus 1 (220K GPUs) with Claude rate limits doubled on the back of it [SECONDARY].
- **xAI Colossus 2:** 1M-GPU scale target (→ §14 for capacity detail).
- **March 2026 purchase-price snapshot (sourcebyspec.com):** L40S 48GB $8,610–$8,900; RTX 6000 Ada 48GB $7,400–$7,800; RTX PRO 6000 96GB $9,450–$9,800; new H100 80GB SXM $25,000–$35,000 (used $18,000–$22,000); A100 80GB discontinued new, $12,000–$18,000 used. Highest-pull SKUs: H200 NVL PCIe, H100 NVL PCIe, L40S, L4, RTX 4000 Ada.
- **Spelling and structure corrections:** brief wrote "Sohgo" → verified product name **Sohu**; brief's "Groq acquisition" → **non-exclusive license**.

