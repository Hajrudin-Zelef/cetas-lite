---
id: collect-261001-rattrapage/rattrapage/windows-server-guide-13
title: "Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["backlog"]
source: docs/RAG/collect-261001-rattrapage/windows_server_guide.md
source_anchor: ""
source_lines: [2053, 2231]
sha256: 3eb7e04a5807f31d350f23eeae2f3a53a2d9178f64335fa743d55c88dd57929c
---

# Windows Server en entreprise — Guide technique ultra-complet

- [ ] Vérifier les sauvegardes hors site / rotation des disques USB
- [ ] WSUS : nouvelles mises à jour critiques à approuver (anneau Test)
- [ ] VSS : clichés bien créés sur les volumes de partages
- [ ] Redémarrage des serveurs non critiques si patchs en attente
- [ ] Revoir les comptes verrouillés (4625) : attaque ou utilisateur étourdi ?
- [ ] Nettoyer les vieux checkpoints Hyper-V (> 7 jours)
- [ ] Vérifier les files d'impression + niveaux de toner (croiser avec la supervision copieurs)

---

## 75. Checklist mensuelle

- [ ] Patch Tuesday : déployer Test → Pilotes → Production (§27)
- [ ] Server Cleanup Wizard WSUS (§28)
- [ ] Réindexer la base WSUS si la console rame
- [ ] Vérifier les certificats expirant dans < 90 jours (PKI + serveurs web)
- [ ] Revoir les quotas FSRM et l'espace des partages
- [ ] Test de restauration : 1 fichier + 1 VM (alterner les serveurs)
- [ ] Revoir les droits NTFS des partages sensibles (qui a accès à la Compta ?)
- [ ] Mettre à jour la documentation (GLPI) : nouveaux serveurs, changements

```powershell
# Certificats expirant dans moins de 90 jours (machine locale)
Get-ChildItem Cert:\LocalMachine\My |
  Where-Object { $_.NotAfter -lt (Get-Date).AddDays(90) } |
  Format-Table Subject, NotAfter -AutoSize

# Idem sur l'AC : certificats émis expirant bientôt
certutil -view -restrict "NotAfter<$( (Get-Date).AddDays(90).ToString('MM/dd/yyyy') )" -out "CommonName,NotAfter"
```

---

## 76. Checklist trimestrielle / annuelle

**Trimestrielle :**
- [ ] Test de restauration **complète** (bare metal d'un serveur non critique)
- [ ] Revoir les GPO : lesquelles sont encore utiles ? (désactiver les mortes)
- [ ] Audit des comptes à privilèges (Domain Admins : qui ? pourquoi ?)
- [ ] Vérifier le quorum du cluster et le témoin
- [ ] Rotation des mots de passe de service critiques

**Annuelle :**
- [ ] Inventaire licences complet (§58) — avant le renouvellement des contrats
- [ ] Révision du PRA/PCA : le document est-il encore vrai ?
- [ ] Renouvellement du certificat de l'AC subordonnée si échéance < 2 ans
- [ ] Vérifier la CRL : publication OK, expiration lointaine (§50)
- [ ] Plan de remplacement matériel (garanties, fin de vie)
- [ ] Formation / habilitation de l'équipe sur les procédures

---

## 77. Checklist de mise en production d'un nouveau serveur

- [ ] Installation : édition LTSC, Core sauf justification (§7, §10)
- [ ] Nom conforme (`SRV-<ROLE>-<NN>`), IP statique, DNS, passerelle (§6)
- [ ] Domaine joint, OU correcte, GPO appliquées (`gpresult /r`)
- [ ] Fuseau horaire + NTP synchronisé
- [ ] Rôles installés **via script PowerShell versionné**
- [ ] Pare-feu : profil Domaine, règles minimales scopées (§61)
- [ ] Durcissement de base appliqué (§60), SMBv1 désactivé
- [ ] LAPS / Windows LAPS actif (§63)
- [ ] Mises à jour : WSUS ciblé sur le bon groupe (§26)
- [ ] Sauvegarde planifiée + testée (§52)
- [ ] Supervision : ajouté à Zabbix (ping, disque, services)
- [ ] Documentation GLPI : rôle, IP, garanties, mots de passe au coffre
- [ ] Recette : redémarrage propre, bascule testée si cluster

---

## 78. Pense-bête de poche : PowerShell

```powershell
# --- Système ---
Get-ComputerInfo | Select-Object WindowsProductName, WindowsVersion
Rename-Computer -NewName "SRV-X-01" -Restart
Get-HotFix | Sort-Object InstalledOn -Descending | Select-Object -First 5
systeminfo | Select-String "Durée"

# --- Réseau ---
Get-NetIPAddress -AddressFamily IPv4 | Where-Object PrefixOrigin -ne "WellKnown"
Test-NetConnection 192.168.10.1 -Port 445
Resolve-DnsName srv-ad-01.entreprise.lan
Get-NetFirewallProfile | Format-Table Name, Enabled

# --- Services & processus ---
Get-Service | Where-Object Status -eq "Running"
Restart-Service Spooler -Force
Get-Process | Sort-Object CPU -Descending | Select-Object -First 5

# --- Disques ---
Get-PSDrive -PSProvider FileSystem
Get-Volume | Format-Table DriveLetter, FileSystem, SizeRemaining, Size -AutoSize
Clear-RecycleBin -Force   # (avec prudence)

# --- Utilisateurs & AD ---
Get-ADUser -Filter * -SearchBase "OU=Utilisateurs,DC=entreprise,DC=lan" | Measure-Object
Get-ADComputer -Filter { Enabled -eq $false } | Disable-ADAccount  # déjà désactivés : à déplacer/supprimer
Unlock-ADAccount -Identity "j.dupont"

# --- Journaux ---
Get-EventLog -LogName System -Newest 20 -EntryType Error
Get-WinEvent -FilterHashtable @{LogName='Security'; ID=4625} -MaxEvents 10

# --- À distance ---
Enter-PSSession -ComputerName SRV-HV-01
Invoke-Command -ComputerName (Get-Content serveurs.txt) -ScriptBlock { Get-Service Spooler }
```

---

## 79. Pense-bête de poche : commandes et raccourcis

| Commande | Effet |
|---|---|
| `sconfig` | Menu de configuration (Server Core) |
| `slmgr.vbs /dlv` | Détail de la licence |
| `wbadmin get versions` | Sauvegardes disponibles |
| `dfsrdiag backlog ...` | Retard de réplication DFS-R |
| `certutil -ping -config ...` | L'AC répond-elle ? |
| `nltest /dclist:entreprise.lan` | Lister les DC du domaine |
| `dcdiag /v` | Diagnostic complet d'un DC |
| `repadmin /replsummary` | Résumé de la réplication AD |
| `gpresult /r` | GPO appliquées |
| `gpupdate /force` | Forcer l'application des GPO |
| `whoami /groups` | Groupes de l'utilisateur courant |
| `quser /server:SRV-RDS-01` | Sessions sur un hôte RDS |
| `vssadmin list shadows` | Clichés instantanés |
| `w32tm /query /status` | État de la synchro NTP |

---

## 80. Pense-bête de poche : ports réseau essentiels

| Port | Service |
|---|---|
| 53 TCP/UDP | DNS |
| 67/68 UDP | DHCP |
| 88 TCP/UDP | Kerberos (AD) |
| 123 UDP | NTP |
| 135 TCP | RPC (endpoint mapper) |
| 139 TCP | NetBIOS Session |
| 389 TCP/UDP | LDAP |
| 443 TCP | HTTPS (WAC, RD Gateway, WSUS SSL) |
| 445 TCP | SMB |
| 464 TCP/UDP | Kerberos password change |
| 636 TCP | LDAPS |
| 1433 TCP | SQL Server (défaut) |
| 3389 TCP | RDP |
| 5985 TCP | WinRM HTTP |
| 5986 TCP | WinRM HTTPS |
| 8530/8531 TCP | WSUS HTTP/HTTPS |
| 9100 TCP | Impression JetDirect (copieurs) |
| 1688 TCP | KMS |
| 8080 TCP | Réplication Hyper-V (Kerberos) |

---

## 81. Glossaire (A–M)

- **ABE** (Access-Based Enumeration) : l'utilisateur ne voit que les dossiers auxquels il a accès.
- **ADBA** : activation des licences via Active Directory, sans seuil.
- **CAL** : licence d'accès client (par utilisateur ou par appareil).
- **CSV** (Cluster Shared Volume) : volume partagé entre nœuds d'un cluster.
- **DC** (Domain Controller) : contrôleur de domaine Active Directory.
- **Dedup** : déduplication, élimine les blocs de données en double.
- **DFS-N / DFS-R** : espaces de noms distribués / réplication de fichiers.
- **DSRM** : mode de restauration des services d'annuaire (restauration DC).
- **FSRM** : gestionnaire de ressources de fichiers (quotas, filtrage).
- **GPO** : stratégie de groupe Active Directory.
- **Hyper-V Replica** : réplication asynchrone de VM entre hôtes.
- **KMS** : service de gestion des clés d'activation en volume.
- **LAPS** : mots de passe d'admin local uniques et rotatifs.
- **LTSC** : canal de support long (10 ans), la référence en production.

---

## 82. Glossaire (N–Z)

