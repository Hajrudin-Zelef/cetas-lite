---
id: collect-260926-mikrotik/mikrotik/questions-41946-port-forwarding-not-working-on-mikrotik-hap-ac-08f369d2
title: "ADDRESS NETWORK INTERFACE"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-41946-port-forwarding-not-working-on-mikrotik-hap-ac-08f369d2.md
source_anchor: ""
source_lines: [1, 26]
sha256: fdad8c63f701ab34a5baca4033211ab5e54cd05bd0e3b5f0c77c678b6cb1cce0
---

# ADDRESS NETWORK INTERFACE

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have set my MikroTik RouterBoard hap AC as the router instead of my Modem/Router that my ISP gave me.
I'm trying to set port forwarding using a tutorial I've found in google.
That's the configuration of my Port Forwarding:
Flags: X - disabled, I - invalid, D - dynamic
0 ;;; defconf: masquerade
chain=srcnat action=masquerade out-interface=ether1 log=no log-prefix=""
1 chain=dstnat action=dst-nat to-addresses=192.168.1.7 to-ports=8081 protocol=tcp in-interface=ether1 dst-port=8081
And That's my ip address configuration:
Flags: X - disabled, I - invalid, D - dynamic
# ADDRESS NETWORK INTERFACE
0 ;;; defconf
192.168.1.1/24 192.168.1.0 ether2-master
1 D 100.100.161.43/23 100.100.160.0 ether1
I have to say that 100.100.161.43 is the IP the modem gave to my router, but it isn't my ISP IP.
The Port Forwarding isn't working at all. I've tried tcpdump to check if 192.168.1.7 is getting any packets while I'm trying to reach it, but no traffic is being captured.
100.64.0.0/10 is a special IP-range defined in RFC6598, commonly known as 'shared address space'. The typical use of this is that ISP's which are running out of IPv4 space assign customers an IP address from this range and use carrier grade NAT to provide access to the public internet.
The most important reason they use this specific range is that it does not overlap with private IP ranges assigned in RFC1918, so application layer protocol inspection will be able to distinguish the ISP's NAT from the client's NAT.
So the problem here is that your NAT is placed behind the ISP's NAT. That's why you see another IP when you check your address on a website. You will have to contact your ISP if you need to run publicly accessable services, because at this moment you do not have your own public IP address.
100.100.161.43 is public address. Let's call it Address M. If you are not using Address M as a destination when trying to access from outside, that means public address you are trying to reach (Address R) does not forward your request to Mikrotik (Address M). Either try accessing from outside to Address M, or you need to ask your ISP to forward port on Address R to Address M.
Other than that your Mikrotik configuration looks good.
