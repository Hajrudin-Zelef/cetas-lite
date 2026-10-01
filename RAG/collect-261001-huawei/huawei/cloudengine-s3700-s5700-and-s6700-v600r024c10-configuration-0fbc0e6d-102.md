---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-102
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [14323, 14515]
sha256: fde7a1bcaa02113a2801ff519b94efe895943e47ed9217256ce6d3e8aaf7753c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                         isis enable 1
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 4.4.4.9
                         mpls te signal-protocol rsvp-te
                         mpls te bandwidth ct0 20000
                         mpls te tunnel-id 1
                        #
                        return

                 ●      LSR2
                        #
                        sysname LSR2
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         network-entity 00.0005.0000.0000.0002.00
                         traffic-eng level-2
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif200
                         ip address 10.1.2.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls te bandwidth max-reservable-bandwidth 100000
                         mpls te bandwidth bc0 100000
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                         isis enable 1
                        #
                        return

                 ●      LSR3
                        #
                        sysname LSR3
                        #
                        vlan batch 100 200



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      243
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         traffic-eng level-2
                         network-entity 00.0005.0000.0000.0003.00
                        #
                        interface Vlanif100
                         ip address 10.1.3.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls te bandwidth max-reservable-bandwidth 100000
                         mpls te bandwidth bc0 100000
                         mpls rsvp-te
                        #
                        interface Vlanif200
                         ip address 10.1.2.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                         isis enable 1
                        #
                        return

                 ●      LSR4
                        #
                        sysname LSR4
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 4.4.4.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         traffic-eng level-2
                         network-entity 00.0005.0000.0000.0004.00
                        #
                        interface Vlanif100
                         ip address 10.1.3.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      244
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ip address 4.4.4.9 255.255.255.255
                         isis enable 1
                        #
                        return



4.8 Steering Traffic to an MPLS TE Tunnel

4.8.1 Understanding Traffic Steering to MPLS TE Tunnels
                 Unlike LDP LSPs, static or dynamic MPLS TE tunnels cannot automatically steer
                 traffic for forwarding. Instead, traffic must be explicitly directed to tunnels in a
                 specific manner. A tunnel can then forward traffic based on labels. The following
                 describes how to steer traffic to an MPLS TE tunnel:
                 ●      Static routing mode: applies to scenarios where the network topology is
                        simple or the network environment is stable.
                 ●      Automatic routing mode: applies to scenarios where the network topology is
                        complex or the network environment changes frequently.
                 ●      Tunnel policy mode: applies to scenarios where TE tunnels need to be
                        selected to transport VPN services.


Static Routing
                 The simplest way to steer traffic to an MPLS TE tunnel is to configure a static
                 route. This static route operates just like a common static route. The only
                 difference is that you need to configure a TE tunnel interface as the static route's
                 outbound interface.


Automatic Routing
                 In automatic routing mode, a node considers a TE tunnel as a logical link and has
                 it participate in IGP route calculation. The tunnel interface is used as the route
                 outbound interface. The tunnel is considered as a point-to-point (P2P) link, and its
                 metric value can be set. The following automatic routing modes are supported:

