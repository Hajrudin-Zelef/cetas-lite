---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-30
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [3715, 3841]
sha256: e7fb9c901699dae4ef1a02f67ff8c734d599e301ebe5dec0593f38f1e4d973d1
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

3.9 Configuring LDP LSPs
3.9.1 Understanding LDP LSPs
Common LDP LSP Establishment
                 The establishment of an LSP is actually a process of binding FECs and labels, and
                 then reporting the binding to neighboring LSRs on the LSP. Figure 3-9 shows the
                 process of establishing an LSP in DU and ordered modes.

                 Figure 3-9 LDP LSP establishment




                 1.     If an edge node on an MPLS network discovers a new direct route due to a
                        network route change and the address carried in the new route does not
                        belong to any existing FEC, the edge node creates a FEC for the address.
                 2.     If the egress has available labels, it distributes labels for FECs and proactively
                        sends a Label Mapping message to the upstream LSR. The Label Mapping
                        message contains distributed labels and bound FECs.
                 3.     The LSR adds the mapping in the Label Mapping message to the LFIB and
                        sends a Label Mapping message with a specified FEC to its upstream LSR.
                 4.     After receiving the Label Mapping message, the ingress also adds the
                        mapping to its LFIB. An LSP is then established, and the packets in this FEC
                        can be forwarded based on the label.

Proxy Egress LSP Establishment
                 A proxy egress is an egress that can trigger the establishment of LSPs based on
                 non-local routes. If penultimate hop popping is enabled on a router, the
                 penultimate hop acts as a special proxy egress. Generally, a proxy egress is
                 configured manually. A proxy egress can be used when a router that does not
                 support MPLS exists on a network. It can also be used to implement load
                 balancing among BGP routes.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                63
MPLS Configuration
MPLS Configuration                                                             3 MPLS LDP Configuration


                 As shown in Figure 3-10, LSR-A, LSR-B, and LSR-C are in the same MPLS domain.
                 MPLS LDP is not enabled on LSR-D. If all IGP routes are used to trigger LDP LSP
                 establishment, LSR-C functions as a proxy egress, which means LSR-A, LSR-B, and
                 LSR-C can still establish LDP LSPs to LSR-D.

                 Figure 3-10 LDP LSP establishment




3.9.2 Configuring Policies for Triggering LDP LSP
Establishment
Context
                 After MPLS LDP is enabled, LSPs are automatically established. If no policy is
                 configured, a large number of LSPs are established, wasting resources.
                 ●      A policy for triggering LDP LSP establishment can be configured on the
                        ingress and egress using the lsp-trigger command, allowing LDP to establish
                        LSPs only for eligible routes. This limits the number of LSPs to be established,
                        thereby reducing network resource consumption.
                        Such policies can be configured using the lsp-trigger command in both the
                        MPLS and MPLS-LDP-IPv4 views. If they are configured in both views, the
                        configuration in the MPLS-LDP-IPv4 view takes effect.
                 ●      A policy for triggering LDP LSP establishment can be configured on a transit
                        node using the propagate mapping command, allowing LDP to establish
                        LSPs only for eligible routes. The local node does not send Label Mapping
                        messages upstream for the routes that are filtered out. This limits the number
                        of LSPs to be established, thereby reducing network resource consumption.
                        Such policies can be configured using the propagate mapping command in
                        both the MPLS-LDP and MPLS-LDP-IPv4 views. If they are configured in both
                        views, only the latter configuration takes effect.
                 Generally, using the lsp-trigger command is recommended. If this command
                 cannot be run on the ingress or egress, use the propagate mapping command to
                 configure a policy.

Procedure
                 ●      Perform the following configuration on the ingress and egress.
                        –   Perform the configuration in the MPLS view.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             64
MPLS Configuration
MPLS Configuration                                                                       3 MPLS LDP Configuration


                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS view.
                             mpls

                        c.   Configure a policy for triggering LDP LSP establishment.
                             lsp-trigger { all | host | ip-prefix prefix-name | none }

                             ▪      By default, LDP uses IP routes with 32-bit addresses to trigger LDP
                                    LSP establishment.

                             ▪      The default configuration is recommended. Running the lsp-trigger
                                    all command is not recommended, as this command enables LDP
                                    LSPs to be established for all static routes and IGP routes. As a result,
                                    a large number of LSPs are established, consuming excessive label
                                    resources and slowing down LSP convergence on the entire network.
                                    If this command must be run, configure a policy for filtering out
                                    unnecessary routes first, thereby reducing the number of LDP LSPs to
                                    be established and system resource consumption.

                             ▪      If the lsp-trigger all command is run and an IGP route that does not
                                    overlap with that of an MPLS LDP session exists, the system
                                    establishes a proxy egress LSP for the IGP route.

                             ▪      The command allows LDP to establish ingress and egress LSPs for
                                    public network routes and for private network IGP routes. To
                                    configure a policy for triggering transit LSP establishment, run the
                                    propagate mapping command.

                             ▪      If the lsp-trigger command is run in both the MPLS and MPLS-LDP-
                                    IPv4 views, the configuration in the MPLS-LDP-IPv4 view takes effect.
                        d.   Disable the device from establishing proxy egress LSPs.
                             proxy-egress disable

                             If a policy allows LDP to establish LSPs for static and IGP routes or for
                             routes within a specified IP prefix list, the policy also allows LDP to
                             establish proxy egress LSPs. However, these proxy egress LSPs may be
                             useless and unnecessarily consume system resources. To prevent such an
                             issue, run this command to disable the device from establishing proxy
                             egress LSPs.
                        –    Perform the configuration in the MPLS-LDP-IPv4 view.
                        a.   Enter the system view.
                             system-view

