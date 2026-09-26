---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7-5
title: "fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7"
domain: storagereview
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Intel", "TensorRT-LLM"]
dates: []
keywords: ["benchmark", "compute", "decode", "deepseek", "distillation", "gpu", "gpus", "inference", "intel", "kv cache", "latency", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7.md
source_anchor: ""
source_lines: [74, 132]
sha256: e5910453085cdea80d3b09a2523359b2997fd25e268c97ecfaa87d69e42e5dc4
---

# fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7

TGI's benchmarking feature allows evaluation of its performance under different configurations and workloads. It provides a more accurate representation of real-world performance, as it accounts for the complexity of managing LLMs in a production environment.
Text generation using LLMs involves two main stages: prefill and decode. Prefill is the initial stage, where the LLM processes the input prompt to generate the necessary intermediate representations. This stage is compute-intensive, as it involves processing the entire input prompt in a single pass through the model.
During the prefill stage, the input prompt is tokenized and converted into a format usable by the LLM. The LLM then computes the KV cache, which stores information related to the input tokens. This KV cache is an essential data structure that facilitates the generation of output tokens.
In contrast, the decode stage is an autoregressive process where the LLM generates output tokens one by one, relying on the intermediate representations generated during the prefill stage. The decode stage relies heavily on the KV cache generated during the prefill stage, which provides the necessary context for generating coherent and contextually relevant output tokens.
Prefill stage
As batch size increases from 1 to 32, latency for all three models increases; DeepSeek-R1-Distill-Qwen-32B latency increases from 29.97 ms for a batch size of 1 to 76.95 ms for a batch size of 32. Similarly, GEMMA-3-27B-IT and Qwen/QwQ-32B latency increases from 51.84 ms and 29.90 ms to 79.58 ms and 76.30 ms, respectively.
In contrast, token throughput improves significantly with increasing batch size. For a batch of 1, throughput for the three models ranges from 192.95 to 334.46 tokens per second. For a batch of 32, they reach 4158.67, 4021.40, and 4194.13 tokens per second for DeepSeek-R1-Distill-Qwen-32B, GEMMA-3-27B-IT, and Qwen/QwQ-32B, respectively.
| LLM prefill stage performance: latency (ms) and token throughput (tokens/s) |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|
| Batch size | DeepSeek-R1-Distillation-Qwen-32B |  | GEMMA-3-27B-IT |  | Qwen/QwQ-32B |  | 
|---|---|---|---|---|---|---|
|  | Latency (ms) | Token rate | Latency (ms) | Token rate | Latency (ms) | Token rate | 
| 1 | 29.97 | 333.64 | 51.84 | 192.95 | 29.90 | 334.46 | 
| 2 | 30.21 | 662.09 | 52.55 | 380.61 | 29.95 | 667.80 | 
| 4 | 32.40 | 1234.72 | 52.62 | 760.12 | 32.12 | 1245.47 | 
| 8 | 36.98 | 2163.46 | 52.66 | 1519.19 | 36.69 | 2180.66 | 
| 16 | 51.63 | 3125.50 | 60.96 | 2624.64 | 51.29 | 3147.61 | 
| 32 | 76.95 | 4158.67 | 79.58 | 4021.40 | 76.30 | 4194.13 | 
Decode stage
Unlike the prefill stage, latency during the decode stage remains relatively stable regardless of batch size. For example, DeepSeek-R1-Distill-Qwen-32B latency varies from 27.14 ms to 29.52 ms as batch size increases from 2 to 32.
Token throughput during the decode phase improves with batch size, but not as dramatically as during the prefill phase. For a batch of 1, throughput is approximately 36-37 tokens per second for DeepSeek-R1-Distill-Qwen-32B and Qwen/QwQ-32B, and 33.96 tokens per second for GEMMA-3-27B-IT. For a batch of 32, throughput increases to 1083.83, 873.39, and 1084.89 tokens per second, respectively.
| LLM decode performance (token): latency (ms) and token throughput (tokens/s) |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|
| Batch size | DeepSeek-R1-Distillation-Qwen-32B |  | GEMMA-3-27B-IT |  | Qwen/QwQ-32B |  | 
|---|---|---|---|---|---|---|
|  | Latency (ms) | Token rate | Latency (ms) | Token rate | Latency (ms) | Token rate | 
| 1 | 27.24 | 36.71 | 29.45 | 33.96 | 27.24 | 36.71 | 
| 2 | 27.14 | 73.70 | 30.80 | 64.93 | 27.14 | 73.69 | 
| 4 | 27.50 | 145.46 | 31.33 | 127.65 | 27.47 | 145.62 | 
| 8 | 27.91 | 286.61 | 32.54 | 245.83 | 27.90 | 286.78 | 
| 16 | 28.31 | 565.07 | 34.71 | 460.92 | 28.44 | 562.56 | 
| 32 | 29.52 | 1083.83 | 36.64 | 873.39 | 29.50 | 1084.89 | 
This is normal, as the prefill stage computes the initial hidden states and key-value caches for the entire input prompt, which can saturate the GPU, as large batch operations can be executed simultaneously. After processing the prompt, the model generates new tokens, typically one at a time. At each step, the model uses the previous token and cached hidden states to produce the next token. Since this stage proceeds token by token, the batch size is often reduced, resulting in frequent GPU underutilization.
Procyon AI computer vision benchmark
Using concrete artificial vision tasks, the Procyon AI Computer Vision benchmark evaluates AI inference performance on CPUs, GPUs, and AI accelerators. It supports multiple inference engines such as TensorRT, OpenVINO, SNPE, Windows ML, and Core ML, providing insights into efficiency, compatibility, and optimization.
The Procyon AI Computer Vision benchmark results also demonstrate excellent AI inference performance. The system achieved low inference times, with MobileNet V3 at 20.64 ms and ResNet 50 at 22.42 ms. Inception V4 and DeepLab ran at 65.23 ms and 41.37 ms, respectively, efficiently handling more complex vision workloads. YOLO V3, a key object detection model, processed in 37.80 ms, making it particularly suitable for real-time AI applications. REAL-ESRGAN, a computationally intensive super-resolution model, recorded 1,159.22 ms, earning us an overall score of 81 in AI Computer Vision.
| AI Computer Vision (shorter time is better) (higher score is better) | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | 
|---|---|
| MobileNet V3 average inference time | 20.64 ms | 
| ResNet 50 average inference time | 22.42 ms | 
| Inception V4 average inference time | 65.23 ms | 
| DeepLab average inference time | 41.37 ms | 
| YOLO V3 average inference time | 37.80 ms | 
| REAL-ESRGAN average inference time | 1,159.22 ms | 
| AI Computer Vision overall score | 81 | 
Hammer DB TPROC-C
We also evaluated the performance of four popular open-source databases (MariaDB 11.4.4, MySQL 8.4.4, MySQL 5.7.44, and PostgreSQL 17.2) using the HammerDB TPROC-C benchmark to simulate OLTP workloads on 500 warehouses.
MariaDB emerged as the highest-performing solution, particularly in dual-socket configurations, where it scaled efficiently and achieved the highest transaction throughput. MySQL 8.4.4 showed notable improvements over the older 5.7.44 version, highlighting enhancements in recent releases. PostgreSQL 17.2 delivered consistent performance but lagged slightly behind MariaDB and MySQL 8.4.4. MariaDB delivered 3.15 million TPM on a single socket and 5.8 million TPM on two sockets, outperforming the others in both scenarios.
Performance comparison table (Transactions per minute, TPM)
| Database engine | Single-socket TPM | Dual-socket TPM | 
|---|---|---|
| MariaDB 11.4.4 | 3,150,000 | 5,800,000 | 
| MySQL 8.4.4 | 2,850,000 | 5,150,000 | 
| PostgreSQL 17.2 | 2,700,000 | 4,900,000 | 
| MySQL 5.7.44 | 2,300,000 | 4,250,000 | 
Despite the R770's hardware power, with its 86 cores per processor (a mix of high and low priority cores), no database recorded significant performance gains when spread across both sockets. This reflects the general preference of open-source databases for running on a single socket, due to better core locality and reduced memory latency.
Given these results, the R770 is better suited for running multiple database instances in a virtualized environment than for scaling a single instance. The system architecture is ideal for supporting a high-density mixed database workload, leveraging both performance and efficiency cores to ensure consistent throughput across multiple instances.
7-Zip
The built-in memory benchmark tool of the popular 7-Zip utility measures a system's CPU and memory performance during compression and decompression tasks, indicating how well the system can handle data-intensive operations.
