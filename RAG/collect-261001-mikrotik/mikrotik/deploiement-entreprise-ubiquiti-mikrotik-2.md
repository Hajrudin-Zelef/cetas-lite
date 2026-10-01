---
id: collect-261001-mikrotik/mikrotik/deploiement-entreprise-ubiquiti-mikrotik-2
title: "Déploiement entreprise — Ubiquiti & MikroTik (référence complète)"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-mikrotik/deploiement_entreprise_ubiquiti_mikrotik.md
source_anchor: ""
source_lines: [186, 358]
sha256: 2daf73e6026aef7b9d606ab36bf374ad351732d1ed2f3d7447236a4f0755ecb7
---

# Déploiement entreprise — Ubiquiti & MikroTik (référence complète)

```shell
UBNT> configure
! --- WAN ---
UBNT# set interfaces ethernet eth0 address dhcp
UBNT# set interfaces ethernet eth0 description WAN1-FAI-FIBRE
UBNT# set interfaces ethernet eth1 address 198.51.100.50/29
UBNT# set interfaces ethernet eth1 description WAN2-FAI-SECOURS
! --- LAN trunk vers switch ---
UBNT# set interfaces ethernet eth2 description TRUNK-SWITCH
! --- VLANs (VIF) ---
UBNT# set interfaces ethernet eth2 vif 10 address 192.168.10.254/24
UBNT# set interfaces ethernet eth2 vif 10 description VLAN10-DATA
UBNT# set interfaces ethernet eth2 vif 20 address 192.168.20.254/24
UBNT# set interfaces ethernet eth2 vif 20 description VLAN20-VOIX
UBNT# set interfaces ethernet eth2 vif 30 address 192.168.30.254/24
UBNT# set interfaces ethernet eth2 vif 30 description VLAN30-INVITES
UBNT# set interfaces ethernet eth2 vif 40 address 192.168.40.254/24
UBNT# set interfaces ethernet eth2 vif 40 description VLAN40-SERVEURS
UBNT# set interfaces ethernet eth2 vif 99 address 192.168.99.254/24
UBNT# set interfaces ethernet eth2 vif 99 description VLAN99-MGMT
! --- Bonding LACP (si 2 uplinks physiques) ---
UBNT# set interfaces bonding bond0 mode 802.3ad
UBNT# set interfaces bonding bond0 member interface eth3
UBNT# set interfaces bonding bond0 member interface eth4
UBNT# set interfaces bonding bond0 vif 10 address 192.168.10.253/24
UBNT# commit
UBNT# save
```

### 5.2 Firewall policies (zones)

```shell
UBNT> configure
! --- Politique LAN vers WAN ---
UBNT# set firewall name LAN-WAN default-action accept
UBNT# set firewall name LAN-WAN description "LAN vers Internet"
! --- Politique INVITES : Internet seul ---
UBNT# set firewall name GUEST-WAN default-action drop
UBNT# set firewall name GUEST-WAN rule 10 action accept
UBNT# set firewall name GUEST-WAN rule 10 description "DNS"
UBNT# set firewall name GUEST-WAN rule 10 destination port 53
UBNT# set firewall name GUEST-WAN rule 10 protocol udp
UBNT# set firewall name GUEST-WAN rule 20 action accept
UBNT# set firewall name GUEST-WAN rule 20 description "Web"
UBNT# set firewall name GUEST-WAN rule 20 destination port 80,443
UBNT# set firewall name GUEST-WAN rule 20 protocol tcp
UBNT# set firewall name GUEST-WAN rule 30 action drop
UBNT# set firewall name GUEST-WAN rule 30 description "Bloque RFC1918"
UBNT# set firewall name GUEST-WAN rule 30 destination address 10.0.0.0/8
UBNT# set firewall name GUEST-WAN rule 31 action drop
UBNT# set firewall name GUEST-WAN rule 31 destination address 172.16.0.0/12
UBNT# set firewall name GUEST-WAN rule 32 action drop
UBNT# set firewall name GUEST-WAN rule 32 destination address 192.168.0.0/16
! --- WAN vers LAN : tout bloqué par défaut ---
UBNT# set firewall name WAN-LAN default-action drop
UBNT# set firewall name WAN-LAN rule 10 action accept
UBNT# set firewall name WAN-LAN rule 10 state established enable
UBNT# set firewall name WAN-LAN rule 10 state related enable
UBNT# set firewall name WAN-LAN rule 20 action drop
UBNT# set firewall name WAN-LAN rule 20 state invalid enable
! --- Application aux interfaces ---
UBNT# set interfaces ethernet eth2 vif 30 firewall in name GUEST-WAN
UBNT# set interfaces ethernet eth2 vif 30 firewall local name WAN-LAN
UBNT# set interfaces ethernet eth0 firewall in name WAN-LAN
UBNT# set interfaces ethernet eth0 firewall local name WAN-LOCAL
UBNT# commit
UBNT# save
UBNT> show firewall
```

### 5.3 NAT

```shell
UBNT> configure
! --- Masquerade (PAT) ---
UBNT# set service nat rule 5000 description "PAT LAN vers WAN1"
UBNT# set service nat rule 5000 outbound-interface eth0
UBNT# set service nat rule 5000 source address 192.168.0.0/16
UBNT# set service nat rule 5000 type masquerade
! --- Port-forward ---
UBNT# set service nat rule 10 description "HTTPS vers DMZ"
UBNT# set service nat rule 10 inbound-interface eth0
UBNT# set service nat rule 10 destination port 443
UBNT# set service nat rule 10 protocol tcp
UBNT# set service nat rule 10 inside-address address 192.168.40.10
UBNT# set service nat rule 10 inside-address port 443
UBNT# set service nat rule 10 type destination
! --- Règle firewall associée (WAN-LAN) ---
UBNT# set firewall name WAN-LAN rule 30 action accept
UBNT# set firewall name WAN-LAN rule 30 destination address 192.168.40.10
UBNT# set firewall name WAN-LAN rule 30 destination port 443
UBNT# set firewall name WAN-LAN rule 30 protocol tcp
UBNT# commit
UBNT# save
UBNT> show nat rules
UBNT> show nat translations
```

### 5.4 DHCP + DNS

```shell
UBNT> configure
UBNT# set service dhcp-server disabled false
UBNT# set service dhcp-server shared-network-name DATA subnet 192.168.10.0/24 default-router 192.168.10.254
UBNT# set service dhcp-server shared-network-name DATA subnet 192.168.10.0/24 dns-server 192.168.10.254
UBNT# set service dhcp-server shared-network-name DATA subnet 192.168.10.0/24 range 0 start 192.168.10.100
UBNT# set service dhcp-server shared-network-name DATA subnet 192.168.10.0/24 range 0 stop 192.168.10.200
UBNT# set service dhcp-server shared-network-name DATA subnet 192.168.10.0/24 static-mapping SRV-WEB ip-address 192.168.40.10
UBNT# set service dhcp-server shared-network-name DATA subnet 192.168.10.0/24 static-mapping SRV-WEB mac-address AA:BB:CC:DD:EE:FF
UBNT# set service dns forwarding listen-on eth2.10
UBNT# set service dns forwarding listen-on eth2.20
UBNT# set service dns forwarding name-server 9.9.9.9
UBNT# set service dns forwarding name-server 1.1.1.1
UBNT# commit
UBNT# save
UBNT> show dhcp leases
```

### 5.5 Routage : OSPF + failover WAN

```shell
UBNT> configure
! --- OSPF interne ---
UBNT# set protocols ospf area 0 network 192.168.0.0/16
UBNT# set protocols ospf area 0 network 10.0.0.0/8
UBNT# set protocols ospf parameters router-id 10.255.255.1
UBNT# set protocols ospf passive-interface eth2.10
UBNT# set protocols ospf passive-interface eth2.20
! --- Failover WAN (load-balance avec health-check) ---
UBNT# set load-balance group WAN-FO interface eth0
UBNT# set load-balance group WAN-FO interface eth1
UBNT# set load-balance group WAN-FO interface eth0 route-test type ping target 9.9.9.9
UBNT# set load-balance group WAN-FO interface eth1 route-test type ping target 1.1.1.1
UBNT# set protocols static table 1 route 0.0.0.0/0 next-hop <GW-WAN1>
UBNT# set protocols static table 2 route 0.0.0.0/0 next-hop <GW-WAN2>
UBNT# set firewall modify WAN-RULE rule 10 action modify
UBNT# set firewall modify WAN-RULE rule 10 modify table main
UBNT# set firewall modify WAN-RULE rule 10 source address 192.168.0.0/16
UBNT# set interfaces ethernet eth2 vif 10 firewall in modify WAN-RULE
UBNT# commit
UBNT# save
UBNT> show ip ospf neighbor
UBNT> show load-balance status
```

### 5.6 VPN : IPsec site-à-site + WireGuard

```shell
UBNT> configure
! --- IPsec IKEv2 vers agence ---
UBNT# set vpn ipsec ipsec-interfaces interface eth0
UBNT# set vpn ipsec nat-traversal enable
UBNT# set vpn ipsec ike-group IKE-AGENCE proposal 1 encryption aes256
UBNT# set vpn ipsec ike-group IKE-AGENCE proposal 1 hash sha256
UBNT# set vpn ipsec ike-group IKE-AGENCE proposal 1 dh-group 14
UBNT# set vpn ipsec esp-group ESP-AGENCE proposal 1 encryption aes256
UBNT# set vpn ipsec esp-group ESP-AGENCE proposal 1 hash sha256
UBNT# set vpn ipsec site-to-site peer 198.51.100.2 authentication mode pre-shared-secret
UBNT# set vpn ipsec site-to-site peer 198.51.100.2 authentication pre-shared-secret CleIPsec
UBNT# set vpn ipsec site-to-site peer 198.51.100.2 ike-group IKE-AGENCE
UBNT# set vpn ipsec site-to-site peer 198.51.100.2 tunnel 1 esp-group ESP-AGENCE
UBNT# set vpn ipsec site-to-site peer 198.51.100.2 tunnel 1 local prefix 192.168.0.0/16
UBNT# set vpn ipsec site-to-site peer 198.51.100.2 tunnel 1 remote prefix 192.168.100.0/24
UBNT# commit
UBNT# save
UBNT> show vpn ipsec sa
```

> WireGuard sur EdgeOS : via `set interfaces wireguard wg0 …`
> (EdgeOS 2.x) — même logique que RouterOS §6.5.

### 5.7 QoS (traffic-policy shaper)

