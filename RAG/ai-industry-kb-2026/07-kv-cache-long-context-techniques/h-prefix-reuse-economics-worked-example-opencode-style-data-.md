---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/h-prefix-reuse-economics-worked-example-opencode-style-data-
title: "H. Prefix-reuse economics — worked example (OpenCode-style /data, June 2026)"
domain: kv-cache-long-context-techniques
role: deep-dive
task: finance
actors: ["AMD", "DeepSeek", "Google", "Groq", "MiniMax", "Moonshot", "Nvidia", "SGLang", "Samsung", "TensorRT-LLM", "Z.ai", "vLLM"]
dates: ["2026-01-22", "2026-03-25", "2026-05", "2026-05-05", "2026-06", "2026-08-29", "2026-09-17"]
keywords: ["acquisition", "agent", "agentic", "amd", "attention", "benchmarks", "consumer", "cost", "decode", "deepseek", "fp8", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3616, 3657]
section: "7. KV Cache & Long-Context Techniques"
sha256: 9853ffb01115fe9ee137016953fef34ba252aa79c54a7c8f3b352a5c1965677b
---

# H. Prefix-reuse economics — worked example (OpenCode-style /data, June 2026)

| Topology | Transfer choice | Measured anchor |
|---|---|---|
| Co-located P/D (same node) | None — keep KV local | N/A (baseline) |
| Cross-node P/D | NIXL / RDMA transport | llm-d: −55.8%→−88.2% prefill deltas; pull beats push at every length |
| Remote / cold tiers | CacheGen-compressed bitstreams over object storage | 3.5–4.3× less bandwidth; fallback to text+recompute |
| Open KV standard (aspirational) | Open KV Cache API (deterministic block identity, tenant/session scoping) | No engines interoperate yet — "a cache index is a convention; an interface is a contract" |

Digest conventions (the interop hazard): requests carry **SHA-256 (or XXH3-128) hashes of canonical KV block digests**; vLLM truncates to the **last 8 bytes big-endian**, SGLang uses the **first 8** — the documented mis-slice produced 0% cache overlap before being fixed. KV events flow over **ZMQ PUB sockets** (one per prefill endpoint in vLLM; one per DP rank in SGLang) with msgpack envelopes.

### H. Prefix-reuse economics — worked example (OpenCode-style /data, June 2026)

Agent loop: 20 calls against a 100K-token repo context. Full prefill once (~100K uncached tokens) + ~1.9M cached tokens; at typical **10:1 cached:uncached pricing → ~9× effective input-cost cut** vs no caching. Multiplicative composition: a 6×-compressed (TurboQuant-class) cache reused across 20 agent turns ≈ **~100× prefill saving** vs naive per-call full prefill. This is why providers price cached input tokens separately and production traces show **95%+ hit ratios** on agentic workloads.

### I. Prefill→decode KV transfer cost math (Wave 1, §7.3)

| Request (Llama 3 70B, BF16) | KV moved P→D | Reading |
|---|---|---|
| 4K context | **~13.4 GB** | For short prompts, transfer latency can dominate total TTFT |
| 128K context | ~430 GB | Long contexts: the price of disaggregation is the transfer itself |
| Same at FP8 KV | ~215 GB | Compression halves the disaggregation tax — the coupling |

This is why KV *compression* and disaggregation are coupled design problems, not separate ones: a 6×-smaller cache is a 6×-cheaper transfer, and every transfer-cost figure above shrinks multiplicatively with the dtype and compression choices in §B–§D.

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

