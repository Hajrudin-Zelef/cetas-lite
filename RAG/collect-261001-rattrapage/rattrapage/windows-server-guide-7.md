---
id: collect-261001-rattrapage/rattrapage/windows-server-guide-7
title: "Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/windows_server_guide.md
source_anchor: ""
source_lines: [1093, 1253]
sha256: 57244949a13a83b92e71eddd6faf4ee609f41c66eb93ecbcf7c035da2bf1a77c
---

# Exclure des extensions (bases, chiffrés)
Set-DedupVolume -Volume "D:" -ExcludeFileType @("edb","vhdx","bak")
```

| À faire | À ne pas faire |
|---|---|
| Volumes de fichiers partagés, VDI | Volume système (C:) |
| Sauvegardes sur disque (avec prudence) | CSV de cluster Hyper-V |
| | VHDX de VM en cours d'exécution |

---

## 37. Dépannage d'accès aux fichiers (méthode en 7 étapes)

Face à "Accès refusé" sur `\\srv\partage` :

```powershell
# 1. Le partage existe-t-il ? Le serveur répond-il ?
Test-NetConnection SRV-FICHIER-01 -Port 445
Get-SmbShare -CimSession SRV-FICHIER-01 | Where-Object Name -eq "Compta"

# 2. Permissions de PARTAGE
Get-SmbShareAccess -Name "Compta" -CimSession SRV-FICHIER-01

# 3. Permissions NTFS + héritage
(Get-Acl "\\SRV-FICHIER-01\Compta").Access | Format-Table IdentityReference, FileSystemRights, IsInherited

# 4. L'utilisateur est-il bien dans le groupe ? (jeton d'accès)
whoami /groups                          # sur le poste de l'utilisateur
# Attention : appartenance ajoutée APRÈS l'ouverture de session = déconnexion/reconnexion requise !

# 5. Verrous de fichiers ouverts
Get-SmbOpenFile -CimSession SRV-FICHIER-01 | Where-Object Path -like "*Compta*"

# 6. Journal côté serveur (échecs d'audit si activés)
Get-WinEvent -FilterHashtable @{LogName='Security'; ID=4656,4658} -MaxEvents 20

# 7. Accès effectif (onglet Sécurité → Avancé → Accès effectif, ou icacls)
icacls "D:\Partages\Compta"
```

**Top 3 des causes réelles :** (1) groupe ajouté mais session non rouverte, (2) héritage NTFS coupé par un sous-dossier, (3) partage restreint + NTFS large (ou l'inverse).

---

---

## 38. Serveur d'impression : rôle et architecture

Le serveur d'impression centralise : files d'attente, pilotes, déploiement GPO, supervision. Un seul point d'administration au lieu de 40 imprimantes installées en local sur les postes.

```powershell
# Installer le rôle
Install-WindowsFeature -Name Print-Services -IncludeManagementTools
# Services de rôle utiles :
Install-WindowsFeature -Name Print-Server, Print-Management   # Print-Management = console dédiée
```

**Architecture :** poste client → file d'attente sur le serveur (`\\SRV-PRINT-01\Copieur-Compta`) → spouleur serveur → imprimante/copieur réseau (port TCP/IP Standard, port 9100 ou LPR).

**LIEN MÉTIER COPIEURS (Zelef) :** le serveur d'impression est l'interface entre ton parc copieurs (Kyocera, etc.) et les utilisateurs. Bonnes pratiques spécifiques copieurs :
- 1 file d'attente par **fonction** (ex. `Copieur-Accueil-NB`, `Copieur-Accueil-Couleur`) plutôt que par machine, pour mutualiser.
- Pilote **PCL** pour la bureautique courante, **PostScript** si PAO/impression graphique, **KX Driver** (Kyocera) en version validée — jamais la dernière version le jour de sa sortie.
- Port TCP/IP Standard avec **SNMP activé** (communauté `public` en lecture) pour remonter les états (toner bas, bourrage) — croiser avec la supervision (voir guide onduleurs/supervision, Zabbix).

---

## 39. Déploiement d'imprimantes et copieurs via GPO

```powershell
# Publier une imprimante dans l'annuaire (pour recherche)
Set-Printer -Name "Copieur-Compta" -Published $true

# Lister les imprimantes du serveur
Get-Printer -ComputerName SRV-PRINT-01 | Format-Table Name, DriverName, PortName, Published -AutoSize
```

**Méthode GPO (recommandée) :**

1. Console **Gestion de l'impression** → clic droit sur l'imprimante → **Déployer avec la stratégie de groupe**.
2. Choisir : **par utilisateur** (suit l'utilisateur) ou **par ordinateur** (suit le poste — idéal pour un copieur d'étage).
3. Lier la GPO à l'OU des utilisateurs/postes concernés.
4. Au prochain `gpupdate`, l'imprimante apparaît.

**En PowerShell (alternative, script d'ouverture de session) :**

```powershell
# Ajouter une imprimante réseau partagée sur un poste (script GPO)
Add-Printer -ConnectionName "\\SRV-PRINT-01\Copieur-Accueil"
# Définir par défaut
(Get-WmiObject -Class Win32_Printer -Filter "Name='\\\\SRV-PRINT-01\\Copieur-Accueil'").SetDefaultPrinter()
```

> Depuis les correctifs **PrintNightmare** (2021), le déploiement de pilotes Type 3 exige des droits administrateur côté client sauf configuration spécifique (Point and Print Restrictions). Préférer les pilotes **Type 4** quand disponibles, ou signer/valider le pipeline (voir §40).

---

## 40. Pilotes Type 3 vs Type 4, Print Management

| | Type 3 (v3) | Type 4 (v4) |
|---|---|---|
| Modèle | Pilote fabricant complet | Basé sur la classe (inbox) + extensions |
| Installation client | Télécharge le pilote depuis le serveur | Utilise le pilote local du poste |
| PrintNightmare | Surface d'attaque (exécution côté serveur) | Réduite |
| Fonctionnalités avancées | Complètes (finisher, codes...) | Parfois limitées |
| Recommandation | Si besoin métier (copieurs avec codes départementaux) | Par défaut si possible |

```powershell
# Voir les pilotes installés et leur type
Get-PrinterDriver -ComputerName SRV-PRINT-01 | Format-Table Name, PrinterEnvironment, MajorVersion -AutoSize
# MajorVersion 3 = Type 3, 4 = Type 4

# Isoler un pilote problématique (mode isolé)
Set-PrinterDriver -Name "Kyocera ECOSYS MXXXX" -PrinterEnvironment "Windows x64" # (via registre/GUI pour l'isolation)
# En pratique : Gestion de l'impression → Pilotes → Propriétés → Isolation : Isolé/Partagé
```

**Règle copieurs :** les fonctions avancées (codes comptables, boîtes personnelles, finisher agrafage) exigent souvent le pilote **Type 3 du fabricant**. Dans ce cas : serveur d'impression dédié, pilotes validés, **jamais** de pilote téléchargé au hasard.

---

## 41. Files d'attente, spooler : redémarrage et migration

```powershell
# Redémarrer le spouleur (LE réflexe quand tout est bloqué)
Restart-Service -Name Spooler -Force

# Purger une file bloquée (document corrompu)
Get-PrintJob -PrinterName "Copieur-Compta" -ComputerName SRV-PRINT-01
Remove-PrintJob -PrinterName "Copieur-Compta" -ID 42 -ComputerName SRV-PRINT-01
# Purge totale d'urgence :
Stop-Service Spooler -Force
Remove-Item "C:\Windows\System32\spool\PRINTERS\*" -Force
Start-Service Spooler

# Lister les travaux en erreur
Get-PrintJob -ComputerName SRV-PRINT-01 | Where-Object JobStatus -like "*Error*"

# MIGRATION vers un nouveau serveur (l'outil officiel)
# Sur l'ancien :
Printbrm -s \\SRV-PRINT-OLD -b -f D:\Migration\printers.printerExport
# Sur le nouveau :
Printbrm -s \\SRV-PRINT-NEW -r -f D:\Migration\printers.printerExport
```

> `Printbrm` migre files, pilotes, ports, formulaires. Tester sur 2 imprimantes avant de tout basculer. Prévoir le changement de nom dans les GPO de déploiement.

---

## 42. Dépannage d'impression (lien métier copieurs)

| Symptôme | Cause probable | Action |
|---|---|---|
| Rien ne s'imprime, file bloquée | Document corrompu / pilote | Purger la file (§41), redémarrer spouleur |
| "Pilote non disponible" sur postes | Type 3 + restrictions Point and Print | GPO Point and Print : serveurs approuvés ; ou passer en Type 4 |
| Impression lente (gros PDF) | Rendu côté client désactivé | Activer "Rendre les travaux sur les ordinateurs clients" (propriétés imprimante → Partage) |
| Le copieur n'est pas joint | Port TCP/IP en erreur | `Test-NetConnection 192.168.20.50 -Port 9100` ; vérifier SNMP |
| Codes départementaux non demandés | Pilote générique au lieu du KX Driver | Installer le pilote fabricant Type 3 validé |
| Bourrages non remontés | SNMP désactivé sur le port | Activer SNMP sur le port TCP/IP Standard |
| File "hors connexion" | Copieur éteint / IP changée (DHCP !) | **IP fixe ou réservation DHCP** sur tous les copieurs |

