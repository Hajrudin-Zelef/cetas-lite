---
id: collect-261001-cisco/cisco/comparing-cisco-vpn-technologies-policy-based-vs-route-based-vpns
title: "comparing-cisco-vpn-technologies-policy-based-vs-route-based-vpns"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/comparing-cisco-vpn-technologies-policy-based-vs-route-based-vpns.md
source_anchor: ""
source_lines: [1, 41]
sha256: bf88c38aff8f133338efb3e674bfe0aa93c0c38192861cd972541a1ef963589e
---

# comparing-cisco-vpn-technologies-policy-based-vs-route-based-vpns

Comparing Cisco VPN Technologies – Policy Based vs Route Based VPNs

Virtual Private Networks constitute a hot topic in networking because they provide low cost and secure communications between sites (site-to-site VPNs) while improving productivity by extending corporate networks to remote users (remote access VPNs).

Cisco must be proud of its VPN solutions. It’s one of the few vendors that support such a wide range of VPN technologies with so many features and flexibility. Cisco Routers and Cisco ASA Firewalls are the two types of devices that are used most often to build Cisco Virtual Private Networks.

In this article we will discuss and compare two general Cisco VPN categories that are utilized by network engineers to build the majority of VPN networks in today’s enterprise environments. These categories are “Policy Based VPNs” (or IPSEC VPNs) and “Route Based VPNs”. Of course Cisco supports additional VPN technologies such as SSL VPNs (Anyconnect SSL VPN, Clientless SSL VPN), Dynamic Multipoint VPN (DMVPN), Easy VPN, Group Encrypted Transport (GET) VPN etc. Many of these VPN technologies are already covered on Firewall.cx and are beyond the scope of this article.

Below is a selection of Cisco VPN articles to which interested users can refer:

The diagram below shows a quick overview of the two VPN Categories we are going to discuss and their Practical Applications in actual networks:

For a Network Engineer or Designer it’s important to know the main differences between these two VPN categories and their practical applications. Knowing these will help professionals choose the right VPN type for their company and customers.

As shown in the diagram above, Policy-Based VPNs are used to build Site-to-Site and Hub-and-Spoke VPN and also remote access VPNs using an IPSEC Client. On the other hand, Route-Based VPNs are used to build only Site-to-Site or Hub-and-Spoke VPN topologies.

Now let’s see a brief description of each VPN Type.

Policy-Based IPSEC VPN

This is the traditional IPSEC VPN type which is still widely used today. This VPN category is supported on both Cisco ASA Firewalls and Cisco IOS Routers. With this VPN type the device encrypts and encapsulates a subset of traffic flowing through an interface according to a defined policy (using an Access Control List). The IPSEC protocol is used for tunneling and for securing the communication flow. Since the traditional IPSEC VPN is standardized by IETF, it is supported by all networking vendors so you can use it to build VPNs between different vendor devices as well.

Sample Configuration on Cisco ASA Firewalls

To illustrate the reason why this VPN type is called Policy-Based VPN, we will see a sample configuration code on a Cisco ASA firewall based on the diagram below.

From the configuration sample above, the access control list VPN-ACL defines the traffic flow that will pass through the VPN tunnel. Although there is other traffic flowing through the outside ASA interface, only traffic between LAN1 and LAN2 will pass through the VPN tunnel according to the traffic policy dictated by VPN-ACL. That’s exactly the reason why this VPN type is called Policy-Based VPN.

Understanding Route-Based VPNs

A route-based VPN configuration uses Layer3 routed tunnel interfaces as the endpoints of the VPN. Instead of selecting a subset of traffic to pass through the VPN tunnel using an Access List, all traffic passing through the special Layer3 tunnel interface is placed into the VPN. Therefore you need to configure routing accordingly. Either a dynamic routing protocol (such as EIGRP or OSPF) or static routing must be configured to divert VPN traffic through the special Layer3 tunnel interface.

This VPN Type is supported only on Cisco Routers and is based on GRE or VTI Tunnel Interfaces. For secure communication, Route-Based VPNs use also the IPSEC protocol on top of the GRE or VTI tunnel to encrypt everything.

Sample Configuration on Cisco Routers

Based on the network diagram below, let’s see a GRE Route-Based VPN with IPSEC protection.

From the configuration above, a GRE Layer3 Tunnel Interface is created (Tunnel0) which will be one of the endpoints of the VPN tunnel. IPSEC Protection is also applied for security. The other end of the VPN tunnel is Tunnel0 of the other site (with IP 10.0.0.2), thus forming a point-to-point VPN link. The static route shown above will divert VPN traffic destined for LAN2 via the Tunnel Interfaces.

Following is the VPN related configuration commands for our second router:
