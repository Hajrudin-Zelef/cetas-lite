---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-23-2
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-23.md
source_anchor: ""
source_lines: [98, 110]
sha256: 66ff474b331bbe9959a3fc221dcab52a88e21722484d6177219d72ea4faae09a
---

# Summary

Create VLAN interfaces:

Add IP addresses to VLANs:

### RouterOS /32 and IP unnumbered addresses

In RouterOS, to create a point-to-point tunnel with addresses you have to use the address with a network mask of '/32' that effectively brings you the same features as some vendors unnumbered IP address.

There are 2 routers RouterA and RouterB where each is part of networks 10.22.0.0/24 and 10.23.0.0/24 respectively and to connect these routers using VLANs as a carrier with the following configuration:

RouterA:

RouterB:
