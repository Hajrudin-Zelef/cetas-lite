---
id: collect-261001-general-networking/general-networking/jellyfin-vs-plex-2026-0-vs-390-et-3-metadonnees-teste-1
title: "Exemple de configuration Nginx pour Jellyfin"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "AWS", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "benchmarks", "intel", "mai", "nvidia", "open source"]
source: docs/RAG/collect-261001-general-networking/jellyfin-vs-plex-2026-0-vs-390-et-3-metadonnees-teste.md
source_anchor: ""
source_lines: [1, 71]
sha256: 6669f0b8d45057e19d5a012a8e9fb8f0b380b373c2b0e567d7f4b3cf63d33346
---

# Exemple de configuration Nginx pour Jellyfin

En août 2026, le choix d’un serveur multimédia auto-hébergé se résume à deux candidats dominants : **Jellyfin**, le projet open source 100 % gratuit avec plus de 50 400 étoiles GitHub, et **Plex**, la solution freemium dont le Plex Pass à vie a bondi à **749,99 $** depuis le 1er juillet 2026 – une hausse de 200 % annoncée par Plex et relayée par AppleInsider et Tweaktown, tout comme par The FPS Review dès mai 2026 et par le comparateur suisse Digitec, qui a chiffré la hausse à plus de 3× (de 219,99 CHF à 749,99 $) – alors qu’un nouveau forfait intermédiaire de 5 ans à 249,99 $ a fait son apparition le 2 juillet 2026. Face à la collecte de données croissante des plateformes de streaming traditionnelles et aux augmentations tarifaires de Netflix, Disney+ et Amazon Prime en Europe, de plus en plus d’utilisateurs français se tournent vers l’auto-hébergement pour reprendre le contrôle de leur bibliothèque multimédia.

Ce comparatif exhaustif met face à face Jellyfin 10.11.11 – la version stable publiée le 6 juin 2026, qui a succédé à la 10.11.9 du 22 mai 2026 selon les notes de version documentées sur Wikipedia et Releasebot – et Plex en 2026, avec des tests de transcodage, une analyse des fonctionnalités, des benchmarks de performance et un guide de migration complet. Que vous gériez 500 films ou 50 000 fichiers musicaux, ce guide vous aide à choisir la solution adaptée à votre usage.

## Jellyfin vs Plex 2026 : tableau comparatif des spécifications

Avant d’entrer dans les détails, voici un aperçu synthétique des différences fondamentales entre Jellyfin et Plex en 2026. Ce tableau couvre les critères qui comptent le plus pour les utilisateurs français : prix, vie privée, compatibilité matérielle et fonctionnalités clés.

| Critère | Jellyfin 10.11.8 | Plex (2026) | 
|---|---|---|
| Prix | 100 % gratuit, open source (GPL v2) | Gratuit (limité) / Plex Pass : 6,99 $/mois, 69,99 $/an, 249,99 $ à vie | 
| Transcodage matériel | Gratuit (VAAPI, NVENC, QSV) | Plex Pass requis | 
| Code source | Open source (GitHub, 50 400+ étoiles) | Propriétaire, code fermé | 
| Collecte de données | Aucune – zéro télémétrie | Collecte de métadonnées, compte cloud obligatoire | 
| Accès distant | Configuration manuelle (reverse proxy) | Intégré, configuration en un clic | 
| Live TV / DVR | Gratuit, intégré | Plex Pass requis | 
| Applications mobiles | Gratuites, toutes fonctionnalités | Gratuites en local, Pass requis pour le distant | 
| Sous-titres | Plugins (Open Subtitles) | Recherche intégrée | 
| Précision des métadonnées | ~95 % (corrections manuelles nécessaires) | ~98 % (identification automatique) | 
| Installation Docker | Un conteneur, ~2 min | Un conteneur, ~2 min | 
| Serveur NAS | Synology, QNAP, TrueNAS | Synology, QNAP, FreeBSD, Netgear, Drobo | 
| Partage de bibliothèque | Comptes locaux illimités | Partage via comptes Plex (6 utilisateurs gérés) | 

## Philosophie et modèle économique : open source contre freemium

La différence fondamentale entre Jellyfin et Plex réside dans leur modèle économique, et cette distinction affecte chaque aspect de l’expérience utilisateur. Jellyfin est un fork d’Emby créé en 2018, entièrement développé par la communauté sous licence GPL v2. Chaque fonctionnalité – transcodage matériel, Live TV, synchronisation mobile – est accessible sans débourser un centime. Le projet est financé par des dons et maintenu par des centaines de contributeurs bénévoles sur GitHub.

Plex, fondé en 2007, a adopté un modèle freemium. La version gratuite permet de diffuser du contenu sur le réseau local, mais les fonctionnalités avancées sont verrouillées derrière le Plex Pass. Selon Bytesized Hosting, dès mars 2026 le Plex Pass s’élevait déjà à 6,99 $/mois et 249,99 $ à vie, contre 4,99 $/mois et 119,99 $ à vie auparavant – soit une hausse de 40 % sur l’abonnement mensuel et de 108 % sur le forfait à vie, documentée en avril 2026, avant même la flambée suivante. L’abonnement mensuel reste à **6,99 $** et l’annuel à **69,99 $** au 2 juillet 2026 selon Thurrott, mais le pass à vie a subi une seconde hausse spectaculaire : passé de 249,99 $ à **749,99 $** le 1er juillet 2026, soit une augmentation de 200 % confirmée par Plex lui-même et rapportée par AppleInsider, Tweaktown et Engadget. Pour amortir le choc, Plex a lancé le lendemain, le 2 juillet 2026, un nouveau **Plex Pass 5 ans à 249,99 $**, selon XDA Developers, tandis que le **Remote Watch Pass** reste disponible à 1,99 $/mois pour les utilisateurs qui accèdent aux serveurs d’autres personnes sans posséder eux-mêmes un Plex Pass.

Pour un utilisateur français, le coût sur 5 ans se résume ainsi : Jellyfin coûte 0 €, tandis que Plex Pass annuel revient à environ 325 € sur la même période (au taux de change actuel). Depuis la hausse du 1er juillet 2026, le pass à vie Plex à 749,99 $ (~695 €) ne devient rentable qu’après plus de 10 ans d’abonnement annuel, ce qui change complètement le calcul : le nouveau **Plex Pass 5 ans** à 249,99 $ (~232 €), lancé le 2 juillet 2026, est désormais l’option la plus avantageuse pour qui prévoit un usage de moyen terme, TechGeeks confirmant que la lecture locale de base reste gratuite quel que soit le forfait choisi.

## Installation et configuration : Docker, NAS et bare metal

L’installation est un facteur décisif pour de nombreux utilisateurs, en particulier ceux qui découvrent l’auto-hébergement. Les deux solutions proposent des images Docker officielles, ce qui simplifie considérablement le déploiement sur un serveur Linux, un NAS Synology ou un Raspberry Pi.

### Installation Docker de Jellyfin

```
docker run -d \
  --name jellyfin \
  --user 1000:1000 \
  --net=host \
  --volume /chemin/config:/config \
  --volume /chemin/cache:/cache \
  --volume /chemin/media:/media:ro \
  --restart unless-stopped \
  jellyfin/jellyfin:latest
```
### Installation Docker de Plex

```
docker run -d \
  --name plex \
  --net=host \
  -e PLEX_CLAIM="votre_claim_token" \
  -v /chemin/config:/config \
  -v /chemin/transcode:/transcode \
  -v /chemin/media:/data \
  --restart unless-stopped \
  plexinc/pms-docker:latest
```
Les deux installations prennent environ 2 minutes avec Docker. La principale différence réside dans la configuration post-installation : Plex nécessite un **compte cloud obligatoire** et une authentification en ligne, même pour un usage local. Jellyfin fonctionne entièrement en local, sans aucune dépendance cloud. Pour les utilisateurs soucieux de la souveraineté numérique – un sujet brûlant en France avec la migration de 2,5 millions d’appareils vers Linux – cette distinction est capitale.

Sur NAS, Plex dispose de packages natifs pour Synology (DSM 7), QNAP, FreeBSD, Netgear et Drobo. Jellyfin est disponible sur Synology, QNAP et TrueNAS, mais nécessite parfois une installation via Docker plutôt qu’un package natif, ce qui ajoute une étape supplémentaire pour les débutants.

## Transcodage et performance : benchmarks 4K HDR

Le transcodage est la fonctionnalité la plus gourmande en ressources d’un serveur multimédia. Il convertit un fichier vidéo dans un format compatible avec l’appareil de lecture en temps réel. C’est ici que la différence entre Jellyfin et Plex se fait le plus sentir.

Jellyfin propose le **transcodage matériel gratuit** via VAAPI (AMD/Intel), NVENC (NVIDIA) et Intel Quick Sync Video (QSV). Aucun abonnement n’est requis. En 2026, le transcodage de Jellyfin s’est considérablement amélioré, bien que certains formats complexes (comme le DTS-HD MA vers AAC) puissent nécessiter une configuration manuelle de l’accélération matérielle.

