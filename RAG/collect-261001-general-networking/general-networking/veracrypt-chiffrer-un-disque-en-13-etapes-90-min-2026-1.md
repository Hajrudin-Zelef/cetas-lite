---
id: collect-261001-general-networking/general-networking/veracrypt-chiffrer-un-disque-en-13-etapes-90-min-2026-1
title: "Windows (PowerShell)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "incident", "open source"]
source: docs/RAG/collect-261001-general-networking/veracrypt-chiffrer-un-disque-en-13-etapes-90-min-2026.md
source_anchor: ""
source_lines: [1, 44]
sha256: 414e9087851881177f7b5534c9268b29e16a3b1408542a119712830a427c188d
---

# Windows (PowerShell)

Un ordinateur portable volé dans un train, un disque externe oublié dans un taxi, un serveur de sauvegarde qui finit sur le marché de l’occasion sans avoir été effacé correctement : dans chacun de ces cas, un simple mot de passe Windows ou macOS ne protège rien. Il suffit de démonter le disque et de le brancher ailleurs pour lire les fichiers en clair. Le chiffrement de disque change complètement la donne, et VeraCrypt reste en 2026 l’outil open source de référence pour le faire gratuitement, sur Windows, macOS et Linux. Ce tutoriel vous accompagne pas à pas, de l’installation jusqu’à la création d’un volume caché et d’un script de sauvegarde chiffrée automatisé.

## Qu’est-ce que VeraCrypt, et pourquoi chiffrer son disque en 2026 ?

VeraCrypt est un logiciel gratuit et open source de chiffrement de disque, disponible pour Windows, macOS et Linux. Le projet est développé par IDRIX et s’appuie sur le code de TrueCrypt 7.1a, le célèbre outil de chiffrement abandonné en 2014 dans des circonstances jamais totalement élucidées. Plutôt que de repartir de zéro, les développeurs de VeraCrypt ont repris cette base et corrigé les faiblesses identifiées lors de l’audit indépendant mené entre 2013 et 2016 par l’Open Crypto Audit Project (OSTIF) et le cabinet QuarksLab. Cet audit n’a pas révélé de porte dérobée volontaire, mais a mis au jour plusieurs vulnérabilités que VeraCrypt a corrigées au fil des versions.

Concrètement, VeraCrypt permet trois choses : créer un conteneur chiffré (un fichier qui se comporte comme un disque virtuel une fois monté), chiffrer une partition ou un disque entier de données, ou chiffrer le disque système lui-même, y compris Windows et ses fichiers temporaires. Un renforcement technique notable distingue VeraCrypt de TrueCrypt : le dérivé de clé pour les partitions système utilise désormais 200 000 itérations, contre seulement 1 000 sur TrueCrypt, ce qui complique nettement les attaques par force brute sur le mot de passe, au prix d’un temps de déverrouillage légèrement plus long.

Pourquoi ce sujet redevient-il d’actualité en France et en Europe en 2026 ? D’abord parce que le règlement général sur la protection des données (RGPD) cite explicitement le chiffrement à l’article 32 comme l’une des mesures techniques envisageables pour protéger les données personnelles, aux côtés de la pseudonymisation. Ensuite parce que la directive européenne NIS2 place elle aussi la cryptographie parmi les mesures de gestion des risques que les entités concernées doivent envisager, une liste qui touche en France plusieurs milliers d’organisations une fois la transposition pleinement appliquée. Ni le RGPD ni NIS2 n’imposent VeraCrypt spécifiquement ni même le chiffrement en toutes circonstances : les deux textes raisonnent par une approche fondée sur le risque, où le chiffrement figure comme une mesure « appropriée » plutôt qu’une obligation absolue et universelle. Mais dans les faits, un disque non chiffré rend la question du respect de ces textes beaucoup plus difficile à défendre en cas de contrôle ou d’incident.

## Prérequis : versions, configuration et éléments à préparer

Avant de vous lancer, vérifiez que votre configuration correspond à ces prérequis. Ce tutoriel s’appuie sur la version stable la plus récente de VeraCrypt au moment de la rédaction, la 1.26.29, dont les fichiers d’installation sont apparus sur la plateforme de distribution SourceForge le 13 juin 2026, succédant à la 1.26.24 qui y était référencée depuis le 26 avril 2026.

- **VeraCrypt** : version 1.26.29 (publiée le 9 juin 2026), qui ajoute la prise en charge d’Argon2id pour les volumes non-système et corrige deux failles de sécurité
- **Windows** : Windows 11 ou Windows 10 à partir de la version 1809 ; Windows Server 2019 ou ultérieur. Les versions antérieures de Windows nécessitent VeraCrypt 1.26.15 (dernière à supporter le 32 bits et Windows Server 2016) ou 1.25.9 (dernière à supporter XP, Vista, 7, 8 et 8.1)
- **macOS** : macOS 12 Monterey ou ultérieur (pour macOS 10.9 à 11 Big Sur, utilisez VeraCrypt 1.25.9)
- **Linux** : architectures x86, x86-64 et ARM64, la plupart des distributions récentes (Debian, Ubuntu, Fedora, Arch)
- **Autres systèmes** : FreeBSD x86-64 à partir de la version 14, OpenBSD x86-64 à partir de la version 7.8, Raspberry Pi OS en 32 et 64 bits
- **Espace disque libre** : au minimum l’équivalent de la taille du conteneur que vous souhaitez créer, plus 10 % de marge
- **Une clé USB vierge** d’au moins 100 Mo, dédiée au disque de secours (Rescue Disk) et, si vous le souhaitez, à un fichier-clé
- **Une sauvegarde complète** de vos données importantes sur un support externe déconnecté, avant toute opération de chiffrement du disque système
- **L’ordinateur portable branché au secteur** pendant toute la durée du chiffrement initial, qui peut prendre de quelques minutes à plusieurs heures selon la taille et la vitesse du disque

Comptez environ 90 minutes pour suivre l’intégralité de ce tutoriel en 13 étapes, en incluant le temps d’attente du chiffrement initial d’un petit volume de test. Le chiffrement complet d’un disque système de plusieurs centaines de gigaoctets se poursuivra en tâche de fond bien après la fin de la lecture.

## VeraCrypt face à BitLocker et FileVault : quelles différences réelles ?

BitLocker (Windows) et FileVault (macOS) sont déjà intégrés à votre système d’exploitation, ce qui les rend plus simples à activer. Mais ce sont des solutions propriétaires : leur code source n’est pas public, et leur fonctionnement interne ne peut être vérifié que par les éditeurs eux-mêmes ou des chercheurs travaillant par rétro-ingénierie. VeraCrypt, à l’inverse, publie l’intégralité de son code source, ce qui permet un audit communautaire indépendant, comme celui mené sur la base TrueCrypt. Le tableau suivant résume les différences pratiques entre les trois solutions.

| Critère | VeraCrypt | BitLocker | FileVault | 
|---|---|---|---|
| Licence | Open source, gratuit | Propriétaire, inclus dans Windows Pro/Enterprise/Education | Propriétaire, inclus dans macOS | 
| Plateformes | Windows, macOS, Linux, FreeBSD, OpenBSD, Raspberry Pi OS | Windows uniquement | macOS uniquement | 
| Algorithme principal | AES, Serpent, Twofish, Camellia, Kuznyechik (cascades possibles) | AES-CBC ou AES-XTS, 128 ou 256 bits | XTS-AES 128 bits | 
| Volumes cachés / déni plausible | Oui | Non | Non | 
| Dépendance à une puce TPM | Non requise | Recommandée, quasi obligatoire pour le déverrouillage automatique | Non requise (clé dérivée du mot de passe) | 
| Auditabilité du code | Code source public, audit communautaire possible | Code fermé | Code fermé | 
| Conteneurs de fichiers chiffrés | Oui, indépendamment du disque système | Non (détournement via VHD chiffré possible) | Non nativement | 

Aucune de ces solutions n’est meilleure dans l’absolu : BitLocker s’intègre mieux dans un parc Windows géré via Active Directory et Intune, FileVault est le choix par défaut cohérent sur un Mac. VeraCrypt s’impose dès que vous avez besoin de portabilité entre systèmes, de conteneurs chiffrés indépendants du disque système, ou de la fonctionnalité de volume caché qu’aucune des deux alternatives propriétaires ne propose.

## Comprendre les algorithmes de chiffrement et de hachage de VeraCrypt

