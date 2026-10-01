---
id: collect-261001-automatisation-infra/automatisation-infra/netops-toolkit-scripts-netmiko-quickstart-scripts-output-examples-md-at-5c988d242f2613e2d4
title: "netops-toolkit-scripts-netmiko-quickstart-scripts-output-examples-md-at-5c988d242f2613e2d4e48ae4556d"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/netops-toolkit-scripts-netmiko-quickstart-scripts-output-examples-md-at-5c988d242f2613e2d4e48ae4556d.md
source_anchor: ""
source_lines: [1, 54]
sha256: c19740c126e4b124e3030e83781d2225416fb5fc4bb65a1599bae85ba917d5dc
---

# netops-toolkit-scripts-netmiko-quickstart-scripts-output-examples-md-at-5c988d242f2613e2d4e48ae4556d

- Router1 (IOL) - 192.168.1.10
- Router2 (IOL) - 192.168.1.11
- Switch1 (IOL-L2) - 192.168.1.12
- External Connector (Bridge mode)
- CML Free + WSL2 + Python 3 + Netmiko

Note: These scripts use default lab credentials (admin/cisco123).
**Never use default credentials on production devices.**

Connects to a single device and runs `show ip interface brief`.

```
Interface              IP-Address      OK? Method Status                Protocol
Ethernet0/0            192.168.1.10    YES NVRAM  up                    up
Ethernet0/1            unassigned      YES NVRAM  administratively down down
Ethernet0/2            unassigned      YES NVRAM  administratively down down
Ethernet0/3            unassigned      YES NVRAM  administratively down down
```
Loops through all lab devices and runs `show ip interface brief` on each.

```
========================================
Device: hostname Router1
========================================
Interface              IP-Address      OK? Method Status                Protocol
Ethernet0/0            192.168.1.10    YES NVRAM  up                    up
Ethernet0/1            unassigned      YES NVRAM  administratively down down
Ethernet0/2            unassigned      YES NVRAM  administratively down down
Ethernet0/3            unassigned      YES NVRAM  administratively down down
========================================
Device: hostname Router2
========================================
Interface              IP-Address      OK? Method Status                Protocol
Ethernet0/0            192.168.1.11    YES NVRAM  up                    up
Ethernet0/1            unassigned      YES NVRAM  administratively down down
Ethernet0/2            unassigned      YES NVRAM  administratively down down
Ethernet0/3            unassigned      YES NVRAM  administratively down down
========================================
Device: hostname Switch1
========================================
Interface              IP-Address      OK? Method Status                Protocol
Ethernet0/0            unassigned      YES unset  up                    up
Ethernet0/1            unassigned      YES unset  up                    up
Ethernet0/2            unassigned      YES unset  up                    up
Ethernet0/3            unassigned      YES unset  up                    up
Ethernet1/0            unassigned      YES unset  up                    up
Ethernet1/1            unassigned      YES unset  up                    up
Ethernet1/2            unassigned      YES unset  up                    up
Ethernet1/3            unassigned      YES unset  up                    up
Vlan1                  192.168.1.12    YES NVRAM  up                    up
```
- Python 3.x
- Netmiko (`pip install netmiko` )
- SSH-enabled Cisco devices (CML Free works great for this)
