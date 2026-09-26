---
id: collect-260926-rattrapage/rattrapage/fr-review-lenovo-thinksystem-sr650-v4-review-versatile-2u-workhorse-b60ae752-4
title: "Test du Lenovo ThinkSystem SR650 V4 : Un serveur 2U polyvalent et performant"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Nvidia"]
dates: []
keywords: ["fp8", "llama", "mixture of experts", "moe", "nvfp4", "nvidia"]
source: docs/RAG/lot-rattrapage/servers-reviews/fr-review-lenovo-thinksystem-sr650-v4-review-versatile-2u-workhorse-b60ae752.md
source_anchor: ""
source_lines: [186, 225]
sha256: 03abaf2bbfe9a09ae3e3db0762e9333d2b9702c9fd0c7c141a3acd7112436136
---

# Test du Lenovo ThinkSystem SR650 V4 : Un serveur 2U polyvalent et performant

Globalement, FP8 permet d'obtenir un débit total environ 14 fois supérieur entre BS=1 et son pic à BS=16-32. La latence reste constante jusqu'à BS=16, puis se stabilise parallèlement au débit à des niveaux de concurrence plus élevés, ce qui souligne que FP8 est un choix pratique pour les charges de travail d'inférence équilibrées et sensibles à la latence plutôt que pour une mise à l'échelle extrême par lots.

## Llama 3.1 8B Standard Précision

Avec une précision standard, le modèle Llama 3.1 8B présente un comportement prévisible en termes de mise à l'échelle à faible concurrence, avec d'excellentes performances par utilisateur pour les petits lots et un débit agrégé en constante augmentation avec la concurrence. En mode mono-utilisateur (BS=1), le modèle atteint environ 45.8 tok/s par utilisateur, soit un débit total de 91.6 tok/s, établissant ainsi une base solide pour les charges de travail sensibles à la latence.

À mesure que la concurrence augmente, le débit total croît efficacement tandis que le débit par utilisateur diminue progressivement. À BS=2, le débit total atteint 176 tok/s, soit 44.0 tok/s par utilisateur, et à BS=4, il atteint 344 tok/s, soit 43.0 tok/s par utilisateur. Ce comportement se maintient jusqu'à BS=8, où le modèle conserve un débit total de 672 tok/s, soit 42.0 tok/s par utilisateur, ce qui indique une bonne utilisation sans dégradation significative du débit par utilisateur.

La mise à l'échelle se poursuit jusqu'à BS=16, où le débit total atteint un pic d'environ 1 280 tok/s, soit 40.0 tok/s par utilisateur. À BS=32, le débit total reste stable autour de 1 280 tok/s, tandis que le débit par utilisateur chute à 20.0 tok/s, signalant la saturation du pipeline d'exécution. Ce plateau suggère qu'à précision standard, le modèle atteint son point de fonctionnement optimal autour de BS=16, offrant un équilibre entre débit et réactivité.

Globalement, la précision standard permet d'obtenir un débit total environ 14 fois supérieur entre BS=1 et son débit maximal, avec des performances stables par utilisateur pour des niveaux de concurrence modérés. Ce profil rend la précision standard particulièrement adaptée aux déploiements privilégiant une latence constante et une mise à l'échelle prévisible plutôt qu'une expansion par lots trop rapide.

### Performances du modèle parcimonieux

Les modèles épars, notamment les architectures Mixture of Experts (MoE), constituent une approche émergente pour la mise à l'échelle efficace des modèles de langage. Ces architectures conservent un nombre total de paramètres élevé tout en n'activant qu'un sous-ensemble de paramètres par jeton, ce qui peut potentiellement améliorer les performances par paramètre actif.

### Performances du Qwen3-Coder-30B-A3B FP8

Le modèle Qwen3-Coder-30B-A3B, exécuté en précision FP8, présente un profil de mise à l'échelle optimisé pour une concurrence modérée, offrant d'excellentes performances par utilisateur tout en atteignant la saturation plus rapidement que les modèles plus petits. En mode mono-utilisateur (BS=1), le modèle atteint environ 98 tok/s par utilisateur, pour un débit total de 196 tok/s, établissant ainsi une base de référence élevée pour un flux unique, témoignant de son optimisation pour les charges de travail de génération de code.

À mesure que la concurrence augmente, le débit total croît efficacement tandis que le débit par utilisateur diminue de manière contrôlée. Avec BS=2, le débit total atteint 317 tok/s, soit 79 tok/s par utilisateur, et avec BS=4, le modèle fournit un débit total de 477 tok/s, soit 59 tok/s par utilisateur. Ces résultats démontrent une utilisation efficace des ressources de calcul disponibles grâce à un traitement par lots faible à modéré, sans dégradation brutale de la réactivité par utilisateur.

Le modèle atteint son débit agrégé maximal à BS=8, se maintenant à environ 905 tok/s de débit total, soit 56 tok/s par utilisateur. La mise à l'échelle se poursuit marginalement jusqu'à BS=16, où le débit total augmente légèrement à 920 tok/s, tandis que le débit par utilisateur chute à 29 tok/s, indiquant la saturation du pipeline d'inférence. Au-delà de ce point, l'augmentation de la concurrence n'apporte que des gains minimes en termes de débit total, tout en impactant significativement les performances par utilisateur.

Globalement, la configuration Qwen3-Coder-30B-A3B FP8 offre un débit total environ 4.7 fois supérieur entre BS=1 et son débit maximal à BS=16. Les caractéristiques de latence et de débit indiquent que cette configuration est idéale pour les charges de travail de codage multi-utilisateurs modérées nécessitant des performances élevées par requête. Toutefois, le traitement par lots à grande échelle n'est pas son objectif principal.

### Performances des types de données à micro-échelle

La micro-échelle représente une approche de quantification avancée qui applique des facteurs d'échelle précis à de petits blocs de poids, plutôt qu'une quantification uniforme à de grands groupes de paramètres. Le format NVFP4 de NVIDIA implémente cette technique grâce à une représentation en virgule flottante par blocs, où chaque bloc de micro-échelle de 8 à 32 valeurs partage un exposant typique comme facteur d'échelle. Cette approche granulaire préserve la précision numérique tout en assurant une représentation sur 4 bits, maintenant ainsi la plage dynamique essentielle aux architectures de transformateurs. Ce format s'intègre à l'architecture Tensor Core de NVIDIA, permettant un calcul efficace en précision mixte avec décompression à la volée lors des opérations matricielles.

### Performances de GPT-OSS-20b

Le modèle gpt-oss-20b présente d'excellentes performances de mise à l'échelle pour des niveaux de concurrence croissants, avec un débit agrégé particulièrement élevé pour les lots de grande taille. En mode mono-utilisateur (BS=1), le modèle atteint environ 138 tok/s par utilisateur, soit un débit total de 276 tok/s, établissant ainsi une base solide pour l'inférence mono-flux.

À mesure que la concurrence augmente, le débit total croît fortement tandis que le débit par utilisateur diminue, comme prévu. Avec BS=2, le débit total atteint 420 tok/s, soit 105 tok/s par utilisateur, et avec BS=4, il atteint 690 tok/s au total, soit 86 tok/s par utilisateur. Cette progression constante témoigne d'une utilisation efficace des ressources de calcul disponibles grâce à un traitement par lots faible à modéré.

La mise à l'échelle se poursuit jusqu'à BS=8, où le modèle atteint un débit total d'environ 1 120 tok/s à raison de 70 tok/s par utilisateur, et jusqu'à BS=16, où le débit total passe à 1 900 tok/s, soit 60 tok/s par utilisateur. Ces résultats montrent que gpt-oss-20b conserve des performances relativement élevées par utilisateur même lorsque la concurrence augmente, ce qui le rend particulièrement adapté aux scénarios de service multi-utilisateurs.

Le modèle atteint son débit agrégé maximal à BS=32, avec un débit total d'environ 3 250 tok/s, le débit par utilisateur chutant à 50 tok/s. Cela représente une augmentation de 11.8 fois du débit total par rapport aux performances d'un utilisateur unique. Bien que le débit par utilisateur continue de diminuer pour des tailles de lots plus importantes, l'efficacité globale de la mise à l'échelle reste élevée, ce qui indique que le modèle tire parti d'une concurrence accrue sans atteindre de point de saturation prématuré.

