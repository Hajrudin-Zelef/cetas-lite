---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-67-5
title: "Fast Path"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-67.md
source_anchor: ""
source_lines: [380, 405]
sha256: 57b76341d83f8aa435a9064bfd1241ffbb6a25d2f17f2280e5166a6f98c67a60
---

# Fast Path

FastTrack can process packets only in the main routing table so it is the system administrator duty to not FastTrack connections that are going through non-main routing table (thus connections that are processed with mangle action=mark-routing rules). Otherwise packets might be misrouted though the main routing table.

These rules appear as soon as there is at least one fast-tracked connection tracking entry and will disappear after the last fast-tracked connection times out in the connection table.

The connection is FastTracked until a connection is closed, timed out or the router is rebooted.

### Configuration example: excluding specific host, from being Fast-Tracked

In this example, we exclude host 192.168.88.111, from being Fast-tracked, by first accepting it with the firewall rule, both for source and destination. The idea is - not allowing the traffic to reach the FastTrack action.

Note: the "exclusion" rules, must be placed before fasttrack filters, order is important.

### Requirements

FastTrack is active if the following conditions are met:

- no mesh, metarouter interface configuration;
- sniffer, torch, and traffic generator are not running;
- "/tool mac-scan" is not actively used;
- "/tool ip-scan" is not actively used;
- FastPath and Route cache are enabled under IP/Settings (route cache condition does not apply to RouterOS v7 or newer);
- bridge FastPath is enabled if a connection is going over the bridge interface;

# Packet flow for the visually impaired

The following document in DOCX format describes the diagram in a way optimized for visually impaired people. The descriptions are by Apex CoVantage care of Benetech. They are not being updated.
