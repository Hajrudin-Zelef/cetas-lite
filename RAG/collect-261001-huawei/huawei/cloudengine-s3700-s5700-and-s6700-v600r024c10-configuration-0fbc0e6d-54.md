---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-54
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [7430, 7571]
sha256: 8b689493c54cbcb538582b1ed17ffa04c39647598d1f54ba3a024502efd0188b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 Effective local interval at which BFD packets are sent = max{Locally configured
                 interval at which BFD packets are sent, Remotely configured interval at which BFD
                 packets are received}

                 Effective local interval at which BFD packets are received = max{Remotely
                 configured interval at which BFD packets are sent, Locally configured interval at
                 which BFD packets are received}

                 Local detection period = Local interval at which BFD packets are received x
                 Remote BFD detection multiplier

                 For example, if: On the local device, the intervals at which BFD packets are sent
                 and received are 200 ms and 300 ms, respectively, and the detection multiplier is
                 4; on the remote device, the intervals at which BFD packets are sent and received
                 are 100 ms and 600 ms, respectively, and the detection multiplier is 5. Then:

                 ●      On the local device, the actual interval for sending BFD packets is 600 ms,
                        which is calculated using the formula max{200 ms, 600 ms}; the interval for
                        receiving BFD packets is 300 ms, which is calculated using the formula max
                        {100 ms, 300 ms}; the detection period is 1500 ms (300 ms × 5).
                 ●      On the remote device, the actual interval for sending BFD packets is 300 ms,
                        which calculated using the formula max{100 ms, 300 ms}; the interval for
                        receiving BFD packets is 600 ms, which is calculated using the formula
                        max{200 ms, 600 ms}; the detection period is 2400 ms (600 ms × 4).

         Step 7 (Optional) Change the minimum interval at which the local device receives BFD
                packets.
                 min-rx-interval rx-interval

                 By default, the minimum interval is 10 milliseconds.

                 If the reverse path is an IP link, this command cannot be run.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    125
MPLS Configuration
MPLS Configuration                                                                          3 MPLS LDP Configuration


         Step 8 (Optional) Change the local BFD detection multiplier.
                 detect-multiplier multiplier

                 By default, the local BFD detection multiplier is 3.

                 ----End

3.15.4 Setting BFD Parameters on the Egress
                 BFD parameters must be configured on the egress before a BFD session is
                 established to monitor an LDP LSP.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Configure a reverse path for notifying the ingress of faults.

                 The reverse path can be an IP link or an LSP. If an LSP exists, use the LSP as the
                 reverse path preferentially. Otherwise, use an IP link. If BFD also needs to be
                 configured for the reverse path, configure another pair of BFD sessions for it as
                 follows:

                 ●      If the reverse path is an IP link:
                        bfd session-name bind peer-ip peer-ip [ vpn-instance vpn-name ] [ source-ip source-ip ]

                        The peer-ip ip-address value should be the LSR ID of the remote device.
                 ●      If the reverse path is an LDP LSP:
                        bfd session-name bind ldp-lsp peer-ip ip-address nexthop ip-address [ interface interface-type
                        interface-number ]
                        The peer-ip ip-address value should be the LSR ID of the remote device.

         Step 3 Set a local discriminator for the BFD session.
                 discriminator local discr-value

         Step 4 Set a remote discriminator for the BFD session.
                 discriminator remote discr-value

                         NOTE

                        The local discriminator of the local device and the remote discriminator of the remote
                        device must be the same. The remote discriminator of the local device and the local
                        discriminator of the remote device must be the same. A discriminator inconsistency causes
                        the BFD session to fail to be established.

         Step 5 (Optional) Change the minimum interval at which the local device sends BFD
                packets.
                 min-tx-interval tx-interval

                 By default, the minimum interval is 10 milliseconds.

                 If the reverse path is an IP link, this command cannot be run.

                 Effective local interval at which BFD packets are sent = max{Locally configured
                 interval at which BFD packets are sent, Remotely configured interval at which BFD
                 packets are received}

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                              126
MPLS Configuration
MPLS Configuration                                                             3 MPLS LDP Configuration


                 Effective local interval at which BFD packets are received = max{Remotely
                 configured interval at which BFD packets are sent, Locally configured interval at
                 which BFD packets are received}

                 Local detection period = Local interval at which BFD packets are received x
                 Remote BFD detection multiplier

                 For example, if: On the local device, the intervals at which BFD packets are sent
                 and received are 200 ms and 300 ms, respectively, and the detection multiplier is
                 4; on the remote device, the intervals at which BFD packets are sent and received
                 are 100 ms and 600 ms, respectively, and the detection multiplier is 5. Then:

                 ●      On the local device, the actual interval for sending BFD packets is 600 ms,
                        which is calculated using the formula max{200 ms, 600 ms}; the interval for
                        receiving BFD packets is 300 ms, which is calculated using the formula
                        max{100 ms, 300 ms}; the detection period is 1500 ms (300 ms × 5).
                 ●      On the remote device, the actual interval for sending BFD packets is 300 ms,
                        which calculated using the formula max{100 ms, 300 ms}; the interval for
                        receiving BFD packets is 600 ms, which is calculated using the formula
                        max{200 ms, 600 ms }; the detection period is 2400 ms (600 ms × 4).

         Step 6 (Optional) Change the minimum interval at which the local device receives BFD
                packets.
                 min-rx-interval rx-interval

                 By default, the minimum interval is 10 milliseconds.

                 If the reverse path is an IP link, this command cannot be run.

         Step 7 (Optional) Change the local BFD detection multiplier.
                 detect-multiplier multiplier

                 By default, the local BFD detection multiplier is 3.

         Step 8 (Optional) Allow the BFD session to modify the port or link state table.
                 process-pst

                 By default, a BFD session does not modify the port or link state table upon
                 detection of a fault.

