---
id: collect-240926-storagereview/storagereview/fr-review-amd-epyc-9754s-review-a-cpu-with-a-very-particular-set-of-skills-6f36d356
title: "fr-review-amd-epyc-9754s-review-a-cpu-with-a-very-particular-set-of-skills-6f36d356"
domain: storagereview
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "compute", "exploit", "latency", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-amd-epyc-9754s-review-a-cpu-with-a-very-particular-set-of-skills-6f36d356.md
source_anchor: ""
source_lines: [1, 52]
sha256: 383a28d0557cdb0858324fe0a9a91cc5b4dde51cf434ac2b289771170d2559c4
---

# fr-review-amd-epyc-9754s-review-a-cpu-with-a-very-particular-set-of-skills-6f36d356

<!-- source: https://www.storagereview.com/fr/review/amd-epyc-9754s-review-a-cpu-with-a-very-particular-set-of-skills -->

Last year, AMD expanded its server processor lineup with the 4th Gen EPYC. While the 128 to 256-core, 9754-thread EPYC occupies the top spot, just below the SKU matrix is the AMD EPYC 9754S. The difference between the two chips is simple but dramatic. The 9754S has simultaneous multithreading (SMT) disabled. This means the 9754S offers the same 128 cores as the 9754, but with SMT disabled, only 128 threads, compared to 256. This change results in a nice reduction for customers already disabling SMT.
| Model | Colorful | Maximum number of threads | Default TDP | Freq. (GHz) | Boost Freq. (GHz) | L3 Cache (MB) | 
|---|---|---|---|---|---|---|
| 9754 | 128 | 256 | 360W | 2.25 | 3.10 | 256 | 
| 9754S | 128 | 128 | 360W | 2.25 | 3.10 | 256 | 
| 9734 | 112 | 224 | 320W | 2.2 | 3.0 | 256 | 
What is AMD SMT and why does the 9754S exist?
With SMT, a single EPYC processor core can process two threads simultaneously, which can lead to more efficient use of processor resources. When one thread is waiting for data to be loaded from memory or is idle, the other thread can execute instructions. This means the core spends less time idle, potentially improving performance. This is especially true in use cases such as virtualization and rendering.
Disabling SMT can allow manufacturers to market these chips as lower-tier products, ensuring they still meet specific performance and stability criteria. Processors with SMT disabled can be influenced by binning processes, market segmentation strategies, and the desire to meet specific performance or efficiency needs, demonstrating the nuanced approach manufacturers take to product planning and positioning.
That said, not all workloads benefit from SMT and it often happens that SMT is disabled in the BIOS on an AMD server. While this can be an effective adjustment, it raises another important point. The 9754S chip with SMT disabled is slightly cheaper than the 9754. In both cases, single-threaded applications, compute workloads, and all use cases where processor latency is of crucial importance can benefit from disabling SMT.
AMD EPYC 9754S and EPYC 9754 Performance
We wanted to run two of our regular tests, y-cruncher and Cinebench 2024, and see what performance differences we get with and without SMT. We compared the 9754S and 9754 while running the 9754 with SMT enabled and disabled to see what advantages the 9754S has without SMT.
Test platform and specifications:
- TYAN Transport HX TN85-B8261
- 512GB DDR5
- Windows Server 2022
Cinebench 2024
The first is Cinebench 2024, with SMT enabled on our non-S model. Here we can see that we are within run-to-run variation differences.
| Cinebench 2024 Processor | 2x EPYC9754S | 2x EPYC9754 | 
|---|---|---|
| Multi-core processor | 2,682 | 2,587 | 
| Single-core processor | 68 | 69 | 
| PM Ratio | 39.19x | 37.64x | 
y-cruncher was specifically selected due to the program's architecture, positioned as a comprehensive system test. By performing a Pi calculation as large as possible in system memory, we sought to prove our long-standing intuition that SMT can have a negative impact on processor and memory-related workloads. Let's first take a look at the results before diving into what it all means.
y-cruncher 0.8.3
| y-cruncher 0.8.3 Total computation time in seconds (lower is better) | 2x EPYC9754S | 2x EPYC 9754 (SMT disabled) | 2x EPYC 9754 (SMT enabled) | 9754 SMT disabled Performance increase | 
|---|---|---|---|---|
| 1 billion | 13.481 | 13.546 | 14.139 | 4.65% | 
| 2.5 billion | 23.818 | 24.144 | 28.111 | 15.27% | 
| 5 billion | 40.760 | 40.797 | 49.271 | 17.27% | 
| 10 billion | 77.409 | 77.959 | 95.420 | 18.88% | 
| 25 billion | 203.303 | 202.124 | 233.629 | 12.98% | 
| 50 billion | 475.557 | 476.949 | 520.349 | 8.61% | 
| 100 billion | 1,248.458 | 1,251.36 | 1,242.419 | -0.49% | 
y-cruncher 0.8.4
| y-cruncher 0.8.4 Total computation time in seconds (lower is better) | 2x EPYC9754S | 2x EPYC 9754 (SMT disabled) | 2x EPYC 9754 (SMT enabled) | 9754 SMT disabled Performance increase | 
|---|---|---|---|---|
| 1 billion | 13.480 | 13.56 | 14.573 | 7.50% | 
| 2.5 billion | 23.680 | 23.501 | 28.649 | 17.34% | 
| 5 billion | 40.819 | 40.547 | 50.082 | 18.50% | 
| 10 billion | 78.523 | 77.466 | 93.842 | 16.32% | 
| 25 billion | 206.399 | 206.078 | 236.070 | 12.57% | 
| 50 billion | 483.797 | 482.79 | 521.867 | 7.29% | 
| 100 billion | 1,269.484 | 1,266.83 | 1,253.446 | -1.28% | 
Analyzing the results
Diving into the subtleties of AMD SMT, there is an exciting dialogue within the tech community about its implications for system performance. At its core, SMT seems like a simple choice for those seeking improved performance. The theory is this: if enabling SMT can lead to ideal scaling, then why not adopt it as a beneficial architectural choice?
The relationship between SMT efficiency and the underlying architecture is not black or white. Poor SMT scaling does not necessarily indicate a flaw in its implementation. In fact, it could hint at a robust core design that leaves little room for SMT to make a notable difference. This paradox highlights a crucial point in the industry: processor manufacturers cannot claim a unique advantage with SMT or similar technologies. They recognize that while SMT can generate additional performance in some use cases, it is not without drawbacks in other scenarios.
Through the prism of high-performance computing and compute-intensive tasks, the limits of SMT become more obvious. While the idea of doubling the number of threads per core may sound promising, the reality is not comparable to doubling the number of cores. In extreme cases, this can lead to a drop in performance when threads contend for cache resources. Nevertheless, for the majority of multithreaded applications, especially those without cache contention, SMT improves performance, particularly in tasks that can fully exploit its potential.
Closing thoughts
AMD SMT is incredibly useful for a wide variety of common enterprise workloads. But not all workloads need or benefit from SMT. Through our testing, we have shown how AMD is able to leverage manufacturing variations to deliver a solid product with a unique value proposition. Organizations designing platforms for specific types of workloads requiring pure cores without SMT can save some money by purchasing the AMD EPYC 9754S, whose SMT is definitively disabled at the factory.
