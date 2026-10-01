---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf-4
title: "c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf.md
source_anchor: ""
source_lines: [205, 298]
sha256: 35dc87e92cc052ef05d74cbdb1e9f3a1a4382cfd6121ba3e70ad16d31505a2a0
---

# c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf

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
                                       		  
                                       Static and dynamic routing tables—Stateful Failover participates in dynamic routing protocols, like OSPF and EIGRP, so routes that are learned through dynamic routing protocols on the active unit are maintained in a Routing Information Base (RIB) table on the standby unit. Upon a failover event, packets travel normally with minimal disruption to traffic because the active secondary unit initially has rules that mirror the primary unit. Immediately after failover, the re-convergence timer starts on the newly active unit. Then the epoch number for the RIB table increments. During re-convergence, OSPF and EIGRP routes become updated with a new epoch number. Once the timer is expired, stale route entries (determined by the epoch number) are removed from the table. The RIB then contains the newest routing protocol forwarding information on the newly active unit. 
 Note
 Routes are synchronized only for link-up or link-down events on an active unit. If the link goes up or down on the standby unit, dynamic routes sent from the active unit may be lost. This is normal, expected behavior. 
-  
                                       		  
                                       DHCP Server—DHCP address leases are not replicated. However, a DHCP server configured on an interface will send a ping to make sure an address is not being used before granting the address to a DHCP client, so there is no impact to the service. State information is not relevant for DHCP relay or DDNS.
-  
                                       		  
                                       Cisco IP SoftPhone sessions—If a failover occurs during an active Cisco IP SoftPhone session, the call remains active because the call session state information is replicated to the standby unit. When the call is terminated, the IP SoftPhone client loses connection with the Cisco Call Manager. This connection loss occurs because there is no session information for the CTIQBE hangup message on the standby unit. When the IP SoftPhone client does not receive a response back from the Call Manager within a certain time period, it considers the Call Manager unreachable and unregisters itself.
-  
                                       		  
                                       RA VPN—Remote access VPN end users do not have to reauthenticate or reconnect the VPN session after a failover. However, applications operating over the VPN connection could lose packets during the failover process and not recover from the packet loss.
Unsupported Features
For Stateful Failover, the following state information is not passed to the standby ASA:
-  
                                       		  
                                       The user authentication (uauth) table
-  
                                       		  
                                       TCP state bypass connections
-  
                                       		  
                                       Multicast routing.
-  
                                       		  
                                       State information for modules, such as the ASA FirePOWER module.
-  
                                       		  
                                       Selected clientless SSL VPN features: 
  -  
                                             				
                                             Smart Tunnels
  -  
                                             				
                                             Port Forwarding
  -  
                                             				
                                             Plugins
  -  
                                             				
                                             Java Applets
  -  
                                             				
                                             IPv6 clientless or Anyconnect sessions
  -  
                                             				
                                             Citrix authentication (Citrix users must reauthenticate after failover)
-  
                                             				
                                             
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
                                    		  
