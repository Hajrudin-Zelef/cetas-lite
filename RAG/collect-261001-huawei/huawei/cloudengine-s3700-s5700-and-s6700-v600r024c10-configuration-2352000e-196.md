---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-196
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [28653, 28775]
sha256: 274232eadec908231a8eb39112665dc73393eb66474ae29ba3b000507393893f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

         Step 7 (Optional) Run either of the following commands as required:
                    ●    Configure the default VLAN of the main interface.
                         mpls l2vpn default vlan vlanid

                         By default, no default VLAN is configured for a main interface.
                    ●    Add a VLAN tag to the packets passing through the main interface.
                         mpls l2vpn vlan-stacking stack-vlan vlanid

                         By default, the system does not add a VLAN tag to a packet passing through
                         a main interface.
                          NOTE

                        ● If the remote PE is configured to accept only tagged packets, run the mpls l2vpn
                          default vlan command to configure the default VLAN of the main interface before
                          binding the local Ethernet interface to the VSI.
                        ● If the remote PE is configured to accept only double-tagged packets, run the mpls
                          l2vpn vlan-stacking stack-vlan command to configure the stacked VLAN of the main
                          interface before binding the local Ethernet interface to the VSI.

         Step 8 (Optional) Configure an L2VPN description for the AC interface.
                    mpls l2vpn description description-text

         Step 9 Create a VPWS connection.
                    mpls l2vc { ip-address | pw-template pw-template-name } * vc-id [ [ control-word | no-control-word ] |
                    [ raw | tagged ] | tunnel-policy policy-name | ignore-standby-state ] *

                          NOTE

                        ● The IDs of VCs using the same encapsulation type must be unique.
                        ● The raw and tagged parameters are available in this command only for Ethernet links.
                        ● If an Ethernet sub-interface is used, run the dot1q termination vid low-pe-vid
                          command in the interface view to configure the encapsulation type and VLAN ID of the
                          Ethernet sub-interface.
                        ● The default tunnel policy uses LDP tunnels for VPWS connections. To use other types of
                          tunnels, you can configure a tunnel policy and set the tunnel-policy policy-name
                          parameter to reference the tunnel policy. For details about how to configure a tunnel
                          policy, see Configuring a Tunnel Policy.

        Step 10 (Optional) Delete the VCCV byte (following the interface parameter) from a Label
                Mapping message.
                    undo interface-parameter-type vccv

                    This step is required when the local device communicates with a device sending a
                    Label Mapping packet that does not contain the VCCV byte following the interface
                    parameter.

        Step 11 (Optional) Configure the name of an L2VPN service so that you can maintain the
                L2VPN service on the network management system (NMS) user interface (UI)
                based on the specified name.
                    mpls l2vpn service-name service-name

        Step 12 Return to the system view.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            459
VPN Configuration
VPN Configuration                                                                   5 VPWS Configuration

                    quit

                    ----End

5.7.4 (Optional) Configuring Flow Label-based Load Balancing
Context
                    Packets of multiple data flows on the same PW carry the same VC labels, which
                    are encapsulated on a PE. As such, when these packets arrive at a P, the P
                    forwards them all along the same path, though there are multiple load-balancing
                    paths available.
                    To load-balance different data flows, configure flow-label-based load balancing on
                    the PE. After the configurations are complete, the PE adds flow labels following
                    private network labels when encapsulating data packets, and the P load-balances
                    these data flows based on flow labels.
                    On the network shown in Figure 5-23, two different data flows need to be
                    transmitted to the remote end through an L2VPN.
                    1.     PE1 calculates flow labels based on the source and destination IP addresses of
                           the two data flows. The following assumes that the flow labels are 1025 and
                           1026.
                    2.     PE1 adds the flow labels following the private network labels in the packets
                           of the two data flows.
                    3.     When the two data flows reach P1, P1 performs a hash calculation based on
                           the flow labels and maps the two data flows onto different paths. The
                           following assumes that the next hop of the data flow with flow label 1025 is
                           P2, and the next hop of the data flow with flow label 1026 is P3.
                    4.     When the two data flows reach PE2, the private network labels and flow
                           labels are sequentially removed. PE2 then forwards the two data flows to
                           their destination CEs based on their private network labels.

                    Figure 5-23 Flow-label-based load balancing




                    Flow-label-based load balancing applies to an L2VPN where there are multiple
                    links between Ps. Flow-label-based load balancing allows data flows on the same

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                           460
VPN Configuration
VPN Configuration                                                                              5 VPWS Configuration


                    VPN to be load-balanced along different paths based on flow labels, which
                    improves network resource utilization.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the AC interface view.
                    interface interface-type interface-number

         Step 3 Switch the interface working mode to Layer 3.
                    undo portswitch

                    Determine whether to perform this step based on the current interface working
                    mode.
         Step 4 Enable flow label-based load balancing for PWs on the interface.
                    mpls l2vpn flow-label { both | send | receive } [ secondary ] [ static ]

                    By default, flow label-based load balancing of PWs is disabled on an interface.

                          NOTE

