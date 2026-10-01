---
id: collect-261001-general-networking/general-networking/fr-review-synology-diskstation-ds1821-review-e49563c1-2
title: "fr-review-synology-diskstation-ds1821-review-e49563c1"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Samsung"]
dates: []
keywords: ["amd", "benchmark"]
source: docs/RAG/collect-261001-general-networking/fr-review-synology-diskstation-ds1821-review-e49563c1.md
source_anchor: ""
source_lines: [63, 92]
sha256: b9f0a5d826586e33ee5c9559b053bb1ba6fed8b766f00e1b841946ae03009b43
---

# fr-review-synology-diskstation-ds1821-review-e49563c1

Notre analyse de charge de travail synthétique d'entreprise comprend quatre profils basés sur des tâches réelles. Ces profils ont été développés pour faciliter la comparaison avec nos références passées ainsi qu'avec des valeurs largement publiées telles que la vitesse de lecture et d'écriture maximale de 4k et 8k 70/30, qui est couramment utilisée pour les disques d'entreprise.
- 4K
- 
  - 100 % de lecture ou 100 % d'écriture
  - 100% 4K
- 8K 70/30
  - 70 % de lecture, 30 % d'écriture
  - 100% 8K
- 8K (séquentiel)
  - 100 % de lecture ou 100 % d'écriture
  - 100% 8K
- 128K (séquentiel)
  - 100 % de lecture ou 100 % d'écriture
  - 100% 128K
Pour nos tests, nous avons configuré le Synology DiskStation DS1821+ en RAID6 en iSCSI et SMB en utilisant à la fois une configuration HDD et SSD en utilisant les disques suivants :
- 8 x Disque dur NAS rouge WD (14 To)
- 8 x Samsung SM863 SSD (960GB)
Dans la première de nos charges de travail d'entreprise, nous avons mesuré un long échantillon de performances 4K aléatoires avec une activité d'écriture à 100 % et de lecture à 100 %. En ce qui concerne les IOPS, la configuration du disque dur Synology DiskStation DS1821+ a enregistré 1,550 595 IOPS en écriture et 1,575 IOPS en lecture en SMB, tout en atteignant 4,153 38,633 IOPS en écriture et 26,209 37,339 IOPS en lecture en iSCSI. Avec les SSD, il a montré, comme on pouvait s'y attendre, des gains significatifs (comme c'est le cas dans le reste de l'examen des performances) avec 99,988 XNUMX IOPS en écriture et XNUMX XNUMX IOPS en lecture en SMB et XNUMX XNUMX IOPS en écriture et XNUMX XNUMX IOPS en lecture en iSCSI.
Avec une latence moyenne de 4K (où plus bas est meilleur), le Synology DS1821+ affiche 165.11 ms en écriture et 429.74 ms en lecture en SMB, et 162.92 ms en écriture avec une lecture bien inférieure à 61.63 ms en iSCSI (configuration HDD). En utilisant des SSD, la latence est tombée à 6.62 ms et 9.77 ms en écriture/lecture en SMB et 6.85 ms et 2.56 ms en écriture/lecture en iSCSI.
Passant à une latence maximale de 4K, le DS1821+ a affiché 3,625.5 1,663.7 ms en écriture et 3,474.1 1,700.5 ms en lecture en SMB et 315.5 40.61 ms en écriture et 285.8 47.7 ms en lecture en iSCSI. Dans notre configuration SSD, le Synology NAS a enregistré des chiffres de XNUMX ms en écriture et XNUMX ms en lecture en SMB et XNUMX ms en écriture et XNUMX ms en lecture en iSCSI.
Pour notre dernier test 4K, nous avons examiné l'écart type. Dans notre configuration HDD, nous avons enregistré des chiffres de 315.5 ms en écriture et 446.1 ms en lecture en SMB tandis que iSCSI a atteint 242.7 ms en écriture et 56.9 ms en lecture en iSCSI. Avec les SSD, les performances sont passées à 10.5 ms en lecture et 9.22 ms en écriture en SMB et 14.5 ms en écriture et 2.17 ms en lecture en iSCSI.
Notre prochain benchmark mesure 100 % de débit séquentiel 8K avec une charge 16T16Q dans des opérations de lecture à 100 % et d'écriture à 100 %. Ici, le DS1821+ a pu atteindre 51,118 56,348 IOPS en écriture et 52,857 57,085 IOPS en lecture dans SMB et 53,992 56,436 IOPS en écriture et 56,181 57,101 IOPS en lecture dans iSCSI (HDD). Dans notre configuration SSD, le Synology NAS a enregistré une légère amélioration avec XNUMX XNUMX IOPS en écriture et XNUMX XNUMX IOPS en lecture en SMB. Nous avons constaté les meilleures performances en utilisant iSCSI avec des écritures et des lectures de XNUMX XNUMX IOPS et XNUMX XNUMX IOPS, respectivement.
Par rapport à la charge de travail fixe à 16 threads et 16 files d'attente maximales que nous avons effectuée lors du test d'écriture 100 % 4K, nos profils de charge de travail mixtes adaptent les performances à une large gamme de combinaisons thread/file d'attente. Dans ces tests, nous couvrons l'intensité de la charge de travail de 2 threads/2 files d'attente à 16 threads/16 files d'attente. Avec le débit du disque dur, SMB a affiché une plage de 344 IOPS à 385 IOPS tandis que iSCSI a atteint une plage de 575 IOPS à 1,839 8,813 IOPS. En utilisant des SSD, iSCSI était de loin la meilleure configuration avec 2 2 IOPS à 47,941T/16Q et finissant à 16 XNUMX IOPS à XNUMXT/XNUMXQ.
En regardant les performances de latence moyennes dans notre configuration SSD, notre configuration HDD a montré une plage de 11.6 ms à 663.49 ms en SMB, tandis que iSCSI a enregistré 6.93 ms à 139.11 ms. Avec les SSD, le DS1821+ affichait une plage de 0.42 ms à 18.14 ms en SMB et de 0.45 ms à 5.33 ms en iSCSI.
Pour une latence maximale, nous avons vu 894.8 ms à 6,057.28 920.6 ms en SMB et 6,891.26 ms à 35.78 425.22 ms en iSCSI dans notre configuration WD Red HDD, tandis que notre configuration Samsung SSD affichait 51.13 ms à 105.24 ms et XNUMX ms à XNUMX ms en SMB et iSCSI, respectivement.
En ce qui concerne l'écart type, notre configuration HDD a enregistré 23.27 ms à 578.69 ms en SMB et 16.67 ms à 321.66 ms en iSCSI, tandis que notre configuration SDD a affiché 0.45 ms à 21.57 ms (SMB) et 0.48 ms à 6.39 ms (iSCSI).
Le dernier benchmark Enterprise Synthetic Workload est notre test 128K, qui est un test séquentiel à grands blocs qui montre la vitesse de transfert séquentielle la plus élevée pour un appareil. Dans ce scénario de charge de travail, la configuration HDD DS1821+ avait 450.3 Mo/s en écriture et 462.6 Mo/s en lecture dans SMB et 450 Mo/s en écriture et 462.8 Mo/s en lecture dans iSCSI. Avec les SSD, les performances ont légèrement augmenté en SMB et en iSCSI, enregistrant respectivement 458 Mo/s en écriture et 463.1 Mo/s en lecture et 457.6 Mo/s en lecture et 462.9 Mo/s en écriture.
Conclusion
Le Synology DiskStation DS1821+ est un NAS à 8 baies alimenté par un processeur AMD Ryzen et destiné aux PME qui ont besoin d'un peu plus de performances et de capacité dans un NAS facile à utiliser. Comme la plupart des périphériques NAS modernes de l'entreprise, le DS1821+ est livré avec deux emplacements M.2 pour le cache SSD NVMe. Le NAS dispose de quatre ports 1GbE mais dispose d'un slot PCIe qui permet aux utilisateurs d'ajouter un 10GbE. Le DS1821+ peut également être équipé d'un maximum de 32 Go de mémoire DDR4 ECC SODIMM. Tout ce qui précède peut se traduire par une vitesse annoncée allant jusqu'à 2.3 Go/s et jusqu'à 113 1821 IOPS en haut de gamme. Ce DiskStation dispose d'un moteur de chiffrement matériel qui ajoute de la sécurité au prix de certaines de ses performances. Et le Synology DiskStation DSXNUMX+ exploite Synology DSM pour son système d'exploitation.
Pour évaluer ses performances, nous avons configuré le NAS avec des disques durs ( WD Red de 14 To ) et un SSD ( Samsung SM863 de 960 Go ), en utilisant les protocoles iSCSI et SMB. Commençons par nos tests sur disque dur. En lecture 4K, le DiskStation a atteint 1 550 IOPS en écriture et 595 IOPS en lecture via SMB, et 1 575 IOPS en écriture et 4 153 IOPS en lecture via iSCSI. Concernant la latence moyenne 4K, le NAS a affiché 165.11 ms en écriture et 429.74 ms en lecture via SMB, et 162.92 ms en écriture et 61.63 ms en lecture via iSCSI. Enfin, en lecture séquentielle 8K, nous avons constaté 51 118 IOPS en écriture et 56 348 IOPS en lecture via SMB, et 52 857 IOPS en écriture et 57 085 IOPS en lecture via iSCSI. En examinant notre séquence de blocs de grande taille de 128K, le 1821+ avait des vitesses de 450.3 Mo/s en écriture et 462.6 Mo/s en lecture en SMB et 450 Mo/s en écriture et 462.8 Mo/s en lecture en iSCSI.
