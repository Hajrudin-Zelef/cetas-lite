---
id: collect-260926-mikrotik/mikrotik/questions-20073-is-it-possible-to-connect-two-networks-using-mikrotik-routeros-w-0edb3e85
title: "questions-20073-is-it-possible-to-connect-two-networks-using-mikrotik-routeros-w-0edb3e85"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/questions-20073-is-it-possible-to-connect-two-networks-using-mikrotik-routeros-w-0edb3e85.md
source_anchor: ""
source_lines: [1, 11]
sha256: 09242cfc03764656e4af98e8232ad83757599f125686acb508358c5e45093ab6
---

# questions-20073-is-it-possible-to-connect-two-networks-using-mikrotik-routeros-w-0edb3e85

Since you don't have any static IPs on both sites then you will need to use Dynamic DNS to be able to have both mikrotik routers 'find' each other.
Mikrotik supports its own Dynamic DNS out of the box (it's not easy to remember, but it does the job and works out of the box). Check IP > Cloud to enable it.
Now, on Office 1 since it can be DMZ will be the VPN server. Mikrotik supports PPTP, L2TP, OpenVPN, SSTP. Use which-ever works better for you.
On Office 2 which will be the 'client' side of the VPN you create a vpn client interface connecting to Office's 1 dynamic DNS.
Make sure to use the latest version of Mikrotik so the above features are fully supported.
The next step, now that you have a vpn connection, is to create an EoIP tunnel over the VPN you created. This will allow you to have both sites talk to each other as if they where on the same physical location.
Finally you bridge the EoIP tunnel interface on both sites with your LAN interface(s) and you should be able to talk from one site to another with full ethernet frames.
Now, beware that this approach has a lot of overhead due to double encapsulation of packets.
Another approach would be to enable DMZ on Office 2 as well, and have both Mikrotik routers use only an EoIP tunnel that will connect directly without any VPN under it.
In this case you will need some scripting that will detect your public IP on both sites and update the EoIP interface when the IPs change.
Personally I use the second solution to reduce the overhead as much as possible. For security I use IPsec (since EoIP does not encrypt anything on its own). Be careful about IPsec since it uses a lot of CPU (I wouldn't recommend it on a MIPSBE routerboard).
