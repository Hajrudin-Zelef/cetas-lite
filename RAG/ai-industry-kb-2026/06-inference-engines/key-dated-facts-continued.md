---
id: ai-industry-kb-2026/06-inference-engines/key-dated-facts-continued
title: "Key dated facts (continued)"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "AWS", "Alibaba", "Apple", "Baseten", "DeepSeek", "Google", "Huawei", "Intel", "LongCat", "Microsoft", "MiniMax", "Moonshot", "Nebius", "Nvidia", "OpenAI", "Oracle", "SGLang", "TensorRT-LLM", "Z.ai", "vLLM", "xAI"]
dates: ["2025-03", "2026-03-15", "2026-05-04", "2026-06", "2026-08-20", "2026-09", "2026-09-04", "2026-09-22"]
keywords: ["agent", "agentic", "agents", "alignment", "amd", "ascend", "attention", "aws", "blackwell", "compute", "consumer", "datacenter"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2627, 2706]
section: "6. Inference Engines"
sha256: 75dd10b9db0c779eca851de27a791d4fb5a6f9a0a146a4a70f423582946027a9
---

# Key dated facts (continued)

## Key dated facts (continued)

### RadixAttention — the core mechanism (verified)

- SGLang's `RadixAttention` keeps a radix tree (compressed trie) of all active KV-cache entries across requests, indexed by layer (`req_to_token_pool.get_kv_buffer(layer_id)`, `cache_loc`).
- Token sequences are stored as variable-length edges; each node points to paged KV blocks in GPU memory, and new requests match the tree for the longest cached prefix at any token boundary — not fixed block multiples.
- Generated tokens are also inserted, so later requests referencing prior model output benefit from the same prefix tree.
- Prefix caching is enabled by default (`--disable-radix-cache` to turn it off).
- Advantage over vLLM's automatic prefix caching (APC): hash table at fixed 16-token blocks vs token-level radix tree; O(log n) tree traversal vs O(k) hash lookups per block; tree-aware LRU eviction (leaf nodes first, shared prefixes preserved) vs flat LRU.
- Worked example (RAG, 100 users, 400-token system prompt, 3×600-token docs, 60% overlap) [COMMUNITY, methodology unaudited]: no caching 225,000 prefill tokens/batch → vLLM APC 125,000 (−44%) → RadixAttention 65,000 (−71%), i.e. ~3.5x throughput, −60% TTFT, −40% KV memory.
- Cache-aware scheduling is opt-in: `--schedule-policy lpm` (longest prefix match) or `dfs-weight` (the two `CacheAwarePolicy` members) in `srt/managers/schedule_policy.py`; the default is `fcfs` (cache-agnostic), so queue reordering favoring long cached prefixes requires an explicit flag.
- `--enable-cache-report` exposes `cached_tokens` in the OpenAI `usage` field — the production way to verify real hit rates.

### Prefix discipline as an operational lever

- Production prefix hit rates run 50–99% depending on prompt discipline — the same engine doubles or halves its cache advantage on how prompts are written, making prompt design an ops lever, not just an engine property.
- Voice-cloning workloads reach 86.4% prefix hit rates; the RAG worked example shows −71% prefill tokens vs no caching only when overlap discipline holds.
- The sgl-router deployment pattern (3.8x hit rate vs round-robin) is what scales prefix affinity to fleet size — the plausible machinery behind the [UNVERIFIED] 400K-GPU-class claims, since no single engine flag explains fleet-scale prefix sharing.
- The verification loop is closed-loop: `--schedule-policy lpm` (or `dfs-weight`) reorders the queue, `--enable-cache-report` surfaces `cached_tokens` in the OpenAI `usage` field, and the operator iterates on prompt structure — cache-aware scheduling is only as good as the discipline it can exploit.
- Page-size note: the radix cache operates at page sizes 1/16/64 (token-level matching with paged backing) — deep radix-tree mechanics belong to §7; this part keeps the scheduling-relevant outcome.

### Scheduler internals — why prefix-heavy load wins (verified, Sept 2026)

- Multi-process, asynchronous runtime: TokenizerManager → Scheduler → TpModelWorker (one GPU-side executor per TP shard), IPC over ZeroMQ; request lifecycle: tokenize → `waiting_queue` → admission to `running_batch` when KV slots free → prefill (`ForwardMode.EXTEND`) → decode (`ForwardMode.DECODE`, CUDA-graph captured) → detokenize/stream.
- The core decision is `get_next_batch_to_run`: merge last prefill batch into the running batch → stash/exclude chunked requests → try `get_new_batch_prefill()`; prefill is prioritized over decode — if a new prefill batch fits the token budget it runs this step, otherwise the scheduler advances the running decode batch, returning a `NextBatchPlan` bundling both.
- `PrefillAdder` greedily packs waiting requests under a token budget with two big optimizations: prefix-cache matching (cached prefix excluded from tokens to compute, via the radix cache) and chunked prefill (long prompts split at `--chunked-prefill-size`, e.g. 8192/6144).
- Without chunked prefill, a single 32K-token prompt blocks all decode steps for its entire prefill — chunked prefill is as much a head-of-line-blocking fix as a throughput optimization.
- Zero-overhead scheduler (v0.4 blog): CPU scheduler runs one batch ahead, overlapping scheduling/prefix matching with GPU compute (`event_loop_overlap`); measured 1.1x throughput over v0.3, GPU idle time ≈ 0 under Nsight. [VENDOR]
- Cache-aware load balancer (v0.4 blog): routes requests to the DP worker with the best predicted prefix hit rate — up to 1.9x throughput, 3.8x hit rate. [VENDOR]
- Parallelism: TP, pipeline, expert, data parallelism; DP attention specifically for DeepSeek-style MLA models (up to 1.9x decode throughput [VENDOR]); MoE expert parallelism with all-to-all backends (DeepEP, Mooncake RDMA, NIXL-EP, MORI on AMD, FlashInfer), TBO/SBO overlap, EPLB load balancing.
- Prefill/decode disaggregation via separate worker pools over Mooncake/NIXL transfer engines + router/gateway.
- Decode side: FlashInfer BatchDecode kernels, continuous batching, KV-cache quantization (FP8), speculative decoding (multiple tokens per KV load).
- Hardware (September 2026): NVIDIA (GB200/B300/H100/A100/5090), AMD (MI355/MI300/MI325X), Intel CPU/XPU, Google TPU, Ascend NPU, Apple Silicon (SGLang Diffusion via MPS), Moore Threads (MTT S5000).

### Jump-forward constrained decoding and XGrammar (verified)

- SGLang compiles schemas to a finite-state automaton in the engine core with GPU-batched FSA evaluation, then jump-forward decoding emits deterministic token spans without model forward passes — "Model Forward only for needed tokens."
- Backends: XGrammar (default, ~10x faster than Outlines for structured output per community docs), plus Outlines and llguidance — vs post-hoc logit masking (vLLM + Outlines): per-token CPU grammar evaluation (1–5 ms overhead), no cross-request batching.
- The frontend exposes typed constrained generation (`select`, `choices`): for classification-style prompts `select` is faster and strictly correct — output cannot be off-list; structured-output decoding reported ~3x faster than logit-masking via the compressed FSM.
- Why it matters for agents: tool-call schemas, function-calling JSON, and extraction become near-free on grammar-constrained spans while forward passes are reserved for genuinely uncertain tokens — a throughput lever compounding with RadixAttention on agentic loops.
- XGrammar is the 2026 universal substrate: default in vLLM (v0.7+), SGLang (v0.4+), TensorRT-LLM (since 2025-01), MLC-LLM, OpenVINO GenAI, and Modular MAX — SGLang's remaining edge is jump-forward decoding, which skips forward passes entirely on deterministic spans (40–60% fewer LLM calls for known JSON schemas), orthogonal to XGrammar's mask efficiency; combined overhead <3% of unconstrained throughput.
- XGrammar-2 (released 2026-05-04, MLC blog): Structural Tag, a composable JSON protocol uniformly expressing OpenAI harmony format, tool calling, reasoning channels, and custom output structures; cross-grammar caching, repetition-state compression, batching + speculative-decoding support; adopted by xAI, Databricks, DeepSeek; integrated by SGLang, vLLM, TensorRT-LLM, MLC-LLM for strict tool calling.
- Claim check: "sub-40-microsecond per-token overhead, up to 80x throughput improvement over older approaches" is vendor-adjacent [VENDOR] — treat as directional; the better-attested figure is ~3x compressed-FSM speedup over standard guided decoding, with JSON compliance rising from 90–94% (unconstrained) to 96–98.2% (constrained).
- Syntax convergence: XGrammar adopted GBNF as its grammar syntax, making GBNF the de facto shared grammar language across vLLM, SGLang, llama.cpp-family, and llguidance (converters both ways).

### The agentic/frontier niche (2026)

- Named flagship adopters in 2026 surveys: xAI (Grok 3 served on SGLang), Microsoft Azure (DeepSeek R1 endpoints on AMD), Oracle Cloud, Google Cloud, AWS, Nebius, DataCrunch, Voltage Park, Cursor (code completion), LinkedIn (AI features), NVIDIA, AMD, Intel, Baseten, RunPod, Novita; academia (Stanford, UC Berkeley, UCLA, MIT, UW, Tsinghua).
- RL/post-training moat: SGLang is the de facto serving backbone of the 2026 RL stack — verl, AReaL, Miles, slime, Tunix — because GRPO rollouts need exactly high-throughput, prefix-heavy serving; post-training researchers default to SGLang, optimizations (MLA, DP attention, kernels) land there first, frontier open-weight models serve best on it, and the loop reinforces itself.
- The workload split (2026 consensus): SGLang for prefix-heavy production — agents, RAG, multi-turn chat, RL rollouts; vLLM for high-concurrency general API serving and predictable latency (see sibling part 06a).
- Head-to-head (AIMultiple Research, 1,000 ShareGPT prompts, H100, bfloat16): SGLang 16,215 tok/s vs vLLM 12,553 tok/s (+29%) — critically, vLLM was tested with the FlashInfer backend (same kernels as SGLang), so the gap is orchestration overhead, not kernel performance; sgl-router: 3.8x cache hit rate vs round-robin.
- Counter-evidence (same survey): vLLM wins fastest TTFT across concurrency levels in GPT-OSS-120B tests and leads at 100+ concurrent requests, with broader non-NVIDIA hardware support and testing — engines trade wins by workload, not by quality.
- Enterprise comparison (privocto, 2026-03-15) [DIRECTIONAL, single secondary]: SGLang "dominates structured generation" with 4.7x speedup in multi-turn conversations, 99.8% JSON validity, 47% less GPU memory (40GB vs 75GB) in agent workflows while maintaining superior throughput.
- Per-model engine choice (2026-09): one production homelab operator parked SGLang for DeepSeek V4 Flash in June 2026 (SGLang's strict-config loader rejects the NVFP4 quant) and serves it via vLLM instead, keeping SGLang for other models — engine choice in 2026 is per-model, not per-religion.
- Community scale: ~25K–30K GitHub stars, ~600 contributors (ranges across sources); funded by a16z's OSS AI Grant; in the PyTorch ecosystem since March 2025; origin LMSYS/UC Berkeley (Chatbot Arena team).
- Enterprise gap (as of September 2026): no Red Hat- or NIM-equivalent SGLang distribution surfaced — SGLang's enterprise route is self-supported or via cloud partners (Azure productizes it inside first-party AMD endpoints); known research issues: tool calling less stable than vLLM's, LoRA hot-swap quirks, thinner docs.
- OME (`sgl-project/ome`): SGLang's own Kubernetes operator for model management, lifecycle, and multi-engine support — the project's answer to the K8s deployment question.

### SGLang v0.5.19 — the September 2026 release

- Released on PyPI 2026-09-04 (GitHub tag v0.5.19): 786 PRs from 214 contributors; the v0.5.18 tag (2026-08-20) remains what official install docs point to for source builds — v0.5.19 is fresh as of the 2026-09-22 research cutoff.
- New models: Qwen3.8 (2.4T-A95B), Qwen3.8-27B, dots3.note, Ling-3.0-flash, Ling-3.0-tiny, Spark2.5, MiniCPM-SALA, Granite 4.2, plus LongCat-Image-Edit & Edit-Turbo (diffusion).
- Cookbook updates: GLM-5.3 deployment guide; PaddleOCR-VL deployment guide; Kimi-K3 on Ascend A3; Kimi-K2.7-Code-MXFP4 on MI355X; Qwen3.5 MXFP4 on MI355X with FP8 KV cache or HiCache host-memory tier; MiniMax-H3 on 24GB GPU or DGX Spark with a consumer-GPU tuning guide; Ling-3.0-flash on DGX Spark.
- The same model set (Qwen3.8, Kimi-K3) appears in both vLLM v0.28.0 and SGLang v0.5.19 release notes — both engines are racing to serve the same late-2026 flagships across NVIDIA, AMD, Ascend, and DGX Spark; engine differentiation is increasingly measured in time-to-serve for new models.
- Related: the sglang-omni line is at v0.1.4 (config refactor, AMD ROCm Qwen3-TTS/ASR on gfx950, SGLang 0.5.18 dependency alignment) — the any-to-any serving track, separate from core 0.5.19.
- MLA-2026 arc: zero-overhead scheduler and cache-aware load balancer (v0.4 blog); FlashInfer SM120 sparse-MLA decode kernels for DeepSeek-V4 (PR #27455, ~June 2026) delivering 2.2–3.7x TPOT reduction vs the Triton kernel; FP4 MLA KV caches in v0.5.6 (Dec 2025); `--attention-backend=dsv4` for hybrid CSA+HCA attention on DeepSeek V4 Flash.
- Hardware reality check: MLA kernels are SM90+-first (Hopper); SM120 (Blackwell workstation) support arrived through 2026 via FlashInfer with community patching (NaN bugs in specific batch ranges on SM121 bisected and fixed in nightlies) — mature on datacenter Hopper/Blackwell, still being tuned on workstation silicon.

### SGLang Diffusion — image/video generation (2026 expansion)

- The `sglang-diffusion` sub-project (`python/sglang/multimodal_gen/`) reuses SGLang's serving muscle (sgl-kernel ops, scheduler loop, CUDA-graph capture, quantization, multi-hardware, OpenAI-compatible API, PD-style disaggregation) for iterative denoising; installable via `uv pip install "sglang[diffusion]"`.
- Diffusion-specific optimizations: TeaCache (skip redundant denoise steps) and Cache-DiT integration — up to 7.4x inference speedup with minimal quality loss via `SGLANG_CACHE_DIT_ENABLED=True` [VENDOR integration claim].
- Model support: Wan series, FastWan, Hunyuan, Qwen-Image / Qwen-Image-Edit, Flux, Z-Image, GLM-Image, Ideogram 4, Krea-2, Cosmos3, LTX-2 / LTX-2.3, MiniMax-H3, MOVA, Hunyuan3D, and more.
- Hardware (verified from the runtime README): NVIDIA, AMD, Intel XPU, Ascend NPU, Apple Silicon (MPS), Moore Threads (MTT S5000) — the broadest hardware list of any 2026 diffusion-serving project.
- Interfaces: `sglang generate`, `sglang serve`, OpenAI-compatible `/v1/videos`, CLI, and Python SDK.
- Production-recipes signal: an Ant Group cookbook for serving MOVA under SGLang Diffusion — the diffusion-serving path moving from demos to documented production recipes.
- Positioning note: this expansion runs parallel to vLLM's any-to-any multimodal track (see sibling part 06a) — orthogonal to the text-serving battle.

- Positioning note: this expansion runs parallel to vLLM's any-to-any multimodal track (see sibling part 06a) — orthogonal to the text-serving battle.

