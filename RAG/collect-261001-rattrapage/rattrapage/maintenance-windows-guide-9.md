---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-9
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [1501, 1679]
sha256: 8a5e4d41dd544107bae3ecc381f6a3935f049edf187428d3cdd2611967d4d7d3
---

# Maintenance et exploitation Windows en entreprise

| Principe | Détail |
|---|---|
| GPT, pas MBR | MBR limité à 2 To ; GPT obligatoire en UEFI |
| Séparer OS et données | C: système, D: données/logs, volumes dédiés par rôle |
| Alignement | Laisser l'assistant Windows gérer (1 Mo par défaut, OK pour SSD/RAID) |
| Taille C: serveur | ≥ 100 Go (120-150 Go confortable avec WinSxS + logs) |
| Lettres stables | Figer les lettres via `Set-Partition -NewDriveLetter` ; documenter |
| ReFS vs NTFS | NTFS par défaut ; ReFS pour gros volumes de données/Sauvegardes (intégrité) |

Redimensionner (étendre un volume sur de l'espace adjacent libre) :

```powershell
# Étendre D: au maximum disponible
$part = Get-Partition -DriveLetter D
Resize-Partition -DriveLetter D -Size (Get-PartitionSupportedSize -DriveLetter D).SizeMax
```

> Pour **réduire** un volume système, l'espace doit être contigu en fin de volume ;
> les fichiers non déplaçables (pagefile, clichés) peuvent bloquer : désactivez-les
> temporairement si besoin, en fenêtre de maintenance.

---

## 50. Quotas NTFS

Deux mécanismes, ne pas confondre :

| Mécanisme | Portée | Outil |
|---|---|---|
| **Quotas NTFS** | Par volume, par utilisateur | `fsutil quota` / propriétés du volume |
| **FSRM** (File Server Resource Manager) | Par dossier, bien plus fin (seuils, notifications, blocage par type de fichier) | Rôle FSRM + `dirquota` |

Quota NTFS express (volume E:, 10 Go par utilisateur, avertissement à 8 Go) :

```cmd
:: Activer le suivi + appliquer les quotas sur E:
fsutil quota enforce E:
:: Définir un seuil par défaut : limite 10 Go, avertissement 8 Go (en octets)
fsutil quota modify E: 10737418240 8589934592 "contoso\utilisateur-fic"
```

**Recommandation** : sur un serveur de fichiers d'entreprise, utilisez **FSRM**
(quotas par dossier partagé, modèles, rapports de fichiers volumineux, File
Screening anti-ransomwares par extension). Le quota NTFS simple suffit pour des
besoins ponctuels (ex. limiter les profils sur un volume).

---

## 51. Optimisation : défragmentation HDD, TRIM SSD

| Média | Action | Fréquence |
|---|---|---|
| HDD | Défragmentation (`-Defrag`) | Hebdo (tâche intégrée `\Microsoft\Windows\Defrag\ScheduledDefrag`) |
| SSD | **ReTrim** (informe le contrôleur des blocs libres), jamais de défragmentation classique | Mensuelle (intégrée) |

```powershell
# État d'optimisation de tous les volumes
Get-Volume | Where-Object DriveType -eq 'Fixed' | ForEach-Object {
    Optimize-Volume -DriveLetter $_.DriveLetter -Analyze -Verbose
}

# Forcer : défragmentation HDD / ReTrim SSD (Windows choisit selon le média avec -ReTrim/-Defrag explicites)
Optimize-Volume -DriveLetter D -Defrag -Verbose      # HDD
Optimize-Volume -DriveLetter C -ReTrim -Verbose       # SSD

# Vérifier que la tâche planifiée intégrée est active
Get-ScheduledTask -TaskName 'ScheduledDefrag' -TaskPath '\Microsoft\Windows\Defrag\'
```

> Sur **VM**, la défragmentation du guest a peu d'intérêt si le stockage sous-jacent
> est déjà optimisé (SAN/SSD) : elle consomme des IOPS pour rien. Désactivez-la sur
> les VM à stockage flash et surveillez plutôt la latence (§57).

---

## 52. Santé des disques : SMART

Windows expose les compteurs de fiabilité via le sous-système de stockage :

```powershell
# Compteurs SMART/fiabilité : usure SSD, secteurs réalloués, température
Get-PhysicalDisk | Get-StorageReliabilityCounter |
  Select-Object DeviceId, Wear, TemperatureCelsius, ReadErrorsTotal,
    WriteErrorsTotal, PowerOnHours | Format-Table -AutoSize
# Wear = % d'usure (SSD) ; au-delà de 80-90 %, planifiez le remplacement
```

**Politique d'exploitation :**

- `Wear` > 80 % → remplacement planifié au prochain trimestre.
- Erreurs de lecture/écriture qui augmentent → remplacement prioritaire +
  vérification des sauvegardes.
- Température > 55 °C en continu → ventilation/positionnement à revoir.
- Complétez avec l'outil du constructeur (diagnostic long) avant de déclarer un
  disque « sain » après des erreurs.

> Le SMART n'est pas une boule de cristal : ~1/3 des pannes disque ne sont pas
> précédées d'alerte SMART. D'où l'importance des sauvegardes (§123) et du RAID /
> Storage Spaces (§53), jamais du « disque unique avec SMART au vert ».

---

## 53. Storage Spaces : concepts et administration

Les **espaces de stockage** mutualisent des disques physiques en pools, puis en
disques virtuels avec résilience (miroir, parité).

| Type de résilience | Disques min | Tolérance | Usage |
|---|---|---|---|
| Simple (sans résilience) | 1 | Aucune | Données temporaires |
| Miroir double | 2 | 1 disque | Usage général |
| Miroir triple | 5 | 2 disques | Données critiques |
| Parité simple | 3 | 1 disque | Archivage (écritures lentes) |
| Parité double | 7 | 2 disques | Archivage critique |

Administration :

```powershell
# État du pool et des disques virtuels
Get-StoragePool | Where-Object IsPrimordial -eq $false |
  Select-Object FriendlyName, HealthStatus, OperationalStatus,
    @{n='TotalTo'; e={ [math]::Round($_.Size / 1TB, 2) }} | Format-Table -AutoSize
Get-VirtualDisk | Select-Object FriendlyName, ResiliencySettingName, HealthStatus,
  OperationalStatus, @{n='To'; e={ [math]::Round($_.Size / 1TB, 2) }} | Format-Table -AutoSize

# Réparer un disque virtuel dégradé après remplacement du disque physique
Repair-VirtualDisk -FriendlyName 'DataMirror'

# Ajouter un disque au pool
Add-PhysicalDisk -StoragePoolFriendlyName 'PoolData' -PhysicalDisks (Get-PhysicalDisk -CanPool $true)
```

**En production** : surveillez `HealthStatus`/`OperationalStatus` (remontée
supervision §106) et testez la **reconstruction** avant d'en avoir besoin.
Pour du SDS sérieux en cluster, regardez **Storage Spaces Direct (S2D)** —
hors périmètre de ce guide, voir §138.

---

## 54. Performance : les outils (Gestionnaire des tâches, Resource Monitor, perfmon)

| Outil | Lancement | Usage |
|---|---|---|
| Gestionnaire des tâches | `taskmgr` / Ctrl+Maj+Échap | Vue instantanée, processus gourmands, démarrage |
| **Analyseur de ressources** | `resmon` | **Le meilleur premier réflexe** : CPU/disque/réseau/mémoire par processus, en temps réel |
| Analyseur de performances | `perfmon` | Compteurs précis, jeux de collecteurs, historique |
| Compteurs PowerShell | `Get-Counter` | Scriptable, à distance |

**Méthode express** (2 minutes) : `resmon` → onglet Disque : trier par « Total
(octets/s) » → le processus en tête est le coupable dans 80 % des « le serveur
est lent ». Onglet Réseau : idem pour la bande passante. Onglet UC : vérifier la
« Longueur de la file du processeur ».

---

## 55. Compteurs clés : CPU

| Compteur | Seuil d'alerte | Interprétation |
|---|---|---|
| `\Processor(_Total)\% Processor Time` | > 85 % sur 15 min | Saturation CPU |
| `\System\Processor Queue Length` | > 4 (par cœur : > 2×nb cœurs) | Vrai goulot CPU (plus fiable que le %) |
| `\Processor(_Total)\% Privileged Time` | > 30 % du total | Activité noyau/pilotes excessive |
| `\Process(*)\% Processor Time` | — | Identifier le processus fautif |

```powershell
# CPU total + file d'attente, 5 échantillons espacés de 10 s
Get-Counter '\Processor(_Total)\% Processor Time', '\System\Processor Queue Length' `
  -SampleInterval 10 -MaxSamples 5 |
  Select-Object -ExpandProperty CounterSamples |
  Select-Object Path, CookedValue, Timestamp | Format-Table -AutoSize
```

> Un CPU à 100 % avec une file d'attente faible = un processus mono-thread qui
> boucle (pas un manque de cœurs). Cherchez le processus, pas du CPU en plus.

---

## 56. Compteurs clés : mémoire

