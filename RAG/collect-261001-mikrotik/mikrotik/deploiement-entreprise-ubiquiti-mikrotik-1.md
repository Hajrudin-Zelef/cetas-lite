---
id: collect-261001-mikrotik/mikrotik/deploiement-entreprise-ubiquiti-mikrotik-1
title: "Déploiement entreprise — Ubiquiti & MikroTik (référence complète)"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "voice"]
source: docs/RAG/collect-261001-mikrotik/deploiement_entreprise_ubiquiti_mikrotik.md
source_anchor: ""
source_lines: [1, 185]
sha256: 2fde9ccdd114be2710f330017e4ba496493e8d6a31276f670e875c85df6d748b
---

# Déploiement entreprise — Ubiquiti & MikroTik (référence complète)

> Référence de déploiement réseau d'entreprise pour **Ubiquiti**
> (UniFi + EdgeMAX/EdgeRouter) et **MikroTik** (RouterOS) : architecture,
> plan d'adressage, configurations complètes, WAN, VPN, WiFi, QoS,
> supervision, interopérabilité, recette et mise en production.
>
> **Conventions** : `UBNT>` = EdgeOS (EdgeRouter, CLI style Vyatta),
> `MT>` = RouterOS (CLI `/...`). UniFi est piloté par **contrôleur**
> (pas de CLI significatif) : les étapes sont décrites en parcours
> d'interface. Versions visées : **EdgeOS 2.x**, **RouterOS v7.x**,
> **UniFi OS / Network 8.x-9.x**.

---

## 1. Positionnement des gammes

| Besoin | Ubiquiti | MikroTik |
|---|---|---|
| Routeur principal / edge | EdgeRouter (EdgeOS CLI) | CCR / RB5009 (RouterOS) |
| Switching managé | UniFi Switch (contrôleur) | CRS (RouterOS, SwitchOS) |
| WiFi entreprise | UniFi AP (contrôleur) | cAP / CAPsMAN (RouterOS) |
| Site simple tout-en-un | UniFi Gateway (UXG/UDM) | hAP ax / RB4011 |
| VPN multi-sites | EdgeRouter / UniFi | RouterOS (WireGuard natif) |

- **UniFi** = écosystème contrôleur : adoption, provisionnement central,
  superbe pour le WiFi et les sites simples. Le CLI y est marginal.
- **EdgeRouter** = vrai routeur CLI (EdgeOS, base Vyatta) : OSPF/BGP,
  firewall zones, QoS — le « Cisco du pauvre » en CLI.
- **MikroTik** = tout en CLI RouterOS (`/ip firewall`, `/routing`…) :
  le plus flexible et le moins cher, scripting puissant, courbe
  d'apprentissage raide.

---

## 2. Architecture de référence

```
                    ┌──────────────┐      ┌──────────────┐
                    │   FAI-1      │      │   FAI-2      │
                    └──────┬───────┘      └──────┬───────┘
                           │                    │
              ┌────────────┴────────────────────┴────────────┐
              │  EDGE : EdgeRouter-12 / CCR2116 (VRRP .253)  │
              │  NAT, firewall, BGP ou failover, VPN         │
              └────────────┬────────────────────┬────────────┘
                           │ trunk              │
              ┌────────────┴────────┐  ┌────────┴───────────┐
              │ SWITCH CORE/ACCÈS   │  │  WiFi              │
              │ UniFi Switch Pro 48 │  │  UniFi AP / cAP    │
              │ ou CRS518           │  │  (CAPsMAN ou       │
              │ VLANs 10/20/30/99   │  │   contrôleur UniFi)│
              └─────────────────────┘  └────────────────────┘
```

## Plan d'adressage (commun aux deux vendors)

| Usage | Plage |
|---|---|
| Loopbacks / router-id | 10.255.255.0/24 |
| WAN1 / WAN2 | bloc FAI |
| VLAN 10 DATA | 192.168.10.0/24, GW .254 |
| VLAN 20 VOIX | 192.168.20.0/24, GW .254 |
| VLAN 30 INVITES | 192.168.30.0/24, GW .254 |
| VLAN 40 SERVEURS | 192.168.40.0/24, GW .254 |
| VLAN 99 MGMT | 192.168.99.0/24, GW .254 |
| VPN WireGuard | 10.8.0.0/24 |

---

## 3. Socle management

### 3.1 EdgeRouter — socle (EdgeOS)

```shell
UBNT> configure
UBNT# set system host-name ER-EDGE-01
UBNT# set system domain-name entreprise.lan
UBNT# set system login user admin authentication plaintext-password AdminFort2026
UBNT# set system login user admin level admin
UBNT# set service ssh port 22
UBNT# set service ssh disable-password-authentication
UBNT# set system ntp server 192.168.99.10
UBNT# set system ntp server 192.168.99.11
UBNT# set system syslog host 192.168.99.50 facility all level info
UBNT# set system time-zone UTC
UBNT# set snmp community NOC authorization ro
UBNT# set snmp contact noc@entreprise.lan
UBNT# set snmp location DATACENTER-RangA12
UBNT# commit
UBNT# save
UBNT# exit
UBNT> show version
```

### 3.2 MikroTik — socle (RouterOS v7)

```shell
MT> /system identity set name=MT-EDGE-01
MT> /user add name=admin group=full password=AdminFort2026
MT> /user disable admin
MT> /ip service disable telnet,ftp,www,api
MT> /ip service set ssh port=22
MT> /ip service set winbox address=192.168.99.0/24
MT> /system ntp client set enabled=yes
MT> /system ntp client servers add address=192.168.99.10
MT> /system ntp client servers add address=192.168.99.11
MT> /system clock set time-zone-name=UTC
MT> /system logging action add name=SIEM target=remote remote=192.168.99.50
MT> /system logging add topics=info action=SIEM
MT> /snmp set enabled=yes contact="noc@entreprise.lan" location="DATACENTER-RangA12"
MT> /snmp community set name=public addresses=192.168.99.50/32 security=authorized read-access=yes
MT> /export file=socle-avant-prod
MT> /system backup save name=socle-avant-prod
```

> **Premier réflexe MikroTik** : `/export` (config lisible) + `/system
> backup save` (binaire). Les deux, à chaque étape qui marche.

---

## 4. Ubiquiti UniFi — déploiement via contrôleur

### 4.1 Prise en main
1. Installer le contrôleur : **UniFi OS** (Cloud Key / UDM / UXG) ou
   self-hosted (Docker/serveur).
2. Adopter les équipements : Devices > **Adopt** (ils doivent être en
   usine et joignables en L2/L3 ; sinon adoption manuelle par SSH :
   `set-inform http://<controleur>:8080/inform`).
3. **Settings > System** : backup auto du contrôleur, mises à jour
   planifiées (jamais en heures ouvrées).

### 4.2 Réseaux (VLANs)
4. **Settings > Networks > Create New** :
   - `DATA` : VLAN 10, `192.168.10.1/24`, DHCP `192.168.10.100–200`.
   - `VOIX` : VLAN 20, `192.168.20.1/24`.
   - `INVITES` : VLAN 30, `192.168.30.1/24` (+ **Guest isolation** /
     portail captif si besoin : Settings > Hotspot).
   - `SERVEURS` : VLAN 40, `192.168.40.1/24`, DHCP avec réservations.
   - `MGMT` : VLAN 99, `192.168.99.1/24`.
5. **Devices > Switch > Ports** : profils de ports — `All` (trunk) sur
   les uplinks, `DATA` (access) sur les ports utilisateurs, `VOIX`
   en voice VLAN sur les ports téléphonie.

### 4.3 WiFi
6. **Settings > WiFi > Create New** :
   - `CORP` : WPA2/3 Enterprise (RADIUS vers NPS/FreeRADIUS) ou PSK
     fort, VLAN 10, bandes 5 GHz優先.
   - `INVITES` : portail captif + VLAN 30 + **client isolation** +
     limite de débit.
7. **Radio management** : laissez l'optimisation auto (canaux, puissance)
   sauf site dense → plan manuel (1/6/11 en 2.4 GHz).

### 4.4 Firewall (UniFi Gateway)
8. **Settings > Security > Firewall Rules** :
   - LAN IN : autoriser DATA→SERVEURS (ports utiles), **refuser
     INVITES→RFC1918** (règle explicite, placée avant l'autorisation
     Internet).
   - GUEST IN : règles du portail invité (auto).
   - WAN IN : **rien** sauf VPN (WireGuard/OpenVPN) et port-forwards
     nécessaires.
9. **Port Forwarding** : WAN 443 → serveur DMZ (avec règle WAN IN
   associée auto).
10. **Traffic & Firewall Rules > Ad Blocking / GeoIP filtering** :
    activer selon politique.

### 4.5 VPN (UniFi)
11. **Settings > VPN** :
    - **WireGuard** : serveur, pairs (télétravailleurs), AllowedIPs.
    - **OpenVPN** : serveur road-warrior (fichiers `.ovpn` exportables).
    - **Site-to-Site IPsec** : vers les agences (PSK ou certificats),
      statut dans **Insights > VPN**.
12. **Teleport / Identity** (selon licence) : accès distant simplifié.

### 4.6 Supervision UniFi
- **Insights** : clients, débits, DPI (applications), alertes.
- **Alerts** : configurer les notifications (gateway down, rogue AP).
- Logs : **System Logs** filtrables par équipement.

---

## 5. EdgeRouter — déploiement complet en CLI

### 5.1 Interfaces, VLANs, bonding

