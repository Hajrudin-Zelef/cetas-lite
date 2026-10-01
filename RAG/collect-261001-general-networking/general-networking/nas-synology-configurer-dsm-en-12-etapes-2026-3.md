---
id: collect-261001-general-networking/general-networking/nas-synology-configurer-dsm-en-12-etapes-2026-3
title: "Exemple de planification de tâche Hyper Backup (via l'interface DSM)"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google", "Microsoft"]
dates: []
keywords: ["ethernet", "mai"]
source: docs/RAG/collect-261001-general-networking/nas-synology-configurer-dsm-en-12-etapes-2026.md
source_anchor: ""
source_lines: [103, 161]
sha256: d05a3ad8cdb6d4105012e36158c3ffa4a3735cc394743a1680e9b4150f7eac17
---

# Exemple de planification de tâche Hyper Backup (via l'interface DSM)

DSM 7.2 et 7.4 ont renforcé la sécurité de ce service : Auto Block bannit désormais automatiquement les adresses IP après plusieurs tentatives de connexion échouées sur les serveurs QuickConnect, en plus de la protection déjà existante sur l’accès local. Ce point mérite d’être vérifié manuellement dans Panneau de configuration > Sécurité > Protection du compte, car le seuil par défaut (généralement 5 tentatives) peut être resserré pour un NAS exposé publiquement.

Alternative plus contrôlée : configurer un accès VPN vers le réseau local (via le paquet VPN Server de Synology ou un routeur compatible WireGuard) plutôt que d’exposer DSM directement sur internet. Cette approche rejoint la logique de l’architecture Zero Trust appliquée à un contexte domestique : on limite la surface d’attaque en n’exposant que ce qui est strictement nécessaire.

## Étape 8 : Mettre en place la sauvegarde avec Hyper Backup

Un NAS n’est pas une sauvegarde en soi — c’est un poste de stockage. Sans copie externe, un incendie, un vol ou un ransomware qui chiffre le volume local emporte tout. Synology propose **Hyper Backup**, un paquet gratuit qui gère la sauvegarde par blocs vers plusieurs types de destinations : un autre NAS Synology, un disque USB externe, ou le service cloud Synology C2 Storage.

Installation depuis le Centre de paquets, puis création d’une tâche de sauvegarde qui cible les dossiers partagés et les paramètres système à protéger. Un point technique à connaître : Hyper Backup fonctionne avec les volumes Btrfs et ext4, mais si la destination est un disque externe utilisé pour sauvegarder des LUN iSCSI, ce disque externe doit obligatoirement être formaté en ext4.

```
# Exemple de planification de tâche Hyper Backup (via l'interface DSM)
Destination : Synology C2 Storage ou NAS distant
Fréquence : quotidienne, 03:00
Rétention : rotation par versions (ex. 7 quotidiennes, 4 hebdomadaires, 6 mensuelles)
Compression : activée
Chiffrement : activé avec phrase de passe séparée
```
Sur la tarification cloud, Synology C2 Storage propose en 2026 des paliers Basic en Europe autour de **9,99 € par an pour 100 Go**, **24,99 € par an pour 300 Go** et **59,99 € par an pour 1 To** (hors TVA), tandis que C2 Backup pour particuliers démarre autour de **34,99 $/an pour 500 Go** avec 15 Go offerts au départ. Pour une structure professionnelle, C2 Backup Business se facture environ **11,99 € par To et par mois** en Europe.

La règle à retenir pour une vraie stratégie de sauvegarde robuste reste la méthode 3-2-1-1-0 : trois copies des données, sur deux supports différents, dont une hors site, avec une copie hors ligne ou immuable, et zéro erreur vérifiée régulièrement. Un NAS Synology avec Hyper Backup vers C2 Storage coche déjà trois de ces cinq cases dès la configuration initiale.

## Étape 9 : Ajouter un cache SSD pour accélérer les accès fréquents

Sur les modèles équipés d’emplacements M.2 NVMe (DS423+, DS923+ et modèles supérieurs), il est possible de configurer un cache en lecture ou en lecture/écriture. Ce cache accélère l’accès aux fichiers consultés fréquemment (miniatures Synology Photos, indexation Synology Drive, bases de données de petites applications) sans passer systématiquement par les disques durs mécaniques, plus lents en accès aléatoire.

Configuration dans Gestionnaire de stockage > SSD Cache > Créer. Le cache en lecture seule est sans risque (une panne du SSD n’affecte pas l’intégrité des données), tandis que le cache en lecture/écriture améliore aussi les performances d’écriture mais nécessite deux SSD en miroir pour éviter la perte de données en cas de panne du cache lui-même. Pour un usage domestique standard, le cache en lecture seule suffit largement et reste la configuration la plus simple à mettre en place.

## Étape 10 : Installer Synology Drive et Synology Photos

Une fois le stockage et la sauvegarde en place, les deux applications les plus utilisées au quotidien sont Synology Drive (synchronisation de fichiers façon Dropbox, avec historique de versions) et Synology Photos (galerie photo avec reconnaissance faciale et albums partagés, la suite de Moments sur les anciennes versions de DSM).

Installation via le Centre de paquets, puis configuration des dossiers à synchroniser depuis Panneau de configuration au sein de chaque application. Sur un volume Btrfs, l’historique de versions de Synology Drive consomme sensiblement moins d’espace disque que sur ext4, grâce à la copie sur écriture (copy-on-write) qui évite de dupliquer intégralement les métadonnées à chaque nouvelle version d’un fichier.

Les applications clientes existent pour Windows, macOS, Linux, Android et iOS, avec synchronisation automatique en arrière-plan. Pour un usage professionnel, cette combinaison représente une alternative crédible et auto-hébergée aux abonnements Google Workspace ou Microsoft 365 pour la partie stockage de fichiers, dans la même logique que Nextcloud en cloud souverain.

## Étape 11 : Optimiser le réseau pour de meilleurs débits

La plupart des NAS Synology d’entrée et milieu de gamme embarquent deux ports Ethernet 1 GbE, qui peuvent être agrégés (link aggregation, 802.3ad) pour atteindre un débit cumulé plus élevé — à condition que le routeur ou switch le supporte également. Configuration dans Panneau de configuration > Réseau > Interface réseau > Créer > Créer une liaison agrégée.

Si votre installation réseau domestique reste en Wi-Fi 6 ou passe au Wi-Fi 7, le NAS lui-même reste presque toujours raccordé en filaire pour des raisons de fiabilité et de débit soutenu — voir notre comparatif Wi-Fi 7 vs Wi-Fi 6 pour la partie réseau sans fil du reste de la maison. Un transfert initial de plusieurs centaines de gigaoctets vers un NAS neuf se fait presque toujours plus vite et de façon plus fiable via câble Ethernet direct plutôt qu’en Wi-Fi, où les micro-coupures peuvent interrompre de longs transferts.

## Étape 12 : Sécuriser durablement le NAS et suivre les mises à jour

Synology publie des avis de sécurité réguliers, et 2026 n’a pas fait exception. Le CERT-FR avait d’ailleurs déjà alerté en mai 2025 sur des vulnérabilités touchant DSM 7.2.1 dans toutes les versions antérieures à 7.2.1-69057-7, un rappel utile que ces failles ne datent pas seulement de cette année. Plusieurs correctifs importants sont sortis en 2026, avec un rythme moyen d’une à deux publications par mois sur le second trimestre :

| Avis | Date | Composant | Correctif | 
|---|---|---|---|
| Synology-SA-26:03 | 19 mars 2026 | telnetd DSM (GNU inetutils) | CVE-2026-32746, exécution de commande à distance jugée critique | 
| Synology-SA-26:06/07 | 15 avril 2026 | DSM (plusieurs modules) | Mise à jour vers 7.3-81180 ou 7.2.2-72806-7 | 
| Synology-SA-26:01 | 27 mai 2026 | Storage Manager | CVE-2026-2237, mise à jour vers 1.0.1-1100 | 
| Synology-SA-26:10 | 26 mai 2026 | Chat Server | Mise à jour vers 2.4.5-22148 | 
| Synology-SA-26:11 | 26 juin 2026 | MailPlus Server | Mise à jour vers 4.0.1-31663 | 
| Synology-SA-26:12 | 3 août 2026 | Synology Assistant (Windows) | CVE-2026-4793, mise à jour vers 7.0.7-50095 | 

Ce rythme de correctifs n’a rien d’inhabituel pour un système exposé à internet, mais il impose une discipline simple : activer les mises à jour automatiques de sécurité dans Panneau de configuration > Mise à jour et restauration, et vérifier manuellement la page des avis de sécurité Synology une fois par trimestre si le NAS gère des données sensibles. La liste complète des avis reste consultable sur la page officielle des avis de sécurité Synology.

## Projet complet : monter un NAS familial avec sauvegarde automatisée

