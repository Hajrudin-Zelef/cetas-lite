---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/lazro-securing-reolink-ip-camera-systems-nvr-vlan-isolation-with-unifi-os-4-4-9-da21d48c-1
title: "lazro-securing-reolink-ip-camera-systems-nvr-vlan-isolation-with-unifi-os-4-4-9--da21d48c"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude"]
source: docs/RAG/collect-261001-unifi-ubiquiti/lazro-securing-reolink-ip-camera-systems-nvr-vlan-isolation-with-unifi-os-4-4-9--da21d48c.md
source_anchor: ""
source_lines: [1, 154]
sha256: ec34d9a0e5efb399e43dcb6e95b069588571b1939b04ee0dc3c4037851fd856a
---

# lazro-securing-reolink-ip-camera-systems-nvr-vlan-isolation-with-unifi-os-4-4-9--da21d48c

Securing Reolink IP Camera Systems: NVR VLAN Isolation with UniFi OS 4.4.9
This article was written with the help of Sonnet 4.5 on Claude.ai. All outputs were reviewed for accuracy and confirmed working on the UniFi OS and Reolink NVR systems.
Overview
IP camera systems and Network Video Recorders (NVRs) represent a unique security challenge in home and small office networks. These devices require internet access for firmware updates, remote viewing, and cloud features, yet they handle sensitive video footage and often have security vulnerabilities. This guide demonstrates how to isolate a Reolink NVR system on a dedicated VLAN using UniFi OS 4.4.9, allowing necessary functionality while maintaining network security through controlled access.
Hardware Used:
- Cloud Gateway Ultra (running UniFi OS 4.4.9)
- USW Flex Mini
- PoE+ Switch
- U6+ Access Point
Camera System:
- Reolink NVR (hardwired to switch)
- IP cameras (connected to NVR via PoE or WiFi)
Key Principle: NVRs need internet access for updates and remote viewing, but should be isolated from other network devices to prevent unauthorized access to video feeds and limit the impact of potential security vulnerabilities.
Understanding the NVR Security Model
Why NVRs Need Special Treatment
Sensitive Data: NVR systems record video footage of your property, potentially capturing sensitive information about your daily routines, when you’re home, security vulnerabilities, and private activities.
Internet Access Requirements: Unlike printers or some IoT devices, NVRs legitimately need internet access for:
- Firmware updates from the manufacturer
- Remote viewing via mobile apps and web portals
- Cloud recording and backup features
- Time synchronization (NTP)
- Push notifications for motion detection events
Security Vulnerabilities: Camera systems and NVRs have a poor security track record:
- Outdated firmware with known vulnerabilities
- Weak default credentials
- Unnecessary services enabled by default (UPnP, P2P)
- Poor encryption implementations
- Lack of timely security updates
Attack Vectors: Compromised camera systems can be used to:
- Access and exfiltrate video footage
- Use as pivot points to attack other network devices
- Participate in botnets (Mirai botnet famously targeted IP cameras)
- Serve as reconnaissance tools for physical security breaches
The Balance: Isolation with Functionality
The challenge is maintaining functionality while implementing security:
Must Allow:
- Internet access for updates and remote viewing
- Controlled access from specific networks (laptops, phones, tablets)
- Communication between NVR and its cameras
Must Prevent:
- Unauthorized access to video feeds
- Lateral movement to other network segments if compromised
- Unnecessary exposure of the entire camera network
The solution: Place the NVR on an isolated VLAN with internet access, but use IP-based firewall policies that target only the NVR itself. This provides necessary functionality while limiting the attack surface.
Understanding Reolink NVR Architecture
How Reolink Systems Work
NVR as Central Hub: The Reolink NVR acts as the central management and recording system for all cameras. You don’t access individual cameras directly — all management, viewing, and recording happens through the NVR interface.
Camera Connection Methods:
PoE Cameras (Most Common):
- Connect directly to NVR’s built-in PoE ports
- Receive both power and data from NVR
- Operate on NVR’s internal network
- No network configuration needed
WiFi Cameras:
- Connect to your WiFi network
- Must be on the same VLAN as the NVR for discovery
- Managed through NVR console once added
- Still record to NVR storage
Part 1: Initial NVR Configuration
Important: Complete these NVR configuration steps BEFORE creating VLANs or changing network assignments in UniFi OS. This ensures you have access to configure the NVR and can verify settings before isolation.
Step 1: Initial NVR Setup
- Connect the Reolink NVR to your network (temporarily on your main network)
- Power on the NVR and complete initial setup wizard
- Set a strong admin password
- Configure basic settings (time zone, date/time)
- Add any cameras that will be connected
Step 2: Configure NVR Network Ports
Access the Reolink NVR through the client software:
- Open Reolink Client
- Navigate to Device Settings → Network → Advanced → Server Settings
- Configure the following ports:
- Basic Service Port: 9000 (default, can be changed if desired)
- Check HTTPS to enable secure connections
- HTTPS Port: 443 (default, can be changed if desired)
4. Click Save or Apply
5. Note these port numbers for the firewall policy configuration
Why Configure Ports First:
- Ensures you know exactly which ports to open in the firewall policy
- Allows you to enable HTTPS before network isolation
- Verifies NVR is functioning properly before moving to isolated VLAN
- Makes troubleshooting easier if issues arise
Step 3: Test Basic Connectivity
Before proceeding with VLAN isolation:
- Verify you can access NVR web interface
- Confirm all cameras are visible and recording
- Test live view and playback
- Verify motion detection is working
- Note the NVR’s current IP address
Once the NVR is properly configured and tested, proceed with VLAN isolation.
Part 2: Creating the Camera VLAN
Step 1: Create the Camera Network
- Click the Settings gear icon on the left sidebar
- Navigate to Networks
- Click New Virtual Network
- Configure the network:
- Name: Camera
- Router: Default (leave as is)
- Zone: Internal
- Protocol: IPv4
- Auto-scale network: Enabled
- Advanced: Manual
- Check Isolate Network
- Check Allow Internet Access
5. Click Add
Understanding the Configuration
Isolate Network (Checked): Prevents devices on this VLAN from communicating with devices on other VLANs. This ensures the NVR and any WiFi cameras cannot access or be accessed by devices on your other networks except through explicit firewall policies.
Allow Internet Access (Checked): Unlike printers, NVRs legitimately need internet access for:
- Manufacturer firmware updates
- Remote viewing through Reolink cloud or app
- Time synchronization via NTP servers
- Push notifications to mobile devices
- Cloud backup features (if enabled)
Part 3: Assigning NVR to Camera VLAN
Step 1: Configure Switch Port
Since your Reolink NVR is hardwired to your network:
- Click UniFi Devices on the left sidebar
- Select your switch (the one the NVR is connected to)
- Click Port Manager on the right-hand side
- Select the port that your NVR is connected to
- Under Native VLAN / Network, select Camera
- Click Apply
Step 2: Set Fixed IP Address for NVR
Setting a static IP is critical for the IP-based firewall policy approach:
- Click Client Devices on the left sidebar
- Find and select your Reolink NVR
- Click Settings in the right-hand menu
- Under IP Settings:
- Check Virtual Network Override
- Set Network to Camera
- Check Fixed IP Address (required for security)
- Assign a static IP address
- Check Local DNS Record (optional)
- If using DNS, enter hostname (e.g., nvr.local )
5. Click Apply
Why Static IP is Required:
With our security approach, the firewall policy targets the NVR’s specific IP address rather than the entire Camera network. If the IP changes:
- The firewall policy breaks
- You lose access to the NVR
- Security is compromised if the old IP is reused
Static IP ensures the firewall policy remains effective and access stays consistent.
Step 3: Restart NVR
Power cycle the Reolink NVR to ensure it picks up the new network configuration:
- Disconnect power from the NVR
- Wait 10 seconds
- Reconnect power
- Wait 2–3 minutes for full boot
Step 4: Verify NVR Network Assignment
- In UniFi, go to Client Devices
- Find your Reolink NVR
- Verify it shows:
- Connected to Camera VLAN
- Has the static IP you assigned
- Shows as “Online”
Part 4: Adding WiFi Cameras (If Applicable)
