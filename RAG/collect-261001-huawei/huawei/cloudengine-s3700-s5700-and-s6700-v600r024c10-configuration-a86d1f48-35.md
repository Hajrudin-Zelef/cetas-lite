---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-35
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [4254, 4402]
sha256: 95434cf44236dff936b1950834a9f267b3aa834efe234ff1ee5a5158f21ad4ea
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

         Step 2 Configure static binding entries on 10GE1/0/1 and 10GE1/0/2 of DeviceA.
                  [DeviceA] user-bind static ip-address 10.0.0.1 mac-address 00e0-fc12-3456 interface 10GE 1/0/1
                  [DeviceA] user-bind static ip-address 10.0.0.2 mac-address 00e0-fc12-3478 interface 10GE 1/0/2

         Step 3 Configure the upstream interface 10GE1/0/4 as a trusted interface.
                  [DeviceA] dhcp enable
                  [DeviceA] dhcp snooping enable
                  [DeviceA] interface 10GE 1/0/4
                  [DeviceA-10GE1/0/4] dhcp snooping trusted
                  [DeviceA-10GE1/0/4] quit

         Step 4 Enable IPSG in VLAN 10.
                  [DeviceA] vlan 10
                  [DeviceA-vlan10] ipv4 source check user-bind enable
                  [DeviceA-vlan10] quit

                  ----End

Verifying the Configuration
                  # Display static binding entries.
                  [DeviceA] display ip source check user-bind status
                  User-bind table on slot 1:
                  --------------------------------------------------------------------------------
                  IP Address        Prefix          Vlan(O/I)         Interface Binding
                               MAC Address           Type             Status
                  --------------------------------------------------------------------------------
                  10.0.0.1         -             - /-          10GE1/0/1        DHCP
                              00e0-fc12-3456          Static           IPv4/-
                  10.0.0.2         -             - /-          10GE1/0/2        DHCP
                              00e0-fc12-3478          Static           IPv4/-
                  --------------------------------------------------------------------------------
                  Total count:           2


Configuration Scripts
                  DeviceA
                  #
                  sysname DeviceA
                  #


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         75
Security Configuration
Security Configuration                                                                             5 IPSG Configuration

                  dhcp enable
                  #
                  dhcp snooping enable
                  user-bind static ip-address 10.0.0.1 mac-address 00e0-fc12-3456 interface 10GE1/0/1
                  user-bind static ip-address 10.0.0.2 mac-address 00e0-fc12-3478 interface 10GE1/0/2
                  #
                  vlan batch 10
                  #
                  vlan 10
                   ipv4 source check user-bind enable
                  #
                  interface 10GE1/0/1
                   port link-type access
                   port default vlan 10
                  #
                  interface 10GE1/0/2
                   port link-type access
                   port default vlan 10
                  #
                  interface 10GE1/0/3
                   port link-type access
                   port default vlan 10
                  #
                  interface 10GE1/0/4
                   port link-type trunk
                   port trunk allow-pass vlan 10
                   dhcp snooping trusted
                  #
                  return



5.6 Configuring IPSG Based on a Dynamic Binding
Table

5.6.1 Enabling IPSG Based on a Dynamic Binding Table
Context
                  IPSG based on a dynamic binding table is applicable to LANs that contain many
                  hosts or the hosts obtain IP addresses through DHCP. The interfaces directly or
                  indirectly connected to the DHCP server are configured as trusted interfaces, and
                  other interfaces are untrusted. The device forwards the packets received by the
                  trusted interfaces without checking them against the binding entries. In this case,
                  authorized hosts can obtain IP addresses only from valid DHCP servers. After
                  enabling DHCP snooping on an interface connected to users or in the VLAN,
                  configure the interface connected to the DHCP server as a trusted interface, so
                  that the dynamic DHCP snooping binding table is generated.
                  After creating a dynamic binding table, you need to enable IPSG on an interface
                  or in a VLAN for IPSG to take effect.
                  ●      Enabling IPSG on an interface: IPSG checks all packets received by the
                         interface against binding entries. Use this method if you want to perform an
                         IPSG check on specified interfaces and trust other interfaces. This method is
                         suitable for scenarios in which an interface belongs to multiple VLANs,
                         because it eliminates the need to enable IPSG in each VLAN.
                  ●      Enabling IPSG in a VLAN: IPSG checks the packets received by all interfaces in
                         the VLAN against binding entries. Use this method if you want to perform an
                         IPSG check in specified VLANs and trust other VLANs. This method is suitable

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          76
Security Configuration
Security Configuration                                                                          5 IPSG Configuration


                         for scenarios in which multiple interfaces belong to the same VLAN, because
                         it eliminates the need to enable IPSG on each interface.
                          NOTE

                         ● If you use both of the preceding methods, only the one that is configured first takes
                           effect.
                         ● If ACL resources are insufficient, IPSG configurations exist on the device but the binding
                           table may fail to be delivered. To check whether IPSG takes effect, run the display ip
                           source check user-bind status dynamic [ { interface interface name | interface-type
                           interface-number | ip-address ip-address | ipv6-address ipv6-address [ ipv6-prefix ipv6-
                           prefix ] | mac-address mac-address | vlan vlan-id } * ] [ valid | invalid ] [ slot slot-id ]
                           command and check the value of Status in the command output.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Configure DHCP snooping so that a dynamic binding table is generated.
                  1.     Enable DHCP globally.
                         dhcp enable

                         By default, DHCP is disabled globally.
                  2.     Enable DHCP snooping globally.
                         dhcp snooping enable

                         By default, DHCP snooping is disabled globally.
                  3.     Configure a trusted interface.
                         –    Configure a trusted interface in the interface view.
                              interface interface-type interface-number
                              dhcp snooping enable
                              dhcp snooping trusted
                              quit

                         –    Configure a trusted interface in the VLAN view.
                              vlan vlan-id
                              dhcp snooping enable
                              dhcp snooping trusted interface interface-type interface-number
                              quit

