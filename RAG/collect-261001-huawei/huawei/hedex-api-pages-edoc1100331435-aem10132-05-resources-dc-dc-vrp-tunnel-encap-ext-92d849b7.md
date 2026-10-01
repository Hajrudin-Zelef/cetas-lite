---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-vrp-tunnel-encap-ext-92d849b7
title: "hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-vrp-tunnel-encap-ext--92d849b7"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-vrp-tunnel-encap-ext--92d849b7.md
source_anchor: ""
source_lines: [1, 33]
sha256: 6744a813d793828f626b8869a2fdd12b24383ee369d5186065d01175d20aec08
---

# hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-vrp-tunnel-encap-ext--92d849b7

Currently, only an EVPN can be used as a service network for SD-WAN EVPN.
On an SD-WAN EVPN network, an EVPN needs to be configured between PEs to function as a service network.
The system view is displayed.
A VPN instance is created, and the VPN instance view is displayed.
A VN ID is bound to the VPN instance.
The IPv4 address family is enabled in the VPN instance, and the VPN instance IPv4 address family view is displayed.
A route distinguisher (RD) is configured for the VPN instance.
A VPN instance takes effect only after an RD is configured for it. The RDs of different VPN instances on the same PE must be different.
EVPN-VPN targets are configured for the VPN instance.
EVPN-VPN targets are BGP extended community attributes used to control the receiving and advertisement of EVPN routes. A maximum of eight EVPN-VPN targets can be configured using the vpn-target evpn command at a time. To configure more EVPN-VPN targets in the EVPN instance address family, run the vpn-target evpn command multiple times.
The VPN instance IPv4 address family is associated with the route-policy created when the Color attribute is set. In this manner, the routes advertised to the EVPN address family carry the Color attribute for tunnel recursion.
The VPN instance IPv4 address family is associated with an import route-policy to filter the EVPN routes imported from the EVPN address family.
To precisely control EVPN routes, an import route-policy must also be configured. An import route-policy filters routes that are received from the EVPN address family.
Return to the VPN instance view.
Return to the system view.
The view of the tunnel interface to be bound to a VPN instance is displayed.
The tunnel interface is bound to the VPN instance.
The one-to-one mapping between the tunnel interface and VPN instance is configured.
A primary IP address is configured for the interface.
A tunnel protocol on tunnel interface is set to SVPN.
By default, the tunnel protocol is none, indicating that packets are not encapsulated.
The weak protocol status check function in the Lone Ranger SVPN mode.
By default, the weak protocol status check function in the Lone Ranger SVPN mode is disabled.
The BGP view is displayed.
The BGP-EVPN address family is enabled, and the BGP-EVPN address family view is displayed.
The capability to exchange EVPN routes with a peer or peer group is enabled.
The device is enabled to advertise SD-WAN tunnel routes to the BGP EVPN peer.
The BGP EVPN peer is added to a peer group.
Adding BGP EVPN peers to a peer group simplifies BGP network configuration and management.
A route-policy is specified for the BGP EVPN peer or peer group to advertise only specified routes.
A route-policy is specified for the BGP EVPN peer or peer group to receive only specified routes.
To precisely control EVPN routes, an import route-policy must also be configured. An import route-policy filters routes that are received from other BGP EVPN peers or peer groups.
The device is disabled from filtering received EVPN routes based on EVPN-VPN targets.
