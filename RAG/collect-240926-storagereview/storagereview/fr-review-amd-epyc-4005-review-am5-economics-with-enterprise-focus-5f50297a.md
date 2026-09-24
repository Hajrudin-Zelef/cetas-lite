---
id: collect-240926-storagereview/storagereview/fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a
title: "fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "EU", "Intel", "Nvidia", "Samsung"]
dates: []
keywords: ["amd", "apache", "benchmark", "benchmarks", "compute", "cost", "ethernet", "gpu", "intel", "latency", "licenses", "liquid cooling"]
source: docs/RAG/clean_en/storagereview/fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a.md
source_anchor: ""
source_lines: [1, 110]
sha256: 5ae9b1c77842d36610deaa71c5a08339c32bd1ba06a430ddd6cc9c2754eb5b32
---

# fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a

<!-- source: https://www.storagereview.com/fr/review/amd-epyc-4005-review-am5-economics-with-enterprise-focus -->

AMD's EPYC 4005 family brings Zen 5 to AM5 with a clear operational goal. Maintain platform consistency for system builders and hosters, then optimize the essentials for daily use. The result is a server-focused AM5 approach that prioritizes predictable deployments, stable firmware, and an OS matrix designed for scalability.
Within AMD's lineup, the EPYC 4005 sits between Ryzen on AM5 and the high-end EPYC platforms on SP5 and SP6. Ryzen can anchor light server roles, but it doesn't offer the validation depth, lifecycle, or feature set expected by data center operators. The big sockets offer more cores, memory channels, and lanes, but at a much higher cost. The EPYC 4005 bridges this gap. You keep the affordability and simplicity of AM5 while benefiting from an enterprise framework compatible with current licenses and power envelopes.
The product lineup reflects this intention. Coverage extends in both directions, allowing buyers to standardize with a modest power budget or opt for a significantly higher cache tier for latency-sensitive services, all without changing chassis or I/O footprint.
AMD EPYC 4004 vs 4005: Practical Improvements
The EPYC 4005 is a straightforward generational evolution from the 4004, retaining the AM5 heritage while improving the platform. With the 4005, you get Zen 5 cores, increased RAM performance up to DDR5-5600, and the same PCIe Gen28 capability up to 5 lanes. AMD also adds AVX-512 with a full 512-bit data path, giving this single-socket server category more headroom for vector-intensive tasks without changing motherboard or chassis.
At the top of the stack, the comparison is simple. The EPYC 4565P from the 4005 is a 16-core processor with 64 MB of L3 cache, a 170 W TDP, and a clock speed of up to 5.7 GHz. Its 4004 counterpart, the EPYC 4564P, is also a 16-core processor with 64 MB of L3 cache and a 170 W TDP, with the same maximum boost of 5.7 GHz. This means generational gains come from core and platform updates rather than a significant jump in specifications.
If you want to work with cache, both families offer a 3D V-Cache option. However, the 4005 places the EPYC 4585PX processor at the top with 170 W and 128 MB of L3 cache, while the 4004 caps out with the EPYC 4584PX at 120 W and 128 MB of cache. The base frequencies of the 4585PX are also slightly higher.
The most significant practical addition of the 4005 is a 16-core, 65 W reference. The EPYC 4545P allows bare metal hosted servers and SMBs to standardize 16 cores for Windows licensing while adhering to stricter power and cooling envelopes. This option did not exist in the 4004, where 16 cores meant 170 W on the standard chip or a 120 W V-Cache component with lower base frequencies.
As previously noted, memory support has been slightly improved. The 4004 validated two channels of DDR5 up to 5200 4005 MT/s. The 5600 brings this speed to 192 128 MT/s and increases the maximum supported capacity at the platform level from XNUMX GB to XNUMX GB, which facilitates mixed virtualized workloads and small database instances.
EPYC 4004 and 4005 Processor Comparison
| AMD EPYC 4004 Series Models | Arch. | Cores / Threads | L3 Cache (MB) | TDP (W) | Base Clock (GHz) | Boost Clock (GHz) | Price (1KU, USD) | 
|---|---|---|---|---|---|---|---|
| EPYC 4124P | Zen 4 | 4/8 | 16 | 65W | 3.8 | 5.1 | $149 | 
| EPYC 4244P | Zen 4 | 6/12 | 32 | 65W | 3.8 | 5.1 | $229 | 
| EPYC 4344P | Zen 4 | 8/16 | 32 | 65W | 3.8 | 5.3 | $329 | 
| EPYC 4364P | Zen 4 | 8/16 | 32 | 105W | 4.5 | 5.4 | $399 | 
| EPYC 4464P | Zen 4 | 12/24 | 64 | 65W | 3.7 | 5.4 | $429 | 
| EPYC4484PX | Zen 4 | 12/24 | 128 | 120W | 4.4 | 5.6 | $599 | 
| EPYC 4564P | Zen 4 | 16/32 | 64 | 170W | 4.5 | 5.7 | $699 | 
| EPYC4584PX | Zen 4 | 16/32 | 128 | 120W | 4.2 | 5.7 | $699 | 
| AMD EPYC 4005 Series Models |  |  |  |  |  |  |  | 
| EPYC 4245P | Zen 5 | 6/12 | 32 | 65W | 3.9 | 5.4 | $239 | 
| EPYC 4345P | Zen 5 | 8/16 | 32 | 65W | 3.8 | 5.5 | $329 | 
| EPYC 4465P | Zen 5 | 12/24 | 64 | 65W | 3.4 | 5.4 | $399 | 
| EPYC 4545P | Zen 5 | 16/32 | 64 | 65W | 3.0 | 5.4 | $549 | 
| EPYC 4565P | Zen 5 | 16/32 | 64 | 170W | 4.3 | 5.7 | $589 | 
| EPYC4585PX | Zen 5 | 16/32 | 128 | 170W | 4.3 | 5.7 | $699 | 
Test Platform: MSI S1102-02 with Liquid Cooling
For this test, we used the MSI S1102-02 motherboard, a 1U AM5 barebone perfectly suited to the target market of EPYC 4005 processors. It supports AM5 processors up to 170 W and our model comes with an optional all-in-one liquid cooling system, a rare feature in this category and ideal for high-end 4005 processors. This system should notably reduce fan speed and noise levels. For cooling, the chassis includes seven 40 mm system fans, a 600 W 1+1 Platinum power supply, and full management via an AST2600 BMC controller compatible with IPMI and Redfish.
Networking is handled via two RJ10 45 GbE ports on the Intel X710 processor. Storage is simple with four hot-swappable SATA bays at the front, two internal 2.5-inch SATA bays, and two NVMe M.2 2280/22110 slots. Expansion includes one Gen5 x16 slot for FHHL cards and a shared Gen4 x4 path that can be routed to a secondary PCIe slot or to one of the M.2 sockets. Memory consists of four DDR5 UDIMM slots with ECC support, validated up to 5600 1 MT/s at 48 DPC and 192 GB per DIMM for a maximum capacity of 26 GB. Physical depth is XNUMX inches, offering great deployment flexibility in short racks.
Full MSI S1102-02 Specifications
| Specifications | DETAILS | 
| Form Factor | 1U | 
| Dimensions | 438.5 mm (17.26″) L x 43.5 mm (1.71″) H x 660 mm (26″) D | 
| Processor | AMD Ryzen™ 7000/9000 Series and EPYC™ 4004/4005 Series Processor, up to 170 W TDP | 
| Socket | (1) AMD AM5 Socket | 
| Chipset | AMD B650 | 
| Memory | (4) DDR5 DIMM slots, 2DPC, ECC/non-ECC UDIMM – Max. Speed 5600MT/s(1DPC) and 3600MT/s(2DPC) – Max. Capacity per DIMM: 48 GB | 
| Drive Bays | (4) Hot-swappable 3.5"/2.5" drive bays support SATA 3.0 | 
| Internal Storage | (2) M.2 2280/22110 PCIe4.0 x4 ports (2) Internal 2.5″ drive bays support SATA 2.0 | 
| Expansion Slots | (1) Processor PCIe 4.0 x16 slot supports FHHL PCIe card | 
| Networking | (2) 10GBase-T Ethernet ports (Intel® X710AT2) *JLAN1,2 supports NCSI | 
| RAID | N/A | 
| Front I/O | (4) Hot-swappable 3.5" drive bays (2) USB3.2 Gen1 Type-A ports (1) System power button (1) UID button (1) UID LED (5) Status LEDs: Power/Fault/HDD/(2)LAN | 
| Rear I/O | (2) 10GBase-T Ethernet ports (1) Dedicated 1000Base-T server management port (4) USB3.2 Gen1 Type-A ports (1) VGA D-Sub port (1) COM DB9 port (1) UID LED button | 
| TPM | (1) TPM header with SPI interface | 
| Security | TPM 2.0 | 
| Server Management | (1) Dedicated 1000Base-T server management port (Realtek RTL8211FD-CG) ASPEED AST2600 with AMI MegaRAC-based firmware supporting IPMI 2.0 and DMTF Redfish API | 
| Environment | System operating temperature: 0°C ~ 35°C Non-operating temperature: -20°C ~ 70°C Non-operating relative humidity: 5% ~ 85% (non-condensing) | 
| Power Supply | (1+1) 450 W 80+ Platinum redundant power supply | 
| Cooling | (1) Closed-loop liquid cooling module for processor max. 170 W (7) 4028 system fans | 
| Certifications | CE, FCC (Class A) | 
| Package Contents | (1) S1102-02 1U Barebone (1) Quick Guide (1) Rail Kit (1) Closed-loop liquid cooling module (assembled inside chassis) (4) Power Cables | 
| Parts List | (1) Quick Guide: G52-S3362X2-Q13 (1) Rail Kit: E21-S336020-C27 (1) Closed-loop liquid cooling module: E34-F000091-AQ0 (4) Power Cables: (2)K33-3001002-I45 for North America, (2)K33-3001005-I45 for EU | 
| Optional Parts | TPM2.0 Module: TPM20-IR C13 Power Cord 250 V/10 A (CN): K33-3001029-I45 C13 Power Cord 125 V/13 A (JP): K33-3001346-I45 | 
AMD EPYC 4005 Performance
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
