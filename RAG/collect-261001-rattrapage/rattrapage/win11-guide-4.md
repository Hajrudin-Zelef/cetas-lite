---
id: collect-261001-rattrapage/rattrapage/win11-guide-4
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: ["2023-10-10", "2024-10-08", "2025-10-14", "2025-11-11", "2026-09-27", "2026-10-13", "2026-11-10", "2027-10-12"]
keywords: ["amd", "intel", "sandbox"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [209, 387]
sha256: 76bf4666b943ef897e577c8fa8cda864a30744d9dbe1036179ce3440ae9e8a22
---

# Windows 11 en entreprise — Guide technique ultra-complet

L'outil officiel « Contrôle d'intégrité du PC » (`PC Health Check`) donne un verdict binaire. Utile pour un parc de quelques dizaines de postes, insuffisant à l'échelle.

### 3.2 Script d'inventaire de compatibilité (le vrai outil à l'échelle)

À exécuter via GPO (script de démarrage), SCCM/MECM, Intune (script PowerShell) ou en remote :

```powershell
<#
.SYNOPSIS
    Audit de compatibilité Windows 11 — à exécuter en administrateur.
    Exporte un CSV par machine vers un partage réseau.
#>
$ErrorActionPreference = 'SilentlyContinue'
$report = [ordered]@{}

$report.Hostname   = $env:COMPUTERNAME
$report.OS         = (Get-CimInstance Win32_OperatingSystem).Caption
$report.Build      = (Get-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion").CurrentBuild

# CPU
$cpu = Get-CimInstance Win32_Processor
$report.CPU = $cpu.Name.Trim()

# RAM (Go)
$report.RAM_GB = [math]::Round((Get-CimInstance Win32_ComputerSystem).TotalPhysicalMemory / 1GB, 1)

# Disque système (Go libres / taille)
$sys = Get-CimInstance Win32_LogicalDisk -Filter "DeviceID='C:'"
$report.DiskFree_GB  = [math]::Round($sys.FreeSpace / 1GB, 1)
$report.DiskSize_GB  = [math]::Round($sys.Size / 1GB, 1)

# TPM
$tpm = Get-Tpm
$report.TPM_Present = $tpm.TpmPresent
$report.TPM_Ready   = $tpm.TpmReady
$report.TPM_Enabled = $tpm.TpmEnabled
try {
    $spec = (Get-CimInstance -Namespace root/cimv2/security/microsofttpm -ClassName Win32_Tpm).SpecVersion
    $report.TPM_Version = ($spec -split ',')[0].Trim()
} catch { $report.TPM_Version = 'Inconnu' }

# Secure Boot
try { $report.SecureBoot = Confirm-SecureBootUEFI }
catch { $report.SecureBoot = 'Non supporté / BIOS Legacy' }

# Type de firmware
$report.Firmware = if (Test-Path "HKLM:\SYSTEM\CurrentControlSet\Control\SecureBoot\State") { 'UEFI' } else { 'Legacy probable' }

# Verdict
$report.Compatible = ($report.TPM_Version -eq '2.0' -and $report.SecureBoot -eq $true -and $report.RAM_GB -ge 4 -and $report.DiskSize_GB -ge 64)

$out = "\\SRV-FICHIERS\Inventaire$\Win11Compat_$($env:COMPUTERNAME).csv"
[pscustomobject]$report | Export-Csv -Path $out -NoTypeInformation -Encoding UTF8
Write-Output "Rapport écrit : $out — Compatible = $($report.Compatible)"
```

> ⚠️ Adaptez le chemin UNC `\\SRV-FICHIERS\Inventaire$` à votre infrastructure et donnez aux comptes machines le droit d'écriture sur le partage.

### 3.3 Inventaire à distance (WinRM)

```powershell
# Lister les postes d'une OU et auditer la compatibilité à distance
$computers = Get-ADComputer -Filter * -SearchBase "OU=Postes,DC=entreprise,DC=local" |
             Select-Object -ExpandProperty Name

Invoke-Command -ComputerName $computers -ScriptBlock {
    [pscustomobject]@{
        Hostname    = $env:COMPUTERNAME
        TPM         = (Get-Tpm).TpmPresent
        TPMVersion  = ((Get-CimInstance -Namespace root/cimv2/security/microsofttpm -ClassName Win32_Tpm).SpecVersion -split ',')[0]
        SecureBoot  = try { Confirm-SecureBootUEFI } catch { 'N/A' }
        RAM_GB      = [math]::Round((Get-CimInstance Win32_ComputerSystem).TotalPhysicalMemory / 1GB, 1)
    }
} | Export-Csv "C:\Admin\CompatWin11.csv" -NoTypeInformation -Encoding UTF8
```

### 3.4 Cas particulier : activer le TPM / passer en UEFI

| Situation | Action | Risque |
|---|---|---|
| TPM désactivé dans l'UEFI (Intel PTT / AMD fTPM off) | Activer dans le BIOS, souvent via outil constructeur (ex. Dell Command, Lenovo Thin Installer) | Nul |
| Disque en MBR / BIOS Legacy | Convertir avec `MBR2GPT.exe` avant migration | Sauvegarder avant ; échec possible si partitions non standard |
| TPM 1.2 uniquement | Remplacement ou contournement non supporté | À traiter comme poste non migrable |

```powershell
# Conversion MBR -> GPT sans perte de données (depuis Windows 10)
# 1. Valider d'abord :
MBR2GPT.EXE /validate /allowFullOS
# 2. Convertir :
MBR2GPT.EXE /convert /allowFullOS
# Redémarrer, passer le firmware en UEFI, activer Secure Boot.
```

---

## 4. Éditions Windows 11 : Pro vs Entreprise, canaux et cycle de vie

### 4.1 Comparatif Pro / Entreprise (ce qui change vraiment)

| Fonctionnalité | Pro | Entreprise |
|---|---|---|
| Domaine AD / Entra join | ✅ | ✅ |
| BitLocker | ✅ | ✅ |
| GPO (gpedit) | ✅ | ✅ |
| Hyper-V, Sandbox | ✅ | ✅ |
| Windows Update for Business | ✅ | ✅ |
| AppLocker (applocker CSP complet) | Partiel | ✅ |
| Credential Guard (stratégie native) | ❌ (registre possible) | ✅ |
| DirectAccess | ❌ | ✅ (déprécié — préférer Always On VPN) |
| App-V / UE-V | ❌ | ✅ (dépréciés — voir FSLogix) |
| Activation | Clé retail/OEM/MAK/KMS | KMS, Active Directory-Based Activation, abonnement M365 |
| Cycle de vie (feature update) | 24 mois | 36 mois |

**Recommandation** : standardisez sur **Entreprise** dès que le parc dépasse ~50 postes ou que la conformité l'exige (36 mois de support = 1 an de marge de migration en plus). Le surcoût se justifie par AppLocker/Credential Guard natifs et le cycle allongé.

### 4.2 Canaux de mise à jour (canaux de maintenance)

| Canal | Cible | Fréquence feature updates | Support |
|---|---|---|---|
| General Availability Channel | Tous | 1/an (H2, ex. 24H2) | Pro 24 mois / Entreprise 36 mois |
| Windows Insider (Canary/Dev/Beta/Release Preview) | Tests uniquement | Continu | Aucun support prod |

> ⛔ **Ne jamais** inscrire un poste de production au programme Insider. Utilisez un petit lot de pilotes/testeurs sur le canal **Release Preview** pour valider les MàJ en avance.

### 4.3 Cycle de vie — dates clés à connaître

| Version | Build | Fin de support Pro | Fin de support Entreprise |
|---|---|---|---|
| 21H2 | 22000 | 10/10/2023 | 08/10/2024 |
| 22H2 | 22621 | 08/10/2024 | 14/10/2025 |
| 23H2 | 22631 | 11/11/2025 | 10/11/2026 |
| 24H2 | 26100 | 13/10/2026 | 12/10/2027 |
| 25H2 | 26200 | (à confirmer) | (à confirmer) |

> 📌 **Lecture** : au 27/09/2026, la 23H2 Pro est déjà hors support et la 23H2 Entreprise se termine le 10/11/2026. **La cible de déploiement en 2026 est 24H2** (ou 25H2 si validée). Vérifiez toujours les dates sur le site officiel du cycle de vie Microsoft, elles peuvent évoluer.

```powershell
# Version installée sur un poste
Get-ComputerInfo | Select-Object WindowsProductName, WindowsVersion, OsBuildNumber
# ou via le registre :
Get-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion" |
    Select-Object ProductName, DisplayVersion, CurrentBuild, UBR
```

---

## 5. Obtenir les médias : ISO, MCT, VLSC, canaux de mise à jour

| Source | Usage | Remarque |
|---|---|---|
| Media Creation Tool (MCT) | Clé USB / ISO grand public | Éditions Famille/Pro ; pas de VL |
| Page de téléchargement Microsoft (ISO direct) | ISO Pro | Simple, sans VL |
| VLSC (Volume Licensing Service Center) | ISO Entreprise VL | **La source officielle en entreprise** |
| Microsoft 365 admin center | ISO Entreprise (abonnement) | Si licences via M365 |
| Visual Studio Subscriptions | ISO toutes éditions | Pour lab/test |
| Windows Update / WUfB | Mise à niveau | Pas d'ISO, déploiement réseau |

```powershell
# Vérifier l'intégrité d'une ISO téléchargée (SHA-256 publié par Microsoft)
Get-FileHash C:\ISO\Win11_24H2_French_x64.iso -Algorithm SHA256
```

> 💡 **Montez toujours vos images de référence depuis l'ISO VLSC du mois en cours** (ISO « mise à jour » mensuelle) plutôt que de patcher une vieille ISO : vous gagnez le temps d'installation des MàJ cumulatives.

### Éditions contenues dans l'ISO : vérifier l'index

```powershell
# Lister les éditions dans install.wim / install.esd
dism /Get-WimInfo /WimFile:E:\sources\install.wim
# ou :
Get-WindowsImage -ImagePath E:\sources\install.wim
```

---

## 6. Installation propre : clé USB bootable, étapes, pilotes

### 6.1 Créer la clé USB (méthode supportée)

