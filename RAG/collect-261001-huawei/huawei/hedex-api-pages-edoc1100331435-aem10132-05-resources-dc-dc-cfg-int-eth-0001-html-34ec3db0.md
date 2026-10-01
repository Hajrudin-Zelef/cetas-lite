---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-int-eth-0001-html-34ec3db0
title: "hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-int-eth-0001-html-34ec3db0"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-int-eth-0001-html-34ec3db0.md
source_anchor: ""
source_lines: [1, 24]
sha256: c20670d5b75b5165acbcfa490b13d451a6f64158ae5e45b39a872cf82d2bd183
---

# hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-int-eth-0001-html-34ec3db0

Due to hardware restrictions of interface cards, some Ethernet interfaces can work only in Layer 2 or Layer 3 mode, whereas other Ethernet interfaces can work in both Layer 2 and Layer 3 modes. An Ethernet interface works as a Layer 2 interface in Layer 2 mode and a Layer 3 interface in Layer 3 mode.
GE0/0/0 to GE0/0/8 on the AR6121E-S, AR6121-S, AR6121EC-S, and AR6121C-S can be changed from Layer 2 mode to Layer 3 mode.
GE0/0/0, GE0/0/1, GE0/0/4, GE0/0/5, and GE0/0/8 on the AR6140-S and AR6140E-S can be changed from Layer 2 mode to Layer 3 mode.
GE0/0/0 to GE0/0/11 on the AR6140H-S can be changed from Layer 2 mode to Layer 3 mode.
GE0/0/0 to GE0/0/3 on the AR611-S, AR611W-S, AR611E-S, AR611, AR631I-LTE4CN, AR631I-LTE4EA, AR611-LTE4EA, AR611W, AR611W-LTE4CN, AR611W-LTE6EA, AR617VW, AR617VW-LTE4, AR617VW-LTE4EA, AR651W-X4, and AR651-X8 can be changed from Layer 2 mode to Layer 3 mode.
GE0/0/0 to GE0/0/7 on the AR651C, AR651U-A4, AR651K, AR651, AR651W-8P, AR651W, AR651EW, AR657W, AR6120, AR6120-S, and AR6120-VW can be changed from Layer 2 mode to Layer 3 mode.
GE0/0/0 to GE0/0/8 on the AR6121K, AR6121E, and AR6121 can be changed from Layer 2 mode to Layer 3 mode.
GE0/0/0 to GE0/0/5, GE0/0/8, and GE0/0/9 on the AR651F-Lite can be changed from Layer 2 mode to Layer 3 mode.
GE0/0/0 to GE0/0/11 on the AR6140-16G4XG can be changed from Layer 2 mode to Layer 3 mode.
GE0/0/0, GE0/0/1, GE0/0/4, GE0/0/5, and GE0/0/8 on the AR6140-9G-2AC, AR6140E-9G-2AC, and AR6140K-9G-2AC can be changed from Layer 2 mode to Layer 3 mode.
LAN interfaces on the SRU-100H and SRU-200H can be changed from Layer 2 mode to Layer 3 mode.
WAN interfaces on the SRU-400H, SRU-400HK, SRU-600HK, and SRU-600H can be changed from Layer 3 mode to Layer 2 mode.
WAN interfaces on the SRU-100HH can be changed from Layer 3 mode to Layer 2 mode.
WAN interfaces on the AR6140-9G-2AC and AR6140E-9G-2AC can be changed from Layer 3 mode to Layer 2 mode.
After the reserved VLAN ID of the 8FE1GE Ethernet electrical interface card and 4ES2G-S Ethernet LAN card of the AR6140-16G4XG, AR6140H-S, AR6200 series, and AR6300 series are using the set reserved-vlan command, the working modes of all interfaces on the card can be changed from Layer 2 mode to Layer 3 mode.
After the reserved VLAN ID of the 4GE-C and 8GE-T Ethernet WAN cards are using the set reserved-vlan command, the working modes of all interfaces on the card can be changed from Layer 2 mode to Layer 3 mode.
Interfaces on the 24GE Ethernet LAN cards of the AR6200 series and AR6300 series can be changed from Layer 2 mode to Layer 3 mode.
The system view is displayed.
The Ethernet interface view is displayed.
The Ethernet interface is switched from Layer 2 mode to Layer 3 mode.
By default, an Ethernet interface works in Layer 2 mode.
When you run this command on an interface, the mode switching configuration takes effect when only attribute configurations (such as shutdown and description configurations) exist on the interface. If service configurations (such as the port link-type trunk configuration) exist on the interface, you need to clear all service configurations before running this command.
IP addresses can be assigned to Ethernet interfaces in Layer 3 mode.
Run the display interface [ interface-type [ interface-number ] ] command in any view or the display this interface command in the interface view to check the running status of an interface. The interface is a Layer 2 interface if the Switch Port field is displayed in the command output and is a Layer 3 interface if the Route Port field is displayed.
