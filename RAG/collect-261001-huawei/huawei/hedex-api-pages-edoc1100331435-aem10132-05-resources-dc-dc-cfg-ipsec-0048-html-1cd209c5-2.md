---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ipsec-0048-html-1cd209c5-2
title: "Assign an IP address to an interface on RouterA."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ipsec-0048-html-1cd209c5.md
source_anchor: ""
source_lines: [122, 220]
sha256: c64947f97b8829f664528b1cf825c8d89ea4aa7268f9487b7443a862cdb0f36a
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
acl number 3101
 rule 5 permit ip source 10.1.1.0 0.0.0.255 destination 10.1.2.0 0.0.0.255
#
ipsec proposal prop
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
ike peer rut
 version 1
 pre-shared-key cipher %^%#JvZxR2g8c;a9~FPN~n'$7`DEV&=G(=Et02P/%\*!%^%#
 ike-proposal 5
 remote-address 60.1.1.1
#
ipsec policy policy1 10 isakmp
 security acl 3101
 ike-peer rut
 proposal prop
#
ipsec policy policy1 shared local-interface LoopBack0
#
interface GigabitEthernet1/0/0
 ip address 70.1.1.1 255.255.255.0
 ipsec policy policy1
#
interface GigabitEthernet2/0/0
 ip address 80.1.1.1 255.255.255.0
 ipsec policy policy1
#
interface GigabitEthernet3/0/0
 ip address 10.1.1.1 255.255.255.0
#
interface LoopBack0
 ip address 1.1.1.1 255.255.255.255
#
ip route-static 10.1.2.0 255.255.255.0 70.1.1.2 preference 10
ip route-static 10.1.2.0 255.255.255.0 80.1.1.2 preference 20
ip route-static 60.1.1.0 255.255.255.0 70.1.1.2 preference 10
ip route-static 60.1.1.0 255.255.255.0 80.1.1.2 preference 20
#
return
Configuration file of RouterB
#
 sysname RouterB
#
acl number 3101
 rule 5 permit ip source 10.1.2.0 0.0.0.255 destination 10.1.1.0 0.0.0.255
#
ipsec proposal prop
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
ike peer rut
 version 1
 pre-shared-key cipher %^%#K{JG:rWVHPMnf;5\|,GW(Luq'qi8BT4nOj%5W5=)%^%#
 ike-proposal 5
 remote-address 1.1.1.1
#
ipsec policy policy1 10 isakmp
 security acl 3101
 ike-peer rut
 proposal prop
#
interface GigabitEthernet1/0/0
 ip address 60.1.1.1 255.255.255.0
 ipsec policy policy1
#
interface GigabitEthernet3/0/0
 ip address 10.1.2.1 255.255.255.0
#
ip route-static 1.1.1.1 255.255.255.255 60.1.1.2
ip route-static 10.1.1.0 255.255.255.0 60.1.1.2
ip route-static 70.1.1.0 255.255.255.0 60.1.1.2
ip route-static 80.1.1.0 255.255.255.0 60.1.1.2
#
return
