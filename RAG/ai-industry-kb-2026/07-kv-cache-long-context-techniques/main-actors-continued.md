---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/main-actors-continued
title: "Main actors (continued)"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "Anthropic", "DeepSeek", "Google", "Groq", "Intel", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Samsung", "TensorRT-LLM", "Z.ai", "vLLM"]
dates: ["2026-01-22", "2026-02-19", "2026-03-16", "2026-03-21", "2026-03-25", "2026-04-15", "2026-04-24", "2026-05", "2026-05-05", "2026-05-12", "2026-05-28", "2026-06", "2026-06-01", "2026-06-12", "2026-06-13", "2026-06-16", "2026-08-13", "2026-08-26", "2026-08-29", "2026-09-08", "2026-09-17", "2026-09-22", "2026-10-20"]
keywords: ["acquisition", "agentic", "agents", "agi", "amd", "attention", "benchmark", "benchmarks", "claude", "compute", "consumer", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3639, 3766]
section: "7. KV Cache & Long-Context Techniques"
sha256: c01e9400d892794d643955cf371b594ab2b6be329681967880518cbaa08a4b62
---

# Main actors (continued)

## Main actors (continued)

- **LMCache (UC Berkeley lineage; v0.3.15, Mar 2026)** — the engine-independent KV-cache layer: tiered GPU → CPU → disk → remote storage, NIXL-based P/D transfer, CacheBlend non-prefix reuse. The 7.43× second-run figure is the canonical "KV cache is storage tier" proof.
- **Lightbits (Arthur Rasmusson)** — RDMA-paged KV to 10M tokens (2026-09-17): the strongest measured prefill-cost-collapse evidence, plus the Open KV Cache API interop effort (Inferra reference implementation).
- **llm-d** — peer-to-peer KV sharing guide (Sep 2026): the measured pull-vs-recompute crossover on gpt-oss-120b (−55.8% → −88.2% prefill latency over RDMA/IB).
- **Mooncake / Mooncake Store (Moonshot AI lineage)** — the KVCache-centric disaggregation blueprint (+115%/+107% on Kimi) and its open-source TCP/RDMA store (17–22% / 26–33% vs Redis, Tsinghua).
- **NIXL (NVIDIA)** — the common high-performance transfer library behind LMCache backends, llm-d p2p, and Mooncake-style stores; abstracts RDMA/IB/GPUDirect.
- **CacheGen (UChicago)** — KV bitstream encoding for the network (3.5–4.3× bandwidth cut), the ancestor of transfer compression.
- **DeepSeek** — FlashMLA (Feb 2025, the FP8-KV origin), DSA (Dec 2025), V4/CSA+HCA efficiency lineage (Apr 2026): the open-weight stack whose KV techniques all others consume.
- **Google Research** — TurboQuant (ICLR 2026, blog 2026-03-25): the reference online 3-bit KV quantization; announcement moved semiconductor stocks.
- **Z.AI** — GLM-5.2 (June 2026) with IndexShare; GLM-5.3-Flash (Aug 2026) at 1M on one node; the open-weight lab pairing MLA-class efficiency with shipped 1M windows.
- **MiniMax** — M3 (June 2026): the 1M open-weight proof with a stated 512K guaranteed floor and the sparse-attention retrieval-fidelity argument.
- **NVIDIA** — Dynamo (P/D orchestration), Groq 3 LPU acquisition (~$20B, Dec 2025; unveiled GTC 2026 — hardware P/D disaggregation), Nemotron-3-Ultra hybrid, KV Cache Connector API in TensorRT-LLM.
- **Inferact ($150M seed, 2026-01-22, a16z + Lightspeed, $800M valuation)** and **RadixArk ($100M seed led by Accel, $400M post-money, formal launch 2026-05-05)** — the January–May 2026 commercialization storyline for vLLM and SGLang respectively; vLLM positioned as the "de facto open-source LLM inference engine" (Futurum Group, 2026-08-29).
- **SK Hynix (~53%), Samsung (~38%), Micron (~9%)** — the HBM oligopoly; Hynix's 72% Q1-2026 margin and reported $28B IPO filing are the pricing-power signals; AMD's MoRI roadmap targets tiered KV as an explicit platform feature for H2-2026.
- **Community/secondary**: KTransformers (consumer-tier disaggregation), onnx-light-cpu (INT8/FP8 KV on CPU roadmap), AMD-fork kernel-fusion benchmarks, turboquant-mlx integrations.

## Timeline and context (continued)

| Date | Event |
|---|---|
| 2025-12 | DeepSeek-V3.2 introduces DeepSeek Sparse Attention (learned lightning indexer, top-K 2048) |
| 2025-12 | NVIDIA acquires Groq technology/team (~$20B) — decode hardware enters NVIDIA's roadmap |
| 2026-01-22 | Inferact launches: $150M seed at $800M valuation to commercialize vLLM (a16z + Lightspeed) |
| 2026-03-21 | HF TGI archived read-only (maintenance mode) — MLA existed only on the Intel-Gaudi branch, never mainline CUDA |
| 2026 (v0.25) | PagedAttention removed from vLLM's new internal engine — legacy path only from here on |
| 2026-02 | SGLang blog: "Unlocking 25× inference performance on GB300 NVL72" [VENDOR headline; baseline unaudited] |
| 2026-02 | Qwen3.5-MoE ships: Gated-DeltaNet/Mamba + 256-expert MoE hybrid |
| 2026-02-19 | Gemini 3.1 Pro preview: 1M input / 64–66K output — 1M enters flagship-closed tier |
| 2026-03 | LMCache v0.3.15: engine-independent KV caching goes hardware-portable (H100/H200/B200/MI300X) |
| 2026-03-25 | Google Research blog announces TurboQuant (ICLR 2026): 3-bit KV, 6× memory, 8× attention speedup |
| 2026-03-16–17 | GTC 2026: Groq 3 LPU unveiled (150 TB/s SRAM decode; LPX rack 35×/MW claim [VENDOR]); Rubin CPX with GDDR7 for prefill |
| 2026-04 | vLLM FP8-KV benchmarks: 54% ITL slope vs BF16, break-even ~7K tokens, sub-0.3% accuracy |
| 2026-04 | FoveatedKV independent benchmark (M3 Max, Apr 4): importance-adaptive fp16/fp8/INT4 tiering |
| 2026-04-15 | Qwen3.6-35B-A3B (~): Gated-DeltaNet + Gated-Attention + MoE at **256K** — the 1M-tier exception |
| 2026-04-24 | DeepSeek V4 (Pro/Flash): 1M context, MIT license, CSA+HCA; SGLang day-zero support |
| 2026-05-05 | RadixArk formal launch: $100M seed (Accel-led, $400M post-money) to commercialize SGLang |
| 2026-05-28 | Claude Opus 4.8: 1M / 128K (secondary-sourced; Foundry caps at 200K on some tiers) |
| 2026-06-01 | MiniMax M3: 1M context open-weight launch (512K guaranteed floor) |
| 2026-06-12 | NVIDIA Nemotron-3-Ultra-550B-A55B: Mamba-2 hybrid MoE, MTP, 1M, NVFP4, OpenMDW-1.1 |
| 2026-06-13/17 | GLM-5.2 API release / open weights: IndexShare (2.9× FLOP cut at 1M), 1M/65K |
| 2026-07 | RadixArk expands Google TPU partnership (SGLang-JAX); SK Hynix $28B IPO filing reported |
| 2026-07 | SGLang day-zero support for Kimi K3 |
| 2026-06 | Zamba2-VL: Mamba2–Transformer hybrid VLM, ~10× TTFT cut (Jun 2026, MarkTechPost) — hybrid-side datapoint |
| 2026-08 | GPT-5.6 family: 1.05M context with pricing cut |
| 2026-08-13 | DeepSeek V4-Pro GA checkpoint (V4-Pro-0813) — secondary-sourced |
| 2026-08-26 | vLLM v0.28.0 tagged (see §6 for engine detail: sparse MLA, tiered KV offload to disk) |
| 2026-08-29 | Futurum Group: vLLM the "de facto open-source LLM inference engine" |
| 2026-09-17 | Lightbits RDMA-paged KV post: 8K→38.8 ms TTFT, 10.5M-token restore in seconds |
| 2026-09-22 | LMCache docs: 7.43× second-run CPU offload; llm-d p2p guide (−55.8%→−88.2%); this wave's FP8-KV default synthesis |
| 2026-10-20–21 | PyTorch Conference NA, San Jose (post-cutoff): vLLM across program tracks — program positioning only as of 2026-09-22 |

## Implications (continued)

1. **The KV cache is now a storage tier, not a scratch buffer.** LMCache (engine-independent, 7.43× reuse), Lightbits RDMA (10M-token bit-identical restoration), and llm-d p2p sharing converge on one 2026 reality: persistent, addressable, cross-instance KV storage is production infrastructure — and the Open KV Cache API effort is trying to make it interoperable.
2. **Prefill cost at 1M+ tokens is a solved problem *given reuse*.** 100–1,000× TTFT improvements are measured, not projected — but entirely conditional on cache reuse. Workload analysis (multi-turn/agentic vs one-shot) now precedes any infrastructure decision.
3. **FP8 KV is the production default; the dtype choice moved down one level.** With `fp8_e4m3` pinned as the recipe standard on both vLLM and SGLang, SGLang's broader matrix (FP4 experimental, INT4/INT2, NVFP4) is where the next production battle is: capacity at 3.56× BF16 is on the table, kernel fusion decides whether it is also faster.
4. **Quantized-KV performance is kernel-fusion-dependent, not a dtype property.** The AMD-fork divergence (SGLang faster, vLLM-ROCm slower at fp8 KV) generalizes: benchmark the fused path on your hardware before standardizing a dtype.
5. **The Hopper accumulation-precision bug is the 2026 case study in silent long-context cliffs.** A 91%→13% NIAH collapse with no error signal is why every KV-stack change needs a 128K needle probe as a validation gate, not just perplexity checks. Verify flash-attention#96/#91 before routing traffic.
6. **Transport is a line item in the serving budget.** Mooncake Store TCP 17–22% / RDMA 26–33% vs Redis, plus llm-d's −55.8%→−88.2% prefill deltas, give the network layer a quantified ROI: provision RDMA where disaggregation crosses node boundaries, CacheGen-compressed bitstreams where it crosses regions.
7. **Cross-engine KV interop is the unpriced risk.** The digest-truncation bug class (vLLM last-8 vs SGLang first-8) and the "cache index is a convention; interface is a contract" lesson mean disaggregated fleets mixing engines will bite harder before the Open KV Cache API lands — a standard KV-block hashing and transfer format does not yet exist.
8. **Disaggregation is the default at scale, converged below it.** Multi-node long-context/agentic fleets: split prefill/decode (vLLM V1, SGLang EPD, Dynamo). Single-GPU: stay converged with chunked prefill + FP8 KV. Workload analysis decides.
9. **"1M context" is a flagship-tier spec, not a universal product reality.** Tier caps (Foundry 200K for Opus 4.8), 256K exceptions (Qwen3.6), and vendor-run NIH-at-1M validation mean procurement should validate the *deployment tier*, not the headline spec — and run needle-style retrieval tests for retrieval-critical workloads rather than trusting length alone.
10. **Cache reuse is a pricing primitive.** At 95%+ hit rates, "cost per token" is meaningless without the cached/uncached split; the ~9× effective input-cost cut on agentic loops makes prefix/persistent caching the first economic lever, before any compression.
11. **Security and tenancy are KV-layer requirements now.** Exact tokens are recoverable from cache dumps; the Open KV Cache API bakes in tenant/session scoping and ciphertext-outside-GPU from the start — compliance teams should treat shared KV tiers as sensitive data stores, not scratch space.
12. **HBM economics dominate technical choices.** Sold out through 2028, +40% in a quarter, 50–60% of GPU BOM: every technique in this part should be costed in $/GB-of-HBM-avoided terms. The KV-cache work of 2026 is HBM-demand destruction.
13. **DSA reduces compute; it does not shrink the cache.** Any plan that budgets KV memory against a sparse-attention mechanism without a separate cache-compression scheme (CSA/HCA-class) is misreading the ledger — the indexer can select any historical token, so full-fidelity KVs must be retained.
14. **The consumer tier is real.** KTransformers (128K–1M on consumer hardware via CPU-resident sparse attention), turboquant-mlx (Qwen 27B at 128K on M4 Pro 48GB), and FoveatedKV (75% KV cut, ≤3% quality loss, independently benchmarked on M3 Max) show the KV stack reaching local hardware — the long-context story is no longer GPU-cluster-only.
15. **Two independent ceilings bind supply: wafer output and CoWoS packaging.** The HBM wall is not only price but physical throughput (6–9 month CoWoS lead times; panel-level CoPoS targeted at the 2028 GPU generation). Infrastructure plans that assume elastic HBM supply on a 12–18 month horizon should treat that as a modeling assumption to defend, not a default.
16. **EPD and encoder-cache connectors are the 2026 answer to multimodal KV.** Splitting encoding out as a third disaggregation stage (SGLang EPD, vLLM encoder-cache connectors with CPU offload) acknowledges that vision/audio encoders are a distinct workload shape — operators sizing pools for multimodal models should plan three stages, not two.
17. **The same architecture serves both cluster and local.** Tiered offload, sparse CPU attention, and online quantization are the same primitives whether the "remote tier" is another node or host DRAM — local/private inference stacks should adopt the disaggregated design, not a separate one.

## Sources and URLs (continued)

- [WAVE1] https://temperature2.com/p/2026-09-08-did-you-know-multi-head-latent-attention/
- [WAVE1] https://github.com/frank-1150/learning-ai-technical/blob/HEAD/docs/machine-learning/inference/prefill-decode-disaggregation-mooncake.md
- [WAVE1] https://github.com/balakreshnan/samples2026/blob/HEAD/AIStrategy/llm-inference-optimization-and-scaling.md
- [WAVE1] https://github.com/loxilb-io/loxilb-inference-gateway/blob/HEAD/docs/load-balancing/16-sglang-vs-vllm-routing-differences.md
- [WAVE1] https://github.com/claude-dev-suite/claude-dev-suite/blob/HEAD/skills/ai-systems/inference-serving-topology/SKILL.md
- [WAVE2] https://github.com/lmcache/lmcache/blob/HEAD/README.md
- [WAVE2] https://github.com/qiuziyan1998/lmcache/blob/HEAD/docs/source/getting_started/quickstart/offload_kv_cache.rst
- [WAVE2] https://www.lightbitslabs.com/blog/follow-up-on-paged-attention-over-rdma/
- [WAVE2] https://github.com/llm-d/llm-d/blob/HEAD/guides/p2p-kv-cache-sharing/README.md
- [WAVE2] https://www.snia.org/sites/default/files/2025-10/SNIA-SDC25-Kaynar-KV-Cache-Storage-Offloading.pdf
- [WAVE2] https://github.com/kvcache-ai/ktransformers/blob/HEAD/doc/en/long_context_introduction.md
- [WAVE2] http://arxiv.org/pdf/2310.07240
- [WAVE2] https://docs.kanaries.net/articles/llm-cache
- [WAVE2] https://medium.com/@adityaj5400/the-kv-cache-is-killing-your-llm-at-scale-heres-the-low-level-physics-nobody-talks-about-b577c4c7549e
- [WAVE2] https://github.com/kingsdigitallab/issa/blob/HEAD/workshops/ws1/docs/03-fp8_kv_cache_analysis.md
- [WAVE2] https://github.com/bayesiansapien/cere-bro/blob/HEAD/wiki/inference-efficiency/kv-cache.md
- [WAVE2] https://madsys.cs.tsinghua.edu.cn/publication/mooncake-a-kvcache-centric-disaggregated-architecture-for-serving/ToS2025-Qin.pdf
- [WAVE2] https://github.com/trishambp/trishambp.github.io/blob/HEAD/_implementations/mooncake-kvcache-centric-architecture-for-serving-llm-chatbot.md
- [WAVE2] https://arxiv.org/pdf/2407.00079v3
- [WAVE2] https://ai-cost-estimator.com/blog/sk-hynix-28b-ipo-hbm-monopoly-ai-inference-cost-impact
- [WAVE2] https://www.techtimes.com/articles/317269/20260527/memory-chip-shortage-hits-3-trillion-market-value-smartphones-cost-14-more-2026.htm
- [WAVE2] https://momoview.com/blog/en/posts/hbm-industry-analysis-sk-hynix-samsung-micron-2026-ai-memory-supercycle-investment-thesis/
- [WAVE2] https://github.com/hczhu/stock-research/blob/HEAD/memos/2026-06-01-memory-super-cycle-x-thread-and-memo.md
- [WAVE2] https://www.ainvest.com/news/ai-bottleneck-isn-gpu-anymore-memory-company-owns-tap-2606/
- [WAVE2] https://www.ainvest.com/news/memory-curve-ai-infrastructure-bottleneck-2026-supercycle-2601/
- [WAVE2.1] https://github.com/sgl-project/sglang/blob/HEAD/docs/docs/advanced_features/quantized_kv_cache.mdx
- [WAVE2.1] https://github.com/sgl-project/sglang/pull/6109
- [WAVE2.1] https://github.com/semianalysisai/inferencex-app/blob/HEAD/packages/app/content/blog/sglang-0-5-6-b200-deepseek-r1-fp4-up-to-1-8x.mdx
- [WAVE2.1] https://inferencex.semianalysis.com/blog/sglang-0-5-6-b200-deepseek-r1-fp4-up-to-1-8x
- [WAVE2.1] https://github.com/vllm-project/vllm/pull/48250
- [WAVE2.1] https://github.com/vllm-project/vllm/releases/
- [WAVE3] https://github.com/vllm-project/vllm/releases/tag/v0.26.0
- [WAVE3] https://github.com/sgl-project/sglang/pull/30514
- [WAVE3] https://github.com/iopsystems/llm-calc/blob/HEAD/docs/superpowers/specs/2026-05-12-dsa-design.md
- [WAVE3] https://github.com/suzeai/transformers/blob/HEAD/docs/source/en/model_doc/glm_moe_dsa.md
- [WAVE3] https://earlyterms.com/term/indexshare
- [WAVE3] https://www.marktechpost.com/2026/02/19/google-ai-releases-gemini-3-1-pro-with-1-million-token-context-and-77-1-percent-arc-agi-2-reasoning-for-ai-agents/
- [WAVE3] https://devtk.ai/en/blog/gemini-3-1-pro-pricing-guide-2026/
- [WAVE3] https://apidog.com/blog/what-is-deepseek-v4/
- [WAVE3] https://dev.to/agdex_ai/deepseek-v4-1m-context-open-source-agentic-coding-sota-what-ai-builders-need-to-know-2026-3m3n
- [WAVE3] https://felloai.com/it/deepseek-v4/
- [WAVE3] https://felloai.com/anthropic-claude-opus-4-8/
- [WAVE3] https://github.com/neetx/ai-research-radar/blob/HEAD/reports/2026-06-16.md
- [WAVE3] https://github.com/pestopoppa/epyc-root/blob/HEAD/wiki/ssm-hybrid.md
- [WAVE3] https://github.com/mtgibbs/pi-cluster/blob/HEAD/docs/model-eval-2026-05.md
- [WAVE3] https://techcrunch.com/2026/01/22/inference-startup-inferact-lands-150m-to-commercialize-vllm/
- [WAVE3] https://www.businesswire.com/news/home/20260505077157/en/RadixArk-Launches-with-%24100-Million-in-Seed-Funding-Led-by-Accel-to-Grow-SGLang-and-Democratize-Frontier-AI-Infrastructure
- [WAVE3] https://github.com/xianlubird/sglang
- [WAVE3] https://github.com/yusanxy/llm_flops/blob/HEAD/operators/references/deepseek_v32_dsa_sparse_attention/README.md
- [WAVE1] https://www.spheron.network/blog/nvidia-groq-3-lpu-explained/
- [WAVE1] https://www.businessworld.in/article/nvidia-gtc-2026-jensen-huang-vera-rubin-groq-lpu-trillion-dollar-forecast-597897
- [WAVE1] https://byteiota.com/nvidia-vera-rubin-20b-groq-lpu-integration-at-gtc-2026/
- [WAVE2] https://www.marktechpost.com/2026/06/12/zyphra-release-zamba2-vl-hybrid-mamba2-transformer-vision-language-models-that-cut-time-to-first-token-by-about-an-order-of-magnitude/

