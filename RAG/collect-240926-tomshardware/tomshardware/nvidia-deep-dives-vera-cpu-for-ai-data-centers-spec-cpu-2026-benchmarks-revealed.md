---
id: collect-240926-tomshardware/tomshardware/nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed
title: "nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "AWS", "Cohere", "Google", "Intel", "Meta", "Nvidia"]
dates: []
keywords: ["benchmark", "benchmarks", "nvidia", "agent", "agentic", "amd", "aws", "chiplet", "decode", "gpu", "gpus", "graviton"]
source: docs/RAG/clean_en/tomshardware/nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed.md
source_anchor: ""
source_lines: [1, 152]
sha256: 60a43d75afde5ccd51d38613f4fe7b6494c92c3b870012528484fea9ee105723
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

En plus des résultats ci-dessus, Nvidia a partagé quelques tests du système mémoire de Vera comparé à Turin. Ces microbenchmarks sont utiles pour valider les spécifications de Nvidia, mais ils examinent la performance architecturale, pas la performance applicative. Un avantage architectural se traduit par un avantage de performance, mais pas toujours de manière linéaire et attendue.

Nvidia a utilisé des outils développés en interne pour les tests mémoire, bien qu’ils soient disponibles sur GitHub pour que quiconque puisse les exécuter.

Le premier est la latence mémoire en charge, sollicitant le sous-système mémoire à mesure que l’utilisation de la bande passante augmente. Vera dispose d’une bande passante globale bien plus élevée, mais vous pouvez voir que la puce Turin atteint un mur de latence en dessous de son maximum, ce que Nvidia attribue au parcours des domaines Non-Uniform Memory Access (NUMA) et à la latence CCD-à-CCD.

En examinant la bande passante par cœur, Nvidia affirme que Vera fournit plus de quatre fois la bande passante du 9755 d’AMD. La suggestion ici est que la bande passante par cœur « réelle » est encore meilleure que ce que les spécifications de Nvidia laissent entendre (ou peut-être pire que celle d’AMD).

Peut-être le plus conséquent de ces tests est celui que vous pouvez voir ci-dessus, examinant la latence cœur-à-cœur. Ce n’est un secret pour personne que traverser le CCD sur l’architecture à chiplets d’AMD entraîne une pénalité de latence importante. Vous pouvez le constater même dans notre __test du Ryzen 9 9950X3D2__, et les pénalités s’accumulent à mesure que vous augmentez le nombre de CCD.

Pour être juste envers AMD ici, les conceptions à base de chiplets ne sont pas conçues pour ce type de traversée inter-CCD, préférant garder les charges de travail localisées et optimiser pour la densité de cœurs. L’objectif de conception de Vera est clairement de maintenir des latences cohérentes sur toute la puce, en sacrifiant la densité de cœurs au passage. Ian Buck de Nvidia nous a déclaré que ce compromis de conception « se fera au détriment de la charge de travail héritée », lors de notre récente visite au siège de Nvidia.

C’est un contexte important. Nvidia ne cherche pas tant à voler des parts de marché existantes à AMD et Intel qu’à saisir des parts de marché dans un marché en expansion avant qu’AMD et Intel ne puissent le faire. Certaines institutions financières (dont Morgan Stanley et Bank of America) suggèrent que le marché des CPU pour serveurs pourrait doubler de taille (ou croître encore davantage) d’ici 2030. Ce contexte est important car il y aura une demande continue de CPU capables de gérer des charges de travail pour lesquelles Vera n’est pas optimisée, et il sera intéressant de voir comment AMD et Intel abordent cette dynamique avec de futurs produits, en essayant de conserver une base de clients héritée tout en progressant vers le marché élargi.

Nvidia a clairement une vision de ce à quoi ressemble ce marché élargi, et à cette fin, n’a pas partagé les résultats en virgule flottante de SPEC CPU. Vraisemblablement, cela est dû au fait que la suite vectorisée de SPEC est principalement axée sur les charges de travail HPC, tandis que Nvidia s’est concentré sur ce qu’elle considère comme des charges de travail agentiques critiques, basées sur les entiers. Vera dispose d’un moteur vectoriel complet avec SVE, mais cela ne semble pas être la priorité de Nvidia.

Dans un système Nvidia de bout en bout, ces charges de travail vectorisées seraient déchargées vers un GPU Rubin. Pourtant, nous n’avons pas encore de résultats vectoriels pour Vera. Jusqu’à présent, nous n’avons vu que des résultats entiers, ce qui est étrange étant donné le système mémoire en jeu dans Vera.

Page actuelle : CPU Nvidia Vera

Next Page Nvidia Vera CPU architecture — A closer look at the Olympus core
- 
Reply
 Only if you're talking about server CPUs. Before that, they made at least two generations of cores for tablet/embedded SoCs (i.e.The article said:Vera is Nvidia’s first CPU with a core design created in-house, which is the Olympus core.**Denver** and**Carmel** ). See: https://en.wikipedia.org/wiki/Project_Denver
 
 Ah, but you missed an interesting detail. The way you reach 1.2 TB/s with 9600 MT/s memory is by scaling the memory datapath up to 1024 bits, which is the equivalent of 16x DDR5 DIMMs! Because SOCAMM2 each have a 128-bit data width, you only need 8 of them. I was staring at a photo of the Vera board and wondering how they possibly reached 1.2 TB/s, and that's when it clicked that they actually packed a 128-bit datapath in those little memory modules!The article said:It uses a SOCAMM2 LPDDR5X memory system with capacity of up to 1.5 TB and speeds up to 9600 MT/s
 
 Looking at the block diagram of the core, I'm struck by a few things.
 The L1i cache is the same size as Intel's Lion Cove, but surprisingly only 4-way. Intel used 8-way associativity for their I-cache for almost the past 20 years!
 At 64 entries, its iTLB is also much smaller than Golden Cove (the last Intel core I could find data for), which has 256 entries.
 However, it regains some ground by being fully-associative, whereas the iTLBs of Intel CPUs usually have only 8-way associativity.
 The dTLB is more comparable to the size recent Intel CPUs use (112 vs. 128 entries in Lion Cove), but also fully-associative whereas Intel typically uses only 4-way associativity.
 At 96k, the L1d is running 2x of Intel's. As with the L1i, its associativity is less (6-way instead of 12-way for Intel).
 Its L2 cache is comparable to Raptor Cove's in size, but it's only 8-way associative whereas Raptor had 16-way associativity.
 Its STLB is similar in size to the latest Intel core I have data on (3k vs. Golden Cove's 2k).
 
 AMD said it used a neural network-based branch predictor in the original Zen cores. Jim Keller briefly discuss it, in an interview. They apparently had an open contest for people to submit branch predictor algorithms and the neural one simply beat all of the others. If you search the patent database, you might be able to find some on it. That's probably the best chance of finding out how it actually works.The article said:Research on neural branch prediction dates back to the late 90s
 
 
 The decoder in Zen 5 is hard-partitioned between threads, limiting each to only 4 instructions per cycle. Even if there's only one thread running on the core (or if you disable SMT in BIOS), it still uses only a single 4-wide decoder.The article said:At the last stage of the front end is that 10-wide decode, feeding more instructions into the execution engine per cycle than the 8-wide decode in AMD’s Zen 5 microarchitecture.
 
 But. We should consider that x86 instructions can encompass more functionality than an ARM64 instruction. So, maybe a typical ARM instruction stream needs like 5 instructions to do the same work as 4 x86 instructions. Just spitballing, here. It's not a huge difference, since most instructions are fairly simple and operating register-to-register, whereas x86 instructions get a lot of their density advantage by combining memory loads/stores with some arithmetic or logical operation, and supporting more sophisticated address arithmetic.
 
 Looking only at decode-width is a little misleading, because CPUs like Zen 5 have much wider dispatch (8-way) from their micro-op cache. Since these are now micro-ops, they're more equivalent to ARM64 instructions, in terms of how much work each represents.
 
 BTW, there are two other obvious points of comparison:
 The Neoverse V2 cores, used in Nvidia's prior Grace CPUs, have a 6-wide decoder* (**source:** https://chipsandcheese.com/p/hot-chips-2023-arms-neoverse-v2 )
 The Cortex-X925 cores, used in Nvidia's RTX Spark, have a 10-wide decoder (**source:** https://chipsandcheese.com/p/arms-cortex-x925-reaching-desktop )
 * Note that Neoverse V2 still has a mOP cache with 8-wide dispatch. So, as with x86 P-cores, the decoder width is a little bit deceptive.
 
 
 You skipped a pretty big detail:The article said:Past the front end, the mid-core rename / allocation engine is built to keep instructions moving while waiting on dependencies.**no mOP cache!**
 Modern x86 P-cores and some 64-bit ARM cores had micro-op caches to avoid having to re-decode the same instructions. Once ARM dropped 32-bit compatibility, they started getting rid of those. Also, Intel either doesn't have them in their E-cores, or perhaps what they did was to put some of that into the I-cache.
 
 So, it's interesting (but not all that surprising) that Nvidia followed ARM's approach of just skipping the mOP cache, entirely.
 
 
 It's darkly ironic that SVE's main selling point was to allow CPUs to scale all the way up to 2048 bits per vector, but all mainstream implementations (except for AWS Graviton 3) are just 128-bit.The article said:For SIMD instructions, the execution engine includes a vector cluster for Arm’s Scalable Vector Extension (SVE), including six vector units that support 128-bit SVE instructions
 
 
 Not really. The way AMD describes SMT in Zen 5 is that certain competitively-shared resources have watermarks that limit how much a single thread is able to use, so that it doesn't starve out the other thread. However, when there's only one thread running on a core, more of those constraints go away and only the statically-partitioned resources (e.g. the decoder) remain exclusive.The article said:Traditional SMT time-slices execution, giving both threads access to all of the core resources and sharing them as instructions execute in parallel.
 
 My read on "Spatial Multithreading" is that it's Nvidia's marking machine trying to spin a weakness to make it sound more like a strength. This being Nvidia's first SMT implementation (AFAIK), it won't have the same sophistication as where Intel and AMD have gotten, over a couple decades of experience implementing and refining theirs.
 
 
 First, Intel has traditionally done the same thing. However, I need to catch up on what they've said about Sierra Forest.The article said:Nvidia’s second-generation Scalable Coherency Fabric (SCF). It underpins Nvidia’s approach of using a monolithic die as opposed to a chiplet-based design, distributing last level cache in a mesh across the die and avoiding the cross-CCD latency penalty with localized L3.
 
 Second, people tend to overestimate the impact of the cross-CCD thing. Unlike Intel CPUs, one AMD CCD won't write data to another's L3 slice. Each is private to that CCD, except for coherency. That said, when you've got multiple threads that are either exchanging or both modifying the same data, and they happen to be scheduled on different CCDs,*that* is when you feel the impact of the die-to-die communication.
 
 Anyway, AMD's approach has clearly scaled better. I think that's why Sierra Forest looks like it moved in the direction of doing the same thing with segmenting its L3 cache.
 
 
 I don't know what happened here, but the slide with the annotated die photo shows it supporting PCIe 6 with only 16 lanes.The article said:For I/O, Vera supports PCIe 6.4 with 88 lanes per CPU and bifurcation support down to x2.
 
 **P.S.** As for the benchmarks, the initial Phoronix review was interesting, but now I'm just waiting for some independent analysis. From a CPU microarchitecture standpoint, what I really want to know is the single-threaded performance across the whole SPEC suite, and how the sub-scores compare with other CPUs. MT scaling is mildly interesting, but it really stacked that deck in its favor by having relatively few cores and huge amounts of memory bandwidth. So, I expect it scales well, but not well enough to outright beat Intel or AMD on anything that's not fundamentally bottlenecked by memory bandwidth.
- 
Reply
 Announced? Venice isusertests said:But their comparisons (obviously) are to Turin rather than newly announced Venice.*launching* in about a week! That's probably why Nvidia is making a bunch of noise about this*now!*
- 
I like NVIDIAs marketing department scaling results per core. This is the height of aerobatics. Interesting how all would compare per core with the AMD 64core Turin 9575F, their king of single core benchmarks ? And how per core comparison will look for just the 64 core Venice? Huang will leave full of tears on leather jacketReply
- 
Reply
 Good question. Where Vera has a big lead over AMD's Turin is that it has basically 2x the memory bandwidth. Also, ARM uses memory bandwidth a little more efficiently, due to its relaxed memory consistency model.Stomx said:I like NVIDIAs marketing department scaling results per core. The height of aerobatics. Interesting how all would compare per core with the AMD 64core Turin 9575F, their king of single core benchmarks ?
 
 Where AMD would win biggest is going to be on AVX-512 stuff that's not memory-bound.
 
 Otherwise, it looks to me like Vera might have a slight edge over Zen 5. I mean it's a 10-way dispatch core vs. an 8-way one. Even Zen 6 will remain 8-way, if the rumors are correct. Not to say dispatch width is everything, since Intel's cores are also wider, but it's at least a good starting point. I think Zen 5 does have 3 integer multiply ports vs. Vera's 2.
 
 
Yeah. With Venice, AMD is going to leap-frog Nvidia's memory bandwidth by approximately 33%, according to what they've said. That's to be accomplished using 16 DIMMs (matching the same data width as Vera of 1024 bits) using MRDDR5-12800. So, on low core-count models, they'd stomp Nvidia on per-core memory bandwidth.Stomx said:And how per core comparison will look for just the 64 core Venice? Huang will leave full of tears on leather jacket
- 
Reply
 I think that explains why it doesn't have a lot of cores, but it doesn't explain why Nvidia went to the trouble of designing their own cores instead of just licensing the next Neoverse V-series core, like they did in their Grace CPU and like Amazon did since Graviton 3.palladin9479 said:Performance isn't really important, it's only real job is to manage the workloads for the GPU like devices connected to it. It only needs to be fast enough to not make the GPUs wait around.
 
 Arm's Neoverse V3 seems to be derived from their Cortex-X4, which is also a 10-way core (source: https://www.androidauthority.com/arm-cortex-x4-explained-3328008/ ), though it has only 4x SVE2 pipes. However, I think the X4 and V3 are probably more efficiency-focused and don't clock very high. Graviton 5 uses Neoverse-V3 at 3.3 GHz, for instance.
 
So, it seems like the main things it seems Nvidia got by making its own core was the ability to achieve higher single-thread perf by hitting higher clock speeds and the ability to add SMT for better utilization.
