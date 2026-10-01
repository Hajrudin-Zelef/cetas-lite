---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1034611-esxi-switch-issue-with-opnsense-eedd33c6
title: "ESXi switch issue with OPNSense"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1034611-esxi-switch-issue-with-opnsense-eedd33c6.md
source_anchor: ""
source_lines: [1, 9]
sha256: 397513e77ba4dd5b9e35e01caaa1b98e5d08558e95a0367e57e3c4d0e3a2a918
---

# ESXi switch issue with OPNSense

*Score : 1 | Source : https://serverfault.com/questions/1034611/esxi-switch-issue-with-opnsense*

I've got an issue with configuration of 6 port ESXi. I read ESXI single vSwitch with 2 physical NICS and that does not really resolve my issues. I just can't get this configuration below to work.
I've got a 6 NIC (WAN+LAN+4xOPT) ESXi The NICs are connected as following:
The ESXi config
Configuring WAN was straightforward but the LAN is giving me headache. I just want to have DeviceA..DeviceZ be able to communicate on the LAN-network of OPNSense, but all I see in ESXi are load-balancing options..
I configured OPNSense with a LAN and a WAN interface
