---
id: collect-250926-servers-hardware/servers-hardware/fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43-3
title: "fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "valuation"]
source: docs/RAG/clean4/fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43.md
source_anchor: ""
source_lines: [90, 115]
sha256: 1af0d991ff1c5bfc8119a687d6a99100fdc060f42d43b89671dd1154f2789751
---

# fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43

L'utilitaire populaire 7-Zip dispose d'un test de mémoire intégré qui démontre très bien les performances du processeur. Dans ce test, nous l'exécutons avec une taille de dictionnaire de 128 Mo lorsque cela est possible. Dans ce test, le 6780E devance à peine le Platinum 8592+ dans la note totale. Pour la décompression, le 6766E a également devancé le 8592+.
| Compression à 7 zips | Xeon 6780E (256 Go DDR5) | Xeon 6766E (256 Go DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 Go DDR5 4400 XNUMX MHz) | 
|---|---|---|---|---|---|
| Compression |  |  |  |  |  | 
| Utilisation actuelle du processeur | 5,891 % | 4,768 % | 5,609 % | 5,732 % | 5,482 % | 
| Courant nominal/utilisation | 4.985 GIPS | 4.614 GIPS | 4.912 GIPS | 3.912 GIPS | 4.628 GIPS | 
| Courant | 293.689 GIPS | 220.001 GIPS | 275,503 GIPS | 224.209 GIPS | 253.724 GIPS | 
| Utilisation résultante du processeur | 5,603 % | 5,103 % | 5,605 % | 5,669 % | 5,475 % | 
| Évaluation/utilisation résultante | 4.954 GIPS | 4.638 GIPS | 4.883 GIPS | 3.923 GIPS | 4.628 GIPS | 
| Note résultante | 277.670 GIPS | 236.910 GIPS | 273.716 GIPS | 222.407 GIPS | 253.382 GIPS | 
| Décompression |  |  |  |  |  | 
| Utilisation actuelle du processeur | 5,962 GIPS | 5,798 % | 6,243 % | 5,852 % | 6,219 % | 
| Courant nominal/utilisation | 4.550 GIPS | 4.152 GIPS | 3.635 GIPS | 3.423 GIPS | 3.745 GIPS | 
| Courant | 271.266 GIPS | 240.693 GIPS | 226.917 GIPS | 200.350 GIPS | 231.916 GIPS | 
| Utilisation résultante du processeur | 5,832 % | 6,029 % | 6,232 % | 5,894 % | 6,129 % | 
| Évaluation/utilisation résultante | 4.540 GIPS | 4.161 GIPS | 3.654 GIPS | 3.385 GIPS | 3.871 GIPS | 
| Note résultante | 264.764 GIPS | 250.853 GIPS | 227.744 GIPS | 199.363 GIPS | 237.259 GIPS | 
| Note totale |  |  |  |  |  | 
| Utilisation totale du processeur | 5,717 % | 5,566 % | 5,919 % | 5,781 % | 5,802 % | 
| Note totale/utilisation | 4.747 GIPS | 4.399 GIPS | 4.269 GIPS | 3.654 GIPS | 4.249 GIPS | 
| Note totale | 271.217 GIPS | 243.882 GIPS | 250.730 GIPS | 210.363 GIPS | 245.320 GIPS | 
Conclusion
La nouvelle gamme E-core Intel Xeon 6, qui fait partie de la famille Sierra Forest, représente un changement de rythme en matière de cadence de révision. Dans les versions précédentes d'Intel, nous avons généralement vu les processeurs haut de gamme en premier ; cette fois-ci, Intel est en tête avec les SKU efficaces de milieu de gamme.
Cela dit, les processeurs Sierra Forest résistent bien, même par rapport aux modèles évolutifs de 5e génération précédents. Mais d'après Intel, les clients cherchant à acheter du nouveau matériel vont les comparer à des systèmes vieux d'environ 5 ans qui sont en train d'être mis hors service, exécutant quelque chose comme les processeurs Intel Xeon 8280. Avec ces systèmes comme composants, les processeurs E-Core Xeon 6 offrent un argument massif en matière de densité et d’économie d’énergie.
Dans cette revue, nous avons affaire à des processeurs très anciens et à un serveur assez performants pour une visite de Sierra Forest, mais pas tout à fait là où nous le souhaitons pour une revue complète. Entre les bugs du système d'exploitation Windows (non présents dans nos tests Ubuntu) et l'attente des révisions du BIOS, il est clair qu'il y a de la place pour grandir et maximiser les performances de ces CPU. Avec ce petit avant-goût, nous sommes optimistes quant à ce que Sierra Forest peut offrir aujourd'hui et à ce que les processeurs Granite Rapids peuvent offrir plus tard cette année pour les charges de travail performantes. De plus, ne dormons pas sur l'éventuel 288 E-Core 6900 Series au cours des prochains trimestres. Intel va offrir à ses clients plus d'options que jamais en matière de réglage de l'infrastructure pour leurs applications.
Nous attendons bientôt le matériel final d'expédition comprenant les processeurs Sierra Forest en laboratoire. Nous fournirons une analyse approfondie de la consommation d’énergie, des performances de stockage et bien plus encore dans les prochaines revues.
