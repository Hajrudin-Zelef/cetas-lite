---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-client-v-61c0cbc7-2
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-client-v-61c0cbc7"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-client-v-61c0cbc7.md
source_anchor: ""
source_lines: [107, 117]
sha256: f06c5925ea3dc42d4c7b7add34364f7ed4ecc23b95214a9869061c60da030125
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-client-v-61c0cbc7

  - Note: Alternatively, run ncpa.cpl directly from Search or Command prompt to quickly access your VPN adapters.
- In the Security tab, under Data encryption > Select Require encryption (disconnect if sever declines)
- Under Authentication > Select Allow these protocols > Tick the box Unencrypted password (PAP)
- Verify that no other protocols are selected
Linux
To configure a Red Hat Linux device to connect to client VPN, see Configuring a VPN connection in Red Hat Documentation.
To configure an Ubuntu Linux device to connect to client VPN, see Connect to a VPN in Ubuntu Documentation.
The following packages, and their dependencies, are minimum requirements for Linux:
- xl2tpd to implement L2TP
- strongswan or libreswan to implement IPSec
GUI management of the connection requires the network-manager-l2tp-gnome VPN plugin.
