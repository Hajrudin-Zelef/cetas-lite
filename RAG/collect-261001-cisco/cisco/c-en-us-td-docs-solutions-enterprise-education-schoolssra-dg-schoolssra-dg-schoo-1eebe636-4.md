---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636-4
title: "c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["cost", "distribution"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636.md
source_anchor: ""
source_lines: [137, 165]
sha256: 84561f0585697eeeb3add78db23ece3220d97e5e580fde8b271a767eeaa4980c
---

# c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636

•Load Balancer—A load balancer such as Cisco ACE Application Control Engine (ACE) load-balances traffic across multiple ESA appliances.
IronPort NIC pairing is the most cost-effective solution (see Figure 4-7), because it does not require the implementation of multiple ESA appliances and other hardware. It does not, however, provide redundancy in case of chassis failure.
Figure 4-7 Cisco IronPort ESA NIC Pairing
For more information on how to configure ESA, refer to the Cisco SAFE Reference Guide (http://www.cisco.com/en/US/docs/solutions/Enterprise/Security/SAFE_RG/SAFE_rg.html) and IronPort ESA User's Guide (http://www.ironport.com/support/).
Web Security Guidelines
The school architecture implements a Cisco IronPort S Series Web Security Appliance (WSA) to block access to sites with content that may be harmful or inappropriate for minors, and to protect the schools from web-based malware and spyware.
Cisco IronPort WSA's protection relies in two independent services:
•Web Proxy—This provides URL filtering, web reputation filters, and optionally anti-malware services. The URL filtering capability defines the handling of each web transaction based on the URL category of the HTTP requests. Leveraging the SensorBase network, the web reputation filters analyze the web server behavior and characteristics to identify suspicious activity and protect against URL-based malware. The anti-malware service leverages anti-malware scanning engines such as Webroot and McAfee to monitor for malware activity.
•Layer 4 Traffic Monitoring (L4TM)—Service configured to monitor all Layer-4 traffic for rogue activity and to detect infected clients.
Note The SensorBase network is an extensive network that monitors global E-mail and web traffic for anomalies, viruses, malware, and other abnormal behavior. The network is composed of the Cisco IronPort appliances, Cisco ASA, and Cisco IPS appliances and modules installed in more than 100,000 organizations worldwide, providing a large and diverse sample of Internet traffic patterns.
As the school design assumes a centralized Internet connection, the WSA is implemented at the distribution layer of the district office network. This allows the inspection and enforcement of web access policies to all students, staff, and faculty residing at any of the school premises. Logically, the WSA sits in the path between web users and the Internet, as shown in Figure 4-8.
Figure 4-8 Cisco IronPort WSA
There are two deployment modes for the Web Proxy service:
•Explicit Forward Proxy—Client applications, such as web browsers, are aware of the Web Proxy and must be configured to point to the WSA. The web browsers can be either configured manually or by using Proxy Auto Configuration (PAC) files. The manual configuration does not allow for redundancy, while the use of PAC files allows the definition of multiple WSAs for redundancy and load balancing. If supported by the browser, the Web Proxy Autodiscovery Protocol (WPAD) can be used to automate the deployment of PAC files. WPAD allows the browser to determine the location of the PAC file using DHCP and DNS lookups.
•Transparent Proxy—Client applications are unaware of the Web Proxy and do not have to be configured to connect to the proxy. This mode requires the implementation of a Web Cache Communications Protocol (WCCP) enable device or a Layer-4 load balancer in order to intercept and redirect traffic to the WSA. Both deployment options provide for redundancy and load balancing.
Explicit forward proxy mode requires the school to have control over the configuration of the endpoints, which may not be always possible. For example, the school may allow students, staff and faculty to use personal laptops, smart-phones and other devices outside the school's administration. Transparent proxy mode, on the other hand, provides a transparent integration of WSA without requiring any configuration control over the endpoints. It also eliminates the possibility of users reconfiguring their web browsers to bypass the appliance without knowledge of the administrators. For these reasons, the school architecture implements transparent proxy with WCCP. In this configuration, the Cisco ASA at the Internet perimeter is leveraged as a WCCP server while the WSA act as a WCCP Traffic Processing Entity.
Note It is recommended to enable both the Layer-4 traffic monitor and transparent proxy during the initial System Setup Wizard. Either of these services can be disabled or reconfigured after initial setup from the web interface. If you do not enable one of the features in the System Setup Wizard and then need to enable it later, you must run the System Setup Wizard again, losing all configurations added to the appliance.
The Cisco ASA uses WCCP version 2, which has a built-in failover and load balancing mechanism. Per WCCPv2 specification, multiple appliances (up to 32 entities) can be configured as part of the same service group. HTTP and HTTPS traffic is load-balanced across the active appliances based on source and destination IP addresses. The server (Cisco ASA) monitors the availability of each appliance in the group, and can identify appliance failures within 30 seconds. After failure, traffic is redirected across the remaining active appliances. In the case no appliances are active, WCCP takes the entire service group offline and subsequent requests bypass redirection. In addition, WCCPv2 supports MD5 authentication for the communications between WCCP server and WSA appliances.
Note In the event the entire service group fails, WCCP automatically bypasses redirection, allowing users to browse the Internet without the Web controls. In case it is desired to handle a group failure by blocking all traffic, an outbound ACL may be configured on the Cisco ASA outside interface to permit HTTP/HTTPS traffic originated from the WSA appliance itself and to block any direct requests from clients. The ACL may also have to be configured to permit HTTP/HTTPS access from IPS and other systems requiring such access.
WCCPv2 supports Generic Route Encapsulation (GRE) and Layer-2-based redirection; however, the Cisco ASA only supports GRE. In addition, WCCP is supported only on the ingress of an interface. The only topology supported is one where both clients and WSA are reachable from the same interface, and where the WSA can directly communicate with the clients without going through the Cisco ASA. For these reasons, the WSA appliance is deployed at the inside segment of the Cisco ASA.
Figure 4-9 illustrates the how WCCP redirection works in conjunction with Cisco ASA.
Figure 4-9 WCCP Redirection
The following steps describe what takes place in Figure 4-9:
Step 1 Client's browser requests connection to http://website.com.
Step 2 Cisco ASA intercepts and redirects HTTP requests over GRE.
Step 3 If content not present in local cache, WSA performs a DNS query on destination domain and checks the received IP address against URL and reputation rules, and allows/denies request accordantly.
Step 4 WSA fetches content from destination web site.
Step 5 Content is inspected and then delivered directly to the requesting client.
The WSA appliance may also be configured to control and block peer-to-peer file-sharing and Internet applications such as AOL Messenger, BitTorrent, Skype, Kazaa, etc. The way WSA handles these applications depends on the TCP port used for transport:
