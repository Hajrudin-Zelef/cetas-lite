---
id: collect-261001-rattrapage/rattrapage/cisco-cli-guide-3
title: "Cisco — Guide CLI complet (IOS / IOS XE)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "attention", "license", "memory"]
source: docs/RAG/collect-261001-rattrapage/cisco_cli_guide.md
source_anchor: ""
source_lines: [514, 759]
sha256: 957926b219702d270022440f4eaa7667ee7140db314b1c735988dcd6e62666e3
---

# Cisco — Guide CLI complet (IOS / IOS XE)

```shell
# Port-security
SW(config)# interface GigabitEthernet 1/0/5
SW(config-if)# switchport port-security
SW(config-if)# switchport port-security maximum 2
SW(config-if)# switchport port-security violation restrict
SW(config-if)# switchport port-security mac-address sticky
SW# show port-security interface GigabitEthernet 1/0/5

# DHCP snooping + DAI + IP Source Guard (anti-spoofing)
SW(config)# ip dhcp snooping
SW(config)# ip dhcp snooping vlan 10
SW(config)# interface GigabitEthernet 1/0/24
SW(config-if)# ip dhcp snooping trust        # uplink uniquement !
SW(config)# ip arp inspection vlan 10
SW(config)# ip verify source vlan dhcp-snooping

# Storm control
SW(config)# interface GigabitEthernet 1/0/5
SW(config-if)# storm-control broadcast level 5.00
SW(config-if)# storm-control action trap
```

---

## 18. Supervision : syslog, SNMP, NetFlow

```shell
# Syslog
R(config)# logging host 192.168.1.100
R(config)# logging trap informational
R(config)# logging source-interface Loopback 0
R# show logging

# SNMP
R(config)# snmp-server community PublicRO RO
R(config)# snmp-server community PriveRW RW
R(config)# snmp-server host 192.168.1.100 version 2c PublicRO
R(config)# snmp-server enable traps
R(config)# snmp-server location SIEGE-Baie1
R(config)# snmp-server contact admin@entreprise.lan

# SNMPv3 (recommandé : chiffré et authentifié, remplace les community)
R(config)# snmp-server group ADMIN v3 priv
R(config)# snmp-server user operateur ADMIN v3 auth sha AuthPass priv aes 256 PrivPass
R(config)# snmp-server host 192.168.1.100 version 3 priv operateur

# NetFlow (trafic)
R(config)# interface GigabitEthernet 0/0
R(config-if)# ip flow ingress
R(config-if)# ip flow egress
R(config-if)# exit
R(config)# ip flow-export version 9
R(config)# ip flow-export destination 192.168.1.100 9996
R# show ip flow interface
R# show ip cache flow
```

---

## 19. Diagnostic

```shell
R# show version                    # modèle, IOS, uptime
R# show inventory                   # matériel / S/N
R# show processes cpu               # CPU (trié)
R# show processes cpu history       # graphique 72h
R# show memory                      # mémoire
R# show interfaces status
R# show ip interface brief
R# show interfaces Gi0/0 | include errors
R# show controllers Gi0/0           # niveau physique
R# ping 8.8.8.8 source Loopback 0
R# ping 8.8.8.8 size 1400 df-bit    # test MTU
R# traceroute 8.8.8.8
R# show ip route 8.8.8.8
R# show arp
R# show mac address-table address aabb.ccdd.eeff
R# show cdp neighbors detail        # voisins Cisco directs
R# show lldp neighbors              # standard LLDP
R# debug ip packet                  # ATTENTION : dangereux, toujours avec ACL !
R# debug ip ospf hello
R# undebug all                      # arrêter tous les debugs
R# show tech-support                # dump complet (pour le TAC)
```

- **Règle d'or du debug** : `access-list 199 permit ip host X any` puis
  `debug ip packet 199` — jamais de debug sans filtre en production.
- `service timestamps debug datetime msec` : horodate les debugs/logs.

---

## 20. AAA centralisé (TACACS+/RADIUS)

```shell
R(config)# aaa new-model
R(config)# tacacs-server host 192.168.1.50 key CleTACACS
R(config)# aaa authentication login default group tacacs+ local
R(config)# aaa authorization exec default group tacacs+ local
R(config)# aaa accounting exec default start-stop group tacacs+
R# test aaa group tacacs+ admin Cisco123 legacy   # test d'authentification
R# show aaa servers
```

- `group tacacs+ local` = fallback local si le serveur ne répond pas.
- RADIUS : `radius-server host 192.168.1.50 key CleRADIUS`.

---

## 21. Fichiers, upgrade IOS, mots de passe

```shell
R# dir flash:
R# show flash:
R# copy tftp: flash:                  # récupère l'image (nom demandé)
R# verify flash:c2900-universalk9-mz.SPA.157-3.M8.bin
R(config)# boot system flash:c2900-universalk9-mz.SPA.157-3.M8.bin
R(config)# config-register 0x2102
R# show boot
R# reload

# Sauvegarde de config sur TFTP
R# copy running-config tftp:
R# copy startup-config tftp:

# Mot de passe perdu (routeur) :
# 1. Redémarrer, Ctrl+Break -> rommon>
# 2. rommon> confreg 0x2142
# 3. rommon> reset  (démarre sans charger la startup-config)
# 4. Router# copy startup-config running-config
# 5. Changer les mots de passe, puis :
#    (config)# config-register 0x2102
#    # copy running-config startup-config + reload

# Licences (Smart Licensing, obligatoire depuis IOS XE 17.x)
R# show license status
R# show license summary
R(config)# license smart transport smart
# Plus de fichiers PAK : l'équipement s'enregistre sur le Smart Account
# (direct, via satellite on-prem, ou réservation hors-ligne).
```

- `0x2142` = ignore la startup-config ; `0x2102` = valeur normale.
- Sur switch Catalyst : mode `switch:` au boot, `flash_init`,
  `rename flash:config.text flash:config.old`, `boot`.

---

## 22. IPv6 (bases)

```shell
R(config)# ipv6 unicast-routing
R(config)# interface GigabitEthernet 0/1
R(config-if)# ipv6 address 2001:db8:1::1/64
R(config-if)# ipv6 enable
R(config)# ipv6 route ::/0 2001:db8:ff::1
R(config)# ipv6 access-list FILTRE6
R(config-ipv6-acl)# permit tcp any any eq 80
R(config-ipv6-acl)# deny ipv6 any any log
R# show ipv6 interface brief
R# show ipv6 route
R# ping ipv6 2001:db8::1
```

- OSPFv3 : `router ospfv3 1`, `address-family ipv6 unicast`.

---

## 23. Exemple complet : routeur de site

```shell
hostname R-AGENCE
no ip domain-lookup
ip domain-name agence.lan
service password-encryption
!
username admin privilege 15 secret AdminFort
!
interface GigabitEthernet 0/0
 description WAN_FAI
 ip address dhcp
 ip nat outside
 ip access-group FILTRE-WAN in
 no shutdown
!
interface GigabitEthernet 0/1
 description LAN
 ip address 192.168.10.1 255.255.255.0
 ip nat inside
 ip policy route-map PBR1
 no shutdown
!
ip dhcp excluded-address 192.168.10.1 192.168.10.10
ip dhcp pool LAN
 network 192.168.10.0 255.255.255.0
 default-router 192.168.10.1
 dns-server 8.8.8.8 1.1.1.1
!
ip access-list extended FILTRE-WAN
 permit tcp any host 202.10.10.10 eq 443
 deny ip any any log
!
access-list 1 permit 192.168.10.0 0.0.0.255
ip nat inside source list 1 interface GigabitEthernet 0/0 overload
!
ip route 0.0.0.0 0.0.0.0 GigabitEthernet 0/0 dhcp
!
router ospf 1
 router-id 10.255.10.1
 network 192.168.10.0 0.0.0.255 area 0
 passive-interface default
 no passive-interface Tunnel 0
!
line vty 0 4
 login local
 transport input ssh
!
crypto key generate rsa modulus 2048
ip ssh version 2
!
ntp server 192.168.1.100
logging host 192.168.1.100
snmp-server community PublicRO RO
!
end
copy running-config startup-config
```

---

## 24. Différences IOS XR / NX-OS (repères)

- **IOS XR** : après chaque modif, `commit` (ou `commit confirmed`) ;
  `show configuration` = changements en attente ; `show running-config`
  existe aussi. Interfaces nommées `GigabitEthernet0/0/0/0`.
- **NX-OS** : `feature ospf` / `feature bgp` / `feature hsrp` à activer
  avant usage ; `show running-config` identique ; VRF par défaut
  `management` ; `copy running-config startup-config` identique.
- Les concepts (VLAN, OSPF, BGP, NAT, IPsec) sont les mêmes, seule la
  syntaxe varie à la marge.

---

*Fin du guide. Pour aller plus loin : les Command Reference officiels
Cisco (IOS Master Command List) détaillent chaque commande avec ses
options exactes par version.*
