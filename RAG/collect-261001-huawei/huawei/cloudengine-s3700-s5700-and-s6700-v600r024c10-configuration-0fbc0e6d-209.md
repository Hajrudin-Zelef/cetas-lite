---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-209
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [30345, 30516]
sha256: cb25dd1500d9191811b24828526dd76edb7dacc46620c7aa3b9ad08e8e3edfc5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      PE3
                        #
                        sysname PE3
                        #
                        ip vpn-instance vpn1
                         route-distinguisher 100:3
                         tnl-policy policy1
                         vpn-target 100:1 export-extcommunity
                         vpn-target 100:1 import-extcommunity
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 2.2.2.2
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls te cspf
                        #
                        explicit-path tope1
                         next hop 10.1.1.1
                         next hop 1.1.1.1
                        #
                        interface vlanif 100
                         ip address 10.1.1.2 255.255.255.252
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      502
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 1.1.1.1
                         mpls te tunnel-id 4
                         mpls te path explicit-path tope1
                         mpls te reserved-for-binding
                        #
                        bgp 100
                         peer 1.1.1.1 as-number 100
                         peer 1.1.1.1 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.1 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 1.1.1.1 enable
                         #
                         ipv4-family vpn-instance vpn1
                          import-route direct
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 10.1.1.0 0.0.0.3
                          network 2.2.2.2 0.0.0.0
                          mpls-te enable
                        #
                        tunnel-policy policy1
                         tunnel binding destination 1.1.1.1 te tunnel 1
                        #
                        return



4.29 Configuring Synchronization Between TE Tunnel
Status and BFD Session Status
Prerequisites
                 Before configuring synchronization between TE tunnel status and BFD session
                 status, complete the following tasks:

                 ●      4.7 Configuring Dynamic MPLS TE Tunnels
                 ●      4.24 Configuring Static BFD for CR-LSP, 4.25 Configuring Dynamic BFD for
                        CR-LSP, or 4.28 Configuring Static BFD for TE Tunnel


Context
                 After synchronization between TE tunnel status and BFD session status is enabled
                 and BFD for TE tunnel/BFD for CR-LSP is deployed in a BGP Add-Path scenario, if
                 the BFD session status is down, the device sets the TE tunnel status to down. In
                 this case, the tunnel will not be selected for forwarding services. Instead, BGP will
                 select a load balancing or backup route that has the same destination (with the
                 same prefix) as the down tunnel, preventing traffic forwarding failures.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        503
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 If BFD for TE tunnel is configured and the BFD session of a TE tunnel is down, the
                 device sets the tunnel's status to down.

                 If BFD for TE tunnel is configured and the BFD session of an RSVP-TE tunnel is up
                 or if BFD for TE tunnel is not configured, the device checks the status of BFD for
                 CR-LSP and the protocol status of an LSP. It then implements different actions
                 based on specific situations as follows:

                 1.     Both primary and backup LSPs are configured, the protocol status is up for
                        both of them, and BFD for CR-LSP is configured for both of them: If both of
                        the BFD sessions are down, the device sets the TE tunnel's status to down.
                 2.     Both primary and backup LSPs are configured but BFD for CR-LSP is
                        configured for only one of them: If the BFD session is down and the LSP
                        without BFD monitoring is also down, the device sets the TE tunnel's status to
                        down. Otherwise, the device does not set the tunnel's status to down.
                 3.     Only a primary LSP is configured, its protocol status is up, and BFD for CR-LSP
                        is configured: If the BFD session for the primary LSP is down, the device sets
                        the TE tunnel's status to down.
                 4.     Both primary and backup LSPs are configured, the protocol status of the
                        primary LSP is down, the protocol status of the backup LSP is up, and BFD for
                        CR-LSP is configured for the backup LSP: If the BFD session for the backup LSP
                        is down, the device sets the TE tunnel's status to down.
                 5.     Both primary and backup LSPs are configured, the protocol status is up for
                        both of them, and BFD for CR-LSP is configured for both of them: If the BFD
                        session status of an LSP is unknown, the previous status of the LSP is
                        considered as the current status of the LSP. If both the primary and backup
                        LSPs are down, the device sets the TE tunnel's status to down.

                 Perform the following configuration on the ingress of an MPLS TE tunnel.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Enable MPLS TE globally on the local node.
                 mpls te

         Step 4 Configure synchronization between P2P RSVP-TE tunnel status and BFD session
                status.
                 mpls rsvp-te tunnel track bfd

                 ----End


Verifying the Configuration
                 ●      Run the display tunnel all command to check the P2P RSVP-TE tunnel states
                        reported to TNLM.



Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          504
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration




4.30 Configuring Synchronization Between TE FRR and
CR-LSP Backup

4.30.1 Configuring Synchronization Between TE FRR and CR-
LSP Backup

