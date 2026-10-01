---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-10
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [806, 968]
sha256: 3930b56bdc2f99dbd2aafeb47886532a2a529b84500301e547168498392337e3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    For the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S, S6730E-
                    H-V2, S5755E-H, S5755-S and S5755-H series, the device supports a maximum of
                    2048 traffic classifiers.

         Step 3 Configure matching rules according to the following table.
                           NOTE

                         For details about precautions for each matching rule, see the corresponding commands in
                         the Command Reference.
                         For S5735-L-V2, S5735-S-V2, S5735I-L-V2, S5735I-S-V2, S5735I-H-V2, S5735R-L-V2, S3710-
                         H, S5735R-S-V2, S5735E-L-V2, S5735E-S-V2 series, a maximum of 1024 if-match
                         commands can be configured in a traffic classifier.
                         For the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S, S6730E-H-V2,
                         S5755E-H, S5755-S and S5755-H series, a maximum of 2048 if-match commands can be
                         configured in a traffic classifier.
                    ●     Link-layer rule (Layer 2 rule)

                           Matching Rule                    Command

                           Destination MAC                  if-match destination-mac mac-address [ mac-
                           address                          address-mask ]
                           Source MAC address               if-match source-mac mac-address [ mac-address-
                                                            mask ]
                           Protocol type in the             if-match l2-protocol { arp | ip | profinet | goose |
                           Ethernet frame                   rarp | protocol-value }
                           header

                           VLAN ID in the                   if-match vlan start-vlan-value [ inner-vlan start-
                           outer tag of VLAN                inner-vlan-value [ to end-inner-vlan-value ] ]
                           packets or VLAN ID               if-match vlan start-vlan-value [ to end-vlan-value ]
                           in the inner and                 [ inner-vlan start-inner-vlan-value ]
                           outer tags of QinQ
                           packets


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                     13
QoS Configuration
QoS Configuration                                                                   3 MQC Configuration


                         Matching Rule           Command

                         802.1p value in the     if-match 8021p 8021p-value &<1-8>
                         outer tag of VLAN
                         packets

                         Inner VLAN ID in        if-match inner-vlan start-inner-vlan-value [ to end-
                         QinQ packets            inner-vlan-value ]
                         802.1p value in the     if-match inner-8021p inner-8021p-value &<1-8>
                         inner tag of QinQ
                         packets

                         Inner and outer         if-match double-tag
                         VLAN IDs in QinQ
                         packets


                    ●   Network-layer rule (Layer 3 rule)
                         Matching Rule           Command

                         DSCP value in IP        if-match dscp dscp-value &<1-8>
                         packets

                         DSCP value in IPv6      if-match ipv6 dscp dscp-value &<1-8>
                         packets

                         IP precedence in IP     if-match ip-precedence ip-precedence &<1-8>
                         packets

                         ECN marking in IP       if-match ecn ecn-value
                         packets

                         ECN marking in IPv6     if-match ipv6 ecn ecn-value
                         packets


                    ●   Transport-layer rule (Layer 4 rule)
                         Matching Rule           Command

                         TCP flag in the TCP     if-match tcp-flag { tcp-flag-value | { ack | fin | psh |
                         packet header           rst | syn | urg } * }


                    ●   ACL rule
                         Matching Rule           Command

                         ACL rule                if-match acl { acl-number | acl-name }




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              14
QoS Configuration
QoS Configuration                                                                 3 MQC Configuration


                         Matching Rule          Command

                         ACL6 rule              if-match ipv6 acl { acl-number | acl6-name }
                                                [ loose-mode | strict-mode ]
                                                Only the S6780-H, S6750-S, S6730-H-V2, S6750-H,
                                                S5732-H-V2, S6750E-S, S6730E-H-V2, S5755E-H,
                                                S5755-S and S5755-H series support the loose-
                                                mode and strict-mode parameters.


                    ●   Forwarding rule
                         Matching Rule          Command

                         Dropped packets        if-match discard

                         All packets            if-match any

                         Inbound interface      if-match inbound-interface { { interface-type
                                                interface-number1 | interface-name1 } [ to
                                                { interface-type interface-number2 | interface-
                                                name2 } ] } &<1-8>
                         Interface              if-match outbound-interface { { interface-type
                                                interface-number1 | interface-name1 } [ to
                                                { interface-type interface-number2 | interface-
                                                name2 } ] } &<1-8>
                         Known unicast          if-match unicast
                         packets

                         Unknown unicast        if-match unknown-unicast
                         packets


                    ●   VXLAN packet rule
                         Matching Rule          Command

                         Inner information in   if-match vxlan [ transit ] acl { acl-number | acl-
                         VXLAN packets to be    name }
                         matched against
                         ACL rules

                         Inner information in   if-match vxlan [ transit ] ipv6 acl { acl-number |
                         VXLAN packets to be    acl-name } [ loose-mode | strict-mode ]
                         matched against        Only the S6730E-H-V2, S6730-H-V2, S6750E-S,
                         ACL6 rules             S6750-S, S5755E-H, S5755-H, and S5732-H-V2 series
                                                support the loose-mode and strict-mode
                                                parameters.

                         Inner information in   if-match vxlan [ transit ] vni vni-id
                         VXLAN packets




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                             15
QoS Configuration
QoS Configuration                                                                    3 MQC Configuration


                        Only the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S,
                        S6730E-H-V2, S5755E-H, S5755-S and S5755-H series support the commands
                        for configuring VXLAN packet matching rules.
                    ●   SRv6 packet rule
                         Matching Rule          Command

