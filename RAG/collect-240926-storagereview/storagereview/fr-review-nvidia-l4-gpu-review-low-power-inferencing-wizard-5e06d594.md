---
id: collect-240926-storagereview/storagereview/fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594
title: "fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["gpu", "nvidia", "amd", "benchmark", "benchmarks", "compute", "fp8", "gpus", "inference", "intel", "latency", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594.md
source_anchor: ""
source_lines: [1, 174]
sha256: 55f1a5c6b8162a016b6d47756028ff70402be6653618a0606cc2ad12e39d0962
---

# fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594

<!-- source: https://www.storagereview.com/fr/review/nvidia-l4-gpu-review-low-power-inferencing-wizard -->

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
In addition to validating the linear scalability of NVIDIA L4 GPUs, our lab tests highlight the practical implications of deploying these units in different operational scenarios. For example, the consistency of performance between server and offline modes across all configurations with L4 GPUs reveals their reliability and versatility.
This aspect is particularly relevant for businesses and research institutions where operational contexts vary considerably. Furthermore, our observations on the minimal impact of interconnect bottlenecks and the efficiency of GPU synchronization in multi-GPU configurations provide valuable information for those seeking to scale their AI infrastructure. This information goes beyond simple benchmark numbers, offering a deeper understanding of how such hardware can be optimally used in real-world scenarios, guiding better architectural decisions and investment strategies in AI and HPC infrastructure.
NVIDIA L4 – Application Performance
We compared the performance of the new NVIDIA L4 to that of the NVIDIA A2 and NVIDIA T4 that preceded it. To showcase this performance improvement over previous models, we deployed all three models on a server in our lab, with Windows Server 2022 and the latest NVIDIA drivers, leveraging our full GPU test suite.
These cards were tested on a Dell Poweredge R760 with the following configuration:
- 2 x Intel Xeon Gold 6430 (32 cores, 2.1 GHz)
- Windows Server 2022
- NVIDIA Driver 538.15
- ECC disabled on all cards for 1x sampling
As we launch performance tests between this group of three enterprise GPUs, it is important to note the unique performance differences between the previous A2 and T4 models. When the A2 was released, it offered notable improvements such as lower power consumption and operation on a smaller PCIe Gen4 x8 slot, instead of the larger PCIe Gen3 x16 slot required by the older T4. From the outset, this allowed it to fit into more systems, particularly with the smaller footprint required.
Blender OptiX 4.0
Blender OptiX is an open-source 3D modeling application. This test can be run for both CPU and GPU, but we only performed the GPU test like most other tests here. This benchmark was run using the Blender Benchmark CLI utility. The score is expressed in samples per minute, with the highest being the best.
| Blender 4.0 (Higher is Better) | Nvidia L4 | Nvidia A2 | Nvidia T4 | 
|---|---|---|---|
| Blender GPU CLI – Monster | 2,207.765 | 458.692 | 850.076 | 
| Blender GPU CLI – Junkshop | 1,127.829 | 292.553 | 517.243 | 
| Blender GPU CLI – Classroom | 1,111.753 | 262.387 | 478.786 | 
Blackmagic RAW Speed Test
We test CPUs and GPUs with Blackmagic's RAW Speed Test, which tests video reading speeds. This is more of a hybrid test including both CPU and GPU performance for actual RAW decoding. These are displayed as separate results, but we are focusing only on the GPUs here, so the CPU results are omitted.
| Blackmagic RAW Speed Test (Higher is Better) | Nvidia L4 | Nvidia A2 | NVIDIA T4 | 
|---|---|---|---|
| CUDA 8K | FPS 95 | FPS 38 | FPS 53 | 
Cinebench 2024 GPU
Maxon's Cinebench 2024 is a CPU and GPU rendering benchmark that utilizes all processor cores and threads. Again, since we are focusing on GPU results, we did not run the CPU portions of the test. Higher scores are better.
| Cinebench 2024 (Higher is Better) | Nvidia L4 | Nvidia A2 | NVIDIA T4 | 
|---|---|---|---|
| GPU | 15,263 | 4,006 | 5,644 | 
GPU PI
GPUPI 3.3.3 is a version of the lightweight benchmarking utility designed to calculate π (pi) to billions of decimal places using hardware acceleration via GPUs and CPUs. It leverages the computing power of OpenCL and CUDA, which includes central and graphics processing units. We ran CUDA only on the 3 GPUs, and the figures here are the computation time without added reduction time. Lower is better.
| GPU PI Computation Time in Seconds (Lower is Better) | Nvidia L4 | Nvidia A2 | NVIDIA T4 | 
|---|---|---|---|
| GPUPI v3.3 – 1B | 3.732s | 19.799s | 7.504s | 
| GPUPI v3.3 – 32B | 244.380s | 1,210.801s | 486.231s | 
While the previous results only covered a single iteration of each card, we also had the opportunity to examine a deployment of 5 NVIDIA L4 cards inside the Dell PowerEdge T560.
| GPU PI Computation Time in Seconds (Lower is Better) | Dell PowerEdge T560 (2x Xeon Gold 6448Y) with 5x NVIDIA L4 | 
|---|---|
| GPUPI v3.3 – 1B | 0 s 850 ms | 
| GPUPI v3.3 – 32B | 50 s 361 ms | 
OctaneBench
OctaneBench is a benchmarking utility for OctaneRender, another 3D rendering engine with RTX support similar to V-Ray.
| Octane (Higher is Better) |  |  |  |  | 
| Scene | Kernel | Nvidia L4 | Nvidia A2 | NVIDIA T4 | 
| Interior | Info channels | 15.59 | 4.49 | 6.39 | 
|  | Direct lighting | 50.85 | 14.32 | 21.76 | 
|  | Path tracing | 64.02 | 18.46 | 25.76 | 
| The Idea | Info channels | 9.30 | 2.77 | 3.93 | 
|  | Direct lighting | 39.34 | 11.53 | 16.79 | 
|  | Path tracing | 48.24 | 14.21 | 20.32 | 
| ATV | Info channels | 24.38 | 6.83 | 9.50 | 
|  | Direct lighting | 54.86 | 16.05 | 21.98 | 
|  | Path tracing | 68.98 | 20.06 | 27.50 | 
| Skybox | Info channels | 12.89 | 3.88 | 5.42 | 
|  | Direct lighting | 48.80 | 14.59 | 21.36 | 
|  | Path tracing | 54.56 | 16.51 | 23.85 | 
| Total Score |  | 491.83 | 143.71 | 204.56 | 
Geekbench 6 GPU
Geekbench 6 is a cross-platform evaluation tool that measures the overall performance of a system. It offers tests for the CPU and graphics card. The higher the score, the better the performance. We focused here on the graphics card results.
You can find comparisons with any system in the Geekbench browser.
| Geekbench 6.1.0 (Higher is Better) | Nvidia L4 | Nvidia A2 | NVIDIA T4 | 
|---|---|---|---|
| GeekbenchGPU OpenCL | 156,224 | 35,835 | 83,046 | 
LuxMark
LuxMark is a cross-platform OpenCL benchmarking tool designed by those who maintain the open-source 3D rendering engine LuxRender. This tool examines GPU performance in 3D modeling, lighting, and video work. For this review, we used the latest version, v4alpha0. In LuxMark, the higher the score, the better.
| Luxmark v4.0alpha0 GPU OpenCL (Higher is Better) | Nvidia L4 | Nvidia A2 | NVIDIA T4 | 
|---|---|---|---|
| Entry Bench | 14,328 | 3,759 | 5,893 | 
| Food Bench | 5,330 | 1,258 | 2,033 | 
GROMACS CUDA
We also sourced GROMACS, a molecular dynamics software, specifically for CUDA. This custom compilation was intended to leverage the parallel processing capabilities of the 5 NVIDIA L4 GPUs, essential for accelerating computational simulations.
The process involved using nvcc, NVIDIA's CUDA compiler, along with numerous iterations of appropriate optimization flags to ensure the binaries were properly suited to the server's architecture. The inclusion of CUDA support in the GROMACS compilation allows the software to interface directly with the GPU hardware, which can significantly improve computation times for complex simulations.
The test: Custom protein interaction in Gromacs
Leveraging a community-provided input file from our various Discord, which contained parameters and structures tailored to a specific protein interaction study, we launched a molecular dynamics simulation. The results were remarkable: the system achieved a simulation rate of 170.268 nanoseconds per day.
| GPU | System | ns/day | base time (s) | 
|---|---|---|---|
| Nvidia A4000 | AMD Ryzen 5950x Whitebox | 84.415 | 163,763 | 
| NVIDIA RTX 4070 | AMD Ryzen 7950x3d Whitebox | 131.85 | 209,692.3 | 
| 5x Nvidia L4 | Dell T560 with 2x Intel Xeon Gold 6448Y | 170.268 | 608,912.7 | 
More than AI
With the AI hype raging, it's easy to get caught up in the performance of models on the NVIDIA L4, but it also has a few other aces up its sleeve, opening up a range of possibilities for video applications. It can host up to 1,040 simultaneous AV1 720p30 video streams. This can transform how content can be streamed live to edge users, enhance creative storytelling, and present interesting uses for immersive AR/VR experiences.
The NVIDIA L4 also excels in optimizing graphics performance, as evidenced by its real-time rendering and ray tracing capabilities. In an edge office, the L4 is capable of delivering robust and powerful graphics compute acceleration in VDI to end users who need it most when high-quality real-time graphics rendering is essential.
Closing Thoughts
The NVIDIA L4 GPU provides a solid platform for cutting-edge AI and high-performance computing, offering unmatched efficiency and versatility across multiple applications. Its ability to handle intensive AI, acceleration, or video pipelines and optimize graphics performance makes it an ideal choice for edge inference or virtual desktop acceleration. The L4's combination of high compute power, advanced memory capabilities, and power efficiency positions it as a key player in accelerating edge workloads, particularly in AI and graphics-intensive industries.
There is no doubt that AI is the eye of the computing hurricane these days, and the demand for monstrous H100/H200 GPUs continues to explode. But there is also a major effort to put a more robust set of computing kits towards the edge, where data is created and analyzed. In these cases, a more appropriate GPU is needed. Here, the NVIDIA L4 excels and should be the default option for edge inference, either as a single unit or scaled globally, as we tested in the T560.
