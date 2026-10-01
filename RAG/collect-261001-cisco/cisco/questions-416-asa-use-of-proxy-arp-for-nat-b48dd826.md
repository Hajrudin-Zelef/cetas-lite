---
id: collect-261001-cisco/cisco/questions-416-asa-use-of-proxy-arp-for-nat-b48dd826
title: "questions-416-asa-use-of-proxy-arp-for-nat-b48dd826"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-416-asa-use-of-proxy-arp-for-nat-b48dd826.md
source_anchor: ""
source_lines: [1, 7]
sha256: 2629192451a3a144875f2e03e58b538eeeb24fcb2f26ecd7cb04cc8394e11393
---

# questions-416-asa-use-of-proxy-arp-for-nat-b48dd826

I will be performing NAT on the outside interface of an ASA to a web server within a DMZ. I would like to disable proxy-arp on all the interfaces that I can. I know that outside interface will need to have proxy-arp enabled because of the NAT statement. Do I need to have proxy-arp enabled on the DMZ interface as well?
4 Answers 4
The web server will only need to ARP the ASA to get the MAC address for its (default) gateway. So, the ASA won't need to reply an ARP for any other IP addresses then its own. You can safely turn off proxy-arp on the DMZ interface. You are correct in your assumption that you need it for the outside interface.
You don't need proxy arp on the DMZ LAN. The web system will answer ARP for itself on that LAN.
Proxy ARP is used on the ASA's to respond to hosts that are used in STATIC NAT's on the same network.
To get around this I would recommend routing any addresses used as STATIC's down to the FW address which will get rid of the need for Proxy ARP on the ASA and anything else.
ASA will use proxy-arp to reply for the incoming request to the NAT-ed IP address, so you will only need to enable it on the outside interface.
