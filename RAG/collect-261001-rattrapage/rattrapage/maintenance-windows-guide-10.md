---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-10
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "full-duplex", "memory"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [1680, 1865]
sha256: f3ac593a9eb3af98c203ba6774c0516eb7c515d58622bef5ec8d38095251a9fc
---

# Maintenance et exploitation Windows en entreprise

| Compteur | Seuil | Sens |
|---|---|---|
| `\Memory\Available MBytes` | < 10 % de la RAM | Mémoire sous pression |
| `\Memory\% Committed Bytes In Use` | > 85 % | Engagement excessif (RAM + pagefile) |
| `\Memory\Pages/sec` | > 1000 soutenu | Pagination dure (disque lent = catastrophe) |
| `\Memory\Pool Nonpaged Bytes` | Croissance continue | **Fuite mémoire noyau/pilote** |
| `\Process(*)\Working Set` | — | Consommation par processus |

```powershell
# Mémoire disponible + % engagé
$os = Get-CimInstance Win32_OperatingSystem
[pscustomobject]@{
    TotalGo      = [math]::Round($os.TotalVisibleMemorySize / 1MB, 1)
    LibreGo      = [math]::Round($os.FreePhysicalMemory / 1MB, 1)
    LibrePct     = [math]::Round($os.FreePhysicalMemory / $os.TotalVisibleMemorySize * 100, 1)
    PageFileUtil = (Get-Counter '\Paging File(_Total)\% Usage').CounterSamples.CookedValue
} | Format-List
```

**Fuite mémoire** : si le pool non paginé croît sans jamais redescendre après
arrêt des applications, suspectez un pilote → `poolmon.exe` (WDK) puis Driver
Verifier (§74). Sur un serveur, une fuite = redémarrage planifié en attendant le
correctif, pas un « on verra ».

---

## 57. Compteurs clés : disque

| Compteur | Seuil | Sens |
|---|---|---|
| `\PhysicalDisk(*)\Avg. Disk sec/Read` et `/Write` | > 20 ms soutenu (HDD), > 5 ms (SSD) | **Latence** : le vrai indicateur |
| `\PhysicalDisk(*)\Disk Reads/sec`, `Disk Writes/sec` | — | IOPS |
| `\PhysicalDisk(*)\% Disk Time` | > 90 % | Disque saturé |
| `\PhysicalDisk(*)\Current Disk Queue Length` | > 2 par disque | File d'attente |

```powershell
# Latence moyenne lecture/écriture par disque, 3 échantillons de 15 s
Get-Counter '\PhysicalDisk(*)\Avg. Disk sec/Read', '\PhysicalDisk(*)\Avg. Disk sec/Write' `
  -SampleInterval 15 -MaxSamples 3 |
  Select-Object -ExpandProperty CounterSamples |
  Where-Object InstanceName -ne '_Total' |
  Select-Object InstanceName, Path,
    @{n='LatenceMs'; e={ [math]::Round($_.CookedValue * 1000, 2) }}, Timestamp |
  Format-Table -AutoSize
```

> Sur VM : la latence vue par l'invité inclut l'hôte et la baie. Si l'invité est
> « lent disque » mais ses compteurs sont bons, regardez l'hôte/Hyper-V et le
> stockage partagé avant d'ajouter des vDisques.

---

## 58. Compteurs clés : réseau

| Compteur | Seuil | Sens |
|---|---|---|
| `\Network Interface(*)\Bytes Total/sec` | > 80 % de la capacité du lien | Saturation |
| `\Network Interface(*)\Output Queue Length` | > 2 | Congestion sortante |
| `\TCPv4\Segments Retransmitted/sec` | > 5 % des segments | Perte / qualité de lien |
| `\Network Interface(*)\Packets Received Errors` | > 0 croissant | Problème physique (câble, duplex) |

```powershell
# Débit total par interface (octets/s), 5 échantillons de 5 s
Get-Counter '\Network Interface(*)\Bytes Total/sec' -SampleInterval 5 -MaxSamples 5 |
  Select-Object -ExpandProperty CounterSamples |
  Where-Object { $_.InstanceName -notmatch 'Loopback|isatap|Teredo' } |
  Select-Object InstanceName,
    @{n='Mo/s'; e={ [math]::Round($_.CookedValue / 1MB, 2) }}, Timestamp |
  Format-Table -AutoSize
```

> Erreurs/duplex : un lien négocié en 100 Mb/s half-duplex au lieu de 1 Gb/s
> full-duplex explique bien des « lenteurs réseau » (§97, cas 12).

---

## 59. Analyseur de performances : jeux de collecteurs

Pour un problème intermittent (« lent tous les jours à 14h »), collectez au lieu
de regarder en direct : **Jeux de collecteurs de données** > **Défini par
l'utilisateur** > Nouveau > « Créer manuellement » > Compteur de performances.

Jeu « Diagnostic standard » (modèle à créer une fois, réutiliser partout) :

- `\Processor(_Total)\% Processor Time`, `\System\Processor Queue Length`
- `\Memory\Available MBytes`, `\Memory\Pages/sec`, `\Memory\Pool Nonpaged Bytes`
- `\PhysicalDisk(*)\Avg. Disk sec/Read`, `\PhysicalDisk(*)\Avg. Disk sec/Write`
- `\Network Interface(*)\Bytes Total/sec`
- Intervalle : 15 s. Durée : 24-72 h. Format : binaire circulaire (limite 500 Mo).

Planifiez le démarrage/arrêt (ou laissez tourner en circulaire) puis analysez dans
perfmon : ajoutez les compteurs, repérez les pics corrélés à l'heure du symptôme.
**Corrélation > valeur absolue** : un pic disque à 14h02 tous les jours = un batch,
pas un disque mourant.

---

## 60. PAL : Performance Analysis of Logs

**PAL (Performance Analysis of Logs)** est un outil gratuit (PowerShell + seuils
documentés) qui analyse un journal de compteurs perfmon et génère un rapport HTML
avec alertes (« seuil dépassé ») : idéal pour objectiver un problème de perfs
devant un éditeur ou la direction.

Usage type :

1. Collectez avec le jeu §59 (format `.blg`).
2. Passez le `.blg` dans PAL avec le seuil correspondant au rôle
   (« File Server », « SQL Server », « Terminal Services »…).
3. Le rapport signale les compteurs hors seuils avec explications.

> Mention : PAL n'est plus maintenu activement mais reste fonctionnel et ses
> **seuils** demeurent une excellente référence documentaire, même lus à la main.

---

## 61. Scripts de collecte de performances

Collecte légère et scriptable sans perfmon (à pousser via PSRemoting, §113) :

```powershell
# Snapshot de santé perf : CPU, mémoire, disque, réseau — une ligne par serveur
$servers = Get-Content C:\Scripts\serveurs.txt
Invoke-Command -ComputerName $servers {
    $cpu  = (Get-Counter '\Processor(_Total)\% Processor Time' -SampleInterval 5 -MaxSamples 3 |
             Select-Object -ExpandProperty CounterSamples | Measure-Object CookedValue -Average).Average
    $os   = Get-CimInstance Win32_OperatingSystem
    $disk = Get-Counter '\PhysicalDisk(_Total)\Avg. Disk sec/Read' -ErrorAction SilentlyContinue |
             Select-Object -ExpandProperty CounterSamples
    [pscustomobject]@{
        Machine   = $env:COMPUTERNAME
        CPU_Pct   = [math]::Round($cpu, 1)
        MemLibrePct = [math]::Round($os.FreePhysicalMemory / $os.TotalVisibleMemorySize * 100, 1)
        LatDiskMs = if ($disk) { [math]::Round($disk.CookedValue * 1000, 1) } else { 'n/a' }
    }
} | Sort-Object CPU_Pct -Descending | Format-Table -AutoSize
```

Historisation : écrivez ce snapshot toutes les 15 min dans un CSV central
(`Export-Csv -Append`) via une tâche planifiée (§43) — vous obtenez un mini
perfmon sans infrastructure, exploitable sous Excel ou Grafana.

---

## 62. WinRE : accéder à l'environnement de récupération

**WinRE** (Windows Recovery Environment) = le mode de réparation hors ligne
(dépannage, invite de commandes, restauration système, réparation du démarrage).

Accès :

| Méthode | Détail |
|---|---|
| Paramètres > Récupération > Démarrage avancé > Redémarrer | Windows 11 |
| `shutdown /r /o /f /t 0` | Redémarre directement sur les options avancées |
| Touche Maj + Redémarrer (menu Démarrer) | Le plus rapide |
| 2-3 démarrages ratés consécutifs | Bascule automatique (ne pas en abuser) |
| Support d'installation > Réparer l'ordinateur | Quand WinRE local est absent/corrompu |

Vérifier l'état de WinRE (doit être **Enabled**) :

```cmd
reagentc /info
```

Si « Disabled » ou image introuvable : `reagentc /enable` (nécessite la partition
de récupération et `winre.wim` présents).

---

## 63. WinRE : activer, désactiver, personnaliser

```cmd
:: État détaillé
reagentc /info

:: Activer / désactiver
reagentc /enable
reagentc /disable

:: Redéfinir l'emplacement de winre.wim (ex. après recréation de la partition)
reagentc /setreimage /path R:\Recovery\WindowsRE
```

**Bonnes pratiques :**

