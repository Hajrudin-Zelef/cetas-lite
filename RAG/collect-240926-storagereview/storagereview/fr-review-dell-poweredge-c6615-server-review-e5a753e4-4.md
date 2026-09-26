---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-c6615-server-review-e5a753e4-4
title: "fr-review-dell-poweredge-c6615-server-review-e5a753e4"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "compute", "exploit", "inference", "intel"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-c6615-server-review-e5a753e4.md
source_anchor: ""
source_lines: [143, 159]
sha256: 2bd161beab7d860017ad482cd8917ef6e547367a9d039aa9ddeb433221a114ff
---

# fr-review-dell-poweredge-c6615-server-review-e5a753e4

UL Procyon AI Inference is designed to evaluate a workstation's performance in professional applications. It is important to note that this test does not exploit the capabilities of multiple processors. Specifically, this tool evaluates the workstation's ability to handle AI-based tasks and workflows, providing a detailed analysis of its efficiency and speed in processing complex AI algorithms and applications.
For this test, we use Procyon V2.7.0. In this test, lower times are better. Across all nodes, the averages were 3.91 ms on MobileNet V3, 8.4.0 ms for Resnet50, and 29.47 ms. On the rest of the scores, we saw 30.96 ms on DeepLab V3, 44.68 ms on YOLO V3, and 2008.65 ms on Real-ESRGAN. For the overall score, the nodes averaged 133.5.
| UL Procyon Computer Vision (Average Inference Time) | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Mobile Net V3 | 3.87 ms | 3.94 ms | 3.84 ms | 4.00 ms | 3.91 ms | 
| ResNet50 | 8.47 ms | 8.45 ms | 8.23 ms | 8.46 ms | 8.40 ms | 
| Inception V4 | 29.76 ms | 29.55 ms | 28.74 ms | 29.84 ms | 29.47 ms | 
| Deep Lab V3 | 30.39 ms | 30.21 ms | 33.18 ms | 30.07 ms | 30.96 ms | 
| YOLO V3 | 44.71 ms | 44.58 ms | 44.79 ms | 44.63 ms | 44.68 ms | 
| Real-ESRGAN | 2003.18 ms | 1971.97 ms | 2018.26 ms | 2041.18 ms | 2008.65 ms | 
| Overall Score | 134 | 134 | 133 | 133 | 133.5 | 
Conclusion
The Dell PowerEdge C6615 nodes offer a single AMD EPYC processor with up to 64 cores and six DDR5 slots supporting 96 GB DIMMs. The C6600 chassis that houses these nodes offers a few storage configurations. Our evaluation system has the 8x E3.S Gen5 SSD backplane. In the C6600 design, each node has access to two of these SSDs; the chassis simply provides power and direct wired access to the drives. For management, each C6615 offers iDRAC; the chassis does not have dedicated management.
We independently evaluated the capabilities of each C6615 node during our performance tests and averaged the scores of the four nodes to identify performance anomalies. The performance data highlights that the nodes operate consistently, without outliers or uneven performance. This predictability is essential for service providers and hyperscale customers who can benefit from dense systems like this.
We found the system well designed for the intended use case; our only complaint is the relatively limited Gen5 SSD support: only two drives per node. Dell would likely suggest that compute-dense customers do not need as much local storage and that cooling more Gen5 drives is a serious technical challenge, and they are probably right, we just prefer more drives than fewer on almost every occasion. Another remark worth mentioning, we are reviewing the C6615 here, but as noted at the top of this review, Dell offers additional node types for this platform, the Intel-based C6620 is available in a liquid-cooled version, which some may find compelling.
The Dell PowerEdge C6615 compute nodes offer service providers an amazing combination of performance per rack U. We have already seen many 2U4N configurations, but this design allows for greater width, and therefore greater expansion flexibility, in each server than many competing systems. Combine the excellent design with management software like iDRAC and OpenManage Enterprise and we are big fans of the final result.
Dell PowerEdge C6615 Product Page
