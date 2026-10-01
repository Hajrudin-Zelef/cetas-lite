---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ipsec-0053-html-287aa4a2-2
title: "Assign an IP address to an interface on RouterA."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ipsec-0053-html-287aa4a2.md
source_anchor: ""
source_lines: [109, 192]
sha256: 4f82c6cdb44317eea386c03e92c754db285c2141ff5889e58c4590864f445ee4
---

# Assign an IP address to an interface on RouterA.

  Flag Description:           
  RD--READY   ST--STAYALIVE   RL--REPLACED   FD--FADING   TO--TIMEOUT
  HRT--HEARTBEAT   LKG--LAST KNOWN GOOD SEQ NO.   BCK--BACKED UP
  M--ACTIVE   S--STANDBY   A--ALONE  NEG--NEGOTIATING   
Configuration file of RouterA
#
 sysname RouterA
#
ipsec efficient-vpn evpn mode client
 remote-address 60.1.2.1 v1
 pre-shared-key cipher %^%#JvZxR2g8c;a9~FPN~n'$7`DEV&=G(=Et02P/%\*!%^%#
 dh group14
#
interface GigabitEthernet1/0/0
 ip address 60.1.1.1 255.255.255.0
 ipsec efficient-vpn evpn
#
interface GigabitEthernet2/0/0
 ip address 10.1.1.1 255.255.255.0
#
ip route-static 60.1.2.0 255.255.255.0 60.1.1.2
ip route-static 10.1.2.0 255.255.255.0 60.1.1.2
#
return
Configuration file of RouterB
#
 sysname RouterB
#
ipsec authentication sha2 compatible enable
#
dhcp enable
#
ipsec proposal prop1
 esp authentication-algorithm sha2-256
 esp encryption-algorithm aes-128
#
ike proposal 5
 encryption-algorithm aes-128
 dh group14
 authentication-algorithm sha2-256
 authentication-method pre-share
 integrity-algorithm hmac-sha2-256
 prf hmac-sha2-256
#
ike peer rut3
 version 1
 exchange-mode aggressive
 pre-shared-key cipher %^%#K{JG:rWVHPMnf;5\|,GW(Luq'qi8BT4nOj%5W5=)%^%#
 ike-proposal 5
 service-scheme schemetest
#
ipsec policy-template temp1 10
 ike-peer rut3
 proposal prop1
#
ipsec policy policy1 10 isakmp template temp1
#
dhcp server group dhcp-ser1
 dhcp-server 10.1.3.2 0
 gateway 100.1.1.3
#
aaa
 service-scheme schemetest
  dns 2.2.2.2
  dns 2.2.2.3 secondary
  dhcp-server group dhcp-ser1
  wins 3.3.3.2
  wins 3.3.3.3 secondary
  dns-name mydomain.com.cn
#
interface GigabitEthernet1/0/0
 ip address 60.1.2.1 255.255.255.0
 ipsec policy policy1
#
interface GigabitEthernet2/0/0
 ip address 10.1.2.1 255.255.255.0
#
interface GigabitEthernet3/0/0  ip address 10.1.3.1 255.255.255.0 # interface GigabitEthernet4/0/0  ip address 100.1.1.3 255.255.255.0
#
ip route-static 60.1.1.0 255.255.255.0 60.1.2.2
ip route-static 10.1.1.0 255.255.255.0 60.1.2.2
ip route-static 100.1.1.0 255.255.255.0 60.1.2.2
#
return
