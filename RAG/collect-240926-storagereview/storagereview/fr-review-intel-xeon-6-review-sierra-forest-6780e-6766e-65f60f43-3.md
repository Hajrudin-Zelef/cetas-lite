---
id: collect-240926-storagereview/storagereview/fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43-3
title: "fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "benchmark", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43.md
source_anchor: ""
source_lines: [43, 114]
sha256: acb19c246e6f14cf1284c86439bdafd31293fb1c42c116db8a2cb375315d3613
---

# fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43

The processors are early samples and not retail boxed processors. As such, we were unable to run our full test suite and did not have time before the embargo lifted to properly resolve some of the performance and stability issues we encountered. As such, the following data should be considered informative and not definitive. We will wait for platforms to ship from Intel's OEM partners before making a more conclusive judgment on the capabilities of the Sierra Forest processors.
Intel shipped us a Quanta QuantaGrid D55Q-2U system as a test platform to showcase the new processors. This server features a 24-bay direct-connect NVMe backplane supporting U.2 Gen4/Gen5 SSDs.
Our test configuration included 16 x 16 GB DDR5-6400 RAM and Micron MTC10F1084S1RC64BDY modules.
Xeon 6780E and Xeon 6766E
- Quanta QuantaGrid D55Q-2U
- 256GB DDR5 6400MHz
Xeon Platinum 8592+
- Dell Poweredge R760
- 1 TB DDR5 4800 XNUMX MHz
Xeon Gold 6430
- Dell Poweredge R760
- 1 TB DDR5 4800 XNUMX MHz
Xeon Platinum 8480+
- HP ML350 Gen 11
- 256GB DDR5 4800MHz
With the earliest test platform gremlins in play, the new Xeon 6780E and 6766E processors were tested in a Windows Server 2025 environment, while the older processors were tested in Windows Server 2022.
Blender OptiX 4.0
In Blender OptiX, we have 3 different tests: Monster, Junkshop, and Classroom. In Monster, we saw the 6780E and 6766E beat the 8592+ by 21% and 14%, with an 18% difference between the two. In Junkshop, the 6780E was 8592% ahead of the 10.5+, with the 6766E behind the 8592+ by only 0.35%. There was a 10.8% difference between the 6780E and 6766E in Junkshop. For the Classroom part, the 6780E and 6766E beat the 8592+ by 20% and 22%, with a 10% difference between the 6780E and 6766E.
| Blender 4.0 Processor | 2x Xeon 6780E (256 GB DDR5) | 2x Xeon 6766E (256 GB DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 
|---|---|---|---|---|---|
| Monster | 1410.463 | 1297.715 | 1115.057 | 943.300 | 540.039 | 
| Junkshop | 862.418 | 777.716 | 780.408 | 627.662 | 361.066 | 
| Classroom | 696.543 | 628.960 | 556.550 | 475.144 | 278.228 | 
Cinebench R23
For Cinebench R23, multi-core performance showed the 6780E and 6766E behind the 8592+ by 38% and 42%, respectively. For single-core performance, the 6780E was 24% behind the 8592+ and the 6766E was 31% behind. The difference between the 6780E and 6766E dropped to about 5% for the multi-core score and 18% for single-core.
| Cinebench R23 | 2x Xeon 6780E (256 GB DDR5) | 2x Xeon 6766E (256 GB DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 
|---|---|---|---|---|---|
| Multi-core CPU | 67,984 | 64,326 | 110,498 | 79,164 | 69,663 | 
| Single-core CPU | 873 | 793 | 1,144 | 1,461 | 1,022 | 
| PM Ratio | 77.91x | 81.10x | 96.63x | 54.20x | 68.17x | 
Cinebench 2024
For Cinebench 2024, the multi-core score saw a drop of about 55% from the 8592+ to the 6780E and a drop of 61% from the 8592+ to the 6766E. This also showed a 13% difference between the 6780E and 6766E on Multi-Core. For single-core scores, the 6780E and 6766E had only a 5% difference, with the 6780E and 6766E behind the 8592+ by 37% and 34%, respectively.
| Cinebench R23 | 2x Xeon 6780E (256 GB DDR5) | 2x Xeon 6766E (256 GB DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 
|---|---|---|---|---|---|
| Multi-core CPU | 2,687 | 2,347 | 6,001 | 3,746 | 4,699 | 
| Single-core CPU | 43 | 45 | 68 | 59 | 76 | 
| PM Ratio | 62.85x | 52.65x | 88.48x | 63.22x | 61.44x | 
Y-Cruncher
Y-cruncher is a popular benchmarking and stress-testing application launched in 2009. This test is multithreaded and scalable, calculating Pi and other constants to billions of digits. Faster is better in this test.
The new 6780E and 6766E run a bit slower than the Emerald Rapids 8592+, but they aren't exactly direct competitors. The 6780E ran about 13% faster than the Xeon Gold 6430 for the 1 billion test, but about 39% slower than the Xeon Platinum 8592+. The 6766E ran 19% slower than the Gold 6439 and 42% slower than the Platinum 8592+.
| Y-Cruncher (lower is better) | Xeon 6780E (256 GB DDR5) | Xeon 6766E (256 GB DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 
|---|---|---|---|---|---|
| 1 billion | 6.927 seconds | 7.254 seconds | 4.239 seconds | 6.060 seconds | 5.136 seconds | 
| 2.5 billion | 17.898 seconds | 19.507 seconds | 11.466 seconds | 16.896 seconds | 13.768 seconds | 
| 5 billion | 38.454 seconds | 41.116 seconds | 25.325 seconds | 36.843 seconds | 29.889 seconds | 
| 10 billion | 81.146 seconds | 87.403 seconds | 54.921 seconds | 80.574 seconds | 65.194 seconds | 
| 25 billion | 217.530 seconds | 238.813 seconds | 156.923 seconds | 229.017 seconds | 186.841 seconds | 
| 50 billion | 565.913 seconds | 502.245 seconds | N/A | N/A | N/A | 
7-Zip
The popular 7-Zip utility has a built-in memory benchmark that demonstrates processor performance very well. In this test, we run it with a 128 MB dictionary size when possible. In this test, the 6780E barely beats the Platinum 8592+ in the overall score. For decompression, the 6766E also beat the 8592+.
| 7-Zip Compression | Xeon 6780E (256 GB DDR5) | Xeon 6766E (256 GB DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 TB DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 GB DDR5 4400 XNUMX MHz) | 
|---|---|---|---|---|---|
| Compression |  |  |  |  |  | 
| Current CPU Usage | 5,891 % | 4,768 % | 5,609 % | 5,732 % | 5,482 % | 
| Current Rating/Usage | 4.985 GIPS | 4.614 GIPS | 4.912 GIPS | 3.912 GIPS | 4.628 GIPS | 
| Current | 293.689 GIPS | 220.001 GIPS | 275,503 GIPS | 224.209 GIPS | 253.724 GIPS | 
| Resulting CPU Usage | 5,603 % | 5,103 % | 5,605 % | 5,669 % | 5,475 % | 
| Resulting Rating/Usage | 4.954 GIPS | 4.638 GIPS | 4.883 GIPS | 3.923 GIPS | 4.628 GIPS | 
| Resulting Rating | 277.670 GIPS | 236.910 GIPS | 273.716 GIPS | 222.407 GIPS | 253.382 GIPS | 
| Decompression |  |  |  |  |  | 
| Current CPU Usage | 5,962 GIPS | 5,798 % | 6,243 % | 5,852 % | 6,219 % | 
| Current Rating/Usage | 4.550 GIPS | 4.152 GIPS | 3.635 GIPS | 3.423 GIPS | 3.745 GIPS | 
| Current | 271.266 GIPS | 240.693 GIPS | 226.917 GIPS | 200.350 GIPS | 231.916 GIPS | 
| Resulting CPU Usage | 5,832 % | 6,029 % | 6,232 % | 5,894 % | 6,129 % | 
| Resulting Rating/Usage | 4.540 GIPS | 4.161 GIPS | 3.654 GIPS | 3.385 GIPS | 3.871 GIPS | 
| Resulting Rating | 264.764 GIPS | 250.853 GIPS | 227.744 GIPS | 199.363 GIPS | 237.259 GIPS | 
| Total Rating |  |  |  |  |  | 
| Total CPU Usage | 5,717 % | 5,566 % | 5,919 % | 5,781 % | 5,802 % | 
| Total Rating/Usage | 4.747 GIPS | 4.399 GIPS | 4.269 GIPS | 3.654 GIPS | 4.249 GIPS | 
| Total Rating | 271.217 GIPS | 243.882 GIPS | 250.730 GIPS | 210.363 GIPS | 245.320 GIPS | 
Conclusion
The new Intel Xeon 6 E-core lineup, part of the Sierra Forest family, represents a change of pace in revision cadence. In previous Intel releases, we typically saw the high-end processors first; this time, Intel is leading with efficient mid-range SKUs.
