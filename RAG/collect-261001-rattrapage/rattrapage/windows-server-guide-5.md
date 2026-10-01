---
id: collect-261001-rattrapage/rattrapage/windows-server-guide-5
title: "Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "arr", "memory"]
source: docs/RAG/collect-261001-rattrapage/windows_server_guide.md
source_anchor: ""
source_lines: [745, 936]
sha256: b03548008e7f315eb0d5a551d7e4df1e5e37dfa3afdf274455760c3dda4d5df1
---

# Windows Server en entreprise — Guide technique ultra-complet

1. Ouvrir la console **Windows Server Update Services**.
2. Assistant de configuration : amont = **Microsoft Update** (ou serveur WSUS parent en aval/réplica).
3. Choisir les **produits** (ex. Windows Server 2022, Windows 11, Office) — ne cocher que l'utile.
4. Choisir les **classifications** : Critical Updates, Security Updates, Definition Updates, Updates, Upgrades (avec parcimonie).
5. Planifier la **synchronisation** (ex. 1×/jour à 3h).
6. Première synchro : longue (plusieurs heures la première fois).

```powershell
# Forcer une synchronisation
$wsus = Get-WsusServer -Name "localhost" -PortNumber 8530
$wsus.GetSubscription().StartSynchronization()

# Vérifier l'état de la synchro
$wsus.GetSubscription().GetLastSynchronizationInfo()
```

> Ports : **8530 (HTTP)** / 8531 (HTTPS) par défaut depuis WSUS 2012. Les vieux guides parlent du port 80 : à oublier.

---

## 26. Groupes d'ordinateurs et ciblage côté client

**Côté serveur :** créer des groupes (ex. `Serveurs-Production`, `Serveurs-Test`, `Postes-Pilotes`, `Postes-Standard`).

**Côté client :** ciblage via **GPO** (recommandé) :

```
Configuration ordinateur → Stratégies → Modèles d'administration →
Composants Windows → Windows Update → Gérer les mises à jour proposées par WSUS
  • Spécifier l'emplacement du service : http://srv-wsus-01:8530
  • Activer le ciblage côté client → nom du groupe : Serveurs-Test
```

```powershell
# Vérifier côté client que la GPO s'applique
gpresult /r | Select-String "WSUS|ciblage" -Context 2
Get-ItemProperty "HKLM:\SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate" -Name WUServer, TargetGroup -ErrorAction SilentlyContinue

# Forcer le client à se signaler au WSUS
$updateSession = New-Object -ComObject Microsoft.Update.Session
wuauclt /reportnow /detectnow   # legacy, toujours utile
UsoClient.exe StartScan         # Windows 10/11 / Server 2019+
```

> Le client apparaît dans la console WSUS sous **Ordinateurs → Tous les ordinateurs** après son premier rapport (délai : jusqu'à ~24 h ou forcer comme ci-dessus).

---

## 27. Approbations : manuelles et règles automatiques

**Manuelle :** clic droit sur la mise à jour → Approuver → choisir groupes → OK.

**Règles d'approbation automatique** (recommandées) :

1. Console WSUS → **Options → Approbations automatiques**.
2. Nouvelle règle : ex. *"Critiques + Sécurité → Serveurs-Test"*.
3. Deuxième règle : *"Critiques + Sécurité → Serveurs-Production, délai 7 jours"* (échéance différée).

```powershell
# Approuver via PowerShell : toutes les mises à jour de sécurité non approuvées pour un groupe
$wsus = Get-WsusServer -Name "localhost" -PortNumber 8530
$groupe = $wsus.GetComputerTargetGroups() | Where-Object Name -eq "Serveurs-Production"
$updates = $wsus.GetUpdates() | Where-Object {
  $_.Classification.Title -match "Security" -and
  ($_.GetUpdateApprovalActions($groupe) | Where-Object { $_.Action -eq 'Install' }).Count -eq 0 -and
  -not $_.IsDeclined
}
$updates | ForEach-Object { $_.Approve("Install", $groupe) }
```

**Cycle de patch recommandé :** Test (J+2) → Pilotes (J+5) → Production (J+9), avec fenêtre de maintenance et validation applicative entre chaque.

---

## 28. Maintenance : Server Cleanup Wizard et scripts

Sans maintenance, la base WID gonfle (10+ Go) et la console devient inutilisable.

**Assistant de nettoyage du serveur** (Options → Assistant Nettoyage du serveur), mensuel :
- Mises à jour expirées, remplacées, refusées
- Fichiers non nécessaires
- Ordinateurs inactifs (30+ jours)

```powershell
# Nettoyage automatisé mensuel (tâche planifiée)
$wsus = Get-WsusServer -Name "localhost" -PortNumber 8530
$cleanup = $wsus.GetCleanupManager()
$cleanup.PerformCleanup(
  [Microsoft.UpdateServices.Administration.CleanupScope]::new(
    $true,  # CleanUnneededContentFiles
    $true,  # CleanExpiredUpdates
    $true,  # CleanSupersededUpdates
    $true,  # CleanExpiredComputers
    $true,  # DeclineExpiredUpdates
    $true   # DeclineSupersededUpdates
  )
)

# Réindexer la base WID (quand la console rame)
sqlcmd -S \\.\pipe\MICROSOFT##WID\tsql\query -E -Q "USE SUSDB; EXEC sp_MSforeachtable 'DBCC DBREINDEX (''?'')'"
```

> Si la console WSUS met > 30 s à s'ouvrir : base à réindexer + pool IIS à augmenter (private memory limit du `WsusPool` → 4 Go+).

---

## 29. Dépannage WSUS (clients qui ne remontent pas, reset base)

**Client qui ne remonte pas — méthode en 6 étapes :**

```powershell
# 1. Le client vise-t-il le bon WSUS ?
Get-ItemProperty "HKLM:\SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate" -Name WUServer

# 2. Connectivité
Test-NetConnection srv-wsus-01 -Port 8530

# 3. Réinitialiser l'agent Windows Update (le classique qui marche)
Stop-Service wuauserv, bits, cryptsvc -Force
Remove-Item "$env:windir\SoftwareDistribution\*" -Recurse -Force
Remove-Item "$env:windir\System32\catroot2\*" -Recurse -Force   # prudent : arrêter cryptsvc avant
Start-Service wuauserv, bits, cryptsvc
wuauclt /resetauthorization /detectnow

# 4. Doublon de SusClientId (image clonée sans sysprep !)
Remove-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\WindowsUpdate" -Name SusClientId, SusClientIdValidation -ErrorAction SilentlyContinue

# 5. Forcer le rapport
UsoClient.exe StartScan; Start-Sleep 60; wuauclt /reportnow

# 6. Logs
Get-WindowsUpdateLog   # génère WindowsUpdate.log lisible sur le bureau
```

**Reset complet de la base WSUS** (quand SUSDB est corrompue) :

```powershell
# Option nucléaire : réinitialiser le contenu + resynchroniser
& 'C:\Program Files\Update Services\Tools\wsusutil.exe' reset
# Puis forcer une synchro complète. Les approbations sont conservées (base), seul le contenu est re-téléchargé.
```

**Erreur 0x8024401c / timeout :** souvent le pool IIS `WsusPool` en mémoire insuffisante → augmenter la limite mémoire privée à 4–8 Go et recycler.

---

---

## 30. Partages réseau : SMB, création, bonnes pratiques

```powershell
# Créer un partage SMB simple
New-SmbShare -Name "Commun" -Path "D:\Partages\Commun" -FullAccess "ENTREPRISE\Utilisateurs du domaine"

# Partage avec accès restreint + description
New-SmbShare -Name "Compta" -Path "D:\Partages\Compta" -ChangeAccess "ENTREPRISE\GRP-Compta" `
  -Description "Partage comptabilité - accès restreint"

# Partage masqué (avec $) pour l'administration
New-SmbShare -Name "Admin$" -Path "D:\Partages\Admin" -FullAccess "ENTREPRISE\Admins du domaine"

# Lister / modifier / supprimer
Get-SmbShare | Format-Table Name, Path, Description -AutoSize
Set-SmbShare -Name "Commun" -Description "Nouvelle description" -Force
Remove-SmbShare -Name "AncienPartage" -Force
```

**Bonnes pratiques :**
- Permissions de partage : **Tout le monde = Contrôle total**, et gérer le fin via **NTFS** (§31). (L'exception : partages sensibles où on restreint aussi au niveau partage.)
- Activer **l'énumération basée sur l'accès (ABE)** : l'utilisateur ne voit que ce à quoi il a accès.
- Volume dédié, quotas FSRM (§35), antivirus avec exclusions réfléchies.
- Désactiver SMBv1 partout (voir §60) :

```powershell
# Vérifier / désactiver SMBv1
Get-WindowsFeature FS-SMB1
Disable-WindowsOptionalFeature -Online -FeatureName SMB1Protocol -NoRestart
Get-SmbServerConfiguration | Select-Object EnableSMB1Protocol, EnableSMB2Protocol
```

---

## 31. Permissions NTFS vs permissions de partage : la combinaison

**Règle fondamentale : l'accès effectif = l'intersection (la plus restrictive) des deux.**

| | Partage : Contrôle total | Partage : Lecture |
|---|---|---|
| **NTFS : Contrôle total** | Contrôle total | Lecture |
| **NTFS : Modification** | Modification | Lecture |
| **NTFS : Lecture** | Lecture | Lecture |

