---
id: collect-261001-ia-llm/ia-llm/fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6-3
title: "fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "gpu", "nvidia"]
source: docs/RAG/collect-261001-ia-llm/fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6.md
source_anchor: ""
source_lines: [51, 112]
sha256: e4fefedd7de05222ee31bce0250294507dbf4db69076007f52ee652f3aa8f82d
---

# fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6

- Stockage->PAGE_CACHE->CPU->GPU : Ce chemin utilise le cache de pages pour les transferts de données, où les données sont d'abord mises en cache en mémoire avant d'être traitées par le CPU puis transférées vers le GPU. Cette configuration est utile pour tester l'impact des mécanismes de mise en cache et de la bande passante mémoire sur les performances globales, ce qui est pertinent lors de la formation lorsque les données peuvent être prétraitées et mises en cache pour plus d'efficacité. Encore une fois, nous pourrions affirmer que ce chemin de données représenterait les performances quel que soit le fournisseur de GPU.
Imiter les étapes du pipeline IA
Les tests de référence sont conçus pour refléter les différentes étapes du pipeline d'IA, garantissant que les mesures de performance obtenues sont pertinentes et complètes.
Préparation des données:
- Tailles d'E/S : Plus petit (128K, 256K, 512K)
- Sujets: 1, 4
- Types de transfert : « Stockage->CPU->GPU », « Stockage->PAGE_CACHE->CPU->GPU »
- Objectif : Évaluez la manière dont les disques SSD gèrent les petits transferts de données fréquents et l'implication du processeur, essentiels lors des phases d'ingestion, de nettoyage et d'augmentation des données.
Formation et mise au point :
- Tailles d'E/S : Moyen à grand (1M, 4M, 16M)
- Sujets: 4, 16, 32
- Types de transfert : « Stockage->GPU (GDS) », « Stockage->CPU->GPU »
- Objectif : Évaluez les performances dans des conditions de débit de données élevé avec plusieurs flux de données simultanés, ce qui représente la gestion intensive des données requise lors de la formation et du réglage précis du modèle.
Inférence :
- Tailles d'E/S : Grand à très grand (16 M, 64 M, 128 M) et 4K
- Sujets: 1, 4, 16
- Types de transfert : Stockage->GPU (GDS)
- Objectif : Mesurez l'efficacité des transferts de données directs à grande échelle vers le GPU, essentiels pour les applications d'inférence en temps réel où un accès rapide aux données et une latence minimale sont primordiaux. 4K est conçu pour examiner les recherches effectuées dans la base de données RAG.
En faisant varier ces paramètres et en testant différentes configurations, nous pouvons obtenir un profil de performances détaillé des SSD Solidigm 61.44 To QLC dans un environnement de serveur d'IA hautes performances, fournissant ainsi un aperçu de leur adéquation et de leur optimisation pour diverses charges de travail d'IA. Nous avons examiné les données en effectuant plus de 1200 XNUMX tests sur quelques semaines.
Configuration du serveur
- Lenovo Think System SR675 V3
- Processeur AMD EPYC 9254 24 cœurs
- 6 X 64 Go DDR5, capacité totale de 384 Go
- 4X GPU NVIDIA L40S
- 4 disques SSD NVMe Solidigm P61.44 QLC de 5336 To
- Ubuntu Server 22.04
- Version du pilote NVIDIA : 535.171.04
- Version CUDA : 12.2
Résultats de référence
Tout d’abord, examinons les charges de travail de type formation et inférence. La taille d'E/S GPU Direct 1024K représente le chargement du modèle, les données d'entraînement chargées sur le GPU et d'autres tâches d'inférence par lots volumineuses comme dans le travail d'image ou de vidéo.
| 4Conduire | Type d'E / S | Type de transfert | Threads | Taille de l'ensemble de données (Kio) | Taille des E/S (Ko) | Débit (Gio/sec) | Latence moyenne (usecs) | 
|---|---|---|---|---|---|---|---|
|  | ÉCRIRE | GPUD | 8 | 777,375,744 | 1024 | 12.31 | 634.55 | 
|  | LIS | GPUD | 8 | 579,439,616 | 1024 | 9.30 | 840.37 | 
|  | RANDÉCRITURE | GPUD | 8 | 751,927,296 | 1024 | 12.04 | 648.67 | 
|  | RANDIRER | GPUD | 8 | 653,832,192 | 1024 | 10.50 | 743.89 | 
Ensuite, examinez les tailles d'E/S plus petites, pour une charge de travail de type RAG, par exemple où un accès rapide et aléatoire aux données 4k à une base de données RAG stockée sur disque. Des E/S aléatoires efficaces sont nécessaires pour les scénarios dans lesquels les charges de travail d'inférence doivent accéder aux données de manière non séquentielle, comme avec des systèmes de recommandation ou des applications de recherche. La configuration RAID0 présente de bonnes performances pour les opérations séquentielles et aléatoires, ce qui est crucial pour les applications d'IA qui impliquent un mélange de modèles d'accès comme RAG. Les valeurs de latence de lecture sont particulièrement faibles, notamment dans les GPUD .
8 threads de travail ont été sélectionnés ici, qui ne saturent pas complètement le SSD, mais fournissent un instantané plus représentatif de ce que vous pouvez trouver dans une charge de travail de type RAG. Cela fournit un contexte d'application prête à l'emploi autour du point de vue du GPU avec un nombre limité de tâches travaillées et une profondeur de file d'attente plus élevée, ce qui mérite de noter que cela montre qu'il reste plus de performances sur la table qui peuvent être obtenues grâce à d'autres optimisations logicielles. .
| 4Conduire | Type d'E / S | Type de transfert | Threads | Taille de l'ensemble de données (Kio) | Taille des E/S (Ko) | Débit (Gio/sec) | Latence moyenne (usecs) | 
|---|---|---|---|---|---|---|---|
|  | ÉCRIRE | GPUD | 8 | 69,929,336 | 4 | 1.12 | 27.32 | 
|  | LIS | GPUD | 8 | 37,096,856 | 4 | 0.59 | 51.52 | 
|  | RANDÉCRITURE | GPUD | 8 | 57,083,336 | 4 | 0.91 | 33.42 | 
|  | RANDIRER | GPUD | 8 | 27,226,364 | 4 | 0.44 | 70.07 | 
Si vous n'utilisez pas GPU Direct en raison de bibliothèques ou de GPU non pris en charge, voici ces deux types si vous utilisez le CPU dans le transfert de données. Dans ce serveur spécifique, le Lenovo ThinkSystem SR675 V3, étant donné que tous les périphériques PCIe passent par le complexe racine du processeur, nous constatons une bande passante comparable mais prenons un coup sur notre latence. On peut s'attendre à une amélioration d'un système avec des commutateurs PCIe.
| 4Conduire | Type d'E / S | Type de transfert | Threads | Taille de l'ensemble de données (Kio) | Taille des E/S (Ko) | Débit (Gio/sec) | Latence moyenne (usecs) | 
|---|---|---|---|---|---|---|---|
|  | ÉCRIRE | CPU_GPU | 8 | 767,126,528 | 1024 | 12.24 | 638.05 | 
|  | LIS | CPU_GPU | 8 | 660,889,600 | 1024 | 10.58 | 738.75 | 
|  | RANDÉCRITURE | CPU_GPU | 8 | 752,763,904 | 1024 | 12.02 | 649.76 | 
|  | RANDIRER | CPU_GPU | 8 | 656,329,728 | 1024 | 10.47 | 746.26 | 
|  | ÉCRIRE | CPU_GPU | 8 | 69,498,220 | 4 | 1.11 | 27.47 | 
|  | LIS | CPU_GPU | 8 | 36,634,680 | 4 | 0.58 | 52.31 | 
Le tableau indique des débits élevés pour les opérations de lecture, en particulier avec le GPUD type de transfert. Par exemple, lisez les opérations dans GPUD Le mode atteint plus de 10.5 Gio/sec. Cela profite aux charges de travail d’IA, qui nécessitent souvent un accès rapide aux données pour entraîner de grands modèles.
Les performances équilibrées entre les opérations aléatoires et séquentielles rendent cette configuration adaptée aux tâches d'inférence, qui nécessitent souvent un mélange de ces modèles d'accès. Même si les valeurs de latence ne sont pas extrêmement faibles, elles restent dans des limites acceptables pour de nombreuses applications d'inférence.
De plus, nous constatons des débits impressionnants, avec des opérations d’écriture atteignant jusqu’à 12.31 Gio/s et des opérations de lecture jusqu’à 9.30 Gio/s. Ce débit élevé profite aux charges de travail d’IA qui nécessitent un accès rapide aux données pour la formation et l’inférence des modèles.
Lectures séquentielles et optimisation
En passant à une taille d'E/S de 128 M et en parcourant les threads de travail, nous pouvons voir le résultat de l'optimisation d'une charge de travail pour une solution de stockage.
| Type de transfert | Threads | Débit (Gio/s) | Latence (usec) | 
|---|---|---|---|
| Stockage->CPU->GPU | 16 | 25.134916 | 79528.88255 | 
| Stockage->CPU->GPU | 4 | 25.134903 | 19887.66948 | 
