---
id: collect-261001-general-networking/general-networking/support-forum-92-ssl-vpn-troubleshooting-103415-3db92e17
title: "support-forum-92-ssl-vpn-troubleshooting-103415-3db92e17"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/support-forum-92-ssl-vpn-troubleshooting-103415-3db92e17.md
source_anchor: ""
source_lines: [1, 8]
sha256: 8f19aa34b0022df4b57f813ee545aefc273a81933551ba721f689601d36307fe
---

# support-forum-92-ssl-vpn-troubleshooting-103415-3db92e17

SSL VPN troubleshooting
Hi,
how could I troubleshoot unsuccessful SSL VPN connections to Fortigate, I have a client that send me screenshot with error "Unable to logon to the server. Your username or password may not be configured properly for this connection. (-12)"
How could I dig into this issue, other users connecting without problem so config is right, we use windows NPS as authentication for users. I checked this user domain account and is enabled and active.
I found a command:
diagnose debug application sslvpn -1
diagnose debug enable
but this is rather for live monitoring, how could I find a reason for this unsuccessful SSL VPN connection from two days ago. On Fortianalyzer in SSL & Dialup Ipsec I see only successful connections.
