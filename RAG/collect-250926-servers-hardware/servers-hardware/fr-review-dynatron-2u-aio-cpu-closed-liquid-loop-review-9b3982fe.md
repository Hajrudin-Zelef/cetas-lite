---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dynatron-2u-aio-cpu-closed-liquid-loop-review-9b3982fe
title: "fr-review-dynatron-2u-aio-cpu-closed-liquid-loop-review-9b3982fe"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "gpu", "intel"]
source: docs/RAG/clean4/fr-review-dynatron-2u-aio-cpu-closed-liquid-loop-review-9b3982fe.md
source_anchor: ""
source_lines: [1, 30]
sha256: 2a802b124684ee06f48ad7f6e37209de5b765e340a6f79b16edf8f93db356470
---

# fr-review-dynatron-2u-aio-cpu-closed-liquid-loop-review-9b3982fe

En matière de refroidissement de serveurs et de stations de travail, Dynatron est une référence. Outre ses refroidisseurs à air, Dynatron propose également de nombreuses solutions de refroidissement liquide. Vous avez peut-être déjà aperçu leur nom, ou même des refroidisseurs sans marque, dans des machines comme le 45Homelab HL15 ou le TYAN Transport HX FT65T-B8050 que nous avons testés. Dynatron fabrique et fournit ces refroidisseurs aux constructeurs (OEM) pour leurs systèmes.
Nous avons reçu un système de refroidissement liquide tout-en-un Dynatron personnalisé pour le tester en laboratoire. Pour ce faire, nous l'avons installé dans le boîtier Tyan HX FT65T-B8050 que nous avons testé il y a quelques semaines. Ce système remplace le ventirad Dynatron J10 d'origine. Dans notre configuration, nous avons retiré les ventilateurs du radiateur et positionné ce dernier contre ces derniers, au centre du boîtier.
Ce système AIO est composé du radiateur L35 de Dynatron et de leur plaque froide SP5 de leur L32 , assemblés dans une configuration personnalisée pour s'adapter à la bête qu'est l' EPYC 9684X.
Spécifications Dynatron AIO
Il est difficile d’obtenir les spécifications exactes de cette unité car il s’agit d’une configuration personnalisée et il n’existe pas de fiche technique formelle pour cette configuration. Les spécifications ci-dessous concernent la configuration du radiateur et du ventilateur L35.
| Catégories | Information | 
|---|---|
| Applications serveur | Serveur 2U et supérieur, serveurs tour | 
| Dimensions du ventilateur | 80 x 80 x 38 mm (3.15 x 3.15 x 1.5 pouces) | 
| Dimensions de l'ensemble radiateur | 323.2 x 44 (82 avec ventilateurs) x 85.15 mm (12.72 x 1.73 (3.23 avec ventilateurs) x 3.35 pouces) | 
| Vitesse du Ventilateur |  | 
| Consommation d'énergie |  | 
| Niveau de bruit |  | 
| Flux d'air |  | 
| Pression de l'air |  | 
Applications en boucle liquide fermée
Les systèmes de refroidissement liquide en circuit fermé peuvent constituer une option intéressante pour les serveurs, car ils offrent un niveau sonore réduit et une plus grande flexibilité d'encombrement que d'autres solutions. OSS a tiré parti de cette flexibilité dans sa configuration Gen 5 SDS présentée à SC23, dont nous avons parlé en novembre dernier. L'utilisation d'un système de refroidissement liquide Dynatron SP5 a permis à OSS de positionner les alimentations à l'emplacement initialement prévu pour un refroidisseur à air, libérant ainsi de l'espace pour les GPU tout en respectant leur format.
En ce qui concerne le bruit, nous avons rencontré des problèmes sur le Tyan où le petit ventilateur du refroidisseur d'air J10 hurlait sous charge. Cependant, une fois le refroidisseur de liquide remplacé, nous avons remarqué un volume beaucoup plus faible sous charge. Le ventilateur du processeur, qui était autrefois lié à la charge du processeur en termes de bruit, a complètement disparu, car nous exploitions les ventilateurs du châssis existants pour la circulation de l'air du radiateur. Ainsi, dans notre système, les niveaux de bruit ont chuté de façon spectaculaire.
Performances de la boucle Dynatron
Pour les avantages en termes de performances, c’est là que les choses sont biaisées entre les refroidisseurs à air et les refroidisseurs à liquide. Les bons refroidisseurs d’air fonctionnent très bien, et il en va de même pour les refroidisseurs de liquide. Lorsqu’ils sont tous deux conçus pour suivre les mêmes spécifications en termes de quantité d’énergie thermique qu’ils peuvent déplacer d’une zone donnée, les résultats seront assez similaires entre les deux.
Dans notre cas, nous avons effectué des tests sur le même matériel, en échangeant uniquement le refroidisseur, et nous avons constaté un changement minime. Nous avons chargé le processeur avec une charge de travail HPC intensive et comparé la vitesse d'horloge moyenne du cœur entre les deux refroidisseurs.
On pourrait dire que cela a peut-être changé les choses dans un très petit pourcentage, mais cela reste dans la marge d’erreur de ces tests. L’avantage du refroidissement liquide n’est pas toujours de meilleures performances, mais il donne au client une plus grande flexibilité en termes de gestion de la chaleur, du bruit et de la consommation électrique dans différents environnements.
Autres offres Dynatron
Comme mentionné précédemment, Dynatron couvre bien plus que les refroidisseurs d'air et de liquide pour processeur ; ils ont leur nom partout sur le marché du refroidissement. Dynatron propose d'autres produits tels que des dissipateurs thermiques passifs, des pièces de boucle de liquide personnalisées, des pièces RVB, des ventilateurs soufflants et des ventilateurs de boîtier. Ces différentes options permettent aux constructeurs OEM d'obtenir tous leurs besoins en refroidissement auprès d'un seul fournisseur, ainsi que la possibilité de disposer de solutions de refroidissement personnalisées.
Les options de refroidisseur de processeur proposées par Dynatron couvrent plusieurs facteurs de forme et formes sur le même socket pour vous offrir une grande flexibilité dans vos configurations. Dynatron propose également des options qui couvrent des tonnes de sockets non seulement pour les processeurs de serveur mais également pour les processeurs de bureau. Les options de socket AMD sont FM1, FM2(+), AM2(+), AM3(+), AM4, AM5, C32, G34, Opteron 6000 et 6100, socket F, SP3, SP5, SP6, sWRX8, sTRX4, TR4, et TR5. Pour les processeurs Intel, Dynatron couvre les sockets LGA1200, 115x, 1356, 1366, 1700, 1851, 2011 étroit et carré, 2066 étroit et carré, 3647 étroit et carré, 4677, 4710 et 7529, PGA479 et PGA988.
Pour les systèmes de refroidissement liquide, Dynatron propose des produits adaptés aux sockets grand public comme l'AM5, ainsi qu'aux sockets serveurs comme le SP5 pour les processeurs AMD. Dynatron offre également un large choix de solutions pour les processeurs Intel. Le socket AM5 dispose de plusieurs options de refroidissement, du radiateur à ventilateur unique de 120 mm du L5 au L25-u , équipé de cinq ventilateurs de 40 mm pour un châssis rackable 1U.
Actuellement, le socket SP5 n'est proposé qu'avec le refroidisseur L32 qui est conçu pour prendre en charge jusqu'à 500 W de TDP en seulement 1U, mais nous nous attendons à voir plus d'options à l'avenir. Dynatron proposera probablement différents facteurs de forme comme des radiateurs 2U, et peut-être jusqu'à 4U ou des tailles de tour plus grandes pour ce socket à l'avenir, similaire à ce qu'ils ont fait avec d'autres sockets.
Toutes ces différentes options de Dynatron vous offrent une couverture presque partout où vous en avez besoin sur le marché du refroidissement.
Conclusion
Dans l’ensemble, nous n’avons peut-être pas constaté d’énormes différences thermiques avec ce refroidisseur de liquide, mais nous avons constaté une réduction du bruit. L'utilisation d'un refroidisseur d'eau vous permet également de déplacer le refroidisseur vers d'autres zones du châssis pour libérer de l'espace au plafond. Un autre avantage du refroidisseur de liquide est que vous devez moins vous soucier de la canalisation du flux d’air puisque le radiateur peut couvrir une plus grande surface. Grâce à la variété d'options de refroidissement proposées par Dynatron, ils couvrent de nombreuses applications sur le marché. Si vous avez des besoins en matière de refroidissement du processeur de serveur, Dynatron a probablement croisé votre radar avec au moins un de ses produits.
