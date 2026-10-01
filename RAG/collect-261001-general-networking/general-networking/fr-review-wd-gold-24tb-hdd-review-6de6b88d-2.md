---
id: collect-261001-general-networking/general-networking/fr-review-wd-gold-24tb-hdd-review-6de6b88d-2
title: "fr-review-wd-gold-24tb-hdd-review-6de6b88d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fr-review-wd-gold-24tb-hdd-review-6de6b88d.md
source_anchor: ""
source_lines: [73, 86]
sha256: 01a0f416118cb13f95cf01e4b3a260f647dd85f69601bdab3a84d6c17979d01e
---

# fr-review-wd-gold-24tb-hdd-review-6de6b88d

Les deux disques ont montré des améliorations significatives des performances dans le test de lecture/écriture 8K par rapport au test de lecture/écriture 4K. Dans la configuration iSCSI, le WD Gold de 24 To affichait 157 95 IOPS en lecture, 215 228 IOPS en écriture pour SMB, XNUMX XNUMX IOPS en lecture et XNUMX XNUMX IOPS en écriture en iSCSI.
8K 70 pour cent de lecture 30 pour cent d'écriture
L'autre type de charge de travail que nous exécutons passe d'un scénario de lecture/écriture séquentielle pure à une charge de travail mixte, démontrant comment les performances évoluent de 2T/2Q à 16T/16Q. Ici, les performances en SMB ont montré un minimum de 589 IOPS (par coïncidence, c'était la même chose que le Seagate Exos de 24 To) et un maximum de 1,648 913 IOPS. Avec iSCSI, nous constatons un minimum de 1,923 IOPS et un maximum de XNUMX XNUMX IOPS, avec des performances généralement constantes.
Pour la latence moyenne, nous constatons des pics notables (dans SMB) allant jusqu'à 152,53 ms tout en descendant jusqu'à 6 ms à 2 Threads 4 Queue. Nous avons constaté des scores similaires avec iSCSI, avec une latence aussi faible que 6.78 ms et un maximum de 132.96 ms avec 16 Threads 2 Queue.
En ce qui concerne la latence maximale, WD Golds a montré une latence nettement plus élevée avec SMB par rapport à iSCSI (comme les disques Seagate). Plus précisément dans SMB, la latence a commencé à 5,395.85 5736.11 ms et a culminé à 16 16 ms dans la file d'attente 1167.74 threads 3,618.97, tandis que iSCSI a montré une latence plus faible et plus cohérente, commençant à XNUMX XNUMX ms et culminant à XNUMX XNUMX ms.
Dans la dernière partie du test 8k 70/30, nous avons observé que la configuration iSCSI fonctionnait mieux en écart type. Plus précisément, la configuration SMB a montré une latence aussi faible que 24.1 ms avec un maximum de 373.75 ms, tandis que l'iSCSI présentait une plage de 12.29 ms à 150.22 ms pendant le test.
128K 100 % lecture/écriture
Notre test final est le test 128K 100% lecture/écriture. Ce test implique de gros blocs et démontre la vitesse de transfert séquentielle la plus élevée que les disques peuvent atteindre. Pour les performances des PME, le WD Gold de 24 To a atteint 1.38 Go/s en lecture et 2.3 Go/s en écriture, atteignant 2.08 Go/s en lecture et 2.27 Go/s en écriture pour iSCSI.
Prix et garantie du WD Gold 24 To
Bien que les détails spécifiques des prix puissent varier en fonction du détaillant et des remises potentielles sur les achats en gros, le disque dur d'entreprise WD Gold 24 To coûte actuellement 630 $ dans la boutique Western Digital et est accompagné d'une garantie limitée de cinq ans.
Conclusion
La gamme WD Gold marque une étape importante dans l'expansion de la capacité de stockage avec son modèle 24 To. Proposé exclusivement dans l'interface SATA, ce disque haute capacité est conçu pour répondre aux besoins rapidement croissants en solutions de stockage plus étendues dans les environnements d'entreprise. Le modèle 24 To conserve les caractéristiques qui ont fait de la série Gold un choix fiable pour les entreprises, notamment une conception SATA CMR 7200 XNUMX tr/min remplie d'hélium. Dans le même temps, son efficacité énergétique représente une consommation électrique réduite par rapport aux modèles précédents.
En ce qui concerne les performances, le disque dur Gold de 24 To présente une image complexe dans notre suite de tests. Au cours de nos tests rigoureux avec huit disques dans notre configuration NAS QNAP standard, le modèle 24 To a indiqué une baisse des performances dans les tests de lecture 4K par rapport à son homologue 22 To, où nous avons noté une diminution marquée des vitesses de lecture et une latence plus élevée. Le Seagate Exos X24 l'a également surpassé.
Cela dit, l'intégration de technologies de pointe, des prix compétitifs et un solide ensemble de garanties en font un choix solide pour les entreprises qui envisagent d'étendre leurs capacités de stockage. La série WD Gold reste une option pertinente et fiable pour les centres de données modernes recherchant un équilibre entre performances, fiabilité et potentiel de croissance future.
