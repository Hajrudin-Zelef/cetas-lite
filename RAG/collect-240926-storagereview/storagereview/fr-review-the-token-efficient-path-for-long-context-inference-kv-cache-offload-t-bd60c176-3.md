---
id: collect-240926-storagereview/storagereview/fr-review-the-token-efficient-path-for-long-context-inference-kv-cache-offload-t-bd60c176-3
title: "fr-review-the-token-efficient-path-for-long-context-inference-kv-cache-offload-t-bd60c176"
domain: storagereview
role: reference
task: reference
actors: ["MiniMax", "Nvidia", "vLLM"]
dates: []
keywords: ["inference", "agent", "agentic", "agents", "blackwell", "compute", "context window", "decode", "disaggregated", "disaggregated serving", "dram", "gpu"]
source: docs/RAG/clean_en/storagereview/fr-review-the-token-efficient-path-for-long-context-inference-kv-cache-offload-t-bd60c176.md
source_anchor: ""
source_lines: [23, 52]
sha256: 5377767cbf47a157b722fa0973b2a20131e146affcbfc27cd4b855b08c4ccee4
---

# fr-review-the-token-efficient-path-for-long-context-inference-kv-cache-offload-t-bd60c176

In a workload where 98% of traffic consists of cache reads, eviction without offloading involves prefilling most of the work the system was about to perform. The bottleneck shifts from decode bandwidth, a phase that is generally optimized, to prefill computation, already performed by the GPU in previous iterations. In fact, it is common to observe, in disaggregated serving architectures, a greater number of prefill processes than decode processes at scale, for the same reason: prefill is the control phase.
The 98% figure represents the traffic of a single user. A short chat workload would generate more cold traffic. The same would be true for a retrieval-augmented system that reinjects different documents at each exchange. This general trend holds when conversations are long, prefixes are stable, and users regularly return to the same context.
Where is the cache and what happens when it is full?
Where exactly is this cache located, and what happens when it is full? Once the model weights are loaded, the remaining VRAM is converted into KV cache space, partitioned into fixed-size blocks. At startup, the engine reports the number of tokens it can hold. This pool is the only possible location for the cache in a standard configuration without KV offloading.
GPU servers are expensive and tokens are the product, so the goal is to feed them continuously. This involves keeping the GPU loaded with a small queue of requests to process in order to avoid idle compute resources, while balancing this queue so that response latency remains in line with service level objectives (SLOs). Not all requests use their full context window; many completed key-value (KV) pairs therefore remain cached in VRAM. However, under sustained load, this cache fills up and, once saturated, the oldest entries are removed to free up space.
In an ideal world, each response would be processed in a single pass. But models are not yet capable enough; we therefore use inter-agent loops, with exchanges with a user or between agents. When the next turn arrives, if its key-value identifier is not yet in the cache, the engine reprocesses the entire conversation, including the new tokens. The context of each previous turn is reprocessed by the model, and the computation already performed by the GPU on those tokens is paid for a second, third time, and so on, for each subsequent turn. This part of the process represents a waste of resources.
What changes does KV cache offloading bring?
Instead of evicting data as soon as VRAM is full, offloading gradually distributes caches across a hierarchy: from VRAM to system memory, then to SSD. Frequently used data remains in VRAM; data that no longer fits is transferred to system memory, and if that fills up, it is in turn transferred to SSD or network storage. Actual data eviction becomes rare, because it is driven by a user-defined time-to-live (TTL) rather than memory pressure.
On the next turn, instead of recomputing tens of thousands of context tokens, we retrieve the key-value pair from RAM or flash storage. Reloading from system memory introduces a slight latency, but is far faster and more economical than recomputing the prefill. Reloading from an NVMe SSD is a little slower than a DRAM read, but remains much faster than the recomputation it replaces. Storage capacity is relatively inexpensive; GPU compute power, on the other hand, is not. Therefore, accepting a slight reload latency at the start of a turn to avoid a multi-second recomputation is overall advantageous.
How we tested
We tested this hypothesis in our lab. The system configuration was as follows:
- Server: Dell PowerEdge XE7740
- GPU: 4x NVIDIA RTX PRO 6000 Blackwell Server Edition (96 GB)
- System memory: 1 TB DDR5 (16 x 64 GB DDR5 5200 MT/s)
- Storage: 8 Solidigm PS1030 12.8 TB E3.S drives (RAID 10)
- Server stack: vLLM 0.22.0 with LMCache 0.5.0
- Model: MiniMax-M2.7
The test platform is perfectly suited to this workload. The Dell PowerEdge XE7740 server is designed specifically for enterprise AI inference. This PCIe Gen5 chassis supports up to eight double-width GPUs. The four-GPU configuration we tested is one of the most popular. It offers sufficient acceleration capacity for a wide range of inference deployments, while allowing expansion to eight GPUs based on demand. Each NVIDIA RTX PRO 6000 Blackwell Server Edition card has 96 GB of GDDR7 memory; thus, four cards constitute a substantial pool of VRAM before the cache is even stressed. Internally, the offload tier relies on eight Solidigm D7-PS1030 12.8 TB drives in RAID 10. These high-endurance Gen5 flash drives are perfectly suited to the sustained writes generated by a KV cache tier. Finally, the choice of model, MiniMax-M2.7, was, at the time of testing, one of the best open coding models that integrated with our four-GPU configuration.
We tested three configurations:
- A baseline VRAM configuration, standard vLLM without offloading, where everything that fits in VRAM is cached, and everything else is evicted.
- LMCache offloading to system memory, with 512 GB allocated to offloading.
- LMCache offloading to flash memory, with the local NVMe RAID10 array as the offload tier, preceded by a 64 GB RAM transit buffer.
We modeled the workload based on real agentic coding traffic rather than a sweep of fixed synthetic prefixes. This traffic is characterized by significant prefill and minimal decoding, with intensive prefix reuse. Sessions arrive according to a Poisson process at a rate of λ = 2.5 sessions per minute. Each session runs over multiple turns. At each turn, the model receives a short addition, typically the output of a tool or command of 250 to 600 tokens, and occasionally the reading of a file of 1,500 to 3,500 tokens. It responds with a brief reply, typically 40 to 200 tokens of tool calls or brief reasoning, and occasionally a code block of 400 to 900 tokens. Turn lengths vary by ±100 tokens. Sessions grow monotonically up to a context ceiling of 64,000 tokens, placing them in the deep context regime, where the working set exceeds VRAM capacity. The same initial load feeds all three tiers, each with a 3-minute cache warmup time.
Note regarding the memory configuration: the XE7740 shipped with 2 TB of DDR5, a high-end configuration designed to cover a wide range of projects, and whose price predated the 2026 memory price surge, which made that amount of DRAM much more expensive. The original configuration was oversized relative to the current needs of an enterprise equipped with four GPUs. To make it more realistic, we removed 1 TB of memory, keeping one DIMM per channel to preserve maximum memory bandwidth.
Methodological note regarding the flash tier: the tests used LMCache's local disk backend with its required 64 GB RAM transit buffer upstream of the array. The KV footprint retained on flash memory far exceeded the total capacity of the host DRAM during the tests; the results obtained therefore reflect disk activity rather than host memory.
Performance
Throughput under load
Total service throughput over time, as the KV working set exceeds the capacity of each tier:
Each memory tier ramps up simultaneously during the first 10 minutes, while the caches fill, with no notable difference in performance. Then, as the memory tiers fill, they diverge. The VRAM-only baseline stalls first: once VRAM is saturated, it stabilizes around 12,000 tokens per second. Each new session replaces a previous session, and the displaced cache must be entirely rebuilt when that session returns. Consequently, a greater portion of time is spent on prefill and less on decoding.
The performance of the DRAM and SSD servers continues to far exceed this threshold because, as long as their caches are not saturated, they process nearly all prefixes directly from memory or disk, without having to recompute them. The time usually spent on prefill is thus reinvested in generating new tokens.
