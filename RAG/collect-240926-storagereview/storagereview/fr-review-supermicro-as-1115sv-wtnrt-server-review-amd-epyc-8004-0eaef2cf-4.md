---
id: collect-240926-storagereview/storagereview/fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf-4
title: "fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "benchmark", "cost", "exploit", "inference", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf.md
source_anchor: ""
source_lines: [146, 181]
sha256: 995464b37d99cffca5361eb3131bf85472995cb5e1a32124bab9b3ced7358462
---

# fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf

UL Procyon AI Inference is designed to evaluate a workstation's performance in professional applications. It is important to note that this test does not exploit the capabilities of multiple processors. Specifically, this tool evaluates the workstation's ability to handle AI-based tasks and workflows, providing a detailed analysis of its efficiency and speed in processing complex AI algorithms and applications.
The results cover a range of AI models, indicating the server's versatility in handling different types of AI workloads.
Inference times range from very fast (4.11 ms for MobileNet V3) to much slower (2,131.81 130 ms for Real-ESRGAN), indicating a wide variation in how the system handles different AI tasks. Given the complexity and variety of functions, a score of XNUMX suggests that while the system performs adequately across a wide range of AI-based operations, there may be limitations when handling more resource-intensive tasks like Real-ESRGAN.
| UL Procyon Average Inference Times (lower is better) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | 
| Mobile Net V3 | 4.11 ms | 
| ResNet 50 | 8.40 ms | 
| Inception V4 | 30.27 ms | 
| Deep Lab V3 | 30.87 ms | 
| YOLO V3 | 45.66 ms | 
| Real-ESRGAN | 2,131.81 ms | 
| Overall Score | 130 | 
7-Zip Compression
The popular 7-Zip utility has a built-in memory benchmark that demonstrates CPU performance. In this test, we run it with a 128 MB dictionary size when possible. Again, we only have the results for the ASG-1115S-NE316R for this test.
|  | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 GB DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 GB DDR5) | 
| Compression |  |  | 
| Current CPU Usage | 5,596 % | 3,574 % | 
| Current Rating/Usage | 4.387 | 5.755 GIPS | 
| Current | 245.508 | 205.670 GIPS | 
| Resulting CPU Usage | 5,615 % | 3,572 % | 
| Resulting Rating/Usage | 4.383 | 5.733 GIPS | 
| Resulting Rating | 246.126 | 204.758 GIPS | 
| Decompression |  |  | 
| Current CPU Usage | 6,198 % | 3,810 % | 
| Current Rating/Usage | 4.224 GIPS | 5.753 GIPS | 
| Current | 261.832 GIPS | 219.191 GIPS | 
| Resulting CPU Usage | 6,057 % | 3,803 % | 
| Resulting Rating/Usage | 4.382 GIPS | 5.779 GIPS | 
| Resulting Rating | 265.381 GIPS | 219.768 GIPS | 
| Total Rating |  |  | 
| Total CPU Usage | 5,836 % | 3,687 % | 
| Total Rating/Usage | 4.383 GIPS | 5.756 GIPS | 
| Total Rating | 255.753 GIPS | 219.768 GIPS | 
Conclusion
The Supermicro Server AS-1115SV-WTNRT stands out as a robust 1U rack-mounted server, capable of handling high-demand applications, including virtualization, database management, and edge computing. Equipped with an AMD EPYC 8004 "Siena" processor, the AS-1115SV-WTNRT appears as an attractive option for service providers looking to improve data center efficiency and performance while controlling costs.
The AMD EPYC 8004 series excels in single-socket platforms, offering many cores at a lower cost than higher-core models, as well as low power consumption (starting at just 70 watts). This makes them particularly suited for dense data centers where space and power are limited. Additionally, these processors support six channels of DDR5 memory in a compact footprint, enabling large configurations that maintain both speed and efficiency.
The integration of these processors into platforms such as the Supermicro Server AS-1115SV-WTNRT aligns with current trends in cost-effective data center operations while efficiently scaling to meet future demands. Service providers who deploy these processors will certainly experience notable improvements in operational efficiency and a reduction in total cost of ownership (TCO). That said, the AS-1115SV-WTNRT offers an efficient and scalable solution that perfectly addresses modern computing challenges, providing flexibility and redundancy in storage options.
