---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-94f1c4e2-1
title: "platform-management-dashboard-administration-design-and-configure-architectures--94f1c4e2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--94f1c4e2.md
source_anchor: ""
source_lines: [1, 76]
sha256: 840f7d40256d320c6576333139b757fc121a8d7397b6dcd94272f5321df2d154
---

# platform-management-dashboard-administration-design-and-configure-architectures--94f1c4e2

Meraki Auto VPN General Best Practices
Click 日本語 for Japanese
Auto VPN Best Practices
The best practices listed here focus on the most common deployment scenario, but is not intended to preclude the use of alternative topologies. The recommended SD-WAN architecture for most deployments is as follows:
- 
    WAN Appliance at the datacenter deployed as a one-armed concentrator
- 
    Warm spare/High Availability at the datacenter
- 
    OSPF route advertisement for scalable upstream connectivity to connected VPN subnets
- 
    Datacenter redundancy
- 
    Split tunnel VPN from the branches and remote offices
- 
    Dual WAN uplinks at all branches and remote offices
Auto VPN at the Branch
Before configuring and building Auto VPN tunnels, there are several configuration steps that should be reviewed.
WAN Interface Configuration
While automatic uplink configuration via DHCP is sufficient in many cases, some deployments may require manual uplink configuration of the WAN appliance at the branch. The procedure for assigning static IP addresses to WAN interfaces can be found in our MX IP assignment documentation.
Some WAN Appliance models have only one dedicated Internet port and require a LAN port be configured to act as a secondary Internet port via the device local status page if two uplink connections are required. WAN Appliance models that require reconfiguring a LAN port as a secondary Internet port currently include the MX64 line, MX67 line, and MX100 devices. This can also be verified per-model in our installation guides online. This configuration change can be performed on the device local status page on the Configure tab.
Subnet Configuration
Auto VPN allows for the addition and removal of subnets from the Auto VPN topology with a few clicks. The appropriate subnets should be configured before proceeding with the site-to-site VPN configuration.
Hub Priorities
Hub priority is based on the position of individual hubs in the list from top to bottom. The first hub has the highest priority, the second hub the second highest priority, and so on. Traffic destined for subnets advertised from multiple hubs will be sent to the highest priority hub that a) is advertising the subnet and b) currently has a working VPN connection with the spoke. Traffic to subnets advertised by only one hub is sent directly to that hub.
Configuring Allowed Networks
To allow a particular subnet to communicate across the VPN, locate the local networks section in the Site-to-site VPN page. The list of subnets is populated from the configured local subnets and static routes in the Addressing & VLANs page, as well as the Client VPN subnet if one is configured.
To allow a subnet to use the VPN, set the Use VPN drop-down to yes for that subnet.
Auto VPN at the Data Center
Deploying a One-Armed Concentrator
A one-armed concentrator is the recommended datacenter design choice for an SD-WAN deployment. The following diagram shows an example of a datacenter topology with a one-armed concentrator:
NAT Traversal
Whether to use Manual or Automatic NAT traversal is an important consideration for the VPN concentrator.
Use manual NAT traversal when:
- 
    There is an unfriendly NAT upstream.
- 
    Stringent firewall rules are in place to control what traffic is allowed to ingress or egress the datacenter.
- 
    It is important to know which port remote sites will use to communicate with the VPN concentrator.
If manual NAT traversal is selected, it is highly recommended that the VPN concentrator be assigned a static IP address. Manual NAT traversal is intended for configurations when all traffic for a specified port can be forward to the VPN concentrator.
Use automatic NAT traversal when:
- 
    None of the conditions listed above that would require manual NAT traversal exist.
If automatic NAT traversal is selected, the WAN Appliance will automatically select a high numbered UDP port to source Auto VPN traffic from. The VPN concentrator will reach out to the remote sites using this port, creating a stateful flow mapping in the upstream firewall that will also allow traffic initiated from the remote side through to the VPN concentrator without the need for a separate inbound firewall rule.
Datacenter Redundancy (DC-DC Failover)
Meraki WAN Appliances support datacenter to datacenter redundancy via our DC-DC failover implementation. The same steps used above can also be used to deploy one-armed concentrators at one or more additional data centers. For further information about VPN failover behavior and route prioritization, refer to our DC-DC Failover documentation.
This section outlines the steps required to configure and implement warm spare high availability (HA) for a WAN Appliance operating in VPN concentrator mode.
Topology
The following is an example of a topology that leverages an HA configuration for VPN concentrators:
Behavior
When configured for high availability (HA), one WAN Appliance is active, serving as the primary, and the other WAN Appliance operates in a passive, standby capacity. The VRRP protocol is leveraged to achieve failover. Check out our MX Warm Spare documentation for more information.
WAN Appliance IP Assignment
In the datacenter, a WAN Appliance can operate using a static IP address or an address from DHCP. WAN appliances will attempt to pull DHCP addresses by default. It is highly recommended to assign static IP addresses to VPN concentrators.
Uplink IPs
Use Uplink IPs is selected by default for new network setups. In order to properly communicate in HA, VPN concentrator WAN Appliances must be set to use the virtual IP (VIP).
Virtual IP (VIP)
Virtual IP is an addressing option that uses an additional (third) IP address that is shared by the HA WAN Appliances. In this configuration, the WAN Appliances will send their cloud controller communications via their uplink IPs, but other traffic will be sent and received by the shared virtual IP address.
WAN Appliance Data Center Routing
The WAN Appliance acting as a VPN concentrator in the datacenter will be terminating remote subnets into the datacenter. In order for bi-directional communication to take place, the upstream network must have routes for the remote subnets that point back to the WAN Appliance acting as the VPN concentrator.
If OSPF route advertisement is not being used, static routes directing traffic destined for remote VPN subnets to the WAN Appliance VPN concentrator must be configured in the upstream routing infrastructure.
If OSPF route advertisement is enabled, upstream routers will learn routes to connected VPN subnets dynamically.
Failover Times
There are several important failover timeframes to be aware of:
| Service | Failover Time | Failback Time | 
| Auto VPN Tunnels | 30-40 seconds | 30-40 seconds | 
| DC-DC Failover | 20-30 seconds | 20-30 seconds | 
| Dynamic Path Selection | Up to 30 seconds | Up to 30 seconds | 
| Warm Spare | 30 seconds or less | 30 seconds or less | 
| WAN connectivity | 300 seconds or less | 15-30 seconds | 
Configuring OSPF Route Advertisement
Meraki WAN Appliances support advertising routes to connected VPN subnets via OSPF.
A WAN Appliance with OSPF route advertisement enabled will only advertise routes via OSPF; it will not learn OSPF routes.
Note: WAN Appliances in Routed mode only support OSPF on firmware versions 13.4+, when using the "Single LAN" LAN setting. OSPF is otherwise supported when the WAN Appliance is in passthrough mode on any available firmware version. This can be set under Security & SD-WAN > Configure > Addressing & VLANs.
When spoke sites are connected to a hub WAN Appliance, the routes to spoke sites are advertised using an LS Update message. These routes are advertised as type 2 external routes.
BGP and Auto VPN
