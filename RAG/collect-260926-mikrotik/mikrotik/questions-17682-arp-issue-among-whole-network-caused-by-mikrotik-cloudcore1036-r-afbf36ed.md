---
id: collect-260926-mikrotik/mikrotik/questions-17682-arp-issue-among-whole-network-caused-by-mikrotik-cloudcore1036-r-afbf36ed
title: "questions-17682-arp-issue-among-whole-network-caused-by-mikrotik-cloudcore1036-r-afbf36ed"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-17682-arp-issue-among-whole-network-caused-by-mikrotik-cloudcore1036-r-afbf36ed.md
source_anchor: ""
source_lines: [1, 19]
sha256: 300afba4a051fa43cc55794cd4c36185381f60bf391f2afbe16ca0abe04dfdf2
---

# questions-17682-arp-issue-among-whole-network-caused-by-mikrotik-cloudcore1036-r-afbf36ed

MY QUESTIONS ARE:
- Why configuration at layer 2 at layer 3 device at different site of LAN caused problem at another part of the network ?!
- What kind of mechanism at CC1036 has been activated that cause gaps in the ARP tables across the entire network many devices further ?!
At above picture You can see part of my network scheme. Main router based on Mikrotik CloudCore1036 within one bridge and HP switches 1810-24G and 1810-8G - all default configured within default VLAN. All CC1036 interfaces were arp-proxy and arp-enabled mode.
- CloudCore 1036 - main router with one bridge and included one link bond (causing issue balanced-rr and not causing issue 802.3ad,
- swt-d-x switches - my distribution layer switches; all HP 1810-24G - J9803A with one default VLAN,
- swt-a-x switches - my acces layer switches; all HP 1810-8G - J9802A with one default VLAN,
- Server NAS- my local storage server,
- Local station x - hosts connected to the network using DHCP server defined at CC1036 MikrotikRouter,
- Bonding-1 - I bounded 3 links between CC1036 and swt-d-1 switch; at swt-d-1 I used trunks at E22,E23,E24 interfaces and at CC1036 at E1,E2,E3 I used balanced-rr protocol and I add that Bond to the bridge.
Root cause analysis:
At first look everything was fine - all hosts had access to the Internet and between them along the LAN.
First symptom was loosing connectivity to the NAS server from whole LAN site and from Internet using VPN.
Also some host from Local station 3 site couldn't localized LAN printers at Local station 1 - there was PING problems, etc. . The same was when I tried to connecto to different hosts at Local station 3 LAN site.
Also ARP tables at hosts at Local station 3 wasn't complete !!! They didn't know about MAC addresses of their closest neighbors !!!
I suspected that there is problem with swt-d-3.
When I disconnected at swt-d-3 switch link E23 to the CC1036 trying to replace that switch - monitoring server showed that problem disappear immediately.
It was strange because problem was isolated to the Local station 3 LAN part. Theoretically frames and packets shouldn't leave swt-d-3 switch because problematic traffic was limited to hosts connected to this switch ports.
When I reconfigured CC1036 and changed bonding protocol from balanced-rr to the 802.3ad between CC1036 and swt-d-1 the entire network began working well !!! All hosts began see each other, connectivity problems with Nas server disappeared, and all ARP tables have become complete.
