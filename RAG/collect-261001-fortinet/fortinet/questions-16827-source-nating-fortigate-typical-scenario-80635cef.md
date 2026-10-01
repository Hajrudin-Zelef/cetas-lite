---
id: collect-261001-fortinet/fortinet/questions-16827-source-nating-fortigate-typical-scenario-80635cef
title: "questions-16827-source-nating-fortigate-typical-scenario-80635cef"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-fortinet/questions-16827-source-nating-fortigate-typical-scenario-80635cef.md
source_anchor: ""
source_lines: [1, 34]
sha256: c6adb53ebc16479c4c92f924be01f07fb1ed17487dbded4497fcbf261a1a1064
---

# questions-16827-source-nating-fortigate-typical-scenario-80635cef

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
7
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a small query with respect to NATing in Fortigate.I'm struck in particular by a scenario where the remote network allows users with a specific IP range with a specific port for RDP over a different set of physical links.
User(192.168.60.0/24) in LAN should connect to 10.48.1.3 on remote network which is connected to Internal ports on Fortigate (Nor WAN since it is used for internet).
Link between Remote network firewall & Fortigate has been established (10.189.254.17-10.189.254.18). I can ping remote firewall interace 10.189.254.17.
Admin wants us to access 10.48.1.3 over 3389 port via 10.189.1.8-10.189.1.15 (allowed IPs on remote firewall).
So basically, user (ex:Source: 192.168.60.15 need to access Destination 10.48.1.4 via allowed IPs (10.189.1.8-10.189.1.15) over physical link between Fortigate & remote firewall (10.189.254.18-18.189.254.17).
I have tried VIP (Static NAT) (Source NAT) port forwarding, IP Pool (Destination NAT), but no help.
Please advice how to proceed. It is a typical scenario , can we achieve it?
Under your Policy Rule to allow access from 192.168.60.15 to 10.48.1.4.
Simply select NAT then create a Pool and put IP range of 10.189.1.8-10.189.1.15.
Key thing for this to work is you have route on your fortigate pointing 10.48.1.4 to 10.189.254.17 and finally 10.189.254.17 knows 10.189.1.8-10.189.1.15 lives back via 10.189.254.18.
Thats a pretty standard rule and sounds achievable. Some packet captures may help you to debug the issue.
Yes I understand your scenario and your requirement ..to access resources on remote firewall on port RDP ie 3389 from fortigate 200d connected switch lan users
For your requirement no natting required.
.
Please configure static route in fortigate 200D as below
Ip route 10.48.1.0 255.255.255.0 points towards gateway 10.189.254.17
And for reverse traffic static route in remote n
/W firewall
Ip route 192.168.60.0 255.255.255.0 pointing towards gateway 10.189.254.18
And have a security policies in firewalls allowing traffic
Policy in fortigate 200D
Source interface : interface Port need to mention Destination interface : interface Port need to mention Source address :192.168.60.15/32 Destination address :10.48.1.4/32 Port :tcp-3389 Action : allow Security profiles : on
Now security policy in remote n/w firewall
Source interface : egress interfàe of firewall Destination interface :ingress interface of firewall Source address : 192.168.60.15/32 Destination address :10.48.1.4/32 Port :3389/TCP Action : allowed Security profiles :on
.
Now user of fortigate 200D lan users can access internal hosted server on remote network firewall on port 3389
For futher security if you wants to hide your ips then you can use source natting in fortigate 200D firewalls but to accomplish this you need to configure static route in fortigate 200d with destination as source nat pool pointing.
Towards gateway 192.189.254.17..likewise..
