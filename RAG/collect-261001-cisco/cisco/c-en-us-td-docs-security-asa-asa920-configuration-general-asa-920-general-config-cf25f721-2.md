---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721-2
title: "c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721.md
source_anchor: ""
source_lines: [83, 141]
sha256: 106e2ed07b2b442f0bd18bffe1c4c480b5c0f09524ef542226467ef214f04cee
---

# c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721

                                          The unit state (active or standby)
- 
                                          
                                          Hello messages (keep-alives)
- 
                                          
                                          Network link status
- 
                                          
                                          MAC address exchange
- 
                                          
                                          Configuration replication and synchronization
Interface for the failover link
A failover link interface is a dedicated communication channel that
- 
                                          
                                          exists solely for failover communication between high availability units
- 
                                          
                                          uses an unused data interface (physical, subinterface, or EtherChannel)
- 
                                          
                                          cannot be shared with normal networking interfaces or user data traffic, and
- 
                                          
                                          requires specific interface sizing based on the device model.
Interface restrictions and requirements
You can use an unused data interface (physical, subinterface, or EtherChannel) as the failover link; however, you cannot specify an interface that is currently configured with a name. The failover link interface is not configured as a normal networking interface; it exists for failover communication only. This interface can only be used for the failover link (and also for the state link). For most models, you cannot use a management interface for failover unless explicitly described below.
The ASA does not support sharing interfaces between user data and the failover link. You also cannot use separate subinterfaces on the same parent for the failover link and for data.
See these guidelines for the failover link:
- 
                                          
                                          5506-X through 5555-X—You cannot use the Management interface as the failover link; you must use a data interface. The only exception is for the 5506H-X, where you can use the management interface as the failover link.
- 
                                          
                                          5506H-X—You can use the Management 1/1 interface as the failover link. If you configure it for failover, you must reload the device for the change to take effect. In this case, you cannot also use the ASA Firepower module, because it requires the Management interface for management purposes.
- 
                                          
                                          Firepower 4100/9300—You cannot use the management-type interface for the failover link.
- 
                                          
                                          See these guidelines for sizing the link. Table 1. Failover link size Model Interface Size for Combined Failover and State Link Firepower 1010 1 Gbps Firepower 1100 1 Gbps Firepower 2100 1 Gbps Secure Firewall 3100 Secure Firewall 3105—1 Gbps Secure Firewall 3110—1 Gbps Secure Firewall 3120—1 Gbps Secure Firewall 3130—10 Gbps Secure Firewall 3140—10 Gbps Firepower 4100 10 Gbps Secure Firewall 4200 10 Gbps Firepower 9300 10 Gbps
The alternation frequency is equal to the unit hold time (the failover polltime unit command).
| Note | If you have a large configuration and a low unit hold time, alternating between the member interfaces can prevent the secondary unit from joining/re-joining. In this case, disable one of the member interfaces until after the secondary unit joins. | 
For an EtherChannel used as the failover link, to prevent out-of-order packets, only one interface in the EtherChannel is used. If that interface fails, then the next interface in the EtherChannel is used. You cannot alter the EtherChannel configuration while it is in use as a failover link.
Connect the failover link
The failover link enables communication and coordination between units in a failover configuration.
You need to establish a physical connection between the failover units to enable proper failover functionality.
Procedure
| Connect the failover link using one of the following methods:  If you do not use a switch between the units and the interface fails, the link is brought down on both peers. This condition may hamper troubleshooting efforts because you cannot easily determine which unit has the failed interface and caused the link to come down. The ASA supports Auto-MDI/MDIX on its copper Ethernet ports, so you can either use a crossover cable or a straight-through cable. If you use a straight-through cable, the interface automatically detects the cable and swaps one of the transmit/receive pairs to MDIX. | 
The failover link is physically connected and ready for failover configuration.
Stateful failover links
A stateful failover link is a network connection that
- 
                                       
                                       passes connection state information between failover units, and
- 
                                       
