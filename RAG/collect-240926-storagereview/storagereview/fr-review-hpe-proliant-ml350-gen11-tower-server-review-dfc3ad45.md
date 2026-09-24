---
id: collect-240926-storagereview/storagereview/fr-review-hpe-proliant-ml350-gen11-tower-server-review-dfc3ad45
title: "fr-review-hpe-proliant-ml350-gen11-tower-server-review-dfc3ad45"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["benchmark", "distribution", "dram", "gpu", "gpus", "inference", "intel", "license", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-hpe-proliant-ml350-gen11-tower-server-review-dfc3ad45.md
source_anchor: ""
source_lines: [1, 116]
sha256: c2a8994ad864476d201e18f389e24e31149b578466d7030aa8945885b4fba032
---

# fr-review-hpe-proliant-ml350-gen11-tower-server-review-dfc3ad45

<!-- source: https://www.storagereview.com/fr/review/hpe-proliant-ml350-gen11-tower-server-review -->

Many companies are looking for alternatives to rack-mounted servers and the physical maintenance requirements they entail. It makes sense to study the capabilities of tower servers, as they combine the power and performance of rack-mounted servers with the deployment considerations of small businesses. The HPE ProLiant ML350 Gen11 tower server is specifically designed to meet these needs of SMBs or edge environments.
HPE generously provided a review unit so that we could put it to the test. Before getting to the test results, let's start with the background and specifications.
HPE ProLiant ML350 Gen11 Server Specifications
A unique characteristic of the ML350 Gen11 is that no two towers are identical. HPE removed the "default" settings option that customers could click on and order. Instead, HPE asks everyone intending to purchase these beasts to come to the table with their needs and specifications in mind in order to customize a machine that precisely matches the workload it will need to handle.
The potential range of specifications for the ML350 Gen11 includes:
| Processor | 4th/5th Gen Intel® Xeon® Scalable processors supporting up to 64 cores | 
| Memory |  | 
| Storage Controllers | Tri-mode RAID controller:  | 
| Drive Bays | Front bays:  | 
| Power Supplies |  | 
| Fans | Standard – 3 fans included | 
| Dimensions | 46.2 (H) x 71.2 (D) x 17.4 (W) cm 18.2 (H) x 28 (D) x 6.85 (W) in | 
| Form Factor | 4U Tower with rack conversion capability | 
| Integrated Management |  | 
| Server Utilities |  | 
| Security |  | 
| HPE iLO Remote Management Network Port | 1 Gb dedicated, rear | 
| Network Options | No native ports are included. Choice of OCP or stand-up card | 
| GPU Options | Up to 8SW or 4DW | 
| Ports |  Internal:  | 
| PCIe |  | 
| Operating Systems and Hypervisors |  | 
HPE ML350 Gen 11 Build and Design
At first glance, the unique redesign of HPE's latest Gen11 tower server looks like something you might encounter on the Chicago horizon, dominating the morning fog that stretches over the icy Lake Michigan.
Like most tower servers, this one stands out in the server room. At 18 inches tall and over 2 feet deep, consideration will need to be given to placement for heat distribution and noise control.
The dual-depth fan system is a notable feature for HPE when it comes to heat distribution. It is an impressive cooling fan system that our team immediately noticed and was impressed by when we opened the side panel of the server.
HPE ships the ML350 with three fans in base configurations where it is a non-hot-plug replacement and up to eight hot-swappable fans in the dual-processor configuration of our server.
The HPE ProLiant ML350 Gen11 offers several backplane configurations for storage, including support for SFF, LFF, and Gen5 E3.S SSDs. Our test machine was equipped with a backplane capacity for HPE MR408i-o Gen11 x8 Lanes.
The front of the tower with the front bezel opened. Our model has 4 of the 8 bays equipped with SATA drives.
As customizable as the rest of the system, the rear of the tower provides 10 PCIe expansion card slots, such as HBAs, GPUs, and network cards, with options on the side rail to configure specific network interfaces and video outputs.
HPE ProLiant ML350 Gen11 Management – iLO 6
The ML350 Gen 11 incorporates iLO 6, HPE's latest version of its Lights-Out technology. iLO 6 brings a range of benefits (such as remote server configuration, health monitoring, and power and temperature control) that improve the manageability, security, and efficiency of the DL320 and simplify complex IT environments. Although iLO is not new and is commonly found on most HPE servers, we note it here because it represents considerable added value for those who may not be accustomed to such in-depth server management.
System Information
Through ILO, you can view general information such as processor, memory, network, and storage configurations, as well as the status of each.
Next, in the ILO Processor tab, is the Memory tab. You can see the status and configuration of your DRAM, down to the slot level, as well as the frequency at which each DIMM is operating.
At the end of the System Information section is the Storage tab. Here, you can see the status of your storage controllers and the drives connected to them. This information can help identify faulty drives in your configuration in the event of a failure.
Power and Thermal
In the Power and Thermal tab, we can see the power reading and power supply status on the server configuration. More detailed power information can be found under the power meter tab, but this requires a paid ILO license not included with this system for us to try it.
ILO contains a temperature information tab under Power and Thermal. This information shows you, based on each temperature sensor on a 3D graph, where your hot spots are in the system. It is important to monitor temperature data to ensure the stability and longevity of your hardware.
Firmware and Operating System Software
A huge advantage of ILO is being able to remotely install operating systems and firmware on the device without having to put your hands on it. These remote management integrations save considerable time since you don't have to worry about video and peripheral input directly on the device.
HPE ProLiant ML350 Gen 11 Performance
Verifying the Configuration
- 2 Intel Xeon 8480+ processors
- 256GB DDR5 Memory
- Windows Server 2022
Blackmagic RAW Speed Test
We started by running the Blackmagic RAW speed test, which tests video playback. This is more of a hybrid test including CPU and GPU performance for actual RAW decoding. Although there are no other HPE machines to show a direct comparison here, to give you an idea of where this one stands, other servers typically only produce at 50-60 fps.
| Blackmagic RAW Speed Test (Higher is better) | HPE ML350 Gen11 (Dual Intel Xeon(r) Platinum 8480+, 112 cores, 2 GHz) | 
| CPU 8K | 136 | 
| CUDA 8K | N/A | 
Cinebench R23
Maxon's Cinebench R23 is a CPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. Higher scores are better.
| Cinebench R23 | HPE ML350 Gen11 (Dual Intel Xeon(r) Platinum 8480+, 112 cores, 2 GHz) | 
| CPU (Multi-Core) (points) | 79164 | 
| CPU (Single-Core) (points) | 1461 | 
| MP Ratio | 54.20x | 
Cinebench 2024
Maxon's Cinebench 2024 is a CPU and GPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. Since this configuration has no GPU, we don't have those figures. Higher scores are better.
| Cinebench 2024 | HPE ML350 Gen11 (Dual Intel Xeon(r) Platinum 8480+, 112 cores, 2 GHz) | 
| CPU (Multi-Core) (points) | 4,699 | 
| CPU (Single-Core) (points) | 76 | 
| MP Ratio | 61.44x | 
Geekbench CPU Benchmark
Geekbench 6 is a cross-platform benchmarking tool measuring the overall performance of a system. However, it would be interesting to analyze single-core and multi-core performance, as well as OpenCL benchmark results. A high score indicates better performance. Note that we only examined CPU results, as this server does not have a graphics card.
You can find comparisons with any system in the Geekbench browser.
| Geekbench 6 | HPE ML350 Gen11 (Dual Intel Xeon(r) Platinum 8480+, 112 cores, 2 GHz) | 
| CPU Benchmark - Single-Core | 1,939 | 
| CPU Benchmark - Multi-Core | 15,218 | 
| GPU Benchmark – OpenCL | N/A | 
y-cruncher
y-cruncher is a multi-threaded and scalable program that can calculate Pi and other mathematical constants to trillions of digits. Since its launch in 2009, it has become a popular benchmarking and stress-testing application for overclockers and hardware enthusiasts.
| y-cruncher (Total Computation Time) | HPE ML350 Gen11 (Dual Intel Xeon(r) Platinum 8480+, 112 cores, 2 GHz) | 
| 1 billion digits (seconds) | 5.136 | 
| 2.5 billion digits (seconds) | 29.889 | 
| 10 billion digits (seconds) | 65.194 | 
| 25 billion digits (seconds) | 186.841 | 
| 50 billion digits (seconds) | 413.722 | 
7-Zip Compression
The popular 7-Zip utility has a built-in memory test that demonstrates CPU performance very well. In this test, we run it with a 128 MB dictionary size when possible.
|  | HPE ML350 Gen11 (Dual Intel Xeon(r) Platinum 8480+, 112 cores, 2 GHz) | 
|---|---|
| Compression |  | 
| Current CPU Usage | 5,482% | 
| Current Rating/Usage | 4.628 GIPS | 
| Current | 253.724 GIPS | 
| Resulting CPU Usage | 5,475% | 
| Resulting Rating/Usage | 4.628 GIPS | 
| Resulting Rating | 253.382 GIPS | 
| Decompression |  | 
| Current CPU Usage | 6,219% | 
| Current Rating/Usage | 3.745 GIPS | 
| Current | 231.916 GIPS | 
| Resulting CPU Usage | 6,129% | 
| Resulting Rating/Usage | 3.871 GIPS | 
| Resulting Rating | 237.259 GIPS | 
| Total Rating |  | 
| Total CPU Usage | 5,802% | 
| Total Rating/Usage | 4.249 GIPS | 
| Total Rating | 245.320 GIPS | 
UL Procyon AI Inference
UL's Procyon AI inference benchmark suite evaluates the performance of different AI inference engines using state-of-the-art neural networks. These tests were run on the CPU only. Each value represents an average inference time (the lower the value, the better the performance), and the last line indicates an overall score (the higher the value, the better the performance).
|  | HPE ML350 Gen11 (Dual Intel Xeon(r) Platinum 8480+, 112 cores, 2 GHz) | 
|---|---|
| Mobile Net V3 | 2.34 | 
| ResNet 50 | 5.76 | 
| Inception V4 | 21.70 | 
| Deep Lab V3 | 23.00 | 
| YOLO V3 | 30.81 | 
| REAL-ESRGAN | 1535.27 | 
| Overall Score | 191 | 
Conclusion
Overall, the HPE ProLiant ML350 Gen11 arrives on the market with a multitude of customization options that will appeal to any business looking to upgrade its server capabilities without sacrificing the physical space that would typically be required for a rack-mounted system. When we began the review, this system supported 4th Gen Intel Xeon Scalable, but during the process, they also picked up 5th Gen chips, bringing even more power to this platform.
Additionally, this server can accommodate more than two dozen hot-plug NVMe SSDs for capacities up to 368 TB, 32 DIMM slots capable of running up to 8 TB of DDR5 RAM, a ton of PCIe expansion, and an impressive fan system to keep everything cool and running efficiently. Tower servers are often overlooked because they are designed "only for SMBs." The HPE ProLiant ML350 Gen11 is another excellent reminder that towers can be much more than that; they are among the most versatile and customizable systems available to meet the full range of today's increasingly diverse use cases.
