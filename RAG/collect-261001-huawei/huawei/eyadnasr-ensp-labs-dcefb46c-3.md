---
id: collect-261001-huawei/huawei/eyadnasr-ensp-labs-dcefb46c-3
title: "EyadNasr/eNSP_Labs"
domain: huawei
role: reference
task: reference
actors: ["China"]
dates: []
keywords: ["agent", "cost"]
source: docs/RAG/collect-261001-huawei/eyadnasr-ensp-labs-dcefb46c.md
source_anchor: ""
source_lines: [646, 998]
sha256: 4f0602c8b43f1c8779adba40ac9b2857100f46ce216d489a13767ff734fa8c54
---

# EyadNasr/eNSP_Labs

```
[V200R003C00]
#
 sysname P16
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
mpls lsr-id 16.16.16.16
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
 local-user admin password cipher %$%$}y4wK8}@n(~'o*09J9u&K:nZ%$%$
 local-user admin service-type http
 local-user test123 password cipher %$%$lE+-G%Vc'0vk;OS9AZ;GKWSE%$%$
 local-user test123 privilege level 15
 local-user test123 ftp-directory flash:
 local-user test123 service-type telnet ssh ftp
#
firewall zone Local
 priority 15
#
interface GigabitEthernet0/0/0
 ip address 192.168.200.2 255.255.255.0 
#
interface GigabitEthernet0/0/1
 ip address 10.0.16.16 255.255.255.0 
 mpls
 mpls ldp
#
interface GigabitEthernet0/0/2
 ip address 10.0.166.16 255.255.255.0 
 mpls
 mpls ldp
#
interface NULL0
#
interface LoopBack0
 ip address 16.16.16.16 255.255.255.255 
#
ospf 1 router-id 16.16.16.16 
 area 0.0.0.0 
  network 10.0.16.16 0.0.0.0 
  network 10.0.166.16 0.0.0.0 
  network 16.16.16.16 0.0.0.0 
#
 sftp server enable 
 stelnet server enable 
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

## Fichier : Hub&Spokes MPLS InterAS L3VPN  - BGP - same AS - Option A/Configs exported/P5.cfg

```
[V200R003C00]
#
 sysname P5
 ftp server enable
#
 board add 0/3 1GEC 
 board add 0/4 1GEC 
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
mpls lsr-id 5.5.5.5
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
 local-user admin password cipher %$%$JTOW#Wouq"2S5LTA{y^)KMtk%$%$
 local-user admin service-type http
 local-user test123 password cipher %$%$/odQ9*N#EXnnD"1h`{&3KWT)%$%$
 local-user test123 privilege level 15
 local-user test123 ftp-directory flash:
 local-user test123 service-type telnet ssh ftp
#
firewall zone Local
 priority 15
#
interface GigabitEthernet0/0/0
 ip address 10.0.59.5 255.255.255.0 
 ospf cost 2
 mpls
 mpls ldp
#
interface GigabitEthernet0/0/1
 ip address 10.0.35.5 255.255.255.0 
 mpls
 mpls ldp
#
interface GigabitEthernet0/0/2
 ip address 10.0.45.5 255.255.255.0 
 mpls
 mpls ldp
#
interface GigabitEthernet3/0/0
 ip address 192.168.200.8 255.255.255.0 
#
interface GigabitEthernet4/0/0
 ip address 10.0.105.5 255.255.255.0 
 mpls
 mpls ldp
#
interface NULL0
#
interface LoopBack0
 ip address 5.5.5.5 255.255.255.255 
#
ospf 1 router-id 5.5.5.5 
 area 0.0.0.0 
  network 5.5.5.5 0.0.0.0 
  network 10.0.35.5 0.0.0.0 
  network 10.0.45.5 0.0.0.0 
  network 10.0.59.5 0.0.0.0 
  network 10.0.105.5 0.0.0.0 
#
 sftp server enable 
 stelnet server enable 
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

## Fichier : Hub&Spokes MPLS InterAS L3VPN  - BGP - same AS - Option A/Configs exported/P8.cfg

```
[V200R003C00]
#
 sysname P8
 ftp server enable
#
 board add 0/4 1GEC 
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
mpls lsr-id 8.8.8.8
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
 local-user admin password cipher %$%$K8m.Nt84DZ}e#<0`8bmE3Uw}%$%$
 local-user admin service-type http
 local-user test123 password cipher %$%$Ydz@6j~Kr"$d#C8j5(f,KWEa%$%$
 local-user test123 privilege level 15
 local-user test123 ftp-directory flash:
 local-user test123 service-type telnet ssh ftp
#
firewall zone Local
 priority 15
#
interface GigabitEthernet0/0/0
 ip address 192.168.200.7 255.255.255.0 
#
interface GigabitEthernet0/0/1
 ip address 10.0.89.8 255.255.255.0 
 mpls
 mpls ldp
#
interface GigabitEthernet0/0/2
 ip address 10.0.108.8 255.255.255.0 
 ospf cost 2
 mpls
 mpls ldp
#
interface GigabitEthernet4/0/0
 ip address 10.0.26.8 255.255.255.0 
 mpls
 mpls ldp
#
interface NULL0
#
interface LoopBack0
 ip address 8.8.8.8 255.255.255.255 
#
ospf 1 router-id 8.8.8.8 
 area 0.0.0.0 
  network 8.8.8.8 0.0.0.0 
  network 10.0.26.8 0.0.0.0 
  network 10.0.89.8 0.0.0.0 
  network 10.0.108.8 0.0.0.0 
#
 sftp server enable 
 stelnet server enable 
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

## Fichier : Hub&Spokes MPLS InterAS L3VPN  - BGP - same AS - Option A/Configs exported/P9.cfg

```
[V200R003C00]
#
 sysname P9
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
mpls lsr-id 9.9.9.9
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
 local-user admin password cipher %$%$Mv[k6a;02*31,-%3{dMPKH.m%$%$
 local-user admin service-type http
 local-user test123 password cipher %$%$<vry--aLRX5tEU2*Mjm*KWq,%$%$
 local-user test123 privilege level 15
 local-user test123 ftp-directory flash:
 local-user test123 service-type telnet ssh ftp
#
firewall zone Local
 priority 15
#
interface GigabitEthernet0/0/0
 ip address 10.0.59.9 255.255.255.0 
 mpls
 mpls ldp
#
interface GigabitEthernet0/0/1
 ip address 10.0.89.9 255.255.255.0 
 mpls
 mpls ldp
#
interface GigabitEthernet0/0/2
 ip address 192.168.200.4 255.255.255.0 
#
interface NULL0
#
interface LoopBack0
 ip address 9.9.9.9 255.255.255.255 
#
ospf 1 router-id 9.9.9.9 
 area 0.0.0.0 
  network 9.9.9.9 0.0.0.0 
  network 10.0.59.9 0.0.0.0 
  network 10.0.89.9 0.0.0.0 
#
 sftp server enable 
 stelnet server enable 
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

## Fichier : Hub&Spokes MPLS InterAS L3VPN  - BGP - same AS - Option A/Configs exported/PE66.cfg

