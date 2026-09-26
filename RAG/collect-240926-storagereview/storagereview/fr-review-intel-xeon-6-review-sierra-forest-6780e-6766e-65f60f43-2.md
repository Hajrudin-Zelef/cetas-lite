---
id: collect-240926-storagereview/storagereview/fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43-2
title: "fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "compute", "consumer", "dram", "energy", "inference", "latency", "memory", "throughput", "training"]
source: docs/RAG/clean_en/storagereview/fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43.md
source_anchor: ""
source_lines: [3, 42]
sha256: c3a80e4895188048d7705d945039cf68ddde2e86dcb9a4efe0c9baa116ca9608
---

# fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43

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
