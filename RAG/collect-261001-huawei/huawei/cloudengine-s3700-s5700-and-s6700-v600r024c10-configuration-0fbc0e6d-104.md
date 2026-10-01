---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-104
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [14652, 14793]
sha256: 37803b50169634197f8ac37aa87e23bcb3ece961eaa1c5ce089f7763512e232e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                                  The hold-time interval parameter indicates the delay after which IS-
                                  IS responds to a TE tunnel down event. If this parameter is not
                                  specified when a TE tunnel goes down, routes are recalculated. If this
                                  parameter is specified, IS-IS delays responding to the TE tunnel down
                                  event. If no other condition triggers route recalculation during this
                                  period, that is, the TE tunnel goes up after the delay, routes are not
                                  recalculated. If the TE tunnel remains down after the delay, routes
                                  are recalculated.
                        d.   Configure an IGP metric value for the TE tunnel.
                             mpls te igp metric { absolute | relative } value

                             When specifying the metric value for a TE tunnel in IGP shortcut mode,
                             note the following points:

                             ▪    If the absolute parameter is configured, the metric value of the TE
                                  tunnel is equal to the configured metric value.

                             ▪    If the relative parameter is configured, the metric value of the TE
                                  tunnel is the sum of the metric value of the corresponding IGP path
                                  and the relative metric value.
                        e.   Use either of the following methods to configure an IS-IS or OSPF process
                             for the tunnel interface.

                             ▪    For IS-IS, enable the IS-IS process on the tunnel interface.
                                  isis enable [ process-id ]


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     248
MPLS Configuration
MPLS Configuration                                                                                4 MPLS TE Configuration


                             ▪    For OSPF, enable the OSPF process on the tunnel interface and
                                  enable IGP shortcut.
                                  ospf enable process-id area { area-id | areaidipv4 }
                                  quit
                                  ospf [ process-id ]
                                  enable traffic-adjustment

                                         NOTE

                                        You can also run the network address wildcard-mask command in the OSPF
                                        view to enable the network segment routes of the tunnel interface.
                 ●      Configure forwarding adjacency.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS TE tunnel interface view.
                             interface tunnel interface-number

                        c.   Configure forwarding adjacency.
                             mpls te igp advertise [ hold-time interval | include-ipv6-isis ] *

                             If IPv6 IS-IS is used, configure the include-ipv6-isis parameter.
                        d.   Configure an IGP metric value for the TE tunnel.
                             mpls te igp metric { absolute | relative } value

                             By default, the metric value of a TE tunnel is the same as that of an IGP
                             path.

                             When specifying the metric value for a TE tunnel in IGP shortcut mode,
                             note the following points:

                             ▪    If the absolute parameter is configured, the metric value of the TE
                                  tunnel is equal to the configured metric value.

                             ▪    If the relative parameter is configured, the metric value of the TE
                                  tunnel is the sum of the metric value of the corresponding IGP path
                                  and the relative metric value.
                                   NOTE

                                  Set a proper IGP metric value to ensure that the LSP link is advertised and used
                                  correctly. The metric value of a TE tunnel should be smaller than that of an IGP
                                  route that is not expected for use.
                                  If relative is configured and IS-IS is used, this step cannot modify the IS-IS metric
                                  value of the TE tunnel. To change the IS-IS metric value, configure absolute in
                                  this step.
                        e.   Use either of the following methods to configure an IS-IS or OSPF process
                             for the tunnel interface.

                             ▪    For IS-IS, enable the IS-IS process on the tunnel interface.
                                  isis enable [ process-id ]

                             ▪    For OSPF, enable the OSPF process on the tunnel interface and
                                  enable forwarding adjacency.
                                  ospf enable process-id area { area-id | areaidipv4 }
                                  quit
                                  ospf [ process-id ]
                                  enable traffic-adjustment advertise


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                           249
MPLS Configuration
MPLS Configuration                                                               4 MPLS TE Configuration


                                      NOTE

                                     You can also run the network address wildcard-mask command in the OSPF
                                     view to enable the network segment routes of the tunnel interface.

                 ----End


Verifying the Configuration
                 ●      Run the display ip routing-table command to check whether an MPLS TE
                        tunnel interface is used as the outbound interface of a route.
                 ●      Run the display ospf [ process-id ] traffic-adjustment command to check
                        tunnel information of the OSPF process related to traffic forwarding (IGP
                        shortcut and forwarding adjacency).

4.8.4 Configuring a Tunnel Policy

Prerequisites
                 Before steering traffic to an MPLS TE tunnel, complete one of the following tasks:

                 ●      Configure a static MPLS TE tunnel.
                 ●      Configure a dynamic MPLS TE tunnel.


Context
                 Generally, when VPN traffic is forwarded through tunnels, LSPs instead of MPLS TE
                 tunnels are used by default. If the default configuration does not meet VPN traffic
                 requirements, a tunnel policy can be configured to steer VPN traffic to an MPLS TE
                 tunnel. Currently, there are two types of tunnel policies. You can configure either
                 of them as required.


Procedure
                 ●      Select-seq: This policy can change the tunnel type selected by a VPN. A TE
                        tunnel can be selected as the public network tunnel of the VPN according to
                        the configured tunnel type priority. For details, see "Configure a tunnel type
                        prioritizing policy" in "VPN Configuration."
                 ●      Tunnel binding: This policy binds a specific destination address to a TE tunnel
                        for VPN traffic to guarantee QoS. For details, see "Configure a tunnel binding
                        policy" in "VPN Configuration."

                 ----End


