---
id: collect-261001-huawei/huawei/questions-42451-56a8ba6f
title: "questions-42451-56a8ba6f"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["asic", "ethernet"]
source: docs/RAG/collect-261001-huawei/questions-42451-56a8ba6f.md
source_anchor: ""
source_lines: [1, 64]
sha256: 34265455e9a409d8cace8fd10e123cc6fc33de110c77972ee4c99c33470af239
---

# questions-42451-56a8ba6f

we are Struggeling with some IPSEC Configuration via a DSL Line on a Huawei AR169 Router.
When the Tunnel is running over an PPPoE Dialer we notice about 70% Packet Loss in the Tunnel. When the Tunnel is running via Ethernet or Cellular Link everything is fine.
We tested this behavior between the AR160 and several vendors on the opposite Side (Cisco Routers, Fortigate Firewalls, further Huawei AR a.s.o.)
Currently we are running on VRP Version 5.160 (updated to this during Troubleshooting process)
Current PPPoE Configuration:
 interface Dialer1
 link-protocol ppp
 ppp chap user XXXXXXXXXXXXX
 ppp chap password XXXXXXXXXXX
 ppp ipcp dns admit-any
 ppp ipcp dns request
 mtu 1456
 ip address ppp-negotiate
 dialer user arweb
 dialer bundle 1
 dialer-group 1
 nat outbound 2000 interface LoopBack 0 
IPSEC Tunnel related configuration
#
ipsec proposal IPSEC-PROP
 esp authentication-algorithm sha2-256 
 esp encryption-algorithm aes-256
#
ike proposal 1
 encryption-algorithm aes-cbc-256
 dh group2
 authentication-algorithm sha2-256
 prf hmac-sha2-256
#
ike peer IKE-PEER v1
 pre-shared-key simple XXXXXXXXXX
 ike-proposal 1
 local-address XXXXXX
 remote-address XXXXXXXX
#
ipsec profile IPSEC-PROFIL
 ike-peer IKE-PEER
 proposal IPSEC-PROP
 sa duration time-based 43200   
#
interface Tunnel0/0/0
 tcp adjust-mss 1300
 ip address unnumbered interface Vlanif1
 tunnel-protocol ipsec
 source LoopBack0
 destination XXXXXXXXXXX
 ipsec profile IPSEC-PROFIL
#
This is my Setup:
SiteA --- AR169 --- Internet --- Cisco C886VA --- SiteB
- Ping between AR169 and C886 (LAN-to-LAN) => OK
- Ping from AR169 to systems in SiteA => OK
- Ping from AR169 to systems in SiteB => NOK
- Ping from C886 to systems in SiteB => OK
- Ping from C886 to systems in SiteA => NOK
- Replacing the DSL Line at AR169 by some other WAN Technology (Ethernet, Cellular) => OK
First i thought of an MTU issue because the DSL is the only media that doesn't support the MTU of 1500 bytes. But that doesn't explain the huge amount of loss during PINGs with standard packet size.
"display ipsec statistics esp" doesn't show any drops.
Can anybody give me a hint how to isolate the root cause of this drops?
Thanks in advance Andreas
--UPDATE--
During troubleshooting we found an interesting behavior. Sending packets that are too large to be transported by tunnel MTU, so that have to be fragmented by the Router, are transported with 100% Success. So: - "ping -l 1363 " => Packet Loss - "ping -l 1364 " => 100% Success
--UPDATE 2--
Doing a packet capture on the AR160 has the Effect that Traffic is running fine. As soon as the Capture stops, we see the PL again. So i think the root cause is to find somewhere in the VPN hardware acceleration (i think that fragmentation and traffic capturing is process-switched by the CPU and not offloaded to an ASIC or some different silicon). Meanwhile we are in contact with some engineers of Huawei. I will keep you informed.
