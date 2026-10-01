---
id: collect-261001-general-networking/general-networking/fr-review-adaptec-series-8-raid-controllers-review-1f1b7292-2
title: "fr-review-adaptec-series-8-raid-controllers-review-1f1b7292"
domain: general-networking
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-general-networking/fr-review-adaptec-series-8-raid-controllers-review-1f1b7292.md
source_anchor: ""
source_lines: [77, 108]
sha256: 16f6c144cbd29050a1f3363dd8cd2d06cde7f385170d92cd13c78cb3e79aac9a
---

# fr-review-adaptec-series-8-raid-controllers-review-1f1b7292

Nous avons installé trois cartes Adaptec 81605ZQ pour se connecter à quatre pools de disques SSD. Dans notre configuration haute performance pour les tests, nous avons inclus 16 SSD Hitachi SSD800MM 400 Go, 8 SSD Micron M500 960 Go, ainsi que 2 SSD Seagate SSD 600 480 Go. Nous avons réparti les groupes jusqu'à 8 SSD SSD800MM sur deux cartes Adaptec 81605ZQ, la dernière étant connectée aux SSD 8 Micron et 2 SSD Seagate. Dans notre configuration orientée vers des performances optimales, nous avions deux pools RAID10 SAS3 composés des SSD Hitachi, un pool RAID50 du Micron SSS et un dernier pool RAID1 des SSD Seagate pour le démarrage.
Pour mesurer les performances du 81605ZQ dans un environnement réel, nous avons créé des partages SMB3, iSCSI et NFS dans Windows Server 2012 R2 et y avons accédé via notre infrastructure 10GbE de bout en bout. Pour les tests HGST SSD800MM, nous avions deux partages ou LUN de 25 Go sur chaque pool de stockage RAID10 (quatre partages ou LUN au total) ou quatre partages ou LUN sur le seul pool de stockage RAID50 Micron M500. Pour ces quatre connexions, nous avons utilisé chacune un port 10GbE dédié pour leur donner autant de bande passante qu'elles en auraient besoin.
Analyse synthétique de la charge de travail d'entreprise
Notre processus de référence d'entreprise préconditionne chaque périphérique de stockage dans un état stable avec la même charge de travail avec laquelle le périphérique sera testé sous une charge lourde de 16 threads avec une file d'attente exceptionnelle de 16 par thread, puis testé à des intervalles définis dans plusieurs threads/profondeur de file d'attente profils pour afficher les performances en cas d'utilisation légère et intensive.
Tests de préconditionnement et d'état stable primaire :
- Débit (agrégat IOPS lecture + écriture)
- Latence moyenne (latence de lecture + écriture moyennée ensemble)
- Latence maximale (latence maximale de lecture ou d'écriture)
- Écart-type de latence (écart-type de lecture + écriture moyenné ensemble)
Notre analyse synthétique de la charge de travail d'entreprise comprend des profils basés sur des tâches réelles. Ces profils ont été développés pour faciliter la comparaison avec nos références passées ainsi qu'avec des valeurs largement publiées telles que 8k 70/30 qui est couramment utilisée pour les produits d'entreprise.
- 8k 70/30
  - 70 % de lecture, 30 % d'écriture
  - 100% 8K
Nous nous sommes concentrés sur notre charge de travail 8k 70/30, car dans notre laboratoire de test d'entreprise virtualisé avec plus de 100 machines virtuelles actives sur différents tests, les performances de la charge de travail aléatoire sont essentielles. Bien que notre environnement VMware exploite les chiffres de performance iSCSI, SMB3 et NFS ont également été inclus pour voir les performances de Windows Server 2012 R2 pour chaque protocole.
Lors du premier test mesurant le débit, les seize SSD Hitachi SSD800MM SAS3 dans deux pools RAID10 accessibles via SMB3 ont offert les meilleures performances globales à chaque profondeur, atteignant plus de 160,000 78 IOPS, battant facilement ses partages iSCSI ou NFS qui dépassaient respectivement 37 500 IOPS et 50 3 IOPS. Le pool Micron M55 RAID52 offrait des performances SMB11.4 dépassant XNUMX XNUMX IOPS, avec iSCSI et NFS à la traîne avec des chiffres supérieurs de XNUMX XNUMX IOPS et XNUMX XNUMX IOPS respectivement.
La configuration Hitachi SSD800MM 2xRAID10 sur SMB3 offrait la latence moyenne globale la plus faible, passant de 0.33 ms à 2T/2Q à 7.05 ms à 16T/16Q. Notre pool de stockage Micron M500 en RAID50 est passé de 0.66 ms à 2T/2Q à 19.24 ms à 16T/16Q.
La plupart des configurations de disque ont obtenu d'excellents résultats lors des tests Max Latency, en particulier le disque Micron, qui n'a montré pratiquement aucun pic. Cependant, le Hitachi SSD800MM 2xRAID10 via le conteneur iSCSI de Windows Server 2012 a atteint un pic assez élevé avec une latence maximale atteignant plus de 10,500 XNUMX ms, ce qui peut être inférieur à la limite supérieure que Windows peut prendre en charge via ce protocole.
Parmi les SSD, le pool de stockage Micron M500 RAID50 sur SMB3 a battu de peu le Hitachi SSD800MM SMB3 pour offrir l'écart type le plus étroit au niveau 16T/16Q. Le pool de stockage HGST sur iSCSI ainsi que le pool de stockage Micron sur NFS ont beaucoup augmenté au cours de ces tests et ont eu un écart type beaucoup plus élevé à la fin.
Conclusion
L'Adaptec Series 8 81605ZQ a beaucoup à offrir en tant que carte RAID de premier plan, avec 16 ports SAS 12 Gbit/s ainsi qu'une protection intégrée contre les pannes de courant sans avoir à ajouter de carte fille supplémentaire. En ce qui concerne les prix catalogue, le prix de la famille de cartes RAID de la série 8 n'est pas trop différent de celui des cartes de la série 7 qu'elles remplacent. Dans le cas du 16ZQ à 81605 ports (1,065 71605 $), il est en fait moins cher que le prix catalogue du 1,100Q (XNUMX XNUMX $), ce qui en fera une option intéressante pour ceux qui envisagent d'acheter de nouveaux équipements pouvant prendre en charge les dernières normes.
Adaptec montre à nouveau aux intégrateurs de systèmes, aux OEM et aux administrateurs informatiques autonomes que les cartes RAID de la série 8 ont beaucoup à offrir. Avec une densité de ports deux fois supérieure à celle des offres concurrentes, les fournisseurs de serveurs peuvent sérieusement envisager de maximiser le débit PCIe Gen3 avec des SSD SAS haut de gamme avec des versions moins complexes qui ignorent les extensions. Adaptec continue également à faire évoluer ses offres logicielles avec max Cache Plus, qui permet à la fois la mise en cache et la hiérarchisation.
Lors de nos tests dans un environnement Windows Server 2012 R2 hébergeant un stockage partagé dans notre environnement virtualisé, l'Adaptec Series 8 81605ZQ s'est plutôt bien comporté, poussant dans certains cas Windows à ses limites pour différents protocoles de stockage. En exploitant les performances de seize HGST SSD800MM dans deux pools RAID10 exploitant autant de cartes 81605ZQ, nous avons constaté des performances 8k 70/30 supérieures à 160k IOPS sur SMB3. Sur notre troisième 81605ZQ utilisé pour contrôler notre disque de démarrage RAID1 et un pool RAID50 de SSD Micron M500, nous avons vu les performances SMB3 dépasser 55 8 IOPS dans ce même test d'E/S aléatoires. Inutile de dire que le laboratoire est enthousiaste à l'idée de continuer à déployer les cartes de la série XNUMX dans une variété de projets à venir qui incluent un stockage haute performance.
Avantages
- Des performances incroyables
- Prix raisonnable pour le marché des entreprises
- Jusqu'à 16 ports SAS de 12 Go/s
Inconvénients
- Pas de modèle à 24 ports comme la série 7
Conclusion
La gamme Adaptec Series 8 convient parfaitement aux entreprises qui cherchent à améliorer considérablement les performances de leurs applications gourmandes en ressources ou de leurs environnements de serveurs denses à un niveau de pointe. Le 8160ZQ en particulier offre une excellente protection des données avec son module flash intégré, qui élimine le besoin d'une carte fille supplémentaire et les 16 ports offrent une énorme flexibilité dans le système.
Cartes RAID Adaptec série 8 sur Amazon
Discutez de cet avis
