---
id: collect-250926-servers-hardware/servers-hardware/fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594-3
title: "fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["gpu", "nvidia", "amd", "intel", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594.md
source_anchor: ""
source_lines: [146, 172]
sha256: 1f3db3beb5b04937ba73715f96a110611a97dd153d64f28537f11f260361aaad
---

# fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594

Geekbench 6 est un outil d'évaluation multiplateforme qui mesure les performances globales d'un système. Il propose des tests pour le processeur et la carte graphique. Plus le score est élevé, meilleures sont les performances. Nous nous sommes concentrés ici sur les résultats de la carte graphique.
Vous pouvez trouver des comparaisons avec n'importe quel système dans le navigateur Geekbench.
| Geekbench 6.1.0 (Plus c'est mieux) | Nvidia L4 | Nvidia A2 | NVIDIA T4 | 
|---|---|---|---|
| GeekbenchGPU OpenCL | 156,224 | 35,835 | 83,046 | 
marque de luxe
LuxMark est un outil d'analyse comparative multiplateforme OpenCL conçu par ceux qui maintiennent le moteur de rendu 3D open source LuxRender. Cet outil examine les performances du GPU dans la modélisation 3D, l'éclairage et le travail vidéo. Pour cette revue, nous avons utilisé la version la plus récente, v4alpha0. Dans LuxMark, plus le score est élevé, mieux c'est.
| Luxmark v4.0alpha0 GPU OpenCL (Plus haut, c'est mieux) | Nvidia L4 | Nvidia A2 | NVIDIA T4 | 
|---|---|---|---|
| Banc d'entrée | 14,328 | 3,759 | 5,893 | 
| Banc de nourriture | 5,330 | 1,258 | 2,033 | 
GROMACS CUDA
Nous nous approvisionnons également en GROMACS, un logiciel de dynamique moléculaire, spécifiquement pour CUDA. Cette compilation sur mesure devait exploiter les capacités de traitement parallèle des 5 GPU NVIDIA L4, essentielles pour accélérer les simulations informatiques.
Le processus impliquait l'utilisation de nvcc, le compilateur CUDA de NVIDIA, ainsi que de nombreuses itérations des indicateurs d'optimisation appropriés pour garantir que les binaires étaient correctement adaptés à l'architecture du serveur. L'inclusion du support CUDA dans la compilation GROMACS permet au logiciel de s'interfacer directement avec le matériel GPU, ce qui peut considérablement améliorer les temps de calcul pour les simulations complexes.
Le test : interaction protéique personnalisée dans Gromacs
En tirant parti d'un fichier d'entrée fourni par la communauté à partir de notre divers Discord, qui contenait des paramètres et des structures adaptés à une étude d'interaction protéique spécifique, nous avons lancé une simulation de dynamique moléculaire. Les résultats ont été remarquables : le système a atteint un taux de simulation de 170.268 nanosecondes par jour.
| GPU | Système | ns/jour | temps de base (s) | 
|---|---|---|---|
| Nvidia A4000 | Boîte blanche AMD Ryzen 5950x | 84.415 | 163,763 | 
| RTX NVIDIA 4070 | Boîte blanche AMD Ryzen 7950x3d | 131.85 | 209,692.3 | 
| 5x Nvidia L4 | Dell T560 avec 2x Intel Xeon Gold 6448Y | 170.268 | 608,912.7 | 
Plus que l'IA
Avec le battage médiatique de l'IA qui fait fureur, il est facile de se laisser prendre aux performances des modèles sur NVIDIA L4, mais il a également quelques autres atouts dans son sac, ouvrant un champ de possibilités pour les applications vidéo. Il peut héberger jusqu'à 1,040 1 flux vidéo AV720 simultanés à 30pXNUMX. Cela peut transformer la façon dont le contenu peut être diffusé en direct pour les utilisateurs périphériques, améliorer la narration créative et présenter des utilisations intéressantes pour des expériences AR/VR immersives.
Le NVIDIA L4 excelle également dans l’optimisation des performances graphiques, comme en témoignent ses capacités de rendu en temps réel et de lancer de rayons. Dans un bureau périphérique, le L4 est capable de fournir une accélération de calcul graphique robuste et puissante en VDI aux utilisateurs finaux qui en ont le plus besoin lorsqu'un rendu graphique de haute qualité en temps réel est essentiel.
Réflexions de clôture
Le GPU NVIDIA L4 fournit une plate-forme solide pour l'IA de pointe et le calcul haute performance, offrant une efficacité et une polyvalence inégalées sur plusieurs applications. Sa capacité à gérer des pipelines intensifs d’IA, d’accélération ou de vidéo et à optimiser les performances graphiques en fait un choix idéal pour l’inférence de périphérie ou l’accélération des bureaux virtuels. La combinaison du L4 entre une puissance de calcul élevée, des capacités de mémoire avancées et une efficacité énergétique le positionne comme un acteur clé dans l'accélération des charges de travail à la périphérie, en particulier dans les secteurs de l'IA et des graphiques à forte intensité.
Il ne fait aucun doute que l’IA est l’œil de l’ouragan informatique ces jours-ci, et la demande pour les GPU monstrueux H100/H200 continue d’exploser. Mais il y a également un effort majeur pour mettre en place un ensemble de kits informatiques plus robustes vers la périphérie, où les données sont créées et analysées. Dans ces cas-là, un GPU plus approprié est nécessaire. Ici, le NVIDIA L4 excelle et devrait être l'option par défaut pour l'inférence de bord, soit en tant qu'unité unique, soit à l'échelle globale, comme nous l'avons testé dans le T560.
