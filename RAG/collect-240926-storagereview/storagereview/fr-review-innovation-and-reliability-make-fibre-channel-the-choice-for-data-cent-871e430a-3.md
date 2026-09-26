---
id: collect-240926-storagereview/storagereview/fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a-3
title: "fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a"
domain: storagereview
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "compute", "latency", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a.md
source_anchor: ""
source_lines: [27, 61]
sha256: 3b87a020d559cc20539f192cd3f1e941a358ff6842e9ca22e9cda11eed8f4fd3
---

# fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a

FPIN is a notification frame transmitted by a fabric port to notify an end device of a condition for another port in its zone. Conditions include the following:
- Link integrity issues that degrade performance
- Lost frame notifications
- Congestion issues
Through a proactive notification mechanism, port issues can be resolved quickly and recovery actions can be defined to mitigate downtime.
Figure 1: vSphere 8.0 with Marvell QLogic FC registers and receives fabric notifications indicating oversubscription in the SAN
Figure 2: vSphere 8.0 with Marvell QLogic FC registers and receives fabric notifications indicating link integrity degradation.
Marvell QLogic Enhanced 16GFC, Enhanced 32GFC, and 64GFC HBAs are fully integrated with vSphere 8.0 and support fabric notification technology that serves as a building block for autonomous SANs.
Productivity with vVols
VMware has emphasized vVols in recent vSphere releases. With vSphere 8.0, core storage integrated vVols support for NVMe-oF, with FC-NVMe support initially limited. VMware will, however, continue to validate and support other protocols compatible with vSphere NVMe-oF. The new vVols specification, the VASA/VC framework, is available here.
With the industry and many array vendors adding NVMe-oF support for improved performance and lower latency, VMware wanted to ensure that vVols stay up to date with recent storage technologies.
In addition to improved performance, configuring NVMe-oF vVols is simplified. Once VASA is registered, the underlying configuration is done in the background; all that's left is to create the datastore. VASA manages all connections to virtual protocol endpoints (vPE). Customers can now manage NVMe-oF storage arrays in a vVols datastore via storage policy-based management in vCenter. vSphere 8 also supports additional namespaces and paths and improves vMotion performance.
Tracking virtual machines with VM-ID technology
Server virtualization was the catalyst for increased link sharing, as evidenced by Fibre Channel. With the growing number of virtual machines (VMs) in the data center, shared links carry data associated with CPU cores, memory, and other system resources, using the maximum available bandwidth. Data sent from any virtual machine and other physical systems gets mixed together. This data travels along the same path as storage network (SAN) traffic, so everything looks the same and cannot be considered as individual data streams.
Using Marvell QLogic's VM-ID (an end-to-end solution using frame tagging to associate different VMs and their I/O flows on the SAN) makes it possible to decipher each VM on a shared link. QLogic has enabled this capability on its latest 64GFC, Enhanced 32GFC, and Enhanced 16GFC host bus adapters (HBAs). This technology has a built-in application services monitor, which collects the unique worldwide ID of VMware ESX. It can then interpret the different identifiers of each virtual machine to perform intelligent monitoring.
VM-ID brings deep visibility into I/O from the originating VM to the fabric, giving SAN managers the ability to control and direct application-level services to each virtual workload within a QLogic Fibre Channel HBA.
Figure 3: The Brocade switch analytics engine can now display per-virtual-machine statistics by counting Fibre Channel frames tagged with an individual virtual machine ID by the Marvell Fibre Channel HBA.
Increased performance with 64GFC
Fibre Channel progress has continued since the protocol's launch in 1988. The first FC SAN products, 1G FC, began shipping in 1997, and evolution continues today, with 128G products on the horizon.
Every three to four years, FC speed doubles. In addition to advances in increased performance, new services such as Fabric Services, StorFusion™ with Universal SAN Congestion Mitigation, NPIV (virtualization), and cloud services have been included. Networking companies and OEMs participate in the development of these standards and continue to work together to deliver reliable and scalable storage network products.
Fibre Channel is considered the most reliable storage connectivity solution on the market, with a tradition of incremental improvements. Server and storage technologies are driving demand for greater SAN bandwidth. Application and storage capacity, 32G and 64G storage arrays supporting SSDs and NVMe, server virtualization, and multi-cloud deployments prove the value of Fibre Channels as they deliver higher throughput, lower latency, and higher link speeds, all with predictable performance.
Marvell recently announced the introduction of its all-new 64GFC HBAs. These include the QLE2870 series of one-, two-, and four-port FC HBAs that double the available bandwidth, run on a faster PCIe 4.0 bus, and simultaneously support FC and FC-NVMe, ideal for time-tested mission-critical applications.
NVMe delivers!
There's no doubt that NVMe devices offer extremely fast read and write access. So the discussion is about connecting these NVMe devices to high-speed networks without considering that they are still storage devices requiring guaranteed delivery. The technology that was developed as a lossless delivery method is Fibre Channel. In many tests conducted by industry leaders, NVMe-oF and NVMe/FC performed better when the underlying technology was Fibre Channel.
Flash arrays enable faster block storage performance in high-density virtualized workloads and reduce response time for data-intensive applications. All of this sounds pretty good, unless the network infrastructure can't perform at the same level as the flash storage arrays.
Flash-based storage requires a deterministic, low-latency infrastructure. Other storage network architectures often increase latency, creating bottlenecks and network congestion. More packets must be sent when this happens, creating even more congestion. With Channel's credit-based flow control, data can be delivered as fast as the destination buffer can receive it, eliminating dropped packets or forcing retransmissions.
We published an in-depth analysis earlier this year of Marvell's FC-NVMe approach. To learn more, see the article "Marvell strengthens its commitment to FC-NVMe technology."
Key NVMe storage features introduced for the first time in vSphere 7.0
A VMware blog described NVMe over Fabrics (NVMe-oF) as a protocol specification that connects hosts to high-speed flash storage via network fabrics using the NVMe protocol. VMware introduced NVMe-oF in vSphere 7 U1. The VMware blog stated that benchmark results showed that Fibre Channel (FC-NVMe) consistently outperformed SCSI FCP in vSphere virtualized environments, delivering higher throughput and lower latency. NVMe over TCP/IP support was added in vSphere 7.0 U3.
Based on the growing adoption of NVMe, VMware added support for shared NVMe storage using NVMe-oF. Given the inherent low latency and high throughput, industries are leveraging NVMe for AI, ML, and compute workloads. Typically, NVMe used a local PCIe bus, which made it difficult to connect to an external array. At the time, the industry had proposed external connectivity options for NVMe-oF based on IP and FC.
In vSphere 7, VMware added support for shared NVMe storage using NVMe-oF with NVMe over FC and NVMe over RDMA.
Fabrics continue to deliver higher speeds while maintaining the guaranteed lossless delivery required by storage networks. vSphere 8.0 supports 64GFC, the fastest FC speed to date.
vSphere 8.0 advances its NVMe goal, prioritizing Fibre Channel
vVols have been the primary focus of VMware storage engineering in recent releases, and with vSphere 8.0, core storage added support for vVols in NVMe-oF. Initially, VMware will only support FC but will continue to validate and support other NVMe-oF protocols.
This is a new vVols specification, for the VASA/VC framework. To learn more, see the VASA 4.0/vVols 3.0 documentation.
