---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636-5
title: "c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "distribution"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636.md
source_anchor: ""
source_lines: [166, 196]
sha256: 90dfd8cc20158b6e96f088e537802e77ca1f590004ace0facacd0eb9e8ab9acc
---

# c-en-us-td-docs-solutions-enterprise-education-schoolssra-dg-schoolssra-dg-schoo-1eebe636

•Port 80—Applications that use HTTP tunneling on port 80 can be handled by enforcing access policies within the web proxy configuration. Application access may be restricted based on applications, URL categories, and objects. Applications are recognized and blocked based on their user agent pattern, and by the use of regular expressions. The user may also specify categories of URL to block, including the predefined chat and peer-to-peer categories. Custom URL categories may also be defined. Peer-to-peer access may also be filtered based on object and MIME Multipurpose Internet Mail Extensions (MIME) types.
•Ports other than 80—Applications using ports other than 80 can be handled with the L4TM feature. L4TM block access to a specific application by preventing access to the server or block of IP addresses to which the client application must connect.
Note In the school design, the Cisco ASA is configured to allow only permitted ports (HTTP and HTTPS), so any connection attempts on other ports should be blocked by the firewall.
Note The Cisco IPS appliances and modules, and the Cisco ASA (using the modular policy framework), may also be used to block peer-to-peer file sharing and Internet applications.
The following are the guidelines for implementing a Cisco IronPort WSA appliance with WCCP on a Cisco ASA:
•Deploy WSA on the inside of the firewall so that the WSA can communicate with the clients without going through the firewall.
•Implement MD5 authentication to protect the communications between the Cisco ASA and the WSA(s).
•Configure a redirect-list on the firewall to indicate what traffic needs to be redirected. Make sure the WSA is always excluded from redirection.
•Ingress ACL on the firewall takes precedence over WCCP redirection, so make sure the ingress ACL is configured to allow HTTP and HTTPS traffic from clients and the WSA itself.
•In an existing proxy environment, deploy the WSA downstream from the existing proxy servers (closer to the clients).
•Cisco ASA does not support WCCP IP source address spoofing, therefore any upstream authentication or access controls based on client IP addresses are not supported. Without IP address spoofing, requests originating from a client are sourced with the IP address of the Web Proxy, and not the one of the client.
•TCP intercept, authorization, URL filtering, inspect engines, and IPS features do not apply to redirected flows of traffic served by the WSA cache. Content requested by the WSA is still subject to all the configured features on the firewall.
•Configure WSA access policies to block access to applications (AOL Messenger, Yahoo Messenger, BitTorrent, Kazaa, etc) and URL categories not allowed by the school's Internet access policies.
•If an out-of-band (OOB) management network is available, use a separate interface for administration.
Note WCCP, firewall, and other stateful features usually require traffic symmetry, whereby is traffic in both directions should flow through the same stateful device. The school architecture is designed with a single Internet path ensuring traffic symmetry. Care should be taken when implementing active-active firewall pairs as they may introduce asymmetric paths.
The Layer-4 Traffic Monitor (L4TM) service is deployed independently from the Web Proxy functionality, and its mission is to monitor network traffic for rogue activity and for any attempts to bypass port 80. L4TM works by listening to all UDP and TCP traffic and by matching domain names and IP addresses against entries in its own database tables to determine whether to allow incoming and outgoing traffic. The L4TM internal database is continuously updated with matched results for IP addresses and domain names. Additionally, the database table receives periodic updates from the IronPort update server (https://update-manifests.ironport.com).
The following are the key guidelines when deploying the L4 Traffic Monitor:
•Determine physical connection—L4TM requires traffic to be directed to the WSA for monitoring. This can be done by connecting a physical network tap, configuring SPAN port mirroring on a switch, or using a hub. Network taps forward packets in hardware, while SPAN port mirroring is generally done in software. On the other hand, SPAN port mirroring can be easily reconfigured, providing further flexibility.
•Location—Deploy L4TM in the network where it can see as much traffic as possible before getting out to the Internet through the firewall. It is important that the L4TM be logically connected after the proxy ports and before any device that performs network address translation (NAT) on client IP addresses.
•Action setting—The default setting for the L4TM is monitor only. Optionally you may configure the L4TM to monitor and block suspicious traffic. TCP connections are reset with the generating of TCP resets, while UDP sessions are tear down with ICMP unreachables. The use of L4TM blocking requires that the L4TM and the Web Proxy to be placed on the same network so that all clients are accessible on routes that are configured for data traffic.
In the school architecture, L4TM is deployed by setting a SPAN session on the distribution switch to replicate all TCP and UDP traffic on the links connecting to the inside interface of the firewall. Using SPAN provides greater flexibility, and inspecting the firewall's inside links ensures traffic is monitored before NAT and before being sent out the Internet. The L4TM deployment is shown in Figure 4-10.
Figure 4-10 L4TM Deployment
L4TM action is set to monitor only. Because the Internet firewall is configured to block any traffic bound to the Internet other than HTTP and HTTPS, there is no additional benefit in using L4TM blocking. If active mitigation is required, consider implementing a Cisco IPS module or appliance in in-line mode. When deployed in inline mode, the Cisco IPS is placed in the traffic path and is capable of stopping malicious traffic before it reaches the intended target. In addition, the Cisco IPS provides multiple configurable response actions including blocking the malicious packet only, blocking the entire session, or blocking any traffic coming from the offending system.
Configuration steps and examples are included in Chapter 11, "District Office Design."
Network Access Security and Control
Some of the most vulnerable points of the network are the access edges where students, staff and faculty connect to the network. With the proliferation of wireless networks, increased use of laptops and smart mobile devices, the school administration cannot simply rely on physical controls hoping to prevent unauthorized systems from being plugged into the ports of the access switches. Protection should be rather embedded into the network infrastructure, leveraging the native security features available in switches and routers. Furthermore, the network infrastructure should also provide dynamic identity or role-based access controls for all systems attempting to gain access.
Implementing role-based access controls for users and devices help reduce the potential loss of sensitive information by enabling schools to verify a user or device identity, privilege level, and security policy compliance before granting network access. Security policy compliance could consist of requiring antivirus software, OS updates or patches. Unauthorized, or noncompliant devices can be placed in a quarantine area where remediation can occur prior to gaining access to the network.
The Schools SRA achieves access security and control by using the following technologies:
•Catalyst Integrated Security Features (CISF)
•Cisco NAC Appliance
•Cisco Identity-Based Network Networking Services (IBNS)
