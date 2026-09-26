---
id: collect-260926-rattrapage/rattrapage/fr-review-lenovo-thinksystem-sr650-v4-review-versatile-2u-workhorse-b60ae752-3
title: "Test du Lenovo ThinkSystem SR650 V4 : Un serveur 2U polyvalent et performant"
domain: rattrapage
role: reference
task: reference
actors: ["Intel", "Nvidia", "vLLM"]
dates: []
keywords: ["benchmark", "benchmarks", "diffusion", "fp8", "gpu", "intel", "llama", "nvidia", "open source", "valuation", "vllm"]
source: docs/RAG/lot-rattrapage/servers-reviews/fr-review-lenovo-thinksystem-sr650-v4-review-versatile-2u-workhorse-b60ae752.md
source_anchor: ""
source_lines: [123, 185]
sha256: 76ffa7f685805fdb21bb3ad746e884a5eb1b1afa7a420727d040a91094effb74
---

# Test du Lenovo ThinkSystem SR650 V4 : Un serveur 2U polyvalent et performant

Cette section examine les résultats des tests de performances réalisés avec Blender, y-cruncher, vLLM et Phoronix. Les tests initiaux ont été effectués sur un Lenovo ThinkSystem SR650 V4 équipé de processeurs Intel Xeon 6740P, ce qui nous a permis d'évaluer l'architecture P-core axée sur les performances et le débit global de la plateforme sous des charges de travail gourmandes en ressources CPU. En cours de test, nous avons modifié la configuration du système pour activer l'accélération GPU. La transition s'est faite sans difficulté grâce au kit de mise à niveau GPU de Lenovo, qui comprend des ventilateurs hautes performances, des dissipateurs thermiques améliorés, un carénage de flux d'air repensé et des risers compatibles GPU conçus pour supporter des accélérateurs plus puissants. L'installation n'a nécessité aucune modification importante du châssis et s'est parfaitement intégrée à la configuration système existante.

Grâce à l'installation des composants de refroidissement et de ventilation améliorés, nous avons intégré un GPU NVIDIA L40S à la plateforme. Ceci nous a permis d'étendre nos tests aux charges de travail accélérées par GPU, telles que le rendu GPU sous Blender et les benchmarks axés sur l'inférence, tout en garantissant une température stable et un comportement système constant. La modularité de cette mise à niveau souligne la flexibilité du SR650 V4, lui permettant de passer facilement d'une configuration CPU axée sur l'efficacité à une plateforme CPU+GPU équilibrée, adaptée aux charges de travail de calcul mixtes.

### Configuration du Lenovo ThinkSystem SR650 V4 :

- **CPU:** 2 processeurs Intel Xeon 6740P
- **Mémoire:** 1 To de RAM
- **Stockage:** 4 SSD de 960 Go
- **GPU:** Nvidia L40S

## Blender 4.5

Blender est une application de modélisation 3D open source. Ce benchmark a été réalisé avec l'utilitaire Blender Benchmark. Le score est mesuré en échantillons par minute, les valeurs les plus élevées indiquant de meilleures performances.

Dans le test de performances CPU de Blender, le Lenovo ThinkSystem SR650 V4 offre d'excellentes performances de rendu sur toutes les scènes, avec un score de 1136.99 échantillons par minute pour Monster, 707.60 pour Junkshop et 562.53 pour Classroom. Ces résultats témoignent d'une forte capacité de traitement multithread grâce aux deux processeurs Intel Xeon 6740P, garantissant des performances constantes malgré la complexité croissante des scènes.

| Processeur de Blender (Échantillons par minute ; plus c'est élevé, mieux c'est) | Lenovo ThinkSystem SR650 V4 (2x Intel Xeon 6740P, 1 To de RAM) |
|---|---|
| Monster | 1136.99 |
| Junkshop | 707.60 |
| Classroom | 562.53 |

Dans le benchmark Blender CPU sans SMT, le SR650 V4 affiche une nette amélioration du débit de rendu, atteignant 4 686,40 échantillons par minute dans Monster, 2 223,96 dans Junkshop et 2 385,98 dans Classroom. Les résultats GPU soulignent la capacité du système à accélérer les charges de travail de rendu complexes, offrant un débit considérablement supérieur au rendu CPU seul.

| Processeur Blender (sans SMT) (Échantillons par minute ; plus c'est élevé, mieux c'est) | Lenovo ThinkSystem SR650 V4 (2x Intel Xeon 6740P, 1 To de RAM) |
|---|---|
| Monster | 4686.40 |
| Junkshop | 2223.96 |
| Classroom | 2385.98 |

## y-cruncher

y-cruncher est un programme multithread et évolutif qui calcule Pi et d'autres constantes mathématiques avec une précision de plusieurs billions de décimales. Depuis son lancement en 2009, il est devenu une application populaire de test de performance et de résistance pour les overclockeurs et les passionnés de matériel informatique.

Dans y-cruncher, le Lenovo affiche une montée en charge fluide face à l'augmentation de la taille des problèmes, soulignant ainsi les excellentes performances de calcul multithread et la bande passante mémoire de la plateforme. Avec deux processeurs Intel Xeon 6740P et 1 To de RAM, le système effectue le calcul de Pi à 1 milliard de décimales en un peu plus de 20 secondes et conserve une efficacité constante jusqu'à 25 milliards de décimales, en 362 secondes. Ces résultats démontrent une montée en charge prévisible sous une pression soutenue sur le processeur et la mémoire, faisant du SR650 V4 un choix idéal pour les charges de travail mathématiques importantes et les tests de résistance.

| Y-Cruncher (temps de calcul total) | Lenovo ThinkSystem SR650 V4 (2x Intel Xeon 6740P, 1 To de RAM) |
|---|---|
| 1 milliard | 20.715 s |
| 2.5 milliard | 44.412 s |
| 5 milliard | 81.937 s |
| 10 milliard | 152.743 s |
| 25 milliard | 362.566 s |

## Performances d'inférence LLM du service en ligne vLLM

vLLM est le moteur d'inférence et de diffusion à haut débit le plus populaire pour les LLM. Le benchmark de diffusion en ligne vLLM est un outil d'évaluation des performances qui mesure les capacités de diffusion réelles de ce moteur d'inférence sous des requêtes simultanées. Il simule les charges de travail de production en envoyant des requêtes à un serveur vLLM en cours d'exécution avec des paramètres configurables, tels que le débit de requêtes, la longueur des entrées/sorties et le nombre de clients simultanés. Le benchmark mesure des indicateurs clés, notamment le débit (jetons par seconde), le temps d'obtention du premier jeton et le temps par jeton de sortie (TPOT), permettant ainsi aux utilisateurs de comprendre les performances de vLLM sous différentes conditions de charge.

Nous avons testé les performances d'inférence sur une suite complète de modèles couvrant diverses architectures, échelles de paramètres et stratégies de quantification, et évalué le débit sous différents profils de concurrence.

### Performances du modèle dense

Les modèles denses suivent l'architecture LLM conventionnelle, dans laquelle tous les paramètres et activations sont utilisés lors de l'inférence, ce qui engendre un traitement plus intensif en ressources de calcul que leurs homologues clairsemés. Afin d'évaluer de manière exhaustive les performances des différents modèles, en fonction de leur échelle et des stratégies de quantification, nous avons comparé plusieurs configurations de modèles denses de la famille Llama 3.1 8B.

## Llama 3.1 8B Precision FP8

Le modèle Llama 3.1 8B, exécuté en précision FP8, présente un profil de mise à l'échelle différent de celui de la configuration en précision standard. Il privilégie l'efficacité à des niveaux de concurrence modérés, au détriment du débit maximal pour des tailles de lots plus importantes. En mode mono-utilisateur (BS=1), le modèle atteint 67.8 tok/s par utilisateur, pour un débit total de 135.6 tok/s et un TPOT d'environ 3.1 ms, établissant ainsi une performance de base inférieure en flux unique par rapport à la précision maximale.

À mesure que la taille des lots augmente, le débit total croît rapidement tandis que le débit par utilisateur diminue de manière contrôlée. À BS=2, le débit total atteint 271 tok/s, soit 66.8 tok/s par utilisateur. À BS=4, le débit total atteint 523 tok/s et à BS=8, le modèle fournit 1 017 tok/s, soit 63.6 tok/s par utilisateur. La latence reste stable à ces faibles niveaux de concurrence, ce qui rend FP8 particulièrement adapté aux scénarios d'inférence multi-utilisateurs légers.

La mise à l'échelle se poursuit jusqu'à BS=16, où le modèle maintient un débit total de 1 918 tok/s, soit 60.0 tok/s par utilisateur. À BS=32, le débit total se stabilise à environ 1 920 tok/s, le débit par utilisateur chutant à 30.0 tok/s. Ce plateau indique que la configuration FP8 atteint son point de saturation plus tôt que la précision standard, privilégiant l'efficacité et la prévisibilité au détriment du débit agrégé maximal.

