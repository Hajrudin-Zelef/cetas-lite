---
id: collect-261001-general-networking/general-networking/fr-review-western-digital-sn861-gen5-ssd-versatile-solutions-for-modern-hypersca-83520240-2
title: "fr-review-western-digital-sn861-gen5-ssd-versatile-solutions-for-modern-hypersca-83520240"
domain: general-networking
role: reference
task: reference
actors: ["Intel", "Oracle", "Samsung"]
dates: []
keywords: ["intel"]
source: docs/RAG/collect-261001-general-networking/fr-review-western-digital-sn861-gen5-ssd-versatile-solutions-for-modern-hypersca-83520240.md
source_anchor: ""
source_lines: [49, 61]
sha256: a8bf1e984e2ecea1022bd11274d0053f16e30dc649d726584f026359d0a99f79
---

# fr-review-western-digital-sn861-gen5-ssd-versatile-solutions-for-modern-hypersca-83520240

Les performances de lecture aléatoire 4K étaient particulièrement bonnes, mesurant 2.11 millions d’IOPS, avec 1.96 millions d’IOPS du KIOXIA CM7-R comme deuxième plus proche. Lorsque nous avons examiné les performances d'écriture 4K aléatoires, le Western Digital SN861 est également arrivé en premier, avec une vitesse de 474 1743 IOPS, le Samsung PM320 avec 861 XNUMX IOPS étant le modèle le plus proche. Dans nos charges de travail aux quatre coins, le Western Digital SNXNUMX a obtenu le meilleur chiffre dans trois des quatre tests.
Pour tester le SSD SN861 Gen5, nous avons utilisé le serveur Dell® PowerEdge® R760 de notre laboratoire. Ce serveur rack 2U, extrêmement polyvalent, prend en charge deux processeurs Intel Xeon de 4e génération et peut accueillir jusqu'à 24 disques NVMe. Il est conçu pour les charges de travail mixtes, les bases de données et les environnements VDI. Il est important de noter que la version du CM7-R testée dans cet article provient d'un serveur Dell équipé du firmware d'origine Dell. Les performances de ce disque peuvent différer avec le firmware standard de KIOXIA.
Configuration Dell PowerEdge R760 :
- Double Intel® Xeon® Gold 6430 (32 cœurs/64 threads, base 1.9 GHz)
- 1 To de RAM DDR5
- Ubuntu 22.04
Pour une flexibilité ultime, nous avons également travaillé avec Serial Cables, qui nous a fourni un JBOF PCIe Gen8 à 5 baies pour les tests SSD U.2/U.3, M.2 et EDSFF. Cela nous permet de tester tous les types de disques actuels et émergents sur le même matériel de test. VDbench a également été utilisé pour comparer les performances évolutives de notre sélection de SSD dans différents types de charges de travail. Notre processus de test pour ces tests de référence remplit toute la surface du disque avec des données, puis partitionne une section de disque égale à 25 % de la capacité du disque pour simuler la manière dont le disque pourrait répondre aux charges de travail des applications. Cela diffère des tests d’entropie complète, qui utilisent 100 % du disque et le mettent dans un état stable. En conséquence, ces chiffres refléteront des vitesses d’écriture plus élevées et soutenues.
Profils:
- Lecture séquentielle 16K : 100 % de lecture, 32 threads, 0-120 % d'iorate
- Écriture séquentielle 16K : 100 % d'écriture, 16 threads, 0-120 % d'iorate
- Mélange aléatoire 4K, 8K et 16K 70R/30W, 64 fils, 0-120 % de vitesse
- Base de données synthétique : SQL et Oracle
- Traces de clone complet et de clone lié VDI
