---
id: collect-240926-storagereview/storagereview/fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a-2
title: "fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "EU", "Intel"]
dates: []
keywords: ["amd", "cost", "ethernet", "intel", "latency", "licenses", "liquid cooling", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a.md
source_anchor: ""
source_lines: [3, 58]
sha256: 1d6b2d6240a7bb53f9c58bfcd40f8d0c10da7fa6fb543124d77b74078500e80f
---

# fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a

AMD's EPYC 4005 family brings Zen 5 to AM5 with a clear operational goal. Maintain platform consistency for system builders and hosters, then optimize the essentials for daily use. The result is a server-focused AM5 approach that prioritizes predictable deployments, stable firmware, and an OS matrix designed for scalability.
Within AMD's lineup, the EPYC 4005 sits between Ryzen on AM5 and the high-end EPYC platforms on SP5 and SP6. Ryzen can anchor light server roles, but it doesn't offer the validation depth, lifecycle, or feature set expected by data center operators. The big sockets offer more cores, memory channels, and lanes, but at a much higher cost. The EPYC 4005 bridges this gap. You keep the affordability and simplicity of AM5 while benefiting from an enterprise framework compatible with current licenses and power envelopes.
The product lineup reflects this intention. Coverage extends in both directions, allowing buyers to standardize with a modest power budget or opt for a significantly higher cache tier for latency-sensitive services, all without changing chassis or I/O footprint.
AMD EPYC 4004 vs 4005: Practical Improvements
The EPYC 4005 is a straightforward generational evolution from the 4004, retaining the AM5 heritage while improving the platform. With the 4005, you get Zen 5 cores, increased RAM performance up to DDR5-5600, and the same PCIe Gen28 capability up to 5 lanes. AMD also adds AVX-512 with a full 512-bit data path, giving this single-socket server category more headroom for vector-intensive tasks without changing motherboard or chassis.
At the top of the stack, the comparison is simple. The EPYC 4565P from the 4005 is a 16-core processor with 64 MB of L3 cache, a 170 W TDP, and a clock speed of up to 5.7 GHz. Its 4004 counterpart, the EPYC 4564P, is also a 16-core processor with 64 MB of L3 cache and a 170 W TDP, with the same maximum boost of 5.7 GHz. This means generational gains come from core and platform updates rather than a significant jump in specifications.
If you want to work with cache, both families offer a 3D V-Cache option. However, the 4005 places the EPYC 4585PX processor at the top with 170 W and 128 MB of L3 cache, while the 4004 caps out with the EPYC 4584PX at 120 W and 128 MB of cache. The base frequencies of the 4585PX are also slightly higher.
The most significant practical addition of the 4005 is a 16-core, 65 W reference. The EPYC 4545P allows bare metal hosted servers and SMBs to standardize 16 cores for Windows licensing while adhering to stricter power and cooling envelopes. This option did not exist in the 4004, where 16 cores meant 170 W on the standard chip or a 120 W V-Cache component with lower base frequencies.
As previously noted, memory support has been slightly improved. The 4004 validated two channels of DDR5 up to 5200 4005 MT/s. The 5600 brings this speed to 192 128 MT/s and increases the maximum supported capacity at the platform level from XNUMX GB to XNUMX GB, which facilitates mixed virtualized workloads and small database instances.
EPYC 4004 and 4005 Processor Comparison
| AMD EPYC 4004 Series Models | Arch. | Cores / Threads | L3 Cache (MB) | TDP (W) | Base Clock (GHz) | Boost Clock (GHz) | Price (1KU, USD) | 
|---|---|---|---|---|---|---|---|
| EPYC 4124P | Zen 4 | 4/8 | 16 | 65W | 3.8 | 5.1 | $149 | 
| EPYC 4244P | Zen 4 | 6/12 | 32 | 65W | 3.8 | 5.1 | $229 | 
| EPYC 4344P | Zen 4 | 8/16 | 32 | 65W | 3.8 | 5.3 | $329 | 
| EPYC 4364P | Zen 4 | 8/16 | 32 | 105W | 4.5 | 5.4 | $399 | 
| EPYC 4464P | Zen 4 | 12/24 | 64 | 65W | 3.7 | 5.4 | $429 | 
| EPYC4484PX | Zen 4 | 12/24 | 128 | 120W | 4.4 | 5.6 | $599 | 
| EPYC 4564P | Zen 4 | 16/32 | 64 | 170W | 4.5 | 5.7 | $699 | 
| EPYC4584PX | Zen 4 | 16/32 | 128 | 120W | 4.2 | 5.7 | $699 | 
| AMD EPYC 4005 Series Models |  |  |  |  |  |  |  | 
| EPYC 4245P | Zen 5 | 6/12 | 32 | 65W | 3.9 | 5.4 | $239 | 
| EPYC 4345P | Zen 5 | 8/16 | 32 | 65W | 3.8 | 5.5 | $329 | 
| EPYC 4465P | Zen 5 | 12/24 | 64 | 65W | 3.4 | 5.4 | $399 | 
| EPYC 4545P | Zen 5 | 16/32 | 64 | 65W | 3.0 | 5.4 | $549 | 
| EPYC 4565P | Zen 5 | 16/32 | 64 | 170W | 4.3 | 5.7 | $589 | 
| EPYC4585PX | Zen 5 | 16/32 | 128 | 170W | 4.3 | 5.7 | $699 | 
Test Platform: MSI S1102-02 with Liquid Cooling
For this test, we used the MSI S1102-02 motherboard, a 1U AM5 barebone perfectly suited to the target market of EPYC 4005 processors. It supports AM5 processors up to 170 W and our model comes with an optional all-in-one liquid cooling system, a rare feature in this category and ideal for high-end 4005 processors. This system should notably reduce fan speed and noise levels. For cooling, the chassis includes seven 40 mm system fans, a 600 W 1+1 Platinum power supply, and full management via an AST2600 BMC controller compatible with IPMI and Redfish.
Networking is handled via two RJ10 45 GbE ports on the Intel X710 processor. Storage is simple with four hot-swappable SATA bays at the front, two internal 2.5-inch SATA bays, and two NVMe M.2 2280/22110 slots. Expansion includes one Gen5 x16 slot for FHHL cards and a shared Gen4 x4 path that can be routed to a secondary PCIe slot or to one of the M.2 sockets. Memory consists of four DDR5 UDIMM slots with ECC support, validated up to 5600 1 MT/s at 48 DPC and 192 GB per DIMM for a maximum capacity of 26 GB. Physical depth is XNUMX inches, offering great deployment flexibility in short racks.
Full MSI S1102-02 Specifications
| Specifications | DETAILS | 
| Form Factor | 1U | 
| Dimensions | 438.5 mm (17.26″) L x 43.5 mm (1.71″) H x 660 mm (26″) D | 
| Processor | AMD Ryzen™ 7000/9000 Series and EPYC™ 4004/4005 Series Processor, up to 170 W TDP | 
| Socket | (1) AMD AM5 Socket | 
| Chipset | AMD B650 | 
| Memory | (4) DDR5 DIMM slots, 2DPC, ECC/non-ECC UDIMM – Max. Speed 5600MT/s(1DPC) and 3600MT/s(2DPC) – Max. Capacity per DIMM: 48 GB | 
| Drive Bays | (4) Hot-swappable 3.5"/2.5" drive bays support SATA 3.0 | 
| Internal Storage | (2) M.2 2280/22110 PCIe4.0 x4 ports (2) Internal 2.5″ drive bays support SATA 2.0 | 
| Expansion Slots | (1) Processor PCIe 4.0 x16 slot supports FHHL PCIe card | 
| Networking | (2) 10GBase-T Ethernet ports (Intel® X710AT2) *JLAN1,2 supports NCSI | 
| RAID | N/A | 
| Front I/O | (4) Hot-swappable 3.5" drive bays (2) USB3.2 Gen1 Type-A ports (1) System power button (1) UID button (1) UID LED (5) Status LEDs: Power/Fault/HDD/(2)LAN | 
| Rear I/O | (2) 10GBase-T Ethernet ports (1) Dedicated 1000Base-T server management port (4) USB3.2 Gen1 Type-A ports (1) VGA D-Sub port (1) COM DB9 port (1) UID LED button | 
| TPM | (1) TPM header with SPI interface | 
| Security | TPM 2.0 | 
| Server Management | (1) Dedicated 1000Base-T server management port (Realtek RTL8211FD-CG) ASPEED AST2600 with AMI MegaRAC-based firmware supporting IPMI 2.0 and DMTF Redfish API | 
| Environment | System operating temperature: 0°C ~ 35°C Non-operating temperature: -20°C ~ 70°C Non-operating relative humidity: 5% ~ 85% (non-condensing) | 
| Power Supply | (1+1) 450 W 80+ Platinum redundant power supply | 
| Cooling | (1) Closed-loop liquid cooling module for processor max. 170 W (7) 4028 system fans | 
| Certifications | CE, FCC (Class A) | 
| Package Contents | (1) S1102-02 1U Barebone (1) Quick Guide (1) Rail Kit (1) Closed-loop liquid cooling module (assembled inside chassis) (4) Power Cables | 
| Parts List | (1) Quick Guide: G52-S3362X2-Q13 (1) Rail Kit: E21-S336020-C27 (1) Closed-loop liquid cooling module: E34-F000091-AQ0 (4) Power Cables: (2)K33-3001002-I45 for North America, (2)K33-3001005-I45 for EU | 
| Optional Parts | TPM2.0 Module: TPM20-IR C13 Power Cord 250 V/10 A (CN): K33-3001029-I45 C13 Power Cord 125 V/13 A (JP): K33-3001346-I45 | 
AMD EPYC 4005 Performance
