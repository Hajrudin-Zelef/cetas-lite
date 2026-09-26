---
id: collect-240926-storagereview/storagereview/fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753-2
title: "fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "benchmark", "gpu", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753.md
source_anchor: ""
source_lines: [3, 93]
sha256: 088812728acfcf1a6843c6fd2d36bec619d67e07b37518abc696ea94e66a2265
---

# fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753

We have already tested the Dell PowerEdge R760 and found it, as always with Dell, extremely powerful. Following the recent launch of the 5th generation Intel Xeon Scalable processors, we were able to get our hands on a set of high-end Intel Xeon Platinum 8592+ processors. We thought: what better than to upgrade one of our PowerEdge R760 to combine the best of both worlds and see if Intel's latest innovation deserves to make others envious in the enterprise world?
Intel Xeon Platinum 8592+ Specifications
We started with the configuration provided by Dell in our R760, which consisted of two Intel 6430 processors and 1 TB of DDR5 4800 RAM, before swapping in the 8592+ processors. Although the Xeon 6340 is not the high end of the Sapphire Rapids family, we will still compare it to the Xeon 8592+. We align these two in this review not to present them as direct competitors, but to show the potential of existing platforms with the release of the new Emerald Rapids processors.
We brought in the figures for the Sapphire Rapids Platinum 8480+ as well as those for an HP ML350 for comparison with the Xeon 8592+ to get more precise figures. The 8480+ in the ML350 Ge11 is among the best processors in the Sapphire Rapids lineup. It should be noted that we are looking here only at the processors, without comparing the servers themselves to each other.
|  | Intel Xeon Platinum 8592+ Processor (R760) | Intel Xeon Gold 6430 Processor (R760) | Intel Xeon Platinum 8480+ Processor (ML350 G11) | 
|---|---|---|---|
| Generation | 5th, Emerald Rapids | 4th, Sapphire Rapids | 4th, Sapphire Rapids | 
| Cores | 64 | 32 | 56 | 
| Threads | 128 | 64 | 112 | 
| Base Frequency | 1.9 GHz | 2.10 GHz | 2.0 GHz | 
| Turbo Frequency | 3.9 GHz | 3.4 GHz | 3.8 GHz | 
| Cache | 320 MB | 60 MB | 105 MB | 
| Intel UPI Speed | 20 GT / s | 16 GT / s | 16 GT / s | 
| Maximum UPI Links | 4 | 3 | 4 | 
| TDP | 350 W | 270 W | 350 W | 
| Launch Date | Q4 '23 | Q1 '23 | Q1 '23 | 
| Maximum Memory Size | 4TB |  |  | 
| Memory Types | DDR5 5600 1 MT/s (XNUMX DPC) | Up to DDR5 4400 MT/s (1DPC and 2DPC) | Up to DDR5 4800 1 MT/s XNUMXDPC Up to DDR5 4400 2 MT/s XNUMXDPC | 
| Maximum Number of Memory Channels | 8 |  |  | 
| PCIe Revision | 5 |  |  | 
| Maximum Number of PCIe Lanes | 80 |  |  | 
| High Priority Cores | 20 @ 2.1 GHz | 12 @ 2.2 GHz | 16 @ 2.1 GHz | 
| Low Priority Cores | 44 @ 1.7 GHz | 20 @ 1.8 GHz | 40 @ 1.7 GHz | 
| Maximum Default EPC Size for Intel SGX | 512 GB | 128 GB | 512 GB | 
| Product Page | Link | Link | Link | 
Intel Xeon Platinum 8592+ Processor Performance
The high-level tested system specifications for each processor are as follows. Each system had two of the listed processors.
Xeon Platinum 8592+
- Dell Poweredge R760
- 1 TB DDR5 4800 XNUMX MHz
Xeon Gold 6430
- Dell Poweredge R760
- 1 TB DDR5 4800 XNUMX MHz
Xeon Platinum 8480+
- HP ML350 Generation 11
- 256GB DDR5 4800MHz
Blender OptiX
The first is Blender OptiX, an open-source 3D modeling application. This benchmark was run using the Blender Benchmark CLI utility. The score is expressed in samples per minute, with the higher being the better.
The Xeon 8592+ get off to a good start with more than double the performance of the Xeon 6430, reflecting the data shown on the spec sheets. The Xeon 8480+, however, comes within 81 points of the Xeon 8592+ in class.
| Blender 4.0 Processor | 2 Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 
|---|---|---|---|
| Monster | 1115.057 | 540.039 | 943.300 | 
| Junkshop | 780.408 | 361.066 | 627.662 | 
| Classroom | 556.550 | 278.228 | 475.144 | 
Cinebench R23
Maxon's Cinebench R23 is a CPU rendering benchmark that uses all CPU cores and threads. We ran it for multi-core and single-core tests. Higher scores are better. In this test, we see about double the results on the Emerald Rapids Xeon 8592+ compared to the Sapphire Rapids 6430. We found better single-core performance from the Xeon 8480+ than from the Xeon 8592+, but this is not really surprising since the Xeon 8480+ had a higher base clock speed. The Xeon 8592+ pulled ahead in the multi-core test.
| Cinebench R23 | 2 Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 
|---|---|---|---|
| Multi-Core CPU | 110,498 | 69,663 | 79,164 | 
| Single-Core CPU | 1,144 | 1,022 | 1,461 | 
| MP Ratio | 96.63x | 68.17x | 54.20x | 
Cinebench 2024
Here are the CPU results for the 2024 version of Cinebench.
We found a similar performance pattern among the three processors and Cinebench R23.
| Cinebench 2024 | 2 Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 
|---|---|---|---|
| Multi-Core CPU | 6,001 | 3,746 | 4,699 | 
| Single-Core CPU | 68 | 59 | 76 | 
| MP Ratio | 88.48x | 63.22x | 61.44x | 
Geekbench 6
Geekbench 6 is a cross-platform evaluation tool measuring the overall performance of a system. The higher the score, the better the performance. Geekbench offers a GPU performance test, but without a GPU, only the CPU results are available.
*We had issues with the Emerald Rapids Xeon 8592+ during this test and it did not complete. We will therefore revisit it when we can obtain performance figures. In the meantime, we only have numbers for the Xeon 6430 and Xeon 8480+*.
| Geekbench 6 | 2 Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 
|---|---|---|---|
| CPU Benchmark – Single Core | N/A | 1,488 | 1,939 | 
| CPU Benchmark – Multi-Core | N/A | 16,054 | 15,218 | 
Y-Cruncher
Y-cruncher is a popular benchmarking and stress testing application launched in 2009. This test is multithreaded and scalable, calculating Pi and other constants up to billions of digits. Faster is better in this test.
We see that the results follow the same pattern as the other tests, with the Xeon 6430 trailing and the Xeon 8480+ just behind the Xeon 8592+ numbers.
| Y-Cruncher (lower is better) | 2 Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 
|---|---|---|---|
| 1 billion | 4.239 seconds | 6.060 seconds | 5.136 seconds | 
| 2.5 billion | 11.466 seconds | 16.896 seconds | 13.768 seconds | 
| 5 billion | 25.325 seconds | 36.843 seconds | 29.889 seconds | 
| 10 billion | 54.921 seconds | 80.574 seconds | 65.194 seconds | 
| 25 billion | 156.923 seconds | 229.017 seconds | 186.841 seconds | 
7-Zip Compression
The popular 7-Zip utility has a built-in memory test that demonstrates CPU performance very well. In this test, we run it with a dictionary size of 128 MB when possible. As expected, we still see better results on the Xeon 8592+ processors.
|  | 2 Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 
|---|---|---|---|
| Compression |  |  |  | 
| Current CPU Usage | 5,609 % | 5,732 % | 5,482 % | 
| Current Rating/Usage | 4.912 GIPS | 3.912 GIPS | 4.628 GIPS | 
| Current | 275,503 GIPS | 224.209 GIPS | 253.724 GIPS | 
| Resulting CPU Usage | 5,605 % | 5,669 % | 5,475 % | 
| Resulting Rating/Usage | 4.883 GIPS | 3.923 GIPS | 4.628 GIPS | 
| Resulting Rating | 273.716 GIPS | 222.407 GIPS | 253.382 GIPS | 
| Decompression |  |  |  | 
| Current CPU Usage | 6,243 % | 5,852 % | 6,219 % | 
| Current Rating/Usage | 3.635 GIPS | 3.423 GIPS | 3.745 GIPS | 
| Current | 226.917 GIPS | 200.350 GIPS | 231.916 GIPS | 
