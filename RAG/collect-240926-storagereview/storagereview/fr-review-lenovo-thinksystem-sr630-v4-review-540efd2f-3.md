---
id: collect-240926-storagereview/storagereview/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f-3
title: "fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Samsung"]
dates: []
keywords: ["benchmark", "cost", "gpus", "intel"]
source: docs/RAG/clean_en/storagereview/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f.md
source_anchor: ""
source_lines: [56, 105]
sha256: f4cd9a567e40ad1bc286db5c5b90c12dea7a576f6d9450fd7adee907cb206692
---

# fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f

The front panel can also be configured with an optional Mini DisplayPort video port, which can be used for quick local monitoring and diagnostics without having to access the rear of the rack. Up to two optional USB 3.0 ports are also available; one is explicitly designated for connection to the Lenovo XClarity Controller (XCC). This connectivity simplifies management tasks, such as downloading firmware updates or running diagnostics directly from a USB device.
Lenovo has integrated an external diagnostic port on our system, which can be very useful for on-site IT staff who need to troubleshoot hardware issues (such as system status and failures). This can speed up problem resolution and minimize downtime. It also has a removable information tag for quick access to essential system details, such as serial numbers, configurations, and network information.
The indicators and controls on the front control panel provide the usual at-a-glance information on system status and activity, including power and reset buttons and LEDs for drive status and health.
Rear Panel
The rear panel includes hot-swap power supplies (from 800 W to 2000 W) located on each side of the system, which provide redundancy and can be replaced without shutting down the system. The Dual OCP 3.0 slots support PCIe Gen 5 x16 for advanced networking, allowing for high-bandwidth adapters such as dual-port 200 GbE cards. The panel also includes a video port, two USB 3.0 ports, and a dedicated XClarity Controller (XCC) management port for local and remote system management. LEDs provide quick visual updates on system status.
The rear panel includes a mix of low-profile and full-height PCIe slots for expansion, allowing users to add GPUs, storage controllers, or other adapters. Storage options include hot-swap 2.5-inch drives and hot-swap M.2 drives, offering flexible configurations for boot devices or additional storage. The rear panel is also available in four air-cooled configurations and two water-cooled configurations.
Internal
When you open the Lenovo ThinkSystem SR630 V4, you will see the two processors surrounded by their respective DIMM slots. This chassis offers 16 DIMMs per processor, for a total of 32, allowing up to 2 TB of RAM to be installed.
From the front, you will notice up to eight hot-swap fans lined up, channeling airflow over critical components and keeping everything cool under pressure. Our system is equipped with directly connected NVMe SSDs, which are cabled directly to the motherboard.
This direct-attach NVMe offers the highest available NVMe performance, although for RAID, users must choose between software or hardware options like Graid.
At the rear, the PCIe slots are ready to accommodate GPUs or other expansion cards, while the OCP slots add an extra layer of versatility for specialized networking needs. The hot-swap power supplies are easy to access and replace, minimizing downtime during maintenance.
XClarity Controller 3
The Lenovo ThinkSystem SR630 V4 is equipped with the XClarity Controller 3 (XCC3) for remote and lifecycle management. It offers out-of-band management capabilities via a dedicated LAN port, allowing users to configure and deploy new hardware, interact with the system if primary networks are down, perform firmware management activities, and carry out countless other tasks.
Users can get a quick overview of all major components and warnings from the main home screen. System status, active events, and power are the key areas here.
You can see how the platform handles updates by exploring several areas, such as firmware update. You import a firmware package into local storage inside the XClarity controller and click Update System to start the process of updating the downloaded firmware payload.
Remote control is another common function that Lenovo XClarity handles well via an HTML5 web interface. This allows a wide range of client systems to manage the platform, including mobile platforms. Although an iPad is not the primary support mechanism, you sometimes have to use what's nearby.
Lenovo ThinkSystem SR630 V4 Performance
This section examines the performance test results of y-cruncher, Cinebench, Blackmagic, 7-Zip, and Geekbench. We compared the dual-processor Lenovo ThinkSystem SR630 V4 with the recently tested single-processor Supermicro Hyper 1U SYS-112H-TN. Both systems are equipped with the Intel Xeon 6780E processor, allowing us to see how the 6780E scales from single- to dual-processor configurations.
In addition to the Supermicro, we added an older Intel Ice Lake server, provided at the launch of the first Ice Lake Xeon 8380 processors. This shows how the E-core models position themselves as a cost-effective upgrade for legacy platforms that prioritize efficiency over raw processing power. This comparison between the dual-processor SR630 V4 platform and Intel Ice Lake and the single-processor Supermicro Hyper 1U illustrates the performance difference.
Here are the configurations for each system.
Lenovo ThinkSystem SR630 V4 Configuration
- CPU: 2 x Intel Xeon 6780E (144 cores)
- RAM: 512GB DDR5
- SSD Samsung MZWL6960HFJA-00AW7
- Operating System: Windows Server 2025
Supermicro Hyper 1U SYS-112H-TN Configuration
- CPU: Intel Xeon 6780E (144 cores)
- RAM: 512GB DDR5
- SSD Micron 7450 NVMe Data Center SSD
- Operating System: Windows Server 2022
Intel Ice Lake Server
- CPU: 2 x Intel Xeon 8380 (80 cores)
- RAM: 512GB DDR5
- SSD
- Operating System: Windows Server 2025
OptiX Blender
First, we will move on to the Blender test, an open-source 3D modeling application. This test was run using the Blender Benchmark utility. The score is expressed in samples per minute, with the higher being the better.
The Blender OptiX tests reveal interesting information about the systems tested. The Lenovo ThinkSystem SR630 V4 excelled with Blender 4.2.0, delivering 1,432 samples per minute in the Monster scene, while the Intel Ice Lake Server reached 569 and the Supermicro Hyper 1U 112H-TN (running Blender 4.0) reached 781. Lenovo reported 914 samples in the Junkshop scene, with 403 for the Intel Ice Lake Server and 514 for Supermicro. The Classroom scene showed Lenovo had 657 samples, Ice Lake with 280, and Supermicro with 371.
| Blender Processor | Supermicro Hyper 1U 112H-TN (1x Xeon 6780E, 512 GB DDR5) Blender 4.0 | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 GB) Blender 4.2 | Intel Ice Lake Server (2 x Intel Xeon 8380, 512 GB) Blender 4.2 | 
| Monster | 781.42 | 1432.09 | 569.10 | 
| Junkshop | 514.658 | 914.75 | 403.96 | 
| Classroom | 370.52 | 656.68 | 280.86 | 
Geekbench 6
Geekbench 6 is a cross-platform performance evaluation tool that measures the overall performance of a system. The Geekbench browser allows you to compare any system using this tool.
Lenovo's single-core test scored 1,173, while the dual-processor Supermicro recorded 1,154, highlighting its potential for tasks that depend on the efficiency of each core. In the multi-core test, Supermicro recorded 15,167 and Lenovo 13,868. Not all applications scaled well with the increase in core count. Geekbench struggled to scale from 144 to 288 cores on the dual-socket Lenovo SR630 V4 platform. The older Ice Lake processors showed higher single-core and multi-core scores, reaching 1,668 and 17,409 respectively.
| Geekbench 6 (Higher is Better) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 GB DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 GB) | Intel Ice Lake Server (2 x Intel Xeon 8380, 512 GB) | 
| Single-Core CPU | 1,154 | 1,173 | 1,668 | 
| Multi-Core CPU | 15,167 | 13,868 | 17,409 | 
Cinebench R23
The Cinebench R23 benchmarking tool evaluates a system's CPU performance by rendering a complex 3D scene using the Cinema 4D engine. It measures single-core and multi-core performance, providing a comprehensive view of the CPU's capabilities in handling 3D rendering tasks.
