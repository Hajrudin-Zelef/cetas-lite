---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-23954985107351-unifi-zero-touch-provisioning-ztp-and-device-re-c5529e2e
title: "hc-en-us-articles-23954985107351-unifi-zero-touch-provisioning-ztp-and-device-re-c5529e2e"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-23954985107351-unifi-zero-touch-provisioning-ztp-and-device-re-c5529e2e.md
source_anchor: ""
source_lines: [1, 47]
sha256: 4a3ad6dcf086086b71c35cdc4b34fcd0830895aef99a67ba9a33c6402e6e66b2
---

# hc-en-us-articles-23954985107351-unifi-zero-touch-provisioning-ztp-and-device-re-c5529e2e

UniFi Zero-Touch Provisioning (ZTP) and Device Replacement
UniFi’s Zero-Touch Provisioning (ZTP) and Device Replacement streamline system administration by automating device adoption and configuration—before the hardware is even powered on—reducing manual steps and accelerating deployments.
With ZTP, you can pre-assign an AP to any site before it’s even unboxed. Simply ship it to the location, and once it’s plugged in, it will automatically adopt, and begin broadcasting the assigned WiFi SSIDs—no manual setup, no on-site commands, and no need for admin approval.
Device Replacement offers similar zero-touch benefits, but goes a step further by automatically provisioning all configurations tied to the original device—making it seamless to replace any UniFi AP or switch without manual reconfiguration.
Note: ZTP offers no additional benefit over the standard adoption process if you are using a local UniFi Cloud Gateway or other UniFi Console.
Zero-Touch Provisioning (ZTP)
Requirements
- 
Supported Devices: 
  - U7-Pro-Max
  - U7-Pro-XG
  - U7-Pro-XGS
  - U7-Pro-XG Black
  - U7-Pro-XGS Black
  - E7
  - E7-Campus
  - E7-Audience
  - U7-LR
  - U7-Lite
  - UDB
  - UDB-Pro
  - UDB-Pro-Sector
- UniFi Network: Version 8.2.71+
- UniFi OS: Version 4.0+
Use-Case
Unlike Cloud Gateways, which automatically discover new UniFi devices across all VLANs, Layer 3 Network Hosts such as CloudKeys, Official UniFi Hosting, or Self-Hosted Network Servers can only discover devices on the same VLAN. Adopting devices across VLANs or remote networks typically requires manual steps like SSH, DHCP options, or DNS redirection.
Zero-Touch Provisioning (ZTP) eliminates this complexity entirely. It enables seamless adoption without any manual admin action—even on Cloud Gateways—removing the need to click “Adopt” during setup.
How It Works
- Peel the section of the box covering the ZTP Code.
- Scan the QR code, or manually enter the 9-character ZTP code into your Site Manager Inventory.
- Assign the device to the appropriate UniFi site.
- Once the device is plugged in, it will auto-adopt—no additional setup required.
ZTP Code Reuse
ZTP codes remain valid unless the device is physically reset using the hardware reset button. To preserve the code for future use, always perform any factory reset through the UniFi management interface, as outlined here.
Device Replacement
Requirements
The replacement must be the exact same model as the original. Support for compatible model replacements is planned in a future release.
Use-Case
Device Replacement lets you swap out a UniFi device without any manual reconfiguration. By pre-assigning a replacement, UniFi ensures the new device automatically inherits all settings—including IP configuration, radio settings, VLAN assignments, and port profiles—making the process seamless and error-free.
How It Works
- Navigate to UniFi Network > UniFi Devices.
- Select the device you want to replace and click on Settings in the property panel.
- Select Set Replacement Device.
- Select the replacement device if it is already present in UniFi Network, or enter the MAC Address of the new device. This device’s status will now become “Waiting Replacement”.
- The new device will be automatically adopted and provisioned with the old device’s configuration as soon as both conditions are met:
  - The new device is in a “Pending Adoption” or “Online” state (i.e., it has network connectivity).
The old device is in an “Offline” state (i.e., it was unplugged, but not “Removed” or “Forgotten” from the UniFi Network Application).
