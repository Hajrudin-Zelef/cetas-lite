---
id: collect-261001-cisco/cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac-4
title: "en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac.md
source_anchor: ""
source_lines: [136, 198]
sha256: 91fbc70204e2b1ae17448bb48a494bf74cfafa06abca1d6986243cf6359c5f9c
---

# en-us-docs-security-asa-asa91-configuration-firewall-nat-objects-html-252f65ac

| xlate per-session { permit \| deny } { tcp \| udp } source_ip [ operator src_port ] destination_ip operator dest_port ciscoasa(config)# xlate per-session deny tcp any4 209.165.201.3 eq 1720 | Creates a permit or deny rule. This rule is placed above the default rules, but below any other manually-created rules. Be sure to create your rules in the order you want them applied. For the source and destination IP addresses, you can configure the following:  The operator matches the port numbers used by the source or destination. The permitted operators are as follows: | 
Monitoring Network Object NAT
To monitor object NAT, enter one of the following commands:
| show nat | Shows NAT statistics, including hits for each NAT rule. | 
| show nat pool | Shows NAT pool statistics, including the addresses and ports allocated, and how many times they were allocated. | 
| show running-config nat | Shows the NAT configuration. Note You cannot view the NAT configuration using the show running-config object command. You cannot reference objects or object groups that have not yet been created in nat commands. To avoid forward or circular references in show command output, the show running-config command shows the object command two times: first, where the IP address(es) are defined; and later, where the nat command is defined. This command output guarantees that objects are defined first, then object groups, and finally NAT. For example: | 
| show xlate | Shows current NAT session information. | 
Configuration Examples for Network Object NAT
This section includes the following configuration examples:
Providing Access to an Inside Web Server (Static NAT)
The following example performs static NAT for an inside web server. The real address is on a private network, so a public address is required. Static NAT is necessary so hosts can initiate traffic to the web server at a fixed address. (See Figure 4-1).
Figure 4-1 Static NAT for an Inside Web Server
Step 1 Create a network object for the internal web server:
Step 2 Define the web server address:
Step 3 Configure static NAT for the object:
NAT for Inside Hosts (Dynamic NAT) and NAT for an Outside Web Server (Static NAT)
The following example configures dynamic NAT for inside users on a private network when they access the outside. Also, when inside users connect to an outside web server, that web server address is translated to an address that appears to be on the inside network. (See Figure 4-2).
Figure 4-2 Dynamic NAT for Inside, Static NAT for Outside Web Server
Step 1 Create a network object for the dynamic NAT pool to which you want to translate the inside addresses:
Step 2 Create a network object for the inside network:
Step 3 Enable dynamic NAT for the inside network:
Step 4 Create a network object for the outside web server:
Step 5 Define the web server address:
Step 6 Configure static NAT for the web server:
Inside Load Balancer with Multiple Mapped Addresses (Static NAT, One-to-Many)
The following example shows an inside load balancer that is translated to multiple IP addresses. When an outside host accesses one of the mapped IP addresses, it is untranslated to the single load balancer address. Depending on the URL requested, it redirects traffic to the correct web server. (See Figure 4-3).
Figure 4-3 Static NAT with One-to-Many for an Inside Load Balancer
Step 1 Create a network object for the addresses to which you want to map the load balancer:
Step 2 Create a network object for the load balancer:
Step 3 Define the load balancer address:
Step 4 Configure static NAT for the load balancer:
Single Address for FTP, HTTP, and SMTP (Static NAT-with-Port-Translation)
The following static NAT-with-port-translation example provides a single address for remote users to access FTP, HTTP, and SMTP. These servers are actually different devices on the real network, but for each server, you can specify static NAT-with-port-translation rules that use the same mapped IP address, but different ports. (See Figure 4-4.)
Figure 4-4 Static NAT-with-Port-Translation
Step 1 Create a network object for the FTP server address:
Step 2 Define the FTP server address, and configure static NAT with identity port translation for the FTP server:
Step 3 Create a network object for the HTTP server address:
Step 4 Define the HTTP server address, and configure static NAT with identity port translation for the HTTP server:
Step 5 Create a network object for the SMTP server address:
Step 6 Define the SMTP server address, and configure static NAT with identity port translation for the SMTP server:
DNS Server on Mapped Interface, Web Server on Real Interface (Static NAT with DNS Modification)
For example, a DNS server is accessible from the outside interface. A server, ftp.cisco.com, is on the inside interface. You configure the ASA to statically translate the ftp.cisco.com real address (10.1.3.14) to a mapped address (209.165.201.10) that is visible on the outside network. (See Figure 4-5.) In this case, you want to enable DNS reply modification on this static rule so that inside users who have access to ftp.cisco.com using the real address receive the real address from the DNS server, and not the mapped address.
When an inside host sends a DNS request for the address of ftp.cisco.com, the DNS server replies with the mapped address (209.165.201.10). The ASA refers to the static rule for the inside server and translates the address inside the DNS reply to 10.1.3.14. If you do not enable DNS reply modification, then the inside host attempts to send traffic to 209.165.201.10 instead of accessing ftp.cisco.com directly.
Figure 4-5 DNS Reply Modification
Step 1 Create a network object for the FTP server address:
Step 2 Define the FTP server address, and configure static NAT with DNS modification:
DNS Server and FTP Server on Mapped Interface, FTP Server is Translated (Static NAT with DNS Modification)
Figure 4-6 shows an FTP server and DNS server on the outside. The ASA has a static translation for the outside server. In this case, when an inside user requests the address for ftp.cisco.com from the DNS server, the DNS server responds with the real address, 209.165.201.10. Because you want inside users to use the mapped address for ftp.cisco.com (10.1.2.56) you need to configure DNS reply modification for the static translation.
Figure 4-6 DNS Reply Modification Using Outside NAT
Step 1 Create a network object for the FTP server address:
Step 2 Define the FTP server address, and configure static NAT with DNS modification:
IPv4 DNS Server and FTP Server on Mapped Interface, IPv6 Host on Real Interface (Static NAT64 with DNS64 Modification)
Figure 4-6 shows an FTP server and DNS server on the outside IPv4 network. The ASA has a static translation for the outside server. In this case, when an inside IPv6 user requests the address for ftp.cisco.com from the DNS server, the DNS server responds with the real address, 209.165.200.225. Because you want inside users to use the mapped address for ftp.cisco.com (2001:DB8::D1A5:C8E1) you need to configure DNS reply modification for the static translation. This example also includes a static NAT translation for the DNS server, and a PAT rule for the inside IPv6 hosts.
Figure 4-7 DNS Reply Modification Using Outside NAT
Step 1 Configure static NAT with DNS modification for the FTP server.
a. Create a network object for the FTP server address.
b. Define the FTP server address, and configure static NAT with DNS modification and, because this is a one-to-one translation, configure the net-to-net method for NAT46.
Step 2 Configure NAT for the DNS server.
a. Create a network object for the DNS server address.
b. Define the DNS server address, and configure static NAT using the net-to-net method.
Step 3 Configure an IPv4 PAT pool for translating the inside IPv6 network.
ciscoasa(config)# object network IPv4_POOL
ciscoasa(config-network-object)# range 203.0.113.1 203.0.113.254
