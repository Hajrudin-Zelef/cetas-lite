---
id: collect-261001-huawei/huawei/fr-review-seven-things-to-love-about-arc-enabled-sql-managed-instances-721979cb-2
title: "fr-review-seven-things-to-love-about-arc-enabled-sql-managed-instances-721979cb"
domain: huawei
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["intel", "open source"]
source: docs/RAG/collect-261001-huawei/fr-review-seven-things-to-love-about-arc-enabled-sql-managed-instances-721979cb.md
source_anchor: ""
source_lines: [27, 35]
sha256: 52e45ebab14f4219bf241c71648eb8c48b30776cbf1375bb9088db3b97c8c715
---

# fr-review-seven-things-to-love-about-arc-enabled-sql-managed-instances-721979cb

Le plus souvent, la surveillance d'une base de données est une réflexion après coup, un coût supplémentaire ou négligée en raison de sa complexité ou de sa disponibilité. Microsoft a pris une décision audacieuse en incluant une pile de surveillance open source qui comprend InfluxDB et Grafana pour les métriques et Elastic et Kibana pour les journaux pour ses instances SQL gérées avec Arc.
Nous avons été surpris et ravis que Microsoft ait décidé d'utiliser des produits open source réputés et facilement extensibles pour la surveillance. Par exemple, Arc fournit un tableau de bord SQL Managed Instance compatible Grafana Arc avec des widgets qui affichent des indicateurs de performance clés et des métriques individuelles.
Un tableau de bord Grafana est également fourni aux hôtes.
Rétrospectivement, nous aurions dû intituler cet article « Les sept choses que nous avons le plus aimées à propos des instances SQL gérées avec Arc, exécutées sur Azure Stack HCI, avec l'intégration d'Arc sur un serveur basé sur Intel sécurisé fourni par DataON », car chacun de ces produits s'appuie sur et complète l'autre.
SQL Managed Instance permet une migration ou une création facile d'une base de données présentée et consommée en tant que PaaS. Azure Stack HCI permet à SQL Managed Instance compatible avec Arc et à d'autres services Azure de s'exécuter sur site. Arc permet à Azure Stack HCI et Azure dans le cloud d'être gérés à partir de la même interface Web. DataON est un partenaire apprécié de Microsoft et d'Intel qui fournit du matériel pour exécuter SQL Managed Instance compatible avec Arc dans le centre de données d'un client, dans un bureau distant ou en périphérie. Les serveurs basés sur Intel offrent une base sécurisée pour cette solution.
En regardant ce dernier paragraphe, il semble qu'il y ait beaucoup de pièces mobiles dans cette solution, mais elles s'emboîtent si bien qu'elles semblent être une solution unique. Peut-être une analogie à cela serait l'automobile. Bien qu'une automobile comprenne de nombreuses sous-sections complexes, elle se présente comme quelque chose dans lequel vous vous asseyez et conduisez avec toute la complexité sous-jacente qui apparaît à travers une interface unique.
Instance gérée SQL compatible Arc
Ce rapport est parrainé par DataON. Tous les points de vue et opinions exprimés dans ce rapport sont basés sur notre vision impartiale du ou des produits à l'étude.
.
