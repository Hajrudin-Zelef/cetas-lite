---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/davescodemusings-unifi-opnsense-blob-head-docs-iot-md-f31a1288-2
title: "davescodemusings-unifi-opnsense-blob-head-docs-iot-md-f31a1288"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/davescodemusings-unifi-opnsense-blob-head-docs-iot-md-f31a1288.md
source_anchor: ""
source_lines: [84, 87]
sha256: 479f879f8c76796d4f8b4c55a9acc663c5e58706b27829378d24200257d762a2
---

# davescodemusings-unifi-opnsense-blob-head-docs-iot-md-f31a1288

Next, navigate to UniFi Devices in the UniFi console. Select the switch between the UniFi access point and the OPNSense firewall. Use Port Manager to examine the switch ports where these devices attach. Ensure the ports have Traffic Restriction turned off. (Users with advanced configurations may need to add the IoT VLAN to the allow list, but for most, disabling traffic restriction is the best option.)
In the OPNSense console, access the Lobby > Dashboard. Examine Interfaces to ensure the IP configuration is correct and the OPT1 interface is up (shown in green.) Also check the Interface Statistics. The OPT1 interface should show packets in and packets. It will be a small number, but should be greater than zero. If no packets are flowing, double check the interface setup.
You can also check the OPNSense configuration from the SSH interface using the command ifconfig. You should see your IoT VLAN as part of the output.
The whole reason for creating a separate IoT network is to improve your network's security. To do this, you'll need to create firewall rules in OPNSense. While the setup for your network is unique, I have a document explaining what I configured for my network that you can use as a guide: Firewall Rules
