---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-1
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "data centre", "energy", "ethernet", "license", "parameters"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [1, 47]
sha256: 8392943bcb9b7aec7e8219f721ea68b7a11df2a26be01dd60a061f23b6876a48
---

# ms-meraki-campus-lan-5d88fe48

Meraki Campus LAN; Planning, Design Guidelines and Best Practices
Introduction
The Enterprise Campus
The enterprise campus is usually understood as that portion of the computing infrastructure that provides access to network communication services and resources to end-users and devices spread over a single geographic location. It might span a single floor, building or even a large group of buildings spread over an extended geographic area. Some networks will have a single campus that also acts as the core or backbone of the network and provide inter-connectivity between other portions of the overall network. The campus core can often interconnect the campus access, the data centre and WAN portions of the network. In the largest enterprises, there might be multiple campus sites distributed worldwide with each providing both end-user access and local backbone connectivity. From a technical or network engineering perspective, the concept of campus has also been understood to mean the high-speed Layer-2 and Layer-3 Ethernet switching portions of the network outside of the data centre. While all of these definitions or concepts of what a campus network is are still valid, they no longer completely describe the set of capabilities and services that comprise the campus network today.
The campus network, as defined for the purposes of the enterprise design guides, consists of the integrated elements that comprise the set of services used by a group of users and end-station devices that all share the same high-speed switching communications fabric. These include the packet-transport services (both wired and wireless), traffic identification and control (security and application optimization), traffic monitoring and management, and overall systems management and provisioning. These basic functions are implemented in such a way as to provide and directly support the higher-level services provided by the IT organization for use by the end-user community.
This document provides best practices and guidelines when deploying a Campus LAN with Meraki which covers both Wireless and Wired LAN.
Wireless LAN
Planning, Design Guidelines and Best Practices
Planning is key for a successful deployment and aims in collecting/validating the required design aspects for a given solution. The following section takes you through the whole design and planning process for Meraki Wireless LAN. Please pay attention to the key design items and how this will influence the design of WLAN but also other components of your architecture (e.g. LAN, WAN, Security, etc).
Planning Your Deployment
The following points summarizes the design aspects for a typical Wireless LAN that needs to be taken into consideration. Please refer to Meraki documentation for more information about each of the following items.
- Get an estimation of the number of users per AP (Influenced by the AP model, should be the outcome of a coverage survey AND a capacity survey)
- Determine the total number of SSIDs that are required (do not exceed 5 per AP or generally speaking per "air" field such that 802.11 probes of no more than 5 SSIDs compete for airtime)
- For each of your SSIDs, determine their visibility requirements (Open, Hidden, Scheduled, etc) based on your policy and service requirements
- For each of your SSIDs, determine their association requirements (e.g Open, PSK, iPSK, etc) based on the clients' compatibility as well as the network policy
- SSID encryption requirements (e.g WPA1&2, WPA2, WPA3 if supported, etc). Please verify that it's supported on all clients connecting to this SSID and choose the lowest common denominator for a given SSID.
- Client Roaming requirements (if any) per SSID and how that will reflect on your Radio and network settings (e.g. Layer 2 roaming, Layer 3 roaming, 802.11r, OKC, etc). Review recommendations and pay attention to caveats mentioned in the following section.
- Wireless security requirements per SSID (e.g Management Frame Protection, Mandatory DHCP, WIPS, etc)
- Splash pages needed on your SSIDs (e.g. Meraki Splash page, External Captive Portal, etc) and what is the format of your splash page (e.g. Click through, sign-on challenge, etc)
- Splash page customizations (e.g welcome message, company logo, special HTML parameters, etc)
- Is it required to have Active Directory Integration for any of your SSIDs (What is the connectivity to AD server? IP route, VPN ,etc)
- Do you require an integration with a Radius Server (What is the connectivity to Radius server, How many Radius servers, Do you need to proxy Radius traffic, Any special EAP timers, etc, Will CoA be required, Dynamic Group Policy assignment via Radius, Dynamic VLAN assignment via Radius, etc)
- Client IP assignment and DHCP (Which SSID mode is suitable for your needs, which VLAN to tag your SSID traffic, etc)
- If you are tagging SSID traffic, ensure that the Access Point is connected to a trunk switch port and that the required VLANs are allowed
- Do you need to tag Radius and other traffic in a separate VLAN other than the management VLAN? (Refer to Alternate Management Interface)
- What are your traffic shaping requirements per SSID (e.g Per SSID, per Client, per Application)
- What are your QoS requirements per SSID (Per Application settings). Remember, you will need to match this on your Wired network.
- Do you need group policies? (e.g. Per OS group policy, Per Client group policy, etc)
- What are the wired security requirements per SSID (e.g Layer 2 isolation, IPv6 DHCP guard, IPv6 RA guard, L3/7 firewall rules, etc)
- Are you integrating Cisco Umbrella with Meraki MR (Requires a valid LIC-MR-ADV license or follow the manual integration guide)
- How do you want your SSIDs to be broadcasted (e.g. On all APs, specific APs with a tag, all the time, per schedule, etc)
- Is BLE scanning required?
- Is BLE beaconing required (What UUIDs and assignment method)
- What Radio profile best suits your AP(s) (e.g Dual-band, Steering, 2.4GHz off, channel width, Min and Max Power, Bit rate, etc)
- Do you need multiple Radio profiles (e.g. per zone, per AP, per location, etc)
- If you require end-to-end segmentation inclusive of the Wireless edge (Classification, Enforcement, etc.) using Security Group Tags (Requires LIC-MR-ADV license)
- If enabling Adaptive Policy, choose the assignment method (Static vs Dynamic via Radius) and SGT per SSID
- Follow the guidance when configuring your Radius server (e.g. Cisco ISE) to enable dynamic VLAN/Group-Policy/SGT assignment
- For energy saving purposes, consider using SSID scheduling
Design Guidelines and Best Practices
To digest the information presented in the following table, please find the following navigation guide:
- Item: Design element (e.g. Wireless roaming)
- Best Practices: Available options and recommended setup for each (e.g. Bridge mode for seamless roaming)
- Notes: Additional supplementary information to explain how a feature works (e.g. For NAT mode SSID, the AP runs an internal DHCP server)
- Caution/Caveat/Consideration: Things to be aware of when choosing a design option and/or implications to the other network components (e.g. Layer 3 roaming mode requires AP to AP access on port UDP 9358)
Please pay attention to all sections in the below table to ensure that you get the best results of your Wireless LAN design
| Item | Best Practices | Notes | Caution/Caveat/Consideration | 
