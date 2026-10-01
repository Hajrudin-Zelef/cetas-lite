---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-b5bea11b-3
title: "platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["omni", "throughput", "voice"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b.md
source_anchor: ""
source_lines: [120, 155]
sha256: 1f2b73ab0487f1b87e4801da6e7ff4ed5f39b5eb62591f1d28b26876d6f5d5c1
---

# platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b

    If access points have to be installed below 8 feet (~3 meters), indoor access points with integrated omni antennas or external dipole/can omni antennas are recommended.
- 
    If access points have to be installed between 8 - 25 feet (3 - 8 meters), indoor access points with external downtilt omni antennas are recommended.
Wall mounted MRs, Cisco San Francisco
When ceiling heights are too high (25+ feet) or not feasible to mount access points (hard ceiling), a wall mounted design is recommended. The access points are mounted on drywall, concrete or even metal on the exterior and interior walls of the environment. Access points are typically deployed 10-15 feet (3-5 meters) above the floor facing away from the wall. Remember to install with the LED facing down to remain visible while standing on the floor. Designing a network with wall mounted omnidirectional APs should be done carefully and should be done only if using directional antennas is not an option.
Pole mounted MR66 with Sector antennas, Cisco San Francisco
Directional Antennas
If there is no mounting solution to install the access point below 26 feet (8 meters), or where ceilings are replaced by the stars and the sky (outdoors), or if directional coverage is needed it is recommend to use directional antennas. When selecting a directional antenna, you should compare the horizontal/vertical beam-width and gain of the antenna.
When using directional antennas on a ceiling mounted access point, direct the antenna pointing straight down. When using directional antennas on a wall mounted access point, tilt the antenna at an angle to the ground. Further tilting a wall mounted antenna to pointing straight down will limit its range.
Cisco Meraki offers 6 types of indoor-rated external antennas (available for MR42E and MR53E):
C/D/E/F series antennas will be automatically detected by the AP. Once an antenna is detected by the AP it cannot be changed in dashboard until the antenna is removed and AP is rebooted.
Cisco Meraki offers 4 types of outdoor external antennas and supports 5 types of outdoor antennas. Cisco Meraki has certified the antennas for use with the Meraki MR84, MR74, MR72, MR66, and MR62 access points. AIR-ANT2514-P4M can only be used with MR84:
Using 3rd party antennas with gain higher than 11 dBi on 2.4 GHz or 13 dBi on 5 GHz may violate regulations in some countries. Meraki certifies only Meraki antennas.
Access Point Placement
Once the number of access points has been established, the physical placement of the AP’s can then take place. A site survey should be performed not only to ensure adequate signal coverage in all areas but to additionally assure proper spacing of APs onto the floorplan with minimal co-channel interference and proper cell overlap. It’s very important to consider the RF environment and construction materials used for AP placement.
Review the designs below from the Cisco Meraki San Francisco office. The 4th Floor was constructed to support Cisco's sales team, customer briefings, and a cafe. In contrast, the 3rd floor was constructed to support Cisco's 24x7 technical support, our small IT department, and Cisco's Collaboration group with applications such as Telepresence and Cisco Spark HD video chat. The density of the 3rd floor is double that of the 4th floor.
High density with 30 access points, Cisco San Francisco, 4th Floor
Ultra High density with 60 access points, Cisco San Francisco, 3rd Floor
SSID Configuration
Making the changes described in this section will provide a significant improvement in overall throughput by following the best practices for configuring SSIDs, IP assignment, Radio Settings, and traffic shaping rules.
Number of SSIDs
The maximum recommended number of SSIDs is 3, and in a high-density environment, this recommendation becomes a requirement. If needed, the number of SSIDs can be increased to 5 but should be done only when necessary. Using more than 5 SSIDs creates substantial airtime overhead from management frames: consuming 20% or more of the bandwidth available and limiting the maximum throughput to less than 80% of the planned capacity. Create a separate SSID for each type of authentication required (Splash, PSK, EAP) and consolidate any SSIDs that use the same type authentication.
Adding several SSIDs has a negative impact on capacity and performance. See the article Multi-SSID Deployment Considerations for more detail.
Enable Bridge Mode
Bridge mode is recommended to improve roaming for voice over IP clients with seamless Layer 2 roaming. In bridge mode, the Meraki APs act as bridges, allowing wireless clients to obtain their IP addresses from an upstream DHCP server. Bridge mode works well in most circumstances, provides seamless roaming with the fastest transitions. When using Bridge mode, all APs in the intended area (usually a floor or set of APs in an RF Profile) should support the same VLAN to allow devices to roam seamlessly between access points.
For seamless roaming in bridge mode, the wired network should be designed to provide a single wireless VLAN across a floor plan. If the network requires a user to roam between different subnets, using L3 roaming is recommended. Bridge mode will require a DHCP request when roaming between two subnets or VLANs. During this time, real-time video and voice calls will noticeably drop or pause, providing a degraded user experience.
NAT mode is not recommended for Voice over IP: With NAT mode enabled, devices will request a new DHCP IP address on each roam. Moving between APs in NAT mode will cause the connection to break when moving AP to AP. Applications requiring continuous traffic streams such as VoIP, VPN or media streams will be disrupted during roaming between APs.
Layer 3 Roaming
Distribute Layer 3 Roaming (D3LR) and MX as a Concentrator are no longer recommended solutions for large scale roaming.
Large wireless networks that require fast and seamles roaming across multiple AP VLANs and subnets will need a centralized wireless deployment with to enable application and session persistence while mobile clients roams at scale.
Within Meraki dashboard, Cisco Campus Gateway is the solution to centralize wireless traffic from Cisco cloud-managed access points and is a cloud-native solution built for large campus wireless deployments. The Campus Gateway tunnels SSID traffic from APs over Virtual eXtensible Local Area Network (VXLAN) to the Campus Gateway, which then bridges the traffic to the upstream network. A single Campus Gateway can support up to 5,000 access points and 50,000 clients with up to 100 Gbps of throughput, and a cluster of two Campus Gateway s doubles that to 200 Gbps.
Please refer to the Campus Gateway Deployment Guide.
Radio Settings & Auto RF
Cisco Meraki access points feature a third radio dedicated to continuously and automatically monitoring the surrounding RF environment to maximize Wi-Fi performance even in the highest density deployment. By measuring channel utilization, signal strength, throughput, signals from non-Meraki APs, and non-WiFi interference, Cisco Meraki APs automatically optimize the radio transmit power and selected operating channels of individual APs to maximize system-wide capacity.
Additionally, it is recommend to use RF profiles to better tune the wireless network to support the performance requirements. A separate RF profile should be created for each area that needs unique set of RF settings. The following details can be set in the RF Profiles:
Band Selection
