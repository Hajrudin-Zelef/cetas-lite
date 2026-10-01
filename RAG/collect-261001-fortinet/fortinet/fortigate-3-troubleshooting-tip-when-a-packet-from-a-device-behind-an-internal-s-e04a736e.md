---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-when-a-packet-from-a-device-behind-an-internal-s-e04a736e
title: "fortigate-3-troubleshooting-tip-when-a-packet-from-a-device-behind-an-internal-s-e04a736e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-when-a-packet-from-a-device-behind-an-internal-s-e04a736e.md
source_anchor: ""
source_lines: [1, 4]
sha256: 5c795daa8e73831ea16decfe842519fae4acb6d79f9cb524ce480c2320e106e4
---

# fortigate-3-troubleshooting-tip-when-a-packet-from-a-device-behind-an-internal-s-e04a736e

Troubleshooting Tip: When a packet from a device behind an internal segmentation firewall was unable to reach a device across the internet or IPsec tunnel
| Description | This article describes the approach when a device behind an internal segmentation firewall fails to reach the destination device across the internet or IPsec VPN. | 
| Scope | FortiGate. | 
| Solution | A device behind the internal segmentation firewall might not be able to reach the destination device due to misconfiguration of NAT on firewall policy, or Phase II dropped the traffic of the IPsec tunnel because the source subnet(s) are not allowed through the VPN.    diagnose sniffer packet any 'host w.x.y.z and icmp' 4 <----- w.x.y.z is the source IP address.    ping w.x.y.z <-- w.x.y.z is the destination IP address    ipconfig /all   diagnose debug reset diagnose debug flow filter clear diagnose debug console timestamp enable diagnose debug flow show function-name enable diagnose debug flow show iprope enable diagnose debug flow filter addr w.x.y.z diagnose debug flow trace start 99 diagnose debug enable  Repeat the ping test from the PC and review the debug flow output. In some cases, the issue was due to a software firewall on the destination device. One way to confirm is to run sniffer and debug flow commands on the destination FortiGate. Alternatively, Wireshark can be used to identify if the packet reached the destination device and if it is replying back.  Note: Do not forget to disable the sniffer and debug flow once done troubleshooting:   Related article: Troubleshooting Tip: How to perform initial troubleshooting when a device is unable to reach the internet |
