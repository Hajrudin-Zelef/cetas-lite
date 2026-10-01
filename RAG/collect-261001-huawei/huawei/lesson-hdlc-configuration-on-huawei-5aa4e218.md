---
id: collect-261001-huawei/huawei/lesson-hdlc-configuration-on-huawei-5aa4e218
title: "lesson-hdlc-configuration-on-huawei-5aa4e218"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/lesson-hdlc-configuration-on-huawei-5aa4e218.md
source_anchor: ""
source_lines: [1, 219]
sha256: 596433662ff734e8206a70cd9dabc98e019960af621fcf11b7d5ad980ecd44f4
---

# lesson-hdlc-configuration-on-huawei-5aa4e218

In this lesson, we will talk about HDLC configuration on Huawei Routers. For our HDLC Configuration example, we will use the below basic topology:
Orderly, we will do the below configuration for our HDLC Configuration Example:
1) Enabling HDLC
2) Verifying HDLC
Let’s start to configure HDLC for this topology.
You can download this configuration on Huawei eNSP Labs Page.
In this basic HDLC Configuration Example, firstly we will configure the serial interfaces to use HDLC. After that, we will assign the ip address of this interface with its subnet mask. We will do this configuration on Router 1 firstly.
system-view
[Huawei-Router-1] interface Serial 0/0/0
[Huawei-Router-1-Serial0/0/0] link-protocol-hdlc
Warning : The encapsulation protocol of the link will be changed
Continue? [Y/N] : y
[Huawei-Router-1-Serial0/0/0] ip address 192.168.1.1 30
[Huawei-Router-1-Serial0/0/0] quit
Now, let’s configure Router 2 like above.
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
				Collapse			
			
				Expand			
		
- 
							Network Basics
- Network Topology Types
- Network Cabling
- TCP/IP Model
- OSI Referance Model
- 
							Network Devices
- Switches
- Hubs, Switches and Routers
- Network Device Types
- 
							Local Area Networks
- Home Network IP Address Allocation and Internet Connection
- Ethernet
- Local Area Networks
- 
							ICMP and ARP
- Proxy ARP
- Gratuitous ARP
- ARP
- ICMP
- 
							IPv4 Addressing
- IPv4 Header versus IPv6 Header
- IPv4 Addressing Basics
- Subnetting in IPv4
- 
							IPv6 Addressing
- IPv6 Addressing Basics
- Address Types of IPv6
- New Features of IPv6
- Subnetting in IPv6
- 
							TCP and UDP
- UDP (User Datagram Protocol)
- TCP Header : TCP Flags
- TCP Header : TCP Window Size, Checksum and Urgent Pointer
- TCP Header : TCP Options
- TCP versus UDP
- TCP Header
- TCP Header : Sequence and Acknowledgement Number
- TCP (Transmission Control Protocol)
- 
							VLANs
- GVRP (GARP VLAN Registration Protocol)
- VLAN Frame Tagging
- Virtual LANs (VLANs)
- VLAN Port Types and Port Assignment
- GVRP Configuration on Huawei
- VLAN Configuration on Huawei Switches
- 
							Spanning Tree
- MSTP Configuration on Huawei eNSP
- Spanning Tree Protocol
- STP Mechanism
- Rapid STP (RSTP)
- RSTP Configuration on Huawei
- STP Configuration on Huawei
- 
							Link Aggregation
- LACP Overview
- Link Aggregation on Huawei Routers
- 
							RIP
- RIPv2 Configuration on Huawei eNSP
- Routing Information Protocol
- RIP Next Generation
- RIPng Configuration on Huawei
- 
							OSPF
- OSPv2 Configuration on Huawei eNSP
- OSPF Basics
- OSPF Adjacency Mechanism
- Packet Types in OSPF
- LSA Types in OSPF
- Area Types in OSPF
- Network Types in OSPF
- OSPF Miscellenaous
- OSPFv3 (Open Shortest Path First Version 3)
- OSPFv3 Configuration on Huawei
- 
							WAN Protocols
- WAN Technologies
- PPPoE (Point-to-Point Protocol over Ethernet)
- HDLC and PPP
- PPP Configuration on Huawei
- HDLC Configuration on Huawei
- 
							GRE and IPSEC VPN
- What is IPSec VPN? How It Works?
- GRE (Generic Routing Encapsulation)
- IPSec VPN Configuration On Huawei
- Huawei GRE Tunnel Configuration
- 
							DHCP and DNS
- DNS Overview
- Dynamic Host Configuration Protocol
- DHCP IP Allocation
- DHCP Configuration On Huawei Routers
- 
							Network Address Translation (NAT)
- NAT Configuration on Huawei Routers
- NAT (Network Address Translation)
- 
							Access-Lists (ACLs)
- Huawei Access-Lists (ACL)
- 
							AAA
- AAA (Authentication, Authorization, Accounting)
- HWTACACS
- AAA Protocols : RADIUS and HWTACACS
- RADIUS
- Huawei AAA Configuration
- 
							Network Management
- SNMP (Simple Network Management Protocol)
- SNMP Configuration on Huawei
- 
							Huawei CLI
- Basic Huawei Router Configuration
- Huawei FTP Configuration
- Basic Huawei Configuration
- Huawei Router File System
- Huawei VRP (Versatile Routing Platform)
- 
							Security
- Huawei Port Security Configuration on Huawei eNSP
- 
							VLAN Routing
- VLAN Routing with Layer 3 Switch on Huawei
- VLAN Routing with Layer 2 Switch and Router on Huawei
- Switch Virtual Interfaces
- Inter VLAN Routing
- 
							Static Routing
- Huawei Router Interface Configuration
- Huawei Static Routing and Load Balancing
- Floating Static Route and Default Route Configuration
- 
										RJ45 Connector: Pinout, Types, Uses and Ethernet Wiring
					
 Part of: CCNA 200-301 v1.1 Course
- 
										WLAN Configuration Using WPA2 PSK With Packet Tracer
					
 Part of: CCNA 200-301 v1.1 Course
- 
										Cisco Single Area OSPF Configuration
					
 Part of: CCNA 200-301 v1.1 Course
- 
										Interpret QoS Configurations
					
 Part of: CCNP Enterprise 350-401 ENCOR v1.2
- 
										Multicast Source Discovery Protocol (MSDP)
					
 Part of: CCNP Enterprise 350-401 ENCOR v1.2
- 
										Inter VLAN Routing Configuration GNS3 Example
					
 Part of: CCIE Enterprise Infrastructure
- 
										IPv6 Configuration with Cisco Packet Tracer
					
 Part of: Cisco Packet Tracer Lab Course
- 
										Cisco Telnet Configuration with Packet Tracer
					
 Part of: Cisco Packet Tracer Lab Course
- 
										IPv4 Address Configuration with Packet Tracer
					
 Part of: Cisco Packet Tracer Lab Course
- 
										IPv4 Address Configuration
					
 Part of: CCNA 200-301 Labs
- More Lessons
		                    
		                        30 Jul, 2026		                    
						
					
		                    
		                        30 May, 2026		                    
						
					
- 250.000+ Students All Over The World
- 8.000+ Questions & Answers
- 100+ Lab Files & Cheat Sheets
- 30+ IT/Network Courses
- A Real Desire To Help You
- Daily Social Media Shares
- %100 Satisfaction
Leave a Reply
