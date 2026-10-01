---
id: collect-261001-huawei/huawei/eyadnasr-ensp-labs-dcefb46c-5
title: "EyadNasr/eNSP_Labs"
domain: huawei
role: reference
task: reference
actors: ["China"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-huawei/eyadnasr-ensp-labs-dcefb46c.md
source_anchor: ""
source_lines: [1314, 1625]
sha256: d87b631e61dbf4a02ad53f5f613526cc4de50fb0228bf97d3a21262ea904e9f5
---

# EyadNasr/eNSP_Labs

```
[V200R003C00]
#
 sysname Spoke-CE22
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
 local-user admin password cipher %$%$3&&Y8NU#V@$u`z9&SSpSKv-6%$%$
 local-user admin service-type http
 local-user test123 password cipher %$%$LJPQ.P#_I=VRF,-&,:t8KV)\%$%$
 local-user test123 privilege level 15
 local-user test123 ftp-directory flash:
 local-user test123 service-type telnet ssh ftp
#
firewall zone Local
 priority 15
#
interface GigabitEthernet0/0/0
 ip address 192.168.200.3 255.255.255.0 
#
interface GigabitEthernet0/0/1
 ip address 10.0.224.22 255.255.255.0 
#
interface GigabitEthernet0/0/2
 ip address 10.0.222.22 255.255.255.0 
#
interface NULL0
#
interface LoopBack0
 ip address 22.22.22.22 255.255.255.255 
#
bgp 100
 router-id 22.22.22.22
 peer 2.2.2.2 as-number 100 
 peer 2.2.2.2 connect-interface LoopBack0
 peer 4.4.4.6 as-number 123 
 peer 4.4.4.6 ebgp-max-hop 2 
 peer 4.4.4.6 connect-interface LoopBack0
 #
 ipv4-family unicast
  undo synchronization
  network 10.0.224.0 255.255.255.0 
  peer 2.2.2.2 enable
  peer 2.2.2.2 next-hop-local 
  peer 4.4.4.6 enable
#
 sftp server enable 
 stelnet server enable 
#
ip as-path-filter noLoop deny 123 123 123 123
#
ip route-static 2.2.2.2 255.255.255.255 10.0.222.2
ip route-static 4.4.4.6 255.255.255.255 10.0.224.4
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

## Fichier : Hub&Spokes MPLS InterAS L3VPN  - BGP - same AS - Option A/Configs exported/Spoke-PE11.cfg

```
[V200R003C00]
#
 sysname Spoke-PE11
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
ip vpn-instance VPN
 ipv4-family
  route-distinguisher 100:1
  vpn-target 100:1 export-extcommunity
  vpn-target 100:1 import-extcommunity
#
mpls lsr-id 11.11.11.11
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
 local-user admin password cipher %$%$'V*6+/AYJ$WH/UX*6}59K)IS%$%$
 local-user admin service-type http
 local-user test123 password cipher %$%$`@d")$0<}"R~o^0Y;X7EKVKr%$%$
 local-user test123 privilege level 15
 local-user test123 ftp-directory flash:
 local-user test123 service-type telnet ssh ftp
#
firewall zone Local
 priority 15
#
interface GigabitEthernet0/0/0
 ip binding vpn-instance VPN
 ip address 10.0.13.11 255.255.255.0 
#
interface GigabitEthernet0/0/1
 ip address 10.0.115.11 255.255.255.0 
 mpls
 mpls ldp
#
interface GigabitEthernet0/0/2
 ip address 192.168.200.14 255.255.255.0 
#
interface NULL0
#
interface LoopBack0
 ip address 11.11.11.11 255.255.255.255 
#
interface LoopBack1
 ip binding vpn-instance VPN
 ip address 11.11.11.12 255.255.255.255 
#
bgp 321
 router-id 11.11.11.11
 peer 33.33.33.33 as-number 321 
 peer 33.33.33.33 connect-interface LoopBack0
 #
 ipv4-family unicast
  undo synchronization
  peer 33.33.33.33 enable
 # 
 ipv4-family vpnv4
  policy vpn-target
  peer 33.33.33.33 enable
 #
 ipv4-family vpn-instance VPN 
  peer 1.1.1.1 as-number 100 
  peer 1.1.1.1 ebgp-max-hop 2 
  peer 1.1.1.1 connect-interface LoopBack1
  peer 1.1.1.1 substitute-as
#
ospf 1 router-id 11.11.11.11 
 area 0.0.0.0 
  network 10.0.115.11 0.0.0.0 
  network 11.11.11.11 0.0.0.0 
#
 sftp server enable 
 stelnet server enable 
#
ip route-static vpn-instance VPN 1.1.1.1 255.255.255.255 10.0.13.1
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

## Fichier : Hub&Spokes MPLS InterAS L3VPN  - BGP - same AS - Option A/Configs exported/Spoke-PE3.cfg

```
[V200R003C00]
#
 sysname Spoke-PE3
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
ip vpn-instance VPN
 ipv4-family
  route-distinguisher 100:3
  vpn-target 100:1 export-extcommunity
  vpn-target 300:1 import-extcommunity
#
mpls lsr-id 3.3.3.3
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
 local-user admin password cipher %$%$,N3U=TDp~@1b#_6*pj<RK~Lv%$%$
 local-user admin service-type http
 local-user test123 password cipher %$%$jnoH8(8ZOQt*ESXv[P;"KWM2%$%$
 local-user test123 privilege level 15
 local-user test123 ftp-directory flash:
 local-user test123 service-type telnet ssh ftp
#
firewall zone Local
 priority 15
#
interface GigabitEthernet0/0/0
 ip address 10.0.35.3 255.255.255.0 
 mpls
 mpls ldp
#
interface GigabitEthernet0/0/1
#
interface GigabitEthernet0/0/1.1
 dot1q termination vid 1
 ip binding vpn-instance VPN
 ip address 10.0.133.3 255.255.255.0 
 arp broadcast enable
#
interface GigabitEthernet0/0/2
 ip address 192.168.200.11 255.255.255.0 
#
interface NULL0
#
interface LoopBack0
 ip address 3.3.3.3 255.255.255.255 
#
interface LoopBack1
 ip binding vpn-instance VPN
 ip address 3.3.3.4 255.255.255.255 
#
bgp 123
 router-id 3.3.3.3
 peer 26.26.26.26 as-number 123 
 peer 26.26.26.26 connect-interface LoopBack0
 #
 ipv4-family unicast
  undo synchronization
  peer 26.26.26.26 enable
 # 
 ipv4-family vpnv4
  policy vpn-target
  peer 26.26.26.26 enable
 #
 ipv4-family vpn-instance VPN 
  router-id 3.3.3.4
  peer 33.33.33.34 as-number 321 
  peer 33.33.33.34 ebgp-max-hop 2 
  peer 33.33.33.34 connect-interface LoopBack1
#
ospf 1 router-id 3.3.3.3 
 area 0.0.0.0 
  network 3.3.3.3 0.0.0.0 
  network 10.0.35.3 0.0.0.0 
#
 sftp server enable 
 stelnet server enable 
#
ip route-static vpn-instance VPN 33.33.33.34 255.255.255.255 10.0.133.33
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

## Fichier : Hub&Spokes MPLS InterAS L3VPN  - BGP - same AS - Option A/Configs exported/Spoke-PE33.cfg

