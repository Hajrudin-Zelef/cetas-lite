---
id: collect-261001-cisco/cisco/deploiement-entreprise-cisco-huawei-4
title: "Déploiement entreprise — Cisco & Huawei (référence complète)"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-cisco/deploiement_entreprise_cisco_huawei.md
source_anchor: ""
source_lines: [577, 791]
sha256: e1428ce97534d4507353f9f8f057cfcc442a37d9f57193f63925fa5f58449310
---

# Déploiement entreprise — Cisco & Huawei (référence complète)

```shell
HUAWEI> system-view
HUAWEI] ike proposal 10
HUAWEI-ike-proposal] encryption-algorithm aes-256
HUAWEI-ike-proposal] dh group14
HUAWEI-ike-proposal] authentication-algorithm sha2-256
HUAWEI-ike-proposal] integrity-algorithm hmac-sha2-256
HUAWEI-ike-proposal] quit
HUAWEI] ike peer PEER-SITEB
HUAWEI-ike-peer] ike-proposal 10
HUAWEI-ike-peer] remote-address 198.51.100.2
HUAWEI-ike-peer] pre-shared-key cipher CleIPsec
HUAWEI-ike-peer] quit
HUAWEI] ipsec proposal PROP1
HUAWEI-ipsec-proposal] esp authentication-algorithm sha2-256
HUAWEI-ipsec-proposal] esp encryption-algorithm aes-256
HUAWEI-ipsec-proposal] quit
HUAWEI] acl number 3001
HUAWEI-acl] rule permit ip source 192.168.0.0 0.0.255.255 destination 192.168.100.0 0.0.0.255
HUAWEI-acl] quit
HUAWEI] ipsec policy MAP1 10 isakmp
HUAWEI-ipsec-policy] security acl 3001
HUAWEI-ipsec-policy] ike-peer PEER-SITEB
HUAWEI-ipsec-policy] proposal PROP1
HUAWEI-ipsec-policy] quit
HUAWEI] interface GigabitEthernet 0/0/0
HUAWEI-GE] ipsec policy MAP1
HUAWEI-GE] quit
HUAWEI] quit
HUAWEI> display ike sa
HUAWEI> display ipsec sa
```

---

## 8. Haute disponibilité

| Fonction | Cisco | Huawei |
|---|---|---|
| Passerelle redondante | HSRP (`standby`) | VRRP (`vrrp vrid`) |
| Load-sharing passerelle | GLBP | VRRP + load-balance / Eth-Trunk |
| Stacking | StackWise (Catalyst) | CSS / iStack (S) |
| Firewall | HSRP / FHRP | HRP (USG, voir guide USG) |

- **Interop passerelle** : utilisez **VRRP des deux côtés** si le
  segment est mixte (HSRP est propriétaire Cisco).
- **Stacking** : simplifie la gestion (1 seul plan de contrôle) mais
  crée un domaine de panne unique — en core, préférez 2 châssis
  indépendants + HSRP/VRRP.

---

## 9. QoS entreprise

### 9.1 Cisco (MQC)

```shell
! --- Classification et marquage à l'accès ---
CISCO(config)# class-map match-any VOIP
CISCO(config-cmap)# match dscp ef
CISCO(config-cmap)# exit
CISCO(config)# class-map match-any SIGNAL
CISCO(config-cmap)# match dscp cs3
CISCO(config-cmap)# exit
CISCO(config)# class-map match-any CRITIQUE
CISCO(config-cmap)# match dscp af31 af32 af33
CISCO(config-cmap)# exit
!
CISCO(config)# policy-map QOS-CORE-OUT
CISCO(config-pmap)# class VOIP
CISCO(config-pmap-c)# priority percent 30
CISCO(config-pmap-c)# exit
CISCO(config-pmap)# class SIGNAL
CISCO(config-pmap-c)# bandwidth percent 5
CISCO(config-pmap-c)# exit
CISCO(config-pmap)# class CRITIQUE
CISCO(config-pmap-c)# bandwidth percent 30
CISCO(config-pmap-c)# random-detect dscp-based
CISCO(config-pmap-c)# exit
CISCO(config-pmap)# class class-default
CISCO(config-pmap-c)# fair-queue
CISCO(config-pmap-c)# random-detect
CISCO(config-pmap-c)# exit
CISCO(config-pmap)# exit
CISCO(config)# interface range TenGigabitEthernet 1/0/1 - 2
CISCO(config-if-range)# service-policy output QOS-CORE-OUT
!
CISCO# show policy-map interface TenGigabitEthernet 1/0/1
```

### 9.2 Huawei (MQC : classifier / behavior / policy)

```shell
HUAWEI> system-view
HUAWEI] traffic classifier VOIP
HUAWEI-classifier] if-match dscp ef
HUAWEI-classifier] quit
HUAWEI] traffic classifier CRITIQUE
HUAWEI-classifier] if-match dscp af31
HUAWEI-classifier] if-match dscp af32
HUAWEI-classifier] quit
HUAWEI] traffic behavior B-VOIP
HUAWEI-behavior] queue af bandwidth pct 30
HUAWEI-behavior] quit
HUAWEI] traffic behavior B-CRITIQUE
HUAWEI-behavior] queue af bandwidth pct 30
HUAWEI-behavior] quit
HUAWEI] traffic policy QOS-CORE-OUT
HUAWEI-policy] classifier VOIP behavior B-VOIP
HUAWEI-policy] classifier CRITIQUE behavior B-CRITIQUE
HUAWEI-policy] quit
HUAWEI] interface 10GE 1/0/1
HUAWEI-10GE] traffic-policy QOS-CORE-OUT outbound
HUAWEI-10GE] quit
HUAWEI] quit
HUAWEI> display traffic policy applied-record
```

> **Modèle DiffServ commun** : EF = voix (LLQ/priority), CS3 = signalisation,
> AF31 = données critiques, Default = best effort. Marquez au plus près
> de la source (`trust dscp` sur les ports voix), appliquez en sortie
> des uplinks.

---

## 10. Sécurité : durcissement entreprise

### 10.1 ACL d'infrastructure (Cisco + Huawei)

```shell
! --- Cisco : protéger le plan de contrôle ---
CISCO(config)# ip access-list extended ACL-INFRA
CISCO(config-ext-nacl)# permit tcp 192.168.99.0 0.0.0.255 any eq 22
CISCO(config-ext-nacl)# permit udp 192.168.99.0 0.0.0.255 any eq snmp
CISCO(config-ext-nacl)# permit tcp 192.168.99.0 0.0.0.255 any eq 443
CISCO(config-ext-nacl)# deny ip any any log
CISCO(config-ext-nacl)# exit
CISCO(config)# line vty 0 15
CISCO(config-line)# access-class ACL-INFRA in
```

```shell
! --- Huawei : ACL + appel sur VTY ---
HUAWEI> system-view
HUAWEI] acl number 3000
HUAWEI-acl] rule permit tcp source 192.168.99.0 0.0.0.255 destination-port eq 22
HUAWEI-acl] rule permit udp source 192.168.99.0 0.0.0.255 destination-port eq snmp
HUAWEI-acl] rule deny ip
HUAWEI-acl] quit
HUAWEI] user-interface vty 0 14
HUAWEI-ui-vty] acl 3000 inbound
HUAWEI-ui-vty] quit
HUAWEI] quit
HUAWEI> save
```

### 10.2 Protections anti-spoofing / anti-attaque

| Menace | Cisco | Huawei |
|---|---|---|
| Usurpation IP source | `ip verify source` (+ `ip verify unicast source reachable-via rx`) | `arp anti-attack check user-bind` / uRPF `ip urpf` |
| DHCP rogue | `ip dhcp snooping` | `dhcp snooping enable` |
| ARP spoofing | `ip arp inspection vlan` (DAI) | `arp anti-attack` / DAI équivalent |
| Storm broadcast | `storm-control broadcast level` | `storm suppression` / `broadcast-suppression` |
| BPDU rogue | `spanning-tree bpduguard` | `stp bpdu-protection` |

### 10.3 Checklist hardening
- [ ] Telnet désactivé partout, SSH v2 + clés.
- [ ] Comptes nominatifs, 2FA sur les accès VPN/admin, TACACS+/RADIUS.
- [ ] VTY/console filtrés par ACL (réseau MGMT uniquement).
- [ ] SNMP v3 authPriv ; communities v2 supprimées.
- [ ] DHCP snooping + DAI + IP Source Guard sur tous les VLANs users.
- [ ] Port-security + BPDU guard sur tous les ports d'accès.
- [ ] Storm-control sur les ports d'accès.
- [ ] NTP authentifié si possible ; logs vers SIEM redondant.
- [ ] Firmware : version validée, `show version` documentée, procédure
      de rollback testée.

---

## 11. Interopérabilité Cisco ↔ Huawei (réseaux mixtes)

| Protocole | Interopérable ? | Points d'attention |
|---|---|---|
| Trunk 802.1Q | ✅ | **Native VLAN identique** des 2 côtés ; `allowed vlan` miroir |
| LACP | ✅ | `mode active` (Cisco) ↔ `mode lacp` (Huawei), les deux en actif |
| RSTP | ✅ | Timers par défaut identiques ; vérifier `display stp brief` |
| MSTP | ✅ | **Region name + revision + mapping VLAN→instance IDENTIQUES** |
| OSPF | ✅ | Mêmes timers hello/dead, même auth MD5, MTU identique |
| BGP | ✅ | eBGP standard ; MD5 compatible |
| VRRP | ✅ | Même VRID, même VIP, priorités cohérentes |
| HSRP | ❌ | Propriétaire Cisco → utilisez VRRP en mixte |
| DHCP snooping | ⚠️ | Chaque vendor gère sa base ; configurer le trust de chaque côté |
| QoS DSCP | ✅ | Le marquage DSCP est standard ; `trust dscp` des deux côtés |

**Exemple : trunk Cisco ↔ Huawei**
```shell
! Cisco
CISCO(config-if)# switchport mode trunk
CISCO(config-if)# switchport trunk allowed vlan 10,20,30
CISCO(config-if)# switchport trunk native vlan 99
```
```shell
! Huawei
HUAWEI-Eth-Trunk] port link-type trunk
HUAWEI-Eth-Trunk] port trunk allow-pass vlan 10 20 30
HUAWEI-Eth-Trunk] port trunk pvid vlan 99
```
> Le `native vlan` Cisco = le `pvid` Huawei. **Mismatch = VLAN hopping
> potentiel et trafic perdu** : vérifiez toujours les deux côtés.

---

## 12. Supervision et télémétrie

