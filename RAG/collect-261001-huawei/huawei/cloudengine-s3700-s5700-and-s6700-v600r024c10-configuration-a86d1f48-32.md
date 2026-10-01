---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-32
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [3800, 3955]
sha256: 5d16b0fef1a90f02719c1285f4601b2f1b81b6113e6a01499ebcc3720b88a5bb
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                          NOTE

                         Only IPv6 static binding entries have masks.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          66
Security Configuration
Security Configuration                                                                       5 IPSG Configuration


                  Table 5-1 Binding tables
                   Type                              Description                     Application Scenario

                   Static binding table              A static binding table is       A network with a few
                                                     manually configured             IPv4/IPv6 hosts that use
                                                     using the user-bind             static IP addresses.
                                                     static command.

                   DHCP snooping binding             After DHCP snooping is          A network with many
                   table                             configured, hosts request       IPv4/IPv6 hosts that
                                                     IP addresses from the           obtain IP addresses from
                                                     DHCP server. The device         the DHCP server.
                                                     dynamically generates
                                                     DHCP snooping binding
                                                     entries according to the
                                                     DHCP reply packets
                                                     returned by the DHCP
                                                     server.


                  After the binding table is generated, IPSG delivers ACL rules to the specified
                  interface or VLAN according to the binding table, and then checks all IP packets
                  against the ACL rules. The device forwards only the packets that match binding
                  entries. When the binding table is modified, IPSG delivers the ACL rules again. By
                  default, if IPSG is enabled but no binding table is generated, the device forwards
                  IP protocol packets (except for IGMP protocol packets) and rejects all IP data
                  packets.

                          NOTE

                         IPSG checks only the IP packets. It does not check non-IP packets such as ARP packets.

                  Figure 5-2 illustrates the IPSG working mechanism. When an unauthorized host
                  forges an authorized host's IP address to send packets to the Device, the Device
                  discards these packets because they do not match binding entries.

                  Figure 5-2 IPSG working mechanism




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       67
Security Configuration
Security Configuration                                                                5 IPSG Configuration


                  Typically, IPSG is configured on the interfaces or VLANs of the user-side access
                  device.

                  ●      After IPSG is enabled on a user-side interface, it checks all IP packets received
                         by the interface against binding entries.
                  ●      After IPSG is enabled in a user-side VLAN, it checks the IP packets received by
                         all interfaces in the VLAN against binding entries.
                  ●      If the user-side access device does not support IPSG, you can configure IPSG
                         on the interfaces or in VLANs of the upper-layer device.


IPSG Interface Roles
                  IPSG can be configured only on Layer 2 physical interfaces or in VLANs. It also
                  checks the packets only on the untrusted interfaces with IPSG enabled and
                  considers all interfaces to be untrusted by default (you can specify trusted
                  interfaces). The trusted and untrusted interfaces in IPSG are the same as those
                  used in DHCP snooping. In addition, these interfaces are also valid for IPSG based
                  on a static binding table.

                  Figure 5-3 shows the IPSG interface roles.

                  ●      Interface1 and Interface2 are untrusted interfaces with IPSG enabled. The
                         Device performs an IPSG check on the packets received by these interfaces.
                  ●      Interface3 is an untrusted interface with IPSG disabled. The device does not
                         perform an IPSG check on the packets received by this interface.
                         Consequently, Interface3 is prone to attacks.
                  ●      Interface4 is a trusted interface specified by users. The device does not
                         perform an IPSG check on the packets received by this interface; however, it is
                         not prone to attacks. On a network with DHCP snooping configured, the
                         interfaces directly or indirectly connected to a valid DHCP server are generally
                         configured as trusted interfaces.


                  Figure 5-3 IPSG interface roles




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                              68
Security Configuration
Security Configuration                                                                    5 IPSG Configuration


IPSG Filtering
                  A binding entry contains the following items: MAC address, IP address, VLAN ID,
                  mask, and inbound interface. IPSG checks received packets against all items in a
                  static binding table. For a dynamic binding table, this is also the default
                  implementation, but you can additionally specify the items against which IPSG
                  performs checks. Table 5-2 describes some common check items.

                          NOTE

                         Only IPv6 static binding entries have masks.


                  Table 5-2 IPSG filtering
                   Item                        Description

                   Source IP address           The device forwards only the packets whose source IP
                                               addresses match binding entries.

                   Source MAC                  The device forwards only the packets whose source MAC
                   address                     addresses match binding entries.

                   Source IP address +         The device forwards only the packets whose source IP and
                   source MAC                  MAC addresses match binding entries.
                   address

                   Source IP address +         The device forwards only the packets whose source IP
                   source MAC                  addresses, source MAC addresses, and interfaces match
                   address + interface         binding entries.

                   Source IP address +         The device forwards only the packets whose source IP
                   source MAC                  addresses, source MAC addresses, interfaces, and VLANs
                   address + interface         match binding entries.
                   + VLAN




5.3 Configuration Precautions for IPSG

5.4 Default Settings for IPSG
                  Table 5-3 describes the default settings for IPSG.

                  Table 5-3 Default settings for IPSG
                   Parameter                                            Default Setting

                   IP packet check                                      Disabled




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                69

