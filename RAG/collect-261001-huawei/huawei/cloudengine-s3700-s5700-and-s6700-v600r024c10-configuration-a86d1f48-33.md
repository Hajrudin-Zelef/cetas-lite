---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-33
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [3956, 4082]
sha256: 2336fa2687a574516a91042c9a243c8320d11731a3781faaf0910bd5fd2610c0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Security Configuration
Security Configuration                                                                        5 IPSG Configuration


                   Parameter                                          Default Setting

                   IP packet check items                              ● IPSG based on a static binding
                                                                        table: checks all items in the
                                                                        binding table.
                                                                      ● IPSG based on a dynamic binding
                                                                        table: checks the source IP address,
                                                                        source MAC address, interface, and
                                                                        VLAN.

                   IP packet check alarm                              Disabled

                   IP packet check alarm threshold                    100

                   Discarding IP packets with identical               Disabled
                   source and destination IP addresses




5.5 Configuring IPSG Based on a Static Binding Table

5.5.1 Enabling IPSG Based on a Static Binding Table
Context
                  IPSG based on a static binding table is applicable to LANs that contain only a few
                  hosts with fixed IP addresses. After creating a static binding table, you need to
                  enable IPSG on an interface or in a VLAN for IPSG to take effect.
                  ●      Enabling IPSG on an interface: IPSG checks all packets received by the
                         interface against binding entries. Use this method if you want to perform an
                         IPSG check on specified interfaces and trust other interfaces. This method is
                         suitable for scenarios in which an interface belongs to multiple VLANs,
                         because it eliminates the need to enable IPSG in each VLAN.
                  ●      Enabling IPSG in a VLAN: IPSG checks the packets received by all interfaces in
                         the VLAN against binding entries. Use this method if you want to perform an
                         IPSG check in specified VLANs and trust other VLANs. This method is suitable
                         for scenarios in which multiple interfaces belong to the same VLAN, because
                         it eliminates the need to enable IPSG on each interface.
                          NOTE

                         ● If you configure IPSG both in VLANs and on interfaces, only the one that is configured
                           first takes effect.
                         ● If ACL resources are insufficient, IPSG configurations exist on the device but the binding
                           table may fail to be delivered. To check whether IPSG takes effect, run the display ip
                           source check user-bind status static [ [ { interface interface-type interface-number |
                           ip-address ip-address | ipv6-address ipv6-address [ ipv6-prefix ipv6-prefix ]| mac-
                           address mac-address | vlan vlan-id } * ] [ valid | invalid ] | summary ] [ slot slot-id ]
                           command and check the value of Status in the command output.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        70
Security Configuration
Security Configuration                                                                                 5 IPSG Configuration


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Configure static binding entries. Static binding entries include IPv4 and IPv6
                entries. Choose one type of entries according to your network type.
                  ●      Configure IPv4 static binding entries.
                  user-bind static { ip-address { start-ip [ to end-ip ] } &<1-10> | mac-address mac-address } * [ interface
                  { interface-type interface-number | interface-name } ] [ vlan vlan-id [ ce-vlan ce-vlan-id ] ]

                  ●      Configure IPv6 static binding entries.
                  user-bind static { { ipv6-address { start-ipv6 [ to end-ipv6 ] } &<1-10> | ipv6-prefix prefix/prefix-length } |
                  mac-address mac-address } * [ interface { interface-type interface-number | interface-name } ] [ vlan vlan-
                  id [ ce-vlan ce-vlan-id ] ]

                  By default, no static binding table is configured.

                          NOTE

                         IPSG matches packets against all items in static binding entries. Ensure that the created
                         binding table is correct and contains all the items to check. The device forwards the packets
                         from hosts only when the packets match all items in binding entries, and discards all other
                         packets.
                         The device can bind IP addresses or IP address segments in batches. For example, it can
                         bind multiple IP addresses to the same interface or MAC address.
                         ● To bind non-contiguous IP addresses, enter 1 to 10 IP addresses in start-ip. For example,
                           run the user-bind static ip-address 10.0.0.1 10.0.0.3 10.0.0.5 interface 10GE 1/0/1
                           command to bind IP addresses 10.0.0.1, 10.0.0.3, and 10.0.0.5 to the same interface.
                         ● To bind contiguous IP addresses, enter 1 to 10 IP address segments using start-ip to
                           end-ip. When the keyword to is used, the IP address segments cannot overlap. For
                           example, run the user-bind static ip-address 10.0.0.1 to 10.0.0.4 mac-address 00e0-
                           fc12-3456 command to bind IP addresses 10.0.0.1 to 10.0.0.4 to the MAC address 00e0-
                           fc12-3456.

         Step 3 (Optional) Configure a trusted interface.
                          NOTE

                         If the hosts on the network use static IP addresses, you do not need to configure trusted
                         interfaces. However, if the upstream interface is in an IPSG-enabled VLAN, you need to
                         configure the interface as a trusted interface. Otherwise, the return packets will be
                         discarded due to mismatched binding entries. For details, see 5.10.3 IP Packets Are
                         Discarded Because the Upstream Interface Is Not Trusted. After the upstream interface is
                         configured as trusted, the device forwards the packets received by the interface without
                         checking them against the binding entries.

                  1.     Enable DHCP globally.
                         dhcp enable

                         By default, DHCP is disabled globally.
                  2.     Enable DHCP snooping globally.
                         dhcp snooping enable

                         By default, DHCP snooping is disabled globally.
                  3.     Configure a trusted interface.
                         –    Configure a trusted interface in the interface view.
                              interface interface-type interface-number
                              dhcp snooping trusted
                              quit


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                    71
Security Configuration
Security Configuration                                                                          5 IPSG Configuration


