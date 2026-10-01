---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-319
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [47512, 47641]
sha256: 6eccc47df1b6b3afdf871195975406124fc38226fcaf0ca912c6c886c7caa379
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the VSI view.
                    vsi vsi-name


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                    763
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration


         Step 3 (Optional) Configure the MTU for the VSI.
                    mtu mtu-value

                    By default, the MTU value for a VSI is 1500.
                    If the MTUs of the same VSI on two PEs are different, the two PEs cannot
                    establish a connection. Unlike the MTU configured in the interface view, the MTU
                    configured in the VSI view takes effect immediately and you do not need to run
                    the reset command.
         Step 4 (Optional) Configure a tunnel policy for the VSI.
                    tnl-policy policy-name

                    By default, no tunnel policy is configured for a VSI.
         Step 5 (Optional) Configure the description of the VSI.
                    description description

                    By default, no description is configured for a VSI.

                    ----End

6.17.2 Configuring MAC Withdraw
Context
                    On an LDP VPLS, BGP AD VPLS, or LDP+BGP AD VPLS network, if an AC fails or a
                    primary/secondary PW switchover occurs, you can configure MAC Withdraw to
                    enable the local PE to send MAC Withdraw messages to the remote PEs, instruct
                    the remote PEs to clear the MAC addresses in their VSIs, or enable the remote PEs
                    to forward MAC Withdraw messages. An LDP VPLS network supports MAC
                    Withdraw messages that carry the FEC 128 TLV, whereas a BGP AD VPLS network
                    supports MAC Withdraw messages that carry the FEC 129 TLV. When MAC
                    Withdraw messages are forwarded on an LDP+BGP AD VPLS network, translation
                    between the FEC 128 TLV and FEC 129 TLV is required.
                    In Figure 6-28, the PW from PE4 to PE1 is the primary one and the PW from PE4
                    to PE2 is the secondary one; the link from CE1 to PE1 is the primary one and the
                    link from CE1 to PE3 is the secondary one. The following examples use the failure
                    of the PW from PE4 to PE1 and failure of the AC from CE1 to PE1 to describe the
                    MAC Withdraw message forwarding process.
                    ●    When the PW from PE4 to PE1 fails, PE4 performs a primary/secondary PW
                         switchover and sends a MAC Withdraw message that carries the FEC 128 TLV
                         to PE2. After receiving the MAC Withdraw message, PE2 removes MAC
                         addresses in its VSI, translates the message to a MAC Withdraw message that
                         carries the FEC 129 TLV, and forwards the new MAC Withdraw message to
                         PE3. PE3 receives the MAC Withdraw message and clears the MAC addresses
                         in its VSI.
                         When the PW from PE4 to PE1 recovers, PE4 performs a primary/secondary
                         PW switchback based on the configured switchback policy and sends a MAC
                         Withdraw message that carries the FEC 128 TLV to PE1. After receiving the
                         MAC Withdraw message, PE1 removes MAC addresses in its VSI, translates the
                         message to a MAC Withdraw message that carries the FEC 129 TLV, and
                         forwards the new MAC Withdraw message to PE3. PE3 receives the MAC
                         Withdraw message and clears the MAC addresses in its VSI.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                     764
VPN Configuration
VPN Configuration                                                                            6 VPLS Configuration


                    ●    When the AC between CE1 and PE1 fails, PE1 sends a MAC Withdraw
                         message that carries the FEC 128 TLV to PE4 after PE1 detects the AC status
                         change. PE3 sends a MAC Withdraw message that carries the FEC 129 TLV to
                         PE2 after detecting the AC interface status change. After receiving the MAC
                         Withdraw message, PE2 removes MAC addresses in its VSI, translates the
                         message to a MAC Withdraw message that carries the FEC 128 TLV, and
                         forwards the new MAC Withdraw message to PE4. After PE4 receives the MAC
                         Withdraw messages sent from PE1 and PE2, PE4 clears the MAC addresses in
                         its VSI.
                         When the AC between CE1 and PE1 recovers, PE1 sends a MAC Withdraw
                         message that carries the FEC 128 TLV to PE4 after PE1 detects the AC status
                         change. PE3 sends a MAC Withdraw message that carries the FEC 129 TLV to
                         PE2 after detecting the AC interface status change. After receiving the MAC
                         Withdraw message, PE2 removes MAC addresses in its VSI, translates the
                         message to a MAC Withdraw message that carries the FEC 128 TLV, and
                         forwards the new MAC Withdraw message to PE4. After PE4 receives the MAC
                         Withdraw messages sent from PE1 and PE2, PE4 clears the MAC addresses in
                         its VSI.

                    Figure 6-42 Forwarding of MAC Withdraw messages on an LDP+BGP AD VPLS
                    network




                         NOTE

                    MAC Withdraw can be configured in either the VSI-LDP view or VSI view, but cannot be
                    configured in both views. For example, if you have configured MAC Withdraw in the VSI-LDP
                    view, this function cannot be configured in the VSI view. This function can be configured in the
                    VSI view only after the MAC Withdraw configuration in the VSI-LDP view is deleted. Similarly, if
                    you have configured MAC Withdraw in the VSI view, this function cannot be configured in the
                    VSI-LDP view.
                    MAC Withdraw applies only to LDP VPLS networks if configured in the VSI-LDP view, but applies
                    to LDP VPLS, BGP AD VPLS, and LDP+BGP AD VPLS networks if configured in the VSI view.


Procedure
                    ●    Configure a PE to send MAC Withdraw messages when the AC or PW status
                         on the PE changes to instruct the remote ends to clear the MAC addresses in

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      765
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


                        their VSIs. Two methods are available for configuring the MAC Withdraw
                        function: Method 1 applies to LDP VPLS, whereas method 2 applies to LDP
                        VPLS and BGP AD VPLS.
                        –   Method 1
                            i.     Enter the system view.
                                   system-view

                            ii.    Enter the VSI view.
                                   vsi vsi-name

                            iii.   Enter the VSI-LDP view.
                                   pwsignal ldp

