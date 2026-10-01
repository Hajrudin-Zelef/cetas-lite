---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-26
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [2753, 2902]
sha256: e266019632f3efe9fa78ca4b79c0d0db726bd9bd1cfd4117cc1b632588e554f1
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                     One-     All routes with the      Method 1:                   Method 1: ASBRs in
                     label-   same next hop and        The one-label-per-          inter-AS VPN Option
                     per-     same VPN label are       next-hop mode               B networking
                     next-    assigned the same        configured in the           Method 2: Devices on
                     hop      label.                   BGP-VPNv4/v6                which VPN instances
                                                       address family view is      are configured
                                                       mainly applicable to
                                                       inter-AS VPN Option
                                                       B networking.
                                                       Method 2:
                                                       The one-label-per-
                                                       next-hop mode
                                                       configured in the VPN
                                                       instance view is
                                                       applicable to all types
                                                       of IPv4 L3VPN
                                                       networking.




Implementation
                    One-label-per-instance
                    After one-label-per-instance label distribution is configured in a VPN instance IPv4
                    or IPv6 address family, all VPN routes from such an address family share the same
                    VPN label. As shown in Figure 3-6, PE1 is configured with two VPN instances. If

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                41
VPN Configuration
VPN Configuration                                                            3 IPv4 L3VPN Configuration


                    PE1 receives 10,000 routes from the sites of each VPN instance, only two labels on
                    the PE are used by default.

                    Figure 3-6 Networking for one-label-per-instance label distribution




                    One-label-per-route
                    After one-label-per-route label distribution is configured in the VPN instance IPv4
                    or IPv6 address family view, each VPN route is assigned a label. During the
                    forwarding of VPN packets, the packets are forwarded directly to the next hop
                    whose information is carried in a label, and the forwarding speed is fast.
                    One-label-per-next-hop
                    After one-label-per-next-hop label distribution is configured on an ASBR or PE, the
                    ASBR or PE re-advertises an MP-BGP Update message to its peers. The MP-BGP
                    Update message carries VPNv4 routes and their labels that are re-assigned based
                    on next hops. After a peer receives the MP-BGP Update message, the peer updates
                    its local label forwarding table and re-establishes LSPs. After the label forwarding
                    tables of the ASBR or PE and its peers are updated, service traffic is forwarded
                    according to the new label forwarding tables.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             42
VPN Configuration
VPN Configuration                                                                    3 IPv4 L3VPN Configuration


                    Figure 3-7 Networking for one-label-per-next-hop label distribution




                    One-label-per-next-hop label allocation is applicable to VPN routes learned by PE1
                    and peer routes learned by ASBRs:
                    ●   Method 1: On the network shown in Figure 3-7, two VPN instances named
                        VPN 1 and VPN 2 are configured on PE1 in the inter-AS VPN Option B
                        scenario, and the label distribution mode is one-label-per-route. If 10,000 VPN
                        routes are imported into CE1 and CE2 belonging to VPN 1 and VPN 2,
                        respectively, 20,000 labels are consumed when ASBR1 advertises 20,000
                        routes learned from PE1 to ASBR2. After one-label-per-next-hop label
                        allocation is enabled on ASBR1, ASBR1 allocates only one label to the VPN
                        routes with the same next hop and same outgoing label. In this case, ASBR1
                        only needs to allocate two labels to the 20,000 routes.
                    ●   Method 2: On the network shown in Figure 3-7, two VPN instances named
                        VPN 1 and VPN 2 are configured on PE1 in the inter-AS VPN Option B
                        scenario. CE1 and CE2 (belonging to VPN 1 and VPN 2, respectively) each
                        sends 10,000 VPN routes to PE1. If the label allocation mode is one-label-per-
                        route, 20,000 labels are consumed when PE1 advertises 20,000 routes to
                        ASBR1. After one-label-per-next-hop label allocation is enabled on PE1, PE1
                        assigns only one label to the VPN routes with the same next hop and same
                        outgoing label. In this case, PE1 only needs to assign two labels to the 20,000
                        routes.
                         NOTE

                        The one-label-per-route mode and one-label-per-next-hop mode can be flexibly switched to
                        each other. During label allocation mode switching, service packets are lost for a short
                        period due to the update of label forwarding tables on PEs and ASBRs.
                        In the inter-AS VPN Option B scenario, one-label-per-instance label distribution must be
                        configured on PEs if one-label-per-next-hop label distribution is configured on ASBRs.


Benefits
                    Using an appropriate label distribution mode significantly conserves label
                    resources.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                         43
VPN Configuration
VPN Configuration                                                                 3 IPv4 L3VPN Configuration




3.3 Configuration Precautions for IPv4 L3VPN

3.4 Default Settings for IPv4 L3VPN
                    Table 3-2 describes the default settings for IPv4 L3VPN.

                    Table 3-2 Default settings for IPv4 L3VPN
                     Parameter                          Default Setting

                     Maximum number of route            No default setting
                     prefixes supported by the
                     VPN instance IPv4 address
                     family




3.5 Configuring Mutual Access Between Local IPv4
L3VPNs

3.5.1 Configuring an IPv4 VPN Instance on a PE
Prerequisites
                    Before configuring an IPv4 VPN instance on a PE, you have completed the
                    following tasks:
                    ●    Configure link layer protocol parameters for interfaces to ensure that these
                         interfaces work properly.

Context
                    A VPN instance is also called a VRF table or a per-site forwarding table.
                    VPN instances are used to isolate VPN routes from public network routes. Routes
                    of different VPN instances are isolated from one another.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a VPN instance and enter the VPN instance view.
                    ip vpn-instance vpn-instance-name

                          NOTE

