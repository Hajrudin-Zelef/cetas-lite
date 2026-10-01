---
id: collect-261001-general-networking/general-networking/fr-review-enabling-gen5-ssd-evaluations-with-serial-cables-57af70e4
title: "fr-review-enabling-gen5-ssd-evaluations-with-serial-cables-57af70e4"
domain: general-networking
role: reference
task: reference
actors: ["Broadcom", "Samsung"]
dates: []
keywords: ["valuation"]
source: docs/RAG/collect-261001-general-networking/fr-review-enabling-gen5-ssd-evaluations-with-serial-cables-57af70e4.md
source_anchor: ""
source_lines: [1, 21]
sha256: 34806ab4a91101b4cb973da18bf9d73a00e24bfe778bec78b05caa77b2737127
---

# fr-review-enabling-gen5-ssd-evaluations-with-serial-cables-57af70e4

Les tests et l'examen des disques SSD présentent un nouvel ensemble de complexités à chaque mise à jour de l'interface ou changement de facteur de forme. Ces défis n'ont jamais été aussi répandus qu'aujourd'hui, alors que l'industrie passe non seulement aux SSD Gen5, mais également à l'ensemble de facteurs de forme SSD le plus diversifié jamais créé.
Dans le domaine des SSD Gen4, nous venons de voir les SSD M.2 et U.2. Les SSD U.2 apparaissaient généralement toujours dans un boîtier de 15 mm, mais il y en avait de temps en temps, notamment Samsung, qui utilisait le boîtier de 7 mm. En ce qui concerne la diversité des facteurs de forme, c'était à peu près tout. Pour être honnête, nous avons examiné un ensemble de SSD E1.S EDSFF, mais cet examen était une valeur aberrante. Avec l’EDSFF qui prend désormais de l’ampleur, non seulement auprès des hyperscalers mais auprès de tous les fournisseurs de serveurs, la donne évolue, tout comme l’équipement de test.
C'est là qu'intervient Serial Cables, qui fournit l'infrastructure essentielle à notre laboratoire pour l'évaluation des SSD de 5e génération. Notre plateforme de test principale est un serveur Dell PowerEdge R760 équipé de baies U.2 en façade. Or, Dell ne prend en charge que les disques de 4e génération avec son fond de panier U.2, et réserve la 5e génération au fond de panier E3.S. On perçoit déjà les difficultés de standardisation de notre banc d'essai pour les SSD. Ce problème n'est pas propre à Dell : Lenovo, HPE, Supermicro et d'autres fabricants font des choix de conception qui rendent quasiment impossible l'utilisation d'une plateforme de test unique pour tous les SSD.
Ce que Serial Cables a fait, c'est créer un JBOF Gen5 avec des plateaux modulaires qui peuvent s'adapter à n'importe quel facteur de forme. L'unité à 8 baies est associée à une carte hôte, un câblage et une série de « cartes à palettes » pour y monter les disques. Et à l'exception de quelques formes de disque étranges comme les énormes SSD E.1L, les plateaux peuvent accueillir la plupart des disques. ce que nous devons tester, très bien. Cela signifie qu’avec un seul serveur moderne, nous pouvons tester n’importe quel SSD Gen5 exactement de la même manière.
Câbles série Équipement Gen5
Comme indiqué, cette suite de tests Gen5 comprend plusieurs éléments d'équipement de câbles série :
- Carte hôte PCIe Gen5 x16 MCIO
- Gen5 PCIO 8 baies E3 Passif JBOF
- Carte à palette Gen5 PCIe U.2
- Carte à palettes Gen5 PCIe M.2
Carte hôte PCIe Gen5 x16 MCIO
La carte hôte est livrée avec un commutateur Broadcom Atlas2 PCIe. Il dispose de 4 prises MCIO offrant chacune une connexion x4 pour connecter directement divers périphériques PCIe. Nous utilisons cette carte à l'intérieur du Dell PowerEdge R760 pour nous connecter en externe aux câbles série JBOF. La carte coûte 3,995.00 XNUMX $.
La carte elle-même a une empreinte FHHL, permettant une large compatibilité avec la plupart des serveurs. Pour les systèmes qui ne prennent en charge que les cartes demi-hauteur, ils proposent une offre similaire dans ce facteur de forme. Du côté du système, aucun pilote n'est requis, ce qui rend l'installation simple quel que soit le système d'exploitation.
PCIe Gen5 8 baies E3 Passif JBOF
Le cœur de ce système est la plate-forme de test universelle Serial Cables Gen5 PCIe 8 baies E3 Passive JBOD. Ce JBOF prend en charge les facteurs de forme E3, U.2, U.3 et M.2, avec un fond de panier EDSFF commun. Les SSD Gen5 sont plus compliqués que les générations précédentes de SSD en raison de la grande quantité de formes et de tailles dans lesquelles ils se présentent. Heureusement, ce JBOF Gen5 est conçu pour résoudre ce problème car il peut prendre en charge de nombreux types de disques différents.
Le JBOF dispose d'un total de huit baies, ce qui signifie que nous pouvons tester des groupes plus importants de SSD si nous le souhaitons, ou des multiples sans avoir à échanger les disques. Une fonction de gestion hors bande est intégrée au JBOD pour les opérations de base telles que l'activation ou la désactivation des emplacements, ainsi que les fonctions de diagnostic. Cet appareil a un prix catalogue de 2,995.00 XNUMX $.
Cartes à palettes PCIe Gen5 U.2, U.3 et M.2
Ces cartes paddle sont compatibles avec le JBOF et coûtent 75.00 $. Nous exploitons actuellement les cartes à palettes U.2/U.3, E3DSFF et M.2 dans le JBOD pour les évaluations PCIe Gen5 d'entreprise et grand public.
Réflexions finales
Dans l'ensemble, la carte hôte Serial Cables PCIe Gen5 et le combo JBOF sont devenus un composant intégral fantastique du laboratoire StorageReview et nous permettent de tester une large gamme de produits Flash émergents sur une plate-forme de test unique. Cela nous permet d'avoir des rapports de performance reproductibles, ce qui est fondamental pour présenter des données fiables.
Bien que ce système ne soit clairement pas destiné au grand public, tout laboratoire comme le nôtre, qui teste une grande variété de disques, connaît déjà la marque Serial Cables. Ils fabriquent depuis de nombreuses années des cartes d'adaptation devenues incontournables dans de nombreux environnements de test. Ce système JBOD Gen5 est un excellent outil supplémentaire sur lequel nous nous appuierons tout au long de nos tests de SSD Gen5. Vous pouvez consulter les derniers résultats obtenus avec ce système de test dans notre test du Memblaze PBlaze7.
