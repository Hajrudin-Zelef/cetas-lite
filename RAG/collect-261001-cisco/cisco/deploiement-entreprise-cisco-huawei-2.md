---
id: collect-261001-cisco/cisco/deploiement-entreprise-cisco-huawei-2
title: "Déploiement entreprise — Cisco & Huawei (référence complète)"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent", "datacenter", "distribution", "voice"]
source: docs/RAG/collect-261001-cisco/deploiement_entreprise_cisco_huawei.md
source_anchor: ""
source_lines: [133, 338]
sha256: fa5591e8ffd63681457372791b8d1fde7eacfa10b8fd1f9e7d4b093efe10404f
---

# Déploiement entreprise — Cisco & Huawei (référence complète)

```shell
HUAWEI> system-view
HUAWEI] sysname CORE-01
HUAWEI] undo telnet server enable
HUAWEI] ssh server enable
HUAWEI] dsa local-key-pair create
! --- utilisateur admin ---
HUAWEI] aaa
HUAWEI-aaa] local-user admin password irreversible-cipher AdminFort2026
HUAWEI-aaa] local-user admin privilege level 15
HUAWEI-aaa] local-user admin service-type ssh terminal
HUAWEI-aaa] quit
HUAWEI] user-interface vty 0 14
HUAWEI-ui-vty] authentication-mode aaa
HUAWEI-ui-vty] protocol inbound ssh
HUAWEI-ui-vty] idle-timeout 10 0
HUAWEI-ui-vty] quit
! --- NTP ---
HUAWEI] ntp-service unicast-server 192.168.99.10 preference
HUAWEI] ntp-service unicast-server 192.168.99.11
HUAWEI] clock timezone UTC add 00:00:00
! --- Syslog ---
HUAWEI] info-center loghost 192.168.99.50
HUAWEI] info-center source default channel loghost log level informational
HUAWEI] info-center loghost source LoopBack 0
! --- SNMPv3 ---
HUAWEI] snmp-agent
HUAWEI] snmp-agent sys-info location DATACENTER-RangA12
HUAWEI] snmp-agent sys-info contact noc@entreprise.lan
HUAWEI] snmp-agent group v3 NOC privacy
HUAWEI] snmp-agent local-user zabbix group NOC
HUAWEI] snmp-agent local-user zabbix auth-mode sha AuthPass
HUAWEI] snmp-agent local-user zabbix priv-mode aes256 PrivPass
HUAWEI] snmp-agent target-host trap address udp-domain 192.168.99.50 params securityname zabbix v3 privacy
! --- HWTACACS ---
HUAWEI] hwtacacs-server template TAC
HUAWEI-hwtacacs] hwtacacs-server authentication 192.168.99.60
HUAWEI-hwtacacs] hwtacacs-server authorization 192.168.99.60
HUAWEI-hwtacacs] hwtacacs-server accounting 192.168.99.60
HUAWEI-hwtacacs] hwtacacs-server shared-key cipher CleTACACS
HUAWEI-hwtacacs] quit
HUAWEI] aaa
HUAWEI-aaa] authentication-scheme default
HUAWEI-aaa-authen] authentication-mode hwtacacs local
HUAWEI-aaa-authen] quit
HUAWEI-aaa] authorization-scheme default
HUAWEI-aaa-author] authorization-mode hwtacacs local
HUAWEI-aaa-author] quit
HUAWEI-aaa] accounting-scheme default
HUAWEI-aaa-accounting] accounting-mode hwtacacs
HUAWEI-aaa-accounting] quit
HUAWEI-aaa] quit
HUAWEI] domain default
HUAWEI-domain] hwtacacs-server TAC
HUAWEI-domain] quit
HUAWEI] quit
HUAWEI> save
```

---

## 4. Cisco — déploiement par couche

### 4.1 Core (Catalyst 9500 × 2)

```shell
! --- VLANs L3 (SVI) ---
CISCO(config)# vlan 10
CISCO(config-vlan)# name DATA
CISCO(config-vlan)# exit
CISCO(config)# vlan 20
CISCO(config-vlan)# name VOIX
CISCO(config-vlan)# exit
CISCO(config)# vlan 40
CISCO(config-vlan)# name SERVEURS
CISCO(config-vlan)# exit
CISCO(config)# vlan 99
CISCO(config-vlan)# name MGMT
CISCO(config-vlan)# exit
!
! --- SVI passerelles + HSRP (CORE-01 actif) ---
CISCO(config)# interface Vlan 10
CISCO(config-if)# ip address 192.168.10.2 255.255.255.0
CISCO(config-if)# standby 10 ip 192.168.10.254
CISCO(config-if)# standby 10 priority 110
CISCO(config-if)# standby 10 preempt
CISCO(config-if)# no shutdown
CISCO(config-if)# exit
! (répéter pour Vlan 20/40/99 — CORE-02 en priority 100, preempt)
!
! --- Uplinks L3 vers distribution (point-à-point, pas de trunk au core) ---
CISCO(config)# interface range TenGigabitEthernet 1/0/1 - 2
CISCO(config-if-range)# no switchport
CISCO(config-if-range)# ip address 10.0.0.1 255.255.255.252
CISCO(config-if-range)# no shutdown
!
! --- OSPF backbone ---
CISCO(config)# router ospf 1
CISCO(config-router)# router-id 10.255.255.1
CISCO(config-router)# passive-interface default
CISCO(config-router)# no passive-interface TenGigabitEthernet 1/0/1
CISCO(config-router)# no passive-interface TenGigabitEthernet 1/0/2
CISCO(config-router)# network 10.0.0.0 0.0.0.255 area 0
CISCO(config-router)# network 192.168.0.0 0.0.255.255 area 0
CISCO(config-router)# exit
!
! --- Loopback ---
CISCO(config)# interface Loopback 0
CISCO(config-if)# ip address 10.255.255.1 255.255.255.255
!
CISCO# show ip ospf neighbor
CISCO# show standby brief
```

### 4.2 Distribution (Catalyst 9300)

```shell
! --- Trunks vers core (LACP) ---
CISCO(config)# interface range TenGigabitEthernet 1/1/1 - 2
CISCO(config-if-range)# switchport mode trunk
CISCO(config-if-range)# switchport trunk allowed vlan 10,20,30,40,99
CISCO(config-if-range)# channel-group 10 mode active
CISCO(config-if-range)# exit
CISCO(config)# interface Port-channel 10
CISCO(config-if)# switchport mode trunk
CISCO(config-if)# switchport trunk allowed vlan 10,20,30,40,99
!
! --- STP : distribution = root secondaire ---
CISCO(config)# spanning-tree mode rapid-pvst
CISCO(config)# spanning-tree vlan 10,20,30,40,99 priority 8192
! (core en 4096 = root primaire)
!
! --- Routage inter-VLAN ici OU au core (choisir UNE frontière L3) ---
! Option B : SVI en distribution (recommandé si core = pur L3 sans SVI)
! (même syntaxe SVI + HSRP que §4.1)
!
! --- DHCP snooping + DAI ---
CISCO(config)# ip dhcp snooping
CISCO(config)# ip dhcp snooping vlan 10,20,30,40
CISCO(config)# interface Port-channel 10
CISCO(config-if)# ip dhcp snooping trust
CISCO(config-if)# exit
CISCO(config)# ip arp inspection vlan 10,20,30,40
CISCO(config)# interface Port-channel 10
CISCO(config-if)# ip arp inspection trust
!
! --- OSPF ---
CISCO(config)# router ospf 1
CISCO(config-router)# router-id 10.255.255.11
CISCO(config-router)# network 10.0.0.0 0.0.0.255 area 0
CISCO(config-router)# network 192.168.0.0 0.0.255.255 area 0
CISCO(config-router)# passive-interface default
CISCO(config-router)# no passive-interface Port-channel 10
```

### 4.3 Accès (Catalyst 9200/9300, PoE)

```shell
! --- Template de port utilisateur ---
CISCO(config)# interface range GigabitEthernet 1/0/1 - 44
CISCO(config-if-range)# switchport mode access
CISCO(config-if-range)# switchport access vlan 10
CISCO(config-if-range)# spanning-tree portfast
CISCO(config-if-range)# spanning-tree bpduguard enable
CISCO(config-if-range)# switchport port-security
CISCO(config-if-range)# switchport port-security maximum 3
CISCO(config-if-range)# switchport port-security violation restrict
CISCO(config-if-range)# switchport port-security mac-address sticky
CISCO(config-if-range)# switchport port-security aging time 60
CISCO(config-if-range)# ip verify source              ! IP Source Guard
CISCO(config-if-range)# exit
!
! --- Ports voix (téléphones IP) ---
CISCO(config)# interface range GigabitEthernet 1/0/1 - 24
CISCO(config-if-range)# switchport voice vlan 20
CISCO(config-if-range)# mls qos trust dscp
!
! --- Uplink vers distribution ---
CISCO(config)# interface range GigabitEthernet 1/0/47 - 48
CISCO(config-if-range)# switchport mode trunk
CISCO(config-if-range)# switchport trunk allowed vlan 10,20,30,99
CISCO(config-if-range)# channel-group 1 mode active
CISCO(config-if-range)# ip dhcp snooping trust
CISCO(config-if-range)# ip arp inspection trust
!
! --- 802.1X (optionnel, niveau avancé) ---
CISCO(config)# aaa new-model
CISCO(config)# radius-server host 192.168.99.61 key CleRADIUS
CISCO(config)# aaa authentication dot1x default group radius
CISCO(config)# aaa authorization network default group radius
CISCO(config)# dot1x system-auth-control
CISCO(config)# interface range GigabitEthernet 1/0/1 - 44
CISCO(config-if-range)# authentication port-control auto
CISCO(config-if-range)# dot1x pae authenticator
```

---

## 5. Huawei — déploiement par couche (miroir)

Même architecture, même adressage. Notez les équivalences :
`standby` → `vrrp vrid`, `channel-group` → `eth-trunk`,
`spanning-tree portfast` → `stp edged-port`.

### 5.1 Core (CloudEngine S6730 × 2, ou S5735)

