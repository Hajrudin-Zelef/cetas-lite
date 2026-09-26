---
id: collect-260926-mikrotik/mikrotik/mikrotik-capsman-with-vlans-full-setup-guide-hex-2x-cap-ax-mikrotik-masters-2
title: "mikrotik-capsman-with-vlans-full-setup-guide-hex-2x-cap-ax-mikrotik-masters"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/mikrotik-capsman-with-vlans-full-setup-guide-hex-2x-cap-ax-mikrotik-masters.md
source_anchor: ""
source_lines: [203, 247]
sha256: ebf4b9f895bc9117490e72566f3bfa4403b8dd0a2f746685ae1508ce373856ef
---

# mikrotik-capsman-with-vlans-full-setup-guide-hex-2x-cap-ax-mikrotik-masters

`/system reset-configuration no-defaults run-after-reset=cap-config.rsc`
### 🔒 Step 8: Secure with Inter-VLAN Firewall Rules

The following firewall rules are set to do the following actions. The Address Lists are used rather than specifying individual IP subnets in each rule.

| Rule No. | From (Address List) | To (Address List) | Action | Exclusions | 
| 1 | guest (10.0.30.0/24) | corporate (10.0.20.0/24) | drop |  | 
| 2 | corporate (10.0.20.0/24) | guest (10.0.30.0/24) | drop |  | 
| 3 | users (10.0.20.0/24, 10.0.30.0/24) | management (10.0.10.0/24) |  | established,related connections | 

These rules are just examples but it does still allow MGMT access to the other 2 VLANs (which is what the established and related is needed for)

```
/ip firewall address-list
add address=10.10.10.0/24 list=management
add address=10.10.20.0/24 list=corporate
add address=10.10.30.0/24 list=guest
add address=10.10.20.0/24 list=users
add address=10.10.30.0/24 list=users
/ip firewall filter
add chain=forward src-address-list=guest dst-address-list=corporate action=drop
add chain=forward src-address-list=corporate dst-address-list=guest action=drop
add chain=forward src-address-list=users dst-address-list=management connection-state=!established,!related action=drop
```
## ✅ Summary

- Centralized CAPsMAN control
- VLAN-based SSID segmentation
- Easy AP deployment and cloning
- Firewall rules for inter-VLAN security

**[PLACEHOLDER: Final CAPsMAN AP list screenshot]**

Radek
Hi,

nice, I followed guide and succeed. Now I have a tip for next chapter or articl – add another cAPax and use it in station-bridge mode with ethernet access ports and both CAPSMan SSIDs on second(or maybe better wifi1) wifi device. I trien with my 1 week Mikrotik experience and failed for now…

ilium007
There is no “Remote CAP” > capsman button on my ax2 (7.20.4) and therefore no capsman > enable checkbox. This tutorial is only 6 months old, what has changed?

ilium007
So as it tuns out….

“The “CAPsMAN” manager enable checkbox was removed in RouterOS 7; you now enable the CAPsMAN manager by using the command line interface (CLI) in the terminal. To do this, go to the terminal and type /caps-man manager set enabled=yes. “
