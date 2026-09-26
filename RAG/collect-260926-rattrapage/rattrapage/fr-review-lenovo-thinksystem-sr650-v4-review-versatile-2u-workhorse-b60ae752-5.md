---
id: collect-260926-rattrapage/rattrapage/fr-review-lenovo-thinksystem-sr650-v4-review-versatile-2u-workhorse-b60ae752-5
title: "Test du Lenovo ThinkSystem SR650 V4 : Un serveur 2U polyvalent et performant"
domain: rattrapage
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["apache", "gpu", "intel", "open source"]
source: docs/RAG/lot-rattrapage/servers-reviews/fr-review-lenovo-thinksystem-sr650-v4-review-versatile-2u-workhorse-b60ae752.md
source_anchor: ""
source_lines: [226, 256]
sha256: abadf1511394e0b941cd6079ce4d9b889ccfe97005e207d8c5f0b25936237ad7
---

# Test du Lenovo ThinkSystem SR650 V4 : Un serveur 2U polyvalent et performant

Globalement, gpt-oss-20b présente un profil de scalabilité équilibré, alliant d'excellentes performances mono-utilisateur à une forte capacité de traitement par lots. Il constitue ainsi une option intéressante pour les déploiements nécessitant à la fois une inférence individuelle réactive et un débit agrégé élevé sous une charge multi-utilisateurs importante.

## Points de repère Phoronix

Phoronix Test Suite est une plateforme d'analyse comparative automatisée et open source qui prend en charge plus de 450 profils de test et plus de 100 suites de tests via OpenBenchmarking.org. Elle gère l'ensemble du processus, de l'installation des dépendances à l'exécution des tests et à la collecte des résultats, ce qui la rend idéale pour les comparaisons de performances, la validation matérielle et l'intégration continue.

**Bande passante de la mémoire de flux** : Le SR650 V4 a fourni une bande passante mémoire de 369.6 Go/s, démontrant un débit élevé pour les charges de travail gourmandes en données et confirmant des performances mémoire multicanaux saines.

**Compression à 7 zips** : Avec 584 499 MIPS, le système a affiché d'excellentes performances de compression et de décompression, reflétant une solide efficacité multicœur et une forte capacité de calcul entier.

**Compilation du noyau** : Le serveur a terminé la compilation du noyau allmod en 256.175 secondes, un résultat respectable qui souligne sa capacité à gérer efficacement la compilation parallèle et les flux de travail des développeurs.

**Serveur Web Apache** : Avec 61 654 R/s, le SR650 V4 a démontré de solides performances de serveur Web, maintenant un débit de requêtes élevé adapté à l'hébergement frontal ou aux piles de virtualisation légères.

**Vérification OpenSSL** : Atteignant 712.6 milliards d'octets par seconde, la plateforme a démontré des performances robustes en matière de vérification cryptographique, confirmant ainsi la capacité du processeur à effectuer des opérations sécurisées nécessitant de nombreux certificats.

| Points de repère Phoronix | Lenovo Think System SR650 V4 |
|---|---|
| Stream (bande passante mémoire) | 369.6 Go/s |
| 7-ZIP | 584 499 MIPS |
| Compilation du noyau (allmod) | 256.175 secondes |
| Apache (requêtes par seconde) | 61 654 R/s |
| OpenSSL | 712 633 776 990 octets/s |

## Conclusion

Le ThinkSystem SR650 V4 représente une mise à jour majeure, car elle ne se limite pas à un simple changement de processeur. Lenovo a profité de cette génération pour apporter des améliorations significatives dans les domaines qui limitent les déploiements en conditions réelles : bande passante mémoire, disponibilité des lignes PCIe, densité de stockage et capacité d'adaptation de la configuration aux besoins évolutifs. Le passage à l'Intel Xeon 6 offre davantage d'options de cœurs et une capacité d'E/S accrue par socket. De plus, les améliorations apportées à la plateforme, notamment la prise en charge de la DDR5 plus rapide avec MRDIMM et l'extension CXL optionnelle, permettent au SR650 V4 de rester compétitif face à des charges de travail toujours plus gourmandes en mémoire.

Côté stockage, malgré notre configuration de test limitée avec U.2, l'adoption par Lenovo de la norme E3.S offre une densité NVMe supérieure avec une connectivité 1:1 et des caractéristiques thermiques optimisées, en phase avec les infrastructures modernes, notamment pour les environnements de virtualisation qui nécessitent davantage de mémoire flash locale, et pour les pipelines liés à l'IA où l'alimentation des GPU est aussi importante que la puissance de calcul brute. Les deux emplacements OCP 3.0 Gen5 x16 constituent également une amélioration pratique : vous pouvez déployer une infrastructure réseau performante sans sacrifier les emplacements PCIe que vous préférez réserver aux accélérateurs, au stockage ou aux adaptateurs spécialisés.

Si vous utilisez un SR650 V3 et que vous vous sentez limité par la vitesse de la mémoire, la marge PCIe ou la densité NVMe, le SR650 V4 représente une avancée majeure. Il conserve la facilité de maintenance et la configurabilité qui ont fait le succès de la gamme SR650, tout en offrant une plateforme évolutive qui permet aux équipes informatiques d'éviter les situations figées. Que vous construisiez un cluster de virtualisation, modernisiez des services à grande échelle ou assembliez un nœud CPU/GPU équilibré, le SR650 V4 répond aux besoins essentiels et ouvre la voie aux évolutions futures.
