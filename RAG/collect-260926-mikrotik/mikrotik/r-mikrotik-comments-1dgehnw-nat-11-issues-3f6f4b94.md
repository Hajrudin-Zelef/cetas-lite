---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-1dgehnw-nat-11-issues-3f6f4b94
title: "r-mikrotik-comments-1dgehnw-nat-11-issues-3f6f4b94"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/r-mikrotik-comments-1dgehnw-nat-11-issues-3f6f4b94.md
source_anchor: ""
source_lines: [1, 40]
sha256: 897ef52fe4c95b30db20444d817fd18d61023ddcee623171b06d33cf5119a129
---

# r-mikrotik-comments-1dgehnw-nat-11-issues-3f6f4b94

NAT 1:1 issues 
        
    Hello,
I'm trying to assign an IP address to a device on the network. I've managed to create the fairly simple firewall rules as this:
      From inside, the ip 172.16.20.241 will use the external IP xx.xx.xx.108
From outside, accessing  xx.xx.xx.108 will point to 172.16.20.241
    
accessing from outside to insode is working as it should, but from the insode when I try this command:
curl ifconfig.co
I get a reply of a differant IP: xx.xx.xx.106. This is also an IP I have on the router, but it shouldn't be in use in this case.
These are the commands I used to add the rules:
/ip firewall nat add chain=dstnat dst-address=xx.xx.xx.108/29 action=netmap to-addresses=172.16.20.241/32
 /ip firewall nat add chain=srcnat src-address=172.16.20.241/32 action=netmap to-addresses=xx.xx.xx.108/29
I appretiate any help on the subject.
Thank you very much!
Section des commentaires
Why are you trying to do 1:1 of a single host address to a whole range?
Just make sure the address exists on the outside interface and then make the rule 1:1 as in a /32 to a /32
Also make sure you have matching Firewall rules to allow the traffic on the inside dst address. If you want true 1:1 you just target IP to IP with any protocol and port. Otherwise it is not 1:1
Thank you for the answer
I've seen the /29 mistake right after I posted here and fixed it. I've tried to edit the post but it didn't work.
Anyway. I am actually pointing xx.xx.xx.108/32 to 172.16.20.241/32 and pointing 172.16.20.241/32 to xx.xx.xx.108/32. It still shows xx.xx.xx.106/32 for some reason.
Your src-nat rule specifies an out-interface-list - why bother?
This means you'd need a separate src-nat rule for LAN>LAN connections
Hairpin NAT config doesn't look correct. https://help.mikrotik.com/docs/display/ROS/NAT#NAT-HairpinNAT
There are a few ways to achieve this but in general you need to:
src-nat LAN traffic bound for [server LAN IP]
dst-nat all traffic that should hit the server to [server LAN IP] - which you're already doing
Not sure on the difference between action=src-nat and netmap tbh
Sadly, using src-nat or netmap doesn't help. I'm not sure wha is wrong with it.
there seems to exist a generic "catch all" SNAT/masquerade rule with a higher priority that's SNATing your traffic to xx.xx.xx.106.
are you sure your specific SNAT rule remapping 172.16.20.241 to .108 has a higher priority (is upper in the list) than the generic one?
You're right. Moving the rule to the top of the list actually worked!! Thank you very much!
netmap needs the same number of host one left and right, I think.
if u have only 1 host on right, use src-nat or dst-nat action instead.
If u use SINGLE address on left and right, use netmask /32 for left too, like
Thank you!
The issue was in rules placment. there was another rule with higher order which took over it.
Placing the rule higher solved the isue.
yes, "personal" rules should be placed higher than common rules in NAT
