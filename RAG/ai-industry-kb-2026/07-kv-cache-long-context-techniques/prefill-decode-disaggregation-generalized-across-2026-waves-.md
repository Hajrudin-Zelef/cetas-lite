---
id: ai-industry-kb-2026/07-kv-cache-long-context-techniques/prefill-decode-disaggregation-generalized-across-2026-waves-
title: "Prefill/decode disaggregation generalized across 2026 (Waves 1–3)"
domain: kv-cache-long-context-techniques
role: deep-dive
task: architecture
actors: ["AMD", "DeepSeek", "Nvidia", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2025-07", "2026-08-26", "2026-09-17"]
keywords: ["decode", "prefill", "amd", "attention", "compute", "deepseek", "disaggregated", "disaggregated serving", "dram", "fp8", "gpu", "gpus"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3423, 3444]
section: "7. KV Cache & Long-Context Techniques"
sha256: 3a66c5fbde74653895273d1a9a914320fb3baf1f4b3cf3dbced5ff824ba54b36
---

# Prefill/decode disaggregation generalized across 2026 (Waves 1–3)

- **Lightbits follow-up (Arthur Rasmusson, 2026-09-17)**: the July 2025 blueprint — NVMe as a seamless extension of HBM, KV pages moved over RDMA so turn 2 skips turn 1's prefill — was built out to a **10.5M-token session on 8×H100** that takes **1 hour 42 minutes to rebuild through prefill** and is **restored in seconds with bit-identical output**, surviving engine and fabric restarts.
- Measured Lightbits figures (Llama-70B FP16 KV, turn-2 TTFT): 8K context 1199→**38.8 ms**; 32K 4935→**88.7 ms**; 128K ~23 s→**~300 ms**. With FarmGPU/ScaleFlux at 131K–1M tokens: **100–280× TTFT improvement; >1,000× at 10M tokens**.
- Four engineering lessons from the build: (1) prefetch volume should track attention *sparsity*, not context length — prefetching every block a dense window could touch is linear and untenable at 10M; (2) **"KV page fault" is the operative framing** — a touched offloaded block stalls the GPU, so the metric that matters is fault *rate*, not fetch bandwidth, and the tail matters more than the mean under multi-tenant load; (3) **fail-closed design** — a GDS→POSIX fallback created a DMA-to-POSIX downgrade path; the fix was automatic topology detection, not config files; (4) **a cache index is a convention; an interface is a contract** — every engine built a globally-addressable KV index, none interoperate, hence the **Open KV Cache API** (deterministic block identity, typed metadata, async data plane, mandatory tenant/session scoping) with **Inferra** as reference implementation.
- Bit-identicality is load-bearing: the streamed windowed form uses online softmax accumulation (running max + sum of exponentials), numerically identical to the monolithic computation — an infrastructure change with no eval regression. Stated limits: decode is untouched; short-prompt/long-answer workloads see little benefit; at 10.5M tokens the prefill baseline "doesn't run" — the comparison is between an offering and a refusal.
- Security is designed in from the start: a KV cache is a **lossy-but-invertible encoding of the prompt** — published work recovers exact tokens from a cache dump — so the design uses ciphertext everywhere except inside the GPU.
- **llm-d peer-to-peer KV sharing guide (Sep 2026)** measured the pull-vs-recompute crossover on **gpt-oss-120b with RDMA/IB**: prefill-latency deltas of **−55.8% (2K), −77.4% (8K), −83.2% (16K), −85.9% (32K), −88.2% (49K)** — the pull wins at every measured length, hence `minCachedTokenDelta: 2048`. The guide notes RDMA is recommended but not required (NIXL/UCX falls back to TCP): **the transport sets the crossover**, the same crossover LMCache exposes as `min_retrieve_tokens`.
- MoE-scale RDMA datapoint (2026 survey): an **8,192-token DeepSeek-R1 FP8-KV prefill moves ~290 MB of KV over RDMA**. Moving DeepSeek-R1 (256 routed experts) from EP8 on one node to **EP64 across 64 GPUs** cuts experts-per-GPU from 32 to 4, **freeing HBM for KV cache** and enabling larger concurrent batches — NVIDIA reporting up to **2.28× higher per-GPU output throughput** from Wide-EP [VENDOR]. Expert parallelism is therefore a KV-capacity lever, not only a compute-sharding one. AMD's **MoRI H2-2026 roadmap** makes tiering explicit: a tiered distributed KV cache (HBM → DRAM → NVMe via SPDK/GDS) with scheduler co-design.
- The 2026 decision rule emerging from measured data: co-located prefill/decode → skip transfer, keep KV local; cross-node disaggregation → NIXL/RDMA transport; cross-region or cold tiers → CacheGen-compressed bitstreams over plain object storage. Layering: **engine → transfer API (NIXL) → storage tier** — "the 2026 answer to the 2024 question 'where does the KV live': *everywhere, with the transport matched to the distance.*"

### Prefill/decode disaggregation generalized across 2026 (Waves 1–3)

- The 2023–2026 systems literature converged: **stop running prefill (compute-bound) and decode (memory-bandwidth-bound) on the same GPU**. Separate pools get separate SLOs (TTFT vs TPOT), separate hardware sizing, independent autoscaling. Production implementations: vLLM V1 (P/D two-stage, `ep_role:1/2`), SGLang (P/D or three-stage **EPD**: Encode-Prefill-Decode, splitting vision/audio encoding out), TensorRT-LLM (KV Cache Connector API), Mooncake/Mooncake Store (RDMA/NVLink), NVIDIA Dynamo (scheduler-level P/D orchestration). Wave 3 verified this as an **industry-wide architectural shift in 2026**, not a single event.
- **Measured reference**: SGLang's disaggregated serving on 96 H100s running a DeepSeek model hits **52,300 input tokens/sec and 22,300 output tokens/sec per node** (2026 disaggregation survey).
- Disaggregation's engineering crux is KV transfer bandwidth: the full cache moves from prefill to decode instances over NVLink/InfiniBand/RDMA — one Llama 3 70B request at 4K context in BF16 transfers **~13.4 GB**. This couples compression to disaggregation: a 6× smaller cache is a 6× cheaper transfer.
- The common 2026 production pairing: **vLLM as the token engine + Triton as the serving shell**, with Ray Serve/KServe + KEDA for multi-node orchestration. Disaggregation pays at scale; single-GPU deployments stay converged with chunked prefill piggy-backed on decode via Sarathi-style scheduling.
- vLLM v0.26.0 (GitHub release notes: **411 commits from 212 contributors**): DeepSeek-V4 performance push plus maturation of **KV offloading and tiered secondary storage** — an object-store secondary tier, DP-replica-aware tiering, encoder-cache connectors including CPU offload. Wave 2.1/§6 covers vLLM v0.28.0 (tagged 2026-08-26) tiered KV offload including disk; version-level engine detail lives in §6, not repeated here.
- **PagedAttention currency: CONTRADICTED as a current feature.** PagedAttention remains vLLM's historically defining technique, but it was **removed from the new internal engine in v0.25.0** and survives only as a legacy attention path; the 2026 serving story is KV offloading, sparse attention, and disaggregation.
- Disaggregation created a new control plane: **KV-cache-aware gateways** that route based on which KV blocks already exist where. The 2026 reference design (loxilb inference gateway): **kvExactMode** selects endpoint topology; requests carry **SHA-256 (or XXH3-128) hashes of canonical KV block digests**; the gateway dual-dispatches to prefill AND decode concurrently, matching prefix hashes to warm decode instances; KV events flow over ZMQ PUB sockets with msgpack envelopes. Real interop hazard documented: **vLLM truncates digests to the *last* 8 bytes big-endian while SGLang uses the *first* 8** — a mis-slice produced 0% cache overlap before being fixed. A standard KV-block hashing and transfer format does not yet exist.
- Complement inside each pool: **iteration-level continuous batching** (Orca) plus **chunked prefill** (Sarathi-Serve) — batch continuously, chunk the prefill, split the phases, reuse the prefixes.

### Prefix caching and persistent KV reuse across requests (Waves 1–2)

