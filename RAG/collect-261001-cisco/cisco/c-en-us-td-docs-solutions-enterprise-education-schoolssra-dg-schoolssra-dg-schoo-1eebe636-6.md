---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636-6
title: "c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636.md
source_anchor: ""
source_lines: [197, 242]
sha256: e7962ffce680f34d294d2991696dcfb42bb22420f172d1a338deda93563164e7
---

# c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636

The CISF is a set of native security features available on Cisco Catalyst Switches and designed to protect the access infrastructure and users from spoofing, man-in-the-middle, DoS and other network-based attacks. CISF includes features such as private VLANs, port security, DHCP snooping, IP Source Guard, secure Address Resolution Protocol (ARP) detection, and dynamic ARP inspection (DAI). CISF features are considered to be part of a security baseline and should be deployed on all access ports.
In addition to using the CISF features to secure the access ports, access control solutions like IBNS or NAC appliance may be deployed to control who can access the network and ensure endpoint compliance with the schools security policies.
The Cisco IBNS solution is a set of Cisco IOS software services designed to enable secure user and host access to enterprise networks powered by Cisco Catalyst switches and wireless LANs. It provides standards-based network access control at the access layer by using the 802.1X protocol to secure the physical ports where end users connect. 802.1X is an IEEE standard for media-level (Layer 2) access control, offering the capability to permit or deny network connectivity based on the identity of the end user or device enabling enterprise policy enforcement of all users and hosts, whether managed or unmanaged. In addition to holistic access-control provided by 802.1X, IBNS also offers device-specific access-control through MAC-Authentication Bypass (MAB) and user-based access-control through Web-Auth.
The Cisco Network Admission Control (NAC) Appliance uses the network infrastructure to enforce security policy compliance on all devices seeking to access network computing resources. With NAC Appliance, network administrators can authenticate, authorize, evaluate, and remediate wired, wireless, and remote users and their machines prior to network access. It identifies whether networked devices such as laptops, IP phones, or game consoles are compliant with your network's security policies and repairs any vulnerability before permitting access to the network. Noncompliant machines are redirected into a quarantine area, where remediation occurs at the discretion of the administrator.
The choice of which access control solution to use depends on the security goals and the direction of the network design. For networks using or moving towards 802.1x-based wired or wireless access and interested in identity-based access control, Cisco IBNS solution should be considered. For networks requiring role-based access control using posture assessments to ensure security compliance, Cisco NAC Appliance should be considered.
The Catalyst Integrated Security Features (CISF), and the Cisco Identity-Based Networking Services (IBNS) and NAC Appliance access control solutions are discussed in Chapter 9, "Access Layer Security Design."
Endpoint Protection
Servers, desktop computers, laptops, printers, and IP phones are examples of the diverse network endpoints commonly present in School environments. The great variety in hardware types, operating systems, and applications represents a clear challenge to security. In addition, students and staff may bring laptops and other portable devices that could also be used outside the school's premises, potentially introducing viruses, worms, spyware, and other type of malware.
Properly securing the endpoints requires not only adoption of the appropriate technical controls but also end-user awareness. While this document focuses on the implementation of the technical controls, the school's security strategy must include security awareness campaigns and programs. Students and staff must be continuously educated on current threats, use best practices, and the security measures needed for keeping endpoints up-to-date with the latest updates, patches, and fixes.
Following the best practices of Cisco SAFE, the school architecture implements a range of security controls designed to protect the endpoints. These include host-based Cisco IPS, network-based intrusion prevention systems, and web and E-mail traffic security.
As Cisco host-based IPS, the architecture leverages the Cisco Security Agent (CSA). CSA takes a proactive and preventative approach, using behavior-based security to focus on preventing malicious activity on the host. Malicious activity is detected and blocked, independent of the type of malware, spyware, adware, or virus affecting the host.
Once deployed on an endpoint, whenever an application attempts an operation, the agent checks the operation against the application's security policy—making a real-time allow or deny decision on the continuation of that operation and determining whether logging the operation request is appropriate. Security policies are collections of rules that IT or security administrators assign to protect servers and desktops, either individually or organization-wide. CSA provides defense-in-depth protection against spyware and adware by combining security policies that implement distributed firewall, operating system lockdown and integrity assurance, malicious mobile code protection, and audit-event collection capabilities in default policies for servers and desktops.
CSAs are centrally managed with the CSA Management Center (CSA-MC), which in the Cisco SAFE design is placed in a secure segment in the data center. The Management Center (MC) also provides centralized reporting and global correlation.
For complete details about deploying CSA in a network, refer to the Rapid Deployment Guide for Cisco Security Agent 6.0 for Desktops at the following URL:
http://www.cisco.com/en/US/prod/collateral/vpndevc/ps5707/ps5057/deployment_guide_c07-501928.htm
Internet Perimeter Security
Internet Border Router
Whether the Internet border router is managed by the school or the ISP, it must be hardened following the best practices listed in the Network Foundation Protection section. This includes restricting and controlling administrative access, protecting the management and control planes, and securing the dynamic exchange of routing information. In addition, the Internet border router may be leveraged as the first layer of protection against outside threats. To that end, edge ACLs, uRPF and other filtering mechanisms may be implemented for anti-spoofing and to block invalid packets.
The following configuration snippet illustrates the structure of an edge ACL applied to the upstream interface of the Internet border router. The ACL is designed to block invalid packets and to protect the infrastructure IP addresses from the Internet. The configuration assumes the school is assigned the 198.133.219.0/24 address block, and that the upstream link is configured in the 64.104.10.0/24 subnet.
! !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!--- Module 1: Anti-spoofing Denies
!--- These ACEs deny fragments, RFC 1918 space,
!--- invalid source addresses, and spoofs of
!--- internal space (space as an external source).
!
!--- Deny fragments.
access-list 110 deny tcp any 198.133.219.0 0.0.0.255 fragments
access-list 110 deny udp any 198.133.219.0 0.0.0.255 fragments
access-list 110 deny icmp any 198.133.219.0 0.0.0.255 fragments
!--- Deny special-use address sources.
!--- See RFC 3330 for additional special-use addresses.
access-list 110 deny ip host 0.0.0.0 any
access-list 110 deny ip 127.0.0.0 0.255.255.255 any
access-list 110 deny ip 192.0.2.0 0.0.0.255 any
access-list 110 deny ip 224.0.0.0 31.255.255.255 any
!--- Filter RFC 1918 space.
access-list 110 deny ip 10.0.0.0 0.255.255.255 any
access-list 110 deny ip 172.16.0.0 0.15.255.255 any
access-list 110 deny ip 192.168.0.0 0.0.255.255 any
!--- Deny packets spoofing the school's public addresses
access-list 110 deny ip 198.133.219.0 0.0.0.255 any
!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!--- Module 2:  Explicit Permit
!--- Permit only applications/protocols whose destination
!--- address is part of the infrastructure IP block.
