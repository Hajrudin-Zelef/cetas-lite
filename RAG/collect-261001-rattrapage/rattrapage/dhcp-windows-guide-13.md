---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-13
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [2089, 2255]
sha256: 5600eef65ce745a2a36ee927febcb05b1eaa9ab7268a59e3a35ead1eb0459660
---

# Guide technique ultra-complet : DHCP sous Windows Server en entreprise

**Symptômes** : étendue pleine en quelques heures alors que le nombre d'équipements est stable. Dans les baux : des dizaines d'entrées avec des **MAC aléatoires** ou le même hostname.

**Cause** : attaque **DHCP starvation** (outil type Yersinia) ou équipement défectueux qui redemande en boucle avec des MAC forgées.

**Diagnostic** :

```powershell
# Baux créés récemment en masse ?
Get-DhcpServerv4Lease -ScopeId 192.0.3.0 |
    Where-Object { $_.AddressState -eq "Active" } |
    Group-Object ClientId | Where-Object { $_.Count -gt 3 } |
    Select-Object Name, Count

# MAC aléatoires = beaucoup de ClientId uniques créés sur une courte période
# Vérifier les logs : rafale de DISCOVER depuis une même source (giaddr / port switch)
```

**Résolution** :

1. Identifier le port du switch (via l'IP source des DISCOVER ou le snooping).
2. Shutdown le port, identifier l'équipement (attaque ? IoT défectueux ? boucle ?).
3. Nettoyer les baux parasites, réduire temporairement le bail pour purger vite.
4. Activer **DHCP snooping + rate limiting** (section 94) pour bloquer les rafales.

## 91. Cas n°15 — Après un changement d'IP du serveur DHCP, plus rien ne fonctionne

**Symptômes** : changement d'IP du serveur DHCP (ou migration) → les clients ne renouvellent plus, les relais n'atteignent plus le serveur.

**Checklist post-changement** :

```powershell
# 1. Ré-autoriser dans AD avec la nouvelle IP (l'autorisation lie DNS + IP)
Remove-DhcpServerInDC -DnsName "srv-dhcp-01.contoso.local" -IPAddress 192.0.2.10
Add-DhcpServerInDC    -DnsName "srv-dhcp-01.contoso.local" -IPAddress 192.0.2.99

# 2. Mettre à jour TOUS les ip helper-address (chaque VLAN !)
# 3. Mettre à jour le DNS dynamique / compte de service si lié à l'IP
# 4. Mettre à jour la supervision (sondes vers la nouvelle IP)
# 5. Mettre à jour le failover (le partenariat référence le serveur : vérifier)
Get-DhcpServerv4Failover | Select-Object Name, PartnerServer, State

# 6. Forcer le renouvellement côté clients (GPO / script de logon temporaire)
# ipconfig /renew
```

> 💡 **Leçon** : l'IP d'un serveur DHCP est une **référence d'infrastructure** (helpers, supervision, doc). La changer = projet mini-migration. Préférer le failover pour absorber les changements (section 76).

## 91bis. Journaux DHCP : activer et exploiter l'audit

Le DHCP loggue quotidiennement dans `C:\Windows\System32\dhcp\DhcpSrvLog-*.log` (un fichier par jour de la semaine).

```powershell
# Vérifier que l'audit est activé
Get-DhcpServerAuditLog -ComputerName "srv-dhcp-01.contoso.local" |
    Select-Object Enable, Path

# Activer si besoin
Set-DhcpServerAuditLog -ComputerName "srv-dhcp-01.contoso.local" -Enable $true -Path "C:\Windows\System32\dhcp"

# Analyser le log du jour : codes d'événements (ID)
# 10 = nouveau bail, 11 = renouvellement, 12 = release, 13 = conflit détecté,
# 14 = réservation, 15 = NACK, 16 = suppression de bail
$log = Get-ChildItem "C:\Windows\System32\dhcp\DhcpSrvLog-$(Get-Date -Format 'ddd').log" |
    Select-Object -First 1
Import-Csv $log.FullName -Header ID,Date,Time,Description,IP,HostName,MAC | 
    Where-Object { $_.ID -eq 13 } |  # conflits détectés
    Format-Table Date, Time, IP, HostName, MAC -AutoSize
```

> 💡 Les logs d'audit sont la **source de vérité** pour : prouver quel équipement avait quelle IP à quel moment (traçabilité légale), détecter les starvations (rafales ID 10), et auditer les rogue (OFFER multiples).

---

# Bloc H — Sécurité

## 92. Menaces contre le DHCP en entreprise

| Menace | Description | Impact |
|--------|-------------|--------|
| Rogue DHCP | Serveur illégitime qui répond aux DISCOVER | Redirection du trafic (MITM), DNS menteur, coupure |
| DHCP starvation | Épuisement volontaire de l'étendue (MAC forgées) | Déni de service (plus d'IP pour les légitimes) |
| DHCP spoofing | Fausse réponse OFFER/ACK ciblée | MITM ciblé |
| Écoute passive | Capture des DISCOVER (noms de machines) | Reconnaissance réseau |

## 93. Défense n°1 : autorisation Active Directory (rappel)

Déjà couverte section 9 : un serveur Windows non autorisé ne distribue pas. **Limites** : ne bloque ni les serveurs Linux, ni les box/routeurs domestiques, ni les téléphones en partage de connexion. → Compléter avec le DHCP snooping.

## 94. Défense n°2 : DHCP snooping sur les switchs (indispensable)

Le **DHCP snooping** fait du switch un pare-feu DHCP : seuls les ports **trusted** (vers les serveurs DHCP légitimes / uplinks) peuvent envoyer des messages **serveur** (OFFER, ACK, NAK). Les ports **untrusted** (utilisateurs) ne peuvent qu'envoyer des messages **client** (DISCOVER, REQUEST).

**Cisco IOS** :

```text
! Activer globalement (par VLAN)
ip dhcp snooping
ip dhcp snooping vlan 10,20,30,40,50

! Ports vers les serveurs DHCP / uplinks : TRUSTED
interface GigabitEthernet1/0/1
 description Uplink vers SRV-DHCP-01
 ip dhcp snooping trust

! Optionnel : limiter le débit DHCP sur les ports utilisateurs (anti-starvation)
interface range GigabitEthernet1/0/2 - 48
 ip dhcp snooping limit rate 15

! Vérifier
show ip dhcp snooping
show ip dhcp snooping binding
```

**HP/Aruba (Comware)** :

```text
dhcp snooping enable
interface GigabitEthernet1/0/1
 dhcp snooping trust
```

Points d'attention :

- Activer le snooping **VLAN par VLAN** (ne pas oublier les nouveaux VLANs).
- Les ports d'**uplink inter-switch** doivent être trusted (sinon les OFFER légitimes venant d'un autre switch sont bloqués).
- Le **rate limiting** (15 paquets/s) bloque les starvations sans gêner l'usage normal.
- La table de **binding** (MAC ↔ IP ↔ port ↔ VLAN) sert aussi à **Dynamic ARP Inspection** et **IP Source Guard** — triple protection avec une seule config.

> **Lien réseau / métier** : en tant que chef de service systèmes & énergies, c'est typiquement un chantier **conjoint** avec l'équipe réseau : le DHCP snooping se configure sur les switchs, pas sur Windows. Prévoir une fenêtre de maintenance (un port mal classé = DHCP coupé sur ce port).

## 95. Défense n°3 : 802.1X (la vraie solution d'accès)

802.1X authentifie l'équipement **avant** même qu'il puisse envoyer un DISCOVER. Un rogue DHCP branché sur un port 802.1X ne peut pas parler au réseau.

```text
[Équipement] --EAP--> [Switch] --RADIUS--> [NPS/AD]
     │                     │
     │   Port bloqué tant   │
     │   que non authentifié│
     ▼                     ▼
  Pas de DHCP tant que l'authentification n'a pas réussi
```

Windows Server fournit le **NPS** (Network Policy Server, rôle à installer) comme serveur RADIUS. C'est un projet en soi, mais c'est **la** réponse structurelle aux rogue DHCP + équipements non autorisés.

## 96. Durcissement du serveur DHCP Windows

```powershell
# 1. Compte de service dédié avec droits minimaux (pas de compte admin du domaine)
# 2. Pare-feu : n'autoriser UDP 67/68 que depuis les relais et le LAN de management
# 3. Désactiver les services inutiles sur le serveur DHCP dédié
# 4. Audit des modifications : qui touche au DHCP ?
# Activer l'audit des cmdlets sensibles via les logs PowerShell (déjà natif)

# 5. Restreindre qui peut administrer le DHCP (groupes locaux)
net localgroup "DHCP Administrators"
# N'y mettre que les admins DHCP, pas "Domain Admins" entier si possible

# 6. Protéger les sauvegardes (elles contiennent tout le plan d'adressage)
icacls "D:\Backup\DHCP" /inheritance:r /grant:r "Administrateurs:F" /grant:r "SYSTEM:F"

# 7. Détection de conflit activée (section 77) — aussi une mesure anti-spoofing légère
Set-DhcpServerSetting -ConflictDetectionAttempts 1
```

Checklist de durcissement :

