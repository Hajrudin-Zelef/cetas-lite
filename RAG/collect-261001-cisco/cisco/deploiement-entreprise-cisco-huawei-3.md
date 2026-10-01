---
id: collect-261001-cisco/cisco/deploiement-entreprise-cisco-huawei-3
title: "Déploiement entreprise — Cisco & Huawei (référence complète)"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["distribution", "voice"]
source: docs/RAG/collect-261001-cisco/deploiement_entreprise_cisco_huawei.md
source_anchor: ""
source_lines: [339, 576]
sha256: a635e895f52520bc222e6cf6a0459d9a9c0e94184bbbc954990f0d1a9e277348
---

# Déploiement entreprise — Cisco & Huawei (référence complète)

```shell
HUAWEI> system-view
HUAWEI] sysname CORE-01-HW
! --- VLANs ---
HUAWEI] vlan batch 10 20 40 99
HUAWEI] vlan 10
HUAWEI-vlan10] description DATA
HUAWEI-vlan10] quit
! --- Vlanif + VRRP (CORE-01 = master) ---
HUAWEI] interface Vlanif 10
HUAWEI-Vlanif10] ip address 192.168.10.2 24
HUAWEI-Vlanif10] vrrp vrid 10 virtual-ip 192.168.10.254
HUAWEI-Vlanif10] vrrp vrid 10 priority 120
HUAWEI-Vlanif10] vrrp vrid 10 preempt-mode timer delay 20
HUAWEI-Vlanif10] quit
! (répéter Vlanif 20/40/99 — CORE-02 en priority 100)
!
! --- Uplinks L3 ---
HUAWEI] interface 10GE 1/0/1
HUAWEI-10GE] undo portswitch
HUAWEI-10GE] ip address 10.0.0.1 30
HUAWEI-10GE] undo shutdown
HUAWEI-10GE] quit
!
! --- OSPF ---
HUAWEI] ospf 1 router-id 10.255.255.1
HUAWEI-ospf] area 0
HUAWEI-ospf-area] network 10.0.0.0 0.0.0.255
HUAWEI-ospf-area] network 192.168.0.0 0.0.255.255
HUAWEI-ospf-area] quit
HUAWEI-ospf] silent-interface Vlanif 10
HUAWEI-ospf] silent-interface Vlanif 20
HUAWEI-ospf] silent-interface Vlanif 40
HUAWEI-ospf] silent-interface Vlanif 99
HUAWEI-ospf] quit
!
HUAWEI] interface LoopBack 0
HUAWEI-LoopBack] ip address 10.255.255.1 32
HUAWEI-LoopBack] quit
HUAWEI] quit
HUAWEI> display ospf peer
HUAWEI> display vrrp
```

### 5.2 Distribution (S5735)

```shell
HUAWEI> system-view
! --- Eth-Trunk LACP vers core ---
HUAWEI] interface Eth-Trunk 10
HUAWEI-Eth-Trunk10] mode lacp
HUAWEI-Eth-Trunk10] trunkport 10GE 1/0/1 to 1/0/2
HUAWEI-Eth-Trunk10] port link-type trunk
HUAWEI-Eth-Trunk10] port trunk allow-pass vlan 10 20 30 40 99
HUAWEI-Eth-Trunk10] quit
!
! --- STP : root secondaire ---
HUAWEI] stp mode rstp
HUAWEI] stp priority 8192
! (core en 4096)
HUAWEI] stp instance 0 priority 8192
!
! --- DHCP snooping + anti-attaque ARP ---
HUAWEI] dhcp snooping enable
HUAWEI] vlan 10
HUAWEI-vlan10] dhcp snooping enable
HUAWEI-vlan10] quit
! (répéter vlan 20/30/40)
HUAWEI] interface Eth-Trunk 10
HUAWEI-Eth-Trunk10] dhcp snooping trusted
HUAWEI-Eth-Trunk10] arp anti-attack check user-bind enable
HUAWEI-Eth-Trunk10] quit
!
! --- OSPF (idem §5.1, router-id 10.255.255.11) ---
HUAWEI] ospf 1 router-id 10.255.255.11
HUAWEI-ospf] area 0
HUAWEI-ospf-area] network 10.0.0.0 0.0.0.255
HUAWEI-ospf-area] network 192.168.0.0 0.0.255.255
HUAWEI-ospf-area] quit
HUAWEI-ospf] silent-interface all
HUAWEI-ospf] undo silent-interface Eth-Trunk 10
HUAWEI-ospf] quit
```

### 5.3 Accès (S5735-L, PoE)

```shell
HUAWEI> system-view
! --- Ports utilisateurs ---
HUAWEI] port-group USER-PORTS
HUAWEI-port-group] group-member GigabitEthernet 0/0/1 to 0/0/44
HUAWEI-port-group] port link-type access
HUAWEI-port-group] port default vlan 10
HUAWEI-port-group] stp edged-port enable
HUAWEI-port-group] port-security enable
HUAWEI-port-group] port-security max-mac-num 3
HUAWEI-port-group] port-security protect-action restrict
HUAWEI-port-group] port-security aging-time 60
HUAWEI-port-group] quit
!
! --- Ports voix ---
HUAWEI] port-group VOICE-PORTS
HUAWEI-port-group] group-member GigabitEthernet 0/0/1 to 0/0/24
HUAWEI-port-group] voice-vlan 20 enable
HUAWEI-port-group] trust dscp
HUAWEI-port-group] quit
!
! --- Uplink ---
HUAWEI] interface Eth-Trunk 1
HUAWEI-Eth-Trunk1] mode lacp
HUAWEI-Eth-Trunk1] trunkport GigabitEthernet 0/0/47 to 0/0/48
HUAWEI-Eth-Trunk1] port link-type trunk
HUAWEI-Eth-Trunk1] port trunk allow-pass vlan 10 20 30 99
HUAWEI-Eth-Trunk1] dhcp snooping trusted
HUAWEI-Eth-Trunk1] quit
!
! --- 802.1X (optionnel) ---
HUAWEI] dot1x enable
HUAWEI] radius-server template RAD
HUAWEI-radius] radius-server shared-key cipher CleRADIUS
HUAWEI-radius] radius-server authentication 192.168.99.61 1812
HUAWEI-radius] radius-server accounting 192.168.99.61 1813
HUAWEI-radius] quit
HUAWEI] aaa
HUAWEI-aaa] authentication-scheme default
HUAWEI-aaa-authen] authentication-mode radius
HUAWEI-aaa] quit
HUAWEI] interface GigabitEthernet 0/0/1
HUAWEI-GE] dot1x enable
```

---

## 6. WAN entreprise — BGP multi-homing

### 6.1 Cisco (routeurs ISR 4000 / ASR)

```shell
CISCO(config)# ip prefix-list PL-LOCAL seq 5 permit 192.168.0.0/16
CISCO(config)# ip prefix-list PL-LOCAL seq 10 permit 203.0.113.0/24
CISCO(config)# ip prefix-list PL-FULL seq 5 deny 0.0.0.0/0
CISCO(config)# ip prefix-list PL-FULL seq 10 permit 0.0.0.0/0 le 24
!
CISCO(config)# route-map RM-OUT-FAI1 permit 10
CISCO(config-route-map)# match ip address prefix-list PL-LOCAL
CISCO(config-route-map)# exit
CISCO(config)# route-map RM-IN-FAI1 permit 10
CISCO(config-route-map)# match ip address prefix-list PL-FULL
CISCO(config-route-map)# set local-preference 200
CISCO(config-route-map)# exit
!
CISCO(config)# router bgp 65001
CISCO(config-router)# bgp router-id 10.255.255.254
CISCO(config-router)# neighbor 203.0.113.9 remote-as 100
CISCO(config-router)# neighbor 203.0.113.9 description FAI-1
CISCO(config-router)# neighbor 203.0.113.9 password MD5BGP
CISCO(config-router)# neighbor 203.0.113.9 route-map RM-IN-FAI1 in
CISCO(config-router)# neighbor 203.0.113.9 route-map RM-OUT-FAI1 out
CISCO(config-router)# neighbor 203.0.113.9 maximum-prefix 1000000
CISCO(config-router)# network 192.168.0.0 mask 255.255.0.0
CISCO(config-router)# network 203.0.113.0 mask 255.255.255.0
! (répéter vers FAI-2 avec local-preference 100 = backup)
CISCO# show ip bgp summary
CISCO# show ip bgp neighbors 203.0.113.9 advertised-routes
```

### 6.2 Huawei (routeurs AR / NetEngine)

```shell
HUAWEI> system-view
HUAWEI] ip ip-prefix PL-LOCAL index 10 permit 192.168.0.0 16 greater-equal 16 less-equal 24
HUAWEI] ip ip-prefix PL-LOCAL index 20 permit 203.0.113.0 24
HUAWEI] ip ip-prefix PL-FULL index 10 deny 0.0.0.0 0
HUAWEI] ip ip-prefix PL-FULL index 20 permit 0.0.0.0 0 greater-equal 1 less-equal 24
!
HUAWEI] route-policy RM-OUT-FAI1 permit node 10
HUAWEI-route-policy] if-match ip-prefix PL-LOCAL
HUAWEI-route-policy] quit
HUAWEI] route-policy RM-IN-FAI1 permit node 10
HUAWEI-route-policy] if-match ip-prefix PL-FULL
HUAWEI-route-policy] apply local-preference 200
HUAWEI-route-policy] quit
!
HUAWEI] bgp 65001
HUAWEI-bgp] router-id 10.255.255.254
HUAWEI-bgp] peer 203.0.113.9 as-number 100
HUAWEI-bgp] peer 203.0.113.9 description FAI-1
HUAWEI-bgp] peer 203.0.113.9 password cipher MD5BGP
HUAWEI-bgp] peer 203.0.113.9 ip-prefix PL-FULL import
HUAWEI-bgp] peer 203.0.113.9 route-policy RM-OUT-FAI1 export
HUAWEI-bgp-af] network 192.168.0.0 255.255.0.0
HUAWEI-bgp-af] network 203.0.113.0 255.255.255.0
HUAWEI-bgp-af] quit
HUAWEI-bgp] quit
HUAWEI] quit
HUAWEI> display bgp peer
HUAWEI> display bgp routing-table
```

> **Règle d'or BGP entreprise** : ne jamais accepter ni annoncer sans
> filtre. `maximum-prefix` (Cisco) / `peer … route-limit` (Huawei)
> protègent contre les fuites de tables complètes.

---

## 7. VPN inter-sites

### 7.1 Cisco — DMVPN (hub & spoke, recommandé multi-sites)

```shell
! --- HUB ---
CISCO(config)# interface Tunnel 0
CISCO(config-if)# ip address 172.16.0.1 255.255.255.0
CISCO(config-if)# ip nhrp network-id 1
CISCO(config-if)# ip nhrp authentication CleNHRP
CISCO(config-if)# tunnel source GigabitEthernet 0/0
CISCO(config-if)# tunnel mode gre multipoint
CISCO(config-if)# tunnel protection ipsec profile PROF-IPSEC
CISCO(config-if)# exit
CISCO(config)# router eigrp DMVPN
CISCO(config-router)# address-family ipv4 unicast autonomous-system 200
CISCO(config-router-af)# network 172.16.0.0 0.0.0.255
CISCO(config-router-af)# network 192.168.0.0 0.0.255.255
! --- SPOKE (site distant) ---
CISCO(config)# interface Tunnel 0
CISCO(config-if)# ip address 172.16.0.11 255.255.255.0
CISCO(config-if)# ip nhrp network-id 1
CISCO(config-if)# ip nhrp authentication CleNHRP
CISCO(config-if)# ip nhrp nhs 172.16.0.1
CISCO(config-if)# ip nhrp map 172.16.0.1 203.0.113.10
CISCO(config-if)# ip nhrp map multicast 203.0.113.10
CISCO(config-if)# tunnel source GigabitEthernet 0/0
CISCO(config-if)# tunnel mode gre multipoint
CISCO(config-if)# tunnel protection ipsec profile PROF-IPSEC
```

### 7.2 Huawei — IPsec site-à-site (IKEv2)

