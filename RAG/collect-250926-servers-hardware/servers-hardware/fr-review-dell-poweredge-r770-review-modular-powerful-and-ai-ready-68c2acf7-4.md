---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7-4
title: "fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Intel", "Nvidia", "Samsung", "TensorRT-LLM"]
dates: []
keywords: ["benchmark", "benchmarks", "deepseek", "distillation", "gpu", "inference", "intel", "nvidia", "open source", "qwen", "tensorrt"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7.md
source_anchor: ""
source_lines: [58, 106]
sha256: 21ed7acd3404e8f4e7aa9f90a23c099db841ff91cbe421d6a8d1fefe11948fe1
---

# fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7

En revanche, les performances d'écriture de ces disques sont bien inférieures à leurs performances de lecture. Il a fallu que les 16 disques atteignent une bande passante d'écriture de 46.7 Gio/s, la vitesse d'écriture moyenne restant quasiment constante. Compte tenu des capacités d'écriture inférieures de la gamme KIOXIA CD8, les versions haute capacité ou les autres SSD PCIe Gen5 s'en sortiront mieux.
Analyse comparative du Dell PowerEdge R770
Concernant les benchmarks, le R770 est le système phare de Dell et, à ce titre, sera déployé dans des environnements très variés. Nous avons donc réalisé une série complète de benchmarks pour cette plateforme afin d'évaluer ses performances dans différents environnements. Le Lenovo ThinkSystem SR630 V4 a été comparé lors de certains tests afin de mettre en évidence la différence entre les processeurs haut de gamme à cœurs multiples et à cœurs multiples.
Configuration du système
- CPU: 2x Intel Xeon 6787P (86 cœurs chacun)
- RAM: Mémoire DDR32 double rang 64x Micron 5 Go 6400 MT/s Mémoire totale : 2 To
- Alimentations: 2x Delta 1500W
- GPU: 1x NVIDIA H100 pour le benchmark TGI, 1x NVIDIA L4 pour les tests restants
- FIL: Carte réseau DELL BRCM 4P 25G SFP 57504S OCP
- Carte BOSS : Disques BOSS-N1 DC-MHS 0 et 1 SK hynix 480 Go Dell NVMe ISE PE9010 RI M.2 480 Go
- Disques: 0-5 dans le fond de panier 1 : Samsung 6.4 To, Dell NVMe PM1745 MU E3.S 6.4 To
Performances de la charge de travail de l'IA
Benchmark d'inférence de génération de texte
Text Generation Inference (TGI) est un serveur d'inférence LLM hautes performances développé par Hugging Face. Conçu pour optimiser le déploiement et l'utilisation des LLM, il constitue un choix idéal pour les environnements de production. TGI prend en charge divers LLM open source et offre des fonctionnalités telles que le parallélisme tensoriel, le streaming de jetons et le traitement par lots continu, qui améliorent ses performances et son efficacité.
La fonction d'analyse comparative de TGI permet d'évaluer ses performances sous différentes configurations et charges de travail. Elle offre une représentation plus précise des performances réelles, car elle prend en compte la complexité de la gestion des LLM en environnement de production.
La génération de texte à l'aide de LLM comprend deux étapes principales : le préremplissage et le décodage. Le préremplissage est l'étape initiale, où le LLM traite l'invite de saisie pour générer les représentations intermédiaires nécessaires. Cette étape est gourmande en ressources de calcul, car elle implique le traitement de l'intégralité de l'invite de saisie en un seul passage dans le modèle.
Lors de l'étape de préremplissage, l'invite de saisie est tokenisée et convertie dans un format exploitable par le LLM. Ce dernier calcule ensuite le cache KV, qui stocke les informations relatives aux jetons d'entrée. Ce cache KV est une structure de données essentielle qui facilite la génération de jetons de sortie.
En revanche, l'étape de décodage est un processus autorégressif où le LLM génère les jetons de sortie un par un, en s'appuyant sur les représentations intermédiaires générées lors de l'étape de préremplissage. L'étape de décodage s'appuie fortement sur le cache KV généré lors de l'étape de préremplissage, qui fournit le contexte nécessaire à la génération de jetons de sortie cohérents et contextuellement pertinents.
Étape de pré-remplissage
À mesure que la taille du lot augmente de 1 à 32, la latence des trois modèles augmente ; la latence de DeepSeek-R1-Distill-Qwen-32 B augmente de 29.97 ms pour une taille de lot de 1 à 76.95 ms pour une taille de lot de 32. De même, la latence de GEMMA-3-27B-IT et Qwen/QwQ-32B augmente de 51.84 ms et 29.90 ms à 79.58 ms et 76.30 ms, respectivement.
En revanche, le débit de jetons s'améliore significativement avec l'augmentation de la taille du lot. Pour un lot de 1, les débits des trois modèles varient de 192.95 à 334.46 jetons par seconde. Pour un lot de 32, ils atteignent respectivement 4158.67 4021.40, 4194.13 1 et 32 3 jetons par seconde pour DeepSeek-R27-Distill-Qwen-32B, GEMMA-XNUMX-XNUMXB-IT et Qwen/QwQ-XNUMXB.
| Performances de l'étape de pré-remplissage LLM : latence (ms) et débit de jetons (jetons/s) |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|
| Taille du lot | DeepSeek-R1-Distillation-Qwen-32B |  | GEMMA-3-27B-IT |  | Qwen/QwQ-32B |  | 
|---|---|---|---|---|---|---|
|  | Latence (ms) | Taux de jeton | Latence (ms) | Taux de jeton | Latence (ms) | Taux de jeton | 
| 1 | 29.97 | 333.64 | 51.84 | 192.95 | 29.90 | 334.46 | 
| 2 | 30.21 | 662.09 | 52.55 | 380.61 | 29.95 | 667.80 | 
| 4 | 32.40 | 1234.72 | 52.62 | 760.12 | 32.12 | 1245.47 | 
| 8 | 36.98 | 2163.46 | 52.66 | 1519.19 | 36.69 | 2180.66 | 
| 16 | 51.63 | 3125.50 | 60.96 | 2624.64 | 51.29 | 3147.61 | 
| 32 | 76.95 | 4158.67 | 79.58 | 4021.40 | 76.30 | 4194.13 | 
Étape de décodage
Contrairement à l'étape de pré-remplissage, la latence pendant l'étape de décodage reste relativement stable quelle que soit la taille du lot. Par exemple, la latence de DeepSeek-R1-Distill-Qwen-32 B varie de 27.14 ms à 29.52 ms lorsque la taille du lot augmente de 2 à 32.
Le débit de jetons lors de la phase de décodage s'améliore avec la taille du lot, mais pas de manière aussi spectaculaire que lors de la phase de préremplissage. Pour un lot de 1, le débit est d'environ 36-37 jetons par seconde pour DeepSeek-R1-Distill-Qwen-32B et Qwen/QwQ-32B, et de 33.96 jetons par seconde pour GEMMA-3-27B-IT. Pour un lot de 32, les débits augmentent respectivement à 1083.83 873.39, 1084.89 et XNUMX XNUMX jetons par seconde.
| Performances de décodage LLM (jeton) : latence (ms) et débit de jetons (jetons/s) |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|
| Taille du lot | DeepSeek-R1-Distillation-Qwen-32B |  | GEMMA-3-27B-IT |  | Qwen/QwQ-32B |  | 
|---|---|---|---|---|---|---|
|  | Latence (ms) | Taux de jeton | Latence (ms) | Taux de jeton | Latence (ms) | Taux de jeton | 
| 1 | 27.24 | 36.71 | 29.45 | 33.96 | 27.24 | 36.71 | 
| 2 | 27.14 | 73.70 | 30.80 | 64.93 | 27.14 | 73.69 | 
| 4 | 27.50 | 145.46 | 31.33 | 127.65 | 27.47 | 145.62 | 
| 8 | 27.91 | 286.61 | 32.54 | 245.83 | 27.90 | 286.78 | 
| 16 | 28.31 | 565.07 | 34.71 | 460.92 | 28.44 | 562.56 | 
| 32 | 29.52 | 1083.83 | 36.64 | 873.39 | 29.50 | 1084.89 | 
Ceci est normal, car l'étape de préremplissage calcule les états cachés initiaux et les caches clé-valeur pour l'intégralité de l'invite de saisie, ce qui peut saturer le GPU, car de grandes opérations par lots peuvent être exécutées simultanément. Après le traitement de l'invite, le modèle génère de nouveaux jetons, généralement un par un. À chaque étape, le modèle utilise le jeton précédent et les états cachés mis en cache pour produire le jeton suivant. Comme cette étape procède jeton par jeton, la taille du lot est souvent réduite, ce qui entraîne une sous-utilisation fréquente du GPU.
Benchmark de vision par ordinateur Procyon AI
À l'aide de tâches de vision artificielle concrètes, le benchmark Procyon AI Computer Vision évalue les performances d'inférence IA sur les processeurs, les GPU et les accélérateurs IA. Il prend en charge plusieurs moteurs d'inférence tels que TensorRT, OpenVINO, SNPE, Windows ML et Core ML, offrant ainsi des informations sur l'efficacité, la compatibilité et l'optimisation.
