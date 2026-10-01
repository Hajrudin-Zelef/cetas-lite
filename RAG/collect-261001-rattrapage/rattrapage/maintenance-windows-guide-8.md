---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-8
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["attention", "incident"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [1334, 1500]
sha256: a6e1d21d0a10d5006c1c56cadf1387010305473f52950e7a3a3c069ccde8f317
---

# Maintenance et exploitation Windows en entreprise

```powershell
# Exemple : purge des logs IIS tous les dimanches à 03:00, en SYSTEM
$action    = New-ScheduledTaskAction -Execute 'powershell.exe' `
               -Argument '-NoProfile -ExecutionPolicy Bypass -File "C:\Scripts\Purge-IISLogs.ps1"'
$trigger   = New-ScheduledTaskTrigger -Weekly -DaysOfWeek Sunday -At '03:00'
$principal = New-ScheduledTaskPrincipal -UserId 'SYSTEM' -LogonType ServiceAccount -RunLevel Highest
$settings  = New-ScheduledTaskSettingsSet -ExecutionTimeLimit '02:00:00' `
               -RestartCount 3 -RestartInterval (New-TimeSpan -Minutes 15) `
               -StartWhenAvailable
Register-ScheduledTask -TaskName 'Maintenance\Purge logs IIS' `
  -Action $action -Trigger $trigger -Principal $principal -Settings $settings `
  -Description 'Purge des logs IIS de +30 jours (guide §30)'
```

Gestion courante :

```powershell
Get-ScheduledTask -TaskPath '\Maintenance\' | Get-ScheduledTaskInfo |
  Select-Object TaskName, LastRunTime, LastTaskResult
# LastTaskResult = 0 (0x0) : OK ; 0x1 : code de sortie 1 ; autre : à investiguer

# Exécuter immédiatement / désactiver / supprimer
Start-ScheduledTask -TaskName 'Maintenance\Purge logs IIS'
Disable-ScheduledTask -TaskName 'Maintenance\Purge logs IIS'
Unregister-ScheduledTask -TaskName 'Maintenance\Purge logs IIS' -Confirm:$false
```

> Organisez vos tâches dans des **dossiers** (`\Maintenance\`, `\Rapports\`,
> `\Sauvegardes\`) : un planificateur avec 80 tâches à la racine est ingérable.

---

## 44. Déclencheurs avancés et conditions

Au-delà de l'horaire simple :

```powershell
# Déclencheur sur événement : à chaque échec du spouleur (ID 372, journal System)
$triggerEvt = New-ScheduledTaskTrigger -AtLogOn  # base, remplacée ci-dessous
# Les déclencheurs sur événement se définissent via CIM/XML :
$task = Get-ScheduledTask -TaskName 'Maintenance\Surveillance spouleur'
$task.Triggers[0].Subscription = @'
<QueryList><Query Id="0" Path="System">
<Select Path="System">*[System[Provider[@Name='Microsoft-Windows-PrintService'] and EventID=372]]</Select>
</Query></QueryList>
'@
Set-ScheduledTask -InputObject $task

# Conditions utiles (objet Settings) :
$settings = New-ScheduledTaskSettingsSet -StartWhenAvailable `
  -RunOnlyIfNetworkAvailable -AllowStartIfOnBatteries:$false `
  -DontStopIfGoingOnBatteries:$false -ExecutionTimeLimit '01:00:00'
```

Cas d'usage : redémarrer un service après un événement précis, lancer un script
d'inventaire à l'ouverture de session, purger uniquement si le réseau est
disponible (copie vers partage).

---

## 45. Comptes d'exécution : SYSTEM, comptes de service, gMSA

| Compte | Droits | Usage |
|---|---|---|
| **SYSTEM** | Locaux illimités, réseau = compte machine `DOMAINE\MACHINE$` | Tâches purement locales |
| **SERVICE LOCAL / SERVICE RÉSEAU** | Restreints | Services Windows, rarement les tâches |
| Compte de service dédié (utilisateur AD) | Ce que vous lui donnez | Tâches accédant à des ressources réseau spécifiques |
| **gMSA** | Géré par AD, mot de passe auto-roté | **Recommandé** pour tâches/services (§46) |
| Compte utilisateur « réel » | — | **À proscrire** (mot de passe qui expire = tâche en échec) |

**Règle** : une tâche qui échoue avec `0x41306` ou « le mot de passe a expiré »
signale presque toujours un compte utilisateur classique : migrez vers gMSA.

Droits minimaux : accordez au compte **exactement** ce qu'il faut (lecture sur le
partage, écriture dans le dossier de logs) via des groupes AD dédiés
(`GRP-Taches-Maintenance`), jamais « Administrateurs du domaine » par facilité.

---

## 46. gMSA : création et utilisation

Les **comptes de service gérés de groupe** (gMSA) : AD gère et fait tourner le mot
de passe automatiquement (tous les 30 jours). Idéal pour tâches planifiées et
services applicatifs.

```powershell
# --- Sur un contrôleur de domaine (une seule fois par forêt) ---
# 1. Créer la clé racine KDS (attendre ~10 h de réplication en prod ; -EffectiveTime immédiat en labo)
Add-KdsRootKey -EffectiveTime ((Get-Date).AddHours(-10))

# 2. Créer le gMSA (autoriser les serveurs à récupérer le mot de passe)
New-ADServiceAccount -Name 'gmsa-taches-maint' -DNSHostName 'gmsa-taches-maint.contoso.local' `
  -PrincipalsAllowedToRetrieveManagedPassword 'GRP-SRV-Maintenance'

# --- Sur chaque serveur autorisé ---
Install-ADServiceAccount -Identity 'gmsa-taches-maint'
Test-ADServiceAccount -Identity 'gmsa-taches-maint'   # doit retourner True

# Utilisation dans une tâche planifiée (le $ final est obligatoire, pas de mot de passe)
$principal = New-ScheduledTaskPrincipal -UserId 'contoso\gmsa-taches-maint$' -LogonType Password
Register-ScheduledTask -TaskName 'Maintenance\Rapport quotidien' -Action $action `
  -Trigger $trigger -Principal $principal -Settings $settings
```

> Le compte gMSA a besoin du droit **« Ouvrir une session en tant que tâche »**
> (SeBatchLogonRight) : accordé via la GPO « Ouvrir une session en tant que tâche ».
> Sans lui, la tâche échoue avec une erreur d'ouverture de session.

---

## 47. Audit et hygiène des tâches planifiées

Revue **hebdomadaire** (§7) : listez les tâches en échec et les tâches suspectes.

```powershell
# Tâches dont la dernière exécution a échoué (résultat != 0)
Get-ScheduledTask | Get-ScheduledTaskInfo |
  Where-Object { $_.LastTaskResult -ne 0 -and $_.LastRunTime -gt (Get-Date).AddDays(-7) } |
  Select-Object TaskName, TaskPath, LastRunTime,
    @{n='ResultHex'; e={ '0x{0:X}' -f $_.LastTaskResult }} |
  Format-Table -AutoSize

# Tâches exécutées sous un compte utilisateur (à migrer vers gMSA/SYSTEM)
Get-ScheduledTask | Where-Object {
    $_.Principal.UserId -notmatch '^(SYSTEM|.*\$|NT AUTHORITY\\|BUILTIN\\)'
} | Select-Object TaskName, TaskPath, @{n='Compte'; e={ $_.Principal.UserId } } |
  Format-Table -AutoSize
```

**Hygiène sécurité** : les tâches planifiées sont une technique de persistance
classique des attaquants. Toute tâche inconnue (nom bizarre, exécutable dans
`%TEMP%` ou `%APPDATA%`, compte SYSTEM + script obscur) = incident de sécurité
potentiel : isolez, ne supprimez pas avant analyse (gardez une copie du script).

---

## 48. Disques : gestion avec PowerShell et MMC

`diskmgmt.msc` pour le visuel, PowerShell pour l'industrialisation :

```powershell
# Vue d'ensemble : disques physiques, bus, santé
Get-PhysicalDisk | Select-Object FriendlyName, MediaType, BusType, HealthStatus,
  OperationalStatus, @{n='Go'; e={ [math]::Round($_.Size / 1GB) }} | Format-Table -AutoSize

# Volumes et espace libre
Get-Volume | Where-Object DriveType -eq 'Fixed' |
  Select-Object DriveLetter, FileSystemType, HealthStatus,
    @{n='TotalGo'; e={ [math]::Round($_.Size / 1GB, 1) }},
    @{n='LibreGo'; e={ [math]::Round($_.SizeRemaining / 1GB, 1) }},
    @{n='LibrePct'; e={ [math]::Round($_.SizeRemaining / $_.Size * 100, 1) }} |
  Format-Table -AutoSize

# Initialiser / partitionner / formater un nouveau disque (ATTENTION : destructif)
Initialize-Disk -Number 2 -PartitionStyle GPT
New-Partition -DiskNumber 2 -UseMaximumSize -DriveLetter E |
  Format-Volume -FileSystem NTFS -NewFileSystemLabel 'Donnees' -Confirm:$false
```

> ⚠️ `Initialize-Disk` / `Clear-Disk` **détruisent les données**. Vérifiez trois
> fois le numéro de disque (`Get-Disk | ft Number, FriendlyName, Size`) avant
> toute opération, surtout à distance.

---

## 49. Partitionnement : bonnes pratiques

