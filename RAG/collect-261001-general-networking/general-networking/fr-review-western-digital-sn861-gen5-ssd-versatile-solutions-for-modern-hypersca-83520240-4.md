---
id: collect-261001-general-networking/general-networking/fr-review-western-digital-sn861-gen5-ssd-versatile-solutions-for-modern-hypersca-83520240-4
title: "fr-review-western-digital-sn861-gen5-ssd-versatile-solutions-for-modern-hypersca-83520240"
domain: general-networking
role: reference
task: reference
actors: ["Nvidia", "Samsung"]
dates: []
keywords: ["gpu", "nvidia"]
source: docs/RAG/collect-261001-general-networking/fr-review-western-digital-sn861-gen5-ssd-versatile-solutions-for-modern-hypersca-83520240.md
source_anchor: ""
source_lines: [66, 73]
sha256: ecebbd8aeb1b38fffd03c5a4177b4fc04cad470923fdcdeaaee3565e3cc5844a
---

# fr-review-western-digital-sn861-gen5-ssd-versatile-solutions-for-modern-hypersca-83520240

Pour obtenir des performances optimales, en particulier dans les environnements gourmands en GPU, il est nécessaire de garantir un échange de données à haut débit entre les GPU et le stockage. Par exemple, pour saturer complètement la bande passante d'un GPU NVIDIA H100, nous devions atteindre un débit d'environ 64 Go/s, ce qui implique l'utilisation de solutions et de technologies de stockage NVMe hautes performances telles que NVIDIA GPUDirect™. Cette intégration réduit la latence et maximise le débit de données, garantissant une utilisation efficace du GPU pour un traitement plus rapide et plus efficace des ensembles de données à grande échelle.
Lorsque nous examinons les différences de bande passante dans ce que le Gen4 SN655 peut faire à 6.8 Go/s de pointe par rapport à 13.7 Go/s du SN861, il est évident de voir les avantages de passer à un SSD Gen5. Pour atteindre 64 Go/s avec le modèle de la génération précédente, vous avez besoin de dix SSD, alors que le SN861 pourrait atteindre cet objectif avec seulement cinq. Cette différence pourrait vous permettre d'augmenter le nombre de disques pour obtenir une bande passante ou une capacité supplémentaire.
Les performances et la capacité seront essentielles pour que le stockage puisse évoluer avec les besoins de l’IA et d’autres applications avancées. L'interface Gen5 et les performances globales améliorées par le SN861 par rapport aux disques Gen4 sont très convaincantes à cet égard, ce qui signifie que ces disques peuvent prendre en charge plus de GPU au sein d'un seul système de stockage et garantir que ces GPU sont alimentés à un rythme suffisamment rapide pour garantir une utilisation complète.
Conclusion
Le SN861 marque un grand pas en avant pour Western Digital. Le disque se présente sous des formats adaptés aux clients hyperscale et professionnels, avec des fonctionnalités telles que FDP dans le disque E1.S, adaptées à leurs cas d'utilisation potentiels. L’interface Gen5 constitue cependant l’avantage le plus apparent pour les disques, offrant un profil de performances globale impressionnant.
Le Western Digital SN861 a offert d’excellentes performances dès le départ, occupant les trois premières places dans nos charges de travail initiales aux quatre coins mesurant la bande passante séquentielle maximale et le débit aléatoire. Les points forts incluent des performances de lecture 4K aléatoires de 2.11 M IOPS et des performances d’écriture 4K aléatoires mesurant 474 1743 IOPS. Les performances de lecture séquentielle ont été solides, arrivant en deuxième position par rapport au Samsung PM13.3 avec 7.7 Go/s, bien qu'il ait pu prendre la tête de la bande passante d'écriture séquentielle mesurant XNUMX Go/s.
Dans nos charges de travail VDbench, principalement axées sur des charges de travail mixtes ou des transferts de plus petites tailles de blocs, le SN861 a continué à fonctionner exceptionnellement bien. Nous avons mesuré une forte vitesse d'écriture séquentielle 16K de 200K IOPS et des avances solides dans les tests de mixage 70/30 R/W couvrant les tailles de transfert 4K, 8K et 16K. Dans nos charges de travail VDI, le SN861 a échangé la première place avec le KIOXIA CM7-R, qui était au coude à coude dans certains domaines. Dans l’ensemble, le Western Digital SN861 a obtenu de bons résultats dans notre gamme de tests.
Il est évident que les derniers SSD Gen5, tels que le Western Digital SN861, influencent les résultats commerciaux. Si vous avez besoin d’une preuve, ne cherchez pas plus loin que leur impact sur la révolution de l’IA. Nous l'avons vu lors de nos tests ; Les systèmes d'IA ont besoin d'un stockage rapide pour que les GPU continuent de fonctionner, que ce soit dans un cache comme l'exemple NVIDIA IndeX ci-dessus ou dans des baies de stockage partagées ou des serveurs GPU. Western Digital a très bien réussi à positionner le SN861 pour ces charges de travail avancées tout en proposant également des SKU compatibles FDP pour les hyperscalers.
