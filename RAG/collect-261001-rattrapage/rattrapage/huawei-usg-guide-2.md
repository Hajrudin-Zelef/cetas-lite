---
id: collect-261001-rattrapage/rattrapage/huawei-usg-guide-2
title: "Huawei USG — Guide CLI complet (USG6000 series, V500R005 / V600)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["sandbox"]
source: docs/RAG/collect-261001-rattrapage/huawei_usg_guide.md
source_anchor: ""
source_lines: [247, 527]
sha256: 6c57d66cac0ffabbf49ff9806c8a0d2b4d697d87a3eb49a955f26ab19ca5c7a6
---

# Huawei USG — Guide CLI complet (USG6000 series, V500R005 / V600)

- Le NAT se configure **entre zones**, comme la security-policy.
- Vérifier : `display nat-policy rule all`, `display nat session all`.

---

## 9. NAT server (publication de serveurs)

```shell
system-view
# Publie le serveur interne 192.168.50.10:80 sur l'IP publique :8080
[USG] nat server protocol tcp global 202.10.10.10 8080 inside 192.168.50.10 80

# Ne pas oublier la règle security-policy untrust -> dmz (voir §6/§7) !

# Vérifications
display nat server
display firewall session table destination-ip 192.168.50.10
```

- `nat server` sans nom = global ; forme nommée possible avec zones.
- Pensez au *hairpin* : pour accès interne via l'IP publique, ajouter une
  règle NAT adaptée ou utiliser le DNS interne (split-DNS).

---

## 10. Routage

```shell
system-view
# Route statique par défaut
[USG] ip route-static 0.0.0.0 0.0.0.0 202.10.10.1

# OSPF
[USG] ospf 1
[USG-ospf-1] area 0
[USG-ospf-1-area-0.0.0.0] network 192.168.1.0 0.0.0.255
[USG-ospf-1-area-0.0.0.0] quit
[USG-ospf-1] quit

# BGP simple
[USG] bgp 65001
[USG-bgp] peer 10.0.0.2 as-number 65002
[USG-bgp] network 192.168.1.0 24
[USG-bgp] quit

# Routage par politique (PBR)
[USG] policy-based-route
[USG-policy-pbr] rule name PBR1
[USG-policy-pbr-rule-PBR1] ingress-interface GigabitEthernet 0/0/1
[USG-policy-pbr-rule-PBR1] action next-hop 10.0.0.254
```

- Vérifier : `display ip routing-table`, `display ospf peer`,
  `display bgp peer`.

---

## 11. DHCP, DNS, NTP

```shell
system-view
[USG] dhcp enable
[USG] ip pool LAN_POOL
[USG-ip-pool-LAN_POOL] network 192.168.1.0 mask 24
[USG-ip-pool-LAN_POOL] gateway-list 192.168.1.1
[USG-ip-pool-LAN_POOL] dns-list 8.8.8.8 1.1.1.1
[USG-ip-pool-LAN_POOL] lease day 3 hour 0 minute 0
[USG-ip-pool-LAN_POOL] quit
[USG] interface GigabitEthernet 0/0/1
[USG-GigabitEthernet0/0/1] dhcp select global
[USG-GigabitEthernet0/0/1] quit

[USG] dns resolve
[USG] dns server 8.8.8.8
[USG] ntp-service unicast-server 192.168.1.100

display ip pool name LAN_POOL
```

---

## 12. AAA et comptes admin

```shell
system-view
[USG] aaa
[USG-aaa] local-user operateur password irreversible-cipher Huawei@123
[USG-aaa] local-user operateur privilege level 1
[USG-aaa] local-user operateur service-type ssh web
[USG-aaa] local-user admin2 password irreversible-cipher Huawei@123
[USG-aaa] local-user admin2 privilege level 3
[USG-aaa] local-user admin2 service-type ssh web
[USG-aaa] quit

# Voir les utilisateurs en ligne
display access-user
display local-user
```

- Niveau 3 = administrateur complet, 0 = visite.
- Pour l'auth 802.1X / portail captif : `authentication-scheme`,
  `portal` (souvent configuré via web UI).

---

## 13. SSH et accès distant sécurisé

```shell
system-view
[USG] stelnet server enable
[USG] ssh user admin
[USG] ssh user admin authentication-type password
[USG] ssh user admin service-type stelnet
[USG] ssh server port 2222   # optionnel : changer le port
[USG] user-interface vty 0 4
[USG-ui-vty0-4] authentication-mode aaa
[USG-ui-vty0-4] protocol inbound ssh
[USG-ui-vty0-4] quit
```

---

## 14. VPN IPsec site-à-site

```shell
system-view
# 1. ACL = trafic à chiffrer
[USG] acl number 3000
[USG-acl-adv-3000] rule 5 permit ip source 192.168.1.0 0.0.0.255 destination 192.168.2.0 0.0.0.255
[USG-acl-adv-3000] quit

# 2. Phase 1 (IKE)
[USG] ike proposal 10
[USG-ike-proposal-10] encryption-algorithm aes-256
[USG-ike-proposal-10] dh group14
[USG-ike-proposal-10] authentication-algorithm sha2-256
[USG-ike-proposal-10] quit
[USG] ike peer PEER_PARIS
[USG-ike-peer-PEER_PARIS] pre-shared-key Huawei@123
[USG-ike-peer-PEER_PARIS] ike-proposal 10
[USG-ike-peer-PEER_PARIS] remote-address 202.20.20.2
[USG-ike-peer-PEER_PARIS] quit

# 3. Phase 2 (IPsec)
[USG] ipsec proposal IPSEC1
[USG-ipsec-proposal-IPSEC1] esp authentication-algorithm sha2-256
[USG-ipsec-proposal-IPSEC1] esp encryption-algorithm aes-256
[USG-ipsec-proposal-IPSEC1] quit
[USG] ipsec policy POLICY1 10 isakmp
[USG-ipsec-policy-isakmp-POLICY1-10] security acl 3000
[USG-ipsec-policy-isakmp-POLICY1-10] ike-peer PEER_PARIS
[USG-ipsec-policy-isakmp-POLICY1-10] proposal IPSEC1
[USG-ipsec-policy-isakmp-POLICY1-10] quit

# 4. Application sur l'interface WAN
[USG] interface GigabitEthernet 0/0/2
[USG-GigabitEthernet0/0/2] ipsec policy POLICY1
[USG-GigabitEthernet0/0/2] quit

# Ne pas oublier : security-policy local/untrust pour IKE (UDP 500/4500)
# et ESP entre les peers !

# Vérifications
display ike sa
display ipsec sa
display ike proposal
```

- Pour du *route-based* : utilisez `ipsec policy-template` + interface
  Tunnel avec `tunnel-protocol ipsec`.
- NAT-T automatique si un NAT est détecté sur le chemin.

---

## 15. GRE, L2TP, SSL VPN

```shell
# GRE simple
system-view
[USG] interface Tunnel 0/0/1
[USG-Tunnel0/0/1] tunnel-protocol gre
[USG-Tunnel0/0/1] source 202.10.10.10
[USG-Tunnel0/0/1] destination 202.20.20.2
[USG-Tunnel0/0/1] ip address 172.16.0.1 30
[USG-Tunnel0/0/1] quit
```

- **L2TP** (accès distant) : serveur L2TP + groupe d'adresses + AAA ;
  configuration typique via l'assistant web UI.
- **SSL VPN** : passerelle virtuelle (Web proxy, partage de fichiers,
  redirection de ports, extension réseau). Se configure essentiellement
  via la web UI (`SSL VPN > Gateway`), avec certificats et utilisateurs
  AAA. Idéal pour le télétravail sans client lourd.

---

## 16. UTM : IPS, antivirus, filtrage URL

```shell
system-view
# Profils (les contenus exacts se règlent en web UI)
[USG] profile type ips name IPS_STRICT
[USG-profile-ips-IPS_STRICT] quit
[USG] profile type av name AV_DEFAULT
[USG-profile-av-AV_DEFAULT] quit
[USG] profile type url-filter name URL_DEFAULT
[USG-profile-url-filter-URL_DEFAULT] quit

# Application dans une règle security-policy
[USG] security-policy
[USG-policy-security] rule name LAN_TO_WAN
[USG-policy-security-rule-LAN_TO_WAN] profile ips IPS_STRICT
[USG-policy-security-rule-LAN_TO_WAN] profile av AV_DEFAULT
[USG-policy-security-rule-LAN_TO_WAN] profile url-filter URL_DEFAULT
[USG-policy-security-rule-LAN_TO_WAN] quit
[USG-policy-security] quit

display utm profile all
display ips signature-database
```

- Autres moteurs : `file-blocking`, `content-filter` (mots-clés),
  `antispam`, `apt` (sandbox), `dns-filter`.
- Les bases de signatures se mettent à jour via l'*Update Center*
  (web UI) ou en planifié ; licence UTM requise.
- Pour HTTPS : inspection SSL (certificat CA à déployer) via
  `ssl-encrypted-traffic-detection`.

---

## 17. Défense anti-attaque et listes

```shell
system-view
# Listes noires / blanches
[USG] firewall blacklist 203.0.113.5
[USG] firewall whitelist 198.51.100.10
display firewall blacklist
display firewall whitelist

# Protections : flood SYN/UDP/ICMP, IP sweep, port scan...
# (réglages fins en web UI : Security > Attack Defense)
```

- Le USG protège aussi le plan de contrôle (zone `local`) contre les
  floods par défaut (CPCAR).
- Journalisez les attaques : `display firewall defend statistics`.

---

## 18. QoS / gestion de bande passante

```shell
system-view
# Politique de trafic : classifier + behavior
[USG] traffic classifier VOIP
[USG-classifier-VOIP] if-match dscp ef
[USG-classifier-VOIP] quit
[USG] traffic behavior PRIORITAIRE
[USG-behavior-PRIORITAIRE] queue af bandwidth pct 30
[USG-behavior-PRIORITAIRE] quit
[USG] traffic policy QOS1
[USG-trafficpolicy-QOS1] classifier VOIP behavior PRIORITAIRE
[USG-trafficpolicy-QOS1] quit

# Limitation de débit par règle (bandwidth) : voir web UI
# Security Policy > règle > Bandwidth
```

---

## 19. Haute disponibilité (HRP / VRRP)

```shell
system-view
# Hot Standby (actif/passif) entre deux USG
[USG] hrp enable
[USG] hrp interface GigabitEthernet 0/0/5 remote 10.10.10.2
[USG] hrp auto-sync config       # synchro auto de la config
[USG] hrp preempt                # préemption du primaire

