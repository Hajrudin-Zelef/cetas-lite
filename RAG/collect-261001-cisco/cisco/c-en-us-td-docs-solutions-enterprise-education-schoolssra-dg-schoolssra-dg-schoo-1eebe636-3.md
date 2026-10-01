---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636-3
title: "c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "cost", "ethernet", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636.md
source_anchor: ""
source_lines: [98, 136]
sha256: 24d792357d7a4c803afea2e71f965893b48237a15c428c2c45c9a831f1ff8b7c
---

# c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636

The Cisco Ironport C Series E-mail Security Appliance (ESA) deployed at the DMZ is responsible for inspecting E-mails and eliminating threats such as E-mail spam, viruses and worms. The ESA can be described as a firewall and threat monitoring system for Simple Mail Transfer Protocol (SMTP) traffic (TCP port 25). Logically speaking, the ESA acts as a Mail Transfer Agent (MTA) within the E-mail delivery chain, as illustrated in Figure 4-4. Upon reception, E-mails are evaluated using a reputation score mechanism based on the SensorBase network. The SensorBase network is an extensive network that monitors global E-mail and web traffic for anomalies, viruses, malware and other and abnormal behavior. The network is composed of Cisco IronPort appliances, Cisco ASA and IPS appliances installed in more than 100,000 organizations worldwide, providing a large and diverse sample of Internet traffic patterns. By leveraging the SensorBase Network, messages originating from domain names or servers known to be the source of spam or malware, and therefore with a low reputation score, are automatically dropped or quarantined by preconfigured reputation filters. Optionally, the school may choose to implement some of the other functions offered by the ESA appliance, including anti-virus protection with virus outbreak filters and embedded anti-virus engines (Sophos and McAfee), encryption to ensure the confidentiality of messages, and data loss prevention (DLP) for E-mail to detect inappropriate transport of sensitive information.
Figure 4-4 E-mail Delivery Chain
Figure 4-4 shows a logical implementation of a DMZ hosting the E-mail server and ESA appliance. This can be implemented physically by either using a single firewall or two firewalls in sandwich.
Figure 4-5 Common ESA Deployments
There are multiple deployment approaches for the security appliance depending on the number of interfaces used (see Figure 4-5):
•Dual-armed configuration—Two physical interfaces used to serve a public mail listener and a private mail listener, each one configured with a separate logical IP address. The public listener receives E-mail from the Internet and directs messages to the internal E-mail servers; while the private listener receives E-mail from the internal servers and directs messages to the Internet. The public listener interface may connect to the DMZ; while the public listener interface may connect to the inside of the firewall.
•One-armed configuration—A single ESA interface configured with a single IP address and used for both incoming and outgoing E-mail. A public mail listener is configured to receive and relay E-mail on that interface. The best practice is to connect the ESA interface to the DMZ where the E-mail server resides.
For simplicity, the school architecture implements the ESA with a single interface. In addition, using a single interface leaves other data interfaces available for redundancy.
Figure 4-6 illustrates the logical location of the ESA within the E-mail flow chain.
Figure 4-6 Typical Data Flow for Inbound E-mail Traffic
The following steps explain what is taking place in Figure 4-6:
Step 1 Sender sends an E-mail to xyz@domain X.
Step 2 What's the IP address of domain X?
Step 3 It's a.b.c.d (public IP address of ESA).
Step 4 E-mail server sends message to a.b.c.d using SMTP.
Step 5 Firewall permits incoming SMTP connection to the ESA, and translates its public IP address.
Step 6 ESA performs a DNS query on sender domain and checks the received IP address in its reputation database, and drops, quarantines E-mail based on policy.
Step 7 ESA forwards E-mail to preconfigured inbound E-mail server.
Step 8 E-mail server stores E-mail for retrieval by receiver.
Step 9 Receiver retrieves E-mail from server using POP or IMAP.
The Internet firewall should be configured to allow communications to and from the Cisco IronPort ESA. Protocols and ports to be allowed vary depending on the services configured on the appliance. For details, refer to the Cisco IronPort User's Guide at the following URL: http://www.ironport.com/support/
The following are some of the most common services required:
•Outbound SMTP (TCP/25) from ESA to any Internet destination
•Inbound SMTP (TCP/25) to ESA from any Internet destination
•Outbound HTTP (TCP/80) from ESA to downloads.ironport.com and updates.ironport.com
•Outbound SSL (TCP/443) from ESA to updates-static.ironport.com and phonehome.senderbase.org
•Inbound and Outbound DNS (TCP and UDP port 53)
•Inbound IMAP (TCP/143), POP (TCP/110), SMTP (TCP/25) to E-mail server from any internal client
In addition, if the ESA is managed in-band, appropriate firewall rules need to be configured to allow traffic such as SSH, NTP, and syslog.
The Cisco IronPort ESA appliance functions as a SMTP gateway, also known as a mail exchange (MX). The following are the key deployment guidelines:
•Ensure that the ESA appliance is both accessible via the public Internet and is the first hop in the E-mail infrastructure. If you allow another MTA to sit at your network's perimeter and handle all external connections, then the ESA appliance will not be able to determine the sender's IP address. The sender's IP address is needed to identify and distinguish senders in the Mail Flow Monitor, to query the SensorBase Reputation Service for the sender's SensorBase Reputation Service Score (SBRS), and to improve the efficacy of the anti-spam and virus outbreak filters features.
•Features like Cisco IronPort Anti-Spam, Virus Outbreak Filters, McAfee Antivirus and Sophos Anti-Virus require the ESA appliance to be registered in DNS. To that end, create an A record that maps the appliance's hostname to its public IP address, and an MX record that maps the public domain to the appliance's hostname. Specify a priority for the MX record to advertise the ESA appliance as the primary (or backup during testing) MTA for the domain. A static address translation entry needs to be defined for the ESA public IP address on the Internet firewall if NAT is configured.
•Add to the Recipient Access Table (RAT) all the local domains for which the ESA appliance will accept mail. Inbound E-mail destined to domains not listed in RAT will be rejected. External E-mail servers connect directly to the ESA appliance to transmit E-mail for the local domains, and the ESA appliance relays the mail to the appropriate groupware servers (for example, Exchange™, Groupwise™, and Domino™) via SMTP routes.
•For each private listener, configure the Host Access Table (HAT) to indicate the hosts that will be allowed to send E-mails. The ESA appliance accepts outbound E-mail based on the settings of the HAT table. Configuration includes the definition of Sender Groups associating groups or users, upon which mail policies can be applied. Policies include Mail Flow Policies and Reputation Filtering. Mail Flow Policies are a way of expressing a group of HAT parameters (access rule, followed by rate limit parameters and custom SMTP codes and responses). Reputation Filtering allows the classification of E-mail senders and to restrict E-mail access based on sender's trustworthiness as determined by the IronPort SennsorBase Reputation Service.
•Define SMTP routes to direct E-mail to appropriate internal mail servers.
•If an out-of-band (OOB) management network is available, use a separate interface for administration.
A failure on the ESA appliance may cause service outage, therefore a redundant design is recommended. There are multiple ways to implement redundancy:
•IronPort NIC Pairing—Redundancy at the network interface card level by teaming two of the Ethernet interfaces on the ESA appliance. If the primary interface fails, the IP addresses and MAC address are assumed by the secondary.
•Multiple MTAs—Consists in adding a second ESA appliance or MTA with an equal cost secondary MX record.
