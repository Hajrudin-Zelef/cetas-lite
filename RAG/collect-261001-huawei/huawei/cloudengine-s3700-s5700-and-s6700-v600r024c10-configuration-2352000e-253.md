---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-253
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [37165, 37288]
sha256: 965ab8419383ed8ac7bf0d7da7e2914f5fa1fef19c9b049299741bf766c85fc3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         Ethernet        Ethernet frames transmitted between CEs and PEs do not
                         access          carry the P-tag. If an Ethernet frame carries a VLAN tag, the
                                         tag is an internal VLAN tag called a user-tag (U-tag) in user
                                         packets. The U-tag is carried in a packet before the packet is
                                         sent to a CE and is not added by the CE. The U-tag is used by
                                         the CE to identify to which VLAN the packet belongs and is
                                         meaningless to PEs.


                    ●   Packet encapsulation on PWs
                        The PW ID and PW encapsulation type combine to uniquely identify a PW.
                        The PW IDs and PW encapsulation types configured on PEs at both end of a
                        PW must be the same. The packet encapsulation types of packets on PWs can
                        be raw or tagged. By default, packets on PWs are encapsulated in tagged
                        mode.

                        Table 6-3 Packet encapsulation types
                         Packet           Description
                         Encapsulati
                         on on PWs

                         Raw              Packets transmitted over a PW cannot carry P-tags. If a PE
                                          receives a packet with the P-tag from a CE, the PE strips the
                                          P-tag and adds two labels (outer tunnel label and inner VC
                                          label) to the packet before forwarding it. If a PE receives a
                                          packet with no P-tag from a CE, the PE directly adds two
                                          labels (outer tunnel label and inner VC label) to the packet
                                          before forwarding it. The PE determines whether to add the
                                          P-tag to a packet, depending on the configuration, before
                                          sending it to a CE. The PE is not allowed to rewrite or
                                          remove an existing U-tag.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          592
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration


                         Packet         Description
                         Encapsulati
                         on on PWs

                         Tagged         Packets transmitted over a PW must carry P-tags. If a PE
                                        receives a packet with the P-tag from a CE, the PE directly
                                        adds two labels (outer tunnel label and inner VC label) to
                                        the packet before forwarding it. If a PE receives a packet
                                        with no P-tag from a CE, the PE adds a null P-tag and two
                                        labels (outer tunnel label and inner VC label) to the packet
                                        before forwarding it. The PE determines whether to rewrite,
                                        remove, or preserve the P-tag of a packet, depending on the
                                        configuration, before forwarding it to a CE.


                    Encapsulation modes of packets transmitted over ACs and PWs can be used
                    together. The following uses Ethernet access in raw mode (without the U-tag) and
                    VLAN access in tagged mode (with the U-tag) as examples to describe the packet
                    exchange process.
                    ●   Ethernet access in raw mode (without the U-tag)

                        Figure 6-4 Ethernet access in raw mode (without the U-tag)




                        As shown in Figure 6-4, ACs use Ethernet encapsulation and PWs use raw
                        encapsulation; packets transmitted from CEs to PEs do not carry U-tags.
                        The packet exchange process in Ethernet access in raw mode (without the U-
                        Tag) is as follows:
                        a.   CE1 sends to PE1 a packet that is encapsulated at Layer 2 and does not
                             carry any U-tag or P-tag.
                        b.   After receiving the packet that does not carry any U-tag or P-tag, PE1
                             searches for a forwarding entry in the corresponding VSI, and selects a
                             tunnel and a PW for the packet. PE1 adds two MPLS labels (outer tunnel

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         593
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


                             label and inner VC label) to the packet based on the selected tunnel and
                             PW, performs Layer 2 encapsulation, and forwards the packet to PE2.
                        c.   PE2 receives the packet from PE1 and decapsulates the packet to remove
                             Layer 2 encapsulation information and the two MPLS labels.
                        d.   PE2 sends the original Layer 2 packet to CE2.
                        The processing for sending a packet from CE2 to CE1 is similar to this process.
                    ●   VLAN access in tagged mode (with the U-tag)

                        Figure 6-5 VLAN access in tagged mode (with the U-tag)




                        As shown in Figure 6-5, ACs use VLAN encapsulation and PWs use tagged
                        encapsulation; packets transmitted from CEs to PEs carry U-tags and P-tags.
                        The packet exchange process in VLAN access in tagged mode (with the U-tag)
                        is as follows:
                        a.   CE1 sends to PE1 a packet that is encapsulated at Layer 2 and carries
                             both a U-tag and a P-tag.
                        b.   After receiving the packet, PE1 does not process the two tags. PE1 retains
                             the U-tag because it treats the U-tag as service data.
                        c.   PE1 retains the P-tag because a packet sent to a PW with the tagged
                             packet encapsulation mode must carry a P-tag.
                        d.   PE1 queries entries in the VSI and selects a tunnel and a PW for the
                             packet.
                        e.   PE1 adds two MPLS labels (outer tunnel label and inner VC label) to the
                             packet based on the selected tunnel and PW, performs Layer 2
                             encapsulation, and forwards the packet to PE2.
                        f.   PE2 receives the packet from PE1 and decapsulates the packet to remove
                             Layer 2 encapsulation information and the two MPLS labels.
                        g.   PE2 sends the decapsulated Layer 2 packet to CE2. The packet contains
                             the U-tag and the replaced P-tag.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            594
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


                        The processing for sending a packet from CE2 to CE1 is similar to this process.

