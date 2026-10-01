---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-94f1c4e2-4
title: "platform-management-dashboard-administration-design-and-configure-architectures--94f1c4e2"
domain: general-networking
role: reference
task: reference
actors: ["China"]
dates: ["2018-02"]
keywords: ["cost", "datacenter"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--94f1c4e2.md
source_anchor: ""
source_lines: [162, 206]
sha256: 5dc56a8db2cfde6ac2223625bd8bab9fe120ded5231f6e74a1cbfc0f7a1cff5a
---

# platform-management-dashboard-administration-design-and-configure-architectures--94f1c4e2

In a DC-DC failover design, a remote site will form VPN tunnels to all configured VPN hubs for the network. For subnets that are unique to a particular hub, traffic will be routed directly to that hub. For subnets that are advertised from multiple hubs, spoke sites will send traffic to the highest priority hub that is reachable.
When a WAN Appliance is configured to connect to multiple VPN concentrators advertising the same subnets, the routes to those subnets become tracked. Hello messages are periodically sent across the tunnels from the remote site to the VPN hubs to monitor connectivity. If the tunnel to the highest priority hub goes down, the route is removed from the route table and traffic is routed to the next highest priority hub that is reachable. This route failover operation only applies when identical routes are advertised from multiple Auto VPN hubs.
Concentrator Priority
When multiple VPN hubs are configured for an organization, the concentrator priority can be configured for the organization. This concentrator priority setting determines the order in which VPN mesh peers will prefer to connect to subnets advertised by multiple VPN concentrators.
This setting does not apply to remote sites configured as VPN spokes.
Other Datacenter Considerations
When implementing an DC-DC architecture, a warm spare concentrator configuration (see warm spare section above) and OSPF route advertisement should always be taken into consideration for WAN Appliances acting as VPN concentrators in a datacenter. Additionally, route flow logic should be considered for all applications in the deployment environment to ensure availability requirements are met. To assist with better understanding DC-initiated flows, please refer below.
Supported VPN Architectures
VPN Topologies
There are several options available for the structure of the VPN deployment.
Split Tunnel
In this configuration, branches will only send traffic across the VPN if it is destined for a specific subnet that is being advertised by another WAN Appliance in the same dashboard organization. The remaining traffic will be checked against other available routes, such as static LAN routes and third-party VPN routes, and if not matched will be NATed and sent out from the branch WAN Appliance unencrypted.
Full Tunnel
In full tunnel mode all traffic that the branch or remote office does not have another route to is sent to a VPN hub.
Hub and Spoke
In a hub and spoke configuration, the WAN appliances at the branches and remote offices connect directly to specific WAN Appliances and will not form tunnels to other WAN Appliance or Teleworker devices in the organization. Communication between branch sites or remote offices is available through the configured VPN hubs. This is the recommended VPN topology for most SD-WAN deployments.
- 
    Hub and Spoke - Total Tunnel Count 
 
Where H is the number of hubs, S is the number of spokes and L is the number of uplinks the WAN Appliance has (LH for the hubs, LS for the spokes). If each WAN Appliance has a different number of uplinks then a sum series, as opposed to a multiplication will be required.
For example, if all WAN Appliances have 2 uplinks, we have 4 hubs and 100 spokes, then the total number of VPN tunnels would be 24 + 1600 = 1624.
Standard Hub & Spoke
How it works:
Utilizing the standard Meraki Auto VPN registry to ascertain how the VPN tunnels configured need to form (i.e. via public address space or via private interface address space) as described in Configuring Site-to-site VPN over MPLS.
When should it be used?
Whenever possible. It is strongly recommended that this model is the 1st, 2nd and 3rd option when designing a network. Only if the deployment has an exceptionally strong requirement should one of the following hub and spoke derivatives be considered.
VPN Mesh
It is also possible to use a VPN mesh configuration in an SD-WAN deployment.
In a mesh configuration, a WAN Appliance at the branch or remote office is configured to connect directly to any other WAN Appliances in the organization that are also in mesh mode, as well as any spoke WAN Appliances that are configured to use it as a hub.
- 
    Full Mesh - Total Tunnel Count 
 
Where H is the number of WAN Appliances and L is the number of uplinks each WAN Appliance has.
For example, if all WAN Appliances have 2 uplinks and there are 50 WAN Appliances, then the total number of VPN tunnels would be 4900 and every WAN Appliance would have to be able to support 196 tunnels from the number of VPN peers for a single WAN Appliance (49) multiplied by the number of VPN tunnels between each peer (4), in this case, we would need 50 MX100s as a minimum.
China Auto VPN
With regulatory constraints imposed by the Chinese government, specific architecture requirements are needed to deploy and interconnect the Auto VPN domain within China to the rest of the world.
Secure VPN technology provides the most cost-effective connectivity under most circumstances. Chinese regulations have placed restrictions affecting VPN technologies across international borders. For enterprises to achieve cross border connections, there are two options.
- 
    The enterprise can directly lease international dedicated lines from the 3 Chinese telecom carriers (China Telecom, China Mobile, China Unicom) in China.
- 
    Additionally, the enterprise can directly delegate a foreign telecom carrier with a presence in China to rent the international dedicated line (including VPN) from the 3 Chinese telecom carriers, and connect the corporate private network and equipment.
Note: The above cross-border connection methods must be used only for internal data exchange and office use. (Current as of 3 February 2018, subject to further regulatory developments.)
All devices located within mainland China will connect to Meraki China servers also located within China. Currently, only enterprise licensing is available for WAN Appliances located within China.
China Auto VPN Architecture
In the above diagram, we are utilizing Meraki Auto VPN to connect the enterprise sites inside China. The above diagram also demonstrates the Chinese government-approved dedicated circuits connecting the Chinese parts of the enterprise to the rest of the global enterprise. Dynamic routing such as BGP or OSPF can be utilized to exchange routing information between the domains.
