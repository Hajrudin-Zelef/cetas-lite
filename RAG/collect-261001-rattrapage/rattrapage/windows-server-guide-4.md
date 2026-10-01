---
id: collect-261001-rattrapage/rattrapage/windows-server-guide-4
title: "Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/windows_server_guide.md
source_anchor: ""
source_lines: [565, 744]
sha256: f998d337b9c13fd9e6bc04209376f8c8cdbf293b9e2d2a99dbe8436e87924b8d
---

# Windows Server en entreprise — Guide technique ultra-complet

```powershell
# Activer la mémoire dynamique (min 512 Mo, max 8 Go, démarrage 2 Go)
Set-VMMemory -VMName "SRV-WEB-01" -DynamicMemoryEnabled $true `
  -MinimumBytes 512MB -StartupBytes 2GB -MaximumBytes 8GB -Priority 50

# Mémoire statique (préférable pour SQL, AD, charges critiques)
Set-VMMemory -VMName "SRV-SQL-01" -DynamicMemoryEnabled $false -StartupBytes 16GB
```

| Charge | Mémoire dynamique ? |
|---|---|
| Web, fichiers, utilitaires | Oui |
| AD DS, DNS | Plutôt statique (petites VM, coût nul) |
| SQL Server, Exchange | **Non** (statique + verrouillée) |
| RDS Session Host | Statique |

**Disques :**

| Type VHDX | Avantage | Inconvénient |
|---|---|---|
| Dynamique (expansible) | Gain de place | Fragmentation, légère perte perf |
| Taille fixe | Meilleures perfs | Espace réservé |
| Différenciant | Lab, VDI | Chaîne parent/enfant fragile |

```powershell
# Convertir un disque dynamique en fixe (VM éteinte)
Convert-VHD -Path "D:\VMs\SRV-SQL-01\disque.vhdx" -DestinationPath "D:\VMs\SRV-SQL-01\disque-fixe.vhdx" -VHDType Fixed

# Étendre un disque (VM allumée possible si pas de snapshot)
Resize-VHD -Path "D:\VMs\SRV-FICHIER-01\data.vhdx" -SizeBytes 500GB
# Puis étendre la partition DANS la VM : Gestion des disques ou
# Resize-Partition -DriveLetter D -Size (Get-PartitionSupportedSize -DriveLetter D).SizeMax
```

---

## 20. Checkpoints : standard vs production

| | Checkpoint standard | Checkpoint de production |
|---|---|---|
| Technique | État mémoire complet | VSS dans la VM (cohérent applicatif) |
| Usage | Lab, tests | **Production** (sauvegarde applicative) |
| Restauration | Retour état exact | Restauration "au redémarrage" |

```powershell
# Passer en checkpoints de production (défaut recommandé en prod)
Set-VM -Name "SRV-SQL-01" -CheckpointType Production
Set-VM -Name "SRV-SQL-01" -AutomaticCheckpointsEnabled $false

# Créer / lister / restaurer / supprimer
Checkpoint-VM -Name "SRV-WEB-01" -SnapshotName "Avant-MAJ-2026-09"
Get-VMSnapshot -VMName "SRV-WEB-01"
Restore-VMSnapshot -Name "Avant-MAJ-2026-09" -VMName "SRV-WEB-01" -Confirm:$false
Remove-VMSnapshot -VMName "SRV-WEB-01" -Name "Avant-MAJ-2026-09"
```

> **Un checkpoint n'est PAS une sauvegarde.** Chaîne AVHDX fragile, perte de perfs, risque de saturation disque. Durée de vie max conseillée : **quelques jours**, jamais des mois. En production, préférer la vraie sauvegarde (§52).

---

## 21. Réplication Hyper-V (Hyper-V Replica)

Réplication asynchrone d'une VM vers un second hôte (même site ou distant), sans stockage partagé. RPO 5 minutes (30 s / 15 min configurables).

```powershell
# --- Sur le site de REPLICA (destination) ---
# Autoriser la réplication entrante (Kerberos = domaine, Certificat = hors domaine)
Set-VMReplicationServer -ReplicationEnabled $true -AllowedAuthenticationType Kerberos `
  -KerberosPort 8080 -DefaultStorageLocation "E:\Replica"

# Ouvrir le pare-feu
Enable-NetFirewallRule -DisplayName "Réplication Hyper-V (TCP 8080 entrant)"

# --- Sur le site PRINCIPAL (source) ---
# Activer la réplication pour une VM
Enable-VMReplication -VMName "SRV-FICHIER-01" -ReplicaServerName "SRV-HV-DR-01" `
  -ReplicaServerPort 8080 -AuthenticationType Kerberos `
  -ReplicaVhdLocation "E:\Replica" -VSSSnapshotFrequencyHour 4

# Démarrer la réplication initiale (réseau ou support externe)
Start-VMInitialReplication -VMName "SRV-FICHIER-01"

# Suivi
Get-VMReplication -VMName "SRV-FICHIER-01" | Format-List *
Measure-VMReplication -VMName "SRV-FICHIER-01"   # statistiques (nécessite le compteur activé)

# Basculement planifié (maintenance) / non planifié (sinistre)
Start-VMFailover -VMName "SRV-FICHIER-01"                    # planifié : sur la source
Start-VMFailover -VMName "SRV-FICHIER-01" -AsTest            # test sans impacter la prod
```

**Limites :** pas de haute disponibilité automatique (bascule manuelle), RPO ≥ 30 s, ne remplace pas un cluster pour la continuité locale.

---

## 22. Cluster de basculement + CSV

Architecture : 2+ nœuds Hyper-V + stockage partagé (SAN, Storage Spaces Direct, ou SMB 3.0 Scale-Out) + réseau dédié.

```powershell
# 1. Installer le rôle sur chaque nœud
Install-WindowsFeature -Name Failover-Clustering -IncludeManagementTools

# 2. Valider la configuration (OBLIGATOIRE avant création)
Test-Cluster -Node SRV-HV-01,SRV-HV-02 -Include "Inventory","Network","Storage","System Configuration"
# Lire le rapport : C:\Windows\Cluster\Reports\Validation Report.html

# 3. Créer le cluster
New-Cluster -Name CLUSTER-HV -Node SRV-HV-01,SRV-HV-02 -StaticAddress 192.168.10.50

# 4. Ajouter le stockage et le passer en CSV (Cluster Shared Volume)
Get-ClusterAvailableDisk | Add-ClusterDisk
Add-ClusterSharedVolume -Name "Cluster Disk 1"   # devient C:\ClusterStorage\Volume1

# 5. Rendre des VM hautement disponibles
Add-ClusterVirtualMachineRole -VMName "SRV-AD-01"
Add-ClusterVirtualMachineRole -VMName "SRV-FICHIER-01"

# 6. Quorum : témoin de partage de fichiers (recommandé pour 2 nœuds)
Set-ClusterQuorum -FileShareWitness "\\SRV-FICHIER-01\Quorum"
```

**Réseaux du cluster :** séparer (1) management, (2) trafic VM, (3) cluster/heartbeat, (4) CSV/redirection. Le heartbeat sur le même lien que les VM = risque de bascules intempestives.

```powershell
# Vérifier l'état
Get-ClusterNode | Format-Table Name, State -AutoSize
Get-ClusterResource | Where-Object ResourceType -eq "Virtual Machine" | Format-Table Name, State, OwnerNode
```

---

## 23. Bonnes pratiques Hyper-V (20 règles d'or)

1. Hôtes en **Server Core**, jamais de Desktop Experience sur un hôte.
2. **Aucun autre rôle** sur l'hôte (pas d'AD, pas de fichiers, pas d'IIS).
3. Volume dédié pour les VM, jamais sur C:.
4. Antivirus : **exclure** les dossiers VM, extensions `.vhd/.vhdx/.avhd/.avhdx`, processus `vmms.exe`/`vmwp.exe`.
5. Checkpoints : durée de vie < 7 jours, jamais comme sauvegarde.
6. Mémoire dynamique sauf charges critiques (SQL).
7. vCPU : commencer petit (2-4), ratio max ~8:1 vCPU/pCPU en prod généraliste.
8. **Secure Boot** activé sur les Gen2 (sauf incompatibilité prouvée).
9. Services d'intégration à jour (via Windows Update dans la VM).
10. Arrêt : configurer l'**action d'arrêt automatique** (Enregistrer l'état ou Arrêter).
11. Réseau : séparer management / VM / cluster / CSV.
12. Teaming (SET) sur 2022/2025 : `New-VMSwitch -Name vSw -NetAdapterName "Team1" -EnableEmbeddedTeaming $true`.
13. Jumbo frames uniquement si toute la chaîne les supporte (sinon perfs dégradées).
14. Réplication Hyper-V ≠ cluster : RPO 5 min, bascule manuelle.
15. Ne pas faire de snapshots sur les **contrôleurs de domaine virtualisés** (risque USN rollback — préférer la réplication AD native).
16. Time sync : désactiver l'intégration "Synchronisation de l'heure" sur les DC (le PDC fait foi via NTP).
17. Dédup : **jamais** sur un CSV ni sur des VHDX de VM en cours d'exécution.
18. Sauvegarder l'hôte (config) ET les VM (données).
19. Documenter : nom VM = rôle, notes Hyper-V remplies (`Set-VM -Notes "..."`).
20. Tester la restauration. Une sauvegarde non testée = pas de sauvegarde.

---

## 24. WSUS : architecture et prérequis

WSUS = référentiel local des mises à jour Microsoft. Les clients se synchronisent dessus au lieu d'aller sur Internet.

**Prérequis :**
- IIS (installé automatiquement avec le rôle)
- Base : **WID** (Windows Internal Database, incluse) ou SQL Server
- Stockage : 100 Go minimum (métadonnées + contenu selon produits/classifications)
- RAM : 4 Go minimum, 8 Go conseillés (le pool IIS WSUS consomme)

```powershell
# Installer WSUS (base WID)
Install-WindowsFeature -Name UpdateServices -IncludeManagementTools
# Puis lancer la configuration post-installation :
& 'C:\Program Files\Update Services\Tools\wsusutil.exe' postinstall CONTENT_DIR=D:\WSUS
```

---

## 25. Installation et configuration initiale de WSUS

Après `wsusutil postinstall` :

