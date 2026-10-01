---
id: collect-261001-general-networking/general-networking/fr-review-lsi-megaraid-sas3-9361-8i-review-f283ce9b-2
title: "fr-review-lsi-megaraid-sas3-9361-8i-review-f283ce9b"
domain: general-networking
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["benchmark"]
source: docs/RAG/collect-261001-general-networking/fr-review-lsi-megaraid-sas3-9361-8i-review-f283ce9b.md
source_anchor: ""
source_lines: [58, 76]
sha256: e430358c748e71cfcba6ab59b98a50267db35977fcd5657cf70dc11254b124df
---

# fr-review-lsi-megaraid-sas3-9361-8i-review-f283ce9b

Dans notre test de charge de travail mixte 8k, 70% lecture, 30% écriture, mesurant le débit, la configuration LSI MegaRAID RAID0 offrait sans surprise les meilleures performances globales à pratiquement toutes les profondeurs avec un pic supérieur à 450,000 10 IOPS, soit presque le double du débit de la deuxième place RAID 9300 configuration. La différence surprenante dans cette revue était la quantité d'amélioration des performances offerte par la carte RAID par rapport à la LSI 8-0i IR en RAID10, qui atteignait environ la moitié du débit. La même différence peut être observée en comparant RAID9361 sur le 8-9300i par rapport au 8-XNUMXi IR.
Le LSI-Supermicro 9300-8i en configuration RAID 0 offrait la latence moyenne globale la plus faible du 16T/16Q. La configuration MegaRAID RAID 5 et 6 de LSI avait le score de latence le plus élevé de 16T/16Q, atteignant presque 30 ms.
La plupart des configurations RAID ont obtenu d'assez bons résultats dans les tests de latence moyenne, en particulier MegaRAID RAID 0, qui n'avait presque pas de pics de latence. Le MegaRAID en configuration RAID 6 faisait cependant exception avec une latence moyenne maximale atteignant près de 10ms. Un élément qui a suivi le graphique était la latence maximale du LSI 9300-8i IR en RAID10, qui flottait au-dessus de 1,000 XNUMX ms.
Parmi les configurations RAID, le MegaRAID RAID50 et le LSI 9300-8i IR RAID0 offraient l'écart type le plus étroit au niveau 16T/16Q. Encore une fois, la configuration MegaRAID RAID 5 et 6 de LSI a un peu augmenté au cours de ces tests et avait un écart type significativement plus élevé à la fin. Le 9300-8i IR en RAID10 a de nouveau tracé ce graphique dans la plage d'écart type de 11 à 37 ms.
Dans notre prochain benchmark, nous nous en tenons toujours à la taille de transfert de 8k, mais en passant à des opérations séquentielles de lecture à 100% et d'écriture à 100%. Sous cette charge de travail, le MegaRAID en configuration RAID 0 a continué de briller, atteignant respectivement 482,722 438,929 IOPS et XNUMX XNUMX IOPS en lecture et en écriture. Nous voyons également l'avantage de passer à une solution RAID matérielle sur le HBA avec des capacités RAID intégrées, où pour une configuration RAID donnée, la solution H/W avait le double des performances de lecture.
Dans notre benchmark synthétique final, nous avons utilisé une taille de transfert beaucoup plus grande de 128k avec des opérations de lecture à 100% et d'écriture à 100%. Sous cette charge de travail, la plupart des configurations RAID pour MegaRAID et Supermicro étaient au coude à coude pour le débit de lecture ; cependant, le MegaRAID RAID 0 est une fois de plus arrivé en tête dans la colonne d'écriture avec une marge significative.
Conclusion
Le LSI MegaRAID SAS 9361-8i apporte une solution RAID SAS3 complète sur le marché pour les serveurs qui veulent tirer le meilleur parti de leur stockage haut de gamme, mais qui ont besoin de la fiabilité, de la compatibilité et de la protection des données qui sont l'héritage de LSI. Avec les extensions, le 9361-8i peut prendre en charge jusqu'à 128 disques SATA ou SAS. Bien que la carte soit conçue pour les disques SAS3, elle est bien sûr rétrocompatible. LSI prend en charge les niveaux RAID 0,1,5,6,10, 50, 60, 9361, 8, XNUMX et XNUMX avec le XNUMX-XNUMXi, et la protection des données d'entreprise via CacheVault (facultatif) qui ajoute une autre couche de protection si un disque semble en panne.
En ce qui concerne les performances, le LSI 9361-8i n'est pas en reste lorsqu'il s'agit de gérer les derniers et les meilleurs SSD SAS3. En gérant huit SSD Hitachi SSD800MM, nous avons testé tous les modes RAID pris en charge sur la carte et les avons comparés à l'ensemble de fonctionnalités RAID intégré sur un LSI/Supermicro 9300-8i. En RAID0, nous avons mesuré plus de 556 573 IOPS en lecture, 4 450 IOPS en écriture dans notre test aléatoire 8 70 et un peu moins de 30 0 IOPS dans notre test de charge de travail mixte 10 XNUMX XNUMX/XNUMX. Cela se traduit par environ le double des performances par rapport au modèle HBA qui prend en charge à la fois RAIDXNUMX et RAIDXNUMX. Pour être juste, le HBA n'est évidemment pas destiné aux configurations les plus performantes, mais l'échelle comparative est pertinente.
En passant à des charges de travail séquentielles, nous avons constaté que le débit a atteint son maximum lors de notre test 8 482 à 438 0 IOPS en lecture et 9361 8 IOPS en écriture en RAID4.9. Dans les transferts de gros blocs, le 50-0i a fait pencher la balance en offrant un peu moins de 4.3 Go/s en lecture dans son mode RAID9361 et RAID8 offrant la bande passante d'écriture la plus élevée à 2 Go/s en continu. Dans l'ensemble, lorsqu'il s'agit de prendre en charge la dernière génération de SSD, le LSI MegaRAID 3-1i a beaucoup à offrir. Il s'appuie sur les modèles SASXNUMX incroyablement populaires qui l'ont précédé, tout en ajoutant la prise en charge des SSD SASXNUMX qui peuvent atteindre des vitesses individuelles allant jusqu'à XNUMX Go/s R/W.
Avantages
- Excellentes performances à coupler avec les SSD SAS3
- Offre la protection du cache flash CacheVault
- Large prise en charge des pilotes
Inconvénients
- Limité à 8 ports internes sur la carte
Conclusion
L'adaptateur de stockage LSI MegaRAID SAS 9361-8i permet aux SSD SAS3 de dernière génération de se concentrer sur la saturation du bus PCIe Gen3 sans avoir à se soucier de la carte RAID qui les retient.
LSI MegaRAID 9361 sur Amazon
