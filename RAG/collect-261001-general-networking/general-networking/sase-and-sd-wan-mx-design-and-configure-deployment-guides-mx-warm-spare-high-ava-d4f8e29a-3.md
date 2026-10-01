---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a-3
title: "sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a.md
source_anchor: ""
source_lines: [93, 133]
sha256: 08602b8e706f88fb185288482199263fa75eb807e8cc79a7d0d1696385f79dcc
---

# sase-and-sd-wan-mx-design-and-configure-deployment-guides-mx-warm-spare-high-ava-d4f8e29a

- Using the warm spare "swap button" on MX that are configured to use virtual IPs will change the shared virtual WAN and LAN MAC addresses, which will most likely cause connectivity disruption. It is not recommended to use it in production hours.
- Connection monitoring and management traffic will still use WAN interface MAC and IP addresses, even if VIP is configured.
Note: MX devices in HA randomly assign a VRRP VRID between 1 and 255 at boot. For a static VRID, please contact Support for assistance.
Failure Detection
There are two failure detection methods for router mode warm spare.
WAN failover: WAN monitoring is performed using the same internet connectivity tests that are used for uplink failover. For more data on these checks, see the Connection Monitoring for WAN Failover article. If the primary appliance does not have a valid internet connection based on these tests, it will stop sending VRRP heartbeats, which will result in a failover. When uplink connectivity on the original primary appliance is restored and the warm spare begins receiving VRRP heartbeats again, it will relinquish the active role back to the primary appliance.
LAN failover: The two appliances share health information over the network via the VRRP protocol. These VRRP heartbeats occur at layer two and are performed on all configured VLANs. When no advertisements reach the spare on all VLANs, it will transition into an active state. When the warm spare begins receiving VRRP heartbeats again, it will transition back into a passive, ready state.
Requirements and Best Practices
When configuring routed HA, it is critical that both MXs have a reliable connection to each other on the LAN, so the heartbeats of the primary MX can be seen reliably by the spare. To ensure this connection is reliable:
- The two MXs should be connected to each other through a downstream switch (or ideally, multiple switches) on the LAN to allow for passing VRRP heartbeats.
    
  - There should be no more than one additional hop between them, and they must be able to communicate on all VLANs.
  - Make sure Spanning-Tree Protocol (STP) is enabled on the downstream switching infrastructure, as a properly-configured HA topology will introduce a loop on the network.
- When first configuring routed HA, the spare should be added and configured in the dashboard before the device is physically deployed, so it will immediately fetch its configuration and behave appropriately.
- Ensure that both MXs have their own uplink IP address for dashboard connectivity as described in the Uplink IP Configuration section.
    
  - If a virtual IP is being used, an additional IP address is needed, and all three IPs must be in the same subnet.
Cellular Failover Behavior
Meraki supports cellular failover with high-availability (HA) pair, limited to the MX67C and MX68CW models with embedded cellular modules. In order to support HA, customers must be using firmware MX 14.53, MX 15.42, or MX 16.11 or higher. At this time, if a cellular uplink is used in an HA pair, the following will occur in order:
- Primary MX WAN 1+2 fails > fails over to secondary MX
- Secondary MX WAN 1+2 fails > fails over to primary MX cellular
- Primary MX cellular fails > fails over to secondary MX cellular
Note: While it is possible to use cellular failover as described above, it is not officially supported by Meraki if leveraging other MX models and USB cellular dongle.
While DDNS is enabled, if all Primary MX WAN links fail but cellular is still active, DDNS will resolve to the Primary MX cellular uplink IP. It is recommended that you perform a swap of the Secondary MX to the Primary to prevent any issues with DDNS updates while the original primary MX WAN links are offline.
Recommended Topologies
There are two physical architectures available for routed warm spare deployments.
Fully Redundant (Two Switches)
In this architecture, the primary and secondary MXs are not directly connected, and VRRP heartbeats are carried between the downstream switches. This is the recommended architecture for most deployments, as there is no single point of failure in this topology.
Fully Redundant (Switch Stack)
In this architecture, the primary and secondary MXs are connected via a downstream switch stack. Each switch has at least one uplink to each MX, ensuring there is no single point of failure in the topology.
High Availability with More Than Two Physical WAN Uplinks
Although only two active uplinks are supported at a time on an active/primary MX, additional uplinks should be utilized for tertiary failover on the secondary MX. One or two additional uplinks may be utilized on the secondary MX, and will become active when all uplinks on the primary MX fail or when a hardware failure occurs on the primary MX. These additional uplinks connected to the secondary MX can be part of a different IP subnet than the uplinks on the primary MX.
Troubleshooting Routed Warm Spare
If there is a problem with the routed HA configuration, there may be various symptoms that will affect the network, and it may not be obvious that the root cause is routed HA. This section outlines what issues with HA typically look like, as well as recommended troubleshooting steps.
Dual Active Issue
The most common sign of a problem with routed HA is a dual active scenario where both the primary and spare MX report in the dashboard as being active. This can be observed in the dashboard under Security & SD-WAN > Monitor > Appliance status and by comparing the current state of each appliance.
This will occur if the primary MX is online and sending heartbeats that aren't seen by the spare, resulting in the spare thinking that the primary is down. If both the primary and spare are in the active state, this will cause various issues with the network, affecting DHCP, routing, VPN, etc.
Recommended Troubleshooting Steps
If network issues appear to be related to routed HA, follow the troubleshooting steps to identify the root cause:
- Check both appliances in the dashboard (under Security & SD-WAN > Monitor > Appliance status) to check if there is a dual active scenario as outlined above.
    
