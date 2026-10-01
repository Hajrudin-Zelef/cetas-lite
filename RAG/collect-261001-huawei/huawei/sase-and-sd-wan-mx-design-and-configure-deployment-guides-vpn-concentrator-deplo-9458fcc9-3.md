---
id: collect-261001-huawei/huawei/sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9-3
title: "sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-huawei/sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9.md
source_anchor: ""
source_lines: [146, 207]
sha256: a39a19d68b152aba798c0254a4d9fc4140d8685097cc756a5eb50498470cdbca
---

# sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9

VLAN and Static Route Configuration
Both VLANs and Static routes can be configured from the Addressing & VLANs page. Both Static routes and VLANs can be advertised into the AutoVPN topology. However, VLANs configured on a Routed mode MX must be unique to each Routed mode MX within the AutoVPN topology.
For a Routed mode concentrator, it is recommended to configure a VLAN with a small subnet for communication between the MX and other downstream infrastructure. Static routes are then used to provide access to other datacenter services downstream.
Defining a VLAN
Begin by navigating to the Security & SD-WAN > Configure > Addressing & VLANs page to define a subnet to be used for communication with other downstream routers.
If VLAN-specific configuration is required for downstream communication out the MX's LAN port, such as tagging traffic with a specific VLAN ID, VLANs must be enabled. First, enable VLANs. Under the Routing heading, check the Use VLANs box to enable VLANs. This allows a VLAN ID to be configured for subnets defined in the Subnets table.
Then, click the Default subnet within the Subnets table. This will bring up the Modify VLAN configuration menu.
From the VLAN configuration, define the VLAN Name, VLAN ID, Group Policy, VPN mode (if Site-to-site VPN is enabled), IPv4 VLAN interface IP and Subnet, and IPv6 config (if enabled).
The Configure Single LAN configuration menu will be presented if VLANs are disabled. The Modify VLAN configuration menu will be presented if VLANs are enabled. VLAN ID is only configurable from the Modify VLAN configuration menu.
Per-port VLAN Configuration
If VLANs are set to enabled from the Addressing & VLANs page and a VLAN has been defined for communication between the WAN appliance acting as a Routed mode VPN concentrator and downstream routers, it is important to set the LAN port's VLAN configuration correctly for proper bi-directional communication.
In the Per-port VLAN Settings table, click on the LAN port connecting the WAN appliance to the downstream infrastructure to bring up the Configure MX LAN ports menu. From here, set Enabled, Type, Native VLAN, and Allowed VLANs.
Defining Static Routes
To define a static route, begin by navigating to the Security & SD- WAN > Configure > Addressing & VLANs page.
Click on the Add Static Route link in the Static Routes table to open the Add Static Route configuration menu.
In the Add Static Route configuration menu, define the Name, IP version, Subnet, Next hop IP, Active state, and the VPN mode status.
The VPN mode configuration option on the static route configuration menu will only appear if VPN has already been enabled on the Security & SD-WAN > Configure > Site-to-site VPN page.
Static routes can also be configured to be allowed in the VPN from the Site-to-site VPN page.
Multiple static routes may be configured. An example is included below:
Static routes that are allowed in VPN will always be advertised into AutoVPN. Static routes configured as active While next hop responds to ping and While host responds to ping will be advertised AutoVPN, independent of whether the static route's active condition is met.
Please see here for more information on configuring static routes on Routed mode WAN appliances.
Configure Site-to-site VPN
Site-to-site VPN configuration settings are managed from the Security & SD-WAN > Configure > Site-to-site VPN page. From the site-to-site VPN page, begin by setting the type to "Hub (Mesh)." In the Local networks table, for each subnet that needs to be accessible over VPN, set VPN mode to "Enabled". NAT traversal can be set to either Automatic or Manual: Port forwarding. See below for more details on these two options. An example screenshot is included below:
NAT traversal
Whether to use Manual or Automatic NAT traversal is an important consideration for the VPN concentrator.
Use manual NAT traversal when:
- 
    There is an unfriendly NAT upstream
- 
    Stringent firewall rules are in place to control what traffic is allowed to ingress or egress the datacenter
- 
    It is important to know which port remote sites will use to communicate with the VPN concentrator
If manual NAT traversal is selected, it is highly recommended that the VPN concentrator be assigned a static IP address. Manual NAT traversal is intended for configurations when all traffic for a specified port can be forward to the VPN concentrator.
Use automatic NAT traversal when:
- 
    None of the conditions listed above that would require manual NAT traversal exist
If automatic NAT traversal is selected, the MX will automatically select a high numbered UDP port to source AutoVPN traffic from. The VPN concentrator will reach out to the remote sites using this port, creating a stateful flow mapping in the upstream firewall that will also allow traffic initiated from the remote side through to the VPN concentrator without the need for a separate inbound firewall rule.
Adding warm spare
This section outlines the steps required to configure and implement warm spare (HA) for an MX WAN appliance operating in Routed mode.
Topology
The following is an example of a topology that leverages an HA configuration for VPN concentrators:
Behavior
When configured for high availability (HA), one MX is active, serving as the active, and the other MX operates in a passive, standby capacity. The VRRP protocol is leveraged to achieve failover. Please see here for more information.
Dashboard Configuration
High availability on MX WAN appliances requires a second MX of the same model. The HA implementation is active/passive and will require the second MX also be connected and online for proper functionality.
High availability (also known as warm spare) can be configured from Security & SD-WAN > Monitor > Appliance status. Begin by setting Warm Spare to Enabled. Next, enter the serial number of the warm spare MX. Finally, select whether to use MX uplink IPs or virtual uplink IPs.
Uplink IPs
Use Uplink IPs is selected by default for new network setups. In order to properly communicate in HA, VPN concentrator MXs must be set to use the virtual IP (vIP).
Virtual IP (vIP)
The virtual uplink IPs option uses an additional IP address that is shared by the HA MXs. In this configuration, the MXs will send their cloud controller communications via their uplink IPs, but other traffic will be sent and received by the shared virtual IP address.
The functionality discussed here is currently only available in beta. To get access to the beta, please contact Meraki Support.
Other Datacenter Configuration
Routing Considerations
The MX acting as a VPN concentrator in the datacenter will be terminating remote subnets into the datacenter. In order for bi-directional communication to take place, the downstream network must have routes for the remote AutoVPN subnets that point back to the MX acting as the VPN concentrator.
In order for traffic received on the LAN side of a Routed mode concentrator to be passed over AutoVPN, traffic must both be sourced from a subnet matching a local VLAN or static route defined on the Addressing & VLANs page of the concentrator and that subnet must be allowed in VPN. If either condition is not met, traffic will not be routed by the MX from the LAN over AutoVPN.
Upstream Considerations
This section discusses configuration considerations for other components of the datacenter network.
Firewall considerations
The MX WAN appliance makes use of several types of outbound communication. Configuration of the upstream firewall may be required to allow this communication.
Dashboard & Cloud
The MX WAN appliance is a cloud managed networking device. As such, it is important to ensure that the necessary firewall policies are in place to allow for monitoring and configuration via the Cisco Meraki Dashboard. The relevant destination ports and IP addresses can be found under the Help > Firewall info page in the Dashboard.
VPN Registry
