---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721-4
title: "c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721.md
source_anchor: ""
source_lines: [203, 282]
sha256: 12b568c24439ede16672392a515038ca9bfea9dc494b46dae71ca957e46fe77b
---

# c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721

                                       When the failed unit comes back online, and you enabled the preempt option, it resumes the failover group.
Active/active configurations implement additional address management:
Virtual MAC addresses provide specialized functionality:
The ASA has multiple methods to configure virtual MAC addresses. We recommend using only one method. If you set the MAC address using multiple methods, the MAC address used depends on many variables, and might not be predictable. Manual methods include the interface mode mac-address command, the failover mac address command, and for Active/Active failover, the failover group mode mac address command, in addition to autogeneration methods described below.
In multiple context mode, you can configure the ASA to generate virtual active and standby MAC addresses automatically for shared interfaces, and these assignments are synced to the secondary unit (see the mac-address auto command). For non-shared interfaces, you can manually set the MAC addresses for Active/Standby mode (Active/Active mode autogenerates MAC addresses for all interfaces).
For Active/Active failover, virtual MAC addresses are always used, either with default values or with values you can set per interface.
MAC address table updates occur during failover events:
During failover, the device that is designated as the new active device generates multicast packets for each MAC address entry in the MAC table and sends them to all the bridge group interfaces. This action prompts the upstream switches in the bridge group to update their routing tables with the new active device's interface to ensure accurate traffic forwarding.
The time taken to generate multicast packets and update the routing tables of the upstream switches depends on the number of entries in the MAC address table and the number of bridge group interfaces. Use the show failover statistics state-switch-delay command to display statistics related to the delays encountered during failover events.
Stateless and Stateful Failover
The ASA supports two types of failover, stateless and stateful for both the Active/Standby and Active/Active modes.
| Note | Some configuration elements for clientless SSL VPN (such as bookmarks and customization) use the VPN failover subsystem, which is part of Stateful Failover. You must use Stateful Failover to synchronize these elements between the members of the failover pair. Stateless failover is not recommended for clientless SSL VPN. | 
Stateless Failover
When a failover occurs, all active connections are dropped. Clients need to reestablish connections when the new active unit takes over.
| Note | Some configuration elements for clientless SSL VPN (such as bookmarks and customization) use the VPN failover subsystem, which is part of Stateful Failover. You must use Stateful Failover to synchronize these elements between the members of the failover pair. Stateless (regular) failover is not recommended for clientless SSL VPN. | 
Stateful failover
A stateful failover is a high availability mechanism that
- 
                                       
                                       continuously passes per-connection state information from the active unit to the standby unit
- 
                                       
                                       maintains connection information after a failover occurs, and
- 
                                       
                                       allows supported end-user applications to continue communication sessions without reconnecting.
Stateful failover operation
When Stateful Failover is enabled, the active unit continually passes per-connection state information to the standby unit, or in Active/Active failover, between the active and standby failover groups. After a failover occurs, the same connection information is available at the new active unit. Supported end-user applications are not required to reconnect to keep the same communication session.
Supported features
Supported features are high availability capabilities that
- 
                                          
                                          enable state information to be passed to the standby ASA device during Stateful Failover
- 
                                          
                                          maintain connection states and critical information for seamless traffic flow, and
- 
                                          
                                          ensure minimal disruption during failover events.
State information types
For Stateful Failover, these state information types are passed to the standby device:
- 
                                          
                                          NAT translation table.
- 
                                          
                                          TCP and UDP connections and states. Other types of IP protocols, and ICMP, are not parsed by the active unit, because they get established on the new active unit when a new packet arrives.
- 
                                          
                                          The HTTP connection table (unless you enable HTTP replication).
- 
                                          
                                          The HTTP connection states (if HTTP replication is enabled)—By default, the ASA does not replicate HTTP session information when Stateful Failover is enabled. We suggest that you enable HTTP replication.
- 
                                          
                                          SCTP connection states are included. SCTP inspection stateful failover is best effort. If any SACK packets are lost during failover, the new active unit drops all other out-of-order packets in the queue until the missing packet is received.
- 
                                          
                                          The ARP table
- 
                                          
                                          The Layer 2 bridge table (for bridge groups)
- 
                                          
                                          The ISAKMP and IPsec SA table
- 
                                          
                                          GTP PDP connection database
- 
                                          
                                          SIP signaling sessions and pin holes.
- 
                                          
                                          ICMP connection state—ICMP connection replication is enabled only if the respective interface is assigned to an asymmetric routing group.
- 
                                          
                                          Static and dynamic routing tables—Stateful Failover participates in dynamic routing protocols, like OSPF and EIGRP, so routes that are learned through dynamic routing protocols on the active unit are maintained in a Routing Information Base (RIB) table on the standby unit. Upon a failover event, packets travel normally with minimal disruption to traffic because the active secondary unit initially has rules that mirror the primary unit. Immediately after failover, the re-convergence timer starts on the newly active unit. Then the epoch number for the RIB table increments. During re-convergence, OSPF and EIGRP routes become updated with a new epoch number. Once the timer is expired, stale route entries (determined by the epoch number) are removed from the table. The RIB then contains the newest routing protocol forwarding information on the newly active unit. Note 
 Routes are synchronized only for link-up or link-down events on an active unit. If the link goes up or down on the standby unit, dynamic routes sent from the active unit may be lost. This is normal, expected behavior. 
- 
                                          
