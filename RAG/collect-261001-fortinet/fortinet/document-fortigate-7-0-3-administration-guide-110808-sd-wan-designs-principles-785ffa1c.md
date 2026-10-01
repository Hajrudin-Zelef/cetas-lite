---
id: collect-261001-fortinet/fortinet/document-fortigate-7-0-3-administration-guide-110808-sd-wan-designs-principles-785ffa1c
title: "document-fortigate-7-0-3-administration-guide-110808-sd-wan-designs-principles-785ffa1c"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "latency"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-0-3-administration-guide-110808-sd-wan-designs-principles-785ffa1c.md
source_anchor: ""
source_lines: [1, 38]
sha256: cd46cea95d3a4d7292d067efc87f1dc934cdaf38c20f80432c58f3bada162a5d
---

# document-fortigate-7-0-3-administration-guide-110808-sd-wan-designs-principles-785ffa1c

SD-WAN designs principles
SD-WAN designs principles
The Five-pillar approach, described in the SD-WAN / SD-Branch Architecture for MSSPs guide, is recommended when designing a secure SD-WAN solution.
| Pillar | Overview | 
|---|---|
| Underlay | Choose the WAN links to use. | 
| Overlay | Choose the topology to interconnect your sites. | 
| Routing | Choose how to propagate routes between your sites. | 
| Security | Choose how to protect each of the available paths. | 
| SD-WAN | Choose the strategy used to pick one of the available paths. | 
Underlay
Determine the WAN links that will be used for the underlay network, such as your broadband link, MPLS, 4G/5G LTE connection, and others.
For each link, determine the bandwidth, quality and reliability (packet loss, latency, and jitter), and cost. Use this information to determine which link to prefer, what type of traffic to send across the each link, and to help you the baselines for health-checks.
Overlay
VPN overlays are needed when traffic must travel across multiple sites. These are usually site-to-site IPsec tunnels that interconnect branches, datacenters, and the cloud, forming a hub-and-spoke topology.
The management and maintenance of the tunnels should be considered when determining the overlay network requirements. Manual tunnel configuration might be sufficient in a small environment, but could become unmanageable as the environment size increases. ADVPN can be used to help scale the solution; see ADVPN for more information.
Routing
Traditional routing designs manipulate routes to steer traffic to different links. SD-WAN uses traditional routing to build the basic routing table to reach different destinations, but uses SD-WAN rules to steer traffic. This allows the steering to be based on criteria such as destination, internet service, application, route tag, and the health of the link. Routing in an SD-WAN solution is used to identify all possible routes across the underlays and overlays, which the FortiGate balances using ECMP.
In the most basic configuration, static gateways that are configured on an SD-WAN member interface automatically provide the basic routing needed for the FortiGate to balance traffic across the links. As the number of sites and destinations increases, manually maintaining routes to each destination becomes difficult. Using dynamic routing to advertise routes across overlay tunnels should be considered when you have many sites to interconnect.
Security
Security involves defining policies for access control and applying the appropriate protection using the FortiGate's NGFW features. Efficiently grouping SD-WAN members into SD-WAN zones must also be considered. Typically, underlays provide direct internet access and overlays provide remote internet or network access. Grouping the underlays together into one zone, and the overlays into one or more zones could be an effective method.
SD-WAN
The SD-WAN pillar is the intelligence that is applied to traffic steering decisions. It is comprised of four primary elements:
- 
                                                    SD-WAN zones SD-WAN is divided into zones. SD-WAN member interfaces are assigned to zones, and zones are used in policies as source and destination interfaces. You can define multiple zones to group SD-WAN interfaces together, allowing logical groupings for overlay and underlay interfaces. Routing can be configured per zone. See SD-WAN zones.
- 
                                                    SD-WAN members Also called interfaces, SD-WAN members are the ports and interfaces that are used to run traffic. At least one interface must be configured for SD-WAN to function. See Configuring the SD-WAN interface.
- 
                                                    Performance SLAs Also called health-checks, performance SLAs are used to monitor member interface link quality, and to detect link failures. When the SLA falls below a configured threshold, the route can be removed, and traffic can be steered to different links in the SD-WAN rule. SLA health-checks use active or passive probing: 
  - 
                                                            Active probing requires manually defining the server to be probed, and generates consistent probing traffic.
  - 
                                                            Passive probing uses active sessions that are passing through firewall policies used by the related SD-WAN interfaces to derive health measurements. It reduces the amount of configuration, and eliminates probing traffic. See Passive WAN health measurement for details.
 See Performance SLA.
- 
                                                            
- 
                                                    SD-WAN rules Also called services, SD-WAN rules control path selection. Specific traffic can be dynamically sent to the best link, or use a specific route Rules control the strategy that the FortiGate uses when selecting the outbound traffic interface, the SLAs that are monitored when selecting the outgoing interface, and the criteria for selecting the traffic that adheres to the rule. When no SD-WAN rules match the traffic, the implicit rule applies. See SD-WAN rules.
