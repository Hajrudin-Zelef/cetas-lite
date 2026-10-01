---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-167
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [24210, 24353]
sha256: 9545a563fdd5f3ded15e6b1d0106cd56e4d29565e5b556b5c270066b41a0ec4b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Bind a BFD session to the primary or backup CR-LSP of a specified tunnel.
                 bfd sessname-value bind mpls-te interface tunnel tunnel-number te-lsp [ backup ]

                 If the backup parameter is specified, the BFD session is bound to a backup CR-LSP.
         Step 3 Set a local discriminator for the BFD session.
                 discriminator local discr-value

         Step 4 Set a remote discriminator for the BFD session.
                 discriminator remote discr-value

                         NOTE

                        The local discriminator of the local device and the remote discriminator of the remote
                        device are the same. The remote discriminator of the local device and the local
                        discriminator of the remote device are the same. A discriminator inconsistency causes the
                        BFD session to fail to be established.

         Step 5 Enable the system to modify the port state table (PST) when the BFD session state
                changes.
                 process-pst


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     402
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 When the BFD session state changes, the BFD module notifies the application
                 protocol of the change. Fast switching between the primary and backup CR-LSPs
                 can be then performed.
         Step 6 (Optional) Set local BFD parameters as required. For details, see Table 4-20.
                 You can perform one or more steps in Table 4-20 as required.

                 Table 4-20 Configuring local BFD parameters

                  Operation                     Command                     Purpose

                  Set the minimum               min-tx-interval interval    The default interval is
                  interval at which the                                     1000 milliseconds.
                  local device sends BFD
                  packets.

                  Set the minimum               min-rx-interval interval    The default interval is
                  interval at which the                                     1000 milliseconds.
                  local device receives BFD
                  packets

                  Set the local BFD             detect-multiplier           The default local BFD
                  detection multiplier.         multiplier                  detection multiplier is 3.




                 Actual local interval for sending BFD packets = max {Locally configured interval
                 for sending BFD packets, Remotely configured interval for receiving BFD packets}
                 Actual local interval for receiving BFD packets = max {Remotely configured
                 interval for sending BFD packets, Locally configured interval for receiving BFD
                 packets}
                 Actual local detection interval = Actual local interval for receiving BFD packets x
                 Remotely configured BFD detection multiplier
                 For example:
                 ●      On the local device, the interval for sending BFD packets is set to 200 ms, the
                        interval for receiving BFD packets is set to 300 ms, and the detection
                        multiplier is set to 4.
                 ●      On the remote device, the interval for sending BFD packets is set to 100 ms,
                        the interval for receiving BFD packets is set to 600 ms, and the detection
                        multiplier is set to 5.
                 Then:
                 ●      On the local device, the actual interval for sending BFD packets is 600 ms, as
                        calculated using the formula max {200 ms, 600 ms}; the actual interval for
                        receiving BFD packets is 300 ms, as calculated using the formula max {100
                        ms, 300 ms}; the detection multiplier is 1500 ms, as calculated using the
                        formula 300 ms x 5.
                 ●      On the remote device, the actual interval for sending BFD packets is 300 ms,
                        as calculated using the formula max {100 ms, 300 ms}; the actual interval for
                        receiving BFD packets is 600 ms, as calculated using the formula max {200

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             403
MPLS Configuration
MPLS Configuration                                                                           4 MPLS TE Configuration


                        ms, 600 ms}; the detection multiplier is 2400 ms, as calculated using the
                        formula 600 ms x 4.

                 ----End

4.24.4 Setting BFD Parameters on an Egress
Context
                 BFD parameters that can be configured on a tunnel egress include the local
                 discriminator, remote discriminator, local interval for sending BFD packets, local
                 interval for receiving BFD packets, and local BFD detection multiplier. These
                 parameters affect the establishment of a BFD session.
                 Perform the following configuration on the egress of an MPLS TE tunnel.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Configure a reverse path to notify the ingress of faults. A reverse path can be an
                IP link, an LSP, or a TE tunnel. If there is a reverse LSP or TE tunnel, use the
                reverse LSP or TE tunnel. If no LSP or TE tunnel is established, use an IP link as a
                reverse path instead. If the configured reverse path requires BFD detection, you
                can configure a pair of BFD sessions for it. Perform one of the following
                configurations as required:
                 ●      Configure an IP link as a reverse path.
                        bfd session-name bind peer-ip peer-ip [ vpn-instance vpn-name ] [ interface interface-type
                        interface-number] [ source-ip source-ip ]
                 ●      Configure an LDP LSP as a reverse path.
                        bfd session-name bind ldp-lsp peer-ip ip-address nexthop ip-address [ interface interface-type
                        interface-number ]
                 ●      Configure a CR-LSP as a reverse path.
                        bfd sessname-value bind mpls-te interface tunnel tunnel-num te-lsp [ backup ]
                 ●      Configure a TE tunnel as a reverse path.
                        bfd sessname-value bind mpls-te interface tunnel tunnel-num

         Step 3 Set a local discriminator for the BFD session.
                 discriminator local discr-value

         Step 4 Set a remote discriminator for the BFD session.
                 discriminator remote discr-value

                         NOTE

                        The local discriminator of the local device and the remote discriminator of the remote
                        device are the same. The remote discriminator of the local device and the local
                        discriminator of the remote device are the same. A discriminator inconsistency causes the
                        BFD session to fail to be established.

         Step 5 (Optional) Enable the system to modify the PST when the BFD session state
                changes.
                 process-pst

                 When the BFD session state changes, the BFD module notifies the application
                 protocol of the change. Fast switching between the primary and backup CR-LSPs

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                              404

