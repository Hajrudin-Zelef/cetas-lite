---
id: collect-261001-general-networking/general-networking/questions-69754-can-a-qsfp28-port-break-out-to-4-sfp-ports-to-separate-hosts-ff0cac38
title: "questions-69754-can-a-qsfp28-port-break-out-to-4-sfp-ports-to-separate-hosts-ff0cac38"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/questions-69754-can-a-qsfp28-port-break-out-to-4-sfp-ports-to-separate-hosts-ff0cac38.md
source_anchor: ""
source_lines: [1, 41]
sha256: f27baa80c1accb22a91e2d6f0e08c26b6a64d2c4afef5492f45b0098acfc1469
---

# questions-69754-can-a-qsfp28-port-break-out-to-4-sfp-ports-to-separate-hosts-ff0cac38

On most devices, that is possible. Check the manual for splitting a QSFP28 port into four SFP28 or SFP+ interfaces.
The datasheet specifically shows 28 10GbT and 8 SFP+ for 10G and lists a 40GbE, 4x10GbE, QSFP+ to 4xSFP+, passive DAC.
According to the user guide, using feature auto-breakout should automatically detect a breakout cable and use split ports:
Breakout auto-configuration
You can globally enable front-panel Ethernet ports to automatically
detect SFP pluggable media in a QSFP+ or QSFP28 port. The port
autoconfigures breakout interfaces for media type and speed. For
example, if you plug a 40G direct attach cable (DAC) with 4x10G
far-side transceivers into a QSFP28 port, the port autoconfigures in
10g-4x Interface-breakout mode.
RJ-45 ports and ports that are members of a port group do not support
breakout auto-configuration. Breakout auto-configuration is disabled
by default.
Enable breakout auto-configuration
OS10(config)# feature auto-breakout
Display breakout auto-configuration
Before you plug a cable in
Ethernet port 1/1/25:
OS10# show interface status
                ------------------------------------------------------------------
                Port       Description  Status Speed Duplex Mode Vlan Tagged-Vlans
                ------------------------------------------------------------------
                Eth 1/1/1               down   0     auto    -
                Eth 1/1/2               down   0     auto    A    1    -
                Eth 1/1/25              down   0     auto    A    1    -
                Eth 1/1/29              down   0     auto    A    1    -
                          
After you enter feature auto-breakout and plug a breakout cable in
Ethernet port 1/1/25:
OS10# show interface status
                ------------------------------------------------------------------
                Port       Description  Status Speed Duplex Mode Vlan Tagged-Vlans
                ------------------------------------------------------------------
                Eth 1/1/1               down   0     auto    -
                Eth 1/1/2               down   0     auto    A    1    -
                Eth 1/1/25:1            down   0     auto    A    1    -
                Eth 1/1/25:2            down   0     auto    A    1    -
                Eth 1/1/25:3            down   0     auto    A    1    -
                Eth 1/1/25:4            down   0     auto    A    1    -
                Eth 1/1/29              down   0     auto    A    1    -
There is no mention that all breakout ports need to link to the same host nor would that make much sense - it would imply a LAG trunk, largely negating the reason for a port breakout.
