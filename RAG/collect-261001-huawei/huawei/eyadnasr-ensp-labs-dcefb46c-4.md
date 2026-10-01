---
id: collect-261001-huawei/huawei/eyadnasr-ensp-labs-dcefb46c-4
title: "EyadNasr/eNSP_Labs"
domain: huawei
role: reference
task: reference
actors: ["China"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-huawei/eyadnasr-ensp-labs-dcefb46c.md
source_anchor: ""
source_lines: [999, 1313]
sha256: 8b977b11251e7b5a8ed0715850127e9c7e29589c0caabe85efaf26fdeb66baf6
---

# EyadNasr/eNSP_Labs

```
[V200R003C00]
#
 sysname PE66
 ftp server enable
#
 snmp-agent local-engineid 800007DB03000000000000
 snmp-agent 
#
 clock timezone China-Standard-Time minus 08:00:00
#
portal local-server load portalpage.zip
#
 drop illegal-mac alarm
#
 lldp enable 
#
 set cpu-usage threshold 80 restore 75
#
ip vpn-instance site1
 ipv4-family
  route-distinguisher 400:1
  vpn-target 100:1 export-extcommunity
  vpn-target 300:1 import-extcommunity
#
ip vpn-instance site2
 ipv4-family
  route-distinguisher 400:2
  vpn-target 200:1 export-extcommunity
  vpn-target 300:1 import-extcommunity
#
mpls lsr-id 66.66.66.66
mpls
#
mpls ldp
#
#
aaa 
 authentication-scheme default
 authorization-scheme default
 accounting-scheme default
 domain default 
 domain default_admin 
 local-user admin password cipher %$%$d3+R.%}wxD6,V}0TMzE-K=*7%$%$
 local-user admin service-type http
 local-user test123 password cipher %$%$#<Jn*:7sJ/Sm%iF;pbw/KWB<%$%$
 local-user test123 privilege level 15
 local-user test123 ftp-directory flash:
 local-user test123 service-type telnet ssh ftp
#
firewall zone Local
 priority 15
#
interface GigabitEthernet0/0/0
#
interface GigabitEthernet0/0/0.1
 dot1q termination vid 1
 ip binding vpn-instance site1
 ip address 10.0.69.66 255.255.255.0 
 arp broadcast enable
#
interface GigabitEthernet0/0/0.2
 dot1q termination vid 2
 ip binding vpn-instance site2
 ip address 10.1.69.66 255.255.255.0 
 arp broadcast enable
#
interface GigabitEthernet0/0/1
 ip address 192.168.200.10 255.255.255.0 
#
interface GigabitEthernet0/0/2
 ip address 10.0.166.66 255.255.255.0 
 mpls
 mpls ldp
#
interface NULL0
#
interface LoopBack0
 ip address 66.66.66.66 255.255.255.255 
#
interface LoopBack1
 ip binding vpn-instance site1
 ip address 66.66.66.67 255.255.255.255 
#
interface LoopBack2
 ip binding vpn-instance site2
 ip address 66.66.66.68 255.255.255.255 
#
bgp 312
 router-id 66.66.66.66
 peer 6.6.6.6 as-number 312 
 peer 6.6.6.6 connect-interface LoopBack0
 #
 ipv4-family unicast
  undo synchronization
  peer 6.6.6.6 enable
 # 
 ipv4-family vpnv4
  policy vpn-target
  peer 6.6.6.6 enable
  peer 6.6.6.6 allow-as-loop 2
 #
 ipv4-family vpn-instance site1 
  router-id 66.66.66.67
  peer 26.26.26.27 as-number 123 
  peer 26.26.26.27 ebgp-max-hop 2 
  peer 26.26.26.27 connect-interface LoopBack1
  peer 26.26.26.27 substitute-as
  peer 26.26.26.27 advertise-ext-community
 #
 ipv4-family vpn-instance site2 
  router-id 66.66.66.68
  peer 26.26.26.28 as-number 123 
  peer 26.26.26.28 ebgp-max-hop 2 
  peer 26.26.26.28 connect-interface LoopBack2
  peer 26.26.26.28 substitute-as
  peer 26.26.26.28 advertise-ext-community
#
ospf 1 router-id 66.66.66.66 
 area 0.0.0.0 
  network 10.0.166.66 0.0.0.0 
  network 66.66.66.66 0.0.0.0 
#
 sftp server enable 
 stelnet server enable 
#
ip route-static vpn-instance site1 26.26.26.27 255.255.255.255 10.0.69.9
ip route-static vpn-instance site2 26.26.26.28 255.255.255.255 10.1.69.9
#
user-interface con 0
 authentication-mode password
user-interface vty 0 4
 authentication-mode aaa
 protocol inbound all
user-interface vty 16 20
#
wlan ac
#
return
```

## Fichier : Hub&Spokes MPLS InterAS L3VPN  - BGP - same AS - Option A/Configs exported/Spoke-CE1.cfg

```
[V200R003C00]
#
 sysname Spoke_CE1
 ftp server enable
#
 snmp-agent local-engineid 800007DB03000000000000
 snmp-agent 
#
 clock timezone China-Standard-Time minus 08:00:00
#
portal local-server load portalpage.zip
#
 drop illegal-mac alarm
#
 lldp enable 
#
 set cpu-usage threshold 80 restore 75
#
aaa 
 authentication-scheme default
 authorization-scheme default
 accounting-scheme default
 domain default 
 domain default_admin 
 local-user admin password cipher %$%$ClHXK<c%F;hR#J@W=n"2KkZ(%$%$
 local-user admin service-type http
 local-user test123 password cipher %$%$0(C=<\w]lG=%|Q"YSTu=KV$>%$%$
 local-user test123 privilege level 15
 local-user test123 ftp-directory flash:
 local-user test123 service-type telnet ssh ftp
#
firewall zone Local
 priority 15
#
interface GigabitEthernet0/0/0
 ip address 10.0.13.1 255.255.255.0 
#
interface GigabitEthernet0/0/1
 ip address 192.168.200.13 255.255.255.0 
#
interface GigabitEthernet0/0/2
#
interface NULL0
#
interface LoopBack0
 ip address 1.1.1.1 255.255.255.255 
#
interface LoopBack1
 ip address 192.168.1.1 255.255.255.0 
#
bgp 100
 router-id 1.1.1.1
 peer 11.11.11.12 as-number 321 
 peer 11.11.11.12 ebgp-max-hop 2 
 peer 11.11.11.12 connect-interface LoopBack0
 #
 ipv4-family unicast
  undo synchronization
  network 10.0.13.0 255.255.255.0 
  network 192.168.1.0 
  peer 11.11.11.12 enable
#
 sftp server enable 
 stelnet server enable 
#
ip route-static 11.11.11.12 255.255.255.255 10.0.13.11
#
user-interface con 0
 authentication-mode password
user-interface vty 0 4
 authentication-mode aaa
 protocol inbound all
user-interface vty 16 20
#
wlan ac
#
return
```

## Fichier : Hub&Spokes MPLS InterAS L3VPN  - BGP - same AS - Option A/Configs exported/Spoke-CE2.cfg

```
[V200R003C00]
#
 sysname Spoke_CE2
 ftp server enable
#
 snmp-agent local-engineid 800007DB03000000000000
 snmp-agent 
#
 clock timezone China-Standard-Time minus 08:00:00
#
portal local-server load portalpage.zip
#
 drop illegal-mac alarm
#
 lldp enable 
#
 set cpu-usage threshold 80 restore 75
#
aaa 
 authentication-scheme default
 authorization-scheme default
 accounting-scheme default
 domain default 
 domain default_admin 
 local-user admin password cipher %$%$5RqT"Qx8g.cd[`$s[a#)K}Vq%$%$
 local-user admin service-type http
 local-user test123 password cipher %$%$:hqx(Y59$515cH~u"6ODKW\3%$%$
 local-user test123 privilege level 15
 local-user test123 ftp-directory flash:
 local-user test123 service-type telnet ssh ftp
#
firewall zone Local
 priority 15
#
interface GigabitEthernet0/0/0
 ip address 10.0.24.2 255.255.255.0 
#
interface GigabitEthernet0/0/1
 ip address 192.168.200.12 255.255.255.0 
#
interface GigabitEthernet0/0/2
 ip address 10.0.222.2 255.255.255.0 
#
interface NULL0
#
interface LoopBack0
 ip address 2.2.2.2 255.255.255.255 
#
interface LoopBack1
 ip address 172.16.0.1 255.255.255.0 
#
bgp 100
 router-id 2.2.2.2
 peer 4.4.4.5 as-number 123 
 peer 4.4.4.5 ebgp-max-hop 2 
 peer 4.4.4.5 connect-interface LoopBack0
 peer 22.22.22.22 as-number 100 
 peer 22.22.22.22 connect-interface LoopBack0
 #
 ipv4-family unicast
  undo synchronization
  network 10.0.24.0 255.255.255.0 
  network 172.16.0.0 255.255.255.0 
  peer 4.4.4.5 enable
  peer 22.22.22.22 enable
  peer 22.22.22.22 next-hop-local 
#
 sftp server enable 
 stelnet server enable 
#
ip as-path-filter noLoop deny 123 123 123 123
#
ip route-static 4.4.4.5 255.255.255.255 10.0.24.4
ip route-static 22.22.22.22 255.255.255.255 10.0.222.22
#
user-interface con 0
 authentication-mode password
user-interface vty 0 4
 authentication-mode aaa
 protocol inbound all
user-interface vty 16 20
#
wlan ac
#
return
```

## Fichier : Hub&Spokes MPLS InterAS L3VPN  - BGP - same AS - Option A/Configs exported/Spoke-CE22.cfg

