---
id: collect-240926-storagereview/storagereview/fr-review-synology-e10m20-t1-review-a0ab7b39-3
title: "fr-review-synology-e10m20-t1-review-a0ab7b39"
domain: storagereview
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "benchmarks", "latency", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-synology-e10m20-t1-review-a0ab7b39.md
source_anchor: ""
source_lines: [58, 63]
sha256: 9ff03d761ef54da2bdd95d1262393ee2cba290f368be7371ca8b71dd2a70c09d
---

# fr-review-synology-e10m20-t1-review-a0ab7b39

The last synthetic enterprise workload benchmark is our 128K test, which is a large-block sequential test that shows the highest sequential transfer speed for a device. In this workload scenario, CIFS had 1.09 GB/s read and 464 MB/s write without cache and 1.14 GB/s read and 484 MB/s with it. For iSCSI, we saw 1.15 GB/s read and 443 MB/s without cache and 1.15 GB/s read and 615 MB/s write with it.
Conclusion
Synology has developed a simple way to add 10GbE connectivity and NVMe cache to a selected number of their NAS devices via the Synology E10M20-T1 AIC. The card integrates perfectly with the company's SSD range and has slots to accommodate two 2280 or 22110 M.2 form factors. This allows users to load the NAS with high-capacity hard drives, then use the card to improve I/O performance. And, of course, the 10GbE port increases network speed compared to integrated GbE ports.
For performance testing, we again used the Synology DS1819+ with 14TB WD Red hard drives, to which we this time added the E10M20-T1 interface controller and two Synology SNV3400-400G SSDs. In most cases, but not all, we saw performance improvement. Rather than detailing performance again, let's first look at the differences with the CIFS configuration. In CIFS 4K throughput, we observed an increase of 3,905 IOPS read and 9,489 IOPS write. 4K average latency decreased by 1,434 ms read and 153 ms write. 4K maximum latency dropped by 4,137 ms read and 3,315 ms write. 4K standard deviation decreased by 220.7 ms read and 286 ms write. In 100% 8K mode, we saw a decrease of 934 IOPS read and an increase of 673 IOPS write. In large-block sequential mode, throughput increased by 50 MB/s read and 20 MB/s write.
With iSCSI performance, we saw highlights in 4K throughput with an increase of 17,124 20,937 IOPS read and 4 75.2 IOPS write. In 4K average latency, we saw a decrease of 158.6 ms read and 4 ms write. In 4K max latency, we saw an increase of 10,900 15,960 ms read and a decrease of 4 127 ms write. 266K standard deviation saw a latency spike increase again of 8 ms read and a decrease in writes of 1,200 ms. In 2,977% 1.15K, we saw an increase of 172 XNUMX IOPS read and XNUMX XNUMX IOPS write. Large-block sequential read holds at XNUMX GB/s with or without cache and increases writes by XNUMX MB/s.
The Synology E10M20-T1 AIC is an easy way to add 10GbE connectivity and SSD cache for certain Synology NAS models. Although it did not improve all performance in all areas, it saw significant improvements in several of our benchmarks.
