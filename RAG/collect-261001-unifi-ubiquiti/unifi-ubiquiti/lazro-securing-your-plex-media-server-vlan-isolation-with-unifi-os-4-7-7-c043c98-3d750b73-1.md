---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/lazro-securing-your-plex-media-server-vlan-isolation-with-unifi-os-4-7-7-c043c98-3d750b73-1
title: "lazro-securing-your-plex-media-server-vlan-isolation-with-unifi-os-4-7-7-c043c98-3d750b73"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Anthropic", "Apple"]
dates: []
keywords: ["claude"]
source: docs/RAG/collect-261001-unifi-ubiquiti/lazro-securing-your-plex-media-server-vlan-isolation-with-unifi-os-4-7-7-c043c98-3d750b73.md
source_anchor: ""
source_lines: [1, 191]
sha256: 46cca827193725011140d849036399f6915608754ab76d3710d098f6c7f87c23
---

# lazro-securing-your-plex-media-server-vlan-isolation-with-unifi-os-4-7-7-c043c98-3d750b73

Securing Your Plex Media Server: VLAN Isolation with UniFi OS 4.4.9
This article was written with the help of Sonnet 4.5 on Claude.ai. All outputs were reviewed for accuracy and confirmed working on the UniFi OS and Plex system.
Overview
This guide walks through securing a Plex media server using VLAN isolation in UniFi OS 4.4.9. By isolating your media server and related services (like Deluge) on separate VLANs, you create network segmentation that limits potential security risks while maintaining controlled access from your other devices.
Hardware Used:
- Cloud Gateway Ultra (running UniFi OS 4.4.9)
- USW Flex Mini
- PoE+ Switch
- U6+ Access Point
Services Covered:
- Plex Media Server (Ubuntu 24.04.3 LTS)
- Deluge BitTorrent Client
- UniFi OS
Part 1: Creating Isolated VLANs
Access your Unifi Site Manager at https://unifi.ui.com/. You can select the site listed in the manager.
Creating a New VLAN
- Click the Settings gear icon on the left sidebar
- Navigate to Networks
- Click New Virtual Network
- Configure the network:
- Name: Plex (or your preferred name)
- Router: Default (leave as is)
- Zone: Internal
- Protocol: IPv4
- Auto-scale network: Enabled
- Advanced: Manual
- Check Isolate Network
- Check Allow Internet Access
5. Click Add
Repeat this process for any additional isolated networks you need (Apple TV network, etc.).
Assigning Devices to VLANs
For Wired Devices (Connected to Switch)
- Click UniFi Devices on the left sidebar
- Select your PoE+ Switch
- Click Port Manager on the right-hand side
- Select the port that your device is connected to (e.g., the port for Plex server)
- Under Native VLAN / Network, select your VLAN (e.g., Plex)
- Click Apply
Setting Fixed IP Address
- Click Client Devices on the left sidebar
- Select the device you want to configure (e.g., Plex server)
- Click Settings in the right-hand menu
- Under IP Settings:
- Check Virtual Network Override
- Set Network to your VLAN (e.g., Plex)
- Check Fixed IP Address (recommended for servers)
- Assign your desired static IP address
5. Click Apply
Forcing IP Update on Ubuntu
If your Ubuntu server doesn’t immediately pick up the new IP:
Restart NetworkManager:
sudo systemctl restart NetworkManager
Alternative — GUI Method:
- Disable the network adapter in Ubuntu’s network settings
- Re-enable it
- The system should request a new DHCP lease
Verify the new IP:
ip a
Part 2: Configuring Firewall Policies
UniFi OS 4.4.9 uses a Zone-based Policy Engine rather than traditional firewall rules.
Understanding Policy Components
- Source Zone: Where traffic originates (e.g., Client VLAN, User network)
- Source Port: Port on the originating device (almost always “Any”)
- Destination Zone: Where traffic is going (e.g., Plex VLAN)
- Destination Port: The service port you’re allowing (e.g., 32400 for Plex)
- Action: Block, Allow, or Reject
Critical Note: Source ports should almost always be set to “Any” because client devices use random high ports when initiating connections. Only the destination port should be restricted.
Creating Internal Access Policies
Policy 1: Allow Multiple Networks to Access Plex
- Navigate to Settings → Policy Engine
- Click Create New Policy
- Select Firewall
- Configure:
- Name: Internal to Plex
- Source Zone: Internal
- Source Networks: Select all networks that need Plex access
- Source Port: Any
- Action: Allow
- Auto Allow Return Traffic: Checked
- Destination Zone: Internal
- Destination Network: Plex
- Destination Port: 32400
- Protocol: All
5. Click Save
Policy 2: Allow Access to Deluge (if applicable)
- Navigate to Settings → Policy Engine
- Click Create New Policy
- Select Firewall
- Configure:
- Name: Internal to Deluge
- Source Zone: Internal
- Source Networks: Select networks that need Deluge access
- Source Port: Any
- Action: Allow
- Auto Allow Return Traffic: Checked
- Destination Zone: Internal
- Destination Network: Plex
- Destination Port: 8112
- Protocol: All
5. Click Save
Part 3: External Access Configuration
Security Considerations
Important: Opening ports to the internet increases your attack surface and should only be done when necessary. Before proceeding with external access configuration, consider whether you truly need it.
Alternative to External Access: Plex supports downloading media directly to mobile devices through the Plex app when you’re connected to your local network. You can sync content to your phone or tablet for offline viewing, eliminating the need to expose your server to the internet for many use cases. This is a much more secure option if your primary goal is accessing media while away from home.
Get Lazro’s stories in your inbox
Join Medium for free to get updates from this writer.
If you decide external access is necessary, proceed with the following steps to configure it as securely as possible.
Setting Up Remote Access for Plex
Using a non-standard port (32401) provides security through obscurity.
Step 1: Create Port Forwarding Rule
- Navigate to Settings → Policy Engine
- Click Create New Policy
- Select Port Forwarding
- Configure:
- Name: Plex External
- WAN Interface: WAN1 (Default)
- WAN Port: 32401 (external port)
- From: Any
- Forward IP Address: Your Plex server’s IP address (Select Device)
- Forward Port: 32400 (Plex’s internal listening port)
- Protocol: TCP/UDP
5. Click Add
Key Concept: Port translation happens here:
- External clients connect to port 32401
- Traffic is forwarded to internal port 32400 where Plex actually listens
Step 2: Create External Firewall Policy
- Navigate to Settings → Policy Engine
- Click Create New Policy
- Select Firewall
- Configure:
- Name: Internet to Plex
- Source Zone: External
- Source Region: Country of residence (or restrict to specific countries)
- Source Port: Any
- Action: Allow
- Auto Allow Return Traffic: Checked
- Destination Zone: Internal
- Destination Network: Plex
- Destination Port: 32400 (internal port after forwarding)
- IP Version: Both
- Protocol: All
5. Click Save
Step 3: Configure Plex Server
- Open Plex web interface
- Go to Settings → Remote Access
- Check “Manually specify public port”
- Enter 32401
- Click Apply
Important: Plex always listens internally on port 32400. The manual port setting only tells Plex what external port to advertise.
In Plex:
- Settings → Remote Access should show “Fully accessible outside your network”
Step 4: Verify External Access
- Visit https://www.canyouseeme.org/
- Enter port 32401
- Click Check Port
- Should show “Success”
Step 5: Fix Mobile App Relay Connection
If the Plex mobile app shows “Relay” instead of direct connection:
- Open Plex app
- Go to Profile (top right) → Settings
- Navigate to Advanced
- Tap Reset Cache
- Restart the app
- Verify: Libraries → Library (top) → See All Libraries → Should show “Remote Connection”
Part 4: Local DNS Configuration
Creating friendly hostnames makes accessing your services easier without remembering IP addresses.
Creating DNS Records
- Click Client Devices on the left sidebar
- Select your device (e.g., Plex server)
- Click Settings in the right-hand menu
- Under IP Settings:
- Check Local DNS Record
- Enter your desired hostname (e.g., plex.local )
5. Click Apply
Accessing Services
Note: DNS only resolves to IP addresses, not ports. You must include the port in the URL.
- Plex: http://plex.local:32400/web
- Deluge: http://plex.local:8112
Recommendation: Create browser bookmarks for these URLs for quick access.
Part 5: Troubleshooting Common Issues
Device Not Getting Correct IP
Problem: Device shows old IP after VLAN assignment
Solutions:
- Disable and re-enable network adapter on the device
- Restart NetworkManager: sudo systemctl restart NetworkManager
- Reboot the device
- Check that VLAN assignment is correct in UniFi
Cannot Access Service Despite Policy
Problem: Firewall policy exists but connection times out
Common Causes:
- Source port is restricted: Should be “Any”, not a specific port
