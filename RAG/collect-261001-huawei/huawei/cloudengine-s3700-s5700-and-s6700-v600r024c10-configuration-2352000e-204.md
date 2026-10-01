---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-204
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [29952, 30092]
sha256: 460466f8a8c89507368745dbbc459762c2d30d917c898a4710077c89685275fa
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

SVC VPWS Topology
                    In SVC VPWS, the method of specifying the outer label (identifying a public
                    network tunnel) is similar to that of LDP VPWS. The inner label is manually
                    configured during VC configuration, so PEs do not need the signaling protocol to
                    transmit VC labels. As such, the network topology model and packet exchange
                    process of SVC VPWS are similar to those of LDP VPWS.

                    When creating a static Layer 2 VC connection of SVC VPWS, you can specify the
                    tunnel type (LDP or TE) through a tunnel policy.

                    Static PWs support the tunnel types of LDP LSPs and CR-LSPs. By default, LDP
                    LSPs are used.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        478
VPN Configuration
VPN Configuration                                                                 5 VPWS Configuration


5.8.2 (Optional) Configuring a PW Template
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

         Step 5 Configure PW attributes. Perform the operations listed in Table 5-6 based on
                service requirements.

                    Table 5-6 Attributes in a PW template

                     Operation                                  Command

                     Specify the IP address of the remote       peer-address ip-address
                     end on a PW.

                     Enable the control word function.          control-word

                     Set the MTU.                               mtu mtu-value

                     Configure the tunnel policy used by a      tnl-policy policy-name
                     PW.



                    ----End

Follow-up Procedure
                    If you modify the attributes in a PW template, you need to make the modification
                    take effect and re-create the PW that references the PW template.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                       479
VPN Configuration
VPN Configuration                                                                                   5 VPWS Configuration

                    reset pw pw-template pw-template-name

                    Running this command may disconnect and re-establish the corresponding PWs. If
                    multiple PWs are referencing this template, system operations are affected.

5.8.3 Configuring an SVC VPWS Connection
Prerequisites
                    Before configuring an SVC VPWS connection, you have completed the following
                    tasks:
                    ●      Configure static routes or an IGP on PEs and Ps of the MPLS backbone
                           network to ensure IP connectivity.
                    ●      Configure basic MPLS functions on the PEs and Ps of the MPLS backbone
                           network.
                    ●      Establish tunnels between PEs based on tunnel policies. If no tunnel policy is
                           configured, LDP tunnels are established by default.

Context
                    Perform the following operations on the PEs at both ends of a PW.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Configure MPLS L2VPN and enter the MPLS L2VPN view.
                    mpls l2vpn

         Step 3 Return to the system view.
                    quit

         Step 4 Enter the AC interface view.
                    interface interface-type interface-number [ .subinterface-number ]

         Step 5 Switch the interface working mode to Layer 3.
                    undo portswitch

                    Determine whether to perform this step based on the current interface working
                    mode.
         Step 6 (Optional) Configure an L2VPN description for the AC interface.
                    mpls l2vpn description description-text

                    By default, no L2VPN description is configured.
         Step 7 Create a VPWS connection.
                    mpls static-l2vc { { destination ip-address | pw-template pw-template-name vc-id } * | destination ip-
                    address [ vc-id ] } transmit-vpn-label transmit-label-value receive-vpn-label receive-label-value [ tunnel-
                    policy tnl-policy-name | [ control-word | no-control-word ] | [ raw | tagged ] ] *

                    The default tunnel policy uses LDP tunnels for VPWS connections. To use other
                    types of tunnels, you can configure a tunnel policy and set the tunnel-policy
                    policy-name parameter to reference the tunnel policy. For details about how to
                    configure a tunnel policy, see Configuring a Tunnel Policy.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                               480
VPN Configuration
VPN Configuration                                                                 5 VPWS Configuration


         Step 8 (Optional) Configure the name of an L2VPN service so that you can maintain the
                L2VPN service on the NMS UI based on the specified name.
                    mpls l2vpn service-name service-name

                    By default, no L2VPN service name is configured in the system.
                    ----End

