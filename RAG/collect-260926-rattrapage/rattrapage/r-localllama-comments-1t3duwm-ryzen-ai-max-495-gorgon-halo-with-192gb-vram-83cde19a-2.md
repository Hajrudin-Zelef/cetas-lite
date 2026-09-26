---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1t3duwm-ryzen-ai-max-495-gorgon-halo-with-192gb-vram-83cde19a-2
title: "r-localllama-comments-1t3duwm-ryzen-ai-max-495-gorgon-halo-with-192gb-vram-83cde19a"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["llama", "amd", "benchmarks", "fp8", "gpu", "llama.cpp", "nvidia"]
source: docs/RAG/lot-rattrapage/servers-reviews/r-localllama-comments-1t3duwm-ryzen-ai-max-495-gorgon-halo-with-192gb-vram-83cde19a.md
source_anchor: ""
source_lines: [118, 154]
sha256: 489e6166525f8df189759258ac20ae9165f95a566c8ed18c78a2f6ef90cf579a
---

# r-localllama-comments-1t3duwm-ryzen-ai-max-495-gorgon-halo-with-192gb-vram-83cde19a

Section des commentaires
Je ne veux pas être ce gars, mais augmenter la RAM ne signifie pas que tu peux faire tourner un plus gros modèle efficacement. La vitesse de préremplissage est le point faible de cet appareil.
Je pense que ce serait bien pour exécuter plusieurs petits modèles pour diverses tâches à la place.
Oui, le préremplissage est une grande faiblesse. J'ai moi-même un Strix Halo.
Cependant, le Medusa Halo aura beaucoup plus de puissance de calcul et de bande passante mémoire. Donc, je suis optimiste.
De plus, cela met NVIDIA sous pression pour sortir un DGX Spark 2.
La mémoire supplémentaire pourrait être un énorme atout pour sa vitesse de traitement de requêtes insuffisante. Maintenant, vous pouvez jongler avec plusieurs contextes sans avoir à user vos SSD et éviter le besoin de retraiter complètement la requête.
Peut-être qu'à l'avenir, il y aura des architectures plus performantes avec un préfil de type Mamba & Co
Bien que je sois sûr que certaines personnes apprécieront la mémoire supplémentaire, quelques notes de quelqu'un qui a effectué des tests très approfondis sur Strixt Halo (et beaucoup de travaux sur le noyau sur RDNA3) :
La bande passante mémoire semble rester la même ? 256 Go/s théorique. Sur Strix Halo, la meilleure bande passante mesurée sur GPU que j'ai obtenue (en utilisant ROCm/rocm_bandwidth_test) était de 212 Go/s (83 % du maximum théorique), et la meilleure sur llama.cpp (tests tg de Llama-2-7B) était d'environ 180 Go/s (70 %)
Ce qui est pire, c'est que bien que le maximum théorique des TFLOPS FP16 soit d'environ 59,4, le plus rapide que j'ai trouvé avec mamf-finder était d'environ 37 TFLOPS (hipBLASLt), soit environ 62 % d'efficacité. Beaucoup de formes sont beaucoup pires.
Note, avec un long contexte, je pense que le calcul est en fait ce qui tue la vitesse de décodage. Tant que les APU AMD restent sur RDNA3, cela ne changera pas. Je serais hésitant à recommander Gorgon Halo même pour l'inférence LLM en 2026/2027
Si Medusa Halo passe à RDNA5 ou quoi que ce soit ayant une meilleure architecture pour l'IA/ML, tant mieux, sinon tu serais bien mieux avec pratiquement n'importe quoi d'autre (Mac Studio, GPU + station de travail/serveur avec K-Transformers, probablement même un DGX Spark).
Oui, ici ça reste surtout le même, c'est juste un Strix Halo Refresh. C'est pourquoi je vais l'ignorer.
Medusa Halo apportera un vrai progrès.
Je ne pense pas que cela puisse être résolu avec le DGX ou le Mac Studio. Le DGX a la même bande passante mémoire que le 395 et des performances de calcul similaires. Le Mac Studio a environ 3 fois la bande passante mémoire, mais souffre de performances de calcul inférieures (26 TFLOPS contre 37 mesurés sur le Strix), ce qui rend le traitement de longs prompts contextuels plus lent sur le Mac.
Le DGX a une bande passante mémoire similaire, mais le calcul n'est pas aussi similaire...
Puisque j'ai déjà créé le tableau il y a un moment à partir de mon guide wiki Strix Halo... https://strixhalo.wiki/AI/AI_Capabilities_Overview
Sur le papier, le BF16/FP16 est assez proche, cependant le FP8 est déjà 2X et l'INT8 est 4X sur le DGX. Ceci n'est qu'un matériel - en pratique, le rocBLAS et le hipBLASLt pour le RDNA3.5 ne sont également pas très performants...
Qu'est-ce que cela signifie pratiquement ? En regardant les benchmarks les plus récents du fil de discussion DGX llama.cpp , j'ai observé des modèles similaires aux benchmarks Strix Halo de kyuz0 et bien que aucun des quants exacts, d'après ce que j'ai regardé, le préremplissage pour le DGX est actuellement environ 2-5X plus rapide que Strix Halo.
Ayant fait beaucoup de travail sur le noyau pour RDNA3/3.5 aussi, je dirais que le max FP16 TFLOPs est inférieur au max théorique car les horloges n'atteignent tout simplement pas cette hauteur. Et hipBLASLt laisse encore de la place pour plus de TFLOPs sur la table, car il est possible d'obtenir en moyenne ~45 TFLOPs. C'est juste dommage qu'il semble qu'il y ait encore du travail à faire pour RDNA3.5 car la performance du noyau semble toujours fluctuer
Ce n'est pas juste le temps, je pense, même quand je regarde l'horloge et que je fais les calculs, tout est encore bien en dessous de ce que ça "devrait" être. Je pense qu'il y a beaucoup de choses, mais l'une des principales est que RDNA3 n'a pas été conçu pour l'IA et leur publicité est en gros un mensonge.
V_WMMA_F16_16X16X16_F16exécute un matmul de 16×16×16 sur la vague, prend 16 cycles à retirer sur un SIMD32 unique. *MAIS* alors que les cœurs tensoriels NVIDIA ont du matériel dédié qui fonctionne en parallèle avec le pipeline SIMD, pour AMD, WMMA est le pipeline SIMD puisque c'est une instruction ALU. Donc chaque opération WMMA bloque les mêmes ports VGPR que les opérations scalaires utiliseraient, et vous ne pouvez pas superposer WMMA avec l'accumulation FP32 (vous devez décompacter). Puisque le nombre maximum de débit suppose un enlèvement WMMA consécutif sans opérations dépendantes entre elles, et puisque des travaux qui ne sont pas WMMA doivent toujours être faits (mise à l'échelle, softmax, masquage, etc.), WMMA va être displacé et vous n'obtiendrez jamais le maximum de FLOPS.
L'autre chose que j'ai trouvée avec mes tests, c'est que les VGPRs sont beaucoup trop bas pour cacher la latence. De plus, le trafic LDS est nul, surtout si vous faites de l'FA. Oh, et le compilateur est toujours nul aussi. Les instructions WMMA sont programmées avec des opérations dépendantes trop rapprochées, un empaquetage redondant, de mauvais temps d'attente mémoire, toutes sortes de choses qui bloquent le pipeline, donc si vous voulez quelque chose de mieux, vous devez régler à la main (eh bien, de nos jours, par l'IA).
À mon avis, pour ceux qui ont 395 en ce moment, Medusa Halo en 2027 est la seule mise à niveau qui en vaille la peine.
Je vais attendre Méduse
les principaux problèmes de l'AI MAX+ 395 ne sont pas un RAM insuffisant, mais un iGPU trop lent.
iGPU trop lent, manquant de compatibilité avec des formats de données importants, avec un bus étroit et une mémoire lente.
Plus de 256 Go, puisque Medusa Halo passe à un bus LPDDR6 de 384 bits. En utilisant la même densité de die qu'un Gorgon Halo de 192 Go, ça donne 288 Go. Avec des dies plus grands (ceux qui seraient nécessaires pour un Gorgon Halo de 256 Go), nous obtenons 384 Go pour Medusa Halo.
192 Go c'est cool mais comment y parviennent-ils ? Y aura-t-il une bande passante supplémentaire ou passent-ils juste à de la RAM plus dense ? La bande passante est le problème critique.
édit : c'est mort à l'arrivée, passez votre chemin, faites la queue pour le train de hype de Medusa Halo.
Même GPU, je vais attendre jusqu'à la génération suivante, j'espère avec une mémoire plus rapide et un GPU plus puissant pour l'inférence.
c'est un changement mineur du point de vue d'amd. les modules de mémoire fournis sont simplement plus denses.
probablement pas d'amélioration réelle de la bande passante de la mémoire.
donc cela ne coûtera pas beaucoup plus.
GPT 120b a10b ne sera probablement que légèrement plus rapide que minimax 230b a10b mais la grande différence réside dans l'intelligence et maintenant vous pouvez charger minimax, c'est la différence.
étant donné ma tendance à naviguer dans un contexte de 200 000 avec minimax tout le temps. Je me demande quelles vitesses j'obtiendrai, mais je vais acheter :)
