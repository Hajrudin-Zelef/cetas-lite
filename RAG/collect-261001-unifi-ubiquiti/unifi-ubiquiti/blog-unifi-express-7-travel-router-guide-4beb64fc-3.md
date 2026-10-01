---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-unifi-express-7-travel-router-guide-4beb64fc-3
title: "blog-unifi-express-7-travel-router-guide-4beb64fc"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-unifi-express-7-travel-router-guide-4beb64fc.md
source_anchor: ""
source_lines: [108, 202]
sha256: e87467e1e96be38d7fba4c434aac97f6a1ac98eb3e16597349747e43a621b5cd
---

# blog-unifi-express-7-travel-router-guide-4beb64fc

Setup time: 15-20 minutes including switch adoption into UniFi console.
Configuration Comparison at a Glance
| Configuration | Cost | Best For | Users | 
|---|---|---|---|
| Basic | ~$210 | Client sites with ethernet | 2-3 | 
| Professional | ~$310 | Multi-device deployments | 5-10 | 
Step-by-Step Setup Guide
Setting up a UX7 as a travel router takes about 15 minutes: adopt it in the UniFi app, connect the WAN port to ethernet, then enable Site Magic SD-WAN back to headquarters.
This walkthrough assumes you're starting with a brand-new UniFi Express 7. It covers the basic setup and the Site Magic configuration for site-to-site VPN connectivity.
Prerequisites
Before you begin:
- Active UniFi account (free to create at ui.com)
- UniFi Network mobile app installed (iOS or Android)
- A compatible UniFi Cloud Gateway or Independent Gateway at headquarters (for Site Magic configuration — switches and access points alone are not sufficient)
- Internet connection for initial setup
Initial Configuration
Step 1: Power and Connect the UX7
Connect the included USB-C power adapter to the UX7 and plug into a power outlet. (Battery operation is possible with a power bank verified to supply 5V/5A, but we recommend the wall adapter for initial setup and production use.)
The front display will illuminate, showing "UniFi" and initialization messages. Wait approximately 60-90 seconds for the device to boot completely. You'll see the WiFi icon appear when ready.
Step 2: Configure Using UniFi Mobile App
The recommended setup method uses the UniFi Network mobile app with Bluetooth discovery:
- Download the UniFi Network app from the App Store (iOS) or Google Play (Android)
- Open the app and sign in to your UniFi account (or create one)
- The app will automatically discover the UX7 via Bluetooth
- Follow the on-screen setup wizard to:
  - Name your site (e.g., "Field Team - Miami")
  - Configure your WiFi SSID and password
  - Set your timezone
- The setup wizard completes and the device reboots (~30 seconds)
Alternative: Connect a laptop to the UX7's LAN port and navigate to http://192.168.1.1 in a web browser to access the setup wizard directly.
Official Documentation
For detailed setup instructions, see Ubiquiti's official UniFi Express 7 Installation Guide.
Connecting to External Internet
Step 3: Connect WAN Port
For travel router operation, you need a wired ethernet connection. The most reliable scenario we've tested is connecting to an unrestricted ethernet jack at a client or supplier site—particularly when their guest WiFi blocks VPN traffic.
Ethernet Availability
Hotel and Airbnb ethernet availability varies significantly. Some locations offer easy plug-and-play access, while others may require working with the venue or have no ethernet available at all. Test before relying on this setup for critical work.
Connect an ethernet cable from your internet source to the UX7's WAN port (marked with a globe icon on the device label).
The UX7 automatically detects the connection and acquires an IP address via DHCP. You'll see the WAN status change to "Connected" on the front display.
Captive Portal Workaround
Some hotel ethernet ports require web login (captive portal). Captive portals typically authenticate a MAC address, not just the physical port. If the UX7 can't authenticate directly, connect a laptop first, complete the portal login, then configure MAC cloning on the UX7's WAN interface to match the laptop's MAC address (available under WAN settings in the UniFi console). Always check the venue's acceptable-use policy — some venues prohibit MAC cloning or router connections.
Step 4: Verify Internet Connectivity
Open the UniFi Network mobile app:
- Tap the site name (top of screen)
- Select "Settings" → "System"
- Check the WAN status shows connected with an IP address
- Verify DNS resolution is working
Test by opening a web browser on a connected device and navigating to any website.
Enabling Site Magic SD-WAN for Site-to-Site Connectivity
Site Magic is what turns the UX7 from a portable router into a remote branch of your network. Modern UniFi Cloud Gateways use Site Magic SD-WAN, managed from the cloud-based UniFi Site Manager.
HQ Network Requirement
Site Magic requires at least one site — normally headquarters — to have a publicly routable IP address. If HQ sits behind double NAT or carrier-grade NAT (CGNAT), tunnel setup can fail. The UX7 itself can sit behind NAT, which is what makes the travel use case work, but one end must be reachable. Our Site Magic setup guide covers subnet planning and the zone-based firewall rules this depends on.
Site Magic, IPsec/OpenVPN, WireGuard or Teleport?
The UX7 supports several VPN options, each suited to a different field scenario:
- Site Magic SD-WAN — best for recurring deployments where you want the UX7's entire field LAN connected back to a UniFi headquarters with automatic tunnel management and subnet routing.
- IPsec or OpenVPN site-to-site — native site-to-site modes for connecting to non-UniFi endpoints, such as a cloud VPN or a third-party concentrator.
- WireGuard client — connects the UX7 to an external WireGuard endpoint (e.g., a self-hosted server or VPN provider). Routing the field LAN through it requires configuring policy-based routing; WireGuard is not listed as a native site-to-site mode on UniFi gateways.
- Teleport — a remote-access VPN that connects individual devices running the WiFiman app back to a UniFi network. It does not route the UX7's field LAN as a whole. Use it for giving a single remote user access, not for connecting an entire field site.
Step 5: Access UniFi Site Manager
Site-to-site connectivity is managed through UniFi Site Manager, not local gateway settings:
- Log into unifi.ui.com with your UniFi account
- Navigate to the SD-WAN tab in Site Manager
- You'll see all UniFi sites associated with your account
Step 6: Configure SD-WAN Topology
- In the SD-WAN dashboard, choose a topology: Mesh (all sites connect to each other, limited to 20 sites) or Hub-and-Spoke (branch sites connect through a central hub, supports up to 1,000 locations per Ubiquiti)
- For hub-and-spoke, select your headquarters gateway as the hub
- Add the travel router site (your UX7) as a spoke or mesh member
- Select which networks/subnets should be accessible between sites
- Site Magic establishes encrypted tunnels automatically
Both sites must share the same UI Account owner. Wait 30-60 seconds for the tunnel to establish. You'll see the connection status change to "Connected" in the SD-WAN dashboard.
Site Magic Setup Guide
For detailed setup instructions, see Ubiquiti's official Site Magic SD-WAN Guide.
Step 7: Verify Site-to-Site Connectivity
From a device connected to the UX7's WiFi:
- Ping an internal server at headquarters (e.g., file server, internal web app)
- Access an internal website or application
- Map a network drive to a headquarters file share
- Confirm you can reach resources that should only be available on the corporate network
If connections fail, check:
- Firewall rules at headquarters allow traffic from the remote subnet
- DNS configuration resolves internal hostnames correctly
- Both sites show "Connected" status in Site Magic settings
Adding a Portable Switch
If you need additional wired ports, integrate a UniFi Flex Mini switch:
Step 8: Connect and Power the Switch
- Connect an ethernet cable from the UX7's LAN port to any port on the Flex Mini
- Power the Flex Mini using its included USB-C adapter (or power bank if operating on battery)
The switch's LED indicators will blink as it initializes, then show solid or pulsing lights indicating activity.
Step 9: Adopt the Switch into UniFi
The Flex Mini automatically appears in your UniFi console when connected:
- Open the UniFi Network app
- Navigate to Devices (bottom menu)
- Look for "USW-Flex-2.5G-5" with status "Pending Adoption"
- Tap the device and select "Adopt"
- Wait 60-90 seconds for adoption to complete
