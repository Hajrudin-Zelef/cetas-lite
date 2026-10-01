---
id: collect-261001-meraki/meraki/questions-777920-dhcp-on-meraki-mx64-or-windows-server-2012r2-fb40766c
title: "questions-777920-dhcp-on-meraki-mx64-or-windows-server-2012r2-fb40766c"
domain: meraki
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2016-05-20"]
keywords: []
source: docs/RAG/collect-261001-meraki/questions-777920-dhcp-on-meraki-mx64-or-windows-server-2012r2-fb40766c.md
source_anchor: ""
source_lines: [1, 10]
sha256: 08eb8c7e85f0c3bbff61e6070b1d7066ba122b67340d7bb3e0275b8d7c9a3d06
---

# questions-777920-dhcp-on-meraki-mx64-or-windows-server-2012r2-fb40766c

I'm new to the Cisco world and I would like some advises. Currently I'm configuring a Windows Server for a small bussiness (less then 50 users). Normally the Windows server would have the AD, DNS and DHCP services installed alltogether and configured, thus disabling the DHCP server on the ISP router. But since we got a new gear (Cisco Meraki MX64), they recommend me using the DHCP server on the Meraki itself and let it control all the ip assignment and routing. One main reason is because of the cloud based managment of the Meraki. My question is: Does it really make a difference who controls the DHCP? If by any reason I use the DHCP on the Cisco Meraki, how do I connect the DNS and AD services with the Meraki? How does it work? I couldn't find any tutorial regarding this. We also have a Cisco Meraki switch controller. Thanks guys for your help!
- 
        One could argue this question is a matter of opinion. There is one relevant fact, which is that you seem to understand Windows DHCP fairly well and Meraki DHCP is not as clear to you. That is at least an entry in the column in favor of using the Windows DHCP. Also, cloud management isn't inherently better, as far as I know. What is the supposed advantage of cloud management over server admin tools or RDP management?Todd Wilcox– Todd Wilcox2016-05-20 04:03:01 +00:00Commented May 20, 2016 at 4:03
2 Answers 2
Use Microsoft's DHCP if you're running Active Directory and DNS.
The Meraki's cloud management is irrelevant. Your environment is going to hinge more on the smooth operation of the Microsoft services than anything else.
A normal use case for running DHCP on the Meraki would be a home or branch office situation where there's no server presence. Not this.
The Meraki MX-series has a very basic interface and set of capabilities. So for anything more complex than the barebones DHCP functionality, you should use a server.
If you have other Meraki gear in your network (or other cloud-managed network equipment), it's a good idea to use the firewall to provide DHCP on your management vLAN. You don't want trouble with your server to cause the entire network to go down.
Let the Meraki firewall handle DHCP for vLAN 1, and let your Windows server handle DHCP for your data and other vLANs.
