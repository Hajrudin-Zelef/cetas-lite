---
id: collect-261001-rattrapage/rattrapage/cisco-cli-guide-1
title: "Cisco — Guide CLI complet (IOS / IOS XE)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["cost", "datacenter", "memory"]
source: docs/RAG/collect-261001-rattrapage/cisco_cli_guide.md
source_anchor: ""
source_lines: [1, 248]
sha256: c63634dea35d2eb1e83520302d0cd427b9f21c14d12b25f82899b21ff7f9356b
---

# Cisco — Guide CLI complet (IOS / IOS XE)

> Synthèse rédigée à partir de la documentation officielle Cisco
> (Command Reference IOS/IOS XE). Couvre l'essentiel du CLI : accès,
> switching, routage, NAT, VPN, QoS, HA, sécurité, supervision et
> maintenance. Focus IOS/IOS XE (routeurs ISR/ASR, switchs Catalyst) ;
> les différences IOS XR / NX-OS sont signalées en fin de document.
>
> **À jour 2025-2026** : les exemples visent **IOS XE 17.x** (17.9 / 17.12 /
> 17.15, versions recommandées actuelles). IOS 15.x est en fin de vie :
> certaines commandes legacy (crypto map IKEv1, PAgP, RIPv1) y existent
> encore mais sont dépréciées. Algorithmes faibles (DES, 3DES, MD5,
> DH groupes 1/2/5) bannis des exemples — voir §13.
> Vérifiez les variantes selon votre version exacte (`show version`).

---

## 1. Gammes et concepts

- **IOS / IOS XE** : routeurs ISR/ASR, switchs Catalyst. CLI historique,
  le plus répandu.
- **IOS XE 17.x** (2025) : ISR 4000, Catalyst 8000 (edge), Catalyst
  9000 (9300/9400/9500/9600, switching). C'est la cible des exemples
  de ce guide.
- **IOS XR** : routeurs opérateurs (ASR 9000, NCS). Syntaxe proche mais
  avec `commit` obligatoire après chaque changement.
- **NX-OS** : switchs datacenter Nexus. Syntaxe proche d'IOS avec
  particularités (`feature` à activer : `feature ospf`, `feature bgp`…).
- Ce guide couvre **IOS/IOS XE**, la base commune (~90 % du CLI).
- Modes : `User EXEC` (`Router>`) → `Privileged EXEC` (`Router#`) →
  `Global config` (`Router(config)#`) → sous-modes
  (`config-if`, `config-router`, `config-line`…).
- `show` = afficher, `no` = annuler (équivalent du `undo` Huawei).
- La config active est en RAM : **`copy running-config startup-config`**
  (ou `write memory`) pour sauvegarder, sinon tout est perdu au reboot.

---

## 2. Premier accès et configuration de base

- Console : 9600 bauds, 8N1. Pas de mot de passe par défaut sur la
  plupart des équipements récents (configuration initiale via le
  dialogue `setup`).
- SSH/HTTPS : à configurer (voir §12).

```shell
Router> enable
Router# configure terminal
Router(config)# hostname R-SIEGE
R-SIEGE(config)# no ip domain-lookup
R-SIEGE(config)# ip domain-name entreprise.lan
R-SIEGE(config)# service password-encryption
R-SIEGE(config)# banner motd # Acces reserve #
R-SIEGE(config)# end
R-SIEGE# show version
R-SIEGE# show running-config
```

---

## 3. Commandes fondamentales

```shell
enable / disable              # monter / descendre de niveau
configure terminal            # passer en configuration (conf t)
end                           # retour en mode privilégié (Ctrl+Z)
exit                          # remonte d'un niveau
show running-config           # config active
show startup-config           # config sauvegardée
copy running-config startup-config   # sauvegarder (write memory)
show ?                        # aide contextuelle partout
do show ip interface brief    # exécuter un show depuis le mode config
```

- Complétion avec `Tab`, rappel avec flèches haut/bas.
- `| include`, `| exclude`, `| begin`, `| section` pour filtrer les
  sorties : `show running-config | include interface`.
- `terminal length 0` : désactive la pagination (utile en script).

---

## 4. Interfaces

```shell
R-SIEGE(config)# interface GigabitEthernet 0/0
R-SIEGE(config-if)# description LIEN_FAI
R-SIEGE(config-if)# ip address 202.10.10.10 255.255.255.252
R-SIEGE(config-if)# no shutdown
R-SIEGE(config-if)# duplex full
R-SIEGE(config-if)# speed 1000
R-SIEGE(config-if)# exit

# Sous-interface 802.1Q (router-on-a-stick)
R-SIEGE(config)# interface GigabitEthernet 0/1.10
R-SIEGE(config-subif)# encapsulation dot1Q 10
R-SIEGE(config-subif)# ip address 192.168.10.1 255.255.255.0
R-SIEGE(config-subif)# exit

# Loopback (toujours UP, idéale pour router-id / management)
R-SIEGE(config)# interface Loopback 0
R-SIEGE(config-if)# ip address 10.255.255.1 255.255.255.255

# Vérifications
R-SIEGE# show ip interface brief
R-SIEGE# show interfaces GigabitEthernet 0/0
R-SIEGE# show interfaces status        # switchs
R-SIEGE# show interfaces description
```

### EtherChannel (agrégat)

```shell
R-SIEGE(config)# interface range GigabitEthernet 0/2 - 3
R-SIEGE(config-if-range)# channel-group 1 mode active    # LACP
R-SIEGE(config-if-range)# exit
R-SIEGE(config)# interface Port-channel 1
R-SIEGE(config-if)# ip address 10.0.0.1 255.255.255.252
R-SIEGE# show etherchannel summary
```

- Modes : `active`/`passive` (LACP), `desirable`/`auto` (PAgP), `on`
  (statique).

---

## 5. Switching : VLAN, trunk, STP

```shell
# VLANs
SW(config)# vlan 10
SW(config-vlan)# name COMPTA
SW(config-vlan)# exit
SW(config)# vlan 20
SW(config-vlan)# name INVITES
SW(config-vlan)# exit

# Port d'accès
SW(config)# interface GigabitEthernet 1/0/5
SW(config-if)# switchport mode access
SW(config-if)# switchport access vlan 10
SW(config-if)# spanning-tree portfast       # edge port
SW(config-if)# exit

# Trunk 802.1Q
SW(config)# interface GigabitEthernet 1/0/24
SW(config-if)# switchport mode trunk
SW(config-if)# switchport trunk allowed vlan 10,20
SW(config-if)# switchport trunk native vlan 99
SW(config-if)# exit

# Spanning Tree
SW(config)# spanning-tree mode rapid-pvst
SW(config)# spanning-tree vlan 10 priority 4096   # root primaire

# Vérifications
SW# show vlan brief
SW# show interfaces trunk
SW# show spanning-tree vlan 10
SW# show mac address-table
```

- `switchport mode dynamic desirable/auto` : négociation DTP (à éviter,
  figez en access/trunk).
- VTP : `vtp mode transparent` recommandé (évite les effacements de
  VLAN accidentels en `server`/`client`).

---

## 6. Adressage, DHCP, DNS, NTP

```shell
# DHCP serveur
R(config)# ip dhcp excluded-address 192.168.1.1 192.168.1.10
R(config)# ip dhcp pool LAN
R(dhcp-config)# network 192.168.1.0 255.255.255.0
R(dhcp-config)# default-router 192.168.1.1
R(dhcp-config)# dns-server 8.8.8.8 1.1.1.1
R(dhcp-config)# lease 3
R(dhcp-config)# exit
R# show ip dhcp binding
R# show ip dhcp pool

# DHCP client (interface WAN)
R(config)# interface GigabitEthernet 0/0
R(config-if)# ip address dhcp

# DNS / NTP
R(config)# ip name-server 8.8.8.8
R(config)# ntp server 192.168.1.100
R# show ntp status
R# ping www.google.com     # teste aussi la résolution DNS
```

---

## 7. Routage statique et par politique

```shell
R(config)# ip route 0.0.0.0 0.0.0.0 202.10.10.9
R(config)# ip route 192.168.2.0 255.255.255.0 10.0.0.2 10   # distance admin
R(config)# ipv6 route ::/0 2001:db8::1
R# show ip route
R# show ip route 8.8.8.8        # quelle route pour cette destination ?
R# show ip cef 8.8.8.8          # forwarding exact (CEF)

# PBR (Policy-Based Routing)
R(config)# access-list 100 permit ip 192.168.1.0 0.0.0.255 any
R(config)# route-map PBR1 permit 10
R(config-route-map)# match ip address 100
R(config-route-map)# set ip next-hop 10.0.0.254
R(config-route-map)# exit
R(config)# interface GigabitEthernet 0/1
R(config-if)# ip policy route-map PBR1
```

---

## 8. OSPF

```shell
R(config)# router ospf 1
R(config-router)# router-id 10.255.255.1
R(config-router)# network 192.168.1.0 0.0.0.255 area 0
R(config-router)# network 10.0.0.0 0.0.0.3 area 0
R(config-router)# passive-interface default
R(config-router)# no passive-interface GigabitEthernet 0/2
R(config-router)# exit

# Authentification MD5
R(config)# interface GigabitEthernet 0/2
R(config-if)# ip ospf authentication message-digest
R(config-if)# ip ospf message-digest-key 1 md5 CleOSPF
R(config-if)# ip ospf cost 10
R(config-if)# ip ospf priority 100

R# show ip ospf neighbor
R# show ip ospf database
R# show ip ospf interface brief
R# show ip route ospf
```

- États : Down → Init → 2-Way → ExStart → Exchange → Loading → **Full**.
- Reste en ExStart/Exchange = MTU mismatch ou duplicate router-id.

---

## 9. EIGRP et BGP

