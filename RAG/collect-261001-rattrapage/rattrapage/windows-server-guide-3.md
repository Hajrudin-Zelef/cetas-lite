---
id: collect-261001-rattrapage/rattrapage/windows-server-guide-3
title: "Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "distribution", "ethernet", "gpu", "intel"]
source: docs/RAG/collect-261001-rattrapage/windows_server_guide.md
source_anchor: ""
source_lines: [360, 564]
sha256: e335535fa0c591f374c65c22748c62c446adc77cbf366434fee3bc00870fff7b
---

# Lister ce qui est installé
Get-WindowsFeature | Where-Object Installed | Format-Table Name, DisplayName -AutoSize
```

---

## 12. Installation via Server Manager

1. **Gérer → Ajouter des rôles et fonctionnalités**.
2. Installation basée sur un rôle ou une fonctionnalité → serveur local ou distant.
3. Cocher le(s) rôle(s) → accepter les fonctionnalités dépendantes.
4. Options par rôle (ex. Hyper-V : créer le commutateur virtuel).
5. **Redémarrage automatique si nécessaire** : à cocher avec prudence en production.
6. Vérifier dans **Notifications** (drapeau) que l'installation a réussi.

> Server Manager est pratique pour découvrir, mais **tout déploiement reproductible passe par PowerShell** (§13) : script versionné, rejouable, auditable.

---

## 13. Installation via PowerShell (Install-WindowsFeature)

```powershell
# Installer un rôle simple
Install-WindowsFeature -Name Hyper-V -IncludeManagementTools -Restart

# Installer plusieurs rôles/fonctionnalités d'un coup
Install-WindowsFeature -Name AD-Domain-Services, DNS, DHCP -IncludeManagementTools

# Avec les outils d'administration, sans redémarrage forcé
Install-WindowsFeature -Name File-Services, FS-FileServer -IncludeManagementTools

# Depuis un support d'installation (source alternative, utile en Core sans internet)
Install-WindowsFeature -Name Net-Framework-Core -Source D:\sources\sxs

# Désinstaller proprement
Uninstall-WindowsFeature -Name Telnet-Client -Restart:$false

# Vérifier + journal
Get-WindowsFeature -Name Hyper-V | Select-Object Installed, InstallState
```

**Noms de rôles à connaître par cœur :**

| Nom PowerShell | Rôle |
|---|---|
| `AD-Domain-Services` | AD DS |
| `DNS` | Serveur DNS |
| `DHCP` | Serveur DHCP |
| `Hyper-V` | Hyper-V |
| `Failover-Clustering` | Clustering de basculement |
| `FS-FileServer` | Serveur de fichiers |
| `FS-DFS-Namespace` / `FS-DFS-Replication` | DFS-N / DFS-R |
| `FS-Data-Deduplication` | Déduplication |
| `FS-Resource-Manager` | FSRM |
| `Print-Services`, `Print-Server` | Serveur d'impression |
| `Web-Server` | IIS |
| `Windows-Server-Backup` | Sauvegarde |
| `UpdateServices` | WSUS |
| `AD-Certificate` | AD CS (PKI) |
| `RDS-RD-Server` (+ sous-rôles) | RDS |

---

## 14. Les 15 rôles courants et quand les installer

| # | Rôle | Quand l'installer |
|---|---|---|
| 1 | AD DS | Contrôleur de domaine (toujours ≥ 2 par domaine) |
| 2 | DNS | Avec AD DS, ou DNS dédié |
| 3 | DHCP | Distribution d'adresses (redondance : basculement DHCP) |
| 4 | Hyper-V | Virtualisation |
| 5 | Clustering de basculement | Haute dispo Hyper-V / fichiers |
| 6 | Serveur de fichiers | Partages SMB, DFS |
| 7 | Déduplication | Volumes de fichiers avec doublons (jamais sur CSV ni disques système) |
| 8 | FSRM | Quotas, filtrage (bloquer .mp3, .exe...) |
| 9 | Serveur d'impression | Centraliser copieurs/imprimantes (§38) |
| 10 | WSUS | Patch management Windows |
| 11 | AD CS | PKI interne (certificats 802.1X, LDAPS, RDP...) |
| 12 | RDS | Bureaux / applis à distance |
| 13 | IIS | Applis web internes |
| 14 | Windows Server Backup | Sauvegarde locale d'appoint |
| 15 | WDS/MDT (Services de déploiement) | Déploiement d'images Windows (postes) |

---

## 15. Hyper-V : architecture et prérequis

Hyper-V est un hyperviseur **de type 1** (bare metal) : il s'installe sous Windows, qui devient lui-même une VM privilégiée (partition parente).

**Prérequis matériels :**
- CPU 64 bits avec **SLAT** (EPT/NPT) — tout CPU serveur depuis ~2012
- **Virtualisation matérielle activée dans le BIOS/UEFI** (Intel VT-x / AMD-V)
- DEP/NX activé
- RAM : 4 Go minimum pour l'hôte, bien plus en pratique

```powershell
# Vérifier la compatibilité AVANT d'installer
systeminfo | Select-String "Hyper-V"

# Détail via WMI
Get-CimInstance Win32_Processor | Select-Object Name, SecondLevelAddressTranslationExtensions,
  VirtualizationFirmwareEnabled, VMMonitorModeExtensions
# VirtualizationFirmwareEnabled doit être True (= activé dans le BIOS)
```

**Installation :**

```powershell
Install-WindowsFeature -Name Hyper-V -IncludeManagementTools -Restart
# Le redémarrage est OBLIGATOIRE (l'hyperviseur se charge au boot)
```

> Sur Windows Server 2025, Hyper-V prend en charge les GPU partitionnés (GPU-P) améliorés et les VM confidentielles de manière plus large. Les concepts ci-dessous restent identiques.

---

## 16. Commutateurs virtuels : externe, interne, privé

| Type | La VM parle au réseau physique ? | L'hôte parle aux VM ? | Usage typique |
|---|---|---|---|
| **Externe** | Oui (via une carte physique liée) | Oui (si "autoriser le partage") | VM de production |
| **Interne** | Non | Oui | Lab, management isolé |
| **Privé** | Non | Non | Lab totalement isolé, malware |

```powershell
# Lister les commutateurs
Get-VMSwitch | Format-Table Name, SwitchType, NetAdapterInterfaceDescription -AutoSize

# Créer un commutateur externe (lié à la carte physique "Ethernet")
New-VMSwitch -Name "vSwitch-Prod" -NetAdapterName "Ethernet" -AllowManagementOS $true

# Créer un commutateur interne
New-VMSwitch -Name "vSwitch-Lab" -SwitchType Internal

# Créer un commutateur privé
New-VMSwitch -Name "vSwitch-Isole" -SwitchType Private

# Identifier la carte physique à lier (si plusieurs)
Get-NetAdapter | Where-Object Status -eq 'Up' | Format-Table Name, InterfaceDescription, LinkSpeed
```

**Bonnes pratiques réseau Hyper-V :**
- Dédier au moins 1 carte physique (ou 1 team) au trafic VM, 1 au management hôte.
- **Ne jamais** mettre le trafic de cluster/CSV et le trafic VM sur le même vSwitch sans QoS.
- Activer le **trunk VLAN** si les VM sont sur plusieurs VLAN :

```powershell
# Port trunk sur une VM (VLAN 10-20 autorisés)
Set-VMNetworkAdapterVlan -VMName "SRV-WEB-01" -Trunk -AllowedVlanIdList "10-20" -NativeVlanId 10
# Port access simple
Set-VMNetworkAdapterVlan -VMName "SRV-AD-01" -Access -VlanId 10
```

---

## 17. VM génération 1 vs génération 2

| Critère | Génération 1 | Génération 2 |
|---|---|---|
| Firmware | BIOS | **UEFI** (Secure Boot) |
| Démarrage | IDE, PXE legacy | SCSI virtuel, PXE UEFI |
| OS supportés | Anciens (2008 R2, Linux legacy...) | 2012+ / 2016+, Linux récents |
| Secure Boot | Non | Oui |
| vTPM | Non | Oui (VM protégées) |
| Performances disque | Moindres (IDE émulé) | Meilleures (SCSI paravirtualisé) |

**Règle : toujours Génération 2** sauf besoin d'un OS antédiluvien. Le choix est **définitif** à la création.

---

## 18. Création d'une VM pas à pas (PowerShell)

```powershell
# 1. Créer la VM (génération 2, 4 Go RAM, nouveau VHDX de 100 Go)
New-VM -Name "SRV-AD-02" -MemoryStartupBytes 4GB -Generation 2 `
  -NewVHDPath "D:\VMs\SRV-AD-02\disque.vhdx" -NewVHDSizeBytes 100GB `
  -SwitchName "vSwitch-Prod" -Path "D:\VMs\SRV-AD-02"

# 2. CPU virtuels
Set-VMProcessor -VMName "SRV-AD-02" -Count 4

# 3. Désactiver le Secure Boot si l'OS ne le supporte pas (rare)
Set-VMFirmware -VMName "SRV-AD-02" -EnableSecureBoot Off

# 4. Monter l'ISO et démarrer
Add-VMDvdDrive -VMName "SRV-AD-02" -Path "D:\ISO\fr_windows_server_2022.iso"
Set-VMFirmware -VMName "SRV-AD-02" -FirstBootDevice (Get-VMDvdDrive -VMName "SRV-AD-02")
Start-VM -Name "SRV-AD-02"

# 5. Après l'install OS : démonter l'ISO
Set-VMFirmware -VMName "SRV-AD-02" -FirstBootDevice (Get-VMHardDiskDrive -VMName "SRV-AD-02")
Remove-VMDvdDrive -VMName "SRV-AD-02" -ControllerNumber 0 -ControllerLocation 1

# 6. Installer les services d'intégration (inclus depuis 2016, via Windows Update)
Get-VMIntegrationService -VMName "SRV-AD-02"
```

> Stocker les VM sur un volume dédié (jamais sur C:). Convention : `D:\VMs\<NomVM>\`.

---

## 19. Mémoire dynamique, processeurs virtuels, disques

**Mémoire dynamique :**

