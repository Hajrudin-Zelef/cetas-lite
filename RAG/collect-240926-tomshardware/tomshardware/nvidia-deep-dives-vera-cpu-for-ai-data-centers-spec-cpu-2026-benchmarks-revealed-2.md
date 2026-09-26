---
id: collect-240926-tomshardware/tomshardware/nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed-2
title: "nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["nvidia", "amd", "decode", "gpu", "intel", "lpddr5x", "memory", "research", "rubin"]
source: docs/RAG/clean_en/tomshardware/nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed.md
source_anchor: ""
source_lines: [56, 100]
sha256: f3e07be84750fc0a3e2e5522b890079d92011d4e9af76d86805a40cd99e675ea
---

# nvidia-deep-dives-vera-cpu-for-ai-data-centers-spec-cpu-2026-benchmarks-revealed

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
 
