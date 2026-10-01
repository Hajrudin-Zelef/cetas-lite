---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-55269-unifi-switch-us-8-60w-with-a-unifi-ap-ac-lr-2f796efd
title: "questions-55269-unifi-switch-us-8-60w-with-a-unifi-ap-ac-lr-2f796efd"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-55269-unifi-switch-us-8-60w-with-a-unifi-ap-ac-lr-2f796efd.md
source_anchor: ""
source_lines: [1, 5]
sha256: 5664e12f3666dc3090f2335fd8b511916eacfd0cdbbc53b10eb31f1bb4d2a56d
---

# questions-55269-unifi-switch-us-8-60w-with-a-unifi-ap-ac-lr-2f796efd

I've bought myself a UniFi AP-AC-LR and a UniFi Switch US 8-60w. Are those two devices compatible to use power over ethernet cable to link them?
With kind regards
The access point UniFi AP-AC-LR uses 803.3/af mode A, and the switch US 8-60w supports it (according to their datasheets).
Ubiquiti publishes its compatibility matrix confirming this, here: https://help.ubnt.com/hc/en-us/articles/115000263008--UniFi-Understanding-PoE-and-How-UniFi-Devices-are-Powered
I believe the switch also supports so-clled "Passive POE" -- which is really just power-over-CAT5 -- but this isn't a good idea for long cable runs, easy management, or standards compliance.
