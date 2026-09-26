---
id: ai-industry-kb-2026/06-inference-engines/hardware-tracks-dated-facts
title: "Hardware tracks — dated facts"
domain: inference-engines
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "CoreWeave", "DeepSeek", "Google", "Intel", "Mistral", "Moonshot", "Nvidia", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2026-04-28", "2026-06-05"]
keywords: ["amd", "aws", "benchmarks", "cost", "decode", "deepseek", "disaggregated", "disaggregated serving", "dram", "gpu", "gpus", "hbm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2311, 2326]
section: "6. Inference Engines"
sha256: 9fbd7032e9bc81b4d77ba635d7938f2e1abf2eac45ea9ccecd6a899f3274f555
---

# Hardware tracks — dated facts

- **Three complementary libraries, not competitors:** **NCCL** answers "what collective operation should these GPUs perform?" (AllReduce/AllGather every layer over NVLink). **UCX** answers "how do data move between these endpoints?" (fastest transport: SHM, verbs, TCP). **NIXL (NVIDIA Inference Xfer Library)** answers "how does an inference system move data between heterogeneous memory/storage resources?" — e.g. "this request's KV blocks (32k context, ~4 GiB) now live on prefill worker P-3; decode worker D-7 needs them; move them asynchronously and signal D-7". NIXL is the KV transfer layer of vLLM (NixlConnector), NVIDIA Dynamo (KVBM), SGLang, TensorRT-LLM, LMCache, and llm-d.
- **KV connector landscape (vLLM, 2026):** **NixlConnector** — single cluster, RDMA/NVLink available; default high-performance P/D; metadata server is a startup SPOF. **LMCacheConnector** — cross-instance cache sharing, HBM→DRAM→NVMe tiering; NIXL under the hood + offload backends; tiered KV + shared prefix index (LMCache, arXiv:2510.09665). **MooncakeConnector** — cluster-scale shared cache pools; RDMA-native; separate KV-cache cluster pulled by many vLLM instances. **MooncakeStoreConnector** — tiered offload through a distributed master store.
- **Practitioner guidance** (terrytangyuan, 2026-06-05): treat KV transfer as a production data path — measure end-to-end latency including queue time, alert on tail latency, verify RDMA driver health per node; NIXL's async send/receive is correct only if prefill workers never block on decode-side acknowledgement.
- **LMCache bidirectional NIXL cache probe:** AWS DLC v1.1.0 (2026-04-28, vLLM 0.19.1) bundled LMCache 0.4.5.dev0 to enable disaggregated prefill/decode with bidirectional cache querying between prefill and decode workers. The framing winning mindshare: think of the KV cache as a **cluster-wide resource with its own scheduling concerns** — tiered HBM/DRAM/NVMe with a global prefix index so requests sharing a prefix share KV blocks regardless of which instance generated them.
- **2025-03:** NVIDIA Dynamo launched at GTC 2025 — inference orchestration (disaggregated P/D, KV-aware routing).
- **~2026-07:** **NVIDIA Dynamo 1.0** ships (6.6K GitHub stars, 281 contributors): disaggregated prefill/decode (scale each phase independently), KV-aware routing, multi-tier KV caching (GPU → CPU → SSD → remote storage), SLA-driven autoscaling. Vendor-claimed [VENDOR]: 7x higher throughput per GPU (DeepSeek R1 on GB200), 7x faster model startup (weight streaming), 2x faster TTFT (KV-aware routing), 80% fewer SLA breaches at 5% lower TCO. Named cases: Moonshot AI's Kimi K2 10x inference speedup on GB200; Mistral Large 3 10x faster inference; Dell PowerScale + NIXL 19x faster TTFT.
- **2026:** NVIDIA folds Triton into Dynamo ("Dynamo-Triton") — disaggregated serving and KV cache management on top of Triton.
- **2026:** **llm-d** (CNCF project; Red Hat, Google Cloud, IBM, NVIDIA, CoreWeave): Kubernetes-native orchestration turning vLLM/SGLang into a disaggregated system — cache-aware request routing (Endpoint Picker & Proxy), prefill/decode worker pools, NIXL KV transfer (default port 5600 since vLLM v0.15.1), tiered KV prefix caching (hot HBM → warm DRAM → cold SSD/shared FS via Lustre/LMCache/Mooncake), Workload Variant Autoscaler scaling prefill/decode pools independently and to zero. Also supports AMD ROCm, Intel XPU, Intel Gaudi/HPU. v0.5 benchmarks claim up to **57x faster TTFT and 2x throughput** vs round-robin under high prefix reuse (8 pods/16×H100) [VENDOR]; operators' own measurements more conservative (~25% over defaults, 2–3x tokens/s/GPU with prefix-cache-hit routing, 3–5x cost-per-token on chat-shaped workloads).
- **Dynamo vs llm-d philosophies (kaito comparison, 2026-09):** Dynamo — Rust+Python, custom discovery + etcd/NATS (Grove K8s operator optional), NVIDIA-optimized (NVLink/NVSwitch/GB200 NVL72), ModelExpress GPU-to-GPU weight streaming, AIConfigurator offline simulation of 10K+ configs, native multimodal E/P/D. llm-d — Go scheduler + Python, Kubernetes-native (Gateway API Inference Extension, EPP filter/scorer), K8s-first Helm, hardware-neutral (NVIDIA GPU, Intel XPU, Google TPU), vLLM primary/SGLang in progress, text-focused. The llm-d proposal explicitly contrasts: P/D decisions **in the scheduler** (not a Dynamo-style pipeline); **RPC rather than Dynamo's async queue** for stronger cancellation semantics; strong operational boundary between in-memory prefix-cache tiers and storage tiers vs Dynamo's unified KVBM memory API. **llm-d intends to reuse NIXL — Dynamo's own transfer library**: competition on orchestration, shared data plane.
- **2026:** Grove (Dynamo's Kubernetes operator — PodCliqueSets, hierarchical gang scheduling, topology-aware placement, multi-level autoscaling, startup ordering); AIBrix (vLLM project — StormService + RoleSet CRDs for high-density LoRA serving, gateway, autoscaling, P/D disaggregation); vLLM production-stack (Helm-based K8s deployment with request router and KV-aware routing).
- **2026:** AWS `awslabs/ai-on-eks` ships Dynamo blueprints for **vLLM, SGLang, and TensorRT-LLM** in aggregated and disaggregated modes; Dynamo's TRT-LLM guide demonstrates disaggregated gpt-oss-120b on a single B200 node (1 prefill worker on 4 GPUs + 1 decode worker on 4 GPUs).
- Practitioner view (2026): disaggregated serving enables **4–10x cost-per-token reduction** and is now the **default architecture for new production deployments** [COMMUNITY].
- **Two connector layers, don't confuse them:** vLLM's **KV connectors** (NixlConnector etc.) move KV blocks between prefill and decode workers inside one engine's disaggregated deployment; vLLM-Omni's **OmniConnector** is the higher-level abstraction routing data between heterogeneous *stages* (encoders, AR stages, DiT stages) with per-stage batching policies. llm-d's EPP (Endpoint Picker & Proxy) sits above both, making routing decisions across instances.

### Hardware tracks — dated facts

