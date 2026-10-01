---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-arp-sec-0019-html-91034d18
title: "hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-arp-sec-0019-html-91034d18"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-arp-sec-0019-html-91034d18.md
source_anchor: ""
source_lines: [1, 24]
sha256: e2044e10340de5c3ac40988313ab67f5a9dfea68023affe83d3a3b9d3bd0e6c6
---

# hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-arp-sec-0019-html-91034d18

Configuring DAI on an access device can prevent MITM attacks and theft on authorized users' information. After DAI is configured, the device compares the source IP address, source MAC address, VLAN ID, and interface number in the received ARP packet with binding entries. If the ARP packet matches a binding entry, the device considers the ARP packet valid and allows the packet to pass through. If the ARP packet does not match a binding entry, the device considers the ARP packet invalid and discards the packet.
You can enable DAI in the interface view or the VLAN view. When DAI is enabled in an interface view, the device checks all ARP packets received on the interface against binding entries. When DAI is enabled in the VLAN view, the device checks the ARP packets received on all interfaces belonging to the VLAN against binding entries.
If you want to receive an alarm when a large number of ARP packets are generated, enable the alarm function for the ARP packets discarded by DAI. After the alarm function is enabled, the device will generate an alarm when the number of discarded ARP packets exceeds a specified threshold.
When ARP learning triggered by DHCP is enabled on the gateway, DAI can be enabled on the gateway.
This function is available only for DHCP snooping scenarios. The device enabled with DHCP snooping generates DHCP snooping binding entries when DHCP users go online. If a user uses a static IP address, you need to manually configure a static binding entry for the user. For details about the DHCP snooping configuration, see DHCP Snooping Configuration. For details on how to configure a static binding entry, see Configuring IPSG Based on a Static Binding Table.
After the DAI function is configured on the router, the port isolation and proxy ARP functions must be configured; otherwise, the DAI function does not take effect. For the configuration of port isolation, see Configuring Interface Isolation in the NetEngine AR600, AR6100, AR6200, and AR6300 Configuration Guide - Interface Management. For the configuration of proxy ARP, see Configuring Proxy ARP in the NetEngine AR600, AR6100, AR6200, and AR6300 Configuration Guide - IP Services.
The system view is displayed.
The interface view or VLAN view is displayed.
DAI is enabled.
By default, DAI is disabled.
Only LAN-side interfaces on the AR6140H-S and AR6140-16G4XG support this function.
Only LAN-side interfaces on the SRU-100H, SRU-200H, SRU-100HH, SRU-400HK, SRU-600HK, SRU-400H, and SRU-600H support this function.
Only LAN-side ports on 8FE1GE, 24GE, and 24ES2GP boards support the SRU-100H, SRU-200H, SRU-100HH, SRU-400HK, SRU-600HK, SRU-400H, and SRU-600H. This function is supported only when the MPU is used.
Or in the VLAN view, run: arp anti-attack check user-bind check-item { ip-address | mac-address | interface }*
Items for checking ARP packets based on binding entries are configured.
By default, the check items consist of IP address, MAC address, VLAN ID, and interface number.
To allow some special ARP packets that match only one or two items in binding entries to pass through, configure the device to check ARP packets according to one or two specified items in binding entries.
Items for checking ARP packets based on binding entries do not take effect on user hosts that are configured with static binding entries. These hosts check ARP packets based on all items in static binding entries.
The alarm function for ARP packets discarded by DAI is enabled.
By default, the alarm function for ARP packets discarded by DAI is disabled.
This type of alarm is generated for the ARP packets discarded by DAI on interfaces. Do not run the arp anti-attack check user-bind enable command in a VLAN and the arp anti-attack check user-bind alarm enable command on an interface in this VLAN at the same time; otherwise, the actual number of discarded ARP packets in the VLAN is different from the number of discarded packets on the interface.
Since the default interval for sending ARP alarms is 0 (that is, no ARP alarm is sent), you must run the arp anti-attack log-trap-timer time command to increase the alarm sending interval after enabling the alarm for packets discarded by DAI.
The alarm threshold of ARP packets discarded by DAI is set.
By default, the threshold on an interface is consistent with the threshold set by the arp anti-attack check user-bind alarm threshold threshold command in the system view. If the alarm threshold is not set in the system view, the default threshold on the interface is 100.
