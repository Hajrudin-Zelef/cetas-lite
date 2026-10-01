---
id: collect-261001-general-networking/general-networking/installer-nextcloud-cloud-souverain-13-etapes-2026-1
title: "Mise à jour complète du système"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: ["agi", "apache", "open source"]
source: docs/RAG/collect-261001-general-networking/installer-nextcloud-cloud-souverain-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 45]
sha256: 0c14a71db06e0f00bbc75c293440f7292c208d9347997e0f744d519926cb5090
---

# Mise à jour complète du système

La **souveraineté numérique** est devenue un impératif stratégique en France et en Europe en 2026. Depuis le lancement par France Stratégie de l’Observatoire de la souveraineté numérique en janvier 2026 et la formalisation d’un indice de résilience numérique évaluant les dépendances technologiques selon 20 critères répartis en huit dimensions, les PME, ETI et collectivités cherchent des alternatives concrètes aux grands fournisseurs américains. Reprendre le contrôle de ses données passe souvent par une décision simple : **installer Nextcloud**, la plateforme open source qui transforme un serveur en cloud privé.

Ce tutoriel vous guide en **13 étapes** pour héberger Nextcloud en France sur votre propre serveur, avec Docker Compose, PostgreSQL, Redis et un reverse proxy Caddy gérant automatiquement le HTTPS. Comptez environ **45 minutes** pour un déploiement complet et fonctionnel, prêt pour une équipe de 10 à 20 personnes. À la fin, vous disposerez d’un **cloud souverain open source** conforme au RGPD, sans boîte noire ni droit extraterritorial, avec édition collaborative de documents et chiffrement TLS. Mis à jour le 16 septembre 2026, suite à la publication par Nextcloud GmbH de la version **35.0.0** via l’installeur web le 16 septembre 2026, désormais la version stable de référence.

## Pourquoi installer Nextcloud pour un cloud souverain en 2026

Le débat public français a basculé. Depuis le Sommet de Berlin du 18 novembre 2025, la doctrine officielle ne considère plus le numérique comme un simple support technique mais comme un « champ de compétition » géopolitique. L’État a réagi avec des outils mesurables : l’Observatoire de la souveraineté numérique, lancé par France Stratégie, a pour mission de cartographier les dépendances technologiques de la France pour mieux les réduire, sa première collecte se poursuivant jusqu’à fin février 2026. En parallèle, le SILL (Socle Interministériel de Logiciels Libres) référençait plus de 600 logiciels libres recommandés à l’administration en avril 2026.

Nextcloud s’inscrit exactement dans cette trajectoire. Là où Google Drive, Microsoft OneDrive ou Dropbox stockent vos fichiers sur des infrastructures soumises au CLOUD Act américain, un **auto-hébergement Nextcloud** place vos données sous votre seul contrôle juridique et technique. Les communications officielles de 2026 sur la souveraineté numérique insistent sur trois principes que Nextcloud respecte nativement : la **réversibilité** (vos données restent exportables), la **portabilité** (formats ouverts) et l’**immunité au droit extraterritorial**. Le mouvement est déjà massif : dans un billet publié le 24 mars 2026, Nextcloud GmbH revendiquait plus de **2 millions de nouveaux utilisateurs entreprise** gagnés en un an, une traction qui confirme l’argument ayant poussé de nombreuses administrations à quitter les suites américaines, comme nous l’avons analysé dans notre dossier sur l’abandon de Teams et Zoom.

Au-delà de la conformité, l’intérêt est aussi financier. Nextcloud est sous licence AGPLv3 : le logiciel est gratuit, sans coût de licence par utilisateur. Vous ne payez que l’infrastructure – un serveur dédié chez un hébergeur français comme Scaleway ou OVHcloud, ou même une machine sur site. Pour un cloud souverain européen, l’équation est imbattable : aucune redevance mensuelle par siège, un stockage limité uniquement par votre disque, et un écosystème de plus de 200 applications (Calendrier, Contacts, Talk pour la visioconférence, Collabora pour la bureautique en ligne). C’est le socle modulaire basé sur des standards ouverts que recommandent explicitement les plans d’action 2026 pour les PME et ETI.

Ce guide privilégie l’approche **Docker Compose manuelle** plutôt que Nextcloud All-in-One (AIO), dont l’installeur Docker vient d’atteindre la version **13.6.0** (tag publié le 18 août 2026 par Nextcloud sur GitHub, 290ᵉ version publiée au total). AIO reste plus simple et opinionné – Nextcloud GmbH a d’ailleurs publié le 23 avril 2026 une mise à jour de son interface de gestion apportant de nouvelles options pour les installeurs Windows et Linux – mais Compose offre un contrôle total sur les versions, la topologie réseau, le reverse proxy et les services externes (base de données, cache) – un atout décisif quand on vise la maîtrise complète exigée par une démarche de souveraineté.

## Prérequis : versions et matériel pour héberger Nextcloud

Avant de commencer, voici les composants logiciels et leurs versions de référence à la mi-2026. La branche stable de production de Nextcloud est désormais la **35.x** (version **35.0.0**, diffusée via l’installeur web le 16 septembre 2026, selon Nextcloud GmbH), et les pages d’installation officielles – y compris leurs versions française et allemande, toutes deux mises à jour le 16 septembre 2026 – ciblent déjà cette version 35.0.0 pour toute nouvelle installation. À titre de comparaison, l’archive serveur complète de la branche précédente, `34.0.3.zip`, pesait 315 Mo lors de sa publication le 13 août 2026, selon l’index des releases du serveur Nextcloud – un poids qui donne une idée du volume à télécharger si vous optez pour une installation manuelle plutôt que Docker. Nous utilisons les images officielles Docker, ce qui évite justement toute installation manuelle de PHP ou de serveur web.

| Composant | Version / Référence | Rôle dans l’architecture | 
|---|---|---|
| Système d’exploitation | Ubuntu Server 24.04 LTS | Hôte du serveur (support jusqu’en 2029) | 
| Nextcloud Server | 33 (image `nextcloud:33-apache` ) | Application principale, branche stable | 
| Docker Engine | Dernière version stable (v27+) | Moteur de conteneurs | 
| Docker Compose | v2 (plugin `docker compose` ) | Orchestration multi-conteneurs | 
| PostgreSQL | 16 (image `postgres:16-alpine` ) | Base de données relationnelle | 
| Redis | 7 (image `redis:7-alpine` ) | Cache mémoire et verrouillage de fichiers | 
| PHP | 8.3 (intégré à l’image Nextcloud) | Runtime applicatif | 
| Caddy | 2 (image `caddy:2-alpine` ) | Reverse proxy + HTTPS Let’s Encrypt auto | 
| Collabora Online | CODE (image `collabora/code` ) | Édition collaborative de documents | 

Côté matériel, le dimensionnement dépend du nombre d’utilisateurs et des fonctionnalités activées (prévisualisation d’images, recherche plein texte, bureautique en ligne). Voici des paliers réalistes pour un **cloud souverain auto-hébergé**.

| Usage | vCPU | RAM | Stockage | Commentaire | 
|---|---|---|---|---|
| Test / 1 à 3 utilisateurs | 2 vCPU | 2 Go | 40 Go SSD | Minimum viable, sans Collabora | 
| Petite équipe (10-20 users) | 4 vCPU | 8 Go | 200 Go SSD NVMe | Recommandé avec Collabora | 
| PME (50 users) | 8 vCPU | 16 Go | 1 To SSD | Prévoir recherche plein texte | 
| Usage intensif | 8+ vCPU | 32 Go | 2 To+ SSD | Collabora + prévisualisations lourdes | 

Vous aurez aussi besoin d’un **nom de domaine** (par exemple `cloud.votreentreprise.fr`) pointant vers l’adresse IP publique de votre serveur via un enregistrement DNS de type A, ainsi que des **ports 80 et 443 ouverts** au pare-feu. Caddy a besoin du port 80 pour valider les certificats Let’s Encrypt et du port 443 pour servir le trafic HTTPS. Si vous débutez avec la conteneurisation, notre tutoriel Docker en 13 étapes pose les fondamentaux utiles avant de poursuivre.

## Étape 1 – Préparer et sécuriser le serveur Ubuntu

Connectez-vous à votre serveur Ubuntu 24.04 LTS en SSH. La première règle de toute démarche souveraine est de partir d’une base saine et à jour. Mettez à jour le système et installez les outils de base, puis configurez le pare-feu UFW pour n’exposer que le strict nécessaire : SSH, HTTP et HTTPS.

