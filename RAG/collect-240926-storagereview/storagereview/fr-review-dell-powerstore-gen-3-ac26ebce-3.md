---
id: collect-240926-storagereview/storagereview/fr-review-dell-powerstore-gen-3-ac26ebce-3
title: "fr-review-dell-powerstore-gen-3-ac26ebce"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["compute", "cost", "ethernet", "intel", "memory", "parameters", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-powerstore-gen-3-ac26ebce.md
source_anchor: ""
source_lines: [34, 65]
sha256: 644ccae9990f3da32b93f02af8e293e566581071ab76b24840998b1a5bd2df51
---

# fr-review-dell-powerstore-gen-3-ac26ebce

The choice between TLC and QLC comes down to price/GB and density rather than performance level. Minimum configurations vary slightly by model: TLC is available starting at 6 × 3.84 TB across the range, while QLC is available starting at 7 × 30.72 TB on the 1500 model and 11 × 30.72 TB on the 5500 and 9500 models. All drives are SED or FIPS certified, and the platform supports DRE +1 and +2 configurations. Thermal dissipation is also improved: Dell claims cooling requirements roughly 50% lower than an equivalent 2.5-inch deployment, thanks in part to the EDSFF airflow geometry. All 40 bays are now user-accessible for data, with cache managed by software-defined persistent memory (SDM) rather than NVRAM drives occupying front bays, as on generation 2.
|  | PowerStore Gen 2 | PowerStore Elite (Gen 3) | 
|---|---|---|
| Chassis and platform |  |  | 
| Base chassis | 2U, 2 nodes | 3U, 2 nodes | 
| I/O bus | PCIe generation 3 | PCIe generation 5 | 
| Memory | DDR4 | DDR5 | 
| Storage |  |  | 
| Drive form factor | Up to 25 2.5-inch U.2 NVMe SSDs (dual port) | Up to 40× E3.S/L 1T NVMe (dual port) | 
| Expandable shelf | Up to 24 2.5-inch U.2 NVMe SSDs | Up to 44× E3.S/L NVMe (post-RTS) | 
| Cache strategy | U.2 NVRAM drives | Cache to local flash memory (SDPM) | 
| I/O and network |  |  | 
| I/O slots | 3× SLIC / OCP 2.0 (PCIe Gen 3 x16) | Up to 5× OCP 3.0 (PCIe Gen 4/5 x16) | 
| Inter-node connection | 2× 10 GbE, RDMA | Up to 200 GbE, RDMA (400 GbE ready) | 
| Management and power |  |  | 
| Out-of-band management (BMC) | EMC GEM | iDRAC (storage version) | 
| Engine Tuning | EMC power supplies / custom power supplies | PowerEdge BBU / Power supplies | 
Inside the PowerStore Gen 3 Hardware Platform
Comparing the front of the new 3U PowerStore Gen 3 server to that of the Gen 2 model, the modernized design adopts the aesthetic of the 17th-generation PowerEdge servers we have already tested. It is offered in Dell's new gray tone, with the matching honeycomb bezel, which has become the visual signature of the brand's current enterprise hardware lineup.
Compute and Memory
On the compute side, the system is powered by Intel processors, with a per-processor TDP ranging from 165 W to 270 W and up to 32 cores per processor on the 9500 model. Dell claims up to 50% more cores per node. Memory moves to DDR5 across the range: the 9500 model ships with 2 TB (64 × 32 GB DIMMs), the 5500 with 1 TB, and the 1500 with 512 GB. The move to DDR5 significantly improves memory throughput, delivering twice the bandwidth of previous DDR4 generations.
The internal architecture moves to PCIe Gen 5, offering four times the per-lane bandwidth of the Gen 3 architecture used in PowerStore Gen 2. Socket configuration is the main differentiator within the lineup: the 1500 model is single-socket, while the 5500 and 9500 models are dual-socket. All use the same chassis and the same PowerStore OS software image, with core count and memory being the parameters used to tailor the configuration.
Chassis Cooling
The cooling system of the 9500, shown below, is an impressive design. The two processor coolers of a single controller are connected by heat pipes and form a large common fin stack, a clever solution given the amount of heat a dual-processor controller must dissipate. The 3U chassis, an earlier design choice, plays a crucial role in this cooling. Each controller is now 1.5U tall, giving air a longer path than with a 1U tray. This ensures each model has a comfortable cooling margin throughout the platform's lifespan. The 1500 uses a more conventional-looking processor cooler but retains all the airflow benefits of the 5500 and 9500 models. Regardless of the model, cooling across the lineup is provided by six fans per controller, for a total of 12 fans that keep the drives and the rest of the unit within their optimal thermal range.
Networking and Storage
Compared to previous generations, the new models adopt a modular, standardized I/O architecture based on OCP 3.0 slots. The 5500 and 9500 models offer up to 5 slots per node (one of which is reserved for expansion), while the 1500 model offers 3 (one of which is reserved). The modules are hot-swappable without tools and replace the proprietary SLIC carrier used in the second generation. Another important point: the I/O interface moves to PCIe Gen 5 on the new models, significantly increasing the per-lane bandwidth available to each card and pushing the limits of the network. At launch, card options include 4-port 32/64 Gb/s FC, 4-port 1/10 Gb/s Ethernet, 4-port 10/25 Gb/s Ethernet, and 2-port 100 Gb/s Ethernet. 200/400 Gb/s Ethernet and 128 Gb/s FC are planned for later releases. Dell is also strengthening network security: all Fibre Channel cards will support the EDIF (Encrypted Data-in-Flight) protocol through an upcoming non-disruptive software update. The result is up to 40 network ports per appliance, double the previous generation and roughly 11% higher port density than the second-generation controller.
The controller layout has also been redesigned. On generation 2, the top controller was inverted relative to the bottom one, which complicated maintenance and increased the risk of removing the wrong controller or handling the wrong component during a hot swap. On generation 3, both controllers are oriented the same way and can be removed interchangeably using the two handles and levers located on the sides of the chassis, sliding them out like a 1.5U tray. This change, though minor in appearance, is significant for ease of maintenance, particularly in racks where top access is limited or when the technician is working under pressure.
PCIe signaling plays an essential role in the new PowerStore Elite, drawing inspiration from the latest-generation PowerEdge servers. When Dell began offering E3.S Gen5 interface support on platforms such as the PowerEdge R770 or PowerEdge R7725, it decided to abandon the use of PCIe switches on the drive backplane. Older models, such as the PowerEdge R760 with a 24-drive backplane, used a PCIe switch to enable all PCIe lanes, which increased complexity and cost while reducing drive performance. On models using the E3.S interface, configurations with up to 16 SSDs allocated 4 PCIe lanes per drive, while configurations with up to 40 bays allocated 2. Dell applies the same philosophy to the new PowerStore Gen 3 models: each controller uses only 2 PCIe lanes to communicate with each dual-port E3.S SSD. The remaining PCIe Gen5 lanes are then used for inter-node communication and rear I/O connectivity. Nearly all PCIe lanes inside the chassis are used, down to the remaining PCIe Gen4 lanes from the processor chipset, which power the OCP slots that do not require high bandwidth.
Inter-Node Interconnect
The node interconnect represents one of the major advances over generation 2. The 5500 and 9500 models support up to 200 GbE RDMA between controllers (the 1500 model uses 100 GbE RDMA), compared to only 2 × 10 GbE for generation 2. The links are wireless point-to-point connections, routed through the backplane and dedicated to write ingestion. The image below shows one of the internal 200 GbE interconnect modules of the 9500 model.
Importantly, the new interconnect layer is processor-independent and decoupled from processor generations and vendors, allowing the chassis to accept upgrades to future processor platforms throughout the unit's lifecycle without disrupting the drives or rethinking the architecture.
Cache and Persistent Memory
