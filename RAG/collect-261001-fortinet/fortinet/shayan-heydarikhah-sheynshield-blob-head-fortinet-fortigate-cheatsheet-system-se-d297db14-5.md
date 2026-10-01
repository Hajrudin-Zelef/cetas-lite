---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-se-d297db14-5
title: "Configure administrator timeout"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-se-d297db14.md
source_anchor: ""
source_lines: [1187, 1222]
sha256: d7f9a887bb63d4310ec4de53ce8779e637572d358048bc7a687b63734e0ddbc5
---

# Configure administrator timeout

CPU -> session establishment/control
NPU -> eligible data-plane forwarding
TPM encrypts the FortiGate disk.
❌ Wrong.
TPM protects cryptographic material.
TPM != Full Disk Encryption
NTPv3 is used when ntpv3 disable is configured.
❌ Wrong.
ntpv3 enable  = NTPv3
ntpv3 disable = NTPv4
Workspace changes are immediately committed globally.
❌ Wrong.
start
  |
modify
  |
commit
Changes become available globally after commit.
execute cfg save means reboot.
❌ Wrong.
execute cfg save   = save configuration
execute cfg reload = reload/reboot
Auxiliary sessions should always be enabled in SD-WAN.
❌ Wrong.
Routing symmetry must be considered. Fortinet specifically recommends evaluating the effect on symmetric return traffic in SD-WAN hub/spoke and ADVPN-type topologies.
FortiGate  · FortiOS 7.2  · FortiGate NSE4 · FortiGate NSE7 · Fortinet NSE4 notes · Fortinet NSE7 notes · FortiGate auxiliary session · FortiGate asymmetric routing · FortiGate NPU offloading · FortiGate NTP · FortiGate PTP · FortiGate TPM · FortiGate workspace mode · FortiGate configuration transaction · FortiGate troubleshooting commands · FortiGate hardening · FortiGate session troubleshooting · FortiOS CLI 
The real skill is not memorizing CLI commands.
Understand the relationship between:
Routing → Session State → CPU → NPU → Auxiliary Session → Configuration State
Once this chain is clear, troubleshooting complex FortiGate behavior becomes much easier.
SheynShield | Engineering Secure Networks
- 
YouTube — SheynShield 
  - Fortinet NSE content
  - FortiGate troubleshooting
  - Network Security Engineering
