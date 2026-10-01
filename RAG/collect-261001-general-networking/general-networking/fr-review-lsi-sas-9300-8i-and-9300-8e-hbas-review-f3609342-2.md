---
id: collect-261001-general-networking/general-networking/fr-review-lsi-sas-9300-8i-and-9300-8e-hbas-review-f3609342-2
title: "fr-review-lsi-sas-9300-8i-and-9300-8e-hbas-review-f3609342"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark"]
source: docs/RAG/collect-261001-general-networking/fr-review-lsi-sas-9300-8i-and-9300-8e-hbas-review-f3609342.md
source_anchor: ""
source_lines: [41, 54]
sha256: 411a23ae2a94eafb0671c861b26afc62dd23a751b5983298a2f9986eb1d56795
---

# fr-review-lsi-sas-9300-8i-and-9300-8e-hbas-review-f3609342

Une fois de plus, le LSI HBA a permis d'importants gains de performances, le disque Hitachi affichant un impressionnant 6,277.8 02 TPS pour prendre la première place. Les Toshiba PX03SN et PX6,193.6SN sont juste derrière avec 6,181.6 XNUMXTPS et XNUMX XNUMXTPS, respectivement.
Les latences moyennes pendant le benchmark SQL Server montrent un écart beaucoup plus important lorsque l'on compare les disques Hitachi et Toshiba avec le PX03SN connaissant la latence moyenne la plus élevée des trois à 111 ms. Ce sont encore de très bons résultats.
Passant à l'analyse de la charge de travail synthétique des disques d'entreprise, plus précisément notre référence 4k 100% lecture et 100% écriture, la plupart des disques ont pu se vanter d'avoir de grands chiffres avec le LSI HBA. Le SSD Hitachi avait les meilleures performances de lecture et d'écriture avec respectivement 149,078 66,367 IOPS et XNUMX XNUMX IOPS.
Conclusion
Les HBA LSI SAS 9300-8e/8i, comme leurs frères (HBA LSI SAS3008), nous ont prouvé qu'ils étaient des produits très performants ; ils ont excellé à presque tous les tests qui leur ont été lancés. Les deux modèles prennent en charge 8 voies (de PCIe 3.0) et incluent des connecteurs mini-SAS HD pour répondre à la norme SAS 3.0 permettant aux serveurs de toute taille de connecter 1024 périphériques SAS ou SATA dans des boîtiers externes/internes. De plus, lorsqu'ils sont connectés aux bons SSD, ces deux HBA LSI suppriment essentiellement le goulot d'étranglement du SAS 6 Gb/s, permettant aux entreprises même les plus gourmandes en ressources d'atteindre leurs objectifs de performances. Cela était assez évident dans les disques testés ci-dessus, car cela permettait aux SSD d'utiliser leur potentiel de 12 Gb/s, atteignant un niveau de performances beaucoup plus élevé par rapport au SAS2, y compris une latence très faible et un débit très élevé. Plus précisément, le Toshiba PX02SS a affiché le meilleur benchmark MySQL que nous ayons enregistré à ce jour, avec plus de 2,150 32 TPS sur 03 threads ; le PX800SN et le SSD02MM ont montré des résultats très similaires et juste là avec le PXXNUMXSS en termes de performances.
Heureusement, ces HBA sont disponibles à l'achat en tant que composant individuel, à un prix raisonnable de 290 $ et 410 $ pour les 9300-8i et 9300-8e respectivement. Les HBA LSI établissent une barre en matière de performances et de fiabilité et valent chaque centime, en particulier pour ceux qui cherchent à obtenir les meilleurs résultats possibles avec leurs SSD basés sur SAS3 de nouvelle génération.
Avantages
- Permet aux SSD SAS3 de fonctionner à leur plein potentiel
- Plate-forme de confiance avec un long historique
- Disponible pour le stockage interne et les JBOD externes
Inconvénients
- Seulement 8 voies SAS natives
Conclusion
Continuant à plaider en faveur de l'adoption du SAS 12 Gb/s grand public dans la sphère de l'entreprise, les HBA LSI SAS 9300-8e/8i ont la capacité d'obtenir les meilleures performances absolues de la prochaine génération de SSD basés sur SAS3.
