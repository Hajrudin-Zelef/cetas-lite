---
id: collect-261001-rattrapage/rattrapage/win11-guide-21
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [3495, 3703]
sha256: 9d0cff7d40e690357754d053041caadba5b2c8b3220fd29b35a80134f08c907d
---

# Windows 11 en entreprise — Guide technique ultra-complet

```powershell
# Configurer un kiosque (ex. : Edge sur un portail)
Set-AssignedAccess -AppUserModelId "Microsoft.MicrosoftEdge.Stable_8wekyb3d8bbwe!App" -UserName "kiosque"
# Le compte "kiosque" doit exister en local (standard, pas admin)

# Pour un kiosque multi-applications : utiliser le XML de configuration
# (Set-AssignedAccess avec -ConfigFile, voir doc "CreateAssignedAccessSettings")
```

### 65.2 Postes partagés : bonnes pratiques

| Point | Réglage |
|---|---|
| Suppression auto des profils | GPO « Supprimer les profils utilisateurs datant de plus de X jours » |
| Compte invité | **Désactivé** (compte Invité intégré = à laisser désactivé) ; créer des comptes locaux nominatifs ou Entra |
| Verrouillage de session | Court (5 min) par GPO |
| Nettoyage à la fermeture | Script de déconnexion : vider Téléchargements, historique navigateur |

```powershell
# Supprimer les profils inactifs depuis plus de 30 jours (GPO équivalente)
$del = "HKLM:\SOFTWARE\Policies\Microsoft\Windows\System"
New-Item -Path $del -Force | Out-Null
# (La GPO native "Supprimer les profils utilisateurs..." se configure dans :
#  Configuration ordinateur\Modèles d'administration\Système\Profils utilisateur)
```

---

## 66. Impression en entreprise : Universal Print, Point and Print

### 66.1 Universal Print (cloud)

Service d'impression cloud Microsoft 365 : les imprimantes compatibles (ou via un connecteur) sont publiées dans Entra, déployées par Intune, **sans serveur d'impression**.

```powershell
# Côté client : installer une imprimante Universal Print (Intune la pousse en général)
# Vérifier les imprimantes :
Get-Printer | Select-Object Name, DriverName, PortName, Shared
```

### 66.2 Point and Print — durcissement (depuis PrintNightmare)

```powershell
# Restreindre Point and Print aux serveurs approuvés (GPO recommandée)
$pp = "HKLM:\SOFTWARE\Policies\Microsoft\Windows NT\Printers\PointAndPrint"
New-Item -Path $pp -Force | Out-Null
Set-ItemProperty -Path $pp -Name "RestrictDriverInstallationToAdministrators" -Value 1 -Type DWord
Set-ItemProperty -Path $pp -Name "TrustedServers" -Value 1 -Type DWord
Set-ItemProperty -Path $pp -Name "ServerList" -Value "srv-print1;srv-print2" -Type String
# + "PackagePointAndPrintOnly" = 1 (n'accepter que les pilotes packagés)
```

### 66.3 Dépannage impression (express)

```powershell
# Redémarrer le spouleur (résout 70 % des cas)
Restart-Service Spooler -Force

# Vider la file bloquée
Stop-Service Spooler -Force
Remove-Item "C:\Windows\System32\spool\PRINTERS\*" -Force
Start-Service Spooler

# Voir les files et les erreurs
Get-PrintJob -PrinterName "Imprimante-Compta" | Select-Object Id, JobStatus, UserName
Get-WinEvent -FilterHashtable @{LogName='Microsoft-Windows-PrintService/Admin'; Level=2; StartTime=(Get-Date).AddDays(-1)} |
    Select-Object -First 5 TimeCreated, Id, Message
```

---

## 67. Langues, régions, claviers : déploiement multilingue

### 67.1 Ajouter une langue en ligne de commande

```powershell
# Installer le pack de langue français (si l'image est en anglais)
$lp = "HKLM:\SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate"  # (via WU ou source locale)
# Méthode DISM (depuis le CAB de langue) :
dism /Online /Add-Package /PackagePath:C:\Langues\Microsoft-Windows-Client-Language-Pack_x64_fr-fr.cab

# Définir la langue d'affichage
Set-WinUILanguageOverride -Language fr-FR
Set-WinUserLanguageList fr-FR -Force
Set-WinSystemLocale fr-FR
Set-WinHomeLocation -GeoId 84   # 84 = France
```

### 67.2 Clavier et paramètres régionaux (image de référence / unattend)

```powershell
# Clavier FR par défaut pour tous les nouveaux utilisateurs
# Dans unattend.xml (passes specialize/oobeSystem) :
# <InputLocale>040c:0000040c</InputLocale>  → 040c = fr-FR

# Copier les paramètres régionaux actuels vers l'écran d'accueil et les nouveaux comptes
# Panneau de configuration → Région → Administration → "Copier les paramètres"
# Équivalent automatisé (sysprep / MDT) :
control.exe intl.cpl
```

### 67.3 Fuseau horaire

```powershell
# Lister / définir
Get-TimeZone -ListAvailable | Where-Object { $_.Id -like "*Romance*" }
Set-TimeZone -Id "Romance Standard Time"   # Paris (UTC+1, heure d'été gérée)
# Vérifier la synchro NTP
w32tm /query /status
w32tm /resync
```

---

## 68. Télémétrie et confidentialité : niveaux, GPO, registre

### 68.1 Niveaux de télémétrie (données de diagnostic)

| Niveau | Données envoyées | Usage |
|---|---|---|
| 0 — Sécurité (Entreprise/Education uniquement) | Minimales (sécurité) | **Recommandé entreprise** |
| 1 — Basique | + données appareil de base | Minimum Pro |
| 2 — Amélioré | + usage | — |
| 3 — Complet | Tout | Jamais en entreprise |

```powershell
# Forcer le niveau Sécurité (Entreprise uniquement, sinon minimum Basique)
$tele = "HKLM:\SOFTWARE\Policies\Microsoft\Windows\DataCollection"
New-Item -Path $tele -Force | Out-Null
Set-ItemProperty -Path $tele -Name "AllowTelemetry" -Value 0 -Type DWord
# 0=Sécurité, 1=Basique, 2=Amélioré, 3=Complet
# + "LimitDiagnosticLogCollection" = 1, "LimitDumpCollection" = 1
```

### 68.2 Autres verrous de confidentialité

```powershell
# Désactiver la publicité ciblée (ID publicitaire)
New-Item -Path "HKLM:\SOFTWARE\Policies\Microsoft\Windows\AdvertisingInfo" -Force | Out-Null
Set-ItemProperty -Path "HKLM:\SOFTWARE\Policies\Microsoft\Windows\AdvertisingInfo" -Name "DisabledByGroupPolicy" -Value 1 -Type DWord

# Désactiver les expériences personnalisées (déjà vu section 23, rappel)
# Désactiver l'accès des apps aux données sensibles (position, micro, caméra) par défaut :
# → GPO "Confidentialité" : "Autoriser les applications Windows à accéder..." → Forcer le refus
```

> 📌 **Note RGPD** : documentez les niveaux de télémétrie choisis dans votre registre des traitements. Le niveau Sécurité + documentation = posture défendable.

---

## 69. Microsoft 365 Apps : déploiement, canaux, licences

### 69.1 Canaux de mise à jour

| Canal | Fréquence | Cible |
|---|---|---|
| Current Channel | Mensuel | Utilisateurs standard (défaut) |
| Monthly Enterprise Channel | Mensuel (décalé) | **Recommandé entreprise** |
| Semi-Annual Enterprise Channel | 2/an | Postes critiques / validation longue |

### 69.2 Déploiement avec l'Outil de déploiement Office (ODT)

`configuration.xml` type :

```xml
<Configuration>
  <Add OfficeClientEdition="64" Channel="MonthlyEnterprise">
    <Product ID="O365ProPlusRetail">
      <Language ID="fr-fr" />
      <ExcludeApp ID="Groove" />
      <ExcludeApp ID="Lync" />
    </Product>
  </Add>
  <Updates Enabled="TRUE" Channel="MonthlyEnterprise" />
  <Display Level="None" AcceptEULA="TRUE" />
  <Property Name="SharedComputerLicensing" Value="0" />
  <Property Name="AUTOACTIVATE" Value="1" />
</Configuration>
```

```powershell
# Télécharger puis installer
.\setup.exe /download configuration.xml
.\setup.exe /configure configuration.xml
```

### 69.3 Licences et activation

| Mode | Usage |
|---|---|
| Activation basée sur l'utilisateur (défaut) | Poste nominatif, licence M365 |
| **Shared Computer Licensing** | RDS, postes partagés (`SharedComputerLicensing=1`) |
| Clé KMS/MAK (Office LTSC) | Postes sans M365, version perpétuelle |

```powershell
# Vérifier l'état d'activation Office
cscript "C:\Program Files\Microsoft Office\Office16\ospp.vbs" /dstatus
```

---

## 70. Navigateurs en entreprise : Edge géré, stratégies

### 70.1 Pourquoi Edge en entreprise (ou Chrome géré)

Edge = Chromium + intégration Entra/SSO + stratégies GPO/Intune natives + IE Mode (pour les vieilles applications intranet).

### 70.2 Stratégies essentielles (modèles ADMX Edge)

