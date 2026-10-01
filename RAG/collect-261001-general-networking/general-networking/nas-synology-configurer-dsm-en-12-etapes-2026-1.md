---
id: collect-261001-general-networking/general-networking/nas-synology-configurer-dsm-en-12-etapes-2026-1
title: "Exemple de planification de tâche Hyper Backup (via l'interface DSM)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Google"]
dates: []
keywords: ["agent", "amd", "ethernet"]
source: docs/RAG/collect-261001-general-networking/nas-synology-configurer-dsm-en-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 42]
sha256: 2a85a239c4d1ee89bc52cbd59165641f17ccd9b9624e53ae5951c06ffe05d0d1
---

# Exemple de planification de tâche Hyper Backup (via l'interface DSM)

Un NAS Synology reste allumé 24 heures sur 24 dans un coin du salon ou du bureau, souvent pendant des années sans qu’on y touche vraiment après l’installation initiale. C’est justement cette phase de démarrage qui coince le plus : entre le choix du système de fichiers, la configuration du volume, le paramétrage de QuickConnect et la mise en place d’une vraie stratégie de sauvegarde, il y a une bonne douzaine de décisions à prendre correctement dès le premier jour. Se tromper sur le système de fichiers ou sur le type de RAID, par exemple, oblige à tout reformater plus tard, ce qui veut dire réinjecter des téraoctets de données depuis zéro.

Ce tutoriel couvre la configuration complète d’un NAS Synology sous DSM, de la sortie de la boîte jusqu’à un système de sauvegarde automatisé et sécurisé, avec les pièges à éviter et les réglages qui comptent vraiment pour un usage domestique ou pour un petit bureau en France. Nous utilisons DSM 7.4.1‑90080 (la version générale la plus récente au 20 août 2026), point de mise à jour de la branche DSM 7.4 lancée le 16 juin 2026 avec l’assistant piloté par IA **DSM Agent** et la déduplication/compression a posteriori des volumes HDD, tandis que la branche de maintenance longue durée DSM 7.2.2‑72806 reste mentionnée à titre de repère pour les modèles plus anciens.

## Pourquoi configurer un NAS Synology en 2026

La question du stockage local revient en force depuis deux ans, portée par trois choses : la hausse des abonnements cloud grand public, les inquiétudes sur la localisation des données personnelles, et l’explosion des volumes de photos et vidéos 4K produits par les smartphones. Un NAS 2 baies comme le **Synology DS224+** se trouve autour de **400 à 430 € TTC** en France mi‑2026 (comparateurs 123Comparer et Labonneconfig, prix relevés en août 2026), disques durs non compris. Un modèle 4 baies comme le **DS423+** tourne plutôt entre **430 et 600 € TTC** pour le boîtier seul, jusqu’à plus de 900 € en kit avec disques.

Sur le papier, ça paraît cher face à 200 Go de stockage cloud à quelques euros par mois. Mais un NAS bien configuré remplace en même temps un service de sauvegarde photo, un serveur de fichiers partagé, un lecteur multimédia pour la télévision et, avec Synology Drive, une alternative à Google Drive ou OneDrive hébergée chez soi. Pour les foyers ou petites structures qui veulent garder la main sur leurs données plutôt que de dépendre d’un cloud tiers, l’angle rejoint directement les problématiques de souveraineté numérique face au cloud public déjà traitées sur ce site, et fait écho aux recommandations générales de la CNIL sur la minimisation des données confiées à des tiers.

Le NAS Synology occupe une place particulière dans ce paysage : DSM (DiskStation Manager) est une interface web complète, proche d’un vrai système d’exploitation, avec un catalogue d’applications (Paquets) qui couvre la sauvegarde, la synchronisation, la surveillance vidéo, le VPN ou même un serveur mail. C’est ce qui distingue Synology d’un NAS plus brut ou d’une solution auto-hébergée : l’installation initiale est guidée de bout en bout, ce qui réduit fortement le risque d’erreur pour un premier NAS.

## Prérequis matériels et logiciels

Avant de commencer, voici ce qu’il faut avoir sous la main. La configuration matérielle influence directement les performances et les options disponibles dans DSM, donc mieux vaut vérifier chaque point avant de sortir le tournevis.

| Élément | Recommandation | Détail | 
|---|---|---|
| Modèle NAS | DS224+ (2 baies) ou DS423+/DS923+ (4 baies) | DS923+ : CPU AMD Ryzen R1600, RAM ECC 4 Go extensible à 32 Go | 
| Disques durs | 2 minimum (idéalement identiques) | Disques NAS dédiés, pas des disques de bureau grand public | 
| Version DSM | DSM 7.4.1‑90080 ou plus récent | Publiée le 23 juillet 2026, branche générale | 
| Réseau | Câble Ethernet Cat 6 minimum | Le Wi-Fi fonctionne mais ralentit fortement les premiers transferts | 
| Ordinateur d’installation | Windows, macOS ou Linux | Navigateur à jour (Chrome, Edge ou Firefox) | 
| Compte Synology | Gratuit, créé pendant l’installation | Nécessaire pour QuickConnect et les notifications | 
| Onduleur (optionnel mais conseillé) | 400 VA minimum | Protège contre les coupures qui corrompent les volumes Btrfs | 

Un point souvent négligé : les disques durs. Un disque grand public conçu pour tourner quelques heures par jour dans un PC de bureau n’est pas fait pour fonctionner 24h/24 dans un NAS. Les fabricants proposent des gammes dédiées, plus tolérantes à la vibration et calibrées pour un fonctionnement continu. Bonne nouvelle pour les propriétaires de modèles récents : DSM 7.3, sorti en octobre 2025, a restauré la prise en charge des disques SATA tiers sur les DiskStation de cette génération, après les restrictions imposées un temps par Synology. Cela dit, utiliser des disques de bureau dans un NAS Synology fonctionne à court terme, mais les taux de panne grimpent nettement après 12 à 18 mois d’usage continu selon les retours d’expérience de la communauté Synology.

## Étape 1 : Préparer le matériel et insérer les disques

Éteignez le NAS avant toute manipulation, même si la plupart des modèles Synology récents supportent l’insertion à chaud. Sur un DS224+ ou un DS423+, les tiroirs de disques se retirent sans outil : un clic sur la languette latérale suffit à les libérer. Insérez le disque dans le tiroir, vissez-le sur les côtés (4 vis fournies) si vous utilisez un modèle sans fixation tool-less pour les baies 3,5 pouces, puis remettez le tiroir en place jusqu’au clic.

Sur les modèles Plus récents comme le DS423+, deux emplacements M.2 NVMe supplémentaires sont disponibles sous le boîtier, séparés des baies principales. Ces emplacements servent presque exclusivement au cache SSD (voir étape 9) et non à créer un volume de stockage principal — Synology limite volontairement cet usage sur l’entrée de gamme Plus. Si vous comptez utiliser un NVMe comme stockage de volume complet, vérifiez la compatibilité exacte du modèle avant achat, car cette fonctionnalité n’est pas activée sur tous les DS Plus.

Une fois les disques en place, branchez le câble Ethernet sur le port LAN1 (ou sur les deux ports si votre routeur supporte l’agrégation de liens), puis le câble d’alimentation. Appuyez sur le bouton d’alimentation pendant une seconde. Le voyant d’état clignote en bleu pendant l’amorçage, ce qui prend généralement 90 secondes à 3 minutes selon le modèle.

## Étape 2 : Trouver le NAS sur le réseau avec Synology Assistant

Depuis un ordinateur connecté au même réseau local, deux méthodes permettent de localiser le NAS fraîchement démarré. La première passe par le site **find.synology.com**, qui lance le Web Assistant directement dans le navigateur : il scanne le réseau local et affiche les NAS Synology détectés avec leur statut (“Non installé” pour un appareil neuf).

La seconde méthode utilise **Synology Assistant**, un utilitaire de bureau téléchargeable depuis le site Synology, qui scanne le réseau local en dehors du navigateur et permet ensuite de gérer plusieurs NAS depuis une interface centralisée. C’est l’outil à privilégier si vous gérez plusieurs appareils Synology ou si le Web Assistant ne détecte rien (pare-feu d’entreprise, réseau segmenté en VLAN, etc.).

