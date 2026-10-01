---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-3
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["disclosure", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [85, 131]
sha256: 478df7cfdb4deea5629fba8578b51c70e55afd1049022d0796f9868c0f9c5eea
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

Small enterprise networks are built with routers, switches, and other network devices that keep the applications and services running. Therefore, properly securing these network devices is critical for continued operation.
The network infrastructure of a small enterprise can be protected by implementing the Cisco SAFE best practices for the following areas:
–Restrict management device access to authorized parties and for the authorized ports and protocols.
–Enforce Authentication, Authorization, and Accounting (AAA) with Terminal Access Controller Access Control System (TACACS+) or Remote Authentication Dial-In User Service (RADIUS) to authenticate access, authorize actions, and log all administrative access.
–Display legal notification banners.
–Ensure confidentiality by using secure protocols like Secure Shell (SSH) and HTTPS.
–Enforce idle and session time-outs and disable unused access lines.
–Restrict routing protocol membership by enabling Message-Digest 5 (MD5) neighbor authentication and disabling default interface membership.
–Enforce route filters to ensure that only legitimate networks are advertised and networks that are not supposed to be propagated are never advertised.
–Log status changes of neighbor sessions to identify connectivity problems and DoS attempts on routers.
•Device resiliency and survivability
–Disable unnecessary services and implement control plane policing (CoPP).
–Enable traffic storm control.
–Implement topological, system, and module redundancy for the resiliency and survivability of routers and switches and to ensure network availability.
–Keep local device statistics.
–Enable Network Time Protocol (NTP) time synchronization.
–Collect system status and event information with Simple Network Management Protocol (SNMP), Syslog, and TACACS+/RADIUS accounting.
–Monitor CPU and memory usage on critical systems.
–Enable NetFlow to monitor traffic patterns and flows.
–Implement access edge filtering.
–Enforce IP spoofing protection with access control lists (ACLs), Unicast Reverse Path Forwarding (uRPF), and IP Source Guard.
–Implement a hierarchical design, segmenting the LAN into multiple IP subnets or virtual LANs (VLANs) to reduce the size of broadcast domains.
–Protect the Spanning Tree Protocol (STP) domain with BPDU Guard and STP Root Guard.
–Use per-VLAN Spanning Tree to reduce the scope of possible damage.
–Disable VLAN dynamic trunk negotiation on user ports.
–Disable unused ports and put them into an unused VLAN.
–Implement Catalyst Infrastructure Security Features (CISF) including port security, dynamic ARP inspection, and DHCP snooping.
–Use a dedicated VLAN ID for all trunk ports.
–Explicitly configure trunking on infrastructure ports.
–Use all tagged mode for the native VLAN on trunks and drop untagged frames.
–Ensure the secure management of all devices and hosts within the enterprise network.
–Authenticate, authorize, and keep record of all administrative access.
–If possible, implement a separate out-of-band (OOB) management network (hardware or VLAN based) to manage systems at the main site.
–Secure the OOB by enforcing access controls, using dedicated management interfaces or virtual routing and forwarding (VRF) tables.
–Provide secure in-band management access for systems residing at the remote sites by deploying firewalls and ACLs to enforce access controls, using Network Address Translation (NAT) to hide management addresses and using secure protocols like SSH and HTTPS.
–Ensure time synchronization by using NTP.
–Secure servers and other endpoint with endpoint protection software and operating system (OS) hardening best practices.
For more detailed information on the NFP best practices, refer to "Chapter 2, Network Foundation Protection" of the Cisco SAFE Reference Guide at: http://www.cisco.com/en/US/docs/solutions/Enterprise/Security/SAFE_RG/chap2.html.
Internet Perimeter Protection
The small enterprise network design assumes the existence of a centralized Internet connection at the headquarters or main site serving users at all locations. Common services include E-mail and Web browsing for employees, the hosting of a company's website accessible to clients and partners over the Internet, and secure remote access for mobile users and remote workers. Other services may also be provided using the same infrastructure.
The network infrastructure that provides Internet connectivity is defined as the Internet perimeter, illustrated in Figure 3.
Figure 3 Internet Perimeter
The primary functions of the Internet perimeter is to allow for safe and secure access for users residing at all locations and to provide public services without compromising the confidentiality, integrity, and availability of the company's resources and data. To that end, the Internet perimeter incorporates the following security functions:
•Internet Border Router—The Internet border router is the Internet gateway responsible for routing traffic between the enterprise network and the Internet. The Internet border router may be administered by the company's personnel or may be managed by an Internet service provider (ISP). The router provides the first line of protection against external threats and should be hardened following the Network Foundation Protection (NFP) best practices.
•Internet Firewall—A Cisco ASA provides stateful access control and deep packet inspection to protect company resources and data from unauthorized access and disclosure. The security appliance is configured to prevent incoming access from the Internet, to protect the enterprise Internet public services, and to control user traffic bound to the Internet. In addition, the Cisco ASA Botnet Traffic Filter feature can be enabled to defend the enterprise against botnet threats. Once enabled, the Botnet Traffic Filter feature monitors network traffic across all ports and protocols for rogue activity and to prevent infected internal endpoints from sending command and control traffic back to external hosts on the Internet. The security appliance may also provide secure remote access to employees with the Cisco AnyConnect Secure Mobility client.
•Intrusion Prevention—An Advanced Inspection and Prevention Security Service Module (AIP SSM) on the Cisco ASA or a separate IPS appliance can be implemented for enhanced threat detection and mitigation. The IPS module or appliance is responsible for identifying and blocking anomalous traffic and malicious packets recognized as well-known attacks. IPS can be deployed either in inline or promiscuous mode. The module or appliance may be configured to participate in Cisco IPS Global Correlation, allowing the IPS to gain visibility on global threats as they emerge in the Internet and to quickly react to contain them. IPS may also be configured to help block certain Internet applications, such as AOL Messenger, BitTorrent, Skype, etc.
•Public Services DMZ—The company's Internet website, mail server, and other public-facing services may be placed on a demilitarized zone (DMZ) for security and control purposes. The DMZ acts as a middle stage between the Internet and the company's private resources, preventing external users from directly accessing any internal servers and data. The Internet firewall is responsible for restricting incoming access to the public services and by limiting outbound access from DMZ resources out to the Internet. Systems residing on the DMZ are hardened with endpoint protection software and operating system (OS) hardening best practices.
