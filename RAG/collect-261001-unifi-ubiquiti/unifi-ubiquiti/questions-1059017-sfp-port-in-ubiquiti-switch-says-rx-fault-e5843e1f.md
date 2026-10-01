---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1059017-sfp-port-in-ubiquiti-switch-says-rx-fault-e5843e1f
title: "questions-1059017-sfp-port-in-ubiquiti-switch-says-rx-fault-e5843e1f"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1059017-sfp-port-in-ubiquiti-switch-says-rx-fault-e5843e1f.md
source_anchor: ""
source_lines: [1, 6]
sha256: 4af6e2e1ee7a4800cd027954efedbbcf7d6b4fea076d7a21309f8028c8acf007
---

# questions-1059017-sfp-port-in-ubiquiti-switch-says-rx-fault-e5843e1f

guys, so I recently bought an SFP cable to hook up to my server and my switch (Ubiquiti), but today, when I tried to hook everything up in the unifi controller in the port, I was connecting the SFP, I got an error saying RX Fault.
I saw some posts saying to check the bios, and the hardware is being initialized as far as I know. After looking into the bios settings, I reset the partitions to defaults, and the rest of the settings looked like it defaulted too (First time using SFP, so I don't really know what to look for).
Controller:
Proxmox:
In proxmox, it's not detecting the connection, but after inspecting the hardware, I saw a green light blink every now, and then so I guess something is being detected, and in proxmox, it detects the hardware/ports.
eno1 and eno2 are from the SFP controller hardware.
