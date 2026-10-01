---
id: collect-250926-servers-hardware/servers-hardware/fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u-7b819f50-5
title: "fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "MiniMax", "Mistral", "Nvidia"]
dates: []
keywords: ["fp4", "fp8", "gpu", "llama", "mistral", "moe", "nvidia"]
source: docs/RAG/clean4/fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50.md
source_anchor: ""
source_lines: [109, 143]
sha256: 468a471d5072e8d9e5f43d3a2eb2d92aac5212547395912a6b15b1702df0a159
---

# fr-review-comino-grando-rtx-pro-6000-review-768gb-of-vram-in-a-liquid-cooled-4u--7b819f50

| Lama 3.1 8B Instruire | FP8 | 12,109 | 20,137 | 7,353 | 
| Lama 3.1 8B Instruire | FP4 | 11,954 | 20,206 | 7,239 | 
| Lama 3.1 8B Instruire | BF16 | 11,752 | 17,346 | 6,155 | 
| Codeur Qwen3 30B A3B | FP8 | 10,985 | 16,659 | 4,907 | 
| Codeur Qwen3 30B A3B | BF16 | 10,588 | 16,680 | 4,829 | 
| Mistral Small 3.1 24B | BF16 | 8,925 | 11,846 | 4,975 | 
| MiniMax M2.5 (230B) | ep_dp1 | 5,753 | 7,357 * | 2,555 | 
| Toutes les valeurs sont exprimées en tok/s, débit maximal à BS=256. *MiniMax M2.5 préremplissage important a atteint un pic à BS=128 (7 357 tok/s) ; BS=256 était de 7 141 tok/s. |  |  |  |  | 
GPT-OSS 120B et 20B
La famille de modèles GPT-OSS a été testée dans les configurations 120B et 20B sur le Comino Grando.
GPT-OSS 120B
Sous une charge de travail égale (256/256), le modèle 120B atteint 268.85 tok/s à BS=1, 6 666,23 tok/s à BS=64 et un pic de 11 726,04 tok/s à BS=256. Avec un préremplissage important (8k/1k), le débit commence à 1 375,69 tok/s, grimpe à 16 374,19 tok/s à BS=64 et 17 944,55 tok/s à BS=128, pour atteindre un pic de 21 636,41 tok/s à BS=256. Le décodage intensif (1k/8k) passe de 196.28 tok/s à BS=1 à 7 569,97 tok/s à BS=256, la latence étant bien contrôlée à des niveaux de concurrence inférieurs.
GPT-OSS 20B
Le modèle 20B atteint un débit de 334.80 tok/s à BS=1 sous une charge de travail égale, 10 303,56 tok/s à BS=64 et un pic de 17 280,12 tok/s à BS=256. En mode préremplissage intensif, le débit commence à 2 007,90 tok/s, grimpe à 24 990,46 tok/s à BS=64 et à 26 866,25 tok/s à BS=128, pour atteindre un pic de 32 060,72 tok/s à BS=256, soit le débit de préremplissage absolu le plus élevé enregistré pour les deux tailles de modèle. Le débit de décodage important passe de 286.08 tok/s à BS=1 à 11 187,36 tok/s à BS=256, offrant environ 1.5 fois le débit de décodage du 120B à la concurrence maximale tout en maintenant une latence plus faible.
Instructions Qwen3 Coder 30B A3B et instructions FP8
Le modèle Qwen3-Coder-30B-A3B-Instruct a été testé avec une précision BF16 et FP8.
Qwen3-Coder-30B-A3B-Instruction (BF16)
Sous une charge de travail égale (256/256), le modèle BF16 atteint 1 902,32 tok/s à BS=8, 6 683,58 tok/s à BS=64 et un pic de 10 587,56 tok/s à BS=256. Avec un préremplissage important (8k/1k), il démarre à 1 256,03 tok/s à BS=1, grimpe à 14 400,57 tok/s à BS=64 et 15 308,35 tok/s à BS=128, pour un pic de 16 679,52 tok/s à BS=256. Le décodage intensif (1k/8k) passe de 169.19 tok/s à BS=1 à 4 828,82 tok/s à BS=256, la latence étant bien contrôlée à des niveaux de concurrence inférieurs.
Qwen3-Coder-30B-A3B-Instruction (FP8)
Le modèle FP8 offre un débit comparable à celui du BF16 dans la plupart des scénarios, avec une charge de travail équivalente atteignant 6 478,54 tok/s à BS=64 et un pic à 10 984,61 tok/s à BS=256, soit une légère amélioration par rapport au BF16 en période de forte concurrence. Avec un préremplissage important, le débit commence à 987.48 tok/s à BS=1, grimpe à 14 036,46 tok/s à BS=64 et à 15 156,69 tok/s à BS=128, pour atteindre un pic à 16 658,98 tok/s à BS=256. Le décodage intensif passe de 130.70 tok/s à BS=1 à 4 906,51 tok/s à BS=256, dépassant légèrement BF16 à la concurrence maximale tandis que les deux configurations restent étroitement liées sur le reste de la plage de concurrence.
Mistral Small 3.1 24B Instructions 2503
Sous une charge de travail égale (256/256), le modèle atteint 1 598,79 tok/s à BS=8, 4 713,84 tok/s à BS=64 et grimpe fortement jusqu'à 8 925,12 tok/s à BS=256. Avec un préremplissage important (8k/1k), le débit commence à 897.84 tok/s à BS=1, grimpe à 9 632,58 tok/s à BS=64 et 11 488,13 tok/s à BS=128, pour atteindre un pic de 11 846,15 tok/s à BS=256. Le débit de décodage intensif (1k/8k) passe de 124.98 tok/s à BS=1 à 2 653,82 tok/s à BS=64, puis s'accélère sensiblement à des niveaux de concurrence plus élevés, atteignant 4 262,53 tok/s à BS=128 et culminant à 4 975,06 tok/s à BS=256, reflétant la capacité du modèle à maintenir un débit de décodage élevé à mesure que la concurrence augmente.
Lama 3.1 8B Instruire
Le modèle Llama-3.1-8B-Instruct a été testé sur trois configurations de précision sur le Comino, offrant une vue claire de la façon dont la quantification affecte le débit pour cette taille de modèle.
Llama 3.1 8B Instruction BF16
Sous une charge de travail égale (256/256), le modèle BF16 atteint 2 776,42 tok/s à BS=8, 7 369,01 tok/s à BS=64 et un pic de 11 751,56 tok/s à BS=256. Avec un préremplissage important (8k/1k), il démarre à 1 645,29 tok/s à BS=1, grimpe à 14 990,47 tok/s à BS=64 et 17 140,71 tok/s à BS=128, et culmine à 17 345,80 tok/s à BS=256. Le décodage intensif (1k/8k) passe de 234.78 tok/s à BS=1 à 6 154,73 tok/s à BS=256.
Llama 3.1 8B Instruction FP8
La quantification FP8 offre un gain significatif dans tous les scénarios. À charge de travail égale, le débit atteint 7 530,39 tok/s à BS=64 et culmine à 12 108,98 tok/s à BS=256. Avec un préremplissage intensif, le débit grimpe à 16 546,53 tok/s à BS=64 et à 19 306,49 tok/s à BS=128, avec un pic à 20 137,35 tok/s à BS=256, soit un gain d'environ 16 % par rapport à BF16 en pleine concurrence. Avec un décodage intensif, le débit culmine à 7 353,40 tok/s à BS=256, soit environ 19 % de plus qu'avec BF16.
Llama 3.1 8B Instruction FP4
Le FP4 offre un débit très compétitif par rapport au FP8 à des niveaux de concurrence élevés, bien qu'il soit légèrement en retrait pour les petits lots. À charge de travail égale, le débit maximal atteint 11 954,40 tok/s à BS=256, et avec un préremplissage intensif, il culmine à 20 205,57 tok/s à BS=256, surpassant de peu le FP8 à pleine concurrence. Avec un décodage intensif, le débit maximal atteint 7 239,29 tok/s à BS=256, restant à quelques pourcents seulement du FP8, ce qui fait du FP4 une option intéressante lorsque l'efficacité de la mémoire est primordiale sans sacrifice significatif du débit.
MiniMax M2.5
Le MiniMax-M2.5 230B, testé sur le Comino Grando, était le modèle le plus grand et le plus exigeant que nous ayons utilisé.
Sous une charge de travail égale (256/256), le modèle démarre à 16.35 tok/s à BS=1, atteint 2 751,25 tok/s à BS=64 et connaît une forte montée en puissance avec une concurrence plus élevée, culminant à 5 753,24 tok/s à BS=256. Avec un préremplissage important (8k/1k), le débit démarre à 606.97 tok/s à BS=1, augmente régulièrement jusqu'à 5 351,02 tok/s à BS=32 et 6 557,92 tok/s à BS=64, atteignant son pic à 7 357,26 tok/s à BS=128 avant de diminuer légèrement à 7 140,74 tok/s à BS=256, ce qui suggère que le modèle approche la saturation en termes de débit de préremplissage au-delà de BS=128. Le décodage intensif (1k/8k) augmente de manière constante de 82.21 tok/s à BS=1 à 1 485,28 tok/s à BS=64, atteignant un pic de 2 554,87 tok/s à BS=256, reflétant les besoins attendus en bande passante mémoire d'une architecture MoE 230B sous des charges de travail de décodage soutenues.
Conclusion
Le Comino Grando est avant tout un système conçu pour exploiter pleinement le potentiel de huit GPU NVIDIA RTX PRO 6000. Chaque choix de conception majeur, de l'agencement inversé de la carte mère au système de refroidissement et à la plateforme de surveillance intégrée, vise à garantir un fonctionnement continu de ces GPU à leur TDP maximal de 600 W, sans contraintes thermiques ni énergétiques.
