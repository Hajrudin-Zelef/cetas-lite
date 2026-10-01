---
id: collect-261001-general-networking/general-networking/ms-layer-3-switching-layer-3-switch-example-0eed29dd
title: "ms-layer-3-switching-layer-3-switch-example-0eed29dd"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "training", "voice"]
source: docs/RAG/collect-261001-general-networking/ms-layer-3-switching-layer-3-switch-example-0eed29dd.md
source_anchor: ""
source_lines: [1, 80]
sha256: a44859339175482268460b7d5c787ce0aedc4366e6cfdbd9790119067e89e08d
---

# ms-layer-3-switching-layer-3-switch-example-0eed29dd

Layer 3 Switch Example
This article outlines a basic example of how layer 3 routing functionality on MS series switches could be implemented. Before proceeding, please refer to the Layer 3 Switch Overview for general information and configuration options.
Learn more with this free online training course on the Meraki Learning Hub:
Initial Topology
In this scenario, an MX security appliance is acting as the network gateway and firewall, performing NAT to a private subnet of 192.168.128.0/24 (VLAN 20). The MX is using an IP address of 192.168.128.254 on VLAN 20.
This leads down to a distribution switch that connects to both an access switch and a one-armed-router performing inter-VLAN routing for the network:
We will be reconfiguring the distribution switch to perform inter-VLAN routing for the network, so we can fully deprecate the legacy router.
Note: When designing a network with a layer 3 switch at the distribution layer, it is very important to understand which device is set as the gateway for clients on each subnet. If the L3 switch is the gateway for clients downstream subnets, any upstream firewall must be configured with a static route to that downstream subnet. If the firewall is configured with a VLAN interface for this downstream subnet, the firewall may receive incorrectly tagged traffic from this subnet (which will then be dropped).
Configuring the Layer 3 Interfaces
On the Distribution Switch, three layer 3 interfaces will be required. One for the uplink to the Firewall (which acts as the switch's default route), one for the data VLAN, and one for the voice VLAN. Configure the uplink interface first using the following steps:
- Navigate to the Distribution Switch's details page from Monitor > Switches.
- Click Initialize layer 3 under the Status section.
- Enter the following settings:
    
  - Name: Uplink
  - Subnet: 192.168.128.0/24
  - Interface IP: 192.168.128.1
  - VLAN: 20
  - Default Gateway: 192.168.128.254
- Click Save.
Note: Keep in mind that the management interface (whose IP can be found on the switch's details page) and this uplink interface are separate. Both interfaces can exist on the same VLAN/subnet, but the management interface must have a different IP configuration that allows it to communicate with the Internet.
Next configure the layer 3 interfaces for the data and voice VLANs by using the following steps:
- Navigate to Configure > Layer 3 routing.
- Click Add an interface.
- Select the Distribution Switch.
- Enter the following settings:
    
  - Name: Data
  - Subnet: 10.1.0.0/23
  - Interface IP: 10.1.1.254
  - VLAN: 5
  - Client Addressing: Relay DHCP to another server
  - DHCP server IPs: 192.168.128.254
- Click Save and add another.
- Select the Distribution Switch.
- Enter the following settings:
    
  - Name: Voice
  - Subnet: 10.1.2.0/23
  - Interface IP: 10.1.3.254
  - VLAN: 10
  - Client Addressing: Relay DHCP to another server
  - DHCP server IPs: 192.168.128.254
- Click Save
Configuring the Switch Ports
To allow for the downstream access switch and connected clients to take advantage of the routed interfaces, the switch port going to the access switch will need to be configured as a trunk to allow for both VLANs to traverse it. Under Configure > Switch ports select the port that will be connected to the access switch, and update the following settings:
- Type: Trunk
- Native VLAN: 1
- Allowed VLANs: All
The uplink port on the access switch should be configured identically, otherwise VLAN mismatches will result. The access switch will also need to be configured appropriately to place client traffic in the voice and data VLANs. Please see the Related KBs section for more details on how to configure switch ports. The LAN port on the firewall and the uplink port on the distribution switch also need to have similar settings, likely a trunk port, though configuration may vary as there is only one VLAN between the two devices.
Once this has been done, we can remove the legacy router from the network, as all routing functionality has been delegated to the distribution switch.
Additional Considerations
Now that the distribution switch is performing inter-VLAN routing for the network, we will need to perform some additional configuration steps on the firewall to allow full network connectivity:
- The firewall needs be configured with static routes (under Configure > Addressing & VLANs), so any inbound traffic destined for the voice or data VLANs will go through the routing interface of the switch. Those routes can be configured as follows:
    
  - Route to data VLAN:
        
    - Enabled: Yes
    - Name: Data
    - Subnet: 10.1.0.0/23
    - Next hop IP: 192.168.128.1
    - Active: Always
    - In VPN: No
  - Route to voice VLAN:
        
    - Enabled: Yes
    - Name: Voice
    - Subnet: 10.1.2.0/23
    - Next hop IP: 192.168.128.1
    - Active: Always
    - In VPN: No
- Route to data VLAN:
        
- Once the static routes have been added to the MX, DHCP scopes will need to be configured for each VLAN. Please reference our existing documentation on configuring DHCP services for configuration steps, using the switch's routing interface for each VLAN as the gateway:
    
  - Gateway IP for data: 10.1.1.254
  - Gateway IP for voice: 10.1.3.254
Once this has been completed, all clients on both VLANs can communicate both within and outside the network. Optionally, some additional configuration can be done to fulfill network requirements:
- The distribution switch can be configured with an IPv4 ACL to restrict communication between VLANs.
- If the firewall is participating in a site-to-site VPN, its static routes can be updated to allow voice and/or data clients to communicate over the VPN tunnel (by changing In VPN to Yes).
