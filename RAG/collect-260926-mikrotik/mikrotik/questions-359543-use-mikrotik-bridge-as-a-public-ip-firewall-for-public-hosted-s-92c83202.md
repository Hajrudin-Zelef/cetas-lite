---
id: collect-260926-mikrotik/mikrotik/questions-359543-use-mikrotik-bridge-as-a-public-ip-firewall-for-public-hosted-s-92c83202
title: "questions-359543-use-mikrotik-bridge-as-a-public-ip-firewall-for-public-hosted-s-92c83202"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "research"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-359543-use-mikrotik-bridge-as-a-public-ip-firewall-for-public-hosted-s-92c83202.md
source_anchor: ""
source_lines: [1, 27]
sha256: 90ab3118da3b271de765c479c26342a586b24e2175c0eac596d44723a70966e5
---

# questions-359543-use-mikrotik-bridge-as-a-public-ip-firewall-for-public-hosted-s-92c83202

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
We want to do this:
Internet<->MikroTik in Bridge Mode with Firewall Filter<->Hosted Server
The primary objective is to allow RDP and FTP in from the outside but block everything else from the outside. From the inside everything must go out.
The problem we are running into is we add these rules and blocking outside to inside is working, but now the hosted server cannot access anything to the outside. The return TCP/IP from the outside is not port 3389 or port 80, but random.
/interface bridge filter> pr
Flags: X - disabled, I - invalid, D - dynamic
0 ;;; Accept ICMP for PING
chain=forward action=accept mac-protocol=ip dst-address=196.x.x.x/32 ip-protocol=icmp
1 ;;; Accept FTP Transfer Port
chain=forward action=accept mac-protocol=ip dst-address=196.x.x.x/32 dst-port=20 ip-protocol=tcp
2 ;;; Accept FTP Control Port
chain=forward action=accept mac-protocol=ip dst-address=196.x.x.x/32 dst-port=21 ip-protocol=tcp
4 ;;; Log everything that is about to get dropped
chain=forward action=log mac-protocol=ip dst-address=196.x.x.x/32 ip-protocol=tcp log-prefix="firewall_drop"
5 ;;; Drop everything
chain=forward action=drop mac-protocol=ip dst-address=196.x.x.x/32 ip-protocol=tcp
Just FYI bridge is set to use firewall and connection tracking is on.
I suggest to use local ip subnet for your server and have private LAN between mikrotik (router) and server. Also have mikrotik directly on public ip. Then do the ip src-nat for outgoing traffic (from server to internet). All incoming traffic should be port forwarded (dst-nat). Also you can use normal layer3 ip firewall and disable layer3 firewall for layer2 / bridge network.
This has something to do with connection tracking, but I don't know exactly how it works in a bridge. I would just add another rule to allow everything from inside with src-address or interface:
I agree with @Matt that bridging was a stupid idea and that we should have routed from the start. The entire reason why we bridged was because our border router was a legacy Cisco VXR 7206 with two interfaces, a WAN (ATM) and LAN (used for public) interface. We wanted to firewall stuff on public without re-sub-netting our network. Also, although @DJ_Kukky's suggestion of using DST-NAT is plausible that would have led to much more configuration and also we can't (or won't) give our clients private IPs as we're a public facing ISP.
We should have either implemented the firewall rules on the Cisco or replaced the technology (ATM). In the end we replaced the ATM with MetroEthernet and we were able to establish the outside interface on an Ethernet interface and as such use MikroTik. It's working really well.
