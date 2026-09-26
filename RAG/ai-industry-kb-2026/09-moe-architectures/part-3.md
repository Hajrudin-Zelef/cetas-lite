---
id: ai-industry-kb-2026/09-moe-architectures/part-3
title: "9. MoE Architectures (part 3)"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Anthropic", "DeepSeek", "Google", "Huawei", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Together AI", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-04-24", "2026-06-01", "2026-06-13", "2026-07-16", "2026-07-24", "2026-07-27", "2026-08-13"]
keywords: ["moe", "agentic", "amd", "apache", "ascend", "attention", "benchmarks", "blackwell", "claude", "compute", "cost", "decode"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4669, 4701]
section: "9. MoE Architectures"
sha256: 609ef7fb9308979b04c1574b37a3c490b8d13e5900689f97e71ec0c2822e3f4f
---

# 9. MoE Architectures (part 3)

- **Specs (verified via HF model card, Together AI, community analyses):** 35B total / ~3B active per token; 40 layers; hidden dim 2048; **256 experts with 8 routed + 1 shared active per token**; expert intermediate dim 512.
- Hybrid attention: 10 repetitions of [3× (Gated DeltaNet → MoE) + 1× (Gated Attention → MoE)] — 30 Gated DeltaNet linear-attention layers + 10 gated attention layers (16 Q heads, 2 KV heads, head dim 256). DeltaNet uses 32 linear attention heads for V and 16 for QK (head dim 128).
- Trained with multi-token prediction (MTP); 262,144 tokens native context, extensible to ~1,010,000 via YaRN RoPE scaling; vision encoder built in (image-text-to-text); Apache 2.0.
- Local inference [COMMUNITY]: Unsloth UD-Q4_K_XL GGUF ~21 GB runs on a single 24 GB GPU (RTX 4090); 20.9 GB Q4 quant runs on MacBook Pro; vLLM ≥ 0.19.0 required (`Qwen3MoeSparseMoeBlock` support); practitioner measurement 61 tok/s on RTX 5070 vs 7 tok/s for a 27B dense model on the same hardware — the canonical active-compute illustration (3B active vs 27B dense, 9× less per-token compute).

#### Google Gemma 4 26B-A4B — 25.2B/3.8B pure-attention MoE (released ~2026-04)

- **Specs (verified via HF model card, community architecture breakdowns):** 25.2B total / 3.8B active per token; 30 layers; hidden dim 2560; **128 experts + 1 shared, 8 routed + 1 shared = 9 active per token**; 32 Q heads, 8 KV heads (GQA); 256K native context; 262K-token vocabulary; ~550M vision encoder (separate, multimodal).
- Attention design: hybrid local/global — sliding-window attention (1024-token window) alternating with global attention; local layers use 8 KV heads (head dim 256), global layers 2 KV heads (head dim 512). Pure-attention MoE with no SSM/recurrence (unlike Qwen3.6's DeltaNet), making the full forward pass parallelizable — prefill scales to full GPU TFLOPS with no sequential bottleneck.
- Expert offloading case study [COMMUNITY]: only ~4.2–4.8 GB must stay hot in VRAM (embeddings ~1.3 GB, per-layer attention+router ~2.4 GB, shared expert ~50 MB); 128 routed experts (~11 GB in Q4_K_M) can sit on NVMe with ~24 MB loaded per forward pass (8 experts × ~3 MB) — servable on an 8 GB RTX 3070 with flash offloading.
- Family context: Gemma 4 ships four sizes — E2B (dense, 2.3B effective), E4B (dense, 4.5B effective), 26B-A4B (MoE), 31B (**dense**, 30.7B) — all Apache 2.0, a licensing shift from earlier Gemma-specific terms. The dense 31B sibling is itself the counterexample to "MoE is the default everywhere" claims.

#### DeepSeek V4 — Pro 1.6T/49B + Flash 284B/13B (released 2026-04-24)

- **Specs (verified via official announcement, model cards, API docs):** V4-Pro: 1.6T total / 49B activated, pre-trained on 33T tokens; V4-Flash: 284B total / 13B activated, 32T tokens. Both: 1M-token native context, up to 384K output tokens, dual thinking/non-thinking modes (high/max/non-think effort levels), text-only.
- Previewed 2026-04-24 (same day as OpenAI GPT-5.5); V4-Pro reached GA 2026-08-13 (V4-Pro-0813 checkpoint); legacy `deepseek-chat`/`deepseek-reasoner` aliases retired 2026-07-24, routing to V4-Flash during grace period.
- Notable industry milestone: launched on **Huawei Ascend chips first** (no NVIDIA required). License: MIT in later sources; early April coverage cited Apache 2.0 — unresolved discrepancy.
- **Reconstructed architecture [PARTIALLY VERIFIED]** (secondary technical analyses and a mirrored snapshot of the V4 technical report; no official DeepSeek report directly verified): V4-Flash — 43 layers, 256 routed experts, 6 active per token + 1 shared; V4-Pro — 61 layers, 384 routed experts, 6 active per token + 1 shared. Every block is MoE; the first three blocks use deterministic token-ID hash routing instead of learned routing. Auxiliary-loss-free load balancing retained from the V3 lineage with a small additional sequence-wise balance loss. Routing affinity changed from sigmoid to `sqrt(softplus(...))`. Routed expert weights reportedly deployed in FP4. V4 replaces the earlier MLA design with CSA/HCA hybrid attention (attention-side detail consolidated in part 09b).

#### MiniMax M3 — 428B/23B natively multimodal MoE (released 2026-06-01)

- **Specs (verified via HF model card, SGLang cookbook, Artificial Analysis):** ~428B total / ~23B activated per token; 60 layers; hidden size 6144; 64 attention heads with `num_key_value_heads=4` (GQA); **128 routed experts with 4 active per token + 1 shared expert**; first 3 layers dense. 1M-token context via MiniMax Sparse Attention (MSA), a block-sparse "lightning indexer" attention (top-k 128-token blocks) keeping decode cost roughly flat in context length — MiniMax reports ~9× prefill and ~15× decode speedup over M2 at 1M context, ~1/20 the per-token compute of the previous generation [VENDOR].
- Multimodality: trained on mixed text, image, and video **from step 0** (not a bolted-on adapter); accepts interleaved text and images; positioned as the first open-weight release combining reasoning, agentic coding, and native multimodality.
- Reasoning/tool use: three reasoning modes via `thinking` param (enabled/adaptive/disabled); chain of thought wrapped in `<mm:think>...</mm:think>`; native XML-namespace tool calling parsed to OpenAI `tool_calls`; served with `--reasoning-parser auto` / `--tool-call-parser auto` in SGLang.
- Serving notes: FP8/MXFP8 and NVFP4 checkpoints exist (`MiniMaxAI/MiniMax-M3-MXFP8`, ~440 GB); served on NVIDIA Blackwell and AMD Instinct via SGLang, and on Hopper (H200) in bfloat16. Community test logs report 7 multi-token-prediction modules in the semi-analysis model card but MTP/NEXTN weights absent from shipped checkpoints (`speculative_enabled: false` mandatory) [COMMUNITY].
- License: MiniMax Community License (custom). Reported to match Claude Sonnet 4.6 on real-world agentic benchmarks via Morph [VENDOR, UNVERIFIED independently].

#### Z.ai GLM-5.2 — coding-plan flagship (released 2026-06-13/17)

- Total parameters undisclosed; the MoE base was reused in GLM-5.3 (753B); 1M context; custom license; coding-plan flagship positioning. Documented here because its base is the one disclosed in GLM-5.3 (see below) — the 5.2→5.3 lineage is the 2026 case study of post-training upgrades on an identical MoE base.

#### Moonshot Kimi K3 — 2.8T/104B frontier MoE (announced 2026-07-16, weights 2026-07-27)

