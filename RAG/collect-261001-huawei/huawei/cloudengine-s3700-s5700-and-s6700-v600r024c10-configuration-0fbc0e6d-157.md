---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-157
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [22703, 22873]
sha256: f38cd64e928ce8712e92cc506d38f3901e6cbeecf98cdc9c99fe584d0962c439
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                             ▪    In the global MPLS view, configure tunnel re-optimization based only
                                  on the IGP metric.
                                  mpls
                                  mpls te reoptimization-aggressive enable
                                  quit

                             ▪    In the tunnel interface view, configure tunnel re-optimization based
                                  only on the IGP metric.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     377
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration

                                  interface tunnel interface-number
                                  mpls te reoptimization-aggressive enable
                                  quit

                        c.   (Optional) Disable tunnel re-optimization that is based only on the IGP
                             metric.
                             interface tunnel interface-number
                             mpls te reoptimization-aggressive block

                             To cancel the configuration for some tunnels, run this command in the
                             tunnel interface view of these MPLS TE tunnels.

                 ----End


Verifying the Configuration
                 ●      Run the display mpls te tunnel-interface command to check tunnel
                        interface information on the local node, including the interval and enabling
                        status of automatic re-optimization.

4.20.3 (Optional) Suppressing Immediate Reestablishment for
Down LSPs

Context
                 By default, a device starts the reestablishment process for all down LSPs
                 immediately when links change. If route flapping occurs on the network, path
                 computation is repeatedly performed for the down LSPs, wasting network
                 resources. To prevent such an issue, suppress immediate reestablishment for down
                 LSPs. Then, the device will not start the reestablishment process for down LSPs
                 immediately when links change. Instead, it waits for the reestablishment timer to
                 expire before starting reestablishment.

                 Perform the following configuration on the ingress of an MPLS TE tunnel.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enable MPLS on the local node and enter the MPLS view.
                 mpls

         Step 3 Enable MPLS TE globally.
                 mpls te

         Step 4 Disable the device from performing immediate reestablishment for down LSPs
                when links change.
                 mpls te rebuild-lsp link-up disable

                 ----End




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                         378
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration




4.21 Configuring Delayed Switching and Deletion in
MPLS TE
Prerequisites
                 Before configuring delayed switching and deletion in MPLS TE, complete the
                 following task:

                 ●      Configure a dynamic MPLS TE tunnel.


Context
                 MPLS TE uses the make-before-break mechanism. When MPLS TE link attributes
                 and tunnel attributes change, a CR-LSP that satisfy new attributes needs to be
                 established. To minimize data loss and bandwidth consumption during traffic
                 switching, a new CR-LSP must be established before the original CR-LSP is torn
                 down.

                 In real-world applications, the service status of each node on an MPLS network
                 varies. If an upstream node on an MPLS network is busy but its downstream node
                 is idle or an upstream node is idle but its downstream node is busy, a CR-LSP may
                 be torn down before the new CR-LSP is established, causing a temporary traffic
                 interruption.

                 The make-before-break mechanism uses switching and deletion delay timers to
                 prevent temporary traffic interruptions. When the two timers are configured, the
                 system switches traffic to a new CR-LSP after the switching delay time, and then
                 deletes the original CR-LSP after the deletion delay time.

                 Perform the following configuration on the ingress of an MPLS TE tunnel.


Procedure
         Step 1 Enter the system view of the tunnel ingress.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Configure switching and deletion delays.
                 mpls te switch-delay switch-time delete-delay delete-time

                 ----End


4.22 Configuring Synchronization Between CR-LSP
Establishment and Overload Status


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                       379
MPLS Configuration
MPLS Configuration                                                          4 MPLS TE Configuration


4.22.1 Understanding Synchronization Between CR-LSP
Establishment and Overload Status
                 Association between CR-LSP establishment and IS-IS overload allows MPLS TE
                 traffic to bypass overloaded nodes. This function allows a tunnel ingress to bypass
                 overloaded nodes when establishing CR-LSPs, improving CR-LSP reliability and
                 service quality.

Context
                 When a device cannot store new LSPs or synchronize the LSDB, the routing
                 information calculated by the device becomes incorrect. In that case, the device
                 enters the overload state. When deploying MPLS TE services, you can configure
                 synchronization between CR-LSP establishment and overload, allowing all CR-LSPs
                 to bypass overloaded nodes. This improves CR-LSP reliability and service quality.

Related Concepts
                 IS-IS overload state
                 When a device cannot store new LSPs or synchronize the LSDB, the routing
                 information calculated by the device becomes incorrect. In that case, the device
                 enters the overload state. For example, if the system memory is insufficient, you
                 can configure whether the device enters the overload state. An IS-IS device may
                 enter the overload state either due to an unexpected condition or it can be
                 manually set to enter the overload state.

Implementation
                 In the topology shown in Figure 4-33, LSR1 supports association between CR-LSP
                 establishment and overload, and LSR3 and LSR4 support IS-IS overload.

                 Figure 4-33 Sychronization between CR-LSP establishment and IS-IS overload
                 status




                 LSR1 establishes a tunnel (such as Tunnel1). The tunnel destination is LSR2, and
                 the LSP path is LSR1 -> LSR3 -> LSR2.

Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                           380
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


