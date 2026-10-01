---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721-6
title: "c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721.md
source_anchor: ""
source_lines: [381, 458]
sha256: 2bdebebecc408ba503951e44d6b3a1c1f254d103e6ea943402bd3dd89e59049a
---

# c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721

                                       The additional heartbeat module works similarly to the control plane module but sends and receives heartbeat messages using the data plane transport infrastructure.
- 
                                       
                                       When the peer receives heartbeat packets in the data plane, a counter gets incremented.
- 
                                       
                                       If the heartbeat transfer in the control plane fails, the node checks the heartbeat counter in the data plane. If the counter is incrementing, then the peer is alive, and the cluster does not perform a failover in this situation.
| Note |  | 
Interface monitoring
Interface monitoring is a failover capability that
- 
                                       
                                       tracks the health of up to 1025 interfaces to detect interface failures
- 
                                       
                                       performs interface tests when hello messages are not received within 15 seconds, and
- 
                                       
                                       triggers failover when the number of failed interfaces meets the defined threshold.
Interface monitoring operation
You can monitor up to 1025 interfaces (in multiple context mode, divided between all contexts). You should monitor important interfaces. For example in multiple context mode, you might configure one context to monitor a shared interface: because the interface is shared, all contexts benefit from the monitoring.
When a unit does not receive hello messages on a monitored interface for 15 seconds (the default), it runs interface tests. (To change the period, see the failover polltime interface command, or for Active/Active failover, the polltime interface command) If one of the interface tests fails for an interface, but this same interface on the other unit continues to successfully pass traffic, then the interface is considered to be failed, and the ASA stops running tests.
If the threshold you define for the number of failed interfaces is met (see the failover interface-policy command, or for Active/Active failover, the interface-policy command), and the active unit has more failed interfaces than the standby unit, then a failover occurs. If an interface fails on both units, then both interfaces go into the "Unknown" state and do not count towards the failover limit defined by failover interface policy.
An interface becomes operational again if it receives any traffic. A failed ASA returns to standby mode if the interface failure threshold is no longer met.
If an interface has IPv4 and IPv6 addresses configured on it, the ASA uses the IPv4 addresses to perform the health monitoring. If an interface has only IPv6 addresses configured on it, then the ASA uses IPv6 neighbor discovery instead of ARP to perform the health monitoring tests. For the broadcast ping test, the ASA uses the IPv6 all nodes address (FE02::1).
| Note | If a failed unit does not recover and you believe it should not be failed, you can reset the state by entering the failover reset command. If the failover condition persists, however, the unit will fail again | 
Interface tests
Interface tests are monitoring mechanisms that:
- 
                                          
                                          evaluate interface operational status during failover scenarios
- 
                                          
                                          run sequentially with approximately 1.5 second duration per test by default, or 1/16 of the failover interface holdtime(see the failover polltime interface command, or for Active/Active failover, the interface-policy command), and
- 
                                          
                                          determine interface failure status on ASA units.
Interface test sequence
The ASA uses these interface tests in sequence:
- 
                                          
                                          Link Up/Down test—A test of the interface status. If the Link Up/Down test indicates that the interface is down, then the ASA considers it failed, and testing stops. If the status is Up, then the ASA performs the Network Activity test.
- 
                                          
                                          Network Activity test—A received network activity test. At the start of the test, each unit clears its received packet count for its interfaces. As soon as a unit receives any eligible packets during the test, then the interface is considered operational. If both units receive traffic, then testing stops. If one unit receives traffic and the other unit does not, then the interface on the unit that does not receive traffic is considered failed, and testing stops. If neither unit receives traffic, then the ASA starts the ARP test.
- 
                                          
                                          ARP test—A test for successful ARP replies. Each unit sends a single ARP request for the IP address in the most recent entry in its ARP table. If the unit receives an ARP reply or other network traffic during the test, then the interface is considered operational. If the unit does not receive an ARP reply, then the ASA sends a single ARP request for the IP address in the next entry in the ARP table. If the unit receives an ARP reply or other network traffic during the test, then the interface is considered operational. If both units receive traffic, then testing stops. If one unit receives traffic, and the other unit does not, then the interface on the unit that does not receive traffic is considered failed, and testing stops. If neither unit receives traffic, then the ASA starts the Broadcast Ping test.
- 
                                          
                                          Broadcast Ping test—A test for successful ping replies. Each unit sends a broadcast ping, and then counts all received packets. If the unit receives any packets during the test, then the interface is considered operational. If both units receive traffic, then testing stops. If one unit receives traffic, and the other unit does not, then the interface on the unit that does not receive traffic is considered failed, and testing stops. If neither unit receives traffic, then testing starts over again with the ARP test. If both units continue to receive no traffic from the ARP and Broadcast Ping tests, then these tests will continue running in perpetuity.
Interface status
Interface status is a monitoring indicator that
- 
                                          
                                          shows the current operational state of monitored interfaces in a failover configuration
- 
                                          
                                          tracks traffic reception and hello packet communication between peer units, and
- 
                                          
                                          determines whether an interface is actively monitored by the failover process.
Interface status types
Monitored interfaces can have these status types:
- 
                                          
                                          Unknown: Initial status. This status can also mean the status cannot be determined.
- 
                                          
                                          Normal: The interface is receiving traffic.
- 
                                          
                                          Testing: Hello messages are not heard on the interface for five poll times.
- 
                                          
                                          Link Down: The interface or VLAN is administratively down.
- 
                                          
