---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-si-5a16b635
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-si-5a16b635"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "latency"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-si-5a16b635.md
source_anchor: ""
source_lines: [1, 21]
sha256: a7c5f2fc96c58a1d9a5b40ffcbaabd0e3ccfd69ec00a35a613cb359c97a7bdbf
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-si-5a16b635

Site-to-Site VPN Failover Behavior
Overview
MX Security Appliances make use of multiple technologies to provide site-to-site VPN redundancy in a variety of different scenarios.
AutoVPN
Meraki AutoVPN is a site-to-site technology that uniquely leverages the Meraki Dashboard for centrally-managed cloud orchestration of site-to-site VPN connections between MX devices.
Uplink Failure
Uplink failover is a feature built into the MX series to keep a constant connection in the event of Primary link failure.
Failover occurs when the primary uplink of the MX is unable to reach the Internet. By default, AutoVPN traffic will egress the primary uplink, as configured under Security & SD-WAN > Configure > SD-WAN & traffic shaping. WAN uplink failover occurs automatically when multiple uplinks are in use. For more information on uplink failover, please see this article.
Datacenter Redundancy
In addition to uplink redundancy, MX security appliances also support datacenter redundancy (DC-DC failover).
In a DC-DC failover design, a spoke MX will form AutoVPN tunnels to multiple VPN hubs in different datacenters. These VPN hubs advertise availability of the same subnet(s) into the AutoVPN topology. The spoke MX will send traffic destined for a shared resource to the highest priority hub advertising the subnet(s). If connectivity to a hub is lost, spoke MX sites will automatically send traffic for shared resources to the next highest priority hub advertising the subnet(s).
SD-WAN
MX appliances complement uplink failover with SD-WAN features to provide additional resiliency to brown-outs and changing WAN performance conditions.
MXs with dual WAN uplinks form concurrently active VPN tunnels across both uplinks. The loss, latency, and jitter of these VPN tunnels are monitored and tracked. The MX compares these metrics with policies that have been configured by the network administrator to dynamically shift traffic to different VPN paths. SD-WAN policies can be configured under Security & SD-WAN > Configure > SD-WAN & traffic shaping by configuring VPN flow preferences and custom performance classes. For more information on SD-WAN, please see this article.
IPsec site-to-site VPN
If the MX in question has an established VPN tunnel with a IPsec VPN peer, the IPsec VPN device will need to have the ability to designate a backup (failover) peer IP. By designating the public IP address of the MX's secondary uplink as the back-up VPN IP on the IPsec VPN peer, you can ensure that the VPN tunnel will be re-established in the event of an uplink failure.
Additional Resources
Warm Spare
In addition to site-to-site VPN redundancy, the MX product family also supports the ability to configure a warm spare appliance. A warm spare configuration can be leveraged in both the branch (spoke), as well as in the datacenter (hub) as part of a highly available VPN architecture. For more information on warm spare, please see this article.
MX Routing Behavior
For more detailed information on how routing on MX appliances occurs please see this article. This also includes a detailed review of how routing is performed in an example DC-DC failover configuration.
