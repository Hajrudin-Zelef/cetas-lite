---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-2
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["cost", "distribution", "ethernet", "safeguards"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [44, 84]
sha256: 1bf710caac4e5ccc922bdb227fe2f3eb70fc156a171b7358d8c54797ca7a1e3b
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

•Defense in depth—Multi-layer security embedded throughout the entire infrastructure, endpoints, and applications.
•Global and local intelligence and collaboration—Cloud-based threat information, reputation-based intelligence, and local event and posture information is shared across safeguards for greater visibility and control under a common strategy.
•Service Availability and Resiliency—Multi-level redundancy and device hardening.
•Modularity and Flexibility—Functional modular designs for maximum flexibility and adaptability.
•Strive for Operational Efficiency—Provide the tools and procedures to verify the effectiveness and proper operation of safeguards.
•Regulatory Compliance—Deliver a rich set of security practices and functions commonly required by regulations and standards.
Underlying Network Design
The Cisco SAFE best practices, designs, and configurations presented in this document were integrated and validated using the network design best practices for small businesses and as documented in the Small Enterprise Design Profile. This CVD is a well-designed and validated network architecture that enables the small business to deliver all of the services required for an enhanced business environment. The Small Enterprise Design Profile includes a routing and switching foundation and integrates services such as WAN connectivity, Video Surveillance, Unified Communications, Digital Media Systems, and Mobility.
The Small Enterprise Design Profile consists of a main site and multiple remote sites of various sizes, all interconnected over a Metro Ethernet core. The architecture is illustrated in Figure 2. At the heart of the architecture is a robust routing and switching network. Operating on top of this network are the services used within the enterprise, which often include the majority of the most business-critical applications, such as databases, payroll, accounting, and customer relationship management (CRM) applications. The core of those services are deployed and managed at the main site, allowing the enterprise to reduce the need for separate services to be operated and maintained at the various remote locations. Centralized systems and applications are served by a main site serverfarm.
Figure 2 Small Enterprise Network Design
One of the main features of the Small Enterprise Design Profile is that it provides a loop-free topology, minimizing re-convergence times and eliminating the complexities of certain technologies, such as spanning tree. The architecture supports different designs for the access layer, including the traditional multi-tier access layer design and the more optimal routed access design. The Small Enterprise Design Profile ensures a loop-free topology for both access designs, however the routed access approach provides for up-link load balancing and virtually eliminates any issues associated with spanning tree.
Another important feature of the Small Enterprise Design Profile is that it uses a collapsed network design where the core and distribution layers at each site are collapsed into one layer, allowing the use of a single device (reducing network cost) while maintaining most of the benefits of the traditional three-tier hierarchical model (core, distribution, and access).
The Small Enterprise Design Profile accommodates different levels of redundancy. This includes the use of redundant supervisor engines with Stateful Switchover (SSO) and Non-Stop Forwarding (NSF) capabilities on Cisco Catalyst 4500 switches, Cisco's Virtual Switching Systems (VSS) technology on the Cisco Catalyst 6500 series, and stackwise technology on the Cisco Catalyst 3700 series switches. The last two features allow two or more distribution switches to be combined into a single virtual switch from a management and data forwarding perspective. Redundancy at the link levels is implemented by using EtherChannel technology, where multiple physical links are bundled into a single EtherChannel. EtherChannel significantly simplifies the network response to an individual link failure. If an individual link in EtherChannel fails, the interface does not trigger any network topology changes, thus minimizing impact to network and application performance and improving network convergence.
Small Enterprise Network Security Design
The small enterprise design presented here implements security following the guidelines of the Cisco SAFE security architecture. As a result, a series of network security technologies and products are strategically deployed throughout the network to protect employees and company assets, to guarantee the confidentiality of sensitive data, and to ensure the availability and integrity of systems and data. Safeguards were carefully chosen to mitigate well-known attacks as well as emerging threats.
Common threats to enterprise environments include:
•Service disruption—Disruption to the infrastructure, applications, and other business resources caused by botnets, worms, malware, adware, spyware, viruses, denial-of-service (DoS) attacks, and Layer 2 attacks.
•Network abuse—Use of non-approved applications by employees; peer-to-peer file sharing and instant messaging abuse; and access to non-business-related content.
•Unauthorized access—Intrusions, unauthorized users, escalation of privileges, and unauthorized access to restricted resources.
•Data loss—Theft or leakage of private and confidential data from servers, endpoints, while in transit, or as a result of spyware, malware, key-loggers, viruses, etc.
•Identity theft and fraud—Theft of personnel identity or fraud on servers and end users through phishing and E-mail spam.
As shown in Figure 2, the application of the Cisco SAFE principles to the small enterprise network design follows a defense-in-depth approach, where multiple layers of protection are built into the architecture. The different security tools are combined together for enhanced visibility and control.
The security design for the small enterprise network focuses on the following key areas:
•Network Foundation Protection (NFP)
–Ensuring the availability and integrity of the network infrastructure, protecting the control and management planes.
•Internet Perimeter Protection
–Ensuring safe Internet connectivity and protecting internal resources and users from malware, viruses, and other malicious software.
–Protecting personnel from harmful and inappropriate content.
–Enforcing E-mail and Web browsing policies.
–Ensuring the availability and integrity of centralized applications and systems.
–Protecting the confidentiality and privacy of sensitive data.
•Network Access Security and Control
–Securing the access edges and enforcing authentication and role-based access for users residing at the main site and remote offices.
–Ensuring systems are up-to-date and in compliance with the enterprise network security policies.
–Providing secure, persistent connectivity to all mobile employees on laptops, smartphones, and other mobile platforms. Enforcing encryption, authentication, and role-based access to all mobile users.
–Delivering consistent protection to all mobile employees from viruses, malware, botnets, and other malicious software.
–Ensuring a persistent enforcement of enterprise network security policies to all users and ensuring systems comply with corporate policies and have up-to-date security.
 Note Network management is out of the scope of this document. For best practices in implementing network management, refer to "Chapter 9, Management" of the Cisco SAFE Reference Guide at:
http://www.cisco.com/en/US/docs/solutions/Enterprise/Security/SAFE_RG/chap9.html. 
The following sections discuss the key areas of the small enterprise network security design.
Network Foundation Protection
