---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1571561-unifi-dream-machine-uk-use-dsl-14ee2e21
title: "questions-1571561-unifi-dream-machine-uk-use-dsl-14ee2e21"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer", "optics"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1571561-unifi-dream-machine-uk-use-dsl-14ee2e21.md
source_anchor: ""
source_lines: [1, 25]
sha256: a1b100e77c9b2a4fa43887eb017650493b20055ae2468957f0e70ee69917c972
---

# questions-1571561-unifi-dream-machine-uk-use-dsl-14ee2e21

First, some terminology and diagrams:
All-In-One Router
The EE Bright Box 1 is a typical "Home Router", which means that includes all of the components necessary to facilitate the customer. ISPs generally provide these as they can have better control over the configuration, and customers don't generally like having multiple boxes. They exist for both Cable and xDSL connections.
They are actually a number of devices in one unit:
- xDSL Modem
- Router
- Switch
- Wireless Access Point
Separate Router and Modem (or Media Converter)
The UniFi Dream Machine aligns with what has been previously marketed as a "Cable Router" - which fundamentally means that it doesn't include a modem, and the user will generally require either an external modem (e.g: xDSL or DOCSIS, etc...) or media converter (i.e: FTTP).
FTTP customers will likely have this setup - I've not yet seen a consumer / home router with built-in optics for FTTP customers, but I imagine that they'll be around in the not too distant future.
Routers in "Bridge Mode"
Many DSL routers support what is known as "Bridge Mode" - where only the modem remains functional, and all of the other internal components are disabled.
This allows you to use their hardware to provide the underlying connection, and your own hardware to manage your internal network.
Summary
Unfortunately, it looks like it is not possible to configure the EE Bright Box 1 into "Bridge Mode", meaning that you have two options available.
Keep the Bright Box
Keep using the Bright Box as it is, and use the Dream Machine as a second (internal) router... this comes with some negative points:
- The Dream Machine may not be able to accurately determine its public IP address - it may instead report the IP address that the Bright Box's DHCP server gave it. This could affect things like Dynamic DNS services that it may offer, as well as reporting in the UI
- Any port forwarding will need configuration on both routers (unless you are able to configure a DMZ on the Bright Box)
- Routing and address conflicts will need to be carefully managed between the Bright Box and Dream Machine (i.e: they can't both use 192.168.0.0/24 ), which can lead to confusion when things don't work.
- You're adding a step to the path of all internet-bound traffic, meaning that latencies may rise.
Replace the Bright Box
Instead, you could purchase another xDSL modem, such as the DrayTek Vigor 130, and replace the Bright Box entirely.
This is the route I'd suggest that you take, and it aligns with the second diagram above. You may even improve your connection stability and speeds by replacing the modem built into the Bright Box as well.
