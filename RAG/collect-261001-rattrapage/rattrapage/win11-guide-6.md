---
id: collect-261001-rattrapage/rattrapage/win11-guide-6
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [574, 792]
sha256: 88ccca9ec167fdcbef659f3b27ab105c80b2b4b8fa146950dafcafa1ed8736b3
---

# Windows 11 en entreprise — Guide technique ultra-complet

```powershell
# Activation basée sur Active Directory (ADBA) — côté serveur, une fois :
# 1. Installer le rôle "Services d'activation en volume" sur un DC/serveur
# 2. slmgr.vbs /ipk <clé hôte KMS> puis /ato
# Côté client : rien à faire, l'activation est automatique à la jointure du domaine.
```

### 8.4 setupconfig.ini (alternative légère à unattend.xml)

Pour une mise à niveau pilotée sans XML complet :

```ini
[SetupConfig]
ShowOobe=None
Telemetry=Disable
DynamicUpdate=Enable
Compat=IgnoreWarning
PostOobe=C:\Admin\post-install.cmd
```

À placer à côté de `setup.exe` ou dans `C:\Users\Default\AppData\Local\Microsoft\Windows\WSUS\SetupConfig.ini`.

---

## 9. Sysprep : généralisation, audit mode, capture d'image

### 9.1 Les 3 modes de Sysprep

| Mode | Commande | Usage |
|---|---|---|
| OOBE + Generalize | `sysprep /generalize /oobe /shutdown` | **Image maître** : nettoie les SID, prépare la capture |
| Audit | `sysprep /audit /reboot` | Personnalisation en mode audit (compte Administrator intégré) |
| OOBE simple | `sysprep /oobe /shutdown` | Sans généralisation (même machine) |

```powershell
# La commande de référence avant capture :
C:\Windows\System32\Sysprep\Sysprep.exe /generalize /oobe /shutdown /unattend:C:\Admin\unattend.xml
```

> ⛔ **Limites impératives** :
> - Sysprep **échoue** si des applications du Store ont été mises à jour pour un utilisateur (erreur `0x80073cf2`). Désinstaller ou ne pas lancer le Store avant sysprep.
> - Maximum **3** réarmements (`slmgr /rearm`) — au-delà, sysprep refuse.
> - Ne jamais sysprepper une machine déjà en production ou jointe au domaine pour en faire une image (problèmes de SID/relations d'approbation).

### 9.2 Mode audit pas à pas

1. Sur le poste de référence, à l'écran OOBE : **Ctrl+Shift+F3** → redémarre en mode audit (compte Administrator).
2. Installer applications, pilotes, personnaliser le profil par défaut si besoin.
3. Nettoyer : vider les journaux, supprimer les fichiers temporaires.
4. `sysprep /generalize /oobe /shutdown`.
5. Capturer (section suivante).

```powershell
# En mode audit : copier un profil personnalisé vers le profil par défaut
# (via unattend.xml, passe specialize) :
# <CopyProfile>true</CopyProfile>
# ATTENTION : CopyProfile est déprécié dans ses effets sur le menu Démarrer
# moderne — préférez les GPO de personnalisation (section 22).
```

### 9.3 Capture de l'image avec DISM

Depuis WinPE (clé bootable) ou un autre OS :

```powershell
# Capturer C: vers un WIM
dism /Capture-Image /ImageFile:D:\Images\Win11-Ref.wim /CaptureDir:C:\ /Name:"Win11 24H2 Ref" /Compress:max /CheckIntegrity /Verify

# Vérifier le contenu
dism /Get-WimInfo /WimFile:D:\Images\Win11-Ref.wim

# Appliquer l'image sur un poste cible (depuis WinPE)
diskpart /s C:\Admin\partitions-uefi.txt
dism /Apply-Image /ImageFile:D:\Images\Win11-Ref.wim /Index:1 /ApplyDir:C:\
bcdboot C:\Windows
```

Fichier `partitions-uefi.txt` type (GPT/UEFI) :

```
select disk 0
clean
convert gpt
create partition efi size=100
format quick fs=fat32 label="System"
assign letter="S"
create partition msr size=16
create partition primary
format quick fs=ntfs label="Windows"
assign letter="C"
exit
```

---

## 10. MDT : installation, deployment share, task sequences

### 10.1 Prérequis et installation

| Composant | Version / remarque |
|---|---|
| Windows ADK | Version correspondant à Windows 11 (ex. ADK pour 24H2) |
| WinPE Add-on pour l'ADK | **Obligatoire** (boot images) |
| MDT | Dernière build (8456+) |
| Serveur | Windows Server 2019/2022/2025, partage SMB |

Ordre d'installation : **ADK → WinPE Add-on → MDT**. Puis console « Deployment Workbench ».

```powershell
# Partage de déploiement : créer le dossier et le partager
New-Item -Path D:\DeploymentShare -ItemType Directory
New-SmbShare -Name "DeploymentShare$" -Path D:\DeploymentShare -FullAccess "Administrateurs"
# (Affiner les ACL : lecture pour "Utilisateurs authentifiés", écriture pour le compte MDT)
```

### 10.2 Créer le deployment share

1. Deployment Workbench → clic droit « Deployment Shares » → New.
2. Chemin : `D:\DeploymentShare`, nom de partage `DeploymentShare$`.
3. Options : cocher les plateformes x64, désactiver x86 (sauf besoin legacy).
4. Le partage génère les images WinPE (`LiteTouchPE_x64.wim`) dans `Boot\`.

```powershell
# Mettre à jour le deployment share après chaque modification (régénère WinPE)
# Via la console : clic droit → "Update Deployment Share" → "Completely regenerate"
```

### 10.3 Importer l'OS

Deployment Workbench → Operating Systems → Import :

- **Full set of source files** : depuis l'ISO VLSC extraite (recommandé).
- **Custom image file** : votre WIM capturé (section 9).
- **WIM existant** du deployment share.

### 10.4 Task sequence « Standard Client »

Clic droit Task Sequences → New → **Standard Client Task Sequence** :

1. ID : `WIN11-24H2`, nom : « Windows 11 24H2 Entreprise ».
2. OS : sélectionner l'image importée.
3. Clé produit : ne rien saisir (activation via KMS/ADBA après) ou GVLK.
4. Compte admin local : définir un mot de passe **fictif/temporaire** (sera écrasé par LAPS ensuite).
5. Étapes par défaut : validation, formatage, application de l'image, pilotes, Windows Update, applications.

Structure d'une task sequence (onglet Task Sequence) :

```
Initialization
  └─ Gather local only
Validation
  └─ Validate (BIOS/UEFI, mémoire...)
State Capture (refresh uniquement)
Preinstall
  ├─ New Computer only → Format and Partition Disk (UEFI)
  └─ Apply Operating System
Postinstall
  ├─ Inject Drivers
  ├─ Apply Windows PE / Apply Network Settings
  └─ Apply Local GPO Package (optionnel)
State Restore
  ├─ Windows Update (Pre-Application / Post-Application)
  ├─ Install Applications
  └─ Tattoo (écrit les infos de déploiement dans le registre)
```

---

## 11. MDT : Lite Touch, pilotes, règles CustomSettings.ini

### 11.1 Boot Lite Touch

1. Copier `D:\DeploymentShare\Boot\LiteTouchPE_x64.wim` vers votre serveur WDS, ou créer une clé USB bootable (`LiteTouchPE_x64.iso`).
2. Au boot : WinPE → assistant Lite Touch → choix de la task sequence → identifiants (`DOMAIN\svc-mdt`) → déploiement.

Avec WDS : ajouter l'image de démarrage, configurer DHCP (options 66/67) ou IP Helper vers le serveur WDS.

### 11.2 Gestion des pilotes (par modèle)

Bonne pratique : **un dossier par modèle**, sélection par WMI :

```
Out-of-Box Drivers
 └─ Win11-x64
     ├─ Dell-Latitude-5440
     ├─ HP-EliteBook-840-G10
     └─ Lenovo-ThinkPad-T14-G4
```

Dans la task sequence, étape « Inject Drivers » → « Install all drivers from the selection profile » avec un **profil de sélection** par modèle, ou « Install only matching drivers from the selection profile ».

```powershell
# Importer des pilotes en masse dans MDT (à exécuter sur le serveur)
$drivers = "D:\DriversSCCM\Dell-Latitude-5440\Win11"
Import-MDTDriver -Path "DS001:\Out-of-Box Drivers\Win11-x64\Dell-Latitude-5440" -SourcePath $drivers -ImportDuplicates $false
```

```powershell
# Identifier le modèle exact (pour nommer le dossier)
Get-CimInstance Win32_ComputerSystem | Select-Object Manufacturer, Model
```

> ⚠️ **Ne jamais** injecter « tous les pilotes pour tout le monde » : conflits, BSOD, task sequences interminables. Un profil de sélection par modèle + requête WMI, c'est la règle.

### 11.3 CustomSettings.ini — le fichier qui automatise tout

`D:\DeploymentShare\Control\CustomSettings.ini` :

```ini
[Settings]
Priority=Model, Default
Properties=MyCustomProperty

[Dell Inc.-Latitude 5440]
DriverGroup001=Dell-Latitude-5440

[HP-EliteBook 840 G10]
DriverGroup001=HP-EliteBook-840-G10

