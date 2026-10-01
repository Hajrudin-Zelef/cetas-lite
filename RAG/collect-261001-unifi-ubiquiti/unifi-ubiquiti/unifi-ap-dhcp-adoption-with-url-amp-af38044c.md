---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/unifi-ap-dhcp-adoption-with-url-amp-af38044c
title: "unifi-ap-dhcp-adoption-with-url-amp-af38044c"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/unifi-ap-dhcp-adoption-with-url-amp-af38044c.md
source_anchor: ""
source_lines: [1, 5]
sha256: 465407963f8a195cc56bd7f7586b10f28c2127be315e6f6fac6ad426918c359c
---

# unifi-ap-dhcp-adoption-with-url-amp-af38044c

UniFi AP DHCP adoption with URL
UniFi APs can automatically annonce themselves to a controller for adoption via either a DHCP option (43) or the unifi DNS name. But sadly both options don't officially allow you to configure a URL instead of just an IP address. If the controller is not hosted locally this is quite annoying as its IP might change and you cannot use non-standard ports.
Luckily there is an undocumented DHCP option code hidden in UniFi's firmware which allows passing the full URL to the inform endpoint. The standard IP-based provisioning uses option 43 code 1 containing an IP address in binary format. But there is also code 2 which takes a full URL in text format. Sadly configuration of these vendor-specific options is very dependant on the used DHCP server, I can only give an example for ISC dhcpd.
With this configured, all unconfigured UniFi APs in the network will send an adoption request to the given inform endpoint.
Since this option is undocumented by Ubiquiti, it could theoretically go away at any time, but it has been there for at least a few years and three major firmware revisions (4, 5 and 6) so it seems like Ubiquiti has no interest in removing it.
