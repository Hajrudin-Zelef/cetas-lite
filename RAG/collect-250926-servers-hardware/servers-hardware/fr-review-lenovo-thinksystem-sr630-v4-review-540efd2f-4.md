---
id: collect-250926-servers-hardware/servers-hardware/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f-4
title: "fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["compute", "intel", "valuation"]
source: docs/RAG/clean4/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f.md
source_anchor: ""
source_lines: [130, 156]
sha256: b18753f258ce7f4b232ece53eeb5e9ed69221579c2bf8a2a29b9af8abd66c080
---

# fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f

L'outil de référence de mémoire intégré de l'utilitaire populaire 7-Zip mesure les performances du processeur et de la mémoire d'un système pendant les tâches de compression et de décompression, indiquant dans quelle mesure le système peut gérer les opérations gourmandes en données.
En ce qui concerne les tâches de compression, le système Supermicro a obtenu des résultats légèrement supérieurs à ceux du SR630 en termes d'utilisation du processeur et de notes obtenues, avec une note de compression totale de 245.823 GIPS et celle de Lenovo de 224.313 GIPS. Cela suggère un léger avantage lors de la gestion de charges de travail de compression fortement threadées. Les tâches de décompression de Lenovo ont montré une note résultante plus élevée de 288.457 GIPS, tandis que la note de Supermicro a indiqué 269.373 GIPS. Cela se traduit par de meilleures performances pour les charges de travail qui nécessitent la lecture et l'extraction de données. Le serveur Intel Ice Lake a démontré des performances équilibrées, avec une note de compression de 235.437 GIPS et une note de décompression de 253.692 GIPS.
Les performances globales des deux systèmes sont quasiment identiques, avec des notes totales de 257.598 GIPS pour le Supermicro à socket unique et de 256.385 GIPS pour le Lenovo à double socket. L'ancien Xeon Ice Lake a atteint 244.565 GIPS. Les trois systèmes sont très performants.
| Compression à 7 zips | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | Lenovo ThinkSystem SR630 V4 (Intel Xeon 6780E, 512 Go) | Serveur Intel Ice Lake (2 x Intel Xeon 8380, 512 Go) | 
| Compression |  |  |  | 
| Utilisation actuelle du processeur | 5287 % | 5064 % | 5835 % | 
| Courant nominal/utilisation | 4.647 GIPS | 4.341 GIPS | 4.030 GIPS | 
| Courant | 245.699 GIPS | 219.840 GIPS | 235.143 GIPS | 
| Utilisation résultante du processeur | 5296 % | 5156 % | 5839 | 
| Évaluation/utilisation résultante | 4.642 GIPS | 4.350 GIPS | 4.032 GIPS | 
| Note résultante | 245.823 GIPS | 224.313 GIPS | 235.437 GIPS | 
| Décompression |  |  |  | 
| Utilisation actuelle du processeur | 6236 % | 6184 % | 6230 % | 
| Courant nominal/utilisation | 4.261 GIPS | 4.688 GIPS | 4.050 GIPS | 
| Courant | 265.709 GIPS | 289.879 GIPS | 252.326 GIPS | 
| Utilisation résultante du processeur | 6236 % | 6205 % | 6245 % | 
| Évaluation/utilisation résultante | 4.341 GIPS | 4.649 GIPS | 4.062 GIPS | 
| Note résultante | 269.373 GIPS | 288.457 GIPS | 253.692 GIPS | 
| Note totale |  |  |  | 
| Utilisation totale du processeur | 5751 % | 5681 % | 6042 % | 
| Note totale/utilisation | 4.491 GIPS | 4.500 GIPS | 4.047 GIPS | 
| Note totale | 257.598 GIPS | 256.385 GIPS | 244.565 GIPS | 
Conclusion
Le Lenovo ThinkSystem SR630 V4 est un serveur rack 1U polyvalent qui, bien qu'il ne représente pas une mise à niveau significative par rapport à son prédécesseur, constitue néanmoins une avancée fiable dans l'évolution des principaux systèmes d'entreprise de Lenovo. Avec la prise en charge des processeurs Intel Xeon série 6700E, il double la densité de cœurs par rapport à son prédécesseur, offrant une meilleure évolutivité pour les charges de travail exigeantes. Les vitesses de mémoire DDR5 améliorées allant jusqu'à 6400 MHz et la prise en charge prévue des technologies émergentes telles que Compute Express Link (CXL) et MCRDIMM montrent que Lenovo a l'intention de maintenir cette gamme de serveurs équipée pour les demandes futures.
L'évolution du SR630 V4 vers le stockage NVMe et les configurations de disques flexibles signifie également qu'il privilégie les performances et l'évolutivité dans les charges de travail gourmandes en données. De plus, les deux emplacements OCP 3.0 prenant en charge PCIe 5.0 permettent des options de mise en réseau avancées, notamment des composants remplaçables à chaud et des systèmes de refroidissement améliorés, simplifiant la maintenance et l'efficacité opérationnelle.
En ce qui concerne ses performances lors de nos tests de référence, le SR630 V4 a obtenu de bons résultats, excellant dans les charges de travail multithread telles que Cinebench R23 et Y-Cruncher. La façon dont les nouveaux processeurs Intel Sierra Forest E-core Xeon s'intègrent dans cet espace est destinée aux entreprises qui cherchent à mettre à niveau leurs plateformes existantes pour une plus grande efficacité. Toutes les charges de travail ne nécessitent pas de performances supérieures, mais elles pourraient bénéficier de nouvelles améliorations telles que la densité et une consommation d'énergie réduite.
Cela dit, même si l'efficacité des processeurs E-core ne répond pas entièrement aux exigences des charges de travail nécessitant des performances élevées sur un seul thread, Lenovo a prévu la prise en charge des processeurs P-core. Des configurations de disques supplémentaires sont également disponibles, prenant en charge une large gamme d'applications d'entreprise.
