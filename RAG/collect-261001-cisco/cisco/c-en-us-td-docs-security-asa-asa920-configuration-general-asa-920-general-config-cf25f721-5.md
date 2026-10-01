---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721-5
title: "c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721.md
source_anchor: ""
source_lines: [283, 380]
sha256: fb9825e44d2c82a00ba0c97f8b6d94a3e1b04447d4f60a540fb9e75d05f4e6e2
---

# c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721

                                          DHCP Server—DHCP address leases are not replicated. However, a DHCP server configured on an interface will send a ping to make sure an address is not being used before granting the address to a DHCP client, so there is no impact to the service. State information is not relevant for DHCP relay or DDNS.
- 
                                          
                                          Cisco IP SoftPhone sessions—If a failover occurs during an active Cisco IP SoftPhone session, the call remains active because the call session state information is replicated to the standby unit. When the call is terminated, the IP SoftPhone client loses connection with the Cisco Call Manager. This connection loss occurs because there is no session information for the CTIQBE hangup message on the standby unit. When the IP SoftPhone client does not receive a response back from the Call Manager within a certain time period, it considers the Call Manager unreachable and unregisters itself.
- 
                                          
                                          RA VPN—Remote access VPN end users do not have to reauthenticate or reconnect the VPN session after a failover. However, applications operating over the VPN connection could lose packets during the failover process and not recover from the packet loss.
- 
                                          
                                          From all the connections, only established ones will be replicated on the Standby device.
Unsupported features
Unsupported features are device capabilities that
- 
                                          
                                          are not passed to the standby ASA in Stateful Failover
- 
                                          
                                          do not maintain state information during failover events, and
- 
                                          
                                          may require re-establishment after a failover occurs.
State information limitations
For Stateful Failover, the following state information is not passed to the standby ASA:
-  
                                          
                                          The user authentication (uauth) table
-  
                                          
                                          Multicast routing.
Bridge Group Requirements for Failover
There are special considerations for failover when using bridge groups.
Bridge Group Requirements for Appliances, ASAv
When the active unit fails over to the standby unit, the connected switch port running Spanning Tree Protocol (STP) can go into a blocking state for 30 to 50 seconds when it senses the topology change. To avoid traffic loss while the port is in a blocking state, you can configure one of the following workarounds depending on the switch port mode:
-  
                                    		  
                                    Access mode—Enable the STP PortFast feature on the switch: 
interface interface_id
  spanning-tree portfast
The PortFast feature immediately transitions the port into STP forwarding mode upon linkup. The port still participates in STP. So if the port is to be a part of the loop, the port eventually transitions into STP blocking mode.
-  
                                    		  
                                    Trunk mode—Block BPDUs on the ASA on a bridge group's member interfaces with an EtherType access rule. 
access-list id ethertype deny bpdu
access-group id in interface name1
access-group id in interface name2
Blocking BPDUs disables STP on the switch. Be sure not to have any loops involving the ASA in your network layout.
If neither of the above options are possible, then you can use one of the following less desirable workarounds that impacts failover functionality or STP stability:
-  
                                    		  
                                    Disable interface monitoring.
-  
                                    		  
                                    Increase interface holdtime to a high value that will allow STP to converge before the ASAs fail over.
-  
                                    		  
                                    Decrease STP timers to allow STP to converge faster than the interface holdtime.
Failover health monitoring
Failover health monitoring is a system capability that monitors each unit for overall health and interface health.
Health monitoring tests
The ASA performs tests to determine the state of each unit. This section includes information about how the ASA performs tests to determine the state of each unit.
Unit health monitoring
Unit health monitoring is a failover mechanism that
- 
                                       
                                       determines the health of peer units by monitoring the failover link with hello messages
- 
                                       
                                       sends LANTEST messages on each data interface when three consecutive hello messages are missed, and
- 
                                       
                                       initiates appropriate failover actions based on peer unit responsiveness.
Unit health monitoring behavior
The ASA determines the health of the other unit by monitoring the failover link with hello messages. If a unit does not receive three consecutive hello messages on the failover link, the sends LANTEST messages on each data interface, including the failover link, to validate whether the peer is responsive. For the Firepower 9300 and 4100 series, you can enable Bidirectional Forwarding Detection (BFD) monitoring, which is more reliable than hello messages. The action that the ASA takes depends on the response from the other unit. These are the possible actions:
- 
                                       
                                       If the ASA receives a response on the failover link, then it does not fail over.
- 
                                       
                                       If the ASA does not receive a response on the failover link, but it does receive a response on a data interface, then the unit does not failover. The failover link is marked as failed. You should restore the failover link as soon as possible because the unit cannot fail over to the standby while the failover link is down.
- 
                                       
                                       If the ASA does not receive a response on any interface, then the standby unit switches to active mode and classifies the other unit as failed.
Heartbeat module redundancy
Heartbeat module redundancy is a high availability feature that
- 
                                       
                                       sends heartbeat messages using the data plane transport infrastructure to supplement control plane heartbeat packets
- 
                                       
                                       increments a counter when the peer receives heartbeat packets in the data plane
- 
                                       
                                       prevents false failover or split-brain scenarios when the control plane is congested with traffic.
Heartbeat module redundancy behavior
Each unit in the HA periodically sends a broadcast keepalive heartbeat packet over the failover link. If the control plane is too busy handling traffic, sometimes the heartbeat packets do not reach the peers, or the peers do not process the heartbeat packets due to CPU overloading. When peers cannot communicate the keepalive status within the configurable timeout period, a false failover or split-brain scenario occurs.
The heartbeat module in the data plane helps to avoid the occurrence of false failover or split-brain due to traffic congestion in the control plane.
- 
                                       
