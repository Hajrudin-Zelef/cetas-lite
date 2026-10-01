---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-179
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [25975, 26107]
sha256: b8d3f61155220cce8f67f52483d1249897c7aeddde91ad4c1d87fcbe3a7a5662
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Procedure
                    ●      Configure ping to test IPv6 network connectivity.
                           ping ipv6 { [ -a source-ipv6-address | -c echo-number | { -s byte-number | -range [ [ min min-value |
                           max max-value | step step-value ] * ] } | -t timeout | { -tc traffic-class-value | -dscp dscp } | vpn-
                           instance vpn-instance-name | -m wait-time | -name | -nexthop nextHopAddr | -h hoplimit | { -brief |
                           [ -system-time | -ri | -detail ] * } | -p pattern | ignore-mtu ] * destination-ipv6-address [ -i { interface-
                           name | interface-type interface-number } ] [ ipv6-forwarding ] }
                    ●      Configure tracert to test an IPv6 network.
                           Test the fault position.
                           tracert ipv6 [ -f first-hop-limit | -m max-hop-limit | -p port-number | -q probes | -w timeout | vpn-
                           instance vpn-instance-name | -s size | -name | { -nexthop nextHopAddr | -passroute }* | -a source-
                           ipv6-address | { -tc tc | -dscp dscp } ] * host-name [ -i { ifName | ifType ifNum } ]

                    ----End

4.15.4 Checking the Integrated Route Statistics of IPv6 VPN
Instances

Procedure
                    ●      Run the display ipv6 routing-table { vpn-instance vpn-instance-name | all-
                           vpn-instance } statistics command to check the integrated route statistics of
                           an IPv6 VPN instance or all IPv6 VPN instances.
                    ----End




Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                                    410
VPN Configuration
VPN Configuration                                                                   5 VPWS Configuration




                                                5          VPWS Configuration


                    5.1 Overview of VPWS
                    5.2 Understanding VPWS
                    5.3 Configuration Precautions for VPWS
                    5.4 Default Settings for VPWS
                    5.5 Configuring CCC VPWS
                    5.6 Configuring BGP VPWS
                    5.7 Configuring LDP VPWS
                    5.8 Configuring SVC VPWS
                    5.9 Configuring Inter-AS VPWS
                    5.10 Configuring VPWS FRR
                    5.11 Configuring BFD for VPWS
                    5.12 Configuring PW Redundancy
                    5.13 Maintaining VPWS
                    5.14 Troubleshooting VPWS


5.1 Overview of VPWS
Definition
                    Virtual private wire service (VPWS) is a Layer 2 service bearer technology that
                    simulates the basic behaviors and characteristics of services, such as Ethernet,
                    synchronous optical network (SONET), and synchronous digital hierarchy (SDH)
                    services, on a packet switched network (PSN). VPWS simulates the traditional
                    leased lines of an IP network and provides asymmetric and low-cost digital data
                    network (DDN) services. For users at both ends of a virtual leased line, it is similar
                    to a traditional leased line. VPWS is a point-to-point virtual leased line technology
                    and supports almost all link layer protocols.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             411
VPN Configuration
VPN Configuration                                                                  5 VPWS Configuration


Purpose
                    As IP networks have developed, their scalability, upgradability, and interworking
                    capability have been greatly enhanced. Nevertheless, the scalability, upgradability,
                    and interworking capability of traditional communications networks are relatively
                    poor. When upgrading and expanding a traditional communications network, it is
                    important to determine whether to construct a traditional communications
                    network, to rely more on existing network resources, or to make use of public
                    network resources. VPWS enables traditional communications networks to
                    interwork with existing PSNs.

Benefits
                    VPWS has the following benefits:
                    ●   Extended network functions and service capabilities for carriers
                        Carriers can provide MPLS L2VPN services with just a single network. Carriers
                        can also use enhanced MPLS-related technologies, such as traffic engineering
                        (TE), to provide users with different classes of services to meet different
                        requirements.
                    ●   Higher scalability
                        On non-MPLS networks, virtual circuits (VCs) are used to provide L2VPN
                        services. For each VC, both the provider edges (PEs) and providers (Ps) on the
                        network need to maintain complete VC information. When the PEs of a
                        carrier connect to multiple customer edges (CEs), multiple VCs are required,
                        and information about these VCs must be maintained on both PEs and Ps.
                        VPWS uses label stacking to multiplex multiple VCs on a label switched path
                        (LSP). As such, Ps only need to maintain information about one LSP. This
                        improves the system scalability.
                    ●   Clearly defined administrative roles
                        On an MPLS L2VPN, carriers provide only Layer 2 connectivity. Users are
                        responsible for Layer 3 connectivity, such as routing. This implementation
                        prevents incorrect configurations from causing route flapping, which will
                        affect the stability of a carrier's network.
                    ●   Support for multiple protocols
                        Carriers provide only Layer 2 connectivity, but users can use any Layer 3
                        protocol, such as IPv4 and IPv6.
                    ●   Smooth network upgrade
                        The MPLS L2VPN is transparent to users. When a carrier upgrades the
                        network from a traditional L2VPN to an MPLS L2VPN, users do not need to
                        change any settings. Other than a brief period during the switchover where
                        there may be some data loss, user services are unaffected by the upgrade.


5.2 Understanding VPWS




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            412
VPN Configuration
VPN Configuration                                                                  5 VPWS Configuration


5.2.1 VPWS Implementation
Basic VPWS Architecture
                    The VPWS architecture consists of attachment circuits (ACs), pseudo wires (PWs),
                    and tunnels (network tunnels), as shown in Figure 5-1.

                    Figure 5-1 Basic VPWS architecture




