---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-115003173168-unifi-gateway-zone-based-firewall-f5daf8f3-1
title: "hc-en-us-articles-115003173168-unifi-gateway-zone-based-firewall-f5daf8f3"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-115003173168-unifi-gateway-zone-based-firewall-f5daf8f3.md
source_anchor: ""
source_lines: [1, 74]
sha256: 73e220a0fda1c5e1a897bfd80d47158c5259ccbfe15f09f9d76be039171252bd
---

# hc-en-us-articles-115003173168-unifi-gateway-zone-based-firewall-f5daf8f3

Zone-Based Firewalls in UniFi
UniFi's Zone-Based Firewalling (ZBF) simplifies firewall management by allowing you to group network interfaces—such as VLANs, WANs, or VPNs—into zones. This approach lets you efficiently define and enforce policies that control how traffic flows between these zones, making it easy to manage network security and segmentation.
For a full overview of UniFi’s Traffic and Policy Management capabilities, see here.
For a full overview of UniFi's Network and Cyber Security capabilities, see here.
| This feature is part of UniFi Network 9.0.108 Official Release.  | 
Requirements
- UniFi Cloud Gateway (or independent UniFi Gateway)
- UniFi Network Application version 9.0 or newer
- UniFi (Cloud) Gateway version 4.1 or newer
What are Firewall Zones?
Firewall zones are logical groupings of network interfaces, such as VLANs, WANs, or VPNs. By applying policies to these zones, you can define and control traffic flow with ease, eliminating the need to create individual policies for each interface. Each zone can represent different segments of your network, such as trusted, semi-trusted, or untrusted areas, enhancing both security and simplicity.
For a simpler guide to implementing network and client isolation, go here.
Advantages of Zone-Based Firewalls
- Simplified Policy Management: Policies are created between zones, reducing complexity and improving clarity compared to managing policies at the interface level.
- Granular Control Over Traffic: Define precise policies based on IP addresses, protocols, applications, or users, ensuring comprehensive traffic management.
- Enhanced Network Segmentation: Establish clear boundaries between zones to protect sensitive areas, such as limiting how traffic moves from an external WAN zone into your internal network.
- Better Visibility: Policies are shown visually in the Zone Matrix, providing greater insight and ease of management.
Built-in Firewall Zones
The UniFi firewall includes several predefined, built-in zones to which networks and interfaces are associated.
- External: For incoming traffic that is untrusted, or requires more strict control, such as general Internet traffic on the WAN, or a connection with a third-party VPN client service.
- Internal: For trusted traffic, such as employee computers and internal servers on the local network.
- Gateway: Handles traffic directed to or from the UniFi Gateway (such as DHCP, DNS, or HTTPS/SSH management requests).
- VPN: For traffic from remote VPN users (Identity One-Click VPN, WireGuard, L2TP, and OpenVPN), or Site-to-Site VPNs (Site Magic, IPsec, and OpenVPN).
- Hotspot: For guest WiFi hotspot networks where devices have restricted access.
- DMZ: For deployments which require outside access to public-facing resources, such as web or mail servers.
Creating and Modifying Zones
Predefined zones are marked with a lock icon to indicate they cannot be removed. However, admins can create custom zones for specialized traffic or more precise control. Network interfaces are limited to a single zone and are initially assigned to a predefined zone by default, but this assignment can be modified in the Firewall section. Note that there is a maximum of 30 zones.
The Zone Matrix: Viewing Traffic Segmentation Between Zones
The zone matrix provides a clear, visual representation of traffic flow between zones, displaying a grid of built-in and custom policies. Rows represent source zones (where traffic originates), and columns represent destination zones (where traffic is headed). The intersections, or cells, show and allow configuration of policies controlling traffic between zones. For example, clicking on the intersection between “Internal” and “External” zones lets you view or adjust specific firewall policies governing that traffic flow, streamlining policy management and enhancing network visibility.
| Built-in Zones |  | Destination Zone |  |  |  |  |  | 
|  |  | Internal | External | Gateway | VPN | Hotspot | DMZ | 
| Source Zone | Internal | Allow All | Policies | Allow All | Allow All | Allow All | Allow All | 
|  | External | Policies | Policies | Policies | Policies | Policies | Policies | 
|  | Gateway | Allow All | Allow All | - | Allow All | Allow All | Allow All | 
|  | VPN | Allow All | Policies | Allow All | Allow All | Allow All | Allow All | 
|  | Hotspot | Allow Return Traffic | Policies | Policies | Allow Return Traffic | Block All | Block All | 
|  | DMZ | Allow Return Traffic | Policies | Policies | Allow Return Traffic | Block All | Block All | 
The following values are shown in the matrix:
- Allow All - All traffic is allowed from the source zone to the destination zone
- Block All - All traffic is blocked from the source zone to the destination zone
- Allow Return Traffic - This value appears when there is a combination of "Allow All" and "Block All" between two zones. The source zone is allowed to send all traffic to the destination zone, but the destination zone can only reply to the traffic.
- Policies - Specific traffic is allowed and blocked from the source zone to the destination zone, controlled via multiple firewall policies. By default, this applies to built-in policies associated with the External zone which is used for traffic coming and going to the internet.
Traffic Directions and Traffic Inside Zones
With zones, limiting of traffic is done in both directions. This means that if traffic is blocked from source "Zone A" to destination "Zone B" but allowed from "Zone B" to "Zone A", the end result is that traffic is still blocked in one direction. Carefully consider both directions of the traffic when creating firewall policies.
In addition to filtering traffic between different zones, it is also possible to filter within the same zone, for example Internal to Internal. This is useful when there are multiple networks assigned to a zone, but traffic needs to be filtered between them.
Assigning Networks to Zones
Networks can only be assigned to a single zone and are placed in one of the built-in zones by default. Upon creation or when editing the network, it is possible to place it in a different zone. This can also be done by editing the zone configuration in the Firewall section.
Configuring Firewall Policies
Firewall policies control the flow of traffic between zones, letting you allow or block specific types of traffic. Follow these steps to set up and customize a firewall policy:
- 
Navigate to Firewall Rules: Follow the path depending on your UniFi Network version:
  - Network 9.4: Settings > Zones > Create Policy or Settings > Policy Table > Create New Policy
  - Network 9.3: Settings > Policy Engine > Zones > Create Policy
- 
Configure Source and Destination Zones: Specify the rule's scope by selecting the source and destination zones. Optionally, refine criteria for matching traffic using:
  - Any, Device, Network, IP or MAC
  - Port (Any, Specific or Object)
  - App, Domain ("Web") or Region
- 
Select Your Action: Choose how the policy will handle matching traffic:
  - 
Allow: Permit the traffic.
    - Auto Allow Return Traffic: Creates an additional built-in firewall policy to allow the return traffic from the destination to the source zone. This is not required if return traffic is already allowed via another policy.
  - Block: Silently drop traffic.
  - Reject: Block traffic and notify the sender.
- 
Allow: Permit the traffic.
- 
Specify Restrictions (Optional): Customize the policy further by selecting:
  - IP Version: Match IPv4, IPv6, or both.
  - Protocol: Target TCP, UDP, or other protocols like ICMP.
  - Connection State: Match established, invalid, or new connections.
- Enable Syslog Logging (Optional): Send traffic flow data to a remote SIEM server by enabling syslog logging. Configure your SIEM server in the Integrations section.
- Set a Custom Schedule (Optional): Define when the policy will be active, such as during work hours or weekends.
