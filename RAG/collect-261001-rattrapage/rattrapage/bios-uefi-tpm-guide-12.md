---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-12
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "intel", "valuation"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [1731, 1936]
sha256: a26c87e359dbe4623cdf9d77fc4edd2f1060cd5a831d26717b7807aacadee75b
---

# BIOS / UEFI — Secure Boot — TPM 2.0

```powershell
# Méthode 1 : les 3 piliers en une fois
[PSCustomObject]@{
    SecureBoot = Confirm-SecureBootUEFI
    TPM_Present = (Get-Tpm).TpmPresent
    TPM_Ready   = (Get-Tpm).TpmReady
    GPT_System  = ((Get-Disk | Where-Object IsSystem).PartitionStyle -eq 'GPT')
    UEFI_Boot   = (Test-Path "$env:SystemRoot\Panther\setupact.log") # indice, voir §4
}

# Méthode 2 : PC Health Check (outil graphique Microsoft, à déployer si besoin)

# Méthode 3 : registre — état de l'évaluation Windows 11
Get-ItemProperty 'HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion\AppCompatFlags\TargetVersionUpgradeExperienceIndicators\NI23H2' |
    Select-Object UpgEx, UpgExU
# UpgEx = "Green" → compatible ; "Red" → bloqué (voir les raisons dans les valeurs)
```

### Chemin de migration type (parc Windows 10 → 11)

```
Pour chaque poste :
1. Inventaire (§40) → classer : OK / TPM à activer / Legacy à convertir / incompatible.
2. Activer TPM + Secure Boot dans le firmware (vague pilote d'abord).
3. Si Legacy : mbr2gpt /convert (§8) + bascule UEFI.
4. Vérifier les 4 piliers (script ci-dessus).
5. Déployer Windows 11 (MDT/SCCM/Intune, image avec firmware à jour).
6. Post-installation : re-vérifier BitLocker + séquestre des clés.
```

> ⚠️ **Les contournements** (installation de Windows 11 sans TPM via clés de registre `BypassTPMCheck`) existent mais sont **à proscrire en entreprise** : pas de support Microsoft, pas de mises à jour garanties, faille de sécurité assumée. Un poste non conforme = un poste à remplacer, pas à bricoler.

---

## 40. Inventaire du parc : script PowerShell de conformité

Script complet : interroge les machines via CIM/WinRM et produit un CSV de conformité Windows 11 + sécurité firmware.

```powershell
#requires -Version 5.1
<#
.SYNOPSIS
    Inventaire firmware/sécurité du parc : TPM, Secure Boot, GPT/UEFI, BitLocker, BIOS.
.EXAMPLE
    .\Invoke-FirmwareInventory.ps1 -ComputerNames (Get-Content .\parc.txt) |
        Export-Csv .\inventaire-firmware.csv -NoTypeInformation -Encoding UTF8
#>
[CmdletBinding()]
param(
    [Parameter(Mandatory)][string[]]$ComputerNames,
    [PSCredential]$Credential
)

$scriptBlock = {
    $tpm = $null; $sb = $null
    try { $tpm = Get-Tpm -ErrorAction Stop } catch { }
    try { $sb  = Confirm-SecureBootUEFI -ErrorAction Stop } catch { }

    $sysDisk = Get-Disk | Where-Object IsSystem | Select-Object -First 1
    $bde     = Get-BitLockerVolume -MountPoint ($env:SystemDrive) -ErrorAction SilentlyContinue
    $bios    = Get-CimInstance -ClassName Win32_BIOS
    $cs      = Get-CimInstance -ClassName Win32_ComputerSystem

    [PSCustomObject]@{
        ComputerName   = $env:COMPUTERNAME
        Manufacturer   = $cs.Manufacturer
        Model          = $cs.Model
        BiosVersion    = $bios.SMBIOSBIOSVersion
        BiosDate       = $bios.ReleaseDate
        TpmPresent     = $tpm.TpmPresent
        TpmReady       = $tpm.TpmReady
        TpmSpecVersion = $tpm.TpmPresent ? '2.0 (à confirmer via tpm.msc)' : 'N/A'
        SecureBoot     = $sb
        PartitionStyle = $sysDisk.PartitionStyle
        BitLocker      = $bde.VolumeStatus
        Win11_Ready    = ($sb -eq $true -and $tpm.TpmReady -and
                          $sysDisk.PartitionStyle -eq 'GPT')
    }
}

foreach ($computer in $ComputerNames) {
    try {
        $params = @{ ComputerName = $computer; ScriptBlock = $scriptBlock;
                     ErrorAction = 'Stop' }
        if ($Credential) { $params.Credential = $Credential }
        Invoke-Command @params
    } catch {
        [PSCustomObject]@{
            ComputerName = $computer; Manufacturer = 'INJOIGNABLE'
            Model = $_.Exception.Message; BiosVersion = ''; BiosDate = $null
            TpmPresent = $null; TpmReady = $null; TpmSpecVersion = ''
            SecureBoot = $null; PartitionStyle = ''; BitLocker = ''
            Win11_Ready = $false
        }
    }
}
```

### Exploiter le CSV

```powershell
$inv = Import-Csv .\inventaire-firmware.csv

# Synthèse
$inv | Group-Object { $_.Win11_Ready } | Select-Object Name, Count

# Postes à traiter : TPM absent ou non prêt
$inv | Where-Object { $_.TpmReady -ne 'True' } |
    Select-Object ComputerName, Manufacturer, Model, TpmPresent, TpmReady |
    Export-Csv .\a-traiter-tpm.csv -NoTypeInformation

# Postes Legacy (non GPT)
$inv | Where-Object { $_.PartitionStyle -ne 'GPT' } |
    Select-Object ComputerName, Model

# BIOS vieux de plus de 2 ans (candidats à la MAJ firmware)
$inv | Where-Object { $_.BiosDate -and ([datetime]$_.BiosDate) -lt (Get-Date).AddYears(-2) } |
    Select-Object ComputerName, Model, BiosVersion, BiosDate
```

> 💡 **Prérequis réseau :** WinRM activé sur les postes (`Enable-PSRemoting`), compte admin du domaine, pare-feu ouvert (TCP 5985). Pour un parc Intune-only (sans AD), préférez les rapports de conformité Intune / Endpoint Analytics.

---

## 41. Ordre de boot : configuration et bonnes pratiques

L'**ordre de boot** (*Boot Order / Boot Priority*) définit dans quel ordre le firmware essaie les périphériques.

### Ordre recommandé en entreprise

```
1. Disque dur / SSD interne (Windows Boot Manager)
2. Réseau (IPv4 PXE)          ← si déploiement réseau utilisé
3. USB                         ← ou désactivé (voir §44)
4. CD/DVD (si présent)
```

### Pourquoi mettre le disque en premier

- **Sécurité :** un attaquant avec une clé USB bootable ne peut pas démarrer dessus si l'USB est après le disque **et** protégé par mot de passe firmware.
- **Vitesse :** pas de timeout PXE/USB à chaque démarrage.
- **Fiabilité :** pas de boot accidentel sur une clé oubliée.

### Modifier l'ordre

```
Setup UEFI → Boot → Boot Priority Order → monter "Windows Boot Manager"
en premier (touches + / - ou F5/F6 selon firmware) → F10.
```

```powershell
# Via Windows : forcer Windows Boot Manager en premier dans l'ordre firmware
bcdedit /set '{fwbootmgr}' displayorder '{bootmgr}' /addfirst

# Boot unique sur le réseau au prochain redémarrage (utile pour réimager) :
# → le plus simple reste la touche F12 (menu de boot) au démarrage.
```

### Piège : l'ordre ne suffit pas

L'ordre de boot **n'est pas une sécurité** à lui seul : sans mot de passe firmware, n'importe qui change l'ordre en 30 secondes. **Ordre de boot + mot de passe admin UEFI + Secure Boot = le trio de base** (sections 43, 16).

---

## 42. Virtualisation matérielle : VT-x / AMD-V et VT-d / AMD-Vi

### Les deux couches à activer

| Technologie | Nom Intel | Nom AMD | Rôle |
|---|---|---|---|
| Virtualisation CPU | **VT-x** | **AMD-V** (SVM) | Permet à l'hyperviseur de s'exécuter (Hyper-V, VirtualBox, VMware...) |
| Virtualisation E/S | **VT-d** | **AMD-Vi** (IOMMU) | Isolation DMA : un périphérique ne peut pas lire la mémoire d'une autre VM (requis pour Credential Guard, VBS) |

### Où les activer

```
Setup UEFI → Advanced → CPU Configuration (ou Virtualization) :
  Intel Virtualization Technology → Enabled
  VT-d → Enabled
AMD : SVM Mode → Enabled ; IOMMU → Enabled
```

### Vérifier dans Windows

```powershell
# Méthode 1 : informations système
systeminfo | Select-String 'Hyper-V'

# Méthode 2 : CIM (détail)
Get-CimInstance -ClassName Win32_Processor |
    Select-Object Name, VirtualizationFirmwareEnabled, SecondLevelAddressTranslationExtensions

# Méthode 3 : le gestionnaire des tâches → onglet Performances → CPU
#   "Virtualisation : Activé"

# Méthode 4 : vérifier que l'hyperviseur tourne
Get-CimInstance -ClassName Win32_ComputerSystem |
    Select-Object HypervisorPresent
```

### Cas d'usage en entreprise

- **Hyper-V / WSL2 / Docker Desktop** : exigent VT-x/AMD-V.
- **Credential Guard / HVCI (VBS)** : exigent VT-x + VT-d (et UEFI + Secure Boot).
- **Postes développeurs** : activer systématiquement dans l'image master.

