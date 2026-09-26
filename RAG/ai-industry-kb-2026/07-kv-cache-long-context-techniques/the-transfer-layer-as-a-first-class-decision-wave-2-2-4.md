---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/the-transfer-layer-as-a-first-class-decision-wave-2-2-4
title: "The transfer layer as a first-class decision (Wave 2, §2.4)"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["AMD", "AWS", "Alibaba", "Google", "Groq", "Huawei", "Meta", "Mistral", "Moonshot", "Nvidia", "SGLang", "Samsung", "vLLM"]
dates: ["2025-10-29", "2025-12", "2026-01-22", "2026-05-05", "2026-07", "2026-08-29", "2026-09-22", "2026-10-20"]
keywords: ["agentic", "amd", "asic", "attention", "benchmark", "blackwell", "compute", "consumer", "decode", "disaggregated", "disaggregated serving", "dram"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3485, 3520]
section: "7. KV Cache & Long-Context Techniques"
sha256: 3c221aadd4e96e02a752dbb6dd0bd9859be64d0ea616d7def51d977eed81f7f7
---

# The transfer layer as a first-class decision (Wave 2, §2.4)

- **December 2025**: NVIDIA acquired Groq's technology and team for **~$20B** (NVIDIA's largest deal ever); Groq was founded by the engineers behind Google's original TPU.
- **GTC 2026 (March 16–17)**: Jensen Huang unveiled the **Groq 3 LPU** — LPU = Language Processing Unit, Groq's term for its inference ASIC; it is not a GPU. The reference deployment routes **prefill → Rubin GPU, decode → LPU**, orchestrated automatically by NVIDIA's **Dynamo** scheduler, with KV-cache transfer between them — hardware-level prefill/decode disaggregation.
- Specs: 500 MB on-chip SRAM per chip at **150 TB/s** bandwidth (~7× a Rubin GPU's 22 TB/s HBM4, at 1/500th the capacity); deterministic compiler-orchestrated execution; FP8 confirmed (1.2 PFLOPS per LPU); 96 C2C links @ 112 Gbps (2.5 TB/s bidirectional); Samsung 4nm; MXM/VXM/SXM execution modules. **Rubin CPX** (same GTC 2026 announcement): the designated prefill workhorse for 1M+ token contexts using cheaper **GDDR7 instead of HBM** — NVIDIA's own admission that prefill (capacity-hungry) and decode (bandwidth-hungry) want different memory technologies. **LPX rack**: 256 LPUs, 128 GB aggregate SRAM, claimed **35× higher throughput per megawatt** vs Blackwell-only systems; combined system targets 1,500 tokens/sec for agentic workloads. **Shipping: Q3 2026**, initially to model builders and frontier labs.
- Strategic reading: NVIDIA's "Attention-FFN Disaggregation" goes further — exchanging intermediate activations per token loop, GPUs execute attention while LPUs handle FFN layers (~80% of compute in decoder-only models). Whether SRAM-based decode beats HBM-GPU fleets on $/token at realistic batch sizes remained unproven outside NVIDIA's marketing figures as of 2026-09-22.

### The transfer layer as a first-class decision (Wave 2, §2.4)

- Three 2026 sources converged: the KV *transfer* API is a first-class architectural decision, separate from engine and storage backend. **NIXL** appears as an LMCache backend, in the llm-d p2p guide, and as the transport under Mooncake-style stores — abstracting RDMA/IB/GPUDirect behind one API so the same disaggregated stack can target TCP in dev and RDMA in production.
- Layered decision table: co-located prefill/decode → no transfer, keep KV local; cross-node disaggregation → NIXL/RDMA (llm-d: pull beats push at every length, −55.8%→−88.2%); cross-region or cold tiers → CacheGen-compressed bitstreams over plain object storage. Tsinghua's 17–22% (TCP) / 26–33% (RDMA) Mooncake-Store-vs-Redis numbers quantify the transport line item inside an otherwise identical stack.

### The 2026 inference commercialization storyline (Wave 3)

- **2026-01-22**: Inferact launched with a **$150M seed at an $800M valuation** (TechCrunch) to commercialize vLLM — co-led by **Andreessen Horowitz and Lightspeed**, formed by vLLM creators with **Simon Mo as CEO**. An unusually large seed, consistently reported (Pulse2, The AI Insider).
- The same TechCrunch piece first surfaced SGLang commercialization talks "seeking a $400M valuation"; formal launch of **RadixArk came 2026-05-05**: **$100M seed led by Accel, co-led by Spark Capital, $400M post-money**, with NVentures, AMD and MediaTek participating (Business Wire). RadixArk expanded its Google TPU partnership (SGLang-JAX) in July 2026 — SGLang-JAX itself announced 2025-10-29 as a native JAX/XLA TPU engine; the 2026 activity is the expansion, not the birth.
- A Futurum Group analysis (2026-08-29) called vLLM the **"de facto open-source LLM inference engine"**; the PyTorch Conference NA 2026 program (2026-10-20–21, post-cutoff) featured vLLM across KV cache, disaggregated serving, hardware portability, MoE inference and production tracks with contributors from Red Hat, IBM, NVIDIA, Mistral AI, Amazon, Huawei, Meta and Google.

### The multimodal wrinkle: EPD and encoder-cache connectors

- SGLang's three-stage **EPD (Encode-Prefill-Decode)** generalizes disaggregation beyond text: vision/audio encoding is split out as its own stage because multimodal encoders are neither prefill-matmul-bound nor decode-bandwidth-bound — they are a third workload shape. vLLM v0.26.0's **encoder-cache connectors including CPU offload** are the parallel answer in the vLLM stack.
- This matters for KV economics because multimodal prompts (HybridKV's domain: each visual input expands into thousands of tokens — see part 07a) are the fastest-growing KV consumers; the disaggregation pattern is expanding to cover them rather than forcing them into P/D-shaped pools.

### Consumer tier: disaggregation reaches local hardware

- KTransformers (chunked prefill, CPU-resident sparse attention: 128K–1M tokens on consumer hardware) plus the turboquant-mlx community path (e.g. Qwen 27B at 128K on M4 Pro 48GB, part 07a) and FoveatedKV's independent M3 Max benchmark (75% KV cut, ≤3% quality loss) together mark the 2026 shift: the KV stack no longer assumes H100-class hardware — tiered offload, importance-adaptive mixed precision, and online quantization compose on consumer machines.
- The pull-vs-recompute crossover logic (LMCache `min_retrieve_tokens`, llm-d `minCachedTokenDelta: 2048`) is the same whether the "remote tier" is another node in a cluster or the host DRAM of a single machine.

### FP8 KV on hybrid models: operator detail (Wave 2, §3.2)

- A Sep 2026 deployment analysis of a hybrid (Qwen3.8, Gated-DeltaNet + attention) under vLLM documents the hybrid-specific details: FP8 KV roughly **doubles KV capacity** (e.g. toward 512K `max_model_len` with YaRN scaling); `fp8_e4m3` vs `fp8_e5m2` is a precision-vs-dynamic-range choice worth trying when subtle visual/textual separators degrade; **calibrated scales** (e.g. via llm-compressor) beat naive per-tensor scales.
- vLLM's **`--kv-cache-dtype-skip-layers`** lets operators keep sensitive layer types (e.g. sliding-window) at bf16 while quantizing the rest; **`--mamba-cache-dtype` accepts only auto/bfloat16/float16/float32 (no FP8 option)** — but the GDN recurrent state is **~0.02 GiB**, so the limitation is a non-issue, not a blocker.

### Mooncake's 2026 lineage: open questions (Wave 2 flag, carried)

- The Mooncake production figures (+115%/+107%, 50–525%, +75% real-trace) are from the paper era (arXiv:2407.00079). The 2026 open-source state lives in Mooncake Store / LMCache backends and the Tsinghua ToS 2025 evaluation. Whether Moonshot still runs the original CPP-based Mooncake on Kimi in 2026 (vs a successor) was not confirmed in any consulted source — the numbers are the disaggregation blueprint reference, not a verified 2026 production claim.
- Similarly, which 2026 production stacks (if any) deploy GEAR-style low-rank+sparse residual correction versus simpler per-channel/per-token schemes was not established — see part 07a.

