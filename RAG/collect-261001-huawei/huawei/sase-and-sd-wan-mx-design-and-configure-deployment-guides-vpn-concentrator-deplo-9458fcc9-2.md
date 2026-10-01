---
id: collect-261001-huawei/huawei/sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9-2
title: "sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9"
domain: huawei
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["cost", "datacenter", "parameters"]
source: docs/RAG/collect-261001-huawei/sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9.md
source_anchor: ""
source_lines: [59, 145]
sha256: 106bb5cf92ff69d43a1981a2e847e552174fec357d9f48f35f7f56cdaa83fe39
---

# sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9

NAT traversal
Whether to use Manual or Automatic NAT traversal is an important consideration for the VPN concentrator.
Use manual NAT traversal when:
- There is an unfriendly NAT upstream
- Stringent firewall rules are in place to control what traffic is allowed to ingress or egress the datacenter
- It is important to know which port remote sites will use to communicate with the VPN concentrator
If manual NAT traversal is selected, it is highly recommended that the VPN concentrator be assigned a static IP address. Manual NAT traversal is intended for configurations when all traffic for a specified port can be forward to the VPN concentrator.
Use automatic NAT traversal when:
- None of the conditions listed above that would require manual NAT traversal exist
If automatic NAT traversal is selected, the MX will automatically select a high numbered UDP port to source AutoVPN traffic from. The VPN concentrator will reach out to the remote sites using this port, creating a stateful flow mapping in the upstream firewall that will also allow traffic initiated from the remote side through to the VPN concentrator without the need for a separate inbound firewall rule.
For more information regarding Automatic NAT traversal for AutoVPN, please review this document.
Adding warm spare
This section outlines the steps required to configure and implement a warm spare (HA) for an MX WAN appliance operating in VPN concentrator mode.
Behavior
When configured for high availability (HA), one MX is active, serving as the active, and the other MX operates in a passive, standby capacity. The VRRP protocol is leveraged to achieve failover. Please see here for more information.
Dashboard Configuration
High availability on MX WAN appliances requires a second MX of the same model. The HA implementation is active/passive and will require the second MX also be connected and online for proper functionality.
High availability (also known as warm spare) can be configured from Security & SD-WAN > Monitor > Appliance status. Begin by setting Warm Spare to Enabled. Next, enter the serial number of the warm spare MX. Finally, select whether to use MX uplink IPs or virtual uplink IPs.
Uplink IPs
Use Uplink IPs is selected by default for new network setups. In order to properly communicate in HA, VPN concentrator MXs must be set to use the virtual IP (vIP).
Virtual IP (vIP)
The virtual uplink IPs option uses an additional IP address that is shared by the HA MXs. In this configuration, the MXs will send their cloud controller communications via their uplink IPs, but other traffic will be sent and received by the shared virtual IP address.
Configuring OSPF Route Advertisement
MX WAN appliances acting in VPN concentrator mode support advertising routes to connected VPN subnets via OSPF.
Behavior
An MX VPN concentrator with OSPF route advertisement enabled will only advertise routes via OSPF; it will not learn OSPF routes.
When spoke sites are connected to the VPN concentrator, the routes to spokes sites are advertised using an LS Update message. These routes are advertised as type 2 external routes.
Dashboard Configuration
In order to configure OSPF route advertisement, navigate to the Security & SD-WAN > Configure > Site-to-Site VPN page. From this page:
- Set Advertise remote routes to Enabled
- Configure the Router ID
- Configure the Area ID
- Adjust the Cost, if desired
- Adjust the Hello timer, if needed
- Adjust the Dead timer, if needed
- Enable and configure MD5 authentication, if needed
For additional details, please see Using OSPF to Advertise Remote VPN Subnets.
Other Datacenter Configuration
Upstream Considerations
This section discusses configuration considerations for other components of the datacenter network.
Routing
The MX acting as a VPN concentrator in the datacenter will be terminating remote subnets into the datacenter. In order for bi-directional communication to take place, the upstream network must have routes for the remote subnets that point back to the MX acting as the VPN concentrator.
If OSPF route advertisement is not being used, static routes directing traffic destined for remote VPN subnets to the MX VPN concentrator must be configured in the upstream routing infrastructure.
If OSPF route advertisement is enabled, upstream routers will learn routes to connected VPN subnets dynamically.
Firewall considerations
The MX WAN appliance makes use of several types of outbound communication. Configuration of the upstream firewall may be required to allow this communication.
Dashboard & Cloud
The MX WAN appliance is a cloud managed networking device. As such, it is important to ensure that the necessary firewall policies are in place to allow for monitoring and configuration via the Cisco Meraki Dashboard. The relevant destination ports and IP addresses can be found under the Help > Firewall info page in the Dashboard.
VPN Registry
Cisco Meraki's AutoVPN technology leverages a cloud-based registry service to orchestrate VPN connectivity. In order for successful AutoVPN connections to establish, the upstream firewall must allow the VPN concentrator to communicate with the VPN registry service. The relevant destination ports and IP addresses can be found under the Help > Firewall info page in the Dashboard.
Uplink Health Monitoring
The MX also performs periodic uplink health checks by reaching out to well-known Internet destinations using common protocols. The full behavior is outlined here. In order to allow for proper uplink monitoring, the following communications must also be allowed:
- 
    ICMP to 8.8.8.8 (Google's public DNS service)
- 
    HTTP port 80
- 
    DNS to the MX's configured DNS server(s)
Deploying a Routed mode concentrator
Example Topology
An MX VPN concentrator can also be configured to operate in Routed mode. The following diagram shows an example of a datacenter topology with a Routed mode concentrator:
Interface Configuration
The MX WAN appliance being configured as a VPN concentrator should be connected to the "upstream" datacenter infrastructure closer to the network edge using its Internet port, and connected to "downstream" infrastructure closer to the datacenter services using a LAN port.
MX IP Assignment
In the datacenter, an MX WAN appliance can operate using a static IP address or an address from DHCP. MX appliances will attempt to pull DHCP addresses by default. It is highly recommended to assign static IP addresses to VPN concentrators.
Static IP assignment can be configured via the device local status page.
The local status page can also be used to configure VLAN tagging on the uplink of the MX. It is important to take note of the following scenarios:
- 
    If the upstream port is configured as an access port, VLAN tagging should not be enabled.
- 
    If the port upstream is configured as a trunk port and the MX should communicate on the native or default VLAN, VLAN tagging should be left as disabled.
- 
    If the port upstream is configured as a trunk and the MX should communicate on a VLAN other than the native or default VLAN, VLAN tagging should be configured for the appropriate VLAN ID.
Public IP assignment
An MX appliance configured as a Routed mode concentrator can be configured with either a publicly routable IP address or be deployed behind another NAT device within the datacenter topology.
Dashboard Configuration
The Cisco Meraki Dashboard configuration can be done either before or after bringing the unit online.
The following configuration steps will be covered in more detail in the sections below:
- 
    Configure the MX to operate in Routed mode.
- 
    Define VLANs and static routes.
- 
    Configure the Site-to-Site VPN parameters.
Operating Mode Configuration
Begin by configuring the MX to operate in Routed mode. This setting is found on the Security & SD-WAN > Configure > Addressing & VLANs page. 
 
