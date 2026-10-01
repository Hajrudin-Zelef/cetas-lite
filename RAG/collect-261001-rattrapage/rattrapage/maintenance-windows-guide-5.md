---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-5
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr", "parameters"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [709, 939]
sha256: db2fdd79705a095144d776b223c259cfce7ba880dbe5a6b420ae6b1971188a1a
---

# Maintenance et exploitation Windows en entreprise

**Si le poste n'a pas accès à Windows Update** (ou si la réparation échoue avec
`0x800f081f` = source introuvable), fournissez une source saine (ISO de la **même
build**) :

```cmd
:: Monter l'ISO puis utiliser install.wim / install.esd comme source
DISM /Online /Cleanup-Image /RestoreHealth /Source:WIM:D:\sources\install.wim:1 /LimitAccess
:: Variante avec un ESD (éditions grand public)
DISM /Online /Cleanup-Image /RestoreHealth /Source:ESD:D:\sources\install.esd:1 /LimitAccess
```

> L'index (`:1`) doit correspondre à **l'édition exacte** (Pro, Entreprise…).
> Vérifiez avec `DISM /Get-WimInfo /WimFile:D:\sources\install.wim`.
> `/LimitAccess` empêche DISM d'aller chercher sur Windows Update.

Journal détaillé : `C:\Windows\Logs\DISM\dism.log`.

---

## 23. DISM : monter et réparer une image WIM

Utile pour réparer une image de référence (déploiement) ou extraire un fichier
système sain :

```cmd
:: Monter l'image en lecture/écriture
DISM /Mount-Wim /WimFile:D:\images\ref.wim /index:1 /MountDir:C:\mnt\wim

:: ... modifications (ajout de pilotes, correctifs, fichiers) ...

:: Démonter en enregistrant
DISM /Unmount-Wim /MountDir:C:\mnt\wim /Commit
:: Ou abandonner les modifications
DISM /Unmount-Wim /MountDir:C:\mnt\wim /Discard
```

Ajouter un correctif à une image montée :

```cmd
DISM /Image:C:\mnt\wim /Add-Package /PackagePath:D:\patches\windows11.0-kb5034765-x64.msu
```

---

## 24. SFC /scannow : mode d'emploi

Le **Vérificateur des fichiers système** compare les fichiers protégés à leur
empreinte et restaure les versions saines depuis le magasin de composants.

```cmd
:: Analyse + réparation (redémarrage parfois nécessaire à la fin)
sfc /scannow

:: Analyse seule (idéal en diagnostic, sans modification)
sfc /verifyonly

:: Réparer un fichier précis (chemin complet requis)
sfc /scanfile=C:\Windows\System32\kernel32.dll
```

- Durée : 10 à 30 minutes. Ne pas interrompre.
- Résultat « des fichiers endommagés ont été réparés » → redémarrez et relancez
  une fois pour confirmer (« aucune violation d'intégrité »).
- Si SFC échoue de façon répétée : passez par DISM (§22) **avant** de relancer SFC
  (SFC puise dans le magasin que DISM répare).
- Journal : `C:\Windows\Logs\CBS\CBS.log` (verbeux ; filtrez sur `[SR]`).

Extraction des lignes SFC du CBS.log :

```powershell
Select-String -Path C:\Windows\Logs\CBS\CBS.log -Pattern '\[SR\]' |
  Select-Object -ExpandProperty Line | Out-File C:\Temp\sfc-sr.txt
```

---

## 25. CHKDSK et vérification du système de fichiers

```cmd
:: Vérification en lecture seule (sans démontage)
chkdsk C:

:: Réparation des erreurs logiques (redémarrage requis sur le volume système)
chkdsk C: /f

:: + recherche des secteurs défectueux et récupération (long sur gros volumes)
chkdsk D: /f /r

:: NTFS uniquement : répare en ligne sans démontage quand possible
chkdsk C: /scan
chkdsk C: /spotfix
```

Équivalent PowerShell (plus scriptable) :

```powershell
# État de santé logique du volume
Repair-Volume -DriveLetter C -Scan
# Réparation en ligne
Repair-Volume -DriveLetter C -SpotFix
```

**Règles pratiques :**

- Un `chkdsk /r` sur un volume de plusieurs To peut durer **des heures** : planifiez
  en fenêtre de maintenance.
- Des erreurs CHKDSK récurrentes = disque en fin de vie → contrôlez le SMART (§52)
  et préparez le remplacement, pas le 4ᵉ chkdsk.
- Sur **ReFS** (Storage Spaces, §53), `chkdsk` est remplacé par la réparation
  automatique (integrity streams) ; `Repair-Volume` reste utilisable.

---

## 26. Nettoyage disque : cleanmgr et Disk Cleanup

> ⚠️ `cleanmgr.exe` est **déprécié** (toujours présent sur Server 2019/2022/2025 et
> Windows 11, mais sans évolution). Privilégiez Storage Sense (§27) pour
> l'automatisation et DISM pour WinSxS (§29).

Nettoyage manuel assisté :

```cmd
:: Ouvre le sélecteur de catégories (à cocher une fois)
cleanmgr /sageset:10
:: Exécute le nettoyage avec le profil n°10 (scriptable)
cleanmgr /sagerun:10
```

Catégories utiles sur serveur : *Fichiers temporaires*, *Fichiers journaux*,
*Nettoyage de Windows Update*, *Fichiers d'optimisation de livraison*,
*Corbeille*, *Miniatures*.

**Automatisation** (tâche planifiée mensuelle, §43) :

```powershell
# Profil silencieux : à exécuter via une tâche planifiée en SYSTEM
Start-Process cleanmgr.exe -ArgumentList '/sagerun:10' -Wait
```

Vérifiez toujours **avant/après** : `Get-PSDrive C` ou §31. Un nettoyage qui ne
libère rien signale un autre problème (logs qui regrossissent, §30).

---

## 27. Storage Sense : configuration centralisée

Storage Sense (Assistant stockage) automatise le nettoyage : fichiers temp,
corbeille, contenu du dossier Téléchargements, OneDrive à la demande.

GPO (Windows 10/11, partiellement Server) :

```text
Configuration ordinateur > Stratégies > Modèles d'administration > Système >
Assistant stockage
  - Autoriser l'Assistant stockage : Activé
  - Configurer la cadence de nettoyage : Chaque mois
  - Supprimer les fichiers de la corbeille après : 30 jours
  - Supprimer les fichiers du dossier Téléchargements après : 60 jours
```

État et déclenchement manuel en PowerShell :

```powershell
# L'assistant est-il activé ?
Get-ItemProperty 'HKCU:\SOFTWARE\Microsoft\Windows\CurrentVersion\StorageSense\Parameters\StoragePolicy' |
  Select-Object 01, 04, 08
# Forcer un passage (Windows 11)
Start-Process "$env:SystemRoot\System32\cleanmgr.exe" -ArgumentList '/verylowdisk' -Wait
```

> Sur **Windows Server**, Storage Sense est limité : préférez des tâches planifiées
> maison (§43) qui purgent les dossiers identifiés au §28.

---

## 28. Dossiers temporaires et caches à purger

| Emplacement | Contenu | Purge sûre ? |
|---|---|---|
| `C:\Windows\Temp` | Temp système | Oui (fichiers non verrouillés) |
| `C:\Windows\SoftwareDistribution\Download` | Cache Windows Update | Oui **si** le service `wuauserv` est arrêté |
| `%TEMP%` / `%LOCALAPPDATA%\Temp` | Temp utilisateur | Oui (session fermée de préférence) |
| `C:\Windows\Logs\CBS` | Logs CBS/DISM | Oui (anciens `.log`, garder le courant) |
| `C:\inetpub\logs\LogFiles` | Logs IIS | Oui avec rotation (§30) |
| `C:\ProgramData\Microsoft\Windows\WER` | Rapports d'erreurs Windows | Oui (anciens) |
| Corbeilles `C:\$Recycle.Bin` | — | Oui via `Clear-RecycleBin` |

Script de purge mensuelle (tâche SYSTEM, à adapter) :

```powershell
# Purge sécurisée des temporaires (ne supprime que les fichiers de +7 jours non verrouillés)
$paths = @("$env:SystemRoot\Temp", "$env:TEMP")
foreach ($p in $paths) {
    Get-ChildItem $p -Recurse -Force -ErrorAction SilentlyContinue |
      Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-7) } |
      Remove-Item -Force -Recurse -ErrorAction SilentlyContinue
}
Clear-RecycleBin -Force -ErrorAction SilentlyContinue

# Purge du cache Windows Update (service arrêté au préalable)
Stop-Service wuauserv -Force
Remove-Item "$env:SystemRoot\SoftwareDistribution\Download\*" -Recurse -Force -ErrorAction SilentlyContinue
Start-Service wuauserv
```

> ⚠️ Ne purgez **jamais** `C:\Windows\WinSxS` à la main : utilisez DISM (§29).
> Ne touchez pas non plus à `C:\Windows\System32\config` ni aux profils utilisateurs
> actifs.

---

## 29. WinSxS : analyse et nettoyage

Analyse de l'encombrement réel :

```cmd
DISM /Online /Cleanup-Image /AnalyzeComponentStore
```

Exemple de sortie : « Taille réelle du magasin », « sauvegardes et fonctionnalités
désactivées » récupérables. Nettoyage :

```cmd
:: Nettoyage standard des versions remplacées
DISM /Online /Cleanup-Image /StartComponentCleanup

:: Nettoyage agressif : supprime TOUTES les versions remplacées
:: (rend impossible la désinstallation des correctifs !)
DISM /Online /Cleanup-Image /StartComponentCleanup /ResetBase
```

