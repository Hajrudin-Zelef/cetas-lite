---
id: collect-261001-rattrapage/rattrapage/windows-server-guide-11
title: "Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "arr", "backlog"]
source: docs/RAG/collect-261001-rattrapage/windows_server_guide.md
source_anchor: ""
source_lines: [1784, 1930]
sha256: 244c1d9d44cccbe6eaffdbb83c10d84ffcf75eea8078101405c306c446eaae76
---

# Windows Server en entreprise — Guide technique ultra-complet

---

## 62. Microsoft Defender Antivirus sur serveur

```powershell
# État
Get-MpComputerStatus | Select-Object AntivirusEnabled, AntispywareEnabled, AMEngineVersion, AntivirusSignatureLastUpdated

# Mise à jour des signatures
Update-MpSignature

# Scan rapide / complet
Start-MpScan -ScanType QuickScan

# Exclusions OBLIGATOIRES par rôle (exemples) :
# Hyper-V : dossiers VM + vmms.exe, vmwp.exe
Add-MpPreference -ExclusionPath "D:\VMs"
Add-MpPreference -ExclusionProcess "C:\Windows\System32\vmms.exe"
# AD/DNS : NTDS, SYSVOL
Add-MpPreference -ExclusionPath "C:\Windows\NTDS", "C:\Windows\SYSVOL"
# SQL : .mdf/.ldf/.bak + dossier data
Add-MpPreference -ExclusionExtension "mdf","ldf","bak"

# Protection en temps réel (ne JAMAIS désactiver durablement)
Set-MpPreference -DisableRealtimeMonitoring $false

# Voir les menaces détectées
Get-MpThreatDetection | Format-Table ThreatID, Resources, InitialDetectionTime -AutoSize
```

> Sur Server Core, Defender se pilote en PowerShell ou via GPO/WAC. La désinstallation de Defender au profit d'un tiers se fait proprement (pas de "désactivation sauvage").

---

## 63. LAPS et gestion des mots de passe locaux

**Problème :** même mot de passe Administrateur local sur 50 serveurs = 1 compromission = tout le parc.

**LAPS (Local Administrator Password Solution)** : mot de passe unique, aléatoire, rotatif, stocké dans l'AD.

```powershell
# Installer LAPS (fonctionnalité + GPO côté admin)
# 1. Étendre le schéma AD (une fois, sur le DC) :
Update-AdmPwdADSchema
# 2. Déléguer le droit d'écrire le mot de passe aux ordinateurs (sur l'OU) :
Set-AdmPwdComputerSelfPermission -OrgUnit "OU=Serveurs,DC=entreprise,DC=lan"
# 3. Déléguer la lecture aux admins :
Set-AdmPwdReadPasswordPermission -OrgUnit "OU=Serveurs,DC=entreprise,DC=lan" -AllowedPrincipals "ENTREPRISE\GRP-Admins-LAPS"
# 4. GPO : Modèles d'administration → LAPS → activer + complexité + durée (30 jours)

# Lire le mot de passe d'un serveur :
Get-AdmPwdPassword -ComputerName SRV-FICHIER-01

# Forcer la rotation :
Reset-AdmPwdPassword -ComputerName SRV-FICHIER-01
```

> **Windows LAPS** (intégré depuis 2023, natif sur 2025) remplace le LAPS legacy : GPO `Configuration ordinateur → Modèles d'administration → Système → LAPS`, stockage dans l'AD ou Entra ID, chiffrement possible. En 2025, utiliser **Windows LAPS natif**.

---

## 64. Audit, journaux et supervision minimale

```powershell
# Activer l'audit des échecs d'ouverture de session + accès objets (GPO recommandé, ici en local)
auditpol /set /subcategory:"Logon" /success:enable /failure:enable
auditpol /set /subcategory:"File System" /failure:enable

# Taille des journaux (éviter l'écrasement en 2 jours)
Limit-EventLog -LogName Security -MaximumSize 1GB -OverflowAction OverwriteAsNeeded

# Recherches utiles
Get-WinEvent -FilterHashtable @{LogName='Security'; ID=4625} -MaxEvents 20 |  # échecs de logon
  Format-Table TimeCreated, @{N='Compte';E={$_.Properties[5].Value}}, @{N='IP';E={$_.Properties[19].Value}} -AutoSize
Get-WinEvent -FilterHashtable @{LogName='System'; ID=1074,1076} -MaxEvents 10  # arrêts/redémarrages

# Transférer vers un collecteur (WEF - Windows Event Forwarding, gratuit)
# Sur le collecteur :
wecutil qc /q
# Abonnement : journaux Security/System des serveurs → collecteur central
```

**Supervision minimale viable :** ping + ports (445, 3389, 8530...) + espace disque + journaux critiques → Zabbix (voir `zabbix_guide.md`), avec remontée SNMP ou agent.

---

---

## 65. Cas pratiques commentés (1 à 5)

### Cas n°1 — Migrer un serveur de fichiers 2019 → 2022 sans coupure visible
**Contexte :** `SRV-FICHIER-01` (2019) saturé, nouveau `SRV-FICHIER-03` (2022).
1. Installer le rôle Fichiers + DFS-N/DFS-R sur le nouveau.
2. Créer le groupe de réplication DFS-R, nouveau serveur en membre (pas primaire).
3. Laisser la réplication initiale se faire (surveiller le backlog, §34).
4. Ajouter `\\SRV-FICHIER-03\partage` comme **cible DFS-N** (désactivée d'abord).
5. Basculer : activer la nouvelle cible, désactiver l'ancienne. Les utilisateurs via `\\entreprise.lan\data\...` ne voient rien.
6. Après 1 semaine de validation : décommissionner l'ancien.
*Leçon : DFS-N (§33) rend les migrations de serveurs de fichiers transparentes.*

### Cas n°2 — Le WSUS ne propose plus rien depuis 3 mois
**Diagnostic :** console WSUS → Synchronisations : dernière réussie il y a 92 jours. **Cause :** le disque D: plein à 100 % (contenu WSUS + logs IIS). **Actions :** libérer de l'espace, `wsusutil reset`, relancer la synchro, mettre en place le nettoyage mensuel (§28) + alerte disque à 85 % dans Zabbix. *Leçon : superviser l'espace disque du WSUS comme un serveur de production.*

### Cas n°3 — Un copieur n'imprime plus après changement de switch
**Diagnostic :** file "hors connexion". `Test-NetConnection 192.168.20.50 -Port 9100` → échec. **Cause :** le copieur était en DHCP, le nouveau switch a un autre VLAN par défaut → le copieur a changé d'IP. **Actions :** IP fixe sur le copieur (ou réservation DHCP), corriger le port TCP/IP sur le serveur d'impression, documenter les IP des copieurs. *Leçon : IP fixe ou réservation DHCP sur TOUS les périphériques d'impression (§42).*

### Cas n°4 — Ransomware : que faire à 3h du matin ?
1. **Isoler** : débrancher le réseau des serveurs touchés (pas d'arrêt brutal si possible — la mémoire contient des preuves).
2. **Ne pas payer**, ne pas "nettoyer" avant d'avoir des copies des logs.
3. Identifier le patient zéro (journaux 4624/4625, EDR).
4. Reconstruire sur infrastructure **saine** (nouveaux disques/VM).
5. Restaurer depuis la copie **hors site** 3-2-1 (§54).
6. Changer **tous** les mots de passe (admin, services, utilisateurs).
7. Post-mortem : comment est-il entré ? (RDP exposé ? phishing ? SMBv1 ?)
*Leçon : le PRA se prépare avant, pas pendant.*

### Cas n°5 — Ajouter un hôte Hyper-V au cluster un dimanche
1. Mêmes patchs que les nœuds existants (sinon `Test-Cluster` échoue).
2. Mêmes vSwitch (noms identiques !), mêmes réseaux/VLAN.
3. `Add-ClusterNode -Name SRV-HV-03 -Cluster CLUSTER-HV`.
4. Vérifier le quorum et migrer une VM de test dessus.
*Leçon : l'homogénéité des nœuds (noms de vSwitch, patchs) est la condition n°1 d'un cluster sain.*

---

## 66. Cas pratiques commentés (6 à 10)

### Cas n°6 — Déployer 200 postes avec une imprimante par défaut selon l'étage
GPO par OU/étage : déploiement **par ordinateur** (§39) du copieur d'étage + script `SetDefaultPrinter`. Filtrage WMI ou groupes de sécurité par étage. Tester avec `gpupdate /force` + redémarrage sur 2 postes pilotes.

### Cas n°7 — Mettre en place la PKI pour le Wi-Fi d'entreprise (802.1X)
1. AC subordonnée (§48), modèles "Ordinateur" et "Utilisateur" avec auto-enrollment (§49).
2. GPO auto-enrollment sur OU Postes et OU Utilisateurs.
3. NPS (serveur de stratégies réseau) + certificats serveur.
4. SSID WPA2/3-Enterprise, validation du certificat serveur côté client.
*Leçon : sans PKI interne, pas de 802.1X sérieux.*

### Cas n°8 — Le serveur RDS refuse les connexions le lundi matin
**Message :** "Aucun serveur de licences...". **Cause :** délai de grâce de 120 jours expiré pendant les vacances (§45). **Action immédiate :** installer/configurer le rôle RDS-Licensing + CAL. **Préventif :** alerte à J+100 dans l'agenda d'exploitation.

### Cas n°9 — Espace disque critique sur le volume des VM
1. Identifier : `Get-VM | Get-VMHardDiskDrive` + tailles VHDX ; checkpoints oubliés (`Get-VMSnapshot`).
2. Supprimer/fusionner les vieux checkpoints (hors heures de pointe — la fusion consomme des I/O).
3. Compacter les VHDX dynamiques (`Optimize-VHD -Mode Full`, VM éteinte).
4. Déplacer une VM vers un autre volume (`Move-VMStorage`).
*Leçon : alerte disque à 80 % sur les volumes Hyper-V, pas à 95 %.*

