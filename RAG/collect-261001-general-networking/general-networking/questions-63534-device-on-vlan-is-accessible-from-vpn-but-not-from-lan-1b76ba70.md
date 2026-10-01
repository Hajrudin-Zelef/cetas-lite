---
id: collect-261001-general-networking/general-networking/questions-63534-device-on-vlan-is-accessible-from-vpn-but-not-from-lan-1b76ba70
title: "questions-63534-device-on-vlan-is-accessible-from-vpn-but-not-from-lan-1b76ba70"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2019-11-14"]
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-63534-device-on-vlan-is-accessible-from-vpn-but-not-from-lan-1b76ba70.md
source_anchor: ""
source_lines: [1, 8]
sha256: 4080f78103cb24ea9d606f1936b494ba50e044c41c85c46edbc8e136cca94153
---

# questions-63534-device-on-vlan-is-accessible-from-vpn-but-not-from-lan-1b76ba70

I have a Meraki MX84 with two VLANs, 192.168.213.0/24 (main subnet), and 10.4.213.0/24. I have an EdgeRouter 12P connected to the MX with static WAN IP 10.4.213.2. The issue is: when I VPN into the MX from home, I can access the ER GUI at 10.4.213.2, as well as several IPs behind the ER that are NAT'd from the ER's WAN interface. When I am on site and connected to the main subnet, however, I am not able to access the ER's GUI or NAT'd IPs behind it. Why am I able to access the ER from the VPN workstation and not the LAN? Hopefully the attached diagram will help explain the situation. I think I might need to configure two virtual adapters on the LAN workstation (with only one of those having a gateway to retain reliable internet access), but haven't tried this.
- 
        2Why are you NATing inside your own private address space? Likely, that's the problem. If not you need to add the basic configurations including routing tables, NAT rules and firewall rules for the routers involved to your question. We can't use a crystal ball.Zac67– Zac67 ♦2019-11-14 11:06:52 +00:00Commented Nov 14, 2019 at 11:06
1 Answer 1
"LAN workstation" is in the 192.168.213.0/24 subnet - the same subnet you're using for "BAY x LAN". That is a very bad design.
Unless you are using destination NAT and source NAT between those locations, no communication is possible. Likely, you're using destination NAT from LAN workstation to reach "Device N", but without source NAT "Device N" tries to send the reply locally, not through the gateway. No go.
The only sane solution is to renumber the network to get rid of NAT entirely (except for WAN/public IP space). Remember that NAT is a kludge is needs to be avoided wherever possible.
At the very least, you need to change the workstation LAN subnet address to something else, so it doesn't seem to be local to the Device N LANs.
