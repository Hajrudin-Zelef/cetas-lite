---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf-3
title: "c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf.md
source_anchor: ""
source_lines: [143, 204]
sha256: ecbfeb8c564dd7ea37e9e508e4513be6ad327703cc8bc95d0b57cfee52b8f135
---

# c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf

                                       When the failed unit comes back online, it is now in a standby state and takes over the standby IP addresses and MAC addresses.
However, if the secondary unit boots without detecting the primary unit, then the secondary unit becomes the active unit and uses its own MAC addresses, because it does not know the primary unit MAC addresses. When the primary unit becomes available, the secondary (active) unit changes the MAC addresses to those of the primary unit, which can cause an interruption in your network traffic. Similarly, if you swap out the primary unit with new hardware, a new MAC address is used.
Virtual MAC addresses guard against this disruption, because the active MAC addresses are known to the secondary unit at startup, and remain the same in the case of new primary unit hardware. If you do not configure virtual MAC addresses, you might need to clear the ARP tables on connected routers to restore traffic flow. The ASA does not send gratuitous ARPs for static NAT addresses when the MAC address changes, so connected routers do not learn of the MAC address change for these addresses.
Active/Active IP Addresses and MAC Addresses
For Active/Active failover, see the following for IP address and MAC address usage during a failover event:
- 
                                       						
                                       The primary unit autogenerates active and standby MAC addresses for all interfaces in failover group 1 and 2 contexts. You can also manually configure the MAC addresses if necessary, for example, if there are MAC address conflicts.
- 
                                       						
                                       Each unit uses the active IP addresses and MAC addresses for its active failover group, and the standby addresses for its standby failover group. For example, the primary unit is active for failover group 1, so it uses the active addresses for contexts in failover group 1. It is standby for the contexts in failover group 2, where it uses the standby addresses.
- 
                                       						
                                       When a unit fails over, the other unit assumes the active IP addresses and MAC addresses of the failed failover group and begins passing traffic.
- 
                                       						
                                       When the failed unit comes back online, and you enabled the preempt option, it resumes the failover group.
Virtual MAC Addresses
The ASA has multiple methods to configure virtual MAC addresses. We recommend using only one method. If you set the MAC address using multiple methods, the MAC address used depends on many variables, and might not be predictable. Manual methods include the interface mode mac-address command, the failover mac address command, and for Active/Active failover, the failover group mode mac address command, in addition to autogeneration methods described below.
In multiple context mode, you can configure the ASA to generate virtual active and standby MAC addresses automatically for shared interfaces, and these assignments are synced to the secondary unit (see the mac-address auto command). For non-shared interfaces, you can manually set the MAC addresses for Active/Standby mode (Active/Active mode autogenerates MAC addresses for all interfaces).
For Active/Active failover, virtual MAC addresses are always used, either with default values or with values you can set per interface.
Intra- and Inter-Chassis Module Placement for the ASA Services Module
You can place the primary and secondary ASASMs within the same switch or in two separate switches.
Intra-Chassis Failover
If you install the secondary ASASM in the same switch as the primary ASASM, you protect against module-level failure.
Even though both ASASMs are assigned the same VLANs, only the active module takes part in networking. The standby module does not pass any traffic.
The following figure shows a typical intra-switch configuration.
Inter-Chassis Failover
To protect against switch-level failure, you can install the secondary ASASM in a separate switch. The ASASM does not coordinate failover directly with the switch, but it works harmoniously with the switch failover operation. See the switch documentation to configure failover for the switch.
For the best reliability of failover communications between ASASMs, we recommend that you configure an EtherChannel trunk port between the two switches to carry the failover and state VLANs.
For other VLANs, you must ensure that both switches have access to all firewall VLANs, and that monitored VLANs can successfully pass hello packets between both switches.
The following figure shows a typical switch and ASASM redundancy configuration. The trunk between the two switches carries the failover ASASM VLANs (VLANs 10 and 11).
| Note | ASASM failover is independent of the switch failover operation; however, ASASM works in any switch failover scenario. | 
If the primary ASASM fails, then the secondary ASASM becomes active and successfully passes the firewall VLANs.
If the entire switch fails, as well as the ASASM (such as in a power failure), then both the switch and the ASASM fail over to their secondary units.
Stateless and Stateful Failover
The ASA supports two types of failover, stateless and stateful for both the Active/Standby and Active/Active modes.
| Note | Some configuration elements for clientless SSL VPN (such as bookmarks and customization) use the VPN failover subsystem, which is part of Stateful Failover. You must use Stateful Failover to synchronize these elements between the members of the failover pair. Stateless failover is not recommended for clientless SSL VPN. | 
Stateless Failover
When a failover occurs, all active connections are dropped. Clients need to reestablish connections when the new active unit takes over.
| Note | Some configuration elements for clientless SSL VPN (such as bookmarks and customization) use the VPN failover subsystem, which is part of Stateful Failover. You must use Stateful Failover to synchronize these elements between the members of the failover pair. Stateless (regular) failover is not recommended for clientless SSL VPN. | 
Stateful Failover
When Stateful Failover is enabled, the active unit continually passes per-connection state information to the standby unit, or in Active/Active failover, between the active and standby failover groups. After a failover occurs, the same connection information is available at the new active unit. Supported end-user applications are not required to reconnect to keep the same communication session.
Supported Features
For Stateful Failover, the following state information is passed to the standby ASA:
-  
                                       		  
                                       NAT translation table.
-  
                                       		  
                                       TCP and UDP connections and states. Other types of IP protocols, and ICMP, are not parsed by the active unit, because they get established on the new active unit when a new packet arrives.
-  
                                       		  
                                       The HTTP connection table (unless you enable HTTP replication).
-  
                                       		  
                                       The HTTP connection states (if HTTP replication is enabled)—By default, the ASA does not replicate HTTP session information when Stateful Failover is enabled. We suggest that you enable HTTP replication.
-  
                                       		  
                                       SCTP connection states. However, SCTP inspection stateful failover is best effort. During failover, if any SACK packets are lost, the new active unit will drop all other out of order packets in the queue until the missing packet is received.
-  
                                       		  
