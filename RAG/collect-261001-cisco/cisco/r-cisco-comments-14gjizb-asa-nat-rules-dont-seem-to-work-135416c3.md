---
id: collect-261001-cisco/cisco/r-cisco-comments-14gjizb-asa-nat-rules-dont-seem-to-work-135416c3
title: "r-cisco-comments-14gjizb-asa-nat-rules-dont-seem-to-work-135416c3"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/r-cisco-comments-14gjizb-asa-nat-rules-dont-seem-to-work-135416c3.md
source_anchor: ""
source_lines: [1, 137]
sha256: 5d4b7784e4d4b96755f3f2a79a03192ba05f41b065fcdb4d0abc52e72837ff71
---

# r-cisco-comments-14gjizb-asa-nat-rules-dont-seem-to-work-135416c3

ASA NAT rules don't seem to work? 
        
        
        
    
    
    So, im learning about firewalls and NGFW. I was trying to configure this topology. I placed static default routes on "Lan Router" to forward to FW, and on FW to forward to ISP, working fine. Then i tried configuring a nat network object on the FW for the "10.0.1.0/24" net. The problem is that the FW doesnt match any NAT rules when a packet from this subnet comes in. Here is the running config:
lan-firewall#sho nat
Auto NAT Policies (Section 2)
1 (inside) to (outside) source dynamic 10.0.1.0-net interface
    translate_hits = 0, untranslate_hits = 0
2 (inside) to (outside) source dynamic 172.10.10.0-net interface
    translate_hits = 10, untranslate_hits = 5
lan-firewall#show running object
0
object network 10.0.1.0-net
 subnet 10.0.1.0 255.255.255.0
 nat (inside,outside) dynamic interface
0
object network 172.10.10.0-net
 subnet 172.10.10.0 255.255.255.0
 nat (inside,outside) dynamic interface
Hits for inside --> outside from the router (172.10.10.0/24 net) do translate correctly. Am i not understanding something? Thanks in advance.
Edit: Here is the full running-config:
lan-firewall(config)#show running-config
: Saved
:
ASA Version 9.6(1)
!
hostname lan-firewall
names
!
interface GigabitEthernet1/1
 nameif outside
 security-level 0
 ip address 10.10.10.1 255.255.255.0
!
interface GigabitEthernet1/2
 nameif inside
 security-level 0
 ip address 172.10.10.2 255.255.255.0
!
interface GigabitEthernet1/3
 no nameif
 no security-level
 no ip address
 shutdown
!
interface GigabitEthernet1/4
 no nameif
 no security-level
 no ip address
 shutdown
!
interface GigabitEthernet1/5
 no nameif
 no security-level
 no ip address
 shutdown
!
interface GigabitEthernet1/6
 no nameif
 no security-level
 no ip address
 shutdown
!
interface GigabitEthernet1/7
 no nameif
 no security-level
 no ip address
 shutdown
!
interface GigabitEthernet1/8
 no nameif
 no security-level
 no ip address
 shutdown
!
interface Management1/1
 management-only
 no nameif
 no security-level
 no ip address
 shutdown
!
object network 10.0.1.0-net
 subnet 10.0.1.0 255.255.255.0
 nat (inside,outside) dynamic interface
object network 172.10.10.0-net
 subnet 172.10.10.0 255.255.255.0
 nat (inside,outside) dynamic interface
!
route outside 0.0.0.0 0.0.0.0 10.10.10.2 1
route inside 10.0.1.0 255.255.255.0 172.10.10.1 1
!
!
!
!
!
class-map inspection_default
 match default-inspection-traffic
!
policy-map type inspect dns preset_dns_map
 parameters
  message-length maximum 512
policy-map global_policy
 class inspection_default
  inspect dns preset_dns_map
  inspect ftp 
  inspect tftp 
policy-map global-policy
 class inspection_default
!
service-policy global_policy global
!
telnet timeout 5
ssh timeout 5
Edit_2: Hey guys thanks for the tips. Although none of the solutions provided solved the problem, i learned a lot from you. In the end, it seems to be a bug with packet tracer sim. I downloaded Eve-NG and setup the same topology, same configs, and it worked!
Thanks again for the help.
Section des commentaires
Does the ASA have return routes for both networks ?
ASA doesn't have routes for 10.0.1.0/24 network. 172.10.10.0/24 is directly connected.
C 10.0.0.0 255.255.255.0 is directly connected, outside, GigabitEthernet1/1 C 10.10.10.0 255.255.255.0 is directly connected, outside, GigabitEthernet1/1 172.10.0.0/24 is subnetted, 2 subnets C 172.10.0.0 255.255.255.0 is directly connected, inside, GigabitEthernet1/2 C 172.10.10.0 255.255.255.0 is directly connected, inside, GigabitEthernet1/2 S* 0.0.0.0/0 [1/0] via 10.10.10.2
The asa needs to have a route inside back to the router or wherever 10.0.1.0/24 lives
You want to run the packet tracer command next and see where it is failing.
packet-tracer input inside tcp 10.0.1.10 80 8.8.8.8 80
Would work. Look at output and see where it leads you
Funny enough, the packet tracer software does not support packet-tracer command on simulated ASA firewall :/
Now that;s annoying. one of the most useful ASA commands
Based upon the running config, the router traffic in theory shouldn’t be working either due to both the inside and outside interfaces having the same security level (e.g. 0). Try “same-security-traffic permit inter-interface”, that should allow traffic the pass between interfaces with the same security level. Other than that your routing and NAT look fine.
https://www.cisco.com/c/en/us/td/docs/security/asa/asa-cli-reference/S/asa-command-ref-S/sa-shov-commands.html#wp3810489248
Although traffic is getting through both interfaces (it just doesnt get nat-ed) i tried your solution. No success ;-; im going to try and setup eve-ng on my home computer and replicate this topology there and try again. Maybe there is something wrong with packet tracer (wouldnt be the first time to be honest).
Edit: Thanks for the tip on the security level
In theory the configuration with same-security-permit, NAT and routing should work. The only other issue could be that the firewall isn't receiving the traffic. The simplest way to confirm this is to run a packet capture on the ASA to verify if it is receiving the traffic.
The ASA itself can tell you what is going on. Run a packet capture on both inside and outside interface each with trace option enabled.
cap in interface inside tr match ip host 10.0.1.20 host x.x.x.x Cap out interface outside tr match ip host x.x.x.x any
Then simulate traffic. Show cap in Show cap in packet 1 tr detail It should show you what it does with the packet 1. If you see packet go out correctly check out packet capture.
