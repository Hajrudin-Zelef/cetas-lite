---
id: collect-260926-mikrotik/mikrotik/questions-1859889-keep-a-bridge-isolated-from-the-os-in-mikrotik-routeros-27902519
title: "questions-1859889-keep-a-bridge-isolated-from-the-os-in-mikrotik-routeros-27902519"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters", "research"]
source: docs/RAG/lot-mikrotik/RouterOS/questions-1859889-keep-a-bridge-isolated-from-the-os-in-mikrotik-routeros-27902519.md
source_anchor: ""
source_lines: [1, 15]
sha256: ad08156a2142825a78e5f8f70a73f1742a505a15886239074154be9bfbb1844c
---

# questions-1859889-keep-a-bridge-isolated-from-the-os-in-mikrotik-routeros-27902519

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am setting up a MikroTik cAP, to be configured as a multitenant access point. The LAN interface is a VLAN trunk interface. For each tenant, as well as for the management LAN, I am configuring a dedicated bridge. Ports on the bridge are a VLAN on ether1 and a virtual wireless interface for each frequency band. Only the bridge for the management LAN has an IP address configured, the others do not so the cAP is not reachable via IP from the tenant networks. (MAC Server is also turned off.)
After configuring the management network and one tenant, I see broadcast traffic from that tenant in the list of IP connections (/ip firewall connection print, or the equivalent in WebFig). That looks to me as if the system is passing them to its own IP protocol stack – which I want to avoid, in order to minimize attack surface.
How do I configure a bridge to just pass traffic through, without ever passing it to the local IP stack?
My initial suggestion would be to remove those dedicated bridges and VLAN interfaces entirely and replace them with a single vlan-filtering=yes bridge. The management VLAN is set through the bridge's pvid= parameter, while wireless interfaces have the vlan-mode=use-tagvlan-id=XXX parameters, and the physical Ethernet port should have all those VLANs defined via /interface/bridge/vlan.
I believe this is the recommended configuration in general, but in your case it also results in the router's networking stack seeing only one VLAN at a time – only the frames that match the bridge's pvid will arrive through the "self port".
Alternatively: /ip/firewall drop all incoming packets from in-interface=!my-mgmt-bridge. Repeat for /ipv6/firewall.
(I don't know why you're seeing those conntrack entries in IPv4 given that the interfaces have no IPv4 addresses... but since every interface has a link-local address in IPv6, firewall filtering is necessary anyway, so might as well do it for IPv4 too. It should be enough to discard in filterchain=input, but if you want to be extra sure, you can also do so in rawchain=prerouting.)
Related: Make sure to disable MAC-Telnet and MAC-Winbox on those interfaces via /tool/mac-server – both as a risk on their own, and possibly as the reason RouterOS still processes IP broadcasts on those interfaces.
Maybe alternatively: /ip/vrf put all those interfaces in a separate VRF for L3 isolation. Most RouterOS services are configured to only be available on the 'main' VRF by default.
