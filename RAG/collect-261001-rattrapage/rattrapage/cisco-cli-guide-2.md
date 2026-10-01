---
id: collect-261001-rattrapage/rattrapage/cisco-cli-guide-2
title: "Cisco — Guide CLI complet (IOS / IOS XE)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/cisco_cli_guide.md
source_anchor: ""
source_lines: [249, 513]
sha256: 612e720dcfb0cba0473e207cf195fecbe61b8115d894bc9bbacece044177dce1
---

# Cisco — Guide CLI complet (IOS / IOS XE)

```shell
# EIGRP (named mode moderne)
R(config)# router eigrp LAN
R(config-router)# address-family ipv4 unicast autonomous-system 100
R(config-router-af)# network 192.168.1.0 0.0.0.255
R(config-router-af)# eigrp router-id 10.255.255.1
R(config-router-af)# exit
R# show ip eigrp neighbors
R# show ip eigrp topology

# BGP
R(config)# router bgp 65001
R(config-router)# bgp router-id 10.255.255.1
R(config-router)# neighbor 10.0.0.2 remote-as 65002
R(config-router)# neighbor 10.0.0.2 description FAI
R(config-router)# network 192.168.1.0 mask 255.255.255.0
R(config-router)# exit
R# show ip bgp summary
R# show ip bgp neighbors 10.0.0.2
R# show ip bgp
```

- BGP : `show ip bgp summary` — l'état doit être un nombre (préfixes),
  pas `Active`/`Idle` (problème de session).
- iBGP : full-mesh ou route-reflector ; eBGP multihop si plusieurs sauts :
  `neighbor x ebgp-multihop 2`.

---

## 10. ACL (listes de contrôle d'accès)

```shell
# Étendue numérotée
R(config)# access-list 100 permit tcp 192.168.1.0 0.0.0.255 any eq 80
R(config)# access-list 100 permit tcp 192.168.1.0 0.0.0.255 any eq 443
R(config)# access-list 100 deny ip any any log

# Nommée (recommandée : modifiable ligne par ligne)
R(config)# ip access-list extended FILTRE-WAN
R(config-ext-nacl)# permit tcp any host 192.168.50.10 eq 443
R(config-ext-nacl)# deny ip any any log
R(config-ext-nacl)# exit
R(config)# interface GigabitEthernet 0/0
R(config-if)# ip access-group FILTRE-WAN in

R# show access-lists
R# show ip access-lists FILTRE-WAN
```

- Le `deny ip any any` implicite existe toujours en fin d'ACL.
- `in`/`out` = sens par rapport au routeur ; testez dans le bon sens.
- `log` en fin de deny aide énormément au diagnostic.

---

## 11. NAT

```shell
R(config)# interface GigabitEthernet 0/0
R(config-if)# ip nat outside
R(config-if)# exit
R(config)# interface GigabitEthernet 0/1
R(config-if)# ip nat inside
R(config-if)# exit

# PAT (overload) : tout le LAN derrière l'IP WAN
R(config)# access-list 1 permit 192.168.1.0 0.0.0.255
R(config)# ip nat inside source list 1 interface GigabitEthernet 0/0 overload

# Pool de PAT
R(config)# ip nat pool PUBLIC 202.10.10.10 202.10.10.20 netmask 255.255.255.0
R(config)# ip nat inside source list 1 pool PUBLIC overload

# NAT statique (publication de serveur)
R(config)# ip nat inside source static tcp 192.168.50.10 80 202.10.10.10 8080 extendable
R(config)# ip nat inside source static 192.168.50.11 202.10.10.11   # 1:1

R# show ip nat translations
R# show ip nat statistics
R# clear ip nat translation *
```

- `inside`/`outside` inversés = NAT qui ne traduit rien (erreur n°1).
- Pour les protocoles à ports dynamiques : `ip nat service` / ALG
  (`show ip nat translations verbose`).

---

## 12. AAA, utilisateurs, SSH

```shell
# Utilisateurs locaux
R(config)# username admin privilege 15 secret MotDePasseFort
R(config)# username operateur privilege 7 secret AutreMotDePasse

# Lignes console / VTY
R(config)# line console 0
R(config-line)# password ConPass
R(config-line)# login
R(config-line)# logging synchronous
R(config-line)# exit
R(config)# line vty 0 4
R(config-line)# login local
R(config-line)# transport input ssh
R(config-line)# exec-timeout 10 0
R(config-line)# exit

# SSH
R(config)# crypto key generate rsa modulus 2048
R(config)# ip ssh version 2
R(config)# ip ssh time-out 60
R# show ip ssh
R# show users
R# show ssh
```

- `transport input ssh` seul = telnet désactivé (bien).
- AAA centralisé : `aaa new-model`, `tacacs-server host`, `radius-server
  host` (voir §20 pour l'exemple TACACS+).

---

## 13. VPN IPsec site-à-site (IKEv1)

```shell
# Phase 1
R(config)# crypto isakmp policy 10
R(config-isakmp)# encryption aes 256
R(config-isakmp)# hash sha256
R(config-isakmp)# authentication pre-share
R(config-isakmp)# group 14
R(config-isakmp)# lifetime 86400
R(config-isakmp)# exit
R(config)# crypto isakmp key ClePartagee address 202.20.20.2

# Phase 2
R(config)# crypto ipsec transform-set TS1 esp-aes 256 esp-sha256-hmac
R(cfg-crypto-trans)# mode tunnel
R(cfg-crypto-trans)# exit
R(config)# access-list 100 permit ip 192.168.1.0 0.0.0.255 192.168.2.0 0.0.0.255
R(config)# crypto map CMAP 10 ipsec-isakmp
R(config-crypto-map)# set peer 202.20.20.2
R(config-crypto-map)# set transform-set TS1
R(config-crypto-map)# set pfs group14
R(config-crypto-map)# match address 100
R(config-crypto-map)# exit
R(config)# interface GigabitEthernet 0/0
R(config-if)# crypto map CMAP

R# show crypto isakmp sa
R# show crypto ipsec sa
R# show crypto map
```

- ACL **miroir** des deux côtés (erreur n°1 de la phase 2).
- UDP 500/4500 + ESP autorisés dans les ACL d'entrée.
- `debug crypto isakmp` / `debug crypto ipsec` (avec parcimonie !).

> **Note crypto 2025** : IKEv1 + crypto maps sont **legacy**.
> Bannissez DES, 3DES, MD5, SHA-1 et les groupes DH 1/2/5 (cassés ou
> faibles). Minimum : AES-128/256, SHA-256, DH groupe 14 (2048 bits) ;
> recommandé : AES-GCM, DH groupes 19/20 (ECP). Préférez **IKEv2**
> (ci-dessous) pour tout nouveau déploiement.

### IKEv2 (moderne, recommandé)

```shell
R(config)# crypto ikev2 proposal PROP1
R(config-ikev2-proposal)# encryption aes-gcm-256
R(config-ikev2-proposal)# prf sha256
R(config-ikev2-proposal)# group 14
R(config-ikev2-proposal)# exit
R(config)# crypto ikev2 keyring KR1
R(config-ikev2-keyring)# peer PEER1
R(config-ikev2-keyring-peer)# address 202.20.20.2
R(config-ikev2-keyring-peer)# pre-shared-key local CleA
R(config-ikev2-keyring-peer)# pre-shared-key remote CleA
R(config-ikev2-keyring-peer)# exit
```

---

## 14. GRE et DMVPN (aperçu)

```shell
# GRE simple
R(config)# interface Tunnel 0
R(config-if)# ip address 172.16.0.1 255.255.255.252
R(config-if)# tunnel source GigabitEthernet 0/0
R(config-if)# tunnel destination 202.20.20.2
R(config-if)# tunnel mode gre ip

# GRE sur IPsec : appliquez la crypto map sur l'interface physique
# avec une ACL qui matche le trafic GRE (permit gre host A host B)
```

- DMVPN : hub `nhrp`, spokes dynamiques + IPsec + EIGRP/OSPF par-dessus ;
  idéal pour les réseaux d'agences (config en 3 phases : mGRE, NHRP,
  IPsec). Dites-moi si vous voulez le guide DMVPN détaillé.

---

## 15. QoS (MQC : class-map / policy-map)

```shell
R(config)# class-map match-any VOIP
R(config-cmap)# match dscp ef
R(config-cmap)# exit
R(config)# class-map match-any CRITIQUE
R(config-cmap)# match dscp af31
R(config-cmap)# match access-group 110
R(config-cmap)# exit
R(config)# policy-map QOS-WAN
R(config-pmap)# class VOIP
R(config-pmap-c)# priority percent 30
R(config-pmap-c)# exit
R(config-pmap)# class CRITIQUE
R(config-pmap-c)# bandwidth percent 40
R(config-pmap-c)# exit
R(config-pmap)# class class-default
R(config-pmap-c)# fair-queue
R(config-pmap-c)# random-detect
R(config-pmap-c)# exit
R(config-pmap)# exit
R(config)# interface GigabitEthernet 0/0
R(config-if)# service-policy output QOS-WAN

R# show policy-map interface GigabitEthernet 0/0
```

- `priority` = LLQ (temps réel), `bandwidth` = CBWFQ (garanti).
- Marquez au plus près de la source (`set dscp`), contrôlez en sortie.

---

## 16. Haute disponibilité : HSRP / VRRP / GLBP

```shell
# HSRP (Cisco propriétaire)
R(config)# interface GigabitEthernet 0/1
R(config-if)# standby 1 ip 192.168.1.254
R(config-if)# standby 1 priority 110
R(config-if)# standby 1 preempt
R(config-if)# standby 1 track GigabitEthernet 0/0 20

# VRRP (standard)
R(config-if)# vrrp 1 ip 192.168.1.254
R(config-if)# vrrp 1 priority 110
R(config-if)# vrrp 1 preempt

# GLBP (répartition de charge entre passerelles)
R(config-if)# glbp 1 ip 192.168.1.254
R(config-if)# glbp 1 priority 110
R(config-if)# glbp 1 preempt
R(config-if)# glbp 1 load-balancing round-robin

R# show standby brief
R# show vrrp brief
R# show glbp brief
```

---

## 17. Sécurité switching

