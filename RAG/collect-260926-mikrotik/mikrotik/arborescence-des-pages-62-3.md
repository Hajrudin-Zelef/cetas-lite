---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-62-3
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-62.md
source_anchor: ""
source_lines: [303, 315]
sha256: f172d0ac88156f4a6569fad2bf63699675b979d222bc20652d12e8d2f85729f7
---

# Summary

Make sure that the Windows firewall is set to allow WinBox connections through Private and/or Public network interfaces in the Windows firewall, it can be changed in *Control Panel\System and Security\Windows Defender Firewall\Allowed applications* or disable the Windows firewall.

#### I get an error '(port 20561) timed out' when connecting to routers mac address

Windows (7/8) does not allow mac connection if file and print sharing is disabled.

#### I can't find my device in WinBox IPv4 Neighbors list or MAC connection fails with "ERROR could not connect to XX-XX-XX-XX-XX-XX"

Most of the network drivers will not enable IP stack unless your host device has an IP configuration. Set IPv4 configuration on your host device.

*Sometimes the device will be discovered due to caching, but MAC connection will still fail with "ERROR: could not connect to XX:XX:XX:XX:XX:XX*

WinBox MAC-ADDRESS connection requires MTU value set to 1500, unfragmented. Other values can perform poorly - loss of connectivity can occur.
