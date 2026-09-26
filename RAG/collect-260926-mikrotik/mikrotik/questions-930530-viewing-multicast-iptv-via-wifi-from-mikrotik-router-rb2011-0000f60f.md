---
id: collect-260926-mikrotik/mikrotik/questions-930530-viewing-multicast-iptv-via-wifi-from-mikrotik-router-rb2011-0000f60f
title: "Viewing multicast iptv via wifi from mikrotik router RB2011"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-930530-viewing-multicast-iptv-via-wifi-from-mikrotik-router-rb2011-0000f60f.md
source_anchor: ""
source_lines: [1, 22]
sha256: 0c5449f58f349be9876c60706818e7bee4c9d871ccb4fed6b762ad7612c4a6ce
---

# Viewing multicast iptv via wifi from mikrotik router RB2011

*Source : https://superuser.com/questions/930530/viewing-multicast-iptv-via-wifi-from-mikrotik-router-rb2011 | Site : superuser.com | Score : 1*

I would like to view multicast iptv stream via wifi from my mikrotik router, but couldn't find such settings in it. 

I've found suggestions to install udproxy on the router. This is not possible because of a proprietary RouterOS on the device.

Is there some other way to get multicast streaming via wifi from mikrotik?

The RouterOS version is 6.29.1

The full board name is RB2011UAS-2HnD

---

## Reponse (ACCEPTEE) — score 2

Set up IGMP proxy and have a look at the multicast helper setting
http://wiki.mikrotik.com/wiki/Manual:Interface/Wireless

Mikrotik doesn't support IGMP snooping so be prepared that the wifi connection will be flooded when some other port in bridge is using multicast
