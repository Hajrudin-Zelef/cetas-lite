---
id: collect-240926-storagereview/storagereview/fr-review-hpe-proliant-dl560-gen11-review-closed-loop-liquid-cooling-664217b8-2
title: "fr-review-hpe-proliant-dl560-gen11-review-closed-loop-liquid-cooling-664217b8"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["benchmark", "dram", "gpu", "gpus", "intel", "license", "liquid cooling", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-hpe-proliant-dl560-gen11-review-closed-loop-liquid-cooling-664217b8.md
source_anchor: ""
source_lines: [3, 72]
sha256: 7cc0f3a459f0ef2c8be8f7c3dc4773572af5d0f4b346b36e795e21aa85fd6e1f
---

# fr-review-hpe-proliant-dl560-gen11-review-closed-loop-liquid-cooling-664217b8

If you've ever considered a standard dual-socket 2U server and wished for twice as many processors and a massive RAM footprint, the HPE ProLiant DL560 Gen11 is here for you. This server supports four 4th Generation Intel Xeon Scalable processors, up to 16 TB of DRAM, and a wide variety of storage and rear bay expansion. What else is cool? For quad-processor configurations with a TDP of 270 W or higher (per chip), HPE provides closed-loop liquid cooling to keep those cores running at optimal performance.
Yes. You read that right. This system ships with a closed-loop liquid cooling system for the quad-CPU configuration! How cool is that? (pun intended)
But seriously, HPE broke the mold with this configuration and managed to pack all that power and flexibility into a standard 2U chassis. What makes it so unique, however, is the stacked arrangement of the processor boards that allows system administrators to relatively easily access any of the four chips and the ability to integrate the liquid cooling systems into a single complete package without having to worry about a drain or a leak over the entire equipment.
HPE ProLiant DL560 Gen11 Hardware
Once the top cover is removed, you can see the two upper processors and their cooling plates, as well as the range of available RAM slots, which in our test model is almost empty; however, to ensure good airflow, it comes with spacer blanks.
The closed-loop system goes from the radiator to the first processor through the second processor, then back up to the radiator to cool the now-heated liquid. Below is a mirror image of what you see on top.
One thing to keep in mind is that the liquid cooling system is covered by a 5-year warranty, after which the system must be replaced (in its entirety) at the customer's expense. No part of the cooling system can be repaired or replaced by the customer and must be done by HPE-certified personnel during the 5-year warranty period.
Looking more closely at the cooling plates, and like most CPU coolers for high-performance systems, access to the processor underneath is straightforward, as always.
A close and personal look at the storage configuration provided to us. Like most HPE systems coming out of their factories, the DL560 Gen11 is very versatile and can accommodate several different controllers and storage paths.
Depending on the storage configuration, the backplane will have different combinations of options. Our test system has only a single "bay" filled, but HPE offers several storage configurations, including E3.S SSD backplanes.
The rear of this machine has enough space for additional components such as GPUs (up to 6!) or NVMe boot drives if that's your thing. Networking and connectivity are easily accessible and configured at the rear.
Additionally, even though our test model didn't need them, you can see where there is room for four power supply modules to power this beast.
HPE ProLiant DL560 Gen11 Management
Through iLO, you can view general information such as processor, memory, network, and storage configurations, as well as the status of each.
In the Power and Thermal tab, we can see the power reading and power supply status on the server configuration. Detailed power information is available in the Power Meter tab. However, this required a paid iLO license, which this system did not include.
Next to the power supplies are the fans and cooling modules. This is where you can monitor the liquid cooling system. Unfortunately, it only displays the general status, whether the pumps are running and at what speed.
iLO contains a temperature information tab under Power and Thermal. Based on each temperature sensor on a 3D graph, this information shows you where your hot spots are in the system. Temperature data is essential to monitor to ensure the stability and longevity of your hardware.
HPE ProLiant DL560 Gen11 Specifications
The potential range of specifications for the DL560 Gen11 includes:
| Processor | 4th Generation Intel® Xeon® Scalable processors supporting up to 60 cores | 
| Memory |  | 
| Storage Controllers |  | 
| Drive Bays | 8, 16, or 24 SFF SAS/SATA/NVMe | 
| Power Supplies |  | 
| Fans | Liquid cooling solution, Hot-plug redundant fans, Performance fan kit, or High-performance fan kit | 
| Dimensions | 3.4 cm x 17.05 cm x 31.75 cm | 
| Form Factor | 2U rackable chassis | 
| Integrated Management |  | 
| Server Utilities |  | 
| Security |  | 
| HPE iLO Remote Management Network Port | 1 Gb dedicated, rear | 
| Network Options | None standard. A choice of OCP network card or stand-up network card is required. BTO models will be preselected with a primary network card. | 
| GPU Options | Up to 6 small GPUs or 2 large ones | 
| USB | Up to 7 total:  | 
| Additional Ports |  | 
| PCIe |  | 
| Operating System and Hypervisors |  | 
Performance
HPE ProLiant DL560 Gen11 Configuration
- 4 Intel Xeon Platinum 8444H processors (16 cores, 2.9 GHz)
- 512 GB DDR5 (8 x 64 GB DDR5-4800)
- Windows Server 2022
We compare it to the DL320, recently tested, for comparison.
Blackmagic RAW Speed Test
We always run the Blackmagic RAW speed test, which tests video playback. It's a good indicator. It is more of a hybrid test including CPU and GPU performance for actual RAW decoding. Although there are no other HPE machines to show a direct comparison here, to give you an idea of where this one stands, the other servers typically only produce at 50-60 fps.
| Blackmagic RAW Speed Test (Higher is better) | HPE ProLiant DL560 Gen11 (Quad Intel Xeon Platinum 8444H, 64 cores, 2.9 GHz) | HPE ProLiant DL320 (4th Gen Intel Xeon-G 6430e processor, 32 cores, 2.1 GHz) | 
| CPU 8K | 86 | 99 | 
| CUDA 8K | N/A | N/A | 
Cinebench R23
Maxon's Cinebench R23 is a CPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. Higher scores are better.
| Cinebench R23 | HPE ProLiant DL560 Gen11 (Quad Intel Xeon Platinum 8444H, 64 cores, 2.9 GHz) | HPE ProLiant DL320 (4th Gen Intel Xeon-G 6430e processor, 32 cores, 2.1 GHz) | 
| CPU (multi-core) (points) | 83,478 | 38,707 | 
| CPU (single-core) (points) | 1,261 | 1,245 | 
| MP Ratio | 66.21x | 31.10x | 
Cinebench 2024
Maxon's Cinebench 2024 is a CPU and GPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. Since this configuration has no GPU, we don't have those numbers. Higher scores are better.
| Cinebench 2024 | HPE ProLiant DL560 Gen11 (Quad Intel Xeon Platinum 8444H, 64 cores, 2.9 GHz) | HPE ProLiant DL320 (4th Gen Intel Xeon-G 6430e processor, 32 cores, 2.1 GHz) | 
| CPU (multi-core) (points) | 3,215 | 2,014 | 
| CPU (single-core) (points) | 67 | 71 | 
| MP Ratio | 48.31x | 28.29x | 
Geekbench CPU Benchmark
Geekbench 6 is a cross-platform benchmarking tool measuring the overall performance of a system. However, it would be interesting to analyze single-core and multi-core performance, as well as OpenCL benchmark results. A high score indicates better performance. Let's note that we only examined CPU results, as this server does not have a graphics card.
You can find comparisons with any system in the Geekbench browser.
| Geekbench 6 | HPE ProLiant DL560 Gen11 (Quad Intel Xeon Platinum 8444H, 64 cores, 2.9 GHz) | HPE ProLiant DL320 (4th Gen Intel Xeon-G 6430e processor, 32 cores, 2.1 GHz) | 
| CPU Benchmark - Single-Core | 3,215 | 2,014 | 
| CPU Benchmark - Multi-Core | 67 | 71 | 
| GPU Benchmark – OpenCL | 48.31x | 28.29x | 
y-cruncher
y-cruncher is a multi-threaded and scalable program that can calculate Pi and other mathematical constants to trillions of digits. Since its launch in 2009, it has become a popular benchmarking and stress-testing application for overclockers and hardware enthusiasts.
| y-cruncher (Total Computation Time) | HPE ProLiant DL560 Gen11 (Quad Intel Xeon Platinum 8444H, 64 cores, 2.9 GHz) | HPE ProLiant DL320 (4th Gen Intel Xeon-G 6430e processor, 32 cores, 2.1 GHz) | 
