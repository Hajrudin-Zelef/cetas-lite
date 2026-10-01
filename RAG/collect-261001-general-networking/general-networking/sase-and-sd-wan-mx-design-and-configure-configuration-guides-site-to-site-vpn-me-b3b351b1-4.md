---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1-4
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1.md
source_anchor: ""
source_lines: [108, 118]
sha256: a9c9c9aaf234345e98fce69fa9a70d5ab55457eb3c2b4289af4a0064d3722fc3
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1

Once we reconfigure the firewall upstream from the Access Point to allow outbound destination port range 32768-61000, peers are able to form a tunnel. Although the first 4 captures are filtered by UDP ports 53654 and 45540, once the firewall is opened two-way traffic can occur on any dynamically chosen ports as shown below on a packet capture taken from the wired interface of the Access Point. Now the Access Point is registered with and using with port 41091 for VPN communication.
Access Point 10.0.8.99:41091 -> WAN Appliance 208.72.143.11:53654
Access Point 10.0.8.99:41091 <- WAN Appliance 208.72.143.11:53654
Below are two examples of ACLs that could be used to allow peer-to-peer communication between Cisco Meraki VPN peers. For the second option, X.X.X.X/32 represents the IP address of the Cisco Meraki device.
| 1 | allow inside to outside, protocol: udp, source ip: any, src port: any, dst ip: any, dst port: 32768-61000 | 
| 2 | allow outside to inside established (may not be necessary with stateful firewalls) | 
-OR-
| 1 | allow inside to outside, protocol: udp, source ip: X.X.X.X/32, src port: 32768-61000, dst ip: any, dst port: 32768-61000 | 
| 2 | allow outside to inside established (may not be necessary with stateful firewalls) | 
Access Point to WAN Appliance Concentrator Testing
With an Access Point to WAN Appliance concentrator connection type, use the Test connectivity button on the Wireless network. Running the test will report which Access Points "failed to connect to the concentrator." Please note if the issue is on the concentrator side, it is likely that all Access Points will fail the test. A packet capture should be taken on the wired interface of each Access Point that failed to connect to the concentrator. Another capture should be taken from the primary Internet interface of the WAN Appliance. These captures can be analyzed to determine which site's firewall is blocking outbound IPsec communication.
