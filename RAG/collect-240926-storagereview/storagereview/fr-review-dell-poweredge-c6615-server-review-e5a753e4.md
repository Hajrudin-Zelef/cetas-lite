---
id: collect-240926-storagereview/storagereview/fr-review-dell-poweredge-c6615-server-review-e5a753e4
title: "fr-review-dell-poweredge-c6615-server-review-e5a753e4"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "China", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "benchmark", "compute", "datacenter", "distribution", "ethernet", "exploit", "gpu", "inference", "intel", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-poweredge-c6615-server-review-e5a753e4.md
source_anchor: ""
source_lines: [1, 159]
sha256: 6a54f2713ffd227ae3d9155ba9686da0a3ec2f4a61aecd29b2310690300cffee
---

# fr-review-dell-poweredge-c6615-server-review-e5a753e4

<!-- source: https://www.storagereview.com/fr/review/dell-poweredge-c6615-server-review -->

The Dell PowerEdge C-Series platform has a 2U chassis supporting four servers in the Dell Modular Infrastructure category. Depending on the workload, the C-series system can be configured with two different node types: a single-socket AMD C6615 node or a dual-socket Intel C6620 node.
Our review will focus on the C-series chassis, which features four single-socket AMD EPYC nodes connected to an E8.S PCIe Gen3 5-bay disk backplane.
From a storage perspective, the platform can be configured with a 2.5-inch SFF disk backplane, which supports up to 24 NVMe SSDs or Gen5 support by leveraging a 3-bay E8.S backplane. Internally, these drives are connected directly to each node, with equal distribution across the four servers. For example, in the 24-bay configuration, each node sees six drives; in the 8-bay configuration, each node sees two drives.
The C6600 chassis provides shared redundant power supplies and cooling for the four installed nodes, though beyond that, each node is managed independently. So, unlike a blade chassis managed with a chassis management portal, this is more like four small PowerEdge servers under one metal roof. Each C6615 node has dedicated network connections, an iDRAC interface, and PCIe slots for expansion.
Dell PowerEdge C6615 Node Specifications
| C6615 Specifications |  | 
|---|---|
| Processor | One AMD EPYC processor with up to 64 cores | 
| Memory | 6 DDR5 DIMM slots, supports up to 576 GB RDIMM (6 x 96 GB), speeds up to 4800 MT/s | 
| Storage Controllers | Internal controllers (RAID): PERC H755N, PERC H355 Internal Boot: Boot Optimized Storage Subsystem (NVMe BOSS-N1): HWRAID 1, 2 x M.2 SSD Internal SAS 12 Gbit/s HBA (non-RAID): HBA355i Software RAID: S160 | 
| Availability | Hot-plug redundant hard drives and power supplies | 
| Drive Bays | Front bays: Up to 16 x 2.5-inch SAS/SATA (HDD/SSD) drives, 61 TB maximum Up to 16 x 2.5-inch SATA/NVMe drives, 15.36 TB maximum on a universal backplane configuration Up to 16 x 2.5-inch on NVMe backplane Up to 8 x E3.s on the NVMe SSD hard drive backplane | 
| Hot-swap, Redundant Power Supplies | 3200 W 277 VAC or 336 VDC 2800 W Titanium 200-240 VAC or 240 VDC Platinum 2400 W 100-240 VAC or 240 VDC 1800 W Titanium 200-240 VAC or 240 VDC | 
| Dimensions | Height – 40.0 mm (1.57 inches) Width – 174.4 mm (6.86 inches) Depth – 549.7 mm (21.64 inches), 561.3 mm (22.10 inches) – SAS/SATA or NVMe or E3.S or universal configuration | 
| Weight | 3.7 kg (8.15 pounds) | 
| Integrated Management | iDRAC9 iDRAC Direct RESTful API iDRAC with Redfish iDRAC Service Module | 
| OpenManage Software | CloudIQ plugin for PowerEdge OpenManage Enterprise OpenManage Enterprise Integration for VMware Vcenter OpenManage Integration for Microsoft System Center OpenManage Integration with Windows Admin Center OpenManage Power Manager plugin OpenManage Service plugin OpenManage Update Manager plugin | 
| Integrations | BMC TrueSight Microsoft System Center OpenManage Integration with ServiceNow OpenManage Integration with Windows Admin Center OpenManage Power Manager plugin OpenManage Service plugin OpenManage Update Manager plugin | 
| Security | AMD Secure Encrypted Virtualization (SEV) AMD Secure Memory Encryption (SME) Cryptographically signed firmware Data encryption at rest (SED with local or external key management) Secure BootSecured Component Verification (hardware integrity verification) Secure Erase Silicon Root of Trust System Lockdown (requires iDRAC9 Enterprise or Datacenter) TPM 2.0 FIPS, CC-TCG certified, TPM 2.0 China NationZ | 
| Integrated Network Card | 1 x 1 Gb | 
| Rear Ports | 1 x USB 3.0 1 iDRAC Ethernet port 1 iDRAC Direct port (Micro-AB USB) 1 x Mini DisplayPort | 
| PCIE Slots | Up to 2 Low-Profile PCIe x16 Gen5 slots 1 x OCP 3.0 x16 Gen5 | 
| Operating Systems and Hypervisors | Canonical Ubuntu LTS Server Microsoft Windows Server with Hyper-V Red Hat Enterprise Linux SUSE Linux Enterprise Server VMware ESXi/vSAN | 
Build and Design
The Dell PowerEdge C6600 chassis and C6615 nodes offer an exceptionally dense computing option for deployment scenarios that need to minimize physical space used in a rack-mount environment. This is suitable for hyperconverged solutions running in a clustered environment, requiring multiple nodes or heavy workloads that do not require consuming 4U or 8U through traditional 1U or 2U server designs. The chassis has a 2U footprint with a depth of 30 inches. The weight of the chassis can increase depending on the final configuration. Dell indicates a maximum weight of a C16 6600-bay configuration with all drives installed at 93.69 pounds.
The front of the system is fairly basic compared to other PowerEdge platforms, without much Dell branding. This type of server does not offer the standard PowerEdge bezel but places the drives and fan intakes at the forefront. The front of the E3.S C6600 version features eight Gen5 NVMe SSDs in the middle, flanked by cooling fan intakes.
The side ears of the chassis contain dedicated power buttons for each node and information buttons indicating the status or issues of that node.
Each C6615 node has a condensed port layout at the rear of the chassis compared to a traditional 1U or 2U server. Ports include USB, iDRAC, a display connector, and a USB service port.
For networking, an OCP slot is available for various interface options (ours has a four-port 25GbE network card), and two PCIe slots are also available. The OCP and dual PCIe slots provide a Gen5 interface.
Opening the PowerEdge C6600 chassis gives you visibility into the layout of how cooling, power distribution, and disk I/O paths are managed. The PCIe/SAS cabling from the disk backplane is routed directly to each node via quick-connect fittings that also transmit data and power.
Depending on the internal configuration of each node, the drive connections connect directly to the motherboard or to a PERC card for hardware RAID options.
Apart from cooling and power, the nodes do not share any other resources.
Dell PowerEdge C6615 Performance
Tested Node Specifications
Our four C6615 nodes have identical configurations. We will compare them and show the average performance across the nodes.
- 1 AMD EPYC 8534P processor 64 cores
- 6 x 96 GB DDR5 4800 576 MB/s (XNUMX GB)
- Windows Server Standard 2022
- Dell RAID1 BOSS boot SSD
- 2 PCIe Gen5 E3.S SSDs
During our performance tests, the nodes ran in parallel to give an overall score taking into account shared power and cooling resources.
Storage Performance
Each of the four Dell Power Edge C6615 nodes includes a RAID1 BOSS SSD for booting and two E3.S bays for Gen5 enterprise SSDs. Although the BOSS card is no slouch, it offers a very different performance profile than the E3.S SSDs.
Although much of this review focuses on overall system-level performance, we briefly covered both types of storage on this system with corner-case workloads. Our first test focused on the BOSS RAID1 boot SSD group.
| Dell BOSS RAID1 | Read Performance | Write Performance | 
|---|---|---|
| Sequential 1 MB Q32/4T | 2,963 MB/s | 1,067 MB/s | 
| Random 4K Q32/8T | 600,786 IOPS (0.426 ms) | 249,819 IOPS (1.024 ms) | 
Next, we looked at a single Gen5 E3.S SSD, which included the KIOXIA CM7 7.68 TB read-intensive SSD in our test system.
| KIOXIA 7.68 TB CM7-R | Read Performance | Write Performance | 
|---|---|---|
| Sequential 1 MB Q32/4T | 13,736 MB/s | 7,089 MB/s | 
| Random 4K Q32/8T | 931,671 IOPS (0.266 ms) | 768,739 IOPS (0.329 ms) | 
Cinebench R23
Maxon's Cinebench R23 is a CPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. Higher scores are better. Here are the results for all EPYC chips.
In Cinebench R23, the four nodes were around 74,000 on the multi-core portion, with node 3 slipping into 75,000. All four nodes remained much closer for single-core scores, with nodes 1 and 4 at 1,088. Node 2 was only 5 points behind and node 3 was 8 points ahead. Overall, all nodes showed only minor performance variations, typical of different processors, even though they all belong to the same model.
| Cinebench R23 | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Multi-core CPU | 74,877 | 74,961 | 75,011 | 74,745 | 74,898.5 | 
| Single-core CPU | 1,088 | 1,093 | 1,084 | 1,088 | 1,088.25 | 
| MP Ratio | 64.84 | 68.60 | 69.17 | 68.70 | 67.83 | 
Cinebench 2024
Maxon's Cinebench 2024 is a CPU and GPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. Since these nodes do not have a GPU, we only have the multi-core and single-core numbers.
In Cinebench 2024, all nodes remained close to each other, with minimal variance on the multi-core and single-core portions. The average performance was 4,509 points for multi-core and 67.25 points for single-core, with an MP ratio of 66.98.
| Cinebench 2024 | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Multi-core CPU | 4,544 | 4,577 | 4,436 | 4,481 | 4,509.5 | 
| Single-core CPU | 68 | 68 | 65 | 68 | 67.25 | 
| MP Ratio | 66.79 | 67.23 | 68.21 | 65.69 | 66.98 | 
Geekbench 6 CPU
Geekbench 6 is a cross-platform benchmark that measures the overall performance of a system. This test includes a dedicated CPU portion and a GPU portion, but since these nodes do not have a GPU, we only have the CPU results. The higher the score, the better the performance.
In Geekbench, we saw tight numbers until we got to node 3, which dipped slightly on single-core and multi-core. The average across all nodes was 1,687 single-core and 19,319.5 multi-core.
| Geekbench 6 CPU | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Single-Core | 1,707 | 1,708 | 1,625 | 1,708 | 1,687 | 
| Multi-Core | 19,544 | 19,234 | 18,999 | 19,501 | 19,319.5 | 
Blender 4.0 CPU
The next step is Blender OptiX, an open-source 3D modeling application. This benchmark was run using the Blender Benchmark CLI utility. The score is expressed in samples per minute, with the highest being the best.
The C6615 nodes recorded fairly consistent numbers. The average scores were 591.79 on Monster, 415.88 on Junkshop, and 311.74 on Classroom.
| Blender 4.0 CPU | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Monster | 595.23 | 593.51 | 584.35 | 594.07 | 591.79 | 
| Junkshop | 415.26 | 415.11 | 418.05 | 415.08 | 415.88 | 
| Classroom | 308.57 | 312.91 | 312.69 | 312.78 | 311.74 | 
Blender 4.1 CPU
Blender OptiX 4.1 brings new features, such as GPU-accelerated denoising, streamlining the rendering process and reducing the time required for denoising tasks. Despite these advances, the overall performance improvements in benchmark scores compared to version 4.0 are minimal, indicating only slight improvements in efficiency.
Again, we see consistent numbers across the board, with averages of 587.22 on Monster, 420.20 on Junkshop, and 306.60 on Classroom.
| Blender 4.1 CPU | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Monster | 590.46 | 590.58 | 584.76 | 583.08 | 587.22 | 
| Junkshop | 418.38 | 416.71 | 426.73 | 419.03 | 420.20 | 
| Classroom | 306.86 | 304.81 | 308.95 | 305.79 | 306.60 | 
7-Zip Compression
The popular 7-Zip utility has a built-in memory test that demonstrates CPU performance. In this test, we run it at a dictionary size of 128 MB when possible.
Fair scores were observed across all nodes. In the total scores, we found a total CPU usage of 5,778.75%, a total rating/usage of 4.355 GIPS, and a total rating of 252 GIPS.
| Blender 4.1 CPU | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Compression |  |  |  |  |  | 
| Current CPU Usage | 5,548% | 5,549% | 5,633% | 5,585% | 5,578.75% | 
| Current Rating/Usage | 4.256 GIPS | 4.210 GIPS | 4.156 GIPS | 4.177 GIPS | 4.20 GIPS | 
| Current | 236.158 GIPS | 233.626 GIPS | 234.092 GIPS | 233.285 GIPS | 234.290 GIPS | 
| Resulting CPU Usage | 5,536% | 5,537% | 5,601% | 5,553% | 5,556.75% | 
| Resulting Rating/Usage | 4.193 GIPS | 4.202 GIPS | 4.172 GIPS | 4.168 GIPS | 4.184 GIPS | 
| Resulting Rating | 232.118 GIPS | 232.631 GIPS | 233.691 GIPS | 231.443 GIPS | 232.470 GIPS | 
| Decompression |  |  |  |  |  | 
| Current CPU Usage | 5,973% | 6,027% | 5,992% | 6,014% | 6,001.5% | 
| Current Rating/Usage | 4.543 GIPS | 4.501 GIPS | 4.565 GIPS | 4.509 GIPS | 4.530 GIPS | 
| Current | 271.343 GIPS | 271.287 GIPS | 273.507 GIPS | 271.196 GIPS | 271.833 GIPS | 
| Resulting CPU Usage | 5,997% | 6,015% | 5,999% | 5,990% | 6,000.25% | 
| Resulting Rating/Usage | 4.537 GIPS | 4.519 GIPS | 4.550 GIPS | 4.499 GIPS | 4.526 GIPS | 
| Resulting Rating | 272.066 GIPS | 271.775 GIPS | 272.946 GIPS | 269.509 GIPS | 271.574 GIPS | 
| Total Rating |  |  |  |  |  | 
| Total CPU Usage | 5,767% | 5,776% | 5,800% | 5,772% | 5,778.75% | 
| Total Rating/Usage | 4.365 GIPS | 4.360 GIPS | 4.361 GIPS | 4.333 GIPS | 4.355 GIPS | 
| Total Rating | 252.092 GIPS | 252.203 GIPS | 253.318 GIPS | 250.476 GIPS | 252.022 GIPS | 
Blackmagic Raw Speed Test
We use the Blackmagic Raw Speed Test to evaluate how machines perform actual RAW decoding. This test can integrate both CPU and GPU usage, but we will only test CPU usage.
All four nodes showed extremely close performance, with an average of 119.75 FPS.
| Blackmagic Raw Speed Test | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| CPU 8K | 121 FPS | 121 FPS | 118 FPS | 119 FPS | 119.75 FPS | 
Blackmagic Disk Speed Test
Next is the Blackmagic Disk Speed Test. This test runs a 5 GB file sample for read and write speeds. Since it is single-threaded, it will not show the highest disk speeds, but it still gives a good perspective.
The C6615s have a BOSS card inside, using two M.2 drives in RAID1, so performance is slightly degraded for reliability. For write speeds, we found an average of 991.6 MB/s and for read speeds, an average of 2,801 MB/s.
| Blackmagic Disk Speed Test | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Write | 999.8 MB/s | 977.4 MB/s | 991.4 MB/s | 997.7 MB/s | 991.6 MB/s | 
| Read | 2,807.4 MB/s | 2,790.1 MB/s | 2,828.0 MB/s | 2,780.4 MB/s | 2,801.5 MB/s | 
Y-Cruncher
y-cruncher is a multi-threaded and scalable program that can calculate Pi and other mathematical constants to trillions of digits. Since its launch in 2009, it has become a popular benchmarking and stress testing application for overclockers and hardware enthusiasts.
For our average speeds, we saw 9.5 seconds for 1 billion, 24.20 seconds for 2.5 billion, and 50.73 seconds for 5 billion. On the most significant digit calculations, we saw 105.73 seconds for 10 billion, 288.85 seconds for 25 billion, and 633.5 seconds for 50 billion.
| Y Cruncher (total computation time, in seconds) | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| 1 billion | 9.587 | 9.459 | 9.350 | 9.633 | 9.507 | 
| 2.5 billion | 24.490 | 24.225 | 23.334 | 24.740 | 24.197 | 
| 5 billion | 51.427 | 50.990 | 49.303 | 51.214 | 50.734 | 
| 10 billion | 107.084 | 107.646 | 103.772 | 107.443 | 105.736 | 
| 25 billion | 291.918 | 290.944 | 280.632 | 291.902 | 288.849 | 
| 50 billion | 641.709 | 640.289 | 619.100 | 640.917 | 635.504 | 
UL Procyon AI Computer Vision Benchmark
UL Procyon AI Inference is designed to evaluate a workstation's performance in professional applications. It is important to note that this test does not exploit the capabilities of multiple processors. Specifically, this tool evaluates the workstation's ability to handle AI-based tasks and workflows, providing a detailed analysis of its efficiency and speed in processing complex AI algorithms and applications.
For this test, we use Procyon V2.7.0. In this test, lower times are better. Across all nodes, the averages were 3.91 ms on MobileNet V3, 8.4.0 ms for Resnet50, and 29.47 ms. On the rest of the scores, we saw 30.96 ms on DeepLab V3, 44.68 ms on YOLO V3, and 2008.65 ms on Real-ESRGAN. For the overall score, the nodes averaged 133.5.
| UL Procyon Computer Vision (Average Inference Time) | Node 1 | Node 2 | Node 3 | Node 4 | Average | 
|---|---|---|---|---|---|
| Mobile Net V3 | 3.87 ms | 3.94 ms | 3.84 ms | 4.00 ms | 3.91 ms | 
| ResNet50 | 8.47 ms | 8.45 ms | 8.23 ms | 8.46 ms | 8.40 ms | 
| Inception V4 | 29.76 ms | 29.55 ms | 28.74 ms | 29.84 ms | 29.47 ms | 
| Deep Lab V3 | 30.39 ms | 30.21 ms | 33.18 ms | 30.07 ms | 30.96 ms | 
| YOLO V3 | 44.71 ms | 44.58 ms | 44.79 ms | 44.63 ms | 44.68 ms | 
| Real-ESRGAN | 2003.18 ms | 1971.97 ms | 2018.26 ms | 2041.18 ms | 2008.65 ms | 
| Overall Score | 134 | 134 | 133 | 133 | 133.5 | 
Conclusion
The Dell PowerEdge C6615 nodes offer a single AMD EPYC processor with up to 64 cores and six DDR5 slots supporting 96 GB DIMMs. The C6600 chassis that houses these nodes offers a few storage configurations. Our evaluation system has the 8x E3.S Gen5 SSD backplane. In the C6600 design, each node has access to two of these SSDs; the chassis simply provides power and direct wired access to the drives. For management, each C6615 offers iDRAC; the chassis does not have dedicated management.
We independently evaluated the capabilities of each C6615 node during our performance tests and averaged the scores of the four nodes to identify performance anomalies. The performance data highlights that the nodes operate consistently, without outliers or uneven performance. This predictability is essential for service providers and hyperscale customers who can benefit from dense systems like this.
We found the system well designed for the intended use case; our only complaint is the relatively limited Gen5 SSD support: only two drives per node. Dell would likely suggest that compute-dense customers do not need as much local storage and that cooling more Gen5 drives is a serious technical challenge, and they are probably right, we just prefer more drives than fewer on almost every occasion. Another remark worth mentioning, we are reviewing the C6615 here, but as noted at the top of this review, Dell offers additional node types for this platform, the Intel-based C6620 is available in a liquid-cooled version, which some may find compelling.
The Dell PowerEdge C6615 compute nodes offer service providers an amazing combination of performance per rack U. We have already seen many 2U4N configurations, but this design allows for greater width, and therefore greater expansion flexibility, in each server than many competing systems. Combine the excellent design with management software like iDRAC and OpenManage Enterprise and we are big fans of the final result.
Dell PowerEdge C6615 Product Page
