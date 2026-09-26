---
id: collect-240926-storagereview/storagereview/fr-review-hci-intel-qlc-ssd-277e2f61-3
title: "fr-review-hci-intel-qlc-ssd-277e2f61"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Microsoft", "Samsung"]
dates: []
keywords: ["intel", "benchmark", "cost", "latency", "lean", "memory", "nand"]
source: docs/RAG/clean_en/storagereview/fr-review-hci-intel-qlc-ssd-277e2f61.md
source_anchor: ""
source_lines: [22, 56]
sha256: 36c1e5be33b05bf824e6e19378721f400bb718a19405ab5df07a44f4990f6138
---

# fr-review-hci-intel-qlc-ssd-277e2f61

The Azure Stack HCI cluster we are working with here was built on the DataON HCI-224 NVMe all-flash platform. These servers had a 24U form factor with 1 NVMe bays at the front, offering many expansion options at the rear for PCIe-based components. The labeling was extensive unlike the matte black drive caddies, which made it easy to spot specific drives in case of a needed replacement. Everything was labeled, which is not so rare, but the extent of the labeling was extraordinary. Our deployment had each node labeled (2 and XNUMX), as well as several other elements, making it easy to deploy and manage DataON systems in the data center.
The nodes in this test included two 2nd Generation Intel® Xeon® Scalable Gold 6248 processors at 2.5 GHz, 20 cores, and 28 MB cache, along with eight Samsung 32 GB DDR4 2933 MHz registered ECC RDIMM modules (256 GB total per node) and two 480 GB Intel S4510 SATA M.2 boot drives.
For storage, each node came with four 750 GB 2.5-inch Intel Optane SSD DC P4800X NVMe drives (used for caching) and four 15.36 TB 2.5-inch Intel SSD D5-P4326 QLC drives (capacity storage tier).
The nodes were connected to each other via Mellanox ConnectX-4 EN dual-port QSFP28 40/56 GbE cards using 3M Mellanox LinkX ETH 40GbE, 40Gb/s, QSFP passive copper cables.
Clearly, DataON spent a lot of time and thought on the configuration and component selection of this system to balance performance and cost. We were very interested to see how the Intel SSD D5-P4326 SSDs would perform as a storage tier. By combining Intel Optane SSDs and Intel QLC 3D NAND SSDs, the D5-P4326 SSDs should provide high-performance, economical flash storage, which was once the domain of slow but large hard drives.
In the StorageReview lab, we deployed the two storage nodes and the switches as shown below.
Testing
To get an idea of how a small cluster like this can perform in an edge use case, we set up several Microsoft SQL Server tests. The goal was to examine the overall performance of the cluster to ensure that DataON could properly leverage Intel Optane technology and Intel QLC SSDs. Second, we wanted to examine the capabilities of a single node, to get an idea of how this solution handles the loss of a node, whether for planned updates or in the event of a more serious failure.
Our test plan relied on Quest's Benchmark Factory using the TPC-C profile as the load generator for the SQL Server virtual machines we deployed. We configured eight virtual machines (four per node), which offered a good balance between CPU and disk activity for the cluster. The workload generators were hosted on a system outside this environment and connected to this cluster via a 10GbE network.
SQL Server Test Configuration (per virtual machine)
- Windows Server 2019
- Storage footprint: 800 GB allocated, 620 GB used
- 8 vCPU
- 60 GB RAM (55 GB in failover mode configuration)
- SQL Server 2019
  - Database size: 1,500 XNUMX scale
  - Virtual client load: 15,000 XNUMX
  - Buffer memory: 48 GB
- Test duration: 3 hours
  - 15 minutes preconditioning
  - 45-minute sampling period
In our tests, we focused on latency performance, with the transaction performance level remaining constant with Benchmark Factory.
With a load of 4 VMs total (2 per node), we measured an average latency of 2.5 ms with an aggregate transaction load of 12,649 XNUMX TPS.
By increasing the load to 6 VMs, the average latency increased slightly to 4 ms with an overall transaction load of 18,967 XNUMX TPS.
At peak load of 8 VMs (4 per node), latency reached an average of 6.5 ms, with a total transaction load of 25,277 XNUMX.
Throughout these tests, we clearly saw the benefit of having Optane SSDs in this mix. They took the bulk of the writes, freeing the QLC SSDs for responsive reads as a high-speed capacity tier. Even though we doubled the workload to eight SQL Server virtual machines hitting this HCI cluster, latency only increased slightly, showing that this configuration is well suited for workloads that can burst from time to time.
While performance in a fully operational environment is important, another consideration is how workloads will perform if a node in the cluster goes down or if workloads need to be migrated for system maintenance. To test this scenario, we kept our full load of 8 VMs and migrated them to a single node. In this configuration, we measured an average latency of only 4.5 ms, which was better than both nodes online. This comes in part from the removal of storage overhead in single-node operation.
Conclusion
For this project, we ran a series of SQL tests on the system to illustrate the performance workloads commonly found in Edge and SMB use cases. Our goal was to understand how effectively Microsoft Azure Stack HCI in this DataON cluster was able to leverage the hardware to achieve the desired results. Specifically, this means providing a solution that offers a rare combination of performance and value.
We can confirm through our tests that DataON's component selection did indeed succeed in creating a cost-effective Azure Stack HCI SDS solution that performs extremely well. This is due in part to their choice to use the Intel D5-P4326 SSD for capacity storage, which effectively leverages Intel Optane SSDs for tiering.
This is an essential notion, as QLC SSDs provide massive, dense capacity to the cluster, while offering the TCO benefits that accompany flash storage. To drive the point home, QLC drives allow for 15.36 TB of capacity per 2.5-inch drive bay. It would take 8 2 TB hard drives in RAID 0 to match the capacity, or moving to a 3.5-inch chassis to take advantage of larger but even slower hard drives. In any case, the performance drop from Intel QLC drives to hard drives is more than considerable; it is an exponential difference when it comes to application responsiveness.
Even though we would like all reads and writes to come from Optane SSDs (as they are the highest-performing media in this configuration), there will sometimes be a miss. In that case, QLC SSD performance will crush hard drives, protecting the HCI cluster from the performance irregularities common in topologies that combine flash drives and hard drives. In fact, we saw such balanced performance here that in the future, businesses in general may need to rethink the hard drive/flash design and lean more toward QLC/Optane design to get the most out of HCI.
The other major concern with 2-node clusters is performance in a degraded state. We tested this by failing one node and giving the entire SQL workload to a single node. In this case, SQL was more responsive and ran somewhat better than with 2 nodes, mainly due to reduced node-to-node communication overhead. Of course, it is not recommended to run for long in a degraded state like this, but it is reassuring to know that it can be done without sacrificing performance.
Overall, the HCI-224 HCI cluster with D5-P4326 QLC SSDs was simple to deploy, easy to use, and powerful enough for a wide range of workloads. Its price also makes it accessible to a wide range of users. Additionally, this system was certified for Microsoft Windows Server 2019 and validated as an Intel Select solution.
This report is sponsored by DataON. All views and opinions expressed in this report are based on our impartial assessment of the product(s) under study.
