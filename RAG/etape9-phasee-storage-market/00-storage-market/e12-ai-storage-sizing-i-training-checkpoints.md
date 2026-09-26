---
id: etape9-phasee-storage-market/00-storage-market/e12-ai-storage-sizing-i-training-checkpoints
title: "E12 — AI storage sizing I: training checkpoints"
domain: step-9-phase-e-storage-memory-market-2026
role: deep-dive
task: training
actors: ["AWS", "Google", "Nvidia"]
dates: []
keywords: ["training", "accelerator", "attention", "aws", "benchmark", "cost", "dram", "embedding", "fp8", "gpu", "gpus", "gqa"]
source: docs/RAG/etape9_phaseE_storage_market.md
source_anchor: ""
source_lines: [147, 184]
section: "Step 9 Phase E — Storage & Memory Market 2026"
sha256: cfa63ba502151d09411538047494bd06bd1c6ee480a3a1552f7231a5eb44fa3b
---

# E12 — AI storage sizing I: training checkpoints

- **Why 2026 is high-risk:** with 1 TB retail SSDs at 2×+ their historical prices, the incentive to sell relabeled, used-as-new, or capacity-spoofed drives rises; community guidance converges on a verification routine [secondary].
- **Detection routine:**
  1. **Verify capacity with a full-surface write test** (f3/h2testw class tools) — spoofed-capacity drives report fake sizes until written past real NAND [secondary].
  2. **Read SMART/NVMe identify data** — check model string, firmware version, and serial against the vendor's format; mismatched or generic strings are a flag [secondary].
  3. **Check TBW/percentage-used on "new" drives** — any nonzero wear on a sealed-new unit indicates a used drive [secondary].
  4. **Weigh and inspect** — counterfeits often use lighter PCBs and fewer NAND packages than genuine units; compare against teardown photos [secondary].
  5. **Buy from traceable channels** — marketplace sellers with SMART reports beat anonymous bulk lots; Amazon Renewed is convenient but pricier, eBay offers selection with buyer protection [secondary].
- **Enterprise-drive-specific traps:** Dell/HP-branded pulls sold without sleds (see E8); drives with **vendor-locked firmware that cannot be updated outside the OEM ecosystem**; and "refurbished" drives that are actually failed-stock returns rather than proactive-refresh pulls — ask for origin explicitly [secondary].
- **Firmware-health check on receipt:** update to the latest vendor/OEM firmware where entitled, then re-run SMART — a drive that cannot complete a firmware update or shows reallocated sectors climbing in the first 48 hours goes back under RMA [secondary].

## E12 — AI storage sizing I: training checkpoints

- **The core formula (mixed-precision training, BF16/FP16 weights + FP32 optimizer states): a full checkpoint needs ~8–12 bytes of storage per parameter** — 2 bytes for weights plus ~8 bytes for Adam optimizer states (momentum + variance in FP32), plus gradients and metadata [secondary]. Source: https://www.cudocompute.com/blog/storage-requirements-for-ai-clusters
- **Worked sizes (8–12 B/param rule):**
  - 7B → weights ~14 GB, **full checkpoint ~70 GB** [secondary].
  - 70B → weights ~140 GB, **full checkpoint ~700 GB** [secondary].
  - 175B → weights ~350 GB, **full checkpoint ~1.75 TB** [secondary].
- **Independent benchmark anchor (Argonne DLIO):** Llama 70B (80 layers, 8192 hidden, 128k vocab) measured **checkpoint 1.1 TB**; Llama 405B **6 TB**; Llama 1T **17 TB** — consistent with the 8–12 B/param band once vocab/embedding overhead is included [secondary]. Source: https://github.com/argonne-lcf/dlio_benchmark/wiki/How-to-run-LLM-benchmark
- **AWS's reference decomposition (100B model):** BF16 weights 200 GB + optimizer 800 GB (8 B/param) = **1 TB single checkpoint per model replica**; with data-parallel synchronization, **only one replica's state needs saving** — checkpoint size does not scale with job size [secondary]. Source: https://aws.amazon.com/blogs/storage/architecting-scalable-checkpoint-storage-for-large-scale-ml-training-on-aws/
- **The key scaling insight (Hammerspace/Blocks&Files):** checkpoint size is a function of *model* size, not *cluster* size — the 405B Llama-3 checkpoint trained on 16,000 GPUs is the same size as on 3 nodes; only "under a hundred terabytes for state-of-the-art LLMs" needs saving per checkpoint [secondary]. Source: https://www.blocksandfiles.com/ai-ml/2025/02/04/very-large-ai-model-training-uses-object-storage/1602990
- **Checkpoint count math:** keep N retained checkpoints → N × checkpoint size of *fast* storage. With failures every ~2.8 hours observed on the Llama-3 405B run, checkpointing is frequent; async/lazy snapshot techniques (DataStates-LLM, 2026 research) decouple state abstraction from data movement to avoid stalling training [secondary]. Source: http://quantumzeitgeist.com/transformer-models-datastates-llm-achieves-scalable-checkpointing/
- **Tiering guidance:** checkpoints land on **node-local NVMe ("Tier 0")** first (scales linearly with GPU count, no cross-node reads needed for tokens — 60 TB of tokens for the entire 405B run, i.e., 3.75 GB per GPU over 54 days), then flush to parallel file/object storage (Lustre, S3-class) for retention [secondary].
- **Procurement takeaway:** size the *write bandwidth* of the checkpoint tier (GB/s aggregate across nodes), not just capacity — DLIO measured **~132 GB/s mean checkpoint I/O** on 1,024 accelerators for a ~1 TB checkpoint (~8 s per checkpoint) [secondary].

## E13 — AI storage sizing II: inference KV-cache and VRAM

- **The baseline industry formula:** Memory (GB) = Parameters (B) × Bytes/param × (1 + Overhead), with **20–50% overhead** for KV cache, activations, and framework buffers [secondary]. Source: http://gpusmith.com/articles/en/pdfs/llm-inference-hardware-sizing-guide.pdf
- **Weight footprints (BF16/FP16):** 7B → 14 GB; 13B → 26 GB; 34B → 68 GB; **70B → 140 GB**; INT8 halves, 4-bit quarters (70B → 35 GB at 4-bit) [secondary]. Source: https://github.com/mosesy5688-cell/ai-nexus/blob/HEAD/src/pages/knowledge/vram.md
- **KV-cache mechanics:** per token, KV bytes = 2 (K+V) × layers × KV-heads × head_dim × bytes/value. Worked example — **Llama-3.1-8B-Instruct BF16: 128 KiB/token**, so 16K context ≈ 2 GB for one sequence; on a 24 GB card, 0.90×24 − 16.1 ≈ **5.5 GB KV pool ≈ ~45K cached tokens** [secondary]. Source: https://github.com/cachebox-project/inference-cache/blob/HEAD/docs/reference-stack/GPU-RUNBOOK.md
- **70B-class KV figures:** **20–40 GB at 32K context** (FP16, GQA); with 10 concurrent users at 32K, ~112 GB FP16 cache (56 GB FP8) [secondary]. GQA (grouped-query attention) saves up to ~4× vs old MHA designs [secondary]. Source: https://github.com/tieubao/til/blob/HEAD/notes/ai/turboquant-kv-cache-compression.md
- **Extreme-context example:** Llama-3.1 70B FP16 chatbot, 16 users × 128K context: **140 GB weights + 16 × ~38 GB KV ≈ 748 GB total** — this is why long-context serving is a rack-scale (NVL72-class) problem [secondary]. Source: https://medium.com/@donmccasland_57353/strategic-gpu-selection-for-llm-serving-on-google-cloud-a-guide-to-vram-planning-and-82d97de2870c
- **Compression relief (2026):** Google Research's **TurboQuant** (ICLR 2026) compresses KV cache 16-bit → 3–4 bits with near-zero quality loss (~5×) — directly targeting what Jensen Huang called the #1 long-context bottleneck at GTC 2026 [secondary].
- **Offload relief:** LMCache-style tiering moves KV blocks to **host DRAM** — it relieves GPU KV pressure but *adds a host-memory requirement*; it does not shrink the weights footprint [secondary].
- **GPU-per-model table (healthy KV pool, BF16):** ~8B → 24 GB (1 card); 13–34B → 48–80 GB (1 card); **70B → 160–200 GB (2–4 cards)**; 70B FP8/INT4 → 48–80 GB (1–2 cards); 100B+/MoE → 320 GB+ (8 cards) [secondary]. Card count is driven first by *weights fitting*, then by KV pool/concurrency; multi-card needs NVLink/NVSwitch (bare metal, not multi-GPU VMs) [secondary]. Source: https://github.com/cachebox-project/inference-cache/blob/HEAD/docs/reference-stack/GPU-RUNBOOK.md
- **Local-inference buying rules (Memeburn, Sept 2026):** 8–12 GB GPU memory for light chat; **24 GB VRAM (or 32–48 GB unified) for a serious coding assistant**; ~48–64 GB accelerator memory for a good 70B without heavy CPU offload; **96–128 GB for gpt-oss-120b-class / large MoE**; "a model that barely fits is already telling you the machine is too small" [secondary]. Source: https://memeburn.com/how-much-ram-and-vram-do-you-actually-need-for-local-ai-in-2026/

## E14 — Cost per GB: VRAM vs HBM vs DDR5 vs NAND

