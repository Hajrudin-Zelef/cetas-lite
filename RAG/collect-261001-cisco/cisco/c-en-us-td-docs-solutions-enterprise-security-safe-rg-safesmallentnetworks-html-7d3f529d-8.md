---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-8
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["cost", "distribution", "ethernet", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [251, 285]
sha256: 712a9328b82107751692023bb30a3e15285967c677f8256320d2be5e07c3e116
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

•Ensure that the ESA appliance is both accessible via the public Internet and is the first hop in the E-mail infrastructure. If you allow another MTA to sit at your network's perimeter and handle all external connections, then the ESA appliance will not be able to determine the sender's IP address. The sender's IP address is needed to identify and distinguish senders in the Mail Flow Monitor, to query the SensorBase Reputation Service for the sender's SensorBase Reputation Service Score (SBRS), and to improve the efficacy of the anti-spam and virus outbreak filters features.
•Features like Cisco IronPort Anti-Spam, Virus Outbreak Filters, McAfee Antivirus, and Sophos Anti-Virus require the ESA appliance to be registered in DNS. To that end, create an A record that maps the appliance's hostname to its public IP address and an MX record that maps the public domain to the appliance's hostname. Specify a priority for the MX record to advertise the ESA appliance as the primary (or backup during testing) MTA for the domain. A static address translation entry needs to be defined for the ESA public IP address on the Internet firewall if NAT is configured.
•Add to the Recipient Access Table (RAT) all the local domains for which the ESA appliance will accept mail. Inbound E-mail destined to domains not listed in RAT will be rejected. External E-mail servers connect directly to the ESA appliance to transmit E-mail for the local domains and the ESA appliance relays the mail to the appropriate groupware servers (for example, Exchange™, Groupwise™, and Domino™) via SMTP routes.
•For each private listener, configure the Host Access Table (HAT) to indicate the hosts that will be allowed to send E-mails. The ESA appliance accepts outbound E-mail based on the settings of the HAT table. Configuration includes the definition of Sender Groups associating groups or users, upon which mail policies can be applied. Policies include Mail Flow Policies and Reputation Filtering. Mail Flow Policies are a way of expressing a group of HAT parameters (access rule, followed by rate limit parameters and custom SMTP codes and responses). Reputation Filtering allows the classification of E-mail senders and to restrict E-mail access based on sender's trustworthiness as determined by the IronPort SensorBase Reputation Service.
•Define SMTP routes to direct E-mail to appropriate internal mail servers.
•If an out-of-band (OOB) management network is available, use a separate interface for administration.
A failure on the ESA appliance may cause service outage, therefore a redundant design is recommended. There are multiple ways to implement redundancy:
•IronPort NIC Pairing—Redundancy at the network interface card level by teaming two of the Ethernet interfaces on the ESA appliance. If the primary interface fails, the IP addresses and MAC address are assumed by the secondary.
•Multiple MTAs—Consists of adding a second ESA appliance or MTA with an equal cost secondary MX record.
•Load Balancer—A load balancer such as Cisco ACE Application Control Engine (ACE) load-balances traffic across multiple ESA appliances.
IronPort NIC pairing is the most cost-effective solution (see Figure 9), because it does not require the implementation of multiple ESA appliances and other hardware. It does not, however, provide redundancy in case of chassis failure.
Figure 9 Cisco IronPort ESA NIC pairing
Finally, the Internet firewall should be configured to accommodate traffic to and from the Cisco IronPort ESA deployed at the DMZ. The protocols and ports to be allowed vary depending on the services configured on the appliance. For details, refer to the Cisco IronPort User's Guide at: http://www.ironport.com/support/.
The following are some of the most common services required:
•Outbound SMTP (TCP/25) from ESA to any Internet destination
•Inbound SMTP (TCP/25) to ESA from any Internet destination
•Outbound HTTP (TCP/80) from ESA to downloads.ironport.com and updates.ironport.com
•Outbound SSL (TCP/443) from ESA to updates-static.ironport.com and phonehome.senderbase.org
•Inbound and Outbound DNS (TCP and UDP port 53)
•Inbound IMAP (TCP/143), POP (TCP/110), SMTP (TCP/25) to E-mail server from any internal client
Also remember that if the ESA is managed in-band, appropriate firewall rules need to be configured to allow traffic such as SSH, NTP, and syslog.
For more information on how to configure the Cisco Ironport ESA, see:
•Cisco SAFE Reference Guide- http://www.cisco.com/en/US/docs/solutions/Enterprise/Security/SAFE_RG/SAFE_rg.html
•Cisco IronPort ESA User Guide-http://www.ironport.com/support
Web Security Guidelines
The small enterprise network design implements a Cisco IronPort S Series Web Security Appliance (WSA) to block access to sites with non-business related content and to protect the enterprise from Web-based malware and spyware.
Cisco IronPort WSA's protection relies on two independent services:
•Web Proxy—This provides URL filtering, Web reputation filters, and optionally anti-malware services. The URL filtering capability defines the handling of each Web transaction based on the URL category of the HTTP requests. Leveraging the SensorBase network, the Web reputation filters analyze the Web server behavior and characteristics to identify suspicious activity and protect against URL-based malware. The anti-malware service leverages anti-malware scanning engines such as Webroot and McAfee to monitor for malware activity.
•Layer 4 Traffic Monitoring (L4TM)—Service configured to monitor all Layer 4 traffic for rogue activity and to detect infected clients.
Note The SensorBase network is an extensive network that monitors global E-mail and Web traffic for anomalies, viruses, malware, and other abnormal behavior. The network is composed of the Cisco IronPort appliances, Cisco ASA, and Cisco IPS appliances and modules installed in more than 100,000 organizations worldwide, providing a large and diverse sample of Internet traffic patterns.
As the small enterprise network design assumes a centralized Internet connection, the WSA is implemented at the core/distribution layer of the main site network. This allows the inspection and enforcement of Web access policies to all users residing at any of the enterprise locations. Logically, the WSA sits in the path between Web users and the Internet, as shown in Figure 10.
Figure 10 Cisco IronPort WSA
There are two deployment modes for the Web Proxy service:
•Explicit Forward Proxy—Client applications, such as Web browsers, are aware of the Web Proxy and must be configured to point to the WSA. The Web browsers can be either configured manually or by using Proxy Auto Configuration (PAC) files. The manual configuration does not allow for redundancy, while the use of PAC files allows the definition of multiple WSAs for redundancy and load balancing. If supported by the browser, the Web Proxy Autodiscovery Protocol (WPAD) can be used to automate the deployment of PAC files. WPAD allows the browser to determine the location of the PAC file using DHCP and DNS lookups.
•Transparent Proxy—Client applications are unaware of the Web Proxy and do not have to be configured to connect to the proxy. This mode requires the implementation of a Web Cache Communications Protocol (WCCP)-enabled device or a Layer 4 load balancer in order to intercept and redirect traffic to the WSA. Both deployment options provide for redundancy and load balancing.
