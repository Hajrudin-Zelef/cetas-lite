---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1-3
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1.md
source_anchor: ""
source_lines: [74, 107]
sha256: bf491fd8ac6ea78688fb3f21ebc141004af1b2068fc840239348d589724ece7e
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1

Expected behavior. If the WAN Appliance loses connectivity to the VPN registry, peer information gets purged over time but not immediately. Connectivity to the registry matters when a node changes its contact information after losing connectivity to the VPN registry.
Both the hub and spoke will still be able to form the tunnel if the contact information remains the same, and they lost registry connectivity. Peer information will purge after a few hours causing the tunnel to be marked down.
Examples
In the example packet capture below, a WAN appliance is attempting to reach the VPN registry on UDP port 9350, but is receiving no response because an upstream firewall is preventing the outbound traffic:
In this example, the appropriate firewall rules have been added to allow the traffic to the VPN registry, and responses can be seen:
Solution
If this occurs, make sure that any upstream firewalls are configured to allow traffic to the IP addresses and ports listed on the Help > Firewall info page. Particularly for the VPN registry. It should also allow return traffic from established connections (this is allowed by default for stateful firewalls):
NAT Type: Unfriendly
UDP hole-punching, the mechanism used to establish the VPN connections between Cisco Meraki devices, relies on a consistent IP address and port for both devices involved. Two VPN registry servers are used for redundancy, and both expect to see the device as available on the same public IP address and port.
However, some NAT devices (such as a firewall) will rewrite the source ports differently for each VPN registry server. Other NAT devices or load balancers will attempt to spread the connections to each VPN registry server across two different public IP addresses. Both of these cases will result in the VPN connection failing, and marking the NAT as unfriendly:
Examples
In this example the upstream firewall rewrites the source port for each outbound connection differently. Notice that the first connection is changed to port 56125 while the second is instead 56126. When the registry servers see different source ports, the NAT unfriendly error will appear:
In this example, the upstream firewall is load balancing connections over two WAN connections, and then performing NAT using two different public IP addresses. Notice that the first connection is sent from the 198.51.100.23 address, while the second is sent from 198.51.100.17 instead. When the registry servers see different source IP addresses, the NAT unfriendly error will appear:
Solutions
If using a load balancer, or NAT across multiple public IP addresses, map traffic from the internal address of the appliance to a single public IP address. This will keep the public IP address seen by the VPN registry consistent.
-OR-
Select an arbitrary port that will be used for all VPN traffic to this WAN Appliance (e.g. UDP port 51625). Manually create a port mapping on the upstream firewall that will forward all traffic received on a specific public IP and port to the internal address of the appliance on the selected port. In Dashboard on the Security & SD-WAN > Configure > Site-to-site VPN page use the Manual: Port forwarding option for NAT traversal, and provide the public IP address and port that was configured. All peers will then connect using this IP address and port combination.
Troubleshooting Automatic NAT Traversal
Cisco Meraki VPN peers can use Automatic NAT Traversal to establish a secure IPsec tunnel through a firewall or NAT. When ACLs on an upstream firewall block source ports or more likely the case destination UDP ports in the range 32768-61000 on outbound traffic, a peer will not be able to punch a hole in the firewall and establish a tunnel with other remote peers.
Note: Cisco Meraki VPN peers must be able to use high number UDP ports to communicate with each other. Security systems such as firewalls that disallow this traffic may prevent successful traffic flow over the VPN. Please follow the diagnostic and troubleshooting steps below to resolve such issues.
In the example below, the firewall blocks peer1 from sending outbound UDP packets in the necessary destination port range. This prevents a hole punch. Now when peer2 tries to send inbound packets, its packets are dropped on the outside interface of the firewall because they do not match an existing outbound session. In this instance the tunnel will not be established between peer1 and peer2.
Failed connectivity tests or a VPN status of 'disconnected' indicates a tunnel failure between peers in Dashboard.
In a site-to-site VPN if two peers are unable to establish a VPN connection they will show as disconnected on each other's VPN status page. In this case, a packet capture should be taken on the primary Internet interface of both peers to analyze which firewall is blocking IPsec communication.
Analyzing a Packet Capture for IPsec Connectivity
Packet captures can be taken from Dashboard and downloaded as a .pcap file for analysis and filtering using Wireshark packet analyzer. They are invaluable for troubleshooting connections between hosts and isolating connectivity issues.
In the example below there is an Access Point to VPN concentrator tunnel that will not establish. We take packet captures from different points in the path to help determine which firewall is blocking the peer-to-peer communication.
The first capture, shown below, was taken from the wired interface of the Access Point 10.0.8.99. We can see the Access Point attempting to punch a hole in its local upstream firewall by sending packets to 208.72.143.11, which is the outside IP address of the NAT that the VPN concentrator sits behind. Notice the the Access Point is sending traffic to the concentrator but there is no return traffic in the capture from the WAN appliance behind the NAT.
Access Point 10.0.8.99:45540 -> WAN Appliance 208.72.143.11:53654
A second capture, shown below, was taken from the inside interface of the WAN Appliance upstream from the VPN concentrator. The VPN concentrator uses IP address 10.0.50.246 on the LAN. We can see the VPN concentrator sending packets to 208.72.143.18 which is the outside IP address of the NAT that the Access Point sits behind in an attempt to punch a hole in its local upstream firewall. Notice the VPN concentrator is sending traffic to the Access Point but no return traffic is present from the Access Point behind the NAT.
WAN Appliance 10.0.50.246:53654 -> Access Point 208.72.143.18:45540
A third capture was then taken, this time from the outside interface of the WAN Appliance upstream from the VPN concentrator. We can the see the VPN concentrator's traffic has been translated to 208.72.143.11, which is the firewall's outside IP address, and that it is being forwarded onto the Internet. This indicates the firewall is not blocking outbound IPsec traffic in the VPN concentrator site. However, we do not see any traffic originating from 208.72.143.18, the IP address of the NAT device the Access Point sits behind. From this we can conclude that the firewall upstream from the Access Point is blocking outbound IPsec traffic within the UDP port range 32768-61000.
WAN Appliance 208.72.143.18:53654 -> Access Point 208.72.143.18:45540
To confirm, we take a final capture from the outside interface of the WAN Appliance upstream from the Access Point, shown below. This capture shows packets originating from the VPN concentrator at 208.72.143.11 and arriving at the Access Point firewall's outside interface at 208.72.143.18. We still do not see any traffic originating from the Access Point being sent from the outside interface. This indicates the WAN Appliance is in fact blocking outbound IPsec traffic on the inside interface, specifically destination UDP port range 32768-61000.
WAN Appliance 208.72.143.18:53654 -> Access Point 208.72.143.18:45540
