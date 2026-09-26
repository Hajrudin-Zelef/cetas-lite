---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-xe7740-inside-the-architecture-of-enterprise-ai-inferen-a8c6e9c2-4
title: "fr-review-dell-poweredge-xe7740-inside-the-architecture-of-enterprise-ai-inferen-a8c6e9c2"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Meta", "SGLang", "TSMC", "vLLM"]
dates: []
keywords: ["accelerator", "benchmarks", "compute", "cost", "decode", "ethernet", "fp8", "hbm", "inference", "intel", "kv cache", "latency"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-xe7740-inside-the-architecture-of-enterprise-ai-inferen-a8c6e9c2.md
source_anchor: ""
source_lines: [67, 94]
sha256: f9f8aa9ab461568627b101bba4277e6b9f1bc8e3836d9ea8b702caa4a3926620
---

# fr-review-dell-poweredge-xe7740-inside-the-architecture-of-enterprise-ai-inferen-a8c6e9c2

It is also worth mentioning the AMX tensor units of the Xeon 6, which handle a significant share of the processor-side work. These include preprocessing, tokenization, and hybrid inference tasks involving matrix operations. This proves particularly useful with inference frameworks like SGLang, which use the processor for radix tree KV cache management and scheduling without overhead.
Intel Gaudi 3 expansion cards: Competitive inference at scale
Intel's Gaudi 3 is the company's flagship AI accelerator, launched in the fourth quarter of 2024. Intel positions these accelerators very aggressively, rather than competing head-on with high-end data center training accelerators. The Gaudi 3 clearly targets the inference segment.
Inference of transformer-based models, in all current popular LLMs, is fundamentally memory-bound. During the decoding phase of autoregressive generation, the model generates tokens one by one, reading the model weights and KV cache entries for each token produced. The bottleneck is not in compute power, but in memory bandwidth, that is, the speed at which the accelerator can transfer data from HBM memory to the compute engines.
The Gaudi 3 carries 128 GB of HBM2e memory providing 3.7 TB/s of bandwidth. Its architecture is based on TSMC's 5 nm process and uses a dual-die design: two identical silicon dies connected by a high-bandwidth interconnect, presenting themselves as a single device to software. Compute power is organized into four deep learning cores (DCORE), each comprising 2 matrix multiplication engines (MME), 16 transactional compute units (TPC), and 24 MB of local SRAM cache. The 96 MB of integrated SRAM provides a total internal bandwidth of 12.8 TB/s. The accelerator also integrates 14 dedicated media decoders (H.265, H.264, JPEG, VP9), enabling fast image preprocessing for multimodal workloads.
A large portion of the cutting-edge open-source AI models released today are either natively trained in FP8 or are hybrid models combining FP8 (E4M3) and BF16 weights. The Gaudi 3 offers native FP8 acceleration for these models thanks to its 8 matrix multiplication engines and 64 tensor processing cores, delivering an FP8 compute power of 1.8 PFlops.
The Gaudi 3 also integrates RDMA over Converged Ethernet (RoCEv2) technology with 24 200 GbE ports on the OAM version, directly integrated into the chip. Although the PCIe expansion card used in the XE7740 does not expose all these ports in the same way, the various expansion card versions allow up to four cards to be connected for faster communication.
Performance and benchmarks
XE7740 configuration details:
- 2 Intel Xeon 6787P processors (86 cores, 2.00 GHz)
- 2 TB DDR5 (32 x 64 GB DDR5 at 5,200 MT/s)
- 4 Intel Gaudi 3 PCIe AI accelerators with 128 GB HBM memory
- Ubuntu 24.04.5 server
vLLM Online Serving Performance
To evaluate the inference capabilities of the Dell XE7740 server equipped with Intel Gaudi 3 accelerators, we compared the online serving performance of vLLM across a range of common models, covering different architectures, parameter counts, and precision formats. Each model was tested across three workload profiles, with the number of concurrent requests ranging from 1 to 128.
LLM inference consists of two distinct phases. The prefill phase processes all input tokens in parallel before any output token generation, making it a compute-intensive operation proportional to the number of input tokens. The decode phase then generates output tokens one by one (autoregressively). Each new token requires reading the entire model weights from memory, but the computation per token is relatively low, making it memory-bandwidth-bound.
These two phases stress fundamentally different parts of the accelerator, so we test three workload profiles that shift the balance between them:
- Equal (1024 input tokens/1024 output tokens) represents balanced chat interactions.
- Prefill Heavy (8192 inputs/1024 outputs) simulates retrieval-augmented generation or long-context summarization, in which the system must process large input contexts.
- Decode Heavy (1024 inputs/8192 outputs) represents long-form content generation where sustained memory bandwidth determines throughput.
In this section, we focus on two main metrics: total token throughput (in tokens per second), which reflects the system's overall capacity under load, and time to first token (TTFT), which measures the elapsed time between submitting a request and receiving the first generated token. Since the model must complete the prefill phase before emitting the first token, TTFT is directly related to the accelerator's compute power. Thus, a prefill-heavy scenario (combined with TTFT) is a particularly relevant indicator of the raw compute capabilities of Gaudi 3 accelerators, since the system must process all 8,192 input tokens before the user receives a response.
Conversely, the decode-heavy scenario tests the memory bandwidth of the accelerators, since the system must sustain high throughput for thousands of generated tokens. TTFT is crucial for interactive applications where users wait for a response before streaming begins. A system can achieve excellent throughput with large batch processing but remain slow if TTFT becomes too high; both metrics are therefore important for production deployments.
Note regarding FP8 precision results: although Intel Gaudi 3 accelerators include native FP8 compute acceleration (and FP8 should, in theory, offer higher throughput than BF16), the FP8 performance measured in our tests is lower than that obtained with BF16. This is not a hardware limitation, but rather a software maturity issue within Intel's vLLM build. The version we tested (vLLM 2.7.1 installer on Gaudi Docker 1.22.2) has not yet fully optimized its FP8 execution paths. Intel offers a new plugin-based vLLM version, currently in beta, which could resolve many of these performance issues.
Lama 3.1 8B Instruct
Llama 3.1 8B Instruct is a dense transformer model from Meta, meaning that every parameter is active for each generated token. With 8 billion parameters, it is among the most widely used open-source models. Models of this size are commonly used for everyday tasks such as summarizing short documents, drafting emails and messages, answering simple questions, and powering basic chatbot interactions where speed and cost-effectiveness take priority over reasoning complexity.
We tested this model with two configurations: TP1 (a single accelerator) and TP4 (all four Gaudi 3 accelerators). On TP1, the model reaches a total throughput of approximately 8,000 tok/s for 128 concurrent requests at equal workload, increasing from approximately 250 tok/s for a single request to a much higher throughput. The prefill-heavy scenario reveals interesting behavior: while TP1 caps at approximately 7,000 tok/s, TP4 exceeds 17,900 tok/s for 128 concurrent requests, thanks to the use of additional accelerators to process the large input context more efficiently.
In terms of latency for a single user, TP1 shows lower TTFT under low concurrency (67 ms versus 98 ms for TP4), reflecting the overhead of coordinating across four accelerators for a model that easily fits on one. However, as load increases, TP4 takes a clear advantage. With 128 concurrent requests, TP4 maintains TTFT at approximately 2 seconds for equal and decode workloads, while TP1 reaches 3.7 seconds and 6.6 seconds respectively. It is in the prefill-intensive scenario that the gap is most pronounced: TP1 reaches nearly 47 seconds of TTFT at 128 requests, while TP4 maintains it at approximately 11 seconds.
Lama 3.1 70B Instruct
