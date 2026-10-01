---
id: collect-261001-rattrapage/rattrapage/opnsense-gui-guide-1
title: "OPNsense — Guide ultra-complet de l'interface web"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-rattrapage/opnsense_gui_guide.md
source_anchor: ""
source_lines: [1, 197]
sha256: 552df833b5f292555ee09db263a0dbbf36a0acc0c6ce0b2e2171384ea53d4e5c
---

# OPNsense — Guide ultra-complet de l'interface web

> Tour d'horizon exhaustif de l'interface web d'administration OPNsense :
> chaque menu, à quoi il sert, et comment s'en servir en pratique.
>
> **À jour 2025-2026** : vise OPNsense **24.x / 25.x** (thème par défaut,
> MVC/API, WireGuard intégré au cœur, Unbound, Suricata, Zenarmor en
> plugin). Les chemins de menu sont donnés en anglais tels qu'affichés
> dans l'interface.

---

## 1. Accès et prise en main

- URL par défaut : `https://192.168.1.1/` (LAN). Port HTTPS 443
  (modifiable), HTTP redirigé vers HTTPS par défaut.
- Identifiants par défaut : `root` / `opnsense` — **à changer
  immédiatement** (System > Access > Users).
- L'assistant de démarrage (wizard) au premier login : langue, hostname,
  DNS, fuseau horaire, WAN/LAN, mot de passe root.
- En haut à droite : recherche globale (loupe) — tapez « nat » ou
  « wireguard » pour sauter directement à la page. C'est le moyen le
  plus rapide de naviguer.
- **Bouton « Apply »** : la plupart des pages exigent un clic sur
  Apply / Save pour activer. Une pastille orange signale les changements
  en attente.
- **Dashboard** (Lobby) : widgets configurables (trafic, gateways,
  services, firewall logs, mises à jour). Personnalisez via le « + ».

---

## 2. System — réglages généraux

### System > Settings > General
- Hostname, domaine, fuseau horaire (indispensable : une horloge fausse
  casse IPsec, les certificats et les logs).
- Serveurs DNS (ou « Allow DNS server list to be overridden by DHCP »).
- Langue, thème.

### System > Settings > Administration
- Protocole web UI (HTTPS recommandé), port, interface d'écoute
  (**ne jamais exposer la web UI sur WAN**).
- **Secure Shell** : activer ici l'accès SSH (port, auth par clé).
- Console : mot de passe requis ou non.

### System > Settings > Logging
- Niveaux de log par composant, syslog distant (cible + protocole
  TCP/UDP/TLS), rotation.

### System > Settings > Cron
- Tâches planifiées : mises à jour des alias, renouvellement DHCP,
  backups automatiques, redémarrages planifiés.

### System > Firmware
- **Updates** : vérifier / installer les mises à jour (toujours
  sauvegarder la config avant — voir §15).
- **Plugins** : catalogue (FRR, ACME, WireGuard est au cœur depuis
  21.x, haproxy, crowdsec, etc.). Installer = 2 clics.
- **Packages** : paquets FreeBSD sous-jacents (usage avancé).

### System > Certificates
- **Authorities** : créer votre CA interne.
- **Certificates** : certificats serveur (web UI, OpenVPN, IPsec).
- **Revocation** : CRL.
- Indispensable pour : HTTPS de la web UI avec un vrai certificat,
  OpenVPN, IPsec par certificats, inspection TLS.

### System > Trust
- Autorités et certificats de confiance pour valider les pairs distants.

### System > Gateways
- **Single** : passerelles WAN (IP, monitor IP pour le failover).
- **Group** : groupes de failover / load-balancing multi-WAN
  (Tier 1 = prioritaire, Tier 2 = secours ; trigger : packet loss,
  latency, down).
- Une gateway « down » ici = bascule automatique du groupe.

### System > High Availability
- Synchronisation XMLRPC entre 2 nœuds (config sync : firewall rules,
  aliases, users…), mot de passe de sync, interfaces à synchroniser.
- À combiner avec CARP (voir §4).

### System > Access
- **Users / Groups** : créer des comptes (admin, lecture seule…),
  groupes, privilèges fins (page-level privileges).
- **Servers** : LDAP, RADIUS pour l'authentification centralisée.
- **2FA** : TOTP pour les comptes (recommandé pour root/admin).
- **Sudo** : déléguer des commandes shell à des non-root.
- **Tester** : bouton de test d'authentification.

---

## 3. Interfaces

### Interfaces > Assignments
- Associer chaque port physique (`em0`, `igb0`, `vtnet0`…) à un rôle
  (WAN, LAN, OPT1…). Le cœur du « menu 1 » de la console, en web.

### Interfaces > [WAN] / [LAN] / [OPTn]
- **IPv4/IPv6 Configuration Type** : DHCP, Static, PPPoE, PPTP, L2TP…
- IP statique + passerelle, MTU/MSS, « Block private networks »
  et « Block bogon networks » sur WAN (à laisser cochés).
- **DHCP client options** : hostname, advanced options.

### Interfaces > Other Types
- **VLAN** : créer (parent + tag 802.1Q + description).
- **LAGG** : agrégat LACP (`laggproto lacp`), membres.
- **Bridge** : pont entre interfaces.
- **GRE / GIF** : tunnels.
- **VXLAN** (selon version/plugin).

### Interfaces > Overview
- Vue d'ensemble : état, IP, MAC, trafic, passerelle — le pendant web
  de `ifconfig`.

### Interfaces > Diagnostics
- **ARP Table**, **NDP Table** : qui est joignable en local.
- **Routes** : table de routage lue en direct.

---

## 4. Firewall — règles

### Firewall > Rules > [interface]
- Le cœur du filtrage. Chaque règle : Action (Pass/Block/Reject),
  Disabled, Interface, Direction (in/out), TCP/IP version, Protocol,
  Source, Destination, Destination port range, Log, Description.
- **L'ordre compte** : la première règle qui matche gagne
  (les règles flottantes « Floating » s'évaluent avant, selon
  « quick »).
- **Schedules** : plages horaires (Firewall > Schedules) pour des
  règles actives seulement à certaines heures.
- **Log** : cochez sur les règles sensibles — les hits apparaissent
  dans Firewall > Log Files > Live View.
- Règles automatiques : « Automatically generated rules » (anti-lockout
  LAN, DHCP, etc.) visibles en lecture seule.

### Firewall > Rules > Floating
- Règles transverses (multi-interfaces), idéales pour : bloquer des
  pays/IP partout, QoS, politique sortante globale.

### Firewall > Aliases
- Listes nommées : **Hosts** (IP), **Networks**, **Ports**, **URLs**,
  **URL Tables** (listes téléchargées : blocklists), **GeoIP**
  (pays via MaxMind — nécessite une clé gratuite).
- Utilisées dans Source/Destination des règles et du NAT.
- **URL Table** avec rafraîchissement auto = blocklist dynamique
  (ex. : liste d'IP malveillantes mise à jour chaque jour).

### Firewall > NAT
- **Port Forward** : `Interface WAN → Destination WAN address →
  Redirect target IP → Redirect target port` + règle firewall auto
  (créez-la en même temps, option « Associated filter rule »).
- **Outbound** : Automatic (défaut, PAT sur IP WAN), Hybrid, Manual.
  Manual = contrôle fin (ex. : exclure le trafic VPN du NAT).
- **One-to-One (1:1)** : mapper toute une IP publique vers une IP
  interne (DMZ).
- **NPTv6** : équivalent NAT pour IPv6.

### Firewall > Virtual IPs
- **IP Alias / CARP / Proxy ARP / Other** : IP supplémentaires sur une
  interface (multi-IP publiques du FAI, VIP CARP pour la HA).
- **CARP** : Virtual IP partagée entre 2 nœuds + advskew (priorité).
  Status : Firewall > Virtual IPs > Status.

### Firewall > Shaper
- **Pipes / Queues / Rules** (dummynet) : limiter/garantir la bande
  passante par IP, par service, par interface. Ex. : pipe 100 Mbit
  down / 20 Mbit up sur WAN, queues par VLAN.

### Firewall > Log Files
- **Live View** : le trafic filtré en temps réel (filtres par
  interface, action, IP, port). L'outil n°1 du diagnostic.
- **Plain View** : historique paginé.
- **Summary** : top parleurs / top bloqués.

### Firewall > Diagnostics
- **States** : table des états en direct (recherche par IP, kill).
- **States Summary**, **Statistics** (`pfctl -s info` en web).
- **pfTables** : contenu des tables (bogons, aliases…).
- **pfInfo**, **Sockets**.

---

## 5. DHCP et DNS

### Services > DHCPv4 > [interface]
- **Enable**, Range (plage), passerelle/DNS par défaut ou custom
  (DNS servers, Domain name).
- **Static Mappings** : réservations par MAC (serveurs, imprimantes).
- **Deny unknown clients** : durcissement (avec static mappings).
- **Leases** : baux actifs vus en direct.

### Services > DHCPv6 / Router Advertisements
- RA : Managed / Assisted / Stateless / Disabled selon votre
  stratégie IPv6 (DHCPv6 + RA combinés en général).

