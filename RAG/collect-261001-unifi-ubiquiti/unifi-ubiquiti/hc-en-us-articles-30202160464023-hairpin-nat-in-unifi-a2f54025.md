---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-30202160464023-hairpin-nat-in-unifi-a2f54025
title: "hc-en-us-articles-30202160464023-hairpin-nat-in-unifi-a2f54025"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-30202160464023-hairpin-nat-in-unifi-a2f54025.md
source_anchor: ""
source_lines: [1, 8]
sha256: a08704880aba1bf9ef2b25686d99f0a7cb38ee98ba4150aeaf592934e899d0f1
---

# hc-en-us-articles-30202160464023-hairpin-nat-in-unifi-a2f54025

Hairpin NAT in UniFi
Hairpin NAT is a feature of NAT that is implemented in UniFi. It allows devices on the internal network to access a local server using the network's public IP address. This is particularly useful for environments where a service is hosted internally but needs to be accessed using the same domain name from both inside and outside the network.
How Hairpin NAT Works in UniFi
By default, UniFi's gateway configuration enables Hairpin NAT. When a device on the local network attempts to connect to the public IP address of the UniFi gateway, the traffic is redirected internally, ensuring that port forwarding rules apply as they would for external requests. This makes it seamless for users to access internal services without requiring separate configurations for internal and external access.
Use Cases for Hairpin NAT
- Hosting a local web or application server that must be accessible via a public domain name.
- Enabling internal devices to connect to externally accessible services without modifying DNS settings.
- Simplifying configurations for remote access to internal resources.
