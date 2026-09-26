---
id: collect-240926-storagereview/storagereview/fr-review-maximizing-vsan-8-esa-potential-with-western-digital-ultrastar-dc-sn65-352ebbbc-3
title: "fr-review-maximizing-vsan-8-esa-potential-with-western-digital-ultrastar-dc-sn65-352ebbbc"
domain: storagereview
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "latency", "parameters", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-maximizing-vsan-8-esa-potential-with-western-digital-ultrastar-dc-sn65-352ebbbc.md
source_anchor: ""
source_lines: [27, 51]
sha256: 97811fdba2d6e29ebf015f418d827d5eb9cad3dc099a6889810b300508b27bd5
---

# fr-review-maximizing-vsan-8-esa-potential-with-western-digital-ultrastar-dc-sn65-352ebbbc

Another variable measured in this report concerns the impact of RAID 1 and RAID 5 on overall storage performance within our vSAN cluster. Historically, RAID 1 has been a common storage strategy for vSAN deployments. The recommended configuration for ESA is "vSAN ESA Default Policy – RAID 5," which offers nearly identical performance with much higher capacity. RAID 1 incurs an immediate 50% loss, while RAID 5 limits parity overhead, providing much greater usable capacity.
To measure the performance of our VMware vSAN cluster running ESXi 8 Update 2, we used HCIbench. HCIbench was developed to simplify testing of HCI clusters, in which you need to deploy worker VMs, generate virtual disks, and orchestrate tests across a cluster while aggregating all these results. In our tests, we used the following parameters for our four ESA configurations:
HCIBench VM Deployment:
- 16 VMs
- 8 data disks of 50GB per VM
- 16 vCPUs per VM
- 8GB of RAM per VM
- Test duration of 3600 seconds per workload
- 8 threads per disk for sequential workloads and 16 threads for random workloads
| Western Digital Ultrastar DC SN655 vSAN ESA DATA | 1024K Sequential Write MB/s | 1024K Sequential Read MB/s | 4K Random Write IOPS | 4K Random Read IOPS | 8K Random 70/30 IOPS | 
|---|---|---|---|---|---|
| vSAN ESA 4-SSD per Node RAID1 | 10,914 | 17,638 | 378,861 | 563,054 | 318,640 | 
| vSAN ESA 4-SSD per Node RAID5 | 10,999 | 17,726 | 476,770 | 524,076 | 301,960 | 
| vSAN ESA 8-SSD per Node RAID1 | 12,702 | 24,128 | 520,632 | 526,504 | 323,674 | 
| vSAN ESA 8-SSD per Node RAID5 | 14,336 | 21,994 | 504,508 | 523,666 | 292,557 | 
We split the performance data into two overlapping sections: four versus eight SSDs per node and RAID1 versus RAID5. In our initial comparison, we focused on the point at which performance begins to saturate as more drives are added to the cluster. With four hosts, we looked at 16 SSDs in the cluster, versus 32 when we scaled each node to eight SSDs. The second aspect is the significant difference between the RAID1 ESA policy and the recommended RAID5 ESA policy. The benefits of using RAID5 are quite obvious: users benefit from a huge amount of available capacity through the data protection policy, paired with Western Digital's dense SSDs.
We noticed a few trends emerging when examining the data collected during our tests. If you focus on total bandwidth, the VMware vSAN ESA cluster experienced a significant performance improvement going from four to eight SSDs per node. We measured about 10.9GB/s in sequential write from the RAID1 and RAID5 configurations with a 4-SSD configuration. In sequential read, we measured 17.6-17.7GB/s in RAID1 and RAID5 modes. Moving to random workloads, we found that RAID5 had an advantage in the 4-SSD configuration, measuring 476K IOPS in 4K random write, versus 379K IOPS in RAID1. In 4K random read, the two configurations began to converge, with RAID1 measuring 563K IOPS to 524K IOPS in RAID5. In 8K 70/30, the difference was much narrower, with RAID1 having an edge measuring 319K IOPS versus 302K IOPS for RAID5.
Moving to eight SSDs per node, we found improved performance, but it did not double performance compared to four SSDs. This is where the upper performance limits of vSAN and the 100GbE bandwidth limit on our nodes began to appear. Yes, faster network connections are available, but that moves further away from HCI deployments and typical of SMBs/SMEs.
Regarding write bandwidth, RAID5 had the upper hand, measuring 14.3GB/s versus 12.7GB/s for RAID1. RAID1 led in terms of read bandwidth with 24.1GB/s versus RAID5 with 22GB/s. With 4K random write transfers, RAID1 and RAID5 were close, although we measured 521K IOPS in R1 versus 505K IOPS in R5. When examining 4K random read performance, we found ourselves facing a vSAN cluster performance limit where four and eight SSD configurations were very close to each other. Here, with our eight-SSD configuration, we measured 527K IOPS in RAID1 and 524K IOPS in RAID5. In 8K 70/30, RAID1 ultimately had the advantage, with 324K IOPS versus 293K IOPS in RAID5.
It is important to note that with VMware vSAN ESA, the RAID1 or RAID5 configuration is not a decision set in stone. It is a storage policy applied at the level of a VM's vDisk, meaning that both can coexist simultaneously. So suppose you have a database VM where every bit of I/O counts; give it a RAID1 policy, where everything else is on RAID5. This way, you optimize your use of space as best as possible.
Final Thoughts
Western Digital's Ultrastar DC SN655 Enterprise PCIe Gen 4.0 dual-port NVMe SSD represents the top tier in dense enterprise-class storage while retaining the affordable price for which Western Digital is known. It is available in the U.2 form factor but is also U.3 backward compatible. This means data administrators and engineers can increase storage density with a familiar form factor in dense 1U servers.
Increased capacity density also means a more significant improvement in storage resource utilization for use cases, such as increasing the number of virtualized hosts per SSD and consolidating larger application datasets onto fewer drives. These Big Data analytics and AI/ML datasets also benefit from moving to higher capacities by unlocking low latency and higher throughput from the SN655 NVMe SSD, resulting in faster time to insights and real-time analytics.
Our performance results show that Western Digital Ultrastar DC SN655 NVMe SSDs are well-suited for VMware vSAN 8 ESA deployments, especially concerning SMBs, edge, retail, and thousands of other use cases. For businesses looking to start with a small deployment and scale as demand increases, we found that just four Western Digital SSDs per node delivered a large portion of the performance of our eight-SSD-per-node configuration, with plenty of room for growth as vSAN applications and data needs increase. This helps increase IT budgets since organizations do not need to overcommit on drives to achieve performance targets, and adding SSDs to vSAN could not be simpler.
With VMware vSAN 8 Update 2 GA, vSAN customers are likely ready to take the leap to the new storage architecture. ESA clearly has many benefits, but we are most pleased with the simplicity that vSAN currently offers; select server nodes and SSDs, and off you go with the same "point-and-click" ease of use that vSAN has always provided. When we examined SSDs for vSAN 8, we found that the Ultrastar DC SN655 NVMe SSD performs extremely well, delivering excellent performance in RAID5, which helps customers maximize their storage footprint in vSAN without sacrificing performance. Better yet, the drives are cost-effective, offering even more value to those looking to maximize their investment in VMware vSAN 8.
