---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-tvdv25-guide-how-to-set-up-wireguard-clients-with-vpn-4946b665-3
title: "r-mikrotik-comments-tvdv25-guide-how-to-set-up-wireguard-clients-with-vpn-4946b665"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters", "research"]
source: docs/RAG/lot-mikrotik/forum/wireguard/r-mikrotik-comments-tvdv25-guide-how-to-set-up-wireguard-clients-with-vpn-4946b665.md
source_anchor: ""
source_lines: [71, 131]
sha256: 3dd15d1de53983aa62421a946d145939d08dd2c0456a22c29c061d7bc69fe3ce
---

# r-mikrotik-comments-tvdv25-guide-how-to-set-up-wireguard-clients-with-vpn-4946b665

/ip firewall mangle add action=mark-routing chain=prerouting connection-mark=VPN-IP-PL dst-address-type=!local in-interface=LAN new-routing-mark=wg-pl passthrough=no
Scenario D – Traffic to the countries based on their IP addresses. You can use https://mikrotikconfig.com/firewall/ to download IP ranges. Then you need to change list names to be different for each country. Note: LAN is my bridge for all LAN traffic, you can be interface-specific here
Upload file(s) to MT
/import IP-Poland.rsc
/import IP-Germany.rsc
/import IP-France.rsc
/import IP-UK.rsc
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark dst-address-list=IP-Poland new-connection-mark=VPN-IP-PL passthrough=yes
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark dst-address-list=IP-Germany new-connection-mark=VPN-IP-DE passthrough=yes
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark dst-address-list=IP-UK new-connection-mark=VPN-IP-UK passthrough=yes
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark dst-address-list=IP-France new-connection-mark=VPN-IP-FR passthrough=yes
/ip firewall mangle add action=mark-routing chain=prerouting connection-mark=VPN-IP-DE dst-address-type=!local in-interface=LAN new-routing-mark=wg-de passthrough=no
/ip firewall mangle add action=mark-routing chain=prerouting connection-mark=VPN-IP-UK dst-address-type=!local in-interface=LAN new-routing-mark=wg-uk passthrough=no
/ip firewall mangle add action=mark-routing chain=prerouting connection-mark=VPN-IP-FR dst-address-type=!local in-interface=LAN new-routing-mark=wg-fr passthrough=no
/ip firewall mangle add action=mark-routing chain=prerouting connection-mark=VPN-IP-PL dst-address-type=!local in-interface=LAN new-routing-mark=wg-pl passthrough=no
Scenario E – Combination of Scenarios C & D. The way I am doing this here is first there are computers in the network that will use tunnels for their all traffic and then the rest will use traffic based on destination address i.e. computer with IP-A will use exclusively tunnel to the UK, IP-B to Germany, IP-C to France, IP-D to Poland. So if IP is not in the local-xx list then it checks the destination address and route to proper tunnels. If IP is outside any of your lists it will be routed to your Internet connection without using VPN (i.e. in my case it is WAN). Please note that you can't do it any other way (destination and then source) as it does not make sense and would create more issues with proper routing) Note: LAN is my bridge for all LAN traffic, you can be interface-specific here
/ip firewall address-list add address=IP-A list=local-uk
/ip firewall address-list add address=IP-B list=local-de
/ip firewall address-list add address=IP-C list=local-fr
/ip firewall address-list add address=IP-D list=local-pl
Upload file(s) to MT
/import IP-Poland.rsc
/import IP-Germany.rsc
/import IP-France.rsc
/import IP-UK.rsc
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark new-connection-mark=VPN-IP-PL passthrough=yes src-address-list=local-pl
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark new-connection-mark=VPN-IP-UK passthrough=yes src-address-list=local-uk
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark new-connection-mark=VPN-IP-FR passthrough=yes src-address-list=local-fr
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark new-connection-mark=VPN-IP-DE passthrough=yes src-address-list=local-de
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark dst-address-list=IP-Poland new-connection-mark=VPN-IP-PL passthrough=yes
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark dst-address-list=IP-Germany new-connection-mark=VPN-IP-DE passthrough=yes
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark dst-address-list=IP-UK new-connection-mark=VPN-IP-UK passthrough=yes
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark dst-address-list=IP-France new-connection-mark=VPN-IP-FR passthrough=yes
/ip firewall mangle add action=mark-routing chain=prerouting connection-mark=VPN-IP-DE dst-address-type=!local in-interface=LAN new-routing-mark=wg-de passthrough=no
/ip firewall mangle add action=mark-routing chain=prerouting connection-mark=VPN-IP-UK dst-address-type=!local in-interface=LAN new-routing-mark=wg-uk passthrough=no
/ip firewall mangle add action=mark-routing chain=prerouting connection-mark=VPN-IP-FR dst-address-type=!local in-interface=LAN new-routing-mark=wg-fr passthrough=no
/ip firewall mangle add action=mark-routing chain=prerouting connection-mark=VPN-IP-PL dst-address-type=!local in-interface=LAN new-routing-mark=wg-pl passthrough=no
Hope that helps!!!
Section des commentaires
For surfshark customers, in case of scenario B, the following 2 lines are needed:
/ip firewall mangle add action=mark-connection chain=prerouting new-connection-mark=VPN passthrough=yes
/ip firewall mangle add action=change-mss chain=forward new-mss=1360 protocol=tcp tcp-flags=syn tcp-mss=1453-65535
I hope i was helpful...
My surfshark instructions and info aren't providing a PSK for this connection. Am I missing something?
No, if your connection is not using PSK then disregard anything that has PSK, do not use preshared-key option
DUDE (or "Dudette", as required)!
I've been trying to do something similar for days now to route a local subnet from another MT (that's successfully working for other peers) to my MT and couldn't get it to work! Everything was set up properly- peers were handshaking, interfaces were set up, routes were in place- but NO traffic; I was pulling my hair out!
I'm pretty sure it was these lines of yours that made the difference:
As this box is also a WG server of its own, and I'm guessing that having only one routing table was preventing my remote MT's local subnet from being routed.
I also needed this:
... which I'd had before, but was missing some parameters.
Again, thanks for this!
Thank you for the gold but this is truly unneccesarry
I treat my writeups on Reddit a bit like a backup. Usually, my setup is far from usual hence I have to do a lot of research to create a single, cohesive, stupid-proof setup (see scenario E). I like also to create solutions that may be complex in a single rule rather than creating 10 simple rules as I see the possibility of unwanted interactions between rules. Once the complex rule is fully tested (and I spend a lot of time testing) I know it will work as designed as long I need that rule to be.
Regardless - thank you and happy WGing :)
I can now- I'd been beating my head against this one final issue for an embarrassingly long time!
Greetings, super grateful with your manual, I am using your case D, which I separate the traffic of some IP depending on some countries, but it gives me error when I want that by default everything that is not marked goes through a specific VPN, is it possible?
So you want to commingle scenario D with B? So the logic (for entire network) is that if the destination IP belongs to country A then it goes to VPN-A, country B then it goes to VPN-B, etc, and then rest to VPN-Z?
Or D with A? So the logic (for specific IP addresses) is that if the destination IP belongs to country A then it goes to VPN-A, country B then it goes to VPN-B, etc, and then rest to VPN-Z?
Yes, both are possible.
You can even create rules that Country A -> VPN-A, Country B -> VPN-B,…, specific sites -> No VPN [like connecting remotely to work], everything else -> VPN-Z
