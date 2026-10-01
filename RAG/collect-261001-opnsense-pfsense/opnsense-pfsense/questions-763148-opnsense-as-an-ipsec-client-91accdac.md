---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-763148-opnsense-as-an-ipsec-client-91accdac
title: "OPNsense as an IPsec client"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-763148-opnsense-as-an-ipsec-client-91accdac.md
source_anchor: ""
source_lines: [1, 8]
sha256: d1a64ba38d459009c37ec8a202bd5235fddcc0fd358e9293c3201140775e37b9
---

# OPNsense as an IPsec client

*Score : 1 | Source : https://unix.stackexchange.com/questions/763148/opnsense-as-an-ipsec-client*

I want to setup a permanent VPN connection from one site to another. I already correctly set up an IPsec server on one site, reachable with a fixed IPv4 and IPv6, and domain.
What I want to do now, is to set up one of my OPNsense box interfaces (let's say OPT1) to be permanently connected to this IPsec VPN, so I don't have to configure anything VPN-related on the connected computers/clients inside the network of OPT1.
Is this possible? And if yes, how?
I already tried to find guidance in this type of VPN, but I couldn't find an exact guide for my use case. Most guides are about how to set up an IPsec SERVER in OPNsense, but what I need is an IPsec CLIENT, which will provide the connection to a physical interface.
