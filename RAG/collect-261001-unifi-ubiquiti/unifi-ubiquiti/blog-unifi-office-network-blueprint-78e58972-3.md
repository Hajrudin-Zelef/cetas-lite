---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-unifi-office-network-blueprint-78e58972-3
title: "blog-unifi-office-network-blueprint-78e58972"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: ["2025-03"]
keywords: ["ethernet", "incident", "license", "pricing", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-unifi-office-network-blueprint-78e58972.md
source_anchor: ""
source_lines: [114, 174]
sha256: ac4e0a8edfaa9d00b72a4145e0d0c461858c50038bab60edf6fc7948b4e3b3ff
---

# blog-unifi-office-network-blueprint-78e58972

- Configure Networks: Create your VLANs as outlined in the Configuration Blueprint above. At minimum, create Corporate, Guest, and IoT networks.
- Set Admin Password: Configure a strong local admin password separate from your UI.com credentials.
- Enable IDS/IPS (post-setup): After the wizard completes, navigate to Settings → CyberSecure → Protection to enable IDS/IPS. Choose Notify or Notify and Block and select detection categories appropriate to your environment. The Pro Max supports up to 5 Gbps IDS/IPS throughput regardless of sensitivity setting. See the IDS/IPS documentation for current configuration options.
Step 3: Connect Your Switches and Access Points
With the Dream Machine Pro Max online, connect the UniFi Switch Pro Max 48 PoE to the gateway using a 10G SFP+ DAC cable between the gateway's SFP+ port and one of the switch's four SFP+ uplinks. This provides a 10 Gbps switch uplink — a standard 1GbE Ethernet cable would bottleneck aggregated traffic from all connected devices. Once connected, verify the link through the Port View page in the Network application — the Etherlighting feature makes physical port identification easy.
Each U7 Series Access Point should be connected to the switch using PoE (Power over Ethernet), simplifying the installation by reducing the need for additional power cables. Choose U7 Pro for most offices, U7 Pro Max for high-density areas, or U7 Lite for smaller spaces. For guidance on AP placement and density, see our WiFi 7 Access Points Guide. After connecting the access points, configure Wi-Fi settings in the management interface to ensure proper coverage across your office.
Step 4: Camera Setup and Integration
Next, set up your security cameras. Wired UniFi G6 Series Cameras (Bullet, Turret, Dome) connect to the switch via PoE. The G6 Instant is Wi-Fi connected and USB-C powered — it does not use an Ethernet PoE data connection (an optional PoE-to-USB-C adapter is available). The G6 generation (released March 2025) features on-device AI processing for Face Recognition and License Plate Recognition (LPR) on the Bullet, Turret, Dome, and Instant models. Choose Turret for ceiling mounts with adjustable aiming, Bullet for wall-mount perimeter coverage, or Dome for IK10 vandal-resistant installations — all three are IP66 all-weather rated.
After mounting, configure the cameras through UniFi Protect to define detection zones, enable smart alerts, and set retention policies. For help calculating storage requirements, use our Protect Storage Planning Guide.
UniFi Protect: AI-Powered Security for Your Office
Platform Capabilities
PoE Simplification
APs, cameras, and VoIP phones all draw power through the same Ethernet cable that carries data. This eliminates separate power runs to each device. See our Power over Ethernet Guide for PoE standards, budgeting, and cable-length considerations.
VPN and Remote Access
The Dream Machine Pro Max includes built-in VPN server support (WireGuard, OpenVPN, L2TP) and Teleport for zero-configuration remote access. Site Magic provides license-free SD-WAN between UniFi sites. These features require no additional licensing.
Firmware and Software Updates
UniFi regularly releases firmware and Network application updates. The current official branch is Network 10.x. Always back up your configuration before updating, test on a non-critical site first if possible, and monitor the Ubiquiti Community releases page for known issues before applying to production.
Scaling and Satellite Offices
As businesses grow, the demands on your network will increase. UniFi's modular approach lets you add switches, APs, and cameras to an existing site without replacing the core infrastructure. Cat6A cabling and multigig switching provide headroom for the expected planning period.
Smaller or Satellite Offices
For smaller satellite offices or home setups that don't require the full Pro Max stack, consider:
- UniFi Express 7 – Compact Wi-Fi 7 Cloud Gateway supporting 30+ managed UniFi devices and 300+ simultaneous users, with 2.3 Gbps IDS/IPS
- Dream Router 7 – Desktop gateway with built-in Wi-Fi 7 AP for home offices
Each Cloud Gateway hosts its own site and runs its own Network application instance. They do not join the headquarters Network application as managed devices. Site Manager and Site Magic SD-WAN provide centralized visibility and inter-site connectivity, respectively. For a small office that does warrant the full Pro Max stack — cameras, door access, and Wi-Fi 7 — see our 7-person professional office case study.
WAN Redundancy vs. Gateway HA
These solve different failure scenarios. WAN redundancy (dual ISP with failover) protects against ISP outages. Gateway HA (Shadow Mode on the Pro Max) protects against gateway hardware failure — it requires a second identical gateway with mirrored WAN and LAN connections. Evaluate which failure mode is more likely for your environment; many offices benefit more from dual-WAN than from gateway HA.
Next Steps
This architecture starts with one Pro Max 48 PoE for approximately 20–30 employees and scales to 80+ employees by adding access switches based on endpoint and PoE requirements. If your office has fewer than 15 employees and no camera/access requirements, a Cloud Gateway Ultra or Dream Router 7 with a smaller switch may be more appropriate — see our UniFi Buyer's Guide for right-sizing guidance.
After deployment, validate your setup:
- Test wired throughput between VLANs — confirm isolation policies block unauthorized cross-VLAN traffic
- Verify Wi-Fi coverage meets the -65 dBm / 25 dB SNR targets at workstation height
- Confirm camera feeds are recording and retention policies are active in Protect
- Test VPN connectivity from a remote location
- Document your VLAN, port, and firewall configuration for future reference
At iFeeltech, we specialize in helping businesses design and implement UniFi-based networks tailored to their specific requirements. Whether setting up a new office, expanding existing infrastructure, or migrating from another vendor, our team provides network design, professional cabling, and ongoing support.
Related Resources
Deep Dives by Topic
Hardware Selection:
- UniFi Gateway Comparison Guide – Compare all 2026 gateways
- UDM Pro Max Review – In-depth review of the flagship gateway
- WiFi 7 Access Points Guide – AP selection and placement
- Pro Max Etherlighting Guide – Master the new switch features
Security & Cameras:
- UniFi Protect Guide – Complete camera system setup
- Protect Storage Planning – Calculate retention and storage
- Network Security Guide – Firewall rules and threat management
Planning Tools:
- UniFi Network Configurator – Build your custom equipment list
- Network Cabling Services – Professional Cat6A installation
- Browse UniFi Store – Shop all UniFi networking equipment
Related Articles
More from UniFi Networks
UniFi Protect vs. Reolink vs. TP-Link: Which Camera System Is Best for Your Business?
UniFi Protect, Reolink, and TP-Link (Tapo vs. VIGI) compared for small business use — current 2026 lineups, real 8-camera pricing, and iFeelTech's multi-year UniFi fleet reliability data.
18 min read
UniFi Dream Machine Beast Review 2026: Hands-On SMB Verdict
A hands-on UniFi Dream Machine Beast review from an SMB integrator: official specs, real deployment notes, 100W power draw, the no-PoE catch, and who the $1,499 gateway is actually for.
13 min read
We Ran 538 Ubiquiti Devices for 4 Years. Here's What Actually Failed.
Real fleet data: 538 UniFi devices tracked over 4 years. 0.74% replacement rate, 99.99% core uptime, and five incident post-mortems from commercial sites.
15 min read
