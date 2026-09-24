---
id: collect-240926-storagereview/storagereview/fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43
title: "fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "benchmark", "compute", "consumer", "dram", "energy", "inference", "latency", "memory", "throughput", "training"]
source: docs/RAG/clean_en/storagereview/fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43.md
source_anchor: ""
source_lines: [1, 117]
sha256: 18ac3851062cdd3e654697612deefa1de26811b5a1b53e8833e7dc80ca0d54cd
---

# fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43

<!-- source: https://www.storagereview.com/fr/review/intel-xeon-6-review-sierra-forest-6780e-6766e -->

Intel's Xeon 6 server processor family launched a few weeks ago, replacing the old "Scalable" brand. There are two lines of chips: Sierra Forest, which features E-cores, and Granite Rapids, which features P-cores. In this review, we sampled some early sets of Sierra Forest processors, including the high-end 6780E, which ships 144 cores.
Xeon 6 processors are designed to support virtually any workload, from obvious AI use cases to less demanding edge deployments. To help customers understand the two swim lanes, Intel says Sierra Forest is "optimized for performance per watt in high-density, scalable compute workloads." The Granite Rapids P-core processor is "optimized for performance per core in compute-intensive workloads." However, all Xeon 6 processors use a common platform and firmware stack.
To give a little more context on processor performance, Intel positions Sierra Forest as a great replacement for systems with a 5-year refresh cycle. In this case, the performance per watt figures should be significantly more favorable.
The rollout of Xeon 6 will happen in stages, with the launch of the 6700E family today. But there's much more to come. Intel Xeon 6900P processors are expected to ship this year in the third quarter. We expect to see a flood of chips, including the 3E, 6900P, 6700P, Xeon 6500 SoC and 6P, in the first quarter of 6300.
Architectural innovations
The Xeon 6 features two microarchitecture designs, a performance-optimized core platform (P-cores) and an efficiency-optimized core platform (E-cores). This hybrid approach allows data centers to design a rack in which compute-intensive workloads, such as AI and HPC, benefit from maximum performance while throughput-oriented tasks, such as microservices and networking, achieve a new level of energy efficiency.
Increased core count
The main feature of the Xeon 6 Sierra Forest processors is the significant increase in core count. By integrating more cores per processor, Intel enables data centers to simultaneously handle a wider range of workloads. This increase in core density is particularly beneficial for applications that require high levels of parallel processing, ensuring that resources are used optimally to maximize throughput and minimize latency.
Improved memory bandwidth
For a few generations, it has felt like we've been stuck with slower server DRAM speeds while the consumer market has moved well above 7000 5 MT/s. The integration of DDR2.0 memory with Ultra Path Interconnect (UPI) 6 in Xeon XNUMX Sierra Forest processors significantly improves memory bandwidth. This enables faster data access and reduces bottlenecks, ensuring high-performance applications run properly. Improved memory bandwidth is essential for applications that handle large data sets and require fast data retrieval and processing, such as AI training and big data analytics.
Advanced I/O capabilities of Intel Xeon 6
Support for PCIe 5.0 and Compute Express Link (CXL) 2.0 gives Xeon 6 Sierra Forest processors advanced I/O capabilities. PCIe 5.0 offers double the bandwidth of its predecessor, enabling faster communication between the processor and peripherals. CXL 2.0 further enhances connectivity by providing high-bandwidth, low-latency interconnect for processors, accelerators, memory, and storage. This ensures data centers can meet the demands of modern high-throughput applications and seamlessly integrate with future technologies.
Modular multi-die architecture
Xeon 6 Sierra Forest processors use a modular multi-chip architecture enabled by Embedded Multi-die Interconnect Bridge (EMIB) technology. This design allows multiple chips to be combined in a single package, delivering high bandwidth and low latency while maintaining efficient power consumption. The modular approach provides flexibility and scalability, allowing data centers to customize their infrastructure to meet specific workload requirements. This architecture also supports better thermal management and energy efficiency, which are essential to maintaining performance and reliability in high-density environments.
Intel Xeon 6 die architecture
The die architecture of the Xeon 6 Sierra Forest processors provides an interesting glimpse into Intel's innovative approach to maximizing performance and efficiency. Unlike a traditional large silicon interposer, EMIB implements a small bridge die with multiple routing layers. By leveraging EMIB technology, Intel can interconnect multiple dies in a single package, reducing overall package size and improving signal integrity. This configuration enables faster data transfer rates between dies, which is essential for maintaining high performance across a wide range of applications.
Importance for data centers
Performance and efficiency
Compared to previous generations, Xeon 6 processors deliver up to 2.7 times better performance per watt, making them ideal for AI inference, media transcoding, and general compute tasks. This balance ensures data centers can handle more workloads with lower power consumption, which directly translates into reduced operating costs and improved sustainability.
Scalability and flexibility
The modular architecture allows data centers to customize their infrastructure based on specific workload demands. The ability to deploy a combination of P-cores and E-cores across multiple platforms offers a tailored approach to managing diverse computing needs, improving both performance and efficiency. This presents an interesting concept in which a chassis can be selected and then deployed, with some having E-core processors available to efficiently run core services and handle operations during less demanding times, while others are equipped with P-core processors to handle peak demand and demanding workloads.
Improved security and reliability
Built with enhanced hardware-level security features, Xeon 6 processors provide robust protection for data integrity and system reliability. This is essential for maintaining trust and compliance in data-sensitive environments.
Future-proofing through advanced technologies
Supporting the latest technologies such as DDR5, PCIe 5.0, and CXL 2.0, Xeon 6 processors ensure data centers are ready for future advancements and can seamlessly integrate with emerging hardware and software solutions.
Modern solutions for modern problems
Intel Xeon 6 Sierra Forest processors represent a major advance in data center architecture. By combining high performance, energy efficiency, and scalability, they meet the multifaceted needs of modern data centers, paving the way for improved operational efficiency, reduced costs, and the ability to meet the growing demands of AI and other compute-intensive applications.
Intel Xeon 6 E-core SKU map
All Sierra Forest SKUs feature E-Cores with 88 PCIe Gen5/CXL lanes.
| SKU | Cores | Base GHz | TurboGHz | Max Turbo GHz | L3 Cache MB | TDP Watts | Max Scala. | DDR5 1DPC Memory Speed | 
|---|---|---|---|---|---|---|---|---|
| 6780E | 144 | 2.2 | 3.0 | 3.0 | 108 | 330 | 2S | 6400 | 
| 6766E | 144 | 1.9 | 2.7 | 2.7 | 108 | 250 | 2S | 6400 | 
| 6756E | 128 | 1.8 | 2.6 | 2.6 | 96 | 225 | 2S | 6400 | 
| 6746E | 112 | 2.0 | 2.7 | 2.7 | 96 | 250 | 2S | 5600 | 
| 6740E | 96 | 2.4 | 3.2 | 3.2 | 96 | 250 | 2S | 6400 | 
| 6731E | 96 | 2.2 | 3.1 | 3.1 | 96 | 250 | 1S | 5600 | 
| 6710E | 64 | 2.4 | 3.2 | 3.2 | 96 | 205 | 2S | 5600 | 
Intel Xeon 6 Performance Tests
Our lab sampled two sets of processors for this review, the 6780E and 6766E. Intel provided a QCT server platform for testing. We want to note a few key caveats in our data. The server platform itself was unstable in many of our test environments. For example, Windows Server 2022 would not run properly, so we used 2025.
The processors are early samples and not retail boxed processors. As such, we were unable to run our full test suite and did not have time before the embargo lifted to properly resolve some of the performance and stability issues we encountered. As such, the following data should be considered informative and not definitive. We will wait for platforms to ship from Intel's OEM partners before making a more conclusive judgment on the capabilities of the Sierra Forest processors.
Intel shipped us a Quanta QuantaGrid D55Q-2U system as a test platform to showcase the new processors. This server features a 24-bay direct-connect NVMe backplane supporting U.2 Gen4/Gen5 SSDs.
Our test configuration included 16 x 16 GB DDR5-6400 RAM and Micron MTC10F1084S1RC64BDY modules.
Xeon 6780E and Xeon 6766E
- Quanta QuantaGrid D55Q-2U
- 256GB DDR5 6400MHz
Xeon Platinum 8592+
- Dell Poweredge R760
- 1 TB DDR5 4800 XNUMX MHz
Xeon Gold 6430
- Dell Poweredge R760
- 1 TB DDR5 4800 XNUMX MHz
Xeon Platinum 8480+
- HP ML350 Gen 11
- 256GB DDR5 4800MHz
With the earliest test platform gremlins in play, the new Xeon 6780E and 6766E processors were tested in a Windows Server 2025 environment, while the older processors were tested in Windows Server 2022.
Blender OptiX 4.0
In Blender OptiX, we have 3 different tests: Monster, Junkshop, and Classroom. In Monster, we saw the 6780E and 6766E beat the 8592+ by 21% and 14%, with an 18% difference between the two. In Junkshop, the 6780E was 8592% ahead of the 10.5+, with the 6766E behind the 8592+ by only 0.35%. There was a 10.8% difference between the 6780E and 6766E in Junkshop. For the Classroom part, the 6780E and 6766E beat the 8592+ by 20% and 22%, with a 10% difference between the 6780E and 6766E.
| Blender 4.0 Processor | 2x Xeon 6780E (256 GB DDR5) | 2x Xeon 6766E (256 GB DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 
|---|---|---|---|---|---|
| Monster | 1410.463 | 1297.715 | 1115.057 | 943.300 | 540.039 | 
| Junkshop | 862.418 | 777.716 | 780.408 | 627.662 | 361.066 | 
| Classroom | 696.543 | 628.960 | 556.550 | 475.144 | 278.228 | 
Cinebench R23
For Cinebench R23, multi-core performance showed the 6780E and 6766E behind the 8592+ by 38% and 42%, respectively. For single-core performance, the 6780E was 24% behind the 8592+ and the 6766E was 31% behind. The difference between the 6780E and 6766E dropped to about 5% for the multi-core score and 18% for single-core.
| Cinebench R23 | 2x Xeon 6780E (256 GB DDR5) | 2x Xeon 6766E (256 GB DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 
|---|---|---|---|---|---|
| Multi-core CPU | 67,984 | 64,326 | 110,498 | 79,164 | 69,663 | 
| Single-core CPU | 873 | 793 | 1,144 | 1,461 | 1,022 | 
| PM Ratio | 77.91x | 81.10x | 96.63x | 54.20x | 68.17x | 
Cinebench 2024
For Cinebench 2024, the multi-core score saw a drop of about 55% from the 8592+ to the 6780E and a drop of 61% from the 8592+ to the 6766E. This also showed a 13% difference between the 6780E and 6766E on Multi-Core. For single-core scores, the 6780E and 6766E had only a 5% difference, with the 6780E and 6766E behind the 8592+ by 37% and 34%, respectively.
| Cinebench R23 | 2x Xeon 6780E (256 GB DDR5) | 2x Xeon 6766E (256 GB DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 
|---|---|---|---|---|---|
| Multi-core CPU | 2,687 | 2,347 | 6,001 | 3,746 | 4,699 | 
| Single-core CPU | 43 | 45 | 68 | 59 | 76 | 
| PM Ratio | 62.85x | 52.65x | 88.48x | 63.22x | 61.44x | 
Y-Cruncher
Y-cruncher is a popular benchmarking and stress-testing application launched in 2009. This test is multithreaded and scalable, calculating Pi and other constants to billions of digits. Faster is better in this test.
The new 6780E and 6766E run a bit slower than the Emerald Rapids 8592+, but they aren't exactly direct competitors. The 6780E ran about 13% faster than the Xeon Gold 6430 for the 1 billion test, but about 39% slower than the Xeon Platinum 8592+. The 6766E ran 19% slower than the Gold 6439 and 42% slower than the Platinum 8592+.
| Y-Cruncher (lower is better) | Xeon 6780E (256 GB DDR5) | Xeon 6766E (256 GB DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 
|---|---|---|---|---|---|
| 1 billion | 6.927 seconds | 7.254 seconds | 4.239 seconds | 6.060 seconds | 5.136 seconds | 
| 2.5 billion | 17.898 seconds | 19.507 seconds | 11.466 seconds | 16.896 seconds | 13.768 seconds | 
| 5 billion | 38.454 seconds | 41.116 seconds | 25.325 seconds | 36.843 seconds | 29.889 seconds | 
| 10 billion | 81.146 seconds | 87.403 seconds | 54.921 seconds | 80.574 seconds | 65.194 seconds | 
| 25 billion | 217.530 seconds | 238.813 seconds | 156.923 seconds | 229.017 seconds | 186.841 seconds | 
| 50 billion | 565.913 seconds | 502.245 seconds | N/A | N/A | N/A | 
7-Zip
The popular 7-Zip utility has a built-in memory benchmark that demonstrates processor performance very well. In this test, we run it with a 128 MB dictionary size when possible. In this test, the 6780E barely beats the Platinum 8592+ in the overall score. For decompression, the 6766E also beat the 8592+.
| 7-Zip Compression | Xeon 6780E (256 GB DDR5) | Xeon 6766E (256 GB DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 
|---|---|---|---|---|---|
| Compression |  |  |  |  |  | 
| Current CPU Usage | 5,891 % | 4,768 % | 5,609 % | 5,732 % | 5,482 % | 
| Current Rating/Usage | 4.985 GIPS | 4.614 GIPS | 4.912 GIPS | 3.912 GIPS | 4.628 GIPS | 
| Current | 293.689 GIPS | 220.001 GIPS | 275,503 GIPS | 224.209 GIPS | 253.724 GIPS | 
| Resulting CPU Usage | 5,603 % | 5,103 % | 5,605 % | 5,669 % | 5,475 % | 
| Resulting Rating/Usage | 4.954 GIPS | 4.638 GIPS | 4.883 GIPS | 3.923 GIPS | 4.628 GIPS | 
| Resulting Rating | 277.670 GIPS | 236.910 GIPS | 273.716 GIPS | 222.407 GIPS | 253.382 GIPS | 
| Decompression |  |  |  |  |  | 
| Current CPU Usage | 5,962 GIPS | 5,798 % | 6,243 % | 5,852 % | 6,219 % | 
| Current Rating/Usage | 4.550 GIPS | 4.152 GIPS | 3.635 GIPS | 3.423 GIPS | 3.745 GIPS | 
| Current | 271.266 GIPS | 240.693 GIPS | 226.917 GIPS | 200.350 GIPS | 231.916 GIPS | 
| Resulting CPU Usage | 5,832 % | 6,029 % | 6,232 % | 5,894 % | 6,129 % | 
| Resulting Rating/Usage | 4.540 GIPS | 4.161 GIPS | 3.654 GIPS | 3.385 GIPS | 3.871 GIPS | 
| Resulting Rating | 264.764 GIPS | 250.853 GIPS | 227.744 GIPS | 199.363 GIPS | 237.259 GIPS | 
| Total Rating |  |  |  |  |  | 
| Total CPU Usage | 5,717 % | 5,566 % | 5,919 % | 5,781 % | 5,802 % | 
| Total Rating/Usage | 4.747 GIPS | 4.399 GIPS | 4.269 GIPS | 3.654 GIPS | 4.249 GIPS | 
| Total Rating | 271.217 GIPS | 243.882 GIPS | 250.730 GIPS | 210.363 GIPS | 245.320 GIPS | 
Conclusion
The new Intel Xeon 6 E-core lineup, part of the Sierra Forest family, represents a change of pace in revision cadence. In previous Intel releases, we typically saw the high-end processors first; this time, Intel is leading with efficient mid-range SKUs.
That said, the Sierra Forest processors hold up well, even compared to previous 5th-generation scalable models. But according to Intel, customers looking to buy new hardware will compare them to systems that are about 5 years old and being retired, running something like Intel Xeon 8280 processors. With those systems as components, the Xeon 6 E-Core processors offer a massive argument in terms of density and energy savings.
In this review, we're dealing with very early processors and a server that's quite capable for a Sierra Forest visit, but not quite where we want it to be for a full review. Between Windows operating system bugs (not present in our Ubuntu tests) and waiting for BIOS revisions, it's clear there's room to grow and maximize the performance of these CPUs. With this small taste, we're optimistic about what Sierra Forest can deliver today and what Granite Rapids processors can deliver later this year for high-performance workloads. Also, let's not sleep on the possible 288 E-Core 6900 Series in the coming quarters. Intel will offer its customers more options than ever when it comes to tuning infrastructure for their applications.
We look forward to final shipping hardware including Sierra Forest processors in the lab soon. We'll provide in-depth analysis of power consumption, storage performance, and much more in upcoming reviews.
