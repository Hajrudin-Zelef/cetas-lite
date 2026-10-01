---
id: collect-261001-mikrotik/mikrotik/deploiement-entreprise-ubiquiti-mikrotik-3
title: "Déploiement entreprise — Ubiquiti & MikroTik (référence complète)"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-mikrotik/deploiement_entreprise_ubiquiti_mikrotik.md
source_anchor: ""
source_lines: [359, 496]
sha256: a55ea002e31b77611c7bf12f991c2119bebbbdb404523273fcab99a414ea5b34
---

# Déploiement entreprise — Ubiquiti & MikroTik (référence complète)

```shell
UBNT> configure
UBNT# set traffic-policy shaper WAN-OUT bandwidth 1000mbit
UBNT# set traffic-policy shaper WAN-OUT class 10 bandwidth 30%
UBNT# set traffic-policy shaper WAN-OUT class 10 priority 7
UBNT# set traffic-policy shaper WAN-OUT class 10 match VOIP ip dscp ef
UBNT# set traffic-policy shaper WAN-OUT class 20 bandwidth 40%
UBNT# set traffic-policy shaper WAN-OUT class 20 match CRIT ip dscp af31
UBNT# set traffic-policy shaper WAN-OUT default bandwidth 30%
UBNT# set traffic-policy shaper WAN-OUT default queue-type fair-queue
UBNT# set interfaces ethernet eth0 traffic-policy out WAN-OUT
UBNT# commit
UBNT# save
```

### 5.8 VRRP (HA EdgeRouter)

```shell
UBNT> configure
UBNT# set interfaces ethernet eth2 vif 10 vrrp vrrp-group 10 virtual-address 192.168.10.254/24
UBNT# set interfaces ethernet eth2 vif 10 vrrp vrrp-group 10 priority 200
UBNT# set interfaces ethernet eth2 vif 10 vrrp vrrp-group 10 preempt true
UBNT# set interfaces ethernet eth2 vif 10 vrrp vrrp-group 10 sync-group LAN-HA
UBNT# commit
UBNT# save
UBNT> show vrrp
```

---

## 6. MikroTik RouterOS — déploiement complet en CLI

### 6.1 Interfaces : bridge, VLANs, bonding

```shell
MT> /interface bridge add name=bridge-trunk vlan-filtering=yes comment="Trunk vers switch"
MT> /interface bridge port add bridge=bridge-trunk interface=ether2 comment="Uplink switch"
MT> /interface bridge vlan add bridge=bridge-trunk vlan-ids=10 tagged=bridge-trunk,ether2
MT> /interface bridge vlan add bridge=bridge-trunk vlan-ids=20 tagged=bridge-trunk,ether2
MT> /interface bridge vlan add bridge=bridge-trunk vlan-ids=30 tagged=bridge-trunk,ether2
MT> /interface bridge vlan add bridge=bridge-trunk vlan-ids=40 tagged=bridge-trunk,ether2
MT> /interface bridge vlan add bridge=bridge-trunk vlan-ids=99 tagged=bridge-trunk,ether2
MT> /interface vlan add name=vlan10-data vlan-id=10 interface=bridge-trunk
MT> /interface vlan add name=vlan20-voix vlan-id=20 interface=bridge-trunk
MT> /interface vlan add name=vlan30-invites vlan-id=30 interface=bridge-trunk
MT> /interface vlan add name=vlan40-srv vlan-id=40 interface=bridge-trunk
MT> /interface vlan add name=vlan99-mgmt vlan-id=99 interface=bridge-trunk
MT> /ip address add address=192.168.10.254/24 interface=vlan10-data
MT> /ip address add address=192.168.20.254/24 interface=vlan20-voix
MT> /ip address add address=192.168.30.254/24 interface=vlan30-invites
MT> /ip address add address=192.168.40.254/24 interface=vlan40-srv
MT> /ip address add address=192.168.99.254/24 interface=vlan99-mgmt
MT> /ip address add address=10.255.255.1/32 interface=lo comment="Loopback"
MT> /interface bonding add name=bond-uplink slaves=ether3,ether4 mode=802.3ad comment="LACP"
```

> **VLAN filtering sur bridge** (RouterOS v7) : la méthode moderne et
> performante. L'ancienne méthode « un bridge par VLAN » est dépréciée.

### 6.2 Firewall filter (pare-feu)

```shell
MT> /ip firewall filter add chain=input action=accept connection-state=established,related comment="Retours"
MT> /ip firewall filter add chain=input action=drop connection-state=invalid
MT> /ip firewall filter add chain=input action=accept protocol=icmp comment="Ping diag"
MT> /ip firewall filter add chain=input action=accept src-address=192.168.99.0/24 comment="MGMT"
MT> /ip firewall filter add chain=input action=accept in-interface=vlan10-data dst-port=53 protocol=udp comment="DNS LAN"
MT> /ip firewall filter add chain=input action=drop comment="Tout le reste en input"
MT>
MT> /ip firewall filter add chain=forward action=accept connection-state=established,related
MT> /ip firewall filter add chain=forward action=drop connection-state=invalid
MT> /ip firewall filter add chain=forward action=accept in-interface=vlan10-data out-interface=ether1-WAN comment="DATA vers Internet"
MT> /ip firewall filter add chain=forward action=accept in-interface=vlan30-invites out-interface=ether1-WAN dst-port=80,443 protocol=tcp comment="Invites web seul"
MT> /ip firewall filter add chain=forward action=drop in-interface=vlan30-invites dst-address=192.168.0.0/16 comment="Invites isoles du LAN"
MT> /ip firewall filter add chain=forward action=accept in-interface=vlan10-data out-interface=vlan40-srv comment="DATA vers serveurs"
MT> /ip firewall filter add chain=forward action=drop comment="Reste du forward"
MT> /ip firewall filter print
```

### 6.3 NAT

```shell
MT> /ip firewall nat add chain=srcnat action=masquerade out-interface=ether1-WAN comment="PAT"
MT> /ip firewall nat add chain=dstnat action=dst-nat to-addresses=192.168.40.10 to-ports=443 protocol=tcp dst-port=443 in-interface=ether1-WAN comment="HTTPS vers DMZ"
MT> /ip firewall nat print
MT> /ip firewall connection print where src-address~"192.168.10."
```

### 6.4 Address-lists (l'équivalent des aliases)

```shell
MT> /ip firewall address-list add list=RFC1918 address=10.0.0.0/8
MT> /ip firewall address-list add list=RFC1918 address=172.16.0.0/12
MT> /ip firewall address-list add list=RFC1918 address=192.168.0.0/16
MT> /ip firewall address-list add list=BOGONS address=0.0.0.0/8
MT> /ip firewall address-list add list=BOGONS address=169.254.0.0/16
MT> /ip firewall filter add chain=forward action=drop in-interface=vlan30-invites dst-address-list=RFC1918
```

### 6.5 DHCP + DNS

```shell
MT> /ip pool add name=pool-data ranges=192.168.10.100-192.168.10.200
MT> /ip pool add name=pool-invites ranges=192.168.30.100-192.168.30.200
MT> /ip dhcp-server add name=dhcp-data interface=vlan10-data address-pool=pool-data
MT> /ip dhcp-server network add address=192.168.10.0/24 gateway=192.168.10.254 dns-server=192.168.10.254
MT> /ip dhcp-server add name=dhcp-invites interface=vlan30-invites address-pool=pool-invites
MT> /ip dhcp-server network add address=192.168.30.0/24 gateway=192.168.30.254 dns-server=192.168.30.254
MT> /ip dhcp-server lease add address=192.168.40.10 mac-address=AA:BB:CC:DD:EE:FF comment="SRV-WEB"
MT> /ip dhcp-server enable dhcp-data,dhcp-invites
MT> /ip dns set servers=9.9.9.9,1.1.1.1 allow-remote-requests=yes
MT> /ip dhcp-server lease print
```

### 6.6 Routage : statique, OSPF, BGP (v7)

```shell
MT> /ip route add dst-address=0.0.0.0/0 gateway=203.0.113.9 comment="Defaut FAI-1"
! --- Failover WAN par routage récursif (méthode éprouvée) ---
MT> /ip route add dst-address=9.9.9.9/32 gateway=203.0.113.9 scope=10 comment="Check FAI-1"
MT> /ip route add dst-address=0.0.0.0/0 gateway=9.9.9.9 check-gateway=ping distance=1 comment="Defaut via check"
MT> /ip route add dst-address=1.1.1.1/32 gateway=198.51.100.1 scope=10 comment="Check FAI-2"
MT> /ip route add dst-address=0.0.0.0/0 gateway=1.1.1.1 check-gateway=ping distance=2 comment="Secours"
! --- OSPF v7 ---
MT> /routing ospf instance add name=ospf1 router-id=10.255.255.1
MT> /routing ospf area add name=backbone area-id=0.0.0.0 instance=ospf1
MT> /routing ospf interface-template add networks=192.168.0.0/16 area=backbone
MT> /routing ospf interface-template add networks=10.0.0.0/8 area=backbone
MT> /routing ospf neighbor print
! --- BGP v7 ---
MT> /routing bgp connection add name=FAI1 remote.address=203.0.113.9 remote.as=100 local.address=203.0.113.10 local.as=65001
MT> /routing bgp template set default as=65001 router-id=10.255.255.1
MT> /routing filter rule add chain=bgp-out rule="if (dst in 192.168.0.0/16) {accept} else {reject}"
MT> /routing bgp connection print
```

### 6.7 VPN : WireGuard (natif v7) + IPsec

