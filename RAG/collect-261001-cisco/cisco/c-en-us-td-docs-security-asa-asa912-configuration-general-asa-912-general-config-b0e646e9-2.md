---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-b0e646e9-2
title: "c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-b0e646e9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-b0e646e9.md
source_anchor: ""
source_lines: [113, 215]
sha256: 009ea72d8ac66f49dbb33ae8ee5ecf6ce217a964d980721f8f514062616b260e
---

# c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-b0e646e9

                                 Syslog message class (equivalent to a functional area)
You customize these criteria by creating a message list that you can specify when you set the output destination. Alternatively, you can configure the ASA to send a particular message class to each type of output destination independently of the message list.
Syslog Message Classes
You can use syslog message classes in two ways:
-  
                                    			 
                                    Specify an output location for an entire category of syslog messages. Use the logging class command.
-  
                                    			 
                                    Create a message list that specifies the message class. Use the logging list command.
The syslog message class provides a method of categorizing syslog messages by type, equivalent to a feature or function of the device. For example, the rip class denotes RIP routing.
All syslog messages in a particular class share the same initial three digits in their syslog message ID numbers. For example, all syslog message IDs that begin with the digits 611 are associated with the vpnc (VPN client) class. Syslog messages associated with the VPN client feature range from 611101 to 611323.
In addition, most of the ISAKMP syslog messages have a common set of prepended objects to help identify the tunnel. These objects precede the descriptive text of a syslog message when available. If the object is not known at the time that the syslog message is generated, the specific heading = value combination does not appear.
The objects are prefixed as follows:
Group = groupname, Username = user, IP = IP_address
Where the group is the tunnel-group, the username is the username from the local database or AAA server, and the IP address is the public IP address of the remote access client or Layer 2 peer.
The following table lists the message classes and the range of message IDs in each class.
| Table 2. Syslog Message Classes and Associated Message ID Numbers |  |  | 
|---|---|---|
| Class | Definition | Syslog Message ID Numbers | 
|---|---|---|
| auth | User Authentication | 109, 113 | 
| — | Access Lists | 106 | 
| — | Application Firewall | 415 | 
| bridge | Transparent Firewall | 110, 220 | 
| ca | PKI Certification Authority | 717 | 
| citrix | Citrix Client | 723 | 
| — | Clustering | 747 | 
| — | Card Management | 323 | 
| config | Command Interface | 111, 112, 208, 308 | 
| csd | Secure Desktop | 724 | 
| cts | Cisco TrustSec | 776 | 
| dap | Dynamic Access Policies | 734 | 
| eap, eapoudp | EAP or EAPoUDP for Network Admission Control | 333, 334 | 
| eigrp | EIGRP Routing | 336 | 
|  | E-mail Proxy | 719 | 
| — | Environment Monitoring | 735 | 
| ha | Failover | 101, 102, 103, 104, 105, 210, 311, 709 | 
| — | Identity-based Firewall | 746 | 
| ids | Intrusion Detection System | 400, 733 | 
| — | IKEv2 Toolkit | 750, 751, 752 | 
| ip | IP Stack | 209, 215, 313, 317, 408 | 
| ipaa | IP Address Assignment | 735 | 
| ips | Intrusion Protection System | 400, 401, 420 | 
| — | IPv6 | 325 | 
| — | Botnet traffic filtering. | 338 | 
| — | Licensing | 444 | 
| mdm-proxy | MDM Proxy | 802 | 
| nac | Network Admission Control | 731, 732 | 
| nacpolicy | NAC Policy | 731 | 
| nacsettings | NAC Settings to apply NAC Policy | 732 | 
| — | Network Access Point | 713 | 
| np | Network Processor | 319 | 
| — | NP SSL | 725 | 
| ospf | OSPF Routing | 318, 409, 503, 613 | 
| — | Password Encryption | 742 | 
| — | Phone Proxy | 337 | 
| rip | RIP Routing | 107, 312 | 
| rm | Resource Manager | 321 | 
| — | Smart Call Home | 120 | 
| session | User Session | 106, 108, 201, 202, 204, 302, 303, 304, 305, 314, 405, 406, 407, 500, 502, 607, 608, 609, 616, 620, 703, 710 | 
| snmp | SNMP | 212 | 
| — | ScanSafe | 775 | 
| ssl | SSL Stack | 725 | 
| svc | SSL VPN Client | 722 | 
| sys | System | 199, 211, 214, 216, 306, 307, 315, 414, 604, 605, 606, 610, 612, 614, 615,701, 711, 741 | 
| — | Threat Detection | 733 | 
| tre | Transactional Rule Engine | 780 | 
| — | UC-IME | 339 | 
| tag-switching | Service Tag Switching | 779 | 
| vm | VLAN Mapping | 730 | 
| vpdn | PPTP and L2TP Sessions | 213, 403, 603 | 
| vpn | IKE and IPsec | 316, 320, 402, 404, 501, 602, 702, 713, 714, 715 | 
| vpnc | VPN Client | 611 | 
| vpnfo | VPN Failover | 720 | 
| vpnlb | VPN Load Balancing | 718 | 
| — | VXLAN | 778 | 
| webfo | WebVPN Failover | 721 | 
| webvpn | WebVPN and AnyConnect Client | 716 | 
| — | NAT and PAT | 305 | 
Custom Message Lists
- 
                                    		  
                                    Severity level
- 
                                    		  
                                    Message IDs
- 
                                    		  
                                    Ranges of syslog message IDs
- 
                                    		  
                                    Message class.
For example, you can use message lists to do the following:
-  
                                 		  
                                 Select syslog messages with the severity levels of 1 and 2 and send them to one or more e-mail addresses.
-  
                                 		  
                                 Select all syslog messages associated with a message class (such as ha) and save them to the internal buffer.
A message list can include multiple criteria for selecting messages. However, you must add each message selection criterion with a new command entry. It is possible to create a message list that includes overlapping message selection criteria. If two criteria in a message list select the same message, the message is logged only once.
Clustering
Syslog messages are an invaluable tool for accounting, monitoring, and troubleshooting in a clustering environment. Each ASA unit in the cluster (up to eight units are allowed) generates syslog messages independently; certain logging commands then enable you to control header fields, which include a time stamp and device ID. The syslog server uses the device ID to identify the syslog generator. You can use the logging device-id command to generate syslog messages with identical or different device IDs to make messages appear to come from the same or different units in the cluster.
