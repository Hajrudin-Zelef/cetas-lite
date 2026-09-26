---
id: collect-260926-mikrotik/mikrotik/wireguard-vpn-between-two-mikrotiks-running-routeros-7
title: "Wireguard in Mikrotik RouterOS 7"
domain: mikrotik
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/wireguard-vpn-between-two-mikrotiks-running-routeros-7.md
source_anchor: ""
source_lines: [1, 26]
sha256: 8d712b6b366b3b17bb324dd07f8caf14b60e2e0ff6ac96951b378716ca491fdd
---

# Wireguard in Mikrotik RouterOS 7

As you might be aware **wireguard** offers speed and security (over 35% faster over Openvpn). Apart from that Wireguard offers clients for most common platforms, ie Windows, Linux, MacOS, iOS, Android. In that in mind decided to upgrade the **hexS** to the RouterOs 7 (beta2) and test it with another **hexS** (which literally takes 3 mins to setup)

## **Wireguard VPN between two Mikrotiks running RouterOS 7:**

In the wireguard tab add new interface (BLUE +), then as per the below screenshot:

name the interface, click apply. Select port, allow it in the firewall.

Add IP address to the wanguard interface (point 2)

Copy public key to the peer section, add other endpoint and other site's IP address of the wanguard interface.

Now you need to do the opposite site, I will leave that bit to you 

That's all, the tunnel is made ! 

If you have networks (LANs) behind your Mikrotik routers you most likely would like to join them too
In that case you need the appropriate subnets to the **Allowed Addresses** in Wireguard Peer section
and add static route on the router: my case subnet on the other side of the tunnel was 192.168.22.0/24
so I added it to Allowed Addresses and added static route too:

**/ip route add disabled=no distance=1 dst-address=192.168.22.0/24  gateway=10.6.6.1**

#mikrotik #MikrotikRouterboard #mikrotikacademy #mikrotiksydney #mikrotiklte #mikrotikvpn #mikrotikaustralia #mikrotiksxtlte #mikrotiktc
