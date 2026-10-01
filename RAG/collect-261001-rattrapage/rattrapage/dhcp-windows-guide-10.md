---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-10
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [1556, 1732]
sha256: 9ff638ccbfe0917d79a0fd9bf2dfae9fcb0108cfc768d79e60bdfe31d411cb24
---

# Guide technique ultra-complet : DHCP sous Windows Server en entreprise

```powershell
# 1. Créer le compte dans l'AD (via ADUC ou PowerShell)
New-ADUser -Name "svc-dhcp-dns" -SamAccountName "svc-dhcp-dns" `
    -UserPrincipalName "svc-dhcp-dns@contoso.local" `
    -AccountPassword (Read-Host -AsSecureString "Mot de passe") `
    -Enabled $true -PasswordNeverExpires $true `
    -Description "Compte de mises a jour DNS dynamiques pour DHCP"

# 2. L'ajouter au groupe "DnsUpdateProxy" (autorise la mise à jour des zones sécurisées)
Add-ADGroupMember -Identity "DnsUpdateProxy" -Members "svc-dhcp-dns"

# 3. Configurer le DHCP pour utiliser ce compte (sur CHAQUE serveur DHCP)
# En console : clic droit serveur -> Propriétés -> onglet Avancé -> Informations d'identification
# En PowerShell :
Set-DhcpServerDnsCredential -ComputerName "srv-dhcp-01.contoso.local" `
    -Credential (Get-Credential -UserName "contoso\svc-dhcp-dns" -Message "Compte DNS DHCP")
```

> ⚠️ Le groupe **DnsUpdateProxy** : ses membres peuvent créer des enregistrements DNS **non sécurisés**. Ne jamais y mettre un compte à privilèges. Et **ne pas** mettre les comptes machines des serveurs DHCP eux-mêmes dans ce groupe si tu utilises le compte dédié (sinon les protections ne s'appliquent pas correctement).
>
> ⚠️ Si le mot de passe du compte change/expire : les mises à jour DNS **échouent silencieusement** → DNS fantômes. Superviser l'événement **1056** (échec d'identification DNS).

## 72. Protection de nom (Name Protection / DHCID)

La **protection de nom** empêche qu'un client malveillant (ou un homonyme) **écrase** l'enregistrement DNS d'une autre machine. Le DHCP ajoute un enregistrement **DHCID** qui lie le nom à la MAC du client légitime.

```powershell
# Activer (inclus dans la config section 70 : -NameProtection $true)
# Vérifier :
Get-DhcpServerv4DnsSetting | Select-Object NameProtection
```

> 💡 À activer systématiquement sur les réseaux avec des équipements non gérés (Wi-Fi invités, salles de formation). Sur un LAN 100 % managé, c'est un plus de sécurité à coût nul.

## 73. Sauvegarde du DHCP — Export-DhcpServer (méthode moderne)

```powershell
# Sauvegarde complète : config + baux + (optionnel) secrets de basculement
Export-DhcpServer `
    -ComputerName "srv-dhcp-01.contoso.local" `
    -File "D:\Backup\DHCP\dhcp-backup.xml" `
    -Leases `
    -Force

# Inclure aussi les paramètres de sécurité (ACL) — Windows Server 2022/2025
Export-DhcpServer `
    -ComputerName "srv-dhcp-01.contoso.local" `
    -File "D:\Backup\DHCP\dhcp-backup-full.xml" `
    -Leases -Force
```

Contenu du fichier XML : étendues, options, réservations, baux, stratégies, filtres, config failover. **Lisible** (XML) → utile aussi comme documentation.

Automatiser (tâche planifiée quotidienne) :

```powershell
# Script C:\Admin\Backup-Dhcp.ps1
$date = Get-Date -Format "yyyy-MM-dd"
$dir  = "D:\Backup\DHCP"
if (-not (Test-Path $dir)) { New-Item -ItemType Directory -Path $dir | Out-Null }
Export-DhcpServer -ComputerName "srv-dhcp-01.contoso.local" `
    -File "$dir\dhcp-$date.xml" -Leases -Force
# Rotation : garder 30 jours
Get-ChildItem "$dir\dhcp-*.xml" | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-30) } |
    Remove-Item
```

```powershell
# Créer la tâche planifiée (quotidienne 02h00)
$action    = New-ScheduledTaskAction -Execute "powershell.exe" -Argument "-File C:\Admin\Backup-Dhcp.ps1"
$trigger   = New-ScheduledTaskTrigger -Daily -At 02:00
$principal = New-ScheduledTaskPrincipal -UserId "SYSTEM" -LogonType ServiceAccount
Register-ScheduledTask -TaskName "Backup DHCP quotidien" -Action $action -Trigger $trigger -Principal $principal
```

## 74. Sauvegarde — netsh dhcp (méthode classique, toujours valable)

```powershell
# Sauvegarde complète via netsh (format binaire propriétaire)
netsh dhcp server \\srv-dhcp-01 export D:\Backup\DHCP\dhcp-netsh.dat all

# Sauvegarde d'une étendue seule
netsh dhcp server \\srv-dhcp-01 scope 192.0.2.0 dump > D:\Backup\DHCP\scope-192.0.2.0.txt
# (le "dump" produit un script netsh rejouable — excellent pour la doc !)
```

> 💡 Le `dump` netsh est un **script texte rejouable** : parfait pour documenter une étendue ou la recréer à l'identique ailleurs.

## 75. Restauration — Import-DhcpServer

```powershell
# Restaurer sur un serveur (neuf ou après crash)
# Prérequis : rôle DHCP installé, serveur autorisé dans AD
Import-DhcpServer `
    -ComputerName "srv-dhcp-02.contoso.local" `
    -File "D:\Backup\DHCP\dhcp-backup.xml" `
    -Leases `
    -ScopeOverwrite `
    -Force

# -ScopeOverwrite : écrase les étendues existantes avec le même ScopeId
# Sans -Leases : restaure la config sans les baux (utile pour une migration propre)
```

Procédure de **reprise après crash** :

1. Installer Windows Server + rôle DHCP sur la machine de remplacement.
2. Lui donner la **même IP** que l'ancien serveur (ou mettre à jour les IP helpers).
3. Autoriser dans AD (`Add-DhcpServerInDC`).
4. `Import-DhcpServer` avec le dernier backup.
5. Reconfigurer le compte DNS dédié (section 71) — **non inclus** dans l'export.
6. Reconfigurer le basculement (section 59) — vérifier le secret partagé.
7. Tester : `ipconfig /renew` depuis un poste de chaque VLAN.

## 76. Migration vers un nouveau serveur (procédure sans coupure)

```powershell
# === Sur l'ANCIEN serveur (srv-dhcp-01) ===
Export-DhcpServer -ComputerName "srv-dhcp-01.contoso.local" `
    -File "\\srv-dhcp-02\Partage\dhcp-migration.xml" -Leases -Force

# === Sur le NOUVEAU serveur (srv-dhcp-02) ===
# 1. Installer le rôle + autoriser
Install-WindowsFeature -Name DHCP -IncludeManagementTools
Add-DhcpServerInDC -DnsName "srv-dhcp-02.contoso.local" -IPAddress 192.0.2.11

# 2. Importer (sans écraser si le nouveau est vierge, -ScopeOverwrite inutile)
Import-DhcpServer -ComputerName "srv-dhcp-02.contoso.local" `
    -File "\\srv-dhcp-02\Partage\dhcp-migration.xml" -Leases -Force

# 3. Basculer les IP helpers vers le nouveau (ou monter le failover entre les deux)
# 4. Désactiver les étendues de l'ancien, attendre l'expiration des baux, décommissionner
```

> 💡 Migration **sans coupure** : monter d'abord un **failover** entre l'ancien et le nouveau (section 59), laisser répliquer, puis passer le nouveau en primaire et retirer l'ancien. Zéro interruption.

---

# Bloc G — Dépannage : 15 cas pratiques commentés

> Méthode générale : **1)** le client obtient-il une IP ? (`ipconfig`) **2)** DORA visible ? (Wireshark `bootp` ou logs serveur) **3)** Quelle étape bloque ? (DISCOVER sans OFFER → serveur/relais ; OFFER sans REQUEST → client ; REQUEST sans ACK → étendue/options).

## 77. Cas n°1 — Conflit d'adresses IP (le classique)

**Symptômes** : message Windows *"Conflit d'adresse IP avec un autre système sur le réseau"*, connectivité intermittente sur deux machines.

**Causes** :

- Une IP statique configurée **dans la plage DHCP** sans exclusion (erreur n°1).
- Un équipement avec IP statique rebranché après un changement de plage.
- Deux serveurs DHCP indépendants sur le même segment.

**Diagnostic** :

```powershell
# 1. Vérifier si l'IP en conflit est dans une plage DHCP
Get-DhcpServerv4Scope | Where-Object {
    $_.StartRange -le [ipaddress]"192.0.2.87" -and $_.EndRange -ge [ipaddress]"192.0.2.87"
} | Select-Object Name, ScopeId, StartRange, EndRange

# 2. Vérifier les exclusions couvrent-elles les IP statiques connues ?
Get-DhcpServerv4ExclusionRange -ScopeId 192.0.2.0

# 3. Détecter la MAC du "squatteur" : ping puis arp
ping 192.0.2.87
arp -a | Select-String "192.0.2.87"

# 4. Le bail DHCP correspondant existe-t-il ?
Get-DhcpServerv4Lease -ScopeId 192.0.2.0 | Where-Object { $_.IPAddress -eq "192.0.2.87" }
```

**Résolution** :

```powershell
# Ajouter l'exclusion pour l'IP statique légitime
Add-DhcpServerv4ExclusionRange -ScopeId 192.0.2.0 -StartRange 192.0.2.87 -EndRange 192.0.2.87

