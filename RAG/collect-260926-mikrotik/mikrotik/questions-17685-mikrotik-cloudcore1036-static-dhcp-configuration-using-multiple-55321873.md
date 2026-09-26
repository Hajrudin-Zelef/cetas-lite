---
id: collect-260926-mikrotik/mikrotik/questions-17685-mikrotik-cloudcore1036-static-dhcp-configuration-using-multiple-55321873
title: "questions-17685-mikrotik-cloudcore1036-static-dhcp-configuration-using-multiple--55321873"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-17685-mikrotik-cloudcore1036-static-dhcp-configuration-using-multiple--55321873.md
source_anchor: ""
source_lines: [1, 11]
sha256: 325d4b1062f672be76f51a35668dc2aa2e16e49a18a1b3a20804121bc3b8208d
---

# questions-17685-mikrotik-cloudcore1036-static-dhcp-configuration-using-multiple--55321873

My question is it possible to configure CC1036 like a DHCP server which would be able to:
- gives to specified host (for example Local Station 5) same IPv4 address based on his MAC LAN and MAC WLAN address - it means that no matter of how host connect using his WLAN NIC or LAN NIC (through switch or AP) it receive same IPv4 address ?
At above picture You can see part of my network scheme. Main router based on Mikrotik CloudCore1036 within one bridge and HP switches 1810-24G and 1810-8G - all default configured within default VLAN. All CC1036 interfaces were arp-proxy and arp-enabled mode.
- CloudCore 1036 - main router with one bridge and included one link bond using 802.3ad protocol,
- rtr-a-x - MikroTik 2011 UiAS-2HnD-IN (WiFi) router configured as ap-bridge without any DHCP server,
- swt-d-x switches - my distribution layer switches; all HP 1810-24G - J9803A with one default VLAN,
- swt-a-x switches - my acces layer switches; all HP 1810-8G - J9802A with one default VLAN,
- Server NAS- my local storage server,
- Local station x - hosts connected to the network using DHCP server defined at CC1036 MikrotikRouter,
- Bonding-1 - I bounded 3 links between CC1036 and swt-d-1 switch; at swt-d-1 I used trunks at E22,E23,E24 interfaces and at CC1036 at E1,E2,E3 I used balanced-rr protocol and I add that Bond to the bridge,
- Bonding-2 - I bounded 2 links between CC1036 and swt-d-3 switch; at swt-d-3 I used trunks at E23,E24 interfaces and at CC1036 at E7,E6 I used 802.3ad protocol and I add that Bond to the bridge.
