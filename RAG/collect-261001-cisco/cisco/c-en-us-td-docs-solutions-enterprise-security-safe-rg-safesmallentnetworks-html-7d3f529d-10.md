---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-10
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [319, 345]
sha256: d93a5be88fabdc1d28d676e4e67ff5ed5a2c24b68aeef7e6fa537a03c59db3e8
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

Small enterprise networks typically include a serverfarm at the main site that hosts the systems that serve business applications and store the data accessible to internal users. The infrastructure supporting it may include application servers, the storage media, routers, switches, load balancers, off-loaders, application acceleration devices, and other systems. In addition, they may also host foundational services as part of the enterprise network such as identity and security services, unified communication services, mobility services, video services, partner applications, and other services.
Depending on the size of the enterprise network, the serverfarm may be constructed following different design models. Figure 12 illustrates a collapsed design, and the less common for small enterprise networks, multi-tier design. In the collapsed design all services are hosted in a shared physical serverfarm and high availability is achieved by using redundant processors and interfaces. Large enterprises may implement a more scalable multi-tier design data center with chassis redundancy.
Figure 12 Serverfarm Designs
Independently from the design model adopted by the enterprise, the following are the primary security guidelines for the serverfarm design:
•Network Foundation Protection—All infrastructure equipment should be protected following the best practices in Network Foundation Protection. This includes restricting and controlling administrative access, protecting the management and control planes, and securing the switching and routing planes.
•Firewall—A stateful firewall may be deployed to limit access to only the necessary applications and services and for the intended users. The firewall should be configured to control and inspect both traffic entering and leaving the serverfarm segments. The firewall may also be leveraged to ensure the appropriate segregation between application layers or groups. In addition, the firewall's deep packet inspection may be used to mitigate DoS attacks and enforce protocol compliance.
•Intrusion Prevention—An IPS module on the Cisco ASA or a separate IPS appliance may be implemented for enhanced threat detection and mitigation. The IPS is responsible for identifying and blocking anomalous traffic and packets recognized as well-known attacks. The Cisco IPS may be configured either in inline or promiscuous mode. When deployed in inline mode, the Cisco IPS is placed in the traffic path and is capable of stopping malicious traffic before it reaches the intended target.
•Service Isolation—Services and applications serving different group of users or under different security requirements should be properly isolated. Isolation helps prevent data leakage and contain possible compromises from spreading across different serverfarm groups. Logical isolation may be achieved by separating applications and services in different VLANs and by assigning them into different firewall interfaces (physical or logical). This is illustrated in Figure 13.
•Switch Security—Private VLANs, port security, storm control, and other switch security features may be leveraged to mitigate spoofing, man-in-the-middle, denial-of-service, and other network-based attacks directed to the serverfarm applications and the switching infrastructure.
•Endpoint Protection—Servers residing at the different layers should be protected with host-based IPS or other endpoint security software.
Figure 13 Service Isolation
SSL termination and inspection, Web Application Firewall (WAF), Application Control Engine (ACE), and other solutions may be leveraged to complement the guidelines described above. For a more detailed discussion of serverfarm security, refer to "Chapter 4, Intranet Data Center" of the Cisco SAFE Reference Guide at: http://www.cisco.com/en/US/docs/solutions/Enterprise/Security/SAFE_RG/chap4.html.
Network Access Security and Control
Some of the most vulnerable points of the network are the access edges where users connect to the network. With the proliferation of wireless networks, increased use of laptops and smart mobile devices, the enterprise cannot simply rely on physical controls hoping to prevent unauthorized systems from being plugged into the ports of the access switches. Protection should rather be embedded into the network infrastructure, leveraging the native security features available in switches, routers, and WLAN systems. Furthermore, the network infrastructure should also provide dynamic identity or role-based access controls for all systems attempting to gain access.
Implementing role-based access controls for users and devices helps reduce the potential loss of sensitive information by enabling enterprises to verify a user or device identity, privilege level, and security policy compliance before granting network access. Security policy compliance could consist of requiring antivirus software, OS updates, or patches. Unauthorized or noncompliant devices can be placed in a quarantine area where remediation can occur prior to gaining access to the network.
Access security and control is implemented in the small enterprise network design by leveraging the following technologies:
•Catalyst Integrated Security Features (CISF)—Wired
•Cisco Unified Wireless Network (CUWN) Integrated Security Features—Wireless
•Cisco NAC Appliance—Wired and wireless
•Cisco Identity-Based Network Networking Services (IBNS)—Wired and wireless
Catalyst Integrated Security Features
Catalyst Integrated Security Features (CISF) is a set of native security features available on Cisco Catalyst Switches and designed to protect the access infrastructure and users from spoofing, man-in-the-middle, DoS, and other network-based attacks. CISF includes features such as private VLANs, port security, DHCP snooping, IP Source Guard, secure Address Resolution Protocol (ARP) detection, and Dynamic ARP Inspection (DAI). CISF features are considered to be part of a security baseline and should be deployed on all access ports.
•Port Security—Mitigates MAC flooding and other Layer 2 CAM overflow attacks by restricting the MAC addresses that are allowed to send traffic on a particular port. After Port Security is enabled on a port, only packets with a permitted source MAC address are allowed to pass through the port. A permitted MAC address is referred to as a secure MAC address.
•DHCP snooping—Inspects and filters DHCP messages on a port to ensure DHCP server messages come only from a trusted interface. Additionally, it builds and maintains a DHCP snooping binding table that contains the MAC address, IP address, lease time, binding type, VLAN number, and interface information corresponding to the local untrusted interfaces of a switch. This binding table is used by the other CISF features.
•IP Source Guard—Restricts IP traffic on a port based on DHCP or static IP address MAC bindings to prevent IP spoofing attacks. IP address bindings are validated using information in the DHCP Snooping binding table.
•Dynamic ARP inspection—Validates that the source MAC and IP address in an ARP packet received on an untrusted interface matches the source MAC and IP address registered on that interface (using the DHCP snooping binding table) to prevent ARP spoofing and MITM attacks.
•ARP rate limiting—Where an excessive rate of ARP request (which must be processed by network hosts CPUs) and the switch responds with access restriction if this rate is exceeded.
