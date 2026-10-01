---
id: collect-261001-cisco/cisco/questions-7014-edit-cisco-ios-acl-without-hanging-the-router-901c98ae-1
title: "questions-7014-edit-cisco-ios-acl-without-hanging-the-router-901c98ae"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-cisco/questions-7014-edit-cisco-ios-acl-without-hanging-the-router-901c98ae.md
source_anchor: ""
source_lines: [1, 176]
sha256: b4fe8f231ce5ad7e2f2b6a8ec76ad4bb8cdbe79de94a0aacfa53ba817e827018
---

# questions-7014-edit-cisco-ios-acl-without-hanging-the-router-901c98ae

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
7
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a Cisco Router 1941 Series and I want to make some changes to the access-list on the router. I got to know I will have to delete one of the list and then re-add everything again with my additional ones included while stating the deny ones at the end of it. On removing, it went well but on adding, it hangs after the first statement. And whenever I try pasting the code at once, it hangs too.
I tried copying the file to tftp, make some edits and then bring it back, I got some errors with the certificate signing side of it making some complaints.. How can I solve the problem with the ACL?
[Edit]
At the present, this is what I have,
interface GigabitEthernet0/0
description ### WAN INTERFACE ###
ip address xxx.xxx.xxx.xxx 255.xxx.xxx.xxx
ip flow ingress
ip nat outside
ip virtual-reassembly
duplex full
speed 100
no cdp enable
!
interface GigabitEthernet0/1
description ### LAN INTERFACE ###
no ip address
ip flow ingress
ip virtual-reassembly
duplex auto
speed auto
!
interface GigabitEthernet0/1.1
description ### 1st FLR NETWORK ###
encapsulation dot1Q 1 native
ip address 192.168.1.1 255.255.255.0
ip access-group 110 in
ip accounting output-packets
ip flow ingress
ip nat inside
ip virtual-reassembly
!
interface GigabitEthernet0/1.2
description ### GROUND FLR NETWORK ###
encapsulation dot1Q 2
ip address 192.168.2.1 255.255.255.0
ip access-group 110 in
ip accounting output-packets
ip flow ingress
ip nat inside
ip virtual-reassembly
!
ip forward-protocol nd
!
ip http server
ip http access-class 23
ip http authentication local
ip http secure-server
ip http timeout-policy idle 60 life 86400 requests 10000
ip flow-export source GigabitEthernet0/1
ip flow-export version 9
ip flow-export destination 192.168.1.120 9996
ip flow-top-talkers
top 10
sort-by bytes
!
ip nat source static 192.168.1.19 interface Loopback100
ip nat inside source list NAT interface GigabitEthernet0/0 overload
ip route 0.0.0.0 0.0.0.0 xxx.vvv.vvv.vvv
!
ip access-list standard NAT
permit 192.168.1.0 0.0.0.255
permit 192.168.2.0 0.0.0.255
!
access-list 110 permit ip 192.168.1.0 0.0.0.255 host 192.168.2.44
access-list 110 permit ip host 192.168.2.44 192.168.1.0 0.0.0.255
access-list 110 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.18
access-list 110 permit ip host 192.168.1.18 192.168.2.0 0.0.0.255
access-list 110 deny ip 192.168.1.0 0.0.0.255 192.168.2.0 0.0.0.255
access-list 110 permit udp any any
access-list 110 permit ip any any
And I would like to add additional ones to it... like this...
access-list 110 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.120
access-list 110 permit ip host 192.168.1.120 192.168.2.0 0.0.0.255
access-list 110 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.222
access-list 110 permit ip host 192.168.1.222 192.168.2.0 0.0.0.255
The reason your router hangs is because you're editing the ACL through the same interface that it's applied to. Usually what happens is you wind up blocking yourself accidentally.
Note that I make a suggestion for reorganizing your ACLs at the bottom of this answer, although most people will be interested in this first section (the easiest way to change ACLs).
The easiest way to change Cisco IOS ACLs
Newer Cisco IOS images (such as the image on your Cisco 1941) support IP access-list line numbering, which means you can modify your ACL without removing it...
I will use the example of your original ACL 110, although as I mentioned you should really consider reorganizing your ACLs into two different ACLs.
Step 1, ACL 110 already on the router:
LAB_RTR#sh access-list 110
Extended IP access list 110
10 permit ip 192.168.1.0 0.0.0.255 host 192.168.2.44
20 permit ip host 192.168.2.44 192.168.1.0 0.0.0.255
30 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.18
40 permit ip host 192.168.1.18 192.168.2.0 0.0.0.255
50 deny ip 192.168.1.0 0.0.0.255 192.168.2.0 0.0.0.255
60 permit udp any any
70 permit ip any any
Take note of the line numbers that Cisco automatically added to ACL 110. Remember that you need to add your ACL entries before line 50, due to the deny statement.
Step 2, Modify ACL 110 in place with new entries, starting after line 40:
LAB_RTR#conf t
Enter configuration commands, one per line. End with CNTL/Z.
LAB_RTR(config)#ip access-list extended 110
LAB_RTR(config-ext-nacl)#41 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.120
LAB_RTR(config-ext-nacl)#42 permit ip host 192.168.1.120 192.168.2.0 0.0.0.255
LAB_RTR(config-ext-nacl)#43 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.222
LAB_RTR(config-ext-nacl)#44 permit ip host 192.168.1.222 192.168.2.0 0.0.0.255
LAB_RTR(config-ext-nacl)#end
LAB_RTR#
LAB_RTR#
LAB_RTR#
LAB_RTR#sh access-list 110
Extended IP access list 110
10 permit ip 192.168.1.0 0.0.0.255 host 192.168.2.44
20 permit ip host 192.168.2.44 192.168.1.0 0.0.0.255
30 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.18
40 permit ip host 192.168.1.18 192.168.2.0 0.0.0.255
41 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.120
42 permit ip host 192.168.1.120 192.168.2.0 0.0.0.255
43 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.222
44 permit ip host 192.168.1.222 192.168.2.0 0.0.0.255
50 deny ip 192.168.1.0 0.0.0.255 192.168.2.0 0.0.0.255
60 permit udp any any
70 permit ip any any
LAB_RTR#
Step 3, Renumber the ACL, automatically:
The automatic line numbers are nice, because Cisco also gives you a way to renumber the config...
LAB_RTR#conf t
LAB_RTR(config)#! Renumber ACL 110, from line 10... increment in steps of 10
LAB_RTR(config)#ip access-list resequence 110 10 10
LAB_RTR(config)#end
LAB_RTR#
LAB_RTR#sh access-list 110
Extended IP access list 110
10 permit ip 192.168.1.0 0.0.0.255 host 192.168.2.44
20 permit ip host 192.168.2.44 192.168.1.0 0.0.0.255
30 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.18
40 permit ip host 192.168.1.18 192.168.2.0 0.0.0.255
50 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.120
60 permit ip host 192.168.1.120 192.168.2.0 0.0.0.255
70 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.222
80 permit ip host 192.168.1.222 192.168.2.0 0.0.0.255
90 deny ip 192.168.1.0 0.0.0.255 192.168.2.0 0.0.0.255
100 permit udp any any
110 permit ip any any
LAB_RTR#
Original way to change Cisco IOS ACLs
Originally, the safest way to change your ACL was by removing it from the interfaces in question (this prevented problems if you were changing the ACL via the same interface it was applied to).
interface GigabitEthernet0/1.1
no ip access-group 110 in
!
interface GigabitEthernet0/1.2
no ip access-group 110 in
Now you can add the new ACL lines, but do it before the other permit any entries...
! Reset the ACL
no access-list 110
!
access-list 110 permit ip 192.168.1.0 0.0.0.255 host 192.168.2.44
access-list 110 permit ip host 192.168.2.44 192.168.1.0 0.0.0.255
access-list 110 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.18
access-list 110 permit ip host 192.168.1.18 192.168.2.0 0.0.0.255
!! New ACL lines before the deny statement
access-list 110 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.120
access-list 110 permit ip host 192.168.1.120 192.168.2.0 0.0.0.255
access-list 110 permit ip 192.168.2.0 0.0.0.255 host 192.168.1.222
access-list 110 permit ip host 192.168.1.222 192.168.2.0 0.0.0.255
!! Deny statement
access-list 110 deny ip 192.168.1.0 0.0.0.255 192.168.2.0 0.0.0.255
access-list 110 permit udp any any
access-list 110 permit ip any any
Finally, reapply the ACL to the switch...
interface GigabitEthernet0/1.1
ip access-group 110 in
!
interface GigabitEthernet0/1.2
ip access-group 110 in
ACL Reorganization
