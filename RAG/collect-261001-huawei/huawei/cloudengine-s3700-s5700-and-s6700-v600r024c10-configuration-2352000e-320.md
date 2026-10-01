---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-320
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [47642, 47767]
sha256: 0dc75a39433a8ded70b1144f37a67a6315530e8268229a77bfd5abd104138f3d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                            iv.    Configure a VSI to instruct all its peers to delete all the MAC
                                   addresses learned from this VSI when an AC or PW fault occurs but
                                   the VSI remains up.
                                   mac-withdraw enable

                                   By default, MAC Withdraw is disabled.
                                   After MAC Withdraw is enabled, the MAC addresses learned from the
                                   remote end may be deleted from the local end when the AC or PW
                                   goes up or down and packet broadcasting may occur.
                            v.     (Optional) Enable the function of sending LDP MAC Withdraw
                                   messages to all peers when the status of the AC interface bound to
                                   the VSI changes.
                                   interface-status-change mac-withdraw enable

                                   By default, a VSI does not send LDP MAC Withdraw messages.
                            vi.    Return to the VSI view.
                                   quit

                            vii. (Optional) Enable the device in a VSI to remove MAC addresses
                                 except those associated with the PW through which the device
                                 receives MAC Withdraw messages with the TLV type 0x404.
                                   local-mac remove all-but-mine

                                   By default, the device removes the MAC addresses of all PWs and AC
                                   interfaces after receiving MAC Withdraw messages with the TLV type
                                   of 0x404.
                        –   Method 2
                            i.     Enter the system view.
                                   system-view

                            ii.    Enter the VSI view.
                                   vsi vsi-name

                            iii.   Configure a VSI to instruct all its peers to delete all the MAC
                                   addresses learned from this VSI when an AC or PW fault occurs but
                                   the VSI remains up.
                                   mac-withdraw enable

                                   By default, MAC Withdraw is disabled.
                                   After MAC Withdraw is enabled, the MAC addresses learned from the
                                   remote end may be deleted from the local end when the AC or PW
                                   goes up or down and packet broadcasting may occur.
                            iv.    (Optional) Enable the function of sending LDP MAC Withdraw
                                   messages to all peers when the status of the AC interface bound to
                                   the VSI changes.
                                   interface-status-change mac-withdraw enable


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                       766
VPN Configuration
VPN Configuration                                                                    6 VPLS Configuration


                                   By default, a VSI does not send LDP MAC Withdraw messages.
                            v.     (Optional) Enable the device in a VSI to remove MAC addresses
                                   except those associated with the PW through which the device
                                   receives MAC Withdraw messages with the TLV type 0x404.
                                   local-mac remove all-but-mine

                                   By default, the device removes the MAC addresses of all PWs and AC
                                   interfaces after receiving MAC Withdraw messages with the TLV type
                                   of 0x404.
                    ●   Configure MAC Withdraw relay. Two methods are available for configuring the
                        MAC Withdraw function: Method 1 applies to LDP VPLS, whereas method 2
                        applies to LDP VPLS and BGP AD VPLS.
                             NOTE

                            The mac-withdraw propagate enable command applies to MAC Withdraw messages
                            carrying either the FEC 128 TLV or the FEC 129 TLV. The upe-upe mac-withdraw
                            enable, upe-npe mac-withdraw enable, and npe-upe mac-withdraw enable
                            commands apply only to MAC Withdraw messages carrying the FEC 128 TLV. The mac-
                            withdraw propagate enable command cannot be used with the upe-upe mac-
                            withdraw enable, upe-npe mac-withdraw enable, or npe-upe mac-withdraw
                            enable command in the same VSI.
                        –   Method 1
                            i.     Enter the system view.
                                   system-view

                            ii.    Enter the VSI view.
                                   vsi vsi-name

                            iii.   Enter the VSI-LDP view.
                                   pwsignal ldp

                            iv.    Configure MAC Withdraw relay. Perform the following operations
                                   based on the source and destination of MAC Withdraw messages.
                                   ○    Enable an SPE to forward the LDP MAC Withdraw messages
                                        received from other NPEs to UPEs.
                                        npe-upe mac-withdraw enable

                                        By default, an SPE does not forward the LDP MAC Withdraw
                                        messages received from other NPEs to UPEs.
                                   ○    Enable an NPE to forward the LDP MAC Withdraw messages
                                        received from a UPE to other UPEs.
                                        upe-upe mac-withdraw enable

                                        By default, an NPE does not forward the LDP MAC Withdraw
                                        messages received from a UPE to other UPEs.
                                   ○    Enable an NPE to forward the LDP MAC Withdraw messages
                                        received from a UPE to other NPEs.
                                        upe-npe mac-withdraw enable

                                        By default, an NPE does not forward the LDP MAC Withdraw
                                        messages received from UPEs to other NPEs.
                        –   Method 2
                            i.     Enter the system view.
                                   system-view

                            ii.    Enter the VSI view.
                                   vsi vsi-name


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                           767
VPN Configuration
VPN Configuration                                                                     6 VPLS Configuration


                                 iii.   Enable a PE to forward received MAC Withdraw messages to its
                                        peers.
                                        mac-withdraw propagate enable

                                        By default, a PE does not forward received MAC Withdraw messages.


