---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/installer-opnsense-pare-feu-open-source-en-14-etapes-2026-1
title: "Sous Linux, décompression de l'image téléchargée"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "open source"]
source: docs/RAG/collect-261001-opnsense-pfsense/installer-opnsense-pare-feu-open-source-en-14-etapes-2026.md
source_anchor: ""
source_lines: [1, 49]
sha256: afacec3a4222882ccccee238cb5626a229ea47884d5c200d65279fff4657fb81
---

# Sous Linux, décompression de l'image téléchargée

OPNsense vient de passer un cap. La version 26.7.3 “Xenial Xenops” est sortie le 27 août 2026, à peine un jour avant la publication de ce guide, et elle embarque FreeBSD 15.1, PHP 8.5 et un moteur de règles de pare-feu entièrement réécrit en MVC. Pour les administrateurs réseau qui cherchent une alternative open source crédible aux pare-feux propriétaires (Fortinet, Sophos, Cisco) ou à pfSense, c’est le bon moment pour s’y mettre. Ce tutoriel couvre l’installation complète, du téléchargement de l’ISO jusqu’à la mise en place du VPN WireGuard, de la détection d’intrusion Suricata et de la haute disponibilité CARP, avec chaque commande et chaque piège documenté.

OPNsense équipe aujourd’hui des dizaines de milliers de box faites maison, de PME et même certaines collectivités qui refusent la dépendance à un éditeur unique. Le projet, porté par l’entreprise néerlandaise Deciso, publie des mises à jour à un rythme largement supérieur à celui de pfSense CE, dont la dernière version communautaire (2.8.1) date de septembre 2025. Voici comment monter votre propre pare-feu OPNsense en partant de zéro, avec un niveau de détail suffisant pour éviter les erreurs qui bloquent la moitié des installations débutantes.

## Qu’est-ce qu’OPNsense et pourquoi l’adopter en 2026

OPNsense est un système d’exploitation de pare-feu et de routage open source, basé sur FreeBSD, né en 2015 d’une scission (fork) du projet pfSense. Les développeurs à l’origine du fork reprochaient à pfSense un manque de transparence dans le développement et une gouvernance trop centralisée autour de Netgate. Onze ans plus tard, OPNsense s’est imposé comme un projet à part entière avec sa propre base de code, son interface web et un calendrier de sortie indépendant remarquablement régulier : la série 25.1 (dotée d’une option d’extinction dans l’installateur) est sortie le 29 janvier 2025, suivie de 25.7 “Visionary Viper” le 23 juillet 2025 puis de 26.1 “Witty Woodpecker” le 28 janvier 2026, selon la documentation officielle du projet, soit une nouvelle branche majeure quasiment tous les six mois.

La branche communautaire (Community Edition, CE) actuelle est la série 26.7 “Xenial Xenops”, sortie le 15 juillet 2026, stabilisée par une première mise à jour 26.7.1 dès le 21 juillet 2026 puis portée à la version 26.7.3 le 27 août 2026 (la dernière en date selon Wikipedia). Elle apporte une base FreeBSD 15.1, l’intégration d’OpenVPN 2.7, un environnement PHP 8.5 pour l’interface web, un moteur de règles de pare-feu MVC désormais activé par défaut, ainsi qu’un support IPv6 amélioré. Pour rappel, la branche précédente 26.1 “Witty Woodpecker” avait elle-même reçu pas moins de onze mises à jour mineures jusqu’à la 26.1.11 du 1er juillet 2026, d’après le calendrier de versions tenu par endoflife.date, ce qui illustre le rythme de correctifs très soutenu du projet. Deciso maintient en parallèle une édition Business (BE), dont la version 26.4.2 a été publiée le 14 août 2026, destinée aux entreprises qui veulent un support contractuel et un rythme de mise à jour plus stable.

Contrairement à pfSense, qui distingue une édition Community (CE) gratuite et une édition Plus payante réservée aux appliances Netgate, OPNsense garde l’intégralité de son code sous licence open source de type BSD, y compris pour l’édition Business. La différence entre CE et BE tient surtout au calendrier de publication et au niveau de support, pas à des fonctionnalités verrouillées derrière un paywall. C’est un argument de poids pour les administrateurs qui veulent auditer leur pare-feu de bout en bout.

## Prérequis matériels et logiciels

Avant de se lancer, il faut choisir la bonne machine. OPNsense tourne sur toute architecture x86-64, mais les besoins réels varient énormément selon les fonctionnalités activées. La documentation officielle distingue plusieurs profils de dimensionnement selon le débit visé et les modules utilisés.

| Usage | CPU | RAM | Stockage | Débit indicatif | 
|---|---|---|---|---|
| Pare-feu basique uniquement | 1 GHz mono/dual-core 64 bits | 2 Go min. / 4 Go recommandé | 4-8 Go (image nano) | 11-150 Mbps | 
| Pare-feu + DNS + DHCP | 1 GHz dual-core | 4 Go min. / 8 Go recommandé | 16-40 Go SSD | 151-350 Mbps | 
| Pare-feu + Suricata IDS/IPS | 1,5 GHz multi-core (Intel J6412 conseillé) | 8 Go min. / 16 Go recommandé | 40-120 Go SSD | 350-750+ Mbps | 
| Toutes fonctions (VPN + IDS + HA + monitoring) | Quad-core 1,6 GHz+ | 16 Go min. / 32 Go recommandé | 120 Go SSD NVMe | 750 Mbps et plus | 

Pour une utilisation domestique ou en petite entreprise avec une connexion fibre à 500 Mbps ou 1 Gbps, un mini-PC équipé d’un processeur Intel de la gamme N100/N305 ou d’un ancien SFF (Small Form Factor) avec 8 Go de RAM et deux cartes réseau Intel (pilotes igb/em, réputés pour leur stabilité sous FreeBSD) constitue un excellent point de départ. Évitez les cartes réseau Realtek premier prix : les pilotes FreeBSD associés sont historiquement moins bien supportés et provoquent des pertes de paquets sous forte charge.

- Machine x86-64 avec au minimum 2 interfaces réseau (une pour le WAN, une pour le LAN)
- 4 Go de RAM minimum, 8 Go recommandé si vous activez Suricata
- 16 à 40 Go de stockage SSD (évitez les clés USB en usage permanent, leur endurance en écriture est trop faible)
- Une clé USB de 8 à 16 Go pour créer le support d’installation
- Un poste de travail avec Rufus (Windows) ou balenaEtcher (Windows/macOS/Linux) pour graver l’image
- Un accès physique ou une console série pour l’installation initiale
- Cartes réseau Intel de préférence, pour la compatibilité des pilotes FreeBSD

## Étape 1 : télécharger l’image ISO d’OPNsense

Rendez-vous sur la page de téléchargement officielle du projet et sélectionnez l’image DVD amd64. Pour la série 26.7, le fichier porte le nom `OPNsense-26.7-dvd-amd64.iso.bz2`. Point important souvent oublié : l’image est compressée en bzip2, il faut donc la décompresser avant de la graver sur la clé USB. Une erreur classique consiste à écrire directement le fichier `.bz2` sur la clé, ce qui produit un support non amorçable. Notez aussi qu’une fois décompressée, l’image ISO occupe environ 3 Go sur la clé USB, mais le système de fichiers se redimensionne automatiquement pour utiliser tout l’espace disponible dès le premier démarrage sur le disque de destination, une précision apportée par le README des miroirs de téléchargement d’OPNsense depuis août 2025.

```
# Sous Linux, décompression de l'image téléchargée
bunzip2 OPNsense-26.7-dvd-amd64.iso.bz2
# Vérification de l'intégrité avec la somme de contrôle SHA256 publiée sur opnsense.org
sha256sum OPNsense-26.7-dvd-amd64.iso
```
Vérifiez systématiquement la somme de contrôle SHA256 fournie sur le site officiel avant de graver l’image. Cette étape, souvent négligée, permet d’éviter d’installer un firmware corrompu ou modifié, particulièrement critique pour un équipement qui filtre tout le trafic réseau de votre organisation.

## Étape 2 : créer une clé USB bootable avec Rufus ou Etcher

Sous Windows, ouvrez Rufus, sélectionnez votre clé USB dans le menu déroulant “Périphérique”, puis cliquez sur “Sélection” pour pointer vers le fichier ISO décompressé. Laissez le schéma de partition sur GPT si votre machine cible démarre en UEFI (le cas le plus fréquent en 2026), ou MBR pour un BIOS Legacy. Lancez l’écriture et patientez : l’opération prend entre 3 et 8 minutes selon la vitesse de la clé.

