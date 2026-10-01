---
id: collect-261001-general-networking/general-networking/questions-87497-firewall-security-policy-for-traffic-that-destined-to-originate-f80ad802
title: "questions-87497-firewall-security-policy-for-traffic-that-destined-to-originate--f80ad802"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-87497-firewall-security-policy-for-traffic-that-destined-to-originate--f80ad802.md
source_anchor: ""
source_lines: [1, 9]
sha256: 35524afcdefbec3cb1edc8f8e0cc2f7a08fa874e1d6fee8b5e7d58ff738a3c45
---

# questions-87497-firewall-security-policy-for-traffic-that-destined-to-originate--f80ad802

I was trying to set up a IPsec tunnel on the firewall. I wonder how do firewall handles the traffic that destined to / originate from the firewall ?
Since Interface Profile do not have a option to allow ike/ ipsec traffic, it only have options like ping,https......it only gives partial control on the traffic that destined to the firewall interface.
My question is do we need an extra security policy on when
- VPN Phase 1 and Phase 2 traffic that originate on FortiGate firewall to remote device
- VPN Phase 1 and Phase 2 traffic that originate on remote device to FortiGate firewall
- VPN Phase 1 and Phase 2 traffic that originate on Palo Alto firewall to remote device
- VPN Phase 1 and Phase 2 traffic that originate on remote device to Palo Alto firewall
Please be aware that I am not talking about the traffic that are passing through the firewall, and not the user traffic after the tunnel is built
The fortigate is F601E on v7.2.11 build1740 The Palo Alto is PA-5420 on 11.1.6-h3
