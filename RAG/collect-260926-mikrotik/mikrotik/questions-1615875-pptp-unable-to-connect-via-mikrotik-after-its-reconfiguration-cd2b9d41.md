---
id: collect-260926-mikrotik/mikrotik/questions-1615875-pptp-unable-to-connect-via-mikrotik-after-its-reconfiguration-cd2b9d41
title: "questions-1615875-pptp-unable-to-connect-via-mikrotik-after-its-reconfiguration-cd2b9d41"
domain: mikrotik
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-1615875-pptp-unable-to-connect-via-mikrotik-after-its-reconfiguration-cd2b9d41.md
source_anchor: ""
source_lines: [1, 9]
sha256: 1d41b73bb299526a4568b029f8f4f549e6d833705b384a13864e3e10e088318a
---

# questions-1615875-pptp-unable-to-connect-via-mikrotik-after-its-reconfiguration-cd2b9d41

My home laptop(s) are configured to use PPTP and that initially worked well (see 1). After reconfiguration of my home connection, any laptop is unable to complete its PPTP connection (see 2), hanging after initial authentication at remote endpoint, throwing error 806 upon timeout. The same PPTP connections still work well if these laptops connect to internet anywhere else (or via mobile phone) (see 3).
The reconfiguration was needed due to poor stability of Huawei optical modem/router provided by the ISP. They advised me to lift any work off this device and give it to the MikroTik. They remotely switched their Huawei router into bridge mode and I did reset MikroTik into factory config, then added PPP module and moved PPPoE endpoint there, and all that works.
But as you can see, this new infrastructure config is no longer transparent to PPTP connections of laptops. Even if I temporarily disable MikroTik's firewall (at least I believe) by adding and enabling rules 1 and 2 (they allow everyting, no conditions given):
Note: .
PPTP service ports (on Service Ports tab above) are enabled, too.
The only NAT rule is the default one: add action=masquerade chain=srcnat comment="defconf: masquerade" ipsec-policy=out,none out-interface-list=WAN and WAN interface list consists of ether1 (as default uplink) and PPPoE (added by me).
.
- What else can I do to make PPTP go through via this router normally as before?
/ip firewall nat export)? What do you see while running a packet capture (/tool sniffer) – is there anything that reaches the router through the LAN interface but doesn't leave it through the PPP interface?/tool sniffer quick ip-protocol=tcp port=1723) and one for the data channel (/tool sniffer quick ip-protocol=gre, note GRE has no ports)... though, alternatively, you could just filter by the VPN server's IP address and that way get everything at once. In general I suspect that GRE is the issue (e.g. maybe it's not being NATed correctly).
