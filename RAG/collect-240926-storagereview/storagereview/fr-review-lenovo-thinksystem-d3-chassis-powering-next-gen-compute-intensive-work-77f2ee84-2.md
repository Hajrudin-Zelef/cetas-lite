---
id: collect-240926-storagereview/storagereview/fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84-2
title: "fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["compute", "amd", "energy", "ethernet", "gpu", "intel", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84.md
source_anchor: ""
source_lines: [3, 44]
sha256: 427ef9bc9116a6fb27c5567e699e9db22f866df4f08fc92d940c14435d03291f
---

# fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84

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
