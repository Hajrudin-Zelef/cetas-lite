---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-17785698254999-assigning-unifi-access-devices-to-a-separate-vl-1a10960e
title: "hc-en-us-articles-17785698254999-assigning-unifi-access-devices-to-a-separate-vl-1a10960e"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-17785698254999-assigning-unifi-access-devices-to-a-separate-vl-1a10960e.md
source_anchor: ""
source_lines: [1, 18]
sha256: 6635463e4d43f8071b868674dd87ee20df6dfc79e4871c8956ca389dfed266e5
---

# hc-en-us-articles-17785698254999-assigning-unifi-access-devices-to-a-separate-vl-1a10960e

Assigning UniFi Access Devices to a Separate VLAN
UniFi Access allows you to enhance security by optionally isolating Access devices on a dedicated VLAN, preventing communication with the main network. For a standard UniFi Access setup guide, click here.
Creating a VLAN for Access Devices
- Navigate to Network application > Settings > Networks > Create New.
- Follow the on-screen instructions. Learn more
Changing Access Control Hub's Port Network
- Navigate to Network application > UniFi Devices and select the console/switch to which the Access Control Hub is connected.
- Navigate to Overview > Port Manager > select a port > Native VLAN / Network and select the network you just created.
- Click Apply Changes.
Assigning Static IP Address to Access Devices (Optional)
- Navigate to Network application > UniFi Devices and select an Access device.
- Navigate to Settings > IP Settings > Fixed IP Address and enter the IP address.
- Click Apply Changes.
- Repeat the steps for all Access devices.
Note: If the IPs are not updated in Network application > UniFi Devices, try unplugging the Ethernet cable and plugging it back in.
Updating Network Settings in UniFi Access
- Navigate to Access application > Settings > General > Network.
- Select a network from the dropdown menu and click Apply Changes.
