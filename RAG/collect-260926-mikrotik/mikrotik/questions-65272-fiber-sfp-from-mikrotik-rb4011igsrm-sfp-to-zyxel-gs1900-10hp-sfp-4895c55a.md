---
id: collect-260926-mikrotik/mikrotik/questions-65272-fiber-sfp-from-mikrotik-rb4011igsrm-sfp-to-zyxel-gs1900-10hp-sfp-4895c55a
title: "questions-65272-fiber-sfp-from-mikrotik-rb4011igsrm-sfp-to-zyxel-gs1900-10hp-sfp-4895c55a"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "wavelength"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-65272-fiber-sfp-from-mikrotik-rb4011igsrm-sfp-to-zyxel-gs1900-10hp-sfp-4895c55a.md
source_anchor: ""
source_lines: [1, 4]
sha256: 375a3e04d860fb1486890b921e7720ae8146315b232124a78ffe6b1c29831124
---

# questions-65272-fiber-sfp-from-mikrotik-rb4011igsrm-sfp-to-zyxel-gs1900-10hp-sfp-4895c55a

There are different Ethernet fiber PHYs for different requirements. For short distances over multi-mode fiber there are -S transceivers (850 nm wavelength) and for longer distances (1+ km) over single mode there are -L transceivers (1300 nm).
You need the same type of transceiver at each end, e.g. 1000BASE-SX (for 1 Gbit/s) or 10GBASE-SR (for 10 Gbit/s). Before buying, check that a transceiver is compatible with your device and its supported port speeds (not all SFP+ 10G ports support 1G SFP modules).
Many device vendors try to lock you in with their "original" transceivers but there's a vast market of compatible 3rd party SFPs that work just as well.
Another option for 10 Gbit/s is a direct-attach copper cable (DAC) that has fitted SFP+ modules on each end. If you use those between devices from different vendors that are both locking in, you need to find a supplier that offers custom-programmed DACs with different compatibilities at their ends.
