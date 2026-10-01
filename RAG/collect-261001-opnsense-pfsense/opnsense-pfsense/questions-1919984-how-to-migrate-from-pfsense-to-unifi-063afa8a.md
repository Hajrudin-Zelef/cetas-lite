---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1919984-how-to-migrate-from-pfsense-to-unifi-063afa8a
title: "questions-1919984-how-to-migrate-from-pfsense-to-unifi-063afa8a"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1919984-how-to-migrate-from-pfsense-to-unifi-063afa8a.md
source_anchor: ""
source_lines: [1, 13]
sha256: b3052f007590d027bfe6246d0e943a9e6bda650381555d13a14b17f4be60c7e6
---

# questions-1919984-how-to-migrate-from-pfsense-to-unifi-063afa8a

I'd be very grateful if I could get a sanity check of my plan?
My current network is:
- PFSense firewall doing DHCP + DNS
- Unifi Controller in a Docker
- 2 Unifi Switches and 2 APs
I have bought a Unifi Cloud Gateway Fibre. Would my steps be:
- Backup my current Controller & PFSense configs.
- Make a note of all the network settings on PFSense.
- Turn off the Controller and PFSense.
- Connect the new Cloud Gateway to my ISP router.
- Using a laptop connected to the Cloud Gateway, setup the Cloud Gateway, loading the backup config from my Controller.
- Once it's all setup, configure DHCP + DNS, assigning the static IPs I have on my network.
- Plug the Cloud Gateway into my network via the switch. Wait until it adopts all my devices.
