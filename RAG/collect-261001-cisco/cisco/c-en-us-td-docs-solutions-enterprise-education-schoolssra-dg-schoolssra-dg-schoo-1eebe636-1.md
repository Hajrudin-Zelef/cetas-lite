---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636-1
title: "c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636.md
source_anchor: ""
source_lines: [1, 69]
sha256: 4bb48f411d5076bdbf5d9a42ecfdec9f37a6f13cd56b79bd01e037fdab1ac55d
---

# c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636

Security Design
Introduction
The Cisco Service Ready Architecture (SRA) for Schools solution is designed with security to provide a safe online environment for teaching and learning. Following the proven guidelines of the Cisco SAFE Security architecture, a series of security technologies and products are strategically deployed throughout the solution to protect minors from harmful and inappropriate content, to guarantee the confidentiality of student, staff and faculty private data, and to ensure the availability and integrity of the systems and data.
Protecting the infrastructure and keeping students and staff safe requires the implementation of security controls capable of mitigating both well-known and new forms of threats. Common threats to school environments include:
•Service disruption—Disruption of the administrative infrastructure and learning resources such as computer labs caused by botnets, worms, malware, adware, spyware, viruses, DoS attacks.
•Harmful or inappropriate content—Pornography, adult, aggressive, offensive and other type of content that could put the physical and psychological well being of minors at risk.
•Network abuse—Peer-to-peer file sharing and instant messaging abuse, use of non-approved applications by students, staff, and faculty.
•Unauthorized access—Intrusions, unauthorized users, escalation of privileges, and unauthorized access to learning and administrative resources.
•Data loss—Theft or leakage of student, staff and faculty private data from servers, endpoints, and while in transit, or as a result of spyware, malware, key-loggers, viruses, etc.
The solution design follows a defense-in-depth approach, whereby multiple layers of protection are built into the architecture. Different security products and technologies are combined together for enhanced security visibility and control. Figure 4-1 illustrates the security design and its product positioning.
Figure 4-1 SRA Security Design
The security design focuses on the following key areas:
•Network Foundation Protection (NFP)— Ensuring the availability and integrity of the network infrastructure, protecting the control and management planes.
•Internet Perimeter Protection— Ensuring safe Internet connectivity, and protecting internal resources and users from malware, viruses, and other malicious software. Protecting students and staff from harmful and inappropriate content. Enforcing E-mail and web browsing policies.
•Network Access Security and Control—Securing the access edges. Enforcing authentication and role-based access for students, staff and faculty residing at school sites and district office. Ensuring systems are up-to-date and in compliance with the school's network security policies.
•Network Endpoint Protection—Protecting students, staff and faculty from harmful and inappropriate content. Enforcing E-mail and web browsing policies.
The design guidelines and best practices for each focus area are discussed next. For more detailed information, refer to the Cisco SAFE Reference Guide at the following URL:
http://www.cisco.com/en/US/docs/solutions/Enterprise/Security/SAFE_RG/SAFE_rg.html
Network Foundation Protection
School networks are built with routers, switches, and other network devices that keep the applications and services running. Therefore, properly securing these network devices is critical for continued operation.
The SRA solution protects the network infrastructure by implementing the Cisco SAFE best practices for the following areas:
•Infrastructure device access
–Restrict management device access to authorized parties and for the authorized ports and protocols.
–Enforce Authentication, Authorization and Accounting (AAA) with TACACS+ or RADIUS to authenticate access.
–Authorize actions and log all administrative access.
–Display legal notification banners.
–Ensure confidentiality by using secure protocols like SSH and HTTPS.
–Enforce idle and session timeouts.
–Disable unused access lines.
•Routing infrastructure
–Restrict routing protocol membership by enabling MD5 neighbor authentication and disabling default interface membership.
–Enforce route filters to ensure that only legitimate networks are advertised, and networks that are not supposed to be propagated are never advertised.
–Log status changes of neighbor sessions to identify connectivity problems and DoS attempts on routers.
•Device resiliency and survivability
–Disable unnecessary services, implement control plane policing (CoPP).
–Enable traffic storm control.
–Implement topological, system and module redundancy for the resiliency and survivability of routers and switches and to ensure network availability.
–Keep local device statistics.
•Network telemetry
–Enable NTP time synchronization.
–Collect system status and event information with SNMP, Syslog, TACACS+/RADIUS accounting.
–Monitor CPU and memory usage on critical systems.
•Network policy enforcement
–Implement access edge filtering.
–Enforce IP spoofing protection with access control lists (ACLs), Unicast Reverse Path Forwarding (uRPF) and IP Source Guard.
•Switching infrastructure
–Implement a hierarchical design, segmenting the LAN into multiple IP subnets or VLANS to reduce the size of broadcast domains.
–Protect the Spanning Tree Protocol (STP) domain with BPDU Guard, STP Root Guard.
–Use Per-VLAN Spanning Tree to reduce the scope of possible damage.
–Disable VLAN dynamic trunk negotiation on user ports.
–Disable unused ports and put them into an unused VLAN.
–Enable Traffic Storm Control.
–Implement Catalyst Infrastructure Security Features (CISF) including port security, Dynamic ARP Inspection, and DHCP snooping.
–Use a dedicated VLAN ID for all trunk ports.
–Explicitly configure trunking on infrastructure ports.
–Use all tagged mode for the native VLAN on trunks and drop untagged frames.
•Network management
–Ensure the secure management of all devices and hosts within the school network architecture.
–Authenticate, authorize and keep record of all administrative access.
–If possible, implement a separate out-of-band (OOB) management network (hardware or VLAN-based) to manage systems local at the District Office.
–Secure the OOB by enforcing access controls, using dedicated management interfaces or VRFs.
–Provide secure in-band management access for systems residing at the school sites by deploying firewalls and ACLs to enforce access controls, using Network Address Translation (NAT) to hide management addresses, and use secure protocols like SSH and HTTPS.
–Ensure time synchronization by using NTP. Secure servers and other endpoint with endpoint protection software and operating system (OS) hardening best practices.
Configurations are shown in the design chapters. For more detailed information on the NFP best practices, refer to "Chapter 2, Network Foundation Protection" of the Cisco SAFE Reference Guide at the following URL:
http://www.cisco.com/en/US/docs/solutions/Enterprise/Security/SAFE_RG/chap2.html
Internet Perimeter Protection
The school architecture assumes the existence of a centralized Internet connection at the district office, serving students, staff and faculty residing at all school premises. Common services typically provided include E-mail for staff and faculty, Internet browsing for everyone, and a school web portal accessible over the Internet. Other services may also be provided using the same infrastructure.
The network infrastructure that provides Internet connectivity is defined as the Internet perimeter, illustrated in Figure 4-2.
Figure 4-2 Internet Perimeter
