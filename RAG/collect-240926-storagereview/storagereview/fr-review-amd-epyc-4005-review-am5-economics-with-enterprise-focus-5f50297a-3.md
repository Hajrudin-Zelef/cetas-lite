---
id: collect-240926-storagereview/storagereview/fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a-3
title: "fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Nvidia", "Samsung"]
dates: []
keywords: ["amd", "apache", "benchmark", "benchmarks", "compute", "cost", "gpu", "latency", "memory", "nvidia", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a.md
source_anchor: ""
source_lines: [59, 110]
sha256: 306ac75ca90ec572270e9934eff3cd4c220f8cc4efabe29949f656c356edc023
---

# fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a

In this test, we will compare the AMD EPYC 4564P processor with 16 cores and 64 MB of cache (4004 generation) to the AMD EPYC 4585PX processor with 16 cores and 128 MB of cache (4005 generation). The tests were performed on the MSI S1102-2 server, as previously indicated. Here are the components of the MSI S1102-02 test server:
- AMD EPYC 4585PX or AMD EPYC 4564P
- NVIDIA L4 GPU
- 4 Seagate Exos M 30 TB hard drives
- Solidigm P44 Pro M.2 boot SSD
- 4 x Samsung 16 GB DDR5-4800 ECC UDIMM
Y-Cruncher
y-cruncher is a multithreaded, scalable program capable of calculating Pi and other mathematical constants to trillions of digits. Since its launch in 2009, it has become a popular benchmarking and stress testing application among overclockers and PC hardware enthusiasts.
In standard tests (1B to 10B), the EPYC 4585PX is consistently faster than the EPYC 4564P, reducing computation time by 4.5 to 10.2%. For example, in the 1 billion digit test, it shows 18.802 s versus 20.933 s for the 4564P.
The BBP test shows much larger gains, up to 40%, highlighting the sensitivity of these workloads to memory latency and L3 capacity. The one billion BBP case clearly shows this: 1 s (0.387 4585 PX) versus 0.630 s (4564 38.6 P), a reduction of XNUMX%. Similar benefits are observed for larger BBP sizes.
This pattern aligns with the larger L4585 cache of the 3PX (128 MB vs 64 MB on the 4564P) and architectural improvements
| y-cruncher Total Computation Time (Lower is Better) | AMD EPYC 4585PX 16 cores | AMD EPYC 4564P 16 cores | 
| 1 billion | 18.802 seconds | 20.933 seconds | 
| 2.5 billion | 54.085 seconds | 58.108 seconds | 
| 5 billion | 121.711 seconds | 128.428 seconds | 
| 10 billion | 269.096 seconds | 281.677 seconds | 
| 1 billion BBP | 0.387 seconds | 0.630 seconds | 
| 10 billion BBP | 4.450 seconds | 7.367 seconds | 
| 100 billion BBP | 50.311 seconds | 84.051 seconds | 
Blender 4.0
Blender 4.0 is an open-source 3D modeling application. This benchmark was performed using the Blender Benchmark CLI utility. The score is measured in samples per minute, with higher values being better.
Although both processors performed well for their category, the increased cache of the EPYC 4585PX allows it to be even faster. For example, in the Monster test, the 4585PX reached 351.90 samples per minute, while the 4564P achieved 306.57 samples per minute. In Junkshop, the scores were 226.59 for the 4585PX and 195.26 for the 4564P. Similarly, in Classroom, with 171.87 and 150.27 respectively.
| Blender 4.0 CPU Samples per Minute (Higher is Better) | AMD EPYC 4585PX 16 cores | AMD EPYC 4564P 16 cores | 
| Monster | 351.90 | 306.57 | 
| Junkshop | 226.59 | 195.26 | 
| Classroom | 171.87 | 150.27 | 
Phoronix Benchmarks
We used the Phoronix test suite to automate installations, run workloads, and collect results across five key platforms: STREAM, 7-Zip, Linux kernel compilation, Apache HTTP server, and OpenSSL. Below, the EPYC 4585PX is compared directly to the previous-generation EPYC 4564P (both 16-core).
STREAM Memory Bandwidth: the 4585PX is slightly behind, with 38,472 40,106.9 MB/s versus 4.1 1,634.9 MB/s (−4564%, −XNUMX XNUMX MB/s). STREAM is highly sensitive to memory clocks/topologies and compiler choices; this indicates a small bandwidth margin for the XNUMXP in our configuration.
7-Zip (compression + decompression): the 4585PX delivers 162,951 176,484 MIPS versus 7.7 13,533 MIPS (−7%, −4564 XNUMX MIPS). XNUMX-Zip relies on integer throughput and cache/memory behavior; here, the XNUMXP is slightly ahead.
Linux Kernel Compilation (allmodconfig): 4585PX runs in 621.967 s versus 747.610 s (-125.643 s, or 16.8% faster). This reduced time reflects better parallel build throughput due to architectural improvements.
Apache Requests/sec: 4585PX serves 181,764.45 132,744.77 R/s versus 36.9 49,019.68 R/s (+XNUMX%, +XNUMX XNUMX R/s), indicating a considerable advantage in HTTP throughput.
OpenSSL Verification: the 4585PX reaches 400,939,420,057 224,487,686,510 78.6 176.45 verifications/s versus 4585 3 XNUMX XNUMX verifications/s (+XNUMX%, +XNUMX billion verifications/s). Cryptographic workloads benefit from the architectural gains of the XNUMXPX and its extended LXNUMX.
Overall: the 4585PX shows substantial advances in build, Web, and crypto tests (gains of 17 to 79%), while the 4564P stands out in raw memory bandwidth and 7-Zip in our configuration.
| Phoronix Benchmarks | AMD EPYC 4585PX 16 cores | AMD EPYC 4564P 16 cores | 
| STREAM | 38,472.0 MB/s | 40,106.9 MB/s | 
| 7-ZIP | 162,951 XNUMX MIP/s | 176,484 XNUMX MIP/s | 
| Kernel Compilation (allmod) | 621.967 XNUMX seconds | 747.610 XNUMX seconds | 
| Apache (requests per second) | 181,764.45 XNUMX R/s | 132,744.77 XNUMX R/s | 
| OpenSSL | 400,939,420,057 XNUMX XNUMX XNUMX Verifications | 224,487,686,510 XNUMX XNUMX XNUMX Verifications | 
Hard Drive FIO Performance
The MSI S1102-02 server supports four 3.5-inch hard drives at the front, offering multiple storage options. Although the fastest storage is provided by the integrated PCIe 4.0 M.2 slots, it is possible to install hard drives up to 30 TB at the front. Using 30 TB Seagate Exos M SATA hard drives, we measured sequential and random performance with FIO. We measured a maximum sequential bandwidth of 1.3 GB/s read and 1.1 GB/s write, with 4K random IOPS of 1,720 IOPS read and 1,990 IOPS write. These results were obtained in JBOD configuration; RAID will therefore influence the final performance.
| Workload | Highest Bandwidth | Highest IOPS | Lowest Latency | 
|---|---|---|---|
| Random Write (4K) | 7.8 MB/s | 1,990 | 2.106 ms | 
| Random Read (4K) | 6.7 MB/s | 1,720 | 6.861 ms | 
| Sequential Write (128K) | 1,102.5 MB/s | 8,820 | 0.486 ms | 
| Sequential Read (128K) | 1,322.3 MB/s | 10,578 | 0.459 ms | 
Conclusion
The EPYC 4005 fulfills its mission. It simplifies AM5 for system builders and hosters, while offering valuable headroom thanks to Zen 5, validated DDR5-5600, AVX-512, and broader SKU coverage. This family sits perfectly between Ryzen on AM5 and the larger SP5 and SP6 platforms, offering service providers an enterprise solution without the cost of big sockets. The SKU lineup covers common deployment models. The EPYC 4545P offers a 16-core, 65 W processor for dense hosting and clean Windows Server licensing. The EPYC 4565P remains the standard 170 W solution, while the EPYC 4585PX adds 3D V-Cache for latency-sensitive services that rely on a larger last-level cache. Memory validation up to 192 GB across four UDIMMs and stable I/O with up to 28 PCIe Gen5 lanes complete the platform.
Our data shows where this matters. The 16-core 4585PX wins clear victories in compute- and cache-sensitive tasks. y-cruncher runs tasks 4.5 to 10.2% faster on 1 to 10 billion, while BBP runs reduce time by up to 38.6%. Blender also benefits, with significant gains on Monster, Junkshop, and Classroom. In the Phoronix suite, the 4585PX reduces Linux kernel compilation time by 16.8%, processes 36.9% more Apache requests per second, and records a 78.6% increase in OpenSSL verification. STREAM and 7-Zip are slightly closer to the previous 4564P in our configuration.
The EPYC 4005 is a moderate but significant update from AMD. It preserves the affordability and simplicity that made the 4004 easy to adopt, while improving practical performance and flexibility across a wide range of hosted and edge use cases. If you are developing single-socket AM5 servers at scale, this default solution is ideal for new deployments and is a wise upgrade when cache or efficiency are at stake, especially when paired with the MSI S1012-02 server featuring the internal liquid loop.
