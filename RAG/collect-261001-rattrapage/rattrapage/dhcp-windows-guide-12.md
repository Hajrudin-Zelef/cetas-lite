---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-12
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr", "attribution", "ethernet"]
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [1908, 2088]
sha256: 90df72fd72105dba9c05c93591d1a95937969af45306511574c49917de879da9
---

# 2. Le service DHCP tourne sur le partenaire ?
Get-Service -ComputerName "srv-dhcp-02.contoso.local" -Name DHCPServer

# 3. Le secret partagé correspond-il toujours ? (après un changement de mot de passe)
# 4. Horloges synchronisées ? (écart > 5 min = échec d'auth Kerberos/secret)
w32tm /query /status

# 5. Journal DHCP-Server : événements 1034/1035 (failover)
Get-WinEvent -LogName "Microsoft-Windows-DHCP-Server/Admin" -MaxEvents 30 |
    Where-Object { $_.Id -in 1034, 1035, 1040 } | Format-Table TimeCreated, Id, Message -AutoSize
```

**Résolution** :

```powershell
# Forcer la réplication
Invoke-DhcpServerv4FailoverReplication -ComputerName "srv-dhcp-01.contoso.local" -Name "Failover-LAN"

# Si le partenariat est cassé : le supprimer et le recréer
Remove-DhcpServerv4Failover -ComputerName "srv-dhcp-01.contoso.local" -Name "Failover-LAN"
# puis Add-DhcpServerv4Failover (section 60/61)
```

## 84. Cas n°8 — Les mises à jour DNS dynamiques échouent

**Symptômes** : les baux sont attribués mais les noms ne résolvent pas (`nslookup poste-123` → introuvable). Événement **1056** (échec d'authentification du compte DNS).

**Diagnostic** :

```powershell
# 1. Le DNS dynamique est-il activé sur l'étendue ?
Get-DhcpServerv4Scope -ScopeId 192.0.2.0 | Select-Object Name, DynamicDnsUpdateEnabled

# 2. Config globale
Get-DhcpServerv4DnsSetting | Format-List *

# 3. Le compte dédié est-il valide ? (mot de passe expiré ? verrouillé ?)
Get-ADUser -Identity "svc-dhcp-dns" -Properties PasswordExpired, LockedOut, PasswordLastSet |
    Select-Object Name, PasswordExpired, LockedOut, PasswordLastSet

# 4. Le compte est-il dans DnsUpdateProxy ?
Get-ADGroupMember -Identity "DnsUpdateProxy" | Select-Object Name

# 5. La zone DNS accepte-t-elle les mises à jour sécurisées ?
Get-DnsServerZone -Name "contoso.local" | Select-Object ZoneName, DynamicUpdate
# Doit être "Secure"
```

**Résolution** : réinitialiser le mot de passe du compte, le reconfigurer (`Set-DhcpServerDnsCredential`, section 71), vérifier que la zone est en mise à jour **sécurisée uniquement**.

## 85. Cas n°9 — Client avec APIPA 169.254.x.x : arbre de décision complet

L'adresse **APIPA** (169.254.0.0/16) = le client n'a reçu **aucune** réponse DHCP.

```text
Client en 169.254.x.x
│
├─ D'autres clients du même VLAN ont-ils une IP ?
│  ├─ NON → problème global : voir cas n°4 (relais) ou n°2 (étendue pleine)
│  └─ OUI → problème local au client : continuer
│
├─ ipconfig /renew → message ?
│  ├─ "Impossible de contacter le serveur DHCP" → réseau/câble/switch/802.1X
│  ├─ Timeout silencieux → filtre Deny ? (section 51) / réservation conflictuelle ?
│  └─ Erreur immédiate → service client DHCP arrêté ? (services.msc → DHCP Client)
│
├─ Wireshark (filtre bootp) sur le client :
│  ├─ DISCOVER émis, pas de OFFER → le broadcast ne sort pas (câble/port/VLAN)
│  ├─ OFFER reçu, pas de ACK → le client refuse (rare : option corrompue ?)
│  └─ Pas de DISCOVER → pile réseau HS (réinstaller le pilote, netsh winsock reset)
│
└─ Dernier recours : netsh int ip reset + reboot
```

```powershell
# Réinitialisation complète de la pile réseau côté client (à faire en local, pas à distance !)
netsh winsock reset
netsh int ip reset
# puis redémarrer
```

## 86. Cas n°10 — Baux fantômes après migration (doublons DNS/IP)

**Symptômes** : après une migration de serveur DHCP, des clients ont des IP en double, des noms DNS pointent vers de vieilles IP.

**Causes** : l'ancien serveur n'a pas été **désactivé** avant la mise en service du nouveau → deux DHCP actifs → baux divergents.

**Résolution** :

```powershell
# 1. S'assurer qu'un seul serveur est actif par étendue
# Sur l'ancien : désactiver les étendues
Set-DhcpServerv4Scope -ComputerName "srv-dhcp-OLD.contoso.local" -ScopeId 192.0.2.0 -State Inactive

# 2. Nettoyer les enregistrements DNS périmés (scavenging)
# Sur le DNS : activer le nettoyage (aging/scavenging) si pas fait
Set-DnsServerScavenging -ScavengingState $true -ScavengingInterval 7.00:00:00 -ApplyOnAllZones
Get-DnsServerZone -Name "contoso.local" |
    Set-DnsServerZoneAging -Aging $true -RefreshInterval 7.00:00:00 -NoRefreshInterval 7.00:00:00

# 3. Forcer le nettoyage manuel une fois
Start-DnsServerScavenging -Verbose
```

**Prévention** : procédure de migration avec failover (section 76), jamais deux serveurs indépendants.

## 87. Cas n°11 — Lenteurs à l'obtention du bail (logon lent)

**Symptômes** : les postes mettent 30–60 s à obtenir une IP au démarrage, l'ouverture de session est lente.

**Causes et réglages** :

| Cause | Correctif |
|-------|-----------|
| Détection de conflit activée (ping avant attribution) | La garder à 1 essai max, ou la désactiver sur LAN fiable |
| Serveur DHCP surchargé / distant (WAN) | Rapprocher (serveur local ou relais efficace) |
| STP sur les ports switch (30 s de blocking) | Activer **PortFast** sur les ports d'accès utilisateurs |
| 802.1X lent | Optimiser le timeout RADIUS |
| Client avec plusieurs NIC (Wi-Fi + Ethernet) | Désactiver la carte inutilisée |

```powershell
# Vérifier le temps de réponse du serveur (depuis un client)
Measure-Command { ipconfig /renew } | Select-Object TotalSeconds
# > 5 s = anormal sur un LAN local
```

> 💡 **PortFast** (ou équivalent) sur les ports des postes utilisateurs : le spanning-tree classique bloque le port 30–50 s au branchement, ce qui retarde le DORA. C'est la cause n°1 des "le réseau met une minute au démarrage".

## 88. Cas n°12 — Options DHCP non reçues par le client

**Symptômes** : le client a une IP mais pas la bonne passerelle / pas de DNS / pas de domaine.

**Diagnostic** :

```powershell
# 1. Quelles options le serveur envoie-t-il vraiment ?
Get-DhcpServerv4OptionValue -ScopeId 192.0.2.0 | Format-Table OptionId, Name, Value -AutoSize

# 2. Y a-t-il une STRATÉGIE qui écrase les options pour ce client ?
Get-DhcpServerv4Policy -ScopeId 192.0.2.0 | Format-Table Name, Enabled, ProcessingOrder

# 3. Une réservation avec options spécifiques ?
Get-DhcpServerv4Reservation -ScopeId 192.0.2.0 |
    Where-Object { $_.IPAddress -eq "192.0.2.87" }

# 4. Côté client : vider le cache et renouveler
ipconfig /release; ipconfig /flushdns; ipconfig /renew; ipconfig /all
```

**Piège** : une option définie au niveau **réservation** ou **stratégie** écrase l'option d'étendue (précédence, section 38/57). Vérifier les niveaux du plus spécifique au plus général.

## 89. Cas n°13 — Le serveur DHCP ne démarre plus (base corrompue)

**Symptômes** : le service DHCPServer ne démarre pas, événement **1014** (base de données corrompue).

**Résolution** :

```powershell
# 1. Arrêter le service
Stop-Service -Name DHCPServer -Force

# 2. Restaurer la base depuis la sauvegarde automatique
# Le DHCP fait une sauvegarde auto dans C:\Windows\System32\dhcp\backup
Copy-Item "C:\Windows\System32\dhcp\dhcp.mdb" "C:\Windows\System32\dhcp\dhcp.mdb.corrompu"
Copy-Item "C:\Windows\System32\dhcp\backup\new\dhcp.mdb" "C:\Windows\System32\dhcp\dhcp.mdb" -Force

# 3. Redémarrer et réconcilier
Start-Service -Name DHCPServer
Get-DhcpServerv4Scope | ForEach-Object {
    Invoke-DhcpServerv4ScopeReconciliation -ScopeId $_.ScopeId
}

# 4. Si pas de backup auto : restaurer l'export XML (section 75)
Import-DhcpServer -ComputerName "srv-dhcp-01.contoso.local" `
    -File "D:\Backup\DHCP\dhcp-backup.xml" -Leases -ScopeOverwrite -Force
```

**Prévention** : export quotidien automatisé (section 73) **+** sauvegarde auto native (vérifier qu'elle fonctionne : le dossier `backup\new` doit contenir un `dhcp.mdb` récent).

## 90. Cas n°14 — Épuisement par un équipement bavard (DHCP starvation)

