---
id: collect-261001-general-networking/general-networking/questions-13253-where-to-natwhere-to-run-bgp-e47f4dba
title: "questions-13253-where-to-natwhere-to-run-bgp-e47f4dba"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-13253-where-to-natwhere-to-run-bgp-e47f4dba.md
source_anchor: ""
source_lines: [1, 12]
sha256: e80dce9ffe0afd0490984175f79faf9466c57ea32c0310b4a84b5ac3ef62c2a1
---

# questions-13253-where-to-natwhere-to-run-bgp-e47f4dba

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
3
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
in this diagram :
we are peering with 2 ISPs, let's say I have a class c IP space like 200.200.200.0/24 and let's say my ASN is 65000, the thing is where should I run NAT? on which device? I mean on Cisco Routers that face the internet or could I run that on the FortiGate? and where on which device should I config the peering? again on Cisco routers facing those ISPs or is it possible to have BGP peering configured on the FortiGate? I know some stuff about BGP and how it works, an I know it's based on the TCP session and it's not working like IGPs in which peers are directly connected, but the thing is how could I have NAT alongside the BGP in this diagram?what's the best to do?
You should run BGP on your Cisco routers connecting to ISP. If FW supports BGP, it should run iBGP with the routers.
I would prefer doing NAT on the FW. All you have to do is make sure that the traffic leaving a FW comes back to the same. I think you would already be doing that.
You don't have to accept all the routes from the Internet. Simply apply an in-bound filter to accept only the default route and a few other prefixes.
You can either use VRRP/HSRP on router's inside interfaces or use E/iBGP on the firewalls. This will allow routers to send the default route to the FW dynamically. If one router goes down, FW will not be black-holing the traffic to it as that route will simply go away.
