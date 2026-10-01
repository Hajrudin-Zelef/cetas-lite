---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-89
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [12385, 12555]
sha256: eb98d03096b772dc1129dc711019419e86598dc664ccfcf0e4be9de30044fc5b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 The make-before-break mechanism uses switching and deletion delay timers to
                 prevent temporary traffic interruptions. When the two timers are configured, the
                 system switches traffic to a new CR-LSP after the switching delay time, and then
                 deletes the original CR-LSP after the deletion delay time.


4.3 Configuration Precautions for MPLS TE

4.4 Default Settings for MPLS TE
                 Table 4-11 describes the default settings for MPLS TE.

                 Table 4-11 Default settings for MPLS TE

                  Parameter                                 Default Setting

                  MPLS TE                                   Disabled

                  RSVP-TE                                   Disabled

                  Resource reservation style                SE

                  CSPF                                      Disabled

                  CSPF tie-breaking policy                  Randomly selected

                  Metric type used for tunnel path          TE
                  selection

                  Affinity attribute of a tunnel            0x0, with the mask 0x0

                  Maximum reservable link bandwidth         0kbit/s

                  Maximum number of hops on a CR-           32
                  LSP

                  Tunnel priority                           7 for both the setup and holding
                                                            priority values

                  Tunnel re-optimization                    Disabled


Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                          211
MPLS Configuration
MPLS Configuration                                                                       4 MPLS TE Configuration


                  Parameter                                          Default Setting

                  Route and label recording                          Disabled

                  Period after which the down status of              0s
                  a TE tunnel is responded




4.5 Configuring Static MPLS TE Tunnels

4.5.1 Enabling MPLS TE
Context
                 A static MPLS TE tunnel is created using a static CR-LSP. The static CR-LSP
                 configuration includes manually specifying labels and reserving resources, but
                 static CR-LSP setup does not involve a signaling protocol or MPLS control packet
                 exchange, thereby consuming less resources. In addition, the setup does not
                 involve IGP TE extension or CSPF. However, a static MPLS TE tunnel does not
                 dynamically adjust to network topology changes, and therefore is generally used
                 on small networks with simple topologies.
                 Enabling MPLS TE is a prerequisite for configuring a static MPLS TE tunnel.
                 Perform the following configuration on each node of an MPLS TE tunnel.

                         NOTE

                        After MPLS TE is disabled in the interface view, all CR-LSPs on the interface become down.
                        After MPLS TE is disabled in the MPLS view, MPLS TE is disabled on all interfaces, and all
                        CR-LSPs are deleted.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Configure an LSR ID for the local node.
                 mpls lsr-id lsr-id

                 When configuring an LSR ID, note the following:
                 ●      LSR IDs must be set before you run other MPLS commands.
                 ●      LSR IDs can only be manually configured, and do not have default values.
                 ●      Using the address of a loopback interface as an LSR ID is recommended.
         Step 3 Enable MPLS and enter the MPLS view.
                 mpls

         Step 4 Enable MPLS TE globally on the local node.
                 mpls te

                 To enable MPLS TE on an interface, you must first enable MPLS TE globally in the
                 MPLS view.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      212
MPLS Configuration
MPLS Configuration                                                                  4 MPLS TE Configuration


         Step 5 Return to the system view.
                 quit

         Step 6 Enter the interface view.
                 interface interface-type interface-number

         Step 7 Switch the interface working mode from Layer 2 to Layer 3.
                 undo portswitch

                 Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                 S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from
                 Layer 2 mode to Layer 3 mode using the undo portswitch command.
                 Determine whether to perform this step based on the current interface mode.
         Step 8 Enable MPLS on the interface.
                 mpls

         Step 9 Enable MPLS TE on the interface.
                 mpls te

        Step 10 (Optional) Configure link bandwidth.
                 Plan link bandwidth before you perform this procedure. The reserved bandwidth
                 must be higher than or equal to the bandwidth required by MPLS TE traffic.
                 In real-world applications, link bandwidth attributes only need to be configured on
                 the outbound interfaces that reside on a TE tunnel link and have specific
                 bandwidth requirements.
                 ●      Configure the maximum reservable link bandwidth.
                        mpls te bandwidth max-reservable-bandwidth max-bw-value
                        This command sets the bandwidth reserved for an MPLS TE tunnel on the
                        local link. By default, the maximum reservable bandwidth is not set. With the
                        default configuration, when the ingress of an MPLS TE tunnel initiates a
                        request to establish a CR-LSP with bandwidth constraints, the required CR-LSP
                        bandwidth (definitely greater than 0 kbit/s) may exceed the maximum
                        reservable link bandwidth. If this is the case, the CR-LSP fails to be
                        established.
                 ●      Configure BC bandwidth for the link.
                        mpls te bandwidth bc0 bc0-bw-value

                         NOTE

                        ● Ensure that the maximum reservable link bandwidth does not exceed the actual physical
                          link bandwidth. You are advised to set the maximum reservable link bandwidth to be
                          less than or equal to 80% of the actual physical link bandwidth.
                        ● Ensure that the BC0 bandwidth does not exceed the maximum reservable link
                          bandwidth.

                 ----End

4.5.2 Configuring MPLS TE Tunnel Interfaces
Context
                 To set up a static MPLS TE tunnel and customize tunnel attributes, you need to
                 first configure a tunnel interface. An MPLS TE tunnel interface is mainly

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 213
MPLS Configuration
MPLS Configuration                                                                        4 MPLS TE Configuration


                 responsible for setting up and maintaining an MPLS TE tunnel, as well as
                 controlling how packets are forwarded across the tunnel.

                         NOTE

                        MPLS TE tunnels forward MPLS packets, not IP packets. As such, IP forwarding-related
                        commands, such as the ip verify source-address command (for verifying source addresses),
                        are ineffective on tunnel interfaces.

