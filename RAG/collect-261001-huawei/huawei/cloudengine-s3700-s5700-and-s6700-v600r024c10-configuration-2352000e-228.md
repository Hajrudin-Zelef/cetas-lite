---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-228
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [33538, 33701]
sha256: 5c8c1582a0bb72e8e5d50415ffb4dfa73ab92590c84c6a62c91b34c8e0091780
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

5.11.1 Configuring Static BFD for VPWS
Prerequisites
                    Before configuring static BFD for VPWS, you have completed the following tasks:
                    ●   Configure network layer parameters for devices to communicate.
                    ●   Configure PWs.
                    ●   If a remote peer cannot identify the BFD CV Type field value 0x08 used for
                        VCCV, run the mpls l2vpn vccv bfd-cv-negotiation command to change the
                        BFD CV Type field value carried in a Label Mapping message to be consistent
                        with that of the remote peer.
                    ●   The control word mode of static BFD for VPWS PWs does not support the BFD
                        CV types of 0x10 and 0x20. If the PW-negotiated CV values on both ends are
                        0x10 or 0x20, the BFD session cannot go up after negotiation.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          534
VPN Configuration
VPN Configuration                                                                                5 VPWS Configuration


Context
                    Configuring BFD for VPWS PW accelerates PW fault detection, resulting in fast
                    switching of upper-layer applications. Static BFD applies to small networks.

                            NOTE

                           ● If the status of a service PW is down, the BFD session can be established but cannot go
                             up.
                           ● BFD for PW must be configured or cancelled on PEs at both ends of a PW. If this
                             function is configured or cancelled only on one end, the PW status between both ends
                             will be inconsistent.
                           ● To modify the parameters of an existing BFD session, run the min-tx-interval, min-rx-
                             interval, and detect-multiplier commands.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enable BFD globally and enter the global BFD view.
                    bfd

         Step 3 Return to the system view.
                    quit

         Step 4 Create a BFD session for the VPWS PW.
                    bfd session-name bind pw interface { interface-name1 | interface-type1interface-number1 } [ secondary ]

                    The following describes the interface parameters used for BFD for service PWs and
                    BFD for management PWs (mPW).

                    ●      When a BFD session detects a service PW, the interface specified by interface
                           interface-type1 interface-number1 is the AC interface on which the PW
                           resides.
                    ●      When a BFD session detects an mPW, the interface specified by interface
                           interface-type1 interface-number1 is a loopback interface.

                    If the PW to be detected is a secondary PW, configure secondary.

         Step 5 Configure discriminators.
                    ●      Configure a local discriminator.
                           discriminator local discr-value

                    ●      Configure a remote discriminator.
                           discriminator remote discr-value

                            NOTE

                           For a BFD session, the local discriminator on one end must be the remote discriminator on
                           the other end.

                    ----End




Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                         535
VPN Configuration
VPN Configuration                                                                                      5 VPWS Configuration


5.11.2 Configuring Dynamic BFD for VPWS
Prerequisites
                    Before configuring dynamic BFD for VPWS, you have completed the following
                    tasks:
                    ●      Configure network layer parameters for devices to communicate.
                    ●      Configure PWs.
                    ●      If a remote peer cannot identify the BFD CV Type field value 0x08 used for
                           VCCV, run the mpls l2vpn vccv bfd-cv-negotiation command to change the
                           BFD CV Type field value carried in a Label Mapping message to be consistent
                           with that of the remote peer.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enable BFD globally and enter the global BFD view.
                    bfd

         Step 3 Return to the system view.
                    quit

         Step 4 Enter the PW template view.
                    pw-template pw-template-name

         Step 5 Enable the control word function.
                    control-word

         Step 6 Adjust BFD detection parameters.
                    bfd-detect [ detect-multiplier multiplier | min-rx-interval rx-interval | min-tx-interval tx-interval ] *

                               NOTE

                           The parameters are described as follows:
                           ●     detect-multiplier multiplier indicates the BFD detection multiplier.
                           ●     min-rx-interval rx-interval indicates the Required Min Rx Interval (RMRI), which is the
                                 supported minimum interval at which the local device receives BFD control packets.
                           ●     min-tx-interval tx-interval indicates the Desired Min Tx Interval (DMTI), which is the
                                 desired minimum interval at which the local device transmits BFD control packets.
                           The BFD detection parameters actually used may be different from the ones configured:
                           ●     Actual local detection interval = Actual interval at which the local device receives BFD
                                 packets x Configured remote BFD detection multiplier
                           ●     Actual interval at which the local device receives BFD control packets = Max
                                 { Configured remote DMTI, Configured local RMRI }
                           ●     Actual interval at which the local device transmits BFD control packets = Max
                                 { Configured local DMTI, Configured remote RMRI }

         Step 7 Return to the system view.
                    quit

         Step 8 Enter the AC interface view.
                    interface interface-type interface-number


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                                 536
VPN Configuration
VPN Configuration                                                                                   5 VPWS Configuration


         Step 9 Switch the interface working mode to Layer 3.
                    undo portswitch

                    Determine whether to perform this step based on the current interface working
                    mode.

        Step 10 Create a BFD session for VPWS PWs.
                    mpls l2vpn pw bfd [ detect-multiplier multiplier | min-rx-interval rx-interval | min-tx-interval tx-
                    interval ] * [ remote-vcid vc-id ] [ secondary ]

                    If the PW to be detected is a secondary PW, configure secondary.

                    ----End


Verifying the Configuration
                    After configuring BFD for VPWS PWs, check the configuration.
                    ●    Run the display bfd session command to check BFD session information.
                    ●    Run the display mpls l2vc interface interface-type interface-number
                         command to check dynamic BFD information.

