---
id: collect-261001-huawei/huawei/sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9-1
title: "sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "ethernet", "license", "parameters"]
source: docs/RAG/collect-261001-huawei/sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9.md
source_anchor: ""
source_lines: [1, 58]
sha256: 937f001e15cc928c291878d281926dcd3b3b55acc04bf03b718c1eaf8faee244
---

# sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9

VPN Concentrator Deployment Guide
Overview
In order to connect AutoVPN sites to a central location, such as a datacenter, MX WAN appliances can be deployed to serve as a VPN concentrator. This guide outlines the configuration and deployment steps necessary for setup.
Key Concepts
Before deploying a one-armed VPN concentrator, it is important to understand several key concepts.
Operating Mode
All MXs can be configured in either Routed or VPN concentrator mode. There are important considerations for both modes. More detailed information on concentrator modes, click here.
One-Armed Concentrator
This configuration utilizes an MX device configured to act in VPN concentrator mode, with a single Ethernet connection to the upstream network. All traffic will be sent and received on this interface. This is the recommended configuration for MX appliances serving as VPN termination points into the datacenter.
MX devices configured as VPN Concentrators must utilize WAN 1 as the single connection to the upstream network. Behavior of additional ports beyond that single WAN port is described further here.
Routed Mode Concentrator
In this mode the MX is configured with a single Ethernet connection to the upstream network and one Ethernet connection to the downstream network. VPN traffic is received and sent on the WAN interfaces connecting the MX to the upstream network and the decrypted, unencapsulated traffic is sent and received on the LAN interface that connects the MX to the downstream network.
Warm Spare (High Availability) for VPN concentrators
When configured for high availability (HA), one MX serves as the primary unit and the other MX operates in a spare mode. All traffic flows through the primary MX, while the spare operates as an added layer of redundancy in the event of failure.
Failover between MXs in an HA configuration leverages VRRP heartbeat packets. These heartbeat packets are sent from the Primary MX to the Spare MX via the singular uplink for MXs operating in VPN concentrator mode in order to indicate that the Primary is online and functioning properly. As long as the Spare is receiving these heartbeat packets, it functions in the passive state. If the Spare stops receiving these heartbeat packets, it will assume that the Primary is offline and will transition into the active state. In order to receive heartbeats in a one-armed concentrator configuration, both VPN concentrator MXs should have uplinks on the same subnet within the datacenter.
For Routed mode configurations, both concentrators must be able to communicate using the LAN ports. More information on Routed mode warm spare can be found here.
Only one MX license is required for the HA pair, as only a single device is in full operation at any given time.
Connection Monitor
Connection monitor is an uplink monitoring engine built into every MX WAN appliance. The mechanics of the engine are described in this article.
SSID Tunneling to an MX VPN Concentrator
The MX WAN appliance is the ideal solution for SSID Tunneling using VPN concentration as it is custom built for mission critical networks. Choose the MX WAN appliance that is best fit for your needs based on the Sizing Guide.
The WAN appliance is ready to concentrate SSIDs out of the box without any additional configuration beyond what is outlined in the quick start guide.
For additional information on how to set this up, please refer to this section.
To increase reliability, a second MX WAN appliance can be paired in HA mode. In the case that the primary WAN appliance becomes unreachable from the Meraki Cloud, the Access Points will failover to the HA standby WAN appliance.
Note: Virtual MX (vMX) devices are not supported as Wireless concentrators. As a result, Security & SD-WAN > Configure > Wireless Concentrator is not available on vMX devices.
Deploying a One-Armed Concentrator
Interface Configuration
The MX WAN appliance being configured as a one-armed VPN concentrator should be connected to the upstream datacenter infrastructure using its Internet port, or using the Internet 1 port on devices models with two Internet uplink ports.
When using the MX as a one-armed VPN concentrator for VPN endpoints, be sure to not connect anything to the MX's LAN ports. If the MX is simply being used as a passthrough device, using its LAN ports will not impact its performance.
MX IP Assignment
In the datacenter, an MX WAN appliance can operate using a static IP address or an address from DHCP. MX appliances will attempt to pull DHCP addresses by default. It is highly recommended to assign static IP addresses to VPN concentrators.
Static IP assignment can be configured via the device local status page.
The local status page can also be used to configure VLAN tagging on the uplink of the MX. It is important to take note of the following scenarios:
- If the upstream port is configured as an access port, VLAN tagging should not be enabled.
- If the port upstream is configured as a trunk port and the MX should communicate on the native or default VLAN, VLAN tagging should be left as disabled.
- If the port upstream is configured as a trunk and the MX should communicate on a VLAN other than the native or default VLAN, VLAN tagging should be configured for the appropriate VLAN ID.
Public IP assignment
Placing an MX appliance configured as a one-armed VPN concentrator at the perimeter of the network with a publicly routable IP address is not recommended and can present security risks. As a best practice, one-armed concentrators MX appliances should always be deployed behind an edge firewall that filters inbound connections.
Dashboard Configuration
The Cisco Meraki Dashboard configuration can be done either before or after bringing the unit online.
- 
    Begin by configuring the MX to operate in VPN Concentrator mode. This setting is found on the Security & SD-WAN > Configure > Addressing & VLANs Page. The MX will be set to operate in Routed mode by default. 
- 
    Next, configure the Site-to-Site VPN parameters. This setting is found on the Security & SD-WAN > Configure > Site-to-site VPN page.
- Begin by setting the type to "Hub (Mesh)."
- Configure the local networks that are accessible upstream of this VPN concentrator.
    
  - 
        For the Name, specify a descriptive title for the subnet.
  - 
        For the Subnet, specify the subnet to be advertised to other AutoVPN peers using CIDR notation
- 
        
- 
    NAT traversal can be set to either Automatic or Manual: Port forwarding. See below for more details on these two options.
- 
    An example screenshot is included below: 
 
