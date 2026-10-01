---
id: collect-261001-general-networking/general-networking/switches-layer-3-switching-layer-3-switch-overview-ed7e193d-2
title: "switches-layer-3-switching-layer-3-switch-overview-ed7e193d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/switches-layer-3-switching-layer-3-switch-overview-ed7e193d.md
source_anchor: ""
source_lines: [45, 57]
sha256: 6a5712cd32edeba27d24654ac037c640043144ed2b85e75b03b9a2375790ffe5
---

# switches-layer-3-switching-layer-3-switch-overview-ed7e193d

To start using layer 3 routing, navigate to the Switching > Configure > Routing & DHCP page. Alternatively, you could go to Switching > Monitor > Switches and click on the switch to be configured. Under L3 routing tab, click Configure - which takes you to the same Routing & DHCP page as above.
On the Routing & DHCP page, you will have the option to either "create interface" or to add an interface, if any L3 interfaces (SVI) or routed ports already exist in the network. Clicking on the available option will bring up the Interface Editor UI. Use the Interface Editor to configure your SVI or routed port.
Configuring an IPv4 L3 SVI Interface
- Interface name: A friendly name/description for the interface/VLAN.
- VLAN: The VLAN this L3 interface is in.
- Subnet: The network that this L3 interface is in, in CIDR notation (ex. 10.1.1.0/24).
- Interface IP: The IP address this switch will use for L3 routing on this VLAN/subnet. This cannot be the same as the switch's management IP.
- Multicast support: Enable multicast support if multicast routing between VLANs is required.
- Default gateway: When creating the first IPv4 interface on a switch, you will be prompted to enter a default gateway address. This is the next hop IPv4 address of another device on the network, used for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a L3 interface, and will be used for the default route next hop IP address.
- DHCP settings: If DHCP on this VLAN should be handled by the switch or forwarded to a server, make the appropriate selections. See the article on Configuring DHCP Services for more details.
- OSPF settings: This VLAN can be distributed via OSPF. See the MS OSPF Overview article for more details.
When complete, click Save or Save and add another to configure additional L3 interfaces.
 
