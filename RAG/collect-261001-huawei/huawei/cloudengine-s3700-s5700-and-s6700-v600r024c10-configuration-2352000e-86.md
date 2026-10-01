---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-86
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [11874, 12019]
sha256: 5be40d9e117feca0ba8c8b8150e2d68ce1168887c211389470191358441666d0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                            ii.    Create a default IPv4 static route for the VPN instance.
                                   ip route-static vpn-instance vpn-source-name 0.0.0.0 { 0.0.0.0 | 0 }{ interface-type
                                   interface-number } [ nexthop-address ] [ preference preference | tag tag ] * [ description
                                   text ]
                            iii.   Enter the BGP view.
                                   bgp as-number
                            iv.    Specify the UPE as a BGP peer.
                                   peer { ipv4-address | group-name } as-number as-number
                            v.     Enter the BGP-VPNv4 address family view.
                                   ipv4-family vpnv4
                            vi.    Specify the UPE as a BGP peer.
                                   peer { ipv4-address | group-name } upe

                                           NOTE

                                          This step can be performed only if a VPNv4 peer relationship has been
                                          established between the SPE and UPE.
                            vii. Return to the BGP view.
                                   quit
                            viii. Enter the BGP VPN instance IPv4 address family view.
                                   ipv4-family vpn-instance vpn-instance-name
                            ix.    Import the default route into the IPv4 VPN instance routing table.
                                   network 0.0.0.0 [ 0.0.0.0 | 0 ] [ route-policy route-policy-name ]
                        –   Configure the SPE to send a summary route to the UPE.
                            i.     Enter the system view.
                                   system-view
                            ii.    Enter the BGP view.
                                   bgp as-number
                            iii.   Enter the BGP VPN instance IPv4 address family view.
                                   ipv4-family vpn-instance vpn-instance-name
                            iv.    Create a summary route.
                                   aggregate ipv4-address { mask | mask-length } [ as-set | attribute-policy route-policy-
                                   name1 | detail-suppressed | origin-policy route-policy-name2 | suppress-policy route-
                                   policy-name3 ]
                            v.     Return to the BGP view.
                                   quit
                            vi.    Return to the system view.
                                   quit
                            vii. Configure an IPv4 prefix list.
                                   ip ip-prefix ip-prefix-name [ index index-number ] { permit | deny } ip-address mask-
                                   length [ greater-equal greater-equal-value ] [ less-equal less-equal-value ]
                            viii. Enter the BGP view.
                                   bgp as-number
                            ix.    Specify the UPE as a BGP peer.
                                   peer { ipv4-address | group-name } as-number as-number
                            x.     Enter the BGP-VPNv4 address family view.
                                   ipv4-family vpnv4
                            xi.    Configure the device to advertise filtered routes to the UPE.
                                   peer { ipv4-address | group-name } ip-prefix ip-prefix-name export
                    ●   (Optional) Configure one-label-per-next-hop label distribution on the SPE.
                        In an HoVPN scenario, if an SPE needs to send large numbers of VPNv4 routes
                        but the MPLS labels are inadequate, configure one-label-per-next-hop label
                        distribution on the SPE.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                               191
VPN Configuration
VPN Configuration                                                               3 IPv4 L3VPN Configuration


                        a.    Enter the system view.
                              system-view

                        b.    Enter the BGP view.
                              bgp as-number

                        c.    Enter the BGP-VPNv4 address family view.
                              ipv4-family vpnv4

                        d.    Enable one-label-per-next-hop label distribution for VPNv4 routes.
                              apply-label per-nexthop




                                    NOTICE

                              After one-label-per-next-hop label distribution is enabled or disabled on
                              an SPE, the labels assigned by the SPE to routes change. As a result,
                              temporary packet loss may occur.

                    ----End

Verifying the Configuration
                    Run the display ip routing-table command to check the IP routing table.

3.14.3 Configuring H-VPN

Context
                    On an H-VPN, SPEs function as RRs and UPEs function as RR clients to receive
                    specific routes from SPEs.
                    This scenario requires the following configurations:
                    ●   Create VPN instances on UPEs and NPEs. For configuration details, see
                        Configuring an IPv4 VPN Instance.
                    ●   Configure an MP-BGP peer relationship between each NPE and SPE and
                        between each SPE and UPE. The configuration is similar to that in an IPv4
                        VPN instance. For details, see 3.6.4 Establishing MP-IBGP Peer Relationships
                        Between PEs.
                    ●   Configure routing protocols for NPEs/UPEs to exchange routes with CEs. This
                        configuration is similar to configuring PEs and CEs to exchange routes in an
                        IPv4 VPN instance. For details, see 3.6 Configuring Basic IPv4 L3VPN over
                        MPLS.
                    ●   Configure SPEs as RRs and UPEs as RR clients.
                    Perform the following steps on each SPE.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                         192
VPN Configuration
VPN Configuration                                                                     3 IPv4 L3VPN Configuration


         Step 3 Enter the BGP-VPNv4 address family view.
                    ipv4-family vpnv4

         Step 4 Specify an RR and its client.
                    peer { ipv4-address | group-name } reflect-client

         Step 5 The device is configured to use its own address as the next-hop address of routes
                when advertising these routes.
                    peer { ipv4-address | group-name } next-hop-local

                    To allow an SPE to use its own address as the next-hop address when advertising
                    routes to UPEs and NPEs, run the peer next-hop-local command on the SPE twice
                    with different parameters specified for the UPE and NPE.
         Step 6 Enable one-label-per-next-hop label distribution for VPNv4 routes.
                    apply-label per-nexthop

                    In an H-VPN scenario, if an SPE needs to send large numbers of VPNv4 routes but
                    the MPLS labels are inadequate, configure one-label-per-next-hop label
                    distribution on the SPE.


                        NOTICE

                    After one-label-per-next-hop label distribution is enabled or disabled on an SPE,
                    the labels assigned by the SPE to routes change. As a result, temporary packet loss
                    may occur.

