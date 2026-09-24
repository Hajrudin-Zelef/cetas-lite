---
id: collect-240926-tomshardware/tomshardware/two-variants-of-nvidia-s-rtx-spark-show-up-on-geekbench-revealing-a-cut-down-18
title: "two-variants-of-nvidia-s-rtx-spark-show-up-on-geekbench-revealing-a-cut-down-18-"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Apple", "Google", "Intel", "Nvidia"]
dates: []
keywords: ["nvidia", "amd", "benchmark", "gpus", "intel"]
source: docs/RAG/clean_en/tomshardware/two-variants-of-nvidia-s-rtx-spark-show-up-on-geekbench-revealing-a-cut-down-18-.md
source_anchor: ""
source_lines: [1, 50]
sha256: 278193a58c51620e72b8bacd3eaba02aa939f83bc7a7d390ce91c6bdf3078f80
---

# two-variants-of-nvidia-s-rtx-spark-show-up-on-geekbench-revealing-a-cut-down-18-

<!-- source: https://www.tomshardware.com/pc-components/cpus/two-variants-of-nvidias-rtx-spark-show-up-on-geekbench-revealing-a-cut-down-18-core-model-full-20-core-beats-most-x86-mobile-chips-across-multi-core-and-single-core-tests -->

Two new Geekbench listings for Nvidia's upcoming RTX Spark laptop superchip have surfaced, confirming a second, cut-down SKU exists with 18 CPU cores, at least in the testing phases. Both parts score remarkably close to each other, posting impressive multi-core numbers, but their single-core performance isn't special. Still, these are preliminary results on unreleased hardware, so take these numbers with a grain of salt.

The model name for the 20-core (10+10) SKU was shown as "OEMQAJ 766_MIS Product Name DV" with its clocks reported at 4.0 GHz. This chip scored 2,570 points in the single-core test and 23,126 points in the multi-core test. The second listing showed up as "OEMQAJ OEMQAJ Product" with this SKU's 18 cores divided across 10+8 core clusters, clocked at 3.9 GHz. It scored 2,541 points in the single-core test, so within 1% of the 20-core part, while the multi-core was 21,776 points, which is about 6% faster than the full-fat chip.

When you compare these numbers to other Windows-running SoCs on the Geekbench database, these multi-core numbers are among the best. The Snapdragon X2 Elite Extreme currently holds the record for the highest multi-core score with 25,075 points, so both SKUs of the RTX Spark lose to it. But when we switch to the AMD AI Max+ 395, which scores 20,609 on average in the same test, even the 18-core Nvidia chip is 5.6% faster here, and the 20-core variant is about 12% faster.

Compared to Apple's lineup, the 20-core RTX Spark is 3.3% faster than the 14-core variant of the M3 Max, which scores 22,385 points in the multi-core bench. Considering that this is pre-release hardware with unoptimized drivers, we can expect it to even beat the maxed-out M4 Pro and the base M4 Max (which hover around 25,000 points) when it officially launches. Moreover, the 18-core SKU's 21,776 points are equal to, if not more than, the base M4 Pro, which also nets around 21,700 points in the multi-core test.

| Multi-core peformance |  |  | 
|---|---|---|
| SKU | Score | Difference | 
|---|---|---|
| **20-core RTX Spark** | 23,126 | *Baseline* | 
| **18-core RTX Spark** | 21,776 | -5.84% | 
| Snapdragon X2 Elite (Extreme X2E-94-100) | 25,075 | +8.43% | 
| Intel Core Ultra 9 275HX | 20,783 | -10.13% | 
| AMD AI Max+ 395 | 20,609 | -10.88% | 
| Apple M3 Max (14-core) | 22,385 | -3.20% | 
| Apple M4 Pro (14-core) | 24,901 | +7.68% | 

That being said, the swan song falls apart when we bring up the single-core numbers. Every Apple Silicon chip from the base M3 onwards beats the 2,570 points the 20-core RTX Spark achieved in this run. Only the M2 Max with its 2,477 points is bested by the Nvidia offering, and that too by just 3.75%. Apple's current-gen base M5 is a whopping 41% faster here. On the Windows side, the Snapdragon X2 Elite and X2 Elite Extreme score roughly 3,000 points, but AMD and Intel's top-end models take a beating.

The Ryzen AI Max+ 395 only achieves 2,429 points in the single-core test, making the 20-core RTX Spark 5.8% faster, and the 18-core variant 4.6% faster. The same goes for the Ryzen AI 9 HX 370 because its 2,432 points are within the margin of error of the Strix Halo chip. Intel's highest single-core score is achieved by the Core Ultra 9 275HX at 2,467 points, making it 3% slower than the 18-core RTX Spark and 4.2% slower than the 20-core SKU. However, Intel's actual fastest chip, the Core Ultra X9 388H, scores 2,767 points, making it about 7% faster than even the 20-core RTX Spark.

| Single-core performance |  |  | 
|---|---|---|
| SKU | Score | Difference | 
|---|---|---|
| Apple M4 Pro (14-core) | 3,315 | +28.99% | 
| Snapdragon X2 Elite (Extreme X2E-94-100) | 3,051 | +18.72% | 
| Apple M3 Max (14-core) | 2,841 | +10.54% | 
| **20-core RTX Spark** | 2,570 | *Baseline* | 
| **18-core RTX Spark** | 2,541 | -1.13% | 
| Intel Core Ultra 9 275HX | 2,467 | -4.01% | 
| AMD AI Max+ 395 | 2,429 | -5.49% | 

All in all, this is a much better showing for Nvidia's upcoming laptop Superchip than previous leaks. It's clear that driver optimization is slowly unlocking the full potential of the silicon, and by the time RTX Spark is in the hands of reviewers, we expect it to compete fiercely with every Windows machine. Nvidia is working hard to ensure a proper launch with a commitment to making DRM and anti-cheat work on day one.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*

Hassam Nasir is a die-hard hardware enthusiast with years of experience as a tech editor and writer, focusing on detailed CPU comparisons and general hardware news. When he’s not working, you’ll find him bending tubes for his ever-evolving custom water-loop gaming rig or benchmarking the latest CPUs and GPUs just for fun.

- 
These results don't seem right. Intel 275HX is much faster than what's in those tables. Maybe these results are done with constrained CPU wattage? But most definitely don't reflect the accurate full CPU multi core performance...Reply
- 
So I have an issue with benchmark for the AMD 395+ Strix Halo; my HP Z2 Mini G1a Workstation Desktop PC was scoring much higher than shown above for Geekbench 7.0.0 Multi was 27784 and single core was 2664 on Windows 11 25H2: https://browser.geekbench.com/v7/cpu/106333 and looking at multiple geekbench scores the score above is at bottom end; if you run Linux I see as high as 30K in multi core performance. It would be nice to see high to low scores ranges so people can see the variability of where they may actually fall for each processor.Reply
