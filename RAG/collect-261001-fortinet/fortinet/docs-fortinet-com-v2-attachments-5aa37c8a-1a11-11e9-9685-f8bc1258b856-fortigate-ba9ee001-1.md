---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate-ba9ee001-1
title: "docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2017-05-26", "2018-03-21"]
keywords: ["ethernet", "license", "memory", "training"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001.md
source_anchor: ""
source_lines: [1, 203]
sha256: 6b269c8cf6d1c619c359f70e55b64fbdf6c3e00575d0001c5ec2751dad794855
---

# docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001

FortiOS™ Handbook - Transparent Mode
for FortiOS 5.6

FORTINET DOCUMENT LIBRARY
http://docs.fortinet.com
FORTINET VIDEO GUIDE
http://video.fortinet.com
FORTINET BLOG
https://blog.fortinet.com
CUSTOMER SERVICE & SUPPORT
https://support.fortinet.com 
http://cookbook.fortinet.com/how-to-work-with-fortinet-support/
FORTIGATE COOKBOOK
http://cookbook.fortinet.com
FORTINET TRAINING SERVICES
http://www.fortinet.com/training
FORTIGUARD CENTER
http://www.fortiguard.com
FORTICAST
http://forticast.fortinet.com
END USER LICENSE AGREEMENT
http://www.fortinet.com/doc/legal/EULA.pdf
FORTINET PRIVACY POLICY
https://www.fortinet.com/corporate/about-us/privacy.html
FEEDBACK
Email: techdocs@fortinet.com
Wednesday, March 21, 2018
FortiOS™ Handbook - Transparent Mode
01-563-479967-20180321

TABLE OF CONTENTS
Change Log 5
Introduction 7
What's new in FortiOS 5.6 8
Using the Management IP address to populate pac-file-url 8
Introduction to Transparent Mode 9
What is Transparent Mode? 9
Transparent Mode Features 9
Transparent Mode Installation 11
Installing a FortiGate in Transparent mode 11
Using a Virtual Wire Pair to Simplify Transparent Mode 12
Management IP configuration 13
In-band management details and example 13
Out-of-band management details and example 14
Networking in Transparent Mode 15
Packet Forwarding 15
MAC learning and L2 Forwarding Table 15
Broadcast, Multicast, and Unicast Forwarding 16
Multicast Processing 17
Source MAC Addresses 18
ARP Table 19
Verifying the Forwarding Database 19
Spanning Tree BPDUs Forwarding 19
Non-IPv4 Ethernet frames forwarding 20
Network Address Translation (NAT) 20
Configuring SNAT 20
Configuring DNAT 21
VLANs and Forwarding Domains 22
VLANs in Transparent Mode 22
Forwarding Domains in Transparent Mode 22
VLANs vs Forwarding Domains 24
VLAN Forwarding 24
Unknown VLANs and VLAN Forwarding 24
VLAN Trunking and MAC Address Learning 24

VLAN Translation 25
Inter-VDOM links between NAT/Route and Transparent VDOMs 25
Replay Traffic Scenario 25
Packet Forwarding using Cisco Protocols 26
Configuration Example 27
Firewalls and Security in Transparent Mode 29
Firewall Policy Look Up 29
Firewall Session List 29
Security Scanning 30
IPsec VPN in Transparent Mode 31
Using IPsec VPNs in Transparent Mode 31
Example 1: Remote Sites with Different Subnets 32
Example 2: Remote Sites on the Same Subnet 38
Using FortiManager and FortiAnalyzer 45
High Availability in Transparent Mode 46
Virtual Clustering 46
MAC Address Assignment 47
Best Practices 48

Change Log
Date Change Description
May 26, 2017 Initial publication.
Sept 1, 2017 Updated "VLAN Forwarding" on page 24.
Dec 13, 2017 Updated "Using FortiManager and FortiAnalyzer" on page 45.
Mar 21, 2018
Updated "Using a Virtual Wire Pair to Simplify Transparent Mode" on page 12 ,
"Forwarding Domains in Transparent Mode" on page 22 , and "Unknown VLANs
and VLAN Forwarding" on page 24.

Introduction
This guide explains how to use a FortiGate in Transparent Mode. It includes the following sections:
l Introduction to Transparent Mode : an overview of Transparent mode, including available features.
l Transparent Mode Installation : instructions for installing a FortiGate in Trnasparent mode.
l Networking in Transparent Mode: how networking is configured in Transparent mode.
l Firewalls and Security in Transparent Mode : information about using firewalls and security scanning in Transparent
mode.
l IPsec VPN in Transparent Mode: configuring IPsec VPNs using FortiGates in Transparent mode.
l Using FortiManager and FortiAnalyzer: using external management and logging.
l High Availability in Transparent Mode : configuring Transparent FortiGates in HA mode.
l Best Practices: the best practices for using Transparent mode.
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
7

What's new in FortiOS 5.6
The following new features have been added for Transparent Mode in FortiOS 5.6:
Using the Management IP address to populate pac-file-url
You can now use the Management IP address to populate pac-file-url in Transparent Mode. Previously,
only the interface IP could be used.
CLI syntax
config vdom
edit root
config system settings
set opmode transparent
set manageip 192.168.0.34/24
end
config web-proxy explicit
set pac-file-server-status enable
get pac-file-url [url.pac]
end
8 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

Introduction to Transparent Mode
This section contains a basic overview of Transparent mode, as well as a chart showing features available in
Transparent mode.
What is Transparent Mode?
A FortiGate unit can operate in one of two modes: Transparent or NAT/Route mode.
In Transparent mode, the FortiGate is installed between the internal network and the router. In this mode, the
FortiGate does not make any changes to IP addresses and only applies security scanning to traffic. When a
FortiGate is added to a network in Transparent mode, no network changes are required, except to provide the
FortiGate with a management IP address. Transparent mode is used primarily when there is a need to increase
network protection but changing the configuration of the network itself is impractical.
In NAT/Route mode, a FortiGate unit is installed as a gateway or router between two networks. This allows the
FortiGate to hide the IP addresses of the private network using network address translation (NAT).
Transparent Mode Features
Different FortiOS features are available depending on whether your FortiGate is in Transparent or NAT/Route
mode. The following table shows which features are available for each mode.
For a FortiGate in Transparent mode, the maximum number of Interfaces
per VDOM is 254. This value includes both physical and virtual interfaces.
For any other maximum values, please consult the Maximum Values Table,
available at docs.fortinet.com .
Feature/capability NAT Transparent Comment
Unicast Routing / Policy
Based routing
YES NO
VIP / IP pools / NAT YES YES Configurable from CLI only in
transparent mode
Multicast routing YES NO - Options
available to forward
multicast packets
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
9

Transparent Mode Features Introduction to Transparent Mode
Feature/capability NAT Transparent Comment
L2 forwarding NO YES
In Transparent mode, other
frames than IP can be
forwarded but with no UTM
processing.
Firewall (packet filtering / NAT
/ Authentication)
YES YES
IPv6 capable YES YES
Traffic shaping - TOS
classification
YES YES
Hardware acceleration YES YES
All Security profile features (ex
IPS, Application Control, Web
Filtering, etc ...)
YES YES
Security Fabric YES NO
FortiView YES YES
IPsec gateway YES YES - Policy based
mode only
SSL gateway YES NO
High-Availability (HA) - Virtual
Cluster YES YES
802.3ad (LACP / port
aggregation)
YES YES
HA port redundancy YES YES FortiGate hardware
dependent
802.1q - VLAN trunking YES YES
802.1d - Spanning Tree NO NO - option to
forward BPDUs
Logging and reporting (disk
and memory logging,
FortiCloud, syslog, and
FortiAnalyzer)
YES YES
Managed by FortiManager YES YES
10 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

