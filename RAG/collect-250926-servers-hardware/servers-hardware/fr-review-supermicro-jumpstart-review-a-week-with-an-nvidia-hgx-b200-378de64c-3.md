---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c-3
title: "fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Meta", "Nvidia", "vLLM"]
dates: []
keywords: ["nvidia", "benchmark", "benchmarks", "deepseek", "diffusion", "fp4", "fp8", "llama", "mixture of experts", "moe", "nvfp4", "valuation"]
source: docs/RAG/clean4/fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c.md
source_anchor: ""
source_lines: [42, 64]
sha256: 3eeba54d2734d0adad47e600b3a00fc73e45d81916e2f00e7b95599b08a628f1
---

# fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c

La latence d'écriture aléatoire était initialement relativement faible, avec des performances monocœur allant d'environ 0.6 à 1.2 ms pour les petits blocs et de 5 à 12 ms pour les transferts les plus importants. À mesure que la concurrence augmentait, la latence augmentait rapidement, atteignant 4 à 9 ms avec huit threads et 18 à 38 ms avec 32 threads.
Au-delà de ce seuil, le chemin d'écriture s'est saturé. Avec des blocs plus volumineux, la latence a atteint 150 à 380 ms à 64 threads, et une augmentation continue de la taille des blocs a fait grimper la latence de façon spectaculaire. Le cas le plus défavorable a été observé avec des blocs de 10 Mo et 256 threads, avec une latence maximale d'environ 4.4 secondes.
Service en ligne vLLM – Performances d'inférence LLM
vLLM est le moteur d'inférence et de diffusion à haut débit le plus populaire pour les LLM. Le benchmark de diffusion en ligne vLLM est un outil d'évaluation des performances qui mesure les capacités de diffusion réelles de ce moteur d'inférence sous des requêtes simultanées. Il simule les charges de travail de production en envoyant des requêtes à un serveur vLLM en cours d'exécution avec des paramètres configurables, tels que le débit de requêtes, la longueur des entrées/sorties et le nombre de clients simultanés. Le benchmark mesure des indicateurs clés, notamment le débit (jetons par seconde), le temps d'obtention du premier jeton et le temps par jeton de sortie (TPOT), permettant ainsi aux utilisateurs de comprendre les performances de vLLM sous différentes conditions de charge.
Nous avons testé les performances d'inférence sur une suite complète de modèles couvrant diverses architectures, échelles de paramètres et stratégies de quantification afin d'évaluer le débit sous différents profils de concurrence.
Performances du modèle dense
Les modèles denses suivent l'architecture LLM classique, où tous les paramètres et activations sont utilisés lors de l'inférence, ce qui engendre un traitement plus intensif en ressources de calcul que leurs homologues clairsemés. Afin d'évaluer de manière exhaustive les performances des différents modèles, en fonction de leur taille et des stratégies de quantification, nous avons comparé plusieurs configurations de modèles denses de la famille Llama 3.1 8B.
Notre suite de tests comprenait des évaluations de Meta Llama 3.1 8B sur trois formats de précision : la configuration standard, ainsi que les versions quantifiées FP8 et FP4 utilisant le format NVFP4 de NVIDIA. Il est important de noter que vLLM utilise actuellement le noyau Marlin pour les modèles quantifiés NVFP4, et que les gains de performance optimaux de ce format de quantification ne sont pas encore pleinement visibles dans ces benchmarks. De futures optimisations de vLLM ciblant les opérations natives du cœur tensoriel NVFP4 pourraient apporter des améliorations de performance supplémentaires. Cette stratégie de sélection de modèles permet une comparaison directe des performances tout en isolant l'impact de la quantification progressive sur le débit d'inférence.
Llama 3.1 8B Performance
Le modèle Llama 3.1 8B, en précision standard, présente les caractéristiques de mise à l'échelle suivantes en fonction du niveau de concurrence. En mode mono-utilisateur (BS=1), il atteint 279.27 tok/s par utilisateur, pour un débit total de 1 727,62 tok/s et un TPOT de 3.37 ms. À mesure que la taille des lots augmente, le débit par utilisateur diminue tandis que le débit total augmente. Avec BS=8, le modèle atteint 82.85 tok/s par utilisateur, pour un débit total de 3 386,48 tok/s et un TPOT de 3.28 ms. Les performances continuent de progresser jusqu'à BS=32 (56.46 tok/s par utilisateur, 8 274,66 tok/s au total) et BS=64 (52.70 tok/s par utilisateur, 13 707,66 tok/s au total).
Le modèle atteint son débit total maximal à la station de base 256 (BS=256), avec 32 797,67 tok/s, soit 30.64 tok/s par utilisateur et un TPOT de 16.13 ms. Cela représente une augmentation de 19 fois du débit total par rapport aux performances d'un utilisateur unique. Les valeurs de TPOT restent comprises entre 3 et 4 ms jusqu'à la station de base 64, puis augmentent jusqu'à 16.13 ms à la station de base 256.
Performances du Llama 3.1 8B FP8
La variante quantifiée FP8 présente des caractéristiques différentes. À BS=1, elle atteint 149.46 tok/s par utilisateur, avec un débit total de 9 565,20 tok/s et un TPOT de 3.44 ms. L’analyse de Pareto révèle trois points optimaux (BS=1, BS=128, BS=256).
À BS=128, le modèle fournit 44.12 tok/s par utilisateur avec un total de 19 198,40 tok/s et un temps de réponse de 11.04 ms.
Le débit total maximal est atteint à BS=256 avec un total de 29 219,67 tok/s, soit 30.13 tok/s par utilisateur, et un TPOT de 13.26 ms. La variante FP8 atteint un débit total maximal inférieur à celui de la précision standard (29.2 K contre 32.8 K tok/s).
Performances du Llama 3.1 8B FP4
La configuration quantifiée FP4 affiche les résultats suivants : en mode mono-utilisateur (BS=1), elle atteint 279.73 tok/s par utilisateur, un débit total de 830.46 tok/s et un TPOT de 3.43 ms.
Le modèle FP4 présente sept points de frontière de Pareto. À la station de base 2 (BS=2), il offre un débit de 159.95 tok/s par utilisateur, soit un débit total de 928.60 tok/s, et un TPOT de 3.43 ms. Les performances continuent de progresser jusqu'à BS=4 (76.36 tok/s par utilisateur, soit un débit total de 1 631,69 tok/s) et BS=32 (76.09 tok/s par utilisateur, soit un débit total de 9 014,29 tok/s). Le modèle atteint son débit total maximal à BS=256, avec 29 340,89 tok/s, 30.13 tok/s par utilisateur et un TPOT de 16.18 ms.
Performances du modèle parcimonieux
Les modèles épars, notamment les architectures Mixture of Experts (MoE), constituent une approche émergente pour la mise à l'échelle efficace des modèles de langage. Ces architectures conservent un nombre total de paramètres élevé tout en n'activant qu'un sous-ensemble de paramètres par jeton, ce qui peut potentiellement améliorer les performances par paramètre actif.
Nous avons évalué deux architectures MoE : DeepSeek-R1, un modèle axé sur le raisonnement, et Qwen3 Coder 30B-A3B, une architecture clairsemée spécialisée dans la génération de code. DeepSeek-R1 est le modèle axé sur le raisonnement le plus répandu, présentant des performances nettement supérieures aux modèles de langage traditionnels. Le modèle Qwen3 Coder conserve l’intégralité de ses 30 milliards de paramètres tout en n’en activant que 3 milliards par jeton généré. Nous avons comparé les performances de Qwen3 Coder avec celles de ses variantes de quantification par défaut et par virgule flottante FP8 afin d’analyser les variations selon la stratégie de quantification.
Performances de DeepSeek-R1
Le modèle DeepSeek-R1 présente un comportement de mise à l'échelle intéressant en fonction de la taille des lots. En mode mono-utilisateur (BS=1), il atteint 30.24 tok/s par utilisateur, un débit total de 88.13 tok/s et un TPOT de 29.85 ms. Avec BS=4, les performances atteignent 29.77 tok/s par utilisateur et un débit total de 266.40 tok/s, pour un TPOT de 32.04 ms, soit le débit total maximal obtenu sur l'ensemble des configurations.
