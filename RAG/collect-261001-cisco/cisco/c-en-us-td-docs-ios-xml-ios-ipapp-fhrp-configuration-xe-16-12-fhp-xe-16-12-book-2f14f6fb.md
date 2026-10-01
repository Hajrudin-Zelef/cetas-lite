---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-ipapp-fhrp-configuration-xe-16-12-fhp-xe-16-12-book-2f14f6fb
title: "c-en-us-td-docs-ios-xml-ios-ipapp-fhrp-configuration-xe-16-12-fhp-xe-16-12-book--2f14f6fb"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-ipapp-fhrp-configuration-xe-16-12-fhp-xe-16-12-book--2f14f6fb.md
source_anchor: ""
source_lines: [1, 96]
sha256: 5fbd54a694aa035ca8a3c8e657fa6443485baa9f772ca3ab1b33b21e5fb28aad
---

# c-en-us-td-docs-ios-xml-ios-ipapp-fhrp-configuration-xe-16-12-fhp-xe-16-12-book--2f14f6fb

ARP—Address Resolution Protocol (ARP). ARP performs a required function in IP routing. ARP finds the hardware address, also known
                           as Media Access Control (MAC) address, of a host from its known IP address. ARP maintains a cache (table) in which MAC addresses
                           are mapped to IP addresses. ARP is part of all Cisco IOS systems running IP. 
                           	 
                        
                        
                        
                           active device—The primary device in an HSRP group that is currently forwarding packets for the virtual device. 
                           	 
                        
                        
                         
                           		active RP—The active RP that controls the system, provides network services, runs the routing protocols, and presents the system management
                           interface. 
                           	 
                        
                        
                        
                           			client group—An HSRP group that is created on a subinterface and linked to the primary group via the group name. 
                        
                        
                        
                           HSRP—Hot Standby Router Protocol. Protocol that provides high network availability and transparent network-topology changes. HSRP
                           creates a router group with a lead device that services all packets sent to the HSRP address. The lead device is monitored
                           by other devices in the group, and if it fails, one of these standby HSRP devices inherits the lead position and the HSRP
                           group address. 
                           	 
                        
                        
                         
                           		ISSU—In Service Software Upgrade. A process that allows Cisco IOS software to be updated or otherwise modified while packet forwarding
                           continues. In most networks, planned software upgrades are a significant cause of downtime. ISSU allows Cisco IOS software
                           to be modified while packet forwarding continues, which increases network availability and reduces downtime caused by planned
                           software upgrades. 
                           	 
                        
                        
                        
                           			primary group—An HSRP group that is required on a physical interface for the purposes of electing active and standby devices. 
                        
                        
                         
                           		RF—Redundancy Facility. A structured, functional interface used to notify its clients of active and standby state progressions
                           and events. 
                           	 
                        
                        
                         
                           		RP—Route Processor. A generic term for the centralized control unit in a chassis.
                           	 
                        
                        
                         
                           		RPR—Route Processor Redundancy. RPR provides an alternative to the High System Availability (HSA) feature. HSA enables a system
                           to reset and use a standby Route Processor (RP) if the active RP fails. Using RPR, you can reduce unplanned downtime because
                           RPR enables a quicker switchover between an active and standby RP if the active RP experiences a fatal error. 
                           	 
                        
                        
                         
                           		RPR+—An enhancement to RPR in which the standby RP is fully initialized. 
                           	 
                        
                        
                         
                           		standby group—The set of devices participating in HSRP that jointly emulate a virtual device. 
                           	 
                        
                        
                         
                           		standby device—The backup device in an HSRP group. 
                           	 
                        
                        
                         
                           		standby RP—The backup RP. 
                           	 
                        
                        
                         
                           		switchover—An event in which system control and routing protocol execution are transferred from the active RP to the standby RP. Switchover
                           may be a manual operation or may be induced by a hardware or software fault. Switchover may include transfer of the packet
                           forwarding function in systems that combine system control and packet forwarding in an indivisible unit. 
                           	 
                        
                        
                         
                           		virtual IP address—The default gateway IP address configured for an HSRP group. 
                           	 
                        
                        
                        
                           virtual MAC address—For Ethernet and FDDI, the automatically generated MAC address when HSRP is configured. The standard virtual MAC address
                           used is: 0000.0C07.ACxy, where 
                           		xy is the group number in hexadecimal. The functional address is used for Token Ring. The virtual MAC address is different for
                           HSRP version 2.
