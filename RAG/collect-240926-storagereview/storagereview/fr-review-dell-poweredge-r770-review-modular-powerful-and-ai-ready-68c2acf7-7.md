---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7-7
title: "fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["benchmark", "compute", "cost", "decode", "gpu", "gpus", "intel", "nvidia"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7.md
source_anchor: ""
source_lines: [198, 219]
sha256: c46d0ddfe6b7f274e5015d4b6e98051c96644ab3584b424cc728881ae165a518
---

# fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7

The Geekbench 6 benchmark results show clear performance differences between the Dell PowerEdge R770 and the Lenovo ThinkSystem SR630 V4. In the single-core CPU test, Dell outperformed Lenovo with a score of 1,797, while Lenovo scored 1,173, a 53% improvement in single-core performance for Dell.
In the multi-core CPU test, Dell again dominated with 15,880 points, while Lenovo scored 13,868 points, giving Dell a 14% advantage in multi-core performance. This suggests that Dell's Intel Xeon 6787P processors offer superior overall compute power, particularly for tasks requiring multiple cores.
The OpenCL GPU test further highlighted Dell's advantage, with a score of 148,730 thanks to the NVIDIA L4 GPU.
| Geekbench 6 (Higher is better) | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 GB RAM) | 
|---|---|---|
| Single-core CPU | 1,797 | 1,173 | 
| Multi-core CPU | 15,880 | 13,868 | 
| GPU OpenCL score | 148,730 | (no GPU) | 
Blackmagic RAW Speed Test
The Blackmagic RAW Speed Test is a performance benchmarking tool designed to measure a system's ability to handle video playback and editing using the Blackmagic RAW codec. It evaluates a system's ability to decode and play high-resolution video files, providing frame rates for both CPU-based and GPU-based processing.
In the CPU-based test, the Dell PowerEdge R770 reached 141 fps, surpassing the Lenovo ThinkSystem SR630 V4, which scored 120 fps. This indicates that the Dell system handles CPU-based video processing more efficiently than the Lenovo. In the GPU-based test, the Dell PowerEdge R770 achieved 157 fps, benefiting from the presence of an NVIDIA GPU.
| Blackmagic RAW Speed Test (higher is better) | Dell PowerEdge R770 (2 Intel Xeon 6787P processors \| 2 TB RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 GB RAM) | 
|---|---|---|
| CPU FPS | 141 FPS | 120 FPS | 
| CUDA FPS | 157 FPS | 0 FPS (no GPU) | 
Blackmagic Disk Speed Test
The Blackmagic Disk Speed Test evaluates a disk's read and write speeds, assessing its performance, particularly for video editing tasks. It helps users ensure their storage is fast enough for high-resolution content, such as 4K or 8K video.
In the Blackmagic speed test, the Dell PowerEdge R770 BOSS card with mirrored SK hynix 480 GB Dell NVMe achieved a read speed of 3,010.3 MB/s and a write speed of 976.3 MB/s.
Conclusion
The Dell PowerEdge R770 particularly excites us, thanks to its adoption of the Open Compute Project's Data Center Modular Hardware System standard and its cutting-edge hardware. The integration of OCP DC MHS offers numerous benefits, including increased modularity, improved serviceability, and potential cost reductions through greater standardization. This design philosophy is reflected in every aspect of the system, from the implementation of iDRAC as OCP DC-SCM to the ports.
The R770 also offers impressive storage capabilities, supporting up to 40 E3.S drives in a single 2U chassis, making it the ideal solution for storage-intensive workloads. Additionally, the server's flexibility is enhanced by support for various configurations, including a cold-aisle-accessible front I/O configuration, providing greater flexibility to adapt to different data center layouts and maintenance requirements.
Compatible with a wide range of GPUs and Intel Xeon 6 Performance processors, the R770 is a powerful and versatile server platform, perfectly suited to the demands of modern data centers. Its cutting-edge hardware, modular design, and robust security features make the R770 an attractive option for organizations looking to deploy AI applications, high-performance computing (HPC), and traditional enterprise workloads.
