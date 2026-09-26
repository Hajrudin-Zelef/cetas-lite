---
id: collect-260926-mikrotik/mikrotik/questions-999196-mikrotik-and-vpn-for-specific-web-sites-only-c6befb06
title: "questions-999196-mikrotik-and-vpn-for-specific-web-sites-only-c6befb06"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vpn/questions-999196-mikrotik-and-vpn-for-specific-web-sites-only-c6befb06.md
source_anchor: ""
source_lines: [1, 23]
sha256: 0ff3a587416acd7b9cc7ff7c46869b12f9e8da5bc062c7a0123c9bbfb71a6173
---

# questions-999196-mikrotik-and-vpn-for-specific-web-sites-only-c6befb06

Let's assume that you will be using a PPTP VPN just to demonstrate the commands you need to run.
Since PPTP's encryption is broken for a long time now, I suggest you use something more secure (like OpenVPN). The principle is the same regardless of which VPN/tunnel technology you use.
So first you create the VPN without adding a default gateway route.
/interface pptp-client
add add-default-route=no allow=pap,chap,mschap1,mschap2 connect-to=VPN_SERVER_IP \
dial-on-demand=no disabled=no max-mru=1440 max-mtu=1440 mrru=dis \
name=VPN_NAME password="MY_STRONG_PASSWORD" profile=default-encryption user=USERNAME
Then you create a new routing table by adding a default gateway via the VPN with a new routing mark vpn. This will allow you to route packets via the VPN.
/ip route add dst-address=0.0.0.0/0 distance=1 gateway=VPN_GATEWAY_IP routing-mark=vpn
The next route is optional in case you want to block outgoing traffic if the VPN is down:
/ip route add dst-address=0.0.0.0/0 type=unreachable distance=2 routing-mark=vpn
We also need to do some NAT for the packets that will be leaving via the VPN interface.
/ip firewall nat add chain=srcnat out-interface=VPN_NAME action=masquerade
Now we add the mangle rule that will match the destination IPs we want and do a mark-routing on them so that they will use the vpn route table we created.
/ip firewall mangle add chain=prerouting dst-address-list=VPN action=mark-routing new-routing-mark=vpn
Finally we create an Address List on the firewall with the IPs that we want to route via the VPN.
/ip firewall address-list add list=VPN address=1.1.1.1
/ip firewall address-list add list=VPN address=2.2.2.2
/ip firewall address-list add list=VPN address=3.3.3.3
/ip firewall address-list add list=VPN address=4.4.4.4
You repeat the last rule as many times as you need for as many IPs as you want to route via the VPN.
Keep in mind that the rules above do not provide any security as to who behind your router will be able to access the VPN etc. You may need to add appropriate source IPs checks on the rules to make them more secure.
Also this method will route whole IPs via the VPN. If you need to route specific ports/protocols via the VPN you simply create additional mangle rules that match whatever you need and do mark-routing on them.
