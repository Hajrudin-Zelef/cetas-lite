---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-88765-it-it-possible-to-create-vlan-on-bridge-interface-in-opnsense-fi-83db1a11
title: "questions-88765-it-it-possible-to-create-vlan-on-bridge-interface-in-opnsense-fi-83db1a11"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-88765-it-it-possible-to-create-vlan-on-bridge-interface-in-opnsense-fi-83db1a11.md
source_anchor: ""
source_lines: [1, 7]
sha256: c2507e5db783f8cdf8f5b34d046fcd84e85b23420d0c2c967f12b637defb7b79
---

# questions-88765-it-it-possible-to-create-vlan-on-bridge-interface-in-opnsense-fi-83db1a11

Is it possible to create VLAN on bridge interface in OPNsense firewall? I have tried, but while creating VLAN, on interface call drop down it does not show bridge interface as an option.

On an OPNsense, you cannot use tagged VLANs on a bridged interface. However, you can create tagged VLAN interfaces on the participating physical ports and bridge them together. Bridging to an untagged (access) port is also possible.

Repeat the above process for each VLAN that you want to work across the OPNsense, practically creating a VLAN trunk for the physical ports.

With the OPNsense being a software device, a bridge may perform somewhat poorly (latency and bandwidth wise). I'd seriously recommend connecting the OPNsense to a bridged VLAN on a hardware switch.
