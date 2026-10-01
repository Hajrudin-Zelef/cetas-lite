---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625-3
title: "docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625.md
source_anchor: ""
source_lines: [131, 173]
sha256: 6c6cd6eeaad581649b8b3d818956c93796d07791d8c026aed75a6f252fc9f1da
---

# docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625

In an OPNsense environment, VLANs operate by using VLAN tagging to distinguish traffic from different virtual networks as it passes through the same physical connection. This tagging process follows the IEEE 802.1Q standard, where a VLAN ID is inserted into the Ethernet frame header. When a packet reaches OPNsense, the firewall reads this tag to determine which virtual interface it belongs to, allowing administrators to apply specific routing, firewall rules, and monitoring policies for each VLAN.
VLAN connectivity in OPNsense relies on the proper configuration of trunk ports and access ports on connected switches. A trunk port is designed to carry traffic for multiple VLANs simultaneously and keeps the VLAN tags intact, making it ideal for links between switches or between a switch and OPNsense. In contrast, an access port carries traffic for only one VLAN and strips away any VLAN tags before delivering the data to an end device. This distinction ensures that endpoint devices, such as PCs or printers, connect to the correct VLAN without requiring VLAN configuration on the device itself.
When integrated with OPNsense, trunk ports are typically used on the firewall interfaces to connect to managed switches, enabling the firewall to receive and process traffic for all defined VLANs. Access ports on the switch, then distribute this traffic to devices within the respective VLANs. This setup allows OPNsense to act as a central policy enforcement point, controlling inter-VLAN routing, applying security rules, and monitoring network activity across all segments.
By combining VLAN tagging, trunking, and access port assignment, OPNsense provides a flexible yet secure network architecture. This ensures clear traffic separation, minimizes broadcast overhead, and allows administrators to manage multiple logical networks from a single firewall platform.
What Equipment is Required to Set Up VLANs on OPNsense?
Setting up VLANs on OPNsense* requires networking hardware that fully supports IEEE 802.1Q VLAN tagging. While OPNsense handles the creation of VLAN interfaces and the application of firewall rules, your physical equipment ensures stable data transfer and proper traffic separation. Choosing the right devices is crucial for achieving reliable segmentation, strong network security, and smooth integration between wired and wireless connections.
- 
Supported Firewall/Router: To implement VLANs, you first need an OPNsense-compatible firewall or router that can handle VLAN traffic efficiently. This is the core device where VLAN interfaces are created, tagged, and routed between network segments. The better the hardware performance, the more responsive your VLAN-based network will be. 
  - Dedicated OPNsense appliance (e.g., Netgate, Protectli, Qotom) or a custom-built PC/server running OPNsense
  - Network interface cards (NICs) that support VLAN tagging (single or multi-port options available)
 Optionally, multiple NICs to physically separate traffic if preferred
- 
Managed Switches: A managed switch that supports 802.1Q VLAN tagging is essential for defining trunk ports (carrying multiple VLANs) and access ports (serving a single VLAN). The switch ensures that VLAN assignments are correctly applied to each connected device. Popular choices include Cisco Catalyst, HP/Aruba, Netgear ProSAFE, TP-Link Omada, and Ubiquiti UniFi. Must support per-port VLAN configuration and tagging/untagging rules
- 
(Optional) Wireless Access Points: If you plan to extend VLANs to Wi-Fi networks, VLAN-aware access points allow you to map different SSIDs to specific VLAN IDs. This lets you separate guest Wi-Fi traffic from internal LAN traffic without additional cabling.
- 
Additional Equipment: A reliable VLAN setup requires high-quality cabling and proper physical organization of the network environment. 
  - Ethernet cables (Cat5e or higher) for gigabit or faster speeds
  - Patch panels or racks for larger deployments to keep cable management clean and accessible
When setting up an OPNsense guest network, assign it to a dedicated VLAN that has firewall rules preventing access to internal resources while still allowing internet connectivity. If wireless access is provided, configure the guest SSID on the access point to use this VLAN and ensure the switch port is set as a trunk carrying both guest and main VLAN traffic. This approach improves security and prevents cross-network interference, making it ideal for business or public environments.
How do I Enable VLAN Support on My Interface?
Enabling VLAN support in OPNsense is the first step to segmenting your network into multiple virtual LANs (VLANs) for better security, traffic management, and performance. By linking a VLAN ID to a physical interface, OPNsense can apply dedicated firewall rules, routing policies, and monitoring to that segment independently from the rest of the network.
To enable VLAN support in OPNsense, use the following steps.
- Open the OPNsense web interface and go to Interfaces → Devices → VLAN .
- Click + Add , choose theparent interface , and assign aVLAN tag based on your network design (e.g., VLAN 10 for guests, VLAN 20 for VoIP).
- Save and apply changes, then assign the VLAN under Interfaces → Assignments .
- Configure network settings such as IP address, subnet, and DHCP if required.
Once enabled, your VLAN is ready for integration into OPNsense firewall rules, DHCP services, and routing policies to ensure traffic isolation and optimized network performance.
How do I Assign a Parent Physical Interface for VLANs?
In OPNsense, a parent physical interface is the network port that carries VLAN-tagged traffic between the firewall and the rest of the network. When creating a VLAN in the GUI, you must choose this parent interface so OPNsense knows where to attach the VLAN ID.
The parent interface serves as the underlay for all VLANs associated with it. For example, if igb0 is your uplink to a managed switch, you can configure multiple VLANs, such as VLAN 10 for guests and VLAN 20 for VoIP, on top of that single physical connection. This approach avoids the need for separate cables or NICs for each VLAN, as the VLAN tags (defined by IEEE 802.1Q) are inserted into Ethernet frames before leaving the interface.
When selecting a parent interface in OPNsense, follow these steps.
- Choose the NIC that is physically connected to a VLAN-capable switch port configured as a trunk.
- Ensure the switch trunk is set to allow all VLAN IDs you plan to use.
- Avoid using an interface that is already assigned to another dedicated network unless you intend to combine it with VLANs on the same link.
By correctly assigning the parent physical interface, you establish a stable underlay for your VLANs, ensuring that tagged traffic is handled efficiently and passed to the correct logical interface. This design allows you to consolidate multiple networks over a single cable while maintaining complete separation between them at the logical level.
How do I Assign an IP Subnet to a VLAN Interface?
In OPNsense, assigning an IP subnet to a VLAN interface allows that VLAN to operate as its own independent network segment with its own addressing scheme. This is done by enabling the VLAN interface, selecting an IP configuration method, and defining the subnet details.
To assign an IP subnet to a VLAN interface:
- Go to Interfaces → Assignments and make sure the VLAN interface is listed and enabled.
- Click on the VLAN interface name (e.g., OPT1 or a custom label) to open its settings.
- In the IPv4 Configuration Type or IPv6 Configuration Type dropdown, choose either Static IPv4, Static IPv6, or DHCP depending on your design.
- If you choose Static IPv4 , enter the IP address and subnet mask in the format192.168.50.1/24 . This example creates a network where devices can have addresses in the 192.168.50.0 to 192.168.50.254 range, with .1 being the gateway.
