---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-103
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [14516, 14651]
sha256: 777ae95731054305491906ba475b777b6e9660a063ccf8210b9139c574d0ddbd
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      IGP shortcut: The node does not advertise a TE tunnel to its neighbor nodes.
                        The TE tunnel can be involved only in local route calculation, but cannot be
                        used by the other nodes.
                 ●      Forwarding adjacency: The node advertises a TE tunnel to its neighbor nodes.
                        The TE tunnel is involved in global route calculation and can be used by the
                        other nodes. Forwarding adjacency advertises the tunnel through IGP
                        neighbor information.
                        If the forwarding adjacency mode is used, the two ends of a tunnel must be in
                        the same area.

                 The following example shows the differences between the two automatic routing
                 modes.



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        245
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


                 Figure 4-11 IGP shortcut and forwarding adjacency




                 On the network shown in Figure 4-11, a TE tunnel is established from LSR7 to
                 LSR2 over the path LSR7 -> LSR6 -> LSR2. The TE metric of the tunnel is shown in
                 the figure. Either of the following configurations can be used:
                 ●      Automatic routing is not configured: On LSR5, the next hop of the route to
                        LSR2 is LSR4. On LSR7, the next hop of the route to LSR1 is LSR6.
                 ●      Automatic routing is configured:
                        –   The TE tunnel Tunnel1 is advertised in shortcut mode. On LSR5, the next
                            hop of the route to LSR2 is still LSR4. On LSR7, the outbound interface of
                            the route to LSR1 is changed to Tunnel1. Only LSR7 has Tunnel1
                            participate in IGP route selection. However, LSR5 is unaware of Tunnel1
                            and does not have it participate in IGP route selection.
                        –   The TE tunnel Tunnel1 is advertised in forwarding adjacency mode. On
                            LSR5, the next hop of the route to LSR2 is changed to LSR7. On LSR7, the
                            outbound interface of the route to LSR1 is changed to Tunnel1. Both
                            LSR5 and LSR7 are aware of Tunnel1 and have it participate in IGP route
                            selection.


Tunnel Policy
                 By default, VPN traffic is forwarded through LDP LSPs. If the default LDP LSPs
                 cannot meet VPN traffic requirements, a tunnel policy can be used to steer VPN
                 traffic to a TE tunnel. Currently, there are two types of tunnel policies. You can
                 configure either of them as required.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                             246
MPLS Configuration
MPLS Configuration                                                               4 MPLS TE Configuration


                 ●      Select-seq: This policy can change the tunnel type selected by a VPN. A TE
                        tunnel can be selected as the public network tunnel of the VPN according to
                        the tunnel type priority configured in the tunnel policy.
                 ●      Tunnel binding: This policy binds a VPN to a specific destination address to a
                        TE tunnel for VPN traffic. This policy, however, is effective only for TE tunnels.

4.8.2 Configuring Static Routing
Prerequisites
                 Before steering traffic to an MPLS TE tunnel, complete one of the following tasks:
                 ●      Configure a static MPLS TE tunnel.
                 ●      Configure a dynamic MPLS TE tunnel.

Context
                 The simplest way to steer traffic to an MPLS TE tunnel is to configure a static
                 route.

Procedure
         Step 1 This static route operates just like a common static route. The only difference is
                that you need to configure a TE tunnel interface as the static route's outbound
                interface. For details, see Configuring an IPv4 Static Route in static route
                configuration.

                 ----End

Verifying the Configuration
                 ●      Run the display ip routing-table command to check whether an MPLS TE
                        tunnel interface is used as the outbound interface of a route.

4.8.3 Configuring Automatic Routing
Prerequisites
                 Before steering traffic to an MPLS TE tunnel, complete one of the following tasks:
                 ●      Configure a static MPLS TE tunnel.
                 ●      Configure a dynamic MPLS TE tunnel.

Context
                 The following automatic routing modes are supported:
                 ●      IGP shortcut: In this mode, a TE tunnel is not advertised to neighboring
                        nodes. Therefore, the TE tunnel can only participate in local route calculation,
                        and cannot be used by other nodes.
                 ●      Forwarding adjacency: In this mode, a TE tunnel is advertised to neighboring
                        nodes. Therefore, the TE tunnel participates in global route calculation, and
                        can be used by other nodes.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              247
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration


                 Configure either mode on the ingress of an MPLS TE tunnel as required.
                         NOTE

                        ● IGP shortcut and forwarding adjacency are mutually exclusive and cannot be both
                          configured.
                        ● In forwarding adjacency mode, a reverse tunnel needs to be configured for a routing
                          protocol to perform bidirectional check after a node advertises LSP links to the other
                          nodes. In addition, this mode needs to be enabled for both the forward and reverse
                          tunnels.


Procedure
                 ●      Configure IGP shortcut.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS TE tunnel interface view.
                             interface tunnel interface-number

                        c.   Use either of the following methods to configure IGP shortcut:

                             ▪    Configure IS-IS or OSPF to use TE tunnels for SPF calculation.
                                  mpls te igp shortcut [ isis | ospf ]

                                  If the IGP type is not specified when IGP shortcut is configured, both
                                  OSPF and IS-IS are supported by default.

                             ▪    Configure IS-IS to use TE tunnels for SPF calculation and delay the
                                  response to a TE tunnel state change.
                                  mpls te igp shortcut isis hold-time interval

