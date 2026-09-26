---
id: collect-260926-mikrotik/mikrotik/questions-1869581-can-i-configure-my-routeros-ap-to-provide-internet-access-from-0c7268a2
title: "questions-1869581-can-i-configure-my-routeros-ap-to-provide-internet-access-from-0c7268a2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "research"]
source: docs/RAG/lot-mikrotik/RouterOS/questions-1869581-can-i-configure-my-routeros-ap-to-provide-internet-access-from-0c7268a2.md
source_anchor: ""
source_lines: [1, 25]
sha256: 6adfd704db9f59b2baab2a9887d60e4ebdd8d18a8edfa82e39536015000fb3dc
---

# questions-1869581-can-i-configure-my-routeros-ap-to-provide-internet-access-from-0c7268a2

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
My setup is as follow:
Modem from ISP is connected to a "primary wifi AP" (AP-A) through its WAN port
AP-A serves wifi to assorted client devices in my house
I would like "secondary wifi AP" (AP-B) (which is a MikroTik hAP ac²) to join AP-A's network as a wifi client and provide internet access to devices plugged into its LAN ports
I do know that it is possible to have AP-B join AP-A's network as a wifi client. I achieved this in the RouterOS console by providing the password for AP-A's network in the Wireless -> Security Profiles menu and then scanning for the network on the main Wireless menu and joining it. I verified it had joined the network by seeing its MAC address on AP-A's client list.
However this is where I got stuck. I do not know what to do next. How can I "connect" the internet access that should be available from AP-B being a client on AP-A's wifi network to devices plugged into LAN ports of AP-B?
Some ideas which have come to mind:
AP-B could create its own network for clients plugged into its LAN ports. I can do this but I don't know how to "share" the internet connection with that network.
AP-B could somehow let clients plugged into its LAN ports join the same network as AP-A. I don't know how to do this since AP-A is the one running the DHCP server for that network.
While I wouldn't turn down a "spoon-fed" solution I would prefer to understand the different options and "learn how to fish for myself". Currently I am at a loss trying to grasp (a) what solutions are viable (b) what "networking entities" those solutions need (c) how to set those up properly.
AP-B could somehow let clients plugged into its LAN ports join the same network as AP-A. I don't know how to do this since AP-A is the one running the DHCP server for that network.
That's usually called a "bridge". Ethernet switches are bridges, APs are bridges – they are "transparent" at Ethernet packet level, i.e. host A on one side can send anything to the MAC address of host B on the other side and it'll go straight through. Importantly, this includes broadcast packets as well, so DHCP requests and replies will also go through a bridge.
In RouterOS, create a bridge interface and add all the Wi-Fi interfaces and Ethernet interfaces as bridge ports.
All home wifi routers come with this kind of configuration, with their Wi-Fi and all their Ethernet ports (except for the one labelled "WAN" of course) as bridge members, creating one large network. Commercial APs are the same except they don't have a "WAN" port at all and everything is bridged together. With RouterOS you can decide whether you want to bridge all ports you have or only some of them.
Note that because all bridge ports form a single network at Ethernet level, they lose their "identity" – only the bridge itself has a MAC address or IP address. So if you're connected to it via its ether1 IP address (for example), and ether1 gets bridged, you may lose the SSH/web/Winbox connection. (RouterOS will automatically move the IP address to the bridge so you should be able to reconnect, but you'll have to adjust the /ip/address section manually anyway.)
AP-B could create its own network for clients plugged into its LAN ports. I can do this but I don't know how to "share" the internet connection with that network.
This is the default mode, i.e. what you get from a router when there isn't a bridge connecting its interfaces. "Sharing" a connection in this way usually just means deciding on a unique subnet, setting up DHCP to issue IP addresses, and making sure the rest of the network has a route towards this router's subnet. (The last part is usually worked around by using NAT i.e. address masquerading instead.)
I wouldn't recommend this for APs, as it prevents seamless roaming by clients – they will lose connection if they roam to another AP, seeing it has the same SSID, but suddenly end up with a different IP address.
But if you really want this for a unique SSID, just set up the device as a home router – with DHCP server on the "inside" (WiFi) interface and everything. The RouterOS template should also enable NAT, which really is not the preferred thing to do anywhere else but the outer boundary, but it'll do for a home network.
