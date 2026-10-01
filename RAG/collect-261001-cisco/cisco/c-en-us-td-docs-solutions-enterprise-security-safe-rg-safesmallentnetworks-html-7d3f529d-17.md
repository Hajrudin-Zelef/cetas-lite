---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-17
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [584, 617]
sha256: f9d771329cdcfd2e1dd6e8ee014ff75241b9ed794aac2dce8313e63e61bf57e0
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

The following are the guidelines and configuration examples of ACLs controlling access and traffic flows:
Allow Internet access to users residing at all enterprise sites for the allowed ports and protocols. This typically includes HTTP and HTTPS access.
access-list Outbound extended permit tcp 10.0.0.0 255.0.0.0 any eq httpaccess-list Outbound extended permit tcp 10.0.0.0 255.0.0.0 any eq https
Allow users access to DMZ services such as the company's website, E-mail, and domain name resolution (HTTP, HTTPS, SMTP, POP, IMAP, and DNS). Note that the previous entries in the ACL already permit HTTP and HTTPS traffic.
! Allow DNS queries to DNS serveraccess-list Outbound extended permit udp 10.0.0.0 255.0.0.0 host 10.25.34.13 eq domain! Allow SMTP, POP3 and IMAP access to DMZ mail serveraccess-list Outbound extended permit tcp 10.0.0.0 255.0.0.0 host 10.25.34.12 eq smtpaccess-list Outbound extended permit tcp 10.0.0.0 255.0.0.0 host 10.25.34.12 eq pop3access-list Outbound extended permit tcp 10.0.0.0 255.0.0.0 host 10.25.34.12 eq imap4! Apply ACL to inside interfaceaccess-group Outbound in interface inside
Restrict connections initiated from DMZ only to the necessary protocols and sources. This typically includes DNS queries and zone transfer from DNS server, SMTP from E-mail server, HTTP/SSL access from the Cisco IronPort ESA for updates, Sensorbase, etc.
! Allow DNS queries and zone transfer from DNS serveraccess-list DMZ extended permit udp host 10.25.34.13 any eq domainaccess-list DMZ extended permit tcp host 10.25.34.13 any eq domain!! Allow SMTP from Cisco IronPort ESAaccess-list DMZ extended permit tcp host 10.25.34.11 any eq smtp!! Allow update and SensorBase access to Cisco IronPort ESAaccess-list DMZ extended permit tcp host 10.25.34.11 any eq httpaccess-list DMZ extended permit tcp host 10.25.34.11 any eq https!! Apply ACL to DMZ interfaceaccess-group DMZ in interface dmz
Inbound Internet access should be restricted to the public services provided at the DMZ such as SMTP, Web, and DNS. Any connection attempts to internal resources and subnets from the Internet should be blocked. ACLs should be constructed using the servers' global IP addresses.
! Allow DNS queries and zone transfer to DNS serveraccess-list Inbound extended permit udp any host 198.133.219.13 eq domainaccess-list Inbound extended permit tcp any host 198.133.219.13 eq domain!! Allow SMTP to Cisco IronPort ESAaccess-list Inbound extended permit tcp any host 198.133.219.11 eq smtp!! Allow HTTP/HTTPS access to the company's public web portalaccess-list Inbound extended permit tcp any host 198.133.219.10 eq httpaccess-list Inbound extended permit tcp any host 198.133.219.10 eq https!! Apply ACL to outside interfaceaccess-group Inbound in interface outside
Botnet Traffic Filter
The small enterprise network design utilizes the Cisco ASA Botnet Traffic Filter on the Internet Firewall to detect malware that attempts network activity such as sending private data (passwords, credit card numbers, key strokes, or other proprietary data) when the malware starts a connection to a known bad IP address. The Botnet Traffic Filter checks incoming and outgoing connections against a dynamic database of known bad domain names and IP addresses (the blacklist) and then logs or blocks any suspicious activity.
Configuring the Botnet Traffic Filter requires the following steps:
2. Enable Use of the Dynamic Database.
4. Enable Traffic Classification and Actions for the Botnet Traffic Filter.
5. Verify and Monitor Botnet Traffic Filter Operation.
The following sections provides configuration examples for each of these steps.
Configure DNS Server
The Botnet Traffic Filter requires a DNS Server to access Cisco's dynamic database update server and to resolve entries in the static database. The following example illustrates this configuration.
! Enable DNS requests to a DNS Server out the outside interfacedns domain-lookup outside! Specify the DNS Server Group and the DNS Serversdns server-group DefaultDNSname-server 68.238.112.12name-server 68.238.96.12domain-name cisco.com
Enable Use of the Dynamic Database
The Botnet Traffic Filter can receive periodic updates for the dynamic database from the Cisco update server. This database lists thousands of known bad domain names and IP addresses. The following configuration enables database updates and also enables use of the downloaded dynamic database by the adaptive security appliance.
! enable downloading of the dynamic database from the Cisco Update serverdynamic-filter updater-client enable! enable use of the dynamic databasedynamic-filter use-database
Enable DNS Snooping
DNS Snooping enables inspection of DNS packets and enables Botnet Traffic Filter Snooping, which compares the domain name with those in the dynamic or static databases and adds the name and IP address to the DNS reverse lookup cache. This cache is then used by the Botnet Traffic Filter when connections are made to the suspicious address.
It is recommended that DNS snooping is only enabled on interfaces where external DNS requests are going. Enabling DNS snooping on all UDP DNS traffic, including that going to an internal DNS server, creates unnecessary loads on the Cisco ASA. For example, if the DNS server is on the outside interface, you should enable DNS inspection with snooping for all UDP DNS traffic on the outside interface.
The following configuration example illustrates enabling DNS Snooping on the outside interface:
! create a class map to identify the traffic you want to inspect DNSclass-map dynamic-filter-snoop-classmatch port udp eq domain! create a policy map to enable DNS inspection with Botnet Traffic Filtering snooping for the class mappolicy-map dynamic-filter-snoop-policyclass dynamic-filter-snoop-classinspect dns preset_dns_map dynamic-filter-snoop! activate the policy map on the outside interfaceservice-policy dynamic-filter-snoop-policy interface outside
Enable Traffic Classification and Actions for the Botnet Traffic Filter
The Botnet Traffic Filter compares the source and destination IP address in each initial connection packet to the IP addresses in the dynamic database, static database, DNS reverse lookup cache, and DNS host cache, and sends a syslog message or drops any matching traffic. When an address matches, the Cisco ASA sends a syslog message and can optionally be configured to drop the connection. You can enable Botnet Traffic filter on a subset of traffic or for all traffic by enabling an access list to classify traffic.
The following configuration example enables the Botnet Traffic Filter feature on all traffic and additionally enables dropping of connections going to IP addresses with a severity of moderate and higher.
! identify the traffic that you want to monitor or drop.access-list btf-filter-acl extended permit ip any any! enable Botnet Traffic Filter on the outside interface for traffic classified by the btf-filter-acl access listdynamic-filter enable interface outside classify-list btf-filter-acl! enable automatic dropping of traffic with threat level moderate or higherdynamic-filter drop blacklist interface outside action-classify-list btf-filter-acl threat-level range moderate very-high
Botnet Traffic Filter Verification
To monitor and verify the operation of the Botnet Traffic Filter feature, the following commands can be used:
•show dynamic-filter updater-client—Shows information about the updater server, including the server IP address, the next time the adaptive security appliance will connect with the server, and the database version last installed.
