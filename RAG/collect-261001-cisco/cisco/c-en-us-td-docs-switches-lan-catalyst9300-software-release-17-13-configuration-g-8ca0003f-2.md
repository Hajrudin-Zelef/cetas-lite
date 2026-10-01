---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-8ca0003f-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-8ca0003f"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-8ca0003f.md
source_anchor: ""
source_lines: [65, 135]
sha256: debe71e2043a70d62b99e87ab479331586f429fa2ef459b76b35da9066f74240
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-8ca0003f

                                    Route packets to specific traffic-engineered paths; you might need to route them to allow a specific quality of service (QoS) through the network.
PBR allows you to classify and mark packets at the edge of the network. PBR marks a packet by setting precedence value. The precedence value can be used directly by devices in the network core to apply the appropriate QoS to a packet, which keeps packet classification at your network edge.
For enabling PBR for IPv6, see the Enabling Local PBR for IPv6 section.
For enabling IPv6 PBR for an interface, see the Enabling IPv6 PBR on an Interface section.
Unsupported IPv6 Unicast Routing Features
The switch does not support these IPv6 features:
- 
                                    				
                                    IPv6 packets that are destined to site-local addresses.
- 
                                    				
                                    Tunneling protocols, such as IPv4-to-IPv6 or IPv6-to-IPv4.
- 
                                    				
                                    The switch as a tunnel endpoint supporting IPv4-to-IPv6 or IPv6-to-IPv4 tunneling protocols.
- 
                                    				
                                    IPv6 Web Cache Communication Protocol (WCCP).
IPv6 Feature Limitations
Because IPv6 is implemented in switch hardware, some limitations occurs due to the IPv6 compressed addresses in the hardware memory. This hardware limitation result in some loss of functionality and limits some features. For example, the switch cannot apply QoS classification on source-routed IPv6 packets in hardware.
IPv6 and Switch Stacks
The switch supports IPv6 forwarding across the stack and IPv6 host functionality on the active switch. The active switch runs the IPv6 unicast routing protocols and computes the routing tables. They receive the tables and create hardware IPv6 routes for forwarding. The active switch also runs all IPv6 applications.
If a new switch becomes the active switch, it recomputes the IPv6 routing tables and distributes them to the member switches. While the new active switch is being elected and is resetting, the switch stack does not forward IPv6 packets. The stack MAC address changes, which also change the IPv6 address. When you specify the stack IPv6 address with an extended unique identifier (EUI) by using the ipv6 address ipv6-prefix/prefix length eui-64 interface configuration command, the address is based on the interface MAC address. See the Configuring IPv6 Addressing and Enabling IPv6 Routing section.
If you configure the persistent MAC address feature on the stack and the active switch changes, the stack MAC address does not change for approximately 4 minutes.
These are the functions of IPv6 active switch and members:
- 
                                    				
                                    Active switch: 
  - 
                                          						
                                          runs IPv6 routing protocols
  - 
                                          						
                                          generates routing tables
  - 
                                          						
                                          distributes routing tables to member switches that use distributed Cisco Express Forwarding for IPv6
  - 
                                          						
                                          runs IPv6 host functionality and IPv6 applications
- 
                                          						
                                          
- 
                                    				
                                    Member switch: 
  - 
                                          						
                                          receives Cisco Express Forwarding for IPv6 routing tables from the active switch
  - 
                                          						
                                          programs the routes into hardware
 Note 
 IPv6 packets are routed in hardware across the stack if the packet does not have exceptions (IPv6 Options) and the switches in the stack have not run out of hardware resources. 
 
  - 
                                          						
                                          flushes the Cisco Express Forwarding for IPv6 tables on active switch re-election
- 
                                          						
                                          
Default IPv6 Configuration
| Table 1. Default IPv6 Configuration |  | 
|---|---|
| Feature | Default Setting | 
|---|---|
| SDM template | Default is advance template | 
| IPv6 routing | Disabled globally and on all interfaces | 
| Cisco Express Forwarding for IPv6 or distributed Cisco Express Forwarding for IPv6 | Disabled (IPv4 Cisco Express Forwarding and distributed Cisco Express Forwarding are enabled by default) | 
| IPv6 addresses | None configured | 
| Note | When IPv6 routing is enabled, Cisco Express Forwarding for IPv6 and distributed Cisco Express Forwarding for IPv6 are automatically enabled. |
