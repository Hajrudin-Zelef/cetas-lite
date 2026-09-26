---
id: collect-240926-storagereview/storagereview/fr-review-300-gb-s-in-2u-the-dell-poweredge-r7725xd-resets-expectations-for-stor-1e3ea655-3
title: "fr-review-300-gb-s-in-2u-the-dell-poweredge-r7725xd-resets-expectations-for-stor-1e3ea655"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Broadcom"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "inference", "latency", "memory", "throughput", "training"]
source: docs/RAG/clean_en/storagereview/fr-review-300-gb-s-in-2u-the-dell-poweredge-r7725xd-resets-expectations-for-stor-1e3ea655.md
source_anchor: ""
source_lines: [23, 56]
sha256: c6ee2083d517e434d0a439c8f0079b3dd79b083701390d4df9ee1f12933a9080
---

# fr-review-300-gb-s-in-2u-the-dell-poweredge-r7725xd-resets-expectations-for-stor-1e3ea655

On the right, the "Drive Summary" panel details the physical drives and associated virtual drives. Since the R7725xd uses a direct NVMe architecture without a traditional RAID controller, all drives are identified as non-RAID and individually addressable, consistent with the system's design for large NVMe pools and SDS platforms.
Below the health summary, the "Recently Recorded Storage Events" section lists the insertion logs for each PCIe SSD, sorted by bay and slot. This record confirms proper drive detection in all bays and allows identification of any insertion, cabling, or hot-swap issues. For large-scale deployments, these logs are useful for tracking drive provisioning or verifying that capacity has been correctly allocated.
The last screenshot presents the detailed view of NVMe devices in iDRAC10. Each NVMe drive installed in the system is listed with its status, capacity, and location. Selecting a drive displays the full detail of its characteristics.
In this example, the drive information panel displays the full model string, device protocol, format, and negotiated PCIe settings. The NVMe devices operate at a link speed of 32 GT/s with a negotiated x4 connection, confirming that the drives are utilizing the full bandwidth of the system's PCIe Gen5 backplane. The information section also indicates the endurance percentage, available drive status, and protocol type, allowing administrators to monitor drive health and expected lifespan.
This detailed drive report is valuable in high-density NVMe configurations where link width, negotiated speed, and media health directly influence workload behavior and storage performance.
Overall, the iDRAC 10 interface provides a clear, hardware-focused view of the R7725xd's NVMe storage architecture, enabling easy validation of link status, drive health, and system integrity at a glance.
Dell PowerEdge R7725xd Performance
Before testing, our system was configured with a balanced yet high-performance setup. It is equipped with two AMD EPYC 9575F processors, each with 64 high-frequency cores, and 24 x 32 GB DDR5 DIMM modules running at 6,400 MT/s. For storage, the chassis is fully populated with 24 Micron 9550 PRO U.2 NVMe SSDs of 15.36 TB each, each connected via a dedicated PCIe Gen5 x4 interface. This provides a total raw capacity of 368.64 TB, and the Micron 9550 PRO drives achieve sequential read speeds of up to 14,000 MB/s and sequential write speeds of up to 10,000 MB/s. Networking is provided by four Broadcom BCM57608 adapters that deliver a total of eight 200 Gb ports, as well as an OCP BCM57412 network card offering two additional 10-gigabit ports.
Test System Specifications
- CPU: 2 AMD EPYC 9575F processors with 64 high-frequency cores
- Memory: 24 x 32 GB DDR5 at 6400 MT/s
- Storage: 24 Micron 9550 PRO U.2 drives of 15.36 TB (each connected on 4 PCIe Gen5 lanes); compatible with drives up to 128 TB currently, and higher capacities to come.
- Network: 4 Broadcom BCM57608 2x200G network cards, 1 OCP BCM57412 2x10Gb network card
- Switch: Dell PowerSwitch Z9664
FIO Performance Benchmark
To measure the storage performance of the PowerEdge R7725xd, we used industry-standard metrics and the FIO tool. In this section, we focus on the following FIO benchmarks:
- Random 4K – 1M
- Sequential 4K – 1M
FIO – Local – Bandwidth
During local access tests on the 24 PCIe Gen5 NVMe drives of the Dell PowerEdge R7725xd, the system displays the expected performance of a platform where each drive is connected to the processors via a full PCIe Gen5 x4 link. Without a network layer, this is the pure internal throughput of Dell's Gen5 storage architecture and the PCIe bandwidth of the AMD EPYC platform utilized without restriction.
Sequential reads start at 184 GB/s with 4 KB blocks and increase rapidly with block size. From 512 KB to 1 MB, the server maintains a consistent throughput of 312 to 314 GB/s, demonstrating the system's excellent ability to aggregate 24 × 4 Gen5 lanes to achieve sustained read bandwidth, with no controller-level bottleneck.
Sequential writes follow a different curve but remain within the expected range. Starting at 149 GB/s, results increase to around one hundred GB/s and reach 182 GB/s at 1 million writes. This corresponds to the write behavior of Micron 9550 PRO SSDs and the inherent overhead of highly parallel NVMe writes across such a large number of independent devices.
Random read performance is another strong point. The system reaches speeds close to 300 GB/s with the smallest block sizes, decreases slightly in the middle, then rises again between 200 and 300 GB/s with larger block sizes. At 1 MB, random reads peak at 318 GB/s, demonstrating the platform's ability to evenly distribute mixed operations across all 24 drives.
Random writes are performed at lower throughput, which is typical of scattered metadata tasks and write allocation across a large NVMe set. Results remain between 140 and 160 GB/s for most of the test and decrease slightly to just under 100 GB/s at 1 Mbit/s.
FIO – Local – IOPS
From an IOPS perspective, the R7725xd demonstrates robust performance for small blocks, with request rates reaching several tens of millions before larger block sizes shift the workload toward a bandwidth-oriented profile.
At 4K, reads reach 44.9 million IOPS and writes 36.3 million. Random reads peak at 71.4 million IOPS, demonstrating the system's ability to efficiently distribute high-queue-depth workloads across all drives. These values naturally decrease as block size increases, but the progression remains consistent for 8K, 16K, and 32K sizes.
For 16 KB and 32 KB blocks, reads stabilize at 17.4 million and 8.35 million IOPS, with random reads showing similar values at 16.5 million and 8.15 million. Writes follow the expected trend, with a slight decrease but relative stability, whether for sequential or random access.
From 64 KB onward, the test shifts from a simple IOPS calculation to a scenario more limited by bandwidth. IOPS drop to a few million, then reach several hundred thousand. For a 1 MB block size, read IOPS are around 300,000, write IOPS around 174,000, and random operations fall within the same range of values.
Overall, the local IOPS results clearly demonstrate the system's ability to support very high queue-depth workloads on small blocks, with predictable scaling as transfers increase and bandwidth becomes the dominant factor.
PEAK:AIO: Why the Dell PowerEdge R7725xd is suited to this workload
PEAK:AIO is designed for environments requiring extremely fast, low-latency access to large datasets, particularly for AI training, inference pipelines, financial modeling, and real-time analytics. The platform takes full advantage of high-density NVMe storage, balanced PCIe bandwidth, and predictable latency at scale. To meet these requirements, the underlying hardware must provide sustained throughput while maintaining consistent and reproducible performance under heavy simultaneous loads.
This is where the Dell PowerEdge R7725xd naturally integrates with PEAK:AIO. Its architecture is designed to optimize PCIe Gen5 resources, exposing the full bandwidth of its 24 front U.2 NVMe bays directly to the processors, without resorting to traditional RAID controllers. This configuration gives PEAK:AIO the parallelism and latency expected from modern NVMe data pipelines. The NVMe SSDs are divided into two RAID 0 groups.
In the test scenario, we used two client systems connected to the R7725xd, each equipped with two Broadcom BCM57608 2x 200G network cards. This created a total of four 200G uplinks feeding each client, pushing the R7725xd into a realistic high-performance configuration, similar to production PEAK:AIO deployments. This level of network bandwidth allowed us to fully exercise the NVMe subsystem, PCIe topology, and processor interconnects without a bottleneck at the network card layer.
