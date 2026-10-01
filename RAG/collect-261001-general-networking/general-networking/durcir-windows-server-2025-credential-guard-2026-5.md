---
id: collect-261001-general-networking/general-networking/durcir-windows-server-2025-credential-guard-2026-5
title: "Vérifier la virtualisation matérielle activée dans le firmware"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr", "benchmark"]
source: docs/RAG/collect-261001-general-networking/durcir-windows-server-2025-credential-guard-2026.md
source_anchor: ""
source_lines: [247, 334]
sha256: b772ddb56e10903f8cb3d43d744d09b5a754a3120fca4d9f5b3944db5eea06e3
---

# Vérifier la virtualisation matérielle activée dans le firmware

```
# Script de durcissement Windows Server 2025 - à exécuter en administrateur
# Étape 1 : audit matériel
$secureBoot = Confirm-SecureBootUEFI
$tpm = Get-Tpm
if (-not $secureBoot -or -not $tpm.TpmPresent) {
    Write-Warning "Prérequis matériels non satisfaits. Arrêt du script."
    exit 1
}
# Étape 2 : installation Hyper-V (fondation VBS)
Install-WindowsFeature -Name Hyper-V -IncludeManagementTools
# Étape 3 : activation VBS et Credential Guard (verrouillage UEFI)
New-Item -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\DeviceGuard' -Force | Out-Null
Set-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\DeviceGuard' `
    -Name 'EnableVirtualizationBasedSecurity' -Type DWord -Value 1
Set-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\Lsa' `
    -Name 'LsaCfgFlags' -Type DWord -Value 1
# Étape 4 : LSA Protection (RunAsPPL)
Set-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\Lsa' `
    -Name 'RunAsPPL' -Type DWord -Value 1
Set-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\Lsa' `
    -Name 'RunAsPPLBoot' -Type DWord -Value 1
# Étape 5 : chiffrement BitLocker (recommandation ANSSI)
Enable-BitLocker -MountPoint "C:" -TpmProtector
# Étape 6 : NTP hiérarchie de domaine (recommandation ANSSI)
w32tm /config /syncfromflags:domhier /update
w32tm /resync
# Étape 7 : application de la baseline OSConfig
Install-Module -Name OSConfig -Scope AllUsers -Force
Import-Module OSConfig
Invoke-OSConfigBaseline -Name 'WindowsServer2025-MemberServer' `
    -ComplianceLogPath 'C:\OSConfig\Logs\MemberServer.json' -Confirm:$false
Write-Output "Durcissement appliqué. Redémarrage requis pour activer VBS/Credential Guard."
Restart-Computer -Confirm
```
## Windows Server 2022 vs 2025 : ce qui change côté sécurité

La différence la plus visible entre les deux versions tient à la posture par défaut. Windows Server 2022 exigeait une activation manuelle complète de Credential Guard et de VBS ; Windows Server 2025 les active automatiquement dès lors que le matériel le permet, ce qui réduit le risque d’oubli lors d’un déploiement rapide. Pour les organisations qui migrent progressivement leur parc de Windows Server 2022 vers 2025, ce changement de posture par défaut signifie qu’un serveur fraîchement migré peut se retrouver avec des protections plus strictes que prévu, ce qui casse parfois des scripts d’administration hérités qui n’avaient jamais été testés avec Credential Guard actif. Planifiez systématiquement une phase de test post-migration dédiée à ce point avant de considérer la bascule comme terminée.

| Aspect | Windows Server 2022 | Windows Server 2025 | 
|---|---|---|
| Credential Guard | Désactivé par défaut, activation manuelle requise | Activé par défaut sur serveurs membres éligibles | 
| Gestion des baselines | Security Compliance Toolkit (GPO uniquement) | Security Compliance Toolkit + module OSConfig (local et centralisé) | 
| Détection de dérive de configuration | Limitée, dépendante d’outils tiers | Intégrée nativement via Get-OSConfigCompliance | 
| LSA Protection (RunAsPPL) | Configuration manuelle recommandée | Intégrée aux baselines de sécurité par défaut | 
| Authentification legacy (WDigest, NTLMv1, RC4) | Encore présente sauf durcissement explicite | Désactivée par défaut dans les baselines mises à jour | 

## Foire aux questions

**Qu’est-ce que Credential Guard et pourquoi est-il activé par défaut sur Windows Server 2025 ?**

Credential Guard isole les identifiants dérivés (hashs NTLM, tickets Kerberos) dans un conteneur mémoire protégé par l’hyperviseur, hors de portée d’un attaquant disposant de privilèges administrateur locaux. Microsoft l’active par défaut sur les serveurs membres éligibles depuis Windows Server 2025 pour réduire le nombre de déploiements qui restent vulnérables faute d’activation manuelle.

**VBS ralentit-il les performances du serveur ?**

Oui, dans une mesure limitée. VBS et HVCI ajoutent une consommation supplémentaire de mémoire et de CPU liée à l’isolation par l’hyperviseur. Sur un serveur correctement dimensionné, l’impact reste marginal ; sur un serveur déjà proche de la saturation (SQL Server, RDS sous forte charge), prévoyez une marge matérielle avant d’activer ces protections.

**Credential Guard doit-il être activé sur les contrôleurs de domaine ?**

Non, pas par la méthode par défaut. Microsoft exclut explicitement les contrôleurs de domaine de l’activation automatique de Credential Guard, en raison de scénarios Kerberos spécifiques qui peuvent être affectés. Suivez les recommandations dédiées aux DC plutôt que d’étendre la configuration serveur membre sans validation préalable.

**Quelle est la différence entre Credential Guard et LSA Protection (RunAsPPL) ?**

Credential Guard protège les secrets eux-mêmes en les déplaçant hors de la mémoire accessible par LSASS. LSA Protection (RunAsPPL) protège le processus LSASS contre l’injection de code malveillant. Les deux mécanismes sont complémentaires et se déploient ensemble dans une configuration durcie.

**Le module OSConfig remplace-t-il le Security Compliance Toolkit ?**

Non, les deux coexistent. Le Security Compliance Toolkit reste pertinent pour un pilotage centralisé via GPO, tandis qu’OSConfig ajoute une couche d’application locale et de détection de dérive, particulièrement utile pour les serveurs Core ou peu connectés au réseau d’entreprise.

**Comment vérifier que Credential Guard fonctionne réellement après activation ?**

Exécutez `Get-ComputerInfo | Select-Object DeviceGuardSecurityServicesRunning` après redémarrage : la sortie doit lister `CredentialGuard` parmi les services actifs. Un simple réglage de stratégie sans ce contrôle post-redémarrage ne garantit pas que la protection est réellement opérationnelle.

**Que faire si mes anciens pilotes ne sont pas compatibles avec HVCI ?**

Identifiez le pilote fautif via l’observateur d’événements après un démarrage en mode sans échec, puis contactez l’éditeur pour une version signée compatible avec l’intégrité du code appliquée par l’hyperviseur. En dernier recours, documentez une exception temporaire tout en planifiant le remplacement du composant incompatible.

**Les recommandations ANSSI sont-elles obligatoires en France ?**

Elles ne constituent pas une obligation légale directe pour toutes les organisations, mais elles deviennent de facto une référence d’exigence pour les entités soumises à la directive NIS2 ou à des obligations sectorielles renforcées. Suivre ce guide ANSSI facilite la démonstration de conformité lors d’un audit de sécurité.

**Combien de temps prend un déploiement complet sur un parc de production ?**

Comptez environ 120 minutes pour dérouler les 13 étapes sur un serveur pilote, puis deux à quatre semaines pour une bascule complète en production, en incluant l’audit des applications legacy, la phase pilote en mode audit OSConfig et le déploiement progressif par vagues d’unités d’organisation. Ne tentez jamais un déploiement massif en une seule opération sur un parc de production sans phase pilote préalable.

### Related Coverage

Sources externes citées dans ce tutoriel : documentation officielle Microsoft sur la configuration de Credential Guard, vue d’ensemble de Credential Guard, le guide ANSSI de sécurisation d’un serveur Windows autonome, le benchmark CIS pour Microsoft Windows Server, le site officiel de l’ANSSI et le bulletin de sécurité Microsoft d’août 2026.
