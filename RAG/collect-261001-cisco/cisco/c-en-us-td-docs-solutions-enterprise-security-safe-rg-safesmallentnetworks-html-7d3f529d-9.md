---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-9
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [286, 318]
sha256: 91d72e43e33ba4483476f7ba8ab310c67a99a66da7021db40dd5cc526b55d40a
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

Explicit forward proxy mode requires the enterprise to have control over the configuration of the endpoints, which may not be always possible. For example, the enterprise may allow the use of personal laptops, smart-phones, and other devices outside the company's administration. Transparent proxy mode, on the other hand, provides a transparent integration of WSA without requiring any configuration control over the endpoints. It also eliminates the possibility of users reconfiguring their Web browsers to bypass the appliance without knowledge of the administrators. For these reasons, the small enterprise network design implements transparent proxy with WCCP. In this configuration, the Cisco ASA at the Internet perimeter is leveraged as a WCCP server while the WSA act as a WCCP Traffic Processing Entity.
The Cisco ASA uses WCCP version 2, which has a built-in failover and load balancing mechanism. Per WCCPv2 specification, multiple appliances (up to 32 entities) can be configured as part of the same service group. HTTP and HTTPS traffic is load-balanced across the active appliances based on source and destination IP addresses. The server (Cisco ASA) monitors the availability of each appliance in the group and can identify appliance failures within 30 seconds. After failure, traffic is redirected across the remaining active appliances. In the case where no appliances are active, WCCP takes the entire service group offline and subsequent requests bypass redirection. In addition, WCCPv2 supports MD5 authentication for the communications between WCCP server and WSA appliances.
Note In the event the entire service group fails, WCCP automatically bypasses redirection, allowing users to browse the Internet without the Web controls. In case it is desired to handle a group failure by blocking all traffic, an outbound ACL may be configured on the Cisco ASA outside interface to permit HTTP/HTTPS traffic originated from the WSA appliance itself and to block any direct requests from clients. The ACL may also have to be configured to permit HTTP/HTTPS access from IPS and other systems requiring such access.
WCCPv2 supports Generic Route Encapsulation (GRE) and Layer 2-based redirection; however, the Cisco ASA only supports GRE. In addition, WCCP is supported only on the ingress of an interface. The only topology supported is one where both clients and WSA are reachable from the same interface and where the WSA can directly communicate with the clients without going through the Cisco ASA. For these reasons, the WSA appliance is deployed at the inside segment of the Cisco ASA.
Figure 11 illustrates the how WCCP redirection works in conjunction with Cisco ASA.
Figure 11 WCCP Redirection
The following steps describe what takes place in Figure 11:
1. Client's browser requests connection to http://website.com.
2. Cisco ASA intercepts and redirects HTTP requests over GRE.
3. If content not present in local cache, WSA performs a DNS query on destination domain and checks the received IP address against URL and reputation rules and allows/denies request accordantly.
4. WSA fetches content from destination Web site.
5. Content is inspected and then delivered directly to the requesting client.
The WSA appliance may also be configured to control and block peer-to-peer file-sharing and Internet applications such as AOL Messenger, BitTorrent, Skype, Kazaa, etc. The way WSA handles these applications depends on the TCP port used for transport:
•Port 80—Applications that use HTTP tunneling on port 80 can be handled by enforcing access policies within the Web proxy configuration. Application access may be restricted based on applications, URL categories, and objects. Applications are recognized and blocked based on their user agent pattern and by the use of regular expressions. The user may also specify categories of URL to block, including the predefined chat and peer-to-peer categories. Custom URL categories may also be defined. Peer-to-peer access may also be filtered based on object and MIME Multipurpose Internet Mail Extensions (MIME) types.
•Ports other than 80—Applications using ports other than 80 can be handled with the L4TM feature. L4TM blocks access to a specific application by preventing access to the server or block of IP addresses to which the client application must connect.
Note The Cisco IPS appliances and modules, and the Cisco ASA (using the modular policy framework), may also be used to block peer-to-peer file sharing and Internet applications.
The following are the guidelines for implementing a Cisco IronPort WSA appliance with WCCP on a Cisco ASA:
•Deploy WSA on the inside of the firewall so that the WSA can communicate with the clients without going through the firewall.
•Implement MD5 authentication to protect the communications between the Cisco ASA and the WSA(s).
•Configure a redirect-list on the firewall to indicate what traffic needs to be redirected. Ensure that the WSA is always excluded from redirection.
•Ingress ACL on the firewall takes precedence over WCCP redirection, so ensure that the ingress ACL is configured to allow HTTP and HTTPS traffic from clients and the WSA itself.
•In an existing proxy environment, deploy the WSA downstream from the existing proxy servers (closer to the clients).
•Cisco ASA does not support WCCP IP source address spoofing, therefore any upstream authentication or access controls based on client IP addresses are not supported. Without IP address spoofing, requests originating from a client are sourced with the IP address of the Web Proxy and not the one of the client.
•TCP intercept, authorization, URL filtering, inspect engines, and IPS features do not apply to redirected flows of traffic served by the WSA cache. Content requested by the WSA is still subject to all the configured features on the firewall.
•Configure WSA access policies to block access to applications (AOL Messenger, Yahoo Messenger, BitTorrent, Kazaa, etc.) and URL categories not allowed by the enterprise Internet access policies.
•If an out-of-band (OOB) management network is available, use a separate interface for administration.
Note WCCP, firewall, and other stateful features usually require traffic symmetry, whereby traffic in both directions should flow through the same stateful device. The small enterprise network design is designed with a single Internet path ensuring traffic symmetry. Care should be taken when implementing active-active firewall pairs as they may introduce asymmetric paths.
The Layer 4 Traffic Monitor (L4TM) service is deployed independently from the Web Proxy functionality and its mission is to monitor network traffic for rogue activity and for any attempts to bypass port 80. L4TM works by listening to all UDP and TCP traffic and by matching domain names and IP addresses against entries in its own database tables to determine whether to allow incoming and outgoing traffic. The L4TM internal database is continuously updated with matched results for IP addresses and domain names. Additionally, the database table receives periodic updates from the IronPort update server (https://update-manifests.ironport.com).
For more information on how to configure the L4TM feature on Cisco Ironport WSA, see:
•Cisco SAFE Reference Guide- http://www.cisco.com/en/US/docs/solutions/Enterprise/Security/SAFE_RG/SAFE_rg.html
•Cisco IronPort WSA User Guide-http://www.ironport.com/support
Configuration steps and examples are included in "Appendix D-Web Security Deployment" section of the document.
Serverfarm Protection
