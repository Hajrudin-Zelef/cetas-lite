---
id: etape9-phasec-memory/00-memory/23-cxl-as-the-optane-successor-positioning
title: "23. CXL as the Optane successor (positioning)"
domain: step-9-phase-c-server-accelerator-memory-hardware
role: deep-dive
task: hardware
actors: ["DeepSeek", "Intel", "vLLM"]
dates: ["2023-01-31", "2025-06-30", "2025-10-04"]
keywords: ["accelerator", "awq", "decode", "deepseek", "dram", "fp8", "gpu", "inference", "int4", "intel", "kv cache", "latency"]
source: docs/RAG/etape9_phaseC_memory.md
source_anchor: ""
source_lines: [426, 485]
section: "Step 9 — Phase C: Server & Accelerator Memory Hardware"
sha256: 252eddddb78b3f8e950d92c39bd3f659994af92e31e7f1c101ff2ffc776c80b3
---

# 23. CXL as the Optane successor (positioning)

## 23. CXL as the Optane successor (positioning)

- The industry narrative positions CXL-attached DRAM/pooled memory as the architectural successor to
  persistent-memory DIMMs for the "cheap capacity near the CPU" role: standard DDR behind a CXL
  controller, multi-vendor, no single-vendor lock-in — the opposite of Optane's fate [secondary].
- Key difference: CXL memory is volatile DRAM (capacity tier), not byte-addressable persistent media;
  persistence use cases moved to NVMe SSDs with PLP and software (WAL, journaling) instead [secondary].

## 24. Intel Optane: discontinued

- Intel wound down the entire Optane business after its Q2 2022 earnings miss, writing off $559M in
  inventory; the company judged it "not sufficiently profitable" and non-core [secondary].
- Persistent Memory 300-series "Crow Pass" (DDR-T2 interface, 4000–4400 MT/s, 128/256/512GB modules,
  up to 4TB per 8-channel Eagle Stream platform) was **cancelled effective January 31, 2023**
  [secondary].
- Lenovo's product guide marks Intel Optane Persistent Memory 200 Series as **withdrawn from marketing**
  (October 4, 2025 update) [vendor-reported].
- Intel dropped Optane Persistent Memory support from 5th Gen Xeon Scalable onwards; PMem generations are
  not mixable and were tied to specific Xeon generations (100 series ↔ 2nd Gen, 200 series ↔ 3rd Gen)
  [secondary].
- Existing deployments keep working (5-year warranty honored from purchase date) but get no firmware or
  tooling updates; the Optane PMem 100 Series reached end-of-interactive-support June 30, 2025
  [secondary].
- Technology note: Optane was 3D XPoint — storage persistence of NAND with near-DRAM speed; its
  4K-random-read latency leadership is still cited by enthusiasts in 2026, but the product line is dead
  [secondary].

## 25. NVDIMM and other persistent-memory options, 2026

- NVDIMM-N (DRAM + NAND + supercapacitor for persistence on power loss): still present in niche
  enterprise/backup-accelerator roles, but no 2026 product news of significance was found; the category is
  effectively in maintenance [unverified].
- No JEDEC-standardized DDR5 NVDIMM-P volume products exist in 2026; the "persistent DIMM" slot is
  unfilled post-Optane [unverified].
- Practical replacements: NVMe SSDs with power-loss protection (PLP capacitors), battery-backed write
  caches on RAID controllers, and CXL-attached DRAM for capacity — persistence handled in software
  (WAL/journaling, fsync discipline) [secondary].

## 26. AI memory sizing: inference

- The three VRAM components: **model weights** (dominant, fixed) + **KV cache** (dynamic, per session) +
  activations/overhead (framework ~15%, CUDA buffers ~8%, safety margin ~10%) [secondary].
- Weight math: `params × bytes_per_param`. For a 70B model: FP32 280GB, FP16/BF16 140GB, FP8/INT8 70GB,
  Q5_K_M ~48GB, Q4_K_M ~39GB, INT4 ~35GB, Q3_K_M ~31GB, Q2_K ~22GB [secondary].
- KV-cache formula: `2 × layers × num_kv_heads × head_dim × seq_len × batch × bytes_per_element`
  [secondary]. Worked example (Llama-3.1-8B, FP16): 2 × 32 × 8 × 128 × 2 bytes = 131,072 bytes ≈
  **0.125 MB per token**; 4096 context × 32 sequences ≈ 16GB KV cache [secondary].
- Rule-of-thumb data points: 70B model at 32K context, batch 1, FP16 KV ≈ **10GB** KV cache on top of
  weights; KV cache doubles with context length — long-context inference is disproportionately expensive
  [secondary].
- MoE example (DeepSeek V4 Flash, 70B params / 13B active): ~32KB per token KV → ~32GB KV cache at 1M
  tokens [secondary].
- Fit guidance (vLLM-era): 7B FP16 ≈ 14GB weights; 70B FP16 ≈ 140GB weights → multi-GPU (e.g., 2× H100)
  or quantization; 70B AWQ 4-bit ≈ 35GB weights + KV fits one 80GB-class card [secondary].
- Why bandwidth rules decode: each generated token re-reads all weights + KV cache, so tokens/sec is
  bounded by GB/s — B200's ~8 TB/s nearly doubles H100's ~3.35 TB/s for decode-heavy serving
  [secondary].
- Total-VRAM quick model: `params×bytes + KV cache + ~2–4GB overhead`; vLLM consumes leftover VRAM for
  KV cache, so more headroom = more concurrent requests [secondary].

