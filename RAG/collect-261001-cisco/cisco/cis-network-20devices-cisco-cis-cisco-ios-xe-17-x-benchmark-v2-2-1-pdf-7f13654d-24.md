---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-24
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [3949, 4145]
sha256: 2c368745a09ab3505a6218d7ff03d616bc95eb7049bb9023ebc4412932070c1a
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 152 
Internal Only - General 
2.4.2 Set AAA 'source-interface' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Force AAA to use the IP address of a specified interface for all outgoing AAA packets 
Rationale: 
This is required so that the AAA server (RADIUS or TACACS+) can easily identify 
routers and authenticate requests by their IP address. 
Impact: 
Organizations should design and implement authentication, authorization, and 
accounting (AAA) services for effective monitoring of enterprise network devices. 
Binding AAA services to the source-interface loopback enables these services. 
Audit: 
Perform the following to determine if AAA services are bound to a source interface: 
Verify a command string result returns 
 
hostname#sh run | incl tacacs source | radius source 
Remediation: 
Bind AAA services to the loopback interface. 
 
Hostname(config)#ip radius source-interface loopback 
{loopback_interface_number} 
or 
Hostname(config)#aaa group server tacacs+ {group_name} hostname(config-sg-
tacacs+)#ip tacacs source-interface {loopback_interface_number} 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/d1/sec-cr-i2.html#GUID-
22E8B211-751F-48E0-9C76-58F0FE0AABA8 
2. http://www.cisco.com/en/US/docs/ios-xml/ios/security/d1/sec-cr-i3.html#GUID-
54A00318-CF69-46FC-9ADC-313BFC436713

Page 153 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 5.6 Centralize Account Management 
 Centralize account management through a directory or identity service.  ● ● 
v7 
16.2 Configure Centralized Point of Authentication 
 Configure access for all accounts through as few centralized points of 
authentication as possible, including network, security, and cloud systems. 
 ● ●

Page 154 
Internal Only - General 
2.4.3 Set 'ntp source' to Loopback Interface (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Use a particular source address in Network Time Protocol (NTP) packets. 
Rationale: 
Set the source address to be used when sending NTP traffic. This may be required if 
the NTP servers you peer with filter based on IP address. 
Impact: 
Organizations should plan and implement network time protocol (NTP) services to 
establish official time for all enterprise network devices. Setting 'ntp source loopback' 
enforces the proper IP address for NTP services. 
Audit: 
Perform the following to determine if NTP services are bound to a source interface: 
Verify a command string result returns 
 
hostname#sh run | incl ntp source 
Remediation: 
Bind the NTP service to the loopback interface. 
 
hostname(config)#ntp source loopback {<em>loopback_interface_number}</em> 
Default Value: 
Source address is determined by the outgoing interface. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/bsm/command/bsm-cr-
n1.html#GUID-DF29FBFB-E1C0-4E5C-9013-D4CE59CA0B88

Page 155 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
8.4 Standardize Time Synchronization 
 Standardize time synchronization. Configure at least two synchronized time 
sources across enterprise assets, where supported. 
 ● ● 
v7 
6.1 Utilize Three Synchronized Time Sources 
 Use at least three synchronized time sources from which all servers and 
network devices retrieve time information on a regular basis so that timestamps 
in logs are consistent. 
 ● ●

Page 156 
Internal Only - General 
2.4.4 Set 'ip tftp source-interface' to the Loopback Interface 
(Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Specify the IP address of an interface as the source address for TFTP connections. 
Rationale: 
This is required so that the TFTP servers can easily identify routers and authenticate 
requests by their IP address. 
Impact: 
Organizations should plan and implement trivial file transfer protocol (TFTP) services in 
the enterprise by setting 'tftp source-interface loopback', which enables the TFTP 
servers to identify routers and authenticate requests by IP address. 
Audit: 
Perform the following to determine if TFTP services are bound to a source interface: 
Verify a command string result returns 
 
hostname#sh run | incl tftp source-interface 
Remediation: 
Bind the TFTP client to the loopback interface. 
 
hostname(config)#ip tftp source-interface loopback 
{<em>loobpback_interface_number</em>} 
Default Value: 
The address of the closest interface to the destination is selected as the source 
address. 
References: 
1. http://www.cisco.com/en/US/docs/ios-
xml/ios/fundamentals/command/F_through_K.html#GUID-9AA27050-A578-
47CD-9F1D-5A8E2B449209

Page 157 
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
 
3 Data Plane 
Services and settings related to the data passing through the router (as opposed to 
direct to it). The data plane is for everything not in control or management planes. 
Settings on a router concerned with the data plane include interface access lists, firewall 
functionality (e.g. CBAC), NAT, and IPSec. Settings for traffic-affecting services like 
unicast RPF verification and CAR/QoS also fall into this area.

Page 158 
Internal Only - General 
3.1 Routing Rules 
Unneeded services should be disabled.

Page 159 
Internal Only - General 
3.1.1 Set 'no ip source-route' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Disable the handling of IP datagrams with source routing header options. 
Rationale: 
Source routing is a feature of IP whereby individual packets can specify routes. This 
feature is used in several kinds of attacks. Cisco routers normally accept and process 
source routes. Unless a network depends on source routing, it should be disabled. 
Impact: 
Organizations should plan and implement network policies to ensure unnecessary 
services are explicitly disabled. The 'ip source-route' feature has been used in several 
attacks and should be disabled. 
Audit: 
Verify the command string result returns 
 
hostname#sh run | incl ip source-route 
Remediation: 
Disable source routing. 
 
hostname(config)#no ip source-route 
Default Value: 
Enabled by default 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/ipaddr/command/ipaddr-
i4.html#GUID-C7F971DD-358F-4B43-9F3E-244F5D4A3A93 
Additional Information: 
Reverting the Artifact logic will satisfy this control and will make this control independent 
of the way the assessor tool used checks the configuration (CIS CAT Pro Assessor 
check in 'sh run' results instead of 'sh run all' results)

