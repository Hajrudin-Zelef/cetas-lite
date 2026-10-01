---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/2025-02-how-to-set-up-vlan-on-opnsense-with-an-access-point-9a251eb8
title: "2025-02-how-to-set-up-vlan-on-opnsense-with-an-access-point-9a251eb8"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["cyber", "cybersecurity", "intel", "latency", "voice"]
source: docs/RAG/collect-261001-opnsense-pfsense/2025-02-how-to-set-up-vlan-on-opnsense-with-an-access-point-9a251eb8.md
source_anchor: ""
source_lines: [1, 116]
sha256: 06ccded9baddd6005a2d80ae4ec9e9446c4780f19d606911fe64b08326e15469
---

# 2025-02-how-to-set-up-vlan-on-opnsense-with-an-access-point-9a251eb8

Introduction
This guide provides a step-by-step tutorial on configuring VLANs on OPNsense, including hardware requirements for WLAN, firewall rules, and NAT setup.
Upgrade Mesh Routers
Here is the video version of this written tutorial:
Hardware Requirements
Before configuring VLANs on OPNsense, ensure you have the following hardware:
Deploy Monitoring Tools
- OPNsense-compatible Hardware:
  - Intel-based system with at least 2-4 NICs or a managed switch with VLAN support.
  - Minimum 4GB RAM and 32GB storage (SSD preferred).
- Managed Switch (preferred, but unmanaged switch should work as well):
  - Supports 802.1Q VLAN tagging.
  - At least one trunk port connecting to OPNsense.
-  Wireless Access Point ( WAP) with VLAN Support:
  - UniFi (requires UniFi controller), TP-Link Omada, or Aruba Instant On are good choices.
  - Should allow VLAN tagging for WLAN networks.
Step 1: Enable VLAN on OPNsense
1.1 Access OPNsense Web Interface
- Log in to OPNsense Web UI (default: https://192.168.1.1 ).
- Navigate to Interfaces > Devices > VLAN.
1.2 Create a VLAN
- Click Add.
- Configure the VLAN settings:
  - Device: vlan0.30 (leave empty if you want OPNSense to generate a device name).
  - Parent Interface: Select the LAN or the interface connected to the managed switch.
  - VLAN Tag: Set a unique VLAN ID (e.g., 30 for Guest WiFi, 802.1Q VLAN tag (between 1 and 4094)
  - VLAN Priority(PCP): Determines priority (0-7). (e.g., Best Effort (0, default).
  - Description: Name it (e.g., GUEST ).
- Device: 
- Click Save and Apply Changes.
| PCP Value | Priority Level | Traffic Type | 
|---|---|---|
| 0 | Best Effort | Default traffic, no special priority | 
| 1 | Background | Low priority (e.g., bulk data transfers) | 
| 2 | Spare | Reserved (not commonly used) | 
| 3 | Excellent Effort | Standard business-critical traffic | 
| 4 | Controlled Load | Video streaming, voice traffic | 
| 5 | Video | Latency-sensitive video applications | 
| 6 | Voice | VoIP, real-time communication | 
| 7 | Network Control | Highest priority (e.g., STP, OSPF, BGP, routing updates) | 
1.3 Assign VLAN to a New Interface and Enable DHCP Server
- Go to Interfaces > Assignments.
- Locate the newly created VLAN interface and assign it.
- Click on the assigned interface and configure:
  - Enable Interface: Check the box.
  - Prevent interface removal: (optional).
  - Description: VLAN30_GUEST.
  - Static IPv4 Configuration: Assign an IP (e.g., 10.10.3.1/24 ).
    - Use Calculator.net to calculate your Subnet.
  - Click Save & Apply Changes.
- Go to Services > ISC DHCPv4 > VLAN30_GUEST.
  - Enable: Check the box.
  - Range:
    - From: 10.10.3.2
    - To: 10.10.3.254
  - DNS servers: specify a DNS Server (e.g., Quad9 DNS: 9.9.9.9).
  - Leave DNS servers blank if you wish to use local DNS.
  - Click Save.
- Click Save & Apply Changes.
Step 2: Configure Firewall Rules for VLAN
To allow traffic from VLAN to the internet while restricting access to LAN:
Compare Wireless Plans
- Create RFC1918_Networks Alias:
  - Name: RFC1918_Networks
  - Type: Network(s)
  - Content: 192.168.0.0/16,10.0.0.0/8,176.17.0.0/12
  - Description: RFC1918_Networks
- Go to Firewall > Rules > VLAN30_GUEST (New Interface).
- Click Add to create a new rule:
  - Action: Pass
  - TCP/IP Version: IPv4+6
  - Source: VLAN_30 net
  - Destination / Invert: Check the box.
  - Destination: RFC1918_Networks
  - Description: Allow VLAN_30_GUEST access to the Internet
- Click Save & Apply Changes.
Step 3: Configure NAT to redirect DNS requests to local DNS resolver
- Go to Firewall > NAT > Port Forward.
- Click Add to create a new rule:
- Interface: VLAN30_GUEST
- TCP/IP Version: IPv4
- Protocol: TCP/UDP
- Destination: VLAN30_GUEST net
- Destination port range: 
  - From: DNS
  - To: DNS
- Redirect target IP: 127.0.0.1
- Redirect target port:
  - DNS (if your resolver listens to the default port (53)).
  - (other): specify the port number (e.g., 5353).
- Description: Redirect DNS requests to local DNS resolver
- Click Save & Apply Changes.
Step 4: Configure Managed Switch & WAP
4.1 Configure VLANs on Switch (skip if you are using an unmanaged switch)
- Access the switch’s web UI.
- Navigate to VLAN settings and Create VLAN 30.
- Assign Tagged VLAN 30 on the Trunk Port connected to OPNsense.
- Assign Untagged VLAN 30 to the port connected to the WiFi AP.
4.2 Configure VLAN on Wireless AP
- Access your WiFi Access Point (UniFi, TP-Link, etc.).
- Create a New SSID (e.g., Guest_WiFi ).
- Assign it to VLAN 30.
- Save settings and restart the AP (if required).
Step 5: Test VLAN Connectivity
- Connect a client device to Guest_WiFi.
- Ensure it gets an IP in the 10.10.3.1/24 range (or whichever range you selected in step 1.3).
- Test internet access.
- Try pinging LAN IP to verify VLAN isolation.
- Test DNS using DNSLeak Test.
Conclusion
You have successfully set up VLANs on OPNsense, configured firewall rules, applied NAT, and set up VLAN WiFi. Now, your VLAN clients have internet while being isolated from the main LAN.
Audit Cyber Security
Got questions or feedback? Drop a comment on our Facebook Page or YouTube!
OPNsenseVLAN #VLANSetup #OPNsenseFirewall #NATConfiguration #WiFiVLAN #SecureNetwork #NetworkSecurity #OPNsenseConfig #FirewallRules #NetworkSetup #GuestWiFi #VLANNetworking #CyberSecurity #ITNetworking #StepByStepGuide #OPNsenseTutorial #TechGuide #NetworkingTips #WiFiSetup #VLANConfiguration #SysAdmin #Networking #Cybersecurity
Want to support my channel? Check out ways to help here:
Your support helps keep valuable tech content coming!
