---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-r770ap-review-84b71047-4
title: "fr-review-dell-poweredge-r770ap-review-84b71047"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["apache", "benchmark", "benchmarks", "compute", "exploit", "gpu", "intel", "latency", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-r770ap-review-84b71047.md
source_anchor: ""
source_lines: [93, 137]
sha256: 63237214e1f5da2a6553fee409c175a864edddb45802f990d4d91f532683c3be
---

# fr-review-dell-poweredge-r770ap-review-84b71047

| 25 billion | 99.324 seconds | 86.298 seconds | 
| 50 billion | 221.255 seconds | 192.128 seconds | 
| 100 billion | 491.737 seconds | 430.208 seconds | 
Blender
An open-source 3D modeling application. This benchmark was run with the Blender Benchmark utility. The score is expressed in samples per minute, with the higher being the better.
In the Blender 4.3 benchmark, the R770AP outperformed the R770 in all three scenes. In the "Monster" scene, the R770AP reached 2,200.116 samples per minute, versus 1,706.002 for the R770. In the "Junkshop" scene, the R770AP achieved 1,565.643 samples per minute, versus 1,169.370 for the R770. Finally, in the "Classroom" scene, the R770AP obtained 1,076.122 samples per minute, versus 791.475 for the R770, a performance gain of approximately 36% on this workload.
| CPU performance test with Blender 4.3 (higher samples per minute is better) | Dell PowerEdge R770 (2x Intel Xeon 6787P \| 2 TB RAM) | Dell PowerEdge R770AP (2x Intel Xeon 6978P \| 3 TB RAM) | 
|---|---|---|
| Monster | 1,076.122 samples/min | 1,076.122 samples/min | 
| Junkshop | 1,076.122 samples/min | 1,076.122 samples/min | 
| Classroom | 1,076.122 samples/min | 1,076.122 samples/min | 
Phoronix Benchmarks
Phoronix Test Suite is an open-source automated benchmarking platform supporting over 450 test profiles and over 100 test suites via OpenBenchmarking.org. It handles the entire process, from installing dependencies to running tests and collecting results, making it ideal for performance comparisons, hardware validation, and continuous integration. Here we will compare the performance of the R770AP and R770 using the Stream, 7-Zip, Linux kernel compilation, Apache, and OpenSSL tests.
Stream
In the Stream memory bandwidth test, the R770AP achieved a clear improvement over the R770, reaching 869,965.3 MB/s versus 472,135.6 MB/s. This represents nearly double the memory bandwidth of the reference system, reflecting the R770AP's larger and faster memory configuration.
7-Zip
In the 7-Zip compression test, the R770AP scored 806,375 MIPS, versus 628,206 MIPS for the R770, a clear improvement due to the higher core count of the 6978P processors.
Kernel Compilation
In the Linux kernel compilation test, where a shorter time is preferable, the R770AP completed the allmod compilation in 176.391 seconds versus 188.793 seconds on the R770, reducing compilation time by approximately 12 seconds.
Apache
The Apache test was the only area where the R770 slightly outperformed the R770AP, with a score of 60,258.5 requests per second versus 48,729.63 for the R770AP. This result is important because web server workloads do not always scale linearly with core count and can be influenced by memory latency and I/O characteristics.
OpenSSL
In the OpenSSL verification test, the R770AP scored 2,515,270,390,853 verifications/s versus 2,216,883,554,350 verifications/s on the R770, a significant gain in cryptographic throughput that highlights the computational efficiency of the 6978P at scale.
| Phoronix Benchmarks | Dell PowerEdge R770 (2x Intel Xeon 6787P 86C) | Dell PowerEdge R770AP (2x Intel Xeon 6978P \| 3 TB RAM) | 
|---|---|---|
| Stream | 472,135.6 MB/s | 869,965.3 MB/s | 
| 7-ZIP | 628,206 MIPS | 806,375 MIPS | 
| Kernel Compilation (allmod) (lower is better) | 188.793 seconds | 176.391 seconds | 
| Apache (requests per second) | 60,258.5 R/s | 48,729.63 R/s | 
| OpenSSL | 2,216,883,554,350 Verifications | 2,515,270,390,853 Verifications | 
Dell PowerEdge R770AP: Deterministic Performance and High-Frequency Trading
While our standard test suite focuses on compute throughput, memory bandwidth, and general scalability, the R770AP's design priorities extend into an area we typically do not test: microsecond-level execution determinism. To illustrate this platform's capabilities for its most demanding target audience, Dell published a technical note in partnership with Metrum AI, evaluating the R770AP specifically for high-frequency trading workloads. We did not perform these tests, nor did we independently audit the results. Nevertheless, we include a summary here, as it most directly demonstrates how this server is a distinct product from the R770.
Metrum AI's methodology relies on a custom tool called jitter-c, which measures per-core wake-up latency jitter—that is, the regularity with which a thread scheduled to execute at a precise moment actually starts. This metric isolates processor scheduling variability from network, memory, and application factors, providing a reliable point of comparison between processor generations. By comparing an R770AP equipped with two Xeon 6980P processors (256 cores total) to a previous-generation R760 with two Xeon Platinum 8592+ processors (128 cores total), the study found that the Granite Rapids-AP architecture reduced p99 wake-up jitter to approximately 1 microsecond—roughly half that of the older platform—while doubling core density. These jitter profiles were then fed into a backtesting simulation engine to model the financial impact. The results are summarized below.
| Metrum AI HFT Backtest Results | Dell PowerEdge R760 (2x Xeon 8592+, 128 cores) | Dell PowerEdge R770AP (2x Xeon 6980P, 256 cores) | 
|---|---|---|
| p99 Wake-up Jitter | ~2 µs | ~1 µs | 
| Mean Reversion: Total Trades | 5,175 | 6,229 (+20.4%) | 
| Mean Reversion: Trades/sec | 819 | 991 (+21.1%) | 
| Market Making: Total Trades | 21,765 | 32,491 (+49.3%) | 
| Market Making: Trades/sec | 2,067 | 3,072 (+48.6%) | 
As Seamus Jones of Dell noted in his commentary on the study, the added value lies not in the speed, but in the predictability of that speed. Indeed, in trading, a fast but inconsistent system is a source of risk. Conversely, a deterministic system is a strategic asset.
Conclusion
The Dell PowerEdge R770AP occupies a well-defined place within the 17th generation PowerEdge lineup. It does not replace the R770, and Dell does not present it as such. The R770 remains the versatile and highly configurable Intel 2U platform it has always been, with GPU support, mixed SAS/SATA/NVMe storage, E-core and P-core processor options, and up to 8 TB of memory across 32 DIMM slots. For organizations running general virtualization solutions, mixed enterprise applications, or workloads that leverage that configuration flexibility, the R770 remains the ideal choice.
The R770AP is designed for workloads for which the R770 was never optimized. By adopting the Granite Rapids-AP platform, with its 12-channel memory architecture, up to 128 processing cores per socket, and 504 MB of L3 cache, Dell has created a 2U system that prioritizes compute density, memory bandwidth, and execution determinism over versatility. Our performance tests reflect this priority: STREAM bandwidth nearly doubled, Blender rendering improved by 29 to 36%, and CPU scaling extended consistently as working sets exceeded cache capacity. The Apache regression is an important point to note, as it demonstrates that the R770AP's NUMA topology requires workload consideration to fully exploit its performance, and that not all applications will benefit from this platform change without optimization.
The Metrum AI tests published by Dell alongside this platform highlight the underlying determinism. Halving p99 scheduling jitter while doubling core density represents a significant architectural improvement for organizations running high-frequency trading operations, real-time risk engines, large-scale in-memory analytics, and massively parallel simulations. For these workloads, the R770AP is a high-performance platform that is perfectly suited. For all other applications, the R770 and R7725 remain the most relevant options within the PowerEdge lineup.
