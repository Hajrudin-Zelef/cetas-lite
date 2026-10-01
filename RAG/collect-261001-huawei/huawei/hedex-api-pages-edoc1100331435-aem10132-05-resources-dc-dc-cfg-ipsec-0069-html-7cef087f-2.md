---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ipsec-0069-html-7cef087f-2
title: "Assign an IP address to each interface on RouterA."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ipsec-0069-html-7cef087f.md
source_anchor: ""
source_lines: [140, 290]
sha256: 871640c1ce6296648f8ed465f36f2b19dd669d5c2495678d323ba667063ec690
---

# Assign an IP address to each interface on RouterA.

[RouterC] interface gigabitethernet 0/0/1
[RouterC-GigabitEthernet0/0/1] ipsec policy policy1
[RouterC-GigabitEthernet0/0/1] quit
# After the configurations are complete, PC C can ping PC A successfully. The data transmitted between PC C and PC A is encrypted.
# Run the display ike sa command on RouterA to check information about the tunnel established with RouterC.
[RouterA] display ike sa
IKE SA information :
   Conn-ID  Peer          VPN   Flag(s)   Phase   RemoteType  RemoteID
  ---------------------------------------------------------------------------
  24366    70.1.1.1:500         RD        v1:2    IP          70.1.1.1
  24274    70.1.1.1:500         RD        v1:1    IP          70.1.1.1
                                   
  Number of IKE SA : 2
  ---------------------------------------------------------------------------
                                                           
  Flag Description:           
  RD--READY   ST--STAYALIVE   RL--REPLACED   FD--FADING   TO--TIMEOUT
  HRT--HEARTBEAT   LKG--LAST KNOWN GOOD SEQ NO.   BCK--BACKED UP
  M--ACTIVE   S--STANDBY   A--ALONE  NEG--NEGOTIATING   
# Run the display ike sa command on RouterC. The command output is displayed as follows:
[RouterC] display ike sa
IKE SA information :
  Conn-ID  Peer          VPN   Flag(s)   Phase   RemoteType  RemoteID
  --------------------------------------------------------------------------
   937    60.1.1.1:500         RD|ST     v1:2    IP          60.1.1.1
   936    60.1.1.1:500         RD|ST     v1:1    IP          60.1.1.1
                                   
  Number of IKE SA : 2
  --------------------------------------------------------------------------
                                                           
  Flag Description:           
  RD--READY   ST--STAYALIVE   RL--REPLACED   FD--FADING   TO--TIMEOUT
  HRT--HEARTBEAT   LKG--LAST KNOWN GOOD SEQ NO.   BCK--BACKED UP
  M--ACTIVE   S--STANDBY   A--ALONE  NEG--NEGOTIATING   
Configuration file of RouterA
#
 sysname RouterA
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
ike peer rut1
 version 1
 pre-shared-key cipher %^%#JvZxR2g8c;a9~FPN~n'$7`DEV&=G(=Et02P/%\*!%^%#
 ike-proposal 5
#
ipsec policy-template use1 10
 ike-peer rut1
 proposal tran1
#
ipsec policy policy1 10 isakmp template use1
#
interface GigabitEthernet0/0/1
 ip address 60.1.1.1 255.255.255.0
 ipsec policy policy1
#
interface GigabitEthernet0/0/2
 ip address 192.168.1.2 255.255.255.0
#
ip route-static 70.1.1.0 255.255.255.0 60.1.1.2
ip route-static 192.168.3.0 255.255.255.0 60.1.1.2
#
return
Configuration file of RouterB
#
 sysname RouterB
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
ike peer rut1
 version 1
 pre-shared-key cipher %^%#K{JG:rWVHPMnf;5\|,GW(Luq'qi8BT4nOj%5W5=)%^%#
 ike-proposal 5
#
ipsec policy-template use1 10
 ike-peer rut1
 proposal tran1
#
ipsec policy policy1 10 isakmp template use1
#
interface GigabitEthernet0/0/1
 ip address 60.1.2.1 255.255.255.0
 ipsec policy policy1
#
interface GigabitEthernet0/0/2
 ip address 192.168.1.3 255.255.255.0
#
ip route-static 70.1.1.0 255.255.255.0 60.1.2.2
ip route-static 192.168.3.0 255.255.255.0 60.1.2.2
#
return
Configuration file of RouterC
#
 sysname RouterC
#
acl number 3002
 rule 5 permit ip source 192.168.3.0 0.0.0.255 destination 192.168.1.0 0.0.0.255
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
ike peer rut1
 version 1
 pre-shared-key cipher %^%#IRFGEiFPJ1$&a'Qy,L*XQL_+*Grq-=yMb}ULZdS6%^%#
 ike-proposal 5
 remote-address 60.1.1.1
 remote-address 60.1.2.1
#
ipsec policy policy1 10 isakmp
 security acl 3002
 ike-peer rut1
 proposal tran1
#
interface GigabitEthernet0/0/1
 ip address 70.1.1.1 255.255.255.0
 ipsec policy policy1
#
interface GigabitEthernet0/0/2
 ip address 192.168.3.2 255.255.255.0
#
ip route-static 0.0.0.0 0.0.0.0 70.1.1.2
#
return
