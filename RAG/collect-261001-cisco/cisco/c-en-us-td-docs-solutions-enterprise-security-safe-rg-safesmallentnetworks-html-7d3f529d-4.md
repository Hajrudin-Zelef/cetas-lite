---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-4
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent", "cost", "distribution"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [132, 165]
sha256: 273a5cfef54000f00894877a26717d3511a648c12d160da3a8388f1c3b3075b9
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

•E-mail Security—A Cisco IronPort C Series E-mail Security Appliance (ESA) is deployed at the DMZ to inspect incoming and outgoing E-mails and eliminate threats such as E-mail spam, viruses, and worms. The ESA appliance also offers E-mail encryption to ensure the confidentiality of messages and data loss prevention (DLP) to detect the inappropriate transport of sensitive information.
•Web Security—A Cisco IronPort S Series Web Security Appliance (WSA) is deployed at the distribution switches to inspect HTTP and HTTPS traffic bound to the Internet. This system enforces URL filtering policies to block access to websites containing non-business related content or that are known to be the source of spyware, botnets, or other type of malware. The WSA may also be configured to block certain Internet applications such as AOL Messenger, BitTorrent, Skype, etc.
Design guidelines for implementing the security functions are presented below.
Internet Border Router Security Guidelines
The Internet border router provides connectivity to the Internet through one or more Internet service providers. The router act as the first line-of-defense against unauthorized access, DDoS, and other external threats. Access control lists (ACLs), uRPF, and other filtering mechanisms may be implemented for anti-spoofing and to block invalid packets. NetFlow, Syslog, and SNMP may be used to gain visibility on traffic flows, network activity, and system status. In addition, the Internet border router should be secured following the practices explained in Network Foundation Protection. This includes restricting and controlling administrative access, protecting the management and control planes, and securing the dynamic exchange of routing information.
Internet Border Router Deployment provides an example of Internet edge ACL. For more information on how to configure the Internet border router, refer to "Chapter 6, Enterprise Internet Edge" of the Cisco SAFE Reference Guide at: http://www.cisco.com/en/US/docs/solutions/Enterprise/Security/SAFE_RG/chap6.html.
Internet Firewall Guidelines
The Cisco ASA deployed at the Internet perimeter is responsible for:
•Protecting the enterprise internal resources and data from external threats by preventing incoming access from the Internet
•Protecting public resources served by the DMZ by restricting incoming access to the public services and limiting outbound access from DMZ resources out to the Internet
•Controlling users' Internet-bound traffic
To that end, the security appliance is configured to enforce access policies, keep track of connection status, and inspect packet payloads following these guidelines:
•Deny any connection attempts originating from the Internet to internal resources and subnets.
•Allow outbound Internet access for users residing at any of the enterprise locations and for the protocols permitted by the organization's policies, e.g., HTTP and HTTPS.
•Allow outbound Internet SSL access for administrative updates, SensorBase, IPS signature updates, etc.
•Allow users access to DMZ services such as company's website, E-mail, and domain name resolution (HTTP, SMTP, POP, IMAP, and DNS).
•Restrict inbound Internet access to the DMZ for the necessary protocols and servers (HTTP to Web server, SMTP to the mail transfer agent, DNS to DNS server, etc.).
•Restrict connections initiated from DMZ to the only necessary protocols and sources (DNS from DNS server, SMTP from the mail server, and HTTP/SSL from Cisco IronPort ESA).
•Enable stateful inspection for the used protocols to ensure returning traffic is dynamically allowed by the firewall.
•Implement Network Address Translation (NAT) and Port Address Translation (PAT) to shield the internal address space from the Internet.
Figure 4 illustrates the protocols and ports explicitly allowed by the Cisco ASA.
Figure 4 Allowed Protocols and Ports
Note Figure 4 does not include any management traffic destined to the firewall. Whenever available, a dedicated management interface should be used. In case the firewall is managed in-band, identify the protocols and ports required prior to configuring the firewall ACLs.
It is also important to remember that the Cisco ASA should be hardened following the best practices in Network Foundation Protection. This includes restricting and controlling administrative access, securing the dynamic exchange of routing information with MD5 authentication, and enabling firewall network telemetry with SNMP, Syslog, and NetFlow.
In the small enterprise network design, high availability may be achieved by using redundant physical interfaces. This represents the most cost-effective solution for high availability. As an alternative, a pair of firewall appliances could be deployed in stateful failover, as discussed in Internet Firewall Deployment.
Cisco ASA Botnet Traffic Filter
The Cisco ASA Botnet Traffic Filter feature can be enabled to monitor network ports for rogue activity and to prevent infected internal endpoints from sending command and control traffic back to an external host on the Internet. The Botnet Traffic Filter on the ASA provides reputation-based control for an IP address or domain name, similar to the control that Cisco IronPort SensorBase provides for E-mail and Web servers.
The Cisco Botnet Traffic Filter is integrated into all Cisco ASA appliances and inspects traffic traversing the appliance to detect rogue traffic in the network. When internal clients are infected with malware and attempt to phone home to an external host on the Internet, the Botnet Traffic Filter alerts the system administrator through the regular logging process and can be automatically blocked. This is an effective way to combat botnets and other malware that share the same phone-home communications pattern.
The Botnet Traffic Filter monitors all ports and performs a real-time lookup in its database of known botnet IP addresses and domain names. Based on this investigation, the Botnet Traffic Filter determines whether a connection attempt is benign and should be allowed or is a risk and should be blocked.
The Cisco ASA Botnet Traffic Filter has three main components:
•Dynamic and administrator blacklist data—The Botnet Traffic Filter uses a database of malicious domain names and IP addresses that is provided by Cisco Security Intelligence Operations. This database is maintained by Cisco Security Intelligence Operations and is downloaded dynamically from an update server on the SensorBase network. Administrators can also configure their own local blacklists and whitelists.
•Traffic classification and reporting—Botnet Traffic Filter traffic classification is configured through the dynamic-filter command on the ASA. The dynamic filter compares the source and destination addresses of traffic against the IP addresses that have been discovered for the various lists available (dynamic black, local white, local black) and logs and reports the hits against these lists accordingly.
•Domain Name System (DNS) snooping—To map IP addresses to domain names that are contained in the dynamic database or local lists, the Botnet Traffic Filter uses DNS snooping in conjunction with DNS inspection. Dynamic Filter DNS snooping looks at User Datagram Protocol (UDP) DNS replies and builds a DNS reverse cache (DNSRC), which maps the IP addresses in those replies to the domain names they match. DNS snooping is configured via the Modular Policy Framework (MPF) policies.
The Botnet Traffic Filter uses two databases for known addresses. Both databases can be used together or the dynamic database can be disabled and the static database can be used alone. When using the dynamic database, the Botnet Traffic Filter receives periodic updates from the Cisco update server on the Cisco IronPort SensorBase network. This database lists thousands of known bad domain names and IP addresses.
