---
id: collect-240926-storagereview/storagereview/fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903-4
title: "fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["acquisition", "compute", "memory", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903.md
source_anchor: ""
source_lines: [93, 99]
sha256: 5c80c1fb56f3c7d7f19eea599c395af19fb12731ccc13e9d8eb84d10d0d03a6b
---

# fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903

Beyond raw performance, the R5715 with HDD storage is the ideal platform for a virtualized backup workload. To confirm this, we deployed Proxmox Backup Server on the R5715 configured with the 8-core EPYC 9015 processor and the same 12x 3.5-inch hard drive array. Proxmox is a representative example of the open-source hypervisor and infrastructure ecosystem that has seen considerable growth among SMBs, and Proxmox Backup Server, in particular, is perfectly suited to this server's storage profile.
We deployed Proxmox Backup Server 4.2.0 and used it to back up the virtual machines that power our Proxmox community Discord server environment.
In our configuration, backup and restore operations were somewhat limited by the system's 1 GbE network connection, which was the main bottleneck during large transfers. However, the platform supports simple network upgrades via OCP expansion cards, facilitating migration to 10 GbE or even 25 GbE connectivity. With a faster network, the R5715 would be capable of handling significantly higher backup throughput and restore performance, particularly in environments with larger virtual machine datasets or more demanding backup windows.
Conclusion
The Dell PowerEdge R4715 and R5715 servers owe their success to a perfectly adapted configuration. Two chassis with clearly differentiated form factors, four processor options covering all SMB needs without overlap, and a range of storage solutions broad enough to meet every need, from economical mass storage to 100% flash performance. This configuration flexibility is not purely theoretical. In our tests, the optimal configuration proved different depending on the workload. The 24-core 9255 processor offered the best compromise for transactional database performance on flash memory. The 8-core 9015 processor provided the compute power needed to deploy a Proxmox backup server with high-capacity hard drives. The R4715 with flash storage was the ideal choice for Windows shared storage, while the R5715 with hard drives was the optimal choice for capacity-focused workloads, where I/O spikes are not a limiting factor.
For SMBs and their partners, the value of these platforms lies in the ability to adapt infrastructure to the workload, rather than investing in oversized capacity. The Dell ecosystem, including iDRAC10 management, ProSupport, the security suite, and supply chain predictability, reinforces this value at the operational level. Agile IT teams and partners both benefit from a thoroughly mastered platform.
Dell's positioning regarding these servers, presented as a solution to consolidate existing infrastructures and reduce per-socket and per-core licensing costs, is confirmed by the data collected. With four processor options ranging from 8 to 32 cores on a single platform, clients and partners benefit from great flexibility to adapt acquisition, licensing, and operating costs to their actual needs. That is the entire added value, and the R4715 and R5715 servers fully deliver on it.
