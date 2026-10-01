---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf-2
title: "c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "latency"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf.md
source_anchor: ""
source_lines: [78, 142]
sha256: 577416f9f69f3576a0b837a3f3e44cf313aefb69981ce3c3d1fc7811e835bbfe
---

# c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf

                                       5506-X through 5555-X—You cannot use the Management interface as the failover link; you must use a data interface. The only exception is for the 5506H-X, where you can use the management interface as the failover link.
- 
                                       				
                                       5506H-X—You can use the Management 1/1 interface as the failover link. If you configure it for failover, you must reload the device for the change to take effect. In this case, you cannot also use the ASA Firepower module, because it requires the Management interface for management purposes.
- 
                                       				
                                       5585-X—Do not use the Management 0/0 interface, even though it can be used as a data interface. It does not support the necessary performance for this use.
- 
                                       				
                                       Firepower 4100/9300—We recommend that you use a 10 GB data interface for the combined failover and state link. You cannot use the management-type interface for the failover link.
- 
                                       				
                                       All other models—1 GB interface is large enough for a combined failover and state link.
For a redundant interface used as the failover link, see the following benefits for added redundancy:
-  
                                       		  
                                       When a failover unit boots up, it alternates between the member interfaces to detect an active unit.
-  
                                       		  
                                       If a failover unit stops receiving keepalive messages from its peer on one of the member interfaces, it switches to the other member interface.
The alternation frequency is equal to the unit hold time (the failover polltime unit command).
| Note | If you have a large configuration and a low unit hold time, alternating between the member interfaces can prevent the secondary unit from joining/re-joining. In this case, disable one of the member interfaces until after the secondary unit joins. | 
For an EtherChannel used as the failover link, to prevent out-of-order packets, only one interface in the EtherChannel is used. If that interface fails, then the next interface in the EtherChannel is used. You cannot alter the EtherChannel configuration while it is in use as a failover link.
Connecting the Failover Link
Connect the failover link in one of the following two ways:
- 
                                       		  
                                       Using a switch, with no other device on the same network segment (broadcast domain or VLAN) as the failover interfaces of the ASA.
- 
                                       		  
                                       Using an Ethernet cable to connect the units directly, without the need for an external switch.
If you do not use a switch between the units, if the interface fails, the link is brought down on both peers. This condition may hamper troubleshooting efforts because you cannot easily determine which unit has the failed interface and caused the link to come down.
The ASA supports Auto-MDI/MDIX on its copper Ethernet ports, so you can either use a crossover cable or a straight-through cable. If you use a straight-through cable, the interface automatically detects the cable and swaps one of the transmit/receive pairs to MDIX.
Stateful Failover Link
To use Stateful Failover, you must configure a Stateful Failover link (also known as the state link) to pass connection state information.
Shared with the Failover Link
Sharing a failover link is the best way to conserve interfaces. However, you must consider a dedicated interface for the state link and failover link, if you have a large configuration and a high traffic network.
Dedicated Interface
You can use a dedicated data interface (physical, redundant, or EtherChannel) for the state link. See Interface for the Failover Link for requirements for a dedicated state link, and Connecting the Failover Link for information about connecting the state link as well.
For optimum performance when using long distance failover, the latency for the state link should be less than 10 milliseconds and no more than 250 milliseconds. If latency is more than 10 milliseconds, some performance degradation occurs due to retransmission of failover messages.
Avoiding Interrupted Failover and Data Links
We recommend that failover links and data interfaces travel through different paths to decrease the chance that all interfaces fail at the same time. If the failover link is down, the ASA can use the data interfaces to determine if a failover is required. Subsequently, the failover operation is suspended until the health of the failover link is restored.
See the following connection scenarios to design a resilient failover network.
Scenario 1—Not Recommended
If a single switch or a set of switches are used to connect both failover and data interfaces between two ASAs, then when a switch or inter-switch-link is down, both ASAs become active. Therefore, the following two connection methods shown in the following figures are NOT recommended.
Scenario 2—Recommended
We recommend that failover links NOT use the same switch as the data interfaces. Instead, use a different switch or use a direct cable to connect the failover link, as shown in the following figures.
Scenario 3—Recommended
If the ASA data interfaces are connected to more than one set of switches, then a failover link can be connected to one of the switches, preferably the switch on the secure (inside) side of network, as shown in the following figure.
Scenario 4—Recommended
The most reliable failover configurations use a redundant interface on the failover link, as shown in the following figures.
MAC Addresses and IP Addresses in Failover
When you configure your interfaces, you can specify an active IP address and a standby IP address on the same network. Generally, when a failover occurs, the new active unit takes over the active IP addresses and MAC addresses. Because network devices see no change in the MAC to IP address pairing, no ARP entries change or time out anywhere on the network.
| Note | Although recommended, the standby address is not required. Without a standby IP address, the active unit cannot perform network tests to check the standby interface health; it can only track the link state. You also cannot connect to the standby unit on that interface for management purposes. | 
The IP address and MAC address for the state link do not change at failover.
Active/Standby IP Addresses and MAC Addresses
For Active/Standby Failover, see the following for IP address and MAC address usage during a failover event:
- 
                                       						
                                       The active unit always uses the primary unit's IP addresses and MAC addresses.
- 
                                       						
                                       When the active unit fails over, the standby unit assumes the IP addresses and MAC addresses of the failed unit and begins passing traffic.
- 
                                       						
