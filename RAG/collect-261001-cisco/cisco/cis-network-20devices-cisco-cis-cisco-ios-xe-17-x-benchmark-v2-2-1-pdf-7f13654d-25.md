---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-25
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [4146, 4365]
sha256: 3910b757fc9a4edbdc8604ecf49d66e8f6461809a0c05962138cd92263f0d134
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 160 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped    
v8 
4.4 Implement and Manage a Firewall on Servers 
 Implement and manage a firewall on servers, where supported. Example 
implementations include a virtual firewall, operating system firewall, or a third-
party firewall agent. 
● ● ● 
v8 
4.5 Implement and Manage a Firewall on End-User 
Devices 
 Implement and manage a host-based firewall or port-filtering tool on end-user 
devices, with a default-deny rule that drops all traffic except those services and 
ports that are explicitly allowed. 
● ● ● 
v7 
9.2 Ensure Only Approved Ports, Protocols and Services 
Are Running 
 Ensure that only network ports, protocols, and services listening on a system 
with validated business needs, are running on each system. 
 ● ●

Page 161 
Internal Only - General 
3.1.2 Set 'no ip proxy-arp' (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
Disable proxy ARP on all interfaces. 
Rationale: 
Address Resolution Protocol (ARP) provides resolution between IP and MAC 
Addresses (or other Network and Link Layer addresses on none IP networks) within a 
Layer 2 network. 
Proxy ARP is a service where a device connected to one network (in this case the Cisco 
router) answers ARP Requests which are addressed to a host on another network, 
replying with its own MAC Address and forwarding the traffic on to the intended host. 
Sometimes used for extending broadcast domains across WAN links, in most cases 
Proxy ARP on enterprise networks is used to enable communication for hosts with mis-
configured subnet masks, a situation which should no longer be a common problem. 
Proxy ARP effectively breaks the LAN Security Perimeter, extending a network across 
multiple Layer 2 segments. Using Proxy ARP can also allow other security controls such 
as PVLAN to be bypassed. 
Impact: 
Organizations should plan and implement network policies to ensure unnecessary 
services are explicitly disabled. The 'ip proxy-arp' feature effectively breaks the LAN 
security perimeter and should be disabled. 
Audit: 
Verify the proxy ARP status 
 
hostname#sh ip int {<em>interface</em>} | incl proxy-arp  
Remediation: 
Disable proxy ARP on all interfaces. 
 
hostname(config)#interface {interface} 
hostname(config-if)#no ip proxy-arp 
  
Default Value: 
Enabled

Page 162 
Internal Only - General 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/ipaddr/command/ipaddr-
i4.html#GUID-AEB7DDCB-7B3D-4036-ACF0-0A0250F3002E 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped    
v8 
4.4 Implement and Manage a Firewall on Servers 
 Implement and manage a firewall on servers, where supported. Example 
implementations include a virtual firewall, operating system firewall, or a third-
party firewall agent. 
● ● ● 
v8 
4.5 Implement and Manage a Firewall on End-User 
Devices 
 Implement and manage a host-based firewall or port-filtering tool on end-user 
devices, with a default-deny rule that drops all traffic except those services and 
ports that are explicitly allowed. 
● ● ● 
v7 
9.2 Ensure Only Approved Ports, Protocols and Services 
Are Running 
 Ensure that only network ports, protocols, and services listening on a system 
with validated business needs, are running on each system. 
 ● ●

Page 163 
Internal Only - General 
3.1.3 Set 'no interface tunnel' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Verify no tunnel interfaces are defined. 
Rationale: 
Tunnel interfaces should not exist in general. They can be used for malicious purposes. 
If they are necessary, the network admin's should be well aware of them and their 
purpose. 
Impact: 
Organizations should plan and implement enterprise network security policies that 
disable insecure and unnecessary features that increase attack surfaces such as 'tunnel 
interfaces'. 
Audit: 
Verify no tunnel interfaces are defined 
 
hostname#sh ip int brief | incl tunnel  
Remediation: 
Remove any tunnel interfaces. 
 
hostname(config)#no interface tunnel {<em>instance</em>} 
Default Value: 
No tunnel interfaces are defined 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/interface/command/ir-
i1.html#GUID-0D6BDFCD-3FBB-4D26-A274-C1221F8592DF

Page 164 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped    
v8 
4.4 Implement and Manage a Firewall on Servers 
 Implement and manage a firewall on servers, where supported. Example 
implementations include a virtual firewall, operating system firewall, or a third-
party firewall agent. 
● ● ● 
v8 
4.5 Implement and Manage a Firewall on End-User 
Devices 
 Implement and manage a host-based firewall or port-filtering tool on end-user 
devices, with a default-deny rule that drops all traffic except those services and 
ports that are explicitly allowed. 
● ● ● 
v7 
9.2 Ensure Only Approved Ports, Protocols and Services 
Are Running 
 Ensure that only network ports, protocols, and services listening on a system 
with validated business needs, are running on each system. 
 ● ●

Page 165 
Internal Only - General 
3.1.4 Set 'ip verify unicast source reachable-via' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Examines incoming packets to determine whether the source address is in the 
Forwarding Information Base (FIB) and permits the packet only if the source is 
reachable through the interface on which the packet was received (sometimes referred 
to as strict mode). 
Rationale: 
Enabled uRPF helps mitigate IP spoofing by ensuring only packet source IP addresses 
only originate from expected interfaces. Configure unicast reverse-path forwarding 
(uRPF) on all external or high risk interfaces. 
Impact: 
Organizations should plan and implement enterprise security policies that protect the 
confidentiality, integrity, and availability of network devices. The 'unicast Reverse-Path 
Forwarding' (uRPF) feature dynamically uses the router table to either accept or drop 
packets when arriving on an interface. 
Audit: 
Verify uRPF is running on the appropriate interface(s) 
 
hostname#sh ip int {<em>interface</em>} | incl verify source  
Remediation: 
Configure uRPF. 
 
hostname(config)#interface {<em>interface_name</em>} 
hostname(config-if)#ip verify unicast source reachable-via rx allow-default 
Default Value: 
Unicast RPF is disabled. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/d1/sec-cr-i3.html#GUID-
2ED313DB-3D3F-49D7-880A-047463632757 
2. https://community.cisco.com/t5/routing/ip-verify-unicast-source-reachable-via-
rx/td-p/1710172

Page 166 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped    
v8 
4.4 Implement and Manage a Firewall on Servers 
 Implement and manage a firewall on servers, where supported. Example 
implementations include a virtual firewall, operating system firewall, or a third-
party firewall agent. 
● ● ● 
v8 
4.5 Implement and Manage a Firewall on End-User 
Devices 
 Implement and manage a host-based firewall or port-filtering tool on end-user 
devices, with a default-deny rule that drops all traffic except those services and 
ports that are explicitly allowed. 
● ● ● 
v7 
9.2 Ensure Only Approved Ports, Protocols and Services 
Are Running 
 Ensure that only network ports, protocols, and services listening on a system 
with validated business needs, are running on each system. 
 ● ●

Page 167 
Internal Only - General 
3.2 Border Router Filtering 
A border-filtering device connects "internal" networks such as desktop networks, DMZ 
networks, etc., to "external" networks such as the Internet. If this group is chosen, then 
ingress and egress filter rules will be required.

