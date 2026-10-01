---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-networks-and-routin-883f67ce-2
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-networks-and-routin-883f67ce"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-networks-and-routin-883f67ce.md
source_anchor: ""
source_lines: [57, 69]
sha256: 3fbf9a0615cb1097955f69f6dcdbf8b49d28d9f88433de99bc237c8ee3fa97d1
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-networks-and-routin-883f67ce

Troubleshooting
The easiest way to verify NAT Exceptions is working will be to: generate traffic from the LAN of the security appliance and use the packet capture tool available on dashboard to sniff traffic on the uplink that has NAT Exceptions enabled. Observe source IP of the traffic flowing through the uplink. If traffic is still being NATed, verify your configuration and ensure the security appliance's configuration is up to date.
Duplicate subnets
Please ensure you do not have duplicate subnets on the LAN and WAN. This usually occurs when the WAN appliance is sitting behind a NAT device, ISP modem, etc. This can cause NAT Exceptions to not function as expected.
Unidirectional traffic flow outbound
If you have unidirectional traffic flow,
- Ensure you have a return route pointing to the right next hop (WAN appliance uplink IP, Virtual IP if you have one configured).
- Ensure that you are permitting the right subnets on your inbound firewall
- Take simultaneous packet captures on the LAN and WAN of the MX security appliance and identify where the traffic is being dropped
Unidirectional traffic flow inbound
- Ensure you are permitting the right traffic on the outbound firewall (Security & SD-WAN > Configure > Firewall)
- Ensure the device on the LAN is reachable (ping, check ARP with live tools on dashboard, ensure the device is on the right VLAN)
- Take simultaneous packet captures on the LAN and WAN of the WAN appliance and identify where the traffic is being dropped
