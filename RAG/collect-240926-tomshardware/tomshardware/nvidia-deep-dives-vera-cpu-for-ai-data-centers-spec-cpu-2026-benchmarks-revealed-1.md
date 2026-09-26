---
id: collect-240926-tomshardware/tomshardware/nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed-1
title: "nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Google", "Intel", "Meta", "Nvidia"]
dates: []
keywords: ["benchmark", "benchmarks", "nvidia", "agent", "agentic", "amd", "intel", "latency", "memory", "throughput"]
source: docs/RAG/clean_en/tomshardware/nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed.md
source_anchor: ""
source_lines: [1, 55]
sha256: 48a73776c44cdfe9777c705aaea16fce3c7a486e6a6e20ccc0c06d1f8b4cf84e
---

# nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed

<!-- source: https://www.tomshardware.com/pc-components/cpus/nvidia-spills-the-beans-on-vera-cpu-spec-benchmarks-revealed-olympus-architecture-detailed-and-more -->

Nvidia’s Vera CPU is its first bid to become a key player in the data center CPU market. Although Grace has seen some success (most notably with Grace standalone deployments at Meta), Vera is Nvidia’s first CPU with a custom core design. It’s arriving at an ideal time, as well, with the server CPU market exploding in the last few months on the back of agentic AI demand.

Vera isn’t a chip built to chip away at the market share of AMD and Intel in the cloud. It’s built to grab market share in an expanding market, as hyperscalers look to widen AI infrastructure beyond legacy clouds. As such, it’s designed in a much different way than Nvidia’s x86 competitors, and it even holds some unique architectural design points compared to the swath of Arm-based designs.

Nvidia has slowly revealed more details about Vera as it ramps into general availability, which is on track for the back half of this year. Now, we have a full picture of the chip. Nvidia shared its Vera white paper, along with unofficial SPEC CPU 2026 results comparing Vera to AMD’s Turin-based Epyc 9755.

We’re going to break down the white paper here, including all of the details about the Olympus core and a look at the benchmarks Nvidia ran. At the end of this piece, we’ll also take a brief look at the larger context of Vera and how it fits into Nvidia’s wider AI ecosystem compared to standalone deployments.

But plenty of ink has been spilled about Vera’s technical capabilities and Nvidia’s next-gen AI infrastructure vision. Let’s start with the important thing: the benchmarks.

## Nvidia Vera CPU benchmarks

We’ve seen Vera in action before, though only through a series of __selected benchmarks ran at Nvidia HQ by Phoronix__. In the Vera white paper, Nvidia shared benchmarks for SPEC CPU 2026, specifically the integer suite from SPECrate, against AMD’s Epyc 9755, with both chips running in a dual-socket configuration. Before getting into the results, there are some important notes about how SPEC runs work, and the reporting criteria for them.

Nvidia’s run here isn’t official, as Vera was tested in a reference system due to the fact that it’s not broadly available yet. It’s ramping for general availability in the second half of the year. Due to that, Nvidia is unable to report its results. That’s why you see “estimated” in some of the charts below. Nvidia ran SPEC CPU 2026; it’s not extrapolating expected performance __like we’ve seen from AMD so far__ with its upcoming Venice chips.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

SPEC CPU 2026 is split into four suites, but Nvidia tested the SPECrate integer suite, which is focused on system throughput with integer-based workloads. The “rate” result is looking at how much work is completed within a certain amount of time. Here, each thread in the system has a copy of the workload. The score is how much time it takes for those workloads to complete, regardless of thread count, naturally giving chips with more cores an advantage.

If you want more detail on the benchmarks included in the suite, make sure to read our __original coverage of SPEC CPU 2026__. Here are the overall results:

| **Test** | **Run Time**  | **Rate** | 
| 706.stockfish_r | 324 | 1370 | 
| 707.ntest_r | 251 | 830 | 
| 708.sqlite_r | 250 | 744 | 
| 710.omnetpp_r | 203 | 842 | 
| 714.cpython_r | 136 | 1240 | 
| 721.gcc_r | 296 | 817 | 
| 723.llvm_r | 196 | 909 | 
| 727.cppcheck_r | 142 | 890 | 
| 729.abc_r | 196 | 823 | 
| 734.vpr_r | 199 | 815 | 
| 735.gem5_r | 131 | 1300 | 
| 750.sealcrypto_r | 231 | 816 | 
| 753.ns3_r | 129 | 1670 | 
| 777.zstd_r | 469 | 483 | 
| **Overall base score** | Row 15 - Cell 1 | **925** | 

Nvidia didn’t share the exact results for the 9755 it tested, short of the overall score of 898. Taking that overall score into account, Vera is 3% ahead of the 9755. It’s worth noting that Vera is ahead here despite a large thread disadvantage. An overall score of 898 for a dual-socket Epyc 9755 system isn’t unreasonable compared to publicly-submitted SPEC CPU 2026 runs, though higher results have been published. SPEC CPU ships as source code, which the tester must compile with their compiler of choice, and that can heavily influence results (particularly with vendor-specific compilers). Nvidia used GNU 15.2 with both systems.

Above, you can see Vera’s results stacked up against the 9755, but these aren’t comparing the numbers directly. Nvidia has normalized the per-core performance, which isn’t how SPECrate results are normally shared. According to the overall numbers, Vera is still completing more work within the same amount of time, despite a thread disadvantage, but the margins aren’t in the range of a 70% or 80% advantage as the above chart suggests.

We asked Nvidia about the results given that they're obfuscated by comparison; we could not reverse-engineer the Epyc 9755's scores with the information Nvidia has provided. Here's the response it gave: "Per-core performance under a fully loaded socket is important because agentic AI and RL run many sandboxes concurrently, while each agent step remains sequential and latency-sensitive. It measures how much performance each core sustains amid contention for shared power, memory, cache, and fabric. We therefore normalize by physical core, with SMT enabled on both systems."

The “agentic” workloads Nvidia has highlighted here are code compilation and interpretation workloads, which is something an agent is often doing, querying repos for dependencies and building source code. Below are data science workloads (or Exploratory Data Analysis), and below that are data processing workloads like SQLite database management. The results here align with Nvidia’s overall messaging of Vera, that it’s highly competent at data-rich, backend operations.

Bien que Nvidia partage des résultats par thread, elle soutient que SPECrate reste le bon benchmark à exécuter. Les résultats par thread ici s’inscrivent dans le contexte d’un socket pleinement chargé. Voici la justification tirée du livre blanc : « Cette métrique est non triviale pour l’IA agentique et les systèmes RL, où de nombreux sandboxes, outils et environnements s’exécutent simultanément plutôt que comme des tests monothread isolés. La performance par cœur en pleine charge capture la capacité de chaque cœur à maintenir son débit tout en partageant l’alimentation au niveau du socket, la bande passante mémoire, le cache et les ressources de fabric. »

En plus d’exécuter les charges de travail, Nvidia a analysé l’exécution du code pour les benchmarks architecturaux, que vous pouvez voir dans la galerie ci-dessus. Nvidia revendique un gain global d’IPC allant jusqu’à 1,9x par rapport à Turin, jusqu’à 2,3x plus de prédictions de branchement et 3,5x plus de branchements pris par cycle, et jusqu’à 2,4x plus d’opérations de fetch d’instructions par cycle.

En dehors de SPEC, Nvidia a partagé quelques benchmarks mettant en avant les capacités du cœur Olympus. En premier lieu PageRank, un algorithme développé par Google pour classer à l’origine les pages web, qui met en avant le moteur de prefetch d’Olympus. Nvidia a mis à l’échelle cette charge de travail vers des nombres de cœurs plus élevés, montrant que Vera maintient une grande partie de sa performance monocœur jusqu’à 32 cœurs, tandis que la puce Turin atteint un mur autour de 20 cœurs.

