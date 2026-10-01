---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625-1
title: "docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters", "voice"]
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625.md
source_anchor: ""
source_lines: [1, 68]
sha256: 8404d7c623b3b07bd51de212754beb7dedd0f03b2f251e6a41d2c0dcb2f13724
---

# docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625

How to Set Up VLAN on OPNsense
Virtual Local Area Networks (VLANs) are one of the most effective ways to segment traffic within a network, improving both security and performance. In OPNsense, VLANs can be configured to isolate departments, services, or user groups, ensuring that each segment has its own dedicated path for communication. This segmentation helps reduce broadcast traffic, prevents unauthorized access between different parts of the network, and simplifies overall network management. By implementing VLANs, administrators gain more granular control over traffic flow, firewall rules, and bandwidth allocation, creating a network infrastructure that is both secure and efficient.
This guide provides a step-by-step walkthrough for OPNsense VLAN setup, from the fundamentals of VLAN functionality to advanced configurations. You will learn how to create VLAN interfaces, configure DHCP for each VLAN, and set up tailored firewall rules to enforce security policies. We will also explore use cases such as building an OPNsense guest network for visitors, maintaining configurations during updates, and expanding or modifying VLANs as your network grows. Whether you are starting from scratch or optimizing an existing environment, these instructions will help you perform a complete OPNsense VLAN configuration while ensuring your network remains reliable, scalable, and easy to manage.
In the following sections, we will begin with step-by-step instructions to create and configure VLANs in OPNsense.
1. Login to OPNsense Web Interface
Before setting up VLANs in OPNsense, you need to sign in to the system’s web interface with the following steps.
- Open your browser and navigate to the OPNsense web interface (e.g., https://192.168.1.1 ).
- Enter your admin credentials to log in. Access to this interface is necessary to perform any network configuration, including VLAN setup.
Once you’re logged in, you’ll be ready to access and manage the settings required for VLAN configuration.
2. Open the VLAN Management Section
After accessing the OPNsense dashboard, the next step is to open the VLAN configuration section. This area allows you to create and manage VLAN interfaces that are essential for segmenting your network. To open the VLAN configuration section, follow the next steps.
- 
In the left-hand menu, click on Interfaces .
- 
From the dropdown options, select Devices .
- 
Click on VLAN from the list of options. This will open the VLAN configuration page, where you can add and manage VLAN entries. All VLAN configurations are managed here.Figure 1. VLAN Interfaces on OPNsense
3. Create a New VLAN
To begin setting up a new VLAN, you’ll need to create an entry that defines its core parameters, such as the VLAN tag and the parent interface it will use. To create a new VLAN, click the + Add button located at the top right of the VLANs page. This opens a form where you can input details about the VLAN.
Proceeding to this step prepares OPNsense to handle tagged traffic as part of a virtual LAN segment.
Figure 2. Adding VLAN Interface on OPNsense
4. Define VLAN Settings
The first step in creating a new VLAN in OPNsense is to choose the parent interface to which the VLAN will be connected. The physical interface that will carry the VLAN-tagged traffic is known as the parent interface. Here are the steps for selecting the parent interface and Configuring the VLAN Tagging.
- 
In the Device field, optionally enter a name for the VLAN device (leave blank for automatic assignment).
- 
From the Parent dropdown, select the physical interface that will carry VLAN traffic (e.g., em0 for WAN or em1 for LAN).
- 
In the VLAN Tag field, enter a unique numeric ID (e.g., 20) to differentiate the VLAN traffic from other networks. This must match the VLAN ID configured on your switch.
- 
From the VLAN Priority dropdown, choose the priority level detailed below for this VLAN’s traffic.
  - Best Effort (0): This is the default priority. Traffic with this priority will be treated equally, regardless of the type of traffic.
  - Background (1): Lowest priority, non-urgent traffic
  - Excellent Effort (2): Higher than default, good for standard important traffic
  - Critical Applications (3): Prioritized for essential applications
  - Video (4): Optimized for video streams. It will be given higher priority than Best Effort and Voice traffic, so that video streaming is not interrupted.
  - Voice (5): Optimized for VoIP traffic. This priority is used for voice traffic. It will be given higher priority than Best Effort traffic, so that voice calls are not interrupted by other types of traffic.
  - Internetwork Control (6): This priority is used for network control protocols, such as DHCP, DNS, and NTP. It will be given higher priority than Best Effort, Voice, and Video traffic so that these protocols can function properly.
  - Network Control (7): This highest priority is used for network control traffic. It is used for protocols such as routing and switching. It will be given higher priority than Best Effort, Voice, Video, and Control traffic so that these protocols can function properly.
- 
Optionally, write a brief Description (e.g., "IoT Devices") to identify the VLAN.
- 
Click Save to store the settings.
These parameters tell OPNsense how to recognize and handle traffic for this VLAN.
Figure 3. Define VLAN Settings
5. Assign the New VLAN to a Network Interface
After configuring the basic VLAN settings and specifying the parent interface, the next step is to assign logical interfaces to the newly created VLAN. These logical interfaces are used to manage VLAN traffic and apply various settings as needed. Multiple VLANs can be assigned to a single parent interface or spread across multiple parent interfaces. Because of this flexibility, you can define separate routing, firewalling, and other networking rules for each VLAN. To assign logical interfaces to a VLAN in OPNsense, you may follow the next steps:
- 
Go to Interfaces -> Assignments .
- 
In the dropdown menu, find and select the new VLAN interface. Figure 4. VLAN Interface Assignment on OPNsense
- 
Click Add to include it in the interface list.
The new interface will typically appear as something like OPT1 or VLAN_20.
Figure 5. Interfaces on OPNsense
6. Enable and Configure IP Settings for the VLAN Interface
After assigning the VLAN to an interface, you must enable it and configure its IP settings so it can communicate with devices in the VLAN network. This IP will act as the default gateway for all devices connected to the VLAN. Follow the steps below to enable and configure IP settings for the VLAN settings.
- In the left sidebar under Interfaces , click the new VLAN interface (it appears with the name you gave it).
- On the Basic configuration panel, tick Enable Interface.
- (Optional) In Description , enter a clear label (e.g., vlanforIoTdevices).
- Leave Identifier (e.g., opt1) andDevice (e.g., vlan0.1) as they are, these are auto-assigned.
- In Generic configuration , setIPv4 Configuration Type to Static IPv4.
- (Optional) If you won’t use IPv6, keep IPv6 Configuration Type as None.
- After choosing Static IPv4, an IPv4 address field appears, enter the gateway address for this VLAN (e.g., 192.168.10.1) and select the prefix (e.g., /24). Make sure this subnet does not overlap with any other network on your firewall.
- (Usually for LAN/VLANs) leave Block private networks and Block bogon networks unchecked. Enable them only on WAN-like interfaces.
- Click Save at the bottom, then click Apply Changes at the top to activate the configuration.
You may configure the following options for your new VLAN interface:
- Block private networks: This option blocks traffic from IP addresses that are reserved for private networks. This is useful for security purposes, as it prevents unauthorized access to your network.
