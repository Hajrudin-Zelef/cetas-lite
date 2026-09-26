---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r5715-review-bc780bb6-4
title: "fr-review-dell-poweredge-r5715-review-bc780bb6"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "apache", "benchmarks", "compute", "energy", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r5715-review-bc780bb6.md
source_anchor: ""
source_lines: [105, 122]
sha256: 0351ecc558f1650a375da003e30d3e0b4606265f2b563a1f246aad659eabc2fc
---

# fr-review-dell-poweredge-r5715-review-bc780bb6

Phoronix Test Suite is an automated, open-source benchmarking platform supporting over 450 test profiles and more than 100 test suites via OpenBenchmarking.org. It manages the entire process, from installing dependencies to running tests and collecting results, making it ideal for performance comparisons, hardware validation, and continuous integration. Here we will compare the R5715 and R4715 processor performance using the Stream, 7-Zip, Linux kernel compilation, Apache, and OpenSSL tests.
In terms of Apache web server throughput, the R4715 reached 177,839.86 requests per second, versus 123,710.75 for the R5715, one of the closest results in the entire series. Apache's ability to deliver acceptable performance even with fewer cores, provided memory bandwidth is sufficient, explains why the gap here is smaller than for more heavily parallelized workloads.
OpenSSL transfer throughput showed a larger gap, with the R4715 reaching 533,318,299,283 bytes per second versus 148,168,050,733 bytes per second for the R5715. Cryptographic throughput is one of the workloads that scales most rapidly with thread count, and this gap clearly reflects that.
The Linux kernel compilation test revealed one of the most pronounced gaps in the suite, with the R4715 finishing in 379.53 seconds versus 1,244.86 seconds for the R5715. Kernel compilation is one of the most direct measures of how many threads a system can run simultaneously.
7-Zip compression reached 260,124 MIPS on the R4715 versus 98,555 MIPS on the R5715, consistent with the results obtained across the rest of the suite.
Stream memory throughput was 370,228.9 MB/s on the R4715, versus 230,123.6 MB/s on the R5715.
| Phoronix Benchmarks | Dell PowerEdge R4715 (AMD EPYC 9335 32 cores \| 384 GB RAM) | Dell PowerEdge R5715 (AMD EPYC 9015 8 cores \| 384 GB RAM) | 
|---|---|---|
| Apache requests per second | 177,839.86 | 123,710.75 | 
| OpenSSL transfer throughput (bytes/s) | 533,318,299,283 | 148,168,050,733 | 
| Kernel compilation time (seconds) (shorter time is better) | 379.531 | 1,244.86 | 
| 7-ZIP MIPS | 260,124 | 98,555 | 
| Stream throughput (MB/s) | 370,228.9 | 230,123.6 | 
Conclusion
The Dell PowerEdge R5715 is a storage-dedicated 2U platform, perfectly designed and fully justifying the use of a single-processor chassis in certain workload contexts. Businesses running file services, backup targets, video surveillance systems, or databases, and prioritizing disk density and I/O expandability over raw compute power, will find the R5715 an ideal solution. Its 12-bay 3.5-inch backplane, supporting up to 288 TB of raw capacity, combined with four PCIe Gen5 slots and support for two OCP 3.0 network cards, gives the platform significant room for growth without requiring the move to a more expensive dual-processor chassis.
The performance results are unequivocal. Tested with the EPYC 9015 processor, the R5715 shows significantly lower performance than the R4715 (32 cores) across all benchmarks, as expected. However, this comparison is somewhat beside the point. The R5715 is not designed for compute-intensive tasks, and the EPYC 9015 is not the processor Dell envisions for most of its customers. Configuring the R5715 with a higher-core-count EPYC 9005 processor can considerably narrow this gap, and the platform architecture is fully compatible.
The R5715 excels in the areas essential to its target use cases: storage density, expansion flexibility, energy efficiency, and management. iDRAC10 Enterprise offers a proven and consistent out-of-band management experience, directly from the 17th generation PowerEdge lineup, reducing operational costs for teams that have already invested in Dell's management suite.
For SMB and midmarket buyers looking to consolidate their storage workloads onto a suitable single-processor platform without oversizing their compute capabilities, the R5715 is an excellent choice and a natural complement to the R4715 in Dell's current AMD-based lineup.
