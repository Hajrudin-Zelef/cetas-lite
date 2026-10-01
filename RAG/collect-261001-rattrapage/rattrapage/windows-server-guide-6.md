---
id: collect-261001-rattrapage/rattrapage/windows-server-guide-6
title: "Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-01-01"]
keywords: ["backlog"]
source: docs/RAG/collect-261001-rattrapage/windows_server_guide.md
source_anchor: ""
source_lines: [937, 1092]
sha256: 8dabdd4efca4879f06dbfcdf87abbddd18fa631c0a698539ad4cab9c4dacd768
---

# Windows Server en entreprise — Guide technique ultra-complet

**Méthode recommandée :**
1. Partage : `Tout le monde` (ou `Utilisateurs authentifiés`) = **Contrôle total**.
2. NTFS : permissions fines par groupe (Lecture / Modification / Contrôle total).
3. Groupes AD par rôle : `GRP-Partage-Compta-L` (lecture), `GRP-Partage-Compta-M` (modification).

```powershell
# Poser des permissions NTFS propres (héritage désactivé, groupes uniquement)
$path = "D:\Partages\Compta"
$acl = Get-Acl $path
$acl.SetAccessRuleProtection($true, $false)  # désactive l'héritage, ne copie pas les règles existantes
$acl.Access | ForEach-Object { $acl.RemoveAccessRule($_) } | Out-Null
$acl.AddAccessRule((New-Object System.Security.AccessControl.FileSystemAccessRule(
  "ENTREPRISE\GRP-Compta-M", "Modify", "ContainerInherit,ObjectInherit", "None", "Allow")))
$acl.AddAccessRule((New-Object System.Security.AccessControl.FileSystemAccessRule(
  "ENTREPRISE\GRP-Compta-L", "ReadAndExecute", "ContainerInherit,ObjectInherit", "None", "Allow")))
$acl.AddAccessRule((New-Object System.Security.AccessControl.FileSystemAccessRule(
  "AUTORITE NT\Système", "FullControl", "ContainerInherit,ObjectInherit", "None", "Allow")))
$acl.AddAccessRule((New-Object System.Security.AccessControl.FileSystemAccessRule(
  "ENTREPRISE\Admins du domaine", "FullControl", "ContainerInherit,ObjectInherit", "None", "Allow")))
Set-Acl -Path $path -AclObject $acl

# Auditer les permissions effectives d'un utilisateur
Get-Acl $path | Select-Object -ExpandProperty Access | Format-Table IdentityReference, FileSystemRights, AccessControlType
```

> **Jamais** de permissions à un utilisateur nommé directement : toujours via des groupes AD. Le jour du départ d'un salarié, on retire le groupe, pas 47 ACL.

---

## 32. Cas pratiques de permissions (tableau de vérité)

| Besoin | Partage | NTFS |
|---|---|---|
| Dossier lecture seule pour tous | Tout le monde = CT | Utilisateurs du domaine = Lecture |
| Écriture pour un service, lecture pour les autres | Tout le monde = CT | GRP-Service = Modification ; Utilisateurs = Lecture |
| Dossier confidentiel RH | GRP-RH = CT (partage restreint) | GRP-RH = Modification ; héritage coupé |
| Dépôt sans suppression (writ-only log) | Tout le monde = CT | Créateur propriétaire = Contrôle total ; autres = Écriture seule (droits spéciaux) |
| Homedirs `\\srv\users\%username%` | Utilisateurs authentifiés = CT | Chaque dossier : utilisateur = Modification, héritage coupé |

```powershell
# Créer des dossiers personnels en masse (à partir d'une liste d'utilisateurs)
$users = Get-ADUser -Filter * -SearchBase "OU=Utilisateurs,DC=entreprise,DC=lan" | Select-Object -ExpandProperty SamAccountName
foreach ($u in $users) {
    $dir = "D:\Partages\Users\$u"
    if (-not (Test-Path $dir)) { New-Item -ItemType Directory -Path $dir | Out-Null }
    $acl = Get-Acl $dir
    $acl.SetAccessRuleProtection($true, $false)
    $acl.AddAccessRule((New-Object System.Security.AccessControl.FileSystemAccessRule(
      "ENTREPRISE\$u", "Modify", "ContainerInherit,ObjectInherit", "None", "Allow")))
    Set-Acl -Path $dir -AclObject $acl
}
New-SmbShare -Name "Users" -Path "D:\Partages\Users" -FullAccess "ENTREPRISE\Utilisateurs authentifiés"
```

---

## 33. DFS-N : espaces de noms

DFS-N = un chemin logique unique (`\\entreprise.lan\data\compta`) qui pointe vers un ou plusieurs partages physiques, avec bascule.

```powershell
# Installer
Install-WindowsFeature -Name FS-DFS-Namespace -IncludeManagementTools

# Créer un espace de noms de domaine
New-DfsnRoot -Path "\\entreprise.lan\data" -TargetPath "\\SRV-FICHIER-01\data" -Type DomainV2

# Ajouter un dossier + cibles (2e cible = haute dispo / site distant)
New-DfsnFolder -Path "\\entreprise.lan\data\compta" -TargetPath "\\SRV-FICHIER-01\compta"
New-DfsnFolderTarget -Path "\\entreprise.lan\data\compta" -TargetPath "\\SRV-FICHIER-02\compta"

# Lister
Get-DfsnRoot | Format-Table Path, State
Get-DfsnFolder -Path "\\entreprise.lan\data\*" | Format-Table Path, State
```

**Pourquoi DFS-N même avec un seul serveur :** abstraction. Le jour où tu migres `SRV-FICHIER-01` → `SRV-FICHIER-03`, les utilisateurs gardent `\\entreprise.lan\data\...` : zéro changement de lecteurs mappés, zéro GPO à retoucher.

---

## 34. DFS-R : réplication

DFS-R réplique le contenu entre membres d'un groupe de réplication (RDC = ne transfère que les blocs modifiés).

```powershell
# Installer
Install-WindowsFeature -Name FS-DFS-Replication -IncludeManagementTools

# Créer un groupe de réplication (ex. entre 2 serveurs de fichiers)
New-DfsReplicationGroup -GroupName "RG-Data"
Add-DfsrMember -GroupName "RG-Data" -ComputerName "SRV-FICHIER-01","SRV-FICHIER-02"
New-DfsReplicatedFolder -GroupName "RG-Data" -FolderName "Compta" -DfsnPath "\\entreprise.lan\data\compta"
Set-DfsrMembership -GroupName "RG-Data" -FolderName "Compta" -ComputerName "SRV-FICHIER-01" `
  -ContentPath "D:\Partages\Compta" -PrimaryMember $true -StagingPathQuotaInMB 16384
Set-DfsrMembership -GroupName "RG-Data" -FolderName "Compta" -ComputerName "SRV-FICHIER-02" `
  -ContentPath "D:\Partages\Compta" -StagingPathQuotaInMB 16384

# Suivi
Get-DfsrState -ComputerName SRV-FICHIER-01 | Format-Table FolderName, State -AutoSize
dfsrdiag backlog /rgname:"RG-Data" /rfname:"Compta" /smem:SRV-FICHIER-01 /rmem:SRV-FICHIER-02
```

**Limites à connaître :**
- Fichiers **ouverts/verrouillés** : conflit = le "perdant" va dans `DfsrPrivate\ConflictAndDeleted` (quota par défaut 660 Mo — à augmenter).
- **Ne pas** répliquer des bases de données ouvertes (SQL, etc.) ni des VHDX de VM.
- Taille max de fichier : ça passe, mais la réplication initiale d'un gros volume prend des heures — pré-répliquer via disque externe (`robocopy` + pré-amorçage).

---

## 35. FSRM : quotas et filtrage de fichiers

```powershell
Install-WindowsFeature -Name FS-Resource-Manager -IncludeManagementTools

# --- QUOTAS ---
# Modèle : 10 Go avec seuil d'alerte à 85 %
New-FsrmQuotaTemplate -Name "Standard 10 Go" -Size 10GB
# Appliquer à un dossier
New-FsrmQuota -Path "D:\Partages\Users" -Template "Standard 10 Go"

# Quota dur (bloque) vs souple (alerte seule) : le template "Standard" est dur par défaut.
# Notification e-mail à 90 % :
# (via la console FSRM → Modèles de quota → Seuils, ou en XML d'action)

# --- FILTRAGE DE FICHIERS (file screening) ---
# Bloquer les fichiers multimédias et exécutables sur un partage de travail
New-FsrmFileScreenTemplate -Name "Blocage Multimédia+Exe" -IncludeGroup "Fichiers audio et vidéo","Fichiers exécutables"
New-FsrmFileScreen -Path "D:\Partages\Commun" -Template "Blocage Multimédia+Exe" -Active:$true

# Groupes de fichiers utiles prédéfinis : "Fichiers audio et vidéo", "Fichiers image",
# "Fichiers exécutables", "Fichiers compressés", "Fichiers de sauvegarde"...

# Rapport d'usage
New-FsrmStorageReport -Name "Rapport Quotas" -Namespace @("D:\Partages") -ReportType @("QuotaUsage")
```

**Cas d'usage n°1 en entreprise :** le file screening anti-ransomware — bloquer les extensions exotiques (`*.locked`, `*.crypt`, etc.) en alerte passive pour détecter une propagation.

---

## 36. Déduplication des données

Gain typique : 30–70 % sur des partages bureautiques (beaucoup de doublons : mêmes pièces jointes, mêmes ISO, mêmes installateurs).

```powershell
# Installer + activer sur un volume (JAMAIS sur le volume système, JAMAIS sur un CSV)
Install-WindowsFeature -Name FS-Data-Deduplication
Enable-DedupVolume -Volume "D:" -UsageType Default

# Planifier l'optimisation (par défaut : tâche planifiée, à ajuster)
Set-DedupSchedule -Name "BackgroundOptimization" -Start (Get-Date "2026-01-01 02:00") -DurationHours 6 -Days Monday,Tuesday,Wednesday,Thursday,Friday

# Forcer une optimisation + mesurer le gain
Start-DedupJob -Volume "D:" -Type Optimization -Priority High
Get-DedupStatus -Volume "D:" | Format-List Volume, SavingsRate, SavedSpace, OptimizedFilesCount

