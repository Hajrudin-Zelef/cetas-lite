---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-superchassis-846be1c-r1k28b-review-82358c79-1
title: "fr-review-supermicro-superchassis-846be1c-r1k28b-review-82358c79"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "arr", "benchmark", "benchmarks", "intel"]
source: docs/RAG/clean4/fr-review-supermicro-superchassis-846be1c-r1k28b-review-82358c79.md
source_anchor: ""
source_lines: [1, 49]
sha256: eb41b61c139a6d9e0a187add2aa378bdc33d63571dc5fc17f45c65f0d453fba3
---

# fr-review-supermicro-superchassis-846be1c-r1k28b-review-82358c79

Le SuperMicro SuperChassis 846BE1C-R1K28B est un JBOD à 24 baies. Les plateaux de disque sont remplaçables à chaud et conçus pour les disques durs de 3.5 pouces (bien qu'un adaptateur puisse être utilisé pour installer des disques durs ou des SSD de 2.5 pouces). Si l'on devait utiliser des disques de 8 To, tels que les disques He8 de HGST, cela porterait la capacité maximale totale à 192 To. Le SuperChassis peut également être utilisé comme unité centrale et prend en charge les processeurs Intel et AMD simples et doubles, et prend en charge une variété de cartes mères.
Le concept d'étagère de stockage, ou JBOD, est l'un des plus basiques de l'architecture de stockage. Le châssis abrite essentiellement les disques, se connectant à une machine hôte via un câble SAS et un HBA dans l'hôte. Ce type d'arrangement continue d'être populaire lorsque les utilisateurs d'entreprise souhaitent conserver le stockage local sur l'hôte, mais ont peut-être dépassé les baies de lecteur disponibles, ou ont d'autres exigences uniques et n'ont pas besoin d'un SAN complet avec son propres contrôleurs de stockage. En fait, dans l'avenir, nous montrerons ces mêmes configurations de disques durs associées à des solutions de mise en cache, pour montrer comment le flash et les logiciels peuvent bénéficier à de larges baies de disques durs dans un environnement d'entreprise. Les cas d'utilisation de JBOD continuent de se développer avec de nouvelles technologies et une puissante puissance de calcul côté hôte.
Le SuperMicro SuperChassis 846BE1C-R1K28B est conçu en tenant compte du budget et de la rentabilité. Ce DAS est conçu pour être facile à utiliser et à entretenir. Les disques sont remplaçables à chaud couplés à des alimentations redondantes ; cela rend le SuperChassis susceptible de causer de graves problèmes de temps d'arrêt. Le JBOD est destiné aux entreprises qui ont besoin d'une méthode peu coûteuse pour ajouter du stockage à leur système existant.
Le SuperMicro SuperChassis 846BE1C-R1K28B a un prix catalogue de 1,400 3 $ (non rempli) et est livré avec une garantie qui couvre la main-d'œuvre de 1 ans, les pièces d'un an et le remplacement avancé de 3 mois.
Spécifications SuperMicro SuperChassis 846BE1C-R1K28B :
- Facteur de forme : 4U
- Prise en charge du processeur : processeurs Intel et AMD simples et doubles
- Capacité de stockage:
  - 24 disques durs SAS ou SATA remplaçables à chaud de 3.5 pouces
  - Des plateaux adaptateurs 2.5" en option peuvent être ajoutés
- Emplacements d'extension : 7 x pleine hauteur et pleine longueur
- Extenseur : extenseurs LSI SAS2
- L'environnement:
  - Température de fonctionnement: 5 ° C ~ 35 ° C (41 ° F ~ 95 ° F)
  - Température hors fonctionnement : -40 °C ~ 60 °C (-40 °F ~ 140 °F)
  - Humidité relative de fonctionnement : 8 % ~ 90 % (sans condensation)
  - Humidité relative hors fonctionnement : 5 % à 95 % (sans condensation)
- Climatisation
  - 3 ventilateurs de refroidissement PVM remplaçables à chaud de 80 mm
  - 2 ventilateurs PWM d'échappement arrière remplaçables à chaud de 80 mm
- Alimentation : Alimentations numériques redondantes à haut rendement de 1280 W avec PMBus 1.2
- Entrée AC:
  - Sortie 1000W @ 100-140V, 8-12A, 50-60Hz
  - Sortie 1280W @ 180-240V, 6-8A, 50-60Hz
- Entrée DC:
  - 1000W: +12V/83A; +5Vsb/4A
  - 1280W: +12V/106.7A, +5Vsb/4A
- Dimensions (LxlxH) : 26.5"x17.2"x7"
- Poids (sans disques) : 75 lbs.
Conception et construction
Le SuperMicro SuperChassis 846BE1C-R1K28B JBOD 24 baies est un stockage à connexion directe 4U. À l'avant de la plate-forme se trouvent les baies de lecteur remplaçables à chaud, quatre rangées de six baies. Pour éteindre un lecteur, appuyez simplement sur le bouton marron, une petite poignée apparaît, et le lecteur doit simplement être retiré et un autre remplacé. En bas à gauche se trouvent les voyants LED. Le côté inférieur droit porte la marque SuperMicro. Au sommet de chaque côté se trouvent des poignées d'oreille en métal pour ranger et défaire l'appareil.
En faisant le tour à l'arrière de l'appareil, il y a deux alimentations redondantes remplaçables à chaud sur le côté gauche. Le centre est dominé par deux ventilateurs de refroidissement. Sur le côté droit se trouvent 4 ports mini-SAS HD et un port RJ45.
Contexte des tests et comparables
Pour tester le SuperMicro SuperChassis 846BE1C-R1K28B JBOD à 24 baies, nous avions l'habitude de tester 24 disques HGST Ultrastar Helium He8 8 To dans un réglage miroir. Nous avons également effectué les mêmes tests avec 20 HGST Ultrastar Helium He8 8 To et 4 HGST Ultrastar SSD800MR SAS3 500 Go pour la hiérarchisation SSD.
Disques utilisés dans le JBOD pour les tests :
Direction
La gestion n'est pas la première chose qui vient à l'esprit quand on pense aux JBOD. Cependant, pour notre test de serveur SQL, nous avons configuré un grand pool de stockage miroir de disques hélium HGST he8 8 To ainsi qu'un pull miroir avec cache SSD à l'aide de HGST Ultrastar SSD800MM SAS3 400 Go.
Sur l'écran principal des espaces de stockage, les volumes et les pools de stockage doivent être configurés.
Les utilisateurs doivent sélectionner les disques physiques à utiliser pour le pool de stockage, dans ce cas les disques He8.
Une fois terminé, les utilisateurs doivent nommer le disque virtuel.
Et sélectionnez la taille du disque virtuel en dehors de la capacité existante.
Une fois que l'emplacement, le nom et la taille ont été décidés, les utilisateurs doivent confirmer les paramètres.
Et confirmez les paramètres de hiérarchisation, la disposition Miroir (RAID10) dans ce cas.
Analyse des performances des applications
Le protocole de test OLTP Microsoft SQL Server de StorageReview utilise la version préliminaire actuelle du benchmark TPC-C (Transaction Processing Performance Council's Benchmark C), un benchmark de traitement transactionnel en ligne qui simule les activités rencontrées dans des environnements applicatifs complexes. Le benchmark TPC-C est plus représentatif que les benchmarks de performance synthétiques des performances réelles, permettant d'évaluer avec précision les points forts et les goulots d'étranglement de l'infrastructure de stockage dans les environnements de bases de données. Notre protocole SQL Server pour cette analyse utilise une base de données SQL Server de 685 Go (échelle 3 000) et mesure les performances transactionnelles et la latence sous une charge de 30 000 utilisateurs virtuels.
En utilisant Storage Spaces configuré Mirror RAID, nous avons vu le JBOD abandonner 3,874.65 6,307.92 TPS et en regardant la hiérarchisation SSD, le nombre a presque doublé avec XNUMX XNUMX TPS
La latence moyenne nous montre une différence plus fantastique. Lorsqu'il n'était rempli que de disques à l'hélium He8, la latence était de 2,990 6 ms. Cependant, lorsque nous ajoutons le cache SSD, la latence est tombée à XNUMX ms presque cinq cents fois plus vite.
Conclusion
Le SuperMicro SuperChassis 846BE1C-R1K28B JBOD à 24 baies est une unité de stockage à connexion directe 4U offrant rentabilité et flexibilité. Avec 24 baies de disques, pouvant accueillir des disques 3.5" ou avec un adaptateur 2.5" et prenant en charge SAS et SATA, le JBOD peut avoir une capacité maximale de 192 To avec des disques durs de 8 To. Comme la grande majorité de ce que propose SuperMicro, le 846BE1C JBOD offre une grande flexibilité quant aux processeurs et aux cartes mères qu'il prend en charge si les utilisateurs choisissent d'utiliser le 846BE1C comme unité centrale.
