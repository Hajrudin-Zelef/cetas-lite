---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-superchassis-846be1c-r1k28b-review-82358c79-2
title: "fr-review-supermicro-superchassis-846be1c-r1k28b-review-82358c79"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/clean4/fr-review-supermicro-superchassis-846be1c-r1k28b-review-82358c79.md
source_anchor: ""
source_lines: [50, 60]
sha256: 6e898082fcd5a664f69ddf0fd0861525eac67d38164274befd15d9d2c8eccfea
---

# fr-review-supermicro-superchassis-846be1c-r1k28b-review-82358c79

Pour tester le JBOD 846BE1C au lieu de le comparer à un autre JBOD, nous avons choisi de comparer l'exécution des 24 baies avec le disque de capacité la plus élevée actuellement disponible, le HGST Ultrastar Helium He8 8 To, par rapport à la même configuration avec 4 des disques remplacés par des SSD, HGST Ultrastar SSD800MR SAS3 500 Go, pour la hiérarchisation. Nous avons également utilisé des espaces de stockage pour effectuer les configurations initiales. Bien que la hiérarchisation SSD ait fourni de meilleurs résultats, ce n'est pas une surprise, ce qui était surprenant, c'est la différence dans les résultats. Dans le test du serveur SQL, nous avons vu presque doubler le TPS du test non hiérarchisé, 3,874.65 6,307.92 TPS, par rapport au multiniveau, 2,990 6 TPS. La plus grande différence résidait dans la latence, le non hiérarchisé exécutant un très haut XNUMX XNUMX ms et le hiérarchisé fonctionnant presque cinq cents fois plus vite à XNUMX ms.
Avantages
- Conception abordable et flexible
- Jusqu'à 192 To de capacité avec des disques de 8 To
- Excellentes performances avec les interconnexions SAS3
Inconvénients
- Des baies supplémentaires de 2.5 pouces à l'arrière amélioreraient l'efficacité de l'espace
Conclusion
Le SuperMicro SuperChassis 846BE1C-R1K28B 24 baies est un JBOD 3U compatible SAS4 offert avec beaucoup de flexibilité et de capacité à un prix abordable.
Page produit SuperMicro SuperChassis 846BE1C-R1K28B
Discutez de cet avis
