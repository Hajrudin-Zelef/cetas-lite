---
id: collect-261001-cisco/cisco/fr-review-scaling-ai-checkpoints-the-impact-of-high-capacity-ssds-on-model-train-5dfc8c98-2
title: "fr-review-scaling-ai-checkpoints-the-impact-of-high-capacity-ssds-on-model-train-5dfc8c98"
domain: cisco
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["benchmark", "gpu", "intel", "llama", "nvidia", "valuation"]
source: docs/RAG/collect-261001-cisco/fr-review-scaling-ai-checkpoints-the-impact-of-high-capacity-ssds-on-model-train-5dfc8c98.md
source_anchor: ""
source_lines: [16, 43]
sha256: 44ab36d8f0400aa15ec745e8a3c6e06865c92d4558908d19de80e2d6e73829f1
---

# fr-review-scaling-ai-checkpoints-the-impact-of-high-capacity-ssds-on-model-train-5dfc8c98

- 2 x Intel Xeon Gold 6430 (32 cœurs, 2.1 GHz)
- 16 x 64GB DDR5-4400
- Disque SSD Dell BOSS de 480 Go
- Câbles série Gen5 JBOF
  - Disque SSD Solidigm D7.68-PS7 de 1010 To
  - Disque SSD Solidigm D61.44-P5 de 5336 To
Pour garantir que notre analyse comparative reflète des scénarios réels, nous avons basé nos tests sur l'architecture du modèle LLAMA 3.1 405B, en implémentant des points de contrôle via torch.save() pour capturer les paramètres du modèle, les états de l'optimiseur et les états des couches. Notre configuration a simulé un système à 8 GPU, mettant en œuvre une stratégie de parallélisme hybride avec un traitement parallèle à 4 tenseurs et un traitement parallèle à 2 pipelines répartis sur les huit GPU. Cette configuration a donné lieu à des tailles de points de contrôle de 1,636 XNUMX Go, représentatives des exigences de formation des modèles de langage modernes à grande échelle.
Notre processus de test pour la charge de travail du point de contrôle DLIO consistait à remplir chaque lecteur à un niveau d'utilisation similaire. Pour le Solidigm D61.44-P5 de 5336 To, chaque passage comprenait 33 intervalles de point de contrôle, pour un total de 54 To. Le D7.68-PS7 plus petit de 1010 To s'adaptait confortablement à trois intervalles de point de contrôle, avec une empreinte totale de 4.9 To. Un point de contrôle supplémentaire pourrait s'adapter au D7-PS1010, bien qu'il ait augmenté son utilisation légèrement au-dessus de ce que nous souhaitions.
La charge de travail du point de contrôle DLIO a donné des résultats intéressants lorsque nous avons comparé le D4-P61.44 de 5 To basé sur QLC Gen5536 au D5-PS7.68 de 7 To basé sur TLC Gen1010. Au cours du premier passage, à mesure que les disques se remplissaient, nous avons constaté un écart de performances plus important entre les deux modèles de SSD. Le PS5 Gen1010 plus rapide a terminé chaque point de contrôle en moyenne en 464 secondes, contre 623 secondes pour le P4 Gen5336. Lors des deuxième et troisième passages, l'écart s'est réduit à 579 et 587 secondes pour le PS1010 et à 676 et 680 secondes pour le P5336.
Pour les entreprises qui cherchent à réduire au minimum les intervalles de points de contrôle, le Gen5 PS1010 basé sur TLC offre l'avantage d'offrir le temps de traitement le plus rapide. Si l'objectif est de conserver de nombreux points de contrôle de manière rentable, le Gen4 P5336 basé sur QLC peut le faire. Nous avons mesuré une différence de temps de point de contrôle moyen de moins de 17 % entre les deux disques lors des passes deux et trois.
Bande passante de stockage GPUDirect
Bien que DLIO montre les performances du flash dans un flux de travail d'IA, la charge de travail est entièrement basée sur l'écriture jusqu'à ce qu'un point de contrôle soit restauré. Pour brosser un tableau plus complet des Solidigm D7-PS1010 et D5-P5336 dans les charges de travail d'IA, nous avons inclus des mesures de bande passante de lecture à l'aide de GDSIO.
Comment fonctionne le stockage direct GPU
Traditionnellement, lorsqu'un GPU traite des données stockées sur un disque NVMe, les données doivent d'abord transiter par le processeur et la mémoire système avant d'atteindre le GPU. Ce processus introduit des goulots d'étranglement, car le processeur devient un intermédiaire, ce qui ajoute de la latence et consomme de précieuses ressources système. Le stockage direct GPU élimine cette inefficacité en permettant au GPU d'accéder directement aux données depuis le périphérique de stockage via le bus PCIe. Ce chemin direct réduit la surcharge associée au déplacement des données, permettant des transferts de données plus rapides et plus efficaces.
Les charges de travail de l’IA, en particulier celles impliquant l’apprentissage profond, sont très gourmandes en données. La formation de grands réseaux neuronaux nécessite le traitement de téraoctets de données, et tout retard dans le transfert de données peut entraîner une sous-utilisation des GPU et des temps de formation plus longs. Le stockage direct GPU relève ce défi en garantissant que les données sont transmises au GPU le plus rapidement possible, en minimisant les temps d’inactivité et en maximisant l’efficacité de calcul.
Comme pour le test DLIO, l’objectif est de mieux comprendre et caractériser les différences entre les SSD Gen5 à grande vitesse et les disques QLC à grande capacité. Toutes les charges de travail d’IA ne sont pas identiques et chaque disque offre des avantages distincts, en fonction des besoins.
Matrice de configuration des tests
Nous avons testé systématiquement chaque combinaison des paramètres suivants avec un NVIDIA L4 dans notre plateforme de test :
- Tailles de blocs : 1 M, 128 K, 64 K, 16 K, 8 K
- Nombre de fils : 128, 64, 32, 16, 8, 4, 1
- Nombre d'emplois : 16
- Tailles des lots : 16
Notre premier aperçu a été le D5-P5336 basé sur QLC, qui a atteint 4.2 Gio/s avec une taille de transfert de 1 M à une profondeur d'E/S de 128. L'effet des tailles de blocs a produit une augmentation substantielle de la bande passante, passant de 8 1 à 32 M. L'avantage d'une profondeur d'E/S accrue a commencé à diminuer à XNUMX, où les charges de travail ont commencé à se stabiliser.
Ensuite, nous examinons le Gen5 PS-1010, qui peut évoluer jusqu'à 6.2 Gio/s avec une taille de bloc de 1 M et une profondeur d'E/S de 128. Dans l'ensemble, il a surpassé le P4 basé sur Gen5336, avec des charges de travail particulières démontrant une amélioration substantielle. Un domaine d'amélioration notable est venu dans la taille de bloc de 128 K, où avec une profondeur d'E/S de 64 et 128, le PS1010 offrait une bande passante de lecture deux fois supérieure à celle du P5336.
Il est important de noter que les deux SSD ont été testés avec le GPU NVIDIA L4. Bien que le Gen4 D5-P5336 soit un modèle haut de gamme, les GPU NVIDIA plus performants, comme le H100, ont démontré une meilleure efficacité avec le D7-PS1010. La vitesse d'un disque est le critère de choix principal pour certains clients, tandis que d'autres privilégient la densité globale. Solidigm propose des solutions pour les deux, avec ses SSD QLC et TLC.
Conclusion
À mesure que l'ampleur et la complexité de la formation de l'IA continuent de croître, l'infrastructure de stockage sous-jacente doit non seulement suivre le rythme, mais également donner le ton. Nos tests avec deux SSD très différents illustrent l'importance d'aligner les solutions de stockage sur des priorités de formation spécifiques, telles que la minimisation de la latence des points de contrôle ou la maximisation de la densité des points de contrôle pour une évolutivité rentable.
Dans notre évaluation, nous avons testé le Solidigm D5-P5336 (61.44 To) et le D7-PS1010 (7.68 To) dans des conditions d'entraînement d'IA réalistes à l'aide du benchmark DLIO et d'un workflow de point de contrôle LLM hybride-parallèle complet. Nous avons capturé des mesures reflétant les performances d'écriture des points de contrôle sur plusieurs exécutions à mesure que les disques se remplissaient, mettant en évidence les différences de temps d'achèvement entre le D4-P5 basé sur QLC Gen5336 et le D5-PS7 basé sur TLC Gen1010.
