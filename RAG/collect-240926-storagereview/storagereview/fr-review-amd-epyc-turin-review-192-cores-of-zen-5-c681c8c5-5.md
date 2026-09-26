---
id: collect-240926-storagereview/storagereview/fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5-5
title: "fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "decode", "distribution", "gpu", "inference", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5.md
source_anchor: ""
source_lines: [127, 143]
sha256: c500013e476884af807ccf470b6cf8591895b09aea1348182a3189a9b9ac96e5
---

# fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5

| Deep Lab V3 | 58.16 ms | 33.02 ms | 19.38 ms | 25.33 ms | 30.57 ms |  | 
| YOLO V3 | 84.20 ms | 38.47 ms | 24.54 ms | 34.13 ms | 41.38 ms |  | 
| REAL-ESRGAN | 2923.75 ms | 1600.73 ms | 1219.26 ms | 2524.03 ms | 2301.35 ms |  | 
| Overall Score (Higher is Better) | 44 | 81 | 148 | N/A | N/A |  | 
The Blackmagic RAW Speed Test performance benchmarking tool measures a system's ability to handle video playback and editing using the Blackmagic RAW codec. It evaluates a system's ability to decode and play high-resolution video files, providing frame rates for both CPU-based and GPU-based processing.
In this test, the EPYC 9755 leads with an impressive 174 fps in 8K CPU decoding, ideal for high-resolution video playback in media production environments. The 9575F (SMT disabled) follows closely with 154 fps, slightly surpassing the 9965 at 134 fps. These results suggest that the 9755 and 9575F offer the best balance between frame rate consistency and decoding speed, crucial for video editing and production workflows that handle large, high-quality video files.
| Blackmagic RAW (Higher is Better) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c) | 
| CPU 8K | 134 fps | 174 fps | 154 fps | 
The AMD EPYC Turin 9005 series marks a significant advancement in enterprise computing, offering exceptional performance, efficiency, and adaptability across various workloads. Powered by the Zen 5 architecture, these processors are specifically designed to meet the growing demands of AI, cloud, and HPC environments while delivering unmatched scalability for data centers looking to optimize performance and power consumption.
The EPYC 9005 series stands out for its ability to meet the diverse needs of modern enterprises with configurations ranging from entry-level high-frequency models to very high cache density chips to multi-core powerhouses. Whether it's real-time AI inference, computational fluid dynamics, large-scale data analytics, or high-resolution 3D rendering, the EPYC lineup offers single-thread responsiveness and multithread efficiency. Advanced features such as 12-channel DDR5 memory support, PCIe 5.0 lanes, and AMD's secure and confidential computing make this series both a performance upgrade and a comprehensive solution for forward-looking data centers.
Our testing with the Dell PowerEdge R7725 and R6725 platforms highlighted the true potential of these processors, particularly thanks to their board design and cooling solution, which offer a clear advantage in terms of bandwidth and overall performance. The combination of AMD's exceptional Zen 5 architecture and Dell's robust server design creates a compelling platform for businesses looking to push the boundaries of computing power.
The ASUS ExpertCenter Pro ET900N G3 is the second GB300 DGX station we have tested and the first we have been able to handle…
In the field of enterprise computing, most security upgrades generally occur in response to data breaches or other major incidents. The recently developed INCITS system…
Mini-PCs have become one of the most important segments of the client computing market. They form the basis of streamlined IT fleets…
The HP ZBook Ultra G1a 14 takes the top spot in our ranking of the best laptops for local AI, among the large models…
The most frequently asked question we received regarding the MSI XpertStation WS300 after our test revolves around a key theme. The…
Eaton is about to bring to market a type of power distribution unit (PDU) that shouldn't have existed five years ago. The HDXL…
