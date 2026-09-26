---
id: collect-240926-tomshardware/tomshardware/cpu-benchmarks-and-hierarchy-2026-cpu-rankings-1
title: "cpu-benchmarks-and-hierarchy-2026-cpu-rankings"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["benchmark", "benchmarks", "amd", "compute", "cost", "gpu", "intel", "nvidia"]
source: docs/RAG/clean_en/tomshardware/cpu-benchmarks-and-hierarchy-2026-cpu-rankings.md
source_anchor: ""
source_lines: [1, 64]
sha256: 5bb3c3b499b99b257ef7483466e051d65879699067591f8fc6627fdd9c5e9e91
---

# cpu-benchmarks-and-hierarchy-2026-cpu-rankings

<!-- source: https://www.tomshardware.com/reviews/cpu-hierarchy,4312.html -->

Our CPU benchmark hierarchy provides a broad view of relative performance for the latest Intel and AMD processors. Over the last 30 years, Tom’s Hardware has been benchmarking CPUs, and we use the rankings here as the basis of our __best CPUs for gaming__ and __best budget CPU__ rankings. We run over 200 individual tests for each CPU we look at, and that comprehensive performance is condensed here for a high-level view of how CPUs compare across gaming, single-threaded, and multithreaded performance.

Each of our CPU benchmarks helps expose different aspects of performance, from heavily-threaded code compilation and data science workloads to lightly-threaded web apps and audio encoding. We’re currently in the process of the biggest refresh to our CPU benchmarks hierarchy ever, spanning over a decade of processor releases. The results here provide the first half of that testing, focusing on DDR5 platforms that span the __AMD vs Intel__ product lineups. As we fill out our legacy benchmarks, you’ll see more CPUs added to our rankings. If you want to check the performance of older CPUs now, you can use the second page of this article to see our legacy benchmarks.

In games, __AMD’s Ryzen 7 9850X3D__ is the fastest CPU on the market, though other Zen 5 X3D offerings like the __Ryzen 9 9950X3D__ and __Ryzen 7 9800X3D__ aren’t far behind. X3D chips dominate the charts for gaming at 1080p, with the other exception being the relatively unpopular (and expensive) Ryzen 9 7900X3D. Otherwise, Intel’s last-gen Core i9-14900K is the fastest offering from Team Blue, with the new Core Ultra 7 270K Plus coming in slightly behind __with Intel’s new iBOT feature__.

Intel pulls out strong positions in applications; however, with the __Core Ultra 7 270K Plus__ topping the charts in single-threaded performance and coming in third in multi-threaded rankings. It’s only beaten by the Ryzen 9 9950X and its X3D variant, and only by a hair. Further, both of those CPUs cost about twice as much. AMD's recent Ryzen 9 9950X3D2 claims the top slot in overall performance, but at $900, it's too expensive for most buyers.

In each section below, we’ll show you the rankings for each CPU, as well as reveal what tests went into creating the rankings. We’ll also give you some pointers for benchmarking your own CPU to see how much performance an upgrade or overclock netted you, along with some common, easy-to-run benchmarks you can perform yourself.

### CPU Benchmarks Rankings 2026

In the album above, you can see our master charts for gaming, single-threaded, and multi-threaded performance for CPUs. For games, all of our testing was done with an Nvidia RTX 5090 FE, and for applications, our testing was done with an Nvidia RTX 2080 Ti FE. For applications, no compute is actively running on the GPU; it’s a glorified display output that shares a driver with our gaming GPU. You can find a full breakdown of the test benches we used at the end of this article.

### Gaming CPU Benchmarks Rankings 2026

| Gaming CPU Benchmarks Rankings 2026 |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|
| **CPU / (MSRP)** | **Street Price** | **1080p Gaming Score** | **Architecture** | **Cores/Threads (P+E)** | **Base/Boost Clock (GHz)** | **TDP / Maximum Power** | 
| Ryzen 7 9850X3D ($500) | $484 | 100% | Zen 5 X3D | 8 / 16 | 4.7 / 5.6 | 120W / 162W | 
| Ryzen 7 9800X3D ($480) | $415 | 97% | Zen 5 X3D | 8 / 16 | 4.7 / 5.2 | 120W / 162W | 
| Ryzen 9 9950X3D ($700) | $569 | 95.7% | Zen 5 X3D | 16 / 32 | 4.3 / 5.7 | 170W / 230W | 
| Ryzen 9 9900X3D ($600) | $510 | 86.9% | Zen 5 X3D | 12 / 24 | 4.4 / 5.5 | 120W / 230W | 
| Ryzen 7 7800X3D ($450) | $330 | 85.6% | Zen 4 X3D | 8 / 16 | 4.2 / 5 | 120W / 162W | 
| Ryzen 9 7950X3D ($700) | $650 | 83.9% | Zen 4 X3D | 16 / 32 | 4.2 / 5.7 | 120W / 162W | 
| Ryzen 5 7600X3D ($300) | $240 | 80.6% | Zen 4 X3D | 6 / 12 | 4.1 / 4.7 | 65W / 88W | 
| Core i9-14900K ($550) | $396 | 78.2% | Raptor Lake Refresh | 24 / 32 (8+16) | 3.2 / 6 | 125W / 253W | 
| Core Ultra 7 270K Plus ($300) | $290 | 77.5% | Arrow Lake Refresh | 24 / 24 (8+16) | 3.7 / 5.5 | 125W / 250W | 
| Ryzen 7 7900X3D ($600) | Out of Stock | 77.1% | Zen 4 X3D | 12 / 24 | 4.4 / 5.6 | 120W / 162W | 
| Ryzen 9 9950X ($650) | $550 | 76.9% | Zen 5 | 16 / 32 | 4.7 / 5/7 | 170W / 230W | 
| Core i9-13900K ($590) | Out of Stock | 76.8% | Raptor Lake | 24 / 32 (8+16) | 3 / 5.8 | 125W / 253W | 
| Core i7-14700K ($410) | $365 | 76.4% | Raptor Lake Refresh | 20 / 28 (8+12) | 3.4 / 5.6 | 125W / 253W | 
| Core i7-13700K ($410) | Out of Stock | 75.8% | Raptor Lake | 16 / 24 (8+8) | 3.4 / 5.4 | 125W / 253W | 
| Ryzen 9 9900X ($500) | $338 | 73.9% | Zen 5 | 12 / 24 | 4.4 / 5.6 | 120W / 162W | 
| Core Ultra 5 250K Plus ($200) | $210 | 73.3% | Arrow Lake Refresh | 18 / 18 (6+12) | 4.2 / 5.3 | 125W / 159W | 
| Core i5-14600K ($320) | $361 | 72.8% | Raptor Lake Refresh | 14 / 20 (6+8) | 3.5 / 5.3 | 125W / 181W | 
| Ryzen 5 9600X ($280) | $173 | 72.6% | Zen 5 | 6 / 12 | 3.9 / 5.4 | 65W / 88W | 
| Core Ultra 9 285K ($590) | $499 | 71.8% | Arrow Lake | 24 / 24 (8+16) | 3.7 / 5.7 | 125W / 250W | 
| Ryzen 9 7950X ($700) | Out of Stock | 71% | Zen 4 | 16 / 32 | 4.5 / 5.7 | 170W / 230W | 
| Core i5-13600K ($320) | $339 | 70.9% | Raptor Lake | 14 / 20 (6+8) | 3.5 / 5.1 | 125W / 181W | 
| Ryzen 7 7700X ($400) | $233 | 70.6% | Zen 4 | 8 / 16 | 4.5 / 5.4 | 105W / 142W | 
| Core Ultra 7 265K ($400) | $284 | 70.3% | Arrow Lake | 20 / 20 (8+12) | 3.9 / 5.5 | 125W / 250W | 
| Ryzen 9 7900X ($550) | $313 | 69.2% | Zen 4 | 12 / 24 | 4.7 / 5.6 | 170W / 230W | 
| Ryzen 5 7600X ($300) | $163 | 67.3% | Zen 4 | 6 / 12 | 4.7 / 5.3 | 105W / 142W | 
| Core Ultra 5 245K ($320) | $193 | 67.1% | Arrow Lake | 14 / 14 (6+8) | 4.2 / 5.2 | 125W / 159W | 
| Core i7-12700K ($410) | $325 | 65.8% | Alder Lake | 12 / 20 (8+4) | 3.6 / 5 | 125W / 190W | 
| Core Ultra 5 225 ($183) | $180 | 62.5% | Arrow Lake | 10 / 10 (6+4) | 3.3 / 4.9 | 65W / 121W | 
| Core i5-12600K ($290) | $189 | 60.8% | Alder Lake | 10 / 16 (6+4) | 3.7 / 4.9 | 125W / 150W | 
| Core i5-14400 ($220) | $214 | 58% | Raptor Lake | 10 / 16 (6+4) | 2.5 / 4.7 | 65W / 154W | 

You can see the relative score for AMD and Intel CPUs above, measured against the Ryzen 7 9850X3D, which is the fastest gaming CPU on the market, per our testing. So, the Ryzen 7 9800X3D offers 97.04% of the performance of the Ryzen 7 9850X3D, while the Ryzen 9 7900X offers 69.28% of the performance. You can set any CPU as a baseline for comparison with Bench, which is available in *Tom’s Hardware Premium.* 

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

All of our gaming tests were run with the RTX 5090 FE at 1080p with a mixture of High and Ultra settings. We run each test multiple times — usually between three and five — and pick the median result. In other words, the results we use are real, recorded runs, not an average of several different runs. This is important as some games, such as *Far Cry 6,* show great CPU scaling but are otherwise inconsistent run-to-run.

In addition to consistent hardware (test benches at the end of this article), we use a consistent test image between platforms. That means the same GPU driver, the same Windows install, the game version, etc. We also tested with Virtualization-Based Security (VBS) turned off, Resizable BAR turned on, and automatic overclocking features disabled. That includes the Intel Extreme power profile and AMD’s PBO, both of which aren’t covered under standard warranty.

For this refresh, we tested 17 games and then calculated a geometric mean of the results. A simple average would provide skewed results with such a large test pool. A geomean provides a more realistic view of how each CPU compares to the others.

Here are the games that we used for testing:

