---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662-5
title: "fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["memory", "amd", "benchmark", "cost", "energy", "gpu", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662.md
source_anchor: ""
source_lines: [133, 178]
sha256: f20f1964517aa61d784e981d61a261bfb9c0885b823a3f23fd8a9ffac94f632b
---

# fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662

The Blender OptiX benchmark results indicate that the EPYC 9665P offers good performance given its core count, but falls behind higher-core models for intensive rendering tasks. In the Monster scene, the 9665P achieves 1,026.50 samples per minute, a solid result, but far from the results of the 9965 (2,558.43) and 9755 (2,606.54). The trend continues in the Junkshop and Classroom scenes, where the 9665P achieves 795.44 and 511.45 samples per minute respectively, well below dual-processor systems like Bergamo (2p/128c) and Genoa (2p/96c).
Although the 9665P offers decent performance for most workloads, it is clear that there are more suitable processors with higher core counts. However, the 9665P holds its own against the EPYC 9575F, particularly in the Junkshop section, where its performance is nearly identical (795.44 vs 802.00).
| Blender 4.0 CPU Samples per Minute (higher is better) | AMD EPYC 9655P (96c) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c) | Genoa (2p/96c) | Bergamo (2p/128c) | 
| Monster | 1,026.50 | 2,558.43 | 2,606.54 | 1,196.15 | 1,700.65 | 2,038.71 | 
| Junkshop | 795.44 | 1,866.65 | 1,843.48 | 802.00 | 1,101.84 | 1,382.58 | 
| Classroom | 511.45 | 1,270.17 | 1,251.54 | 637.13 | 869.48 | 1,045.96 | 
Hammer DB TPROC-C
We also tested database performance with Hammer DB on the PowerEdge R7715. This system demonstrated excellent OLTP performance across all databases tested under the HammerDB TPROC-C workload (based on TPC-C, 800 warehouses).
| Database Engine | Transaction Performance (TPM) | 
|---|---|
| MariaDB 11.4.4 | 3,600,000 | 
| MySQL 8.4.4 | 3,300,000 | 
| PostgreSQL 17.2 | 3,100,000 | 
| MariaDB 10.11.12 (MDEV-21923) | 2,950,000 | 
| MariaDB 10.6.22 | 2,850,000 | 
| MySQL 5.7.44 | 2,700,000 | 
MariaDB 11.4.4 showed the best transactional performance. It outperformed older versions of MariaDB, such as 10.6.22 and the 10.11.12 version optimized for specific needs (MDEV-21923). MySQL 8.4.4 also showed excellent performance, closely trailing MariaDB 11.4.4. PostgreSQL 17.2 achieved competitive results but remained slightly behind MariaDB and the new MySQL version. MySQL 5.7.44 was the weakest database among those tested.
7-Zip Compression Benchmark
The built-in memory test in the 7-Zip utility measures a system's CPU and memory performance during compression and decompression tasks, indicating how well the system can handle data-intensive operations. We run this test with a 128 MB dictionary size when possible.
Although it has fewer cores than some other processors in the EPYC 9005 series, it achieved a total score of 378.469 GIPS. While this score is undoubtedly respectable, it is significantly lower than that of the EPYC 9755 (443.029 GIPS) and EPYC 9575F (394.900 GIPS). Interestingly, the 9965 (266.740 GIPS) lags behind in this benchmark, suggesting that a higher core count does not always translate to better compression performance.
In decompression tasks, the 9665P maintains its current level with 395.502 GIPS but remains behind the 9755 (487.263 GIPS) and 9575F (425.580 GIPS). Its single-socket architecture and high clock frequencies allow it to remain competitive, but higher-throughput models have the advantage in raw throughput.
| 7-Zip Compression Benchmark (Higher is Better) | AMD EPYC 9655P (96c) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c, SMT disabled) |  | 
| Compression |  |  |  |  |  | 
| Current CPU Usage | 5881 % | 4302 % | 5233 % | 4406 % |  | 
| Current Rating/Usage | 6.112 GIPS | 5.830 GIPS | 7.597 GIPS | 7.975 GIPS |  | 
| Current | 359.471 GIPS | 250.827 GIPS | 397.536 GIPS | 351.358 GIPS |  | 
| Resulting CPU Usage | 5875 % | 4041 % | 5306 % | 4555 % |  | 
| Resulting Rating/Usage | 6.140 GIPS | 5.804 GIPS | 7.720 GIPS | 8.070 GIPS |  | 
| Resulting Rating | 360.695 GIPS | 234.317 GIPS | 409.652 GIPS | 367.358 GIPS |  | 
| Decompression |  |  |  |  |  | 
| Current CPU Usage | 6168 % | 4322 % | 6041 % | 5017 % |  | 
| Current Rating/Usage | 6.412 GIPS | 7.078 GIPS | 8.065 GIPS | 8.483 GIPS |  | 
| Current | 395.502 GIPS | 305.909 GIPS | 487.263 GIPS | 425.580 GIPS |  | 
| Resulting CPU Usage | 6159 % | 4556 % | 5921 % | 4940 % |  | 
| Resulting Rating/Usage | 6.434 GIPS | 6.577 GIPS | 8.045 GIPS | 8.569 GIPS |  | 
| Resulting Rating | 396.243 GIPS | 299.163 GIPS | 476.405 GIPS | 422.441 GIPS |  | 
| Total Rating |  |  |  |  |  | 
| Total CPU Usage | 6017 % | 4298 % | 5613 % | 4747 % |  | 
| Total Rating/Usage | 6.287 GIPS | 6.190 GIPS | 7.883 GIPS | 8.319 GIPS |  | 
| Total Rating | 378.469 GIPS | 266.740 GIPS | 443.029 GIPS | 394.900 GIPS |  | 
Conclusion
The Dell PowerEdge R7715 offers an impressive balance of performance, scalability, and efficiency for modern enterprise workloads. Compatible with AMD EPYC 9005 series processors, offering up to 160 cores and 24 DDR5 DIMM slots for up to 6 TB of memory, the R7715 is perfectly equipped to handle data-intensive applications in virtualization, analytics, and software-defined storage environments.
With high clock frequencies and a high-performance architecture, the R7715 excels in single-core performance and offers competitive multi-core performance thanks to its single-socket design. Although it does not match dual-socket systems in raw parallel computing, it comes surprisingly close, offering a more energy-efficient and cost-effective alternative for many real-world workloads.
This configuration is not ideal for all use cases, particularly when maximum core density or GPU acceleration is essential. However, Dell is expected to roll out GPU support for the R7715 later this year. Additionally, businesses needing larger capacity drives for specific workloads might consider more extensive storage configurations.
Ultimately, the R7715 is an ideal platform for IT environments that prioritize high throughput, fast memory, and Gen5 I/O flexibility, without the complexity or cost of dual-socket deployments. The R7715 stands out as a wise option for businesses looking to optimize their efficiency without sacrificing capabilities.
Product configuration page
