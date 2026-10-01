---
id: collect-261001-general-networking/general-networking/durcir-windows-server-2025-credential-guard-2026-2
title: "Vérifier la virtualisation matérielle activée dans le firmware"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-general-networking/durcir-windows-server-2025-credential-guard-2026.md
source_anchor: ""
source_lines: [48, 135]
sha256: 391fc025b733a1a8bae06e002520b5a82074020a06f645a83abec7479f86dfcf
---

# Vérifier la virtualisation matérielle activée dans le firmware

**Étape 3 : confirmer que le TPM 2.0 est actif.** Si `Get-Tpm` renvoie `TpmPresent : False` alors que le composant physique existe, le TPM est probablement désactivé dans le firmware ou en mode compatibilité 1.2. Sur les VM Hyper-V Generation 2, activez le module TPM virtuel via `Set-VMKeyProtector` et `Enable-VMTPM` avant de démarrer la machine.

## Étape 4 : installer Hyper-V, fondation de la sécurité basée sur la virtualisation

Même sur un serveur qui n’hébergera jamais de machine virtuelle, le rôle Hyper-V doit être présent, car VBS s’appuie sur l’hyperviseur Windows pour créer un conteneur mémoire isolé du système d’exploitation principal. Cette isolation est ce qui empêche un attaquant disposant de privilèges administrateur locaux de lire les secrets protégés par Credential Guard.

```
# Installer le rôle Hyper-V (sans l'interface de gestion si le serveur est en Core)
Install-WindowsFeature -Name Hyper-V -IncludeManagementTools -Restart
```
C’est précisément à cette étape que surgissent les conflits avec des hyperviseurs tiers. Si le serveur physique héberge déjà VMware ESXi ou un autre hyperviseur en cohabitation, l’installation de Hyper-V échouera ou provoquera une instabilité au redémarrage. Sur un serveur dédié à la charge applicative (pas d’hyperviseur tiers), cette étape se déroule sans friction.

## Étape 5 et 6 : activer VBS et Credential Guard via stratégie de groupe

VBS isole les identifiants dérivés (hashs NTLM, tickets Kerberos TGT/TGS) et d’autres secrets gérés par LSASS dans un environnement protégé par l’hyperviseur. Credential Guard s’appuie directement sur cette isolation : LSASS ne stocke plus les secrets en clair ou sous une forme réversible, il communique avec un « LSASS isolé » via un canal RPC restreint aux composants système privilégiés.

**Étape 5 : activer VBS.** Dans la console de gestion des stratégies de groupe, le chemin unifié pour Windows Server 2025 est :

`Configuration ordinateur → Modèles d'administration → Système → Device Guard → Activer la sécurité basée sur la virtualisation`

Réglez cette stratégie sur **Activé**, puis, sous « Configuration Credential Guard », sélectionnez **Activé avec verrouillage UEFI** pour les serveurs membres. Ce verrouillage empêche un attaquant, ou même un administrateur mal intentionné, de désactiver Credential Guard à distance : la désactivation exige un accès physique et une confirmation au niveau du firmware.

**Étape 6 : configurer Credential Guard.** Pour un déploiement en laboratoire ou une vérification rapide, la configuration peut aussi passer par le registre :

```
# Activer VBS
New-Item -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\DeviceGuard' -Force | Out-Null
Set-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\DeviceGuard' `
    -Name 'EnableVirtualizationBasedSecurity' -Type DWord -Value 1
# Activer Credential Guard avec verrouillage UEFI (valeur 1)
Set-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\Lsa' `
    -Name 'LsaCfgFlags' -Type DWord -Value 1
# Vérifier l'état après redémarrage
Get-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\Lsa' | Select-Object LsaCfgFlags
```
En environnement de production, préférez toujours la stratégie de groupe ou le module OSConfig décrit à l’étape suivante : la modification directe du registre ne survit pas toujours aux réévaluations de baseline et peut créer des divergences de configuration difficiles à auditer.

## Étape 7 : déployer les baselines avec le module OSConfig

OSConfig est la nouveauté la plus utile de Windows Server 2025 pour les équipes qui gèrent des dizaines de serveurs membres. Ce module PowerShell applique des baselines prédéfinies (contrôleur de domaine, serveur membre, groupe de travail), surveille la conformité en continu et signale toute dérive de configuration, y compris sur des serveurs Core ou déconnectés du réseau d’entreprise.

```
# Installer OSConfig depuis PowerShell Gallery
Install-Module -Name OSConfig -Scope AllUsers
Import-Module OSConfig
# Lister les baselines disponibles
Get-OSConfigBaseline
# Simuler l'application de la baseline serveur membre (sans rien modifier)
Invoke-OSConfigBaseline -Name 'WindowsServer2025-MemberServer' -WhatIf
# Appliquer la baseline après validation
Invoke-OSConfigBaseline -Name 'WindowsServer2025-MemberServer' `
    -ComplianceLogPath 'C:\OSConfig\Logs\MemberServer.json' -Confirm:$false
# Vérifier la conformité après application
Get-OSConfigCompliance -Name 'WindowsServer2025-MemberServer'
```
La baseline serveur membre couvre plus de 300 réglages : Credential Guard, protection LSASS/PPL, imposition de TLS 1.2 minimum, SMB 3.0 obligatoire, Kerberos en AES uniquement, et suppression de WDigest des méthodes d’authentification actives. Lancez toujours l’option `-WhatIf` en premier sur un serveur de test : certaines applications legacy s’appuient encore sur des protocoles que la baseline désactive par défaut.

## Étape 8 : importer les baselines via le Security Compliance Toolkit

Le Security Compliance Toolkit (SCT) reste pertinent pour les organisations qui pilotent leur durcissement à travers des GPO centralisées plutôt que par un agent local. Microsoft publie des paquets de baseline pour Windows Server 2025 sous forme de sauvegardes de GPO, accompagnées d’une documentation Excel détaillant chaque réglage.

Le flux de travail classique :

- Téléchargez et extrayez le paquet SCT correspondant à Windows Server 2025.
- Dans la Console de gestion des stratégies de groupe, faites un clic droit sur « Objets de stratégie de groupe » puis « Importer les paramètres ».
- Sélectionnez la sauvegarde de GPO correspondant à votre rôle (serveur membre, contrôleur de domaine).
- Liez la GPO importée à l’unité d’organisation appropriée, en respectant la logique de tiering ANSSI (Tier 0, 1, 2).

SCT et OSConfig ne s’excluent pas : dans une architecture mature, SCT définit la politique centrale via GPO, tandis qu’OSConfig assure l’application locale et la détection de dérive sur les serveurs qui reçoivent difficilement les mises à jour de GPO (serveurs cloud isolés, environnements déconnectés).

## Étape 9 : activer LSA Protection (RunAsPPL) en complément de Credential Guard

LSA Protection, aussi appelée RunAsPPL, exécute le processus LSASS en tant que Protected Process Light : seul du code signé et fiable peut s’y injecter ou en lire la mémoire. C’est un mécanisme complémentaire, pas redondant, avec Credential Guard.

Credential Guard déplace les secrets eux-mêmes dans un conteneur isolé par l’hyperviseur : même en cas de compromission de LSASS, l’attaquant ne trouve rien à extraire. RunAsPPL protège LSASS en tant que processus : il empêche l’injection de code malveillant qui tenterait de manipuler ou de contourner ce mécanisme. Combinés, les deux rendent une attaque de type Mimikatz pratiquement inopérante sur un serveur correctement configuré.

```
# Activer LSA Protection (RunAsPPL)
Set-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\Lsa' `
    -Name 'RunAsPPL' -Type DWord -Value 1
# Forcer l'application dès le démarrage
Set-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\Lsa' `
    -Name 'RunAsPPLBoot' -Type DWord -Value 1
```
Sur Windows Server 2025, ces deux clés sont généralement déjà positionnées par la baseline de sécurité ou par OSConfig ; la configuration manuelle sert surtout à la vérification ou à des environnements qui n’utilisent pas encore ces outils d’automatisation.

## Étape 10 et 11 : appliquer les recommandations ANSSI pour serveurs membres AD DS

