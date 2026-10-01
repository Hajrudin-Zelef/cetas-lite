---
id: collect-261001-general-networking/general-networking/switches-layer-3-switching-layer-3-switch-overview-ed7e193d-3
title: "switches-layer-3-switching-layer-3-switch-overview-ed7e193d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/switches-layer-3-switching-layer-3-switch-overview-ed7e193d.md
source_anchor: ""
source_lines: [58, 145]
sha256: 83a9c7d7aa2f8d0b48f05defb4047600984f2a525e18887f1127a68b8a848388
---

# switches-layer-3-switching-layer-3-switch-overview-ed7e193d

L3 configuration changes on MS210, MS225, MS250, MS350, MS355, MS410, MS425, MS450 require the flushing and rebuilding of L3 hardware tables. As such, momentary service disruption may occur. We recommend making such changes only during scheduled downtime/maintenance window.
Configuring an IPv4 L3 Routed Port Interface
Note: This feature is only supported on the MS390 and Cloud Managed Catalyst switches.
- Interface mode: Select "Routed port".
- Switch or switch stack: Select the switch or switch stack for the routed port.
- Select Module (optional): Select if you have a modular Catalyst switch. If fixed then this will be greyed out.
- Switch ports: Select the routed port interface number.
- Name: A friendly name/description for the interface/routed port.
- VRF: Select the VRF if other than the default.
- IP toggle: Select if this routed port will be IPv4, IPv6, or Both
- Uplink: Select if this routed port will be the uplink interface to Dashboard.
- Subnet: The network that this L3 interface is in, in CIDR notation (ex. 10.1.1.0/24).
- Interface IP: The IP address this switch will use for layer 3 routing on this VLAN/subnet. This cannot be the same as the switch's management IP.
- Multicast support: Enable multicast support if multicast routing between VLANs is required.
- Default gateway (IPv4): When creating the first IPv4 interface on a switch, you will be prompted to enter a default gateway address. This is the next hop IPv4 address of another device on the network, used for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a L3 interface, and will be used for the default route next hop IP address.
- DNS Server: If this interface is selected as uplink, configure DNS servers.
- High Availibility (VRRP): Select if runing VRRP on this routed port.
- DHCP settings: If DHCP on this VLAN should be handled by the switch or forwarded to a server, make the appropriate selections. See the article on Configuring DHCP Services for more details.
- OSPF settings: This VLAN can be distributed via OSPF. See the MS OSPF Overview article for more details.
When complete, click Save changes.
Note: VRRP is not currently supported on routed ports.
Configuring an IPv6 L3 SVI Interface
Note: This feature is only supported on the MS390 and Cloud Managed Catalyst switches.
- Interface name: A friendly name/description for the interface/VLAN.
- VLAN: The VLAN this L3 interface is in.
- Prefix: The IPv6 subnet that this L3 interface is in, in CIDR notation (ex. 2001:db8::/32).
- IPv6 EUI64: Option to use EUI (extended unique identifier) allowing the switch to automatically dervice the interface IPv6 address from the switch's MAC address. This option can only be used if the prefix length is /64.
- Interface IPv6: The IPv6 address this switch will use for L3 routing on this VLAN/subnet. This cannot be the same as the switch's management IPv6 address. If the interace is configure to use EUI64, this option will be disabled.
- Default gateway: When creating the first IPv6 interface on a switch, you will be prompted to enter a default gateway address. This is the next hop IPv6 of a another device on the network, used address for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a L3 interface, and will be used for the default route next hop IP address.
Once created, any layer 3 interfaces or static routes will appear under Switching > Configure > Routing & DHCP.
Note: Each switch can only have a single L3 interface per VLAN.
Configuring an IPv6 L3 Routed Port Interface
Note: This feature is only supported on the MS390 and C9300-M models.
- Interface mode: Select "Routed port".
- Switch or switch stack: Select the switch or switch stack for the routed port.
- Select Module (optional): Select if you have a modular Catalyst switch. If fixed then this will be greyed out.
- Switch ports: Select the routed port interface number.
- Name: A friendly name/description for the interface/routed port.
- VRF: Select the VRF if other than the default.
- IP toggle: Select if this routed port will be IPv4, IPv6, or Both
- Uplink: Select if this routed port will be the uplink interface to Dashboard.
- Prefix: The network that this L3 interface is in, in CIDR notation (ex. 10.1.1.0/24).
- Interface IPv6: The IP address this switch will use for L3 routing on this VLAN/subnet. This cannot be the same as the switch's management IP.
- Multicast support: Enable multicast support if multicast routing between VLANs is required.
- Default gateway (IPv6): When creating the first IPv4 interface on a switch, you will be prompted to enter a default gateway address. This is the next hop IPv4 address of another device on the network, used for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a L3 interface, and will be used for the default route next hop IP address.
- DNS Server: If this interface is selected as uplink, configure DNS servers.
- High Availibility (VRRP): Select if runing VRRP on this routed port.
When complete, click Save changes.
Configuring an L3 Aggregate Routed Port Interface
- 
    Navigate to the Switch Summary page.
- 
    Click on Ports
- 
    Select the ports you want to convert into a L3 aggregate interface and click Aggregate
- 
    Select the Aggregate interface you just created and click Edit
- 
    Select Routed Port for Interace mode and then click Update (this will bring you to the Interface Editor in the Routing & DHCP page to begin configuring the L3 Aggregate interface.
- 
    Name: A friendly name/description for the interface/routed port.
- 
    VRF: Select the VRF if other than the default.
- 
    IP toggle: Select if this routed port will be IPv4, IPv6, or Both
- 
    Uplink: Select if this routed port will be the uplink interface to Dashboard.
- 
    Subnet: The network that this L3 interface is in, in CIDR notation (ex. 10.1.1.0/24).
- 
    Interface IP: The IP address this switch will use for L3 routing on this VLAN/subnet. This cannot be the same as the switch's management IP.
- 
    Multicast support: Enable multicast support if multicast routing between VLANs is required.
- 
    Default gateway (IPv4): When creating the first IPv4 interface on a switch, you will be prompted to enter a default gateway address. This is the next hop IPv4 address of another device on the network, used for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a L3 interface, and will be used for the default route next hop IP address.
- 
    DNS Server: If this interface is selected as uplink, configure DNS servers.
- 
    High Availibility (VRRP): Select if runing VRRP on this routed port.
- 
    DHCP settings: If DHCP on this VLAN should be handled by the switch or forwarded to a server, make the appropriate selections. See the article on Configuring DHCP Services for more details.
- 
    OSPF settings: This VLAN can be distributed via OSPF. See the MS OSPF Overview article for more details.
When complete, click Save changes.
Note: When deleting an L3 aggregate routed port interface from the Routing and DHCP page, the interface and its member ports are converted to an L2 aggregate. You can then convert the aggregate back to an L2 interface, split the member ports to remove the aggregation, or reconfigure the interface as a routed port.
Configuring Static Routes
In order to route traffic elsewhere in the network, static routes must be configured for subnets that are not being routed by the switch or would not be using the default route already configured, such as if another portion of the network was located behind a router or another L3 switch is downstream from the Cisco Meraki L3 switch being configured.
To create a new static route:
