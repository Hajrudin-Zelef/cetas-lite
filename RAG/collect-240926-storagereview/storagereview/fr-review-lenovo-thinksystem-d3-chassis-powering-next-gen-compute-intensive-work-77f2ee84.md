---
id: collect-240926-storagereview/storagereview/fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84
title: "fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["compute", "accelerator", "amd", "benchmark", "benchmarks", "dram", "energy", "ethernet", "gpu", "gpus", "intel", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84.md
source_anchor: ""
source_lines: [1, 154]
sha256: 06919fae4edacccbe458be74a67646cb5cb586b1bbe9a7e562399b63e7b718bf
---

# fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84

<!-- source: https://www.storagereview.com/fr/review/lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-workloads -->

We are examining Lenovo's three new multi-node servers, which share the new 2U ThinkSystem D3 chassis: the Intel-based ThinkSystem SD530 V3 (1U2S), the Intel-based ThinkSystem SD550 V3 (2U2S), and the AMD-based ThinkSystem SD535 V3 (1U1S). .
Lenovo ThinkSystem Multi-Node Servers
A multi-node server is a half-width modular server that fits into a common blade form factor chassis, enabling maximum compute density. Instead of having an entire 2U blade for two processors, you can have up to four dual-processor node servers in the same space.
Lenovo's new multi-node servers are aimed at dense compute workloads, such as energy-efficient transaction processing, cloud computing, HPC, big data analytics, hyperconverged infrastructure, and content delivery. Some of these servers have condensed storage and memory capacity to maximize compute density.
Here is the 2U ThinkSystem D3 chassis with three multi-node servers installed. The chassis has the unique ability to accommodate mixed 1U and 2U nodes and different processor types.
The D3 chassis supports three power supplies (typically, only two are installed) shared among all servers in the chassis. The D3 chassis power supplies have an 80 PLUS Platinum rating for energy efficiency. The servers are also ASHRAE A2 compliant for operation in data centers at 35 degrees Celsius. These ThinkSystem servers are designed to operate 24/7.
In addition to the power supplies, everything else (even cooling) is node-specific.
Thanks to hot-swappable drives and tool-less access to upgrades, such as fans, adapters, processors, and memory, Lenovo makes maintenance of these new models easy. These servers are also managed with Lenovo's integrated XClarity Controller management engine, which has a GUI and standard Redfish REST APIs.
The multi-node servers can be easily removed, as shown below.
Lenovo ThinkSystem D3 Chassis Specifications
| Components | Specifications | 
| Machine type | 7DD0 – 3-year warranty 7DD7 – 1-year warranty | 
| Form factor | 2U rack-mounted chassis | 
| Server support | Up to two SD550 V3 servers per chassis Up to four SD530 V3 servers per chassis Servers can be mixed in the same chassis | 
| Servers per rack | Up to 42 SD550 V3 servers in 21 chassis per 42U rack Up to 48 SD550 V3 servers in 24 chassis per 48U rack | 
| Systems management | None. Management is handled by each node. An incoming remote management connection can be shared using an optional daisy-chain adapter. | 
| Ports | None. | 
| I/O architecture | None integrated. Use top-of-rack network and storage switches. | 
| Power supply | Three hot-swappable power supplies power all nodes installed in the chassis. They are either 1300 1600 W, 2000 2700 W (certain markets), 1 80 W or 80 200 W, with N+240 redundancy. They are in CRPS form factor and are certified 50 PLUS Platinum or 60 PLUS Titanium. All installed power supplies must have identical part numbers. They require XNUMX-XNUMX V AC power, XNUMX or XNUMX Hz, and are installed at the rear of the enclosure. | 
| Power cords | One AC power cord for each power supply, C13 or C19, depending on the selected power supplies | 
| Cooling | None. Fans are located in each server node. | 
| Enclosure LEDs | Each power supply has AC, DC, and error LEDs. | 
| Hot-swappable parts | Power supply | 
| Limited warranty | Three-year customer-replaceable unit and on-site limited warranty with 9x5/NBD coverage. | 
| Service and support | Optional service upgrades are available through Lenovo services: 4-hour or 2-hour response time, 6-hour repair time, 1-year or 2-year warranty extension, and software support for Lenovo hardware and certain third-party applications. . | 
| Dimensions | Height: 87 mm (3.43 inches), depth: 898 mm (35.36 inches), width: 448 mm (17.64 inches). See Physical and Electrical Specifications for more details. | 
| Weight | Empty (without servers or power supplies): 11.8 kg (26.1 lb) Maximum (4 1U servers and 3 power supplies): 47.8 kg (105.4 lb) | 
Introducing Lenovo ThinkSystem SD530 V3 and SD550 V3
The ThinkSystem SD530 V3 and SD550 V3 are Intel servers based on 5th Gen Intel Scalable processors. The SD530 V3 is based on the 2U4N mini-node form factor, while the SD550 V3 is 2U2N. These servers can be mixed and matched in the ThinkSystem D3 chassis. For example, a customer may choose to have two SD530 V3s and one SD550 V3, although it is more common to have the same model in the same chassis. (A ThinkSystem D3 chassis can accommodate four SD530 V3s or two SD550 V3s.)
Let's start with the SD530 V3. Removing it from the 2U chassis, we can see that it has an asymmetrical shape; the cutout next to the rear panel is where the shared power supply connects. These servers are intended to be installed upside down on the right side of the 2U chassis, so there is no left or right version. The vapor-phase heat sinks were designed to work properly regardless of orientation.
Front ports include USB, serial, VGA, and a diagnostic port for the XClarity controller. The SD530 V3 supports two hot-swappable E3.S 1T EDSFF drives accessible from the front panel. There is also a pull-out tab for network information, the server model, and serial number.
At the rear, the SD530 V3 has one OCP 3.0 slot and one discrete PCIe x16 Gen 5 slot for GPU, network, or storage expansion. There is also a remote management Ethernet port, additional USB ports, and a mini-DisplayPort.
Access to the inside of the servers is much like that of a classic blade server, as the top panel slides backward and outward. Below are the sockets for the two 5th Gen Xeon Scalable processors. Each processor has eight DIMM slots for 16 DIMMs; the server supports 2 TB of memory with 128 GB 3DS RDIMMs. The SD530 V3 supports one processor up to 350 W/64 cores/3.9 GHz or two less powerful processors (each up to 205 W/32 cores/3.9 GHz). The rear processor has an additional heat sink, shown in the image below, Lenovo's Neptune thermal transfer module, for improved cooling, because the airflow reaching it is not as cool as that of the first processor.
Two M.2 drives at the front of the server are intended for boot drives, which allow RAID 1. There is also a Root of Trust module and a MicroSD card slot for local XCC storage.
Four 40 mm easy-swap cooling fans are located behind the processors and draw air from the front. It is rare in multi-node servers to place cooling fans in the node itself, which is why Lenovo has done something unique here. This means that the nodes do not share fans, making them more independent. This design makes the nodes the most efficient because their workload requirements change relative to one another.
Now, let's look at the SD550 V3. This server stands out from the SD530 V3 at the front by offering six 2.5-inch SSD drives (SAS, SATA, or NVMe). The front panel is otherwise the same as that of the SD530 V3. The rear of the server is similar to the SD530 V3, but the SD550 V3 adds a full-height/half-length PCIe x16 Gen 5 slot, allowing even more networking and GPU options.
Internally, the SD550 V3 has two sockets and eight DIMM slots per socket, just like the SD530 V3. The difference is that the SD550 V3 has much larger heat sinks and processor fans (three, not four), so it will run cooler and have higher energy efficiency. Both processors can reach 350 W/64 cores/3.9 GHz without a penalty for installing a second processor. The SD550 V3 also supports an internal RAID adapter.
Lenovo ThinkSystem SD530 V3 Specifications
| Components | Specifications | 
| Machine type | 7DD3 – 3-year warranty 7DDA – 1-year warranty | 
| Form factor | 1U half-width compute node. | 
| Supported enclosure | ThinkSystem D3 chassis, 2U height; up to 4 servers per chassis. | 
| Processor | One or two 5th Gen Intel Xeon Scalable processors (formerly named "Emerald Rapids"). With one processor installed, supports processors up to 64 cores, core speeds up to 3.9 GHz, and TDPs up to 350 W. With two processors installed, supports processors up to 32 cores, core speeds up to 3.9 GHz, and TDPs up to 205 W. | 
| Chipset | Intel C741 "Emmitsburg" chipset, part of the platform named "Eagle Stream." | 
| Memory | 16 DIMM slots with two processors (8 DIMM slots per processor) per node. Each processor has eight memory channels, with 1 DIMM per channel (DPC). Lenovo TruDDR5 and 3DS RDIMMs are supported up to 5600 MHz | 
| Persistent memory | Not supported | 
| Maximum memory | Up to 1 TB of system memory (using either 8 x 128 GB 3DS RDIMMs or 16 x 64 GB RDIMMs) | 
| Memory protection | ECC, SDDC, patrol/on-demand scrubbing, limited fault, DRAM address command parity with replay, DRAM uncorrected ECC error retry, on-chip ECC, ECC error checking and scrubbing (ECS), post-package repair | 
| Drive bays |  | 
| Maximum internal storage | 30.72 TB with 2 x 15.36 TB E3.S EDSFF NVMe SSDs | 
| Storage controller | 2x integrated NVMe ports (optional Intel VROC NVMe for RAID) | 
| Optical drive bays | There are no internal bays; use an external USB key. | 
| Tape drive bays | There are no internal bays. Use an external USB key. | 
| Network interfaces | Dedicated OCP 3.0 SFF slot with PCIe 5.0 x16 host interface. Supports a variety of 2- and 4-port adapters with 1, 10, 25, or 100 GbE network connectivity. Optionally, one port can be shared with the XClarity Controller 2 (XCC2) management processor for Wake-on-LAN and NC-SI support. | 
| PCIe slots | One PCIe 5.0 x16 slot with a discrete form factor | 
| GPU support | Supports 1x single-width GPU | 
| Ports | Front: one VGA port for video, one USB 3.2 G1 port (5 Gb/s), one external diagnostic port, and one DB9 serial port for local connectivity Rear: one MiniDP port for video, one USB 3.2 G1 port (5 Gb/s), 1 USB 2.0 port (also for local XCC management), 1 RJ-45 1GbE systems management port for remote XCC management | 
| Cooling | 4 x 40 mm easy-swap dual-rotor fans with N+1 rotor redundancy | 
| Power source | Provided by the D3 chassis. | 
| Hot-swappable parts | Drives | 
| Systems management | Control panel with status LEDs. Optional external diagnostic combo with LCD display. Integrated XClarity Controller 2 (XCC2) management based on the ASPEED AST2600 baseboard management controller (BMC), XClarity Administrator centralized infrastructure provisioning, XClarity Integrator plug-ins, and XClarity Energy Manager centralized server power management - XCC Platinum to enable remote control functions and other features. | 
| Video | Embedded graphics with 16 MB of memory with 2D hardware accelerator, integrated into the XClarity Controller 2 management controller. Two video ports, front VGA and rear Mini DisplayPort. Both ports can be used simultaneously if desired. The maximum resolution of both ports is 1920x1200 at 60 Hz. | 
| Security | Power-on password, administrator password, Trusted Platform Module (TPM), supporting TPM 2.0. | 
| Supported operating systems | Microsoft Windows Server, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware ESXi, and Ubuntu Server. For details, including other vendor-certified or tested operating systems, see the operating system support section. | 
| Limited warranty | Three-year customer-replaceable unit and on-site limited warranty with 9x5 next business day (NBD). | 
| Service and support | Optional service upgrades are available through Lenovo services: 4-hour or 2-hour response time, 6-hour repair time, 1-year or 2-year warranty extension, software support for Lenovo hardware and certain third-party applications. | 
| Ambient temperature | Up to ASHRAE Class A2: 10°C – 35°C (50°F – 95°F) | 
| Dimensions | Width: 222 mm (8.7 inches), height: 41 mm (1.6 inches), depth 908 mm (35.7 inches) | 
| Weight | Maximum: 7.6 kg (16.76 lb) | 
Lenovo ThinkSystem SD550 V3 Specifications
| Components | Specifications | 
| Machine type | 7DD2 – 3-year warranty 7DD9 – 1-year warranty | 
| Form factor | 2U half-width compute node. | 
| Supported enclosure | ThinkSystem D3 chassis, 2U height; up to 2 SD550 V3 servers per chassis. | 
| Processor | One or two 5th Gen Intel Xeon Scalable processors (formerly named "Emerald Rapids"). Supports processors up to 64 cores, core speeds up to 3.9 GHz, and TDPs up to 350 W. | 
| Chipset | Intel C741 "Emmitsburg" chipset, part of the platform named "Eagle Stream" | 
| Memory | 16 DIMM slots with two processors (8 DIMM slots per processor) per node. Each processor has eight memory channels, with 1 DIMM per channel (DPC). Lenovo TruDDR5 and 3DS RDIMMs are supported up to 5600 MHz | 
| Persistent memory | Not supported | 
| Maximum memory | Up to 2 TB using 16 x 128 GB 3DS RDIMMs | 
| Memory protection | ECC, SDDC, patrol/on-demand scrubbing, limited fault, DRAM address command parity with replay, DRAM uncorrected ECC error retry, on-chip ECC, ECC error checking and scrubbing (ECS), post-package repair | 
| Drive bays |  | 
| Maximum internal storage | 92.16 TB with 6 x 15.36 TB 2.5-inch SAS/SATA SSDs 92.16 TB with 6 x 15.36 TB 2.5-inch NVMe SSDs | 
| Storage controller | Integrated NVMe ports (optional Intel VROC NVMe for RAID) Integrated SATA ports Support for internal RAID adapter (CFF) for SAS/SATA drive support | 
| Optical drive bays | No internal bays; use an external USB key. | 
| Tape drive bays | No internal bays. Use an external USB key. | 
| Network interfaces | Dedicated OCP 3.0 SFF slot with PCIe 5.0 x16 host interface. Supports a variety of 2- and 4-port adapters with 1, 10, 25, or 100 GbE network connectivity. Optionally, one port can be shared with the XClarity Controller 2 (XCC2) management processor for Wake-on-LAN and NC-SI support. | 
| PCIe slots | One or two full-height half-length (FHHL) PCIe x16 slots, depending on the number of processors installed: 1-processor configurations:  2-processor configurations:  | 
| GPU support | Supports 2x single-width GPUs | 
| Ports | Front: one VGA port for video, one USB 3.2 G1 port (5 Gb/s), one external diagnostic port, and one DB9 serial port for local connectivity Rear: one MiniDP port for video, one USB 3.2 G1 port (5 Gb/s), 1 USB 2.0 port (also for local XCC management), 1 RJ-45 1GbE systems management port for remote XCC management | 
| Cooling | 3 x 60 mm easy-swap dual-rotor fans with N+1 rotor redundancy | 
| Power source | Provided by the D3 chassis. | 
| Hot-swappable parts | Drives | 
| Systems management | Control panel with status LEDs. Optional external diagnostic combo with LCD display. Integrated XClarity Controller 2 (XCC2) management based on the ASPEED AST2600 baseboard management controller (BMC), XClarity Administrator centralized infrastructure provisioning, XClarity Integrator plug-ins, and XClarity Energy Manager centralized server power management - optional XCC Platinum to enable remote control functions and other features. . | 
| Video | Embedded graphics with 16 MB of memory with 2D hardware accelerator, integrated into the XClarity Controller 2 management controller. Two video ports (front VGA and rear Mini DisplayPort); both can be used simultaneously if desired. The maximum resolution of both ports is 1920x1200 at 60 Hz. | 
| Security | Power-on password, administrator password, Trusted Platform Module (TPM), supporting TPM 2.0. | 
| Supported operating systems | Microsoft Windows Server, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware ESXi, and Ubuntu Server. For details, including other vendor-certified or tested operating systems, see the operating system support section. | 
| Limited warranty | Three-year customer-replaceable unit and on-site limited warranty with 9x5 next business day (NBD). | 
| Service and support | Optional service upgrades are available through Lenovo services: 4-hour or 2-hour response time, 6-hour repair time, 1-year or 2-year warranty extension, software support for Lenovo hardware and certain third-party applications. | 
| Ambient temperature | Up to ASHRAE Class A2: 10°C – 35°C (50°F – 95°F) | 
| Dimensions | Width: 222 mm (8.7 inches), height: 82 mm (3.2 inches), depth 898 mm (35.4 inches) | 
| Weight | Maximum: 11.76 kg (25.93 lb) | 
Introducing the Lenovo ThinkSystem SD535 V3
The ThinkSystem SD530 V3 is an AMD version of the Intel-based ThinkSystem SD530 V3. The main difference is that it supports only one processor, but this is not necessarily a disadvantage since its single AMD EPYC 9004 series chip can have up to 128 cores, which is more than the SD530 V3 can support even when equipped with two processors. The ThinkSystem SD535 V3 can support more cores in total than the SD530 V3; the latter supports only 64, even with two processors installed. The SD535 V3 also supports more memory per socket, with 12 DIMM slots supporting 1.5 TB using 128 GB 3DS RDIMMs.
Here, we can see the single CPU socket and the surrounding DIMM slots. The air cooling of the SD535 V3 is the same as that of the SD530 V3, with four 40 mm easy-swap rear fans. Integrated storage includes two M.2 boot drives. A SAS/SATA RAID adapter is also supported for configurations with six front 2.5-inch SSDs.
The front and rear ports of the SD535 V3 are identical to those of the SD530 V3, with USB, serial, VGA, and a diagnostic port for the XClarity controller; similarly, at the rear, you will find remote management Ethernet, two USB-A ports, a mini-DisplayPort video output, an OCP 3.0, and a PCIe x16 Gen 5 PCIe slot.
Lenovo ThinkSystem SD535 V3 Specifications
| Components | Specifications | 
| Machine type | 7DD1 – 3-year warranty 7DD8 – 1-year warranty | 
| Form factor | 1U half-width compute node. | 
| Supported enclosure | ThinkSystem D3 chassis, 2U height; up to 4 servers per chassis. | 
| Processor | One 4th Gen AMD EPYC 9004 processor. Supports processors up to 128 cores, core speeds up to 4.1 GHz, and TDPs up to 400 W. | 
| Chipset | Not applicable (platform controller hub functions are integrated into the processor SOC) | 
| Memory | 12 DIMM slots with 12 processor memory channels (1 DIMM per channel, DPC). Lenovo TruDDR5 and 3DS RDIMMs are supported up to 4800 MHz | 
| Persistent memory | Not supported | 
| Maximum memory | Up to 1.5 TB of system memory using 12 x 128 GB 3DS RDIMMs | 
| Memory protection | ECC, SDDC, patrol/on-demand scrubbing, limited fault, DRAM address command parity with replay, DRAM uncorrected ECC error retry, on-chip ECC, ECC error checking and scrubbing (ECS), post-package repair | 
| Drive bays |  | 
| Maximum internal storage |  | 
| Storage controller | Integrated NVMe ports (no RAID support) Integrated SATA ports (no RAID support) Support for an internal RAID adapter (CFF) or a PCIe adapter for SAS/SATA drive support | 
| Optical drive bays | No internal bays; use an external USB key. | 
| Tape drive bays | No internal bays. Use an external USB key. | 
| Network interfaces | Dedicated OCP 3.0 SFF slot with PCIe 5.0 x16 host interface. Supports a variety of 2- and 4-port adapters with 1, 10, 25, or 100 GbE network connectivity. One port can be shared with the XClarity Controller 2 (XCC2) management processor for Wake-on-LAN and NC-SI support. | 
| PCIe slots | One PCIe 5.0 x16 slot with a discrete form factor | 
| GPU support | Supports 1x single-width GPU | 
| Ports | Front: none; Rear: one MiniDP port for video, one USB 3.2 G1 port (5 Gb/s), 1 USB 2.0 port (also for local XCC management), 1 RJ-45 1GbE systems management port for remote XCC management | 
| Cooling | 4 x 40 mm easy-swap dual-rotor fans with N+1 rotor redundancy | 
| Power source | Provided by the D3 chassis. | 
| Hot-swappable parts | Drives | 
| Systems management | Control panel with status LEDs. Integrated XClarity Controller 2 (XCC2) management based on the ASPEED AST2600 baseboard management controller (BMC), XClarity Administrator centralized infrastructure provisioning, XClarity Integrator plug-ins, and XClarity Energy Manager centralized server power management - optional XCC Platinum to enable remote control functions and other features. . | 
| Video | Embedded graphics with 16 MB of memory with 2D hardware accelerator, integrated into the XClarity Controller 2 management controller. Rear Mini DisplayPort video port. The maximum resolution of both ports is 1920x1200 at 60 Hz. | 
| Security | Power-on password, administrator password, Trusted Platform Module (TPM), supporting TPM 2.0. | 
| Supported operating systems | Microsoft Windows Server, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware ESXi, and Ubuntu Server. For details, including other vendor-certified or tested operating systems, see the operating system support section. | 
| Limited warranty | Three-year customer-replaceable unit and on-site limited warranty with 9x5 next business day (NBD). | 
| Service and support | Optional service upgrades are available through Lenovo services: 4-hour or 2-hour response time, 6-hour repair time, 1-year or 2-year warranty extension, software support for Lenovo hardware and certain third-party applications. | 
| Ambient temperature | Up to ASHRAE Class A2: 10°C – 35°C (50°F – 95°F) | 
| Dimensions | Width: 222 mm (8.7 inches), height: 41 mm (1.6 inches), depth 898 mm (35.4 inches) | 
| Weight | Maximum: 8.32 kg (18.34 lb) | 
ThinkSystem Ordering and Customization
Lenovo's ThinkSystem SD530 V3, SD550 V3, and SD535 V3 servers are configurable through Lenovo's Data Center Solution Configurator (DCSC) . Base configurations are available, including for general purpose and for AI and high-performance computing (HPC). Custom configurations can be created using the System x and Cluster Solutions configurator . Base models come with a one- or three-year warranty.
ThinkSystem Performance and Benchmarks
Although we are not (yet) testing these new ThinkSystem multi-node servers in our lab, Lenovo has performed benchmarks on them and broken several records. You can view the results via the links below.
- 
  - ThinkSystem SD535 V3 sets 2 world records with new SPECPower benchmark result on Windows
  - ThinkSystem SD535 V3 sets a world record with new SPECPower benchmark result on Linux
  - ThinkSystem SD550 V3 sets 6 world records with new SPECjbb benchmark result on Linux and Windows
  - ThinkSystem SD535 V3 sets 6 world records with new SPECjbb benchmark result on Linux and Windows
  - ThinkSystem SD530 V3 and SD550 V3 deliver impressive multi-node SPECpower benchmark results
Final Thoughts
The ThinkSystem SD530 V3, SD550 V3, and SD535 V3 deliver excellent performance for compute-intensive workloads. A single 2U ThinkSystem D3 chassis can accommodate four SD530 V3 or SD535 V3 servers or two SD550 V3 servers, providing incredibly high compute density compared to traditional blade servers. It is important to note that these servers are truly focused on compute density and are therefore not ideal for storage- or GPU-intensive applications.
(For that, see our related article on Lenovo ThinkSystem SR685a V3 and SR680a V3 GPU servers )
Reliability and availability are also paramount on these servers. The only components that the nodes share in the ThinkSystem D3 chassis are the power supplies, with everything else, even cooling, being node-specific, so component failures are localized to the node itself.
Even though we have not yet performed performance testing on the new ThinkSystem servers, Lenovo has already published record-breaking results in several applications. Overall, we find these new ThinkSystem multi-node servers very impressive for their intended use, and we look forward to bringing them into our lab.
