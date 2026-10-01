---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5-3
title: "fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5"
domain: servers-hardware
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["gpu", "mlperf", "nvidia"]
source: docs/RAG/clean4/fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5.md
source_anchor: ""
source_lines: [94, 107]
sha256: 5e7e7141665f9bdf872723bb1dd0937cdff9a177635d2d455c778a82d66a0971
---

# fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5

Le XR7620 prend en charge jusqu'à 2 cartes accélératrices de 300 W pour les GPU afin de gérer les charges de travail exigeantes, offrant une classification d'images 45 % plus rapide que le serveur Dell XR 12 avec un seul accélérateur GPU de 300 W. La combinaison d'une faible latence et d'une puissance de traitement élevée permet une analyse des données plus rapide et plus efficace, permettant aux organisations de prendre des décisions en temps réel pour une monétisation accrue.
Pour tester les performances de l'inférence, nous avons exécuté l'inférence MLPerf 3.1, à la fois hors ligne et sur le serveur. BERT (Bidirectionnel Encoder Representations from Transformers) est un modèle basé sur un transformateur principalement utilisé pour les tâches de traitement du langage naturel telles que la réponse aux questions, la compréhension du langage et la classification des phrases. ResNet50 est un modèle de réseau neuronal convolutif (CNN) largement utilisé pour les tâches de classification d'images. Il s’agit d’une variante du modèle ResNet à 50 couches, connue pour son architecture profonde mais ses performances efficaces.
Afin de tester MLPerf 3.1, nous avons équipé notre XR7620 d'un seul GPU NVIDIA L4. La L4 est une carte très performante, conçue spécifiquement pour les charges de travail d’IA et d’apprentissage profond. Avec un énorme 24 Go de VRAM sur une carte mi-hauteur mi-longueur de 70 watts économe en énergie, le L4 est un accélérateur parfait pour ce serveur Edge robuste. Le L4, qui fait partie de la dernière génération de GPU pour centres de données Ada Lovelace de NVIDIA, est conçu pour offrir des performances exceptionnelles dans les tâches d'apprentissage automatique, ce qui en fait un choix idéal pour nos objectifs d'analyse comparative.
| Resnet50 – Hors ligne : | 13,010.2 | 
| Resnet50 – Serveur : | 12,204.4 | 
| Bert K99 – Hors ligne : | 973.465 | 
| Bert K99 – Serveur : | 898.945 | 
- Mode hors ligne : ce mode mesure les performances d'un système lorsque toutes les données sont disponibles pour un traitement simultané. Cela s’apparente au traitement par lots, dans lequel le système traite un grand ensemble de données en un seul lot. Ce mode est crucial pour les scénarios dans lesquels la latence n’est pas une préoccupation majeure, mais le débit et l’efficacité le sont.
- Mode serveur : en revanche, le mode serveur évalue les performances du système dans un scénario qui imite un environnement de serveur réel, dans lequel les requêtes arrivent une par une. Ce mode est sensible à la latence et mesure la rapidité avec laquelle le système peut répondre à chaque demande. C’est crucial pour les applications en temps réel où une réponse immédiate est nécessaire, comme dans les serveurs Web ou les applications interactives.
Notre XR7620, équipé de deux processeurs Xeon Scalable Gold de 4e génération, a parfaitement géré le cache L4, avec des résultats conformes à ceux validés par NVIDIA lors de son test d'inférence MLPerf 3.1 . Nos validations MLPerf 3.1 étant encore en cours, ce test servira de référence pour nos futurs tests.
Conclusion
Le Dell PowerEdge XR7620 est un serveur bien conçu et robuste connu pour sa durabilité et sa fiabilité. Sa conception réfléchie et sa qualité de fabrication en font un choix fiable pour les entreprises et les organisations qui ont besoin d'un calcul haute performance dans des environnements difficiles. Sa conception robuste et résiliente et ses avancées technologiques de pointe le positionnent comme un acteur redoutable dans des environnements difficiles. De plus, les options polyvalentes de connectivité et d’extension du serveur, combinées aux services d’assistance et de garantie fiables de Dell, garantissent qu’il répond non seulement, mais qu’il les dépasse, aux exigences des scénarios informatiques de pointe modernes.
Même si le prix est un facteur à prendre en compte, le Dell PowerEdge XR7620 offre un rapport qualité-prix indéniable, en particulier pour les secteurs exigeant une fiabilité et des performances sans faille dans des conditions extrêmes. Dans l’ensemble, le PowerEdge XR7620 est une solution complète conçue pour relever les complexités et les défis de l’informatique de pointe moderne, s’avérant être un investissement judicieux pour les entreprises tournées vers l’avenir.
Ensuite, nous poussons ce système à l’extrême pour un projet de recherche majeur, d’autres à venir !
