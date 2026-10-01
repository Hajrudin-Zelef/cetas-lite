---
id: collect-261001-general-networking/general-networking/durcir-windows-server-2025-credential-guard-2026-3
title: "Vérifier la virtualisation matérielle activée dans le firmware"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr", "exploit"]
source: docs/RAG/collect-261001-general-networking/durcir-windows-server-2025-credential-guard-2026.md
source_anchor: ""
source_lines: [136, 206]
sha256: 5f3881c15dfb9eb2fcae0b3e9397b5017b61aca9a66020f7bf05790eaabdd283
---

# Vérifier la virtualisation matérielle activée dans le firmware

Le guide ANSSI « Back to basics » pour les serveurs membres insiste sur des contrôles simples mais à fort impact, qui complètent Credential Guard sans le remplacer. Quatre priorités ressortent : le chiffrement des disques, la réduction des services actifs, la maîtrise du temps réseau et le filtrage réseau strict.

**Étape 10 : chiffrer les volumes avec BitLocker.** ANSSI recommande de chiffrer les disques système et données pour prévenir le vol physique, en particulier sur les serveurs situés dans des salles techniques peu surveillées ou chez un hébergeur tiers.

```
# Chiffrer le volume système avec protecteur TPM
Enable-BitLocker -MountPoint "C:" -TpmProtector
# Chiffrer un volume de données
Enable-BitLocker -MountPoint "D:" -TpmProtector
# Configurer la synchronisation NTP sur la hiérarchie du domaine
w32tm /config /syncfromflags:domhier /update
w32tm /resync
```
**Étape 11 : isoler les services et restreindre le pare-feu.** N’activez que les rôles et fonctionnalités strictement nécessaires au service rendu par le serveur. Un serveur de fichiers n’a aucune raison d’exposer IIS ou un rôle SQL Server actif. Restreignez l’accès en gestion à distance (RDP, WinRM) aux seuls sous-réseaux d’administration et bloquez SMB et LDAP sortants vers l’extérieur du périmètre autorisé. Cette isolation, associée à la logique de tiering (comptes Tier 0 réservés aux contrôleurs de domaine, jamais utilisés pour se connecter à un serveur membre Tier 2), limite le rayon d’action d’un attaquant même s’il parvient à compromettre un premier serveur.

## Étape 12 et 13 : Attack Surface Reduction, WDAC et audit de conformité final

**Étape 12 : activer les règles de réduction de la surface d’attaque (ASR) et déployer WDAC.** Credential Guard et RunAsPPL protègent les secrets une fois qu’un attaquant est déjà sur la machine. ASR et Windows Defender Application Control (WDAC) réduisent la probabilité qu’un tel accès soit obtenu, en bloquant l’exécution de code non fiable.

```
# Activer une règle ASR de blocage du vol d'identifiants depuis les composants système
Add-MpPreference -AttackSurfaceReductionRules_Ids `
    '9e6c4e1f-7d60-472f-ba1a-a39ef669e4b2' `
    -AttackSurfaceReductionRules_Actions Enabled
# Générer une politique WDAC en mode audit sur un serveur de référence
New-CIPolicy -Level SignedVersion -FilePath C:\WDAC\BasePolicy.xml -UserPEs 3
# Convertir la politique en binaire déployable
ConvertFrom-CIPolicy C:\WDAC\BasePolicy.xml C:\WDAC\BasePolicy.bin
# Déployer la politique
Copy-Item C:\WDAC\BasePolicy.bin `
    C:\Windows\System32\CodeIntegrity\CiPolicies\Active\BasePolicy.bin
```
**Étape 13 : vérifier l’application et patcher.** Terminez toujours par une vérification que les correctifs de sécurité du mois sont installés, en particulier ceux liés à LSASS. Credential Guard n’est pas un substitut au patch management, et son historique de correctifs le rappelle : la mise à jour KB5055523 du 8 avril 2025 avait désactivé par erreur les comptes de machine protégés par Credential Guard à cause d’un problème de rotation des mots de passe Kerberos, un défaut que Microsoft n’a corrigé côté catalogue x64 (chaîne .msu erronée) que le 1er juin 2026. La CVE-2026-62784 illustre par ailleurs qu’une faille non corrigée dans lsasrv.dll peut compromettre la logique même que Credential Guard protège.

```
# Vérifier la présence du correctif d'août 2026
Get-HotFix | Where-Object {$_.HotFixID -eq 'KB5120233'}
# Contrôle final de conformité de la baseline
Get-OSConfigCompliance -Name 'WindowsServer2025-MemberServer' | Format-List
```
## Comment Credential Guard bloque une attaque de vol d’identifiants : scénario concret

Pour comprendre l’intérêt réel de ce durcissement, il est utile de comparer le déroulement d’une attaque type sur un serveur non protégé et sur un serveur qui applique les 13 étapes de ce tutoriel. Le scénario suivant reprend une chaîne d’attaque observée régulièrement lors d’incidents de type ransomware : un attaquant obtient un accès initial via un poste utilisateur compromis, puis se déplace latéralement jusqu’à un serveur membre du domaine.

### Avant durcissement : la mémoire LSASS en libre accès

Sur un serveur Windows Server 2025 encore en configuration par défaut désactivée ou sur un serveur legacy migré sans durcissement, un attaquant disposant de privilèges administrateur locaux peut dumper la mémoire du processus LSASS avec un outil comme Mimikatz ou un simple export via l’outil de diagnostic Task Manager (« Créer un fichier de vidage »). Ce fichier contient les hashs NTLM et parfois les tickets Kerberos en cache des comptes récemment authentifiés sur la machine, y compris des comptes à privilèges élevés si un administrateur du domaine s’est connecté récemment pour une opération de maintenance. L’attaquant rejoue ensuite ces identifiants via pass-the-hash ou pass-the-ticket pour progresser vers le contrôleur de domaine, souvent en quelques minutes seulement après l’accès initial.

### Après durcissement : un dump inexploitable

Sur un serveur où Credential Guard, VBS et LSA Protection (RunAsPPL) sont actifs, ce même dump mémoire ne contient plus les secrets exploitables. Credential Guard a déplacé les identifiants dérivés dans le conteneur isolé par l’hyperviseur, hors de portée du processus LSASS classique. RunAsPPL empêche par ailleurs l’injection du code de dump lui-même dans le processus, puisque seul du code signé et approuvé peut s’y attacher. L’attaquant obtient un fichier de vidage vide de tout secret utilisable, ce qui casse la chaîne d’attaque à cette étape précise et l’oblige à chercher un autre vecteur, généralement plus lent et plus détectable par les outils de surveillance (EDR, SIEM). C’est cette rupture de chaîne, bien plus qu’une case cochée dans un audit, qui justifie le temps investi dans ce tutoriel.

## Superviser le durcissement : journaux d’événements et intégration SIEM

Activer Credential Guard et VBS ne suffit pas si personne ne surveille leur état dans la durée. Un pilote mis à jour, une réinstallation de firmware ou une modification manuelle malencontreuse peuvent désactiver silencieusement ces protections. Windows Server 2025 journalise l’état de VBS et de Credential Guard dans le journal d’événements système, sous le fournisseur `Microsoft-Windows-DeviceGuard`. Les identifiants d’événements 7000 à 7003 signalent respectivement le démarrage, l’échec de démarrage, l’arrêt et les changements de configuration de VBS.

```
# Extraire les événements liés à VBS et Credential Guard des dernières 24 heures
Get-WinEvent -FilterHashtable @{
    LogName = 'System'
    ProviderName = 'Microsoft-Windows-DeviceGuard'
    StartTime = (Get-Date).AddHours(-24)
} | Select-Object TimeCreated, Id, Message | Format-Table -AutoSize
# Exporter le résultat pour ingestion par un SIEM (Sentinel, Splunk, Elastic)
Get-WinEvent -FilterHashtable @{LogName='System'; ProviderName='Microsoft-Windows-DeviceGuard'} |
    Export-Csv -Path 'C:\OSConfig\Logs\deviceguard-events.csv' -NoTypeInformation
```
Pour les environnements qui utilisent Microsoft Sentinel, connectez la source Windows Security Events et créez une règle d’alerte sur l’identifiant d’événement 7001 (échec de démarrage de VBS), qui signale généralement un problème matériel, un pilote incompatible ou une modification de firmware non planifiée. Combinez cette surveillance avec le rapport de conformité OSConfig planifié en tâche récurrente, par exemple toutes les six heures via le Planificateur de tâches Windows, pour détecter toute dérive de baseline avant qu’elle ne soit exploitée.

## Pièges courants à éviter lors du durcissement

