---
id: collect-240926-storagereview/storagereview/fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594-2
title: "fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["gpu", "nvidia", "benchmark", "benchmarks", "fp8", "gpus", "inference", "intel", "latency", "memory", "mlperf", "research"]
source: docs/RAG/clean_en/storagereview/fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594.md
source_anchor: ""
source_lines: [3, 91]
sha256: 7e5213213f4184e154433cd1ff5452c1503dba4f51adb9a33214e2b569867378
---

# fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594

In today's relentless torrent of innovation in the AI world, it is essential to measure and understand the capabilities of various hardware platforms. Not all AI requires massive GPU training farms; there is a significant segment of inference AI, which often requires less GPU power, especially at the edge. In this review, we examine several NVIDIA L4 GPUs, across three different Dell servers and various workloads, including MLPerf, to see how the L4 compares.
NVIDIA L4 GPU
At its core, the L4 offers an impressive 30.3 teraFLOP of FP32 performance, ideal for high-precision computing tasks. Its prowess extends to mixed-precision computations with TF32, FP16, and BFLOAT16 Tensor cores, crucial for deep learning efficiency; the L4 datasheet cites performance ranging from 60 to 121 teraFLOP.
In low-precision tasks, the L4 shines with 242.5 teraFLOP in FP8 and INT8 Tensor cores, thereby enhancing neural network inference. Its 24GB GDDR6 memory, complemented by 300 GB/s bandwidth, makes it capable of handling large datasets and complex models. The L4's power efficiency is most remarkable here, with a TDP of 72W making it suitable for various computing environments. This blend of high performance, memory efficiency, and low power consumption makes the NVIDIA L4 a must-have choice for cutting-edge computing challenges.
| NVIDIA L4 Specifications |  | 
|---|---|
| FP 32 | 30.3 teraFLOP | 
| TF32 Tensor Core | 60 teraFLOP | 
| FP16 Tensor Core | 121 teraFLOP | 
| BFLOAT16 Tensor Core | 121 teraFLOP | 
| FP8 Tensor Core | 242.5 teraFLOP | 
| INT8 Tensor Core | 242.5 TOPS | 
| GPU Memory | 24GB GDDR6 | 
| GPU Memory Bandwidth | 300GB / s | 
| Maximum Thermal Power (TDP) | 72W | 
| Form Factor | Low-profile PCIe single-slot | 
| Interconnect | PCIe Gen4x16 | 
| Specifications Table | L4 | 
Of course, with the L4 priced close to $2,500, the A2 costing about half the price, and the aging (but still quite performant) T4 available for less than $1,000, the obvious question is what the difference is between these three inference GPUs.
| NVIDIA L4, A2, and T4 Specifications | Nvidia L4 | Nvidia A2 | NVIDIA T4 | 
|---|---|---|---|
| FP 32 | 30.3 teraFLOP | 4.5 teraFLOP | 8.1 teraFLOP | 
| TF32 Tensor Core | 60 teraFLOP | 9 teraFLOP | N/A | 
| FP16 Tensor Core | 121 teraFLOP | 18 teraFLOP | N/A | 
| BFLOAT16 Tensor Core | 121 teraFLOP | 18 teraFLOP | N/A | 
| FP8 Tensor Core | 242.5 teraFLOP | N/A | N/A | 
| INT8 Tensor Core | 242.5 TOPS | 36 TOPS | 130 TOPS | 
| GPU Memory | 24GB GDDR6 | 16GB GDDR6 | 16GB GDDR6 | 
| GPU Memory Bandwidth | 300GB / s | 200GB / s | 320+ GB/s | 
| Maximum Thermal Power (TDP) | 72W | 40-60W | 70W | 
| Form Factor | Low-profile PCIe single-slot |  |  | 
| Interconnect | PCIe Gen4x16 | PCIe Gen4x8 | PCIe Gen3x16 | 
| Specifications Table | L4 | A2 | T4 | 
One thing to understand when looking at these three cards is that they are not exactly individual generational replacements, which explains why the T4 remains, many years later, a popular choice for certain use cases. The A2 came to replace the T4 as a lower-power and more compatible option (x8 vs x16 mechanical). Technically, the L4 then replaces the T4, with the A2 straddling an intermediate position that may or may not be refreshed at some point in the future.
MLPerf 3.1 Inference Performance
MLPerf is a consortium of AI leaders from academia, research, and industry, created to provide fair and relevant AI hardware and software benchmarks. These benchmarks are designed to measure the performance of machine learning hardware, software, and services across various tasks and scenarios.
Our tests focus on two specific MLPerf benchmarks: Resnet50 and BERT.
- Resnet50: This is a convolutional neural network used primarily for image classification. It is a good indicator of a system's ability to handle deep learning tasks related to image processing.
- BERT (Bidirectional Encoder Representations from Transformers): This benchmark focuses on natural language processing tasks, providing insight into how a system performs in understanding and processing human language.
These two tests are crucial for evaluating the capabilities of AI hardware in real-world scenarios involving image and language processing.
Evaluating the NVIDIA L4 with these benchmarks is essential to help understand the capabilities of the L4 GPU in specific AI tasks. It also provides insight into how different configurations (single, dual, and quad configurations) influence performance. This information is vital for professionals and organizations seeking to optimize their AI infrastructure.
The models run under two key modes: server and offline.
- Offline mode: This mode measures a system's performance when all data is available for simultaneous processing. This is similar to batch processing, where the system processes a large dataset in a single batch. Offline mode is crucial for scenarios where latency is not a major concern, but throughput and efficiency are.
- Server mode: In contrast, server mode evaluates the system's performance in a scenario mimicking a real server environment, where requests arrive one by one. This mode is latency-sensitive and measures how quickly the system can respond to each request. It is essential for real-time applications, such as web servers or interactive applications, where an immediate response is necessary.
1 x NVIDIA L4 – Dell PowerEdge XR7620
As part of our recent testing of the Dell PowerEdge XR7620, equipped with a single NVIDIA L4, we pushed it to its limits to run several tasks, including MLPerf.
The configuration of our test system included the following components:
- 2 Xeon Gold 6426Y – 16 cores 2.5 GHz
- 1 x Nvidia L4
- 8 x 16GB DDR5
- BOSS RAID480 1GB
- Ubuntu Server 22.04
- NVIDIA Driver 535
| Dell PowerEdge XR7620 1x NVIDIA L4 | Score | 
|---|---|
| Resnet50 – Server | 12,204.40 | 
| Resnet50 – Offline | 13,010.20 | 
| BERT K99 – Server | 898.945 | 
| BERT K99 – Offline | 973.435 | 
The performance in server and offline scenarios for Resnet50 and BERT K99 are almost identical, indicating that the L4 maintains consistent performance across different server models.
1, 2, and 4 NVIDIA L4 – Dell PowerEdge T560
The configuration of our review unit included the following components:
- 2 x Intel Xeon Gold 6448Y (32 cores/64 threads each, 225-watt TDP, 2.1-4.1 GHz)
- 8 x Solidigm P5520 1.6TB SSDs with PERC 12 RAID card
- 1 to 4x NVIDIA L4 GPUs
- 8 x 64GB RDIMM modules
- Ubuntu Server 22.04
- NVIDIA Driver 535
| Dell PowerEdge T560 1x NVIDIA L4 | Score | 
|---|---|
| Resnet50 – Server | 12,204.40 | 
| Resnet50 – Offline | 12,872.10 | 
| Bert K99 – Server | 898.945 | 
| Bert K99 – Offline | 945.146 | 
During our tests with two L4s in the Dell T560, we observed this near-linear performance scaling for the Resnet50 and BERT K99 benchmarks. This scaling demonstrates the efficiency of the L4 GPUs and their ability to operate in tandem without significant losses due to overhead or inefficiency.
| Dell PowerEdge T560 2x NVIDIA L4 | Score | 
|---|---|
| Resnet50 – Server | 24,407.50 | 
| Resnet50 – Offline | 25,463.20 | 
| BERT K99 – Server | 1,801.28 | 
| BERT K99 – Offline | 1,904.10 | 
The consistent linear scaling we witnessed with two NVIDIA L4 GPUs extends impressively to configurations with four L4 units. This scaling is particularly remarkable because maintaining linear performance gains becomes increasingly difficult with each added GPU due to the complexity of parallel processing and resource management.
| Dell PowerEdge T560 4x NVIDIA L4 | Score | 
|---|---|
| Resnet50 – Server | 48,818.30 | 
| Resnet50 – Offline | 51,381.70 | 
| BERT K99 – Server | 3,604.96 | 
| BERT K99 – Offline | 3,821.46 | 
These results are provided for informational purposes only and do not constitute official or competitive MLPerf results. To view the full list of official results, please visit the MLPerf results page.
