---
id: collect-240926-storagereview/storagereview/fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594-3
title: "fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["gpu", "nvidia", "amd", "benchmark", "gpus", "intel", "parameters", "protein", "research"]
source: docs/RAG/clean_en/storagereview/fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594.md
source_anchor: ""
source_lines: [92, 169]
sha256: b26c54ff32d4e16c9f7eb70a3c2a4cfd3fa2f8e6979424e514bde1da121e5b35
---

# fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594

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
