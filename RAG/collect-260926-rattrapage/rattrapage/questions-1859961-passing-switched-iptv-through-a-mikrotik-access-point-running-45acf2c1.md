---
id: collect-260926-rattrapage/rattrapage/questions-1859961-passing-switched-iptv-through-a-mikrotik-access-point-running-45acf2c1
title: "questions-1859961-passing-switched-iptv-through-a-mikrotik-access-point-running--45acf2c1"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-rattrapage/ai-llm/questions-1859961-passing-switched-iptv-through-a-mikrotik-access-point-running--45acf2c1.md
source_anchor: ""
source_lines: [1, 21]
sha256: 18f2a3aa1514ba1340017f793f932fc8200147b5de023fbfac2be8133c2a279b
---

# questions-1859961-passing-switched-iptv-through-a-mikrotik-access-point-running--45acf2c1

My ISP delivers IPTV over a dedicated VLAN on the WAN line. In the simplest setup (as used by the provider-supplied router), that VLAN on the external interface is just terminated and bridged through to some of the internal Ethernet interfaces – the IPTV network is entirely switched and the router doesn’t do any actual routing on that network (it doesn’t even have an IP address bound to that network).
Since running a wire from the router to my TV is difficult, I have a pair of wifi bridges. One connects to the IPTV port on the router, the other to the set-top box. Both have identical hardware, one being configured as the access point, the other one as the client, paired via WPS Push Button Connect.
So the “old” setup now is:
ISP ----- Router ----- WiFi bridge 1 · · · WiFi bridge 2 ----- IPTV receiver
                \
                 ----- Switch ----- AP (old) · · · Smartphone
                             \
                               ----- PC
----- LAN link    ===== LAN trunk link    · · · WiFi link
The router is my own, configured from scratch. This took some experimentation, see Configuring pfSense for IPTV delivered via separate VLAN on WAN link. So far this setup has worked.
Now I would like to set up a MikroTik access point (RouterOS-based) that can serve multiple wifis and will also replace the router-side wifi bridge. I have reconfigured the router interface for the wifi bridge as a trunk interface, with my home LAN and the IPTV network on different VLANs, and also set up the first Ethernet interface on the MikroTik AP in the same way. On the wireless side I have two virtual wifis configured, each bridged to the respective VLAN. I then associated the receiver-side wifi bridge with the IPTV SSID on the new AP.
So the “new” setup is:
ISP ----- Router ===== MikroTik AP · · · WiFi bridge 2 ----- IPTV receiver
                \                    ·
                 ----- Switch        · · Smartphone
                             \
                               ----- PC
----- LAN link    ===== LAN trunk link    · · · WiFi link
I can use the home LAN wifi with my PC and smartphone, but the IPTV receiver doesn’t pick up any TV streams. A packet capture on the MikroTik’s IPTV wifi interface show some traffic that is clearly IPTV traffic from the ISP (such as IGMP membership queries), so at least I know I am connected to the right network.
From my experience with the router I suspect the MikroTik AP may be causing similar problems, discarding certain packets as invalid rather than forwarding them, although the packets in question are needed for IP multicasts to work.
The bridge interface has some options related to IGMP, but I don’t know if they’re relevant or how to configure them so that everything goes through. Does anyone have an idea?
