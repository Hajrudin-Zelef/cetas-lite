---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1-1
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1"
domain: general-networking
role: reference
task: reference
actors: ["China"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1.md
source_anchor: ""
source_lines: [1, 26]
sha256: 730930388996e58580b47860fcab8400b8e623c207c4fb5b096c47cdd1c63747
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1

Meraki Auto VPN - Configuration and Troubleshooting
Click 日本語 for Japanese
Overview
Auto VPN is a proprietary technology developed by Meraki that allows you to quickly and easily build VPN tunnels between Meraki WAN Appliances at your separate network branches with just a few clicks. Auto VPN performs the work normally required for manual VPN configurations with a simple cloud based process. This article outlines how the Auto VPN mechanisms work and how Meraki manages the cloud processes for Auto VPN.
Definitions
- VPN Registry: This is the main server mechanism that allows Auto VPN to happen. It is a cloud service that is used to keep track of the contact information for all the WAN Appliance participating in Auto VPN for an organization.
WAN Appliances in warm spare with Virtual IP address (VIP) will use VIP to communicate with the VPN registry.
- Hub: Hubs are devices in a VPN topology that service connectivity from a remote peer site (such as a spoke) to the hub and the hub to the remote peer site. Hubs also act as a gateway for remote peer sites to communicate with each other via the hub.
- Spoke: Remote sites that connect to a central hub and communicate with each other only through the hub.
- Peer: This refers to another WAN Appliance within the same organization that a local WAN Appliance will form or has formed a VPN tunnel to.
- Contact: This is the public IP and the UDP port that the WAN Appliance will communicate on for Auto VPN.
How Auto VPN Works
- MX1 and MX2 are part of the same organization. MX1 and MX2 are configured to participate in Auto VPN. Both MX1 and MX2 send a Register Request message to their VPN registry in order to share their own contact information, and to get the contact information of the peer WAN Appliance(s) that it should form a VPN tunnel with. The Register Request message contains the IP address and the UDP port that the WAN Appliance communicates on, and the WAN Appliance requests the contact information of its peer WAN Appliance(s).
- VPN registries send the Register Response messages to the WAN Appliances with the contact information of the peers the WAN Appliances should establish a tunnel with.
- Once the information is shared with the WAN Appliance about its peers, a VPN tunnel is formed WAN Appliance to WAN Appliance. The Meraki cloud already knows the subnet information for each WAN Appliance, and now the IP addresses to use for tunnel creation. The cloud pushes a key to the WAN Appliances in their configuration which is used to establish an AES encrypted IPsec-like tunnel. Local subnets specified by dashboard admins are exported/shared across VPN. During this process, VPN routes are pushed from the dashboard to the WAN Appliances. Finally, the dashboard will dynamically push VPN peer information (e.g., exported subnets, tunnel IP information) to each WAN Appliance. Every WAN Appliance stores this information in a separate routing table.
Any devices sitting upstream of a WAN Appliance will need the following destinations whitelisted so the WAN Appliance can communicate with the Auto VPN registries:
- Port
    
  - UDP 9350-9381
- IP range for non-China cloud (meraki.com):
    
  - 209.206.48.0/20
  - 158.115.128.0/19
  - 216.157.128.0/20
- IP range for China cloud (meraki.cn):
    
