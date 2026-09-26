---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf-4
title: "fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "apache", "benchmarks", "compute", "energy", "gpus", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf.md
source_anchor: ""
source_lines: [111, 127]
sha256: 4aac050a5c98f33385343eabf0f68136df394ff5de68c17587c3f4e40839ba3a
---

# fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf

Apache web server throughput was among the closest results in the suite, with the R4715 reaching 177,839.86 requests per second versus 123,710.75 for the R5715. Apache can maintain reasonable throughput with fewer cores when memory bandwidth is sufficient, which explains a smaller gap here than with more heavily parallelized workloads.
OpenSSL transfer throughput showed a larger gap, with the R4715 reaching 533,318,299,283 bytes/s versus 148,168,050,733 bytes/s for the R5715. Cryptographic throughput increases strongly with the number of threads, and this gap directly reflects that.
The Linux kernel compilation test revealed one of the most pronounced gaps in the suite. The R4715 completed in 379.53 seconds versus 1,244.86 seconds for the R5715; kernel compilation being one of the most direct measures of how many threads a system can execute simultaneously.
7-Zip compression reached 260,124 MIPS on the R4715 versus 98,555 MIPS on the R5715, tracking consistently with the rest of the suite.
Stream memory throughput was measured at 370,228.9 MB/s on the R4715, versus 230,123.6 MB/s on the R5715.
| Phoronix Benchmarks | Dell PowerEdge R4715 (AMD EPYC 9335 32 cores \| 384 GiB RAM) | Dell PowerEdge R5715 (AMD EPYC 9015 8 cores \| 384 GiB RAM) | 
|---|---|---|
| Apache requests per second | 177,839.86 | 123,710.75 | 
| OpenSSL transfer throughput (bytes/s) | 533,318,299,283 | 148,168,050,733 | 
| Kernel compilation time (seconds) (shorter time is better) | 379.531 | 1,244.86 | 
| 7-ZIP MIPS | 260,124 | 98,555 | 
| Stream throughput (MB/s) | 370,228.9 | 230,123.6 | 
Conclusion
The Dell PowerEdge R4715 is a high-performance 1U platform that fully justifies a single-processor architecture for typical SMB workloads. Businesses using virtualization, large-scale databases, and network edge deployments, for which licensing efficiency and ease of use are paramount, will find the R4715 perfectly suited. Its 1U form factor, three PCIe Gen5 slots, 24 DDR5 RDIMM slots, and flexible 2.5-inch and 3.5-inch storage options give this platform great versatility without the additional costs of a dual-processor chassis.
The performance results demonstrate the platform's capabilities. Tested with the EPYC 9335 processor, the R4715 consistently outperformed the R5715 across all benchmarks, with the core count advantage particularly visible in heavily parallelized workloads such as kernel compilation, OpenSSL, and Blender. Users who don't need 32 cores can opt for the EPYC 9255 (24 cores), EPYC 9135 (16 cores), or EPYC 9015 (8 cores), choosing a processor suited to their needs and budget.
It is important to clarify what the R4715 is not. It does not support GPUs or DPUs, a deliberate choice to limit costs and footprint. For workloads requiring accelerators, the AMD-based R6715 and R7715 models are best suited. Faster network connectivity is available via PCIe AIC, with 100 GbE and 400 GbE options.
The R4715 consistently excels in areas essential to its target use cases. It offers high compute density in a 1U chassis, optimal energy efficiency thanks to its 800 W and 1,100 W Platinum and Titanium power supplies, flexible NVMe and SAS/SATA storage configurations, and a proven iDRAC10 Enterprise management experience, perfectly integrated into the 17th generation PowerEdge lineup. For SMBs and mid-sized businesses seeking a single-socket compute platform suited to their needs, without overspending on storage or expansion capacity, the R4715 is an excellent choice and a natural complement to the R5715 in Dell's current AMD-based server lineup.
