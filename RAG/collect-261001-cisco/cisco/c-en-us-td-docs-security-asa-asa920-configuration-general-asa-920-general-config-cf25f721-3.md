---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721-3
title: "c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721.md
source_anchor: ""
source_lines: [142, 202]
sha256: af4aa9c34622a51be95023447ceea2197b41980491d9019e9deb76b8e9711346
---

# c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721

                                       must be configured to enable stateful failover functionality.
State link configuration requirements
To use Stateful Failover, you must configure a Stateful Failover link (also known as the state link) to pass connection state information.
Shared failover links
A shared failover link is an interface conservation method that enables multiple network functions to use the same failover connection, though dedicated interfaces for state and failover links should be considered for large configurations and high traffic networks.
Dedicated interface
A dedicated interface is a data interface configuration that uses a dedicated physical or EtherChannel interface specifically for the state link in failover deployments.
Performance considerations
You can use a dedicated data interface—such as a physical or EtherChannel interface—for the state link. For requirements about a dedicated state link, refer to Interface for the failover link. See the section on connecting the state link for additional information Connect the failover link.
For optimum performance with long-distance failover, the state link latency should be less than 10 milliseconds but no more than 250 milliseconds. If latency exceeds 10 milliseconds, performance might degrade due to retransmission of failover messages.
Avoiding Interrupted Failover and Data Links
We recommend that failover links and data interfaces travel through different paths to decrease the chance that all interfaces fail at the same time. If the failover link is down, the ASA can use the data interfaces to determine if a failover is required. Subsequently, the failover operation is suspended until the health of the failover link is restored.
See the following connection scenarios to design a resilient failover network.
Scenario 1—Not Recommended
If a single switch or a set of switches are used to connect both failover and data interfaces between two ASAs, then when a switch or inter-switch-link is down, both ASAs become active. Therefore, the following two connection methods shown in the following figures are NOT recommended.
Scenario 2—Recommended
We recommend that failover links NOT use the same switch as the data interfaces. Instead, use a different switch or use a direct cable to connect the failover link, as shown in the following figures.
Scenario 3—Recommended
If the ASA data interfaces are connected to more than one set of switches, then a failover link can be connected to one of the switches, preferably the switch on the secure (inside) side of network, as shown in the following figure.
MAC addresses and IP addresses in Failover
MAC addresses and IP addresses in Failover are network addressing mechanisms that
- 
                                    
                                    provide unique identification for network interfaces during failover scenarios
- 
                                    
                                    maintain network connectivity when primary devices become unavailable, and
- 
                                    
                                    ensure seamless traffic flow between active and standby configurations.
Address configuration types
When you configure your interfaces, you can specify an active IP address and a standby IP address on the same network. Generally, when a failover occurs, the new active unit takes over the active IP addresses and MAC addresses. Because network devices see no change in the MAC to IP address pairing, no ARP entries change or time out anywhere on the network.
| Note | Although recommended, the standby address is not required. Without a standby IP address, the active unit cannot perform network tests to check the standby interface health; it can only track the link state. You also cannot connect to the standby unit on that interface for management purposes. | 
The IP address and MAC address for the state link do not change at failover.
Active/standby configurations use specific address handling methods:
For Active/Standby Failover, see the following for IP address and MAC address usage during a failover event:
- 
                                       						
                                       The active unit always uses the primary unit's IP addresses and MAC addresses.
- 
                                       						
                                       When the active unit fails over, the standby unit assumes the IP addresses and MAC addresses of the failed unit and begins passing traffic.
- 
                                       						
                                       When the failed unit comes back online, it is now in a standby state and takes over the standby IP addresses and MAC addresses.
However, if the secondary unit boots without detecting the primary unit, then the secondary unit becomes the active unit and uses its own MAC addresses, because it does not know the primary unit MAC addresses. When the primary unit becomes available, the secondary (active) unit changes the MAC addresses to those of the primary unit, which can cause an interruption in your network traffic. Similarly, if you swap out the primary unit with new hardware, a new MAC address is used.
If you disable failover and set the failover configurations to a disabled state, you will need to manually resume failover, or reboot the device. It is recommended to use the command failover reset and resume the failover instead of rebooting the device. If you reload the standby unit with the failover configuration disabled, the standby unit boots up as the active unit and uses the primary unit's IP addresses and MAC addresses. This leads to duplicate IP addresses and causes network traffic disruptions. Use the command failover reset to enable failover and restore the traffic flow.
| Note | If you enable failover on a standalone device, the data interfaces go down at negotiation state of failover, interrupting traffic. | 
Virtual MAC addresses guard against this disruption, because the active MAC addresses are known to the secondary unit at startup, and remain the same in the case of new primary unit hardware. We recommend that you configure the virtual MAC address on both the primary and secondary units to ensure that the secondary unit uses the correct MAC addresses when it is the active unit, even if it comes online before the primary unit. If you do not configure virtual MAC addresses, you might need to clear the ARP tables on connected routers to restore traffic flow. The ASA does not send gratuitous ARPs for static NAT addresses when the MAC address changes, so connected routers do not learn of the MAC address change for these addresses.
For Active/Active failover, see the following for IP address and MAC address usage during a failover event:
- 
                                       						
                                       The primary unit autogenerates active and standby MAC addresses for all interfaces in failover group 1 and 2 contexts. You can also manually configure the MAC addresses if necessary, for example, if there are MAC address conflicts.
- 
                                       						
                                       Each unit uses the active IP addresses and MAC addresses for its active failover group, and the standby addresses for its standby failover group. For example, the primary unit is active for failover group 1, so it uses the active addresses for contexts in failover group 1. It is standby for the contexts in failover group 2, where it uses the standby addresses.
- 
                                       						
                                       When a unit fails over, the other unit assumes the active IP addresses and MAC addresses of the failed failover group and begins passing traffic.
- 
                                       						
