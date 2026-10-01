---
id: collect-261001-automatisation-infra/automatisation-infra/api-deki-pages-1268-pdf-troubleshooting-2blocal-2bconnection-2bissues-2busing-2b-080d2af8-2
title: "api-deki-pages-1268-pdf-troubleshooting-2blocal-2bconnection-2bissues-2busing-2b-080d2af8"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/api-deki-pages-1268-pdf-troubleshooting-2blocal-2bconnection-2bissues-2busing-2b-080d2af8.md
source_anchor: ""
source_lines: [94, 150]
sha256: e3ecd53818458900ded885bd0039b126d810841d54474d338e8beae8cd3c96b3
---

# api-deki-pages-1268-pdf-troubleshooting-2blocal-2bconnection-2bissues-2busing-2b-080d2af8

Expected outcome
The AP obtains an IP address, passes Layer 1, 2, and 3 checks on the switch port, and rejoins as a gateway without the "-scanning" SSIDappearing.
Troubleshooting local connection using the default SSID
Bothap.meraki.com andmy.meraki.com are locally hosted sites useful for configuring an access point (AP) when it cannot reach the Meraki Cloud.
Possible causes
• The AP is on a static, non-DHCP network.
• Strictfirewallrules block the AP's connection to the Meraki Cloud.
• The AP has lost its Internet connection while still powered, causing it to broadcast a default SSID.
Troubleshooting steps
You canconnect to the default SSIDfor administrative tasks in the Local Status Pageby completing these steps:
1. Physically inspect the AP.
1. Check that the AP has power (refer to the LED codes section of theMR installation Guides).
2. Copy the MAC address (refer to the Locatingthe MAC Address of Cisco Meraki Devicesarticle).
2. Check for available wireless networks.
1. Check whether a known default SSID is being broadcast.
2. If the AP has no configuration from the Meraki Cloud controller, the following is expected behavior:
1. AP broadcasts default SSID: “Meraki-Scanning.”
2. AP uses address10.128.128.128, runs DHCP on SSID, and assigns an address to any associated
client.
3. AP and client are connected for local configuration only.
3. If a default SSID is broadcast, connect your device to it.
4. If no known default SSIDs are present, set up a manual wireless network connection.
1. For the SSID name, usemeraki-<MAC_Address> (for example,meraki-xx:xx:xx:xx:xx:xx). Replace the x's with theMAC addressof the APin
lowercase.
5. After connecting, open a web browser and go to one of the local status page addresses.
6. Find the list of available administrative tasks in theUsing the Cisco Meraki Device Local Status Pagearticle.
Default SSIDs
Known default SSID names with their potential causes and solutions:
• <SSID_name>-bad-gateway
◦ Cause: The AP's configured default gatewayfailed torespond to 15 consecutive ARP requests.
4

◦ Solution: Check the AP's IP address configuration and reachability to its default gateway.
• <SSID_name>-connecting
◦ Cause: The AP's SSID configured to use a VPN concentrator cannot connect.
◦ Solution: Verify connectivity to the concentrator using the tools in dashboard.Confirmyour localfirewalldoes
not block the connection.
• <SSID_name>-scanning
◦ Cause:Similar tobad-gateway, the AP cannot connect to its default gateway.
◦ Solution: Check the AP's IP address configuration and reachability to its default gateway.
• meraki
• Cause: The default out-of-the-box SSID broadcasts the APs.
• Solution: Connect the AP to a network with Internet access.
• Meraki Setup
• Cause: The AP has never connected to a Meraki network.
• Solution: Add the AP to a Meraki network.
Expected outcome
You connect to the default SSID, reach the local status page, and complete administrative configuration tasks on the AP.
Troubleshooting notes (firmware-specific behavior)
• MR46 (and other Wi-Fi 6 and newer APs) might not broadcast any default SSIDs out of the box when running factory firmware if the AP cannotacquirean
IP address (for example, on networks without a DHCP server). In this scenario, you cannot use the local status page forinitialIP configuration. Connect
the AP to a network with a DHCP server so it can connect todashboard.
Additional resources
• How do I configure my access point with a static IP address?
• Troubleshooting Meraki AP's
• Static IP Assignment on a Cisco Meraki Access Point
5
