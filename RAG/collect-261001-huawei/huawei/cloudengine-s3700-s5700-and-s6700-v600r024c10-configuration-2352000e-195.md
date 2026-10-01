---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-195
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [28495, 28652]
sha256: f3a007b865f44cec9633e708ddb4abe8448b44231c5e0f61ee936abf2831ef6c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        Figure 5-22 Process of tearing down a PWE3 VPWS PW




                        Figure 5-22 shows the process of tearing down a single-segment PW.
                        a.   When the PW configuration is deleted from PE1, PE1 withdraws its VC
                             label and sends Label Withdraw and Label Release messages to PE2 in
                             succession.
                        b.   PE2 receives the Label Withdraw and Label Release messages from PE1,
                             withdraws its VC label, and sends a Label Release message to PE1.
                        c.   After PE1 and PE2 receive the Label Release message from each other,
                             the PW deletion is complete on PE1 and PE2.
                                  NOTE

                             The Label Withdraw message instructs a PE to withdraw the label. The Label Release
                             message is a response to the Label Withdraw message to inform the PE sending the
                             Label Withdraw message that the label has been withdrawn on the peer. To tear
                             down the PW faster, PE1 can send the Label Withdraw and Label Release messages in
                             succession.
                    ●   Extensions on the PWE3 control plane
                        –    Signaling extension
                             A way to send a Label Notification message has been added to LDP
                             signaling. It can just advertise the status but not tear down a signaling
                             connection. There is a signaling disconnection only when the PW
                             configuration is deleted or if the signaling is interrupted. This signaling
                             extension allows for fewer control packet exchanges, reduces the
                             signaling cost, and is compatible with the original LDP mode.
                        –    Other extensions
                             Other extensions on the control plane are as follows:

                             ▪   Mechanism for negotiating fragmentation

                             ▪   PW connectivity detection, such as VCCV, has been added, improving
                                 the fast network convergence capability and network reliability.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 456
VPN Configuration
VPN Configuration                                                                       5 VPWS Configuration


                    ●      Extensions on the PWE3 data plane
                           –     Real-time information extension
                           –     Bandwidth, jitter, and delay assurance of electrical signals
                           –     Retransmission of disordered packets

5.7.2 (Optional) Configuring a PW Template
Context
                    A PW template is a collection of common PW attributes. Configuring PWs using a
                    PW template reduces the configuration workload. For example, if multiple VPWS
                    connections with the same peer address, control word setting, tunnel policy, and
                    MTU need to be configured between two devices, you can define these attributes
                    in a PW template and use the PW template to configure these PW attributes on
                    different interfaces.
                    Perform the following operations on the PEs at both ends of a PW.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Configure MPLS L2VPN.
                    mpls l2vpn

         Step 3 Return to the system view.
                    quit

         Step 4 Create a PW template and enter the PW template view.
                    pw-template pwname

                    If some PW attributes configured using the mpls l2vc command on an interface
                    are different from those specified in the PW template, the PW attributes
                    configured using the mpls l2vc command take precedence.
         Step 5 Configure PW attributes. Perform the operations listed in Table 5-5 based on
                service requirements.

                    Table 5-5 Attributes in a PW template
                     Operation                                      Command

                     Specify the IP address of the remote           peer-address ip-address
                     end on a PW.

                     Enable the control word function.              control-word

                     Set the MTU.                                   mtu mtu-value

                     Configure the tunnel policy used by a          tnl-policy policy-name
                     PW.



                    ----End

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                           457
VPN Configuration
VPN Configuration                                                                        5 VPWS Configuration


Follow-up Procedure
                    If you modify the attributes in a PW template, you need to make the modification
                    take effect and re-create the PW that references the PW template.
                    reset pw pw-template pw-template-name

                    Running this command may disconnect and re-establish the corresponding PWs. If
                    multiple PWs are referencing this template, system operations are affected.

5.7.3 Configuring an LDP VPWS Connection
Prerequisites
                    Before configuring an LDP VPWS connection, you have completed the following
                    tasks:
                    ●      Configure static routes or an IGP on PEs and Ps of the MPLS backbone
                           network to ensure IP connectivity.
                    ●      Configure basic MPLS functions on the PEs and Ps of the MPLS backbone
                           network.
                    ●      Establish LDP sessions between PEs. If the PEs are indirectly connected,
                           establish remote LDP sessions between them.
                    ●      Establish tunnels between PEs based on tunnel policies. If no tunnel policy is
                           configured, LDP tunnels are established by default.

Context
                    LDP VPWS can be either PWE3-compatible VPWS or PWE3 VPWS.
                    ●      PWE3-compatible VPWS does not use Label Notification messages.
                    ●      PWE3 VPWS uses Label Notification messages. By default, PWE3 VPWS is
                           used.
                    Perform the following operations on the PEs at both ends of a PW.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Configure MPLS L2VPN and enter the MPLS L2VPN view.
                    mpls l2vpn

         Step 3 (Optional) Disable the function of sending L2VPN Label Request messages to the
                remote end.
                    mpls l2vpn no-request-message peer ip-address

                    By default, the device sends L2VPN Label Request messages to all remote ends.
         Step 4 Return to the system view.
                    quit

         Step 5 Enter the AC interface view.
                    interface interface-type interface-number [ .subinterface-number ]


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                             458
VPN Configuration
VPN Configuration                                                                                 5 VPWS Configuration


         Step 6 Switch the interface working mode to Layer 3.
                    undo portswitch

                    Determine whether to perform this step based on the current interface working
                    mode.

