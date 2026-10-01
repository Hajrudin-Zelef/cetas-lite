---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ipsec-0063-html-e03cd1d8-2
title: "Assign an IP address to an interface on RouterA."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ipsec-0063-html-e03cd1d8.md
source_anchor: ""
source_lines: [118, 209]
sha256: 54c77f5946847e48f3382367a8d85c8d74a8f4878d7064568b92c4debf12132c
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
ike local-name rta
#
acl number 3101
 rule 5 permit ip source 10.1.0.0 0.0.0.255 destination 10.2.0.0 0.0.0.255
#
ipsec proposal tran1
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
ike peer rta
 version 1
 exchange-mode aggressive
 pre-shared-key cipher %^%#JvZxR2g8c;a9~FPN~n'$7`DEV&=G(=Et02P/%\*!%^%#
 ike-proposal 5
 local-id-type fqdn
 remote-id rtb
 remote-address 1.2.0.1
#
ipsec policy policy1 10 isakmp
 security acl 3101
 ike-peer rta
 proposal tran1
#
interface GigabitEthernet1/0/0
 ip address 192.168.0.2 255.255.255.0
 ipsec policy policy1
#
interface GigabitEthernet2/0/0
 ip address 10.1.0.1 255.255.255.0
#
ip route-static 0.0.0.0 0.0.0.0 192.168.0.1
#
return                                                                               
Configuration file of RouterB
#
 sysname RouterB
#
ike local-name rtb
#
ipsec proposal tran1
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
ike peer rtb
 version 1
 exchange-mode aggressive
 pre-shared-key cipher %^%#K{JG:rWVHPMnf;5\|,GW(Luq'qi8BT4nOj%5W5=)%^%#
 ike-proposal 5
 local-id-type fqdn
 remote-id rta
#
ipsec policy-template temp1 10
 ike-peer rtb
 proposal tran1
#
ipsec policy policy1 10 isakmp template temp1
#
interface GigabitEthernet1/0/0
 ip address 1.2.0.1 255.255.255.0
 ipsec policy policy1
#
interface GigabitEthernet2/0/0
 ip address 10.2.0.1 255.255.255.0
#
ip route-static 10.1.0.0 255.255.255.0 1.2.0.2
ip route-static 192.168.0.0 255.255.255.0 1.2.0.2
#
return
