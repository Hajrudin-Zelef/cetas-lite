---
id: collect-240926-storagereview/storagereview/fr-review-synology-e10m20-t1-review-a0ab7b39-2
title: "fr-review-synology-e10m20-t1-review-a0ab7b39"
domain: storagereview
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "benchmarks", "ethernet", "latency", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-synology-e10m20-t1-review-a0ab7b39.md
source_anchor: ""
source_lines: [3, 57]
sha256: 58c63ef5532d7e0f38aa2e672c9fd6bfcbeca276be0df4790e688014da1e8c12
---

# fr-review-synology-e10m20-t1-review-a0ab7b39

Synology offers a nice range of NAS devices that are perfect for SMBs. The only downside is that mid-range NAS are usually not designed for maximum performance and network speed. Thanks to a new expansion card, they no longer need to be. With the introduction of the Synology E10M20-T1 AIC, users can get much higher I/O and bandwidth with a single card.
For bandwidth, the Synology E10M20-T1 comes with a 10GbE port, greatly increasing the I/O of a NAS that has integrated GbE. In addition to improved network performance, the card comes with two NVMe SSD slots for M.2 drives (2280 and 22110 form factors). This gives users high I/O performance and adds two NVMe slots for SSD cache without giving up drive bays. This way, users can load their NAS with high-capacity hard drives and still have SSD cache. This card is also particularly useful in some of Synology's older NAS systems, which do not have integrated SSD slots.
The Synology E10M20-T1 card is available today for $250. For those who wish, we also offer a video presentation of the card.
Synology E10M20-T1 Specifications
| General |  | 
| Host Bus Interface | PCIe 3.0 x8 | 
| Bracket Height | Low profile and full height | 
| Size (Height x Width x Depth) | 71.75 mm x 200.05 mm x 17.70 mm | 
| Operating Temperature | 0 ° C to 40 ° C (32 ° F to 104 ° F) | 
| Storage Temperature | -20 ° C to 60 ° C (-5 ° F to 140 ° F) | 
| Relative Humidity | 5% to 95% RH | 
| Warranty | 5 years | 
| Storage |  | 
| Storage Interface | PCIe NVMe | 
| Supported Form Factor | 22110/2280 | 
| Connector Type and Quantity | M key, 2 slots | 
| Network |  | 
| IEEE Specification Compliance | IEEE 802.3an 10 Gb/s Ethernet IEEE 802.3bz 2.5 Gb/s / 5 Gb/s Gigabit Ethernet IEEE 802.3ab Fast Ethernet IEEE 802.3u IEEE 802.3x Flow Control | 
| Data Transfer Rate | 10 Gbps | 
| Network Operating Mode | Full Duplex | 
| Supported Features | 9 KB Jumbo Frame TCP/UDP/IP Checksum Offload Auto-negotiation between 100 Mb/s, 1 Gb/s, 2.5 Gb/s, 5 Gb/s, and 10 Gb/s | 
| Compatibility |  | 
| Applied Models NVMe SSD | SA Series: SA3600, SA3400 20 Series: RS820RP +, RS820 + 19 Series: DS2419+, DS1819+ 18 Series: RS2818RP+, DS3018xs, DS1618+ | 
Design and Build
The Synology E10M20-T1 is an HHFL AIC that will fit certain Synology NAS models. On one side are heat sinks that extend the full length of the card.
Removing the heat sink with four screws at the rear gives access to the two M.2 NVMe SSD bays. Overall, it is easy to secure the drives on the card.
The back of the card is relatively spartan. Synology also includes a full-size bracket in the box if needed, and thermal contact pads for the SSDs.
Performance
To test the Synology E10M20-T1, we installed it in a Synology DS1819+. We installed 14TB WD Red hard drives in the bays. For cache, we used Synology SNV3400-400G SSDs. We tested the drives in iSCSI and CIFS configurations in RAID 6, with and without cache.
Synthetic Enterprise Workload Analysis
Our enterprise hard drive benchmark process preconditions each drive to a steady state with the same workload with which the device will be tested under a heavy load of 16 threads with an outstanding queue of 16 per thread. It is then tested at defined intervals across multiple thread/queue depth profiles to show performance under light and heavy use. Since hard drives reach their rated performance level very quickly, we only graph the main sections of each test.
Preconditioning and Primary Steady-State Tests:
- Throughput (aggregate read + write IOPS)
- Average Latency (read + write latency averaged together)
- Maximum Latency (maximum read or write latency)
- Latency Standard Deviation (read + write standard deviation averaged together)
Our synthetic enterprise workload analysis includes four profiles based on real-world tasks. These profiles were developed to facilitate comparison with our past benchmarks as well as with widely published values such as maximum 4K and 8K 70/30 read and write speed, which is commonly used for enterprise drives.
- 4K
  - 100% read or 100% write
  - 100% 4K
- 8K70/30
  - 70% read, 30% write
  - 100% 8K
- 128K (sequential)
  - 100% read or 100% write
  - 100% 128K
In the first of our enterprise workloads, we measured a long sample of random 4K performance with 100% write and 100% read activity to obtain our main results. For CIFS, we saw 170 read IPS and 1,461 4,075 write IOPS without cache and 10,950 2,897 read IOPS and 1,502 20,021 write IOPS with cache active. For iSCSI, we saw 22,439 XNUMX IOPS read and XNUMX XNUMX IOPS write without cache and leveraging cache in the AIC, we saw XNUMX XNUMX IOPS read and XNUMX XNUMX IOPS write.
With 4K average latency, CIFS gave us 1,497 176 ms read and 63 ms write without cache, then with it enabled, it dropped to 23 ms read and 88 ms write. iSCSI saw 170 ms read and 12.8 ms write, then with cache enabled, it dropped to 11.4 ms read and XNUMX ms write.
Next, 4K maximum latency. Here, CIFS reached 4,476 3,360 ms read and 339 45 ms write without cache; leveraging the AIC, the figures dropped to 1,051 ms read and 6,131 ms write. iSCSI had 11,951 171 ms read and XNUMX XNUMX ms write without cache, and with it the read latency rose to XNUMX XNUMX ms but write latency dropped to XNUMX ms.
Our last 4K test is standard deviation. Here, without cache, CIFS gave us 228 ms read and 288 ms write, with cache-enabled latency reduced to 7.3 ms read and 2 ms write. For iSCSI, we again saw a spike instead of a drop in reads, going from 69 ms without cache to 196 ms with cache. Writes showed improvement, going from 282 ms to 16 ms.
Our next benchmark measures 100% 8K sequential throughput with a 16T16Q load in 100% read and 100% write operations. Here, the CIFS configuration without cache had 13,989 10,770 IOPS read and 13,055 11,443 IOPS write; after enabling cache, the figures rose to 56,579 30,288 IOPS read and 57,774 33,265 IOPS write. With iSCSI, we saw XNUMX XNUMX IOPS read and XNUMX XNUMX IOPS without cache enabled; with it enabled, we saw performance increase slightly to XNUMX XNUMX IOPS read and XNUMX XNUMX IOPS write.
Compared to the fixed workload at 16 threads and 16 maximum queues that we performed during the 100% 4K write test, our mixed workload profiles adapt performance across a wide range of thread/queue combinations. In these tests, we cover workload intensity from 2 threads and 2 queues up to 16 threads and 16 queues. In uncached CIFS, we saw throughput start at 221 IOPS and end at 219 IOPS, fairly stable throughout. With caching enabled, we saw CIFS start at 4,597 4,844 IOPS and end at 519 1,751 IOPS. For iSCSI, non-cache started at 8,308 IOPS and ended at 1,340 XNUMX IOPS. With cache enabled, we saw iSCSI start at XNUMX XNUMX IOPS and end at XNUMX XNUMX IOPS.
Looking at average 8K 70/30 response times, the CIFS configuration started at 18 ms and ended at 1,161 860 ms without cache, and with the card active, it dropped to 53 µs at the start and ended at 7.7 ms. For iSCSI, we saw 146 ms at the start and 470 ms at the end without the card; with the card it was 191 µs at the start and XNUMX ms at the end.
With 8K 70/30 maximum latency, the CIFS configuration started around 1,009 4,799 ms and rose to 523 260 ms. With cache enabled, the figures went from 1,436 ms to 5,614 ms. With iSCSI, we saw latency go from 640 13,588 ms to XNUMX XNUMX ms without cache and from XNUMX ms to XNUMX XNUMX ms with cache.
For 8K 70/30 standard deviation, the CIFS configuration started at 26 ms and ran at 477 ms without cache; with the network card, it went from 1.3 ms to 10.1 ms. For iSCSI, we saw 17 ms to 299 ms without cache, and 920 µs to 1,155 XNUMX ms with it.
