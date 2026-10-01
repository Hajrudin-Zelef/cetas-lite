---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-40
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [5024, 5171]
sha256: ff191005d31bfa96865180348581315bfa3f8d79b509699307633119287e7c20
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

5.10.4 IP Packets Are Discarded After IPSG Is Enabled on the
Upstream Interface
Fault Symptom
                  After IPSG is enabled on the interfaces connected to hosts, none of the hosts can
                  access the Internet, and IP packets from authorized hosts are discarded.
                  In Figure 5-9, static binding entries are configured on DeviceA for PC1 and PC2,
                  IPSG is enabled on Interface1, Interface2, and Interface3, and the PCs can
                  communicate with each other. However, the PCs cannot access the Internet. The
                  following uses PC1 as an example.
                  ●      PC1 sends a packet to the Internet. When the packet reaches Interface1 on
                         DeviceA, DeviceA detects that the packet matches a binding entry, and
                         forwards it.
                  ●      A packet is sent from the Internet to PC1. When the packet reaches Interface3
                         on DeviceA, DeviceA detects that the packet does not match any binding
                         entry, and discards the packet.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                               88
Security Configuration
Security Configuration                                                                              5 IPSG Configuration


                  Figure 5-9 Network diagram for scenarios where IP packets are discarded after
                  IPSG is enabled on the upstream interface




Possible Causes
                  IPSG is enabled on the upstream interface, but this interface is not configured as a
                  trusted interface.


Procedure
         Step 1 Check whether IPSG is enabled on Interface3.
                  display ip source check user-bind status static [ { interface interface-type interface-number | ip-address
                  ip-address | ipv6-address ipv6-address [ ipv6-prefix ipv6-prefix ]| mac-address mac-address | vlan vlan-
                  id } * ] [ valid | invalid ] [ slot slot-id ]

         Step 2 If IPSG is enabled on the interface, that is, ipv4 source check user-bind enable or
                ipv6 source check user-bind enable is displayed in the command output, disable
                IPSG in the interface view.
                  ●      Disable the IPv4 packet check function.
                  undo ipv4 source check user-bind enable

                  ●      Disable the IPv6 packet check function.
                  undo ipv6 source check user-bind enable

                  ----End

5.10.5 IPSG Does Not Take Effect on Hosts with Dynamic IP
Addresses Because DHCP Snooping Is Not Configured

Fault Symptom
                  Hosts can dynamically obtain IP addresses from the DHCP server, but IPSG does
                  not take effect after being enabled.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                               89
Security Configuration
Security Configuration                                                                             5 IPSG Configuration


Possible Causes
                  DHCP is configured, but DHCP snooping is not configured.

Procedure
         Step 1 Check whether the DHCP snooping binding table exists. When hosts obtain IP
                addresses through DHCP, IPSG checks packets received by interfaces based on the
                DHCP snooping binding table. The device automatically generates a DHCP
                snooping binding table for online hosts only after DHCP snooping is enabled.
                  display ip source check user-bind status dynamic [ { interface interface-type interface-number | ip-
                  address ip-address | ipv6-address ipv6-address [ ipv6-prefix ipv6-prefix ]| mac-address mac-address |
                  vlan vlan-id } * ] [ valid | invalid ] [ slot slot-id ]

         Step 2 If the DHCP snooping binding table does not exist, configure it according to
                Configure a dynamic binding table.
                          NOTE

                         After DHCP snooping is configured and the hosts go online again, the device generates
                         DHCP snooping entries for the hosts. IPSG then takes effect. If you enable IPSG but the
                         device does not generate a DHCP snooping binding table, the device rejects all IP packets
                         except DHCP request packets. In this situation, communication between the DHCP hosts
                         and servers is affected. Therefore, before enabling IPSG, configure DHCP snooping to enable
                         the device to generate dynamic binding entries.

                  ----End




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                               90
Security Configuration
Security Configuration                                                    6 Port Security Configuration




                              6          Port Security Configuration


                  6.1 Overview of Port Security
                  6.2 Understanding Port Security
                  6.3 Configuration Precautions for Port Security
                  6.4 Default Settings for Port Security
                  6.5 Configuring Port Security


6.1 Overview of Port Security
Definition
                  Port security converts the dynamic MAC addresses learned on an interface into
                  secure MAC addresses.

Purpose
                  When unauthorized users acquire an interface's MAC address, they may use this
                  address as a destination MAC address to communicate with the device, initiating
                  attacks.
                  To prevent such attacks, a device can enable port security to convert the dynamic
                  MAC addresses learned on an interface into secure MAC addresses. After port
                  security is enabled, dynamic MAC address entries that have been learned on the
                  interface are deleted. When the number of MAC addresses learned again by the
                  interface reaches the limit, the interface stops learning MAC addresses. If an
                  interface receives packets whose source MAC address is not in the MAC address
                  entries, these packets are considered unauthorized. Then the actions including
                  discarding the packets, reporting alarms, or shutting down the interface are
                  performed.




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                             91
Security Configuration
Security Configuration                                                    6 Port Security Configuration




6.2 Understanding Port Security
Classification of Secure MAC Addresses

                  Table 6-1 Classification of secure MAC addresses
                   Type      Description           Characteristic             Scenario

