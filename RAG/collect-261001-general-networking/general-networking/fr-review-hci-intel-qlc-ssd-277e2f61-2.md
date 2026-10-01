---
id: collect-261001-general-networking/general-networking/fr-review-hci-intel-qlc-ssd-277e2f61-2
title: "fr-review-hci-intel-qlc-ssd-277e2f61"
domain: general-networking
role: reference
task: reference
actors: ["Intel", "Microsoft", "Samsung"]
dates: []
keywords: ["intel", "benchmark", "nand"]
source: docs/RAG/collect-261001-general-networking/fr-review-hci-intel-qlc-ssd-277e2f61.md
source_anchor: ""
source_lines: [17, 49]
sha256: 6405d99619f651b42d992017c060c4ff04560f38073923b90068365ef2746083
---

# fr-review-hci-intel-qlc-ssd-277e2f61

Nous avons récemment publié un article sur les clusters à deux nœuds (2NC) de Microsoft Azure Stack HCI . Vous trouverez ci-dessous un résumé de cet article. Nous avons constaté qu'un cluster 2NC pouvait, dans de nombreux cas d'utilisation, fournir la résilience nécessaire à une organisation et qu'il était moins complexe et moins coûteux qu'un cluster traditionnel à trois ou quatre nœuds. DataON a été parmi les premiers fournisseurs à reconnaître la valeur des clusters 2NC et à en adopter l'intégration. Cependant, les clusters 2NC ne sont pas une nouveauté pour DataON : en septembre 2017, l'entreprise a annoncé la commercialisation des deux premiers systèmes Kepler-47 HCI pour Windows Server 2016 Storage Spaces Direct (désormais Azure Stack HCI).
L'implémentation 2NC de DataON prend en charge une panne de disque et une panne de serveur en même temps. Pour ce faire, il utilise RAID 5 + 1 pour assurer la résilience de la parité et la refléter sur l'autre serveur. Microsoft appelle cette capacité « résilience imbriquée » et a ajouté cette capacité à Storage Spaces Direct dans Windows Server 2019. Encore une fois, 2NC n'est pas le bon choix technologique pour tout le monde, mais ils peuvent fournir une solution fiable et rentable à de nombreuses organisations.
Construire et concevoir
Le cluster Azure Stack HCI avec lequel nous travaillons ici a été construit sur la plate-forme NVMe 224 % flash DataON HCI-2. Ces serveurs avaient une taille de 24U avec 1 baies NVMe à l'avant, offrant de nombreuses extensions à l'arrière pour les composants basés sur PCIe. L'étiquetage était élevé contrairement aux caddies de lecteur noir mat, ce qui permettait de repérer facilement des lecteurs spécifiques en cas de remplacement nécessaire. Tout était étiqueté, ce qui n'est pas si rare, mais l'étendue de l'étiquetage était extraordinaire. Notre déploiement avait chaque nœud étiqueté (2 et XNUMX), ainsi que plusieurs autres éléments, facilitant le déploiement et la gestion des systèmes DataON dans le centre de données.
Les nœuds de ce test comprenaient deux processeurs Intel® Xeon® Scalable Gold 6248 de 2e génération à 2.5 GHz, 20 cœurs et 28 Mo de cache, ainsi que huit modules RDIMM ECC enregistrés Samsung de 32 Go DDR4 à 2933 MHz (256 Go au total par nœud) et deux disques de démarrage Intel S4510 SATA M.2 de 480 Go.
Pour le stockage, chaque nœud était livré avec quatre disques Intel Optane SSD DC P4800X NVMe 750 Go 2.5 pouces (utilisés pour la mise en cache) et quatre disques Intel SSD D5-P4326 15.36 To 2.5 pouces QLC (niveau de stockage de capacité).
Les nœuds ont été connectés les uns aux autres via des cartes Mellanox ConnectX-4 EN double port QSFP28 40/56 GbE à l'aide de câbles en cuivre passifs 3M Mellanox LinkX ETH 40GbE, 40Gb/s, QSFP.
De toute évidence, DataON a passé beaucoup de temps et réfléchi à la configuration et à la sélection des composants de ce système pour équilibrer les performances et les coûts. Nous étions très intéressés de voir comment les SSD Intel SSD D5-P4326 fonctionneraient en tant que niveau de stockage. En combinant les SSD Intel Optane et les SSD Intel QLC 3D NAND, les SSD D5-P4326 devraient fournir un stockage flash hautes performances et économique, qui était autrefois le domaine des disques durs lents mais volumineux.
Dans le laboratoire StorageReview, nous avons déployé les deux nœuds de stockage et les commutateurs comme illustré ci-dessous.
Tests
Pour avoir une idée de la façon dont un petit cluster comme celui-ci peut fonctionner dans un cas d'utilisation périphérique, nous avons configuré plusieurs tests Microsoft SQL Server. L'objectif était d'examiner les performances complètes du cluster pour s'assurer que DataON pouvait utiliser correctement la technologie Intel Optane et les SSD Intel QLC. Deuxièmement, nous voulions examiner les capacités d'un seul nœud, pour avoir une idée de la façon dont cette solution gère la perte d'un nœud, que ce soit pour les mises à jour planifiées ou en cas de panne plus grave.
Notre plan de test s'est appuyé sur Benchmark Factory de Quest en utilisant le profil TPC-C comme générateur de charge pour les machines virtuelles SQL Server que nous avons déployées. Nous avons configuré huit machines virtuelles (quatre par nœud), ce qui offrait un bon équilibre entre l'activité CPU et disque pour le cluster. Les générateurs de charge de travail étaient hébergés sur un système en dehors de cet environnement et connectés à ce cluster via un réseau 10GbE.
Configuration des tests SQL Server (par machine virtuelle)
- Windows Server 2019
- Empreinte de stockage : 800 Go alloués, 620 Go utilisés
- 8 vCPU
- 60 Go de RAM (55 Go en configuration en mode échec)
- SQL Server 2019
  - Taille de la base de données : échelle 1,500 XNUMX
  - Charge de client virtuel : 15,000 XNUMX
  - Mémoire tampon : 48 Go
- Durée du test : 3 heures
  - 15 minutes de préconditionnement
  - Période d'échantillonnage de 45 minutes
Dans nos tests, nous nous sommes concentrés sur les performances de latence, le niveau de performance des transactions restant constant avec Benchmark Factory.
Avec une charge de 4 VM au total (2 par nœud), nous avons mesuré une latence moyenne de 2.5 ms avec une charge de transaction agrégée de 12,649 XNUMXTPS.
En augmentant la charge à 6 VM, la latence moyenne a légèrement augmenté à 4 ms avec une charge de transaction globale de 18,967 XNUMX TPS.
Au pic de charge de 8 VM (4 par nœud), la latence a atteint une moyenne de 6.5 ms, avec une charge de transaction totale de 25,277 XNUMX.
Tout au long de ces tests, nous avons clairement vu l'avantage d'avoir les SSD Optane dans ce mix. Ils ont pris le gros des écritures, libérant les SSD QLC pour des lectures réactives en tant que niveau de capacité haute vitesse. Même si nous avons doublé la charge de travail à huit machines virtuelles SQL Server frappant ce cluster HCI, la latence n'a augmenté que légèrement, ce qui montre que cette configuration est bien adaptée aux charges de travail qui peuvent éclater de temps à autre.
Bien que les performances dans un environnement pleinement opérationnel soient importantes, une autre considération est la manière dont les charges de travail fonctionneront si un nœud du cluster se déconnecte ou si des charges de travail doivent être migrées pour la maintenance du système. Pour tester ce scénario, nous avons conservé notre charge complète de 8 VM et les avons migrées vers un seul nœud. Dans cette configuration, nous avons mesuré une latence moyenne de seulement 4.5 ms, ce qui était mieux que les deux nœuds en ligne. Cela provient en partie de la suppression de la surcharge de stockage dans le fonctionnement à nœud unique.
Conclusion
Pour ce projet, nous avons exécuté une série de tests SQL sur le système pour illustrer les charges de travail de performance que l'on trouve couramment dans les cas d'utilisation Edge et SMB. Notre objectif était de comprendre avec quelle efficacité Microsoft Azure Stack HCI dans ce cluster DataON a pu tirer parti du matériel pour obtenir les résultats souhaités. Plus précisément, cela signifie fournir une solution qui offre une combinaison rare de performances et de valeur.
Nous pouvons confirmer grâce à nos tests que la sélection de composants de DataON a effectivement réussi à créer une solution Azure Stack HCI SDS rentable qui fonctionne extrêmement bien. Cela est en partie dû à leur choix d'utiliser le SSD Intel D5-P4326 pour le stockage de capacité, qui tire efficacement parti des SSD Intel Optane pour la hiérarchisation.
