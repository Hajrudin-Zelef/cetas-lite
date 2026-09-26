---
id: collect-260926-mikrotik/mikrotik/questions-1056639-yealink-phones-fail-to-get-ip-once-assigned-voice-vlan-via-dhc-a444932f
title: "questions-1056639-yealink-phones-fail-to-get-ip-once-assigned-voice-vlan-via-dhc-a444932f"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-1056639-yealink-phones-fail-to-get-ip-once-assigned-voice-vlan-via-dhc-a444932f.md
source_anchor: ""
source_lines: [1, 13]
sha256: 59b98c37eaec1994843b88a6b655e517d467b3d16d25b49f3e8ff5861b9306e8
---

# questions-1056639-yealink-phones-fail-to-get-ip-once-assigned-voice-vlan-via-dhc-a444932f

I am deploying Yealink IP phones (T40G, T23G) on following newtork equipment:
CCR1009-7G-1C-1S+ as a router
CRS328-24P-4S+ as switches
Config of the switch in trouble is here.
I set up DHCP option 132 to configure VLAN for phones.
Before deploying such config in production, it was proved working on hAP AC^2. However while deploying it on site I encountered strange bug, which looks like this.
A phone successfully gets IP from untagged VLAN with option 132. Then it releases this IP and requests a new one from voice VLAN. The DHCP server assigns this new IP and sends "ACK" message, which is never received by the phone.
The problem seems to be switch-related. Here (careful, it's in Russian!) an admin managed to overcome this problem by disabling VLAN MAC learning. However this option is valid for SwitchOS, while my switches run RouterOS.
Also, this reddit thread has a brief summary of what needs to be configured, but it is quite outdated (5 yrs today).
This is what sniffing at phone port shows.
Notice reply packet at 49.348 from 192.168.10.40 and its absence at 50.0 subnet:
Switches had 6.48.1 stable ROS. I tried long-term 6.47.9, it didn't help. Switch in question has its own VLAN 50 IP address, which can be pinged from the router. Firewall is disabled. Phone is updated with latest firmware.
What should I do to find the cause and resolve the problem?
