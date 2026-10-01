---
id: collect-261001-general-networking/general-networking/durcir-windows-server-2025-credential-guard-2026-1
title: "Vérifier la virtualisation matérielle activée dans le firmware"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "datacenter", "incident", "intel"]
source: docs/RAG/collect-261001-general-networking/durcir-windows-server-2025-credential-guard-2026.md
source_anchor: ""
source_lines: [1, 47]
sha256: 545015e10c768a6ec597b95f271f1a650ca046078b13ac42aa020356e8c9f060
---

# Vérifier la virtualisation matérielle activée dans le firmware

Depuis le Patch Tuesday d’août 2026, Microsoft a corrigé 751 failles, dont la CVE-2026-62784, une vulnérabilité d’exécution de code à distance touchant directement le processus LSASS sur lequel repose toute la protection des identifiants Windows. Le même correctif KB5120233 a aussi mis fin, selon des confirmations d’utilisateurs remontées sur Microsoft Q&A le 14 août 2026, à des dysfonctionnements persistants de Remote Credential Guard qui cassaient l’authentification Kerberos en double saut (double-hop) sur les serveurs membres. Sur les parcs Windows Server encore configurés avec les réglages par défaut, ce type de faille ouvre une porte royale au vol d’identifiants et aux mouvements latéraux. Windows Server 2025 change pourtant la donne : depuis novembre 2025, selon la documentation Microsoft Learn, Credential Guard est activé par défaut sur les serveurs membres éligibles, et un nouveau module PowerShell baptisé OSConfig permet d’appliquer et de surveiller des baselines de sécurité sans dépendre uniquement des stratégies de groupe. Ce tutoriel détaille, étape par étape, comment durcir Windows Server 2025 en combinant Credential Guard, la sécurité basée sur la virtualisation (VBS), les recommandations ANSSI pour les serveurs membres Active Directory et les nouveaux outils d’automatisation de Microsoft.

L’objectif n’est pas de cocher une case de conformité, mais de fermer la voie d’attaque la plus rentable pour un opérateur de ransomware en 2026 : le vol d’identifiants depuis la mémoire LSASS, suivi d’un pass-the-hash ou d’un pass-the-ticket vers le contrôleur de domaine. Comptez environ 120 minutes pour dérouler l’intégralité des 13 étapes sur un serveur de test, plus une phase de validation en environnement de préproduction avant tout déploiement en production.

## Pourquoi durcir Windows Server 2025 est devenu prioritaire en 2026

Windows Server 2025 est la version LTSC qui succède à Windows Server 2022, avec une intégration beaucoup plus profonde de la sécurité basée sur la virtualisation. Credential Guard lui-même n’est pas propre à cette édition : la documentation Microsoft Learn confirme, depuis juin 2025, sa compatibilité avec Windows Server 2016, 2019, 2022 et 2025. Ce qui change avec la version 2025, c’est le comportement par défaut : sur les machines qui remplissent les prérequis matériels et de licence (UEFI, Secure Boot, TPM, jointes à un domaine, hors contrôleur de domaine), Credential Guard démarre automatiquement, sans intervention de l’administrateur, un comportement que Microsoft Learn documentait encore le 27 avril 2026 et que l’équipe iPurple Team avait déjà signalé le 17 mars 2026 comme hérité de Windows 11 22H2. Un tour d’horizon myITforum de mars 2026 confirmait d’ailleurs que les nouveaux déploiements de Windows Server 2025 activent Credential Guard par défaut via VBS pour isoler les identifiants, et depuis le 23 avril 2026, les montées de version vers Windows Server 2025 héritent elles aussi de cette activation par défaut, sauf désactivation explicite. C’est une rupture avec Windows Server 2022, où l’activation restait manuelle et souvent oubliée lors des déploiements pressés.

Le contexte de menace rend ce changement urgent. Le rapport de mise à jour du 11 août 2026 (KB5120233) liste plusieurs failles d’élévation de privilèges qui servent typiquement de tremplin vers une attaque sur LSASS, en plus de la CVE-2026-62784 mentionnée plus haut. Les équipes de réponse à incident continuent d’observer des outils de type Mimikatz utilisés pour extraire des tickets Kerberos et des hashs NTLM directement depuis la mémoire d’un serveur membre mal protégé. Ce risque autour de NTLM n’est pas propre à Credential Guard : Microsoft avait planifié dès août 2025 un durcissement de NTLMv1, dont le déploiement sur Windows Server 2025 a débuté en novembre 2025 et qui affecte directement certains parcours d’authentification liés à Credential Guard. Durcir Windows Server 2025 avec Credential Guard, VBS et une politique LSA Protection cohérente réduit drastiquement la surface d’attaque exploitable une fois qu’un attaquant a obtenu un accès local.

Le guide ANSSI de sécurisation des serveurs membres Active Directory reste la référence en France pour ce type de durcissement. Il insiste sur des contrôles simples à fort impact : chiffrement BitLocker, isolation des services, synchronisation NTP maîtrisée, et pare-feu restrictif. Ce tutoriel combine ces recommandations avec les nouveaux mécanismes natifs de Windows Server 2025 pour produire une configuration robuste, applicable aussi bien sur un serveur physique que sur une machine virtuelle Hyper-V.

Ce guide s’adresse aux administrateurs systèmes, ingénieurs infrastructure et responsables sécurité qui gèrent un parc Windows Server en environnement d’entreprise, qu’il s’agisse de trois serveurs sur site ou de plusieurs centaines répartis entre datacenter et cloud hybride. Chaque étape est reproductible en environnement de test avant tout déploiement, avec des commandes PowerShell prêtes à l’emploi et les chemins exacts de stratégie de groupe. Aucune connaissance préalable de VBS ou de Credential Guard n’est nécessaire : nous expliquons chaque mécanisme avant de montrer comment l’activer.

## Prérequis : matériel, firmware et versions logicielles

Avant de vous lancer, vérifiez que votre infrastructure remplit les conditions suivantes. Credential Guard et VBS reposent sur des extensions matérielles de virtualisation : sans elles, aucune stratégie de groupe ne les activera.

| Composant | Exigence minimale | Remarque | 
|---|---|---|
| Système d’exploitation | Windows Server 2025 (Standard ou Datacenter, 64 bits) | Édition Core prise en charge | 
| Firmware | UEFI natif | Le mode Legacy/CSM empêche VBS de démarrer | 
| Secure Boot | Activé dans le firmware | Vérifiable via `Confirm-SecureBootUEFI` | 
| Virtualisation matérielle | Intel VT-x + EPT ou AMD-V + NPT/RVI | SLAT obligatoire | 
| TPM | TPM 2.0 recommandé | TPM 1.2 techniquement supporté mais déconseillé pour les baselines ANSSI | 
| Mémoire vive | 8 Go minimum | Marge pour l’overhead VBS/HVCI | 
| Rôle Hyper-V | Activé, même sans VM hébergée | VBS s’appuie sur l’hyperviseur Windows | 
| Statut du serveur | Joint à un domaine, non contrôleur de domaine | Condition de l’activation par défaut de Credential Guard | 

Si votre serveur est virtualisé sous Hyper-V, VMware ou un autre hyperviseur, la virtualisation imbriquée doit être explicitement exposée à la machine virtuelle invitée. Sans cela, VBS échouera silencieusement au démarrage, même si tous les autres prérequis sont réunis. C’est l’un des pièges les plus fréquents que nous détaillons plus loin.

## Étape 1 à 3 : vérifier la compatibilité et activer UEFI, Secure Boot, TPM 2.0

**Étape 1 : auditer le matériel avant toute modification.** Ne touchez à aucun réglage avant d’avoir confirmé que le serveur remplit les prérequis. Exécutez le script suivant en PowerShell administrateur :

```
# Vérifier la virtualisation matérielle activée dans le firmware
systeminfo | Select-String "Virtualization Enabled In Firmware"
# Vérifier l'état de Secure Boot
Confirm-SecureBootUEFI
# Vérifier la présence et la version du TPM
Get-Tpm
# Vérifier l'état d'installation du rôle Hyper-V
Get-WindowsFeature Hyper-V | Format-Table DisplayName, InstallState
```
**Étape 2 : activer Secure Boot et passer en mode UEFI.** Sur un serveur physique, entrez dans le firmware au démarrage (touche variable selon le constructeur : F2, F10, Del) et vérifiez que le mode de démarrage est réglé sur UEFI natif, avec Secure Boot activé. Sur une VM Hyper-V, utilisez plutôt Generation 2 lors de la création de la machine, qui active nativement UEFI et Secure Boot.

