---
id: collect-261001-huawei/huawei/questions-1788657-would-there-be-a-performance-difference-between-a-single-10g-l-f9c5ddf4
title: "Would there be a performance difference between a single 10G LAN going to a managed switch with 3 VLANS vs a 4-port 2.5G network card?"
domain: huawei
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["ethernet", "intel", "throughput"]
source: docs/RAG/collect-261001-huawei/questions-1788657-would-there-be-a-performance-difference-between-a-single-10g-l-f9c5ddf4.md
source_anchor: ""
source_lines: [1, 13]
sha256: 5d6ed254dee109676abfb894a82bd67c667813624930743f20cc0f02b1d36dfb
---

# Would there be a performance difference between a single 10G LAN going to a managed switch with 3 VLANS vs a 4-port 2.5G network card?

*Score : 1 | Source : https://superuser.com/questions/1788657/would-there-be-a-performance-difference-between-a-single-10g-lan-going-to-a-mana*

I'm new to networking and virtualization, but I've been making progress. One thing I've definitely learned is that networking mistakes can be quite costly!
Currently, I'm using OPNsense as my router which is hosted on Proxmox. I have an Intel X550-T2 (2 port, 10g/multi-gig ports). The card is set to passthrough from Proxmox.
I have one port as my WAN (currently Verizon Fios 1g, but I plan to upgrade to 2.5g when it's available).
The other port is going to a QNAP managed switch (QSW-M2106-4C) which has 10g and 2.5g ports. The 10g ports aren't on a VLAN and are going to a Netgear unmanaged 10g/multi-gig switch and my Synology NAS.
Several of the 2.5g ports are going to either 2.5g unmanaged switches or a wireless AP with a 2.5g connection.
I realize that in terms of internet access, I'm limited to my ISP connection of 1g. However, I have several clients (NAS, servers, PCs) that either have 2.5g, 5g, or 10g connections and the throughput internally has been great, but the throughput externally has been mixed with it ranging from 200mbps to 930mbps (using speedtest.net).
So, it's made me start thinking if the VLANs are causing a performance hit running through a single 10g connection and have been considering getting a 4-port 2.5g ethernet card (using Intel 225 or Intel 226) and separating out the interfaces on OPNsense to use physical ports rather than a VLAN.
So, my questions are:
If I'm missing a piece of the puzzle that you need to answer this, please let me know. I sincerely appreciate any feedback.
