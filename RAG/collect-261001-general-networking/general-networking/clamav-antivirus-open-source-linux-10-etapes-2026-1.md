---
id: collect-261001-general-networking/general-networking/clamav-antivirus-open-source-linux-10-etapes-2026-1
title: "Debian / Ubuntu"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "open source"]
source: docs/RAG/collect-261001-general-networking/clamav-antivirus-open-source-linux-10-etapes-2026.md
source_anchor: ""
source_lines: [1, 50]
sha256: c8623178e34117ea101168bc5553939ad5d69cd788da1b31f8e1294a496ff0d3
---

# Debian / Ubuntu

Le 7 août 2026, l’équipe de Cisco Talos publiait coup sur coup deux correctifs de sécurité pour ClamAV : les versions 1.5.4 et 1.4.6 LTS. Deux failles corrigées d’un coup, CVE-2026-20337 et CVE-2026-20345, toutes deux capables de provoquer une corruption mémoire lors de l’analyse d’archives ZIP ou d’images disque au format GPT. Pour un outil que beaucoup d’administrateurs considèrent comme acquis, tourné en tâche de fond sur leur passerelle mail depuis des années, c’est un rappel utile : ClamAV reste un logiciel vivant, activement maintenu, et qui mérite une configuration sérieuse plutôt qu’une installation par défaut oubliée dans un coin.

Ce tutoriel vous guide du premier `apt install` jusqu’à un déploiement complet de ClamAV en passerelle antivirus pour un serveur mail Postfix, avec surveillance, quarantaine et automatisation. Le contexte réglementaire y ajoute une urgence supplémentaire : selon le rapport annuel 2026 du ministère de l’Intérieur sur la cybercriminalité, la France a enregistré 453 200 cyberattaques en 2025, soit une hausse de 87 % en cinq ans. Pour les PME concernées par la directive NIS2, disposer d’un antivirus documenté sur les points d’entrée critiques (serveurs de mail, partages de fichiers, passerelles d’upload) n’est plus une option confortable mais une exigence de conformité.

ClamAV coche une case particulière : c’est un moteur antivirus open source, gratuit, maintenu par Cisco Talos, et intégré nativement dans la quasi-totalité des distributions Linux. Il n’a pas vocation à remplacer un EDR sur les postes de travail, mais sur un serveur mail, un NAS ou une passerelle de dépôt de fichiers, c’est une brique de défense en profondeur qui coûte zéro euro de licence. Voici comment l’installer, le configurer et l’exploiter correctement en 2026.

## Qu’est-ce que ClamAV et pourquoi l’installer en 2026

ClamAV est un moteur antivirus open source né à la fin des années 1990, aujourd’hui développé et maintenu par Cisco Talos, l’équipe de renseignement sur les menaces de Cisco. Contrairement aux suites commerciales grand public comme Bitdefender ou Kaspersky, ClamAV cible en priorité les serveurs : passerelles de messagerie, serveurs de fichiers, systèmes de dépôt de documents, NAS d’entreprise. Son fonctionnement repose sur une analyse par signatures, complétée par des heuristiques et la détection de format de fichiers malveillants (PDF piégés, macros Office, archives contenant des exécutables suspects).

Trois raisons expliquent pourquoi ClamAV reste pertinent en 2026 plutôt que d’être une relique technique. D’abord, le rythme de maintenance ne faiblit pas : entre juin 2025 et août 2026, l’équipe a publié au moins quatre vagues de correctifs de sécurité (1.4.3/1.0.9 en juin 2025, 1.5.3/1.4.5 début juillet 2026, puis 1.5.4/1.4.6 le 7 août 2026), corrigeant des vulnérabilités dans les parseurs PDF, UDF, PE et ZIP. Ensuite, le projet a entamé en décembre 2025 un effort de réduction de la taille des fichiers de signatures (`main.cvd` et `daily.cvd`), ce qui allège la consommation disque et RAM sur les machines modestes. Enfin, l’intégration native dans Debian, Ubuntu, RHEL et leurs dérivés en fait un choix quasi gratuit à déployer, sans négociation de licence ni portail client à gérer.

ClamAV n’est pas un remplaçant direct d’une solution EDR ou d’un antivirus commercial pour poste de travail. Son terrain de jeu naturel reste le scan côté serveur : filtrage des pièces jointes avant remise en boîte mail, analyse des fichiers déposés sur un partage Samba ou NFS, vérification des uploads sur une application web. Utilisé à cet endroit précis, dans une architecture de défense en profondeur, il rend un vrai service, gratuitement et sans dépendance à un éditeur tiers.

## Prérequis : versions, matériel et compétences requises

Avant de commencer, vérifiez que votre environnement correspond aux versions et ressources recommandées. ClamAV tourne sur la quasi-totalité des distributions Linux, ainsi que sur FreeBSD, macOS et Windows, mais ce tutoriel se concentre sur un déploiement Linux serveur, le cas d’usage le plus courant en entreprise.

| Élément | Recommandation 2026 | Remarque | 
|---|---|---|
| Distribution | Debian 12/13, Ubuntu 22.04/24.04, RHEL/Rocky 9 | Paquets à jour dans les dépôts officiels | 
| Version ClamAV | 1.5.4 (branche stable) ou 1.4.6 LTS | 1.4 LTS supportée jusqu’au 15 août 2027 | 
| RAM | 2 Go minimum, 4 Go recommandés | La base de signatures chargée en mémoire par clamd pèse plusieurs centaines de Mo | 
| Espace disque | 3 Go libres | Base virale + logs + quarantaine | 
| Accès réseau | Sortant HTTPS vers les miroirs de signatures | freshclam a besoin d’accéder aux CDN ClamAV | 
| Droits | Accès root ou sudo | Installation de paquets et services systemd | 
| Compétences | Ligne de commande Linux de base | Aucune connaissance de programmation requise | 

Pour la partie intégration Postfix (l’étape « projet complet » de ce tutoriel), vous aurez besoin d’un serveur mail Postfix déjà fonctionnel, avec accès aux fichiers `main.cf` et `master.cf`. Si vous partez de zéro, installez d’abord Postfix seul, vérifiez qu’il route correctement le courrier, puis revenez à cette section. Comptez environ 60 minutes pour l’ensemble du tutoriel, installation, configuration, tests et intégration mail comprises.

## Comprendre l’architecture de ClamAV (clamd, freshclam, clamscan, clamonacc)

ClamAV n’est pas un binaire unique mais un ensemble de quatre composants qui travaillent ensemble. Comprendre leur rôle respectif évite bien des erreurs de configuration par la suite.

### Le démon clamd et le scanner clamscan

**clamd** est le démon de scan qui tourne en arrière-plan. Il charge la base de signatures en mémoire une seule fois au démarrage, puis répond aux demandes de scan via un socket Unix ou TCP, ce qui le rend nettement plus rapide qu’un scan à froid pour chaque fichier. C’est le composant que vous connecterez à Postfix via Milter, ou que vous interrogerez depuis vos propres scripts.

**clamscan**, à l’inverse, est un outil en ligne de commande autonome qui recharge la base de signatures à chaque exécution. Plus lent sur de gros volumes, il reste idéal pour un scan ponctuel, un test rapide ou une tâche cron peu fréquente sur un petit répertoire.

### freshclam et la base de signatures

**freshclam** est le service de mise à jour automatique des signatures. Il interroge les miroirs officiels ClamAV et télécharge les fichiers `main.cvd`, `daily.cvd` et `bytecode.cvd`. Depuis le 16 décembre 2025, le projet a entamé un effort de réduction de la taille de ces fichiers, notamment via un retrait progressif des signatures obsolètes, ce qui allège la charge disque et RAM sur les petites machines.

Un cinquième composant, **clamonacc**, mérite d’être cité séparément : il s’appuie sur les mécanismes fanotify et inotify du noyau Linux pour scanner les fichiers au moment même où ils sont créés ou modifiés sur le disque, une forme de protection en temps réel proche de ce que proposent les antivirus grand public.

## Étape 1 : installer ClamAV sur Debian, Ubuntu et RHEL/Fedora

L’installation passe par le gestionnaire de paquets natif de votre distribution. Sur les systèmes basés Debian, un seul appel apt installe le moteur, le démon et l’outil de mise à jour :

