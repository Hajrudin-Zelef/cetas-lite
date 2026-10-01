---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-36
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [4403, 4546]
sha256: d17d778cbb78163bd8a929d7528252a2fdf3319390c5478a94a7c010e1ee13e4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

         Step 3 Enable IPSG.
                  ●      Enable IPSG in the interface view. Use either of the following methods based
                         on the network environment.
                         interface interface-type interface-number
                         ipv4 source check user-bind enable
                         quit
                         interface interface-type interface-number
                         ipv6 source check user-bind enable
                         quit

                  ●      Enable IPSG in the VLAN view. Use either of the following methods based on
                         the network environment.
                         vlan vlan-id
                         ipv4 source check user-bind enable
                         quit
                         vlan vlan-id
                         ipv6 source check user-bind enable
                         quit

                  By default, IPSG is disabled.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                        77
Security Configuration
Security Configuration                                                                                5 IPSG Configuration


         Step 4 (Optional) Configure the IP packet check items.

                  After IPSG is enabled, the device checks received IP packets against the binding
                  table and forwards only those that match the binding entries. Among the four
                  check items of IP packets, the interface is mandatory, and the source IP address,
                  source MAC address, and VLAN are optional.

                  ●      Configure IP packet check items in the interface view.
                         interface interface-type interface-number
                         ip source check user-bind check-item { ip-address | mac-address | vlan } *
                         quit

                  ●      Configure IP packet check items in the VLAN view.
                         vlan vlan-id
                         ip source check user-bind check-item { ip-address | mac-address | interface } *
                         quit

                  By default, the IP packet check items include the IP address, MAC address, VLAN,
                  and interface. If some items are trusted or unfixed (for example, packets from
                  hosts may be received by different interfaces), you can perform this step. The
                  default items are recommended.

                  ----End

Verifying the Configuration
                  ●      Run the display ip source check user-bind configuration [ vlan vlan-id |
                         interface interface-type interface-number ] command to check the IPSG
                         configuration on an interface or in a VLAN.
                  ●      Run the display ip source check user-bind statistics [ interface interface-
                         type interface-number ] command to check statistics on packets discarded
                         due to IPSG.
                  ●      Run the display ip source check user-bind status [ [ dynamic [ { interface
                         interface-type interface-number | ip-address ip-address | ipv6-address ipv6-
                         address [ ipv6-prefix ipv6-prefix ] | mac-address mac-address | vlan vlan-id }
                         * ] [ valid | invalid ] ] | summary ] [ slot slot-id ] command to check
                         information and status of IPSG dynamic binding entries.

5.6.2 Example for Configuring IPSG Based on a Dynamic
Binding Table in a VLAN

Networking Requirements
                  In Figure 5-6, PC1 and PC2 access the network through DeviceA. The
                  administrator wants the PCs to use dynamically allocated IP addresses to access
                  the Internet and deny the access to the Internet if statically configured IP
                  addresses are used.

                  Figure 5-6 Network diagram of configuring IPSG based on a dynamic binding
                  table in a VLAN
                          NOTE

                         In this example, interface1, interface2, and interface3 represent 10GE1/0/1, 10GE1/0/2, and
                         10GE1/0/3, respectively.


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                           78
Security Configuration
Security Configuration                                                                               5 IPSG Configuration




Procedure
         Step 1 Create a VLAN and add interfaces to the VLAN.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceA
                  [DeviceA] vlan batch 10
                  [DeviceA] interface 10GE 1/0/1
                  [DeviceA-10GE1/0/1] port link-type access
                  [DeviceA-10GE1/0/1] port default vlan 10
                  [DeviceA-10GE1/0/1] quit
                  [DeviceA] interface 10GE 1/0/2
                  [DeviceA-10GE1/0/2] port link-type access
                  [DeviceA-10GE1/0/2] port default vlan 10
                  [DeviceA-10GE1/0/2] quit
                  [DeviceA] interface 10GE 1/0/3
                  [DeviceA-10GE1/0/3] port link-type trunk
                  [DeviceA-10GE1/0/3] port trunk allow-pass vlan 10
                  [DeviceA-10GE1/0/3] quit

         Step 2 Enable DHCP snooping and configure 10GE1/0/3 for connecting to the DHCP
                server as a trusted interface.
                  [DeviceA] dhcp enable
                  [DeviceA] dhcp snooping enable
                  [DeviceA] vlan 10
                  [DeviceA-vlan10] dhcp snooping enable
                  [DeviceA-vlan10] dhcp snooping trusted interface 10GE 1/0/3

         Step 3 Enable IPSG in VLAN 10 of DeviceA.
                  [DeviceA-vlan10] ipv4 source check user-bind enable
                  [DeviceA-vlan10] quit

                  ----End

Verifying the Configuration
                  # Display dynamic binding entries.
                  [DeviceA] display ip source check user-bind status
                  User-bind table on slot 1:
                  --------------------------------------------------------------------------------


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         79
Security Configuration
Security Configuration                                                                               5 IPSG Configuration

                  IP Address               Prefix          Vlan(O/I)        Interface Binding
                                       MAC Address          Type           Status
                  --------------------------------------------------------------------------------
                  10.1.1.254               -              10 /-          10GE1/0/1       DHCP
                                      00e0-fc12-3456        Dynamic            IPv4/-
                  10.1.1.253               -              10 /-          10GE1/0/2       DHCP
                                      00e0-fc12-3478        Dynamic            IPv4/-
                  --------------------------------------------------------------------------------
                  Total count:          2


