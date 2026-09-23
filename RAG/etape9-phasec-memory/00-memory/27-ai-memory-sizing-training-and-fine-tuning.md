---
id: etape9-phasec-memory/00-memory/27-ai-memory-sizing-training-and-fine-tuning
title: "27. AI memory sizing: training and fine-tuning"
domain: step-9-phase-c-server-accelerator-memory-hardware
role: deep-dive
task: hardware
actors: ["Intel", "Nvidia"]
dates: ["2026-04"]
keywords: ["fine-tuning", "memory", "training", "blackwell", "clearwater forest", "compute", "consumer", "cost", "dram", "fp8", "gpu", "gpus"]
source: docs/RAG/etape9_phaseC_memory.md
source_anchor: ""
source_lines: [486, 543]
section: "Step 9 — Phase C: Server & Accelerator Memory Hardware"
sha256: 6480f51d8259200ede8e26ea77c282768babd2f65bcfb6d190f130803c86d14c
---

# 27. AI memory sizing: training and fine-tuning

## 27. AI memory sizing: training and fine-tuning

- Multipliers over inference weights (community calculator): inference 1.05×, **fine-tuning 2.5×**,
  **training from scratch 4.0×** [secondary].
- Why: Adam optimizer states (~8 bytes/param in FP32: m + v), gradients, and activations; activation
  checkpointing trades compute for memory [secondary].
- Concrete: 70B full fine-tune in BF16 ≈ 140GB weights × 2.5 ≈ 350GB → minimum ~5× 80GB GPUs (practical
  configs use 8×); 7B full fine-tune ≈ 14 × 2.5 ≈ 35GB → fits one 48GB-class card [secondary].
- LoRA/QLoRA changes the math: quantized base weights (e.g., 4-bit) + low-rank adapters make 70B
  fine-tuning feasible on 1–2× 80GB GPUs — the standard homelab/enterprise fine-tune path in 2026
  [secondary].
- Reference fit table (typical, FP16, batch 1, moderate context): 7B → 24GB GPU; 13B → 24–40GB;
  34B → 48GB+; 70B → 2× 80GB or 1× 141GB (H200) / quantized on 1× 80GB; 405B → 8× 80GB-class or
  rack-scale [secondary][independent].

## 28. The GPU memory hierarchy (why HBM exists)

- Latency/bandwidth ladder (April 2026 community reference): L1 ~1ns; L3 ~10–15ns; local DDR5 80–90ns
  at ~50 GB/s/channel (400–600 GB/s/socket); HBM3e in-package 100–150ns at ~7.2 TB/s (Hopper/Blackwell
  class); NVLink 5.0 intra-rack ~150ns at 1.8 TB/s/GPU bidirectional; CXL 2.0 DRAM 170–250ns; NVMe-oF SSD
  50–200µs [secondary].
- Design logic: HBM buys ~10–15× the bandwidth of socket DDR5 at similar latency by stacking DRAM dies
  on a silicon interposer next to the compute die with a 1024/2048-bit interface — at 3–4× the wafer
  cost per GB and with strict thermal limits (HBM4 ~75W per stack projected) [secondary].
- Failure note: HBM failures have become the leading cause of GPU malfunctions in data centers — memory
  reliability is now a first-order fleet-operations concern [secondary].

## 29. VRAM-per-model quick reference (FP16 weights + headroom)

| Model size | FP16 weights | Practical minimum GPU config (2026) |
|---|---|---|
| 1–3B | 2–6 GB | Consumer 8–12GB GPU |
| 7–8B | 14–16 GB | 24GB GPU (comfortable) |
| 13–14B | 26–28 GB | 24GB tight / 40–48GB comfortable |
| 32–34B | 64–68 GB | 80GB GPU or 2× 48GB |
| 70–72B | 140–144 GB | 2× 80GB, or 1× 141GB H200 / quantized on 1× 80GB |
| 140B | ~280 GB | 4× 80GB or 2× 141GB |
| 405B | ~810 GB | 8× 80–141GB (rack-scale) |
[secondary][independent]

- Quantization roughly halves these numbers per step (FP8 ≈ ½, INT4 ≈ ¼ of FP16 weight footprint);
  KV cache at long context can rival weights — always budget it separately (Section 26) [secondary].

## 30. Decision guide: server DRAM in 2026

- **General virtualization / enterprise**: DDR5-5600/6400 RDIMMs, 1DPC, all channels populated; ECC
  RDIMM mandatory; prefer 32Gb-die 128GB modules for density [independent].
- **AI host servers (GPU-dense)**: prioritize bandwidth — DDR5-6400 now, 8000 MT/s RDIMMs on Xeon 6700P/
  Clearwater Forest where validated; plan MRDIMM 8800 for 2027 refreshes [secondary].
- **In-memory databases / SAP HANA-class**: capacity first — LRDIMM/3DS up to 256GB+/slot, accept the
  latency trade; CXL Type 3 E3.S modules as a second capacity tier where the platform supports it
  [independent].
- **Homelab**: used DDR4 ECC RDIMMs remain the value king (~$8/GB secondary); DDR5 ECC UDIMM only on
  boards that support it (note: DDR5 ECC UDIMMs are physically keyed/notch-compatible only with
  ECC-capable boards — verify QVL) [secondary].
- **Do not** mix RDIMM/LRDIMM/UDIMM generations or speeds in one system; follow the platform vendor's
  population rules for Chipkill/SDDC modes [secondary].

