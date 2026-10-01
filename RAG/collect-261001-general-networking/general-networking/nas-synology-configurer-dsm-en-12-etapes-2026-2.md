---
id: collect-261001-general-networking/general-networking/nas-synology-configurer-dsm-en-12-etapes-2026-2
title: "Exemple de planification de tâche Hyper Backup (via l'interface DSM)"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-general-networking/nas-synology-configurer-dsm-en-12-etapes-2026.md
source_anchor: ""
source_lines: [43, 102]
sha256: 44ede5fc72e8831e99ec5e439a9f8f38caef8f67f7cb1ec5cf483f4d0313cfa0
---

# Exemple de planification de tâche Hyper Backup (via l'interface DSM)

**Point de sécurité important** : Synology a publié l’avis **Synology-SA-26:12** le 3 août 2026, concernant une faille CVE-2026-4793 dans Synology Assistant pour Windows, qui permettait à un utilisateur local de lire ou écrire des fichiers arbitraires et de provoquer un déni de service. Avant de télécharger l’outil, vérifiez que vous récupérez bien la version 7.0.7-50095 ou plus récente, uniquement depuis le site officiel Synology.

Si aucun outil n’est disponible, vous pouvez aussi accéder directement au NAS via son adresse IP locale, port 5000 en HTTP ou 5001 en HTTPS, par exemple :

```
http://192.168.1.100:5000
https://192.168.1.100:5001
```
L’adresse IP exacte se trouve dans l’interface d’administration de votre routeur (liste des appareils connectés), ou directement dans les résultats du scan Synology Assistant.

## Étape 3 : Installer DSM 7.4

Une fois le NAS détecté, cliquez sur “Configurer” ou “Installer” en suivant les mêmes étapes que le guide officiel d’installation de DSM. L’assistant propose de télécharger automatiquement la dernière version stable de DSM — en août 2026, il s’agit de **DSM 7.4.1-90080**, publiée le 23 juillet 2026. Le téléchargement et l’installation prennent entre 5 et 15 minutes selon la vitesse de votre connexion internet, le NAS redémarrant une ou deux fois pendant le processus.

Pendant cette phase, ne débranchez jamais le NAS et ne fermez pas la fenêtre du navigateur. Une coupure d’alimentation pendant l’écriture du firmware peut rendre le NAS inutilisable et nécessiter une réinstallation via un mode de récupération plus contraignant.

Si vous préférez rester sur une branche de maintenance longue durée plutôt que sur la dernière version générale, sachez que Synology a basculé le support LTS sur **DSM 7.3** depuis octobre 2025, avec une maintenance prévue jusqu’en octobre 2027 et une fin de vie annoncée pour octobre 2028 ; l’ancienne branche **DSM 7.2**, disponible depuis juin 2023, voit de son côté sa maintenance s’arrêter en décembre 2025, même si des correctifs ponctuels comme **DSM 7.2.2-72806 Update 9** (30 juin 2026) ou le plus ancien **DSM 7.2-64570 Update 4** (11 février 2025) restent disponibles pour les modèles qui ne supportent pas toujours DSM 7.4. Vérifiez la compatibilité de votre modèle exact sur la page officielle des notes de version DSM avant de choisir.

## Étape 4 : Créer le compte administrateur et sécuriser l’accès

DSM demande ensuite de nommer le serveur (par exemple “NAS-Maison” ou “NAS-Bureau”) et de créer le compte administrateur. Trois règles à respecter absolument à cette étape :

- Ne gardez jamais le compte “admin” par défaut actif après l’installation — créez un compte administrateur avec un nom personnalisé, puis désactivez le compte “admin” générique dans Panneau de configuration > Utilisateur.
- Utilisez un mot de passe d’au moins 12 caractères avec majuscules, chiffres et symboles ; DSM refuse les mots de passe trop simples mais un gestionnaire de mots de passe reste préférable pour en générer un vraiment robuste.
- Activez l’authentification à deux facteurs (2FA) dès cette étape via Panneau de configuration > Compte utilisateur > Authentification à 2 facteurs, avec une application comme Google Authenticator ou Microsoft Authenticator.

Ces réglages ne sont pas cosmétiques : les NAS exposés sur internet sans 2FA ni changement du compte admin par défaut restent une cible privilégiée pour les campagnes de ransomware ciblant spécifiquement les NAS grand public, un vecteur d’attaque documenté à plusieurs reprises ces dernières années. L’ANSSI, l’agence française de cybersécurité, recommande systématiquement l’activation du 2FA et la suppression des comptes génériques sur tout équipement exposé, une règle qui s’applique aussi bien en entreprise qu’à la maison.

## Étape 5 : Choisir entre Btrfs et ext4

C’est la décision la plus structurante du tutoriel, car elle ne se change pas sans tout reformater. DSM 7.4 propose deux systèmes de fichiers pour le volume principal :

| Critère | Btrfs | ext4 | 
|---|---|---|
| Snapshots natifs | Oui, quasi instantanés | Non | 
| Checksums de données | Oui (détecte la corruption silencieuse) | Non | 
| Versionning Synology Drive | Plus économe en espace disque | Nécessite un espace de stockage double pour les métadonnées de version | 
| Quotas par dossier partagé | Oui | Non | 
| Compatibilité LUN iSCSI sur disque externe | Non pour Hyper Backup | Requis pour la sauvegarde de LUN iSCSI vers disque externe | 
| Charge CPU | Légèrement supérieure | Plus légère, utile sur NAS d’entrée de gamme | 

Pour la quasi-totalité des usages domestiques ou de petit bureau, **Btrfs est le choix recommandé** en 2026 : les snapshots protègent contre les erreurs de manipulation et les débuts de ransomware, et les checksums repèrent la corruption de données avant qu’elle ne devienne un vrai problème. Depuis DSM 7.4, l’outil Storage Efficiency affiche en plus, dossier partagé par dossier partagé, le gain réel obtenu par la déduplication et la compression a posteriori des volumes HDD, exprimé directement en pourcentage d’espace récupéré. ext4 garde un intérêt pour les configurations avec CPU limité (vieux modèles d’entrée de gamme) ou pour des besoins spécifiques de compatibilité iSCSI avec Hyper Backup vers un disque externe.

Si vous avez déjà un volume ext4 existant et voulez migrer vers Btrfs, Synology ne propose pas de conversion en place : il faut sauvegarder les données avec Hyper Backup, supprimer le volume ext4, créer un nouveau volume Btrfs, puis restaurer depuis la sauvegarde. Prévoyez le temps et l’espace de stockage nécessaires avant de vous lancer.

## Étape 6 : Créer le pool de stockage et choisir le type de RAID

Direction Gestionnaire de stockage > Stockage > Créer. DSM propose plusieurs schémas de redondance selon le nombre de disques installés :

- **SHR (Synology Hybrid RAID)** : le format maison de Synology, qui permet de mélanger des disques de tailles différentes sans perdre trop d’espace. SHR-1 tolère la panne d’un disque (équivalent RAID 5), SHR-2 tolère deux pannes simultanées (équivalent RAID 6) mais demande au moins 4 disques.
- **RAID 1** : miroir simple, recommandé sur un NAS 2 baies pour une protection basique contre la panne d’un disque.
- **RAID 5** : nécessite 3 disques minimum, tolère une panne, bon compromis capacité/sécurité sur 4 baies.
- **RAID 6** : nécessite 4 disques minimum, tolère deux pannes simultanées, recommandé si le NAS stocke des données critiques sans sauvegarde externe immédiate.

Pour un premier NAS 2 baies, SHR-1 ou RAID 1 sont les deux options les plus sûres. La différence pratique : RAID 1 exige deux disques strictement identiques en capacité, alors que SHR-1 accepte un disque de 4 To et un disque de 6 To en n’utilisant “que” 4 To exploitables sur le second, ce qui laisse une marge d’évolution si vous rachetez un disque plus gros plus tard.

## Étape 7 : Configurer QuickConnect et l’accès distant

QuickConnect est le service de relais et de DNS dynamique de Synology, qui permet d’accéder à DSM et aux applications (Drive, Photos, etc.) depuis l’extérieur du réseau local sans avoir à configurer manuellement une redirection de port sur la box internet. Le paramétrage se fait dans Panneau de configuration > QuickConnect : connectez-vous avec votre compte Synology, choisissez un identifiant QuickConnect unique, et activez les applications que vous souhaitez rendre accessibles.

