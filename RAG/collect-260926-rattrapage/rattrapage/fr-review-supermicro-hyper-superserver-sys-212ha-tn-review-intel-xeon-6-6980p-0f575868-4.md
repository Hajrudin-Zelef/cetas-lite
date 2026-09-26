---
id: collect-260926-rattrapage/rattrapage/fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868-4
title: "fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868"
domain: rattrapage
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "gpu", "valuation"]
source: docs/RAG/lot-rattrapage/servers-reviews/fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868.md
source_anchor: ""
source_lines: [129, 153]
sha256: 7926a14da7669511b9affde0357705d7be2f9f0bcf9751924fa2621426c23198
---

# fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868

- Performances de compression:Le SYS-212HA-TN atteint un taux de compression résultant de 302.712 GIPS, surpassant considérablement les autres systèmes Lenovo SR630 : 224.313 GIPS, Supermicro 1U 112H-TN : 245.823 GIPS.
- Performances de décompression:Le SYS-212HA-TN est presque à égalité avec le Lenovo SR630 en décompression, avec une note résultante de 287.823 GIPS, bien que suivi de près par le Lenovo SR630 288.457 GIPS, avec le Supermicro 1U 112H-TN derrière à 269.373 GIPS.
- Performance totale:Dans l'ensemble, le SYS-212HA-TN enregistre une note totale de 295.268 GIPS, surpassant à la fois le Lenovo SR630 256.385 GIPS et le Supermicro 1U 112H-TN 257.598 GIPS.
| Compression 7-Zip et Décompression | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | Lenovo ThinkSystem SR630 V4 (Intel Xeon 6780E, 512 Go) | Serveur Hyper Super SYS-212HA-TN (1x Intel Xeon 6980P, 384 Go de RAM) | 
|---|---|---|---|
| Compression – Utilisation actuelle du processeur | 5287 % | 5064 % | 5300 % | 
| Compression – Courant nominal/utilisation | 4.647 GIPS | 4.341 GIPS | 5.795 GIPS | 
| Compression – Courant nominal | 245.699 GIPS | 219.840 GIPS | 307.117 GIPS | 
| Compression – Utilisation du processeur résultante | 5296 % | 5156 % | 5214 % | 
| Compression – Évaluation/utilisation résultante | 4.642 GIPS | 4.350 GIPS | 5.806 GIPS | 
| Compression – Évaluation résultante | 245.823 GIPS | 224.313 GIPS | 302.712 GIPS | 
| Décompression – Utilisation actuelle du processeur | 6236 % | 6184 % | 6195 % | 
| Décompression – Évaluation/utilisation actuelle | 4.261 GIPS | 4.688 GIPS | 4.703 GIPS | 
| Décompression – Courant nominal | 265.709 GIPS | 289.879 GIPS | 291.363 GIPS | 
| Décompression – Utilisation du processeur résultante | 6236 % | 6205 % | 6133 % | 
| Décompression – Évaluation/utilisation résultante | 4.341 GIPS | 4.649 GIPS | 4.693 GIPS | 
| Décompression – Évaluation résultante | 269.373 GIPS | 288.457 GIPS | 287.823 GIPS | 
| Total – Utilisation totale du processeur | 5751 % | 5681 % | 5674 % | 
| Total – Note totale/Utilisation | 4.491 GIPS | 4.500 GIPS | 5.250 GIPS | 
| Total – Note totale | 257.598 GIPS | 256.385 GIPS | 295.268 GIPS | 
Conclusion
Il s'agit du premier système P-core que nous avons eu entre les mains, et le Supermicro Hyper SuperServer SYS-212HA-TN impressionne par ses processeurs Intel Xeon 6e génération (Granite Rapids), notamment le processeur 128 cœurs 6980P tout P-core, délivrant 3.906 W par cœur. Ce serveur est conçu pour gérer des charges de travail exigeantes telles que la virtualisation, l'inférence IA, l'apprentissage automatique et le cloud computing.
Doté de plusieurs emplacements PCIe 5.0, dont un emplacement AIOM PCIe 3.0 x5.0 compatible OCP 16 pour la mise en réseau, il offre une connectivité évolutive pour les GPU, les cartes complémentaires et d'autres composants d'extension. Ses capacités de stockage haute densité et de traitement de base de données en font un choix judicieux pour les applications gourmandes en stockage.
En analysant les performances, vous pouvez constater le contraste frappant entre les processeurs Intel Sierra Forest E-Core et Intel Granite Rapids P-Core. Si les unités E-core peuvent largement surpasser les unités P-core en termes de nombre de cœurs bruts, les performances monothread ne rivalisent pas avec le même niveau. Dans notre charge de travail y-cruncher, en examinant les performances de calcul brutes, le processeur monosocket 6980P a fait des tours autour des configurations monosocket ou double socket 6780E. Cela donne au Hyper SuperServer SYS-212HA-TN un potentiel incroyable lorsqu'il s'agit de cibler des charges de travail à forte intensité de calcul, même avec un seul processeur.
Avec ses performances et sa flexibilité, le Supermicro Hyper SuperServer SYS-212HA-TN est une excellente solution monoprocesseur pour les entreprises et les centres de données à la recherche d'un serveur hautes performances et évolutif pour gérer des charges de travail complexes. De plus, ce système est plus performant que les systèmes 144E à 6780 cœurs à un ou deux sockets (Sierra Forest), offrant un avantage concurrentiel pour les charges de travail qui peuvent bénéficier de cette puissance supplémentaire.
