---
id: collect-261001-general-networking/general-networking/hc-en-us-articles-29887064407319-using-ppsk-radius-for-multiple-vlans-on-an-ssid-0f045e8f
title: "hc-en-us-articles-29887064407319-using-ppsk-radius-for-multiple-vlans-on-an-ssid-0f045e8f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/hc-en-us-articles-29887064407319-using-ppsk-radius-for-multiple-vlans-on-an-ssid-0f045e8f.md
source_anchor: ""
source_lines: [1, 22]
sha256: 11527e56341710fcff0a5be46770b4efd683e3a6f520dab1fcd2e7a0a5bf536f
---

# hc-en-us-articles-29887064407319-using-ppsk-radius-for-multiple-vlans-on-an-ssid-0f045e8f

Using PPSK / RADIUS for Multiple VLANs On an SSID in UniFi Network
UniFi Network’s Private Pre-Shared Keys (PPSK) and RADIUS-Assigned VLAN are two powerful features enabling dynamic network segmentation on a single WiFi SSID. Whether you're managing guest access, IoT devices, or enterprise networks, these features enhance security, segment traffic, and reduce airtime utilization.
- With Private Pre-Shared Keys (PPSK), you connect to WiFi using just the SSID and a shared password—no individual usernames required. The password determines your VLAN assignment, making it compatible with all client types, including legacy devices and IoT devices.
- RADIUS uses unique user profiles with custom VLANs set for each. This will only work for devices supporting WPA2 Enterprise or WPA3 Enterprise encryption.
Private Pre-Shared Keys (PPSK)
With Private Pre-Shared Keys, you can keep all users on the same SSID without requiring a unique user profile for each user. This is a simple and effective way to segment users without requiring additional authentication infrastructure. This option is WPA2 only, so it will not work on 6 GHz wireless bands.
- Ensure your VLANs are configured. Learn more here.
- Go to Settings > WiFi and add a new wireless network.
- Enable Private Pre-Shared Keys (PPSK) in the WiFi settings.
- Create a custom password for each VLAN you want users to access.
- Click Add WiFi Network.
Note: If a user needs to switch VLANs, they can do so by “forgetting” the network in their device’s WiFi settings and then re-connecting with a new password.
RADIUS Authentication
A RADIUS server allows you to assign VLANs dynamically based on user credentials. This method is ideal for organizations needing secure, user-specific network access.
- Ensure your VLANs are configured. Learn more here.
- Create a new RADIUS profile in UniFi. Learn how here.
- Configure users with unique credentials and VLAN assignments.
- Go to Settings > WiFi and add a new wireless network.
- Under Advanced Settings, set the Security Protocol to WPA2 Enterprise or WPA3 Enterprise.
- Select the RADIUS profile created in Step 2.
- Click Add WiFi Network.
Additionally, RADIUS can assign VLANs based on MAC addresses for seamless device authentication. For more information about MAC-based VLAN assignments, click here.
