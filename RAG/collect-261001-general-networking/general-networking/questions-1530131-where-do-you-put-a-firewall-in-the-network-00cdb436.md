---
id: collect-261001-general-networking/general-networking/questions-1530131-where-do-you-put-a-firewall-in-the-network-00cdb436
title: "Where do you put a firewall in the network?"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/questions-1530131-where-do-you-put-a-firewall-in-the-network-00cdb436.md
source_anchor: ""
source_lines: [1, 15]
sha256: 74e850b5fb64ed21a67baa836fe5d00dc6a6d00ac562bc79e4217dea22a9aa46
---

# Where do you put a firewall in the network?

*Score : 1 | Source : https://superuser.com/questions/1530131/where-do-you-put-a-firewall-in-the-network*

I have installed opnsense on a pc to use as my business firewall where do I put this in my network
Do I put it just off of a port on my switch or do I need a 2 port ethernet card and plug the internet into one and my router into the other?

---

### Reponse (acceptee) — score 3

A firewall is only useful if connections actually go through it. It'll do nothing at all if it just dangles off a single switch – it cannot somehow "attract" traffic1, all your client devices will just continue to directly talk to the router/gateway.
To work properly, the firewall needs to have separate connections for "outside" and "inside", i.e. it needs to sit between the Internet and the LAN. This usually means two Ethernet ports – one from your ISP connection, one to your LAN switch.
But it is possible to use two VLANs over a single Ethernet port (also known as "on a stick" configuration), although your switch must also support VLAN tagging and it won't necessarily provide as good performance as separate physical ports would.
1 (Except by using ARP spoofing, and yes, I've heard of one or two "plug and play" firewall products which do exactly that, but... I would not want to use a security product that is literally performing an attack against the network.)
