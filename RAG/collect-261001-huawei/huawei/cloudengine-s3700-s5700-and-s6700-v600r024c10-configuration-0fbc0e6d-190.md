---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-190
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [27561, 27716]
sha256: 73e7f03e273c45eca1c61f168cb2892363a06af8bdeff7266353656643b83531
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

         Step 6 (Optional) Choose one of the following methods to set the Node protection flag
                in messages sent by the tunnel.
                 ●      Set the Node protection flag in messages sent by the current tunnel.
                        mpls te frr { node-protection | no-node-protection }

                 ●      Set the Node protection flag in messages sent by all P2P RSVP-TE tunnels.
                        quit
                        mpls
                        mpls te frr { node-protection | no-node-protection }




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        460
MPLS Configuration
MPLS Configuration                                                                         4 MPLS TE Configuration


                         NOTE

                        By default, when a Huawei device functions as the ingress of a tunnel in facility backup
                        mode, the device always sets the Node protection flag to 1 in sent messages, indicating
                        that node protection is desired.
                        If link protection is required, run the mpls te frr no-node-protection command on the
                        ingress. If a PLR is a Huawei device, you also need to run the mpls te plr apply node-
                        protection-flag command on the PLR.
                        If node protection is required, you do not need to run this command on the ingress or run
                        the mpls te plr apply node-protection-flag on the PLR.
                        After the mpls te frr { node-protection | no-node-protection } command is run in the
                        MPLS view, the configuration takes effect on all P2P RSVP-TE tunnel interfaces for which
                        the mpls te frr { node-protection | no-node-protection } command is not run but TE FRR
                        is enabled.

                 ----End

4.27.4 (Optional) Configuring Automatic Bypass Tunnel Re-
optimization

Context
                 Network changes often cause the changes in optimal paths. Automatic bypass
                 tunnel re-optimization allows the system to re-optimize an automatic bypass
                 tunnel if an optimal path to the same destination is found due to some reasons,
                 such as the changes in the cost. In this manner, network resources are optimized.

                 Perform the following steps on the PLR of an automatic bypass tunnel:

                         NOTE

                        This configuration task is invalid for LSPs in the FRR-in-use state.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Configure automatic re-optimization for automatic bypass tunnels.
                 mpls te auto-frr reoptimization [ frequency interval ]

                 This function allows the system to periodically re-optimize automatic bypass
                 tunnels, eliminating the need for manual intervention and thereby saving
                 manpower.

         Step 4 (Optional) Configure manual re-optimization for automatic bypass tunnels.
                 return
                 mpls te reoptimization [ port-type port-num | TunnelName | auto-tunnel name auto-tunnel-name ]

                 This function immediately re-optimizes all tunnels marked with the re-
                 optimization attribute as soon as the command is executed. After performing a

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        461
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


                 manual re-optimization, the automatic re-optimization timer is reset and starts
                 counting anew.

                 ----End

4.27.5 Verifying the Configuration
Procedure
                 ●      Run the display interface tunnel [ interface-number ] command to check the
                        operating status of tunnel interfaces.
                 ●      Run the display mpls te tunnel [ verbose ] command to check MPLS TE
                        tunnel information.
                 ●      Run the display mpls te tunnel { bypass-inuse { inuse | not-exists | exists-
                        not-used } } [ verbose ] command to check information about TE FRR-
                        enabled tunnels.
                 ●      Run the display mpls te tunnel-interface [ tunnel interface-number | auto-
                        bypass-tunnel [ tunnel-name ] ] command to check detailed information
                        about a primary or bypass tunnel.
                 ●      Run the display mpls lsp command to check CR-LSP information.
                 ●      Run the display mpls te tunnel path command to check path information of
                        the primary and bypass tunnels on the local node.
                 ----End

4.27.6 Example for Configuring MPLS TE Auto FRR
Networking Requirements
                 On the network shown in Figure 4-50, establish a primary tunnel along the
                 explicit path LSR1 -> LSR2 -> LSR3 -> LSR4. Create a bypass tunnel on the ingress
                 LSR1 to protect LSR2, and a bypass tunnel on the transit node LSR2 to protect the
                 link LSR2 -> LSR3.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           462
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 Figure 4-50 Network diagram for MPLS TE Auto FRR




Precautions
                 During the configuration, note the following:

                 ●      LSR IDs must be set before you run other MPLS commands.
                 ●      LSR IDs can only be manually configured, and do not have default values.
                 ●      When specifying an LSR ID, you are advised to use a reachable loopback
                        interface address.
                 ●      To avoid loops in this scenario, ensure that all connected interfaces have STP
                        disabled and are removed from VLAN1. If STP is enabled and VLANIF
                        interfaces of switches are used to construct a Layer 3 ring network, an
                        interface on the network will be blocked. As a result, Layer 3 services on the
                        network cannot run properly.

Configuration Roadmap
                 The configuration roadmap is as follows:

                 1.     Assign an IP address to each interface on each node and configure a loopback
                        interface address as an MPLS LSR ID on each node. Configure OSPF to ensure
                        that public network routes between nodes are reachable.
                 2.     Configure MPLS TE Auto FRR and node protection in the MPLS view on the
                        ingress of the primary tunnel. Configure MPLS TE Auto FRR and link
                        protection in the MPLS view on the ingress of the bypass tunnel.
                 3.     Configure a dynamic primary MPLS TE tunnel and enable TE FRR on the
                        tunnel interface of the ingress of the primary tunnel.

Procedure
         Step 1 Configure interface IP addresses for the nodes.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           463
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration


