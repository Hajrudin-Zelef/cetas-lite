---
id: collect-261001-general-networking/general-networking/fr-review-lsi-sas-9207-8i-e-host-bus-adapter-review-f7ef8c4b-2
title: "fr-review-lsi-sas-9207-8i-e-host-bus-adapter-review-f7ef8c4b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "exploit"]
source: docs/RAG/collect-261001-general-networking/fr-review-lsi-sas-9207-8i-e-host-bus-adapter-review-f7ef8c4b.md
source_anchor: ""
source_lines: [48, 65]
sha256: 1d808e0b82042004d45b18ee60336d2df9e2556757ce1ae5f01bf2b4c624ea86
---

# fr-review-lsi-sas-9207-8i-e-host-bus-adapter-review-f7ef8c4b

Notre processus de référence de stockage d'entreprise commence par une analyse des performances du disque au cours d'une phase de préconditionnement approfondie. Chacune des baies de disques durs comparables est configurée en RAID10, autorisée à se synchroniser entièrement, puis testée sous une charge importante de 16 threads avec une file d'attente exceptionnelle de 16 par thread jusqu'à notre charge légère de 2 threads avec une file d'attente exceptionnelle de 2 par fil de discussion.
Nous examinerons deux des six analyses de charge de travail synthétiques d'entreprise qui ont été testées :
- Séquentiel
  - 8K
    - 100 % de lecture ou 100 % d'écriture
    - 100% 8K
  - 128K
    - 100 % de lecture ou 100 % d'écriture
    - 100% 128K
- 8K
Dans les performances de lecture et d'écriture séquentielles 8K, la vitesse de broche et la densité surfacique sont utilisées. Ici, le Hitachi Ultrastar 7,200K7 RAID4000 à 10 8 tr/min offrait la vitesse de lecture 911K la plus rapide, mesurant 82 Mo/s (15 Mo/s en écriture), tandis que le Toshiba 10K RAID811 mesurait 186 Mo/s (10 Mo/s en écriture) et le Toshiba 10K RAID612 mesurait 178 Mo. /s (XNUMX Mo/s en écriture).
Notre prochain test séquentiel a mesuré les vitesses de transfert de gros blocs. Dans ce profil de charge de travail utilisant le HBA LSI SAS 9207-8i/e, la baie SAS 15K mesurait 1,535 839 Mo/s en lecture et 7.2 Mo/s en écriture, tandis que la baie SATA 1,361 K enregistrait 912 10 Mo/s en lecture et 1,142 Mo/s en écriture, et la matrice SAS 540K avec XNUMX XNUMX Mo/s en lecture et XNUMX Mo/s en écriture.
Conclusion
Les HBA LSI SAS 9207 apportent une solution RAID SAS2 impressionnante pour les entreprises disposant de serveurs de stockage haut de gamme, en particulier pour ceux qui recherchent un équilibre entre prix abordable, performances, évolutivité ainsi que la fiabilité donnée de LSI. Conçu pour les applications de stockage de serveur de milieu de gamme, y compris le stockage hiérarchisé, la sauvegarde et la restauration, le 9207-8i connecte jusqu'à 256 périphériques SAS et SATA avec ses 8 ports SAS internes 6 Gb/s, tandis que le 9207-8e prend en charge jusqu'à 1024 périphériques SAS ou SATA avec ses 8 ports externes ; cependant, il doit s'agir de périphériques non RAID
Dans l'ensemble, la série 9207 est à la fois extrêmement stable et axée sur les performances. Le contrôleur d'E/S SAS LSISAS2308 6 Gb/s équipé, ainsi que son processeur PowerPC double cœur 800 MHz, permettent à ces HBA de se vanter de ces performances impressionnantes. Cela était certainement évident dans pratiquement toutes les applications pour lesquelles nous avons exploité les HBA LSI, comme une latence très faible et un débit très élevé avec l'extension de stockage JBOD iXsystems Titan 316J et le SSD Toshiba HK3R2, comme démontré ci-dessus. Par exemple, le lecteur Toshiba a affiché le meilleur benchmark de base de données MarkLogic NoSQL que nous ayons vu à ce jour, avec plus de 2.122 ms sur 32 threads tout en atteignant presque 1,700 32 TPS sur XNUMX threads dans notre benchmark MySQL.
Globalement, ces cartes LSI HBA sont difficiles à battre en matière de technologie SAS2, bien que les utilisateurs de SAS3 devraient se tourner vers la série LSI 9300.
Conclusion
LSI SAS 9207-8i/e est une série impressionnante de HBA offrant un équilibre entre accessibilité, performances et évolutivité tout en offrant aux entreprises une solution idéale pour les serveurs de stockage haut de gamme.
