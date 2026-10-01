---
id: collect-261001-cisco/cisco/r-networking-comments-twe21s-cisco-asa-nat-rule-3d98fd22
title: "r-networking-comments-twe21s-cisco-asa-nat-rule-3d98fd22"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/r-networking-comments-twe21s-cisco-asa-nat-rule-3d98fd22.md
source_anchor: ""
source_lines: [1, 42]
sha256: 48de50916e3005d2c4774e1d27b0b523883763abc3ef25987086c5c41511ad3e
---

# r-networking-comments-twe21s-cisco-asa-nat-rule-3d98fd22

Cisco ASA NAT rule 
        
        
        
    
    
    Hello all,
I was configuring a NAT rule today and a colleague of mine asked two questions about the NAT rules in Cisco firewall ASAs.
So for example: I configured a NAT rule for a host IP in a DMZ network and that host IP needed to get out on the internet by NATing it to a specific NAT IP address (not the internet interface IP address).
So what I configured (an example): nat (DMZ,INTERNET) source static 10.10.10.10 193.150.150.100
Question 1: If I need to reach that host IP (10.10.10.10) from the internet, how will that work in the ASA? How do the the firewall know that I want to specifically reach that host IP?
Question 2: Let say I want to permit an amount of servers in the DMZ that need to reach out on the internet, let say 10.10.10.10, 10.10.10.11, 10.10.10.12 and 10.10.10.13 need to be NAT:ed to 193.150.150.100. If I want to reach, let say, 10.10.10.12 from the internet, how would the firewall know that I specifically want to reach the 10.10.10.12, i.e the firewall need to forward the traffic to 10.10.10.12?
So what I replied to him was that (I was thinking about the command "show ip nat translations" in a Cisco router/switch and used that logic) I think that the firewall does a NAT mapping and maps these different host IPs (10.10.10.10, 10.10.10.11 etc) to the public IP 193.150.150.100 and assigns each host IPs different ports and that way it knows how to forward the traffic. But since I was unsure, I told my colleague that I need to look this up because I could totally be wrong and quite frankly, Im not so sure how this exactly works.
If someone could help me to understand and explain this to me, I would appreciate it a lot!
Thanks.
Section des commentaires
Question 1: Because there's a 1:1 NAT relationship between 10.10.10.10 and 193.150.150.100 - if you browse to the public IP, the ASA will translate that address to the inside address.
Question 2: You would have to configure individual ports. You cannot have multiple 1:1 NAT rules configured for the same external IP, just like you cannot have multiple devices with the same IP address.
Use the packet tracer on the ASA to confirm your Nats routes and acls all work. It’ll also show you Cisco’s order of operation if you aren’t familiar with it
CCNA doesn't talk about NAT?
Edit after rereading:
Doing multiple whole IP mappings, like:
nat (DMZ,INTERNET) source static 10.10.10.10 193.150.150.100
nat (DMZ,INTERNET) source static 10.10.10.11 193.150.150.100
nat (DMZ,INTERNET) source static 10.10.10.12 193.150.150.100
I'm not sure if this is fully supported, as you could max it out and have issues (3IPs mapping to one)
You could do port forwarding, and that way you would know how each port exactly maps.
Like:
object service port-8433
service tcp source eq 8433
exit
object network obj-2.2.2.2
host 2.2.2.2
exit
object network obj-38.93.235.177
host 38.93.235.177
nat (inside,outside) 1 source static obj-38.93.235.177 obj-200.200.200.50 service port-8433 port-8433
Commentaire supprimé par le membre
Thanks for your interest in posting to this subreddit. To combat spam, new accounts can't post or comment within 24 hours of account creation.
Please DO NOT message the mods requesting your post be approved.
You are welcome to resubmit your thread or comment in ~24 hrs or so.
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
