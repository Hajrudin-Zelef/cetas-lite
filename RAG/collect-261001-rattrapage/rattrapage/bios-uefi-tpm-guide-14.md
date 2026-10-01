---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-14
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: ["4011-69-68"]
keywords: ["datacenter", "distribution", "open source"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [2122, 2335]
sha256: d83dcb742df1f4579bd89be0a87d1b2027aebfd727afa25a72f0864be20d910f
---

# BIOS / UEFI — Secure Boot — TPM 2.0

## 47. WDS : installation du rôle

**WDS** (*Windows Deployment Services*) est le rôle Windows Server qui fournit le boot PXE et la distribution d'images.

### Prérequis serveur

- [ ] Windows Server (2016/2019/2022/2025).
- [ ] **AD DS** (recommandé ; le mode autonome existe mais limité).
- [ ] **DHCP** (sur le serveur ou ailleurs — voir les options §46).
- [ ] **DNS**.
- [ ] Volume **NTFS** dédié pour le dépôt d'images (pas le volume système).

### Installation (PowerShell)

```powershell
# Installer le rôle WDS + outils d'administration
Install-WindowsFeature WDS -IncludeManagementTools

# Vérifier
Get-WindowsFeature WDS | Select-Object Name, Installed
```

### Initialisation du serveur

```powershell
# Initialiser WDS (mode intégré AD, dépôt sur E:\RemoteInstall)
wdsutil /Initialize-Server /Server:WDS-01 /RemInst:"E:\RemoteInstall"

# Mode autonome (sans AD) :
wdsutil /Initialize-Server /Server:WDS-01 /RemInst:"E:\RemoteInstall" /Standalone
```

### Structure du dépôt

```
E:\RemoteInstall
├── Boot
│   └── x64
│       ├── wdsmgfw.efi      ← chargeur PXE UEFI
│       ├── wdsnbp.com       ← chargeur PXE Legacy
│       └── boot.wim         ← image de démarrage (WinPE)
├── Images
│   └── install.wim          ← images d'installation
├── MGMT, Tmp, ...
└── WdsClientUnattend        ← fichiers de réponse (unattend.xml)
```

---

## 48. WDS : configuration initiale

### Assistant / paramètres essentiels (console `wdsmgmt.msc`)

```
1. Clic droit sur le serveur → Configurer le serveur.
2. Emplacement RemoteInstall : E:\RemoteInstall (volume NTFS non système).
3. Réponse PXE :
     ○ Ne répondre à aucun client  (désactivé)
     ○ Répondre uniquement aux clients connus
     ● Répondre à tous les clients  ← labo
   Avec "Exiger l'approbation de l'administrateur" en production
   (les machines inconnues attendent une validation).
4. DHCP :
   - Si DHCP sur le même serveur : ☑ "Ne pas écouter sur le port 67"
     ET ☑ "Configurer l'option DHCP 60 sur PXEClient".
   - Si DHCP ailleurs : ne rien cocher ici, configurer les options
     66/67 sur le serveur DHCP (ou utiliser un redirecteur).
```

### Équivalent en ligne de commande

```powershell
# Ne pas écouter le port 67 (DHCP cohabitant) + option 60
wdsutil /Set-Server /Server:WDS-01 /UseDHCPPorts:No /DHCPOption60:Yes

# Politique de réponse : répondre à tous (labo)
wdsutil /Set-Server /Server:WDS-01 /AnswerClients:All

# En production : uniquement les clients connus + approbation
wdsutil /Set-Server /Server:WDS-01 /AnswerClients:Known
wdsutil /Set-Server /Server:WDS-01 /NewMachinePolicy:AdminApproval

# Démarrer le service
Start-Service WDSServer
Set-Service WDSServer -StartupType Automatic
```

### Fichiers de réponse (automatisation)

```
- WdsClientUnattend.xml : automatise la phase PXE (langue, identifiants,
  choix de l'image) → déposé dans RemoteInstall\WdsClientUnattend.
- Unattend.xml d'image : automatise l'installation Windows elle-même
  (créé avec Windows System Image Manager, SIM).
```

---

## 49. WDS : images de démarrage et d'installation

### Les deux types d'images

| Image | Fichier | Rôle |
|---|---|---|
| **Image de démarrage** (*boot image*) | `boot.wim` | WinPE : l'environnement minimal qui démarre en PXE et lance l'installation |
| **Image d'installation** (*install image*) | `install.wim` | Le Windows à installer (avec vos customisations) |

### Ajouter les images

```powershell
# 1. Image de démarrage : depuis l'ISO Windows (sources\boot.wim)
wdsutil /Add-Image /Server:WDS-01 /ImageFile:"D:\sources\boot.wim" /ImageType:Boot

# 2. Image d'installation : depuis l'ISO (sources\install.wim)
wdsutil /Add-Image /Server:WDS-01 /ImageFile:"D:\sources\install.wim" /ImageType:Install /ImageGroup:"Windows 11"

# Lister les images
wdsutil /Get-Image /Server:WDS-01 /ImageType:Boot
wdsutil /Get-Image /Server:WDS-01 /ImageType:Install /ImageGroup:"Windows 11"
```

### Créer une image de capture (pour cloner un poste master)

```
1. Préparer le poste master (applications, réglages).
2. sysprep : C:\Windows\System32\Sysprep\sysprep.exe /oobe /generalize /shutdown
3. Créer une image de capture dans WDS (à partir de boot.wim) :
   console WDS → Images de démarrage → clic droit sur boot.wim →
   "Créer une image de capture" → la déployer sur le WDS.
4. PXE-booter le poste master → choisir l'image de capture →
   capturer vers install.wim sur le serveur.
```

### Pilotes dans l'image de démarrage (indispensable !)

Sans le pilote réseau/NVMe dans `boot.wim`, le WinPE ne voit ni le réseau ni le disque :

```powershell
# Ajouter un groupe de pilotes au serveur
wdsutil /Add-DriverPackage /Server:WDS-01 /DriverFile:"E:\Drivers\NIC\net.inf"

# Injecter dans l'image de démarrage
wdsutil /Add-ImageDriverPackage /Server:WDS-01 /Image:"Microsoft Windows Setup (x64)" `
    /ImageType:Boot /DriverGroup:"WinPE x64"
```

> 💡 **En 2026 :** pensez aussi à **MDT** (Microsoft Deployment Toolkit, gratuit) au-dessus de WDS pour les séquences de tâches, ou **SCCM/MECM** pour les grands parcs. WDS seul = déploiement basique ; MDT/SCCM = déploiement industrialisé.

---

## 50. Dépannage PXE / WDS

### Tableau des erreurs PXE courantes

| Code / message | Cause probable | Solution |
|---|---|---|
| `PXE-E53: No boot filename received` | Pas de réponse du serveur WDS / options DHCP | Vérifier que WDSServer tourne, options 66/67 ou 60, pare-feu (ports 67/68/69/4011) |
| `PXE-M0F: Exiting PXE ROM` | Aucun serveur PXE trouvé | Câble/VLAN, DHCP fonctionnel ?, WDS démarré ? |
| `PXE-E51: No DHCP or proxyDHCP offers` | Pas de DHCP | Serveur DHCP joignable ? Relais DHCP (ip helper) sur le routeur inter-VLAN ? |
| `PXE-E32: TFTP open timeout` | TFTP bloqué | Pare-feu UDP 69, service WDS, fichier de boot présent |
| `0xc000000f` après chargement | BCD de l'image corrompu | Régénérer l'image de démarrage |
| F12 sans effet / boot direct disque | Ordre de boot, ou touche différente (F12 Dell/HP, F8/F12 Lenovo) | Menu de boot au démarrage |
| WinPE sans réseau | Pilote NIC absent de boot.wim | Injecter le pilote (§49) |
| WinPE ne voit pas le disque | Pilote NVMe/RAID absent | Injecter le pilote ; vérifier mode SATA (AHCI vs RAID) |

### Check-list de diagnostic (dans l'ordre)

```powershell
# Côté serveur WDS
Get-Service WDSServer | Select-Object Status, StartType
# → Running + Automatic ?

# Le port TFTP écoute-t-il ?
Get-NetTCPConnection -LocalPort 69 -ErrorAction SilentlyContinue
# (TFTP = UDP : vérifier avec netstat -ano | findstr ":69")

# Journal WDS
Get-WinEvent -LogName 'Microsoft-Windows-Deployment-Services-Diagnostics/Operational' `
    -MaxEvents 30 | Select-Object TimeCreated, Id, Message | Format-Table -Wrap

# Côté client : tester le DHCP
# (depuis un poste du même VLAN)
ipconfig /all   # passerelle, DHCP visibles ?
```

### Le cas du DHCP sur le même serveur (piège n°1)

```
Symptômes : PXE-E53 ou pas de réponse PXE, alors que le DHCP fonctionne.
Cause : le service DHCP et WDS se disputent le port UDP 67.
Solution : dans les propriétés WDS → onglet DHCP :
  ☑ "Ne pas écouter sur les ports DHCP"
  ☑ "Configurer l'option DHCP 60 sur PXEClient"
En CLI :
  wdsutil /Set-Server /UseDHCPPorts:No /DHCPOption60:Yes
```

### Inter-VLAN (voir aussi cas pratique 13, §76)

```
Le broadcast DHCP ne traverse pas les routeurs.
→ Configurer un relais DHCP (ip helper-address) sur l'interface du VLAN
  client, pointant vers le serveur DHCP ET vers le serveur WDS
  (ou utiliser l'option DHCP 60/66/67 + redirecteur PXE).
```

---

## 51. iPXE et boot réseau avancé

**iPXE** est un firmware PXE open source avancé : il ajoute **HTTP**, les scripts, et le boot depuis des URL — là où le PXE classique est limité à TFTP.

### Pourquoi iPXE en entreprise/datacenter

