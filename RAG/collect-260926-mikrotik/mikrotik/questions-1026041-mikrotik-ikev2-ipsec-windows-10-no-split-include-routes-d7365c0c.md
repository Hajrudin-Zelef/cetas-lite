---
id: collect-260926-mikrotik/mikrotik/questions-1026041-mikrotik-ikev2-ipsec-windows-10-no-split-include-routes-d7365c0c
title: "questions-1026041-mikrotik-ikev2-ipsec-windows-10-no-split-include-routes-d7365c0c"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-1026041-mikrotik-ikev2-ipsec-windows-10-no-split-include-routes-d7365c0c.md
source_anchor: ""
source_lines: [1, 12]
sha256: 5bb10beb3074ef1229f4ff7b14dbcf74ace2ce7de6ce5edf90ae4ed74135bc46
---

# questions-1026041-mikrotik-ikev2-ipsec-windows-10-no-split-include-routes-d7365c0c

I am deploying a solution using IKEv2+ipsec with certificates to connect roadwarriors to corporate network. Mikrotik CHR is used as entry point.
All was swift until I started deploying the solution on Dell notebooks. Once connection to the router is established, laptop doesn't get split includes, and only VPN subnet is available. On the contrary, my admin PC, which is stationary workstation, has no such problems.
Windows 10 receiving split includes using DHCP. After some research I found out that for some reason, Dell-provided Windows 10 Pro 1909 fails to send DHCP request to the router. Laptops get their address, DNS, only split include routes are lost. Also, DHCP works well on Wi-Fi adapter.
What was done:
- Logs at router were examined for both laptop and admin PC. No DHCP requests was found when laptop connects.
- Traffic was sniffed at Microtik CHR: DHCP request, which comes from admin machine, doesn't come from notebook.
- Traffic was sniffed at notebook, and no DHCP requests were detected.
Rebooting, resetting ip and winsock using netsh, reverting to older wi-fi driver, deleting and re-creating WAN Miniports, enforcing DHCP for a connection, dancing around a laptop - all that didn't help.
Currently the only solution that works is a clean MSDN version of Windows 10 1909 installation. With this one, laptops get their split includes well. However, it doesn't seem a sound solution to me.
My questions are:
- What is possible cause of the problem?
- What can be done to fix it?
