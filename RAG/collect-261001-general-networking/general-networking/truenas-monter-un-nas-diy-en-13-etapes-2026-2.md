---
id: collect-261001-general-networking/general-networking/truenas-monter-un-nas-diy-en-13-etapes-2026-2
title: "Identifier la clé USB (attention à bien cibler le bon disque)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Intel"]
dates: []
keywords: ["attention", "intel"]
source: docs/RAG/collect-261001-general-networking/truenas-monter-un-nas-diy-en-13-etapes-2026.md
source_anchor: ""
source_lines: [38, 94]
sha256: 377c45403abd0b0d8abe24c31f505f5fe9425f42f032eef7e70a731d06e7852f
---

# Identifier la clé USB (attention à bien cibler le bon disque)

Côté budget, comptez large. D’après les relevés de prix idealo.de et Amazon.fr/de au 15 août 2026, un disque dur NAS/entreprise coûte en moyenne 48 à 50 €/To, avec quelques promotions ponctuelles descendant vers 25 €/To sur de gros volumes externes (comme un WD Elements 22 To à 549 €). Une carte mère mini-ITX avec processeur Intel N100 intégré, 6 ports SATA et réseau 2,5 GbE se trouve entre 130 € et 210 € selon la marque (ASRock N100DC-ITX autour de 130 $, cartes Topton/SZBOX entre 139 € et 206 €). Les modèles à base d’Intel N150, plus récents, montent plutôt entre 180 € et 280 €.

| Composant | Option d’entrée de gamme | Option recommandée 2026 | Prix approximatif | 
|---|---|---|---|
| Carte mère + CPU | Intel N100 mini-ITX (ASRock N100DC-ITX) | Intel N150 mini-ITX 6 baies SATA (Topton N18) | 130 € – 280 € | 
| RAM | 8 Go DDR4/DDR5 | 32 Go DDR5 (pour ARC et apps Docker) | 35 € – 110 € | 
| Disque de démarrage | Clé USB 32 Go | SSD SATA 120-256 Go | 15 € – 25 € | 
| Stockage principal (par To) | HDD grand public | HDD NAS/entreprise (WD Red Pro, Seagate IronWolf Pro) | 48 € – 50 €/To | 
| Boîtier | Boîtier mini-ITX 4 baies | Boîtier NAS 6-8 baies avec ventilation dédiée | 70 € – 180 € | 
| Alimentation | SFX 300-450W 80+ Bronze | SFX 450-550W 80+ Gold | 50 € – 100 € | 

## Étape 1 : Choisir le processeur et la carte mère

Le choix du processeur détermine tout le reste de la construction. Pour un usage NAS pur (stockage, partages SMB/NFS, sauvegardes), un Intel N100 suffit largement : il consomme peu (6W de TDP), reste silencieux avec un simple dissipateur passif et gère sans problème le chiffrement AES-NI nécessaire au chiffrement des pools ZFS. Si vous prévoyez de transcoder de la vidéo avec Plex ou Jellyfin, ou de faire tourner plusieurs containers Docker en parallèle (Nextcloud, Immich, Home Assistant), montez plutôt vers un Intel N150 ou un Core i3-N305, qui offre 8 cœurs contre 4 pour le N100.

Sur le choix de la carte mère, privilégiez les modèles mini-ITX orientés NAS qui intègrent directement 6 à 8 ports SATA III, comme les cartes Topton N18 ou les modèles ASRock équivalents. Elles évitent d’avoir à acheter une carte contrôleur HBA séparée. Vérifiez systématiquement trois points avant l’achat : la présence d’un port réseau 2,5 GbE minimum (le 10 GbE reste un vrai plus si votre switch le supporte), un slot M.2 NVMe pour le cache ou le disque de démarrage, et la compatibilité mémoire DDR5 si vous visez 32 Go ou plus.

## Étape 2 : Dimensionner la RAM et le stockage de démarrage

Le guide matériel officiel de TrueNAS fixe la barre basse à 8 Go de RAM pour une installation basique avec jusqu’à 8 disques, avec 1 Go supplémentaire par disque au-delà. C’est un minimum technique pour que le système démarre et fonctionne, pas une recommandation pour un usage confortable. En pratique, la communauté homelab et plusieurs guides matériels 2026 convergent vers 16 Go comme plancher réaliste pour un NAS familial, et 32 Go ou plus si vous comptez sauvegarder des VM via iSCSI ou faire tourner plusieurs applications Docker simultanément.

### Le rôle du cache ARC dans ZFS

ZFS, le système de fichiers utilisé par TrueNAS, exploite la RAM disponible comme cache de lecture appelé ARC (Adaptive Replacement Cache). Plus vous avez de RAM, plus ZFS peut garder de données fréquemment consultées en mémoire, ce qui réduit les accès disque et accélère les lectures. Ce n’est pas un luxe : sur un NAS avec seulement 8 Go, l’ARC dispose de très peu de marge une fois le système d’exploitation et les services lancés, ce qui dégrade sensiblement les performances dès que plusieurs utilisateurs accèdent au NAS en même temps.

Pour le disque de démarrage, TrueNAS impose un minimum de 16 à 20 Go selon la version. Évitez la clé USB en usage permanent : elle fonctionne pour l’installation, mais sa durée de vie en écriture continue est nettement inférieure à celle d’un SSD SATA ou NVMe bon marché, qui coûte aujourd’hui entre 15 € et 25 € pour 120 à 256 Go. C’est un des postes de dépense les plus rentables de toute la construction.

## Étape 3 : Choisir la topologie ZFS – RAIDZ1, RAIDZ2 ou Mirror

C’est la décision la plus structurante de tout le projet, car elle détermine votre capacité utile, votre tolérance aux pannes et vos possibilités d’extension future. ZFS propose trois topologies de vdev principales, chacune avec un nombre minimum de disques imposé par la documentation officielle.

| Topologie | Disques minimum | Tolérance de panne | Capacité utile (avec 4 disques de 4 To) | Cas d’usage recommandé | 
|---|---|---|---|---|
| Mirror (miroir) | 2 | 1 disque par paire | 8 To (sur 2 paires) | Lectures aléatoires rapides, bases de données, VM | 
| RAIDZ1 | 3 | 1 disque | 12 To | Stockage familial, budget serré, disques ≤ 8 To | 
| RAIDZ2 | 4 | 2 disques | 8 To | Données critiques, disques > 8 To, plus de 6 disques | 

La documentation ZFS officielle est claire sur la hiérarchie de fiabilité : RAIDZ2 offre une meilleure disponibilité des données et un MTTDL (temps moyen avant perte de données) nettement supérieur à RAIDZ1, car il tolère la perte simultanée de deux disques. En contrepartie, un mirror reste la topologie la plus performante pour les charges de lecture aléatoire, en particulier sur des workloads qu’on ne peut pas mettre en cache. Le livre blanc TrueNAS sur les layouts ZFS recommande de rester entre 3 et 9 disques par vdev, et déconseille formellement de dépasser 12 disques dans un seul vdev, quelle que soit la topologie choisie.

Pour un premier NAS familial avec 4 à 6 disques de 8 à 16 To, RAIDZ2 reste le choix le plus raisonnable en 2026 : avec la hausse des capacités unitaires, le temps de reconstruction (resilver) après panne s’allonge, et RAIDZ1 devient risqué puisqu’une seconde panne pendant la reconstruction entraîne une perte totale du pool.

## Étape 4 : Préparer la clé USB et installer TrueNAS 25.10.6 Goldeye

Téléchargez l’image ISO officielle depuis la page de téléchargement TrueNAS (25.10.6 pour la branche stable, évitez la 26.0.0-BETA.3 tant qu’elle reste en bêta). Si vous partez d’une installation existante encore sur l’ancienne branche 25.04, sachez qu’elle a reçu deux dernières mises à jour de maintenance, la 25.04.2.4 en septembre 2025 puis la 25.04.2.6 en octobre 2025, avant qu’iXsystems ne recommande officiellement de migrer vers la branche 25.10.x. Utilisez un outil comme Balena Etcher ou Rufus pour flasher l’ISO sur une clé USB de 8 Go minimum. En ligne de commande sous Linux ou macOS, la méthode `dd` reste la plus directe :

```
# Identifier la clé USB (attention à bien cibler le bon disque)
lsblk
# Flasher l'image ISO sur la clé USB (remplacez /dev/sdX par votre clé)
sudo dd if=TrueNAS-25.10.6.iso of=/dev/sdX bs=4M status=progress conv=fsync
# Vérifier le checksum SHA256 de l'ISO avant de flasher
sha256sum TrueNAS-25.10.6.iso
```
Une fois la clé prête, branchez-la sur votre futur NAS, entrez dans le BIOS/UEFI (touche Suppr ou F2 selon la carte mère), désactivez le Secure Boot s’il pose problème au démarrage, et sélectionnez la clé USB comme périphérique de boot. L’installeur TrueNAS se lance en mode texte : choisissez le disque de démarrage (jamais un disque que vous comptez utiliser pour le stockage de données), définissez le mot de passe root, et laissez l’installation se terminer. Redémarrez ensuite en retirant la clé USB.

## Étape 5 : Configuration réseau et première connexion à l’interface web

