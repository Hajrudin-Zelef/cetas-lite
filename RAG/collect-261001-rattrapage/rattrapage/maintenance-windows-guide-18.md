---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-18
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [3345, 3538]
sha256: 91459d5d63e55ca4c628dd6713abfc551be21d0ae708a88e6955edfc1c3e4031
---

# Maintenance et exploitation Windows en entreprise

```powershell
# Fiche matérielle complète d'une machine (locale ou via Invoke-Command)
$cs  = Get-CimInstance Win32_ComputerSystem
$os  = Get-CimInstance Win32_OperatingSystem
$bios = Get-CimInstance Win32_BIOS
$cpu = Get-CimInstance Win32_Processor | Select-Object -First 1
[pscustomobject]@{
    Machine      = $env:COMPUTERNAME
    Fabricant    = $cs.Manufacturer
    Modele       = $cs.Model
    Serie        = $bios.SerialNumber
    CPU          = $cpu.Name
    Coeurs       = $cpu.NumberOfCores
    RAM_Go       = [math]::Round($cs.TotalPhysicalMemory / 1GB, 1)
    OS           = $os.Caption
    Build        = "$($os.Version) ($($os.BuildNumber))"
    DernierBoot  = $os.LastBootUpTime
    Disques      = (Get-PhysicalDisk |
                    ForEach-Object { "$($_.FriendlyName) $([math]::Round($_.Size/1GB)) Go $($_.MediaType)" }) -join ' | '
} | Format-List
```

Industrialisation : bouclez sur la liste des serveurs avec `Invoke-Command`
(§113), `Export-Csv -Append` vers un fichier central, planifiez mensuellement
(§121). C'est votre CMDB « v1 » en attendant mieux.

---

## 116. Inventaire logiciels installés

Deux sources complémentaires (aucune n'est exhaustive seule) :

```powershell
# 1. Registre Uninstall (rapide, couvre 95 % des logiciels classiques)
$paths = @(
  'HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\*',
  'HKLM:\SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall\*'
)
Get-ItemProperty $paths -ErrorAction SilentlyContinue |
  Where-Object DisplayName |
  Select-Object DisplayName, DisplayVersion, Publisher, InstallDate |
  Sort-Object DisplayName | Format-Table -AutoSize

# 2. Win32_Product : À ÉVITER (déclenche une reconfiguration MSI à chaque appel !)
# Préférez le registre ci-dessus, ou Get-Package :
Get-Package | Select-Object Name, Version, ProviderName | Sort-Object Name |
  Format-Table -AutoSize
```

> ⚠️ `Get-CimInstance Win32_Product` / `Get-WmiObject Win32_Product` provoque une
> **vérification de chaque MSI** (lent, peut casser des applis). Ne l'utilisez
> jamais en inventaire. C'est l'erreur classique n°1 des scripts trouvés sur le web.

---

## 117. Inventaire des correctifs installés

```powershell
# Correctifs (KB) installés, du plus récent au plus ancien
Get-HotFix | Sort-Object InstalledOn -Descending |
  Select-Object HotFixID, Description, InstalledOn, InstalledBy |
  Format-Table -AutoSize

# Conformité : telle machine a-t-elle le cumulatif du mois ? (ex. KB5034765)
$kb = 'KB5034765'
if (Get-HotFix -Id $kb -ErrorAction SilentlyContinue) {
    "OK : $kb installé"
} else {
    "MANQUANT : $kb"
}

# À l'échelle du parc (via PSRemoting)
Invoke-Command -ComputerName (Get-Content C:\Scripts\serveurs.txt) {
    [pscustomobject]@{
        Machine = $env:COMPUTERNAME
        DernierKB = (Get-HotFix | Sort-Object InstalledOn -Descending |
                     Select-Object -First 1).HotFixID
        DateKB  = (Get-HotFix | Sort-Object InstalledOn -Descending |
                     Select-Object -First 1).InstalledOn
    }
} | Format-Table -AutoSize
```

Croisez avec la date du Patch Tuesday (§10) : un serveur dont le dernier KB date
de 3 mois = non-conforme → à intégrer au rattrapage (§8).

---

## 118. Rapport espace disque multi-serveurs

Le rapport quotidien du §5, version complète :

```powershell
# Rapport-EspaceDisque.ps1 — à planifier chaque matin (§121)
$servers = Get-Content C:\Scripts\serveurs.txt
$seuilPct = 20
$result = Invoke-Command -ComputerName $servers {
    Get-CimInstance Win32_LogicalDisk -Filter "DriveType=3" | ForEach-Object {
        [pscustomobject]@{
            Machine   = $env:COMPUTERNAME
            Volume    = $_.DeviceID
            TotalGo   = [math]::Round($_.Size / 1GB, 1)
            LibreGo   = [math]::Round($_.FreeSpace / 1GB, 1)
            LibrePct  = [math]::Round($_.FreeSpace / $_.Size * 100, 1)
        }
    }
}
$alertes = $result | Where-Object LibrePct -lt $seuilPct
$date = Get-Date -Format 'yyyy-MM-dd'
$result | Export-Csv "C:\Rapports\disques-$date.csv" -NoTypeInformation -Encoding UTF8
if ($alertes) {
    $corps = $alertes | Format-Table -AutoSize | Out-String
    Send-MailMessage -To 'exploitation@contoso.local' -From 'rapports@contoso.local' `
      -Subject "[ALERTE] Espace disque < $seuilPct % ($date)" -Body $corps `
      -SmtpServer 'smtp.contoso.local' -Encoding UTF8
}
```

> Adaptez `Send-MailMessage` à votre relais (authentification, TLS). Sur les
> environnements récents, préférez un connecteur ou un script d'envoi via votre
> solution de supervision plutôt que SMTP anonyme.

---

## 119. Rapport comptes AD expirés / inactifs

```powershell
# Prérequis : module ActiveDirectory (RSAT) sur la machine d'exécution
Import-Module ActiveDirectory

# Comptes inactifs depuis plus de 90 jours (hors comptes de service connus)
$limite = (Get-Date).AddDays(-90)
Search-ADAccount -UsersOnly -AccountInactive -TimeSpan 90 |
  Where-Object { $_.Enabled -and $_.SamAccountName -notlike 'svc-*' } |
  Select-Object Name, SamAccountName, LastLogonDate, DistinguishedName |
  Export-Csv C:\Rapports\comptes-inactifs.csv -NoTypeInformation -Encoding UTF8

# Comptes expirés ou expirant dans moins de 30 jours
Get-ADUser -Filter 'Enabled -eq $true' -Properties AccountExpirationDate, PasswordLastSet |
  Where-Object { $_.AccountExpirationDate -and $_.AccountExpirationDate -lt (Get-Date).AddDays(30) } |
  Select-Object Name, SamAccountName, AccountExpirationDate |
  Sort-Object AccountExpirationDate | Format-Table -AutoSize

# Comptes verrouillés actuellement
Search-ADAccount -LockedOut | Select-Object Name, SamAccountName, LastLogonDate |
  Format-Table -AutoSize
```

**Processus** (mensuel, §8) : le rapport part au chef de service → validation
métier → **désactivation** (pas suppression immédiate : 30 jours de quarantaine
en OU dédiée avant suppression, au cas où).

---

## 120. Rapport certificats qui expirent

```powershell
# Certificats machine expirant dans moins de 60 jours (magasin LocalMachine\My)
$seuil = (Get-Date).AddDays(60)
Get-ChildItem Cert:\LocalMachine\My |
  Where-Object { $_.NotAfter -lt $seuil } |
  Select-Object Subject,
    @{n='ExpireLe'; e={ $_.NotAfter }},
    @{n='JoursRestants'; e={ [math]::Round(($_.NotAfter - (Get-Date)).TotalDays) }},
    Thumbprint |
  Sort-Object ExpireLe | Format-Table -AutoSize

# Multi-serveurs : à exécuter via Invoke-Command sur les serveurs à certificats
# (RDS, IIS, ADFS, VPN...) listés dans C:\Scripts\serveurs-certificats.txt
```

**Politique** : alerte à 60 j (warning), 14 j (critique) — §107. Documentez pour
chaque certificat : usage, autorité, procédure de renouvellement, responsable.
Un certificat expiré un dimanche à 3h du matin sur le VPN d'astreinte, c'est du
vécu — la liste à jour évite ça.

---

## 121. Exécution planifiée des scripts de rapport

Assemblez les rapports §115-120 en une tâche quotidienne :

```powershell
# Tâche "Rapports quotidiens" : 07:30, en gMSA, avec limite de durée
$action = New-ScheduledTaskAction -Execute 'powershell.exe' `
  -Argument '-NoProfile -ExecutionPolicy Bypass -File "C:\Scripts\Rapports-Quotidiens.ps1"'
$trigger = New-ScheduledTaskTrigger -Daily -At '07:30'
$principal = New-ScheduledTaskPrincipal -UserId 'contoso\gmsa-taches-maint$' -LogonType Password
$settings = New-ScheduledTaskSettingsSet -ExecutionTimeLimit '01:00:00' -StartWhenAvailable
Register-ScheduledTask -TaskName 'Rapports\Quotidiens' -Action $action `
  -Trigger $trigger -Principal $principal -Settings $settings `
  -Description 'Inventaire + espace disque + correctifs + certificats (guide §115-120)'
```

