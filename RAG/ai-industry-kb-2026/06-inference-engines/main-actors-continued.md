---
id: ai-industry-kb-2026/06-inference-engines/main-actors-continued
title: "Main actors (continued)"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["AMD", "AWS", "Alibaba", "Baseten", "DeepSeek", "Google", "Huawei", "Hugging Face", "Intel", "LongCat", "Meta", "Microsoft", "Moonshot", "Nebius", "Nvidia", "OpenAI", "Oracle", "SGLang", "TensorRT-LLM", "Unsloth", "Z.ai", "vLLM", "xAI"]
dates: ["2025-03", "2025-10-29", "2025-12-15", "2026-01", "2026-01-22", "2026-03-15", "2026-05", "2026-05-04", "2026-05-05", "2026-05-12", "2026-07", "2026-07-14", "2026-08", "2026-08-05", "2026-08-20", "2026-08-26", "2026-09", "2026-09-04", "2026-09-05", "2026-09-20", "2026-09-22"]
keywords: ["agent", "agentic", "amd", "apache", "ascend", "attention", "aws", "benchmark", "benchmarks", "blackwell", "datacenter", "decode"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2882, 2991]
section: "6. Inference Engines"
sha256: 506c270b2396c9b4b2cae8bd1992dd0de7418931b094a23ce665314dd25961e1
---

# Main actors (continued)

## Main actors (continued)

### SGLang and RadixArk

- SGLang team (LMSYS/UC Berkeley, Chatbot Arena lineage); RadixArk founders Ying Sheng and Banghua Zhu (xAI and NVIDIA backgrounds); SGLang continuity as Apache 2.0.
- RadixArk investors: Accel (lead), Spark Capital (co-lead); NVentures (NVIDIA), AMD, MediaTek, Salience Capital, A&E Investments, HOF Capital, Walden Catalyst Ventures, LDV Partners, WTT Investment.
- RadixArk angels: Igor Babuschkin, Lip-Bu Tan, Hock Tan, John Schulman, Soumith Chintala, Olivier Pomel, Thomas Wolf, William Fedus, Robert Nishihara, Eric Zelikman, Logan Kilpatrick.
- a16z OSS AI Grant funder of SGLang; SGLang in the PyTorch ecosystem since March 2025.

### Frontier adopters

- xAI (Grok 3 on SGLang), Microsoft Azure (DeepSeek R1 on AMD), Oracle Cloud, Google Cloud, AWS, Nebius, DataCrunch, Voltage Park, Cursor (code completion), LinkedIn (AI features); NVIDIA, AMD, Intel as both vendors and users; Baseten, RunPod, Novita.
- Academia: Stanford, UC Berkeley, UCLA, MIT, UW, Tsinghua.
- RL stack: verl, AReaL, Miles (RadixArk's own framework), slime, Tunix; SGLang's OME Kubernetes operator.

### Adjacent engine actors

- MLC (XGrammar-2 authors); xAI, Databricks, DeepSeek (XGrammar-2 adopters).
- Arcfra (Neutree 1.2 / Flex Engine); TileRT; Tiny-vLLM community; Hugging Face (TGI maintenance); Meta (ExecuTorch 1.0 GA).
- Ant Group (MOVA cookbook for SGLang Diffusion); Google (SGLang-JAX TPU partnership, July 2026 expansion).
- Kernel and fabric vendors inside SGLang's stack: FlashInfer (BatchPrefill/BatchDecode, sparse-MLA kernels), DeepEP (MoE all-to-all), MORI on AMD, NIXL/Mooncake (P/D transfer) — the 2026 engine is as much a composition of specialized kernels as a codebase.
- Project tooling: sgl-router (cache-aware routing), sgl-kernel (ops), OME (K8s operator), sglang-omni (any-to-any track) — the SGLang surface extends well beyond the core server.

### Dedup pointers for this part

- vLLM core internals, vllm-gguf-plugin, vLLM-Omni, Inferact, NIM 2.0, Dynamo vs llm-d, spend flip, per-token economics → sibling part 06a (one-line cross-refs only in this part).
- KV-cache techniques (radix-tree deep mechanics, paged-KV internals, tiered offload, FP8 KV dtype, APC mechanics) → §7 authoritative; this part retains only engine-level outcome figures.
- TPU hardware specifics → §15; Unsloth platform internals → the Unsloth-track source (§4 of wave 1) where needed.
## Timeline and context (continued)

| Date | Event |
|---|---|
| 2025-08 | RadixArk first announced as a startup (~$400M valuation, Accel-led round reported) — precursor to the formal launch |
| 2025-10-29 | SGLang-JAX announced (native JAX/XLA TPU engine) |
| 2025-12 | DeepSeek-V3.2 introduces DeepSeek Sparse Attention (lightning indexer, top-K 2048); Hugging Face TGI enters maintenance mode |
| 2025-12-15 | llama.cpp model router announced (multi-process, LRU eviction, OpenAI-compatible dynamic loading) |
| 2026-01-22 | TechCrunch reports SGLang commercialization "seeking a $400M valuation" — the first $400M press datapoint; Inferact launches for vLLM the same day ($150M seed, $800M valuation — see 06a) |
| 2026-02 | SGLang blog: "Unlocking 25x Inference Performance with SGLang on NVIDIA GB300 NVL72" [VENDOR headline, baseline unaudited] |
| 2026-03-15 | Enterprise comparison published: vLLM vs SGLang for production (privocto) — 4.7x multi-turn speedup, 99.8% JSON validity [DIRECTIONAL] |
| 2026-04 | SGLang day-zero support for DeepSeek-V4 |
| 2026-05-04 | XGrammar-2 released — Structural Tag protocol; adopted by xAI, Databricks, DeepSeek |
| 2026-05-05 | **RadixArk formal launch**: $100M seed at $400M post-money, led by Accel, co-led by Spark Capital; NVentures, AMD, MediaTek participate |
| 2026-06 | SGLang FlashInfer SM120 sparse-MLA decode PR #27455 (~2.2–3.7x TPOT); Nemotron and Higgs Audio added to SGLang; DeepSeek V4 day-zero GLM-5.2 IndexShare ships elsewhere as sparse-attention context |
| 2026-06 | One operator parks SGLang for DeepSeek V4 Flash (NVFP4 loader issue), serves via vLLM — per-model engine choice |
| 2026-07 | SGLang day-zero support for Kimi K3; RadixArk expands Google TPU partnership (SGLang-JAX as commercial vehicle) |
| 2026-07-14 | vLLM blog announces vLLM prefill + TileRT decode pairing |
| 2026-07 | TensorRT-LLM v1.3.0rc21–rc23; AutoDeploy backend deprecated (rc21) |
| 2026-08-05 | TensorRT-LLM day-0 support for OpenAI GPT-OSS-120B/20B |
| 2026-08-20 | SGLang v0.5.18 tagged |
| 2026-09-04 | **SGLang v0.5.19 on PyPI** (786 PRs, 214 contributors): Qwen3.8, dots3.note, Ling-3.0, Spark2.5, MiniCPM-SALA, Granite 4.2, LongCat-Image-Edit |
| ~2026-09-05 | llama.cpp server v0.4.0 |
| 2026-09 | Neutree 1.2 / Flex Engine (Arcfra): vLLM + SGLang under one gateway |
| 2026-09-20 | Architecture review: up to 2.5x throughput gain under strict JSON constraints (jump-forward decoding) — the origin of the misread "2.5x cache hit rate" |
| 2026-09-22 | Research cutoff for all wave documents |

- Context arc: SGLang spent 2026 converting its 2025 scheduler advantages (RadixAttention, zero-overhead scheduling) into a full platform story — diffusion serving, a K8s operator (OME), a native TPU engine (SGLang-JAX), and a funded commercial steward (RadixArk) — while TensorRT-LLM converged onto the same disaggregated/NIXL patterns as the OSS engines and its commercial center of gravity shifted toward NIM packaging of non-LLM workloads.
- The two-engine convergence pattern of 2026: each engine absorbs the other's differentiators (SGLang adds FP4/quantized KV, multi-hardware breadth, K8s ops; vLLM absorbs disaggregation, diffusion, decode-speed — see 06a) while deepening its own moat — scheduler/prefix-affinity for SGLang, enterprise distribution for vLLM.

## Implications (continued)

1. **Engine selection is workload-shaped, not per-religion:** the same-kernel +29% SGLang win, vLLM's TTFT/100+-concurrency counter-evidence, and the operator serving DeepSeek V4 Flash on vLLM while keeping SGLang elsewhere all point to per-workload, per-model choice behind a gateway — not engine loyalty.
2. **The `lpm` opt-in gap is the cheapest win in the SGLang track:** cache-aware scheduling (`--schedule-policy lpm`/`dfs-weight`) is off by default and likely underused in the wild; `--enable-cache-report` turns hit rates into an observable metric before any tuning spend.
3. **Don't conflate throughput gains with cache-hit ratios:** the brief's "2.5x cache hit rate vs competition" was a misread of a 2.5x jump-forward throughput figure — the consolidation should cite 3.8x-vs-round-robin router figures and the paper's ~96%-of-optimal figure instead.
4. **Commercialization changes procurement calculus:** RadixArk ($100M, May 2026) narrows SGLang's enterprise-packaging gap vs Red Hat/IBM vLLM; as of September 2026 the comparison is a two-commercial-stack one, still favoring vLLM's shipped server — revisit after RadixArk's paid-hosting tiers materialize.
5. **"4x batch size per GPU" is a planning figure, not a benchmark:** until a controlled OOM/SLO-breaching batch test is published, budget MLA capacity with the paper's 93.3% KV reduction as the ceiling and treat ~4x as conservative-derivative guidance.
6. **Time-to-serve is the 2026 differentiation metric:** both vLLM v0.28.0 and SGLang v0.5.19 raced the same flagships (Qwen3.8, Kimi-K3) across NVIDIA/AMD/Ascend in the same month — watch release-note race dynamics, not just architecture.
7. **Measure on your hardware class:** the August 2026 TRT-LLM vs vLLM measurement and the SM121 NaN-bisecting community patches show workstation-Blackwell results don't transfer from datacenter-Hopper claims — the SM90+-first reality of 2026 MLA/sparse kernels demands per-silicon validation.
8. **The post-training feedback loop is SGLang's durable moat:** verl/AReaL/slime/Tunix run on SGLang, optimizations land there first, frontier open-weight models serve best there — this is where tomorrow's production models are born, and vLLM cannot match it through distribution alone.
9. **Date-stamp every version claim:** vLLM moved v0.27 → v0.28.0 (2026-08-26) with v0.29.0 already attempted; SGLang v0.5.18 → v0.5.19 (2026-09-04) — any "latest version" claim without a date rots within weeks.
10. **Structural decoding is commoditized; the edge is forward-pass elimination:** XGrammar is universal in 2026 — SGLang's differentiator is specifically jump-forward decoding, and XGrammar-2's Structural Tag is the protocol the agent stack (xAI, Databricks, DeepSeek) is converging on.
11. **January 2026 priced the duopoly's commercial layer:** $100M-at-$400M (SGLang/RadixArk) vs $150M-at-$800M (vLLM/Inferact) on the same day — future buyers should read engine choice against these two commercial stacks, not against "open vs proprietary."
12. **Prefix discipline is an operator skill:** with hit rates spanning 50–99% on prompt design, the 2026 practitioner's first tuning pass is prompt templating and `--enable-cache-report` instrumentation — engine upgrades come second.

## Sources and URLs (continued)

- [VENDOR] https://www.businesswire.com/news/home/20260505077157/en/RadixArk-Launches-with-%24100-Million-in-Seed-Funding-Led-by-Accel-to-Grow-SGLang-and-Democratize-Frontier-AI-Infrastructure
- [VENDOR] https://github.com/sgl-project/sglang/releases/tag/v0.5.19
- [VENDOR] https://github.com/sgl-project/sglang-omni/releases/tag/v0.1.4
- [VENDOR] https://github.com/lm-sys/lm-sys.github.io/blob/HEAD/blog/2025-10-29-sglang-jax.md
- [VENDOR] https://github.com/sgl-project/sglang/blob/HEAD/docs/docs/hardware-platforms/tpu.mdx
- [VENDOR] https://github.com/sgl-project/sglang/pull/30514
- [VENDOR] https://github.com/vllm-project/vllm/releases/tag/v0.28.0
- [VENDOR] https://github.com/mlc-ai/blog/blob/HEAD/_posts/2026-05-04-xgrammar-2-fast-customizable-structured-generation.md
- [VENDOR] https://unsloth.ai/docs/basics/unsloth-dynamic-2.0-ggufs
- [VENDOR] https://huggingface.co/unsloth/Qwen3-0.6B-GGUF
- [VENDOR] https://openreview.net/pdf?id=VqkAKQibpq
- [DIRECTIONAL] https://www.morningstar.com/news/business-wire/20260505077157/radixark-launches-with-100-million-in-seed-funding-led-by-accel-to-grow-sglang-and-democratize-frontier-ai-infrastructure
- [DIRECTIONAL] https://pulse2.com/radixark-launches-with-100-million-in-seed-funding-to-democratize-frontier-ai-infrastructure/
- [DIRECTIONAL] https://techcrunch.com/2026/01/22/inference-startup-inferact-lands-150m-to-commercialize-vllm/
- [DIRECTIONAL] https://pulse2.com/inferact-launches-with-150-million-funding-at-800m-valuation-to-commercialize-vllm-as-inference-demand-surges/
- [DIRECTIONAL] https://theaiinsider.tech/2026/01/27/inferact-launches-with-150m-seed-round-to-commercialize-vllm-inference-engine/
- [DIRECTIONAL] https://www.bbw9n.io/blog/sglang_vs_vllm
- [DIRECTIONAL] https://privocto.com/blog/vllm-sglang
- [DIRECTIONAL] https://felloai.com/it/deepseek-v4/
- [DIRECTIONAL] https://apidog.com/blog/what-is-deepseek-v4/
- [COMMUNITY] https://explore.n1n.ai/blog/sglang-vs-vllm-architecture-radixattention-benchmarks-2026-09-20
- [COMMUNITY] https://github.com/xianlubird/sglang
- [COMMUNITY] https://github.com/rishikinger10/radixscope/blob/HEAD/Docs/sglang-integration.md
- [COMMUNITY] https://github.com/softwealth/eval-report-skills/blob/HEAD/skills/inference/xgrammar-structured-output.md
- [COMMUNITY] https://github.com/tabnas/parser/blob/HEAD/ts/doc/gbnf-feasibility.md
- [COMMUNITY] https://github.com/jedmund/homelab/blob/HEAD/roles/sglang/README.md
- [COMMUNITY] https://github.com/profsynapse/synaptic-tuner/blob/HEAD/docs/preparation/vllm-vs-sglang-inference-serving-research.md
- [COMMUNITY] https://github.com/shurankain/agentic-ai-course/blob/HEAD/17_Production_Inference/07_SGLang_and_Alternatives.md
- [COMMUNITY] https://github.com/iopsystems/llm-calc/blob/HEAD/docs/superpowers/specs/2026-05-12-dsa-design.md
- [COMMUNITY] https://github.com/yusanxy/llm_flops/blob/HEAD/operators/references/deepseek_v32_dsa_sparse_attention/README.md
- [COMMUNITY] https://media.patentllm.org/news/hardware/tensorrt-llm-v1-3-0rc21-released-nvidia-gb300-achieves-moe-w-20260728
- [COMMUNITY] https://media.patentllm.org/news/hardware/tensorrt-llm-v1-3-0rc23-released-amd-mi450-nvidia-rtx-5090-o-20260731
- [COMMUNITY] https://github.com/dcharlot-physicalai-bmi/ferric/blob/HEAD/docs/runtime-landscape-2026-08.md
- [COMMUNITY] https://data.safetycli.com/packages/pypi/tensorrt-llm/changelog?page=2

