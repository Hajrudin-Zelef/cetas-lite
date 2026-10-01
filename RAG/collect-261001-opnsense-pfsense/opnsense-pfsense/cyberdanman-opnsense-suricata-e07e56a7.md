---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/cyberdanman-opnsense-suricata-e07e56a7
title: "cyberdanman-opnsense-suricata-e07e56a7"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/cyberdanman-opnsense-suricata-e07e56a7.md
source_anchor: ""
source_lines: [1, 3]
sha256: 236ce4ed187674a6b8633ca4c33137cc30fe9f0c4bae341648d3a28bb815eaf7
---

# cyberdanman-opnsense-suricata-e07e56a7

Configure Suricata IDS/IPS for OPNsense GOAL: Enable Intrusion detection and intrusion prevention services to monitor your network for sybsecurity threats and prevent unauthorized access.
This walk through assumes you have a functioning OPNsense firewall device configured and functioning. You will need root level access to complete the following steps.
Step 1. Installing the Suricata Plugin- Navigate to “System -> Firmware -> Plugins” and install the two intrusion detection open plugins. It is possible that Instrusion detection might already be installed on your version of OPNsense. To check, navigate to services and choose Intrustion Detection. There will be a checkbox to also enable IDS, a second checkbox to enable IPS mode (please also check promiscious mode if you intend to use VLAN). In my installation, I chose to use Suricata to monitor WAN traffic(choose WAN interface), and later I setup ZENARMOR to monitor LAN traffic.
