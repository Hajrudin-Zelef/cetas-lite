---
id: collect-260926-mikrotik/mikrotik/questions-1134601-unable-to-establish-link-between-mikrotik-router-and-mellanox-d2c9e7cd
title: "questions-1134601-unable-to-establish-link-between-mikrotik-router-and-mellanox--d2c9e7cd"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-1134601-unable-to-establish-link-between-mikrotik-router-and-mellanox--d2c9e7cd.md
source_anchor: ""
source_lines: [1, 11]
sha256: d7231a05c55a11669062b8573d23b3ef62a81bae830cacca4eb1f2ad03e0a759
---

# questions-1134601-unable-to-establish-link-between-mikrotik-router-and-mellanox--d2c9e7cd

I am trying to establish a 25 GbE link between Mikrotik CCR2004-1G-12S+2XS router and a server running Windows Server 2022 with Mellanox NIC (MCX512A-ADAT).
For cabling, I tried using the official Mikrotik DAC (XS+DA0003), unofficial FS.com with one side generic and one side "Mellanox-compatible", a pair of SFP28 from FS.com with SMF link and various other SFP/SFP+ modules just for testing.
Each time, the Mikrotik RouterOS shows a connected DAC/SFP incl. its serial number:
However, on the server (Mellanox) side, I only see information about the link being down or the cable unplugged:
I tried changing FEC mode, rate (low/high), flow control and auto-negotiation on Mikrotik, but the link still did not work.
What am I doing wrong?
Edit:
I have also tried setting the FEC and link speed manually:
mlxlink -d mt4121_pciconf0.1 --link_mode_force --speeds 25G
mlxlink -d mt4121_pciconf0.1 -k FC --fec_speed 25G
mlxlink -d mt4121_pciconf0.1 -a TG
