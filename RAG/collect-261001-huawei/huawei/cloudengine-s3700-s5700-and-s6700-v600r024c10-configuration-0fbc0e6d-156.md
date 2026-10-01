---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-156
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [22576, 22702]
sha256: 34e01d3411ce8f12325bae818a4d9a9d7c50f87878e2a9d95ba1890812d651a9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        ● Tunnel re-optimization is performed based on tunnel path constraints. During path
                          calculation for re-optimization, path constraints, such as explicit path constraints and
                          bandwidth constraints, are also considered.
                        ● Tunnel re-optimization cannot be used on tunnels for which CSPF selects paths in most-
                          fill tie-breaking mode.
                        ● Tunnel re-optimization takes effect on the primary and hot-standby paths, but does not
                          take effect on ordinary backup and best-effort paths.
                        ● If a bypass tunnel is configured or the resource reservation style of the tunnel is FF,
                          tunnel re-optimization cannot be configured.

                 Re-optimization is classified into the following types based on the triggering
                 mode:
                 ●      Automatic re-optimization

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         375
MPLS Configuration
MPLS Configuration                                                                        4 MPLS TE Configuration


                        The ingress automatically triggers CSPF to calculate a path for a tunnel when
                        the re-optimization interval configured by a network administrator is reached
                        or when changes occur in the network topology (such as link addition, link
                        removal, or IGP link cost adjustment). If the calculated path has a smaller
                        metric value than the existing path, a CR-LSP is set up over the new path. If
                        the CR-LSP is successfully set up, the forwarding plane is instructed to switch
                        traffic to the new CR-LSP and delete the original CR-LSP. The re-optimization
                        is then complete. If the CR-LSP is not set up, the traffic is still forwarded along
                        the original CR-LSP.
                 ●      Manual re-optimization
                        The re-optimization command is run in the user view to trigger re-
                        optimization on the tunnel ingress.

                 The make-before-break mechanism is used to ensure nonstop service
                 transmission during re-optimization. This means that a new CR-LSP must be
                 established first. Traffic is switched to the new CR-LSP before the original CR-LSP is
                 torn down.

4.20.2 Configuring MPLS TE Tunnel Re-optimization

Prerequisites
                 Before configuring MPLS TE tunnel re-optimization, complete the following task:

                 ●      Configure a dynamic MPLS TE tunnel.


Context
                 An MPLS TE tunnel can be re-optimized using either of the following methods:
                 ●      Automatic re-optimization: The system automatically re-optimizes TE tunnels
                        when the re-optimization interval is reached or when changes occur in the
                        network topology (such as link addition, link removal, or IGP link cost
                        adjustment). This implementation does not involve manual intervention and
                        reduces labor cost.
                 ●      Manual re-optimization: A user configures the system to attempt to
                        reestablish TE tunnels over better paths if there are.

                         NOTE

                        ● Tunnel re-optimization is performed based on tunnel path constraints. During path
                          calculation for re-optimization, path constraints, such as explicit path constraints and
                          bandwidth constraints, are also considered.
                        ● Tunnel re-optimization cannot be used on tunnels for which CSPF selects paths in most-
                          fill tie-breaking mode.
                        ● Tunnel re-optimization takes effect on the primary and hot-standby paths, but does not
                          take effect on ordinary backup and best-effort paths.
                        ● If a bypass tunnel is configured or the resource reservation style of the tunnel is FF,
                          tunnel re-optimization cannot be configured.

                 Perform the following configuration on the ingress of an MPLS TE tunnel.



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         376
MPLS Configuration
MPLS Configuration                                                                       4 MPLS TE Configuration


Procedure
                 ●      Configure automatic tunnel re-optimization.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS TE tunnel interface view.
                             interface tunnel interface-number

                        c.   Enable automatic tunnel re-optimization.
                             mpls te reoptimization [ frequency interval ]

                        d.   (Optional) Disable link change-triggered RSVP-TE tunnel re-optimization.
                             quit
                             mpls
                             mpls te reoptimization link-up disable

                             If route flapping occurs on a network, tunnel re-optimization is
                             repeatedly triggered, wasting network resources. To prevent the problem,
                             run this command to disable link change-triggered RSVP-TE tunnel re-
                             optimization so that tunnel re-optimization will not be performed.

                                   NOTE

                                  This command takes effect only for link change-triggered tunnel re-optimization.
                                  It does not take effect if the mpls te reoptimization (tunnel interface view)
                                  command is run with the frequency interval parameter specified.
                                  This command disables the re-optimization function configured for all tunnels.
                 ●      Configure manual tunnel re-optimization.
                        mpls te reoptimization [ auto-tunnel name auto-tunnel-name | tunnel tunnel-number ]

                        After this command is run in the user view, re-optimization can be triggered
                        for all tunnels or a specified tunnel on the local node.
                 ●      (Optional) Enable IGP metric-based re-optimization for an MPLS TE tunnel.

                        Perform this step if you want to re-optimize an MPLS TE tunnel based only on
                        the IGP metric. The following constraints are ignored during re-optimization:

                        –    Bandwidth usage: A link is selected based on the percentage of the used
                             reservable bandwidth to the maximum reservable bandwidth.
                        –    Hop-counts: A link is selected based on the number of hops on the path.
                        a.   Enter the system view.
                             system-view

                        b.   Configure tunnel re-optimization based only on the IGP metric.

                             The global MPLS configuration takes effect on all MPLS TE tunnels and is
                             used for batch configuration. A single tunnel configuration takes
                             precedence over the global MPLS configuration.

