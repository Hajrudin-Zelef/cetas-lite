---
id: collect-250926-servers-hardware/servers-hardware/aic-gets-flashy-with-32-ssd-bay-jbof-server-for-key-value-caching
title: "aic-gets-flashy-with-32-ssd-bay-jbof-server-for-key-value-caching"
domain: servers-hardware
role: reference
task: reference
actors: ["Broadcom", "Nvidia"]
dates: []
keywords: ["alignment", "compute", "ethernet", "gpu", "gpus", "inference", "kv cache", "memory", "nand", "nvidia", "rubin", "vera rubin"]
source: docs/RAG/clean4/aic-gets-flashy-with-32-ssd-bay-jbof-server-for-key-value-caching.md
source_anchor: ""
source_lines: [1, 47]
sha256: cc354962fb817f4354243667d9cf9e687c63cafb72a2e4105771eaab3923f1cf
---

# aic-gets-flashy-with-32-ssd-bay-jbof-server-for-key-value-caching

Although the recent Computex tradeshow certainly came with a heavy emphasis on processors and compute servers of all kinds, not everything in the server world is solely about computational workloads. Data storage has always been an important part of server design, and fast storage in particular has taken on an increased importance in recent years as AI servers not only need significant amounts of storage for the models they operate from, but these days they need significant dynamic storage as well for caching key-value information.

To that end, at this year’s show, AIC was showing off its F2032-01-G6 or F2032-G6, just a bunch of flash (JBOF) appliance, a 2U storage server that can hold 32 PCIe Gen6 E3 drives and is designed to be backed by NVIDIA’s BlueFiend-4 DPUs. When fully equipped with 256TB SSDs, a single F2032-G6 can store up to 8PB of data.

## AIC F2032-01-G6 JBOF Server

The F2032-01-G6 is a successor of sorts to the company’s previous JBOF server, the PCIe Gen5-era F2026-01-G5. While that system used an older generation of PCIe technology and housed 26 2.5-inch U.2 bays for storage, the F2032-01-G6 is a complete modernization of the platform, upgrading the storage, I/O, and networking capabilities to the latest generation of technologies. The latest JBOF box from AIC also brings the company closer in alignment to NVIDIA’s Context Memory Storage (CMX) platform standard, which, in turn, is designed to be integrated into Vera Rubin racks.

From a storage perspective, the F2032-01-G6 crams just about as much storage as possible into a 2U server. AIC has designed this latest JBOF appliance to use E3-form-factor SSDs, with the system able to accommodate both long (E3.L) and short (E3.S) versions of the 7.5mm-thick SSDs.

The resulting design of the server means that the bulk of the front of the system is the hot swappable storage trays, with thin vents separating every pair of drives to provide for some airflow to the rest of the system.

With E3.L drives available in capacities up to 256TB each, the F2032-01-G6 offers immense storage. For data centers happy with QLC memory, this puts the current limit on the box at 8PB in 2U, while TLC NAND can be had at around half that capacity.

Meanwhile, the system’s internal I/O fabric is built around Broadcom’s massive 144-lane PEX90144 PCIe Gen6 switch. While AIC did not have its showfloor server fully opened at Computex, we had the opportunity to look at the Dell PowerEdge system with the controller back at SC25.

AIC has not published a detailed architectural diagram for the upcoming storage server, but based on the previous-generation box (which used Broadcom’s PCIe Gen5 equivalent), AIC is likely routing two lanes to each SSD to maximize the system’s storage capacity. In which case, the bump to PCIe Gen6 will help offset the bandwidth deficit that a x2 configuration would otherwise entail.

As for the networking side of things, AIC has designed the server to accommodate a mix of NICs and DPUs, though the emphasis is clearly on the latter. Specifically, the system is driven by a pair of nodes, each with its own power supply and capable of housing up to two network expansion devices (or, interestingly enough, even a GPU).

The JBOF appliance is designed to be controlled by at least one of NVIDIA’s BlueField-4 DPUs. The Grace CPU + ConnectX-9 NIC-based DPUs feature 800Gb Ethernet (or InfiniBand) networking, and like the Broadcom PCIe switch at the heart of the I/O fabric, implement PCIe Gen6 for their I/O connections. These DPUs serve as the brains of the system, acting as storage controllers and gatewaying all access to the SSDs beyond them.

AIC is specifically positioning the F2032-01-G6 as a high-availability system. So the support for DPUs in multiples of two is not a coincidence. With up to two DPUs forming a single node, the second node allows for fully redundant operation so that even if one DPU (or part of one) fails, the JBOF appliance is able to stay up.

With the heavy reliance on NVIDIA hardware here to operate the F2032-01-G6, the chief use for the JBOF appliance is to serve as a KV cache as part of NVIDIA’s Context Memory Storage platform, which is NVIDIA’s homegrown KV caching solution for backing Vera Rubin systems.

With the server’s flexible configuration, AIC is also pitching it as a solution for more general use cases, such as object storage, scale-out file systems, and flash storage tiering.

### Final Words

With sizable KV caches quickly becoming a preferred (if not mandatory) component for large inference clusters to bolster their performance, JBOF servers such as AIC’s F2032-01-G6 are slated to become an increasingly common sight within AI data centers. Storing key values close to the GPUs performing inference is increasingly critical to scaling performance, which means placing SSDs as close to those servers as is reasonably possible, a role the F2032-01-G6 was explicitly designed to fill. AIC, in turn, is well positioned to capture a chunk of this market, leveraging its experience with PCIe Gen5 and BlueField-3-era hardware to hit the ground running with the BlueField-4/Vera Rubin generation.

With E3.L drives available in capacities up to 256GB each, the F2032-01-G6 offers immense storage.

Feels like it’s 2006.

Well 2006 was a very good year. So that’s not necessarily a bad thing.

(Though to be sure, the typo is fixed. Thanks!)

Somebody’s Postgres would absolutely scream on this.

Excellent overview of an innovative storage solution designed for modern AI and high-performance workloads. The article does a great job explaining how JBOF architecture and high-density NVMe storage can improve key-value caching performance in today’s data centers.

interesting that it has no traditional CPU, just the ARM of the bluefield running the show.

I’m now wondering what happens if you directly attach a BF2 to a PCIe switch…
