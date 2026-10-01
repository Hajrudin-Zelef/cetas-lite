---
id: collect-261001-general-networking/general-networking/fr-review-wd-gold-22tb-and-qnap-tvs-h1288x-review-9d9f01ab-2
title: "fr-review-wd-gold-22tb-and-qnap-tvs-h1288x-review-9d9f01ab"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark"]
source: docs/RAG/collect-261001-general-networking/fr-review-wd-gold-22tb-and-qnap-tvs-h1288x-review-9d9f01ab.md
source_anchor: ""
source_lines: [90, 121]
sha256: a7836a4c29c2f051cdbe5353bc666264864f5afcda6e84a139985ea9e01bd5cc
---

# fr-review-wd-gold-22tb-and-qnap-tvs-h1288x-review-9d9f01ab

Notre processus de référence de disque dur d'entreprise préconditionne chaque ensemble de disques dans un état stable avec la même charge de travail avec laquelle l'appareil sera testé sous une lourde charge de 16 threads, avec une file d'attente exceptionnelle de 16 par thread. L'appareil est ensuite testé à des intervalles définis dans plusieurs profils de profondeur de thread/file d'attente pour montrer les performances en cas d'utilisation légère et intensive. Étant donné que les disques durs atteignent très rapidement leur niveau de performance nominal, nous ne représentons graphiquement que les principales sections de chaque test.
Tests de préconditionnement et d'état stable primaire :
- Débit (agrégat IOPS lecture + écriture)
- Latence moyenne (latence de lecture + écriture moyennée ensemble)
- Latence maximale (latence maximale de lecture ou d'écriture)
- Écart-type de latence (écart-type de lecture + écriture moyenné ensemble)
Notre analyse de charge de travail synthétique d'entreprise comprend quatre profils basés sur des tâches réelles. Ces profils ont été développés pour faciliter la comparaison avec nos références passées, ainsi qu'avec des valeurs largement publiées telles que la vitesse de lecture et d'écriture maximale de 4K et 8K 70/30, qui est couramment utilisée pour les disques d'entreprise.
4K
- 100 % de lecture ou 100 % d'écriture
- 100% 4K
8K70/30
- 70 % de lecture, 30 % d'écriture
- 100% 8K
128K (séquentiel)
- 100 % de lecture ou 100 % d'écriture
- 100% 128K
Notre premier test de débit mesure les performances aléatoires 4K à l'intérieur du QNAP TVS-h1288x. Ici, le WD Gold 22 To a affiché des performances similaires entre les configurations iSCSI et SMB, ce qui n'était pas le cas avec le Synology NAS. Plus précisément, il a montré 5,578 804 IOPS en lecture et 5,774 IOPS en écriture (SMB) et 978 XNUMX IOPS en lecture et XNUMX IOPS en écriture (iSCSI).
En latence moyenne, le WD Gold de 22 To a pu atteindre respectivement 45.9 ms et 318.2 ms en lecture et écriture en SMB, tandis que l'iSCSI affichait 44.3 ms en lecture et 261.44 ms en écriture.
Passant à la latence maximale, le WD Gold atteint 2,188.9 17,492 ms en lecture et 802.88 6,647 ms en écriture en SMB, tandis qu'il atteint XNUMX ms en lecture et XNUMX XNUMX ms en écriture en iSCSI.
Pour l'écart type, le WD Gold a enregistré 40 ms de lecture et 1,075.1 44.86 ms d'écriture en SMB, et 203.95 ms de lecture et XNUMX ms d'écriture en iSCSI.
Notre prochain benchmark soumet les disques à une activité de lecture et d'écriture de 100 % à un débit séquentiel de 8K. À l'intérieur du QNAP TVS-h1288x, le disque WD Gold de 22 To affichait 151,481 96,064 IOPS en lecture et 206,476 142,538 IOPS en écriture en SMB, tout en atteignant XNUMX XNUMX IOPS en lecture et XNUMX XNUMX IOPS en écriture en iSCSI.
Notre prochain test passe d'un scénario de lecture/écriture séquentielle 8K pure à 100 % à une charge de travail mixte 8K 70/30, qui démontrera comment les performances évoluent dans un environnement allant de 2T/2Q à 16T/16Q.
Le premier est le débit, où le WD Gold 22 To avait une plage de 765 IOPS à 2,137 639 IOPS en SMB et de 1,724 IOPS à XNUMX XNUMX IOPS en iSCSI.
Avec une latence moyenne de 8K 70/30 en utilisant un QNAP TVS-h1288x, le 22 To Gold 22 To a affiché 5.22 ms à 118.83 ms en SMB, tandis que l'iSCSI a enregistré une plage de 6.25 ms à 148.32 ms.
Passant aux chiffres de latence maximale, le WD Gold 22 To avait une plage de 1,713.61 6,979.94 ms à 500.01 7,080.62 ms en SMB tandis que l'iSCSI affichait XNUMX ms à XNUMX XNUMX ms.
Pour nos résultats de latence d'écart type, le 22 To Gold a atteint 20.99 ms à 283.28 ms (SMB) et 12.56 ms à 307.77 ms (iSCSI).
Notre dernier test est le benchmark 128K, un test séquentiel à gros blocs montrant la vitesse de transfert séquentielle la plus élevée. Le WD Gold 22 To affiche 1.03 Go/s en lecture et 947 Mo/s en écriture en SMB, tandis que l'iSCSI enregistre 1.37 Go/s en lecture et un impressionnant 2.31 Go/s en écriture.
Conclusion
Dans l'ensemble, il s'est avéré très utile pour analyser les performances des principaux facteurs de forme NAS de bureau (ou SMB) comme le QNAP TVS-h1288x polyvalent lorsqu'il est combiné avec les disques WD Gold 22 To, car les résultats de nos tests prouvent certainement qu'il apportera un beaucoup de flexibilité aux organisations. Avec le point de capacité massif de Gold et les performances de pointe pour les disques durs 3.5″, les organisations seront en mesure de réduire leur TCO global et de réduire leur empreinte physique. Les PME qui ont d'énormes besoins en données bénéficieront le plus des disques WD Gold.
Pour tester ses performances, nous avons ajouté huit disques WD Gold 22 To à l'intérieur du NAS QNAP et l'avons soumis à nos tests habituels. Dans l'ensemble, nous avons constaté de solides résultats dans l'ensemble de nos analyses comparatives, en particulier lorsqu'elles sont configurées avec le protocole iSCSI.
Les spécificités incluent 5,578 804 IOPS en lecture / 5,774 IOPS en écriture (SMB) et 978 4 IOPS en lecture / 151,481 IOPS en écriture (iSCSI) en 96,064K aléatoire, et 206,476 142,538 IOPS en lecture / 8 765 IOPS en écriture (SMB) et 2,137 639 IOPS en lecture / 1,724 1.03 IOPS en écriture (iSCSI) pendant la charge de travail séquentielle 947K. De plus, le WD Gold a montré une plage de 1.37 IOPS à 2.31 XNUMX IOPS en SMB et de XNUMX IOPS à XNUMX XNUMX IOPS en iSCSI. Pour notre test séquentiel à grands blocs, il a affiché XNUMX Go/s en lecture et XNUMX Mo/s en écriture en SMB, tandis que iSCSI a affiché XNUMX Go/s en lecture et XNUMX Go/s en écriture.
Comme dans le cas du Synology DiskStation, le WD Gold 22 To s'est très bien comporté lorsqu'il est combiné avec le QNAP TVS-h1288x de qualité SMB, affichant des chiffres solides tout au long (en particulier dans nos tests aléatoires 4K et séquentiels 8K). Avec ses 8 baies HDD disponibles et ses 4 baies flash, le NAS QNAP offrira aux bureaux professionnels une tonne de flexibilité pour se développer ainsi que les performances qui vont avec.
