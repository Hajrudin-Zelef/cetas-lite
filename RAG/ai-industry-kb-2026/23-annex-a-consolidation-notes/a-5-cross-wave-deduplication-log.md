---
id: ai-industry-kb-2026/23-annex-a-consolidation-notes/a-5-cross-wave-deduplication-log
title: "A.5 Cross-wave deduplication log"
domain: appendix
role: appendix
task: reference
actors: ["Alibaba", "Anthropic", "ByteDance", "China", "DeepSeek", "EU", "Meta", "MiniMax", "Mistral", "Moonshot", "OpenAI", "SGLang", "TensorRT-LLM", "United States", "Unsloth", "Z.ai", "vLLM", "xAI"]
dates: ["2025-04-05", "2026-05", "2026-07", "2026-08-19"]
keywords: ["apache", "attention", "awq", "benchmark", "benchmarks", "bitnet", "chatgpt", "cost", "deepseek", "fable 5", "fp4", "fp8"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [11262, 11331]
section: "Annex A — Consolidation notes"
sha256: 39461648944b63d9238c31b488f3b04eae80c7048320cbe92b716e2905653281
---

# A.5 Cross-wave deduplication log

48. **Qwen3.8-Max weights filed under Apache-2.0/MIT** — "open-weight with custom restrictive license" (qwen3.8-max), alongside Llama Community and GLM-5.3 License — never under Apache-2.0/MIT; the final license index needs the three-way taxonomy (permissive / custom-restrictive / closed-API) plus the MiniMax **territorially-restricted community license** category (H3 excludes US/EU/UK/South Korea local deployment) — §4, §11.
49. **"95B active" as an official Qwen figure** — [third-party-reported] everywhere, never officially disclosed — §4.
50. **"Sora shut down" as full shutdown** — standalone product + public API only; model survives in ChatGPT paid tiers — §19.

---

## A.5 Cross-wave deduplication log

Content appearing in multiple waves was merged into a single canonical home; other sections carry a one-line cross-reference only. (Direction: canonical home ← sections that cross-ref.)

### §1 — Frontier closed models
- Fable 5 / Mythos 5 export-control suspension (June 12 → July 1): full incident record in **§17**; one-line cross-ref in §1.
- Company funding rounds: canonical in **§20**; cross-ref in §1.
- Nano Banana 2 Lite: deep dive in **§12**; one-line record in §1.
- Mistral Small 4 / Medium 3.5 / Large 3 details: open-model sections; recorded in §1 only to debunk the April 11 fiction.
- Gemma 4 12B specs: open-weight sections; §1 carries only the dating correction.
- ByteDance Seed models: canonical in **§4**; §1 cross-ref only.
- Qwen3.5: canonical in the open-weight chronology (**§2**); §1 drops it entirely (only the mid-February pre-release reporting remains).
- DeepSeek V4: canonical in the DeepSeek chronology (**§5**); §1 drops it.
- Muse Spark 1.3 (Sept 2) and Grok 4.7: canonical in the Wave 2.1 deltas; §1 includes them only for timeline completeness.
- Llama 4 (April 5, 2025) and Behemoth: **explicitly excluded** — 2025 events, outside §1 scope.
- Galaxy S26: excluded — hardware, not a frontier model release.

### §2 — Open-weight chronology
- Architecture deep dives: GLM → §3; Qwen Gated DeltaNet / KDA → §4; DeepSeek DSA/CED/V4.1 → §5. This section's Spec table is the canonical per-model 1-line spec record — other sections cross-reference, never re-specify.
- Inference engines → §6; quantization (GGUF, NVFP4, MXFP4, Unsloth v3.0 mechanics) → §8; full pricing tables → §14; frontier closed-model details → §1.
- Dated release rows live in §2 (1–2 lines each); §5's DeepSeek material collapses to 1–2-line rows + cross-ref to §5.
- §4's Qwen "Key dated facts" bullets collapse to one-line §2 rows with a "→ §4" pointer; the spec table, benchmark reality checks, license map, pricing chronology, and Omni strategy deep dive stay in §4.

### §3 — Chinese GLM/Kimi/MiniMax (deep file)
- GLM-5.2 architecture detail and IndexShare mechanics beyond the parameter reconciliation: canonical in wave3/02-openweight-chronology.md (Claim 3).
- GLM-5.2 benchmarks and text-only status in the frontier comparison: wave1/07 §2.4 and wave3/06 §1.3.
- GLM-5.3 release and safety-hold narrative: wave1/07 §2.4 and wave2.1/05 §4 — this section adds only the full benchmark verification, the always-on-reasoning breaking change, and the rumor-phase dating.
- GLM-5.3-Flash / "Ox Alpha": wave1/07 §2.4 — nothing new here.
- Kimi K3 chronology, launch, and "beats Opus 4.8" evidence: wave3/02 Claim 4 — this section keeps only technical-report architecture detail plus flagged conflicts.
- MiniMax H3 model detail (H3-Omni-Transformer, 2K upscaling, H3-Context-IR, open-weights Aug 2–3): wave2.1/07 §1.1 — this section records only the alias resolution, the territorial license restriction, and the 2.7T rumor.
- Seedance 2.0 / 2.5: wave3/06 §1.3 — cross-reference only.
- OpenVuln / VulnHunter narrative: wave1/07 §2.4 — figures confirmed here, not duplicated.

### §4 — Qwen / Seed / Ling (deep file)
- Feb→May 2026 Qwen skeleton: wave3/02 (not re-derived).
- Gated-DeltaNet hybrid mapping, Qwen3.8-Max-0902 detail, Qwen3.8-27B detail: wave2.1/05 (cross-ref only).
- Seedance 2.5: wave3/06 §1.3 (cross-ref only).
- Artifact/service split enforced: Qwen3.8-Max (hosted) vs Qwen3.8-2.4T-A95B (weights) vs Qwen3.8-Max-0902 (checkpoint) stay separate index entries in every index.

### §5 — DeepSeek (deep file)
- Dated release rows: canonical in §2 (this section keeps the full narrative).
- Serving/pricing economics: summarized here; full treatment in §9 / §14 of the final document.
- V4 fine-grained-MoE architecture reconstruction (256/384 routed experts, 6 active, hash-routed first blocks, sqrt(softplus) affinity, FP4 expert weights): canonical in wave2/03 §4 — cross-referenced, all figures [PARTIALLY VERIFIED].

### §6 — Inference engines (two parts, 06a vLLM / 06b SGLang)
- SGLang internals and the vLLM-vs-SGLang shootout: canonical in **06b** (06a keeps minimal comparative framing: the +29% figure cited once, per-model engine choice).
- KV-cache mechanics, FP8 KV dtype, O(n)/O(n²) analysis: canonical in **§7** (06b keeps only engine-level outcome figures: hit-rate ratios, throughput deltas, release versions).
- Quantization dtype mechanics (FP8/MXFP4/TurboQuant/UD quants): canonical in **§8**.
- GPU pricing and cloud inference markets: canonical in **§14**.
- TPU hardware specifics: **§15** (SGLang-JAX recorded in 06b only as the engine-backend fact).
- vLLM internals, vLLM-Omni, vllm-gguf-plugin, Inferact, NIM 2.0, Dynamo vs llm-d, spend flip, per-token economics: canonical in **06a**.
- SGLang (RadixAttention, scheduling, releases, commercialization, claims) + adjacent runtimes (TensorRT-LLM, llama.cpp server, emerging engines): canonical in **06b**.
- Wave-1 archive flags (MiMo-V2.6 spec discrepancies, 1M-context claim harmonization): explicitly out of §6's scope — harmonized centrally.

### §7 — KV cache & attention (two parts, 07a mechanisms / 07b disaggregation)
- **07a** is the home for: KV-cache sizing law; attention variants (MHA/MQA/GQA/MLA) + MLA engine-support who/when; shared/hybrid layouts (Gemma 4, MSA, IndexShare, DSA); post-hoc compression zoo; error-correction canon (GEAR, MiKV, ResQ, WKVQuant, ZipCache/PrefixQuant/MiniKV, LESS, KVLinC, Kitty, FoveatedKV); quantized-KV production dtypes; Mamba-2/Gated-DeltaNet; sizing ladders. Engine per-release kernel detail lives in **§6**.
- **07b** is the home for: disaggregation, tiered storage, RDMA, prefix caching, 1M-context product status; attention-variant mechanisms appear only as cross-refs to 07a; vLLM/SGLang version specifics as timeline anchors only (detail in **§6**).

### §8 — Quantization (two parts, 08a formats/tooling / 08b precision/production)
- **08a** owns: formats and tooling (GGUF, K-quants/IQ-quants, Dynamic 3.0, EXL2→EXL3, AWQ/GPTQ toolchain turnover, deployment decision matrix, llama.cpp, edge/ARM variants, NF4/QTIP, workflows). The NVFP4≠MXFP4 anatomy table stays in 08a (format-definition fact); the vLLM-vs-EXL3 shootout stays in 08a (adjudicates a format/engine claim, [COMMUNITY]).
- **08b** owns: the FP8/FP4/NVFP4/MXFP4 **production arc**, TurboQuant production verdicts, KV-cache quantization methods, BitNet, Bonsai-27B, cost arithmetic, quality measurement.
- TurboQuant method detail: canonical in **§7** (08b carries a one-line cross-ref).
- Dynamic UD format mechanics: canonical in **§8** (§10 records only versioning facts: 2.0 current July 2026, superseded by v3.0 2026-08-19).

