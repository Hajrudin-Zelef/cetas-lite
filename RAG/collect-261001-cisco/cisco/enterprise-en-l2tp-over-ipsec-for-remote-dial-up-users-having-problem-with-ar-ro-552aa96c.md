---
id: collect-261001-cisco/cisco/enterprise-en-l2tp-over-ipsec-for-remote-dial-up-users-having-problem-with-ar-ro-552aa96c
title: "ike peer Remote_Connect"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-l2tp-over-ipsec-for-remote-dial-up-users-having-problem-with-ar-ro-552aa96c.md
source_anchor: ""
source_lines: [1, 67]
sha256: 71bfc24729b0310056ede5bb67a54b014928d7bc5ca342b8082fd99a4b4a35fe
---

# ike peer Remote_Connect

Hi, I am using the AR2200 router to establish L2TP over IPSec connection to connect from my Laptop at Home to the Office.
My configuration is completed and as per the link below.
https://support.huawei.com/enterprise/en/doc/EDOC1000138063/47f7bc84/example-for-configuring-l2tp-over-ipsec-for-remote-dial-up-users-to-connect-to-the-headquarters
Once the configurations is completed and when I connect the VPN it starts on connecting and after some time it stops automatically without giving error (Windows 10 used in my laptop and I am using the built in client). I dont have direct Public IP Address, it is connected through the Wifi. On the router Side, what I see there is a valid ike sa and IPSec sa for some time but there is no l2tp session.
Router
====
<HQ-R>display ike sa
Conn-ID Peer VPN Flag(s) Phase
----------------------------------------------------------------------------------------------
3189 49.178.52.144:64524 RD|A v1:2
3188 49.178.52.144:64524 RD|A v1:1
Number of IKE SA : 2 --------------------------------------------------------------------------------------------
<HQ-R>display ipsec sa
ipsec sa information:
===============================
Interface: GigabitEthernet0/0/0
-----------------------------
IPSec policy name: "Remote1_Connect"
Sequence number : 10
Acl group : 0
Acl rule : 0
Mode : Template
Connection ID : 3193
Encapsulation mode: Transport
Tunnel local : 87.55.32.210
Tunnel remote : 49.178.52.144
Flow source : 87.55.32.210/255.255.255.255 17/1701
Flow destination : 49.178.52.144/255.255.255.255 17/63890
[Outbound ESP SAs]
SPI: 1597022986 (0x5f30a30a)
Proposal: ESP-ENCRYPT-3DES-192 ESP-AUTH-SHA1
SA remaining key duration (kilobytes/sec): 250000/3598
Outpacket count : 0
Outpacket encap count : 0
Outpacket drop count : 0
Max sent sequence-number: 0
UDP encapsulation used for NAT traversal: Y
[Inbound ESP SAs]
SPI: 3932576243 (0xea6659f3)
Inpacket count : 2
Inpacket decap count : 2
Inpacket drop count : 0
Max received sequence-number: 2
Anti-replay : Enable
Anti-replay window size: 1024
Username and password are correct.
#
l2tp-group 1
undo tunnel authentication
allow l2tp virtual-template 1
interface Virtual-Template1
ppp authentication-mode chap
remote address pool lns
ppp chap user l2tp
ppp chap password cipher %^%#XW@B!!ckF.Vq<LGmzjvQhZ%/U(gBrHpy.x/=a^dT%^%#
ip address 192.168.200.1 255.255.255.0
l2tp-auto-client enable
===================================================================
IPSec Connectivity
# ike peer Remote_Connect
undo version 2
pre-shared-key cipher %^%#S%}b$0&d!8fkeG2+M^8:}1tcR(7m<%29&1;#i=KX%^%#
ike-proposal 2
ipsec policy Remote1_Connect 10 isakmp template Remote_Connect
ipsec policy-template Remote_Connect 10
ike-peer Remote_Connect
proposal Remote_Connect
