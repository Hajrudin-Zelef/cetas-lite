---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-115004589707-mac-based-vlan-assignment-using-802-1x-in-unifi-n-eb5fa186
title: "hc-en-us-articles-115004589707-mac-based-vlan-assignment-using-802-1x-in-unifi-n-eb5fa186"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-115004589707-mac-based-vlan-assignment-using-802-1x-in-unifi-n-eb5fa186.md
source_anchor: ""
source_lines: [1, 44]
sha256: 90bf480dd217f96098aea7a161babbbd1f47844436cf3ae0f93049d7cd16b4f2
---

# hc-en-us-articles-115004589707-mac-based-vlan-assignment-using-802-1x-in-unifi-n-eb5fa186

MAC-Based VLAN Assignment Using 802.1x in UniFi Network
RADIUS-based MAC Authentication (802.1X) allows you to use your database of MAC Addresses to authenticate wired and wireless clients connecting to your network.
Note: If you don't already have a RADIUS server configured with MAC addresses, or you have a small quantity of devices, consider using the MAC Access Control List option.
Configure a RADIUS Profile
- Navigate to Settings > Networks > RADIUS Servers.
  - If using a UniFi Gateway, enable/resume the Default RADIUS profile.
  - If using a third-party RADIUS server, select Create New.
- Create a new RADIUS User (using Local Credentials for UniFi's native RADIUS server) with the following settings:
  - 
Username & Password: MAC Address of the device
    - Every User’s MAC Address must be formatted the same way: AABBCCDDEEFF (no separators)
  - 
VLAN ID: Optionally add a VLAN ID to assign the client. If it is left blank, the client will be assigned to the VLAN associated with the switch port or WiFi it is connected to.
    - If a VLAN is added:
      - Tunnel Type: 13
      - Tunnel Medium Type: 6
    - If no VLAN is added:
      - Tunnel Type: None
      - Tunnel Medium Type: None
  - If a VLAN is added:
- 
Username & Password: MAC Address of the device
- Navigate to Settings > Networks.
  - Check the box beside 802.1X Control.
  - Under RADIUS Profile, select the profile you configured in Step 1.
  - Click Apply Changes.
Note: MAC-based authentication accounts can only be used for wireless and wired clients. L2TP remote access does not apply.
Apply the Profile
Wireless Devices
- Navigate to Settings > WiFi and select your WiFi
- In your WiFi Settings, enable RADIUS MAC Authentication.
  - Select the MAC Address Format that matches the format you’ve used (see point 2.a.i, above)
Wired Devices
To apply this globally, go to Settings > Networks > Global Switch Settings. To individually configure a port, follow these steps:
- Navigate to Settings > Profiles > Ethernet Ports
- 
Create a New Profile with the following settings:
  - Primary Network: Default or another specific network
  - 802.1X Control: MAC-based
- Navigate to a UniFi Switch’s Port Manager.
  - UniFi Devices > Select a Switch > Port Manager
- Select your port.
- Select Ethernet Port Profile and choose the profile you’ve just built.
- Apply Changes.
