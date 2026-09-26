---
id: ai-industry-kb-2026/06-inference-engines/key-dated-facts-continued
title: "Key dated facts (continued)"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "AWS", "Alibaba", "Google", "OpenAI", "SGLang", "TensorRT-LLM", "Unsloth", "vLLM"]
dates: ["2025-12-15", "2026-05-05", "2026-08", "2026-09", "2026-09-17"]
keywords: ["amd", "aws", "benchmark", "benchmarks", "decode", "diffusion", "disaggregated", "disaggregated serving", "exploit", "fp4", "fp8", "funding"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2598, 2649]
section: "6. Inference Engines"
sha256: 4e6fb47a570c35ff02c5e75c405cc9d057e897a3e5cbb855a560ba0adb0ddd47
---

# Key dated facts (continued)

- AWS EKS starter kit vLLM doc with Neuron/Inferentia2/Trainium support: https://github.com/aws-samples/sample-genai-on-eks-starter-kit/blob/HEAD/docs/components/llm-model/vllm.md
- vLLM tpu-inference plugin README (JAX/XLA backend, Google): https://github.com/vllm-project/tpu-inference/blob/HEAD/README.md
- vLLM Neuron parity / TPU support-matrix analysis (pave planning records): https://github.com/jinhuang12/pave/blob/HEAD/planning-records/vllm-neuron-parity/exploration/upstream-and-tpu-sources.md
- AccelMark vLLM-TPU runner notes (offline/batch strengths, async limits): https://github.com/freedomintelligence/accelmark/blob/HEAD/runners/google_vllm_tpu_68cc9ffa/README.md
- Gemma 3 on vLLM-TPU GKE autoscaling guide: https://xprilion.com/gemma3-vllm-tpu-gke-autoscaling/
- RadixArk launch, $100M seed led by Accel, co-led by Spark Capital (Business Wire, 2026-05-05) — SGLang commercial context: https://www.businesswire.com/news/home/20260505077157/en/RadixArk-Launches-with-%24100-Million-in-Seed-Funding-Led-by-Accel-to-Grow-SGLang-and-Democratize-Frontier-AI-Infrastructure
- TensorRT-LLM vs vLLM measured (August 2026, Qwen3-Coder-30B-A3B, marianvid): https://github.com/marianvid/ai-lab-benchmarks/blob/HEAD/docs/engines-2026-08.md
- TensorRT-LLM release notes (B300/GB300, disaggregated benches, KV Cache Connector): https://nvidia.github.io/TensorRT-LLM/release-notes.html
- TensorRT-LLM v1.3.0rc21 release (AutoDeploy deprecation, GB300 MoE): https://media.patentllm.org/news/hardware/tensorrt-llm-v1-3-0rc21-released-nvidia-gb300-achieves-moe-w-20260728
- TensorRT-LLM v1.3.0rc23 release notes: https://media.patentllm.org/news/hardware/tensorrt-llm-v1-3-0rc23-released-amd-mi450-nvidia-rtx-5090-o-20260731
- TensorRT-LLM runtime landscape 2026-08 (rc22/23, day-0 GPT-OSS-120B/20B): https://github.com/dcharlot-physicalai-bmi/ferric/blob/HEAD/docs/runtime-landscape-2026-08.md
- TensorRT-LLM changelog (KV reuse v2, FP4/FP8 decode kernels, disaggregated serving): https://data.safetycli.com/packages/pypi/tensorrt-llm/changelog?page=2
- TensorRT-LLM production stack reference (Client → Triton → tensorrtllm_backend): https://github.com/cyyeh/skills-playground/blob/HEAD/examples/system-explorer/nvidia-tensorrt-llm/07-ecosystem.md
- llama.cpp v0.4.0 release (server changes, ggml v0.23.0): https://github.com/ggml-org/llama.cpp/releases/tag/v0.4.0
- llama.cpp model router (multi-process, LRU eviction, 2025-12-15): https://Fwnbc.marketminute.com/article/tokenring-2025-12-15-llamacpp-unveils-revolutionary-model-router-a-leap-forward-for-local-llm-management
- Unsloth official changelog 2026-09-17 (Docker + MultiUser + AMD): https://unsloth.ai/docs/new/changelog
- Unsloth v0.1.808-beta release notes (diffusion 1.2–1.7x, Vulkan AMD, PyTorch 2.11): https://github.com/unslothai/unsloth/releases/tag/v0.1.808-beta
- Unsloth v0.1.802-beta release notes (Dynamic v3.0, AMD fixes): https://github.com/unslothai/unsloth/releases/tag/v0.1.802-beta
- Unsloth PyPI page (2x faster / 70% less VRAM, unsloth start): https://pypi.org/project/unsloth/2026.8.21/
- Unsloth Dynamic v3.0 GGUF docs (UD per-layer scheme): https://unsloth.ai/docs/basics/dynamic-3.0-ggufs
- Unsloth Dynamic v2.0 GGUF docs (superseded generation): https://unsloth.ai/docs/basics/unsloth-dynamic-2.0-ggufs
- Unsloth UD quant mechanics explainer: https://github.com/orlandoluque/ai_assistant/blob/HEAD/docs/LOCAL_MODELS.md
- Unsloth UD quant benchmark notes (Q4_K_XL within 0.8 pt of original): https://github.com/allanschramm/local-model-autotuning/blob/HEAD/docs/discovery/unsloth-qwen-guides.md
- vLLM vs SGLang vs LMDeploy 2026 (engine comparison): https://dev.to/jaipalsingh/vllm-vs-sglang-vs-lmdeploy-fastest-llm-inference-engine-in-2026-5h04
- 10 best vLLM alternatives 2026 (v0.15.1, 400k+ GPUs claim, TGI maintenance): https://dev.to/jaipalsingh/10-best-vllm-alternatives-for-llm-in-production-2026-530k
- SGLang vs vLLM 2026 comparison (benchmarks, 400k+ GPUs claim): https://www.aimadetools.com/blog/sglang-vs-vllm/
- Arcfra Neutree 1.2 / Flex Engine (unified vLLM+SGLang gateway, September 2026): https://www.prnewswire.com/apac/news-releases/arcfra-releases-neutree-1-2-to-add-non-llm-model-support-and-simplify-ai-operations-302881390.html
- Tiny-vLLM launch (lightweight C++/CUDA engine, 2026): https://dev.to/eli_9c82b7dfe52c1bc371ffe/developers-launch-tiny-vllm-a-lightweight-engine-for-fast-ai-model-inference-1eim

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

