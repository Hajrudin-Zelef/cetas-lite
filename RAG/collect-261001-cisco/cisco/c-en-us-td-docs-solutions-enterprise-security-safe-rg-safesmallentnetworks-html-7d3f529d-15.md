---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-15
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [473, 548]
sha256: aaf7a80d9d31e1ea064396adae59eba64e1f2e0ae2fbb9708a15decf9c64d734
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

Figure 22 illustrates the relationship between the various elements of the Cisco AnyConnect Secure Mobility solution. Remote and mobile users use the Cisco AnyConnect Secure VPN client to establish VPN sessions with the Cisco ASA appliance. The Cisco ASA sends Web traffic to the WSA appliance along with information identifying the user by IP address and user name. The WSA scans the traffic, enforces acceptable use policies, and protects the user from security threats. The Cisco ASA returns all traffic deemed safe and acceptable to the user.
All Internet traffic scanning is done by the WSA, not the client on the mobile device. This improves overall performance by not burdening the mobile device, some of which have limited processing power. In addition, by scanning Internet traffic on the network, the enterprise can more easily and quickly update security updates and acceptable use policies since the enterprise does not have to wait days, weeks, or even months to push the updates to the client. The WSA tracks the requests it receives and applies policies configured for remote users to traffic received from remote users.
For complete details about the Cisco AnyConnect Secure Mobility solution, refer to the documentation available at: http://www.cisco.com/en/US/netsol/ns1049/index.html.
Threats Mitigated
The success of the security tools and measures in place ultimately depends on the degree they enhance visibility and control. Simply put, security can be defined as a function of visibility and control. Without any visibility, it is difficult to enforce any control, and without any control, it is hard to achieve an adequate level of security. Therefore, the security tools selected in the enterprise network design were carefully chosen not only to mitigate certain threats, but also to increase the overall visibility and control.
Table 2 summarizes how the security tools and measures used in the small enterprise network design help mitigate certain threats and how they contribute to increasing visibility and control. Note the table is provided for illustration purposes and it is not intended to include all possible security tools and threats.
Table 2 Security Measures of the Enterprise Design Profile for Small Enterprise Networks
Network Foundation Protection
X
X
X
X
X
Stateful Firewall
X
X
X
X
X
IPS
X
X
X
X
X
X
Mobile Security
X
X
X
X
X
X
X
Web Security
X
X
X
X
X
X
E-mail Security
X
X
X
X
X
Access Security and Control
X
X
X
X
Network Security Deployment
This section of the document describes the deployment best practices for the key security platforms and features used in the small enterprise network design, including deployment and setting guidelines and configuration examples.
Internet Border Router Deployment
Whether the Internet border router is managed by the enterprise or the ISP, it must be hardened following the best practices listed in Network Foundation Protection. This includes restricting and controlling administrative access, protecting the management and control planes, and securing the dynamic exchange of routing information. In addition, the Internet border router may be leveraged as the first layer of protection against outside threats. To that end, edge ACLs, uRPF, and other filtering mechanisms may be implemented for anti-spoofing and to block invalid packets.
The following configuration snippet illustrates the structure of an edge ACL applied to the upstream interface of the Internet border router. The ACL is designed to block invalid packets and to protect the infrastructure IP addresses from the Internet. The configuration assumes the enterprise is assigned the 198.133.219.0/24 address block for its public-facing services and that the upstream link is configured in the 64.104.10.0/24 subnet.
! !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!--- Module 1: Anti-spoofing Denies!--- These ACEs deny fragments, RFC 1918 space,!--- invalid source addresses, and spoofs of!--- internal space (space as an external source).!!--- Deny fragments.access-list 110 deny tcp any 198.133.219.0 0.0.0.255 fragmentsaccess-list 110 deny udp any 198.133.219.0 0.0.0.255 fragmentsaccess-list 110 deny icmp any 198.133.219.0 0.0.0.255 fragments!--- Deny special-use address sources.!--- See RFC 3330 for additional special-use addresses.access-list 110 deny ip host 0.0.0.0 anyaccess-list 110 deny ip 127.0.0.0 0.255.255.255 anyaccess-list 110 deny ip 192.0.2.0 0.0.0.255 anyaccess-list 110 deny ip 224.0.0.0 31.255.255.255 any!--- Filter RFC 1918 space.access-list 110 deny ip 10.0.0.0 0.255.255.255 anyaccess-list 110 deny ip 172.16.0.0 0.15.255.255 anyaccess-list 110 deny ip 192.168.0.0 0.0.255.255 any!--- Deny packets spoofing the enterprise public addressesaccess-list 110 deny ip 198.133.219.0 0.0.0.255 any!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!--- Module 2: Explicit Permit!--- Permit only applications/protocols whose destination!--- address is part of the infrastructure IP block.!--- The source of the traffic should be known and authorized.!!--- Permit external BGP to peer 64.104.10.113access-list 110 permit tcp host 64.104.10.114 host 64.104.10.113 eq bgpaccess-list 110 permit tcp host 64.104.10.114 eq bgp host 64.104.10.113!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!--- Module 3: Explicit Deny to Protect Infrastructureaccess-list 110 deny ip 64.104.10.0 0.0.0.255 any!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!--- Module 4: Explicit Permit for Traffic to Company's Public!--- Subnet.access-list 110 permit ip any 198.133.219.0 0.0.0.255!
Note The 64.104.0.0/16 and 198.133.219.0/24 address blocks used in the examples in this document are reserved for the exclusive use of Cisco Systems, Inc.
Internet Firewall Deployment
The mission of the Internet firewall is to protect the company's internal resources and data from external threats, secure the public services provided by the DMZ, and to control users' traffic to the Internet. The small enterprise network design uses a Cisco ASA appliance as illustrated in Figure 23.
Figure 23 Internet Edge Firewall
The Cisco ASA is implemented with three interface groups, each one representing a distinct security domain:
•Inside—The interface connecting to the core/distribution switch that faces the interior of the network where internal users and resources reside. The inside interface connects to the internal trusted networks, therefore it is given the highest security level, 100.
•Outside—Interface connecting to the Internet border router. The router may be managed either by the enterprise or a service provider. The outside interface connects to the Internet, hence it is given the lowest security level, 0.
•Demilitarized Zone (DMZ)—The DMZ hosts services that are accessible over the Internet. These services may include the company's website and E-mail services. The DMZ serves a medium level security segment, therefore should be given any security value between the ones defined for the inside and the outside interfaces, for example, 50.
The Internet firewall acts as the primary gateway to the Internet; therefore, its deployment should be carefully planned. The following are key aspects to be considered when implementing the firewall:
•Firewall hardening and monitoring
•Network Address Translation (NAT)
Firewall Hardening and Monitoring
The Cisco ASA should be hardened in a similar fashion as the infrastructure routers and switches. According to the Cisco SAFE security best practices, the following is a summary of the measures to be taken:
•Implement dedicated management interfaces to the OOB management network.
•Present legal notification for all access attempts.
•Use HTTPS and SSH for device access. Limit access to known IP addresses used for administrative access.
•Configure AAA for role-based access control and logging. Use a local fallback account in case AAA server is unreachable.
•Use NTP to synchronize the time.
