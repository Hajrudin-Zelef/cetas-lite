---
id: collect-260926-mikrotik/mikrotik/questions-1186386-wireless-client-on-unifi-doesnt-dhcprequest-from-mikrotik-dhcp-a7dac3f5
title: "questions-1186386-wireless-client-on-unifi-doesnt-dhcprequest-from-mikrotik-dhcp-a7dac3f5"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-1186386-wireless-client-on-unifi-doesnt-dhcprequest-from-mikrotik-dhcp-a7dac3f5.md
source_anchor: ""
source_lines: [1, 26]
sha256: 17953aea9419b3e9bb4a62edb55f735f9dc2ffcc607108815c7ab250a65fb0da
---

# questions-1186386-wireless-client-on-unifi-doesnt-dhcprequest-from-mikrotik-dhcp-a7dac3f5

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I'm having problem that sometimes wireless clients doesn't receive IP from the DHCP server. The network map looks like this:
All IP are static leases configured in hEX built-in DHCP server. I have the "always broadcast" option turned off, but it seems that it will turn on on its own. All LAN clients receive IP just fine.
The UniFi AP AC Lite is flashed with LEDE Reboot. However, the same problem also occurs to the original UniFi firmware. Sometimes even the UniFi itself doesn't response to DHCP, but if I run DHCP server on my own computer it works just fine. I suspect that Mikrotik's DHCP server is at fault here.
Also, I tried to run dnsmasq DHCP server on the Unifi and there is no DHCP problem.
I enabled port mirroring on the hEX, and ran Wireshark. It seems that the router did offered DHCP, but there are no response.
(please let me know if you need more details on packet capture)
Is there are anything I should check on both Mikrotik and the LEDE UniFi AP AC?
I had this problem, appears to be a bug in the latest Mikrotik update, as described here:
There is a bug in the current RC releases which causes DHCP packets to
be sent with an invalid source MAC...
You can also force the admin MAC
on your bridge interface to be the same as the actual MAC - this seems
to work around the issue. Mikrotik support are aware of the problem
(and suggested the workaround)" - [tulluk]
. I rolled back to bugfix release and she's all good =D
If the issue is with the UniFi AP only, and it happens regardless of the firmware on the UniFi device, it may be the configuration that you have on the UniFi.
If the UniFi AP is set up with multiple, non-bridged interfaces, make sure that the IP addresses assigned to the ports not connected to the hEX network do not have overlapping subnets. If so, it may be that the response from the UniFI's DHCP client's response to the DHCP server is using the wireless interface.
You could also check to see what options the DHCPDISCOVER from the UniFI is requesting - if the MikroTik DHCP server is not fulfilling required options, the UniFi could ignore the DHCPOFFER. This is not as likely since you're seeing this with both the UniFi firmware and the LEDE firmware.
You could post the wireshark details (fully expand the DHCP Protocol portion of the dissector display) for the DHCPDISCOVER and the DHCPOFFER if the problem persists after checking these.
In version 6.43 and 6.44 of RouterOS they now add the correct SRC MAC address, but also changed the DST MAC address for their broadcast DHCP offer. There exists a unicast DHCPoffer with source and destination MAC address but that one will not work for clients behind a wireless universal repeater. (A universal repeaters replaces the clients MAC address by its own. The unicast DHCPoffer never reaches the client.) The DHCPserver uses a broadcast DHCPoffer if requested or as second option. But the Mikrotik broadcast DHCPoffer contains DST IP 255.255.255.255 (OK) , but the specific DST MAC (not OK). Destination MAC should have been ff:ff:ff:ff:ff:ff .
