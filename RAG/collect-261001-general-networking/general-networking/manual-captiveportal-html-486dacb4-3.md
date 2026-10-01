---
id: collect-261001-general-networking/general-networking/manual-captiveportal-html-486dacb4-3
title: "manual-captiveportal-html-486dacb4"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-captiveportal-html-486dacb4.md
source_anchor: ""
source_lines: [169, 199]
sha256: 8a09aa6e6a0b8a73fed7969be4a425941c0a624e70851a1a07f952cceda37461
---

# manual-captiveportal-html-486dacb4

The redirection rules still need associated pass rules. These rules are also used for clients directly accessing the captive portal webserver, e.g. using port 8000 in the URL.
| Type | Firewall rule | 
| Action | Pass | 
| Interface | <Zone interface> | 
| Version | IPv4+IPv6 | 
| Protocol | TCP | 
| Direction | In | 
| Source | <Zone net> | 
| Destination | This Firewall | 
| Destination port range | 8000 + zone id | 
| Type | Firewall rule | 
| Action | Pass | 
| Interface | <Zone interface> | 
| Version | IPv4+IPv6 | 
| Protocol | TCP | 
| Direction | In | 
| Source | <Zone net> | 
| Destination | This Firewall | 
| Destination port range | 9000 + zone id | 
Default block rule for non-authenticated users
Any traffic originating from a client that is not DNS or access to the portal web page, is blocked according to the rule below.
| Type | Firewall rule | 
| Action | Block | 
| Interface | <Zone interface> | 
| Protocol | Any | 
| Direction | In | 
| Source Invert | Yes | 
| Source | __captiveportal_zone_<zone id> | 
| Destination Invert | Yes | 
| Destination | __captiveportal_zone_<zone id> | 
After the above rules, an explicit pass rule is still required to allow clients to go to the internet, as would normally be the case on any interface that has no firewall rules defined.
