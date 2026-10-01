---
id: collect-261001-general-networking/general-networking/fr-review-broadcom-megaraid-9670w-16i-raid-card-review-f4ec557a-2
title: "fr-review-broadcom-megaraid-9670w-16i-raid-card-review-f4ec557a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["gpu"]
source: docs/RAG/collect-261001-general-networking/fr-review-broadcom-megaraid-9670w-16i-raid-card-review-f4ec557a.md
source_anchor: ""
source_lines: [33, 38]
sha256: 62211f59d137bce534385cf7aca0885299b538b5f869c4172e1816e7bf893b80
---

# fr-review-broadcom-megaraid-9670w-16i-raid-card-review-f4ec557a

Nous avons non seulement évalué les performances globales de chaque mode grâce à des mesures de performances ponctuelles, qui comprenaient également les performances de la carte RAID lors d'une opération de reconstruction, mais nous avons également effectué des tests pour déterminer le temps total nécessaire à la reconstruction. Ici, dans RAID10, supprimer un SSD de 6.4 To du groupe RAID et le rajouter a pris 60.7 minutes pour RAID10 avec une vitesse de reconstruction de 10.4 Min/To. Le groupe RAID5 a pris 82.3 minutes avec une vitesse de 14.1 Min/TB.
Réflexions finales
Pour être honnête, nous sommes entrés dans cette revue avec des têtes légèrement inclinées et un sourcil levé. Nous n'avons pas entendu de présentation de carte RAID pour les SSD NVMe depuis un moment, en dehors de la classe émergente de solutions conçues autour des GPU. Nous avons donc dû poser la question fondamentale, le RAID matériel peut-il même être une chose pour les SSD NVMe ?
La réponse est clairement oui. Les performances de PCIe Gen4 permettent à la carte RAID MegaRAID 9670W-16i de suivre le rythme des SSD modernes sur une variété de charges de travail. Oui, certains domaines comme la bande passante seront limités avec moins de voies PCIe, mais encore une fois, la plupart des environnements de production ne se situent pas à ces niveaux.
En bande passante maximale, nous avons vu le MegaRAID 9670W-16i aller jusqu'à la limite x16 PCIe Gen4 de 28 Go/s en lecture et offrir jusqu'à 13 Go/s en RAID5 en bande passante en écriture. Du côté du débit, les performances de lecture aléatoire 4K ont atteint 7 millions d'IOPS avec une écriture allant de 1 à 2.1 millions d'IOPS entre RAID5 et RAID10. Pour les déploiements cherchant à consolider le flash dans des volumes plus importants ou à contourner les systèmes qui ne prennent pas en charge le RAID logiciel, le MegaRAID 9670W a beaucoup à offrir.
Si vous aimez les adaptateurs de stockage, vous êtes sur le point d'obtenir davantage de ce type de couverture. Nous explorons déjà les serveurs de dernière génération, comme le Dell PowerEdge R760, qui propose une configuration de carte RAID double basée sur le même silicium que cette carte. Dans le boîtier R760, Dell attache 8 SSD NVMe à chaque carte, ce qui nous offre une solution d'entreprise plus robuste que celle que nous avons testée ici pour validation. Il y a donc beaucoup plus à venir maintenant qu'il semble que les cartes RAID soient de retour au menu pour les serveurs avec des SSD NVMe.
