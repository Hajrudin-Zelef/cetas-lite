---
id: collect-261001-general-networking/general-networking/wireless-troubleshooting-and-support-troubleshooting-troubleshooting-local-conne-df19f7e1-2
title: "wireless-troubleshooting-and-support-troubleshooting-troubleshooting-local-conne-df19f7e1"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/wireless-troubleshooting-and-support-troubleshooting-troubleshooting-local-conne-df19f7e1.md
source_anchor: ""
source_lines: [103, 142]
sha256: 681517288644b01ab851debf22a61100e8b6c903d41bbf3326bae8bda583b13e
---

# wireless-troubleshooting-and-support-troubleshooting-troubleshooting-local-conne-df19f7e1

    - AP broadcasts default SSID: “Meraki-Scanning.”
    - AP uses address 10.128.128.128, runs DHCP on SSID, and assigns an address to any associated client.
    - AP and client are connected for local configuration only.
- 
    If a default SSID is broadcast, connect your device to it.
- 
    If no known default SSIDs are present, set up a manual wireless network connection.
- 
    For the SSID name, use meraki-<MAC_Address> (for example, meraki-xx:xx:xx:xx:xx:xx). Replace the x's with the MAC address of the AP in lowercase.
- 
    After connecting, open a web browser and go to one of the local status page addresses.
- 
    Find the list of available administrative tasks in the Using the Cisco Meraki Device Local Status Page article.
Default SSIDs
Known default SSID names with their potential causes and solutions:
- 
    <SSID_name>-bad-gateway 
  - Cause: The AP's configured default gateway failed to respond to 15 consecutive ARP requests.
  - Solution: Check the AP's IP address configuration and reachability to its default gateway.
- 
    <SSID_name>-connecting 
  - Cause: The AP's SSID configured to use a VPN concentrator cannot connect.
  - Solution: Verify connectivity to the concentrator using the tools in dashboard. Confirm your local firewall does not block the connection.
- 
    <SSID_name>-scanning 
  - Cause: Similar to bad-gateway, the AP cannot connect to its default gateway.
  - Solution: Check the AP's IP address configuration and reachability to its default gateway.
- 
    meraki
- Cause: The default out-of-the-box SSID broadcasts the APs.
- Solution: Connect the AP to a network with Internet access.
- 
    Meraki Setup
- Cause: The AP has never connected to a Meraki network.
- Solution: Add the AP to a Meraki network.
Expected outcome
You connect to the default SSID, reach the local status page, and complete administrative configuration tasks on the AP.
Troubleshooting notes (firmware-specific behavior)
- 
    MR46 (and other Wi-Fi 6 and newer APs) might not broadcast any default SSIDs out of the box when running factory firmware if the AP cannot acquire an IP address (for example, on networks without a DHCP server). In this scenario, you cannot use the local status page for initial IP configuration. Connect the AP to a network with a DHCP server so it can connect to dashboard.
