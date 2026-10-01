---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/guides-guide-carp-haute-disponibilite-opnsense-pdf-d2e0a144-1
title: "OPNsense-1 (MASTER)"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/guides-guide-carp-haute-disponibilite-opnsense-pdf-d2e0a144.md
source_anchor: ""
source_lines: [1, 168]
sha256: bc08ca8e796d5157da78884591fe79d2cd583b520ce35e0c784c4dc967bda892
---

# OPNsense-1 (MASTER)

FR
Haute disponibilite avec
OPNsense et CARP
Paire actif/passif sans interruption
Mars 2026  |  Serie Stack OPNsense — Billet 10
BOTUM INC.  |  www.botum.ca  |  contact@botum.ca
Sommaire
1. Qu'est-ce que CARP ?
2. Prerequis : 2 VMs OPNsense sur Proxmox
3. Configurer les VIPs CARP
4. pfsync : synchronisation des etats
5. XMLRPC Config Sync
6. Tester le basculement (failover)
7. Cas d'usage production
8. Monitoring CARP avec Grafana
9. Prochaines etapes — SIEM Wazuh

1. Qu'est-ce que CARP ?
CARP (Common Address Redundancy Protocol) est un protocole reseau qui permet a plusieurs machines de
partager une adresse IP virtuelle (VIP). Developpe initialement pour OpenBSD et integre dans
OPNsense/pfSense, CARP garantit la haute disponibilite d'un firewall : si le noeud maitre tombe, le noeud de
secours prend le relai en quelques secondes, de maniere totalement transparente pour les connexions en
cours.
Dans mon infra BOTUM, j'utilise CARP pour deux cas critiques :
 Mises a jour OPNsense sans interruption : basculer sur le backup, mettre a jour le master, rebasculer
 Pannes materielles ou VM Proxmox : basculement automatique en moins de 3 secondes
 Maintenance planifiee : arret propre du master, le backup prend le relai
2. Prerequis : 2 VMs OPNsense sur Proxmox
Pour mettre en place la haute disponibilite CARP, il vous faut :
 2 VMs OPNsense identiques sur Proxmox (meme version, meme configuration de base)
 3 interfaces reseau par VM : WAN, LAN, et SYNC (pfsync)
 Un switch ou VLAN dedie pour le lien de synchronisation
 Acces SSH aux deux noeuds
 Meme version OPNsense sur les deux noeuds (imperatif)
Architecture reseau recommandee
# OPNsense-1 (MASTER)
# - vtnet0 : WAN  -> IP: 203.0.113.2/28 (reelle)
# - vtnet1 : LAN  -> IP: 192.168.1.1/24 (reelle)
# - vtnet2 : SYNC -> IP: 192.168.254.1/30
# OPNsense-2 (BACKUP)
# - vtnet0 : WAN  -> IP: 203.0.113.3/28 (reelle)
# - vtnet1 : LAN  -> IP: 192.168.1.2/24 (reelle)
# - vtnet2 : SYNC -> IP: 192.168.254.2/30
# VIPs CARP partagees (les clients utilisent ces IPs)
# - WAN VIP : 203.0.113.4/28  (VHID 1)
# - LAN VIP : 192.168.1.254/24 (VHID 2)
3. Configurer les VIPs CARP
Creer la VIP WAN sur le MASTER
Dans OPNsense MASTER : Interfaces > Virtual IPs > Add
Guide Pratique — Haute Disponibilite OPNsense CARP
BOTUM INC.
www.botum.ca  |  contact@botum.ca
Page 2

# Interfaces > Virtual IPs > Settings > Add
Type         : CARP
Interface    : WAN
IP Address   : 203.0.113.4 / 28
Virtual VHID : 1
VHID Password: MonMotDePasseCarp2026
Advertising frequency - Base: 1
Advertising frequency - Skew: 0   <- 0 = MASTER (priorite haute)
Description  : WAN-VIP-CARP
Creer la VIP LAN sur le MASTER
# Repeter pour l'interface LAN
Type         : CARP
Interface    : LAN
IP Address   : 192.168.1.254 / 24
Virtual VHID : 2
VHID Password: MonMotDePasseCarp2026
Advertising frequency - Base: 1
Advertising frequency - Skew: 0
Description  : LAN-VIP-CARP
Configurer le BACKUP (Skew 100)
Sur OPNsense BACKUP, creer les memes VIPs avec Skew = 100 :
# Sur OPNsense-2 (BACKUP) — meme config, seul le skew change :
# WAN VIP : VHID 1, Skew 100  <- 100 = BACKUP (priorite basse)
# LAN VIP : VHID 2, Skew 100
# Regles CARP :
# - Skew 0   = priorite haute -> MASTER actif
# - Skew 100 = priorite basse -> BACKUP en attente
# - Plus le skew est bas, plus le noeud est prioritaire
4. pfsync : Synchronisation des etats de connexion
pfsync synchronise les tables d'etats du firewall entre les deux noeuds en temps reel. Ainsi, lors d'un
basculement, les connexions TCP actives (SSH, HTTPS, VPN) ne sont pas interrompues — le backup connait
deja tous les etats.
Activer pfsync sur le MASTER
System > High Availability > Settings
Guide Pratique — Haute Disponibilite OPNsense CARP
BOTUM INC.
www.botum.ca  |  contact@botum.ca
Page 3

# System > High Availability > Settings
[High Availability Sync]
  Synchronize States (pfsync)   : checked
  Synchronize Interface         : SYNC  (vtnet2)
  pfsync Synchronize Peer IP    : 192.168.254.2  <- IP SYNC du BACKUP
[Firewall]
  Synchronize firewall rules    : checked
  Synchronize NAT               : checked
  Synchronize static routes     : checked
Activer pfsync sur le BACKUP
# System > High Availability > Settings (sur BACKUP)
[High Availability Sync]
  Synchronize States (pfsync)   : checked
  Synchronize Interface         : SYNC  (vtnet2)
  pfsync Synchronize Peer IP    : 192.168.254.1  <- IP SYNC du MASTER
i Le lien SYNC doit etre isole sur un VLAN ou reseau dedie. Ne jamais faire transiter du trafic utilisateur sur cette interface.
5. XMLRPC Config Sync
XMLRPC Config Sync permet de propager automatiquement la configuration OPNsense du MASTER vers le
BACKUP. Chaque modification sur le MASTER (nouvelle regle firewall, NAT, alias) est synchronisee sans
intervention manuelle.
Configurer XMLRPC sur le MASTER
# System > High Availability > Settings
[Configuration Synchronization]
  Synchronize Config to IP    : 192.168.254.2  <- IP SYNC du BACKUP
  Remote System Username      : root  (ou admin)
  Remote System Password      : MotDePasseBackup2026
# Sections a synchroniser (cocher toutes) :
  Aliases             Certificates      DHCP Server
  DNS Resolver        Firewall Rules    Gateways
  Interfaces          NAT               OpenVPN / WireGuard
  Routes              Users and Groups
i Apres chaque modification sur le MASTER, cliquer Save & Sync pour propager immediatement.
6. Tester le basculement (Failover)
Le test de basculement est OBLIGATOIRE avant de passer en production.
Test 1 : Basculement manuel (arret propre)
Guide Pratique — Haute Disponibilite OPNsense CARP
BOTUM INC.
www.botum.ca  |  contact@botum.ca
Page 4

# System > High Availability > Status
# 1. Verifier l'etat initial :
#    MASTER : CARP State = MASTER (VIPs actives)
#    BACKUP : CARP State = BACKUP (VIPs en attente)
# 2. Sur le MASTER : passer en mode BACKUP manuellement
pfctl -d  # Desactiver le firewall (force le basculement)
# OU : System > High Availability > Forcefully become BACKUP
# 3. Observer le BACKUP prendre le relai (~2-3 secondes)
Test 2 : Simulation panne (arret brutal)
# Dans Proxmox, eteindre la VM OPNsense-1 (MASTER) :
# qm stop 100  (ou via l'interface Proxmox)
# Sur un client reseau, lancer un ping continu vers la VIP :
ping 203.0.113.4
# -> Vous devriez voir 1-2 pertes max pendant le basculement
# -> Le backup reprend en 1-3 secondes
# Verifier sur le BACKUP :
# System > High Availability > Status
# VIP WAN : MASTER | VIP LAN : MASTER  OK
Test 3 : Connexions actives survivent au failover
# Sur un poste client, ouvrir une session SSH via la VIP LAN :
ssh admin@192.168.1.254
# Pendant que la session SSH est ouverte, forcer le basculement
# La session SSH doit SURVIVRE grace a pfsync
# (les etats TCP sont synchronises en temps reel)
# Si la session SSH survit  -> pfsync fonctionne correctement
# Si la session SSH se coupe -> verifier la config pfsync
7. Cas d'usage : Production sans interruption
La principale valeur ajoutee de CARP dans mon infra BOTUM est la possibilite de faire des maintenances
sans fenetre d'indisponibilite.
Mise a jour OPNsense sans downtime
Guide Pratique — Haute Disponibilite OPNsense CARP
BOTUM INC.
www.botum.ca  |  contact@botum.ca
Page 5

