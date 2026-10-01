---
id: collect-261001-rattrapage/rattrapage/windows-server-guide-10
title: "Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["license", "open source", "parameters"]
source: docs/RAG/collect-261001-rattrapage/windows_server_guide.md
source_anchor: ""
source_lines: [1606, 1783]
sha256: bb842e4399cc375cbd5cad000efcb129c4105b3bc3d2613d932d5361c9a87de4
---

# Windows Server en entreprise — Guide technique ultra-complet

**Passer à Veeam Backup & Replication (ou équivalent) quand :**
- Parc > 5 serveurs / dizaines de VM
- Besoin de réplication, de SureBackup (test auto des restaurations), de console centrale
- RPO/RTO contractuels (PRA/PCA)
- Sauvegarde applicative fine (SQL, AD, Exchange)

**Alternatives :** Nakivo, Altaro (PME), Bacula/Bareos (open source, voir guide Debian/Ubuntu), Azure Backup (cloud).

> Même avec Veeam : garder la règle 3-2-1, chiffrer, tester les restaurations, documenter le PRA.

---

## 57. Licences : KMS vs ADBA

Deux modes d'activation en volume pour Windows Server (et Windows 10/11) :

| | **KMS** (Key Management Service) | **ADBA** (Active Directory-Based Activation) |
|---|---|---|
| Principe | Un hôte KMS active les clients (port 1688) | L'activation est stockée dans l'AD |
| Seuil d'activation | 25 postes / 5 serveurs | Aucun seuil |
| Renouvellement | Tous les 180 jours (contact KMS) | Persistant (tant que joint au domaine) |
| Hors domaine | Fonctionne | Ne fonctionne pas |
| Idéal | Parc mixte, workgroup | Parc 100 % joint au domaine (le cas courant) |

```powershell
# --- KMS : installer l'hôte ---
slmgr.vbs /ipk XXXXX-XXXXX-XXXXX-XXXXX-XXXXX   # clé hôte KMS (fournie par Microsoft/contrat)
slmgr.vbs /ato                                 # activer l'hôte en ligne
# Publier dans le DNS : automatique (enregistrement _vlmcs._tcp)
# Forcer la publication :
slmgr.vbs /cdns

# --- Côté client : pointer vers le KMS ---
slmgr.vbs /skms srv-kms-01.entreprise.lan:1688
slmgr.vbs /ato
slmgr.vbs /dlv    # détail de licence

# --- ADBA : une seule fois dans la forêt ---
# Sur un DC : installer la clé via l'outil d'activation en volume (VAMT) ou :
slmgr.vbs /ipk <clé ADBA>   # sur le serveur qui héberge le service (souvent le DC)
# Les clients joints au domaine s'activent automatiquement, sans seuil.
```

> Depuis Windows Server 2012 R2, **ADBA est le choix par défaut** pour un parc joint au domaine. KMS reste utile pour les DMZ/workgroup.

---

## 58. CAL utilisateur vs appareil, inventaire

**CAL Windows Server** (Client Access License) : chaque utilisateur **ou** appareil accédant au serveur doit être couvert.

| | Par utilisateur | Par appareil |
|---|---|---|
| Compte | Chaque **personne** | Chaque **terminal** |
| Idéal | 1 personne, N appareils (nomades) | N personnes, 1 poste (3×8, ateliers) |
| Mixable | **Non** : un seul mode par... en pratique on choisit le mode majoritaire et on documente |

**Rôles exigeant des CAL supplémentaires :** RDS (§45), parfois SQL Server.

```powershell
# Inventaire : lister les serveurs et leurs licences (base d'un audit)
Get-ADComputer -Filter { OperatingSystem -like "Windows Server*" } -Property OperatingSystem |
  Group-Object OperatingSystem | Format-Table Name, Count -AutoSize

# Vérifier l'état d'activation de tout le parc (via CIM)
$serveurs = Get-ADComputer -Filter { OperatingSystem -like "Windows Server*" } | Select-Object -ExpandProperty Name
Invoke-Command -ComputerName $serveurs -ScriptBlock {
  (Get-CimInstance SoftwareLicensingProduct | Where-Object { $_.PartialProductKey }).LicenseStatus
} | Group-Object | Format-Table Name, Count
# LicenseStatus : 1 = Licensed (activé)
```

**Documenter :** un tableau `Serveur | Édition | Clé (coffre) | Mode activation | CAL affectées`. Un contrôle de conformité sans inventaire = amende assurée.

---

## 59. Activation : slmgr, DISM, dépannage

```powershell
# Les commandes slmgr.vbs à connaître
slmgr.vbs /dli          # résumé licence
slmgr.vbs /dlv          # détail (canal : Retail/OEM/Volume:GVLK)
slmgr.vbs /xpr          # date d'expiration (KMS : 180 jours renouvelables)
slmgr.vbs /ipk <clé>    # installer une clé
slmgr.vbs /ato          # activer en ligne
slmgr.vbs /upk          # désinstaller la clé
slmgr.vbs /skms <srv>   # définir le KMS

# Activation par téléphone (serveur isolé) :
slmgr.vbs /dti          # affiche l'ID d'installation à dicter

# Changer d'édition avec DISM (voir §2)
DISM /Online /Get-TargetEditions
```

| Erreur | Sens | Action |
|---|---|---|
| `0xC004F074` | KMS introuvable | DNS `_vlmcs._tcp` ? Port 1688 ? `slmgr /skms` correct ? |
| `0xC004C008` | Clé déjà utilisée (quota MAK) | Demander une extension de quota / passer en KMS |
| `0x80072EE7` | Pas de réseau / DNS | Vérifier connectivité vers Microsoft ou KMS |
| Délai de grâce expiré | Non activé | Activer sous 30 jours (notifications puis restrictions) |

---

## 60. Durcissement de base (hardening)

**Les 15 premiers gestes sur tout serveur :**

```powershell
# 1. Désactiver SMBv1 (rançongiciels type WannaCry)
Disable-WindowsOptionalFeature -Online -FeatureName SMB1Protocol -NoRestart

# 2. Désactiver les protocoles obsolètes (LLMNR, NetBIOS si possible)
# GPO : Configuration ordinateur → Modèles d'administration → Réseau → Client DNS → Désactiver la résolution multidiffusion (LLMNR)
Set-ItemProperty "HKLM:\SOFTWARE\Policies\Microsoft\Windows NT\DNSClient" -Name EnableMulticast -Value 0 -Type DWord -Force

# 3. Renommer le compte Administrateur intégré (GPO)
Rename-LocalUser -Name "Administrator" -NewName "AdminLocal-XYZ"

# 4. Désactiver le compte Invité
Disable-LocalUser -Name "Guest"

# 5. Verrouillage de compte (GPO : Stratégie de sécurité locale)
net accounts /lockoutthreshold:5 /lockoutduration:30 /lockoutwindow:30

# 6. Mot de passe : longueur mini 14 (GPO)
net accounts /minpwlen:14

# 7. Désactiver les services inutiles
$servicesInutiles = @("XblGameSave","XboxGipSvc","Print Spooler")  # adapter ! spooler à garder sur SRV-PRINT
foreach ($s in $servicesInutiles) { Set-Service -Name $s -StartupType Disabled -ErrorAction SilentlyContinue }

# 8. RDP : NLA obligatoire + groupe restreint
Set-ItemProperty "HKLM:\SYSTEM\CurrentControlSet\Control\Terminal Server\WinStations\RDP-Tcp" -Name UserAuthentication -Value 1

# 9. Désactiver le partage administratif C$ si non nécessaire (avec prudence)
# HKLM\SYSTEM\CurrentControlSet\Services\LanmanServer\Parameters → AutoShareWks=0 (postes) / AutoShareServer=0 (serveurs)

# 10. UAC au maximum (serveurs avec Desktop)
Set-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System" -Name ConsentPromptBehaviorAdmin -Value 2
```

**Suites :** appliquer un **baseline** (Microsoft Security Compliance Toolkit : GPO "MS Security Guide"), auditer avec **PingCastle** (AD) et **Purview/Secure Score**.

> Avertissement version : les chemins de registre GPO évoluent peu, mais vérifie les noms de stratégies sur 2025 (quelques renommages, ex. "Windows Defender" → "Microsoft Defender").

---

## 61. Pare-feu Windows Defender : règles et profils

3 profils : **Domaine** (le plus permissif, réseau AD), **Privé**, **Public** (le plus restrictif).

```powershell
# État des profils
Get-NetFirewallProfile | Format-Table Name, Enabled, DefaultInboundAction -AutoSize

# Créer une règle : autoriser SMB entrant depuis le LAN uniquement
New-NetFirewallRule -DisplayName "SMB - LAN uniquement" -Direction Inbound `
  -Protocol TCP -LocalPort 445 -RemoteAddress 192.168.10.0/24 -Action Allow -Profile Domain

# Autoriser WinRM (admin à distance)
Enable-NetFirewallRule -DisplayName "Windows Remote Management (HTTP-In)"

# Bloquer un exécutable sortant (ex. télémétrie d'une appli)
New-NetFirewallRule -DisplayName "Blocage appli X" -Direction Outbound `
  -Program "C:\Program Files\X\x.exe" -Action Block

# Auditer : règles actives, triées
Get-NetFirewallRule -Enabled True -Direction Inbound |
  Get-NetFirewallPortFilter | Format-Table Protocol, LocalPort -AutoSize

# Exporter / importer la configuration (pour déploiement)
netsh advfirewall export "D:\Securite\parefeu.wfw"
netsh advfirewall import "D:\Securite\parefeu.wfw"
```

**Règle d'or :** `DefaultInboundAction = Block` sur les profils Public/Privé, n'ouvrir que le nécessaire, **scoper par IP source** (jamais "tout Internet" sauf besoin prouvé comme la passerelle RDS en 443).

