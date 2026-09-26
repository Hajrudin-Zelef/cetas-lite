---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb-4
title: "fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["energy", "gpus", "inference", "intel"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb.md
source_anchor: ""
source_lines: [124, 137]
sha256: dd9b8e74ce06ef839ab0d1e43b0e19636508488ddf11000a7ef5d107ad58e191
---

# fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb

UL Procyon AI Inference is designed to evaluate a workstation's performance in professional applications. It is important to note that this test does not leverage the capabilities of multiple processors. Specifically, this tool evaluates the workstation's ability to handle AI-based tasks and workflows, providing a detailed analysis of its efficiency and speed in processing complex AI algorithms and applications.
The results for the Supermicro Hyper 112h-tn (Xeon 6780E, 512 GB DDR5) provide insight into the system's ability to handle various AI-driven tasks. From the table below, it is clear that the server can efficiently process simpler models like MobileNet V3 and ResNet 50, with inference times of 5.23 ms and 6.60 ms, respectively. However, as model complexity increases, such as with Real-ESRGAN, the inference time increases significantly to 966.48 ms. The overall score of 168 indicates that the Xeon 6780E is capable but may be better suited for moderate AI workloads and would benefit from a dual-processor configuration.
| UL Procyon Average Inference Times (lower is better) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | 
| Mobile Net V3 | 5.23 ms | 
| ResNet 50 | 6.60 ms | 
| Inception V4 | 23.12 ms | 
| Deep Lab V3 | 24.61 ms | 
| YOLO V3 | 35.66 ms | 
| Real-ESRGAN | 966.48 ms | 
| Overall Score | 168 | 
Conclusion
Powered by the Intel Xeon 6 series, the Supermicro Hyper SuperServer SYS-112H-TN offers a compelling combination of performance and efficiency, making it a serious contender for various mainstream enterprise applications. Its equipped 6780E processor allows it to excel in multithreaded workloads, delivering robust performance while maintaining energy efficiency. The efficiency gains are particularly impressive compared to the older systems they are likely to replace.
Beyond the capabilities of the Xeon 6 processors, the Supermicro Hyper 1U SYS-112H-TN is equipped with a range of features that enhance its versatility and performance. It features two PCIe 5.0 x16 slots and an additional PCIe 5.0 AIOM slot, enabling high-bandwidth connectivity for GPUs, network interfaces, and other expansion cards.
Xeon 6 offers up to 144 efficiency cores in the Supermicro Hyper 1U SYS-112H-TN, which is excellent for workloads that don't need the most powerful silicon. It should serve as a flexible building block for system integrators and businesses looking for an affordable yet modern server for a wide variety of tasks.
