---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-99
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [13937, 14059]
sha256: 3056537099c103f80df83dafdf387537ee6f94650cf6c20614122a75fff328c4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      Run the display mpls te link-administration bandwidth-allocation
                        [ interface interface-type interface-number ] command to check information
                        about link bandwidth allocation.
                 ●      Run the display ospf [ process-id ] mpls-te [ area { area-id | area-id-uint } ]
                        [ self-originated ] command to check information about all TE LSAs in the
                        LSDB.
                 ●      Run the display isis traffic-eng advertisements [ lsp-id | local ] [ level-1 |
                        level-2 | level-1-2 ] [ process-id | vpn-instance vpn-instance-name ]
                        command to check IS-IS TE information.
                 ●      Run the display isis traffic-eng statistics [ process-id | vpn-instance vpn-
                        instance-name ] command to check IS-IS TE statistics.
                 ●      Run the display mpls te cspf destination ip-address [ bandwidth ct0 ct0-
                        bandwidth | tie-breaking { random | most-fill | least-fill } ] command to
                        check CSPF-computed paths that meet the tunnel bandwidth constraints or
                        CSPF arbitration policy.
                 ●      Run the display mpls te cspf tedb all command to check all information in
                        the TEDB.
                 ●      Run the display mpls rsvp-te command to check the RSVP-TE configuration.
                 ●      Run the display mpls rsvp-te established [ interface { interface-type
                        interface-number | interface-name } peer-ip-address ] command to check
                        interface-based RSVP-TE resource reservation and basic information about
                        LSPs passing through the interface.
                 ●      Run the display mpls rsvp-te statistics { global | interface { interface-type
                        interface-number | interface-name } } command to check RSVP-TE statistics.
                 ●      Run the display mpls te link-administration admission-control [ interface
                        { interface-type interface-number | interface-name } ] command to check
                        information about the CR-LSPs that pass through a specified link.
                 ----End

4.7.7 Example for Configuring Dynamic MPLS TE Tunnels
Networking Requirements
                 On the network shown in Figure 4-10, LSR1, LSR2, LSR3, and LSR4 run IS-IS and
                 belong to Level-2.
                 Configure RSVP-TE to establish a TE tunnel from LSR1 to LSR4, with the
                 bandwidth being 20 Mbit/s. The maximum reservable bandwidth for each link
                 along the TE tunnel is 100 Mbit/s, and the BC0 bandwidth is 100 Mbit/s.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              237
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 Figure 4-10 Network diagram of dynamic MPLS TE tunnels




Configuration Roadmap
                 The configuration roadmap is as follows:

                 1.     Configure an IP address for each interface and a loopback address to be used
                        as an MPLS LSR ID on each node.
                 2.     Configure IS-IS to ensure that nodes can reach each other over public network
                        routes.
                 3.     Configure an MPLS LSR ID and enable MPLS, MPLS TE, and MPLS RSVP-TE
                        globally on each node. Enable CSPF on the ingress.
                 4.     Enable MPLS, MPLS TE, and MPLS RSVP-TE on each interface.
                 5.     Configure the maximum reservable bandwidth and BC bandwidth for the link
                        on the outbound interface of each node along the tunnel.
                 6.     Create a tunnel interface on the ingress of each tunnel. Set the tunnel IP
                        address, tunneling protocol, destination address, tunnel bandwidth, tunnel ID,
                        and signaling protocol used to establish the tunnel.


Procedure
         Step 1 Configure interface IP addresses for the devices.

                 # Configure LSR1.
                 <HUAWEI> system-view
                 [HUAWEI] sysname LSR1
                 [LSR1] vlan batch 100
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] ip address 10.1.1.1 24
                 [LSR1-Vlanif100] quit
                 [LSR1] interface 10ge 1/0/1
                 [LSR1-10GE1/0/1] port link-type trunk
                 [LSR1-10GE1/0/1] port trunk allow-pass vlan 100
                 [LSR1-10GE1/0/1] quit
                 [LSR1] interface loopback 1
                 [LSR1-loopback1] ip address 1.1.1.9 32
                 [LSR1-loopback1] quit

                 The configurations of LSR2, LSR3, and LSR4 are similar to the configuration of
                 LSR1. For detailed configurations, see Configuration Scripts.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           238
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration


         Step 2 Configure IS-IS to advertise routes.

                 # Configure LSR1.
                 [LSR1] isis 1
                 [LSR1-isis-1] is-level level-2
                 [LSR1-isis-1] network-entity 00.0005.0000.0000.0001.00
                 [LSR1-isis-1] quit
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] isis enable 1
                 [LSR1-Vlanif100] quit
                 [LSR1] interface loopback 1
                 [LSR1-LoopBack1] isis enable 1
                 [LSR1-LoopBack1] quit

                 The configurations of LSR2, LSR3, and LSR4 are similar to the configuration of
                 LSR1. For detailed configurations, see Configuration Scripts.

                 # After the configuration is complete, check the IP routing table on each node.
                 The following example uses the command output on LSR1.
                 [LSR1] display ip routing-table
                 Proto: Protocol        Pre: Preference
                 Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                 ------------------------------------------------------------------------------
                 Routing Table : _public_
                        Destinations : 13        Routes : 13

                 Destination/Mask     Proto Pre Cost        Flags NextHop                          Interface

