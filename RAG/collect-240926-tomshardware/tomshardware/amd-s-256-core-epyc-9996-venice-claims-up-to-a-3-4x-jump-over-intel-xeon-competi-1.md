---
id: collect-240926-tomshardware/tomshardware/amd-s-256-core-epyc-9996-venice-claims-up-to-a-3-4x-jump-over-intel-xeon-competi-1
title: "amd-s-256-core-epyc-9996-venice-claims-up-to-a-3-4x-jump-over-intel-xeon-competi"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "AWS", "Intel", "Nvidia", "TSMC"]
dates: []
keywords: ["amd", "intel", "agentic", "agents", "aws", "benchmark", "benchmarks", "clearwater forest", "compute", "consumer", "lpddr5x", "memory"]
source: docs/RAG/clean_en/tomshardware/amd-s-256-core-epyc-9996-venice-claims-up-to-a-3-4x-jump-over-intel-xeon-competi.md
source_anchor: ""
source_lines: [1, 88]
sha256: 3362b58be44b71b9f729ca8e9fedbdcdaa3ee3b4138afb6b73c585ea0c477b8c
---

# amd-s-256-core-epyc-9996-venice-claims-up-to-a-3-4x-jump-over-intel-xeon-competi

<!-- source: https://www.tomshardware.com/pc-components/cpus/amds-256-core-epyc-9996-venice-claims-up-to-a-3-4x-jump-over-intel-xeon-competition-20-percent-over-nvidia-vera-zen-6-comes-with-up-to-1024mb-of-l3-16-channel-memory-and-5ghz-clock-speeds -->

AMD is finally providing some details on its first Zen 6 CPU, which it has been teasing for over a year. The Epyc 9996 is a 256-core / 512-thread chip, packing AMD’s new Zen 6 architecture, and it’s the first to launch in what AMD describes as a “broad portfolio” for Venice. In addition to claiming significant performance advantages over the impending Nvidia Vera and Intel’s Xeon 6, AMD says it will continue to build out the Venice range with bespoke designs over the next year.

“It’s not just a single processor,” said AMD’s Ravi Kuppuswany, corporate VP of compute and enterprise solutions.q “It’s a portfolio.” AMD says it has purpose-built solutions, splitting its offerings depending on the application, not dissimilar to how Intel has split its Xeon ranges over the past few generations (nor how AMD has softly segmented its Epyc offerings). The roadmap starts with the main Venice lineup on the SP7 socket, which is what AMD has been teasing for so long. It scales up to 256 cores and 512 threads, 1.6 TB/s of memory bandwidth with fast MRDIMMs, and 128 PCIe 6 lanes in 1P configuration (160 lanes in 2P).

**Note:** When scaling up to 256 cores, AMD uses its Zen 6c “dense” design. With a standard Zen 6 design, AMD says Venice scales up to 128 cores and 256 threads, while high-frequency variations top out at 96 cores.

In the first half of next year, AMD plans to launch Venice on its SP8 socket, offering as few as eight cores and up to 128, focused on smaller deployments. These chips support eight-channel memory with two DIMMs per channel, and the same 128 PCIe 6 lanes.

Venice-X is expected in the second half of 2027, on the SP7 socket. We didn’t see Turin-X, but the last, last-gen Genoa-X came with 96 cores and up to 1152 MB of stacked L3 cache. Those specs haven’t changed (short of the Zen 6 microarchitecture), but AMD says it's able to clock Venice-X up to 5.15 GHz.

Finally, Verano should arrive in the second half of next year on the SP8 socket, and it looks like the most direct competitor to Vera (AMD’s Kuppuswamy had some fun with calling it “Vera-No”). It’s optimized to be an AI host node, says AMD, packing up to 72 cores and 5 GHz peak clocks. Critically, it comes with a 24-channel LPDDR5X memory system, leveraging SOCAMM2 modules.

| AMD Epyc 9006 SP7 specifications |  |  |  |  | 
|---|---|---|---|---|
| **Chip** | **Cores / Threads** | **Base / Boost Clock (GHz)** | **L3 Cache** | **TDP** | 
| Epyc 9996 | 256 / 512 | 2.55 / 4.1 | 1024 MB | 600W | 
| Epyc 9966 | 192 / 384 | 2.9 / 4 | 768 MB | 600W | 
| Epyc 9846 | 168 / 336 | 2.85 / 3.7 | 768 MB | 500W | 
| Epyc 9756 | 128 / 256 | 3.15 / 4 | 512 MB | 500W | 
| Epyc 9G76 | 96 / 192 | 3.4 / 4.8 | 384 MB | 500W | 
| Epyc 9656 | 96 / 192 | 3.05 / 3.7 | 512 MB | 400W | 
| Epyc 9686F | 96 / 192 | 3.4 / 5 | 384 MB | 500W | 
| Epyc 9556 | 64 / 128 | 2.75 / 4.3 | 384 MB | 300W | 
| Epyc 9586F | 64 / 128 | 3.75 / 5 | 384 MB | 500W | 

One of the advantages __Nvidia claims with its Vera chip__ is lots of memory bandwidth through the LPDDR5X system. AMD’s approach is different with Venice SP7. It’s scaling up to 16-channel memory with Venice SP7, with support for MRDIMMs running at 12,800 MT/s, or standard DDR5 RDIMMs running at 8000 MT/s.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

It’s a significant jump over Turin, which uses 12-channel memory, with support for RDIMMs running at 6400 MT/s. AMD claims per-socket bandwidth of 1.6 TB/s, significantly higher than the 1.2 TB/s available on Vera, and nearly triple the 576 GB/s per-socket bandwidth of Turin. Intel recently enabled 8000 MT/s RDIMMs on select Granite Rapids and Clearwater Forest SKUs, and it says support for MRDIMMs with speeds up to 8800 MT/s is coming in Q1 2027.

| AMD Epyc 9006 'Venice' SP8 specifications |  |  |  |  | 
|---|---|---|---|---|
| **Chip** | **Cores / Threads** | **Base / Boost Clock (GHz)** | **L3 Cache** | **TDP** | 
| Epyc 9746 | 128 / 256 | 2.9 / 4 | 512 MB | 400W | 
| Epyc 9736P | 128 / 256 | 2.7 / 3.7 | 256 MB | 360W | 
| Epyc 9736 | 128 / 256 | 2.7 / 3.7 | 256 MB | 360W | 
| Epyc 9676F | 96 / 192 | 2.8 / 3.7 | 256 MB | 400W | 
| Epyc 9646P | 96 / 192 | 2.8 / 3.7 | 256 MB | 300W | 
| Epyc 9646 | 96 / 192 | 2.8 / 3.7 | 256 MB | 300W | 
| Epyc 9576F | 64 / 128 | 3.55 / 5 | 384 MB | 400W | 
| Epyc 9536P | 64 / 128 | 3.25 / 4 | 256 MB | 300W | 
| Epyc 9526 | 64 / 128 | 3.25 / 4 | 256 MB | 300W | 
| Epyc 9476F | 48 / 96 | 3.65 / 5 | 192 MB | 330W | 
| Epyc 9456P | 48 / 96 | 3.2 / 3.7 | 256 MB | 265W | 
| Epyc 9456 | 48 / 96 | 3.2 / 3.7 | 256 MB | 265W | 
| Epyc 9376F | 32 / 64 | 3.8 / 5 | 192 MB | 285W | 
| Epyc 9356P | 32 / 64 | 3.6 / 4.5 | 192 MB | 250W | 
| Epyc 9356 | 32 / 64 | 3.6 / 4.5 | 192 MB | 250W | 
| Epyc 9336 | 32 / 64 | 3.15 / 3.7 | 128 MB | 195W | 
| Epyc 9276F | 24 / 48 | 3.8 / 5 | 96 MB | 230W | 
| Epyc 9256 | 24 / 48 | 2.85 / 4.5 | 96MB | 190W | 
| Epyc 9176F | 16 / 32 | 3.9 / 5 | 192 MB | 200W | 
| Epyc 9116 | 16 / 32 | 2.85 / 4.5 | 48 MB | 160W | 
| Epyc 9016 | 8 / 16 | 3.05 / 4.8 | 48 MB | 130W | 

Zen 6 is built on TSMC’s N2 (this has been previously confirmed). AMD confirmed that there are 32 cores on a CCD, along with two IODs. Keep in mind that the 32-core CCD is using Zen 6c, not full Zen 6. There has been plenty of speculation about 32-core CCDs in consumer Zen 6 CPUs, but that seems unlikely.

The 256-core configuration comes with a massive 1,024 MB of L3, nearly triple the amount of the Epyc 9965. This isn’t stacked cache, either; that will come with Venice-X. Each CCD has access to 128 MB or L3, or 4 MB per core, double what was available on Turin.

Although AMD has focused a lot of its teases on the 256-core Venice, the initial SP7 offerings will also hold a 96-core, high-frequency model that can clock up to 5 GHz.

## AMD shares first 256-core Epyc ‘Venice’ benchmarks

Unlike the __extrapolated performance AMD shared__ a few weeks back, we have some concrete benchmarks for the Epyc 9996 now. AMD has, unsurprisingly, focused the workloads around agentic AI. However, many of the workloads applicable for agentic AI are applicable elsewhere, as well, including high-concurrency networking tasks, code compilation, and media processing.

Note that AMD includes just the Epyc 9965 as a gen-on-gen comparison point in the charts above. This is a “dense” Zen 5 design with 192 cores. Results for the 128-core 9755 are included in the tables below.

Starting with front-end operations, AMD claims a 1.2x gen-on-gen improvement and a 2.8x improvement compared to Intel Xeon 6980P, with an NGINX web server using the WRK load generator. Unlike most of these competitive performance figures, AMD included the actual numbers for the benchmarks it ran in the footnotes, which you can see in the table below.

| **Chip** | **Max Request Per Second** | 
| Intel Xeon 6980P | 10,162,179 | 
| AWS Graviton5 | 15,331,108 | 
| AMD Epyc 9755 | 17,906,196 | 
| AMD Epyc 9965 | 24,320,476 | 
| AMD Epyc 9996 | 28,789,170 | 

In data-heavy workloads that are common among AI agents, AMD claims a 1.7x improvement over Turn, and a massive 3.4x over the Xeon 6980P. AMD used the __TPCx-AI benchmark__ to gather these results. The primary metric for this test is AI use cases per minute (AIUCpm), for which AMD shared the median result2. If you’re interested in more about the reporting of this benchmark, __Dell has published an extensive breakdown__.

| **Chip** | **AIUCpm** | 
| Intel Xeon 6980P | 1,750.36 | 
| AWS Graviton5 | 2,444.8 | 
| AMD Epyc 9755 | 2,704.19 | 
| AMD Epyc 9965 | 3,458.79 | 
| AMD Epyc 9996 | 5,982.91 | 

