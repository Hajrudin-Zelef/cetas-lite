---
id: collect-260926-mikrotik/mikrotik/questions-12455-mikrotik-same-static-dhcp-lease-with-one-ip-address-for-lan-and-1d7e799d
title: "questions-12455-mikrotik-same-static-dhcp-lease-with-one-ip-address-for-lan-and--1d7e799d"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-12455-mikrotik-same-static-dhcp-lease-with-one-ip-address-for-lan-and--1d7e799d.md
source_anchor: ""
source_lines: [1, 14]
sha256: 35538ac2f9b6a27d576767cec388b69d42d8961bc44ceef85968eee611675858
---

# questions-12455-mikrotik-same-static-dhcp-lease-with-one-ip-address-for-lan-and--1d7e799d

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
MY QUESTION IS:
Is it possible to use static DHCP leases for same IP addres and two MAC addresses - LAN and WAN - that the host would be able to receive all the time the same IP address (192.168.1.100) regardless of LAN or WAN connection ?
I have problem with DHCP static leases in MikrtoTik CloudCore 1036. When I use static lease for WAN MAC address everything goes fine and host get IP 192.168.1.100. When I switch from WAN to LAN connection host don't receive IP 192.168.1.100 because there is already static lease for this IP address. I can`t enable second MAC to IP record because there is an information that this address already belongs to the host.
If you want to go with static IP addresses then you can specify the dhcp-server by name against each lease. As user3799089 pointed out, the latest RouterOS supports this, you may need to upgrade if the option doesn't exist. I'm running 6.23.
I just checked here on my RB951 (don't know if you're running routerOS also) and there's no specification of interfaces in IP>DHCP SERVER>LEASES.
The only thing I could see causing the problem is if you have separate DHCP address pool for each interface, even if overlapping. I tested changing ports and the host still gets the same reserved address.
For what I have seen, the only difference from a LAN port and a WAN port in a MikroTik router is the DHCP, in the former it's configured for DHCP server and the latter for DHCP client.
Create a bridge and assign WAN and LAN PORT in same bridge. Assign your DHCP to new bridge interface. Start leasing and make static leases. The lease should be the same between interfaces in new Bridge.
