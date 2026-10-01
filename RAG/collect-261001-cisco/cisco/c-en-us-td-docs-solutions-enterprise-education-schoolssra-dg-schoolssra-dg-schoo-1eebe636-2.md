---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636-2
title: "c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent", "cost", "disclosure", "distribution"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636.md
source_anchor: ""
source_lines: [70, 97]
sha256: 60bc448b59206c670f2601224a2777c2a704f931e0d374379c6cac2451b470d5
---

# c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636

The primary functions of the Internet perimeter is to allow for safe and secure access to students, staff, and faculty, and to provide public services without compromising the confidentiality, integrity, and availability of school resources and data. To that end, the Internet perimeter incorporates the following security functions:
•Internet Border Router—This is the Internet gateway responsible for routing traffic between the school and the Internet. The Internet border router may be administered by school personnel or may be managed by the Internet service provider (ISP). The router provides the first line of protection against external threats and should be hardened using the Network Foundation Protection (NFP) best practices.
•Internet Firewall—A Cisco Adaptive Security Appliance provides Stateful access control and deep packet inspection to protect the school resources and data from unauthorized access and disclosure. The security appliance is configured to prevent incoming access from the Internet, to protect the school web portal and other Internet public services, and to control student, staff and faculty traffic bound to the Internet. The security appliance may also implement an Advanced Inspection and Prevention Security Services Module (AIP SSM) for enhanced threat detection and mitigation. This IPS module may be configured either in inline or promiscuous mode. The security appliance may also provide secure remote access to faculty, staff, and students in the form of IPSec or SSL VPN.
•Public Services DMZ—The Internet school web portal, mail server, and other public-facing services may be placed on a demilitarized zone (DMZ) for security and control purposes. The DMZ acts as a middle stage between the Internet and school's private resources, preventing external users from directly accessing any internal servers and data. The Internet firewall is responsible for restricting incoming access to the public services and by limiting outbound access from DMZ resources out to the Internet. Systems residing on the DMZ are hardened with endpoint protection software (i.e., Cisco Security Agent) and operating system (OS) hardening best practices.
•E-mail Security—A Cisco Ironport C Series E-mail Security Appliance (ESA) is deployed at the DMZ to inspect incoming and outgoing E-mails and eliminate threats such as E-mail spam, viruses, and worms. The ESA appliance also offers E-mail encryption to ensure the confidentiality of messages, and data loss prevention (DLP) to detect the inappropriate transport of sensitive information.
•Web Security—A Cisco IronPort S Series Web Security Appliance (WSA) is deployed at the distribution switches to inspect HTTP and HTTPS traffic bound to the Internet. This system enforces URL-filtering policies to block access to websites containing content that may be harmful or inappropriate for minors or that are known to be the source of spyware, botnets, or other type of malware. The WSA is also responsible for the monitoring of Layer-4 traffic for rogue activity and infected systems.
The following subsections describe the design guidelines for implementing the security functions.
Note For implementation details on Remote Access VPN, IPS, CSA, and Internet border router hardening, refer to the Cisco SAFE Reference Guide at the following URL: http://www.cisco.com/en/US/docs/solutions/Enterprise/Security/SAFE_RG/SAFE_rg.html. Firewall and web security configurations can be found in Chapter 11, "District Office Design."
Internet Border Router Guidelines
The Internet border router provides connectivity to the Internet via one or more Internet service providers. The router act as the first line of defense against unauthorized access, DDoS, and other external threats. Access control lists (ACLs), uRPF, and other filtering mechanisms may be implemented for anti-spoofing and to block invalid packets. NetFlow, Syslog, and SNMP may be used to gain visibility on traffic flows, network activity and system status. In addition, the Internet border router should be secured.This includes restricting and controlling administrative access, protecting the management and control planes, and securing the dynamic exchange of routing information.
Chapter 11, "District Office Design"provides an example of Internet edge ACL. For more information on how to configure the Internet border router, refer to "Internet Edge" chapter of the Cisco SAFE Reference Guide at the following URL: http://www.cisco.com/en/US/docs/solutions/Enterprise/Security/SAFE_RG/SAFE_rg.html.
Internet Firewall Guidelines
The Cisco Adaptive Security Appliance (Cisco ASA) deployed at the Internet perimeter is responsible for protecting the school's internal resources and data from external threats by preventing incoming access from the Internet; protecting public resources served by the DMZ by restricting incoming access to the public services and by limiting outbound access from DMZ resources out to the Internet; and controlling user's Internet-bound traffic.
To that end, the security appliance is configured to enforce access policies, keep track of connection status, and inspect packet payloads following these guidelines:
•Deny any connection attempts originating from the Internet to internal resources and subnets.
•Allow outbound Internet HTTP/HTTPS access for students, staff and faculty residing at any of the school premises.
•Allow outbound Internet SSL access for administrative updates, SensorBase, IPS signature updates, etc.
•Allow students, staff and faculty access to DMZ services such as school web portal, E-mail, and domain name resolution (HTTP, SMTP, POP, IMAP, and DNS).
•Restrict inbound Internet access to the DMZ for the necessary protocols and servers (HTTP to Web server, SMTP to Mail Transfer Agent, DNS to DNS server, etc.).
•Restrict connections initiated from DMZ to the only necessary protocols and sources (DNS from DNS server and mail server, SMTP from mail server, SSL for Cisco IronPort ESA).
•Enable stateful inspection for the used protocols to ensure returning traffic is dynamically allowed by the firewall.
•Implement Network Address Translation (NAT) and Port Address Translation (PAT) to shield the internal address space from the Internet.
Figure 4-3 illustrates the protocols and ports explicitly allowed by the Cisco ASA.
Figure 4-3 Allowed Protocols and Ports
Note Allowed Protocols and Ports does not include any management traffic destined to the firewall. Whenever available, a dedicated management interface should be used. In case the firewall is managed in-band, identify the protocols and ports required prior to configuring the firewall ACLs.
In addition, the Cisco ASA should be hardened following the NFP best practices. This includes restricting and controlling administrative access, securing the dynamic exchange of routing information with MD5 authentication, and enabling firewall network telemetry with SNMP, syslog, and NetFlow.
In the school design, higher availability is achieved by using redundant physical interfaces. This represents the most cost-effective solution for high availability. As an alternative, a pair of firewall appliances could be deployed in stateful failover, as discussed in Chapter 11, "District Office Design."
E-mail Security Guidelines
