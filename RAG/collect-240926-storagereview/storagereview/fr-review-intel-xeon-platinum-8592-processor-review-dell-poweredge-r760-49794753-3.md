---
id: collect-240926-storagereview/storagereview/fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753-3
title: "fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["intel", "amd", "benchmark", "compute", "inference", "latency", "memory", "parameters", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753.md
source_anchor: ""
source_lines: [94, 124]
sha256: b1c3569c37112e6561dec93b5dba472a987cc1c9537cd392ba10a72da93747c4
---

# fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753

| Resulting CPU Usage | 6,232 % | 5,894 % | 6,129 % | 
| Resulting Rating/Usage | 3.654 GIPS | 3.385 GIPS | 3.871 GIPS | 
| Resulting Rating | 227.744 GIPS | 199.363 GIPS | 237.259 GIPS | 
| Total Rating |  |  |  | 
| Total CPU Usage | 5,919 % | 5,781 % | 5,802 % | 
| Total Rating/Usage | 4.269 GIPS | 3.654 GIPS | 4.249 GIPS | 
| Total Rating | 250.730 GIPS | 210.363 GIPS | 245.320 GIPS | 
UL Procyon AI Inference
UL's Procyon AI inference test suite evaluates the performance of different AI inference engines using state-of-the-art neural networks. These tests were performed on the CPU only. Each value represents an average inference time (the shorter the time, the better the performance), and the last row indicates an overall score (the higher the score, the better the performance). The differences between these processors are less pronounced in this test: the Xeon 8480+ even stands out in some tests, but in the end, the Xeon 8592+ still achieves a higher overall score.
|  | 2 Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 
|---|---|---|---|
| Mobile Net V3 | 3.17 | 3.71 | 2.34 | 
| ResNet 50 | 5.23 | 4.37 | 5.76 | 
| Creation V4 | 18.66 | 22.59 | 21.70 | 
| Deep Lab V3 | 23.99 | 28.25 | 23.00 | 
| YOLO V3 | 35.47 | 40.28 | 30.81 | 
| REAL-ESRGAN | 1021.49 | 1277.08 | 1535.27 | 
| Overall Score | 196 | 161 | 191 | 
CPU-Z
Although it is not exactly a benchmark, we thought it was worth mentioning the memory speed differences, with the only change being the installed processor.
Here first is a photo of the memory tab of the Sapphire Rapids Xeon 6430 that were initially installed in our R760.
Here next is a photo of the memory tab of the Emerald Rapids Xeon 8592+.
There is a clear difference in memory timing on the Xeon 8592+, making it slightly inferior to that of our Xeon 6430, but probably not enough to impact performance in most real-world applications.
Memory Test
By analyzing the memory information provided by CPU-Z, we can run microbenchmarks from Clamchowder. This tool is designed to evaluate processor- and memory-related parameters, including ROB and register file sizes, lock and cache coherence latency, as well as cache and memory performance. We chose this test to analyze performance gains between different processor generations and for different access sizes.
For this test, we used 64 threads and ran a full profile from 64 KB to 3,145,728 5 XNUMX KB. By focusing on performance once we leave the processor cache and access system memory, we can see a significant increase in throughput with the 5th generation Xeon in all areas during the write test.
As for read bandwidth, the story is not as dramatic, but still notable. The clock speed difference between the 4th generation processor and the 5th generation processor we tested becomes a little clearer here in this test. Despite the slight disadvantage of the 5th generation, we can still see that the 5th generation Xeon Scalable can deliver a reduction in performance increase.
Conclusion
As expected, the Intel Xeon Platinum 8592+ processor showed significantly higher performance than the Xeon 6430 in most tests, which makes sense since they are not direct competing processors. By including the Xeon 8480+, the performance gap narrows considerably, the latter being a direct predecessor of the Xeon 8592+. A major advantage of the 5th generation Xeon processors is their compatibility with the same socket as the 4th generation, thus allowing older platforms to be upgraded on-site, as we did for some of our servers.
We found significant performance improvements across all our workloads between the two 4th generation Intel Xeon processors and the 5th generation Xeon 8592+. The largest improvements were visible in our compute workloads, such as y-cruncher. During our 25 billion digit runs, we measured times on the Gold 6430 and Platinum 8480+ of 229 and 187 seconds respectively, dropping to only 157 seconds on the Platinum 8592+. Performance was also significantly improved in Cinebench R23 and 2024, going from the 4th generation Platinum 8480+ to the 5th generation 8592+. In R23, we saw performance go from 79 110 to 2024 4,699 km, and in 6,001, speeds went from XNUMX XNUMX to XNUMX XNUMX.
The Intel Xeon Platinum 8592+ marks a big step forward for the enterprise and even if it cannot go core to core with AMD's robust chips, the 8592+ has unique strengths that make it a compelling choice. Intel has significantly optimized these processors for AI and HPC workloads, delivering robust performance at all levels.
