---
id: collect-261001-general-networking/general-networking/mx-site-to-site-vpn-configuring-hub-and-spoke-vpn-connections-on-the-mx-security-35f36c1d
title: "mx-site-to-site-vpn-configuring-hub-and-spoke-vpn-connections-on-the-mx-security-35f36c1d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-general-networking/mx-site-to-site-vpn-configuring-hub-and-spoke-vpn-connections-on-the-mx-security-35f36c1d.md
source_anchor: ""
source_lines: [1, 35]
sha256: bee91118cf660d57aaa9c34ca347efab9dfd47f8cccc28068158459c6500d212
---

# mx-site-to-site-vpn-configuring-hub-and-spoke-vpn-connections-on-the-mx-security-35f36c1d

Configuring Hub-and-spoke VPN Connections on the MX Security Appliance
Site-to-site VPN connections between MX Security Appliances and/or Z-series Teleworker Gateways will automatically form a mesh topology between all VPN-enabled peers in the same Dashboard organization by default. This is often undesirable because such connections establish unnecessary IPSec tunnels between remote sites and create performance-degrading networking overhead.
In these cases it is best to configure Site-to-site VPN topology for Hub and spoke, which designates the datacenter MX as the "hub" and all remote sites as the "spoke". This model can be useful in organizations where several auxiliary sites require a connection to the HQ or datacenter-located concentrator, pictured below.
Figure 1. Split tunnel w/ Hub-and-Spoke (connect directly to one peer). VPN connections (blue) are established to only one peer (top). Traffic to the internet (black) goes out locally from each site.
Figure 2. Full tunnel w/ Hub-and-Spoke (connect directly to one peer). VPN connections (blue) are established to only one peer (top). Traffic to the internet (black) goes out from a central concentrator/hub (top).
Although each remote location is not connected directly in this method, remote sites can still connect with each other via the hub by default. This article covers:
- Hub and spoke VPN setup and configuration
- Limiting connections between Spokes
Hub and Spoke VPN Setup and Configuration
Note: Hub and spoke topologies are currently only supported between Meraki MXes, non-Meraki VPN peers cannot be configured as spokes.
The MX features a hub-and-spoke option for its MX to MX VPN. To implement Hub-and-spoke the network administrator needs to follow these steps:
- Set up the hub MX Device
    
  - Navigate to the Dashboard Network of the MX that will act as the hub.
  - Navigate to Security & SD-WAN > Configure > Site-to-Site VPN.
  - Set the Type to Hub (Mesh):
  - (Optional) If another MX in the organization is also configured as a hub, it can be added as an Exit hub. If configured, all VPN client traffic to this MX will be tunneled to the specified exit hub.
  - Configure any other VPN settings desired (local networks, NAT traversal, etc)
  - Save.
- Set up the spoke MX Device
    
  - Navigate to the Dashboard Network of the MX that will act as the spoke.
  - Navigate to Security & SD-WAN > Configure > Site-to-site VPN.
  - Set the Type to Spoke:
  - Select the hub MX under the Name drop-down. Multiple hubs can be added and prioritized in descending order.
  - Select at least one hub for a Default route:
        
    - If a hub is not configured as a default route, the spoke will only send traffic to this hub when the destination subnet is advertised by the hub.
    - If a hub is configured as a default route, any traffic that is not destined for a higher-priority hub will be sent by default to this hub.
  - Configure any other VPN settings desired (local networks, NAT traversal, etc)
  - Save.
Once Saved, the MX set as "Spoke" will form a VPN tunnel with the specified hub(s).
Limiting Connections Between Spokes
In the event you need to limit branch office communication, configure Site-to-site firewall rules. You can edit the Site-to-site firewall from any MX network. This can be done at Security & SD-WAN > Configure > Site-to-site VPN > Organization Wide Settings > Site-to-site outbound firewall.
This will allow you to limit the communication between spokes in the way you desire. In this particular example we have prevented the New York site from communicating with the Chicago site and vice versa, while allowing both sites to communicate with the hub via the inherent allow any rule at the bottom. Rules are processed in a top-down fashion. If you want to deny the traffic from each spoke, you must set up a rule both ways.
