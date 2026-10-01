---
id: collect-261001-rattrapage/rattrapage/win11-guide-12
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["packaging", "sandbox"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [1731, 1927]
sha256: 1e28fb248599fea95b32abf10d91c3ad10c9e39b90116ab206ca8347c8b45828
---

# Windows 11 en entreprise — Guide technique ultra-complet

**Verdict 2026** : pour un parc Windows 11 moderne, **WUfB + Delivery Optimization** est la cible. WSUS ne se justifie plus que pour des environnements isolés (sans Internet) ou des contraintes réglementaires strictes.

### 27.2 Coexistence / migration WSUS → WUfB

```
1. Inventorier les GPO WSUS existantes (WUServer, TargetGroup...)
2. Créer les anneaux WUfB en parallèle sur une OU pilote
3. Basculer l'OU pilote : supprimer le pointage WSUS (GPO "Spécifier l'emplacement...")
   → les postes repassent sur Windows Update avec les délais WUfB
4. Valider 1 cycle mensuel complet
5. Généraliser par vagues, puis décommissionner WSUS
```

```powershell
# Vérifier vers quoi pointe un poste (WSUS ou WU public)
Get-ItemProperty "HKLM:\SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate" -Name WUServer -ErrorAction SilentlyContinue
# Si vide → Windows Update public / WUfB (selon les autres stratégies)
```

### 27.3 Delivery Optimization (le remplaçant du stockage WSUS)

```powershell
# Limiter le P2P au LAN (pas de pairs Internet)
$do = "HKLM:\SOFTWARE\Policies\Microsoft\Windows\DeliveryOptimization"
New-Item -Path $do -Force | Out-Null
Set-ItemProperty -Path $do -Name "DODownloadMode" -Value 1 -Type DWord
# 0=Bypass, 1=LAN (groupe), 2=Groupe, 3=Internet, 99=Simple, 100=Bypass

# Cache local : taille max (Go)
Set-ItemProperty -Path $do -Name "DOMaxCacheSize" -Value 20 -Type DWord
```

---

## 28. Déploiement d'applications : winget en entreprise

### 28.1 Winget — l'essentiel

`winget` (Gestionnaire de package Windows) est intégré à Windows 11. Il installe depuis des manifests (dépôt community + Microsoft Store).

```powershell
# Rechercher / installer
winget search "7-Zip"
winget install --id 7zip.7zip --silent --accept-package-agreements --accept-source-agreements

# Lister et mettre à jour
winget list
winget upgrade --all --silent --accept-package-agreements
```

### 28.2 Script de post-installation type (MDT / Intune / GPO)

```powershell
<#
.SYNOPSIS
    Installe le socle applicatif via winget. À exécuter en SYSTEM ou admin.
    Idempotent : ne réinstalle pas ce qui est déjà présent.
#>
$ErrorActionPreference = 'Continue'
$apps = @(
    "7zip.7zip",
    "VideoLAN.VLC",
    "Adobe.Acrobat.Reader.64-bit",
    "Mozilla.Firefox",
    "Notepad++.Notepad++"
)

# winget en contexte SYSTEM : forcer le chemin et les accords
$winget = (Get-ChildItem "$env:ProgramFiles\WindowsApps\Microsoft.DesktopAppInstaller_*" |
           Sort-Object LastWriteTime -Descending | Select-Object -First 1).FullName + "\winget.exe"

foreach ($id in $apps) {
    if (-not (& $winget list --id $id --exact 2>$null | Select-String $id)) {
        Write-Output "Installation de $id ..."
        & $winget install --id $id --exact --silent `
            --accept-package-agreements --accept-source-agreements `
            --disable-interactivity
    } else {
        Write-Output "$id déjà installé."
    }
}
```

> ⚠️ **Limites de winget en entreprise** : le dépôt community n'est pas « validé entreprise » (un manifest peut changer). Pour un socle critique : **miroir interne** ou packages validés (Intune Win32, MDT). Winget = excellent pour le socle bureautique standard, pas pour les applications métier critiques sans validation.

### 28.3 Stratégies winget (désactiver / restreindre)

```powershell
# Désactiver winget (postes ultra-sensibles) via stratégie :
$pol = "HKLM:\SOFTWARE\Policies\Microsoft\Windows\AppInstaller"
New-Item -Path $pol -Force | Out-Null
Set-ItemProperty -Path $pol -Name "EnableAppInstaller" -Value 0 -Type DWord

# Restreindre aux sources approuvées : n'autoriser que le dépôt interne
# (via "AllowedSources" / paramètres d'entreprise — voir doc winget Enterprise)
```

### 28.4 Mettre à jour en masse (maintenance mensuelle)

```powershell
# Via Intune : script PowerShell planifié, ou via GPO (tâche planifiée) :
$action = New-ScheduledTaskAction -Execute "powershell.exe" `
    -Argument "-NoProfile -File C:\Admin\winget-upgrade.ps1"
$trigger = New-ScheduledTaskTrigger -Weekly -DaysOfWeek Sunday -At "03:00"
Register-ScheduledTask -TaskName "Winget-Upgrade-Hebdo" -Action $action `
    -Trigger $trigger -User "SYSTEM" -RunLevel Highest
```

---

## 29. Microsoft Store en entreprise, applications privées, packaging

### 29.1 Le Store en entreprise : état des lieux

- Le **Microsoft Store for Business** historique est **fermé**. Remplacé par l'intégration Intune du nouveau Microsoft Store + **dépôt privé**.
- Via Intune : Applications → ajouter une « application du Microsoft Store (nouveau) » → recherche dans le catalogue → affectation aux groupes.

### 29.2 Bloquer ou autoriser le Store (GPO)

| Objectif | Réglage |
|---|---|
| Bloquer le Store grand public | `Désactiver l'application Store` (Composants Windows\Store) |
| Autoriser mais sans applications personnelles | Intune : n'affecter que les apps approuvées + AppLocker |
| Dépôt privé | Intune → Store → applications privées de l'entreprise |

```powershell
# Désactiver le Store via registre (kiosques, postes sensibles)
New-Item -Path "HKLM:\SOFTWARE\Policies\Microsoft\WindowsStore" -Force | Out-Null
Set-ItemProperty -Path "HKLM:\SOFTWARE\Policies\Microsoft\WindowsStore" `
    -Name "RemoveWindowsStore" -Value 1 -Type DWord
```

### 29.3 Packaging : les formats

| Format | Usage | Outil |
|---|---|---|
| **Win32** (.exe/.msi + script) | Standard entreprise | IntuneWinAppUtil → `.intunewin` |
| MSIX | Moderne, sandboxé | MSIX Packaging Tool |
| APPX | Legacy Store | — |
| MSI | Toujours valable | Orca / éditeurs MSI |

```powershell
# Créer un package Intune (.intunewin) depuis un dossier source
.\IntuneWinAppUtil.exe -c C:\Packages\MonApp\Source -s setup.exe -o C:\Packages\MonApp\Output
# Commandes d'install/désinstall à déclarer dans Intune :
#   Install : setup.exe /S
#   Uninstall : "%ProgramFiles%\MonApp\uninstall.exe" /S
# Règle de détection : fichier / registre / version
```

### 29.4 Règles de packaging (checklist)

- [ ] Installation **silencieuse** (aucune interaction)
- [ ] Fonctionne en contexte **SYSTEM** (Intune)
- [ ] Code de retour géré (0 = succès ; 3010 = reboot requis)
- [ ] Règle de **détection** fiable (pas seulement « le dossier existe »)
- [ ] Commande de **désinstallation** testée
- [ ] Testé sur 23H2 **et** 24H2

---

## 30. Sécurité : Microsoft Defender Antivirus — configuration entreprise

### 30.1 Vérifier l'état

```powershell
# État global
Get-MpComputerStatus | Select-Object AntivirusEnabled, RealTimeProtectionEnabled,
    AntispywareEnabled, AMServiceEnabled, AntivirusSignatureLastUpdated,
    QuickScanAge, FullScanAge

# Préférences configurées
Get-MpPreference | Select-Object -ExpandProperty ExclusionPath
```

### 30.2 Durcissement via PowerShell (ou GPO / Intune)

```powershell
# Activer la protection cloud + niveau élevé (recommandé)
Set-MpPreference -MAPSReporting Advanced
Set-MpPreference -SubmitSamplesConsent SendAllSamples

# Protection contre les altérations (Tamper Protection) — via Intune de préférence
Set-MpPreference -DisableTamperProtection $false

# Analyses planifiées : rapide quotidienne, complète hebdo
Set-MpPreference -ScanScheduleDay 0 -ScanScheduleTime "02:00"   # 0 = tous les jours
# Analyse complète le dimanche :
# (via GPO : "Spécifier le jour de la semaine pour l'analyse complète")

# Exclusions : avec parcimonie, documentées, revues trimestriellement
Add-MpPreference -ExclusionPath "D:\BasesDeDonnees"
Add-MpPreference -ExclusionProcess "C:\Apps\Metier\app.exe"
```

### 30.3 Règles de réduction de la surface d'attaque (ASR) — activation

