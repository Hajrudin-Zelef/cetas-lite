---
id: collect-261001-general-networking/general-networking/fr-review-hci-intel-qlc-ssd-277e2f61-3
title: "fr-review-hci-intel-qlc-ssd-277e2f61"
domain: general-networking
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["intel"]
source: docs/RAG/collect-261001-general-networking/fr-review-hci-intel-qlc-ssd-277e2f61.md
source_anchor: ""
source_lines: [50, 54]
sha256: 498ed07db8258ced0ad198de5ad4602d5391b77cff3dffdd051cf8b63fc0760c
---

# fr-review-hci-intel-qlc-ssd-277e2f61

Il s'agit d'une notion essentielle, car les SSD QLC fournissent une capacité massive et dense au cluster, tout en offrant les avantages du TCO qui accompagnent le stockage flash. Pour enfoncer le clou, les disques QLC permettent une capacité de 15.36 To par baie de disque 2.5 pouces. Il faudrait 8 disques durs de 2 To en RAID 0 pour correspondre à la capacité, ou passer à un châssis de 3.5 pouces pour profiter de disques durs plus grands, mais encore plus lents. Dans tous les cas, la baisse des performances du lecteur Intel QLC vers les disques durs est plus que considérable ; c'est une différence exponentielle en ce qui concerne la réactivité des applications.
Même si nous aimerions que toutes les lectures et écritures proviennent des SSD Optane (car ce sont les supports les plus performants dans cette configuration), il y aura parfois un échec. Dans ce cas, les performances du SSD QLC écraseront les disques durs, protégeant le cluster HCI des irrégularités de performances courantes dans les topologies qui combinent des disques flash et des disques durs. En fait, nous avons vu des performances si équilibrées ici qu'à l'avenir, les entreprises en général devront peut-être repenser la conception du disque dur/flash et se pencher davantage vers la conception QLC/Optane pour tirer le meilleur parti du HCI.
L'autre préoccupation majeure concernant les clusters à 2 nœuds est la performance dans un état dégradé. Nous l'avons testé en faisant échouer un nœud et en donnant toute la charge de travail SQL à un seul nœud. Dans ce cas, SQL était plus réactif et fonctionnait un peu mieux que dans 2 nœuds, principalement en raison de la surcharge réduite des communications nœud à nœud. Bien sûr, il n'est pas recommandé de rouler longtemps dans un état dégradé comme celui-ci, mais il est rassurant de savoir que cela peut être fait sans sacrifier les performances.
Dans l'ensemble, le cluster HCI-224 HCI avec des SSD D5-P4326 QLC était simple à déployer, facile à utiliser et suffisamment puissant pour une large gamme de charges de travail. Son prix le rend également accessible à un large éventail d'utilisateurs. De plus, ce système a été certifié pour Microsoft Windows Server 2019 et validé en tant que solution Intel Select.
Ce rapport est parrainé par DataON. Tous les points de vue et opinions exprimés dans ce rapport sont basés sur notre vision impartiale du ou des produits à l'étude.
