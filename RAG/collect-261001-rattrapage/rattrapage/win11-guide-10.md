---
id: collect-261001-rattrapage/rattrapage/win11-guide-10
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "copilot"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [1364, 1541]
sha256: a6cd3dd1e964bc594773d5b3b972f4c2cc94313bef27f677df99af9653b88c22
---

# Windows 11 en entreprise — Guide technique ultra-complet

FSLogix encapsule le profil utilisateur dans un **conteneur VHDX** monté à l'ouverture de session. Le profil « suit » l'utilisateur en quelques secondes, sans copie réseau.

**Cas d'usage** : Azure Virtual Desktop, RDS/Terminal Server, postes partagés, environnements non persistants. Sur poste physique fixe : inutile (profil local + OneDrive suffit).

### 20.2 Architecture

```
Ouverture de session
   → FSLogix monte \\SRV-FSLOGIX\Profils$\%username%_%sid%\Profile_%username%.vhdx
   → le VHDX apparaît comme C:\Users\<nom>
   → à la fermeture : démontage, le VHDX reste sur le partage
```

### 20.3 Installation et configuration (résumé)

1. Installer l'agent FSLogix sur l'image/RDS.
2. Configurer via GPO (modèles ADMX FSLogix) ou registre :

```powershell
# Configuration minimale via registre (équivalent GPO)
$reg = "HKLM:\SOFTWARE\FSLogix\Profiles"
New-Item -Path $reg -Force | Out-Null
Set-ItemProperty -Path $reg -Name "Enabled" -Value 1 -Type DWord
Set-ItemProperty -Path $reg -Name "VHDLocations" -Value "\\SRV-FSLOGIX\Profils$" -Type MultiString
Set-ItemProperty -Path $reg -Name "SizeInMBs" -Value 30000 -Type DWord        # 30 Go
Set-ItemProperty -Path $reg -Name "IsDynamic" -Value 1 -Type DWord           # VHDX dynamique
Set-ItemProperty -Path $reg -Name "VolumeType" -Value "vhdx" -Type String
Set-ItemProperty -Path $reg -Name "FlipFlopProfileDirectoryName" -Value 1 -Type DWord
```

3. Droits sur le partage : les utilisateurs doivent pouvoir **créer** leur dossier (Creator Owner : contrôle total sur sous-dossiers).

### 20.4 Bonnes pratiques

| Point | Recommandation |
|---|---|
| Taille | 30 Go par défaut ; surveiller la croissance |
| Redirections | Exclure les caches lourds (Teams, navigateurs) via `redirections.xml` |
| Haute dispo | Partage sur stockage redondé (Storage Spaces Direct, Azure Files) |
| Office | Utiliser le **conteneur Office** séparé si besoin de cache partagé |
| Antivirus | Exclure les VHDX de l'analyse temps réel sur le serveur de fichiers |

---

## 21. Migration de profils : USMT (ScanState / LoadState)

### 21.1 USMT — quand l'utiliser

USMT (User State Migration Tool, inclus dans l'ADK) migre profils et paramètres lors d'un **remplacement de poste** ou d'une **réinstallation**. Pour une simple mise à niveau Windows 10 → 11 sur le même poste : inutile (le setup conserve tout).

### 21.2 ScanState (capture) et LoadState (restauration)

```powershell
# 1. Capture sur l'ancien poste (depuis WinPE ou session admin)
# /i: fichiers de règles, /o: écrase le magasin existant
C:\USMT\amd64\scanstate.exe \\SRV-MIGR\USMT$\%COMPUTERNAME% /i:C:\USMT\amd64\MigApp.xml /i:C:\USMT\amd64\MigDocs.xml /o /c /v:5 /l:C:\Logs\scanstate.log

# Options clés :
# /ue:DOMAINE\* /ui:DOMAINE\jdupont  : exclure tous sauf jdupont
# /c                                 : continuer malgré les erreurs
# /v:5                               : verbosité
# /encrypt /key:xxx                  : chiffrer le magasin (recommandé sur partage réseau)

# 2. Restauration sur le nouveau poste
C:\USMT\amd64\loadstate.exe \\SRV-MIGR\USMT$\ANCIEN-PC /i:C:\USMT\amd64\MigApp.xml /i:C:\USMT\amd64\MigDocs.xml /c /v:5 /l:C:\Logs\loadstate.log
```

### 21.3 Fichiers de règles personnalisés

```xml
<!-- MigCustom.xml : inclure un dossier métier, exclure les caches -->
<migration urlid="http://www.microsoft.com/migration/1.0/migxmlext/custom">
  <component type="Documents" context="User">
    <displayName>Dossier métier</displayName>
    <role role="Data">
      <rules>
        <include>
          <objectSet><pattern type="File">%CSIDL_LOCAL_APPDATA%\MonAppMetier\* [*]</pattern></objectSet>
        </include>
        <exclude>
          <objectSet><pattern type="File">%CSIDL_LOCAL_APPDATA%\*\Cache\* [*]</pattern></objectSet>
        </exclude>
      </rules>
    </role>
  </component>
</migration>
```

### 21.4 Alternative moderne

Pour les postes Entra join : **OneDrive Known Folder Move** + réinstallation des applications via Intune/winget remplace USMT dans 90 % des cas, sans infrastructure de magasin.

---

## 22. GPO Windows 11 : menu Démarrer, barre des tâches, disposition

### 22.1 Verrouiller la disposition du menu Démarrer

Windows 11 utilise un fichier JSON de disposition (`LayoutModification.json`), plus le XML historique :

```powershell
# Exporter la disposition d'un poste modèle
Export-StartLayout -Path C:\Admin\StartLayout.json
```

GPO : `Configuration utilisateur\Stratégies\Modèles d'administration\Menu Démarrer et barre des tâches` → **« Configurer la disposition de l'écran de démarrage »** → chemin du JSON (partage réseau en lecture seule).

Exemple de `StartLayout.json` minimal :

```json
{
  "pinnedList": [
    { "desktopAppLink": "%ALLUSERSPROFILE%\\Microsoft\\Windows\\Start Menu\\Programs\\Microsoft Edge.lnk" },
    { "desktopAppLink": "%APPDATA%\\Microsoft\\Windows\\Start Menu\\Programs\\Word.lnk" },
    { "packagedAppId": "Microsoft.WindowsCalculator_8wekyb3d8bbwe!App" }
  ]
}
```

### 22.2 Barre des tâches : épingles et alignement

| Réglage | GPO / Registre |
|---|---|
| Alignement (gauche vs centré) | `HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced\TaskbarAl` (0 = gauche, 1 = centré) |
| Épingles imposées | Même JSON de disposition (section `taskbar` via stratégies MDM, ou GPO « Configurer la barre des tâches ») |
| Empêcher l'utilisateur de modifier | `LockTaskbar` / stratégies « Verrouiller la barre des tâches » |

```powershell
# Forcer l'alignement à gauche pour tous les nouveaux profils (image de référence)
$reg = "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced"
# Via le profil par défaut (Default User) :
reg load HKU\DefaultUser "C:\Users\Default\NTUSER.DAT"
reg add "HKU\DefaultUser\Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced" /v TaskbarAl /t REG_DWORD /d 0 /f
reg unload HKU\DefaultUser
```

### 22.3 Menu Démarrer : recommandations entreprise

- Épingler : Edge/Chrome géré, Office, portail d'entreprise, outil de tickets.
- **Retirer** : jeux, applications grand public, suggestions.
- Ne pas verrouiller à 100 % : laissez une zone libre pour l'appropriation (moins de tickets au support).

---

## 23. GPO Windows 11 : désactiver Copilot, Widgets, recommandations

### 23.1 Désactiver Windows Copilot (entreprise)

> ⚠️ Selon la build (23H2 vs 24H2) et la région, Copilot peut apparaître comme application ou comme bouton de barre des tâches. Appliquez les deux verrous.

| Méthode | Chemin |
|---|---|
| GPO | `Configuration ordinateur\Modèles d'administration\Composants Windows\Windows Copilot` → **« Désactiver Windows Copilot »** : Activé |
| Registre | `HKLM\SOFTWARE\Policies\Microsoft\Windows\WindowsCopilot\TurnOffWindowsCopilot` = 1 (DWORD) |
| Intune | Catalogue de paramètres → Windows AI / Copilot |

```powershell
# Via registre (déploiement scripté)
New-Item -Path "HKLM:\SOFTWARE\Policies\Microsoft\Windows\WindowsCopilot" -Force | Out-Null
Set-ItemProperty -Path "HKLM:\SOFTWARE\Policies\Microsoft\Windows\WindowsCopilot" `
    -Name "TurnOffWindowsCopilot" -Value 1 -Type DWord
```

### 23.2 Désactiver les Widgets et le fil d'actualités

```powershell
# Widgets : GPO "Autoriser les widgets" → Désactivé
# (Modèles d'administration\Composants Windows\Widgets)
New-Item -Path "HKLM:\SOFTWARE\Policies\Microsoft\Dsh" -Force | Out-Null
Set-ItemProperty -Path "HKLM:\SOFTWARE\Policies\Microsoft\Dsh" `
    -Name "AllowNewsAndInterests" -Value 0 -Type DWord

# Bouton Widgets de la barre des tâches (par défaut utilisateur) :
reg add "HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced" /v TaskbarDa /t REG_DWORD /d 0 /f
```

### 23.3 Recommandations, suggestions, publicité dans le Démarrer

