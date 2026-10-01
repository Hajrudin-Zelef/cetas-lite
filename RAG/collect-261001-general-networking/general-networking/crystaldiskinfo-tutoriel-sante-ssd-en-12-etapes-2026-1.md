---
id: collect-261001-general-networking/general-networking/crystaldiskinfo-tutoriel-sante-ssd-en-12-etapes-2026-1
title: "Calculer l'empreinte SHA-256 de l'installeur téléchargé"
domain: general-networking
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["attention", "gpu", "intel", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/crystaldiskinfo-tutoriel-sante-ssd-en-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 38]
sha256: 90b387edc558659163f11f8c07052ae9f5fef5243455368268f3dc2ba26a02bc
---

# Calculer l'empreinte SHA-256 de l'installeur téléchargé

Un SSD ou un disque dur ne prévient presque jamais avant de rendre l’âme : un matin, la machine ne démarre plus, et des années de données s’évaporent. Pourtant, la quasi-totalité des disques modernes tiennent un journal de bord interne – la technologie **S.M.A.R.T.** – qui accumule des dizaines d’indicateurs sur leur usure réelle. Le problème, c’est que ces valeurs restent invisibles tant qu’un outil ne va pas les lire. C’est précisément le rôle de **CrystalDiskInfo**, l’utilitaire gratuit et open source qui traduit ces compteurs bruts en un verdict clair : *Bon*, *Attention* ou *Mauvais*.

Ce tutoriel vous guide pas à pas, en 12 étapes et une trentaine de minutes, pour installer **CrystalDiskInfo** dans sa version 9.9.2, publiée le 25 juillet 2026 par Crystal Dew World, décrypter les attributs S.M.A.R.T. qui comptent vraiment, configurer des alertes automatiques par e-mail et mettre en place un pipeline de surveillance qui vous prévient avant la panne. Que vous administriez un parc de PME en France, un homelab sous Proxmox ou simplement votre PC de bureau, vous saurez à la fin repérer un disque agonisant des semaines avant qu’il ne lâche.

## Qu’est-ce que CrystalDiskInfo et pourquoi surveiller la santé de vos disques ?

**CrystalDiskInfo** est un logiciel de diagnostic de stockage développé par le studio japonais Crystal Dew World (le développeur connu sous le pseudonyme hiyohiyo). Il lit les données S.M.A.R.T. (*Self-Monitoring, Analysis and Reporting Technology*) directement dans le micrologiciel de chaque disque, puis les présente dans une interface lisible. Contrairement à un simple gestionnaire de fichiers, il ne s’intéresse pas à ce que contient le disque, mais à son état physique : nombre de secteurs défectueux, heures de fonctionnement, cycles d’écriture, température, taux d’erreur. Le code source est publié sur GitHub, ce qui en fait l’un des rares outils de ce type entièrement auditable.

La question qui revient toujours : *pourquoi ne pas attendre que Windows signale un problème ?* Parce que Windows n’affiche un avertissement S.M.A.R.T. que lorsque le disque a déjà franchi un seuil critique – souvent trop tard pour sauvegarder sereinement. Les grandes études de fiabilité le confirment : dans son rapport Drive Stats, l’hébergeur Backblaze, qui surveille plus de 300 000 disques, a identifié cinq attributs S.M.A.R.T. particulièrement corrélés à une panne imminente : les secteurs réalloués (ID 05), les erreurs non corrigibles signalées (187), les délais de commande dépassés (188), les secteurs en attente (197) et les secteurs non corrigibles hors ligne (198). Quand l’un de ces compteurs commence à grimper, la panne n’est plus une question de « si », mais de « quand ».

Surveiller ces indicateurs avec **CrystalDiskInfo** transforme une panne subie en panne anticipée. Vous gagnez la fenêtre de quelques jours à quelques semaines nécessaire pour cloner le disque, commander un remplacement ou activer votre garantie constructeur – le tout avant la perte de données. Pour une surveillance de l’ensemble du système (CPU, GPU, cartes mères, ventilateurs), reportez-vous à notre tutoriel HWiNFO ; CrystalDiskInfo, lui, se concentre exclusivement et en profondeur sur le stockage.

## CrystalDiskInfo en 2026 : version 9.9.1, éditions et licence

La version stable la plus récente est **CrystalDiskInfo 9.9.2**, mise en ligne le 25 juillet 2026 par Crystal Dew World sous la forme d’un installeur EXE de 6,2 Mo et d’une archive ZIP portable de 8,5 Mo (chiffres SourceForge ; build EXE daté du 29 juillet sur Uptodown, où l’outil affiche une note de 4,5/5), succédant à la 9.9.1 du 23 mai 2026 – déjà créditée de 159 438 téléchargements hebdomadaires sur SourceForge pour son seul installeur EXE –, à la 9.9.0 du 18-19 mai 2026 et à la 9.8.0 du 16 février 2026 (distribuée en archive ZIP de 8,1 Mo). Le logiciel pèse ainsi à peine quelques mégaoctets, se lance en une seconde et reste entièrement gratuit – un modèle rare pour un outil aussi complet. Il ne faut pas le confondre avec **CrystalDiskMark**, son outil frère du même développeur : ce dernier mesure les *débits* (vitesses de lecture/écriture) du disque, tandis que CrystalDiskInfo mesure sa *santé*. Les deux sont complémentaires, mais répondent à des questions différentes.

Un point déroute souvent les nouveaux venus : CrystalDiskInfo se décline en quatre « éditions » – Standard, Shizuku, Kurei Kei et Aoi. Rassurez-vous, elles sont **fonctionnellement identiques**. Seule change l’habillage graphique : les éditions Shizuku, Kurei Kei et Aoi embarquent des thèmes illustrés par des mascottes de l’univers du logiciel, ce qui alourdit nettement le téléchargement – la build Aoi de la 9.9.2, datée elle aussi du 25 juillet 2026, pèse 65,0 Mo en EXE et 67,0 Mo en ZIP sur SourceForge, contre seulement 6,2 Mo et 8,5 Mo pour l’édition Standard équivalente. Pour une utilisation professionnelle ou un déploiement en entreprise, l’édition **Standard** est le choix évident : plus légère, sobre, sans illustration superflue.

Côté couverture matérielle, la 9.9.2 reconnaît les plateformes récentes – SSD NVMe PCIe 5.0, contrôleurs Intel les plus récents, disques SATA de grande capacité – et lit aussi bien les disques internes que la plupart des boîtiers externes USB. Elle prend en charge les architectures x86, x64 et ARM64, ce qui la rend utilisable jusque sur les PC Windows on ARM. Cette régularité de mises à jour – de la 9.6.0 du 25 février 2025 à la 9.6.3 du 12 mars 2025, puis la 9.7.0 des 16-17 juin 2025, la 9.7.1 du 28 juillet 2025 et la 9.7.2 du 1er septembre 2025, avant les 9.8.0 du 16 février 2026, les 9.9.0 et 9.9.1 – dont l’archive ZIP portable cumulait à elle seule 64 338 téléchargements hebdomadaires sur SourceForge fin mai 2026 – puis la 9.9.2 – est un gage de fiabilité : le support des nouveaux disques suit de près leur sortie.

## Prérequis et configuration système requise

Avant de commencer, vérifiez que votre environnement coche les cases suivantes. CrystalDiskInfo est peu gourmand, mais quelques prérequis conditionnent l’accès à certaines fonctions avancées (lecture NVMe, alertes par e-mail).

| Élément | Exigence minimale | Recommandé pour ce tutoriel | 
|---|---|---|
| Système d’exploitation | Windows 7 / 8 / 8.1 / 10 / 11, Server 2003 à 2025 | Windows 10 22H2 ou Windows 11 24H2 | 
| Architecture | x86, x64 ou ARM64 | x64 | 
| Lecture des disques NVMe | Windows 10 / Server 2016 ou ultérieur | Windows 11 | 
| .NET Framework | Non requis pour l’usage de base | .NET Framework 4.8+ (obligatoire pour les alertes e-mail) | 
| Privilèges | Compte administrateur (accès S.M.A.R.T. bas niveau) | Administrateur local | 
| Espace disque | ~15 Mo | ~15 Mo (version portable possible) | 
| Version du logiciel | CrystalDiskInfo 9.9.1 (édition Standard) | CrystalDiskInfo 9.9.1 Standard | 

Deux limites techniques méritent votre attention. D’abord, les **ponts USB** : tous ne relaient pas les commandes S.M.A.R.T. Un même SSD affichera ses attributs en interne (SATA/NVMe) mais restera muet dans un boîtier externe bon marché dont le contrôleur ne transmet pas le protocole. Ensuite, les **configurations RAID** : CrystalDiskInfo lit une partie des grappes Intel RST, mais pas tous les contrôleurs RAID matériels, qui masquent les disques physiques derrière un volume logique. Ces deux cas de figure ne sont pas des bugs, mais des limitations du chemin d’accès matériel.

## Étape 1 – Télécharger CrystalDiskInfo depuis la source officielle

