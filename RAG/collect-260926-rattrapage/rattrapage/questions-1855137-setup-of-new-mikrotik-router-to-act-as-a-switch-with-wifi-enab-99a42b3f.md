---
id: collect-260926-rattrapage/rattrapage/questions-1855137-setup-of-new-mikrotik-router-to-act-as-a-switch-with-wifi-enab-99a42b3f
title: "questions-1855137-setup-of-new-mikrotik-router-to-act-as-a-switch-with-wifi-enab-99a42b3f"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-rattrapage/ai-llm/questions-1855137-setup-of-new-mikrotik-router-to-act-as-a-switch-with-wifi-enab-99a42b3f.md
source_anchor: ""
source_lines: [1, 16]
sha256: bab054e4d8dcf9c586e7e0051aa9cbf59956e77e57c601be9267bde83a05f8b7
---

# questions-1855137-setup-of-new-mikrotik-router-to-act-as-a-switch-with-wifi-enab-99a42b3f

"Switch with Wi-Fi enabled" is typically called an 'access point'. More generally, full Ethernet pass-through (both the switch and the Wi-Fi AP) means a 'bridge'; the switch functions as a bridge, the Wi-Fi AP functions as a bridge, and the router connects the two using a software bridge.
If you start with the default Mikrotik "wireless router" configuration, it already has most of it – i.e. it has all of the "LAN" switch-ports in bridge1 with hardware offload enabled – so there's not much to change (and fundamentally it's the same as many existing topics on repurposing various low-end routers as APs, only a bit tidier):
- Disable the DHCP server on bridge1. I would remove all soon-to-be-useless DHCP server configuration, including even the IP pool.
- Enable DHCP client on bridge1, and disable DHCP client on ether1. (It's fine if you just set the interface= on the existing DHCP client.)
- Add ether1 to bridge1.
- From now ether1 is no longer "the WAN port".
- Review firewall rules; I would delete all of the 'forward' rules as they are ineffective on a bridge (can leave empty or replace with a single action=drop, doesn't really matter either way) and I don't like keeping junk around. For the same reason you can remove all NAT rules.
Side note: You don't have to use ether1 specifically – since it's a bridge, any port works equally well as any other, so e.g. connecting Ethernet to the already-bridged ether2 will achieve the same results. (What's important is that you turn off the DHCP service on the bridge.)
On the other hand, if you start with a blank configuration (let's say you've factory-reset it, then connected via mactelnet/macwinbox and opted to remove the defconf), there's even less to do:
- Create a bridge; add all ether ports (as well as the wlan interfaces) to the bridge.
- Create a DHCP client on the bridge.
- Set the wlan interfaces to mode=ap, configure a SSID and a security profile.
- That's pretty much it.
In both cases, the DHCP client is purely for accessing the device itself; with all ports being bridged, the device's IP configuration has no influence on traffic that goes through.
What about using the DHCP on the first router?
Already covered by "act as a switch". The DHCP packets will pass through the bridge like any other packet.
