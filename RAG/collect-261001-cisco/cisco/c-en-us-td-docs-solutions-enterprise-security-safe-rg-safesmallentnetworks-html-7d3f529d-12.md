---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-12
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [382, 415]
sha256: bf9bf7d9d37574818cab488b20b851f7973c9f5f7c4e6c9ad37c51d98d67b6c5
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

To deploy the NAC Appliance in this manner, the client device must be directly connected to the network via a Catalyst switch port. After the user is authenticated and passes posture assessment, the Clean Access Manager (CAM) instructs the switch to map the user port from an unauthenticated VLAN (which switches or routes user traffic to the NAC) to an authenticated (authorized) VLAN that offers full access privileges. For example, as shown Figure 16, the client PC is connected through VLAN 110 to the NAC Clean Access Server for the authentication and posture assessment and is moved to VLAN 10 once it successfully completes the authentication and authorization, scan, and evaluation phases of the NAC solution.
In-Band Modes
When the NAC Appliance is deployed in-band, all user traffic, both unauthenticated and authenticated, passes through the NAC Appliance, which may be positioned logically or physically between end users and the network(s) being protected. See in-band virtual gateway topology (Figure 17) for a logical in-band topology example and physical in-band topology (Figure 18) for a physical in-band topology example.
Figure 17 In-Band Virtual Gateway Topology
Figure 18 Physical In-Band Topology
In-Band Virtual Gateway
When the NAC Appliance is configured as a virtual gateway, it acts as a bridge between end users and the default gateway (router) for the client subnet being managed. The following two bridging options are supported by the NAC Appliance:
•Transparent—For a given client VLAN, the NAC Appliance bridges traffic from its untrusted interface to its trusted interface. Because the appliance is aware of "upper layer protocols", by default it blocks all traffic except for Bridge Protocol Data Unit (BPDU) frames (spanning tree) and those protocols explicitly permitted in the "unauthorized" role, e.g., DNS and DHCP. In other words, it permits those protocols that are necessary for a client to connect to the network, authenticate, undergo posture assessment, and remediation. This option is viable when the NAC Appliance is positioned physically in-band between end users and the upstream network(s) being protected, as shown in Figure 18.
•VLAN mapping—This is similar in behavior to the transparent method, except that rather than bridging the same VLAN from the untrusted side to the trusted side of the appliance, two VLANs are used. For example, Client VLAN 131 is defined for the untrusted interface of the NAC Appliance. There is no routed interface or switched virtual interface (SVI) associated with VLAN 131. VLAN 31 is configured between the trusted interface of the NAC Appliance and the next-hop router interface/SVI for the client subnet. A mapping rule is made in the NAC Appliance that forwards packets arriving on VLAN 131 and forwards them out VLAN 31 by swapping VLAN tag information. The process is reversed for packets returning to the client. Note that in this mode, BPDUs are not passed from the untrusted-side VLANs to their trusted-side counterparts.
The VLAN mapping option is usually selected when the NAC Appliance is positioned logically in-band between clients and the networks being protected. This is the bridging option that should be used if the NAC Appliance is going to be deployed in the virtual gateway mode.
In-Band Real IP Gateway
When the NAC Appliance is configured as a "real" IP gateway, it behaves like a router and forwards packets between its interfaces. In this scenario, one or more client VLAN/subnets reside behind the untrusted interface. The NAC Appliance acts as a default gateway for all clients residing on those networks. Conversely, a single VLAN/subnet is defined on the trusted interface, which represents the path to the protected upstream network(s).
After successful client authentication and posture assessment, the NAC Appliance by default routes traffic from the untrusted networks to the trusted interface, where it is then forwarded based on the routing topology of the network.
The NAC Appliance is not currently able to support dynamic routing protocols. As such, static routes must be configured within the trusted side of the Layer 3 network for each client subnet terminating on or residing behind the untrusted interface. These static routes should reference, as a next hop, the IP address of the trusted interface of the NAC.
If one or more Layer 3 hops exist between the untrusted NAC interface and the end-client subnets, static routes to the client networks must be configured in the NAC Appliance. Likewise, a static default route (0/0) is required within the downstream Layer 3 network (referencing the IP address of the untrusted NAC interface) to facilitate default routing behavior from the client networks to the NAC Appliance.
Depending on the topology, multiple options exist to facilitate routing to and from the NAC Appliance, including static routes, VRF-Lite, MPLS VPN, and other segmentation techniques. It is beyond the scope of this document to examine all possible methods.
In-Band Versus Out-of-Band
Table 1 summarizes different characteristics of each type of deployment.
Table 1 In-Band Versus Out-of-Band Deployment Characteristics
The CAS is always inline with user traffic (both before and following authentication, posture assessment and remediation). Enforcement is achieved through being inline with traffic.
The CAS is inline with user traffic only during the process of authentication, assessment and remediation. Following that, user traffic does not come to the CAS. Enforcement is achieved through the use of SNMP to control switches and VLAN assignments to ports.
The CAS can be used to securely control authenticated and unauthenticated user traffic by using traffic policies (based on port, protocol, subnet), bandwidth policies, and so on.
The CAS can control user traffic during the authentication, assessment, and remediation phase, but cannot do so post-remediation since the traffic is out-of-band.
Does not provide switch port level control.
Provides port-level control by assigning ports to specific VLANs as necessary using SNMP.
In-band deployment is supported for wired and wireless clients.
OOB deployments support wired and wireless clients. Wireless OOB requires a specific network topology.
Cisco NAC Appliance In-Band deployment with supported Cisco switches is compatible with 802.1x.
Cisco does not recommend using 802.1x in an OOB deployment, as conflicts will likely exist between Cisco NAC Appliance OOB and 802.1x to set the VLAN on the switch interfaces/ports.
Out-of-Band Requirements
OOB implementation of Cisco NAC Appliance requires the switches and Wireless LAN Controllers to be supported by the Cisco NAC Appliance software. All the switches tested as part of the development of the Enterprise Design Profile for Small Enterprise Networks, apart from the Cisco Catalyst 2975, are supported by the Cisco NAC OOB and the Wireless LAN Controllers are also supported by the NAC Appliance software used in this document. If the Catalyst 2975 is to be used as an access switch with the Cisco NAC Appliance, the NAC solution must be an in-band solution.
Note To obtain the latest list of supported devices, check the latest version of the Cisco NAC Appliance-Clean Access Manager Installation and Administration Guide at: http://www.cisco.com/en/US/docs/security/nac/appliance/configuration_guide/45/cam/45cam-book.html.
Out-Of-Band, Layer 2, and Layer 3
The proposed design for the small enterprise network is an OOB design, in order to get the highest possible performance and scalability for traffic that has passed through the authentication, posture assessment, and remediation stages of NAC. The small enterprise network design offers two different access layer options, a Layer 2 access layer for smaller sites and a hybrid Layer 2/Layer 3 access layer for larger sites. This means that either a Layer 2 OOB solution or a Layer 3 OOB NAC solution may be deployed.
