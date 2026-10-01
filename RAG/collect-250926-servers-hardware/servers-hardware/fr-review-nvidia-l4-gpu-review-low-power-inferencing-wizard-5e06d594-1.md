---
id: collect-250926-servers-hardware/servers-hardware/fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594-1
title: "fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["gpu", "nvidia", "benchmarks", "fp8", "intel", "mlperf", "valuation"]
source: docs/RAG/clean4/fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594.md
source_anchor: ""
source_lines: [1, 74]
sha256: 4a3fb1424a6d9aaddf10968483740cdf615badc53a81c83aef283b184c9d4187
---

# fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594

Dans le torrent incessant d’innovation du monde de l’IA d’aujourd’hui, il est essentiel de mesurer et de comprendre les capacités des diverses plates-formes matérielles. Toutes les IA ne nécessitent pas d'énormes fermes de GPU de formation, il existe un segment important de l'IA d'inférence, qui nécessite souvent moins de puissance GPU, en particulier à la périphérie. Dans cette revue, nous examinons plusieurs GPU NVIDIA L4, sur trois serveurs Dell différents et diverses charges de travail, y compris MLperf, pour voir comment le L4 se compare.
GPU NVIDIA L4
À la base, le L4 offre une impressionnante capacité de 30.3 téraFLOP en performances FP32, idéale pour les tâches informatiques de haute précision. Ses prouesses s'étendent aux calculs de précision mixte avec les cœurs Tensor TF32, FP16 et BFLOAT16, cruciaux pour l'efficacité de l'apprentissage en profondeur, la fiche technique L4 cite des performances comprises entre 60 et 121 téraFLOP.
Dans les tâches de faible précision, le L4 brille avec 242.5 téraFLOP dans les cœurs Tensor FP8 et INT8, améliorant ainsi l'inférence du réseau neuronal. Sa mémoire GDDR24 de 6 Go, complétée par une bande passante de 300 Go/s, le rend capable de gérer de grands ensembles de données et des modèles complexes. L'efficacité énergétique du L4 est ce qui est le plus remarquable ici, avec un TDP de 72 W le rendant adapté à divers environnements informatiques. Ce mélange de hautes performances, d'efficacité de la mémoire et de faible consommation d'énergie fait du NVIDIA L4 un choix incontournable pour les défis informatiques de pointe.
| Spécifications NVIDIA L4 |  | 
|---|---|
| FP 32 | 30.3 téraFLOP | 
| Noyau tenseur TF32 | 60 téraFLOP | 
| Noyau tenseur FP16 | 121 téraFLOP | 
| Noyau tenseur BFLOAT16 | 121 téraFLOP | 
| Noyau tenseur FP8 | 242.5 téraFLOP | 
| Noyau tenseur INT8 | TOP 242.5 | 
| Mémoire GPU | 24GB GDDR6 | 
| Bande passante mémoire GPU | 300GB / s | 
| Puissance thermique maximale (TDP) | 72W | 
| Facteur de forme | PCIe profil bas à 1 emplacement | 
| Interconnect | PCIe Gen4x16 | 
| Tableau des spécifications | L4 | 
Bien sûr, avec le prix du L4 proche de 2500 2 $, l'A4 coûtant environ la moitié du prix et le T1000 vieilli (mais toujours assez performant) disponible pour moins de XNUMX XNUMX $, la question évidente est de savoir quelle est la différence entre ces trois GPU d'inférence.
| Spécifications NVIDIA L4, A2 et T4 | Nvidia L4 | Nvidia A2 | NVIDIA T4 | 
|---|---|---|---|
| FP 32 | 30.3 téraFLOP | 4.5 téraFLOP | 8.1 téraFLOP | 
| Noyau tenseur TF32 | 60 téraFLOP | 9 téraFLOP | N/D | 
| Noyau tenseur FP16 | 121 téraFLOP | 18 téraFLOP | N/D | 
| Noyau tenseur BFLOAT16 | 121 téraFLOP | 18 téraFLOP | N/D | 
| Noyau tenseur FP8 | 242.5 téraFLOP | N/D | N/D | 
| Noyau tenseur INT8 | TOP 242.5 | 36 TOPS | 130 TOPS | 
| Mémoire GPU | 24GB GDDR6 | 16GB GDDR6 | 16GB GDDR6 | 
| Bande passante mémoire GPU | 300GB / s | 200GB / s | 320+ Go/s | 
| Puissance thermique maximale (TDP) | 72W | 40-60W | 70W | 
| Facteur de forme | PCIe profil bas à 1 emplacement |  |  | 
| Interconnect | PCIe Gen4x16 | PCIe Gen4x8 | PCIe Gen3x16 | 
| Tableau des spécifications | L4 | A2 | T4 | 
Une chose à comprendre en regardant ces trois cartes est qu'elles ne sont pas exactement des remplacements générationnels individuels, ce qui explique pourquoi le T4 reste encore, de nombreuses années plus tard, un choix populaire pour certains cas d'utilisation. L'A2 est venu remplacer le T4 en tant qu'option à faible consommation et plus compatible (mécanique x8 vs x16). Techniquement, le L4 remplace alors le T4, l'A2 étant à cheval sur un intermédiaire qui pourrait ou non être actualisé à un moment donné dans le futur.
Performances de l'inférence MLPerf 3.1
MLPerf est un consortium de leaders de l'IA issus du monde universitaire, de la recherche et de l'industrie, créé pour fournir des références matérielles et logicielles d'IA justes et pertinentes. Ces benchmarks sont conçus pour mesurer les performances du matériel, des logiciels et des services d'apprentissage automatique sur diverses tâches et scénarios.
Nos tests se concentrent sur deux benchmarks MLPerf spécifiques : Resnet50 et BERT.
- Resnet50 : Il s'agit d'un réseau neuronal convolutif utilisé principalement pour la classification d'images. C'est un bon indicateur de la capacité d'un système à gérer les tâches d'apprentissage en profondeur liées au traitement d'images.
- BERT (Bidirectionnel Encoder Representations from Transformers) : cette référence se concentre sur les tâches de traitement du langage naturel, offrant un aperçu de la façon dont un système fonctionne dans la compréhension et le traitement du langage humain.
Ces deux tests sont cruciaux pour évaluer les capacités du matériel d’IA dans des scénarios réels impliquant le traitement d’images et de langage.
L'évaluation de NVIDIA L4 avec ces benchmarks est essentielle pour aider à comprendre les capacités du GPU L4 dans des tâches d'IA spécifiques. Il offre également un aperçu de la façon dont différentes configurations (configurations simples, doubles et quadruples) influencent les performances. Ces informations sont vitales pour les professionnels et les organisations qui cherchent à optimiser leur infrastructure d’IA.
Les modèles fonctionnent sous deux modes clés : serveur et hors ligne.
- Mode hors ligne : ce mode mesure les performances d'un système lorsque toutes les données sont disponibles pour un traitement simultané. Cela s'apparente au traitement par lots, dans lequel le système traite un grand ensemble de données en un seul lot. Le mode hors ligne est crucial pour les scénarios dans lesquels la latence n’est pas une préoccupation majeure, mais le débit et l’efficacité le sont.
- Mode serveur : en revanche, le mode serveur évalue les performances du système dans un scénario imitant un environnement de serveur réel, dans lequel les requêtes arrivent une par une. Ce mode est sensible à la latence et mesure la rapidité avec laquelle le système peut répondre à chaque demande. C'est essentiel pour les applications en temps réel, telles que les serveurs Web ou les applications interactives, pour lesquelles une réponse immédiate est nécessaire.
1 x NVIDIA L4 – Dell PowerEdge XR7620
Dans le cadre de notre récent test du Dell PowerEdge XR7620 , équipé d'un seul NVIDIA L4, nous l'avons poussé dans ses retranchements pour exécuter plusieurs tâches, dont MLPerf.
La configuration de notre système de test comprenait les composants suivants :
- 2 Xeon Gold 6426Y – 16 cœurs 2.5 GHz
- 1 x Nvidia L4
- 8 x 16GB DDR5
- BOSS RAID480 de 1 Go
- Ubuntu Server 22.04
- Pilote NVIDIA 535
| Dell PowerEdge XR7620 1x NVIDIA L4 | Score | 
|---|---|
| Resnet50 – Serveur | 12,204.40 | 
| Resnet50 – Hors ligne | 13,010.20 | 
| BERT K99 – Serveur | 898.945 | 
| BERT K99 – Hors ligne | 973.435 | 
Les performances dans les scénarios de serveur et hors ligne pour Resnet50 et BERT K99 sont presque identiques, ce qui indique que le L4 maintient des performances cohérentes sur différents modèles de serveur.
1, 2 et 4 NVIDIA L4 – Dell PowerEdge T560
La configuration de notre unité d'examen comprenait les composants suivants :
- 2 x Intel Xeon Gold 6448Y (32 cœurs/64 threads chacun, TDP 225 watts, 2.1-4.1 GHz)
- 8 disques SSD Solidigm P1.6 de 5520 To avec carte RAID PERC 12
- 1 à 4x GPU NVIDIA L4
- 8 modules RDIMM de 64 Go
- Ubuntu Server 22.04
- Pilote NVIDIA 535
| Dell PowerEdge T560 1x NVIDIA L4 | Score | 
|---|---|
| Resnet50 – Serveur | 12,204.40 | 
| Resnet50 – Hors ligne | 12,872.10 | 
| Bert K99 – Serveur | 898.945 | 
| Bert K99 – Hors ligne | 945.146 | 
