---
id: collect-261001-cisco/cisco/deploiement-entreprise-cisco-huawei-1
title: "Déploiement entreprise — Cisco & Huawei (référence complète)"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["datacenter", "distribution"]
source: docs/RAG/collect-261001-cisco/deploiement_entreprise_cisco_huawei.md
source_anchor: ""
source_lines: [1, 132]
sha256: 1104315ef3d84320545139c3ab95d6b501c31e6b0e60b259ac78a5b94c3bf4f5
---

# Déploiement entreprise — Cisco & Huawei (référence complète)

> Référence de déploiement réseau d'entreprise couvrant **tous les
> aspects** : architecture, adressage, configuration complète par
> couche (core / distribution / accès), WAN, haute disponibilité,
> sécurité, QoS, supervision, interopérabilité Cisco ↔ Huawei,
> recette et mise en production.
>
> **Conventions** : `CISCO>` = IOS/IOS XE, `HUAWEI>` = VRP (Comware-like
> des switchs S / routeurs AR, syntaxe vue dans vos Command Reference).
> Les exemples forment un tout cohérent : même plan d'adressage des
> deux côtés pour faciliter la comparaison et l'interopérabilité.

---

## 1. Architecture de référence

```
                    ┌──────────────┐      ┌──────────────┐
                    │   FAI-1      │      │   FAI-2      │
                    │ AS 100       │      │ AS 200       │
                    └──────┬───────┘      └──────┬───────┘
                           │ eBGP               │ eBGP
                    ┌──────┴────────────────────┴──────┐
                    │        EDGE (routeurs)           │
                    │   AS 65001 — BGP + NAT + IPsec   │
                    └──────┬────────────────────┬──────┘
                           │                    │
                    ┌──────┴──────┐      ┌───────┴──────┐
                    │  FIREWALL   │      │  DMZ         │
                    │  (USG/pair) │      │  serveurs    │
                    └──────┬──────┘      │  publics     │
                           │            └──────────────┘
              ┌────────────┴────────────┐
              │      CORE (L3)          │  2 switchs, VRRP/HSRP,
              │  10.0.0.0/30 transit   │  OSPF area 0
              └────┬──────────────┬─────┘
                   │ Eth-Trunk/LACP      │ Eth-Trunk/LACP
        ┌──────────┴────────┐  ┌─────────┴─────────┐
        │ DISTRIBUTION-BAT-A │  │ DISTRIBUTION-BAT-B │  routage inter-VLAN,
        │  OSPF area 0       │  │  OSPF area 0        │  STP, DHCP snooping
        └────┬──────────┬───┘  └────┬──────────┬───┘
             │          │           │          │
        ┌────┴───┐ ┌────┴───┐  ┌────┴───┐ ┌────┴───┐
        │ ACCESS │ │ ACCESS │  │ ACCESS │ │  WiFi  │  port-security,
        │  Sw-1  │ │  Sw-2  │  │  Sw-3  │ │  (AP)  │  802.1X, PoE
        └────────┘ └────────┘  └────────┘ └────────┘
```

- **Core** : commutation L3 rapide, aucune ACL lourde, redondance
  maximale. Jamais de postes utilisateurs en direct.
- **Distribution** : frontière L2/L3 (passerelles VLAN), politiques
  (ACL, QoS), agrégation des switchs d'accès.
- **Accès** : ports utilisateurs, PoE, sécurité de port.
- **Edge/WAN** : BGP vers 2 FAIs, NAT, VPN IPsec/DMVPN vers les sites.
- **DMZ** : serveurs publics derrière le firewall, jamais routés
  directement vers le LAN.

---

## 2. Plan d'adressage et conventions

| Usage | Plage | Exemple |
|---|---|---|
| Loopbacks (router-id, management) | 10.255.255.0/24 | CORE-1 = .1, CORE-2 = .2 |
| Transits point-à-point | 10.0.0.0/24 en /30 | CORE-1↔EDGE-1 = 10.0.0.0/30 |
| Management (OOB ou VLAN 99) | 192.168.99.0/24 | — |
| VLAN 10 DATA | 192.168.10.0/24 | GW .254 (HSRP/VRRP) |
| VLAN 20 VOIX | 192.168.20.0/24 | GW .254 |
| VLAN 30 INVITES | 192.168.30.0/24 | GW .254 |
| VLAN 40 SERVEURS | 192.168.40.0/24 | GW .254 |
| VLAN 99 MGMT | 192.168.99.0/24 | GW .254 |
| WAN publics | bloc FAI | — |

**Conventions de nommage** : `CORE-01`, `DIST-A-01`, `ACC-A-01-03`
(bâtiment A, local 01, switch 03). Descriptions systématiques sur les
interfaces (`description UPLINK_DIST-A-01`).

---

## 3. Socle management (les deux vendors)

À déployer **en premier** sur chaque équipement : sans NTP/SNMP/syslog/
AAA homogènes, pas de supervision ni d'audit possibles.

### 3.1 Cisco — socle

```shell
CISCO> enable
CISCO# configure terminal
CISCO(config)# hostname CORE-01
CISCO(config)# no ip domain-lookup
CISCO(config)# ip domain-name entreprise.lan
CISCO(config)# service password-encryption
CISCO(config)# service timestamps log datetime msec
CISCO(config)# service timestamps debug datetime msec
! --- NTP ---
CISCO(config)# ntp server 192.168.99.10 prefer
CISCO(config)# ntp server 192.168.99.11
CISCO(config)# clock timezone UTC 0
! --- SSH ---
CISCO(config)# username admin privilege 15 secret AdminFort2026
CISCO(config)# crypto key generate rsa modulus 2048
CISCO(config)# ip ssh version 2
CISCO(config)# line vty 0 15
CISCO(config-line)# login local
CISCO(config-line)# transport input ssh
CISCO(config-line)# exec-timeout 10 0
CISCO(config-line)# exit
! --- Syslog ---
CISCO(config)# logging host 192.168.99.50
CISCO(config)# logging trap informational
CISCO(config)# logging source-interface Loopback 0
CISCO(config)# logging buffered 32768
! --- SNMPv3 ---
CISCO(config)# snmp-server group NOC v3 priv
CISCO(config)# snmp-server user zabbix NOC v3 auth sha AuthPass priv aes 256 PrivPass
CISCO(config)# snmp-server host 192.168.99.50 version 3 priv zabbix
CISCO(config)# snmp-server location DATACENTER-RangA12
CISCO(config)# snmp-server contact noc@entreprise.lan
! --- AAA TACACS+ ---
CISCO(config)# aaa new-model
CISCO(config)# tacacs-server host 192.168.99.60 key CleTACACS
CISCO(config)# aaa authentication login default group tacacs+ local
CISCO(config)# aaa authorization exec default group tacacs+ local
CISCO(config)# aaa accounting exec default start-stop group tacacs+
CISCO(config)# end
CISCO# copy running-config startup-config
```

### 3.2 Huawei — socle

