---
id: collect-260926-mikrotik/mikrotik/questions-20387-mikrotik-crs-vlaning-and-accounting-0604d9be
title: "questions-20387-mikrotik-crs-vlaning-and-accounting-0604d9be"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-20387-mikrotik-crs-vlaning-and-accounting-0604d9be.md
source_anchor: ""
source_lines: [1, 14]
sha256: 55028f8a90d75c479aa211d0a9b6f1b50d2d03606098c086da73ad5605cdfa59
---

# questions-20387-mikrotik-crs-vlaning-and-accounting-0604d9be

This is from the port isolation part for CRS on Mikrotik wiki:
With default configuration for CRS ether2 is the uplink (gateway) port for local users so you need to add profile for it:
/interface ethernet switch port
    set ether2 isolation-leakage-profile-override=0
Then you need to set isolated profile for the users (example here is for ports 3 to 6 only, you need to add more if needed):
/interface ethernet switch port
    set ether3 isolation-leakage-profile-override=1
    set ether4 isolation-leakage-profile-override=1
    set ether5 isolation-leakage-profile-override=1
    set ether6 isolation-leakage-profile-override=1
And let only uplink port communicate with them:
/interface ethernet switch port-isolation
    add port-profile=1 ports=ether2 type=dst
This way they can reach internet, but not communicate among themselves.
