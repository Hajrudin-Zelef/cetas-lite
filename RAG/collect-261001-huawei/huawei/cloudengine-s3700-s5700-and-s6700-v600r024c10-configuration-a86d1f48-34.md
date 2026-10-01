---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-34
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [4083, 4253]
sha256: 34c2df8cbe7eba26a5a8515c34127841e9380126446e9508033fb5025f9f4200
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                         –    Configure a trusted interface in the VLAN view.
                              vlan vlan-id
                              dhcp snooping trusted interface interface-type interface-number
                              quit

                         By default, an interface is untrusted.

         Step 4 Enable IPSG.
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

                  ----End


Verifying the Configuration
                  ●      Run the display ip source check user-bind statistics [ interface interface-
                         type interface-number ] command to check statistics on packets discarded
                         due to IPSG.
                  ●      Run the display ip source check user-bind status [ [ static [ { interface
                         interface-type interface-number | ip-address ip-address| ipv6-address ipv6-
                         address [ ipv6-prefix ipv6-prefix ] | mac-address mac-address | vlan vlan-id }
                         * ] [ valid | invalid ] ] | summary ] [ slot slot-id ] command to check
                         information and status of IPSG static binding entries.

5.5.2 Example for Configuring IPSG Based on a Static Binding
Table on an Interface

Networking Requirements
                  In Figure 5-4, PC1 and PC2 access the network through DeviceA, and they both
                  use static IP addresses. The administrator wants users to use fixed IP addresses to
                  access the Internet.


                  Figure 5-4 Network diagram of configuring IPSG based on a static binding table
                  on an interface
                          NOTE

                         In this example, interface1 and interface2 represent 10GE1/0/1 and 10GE1/0/2, respectively.


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     72
Security Configuration
Security Configuration                                                                                           5 IPSG Configuration




Procedure
         Step 1 Create static binding entries on Device A.
                  # Create static binding entries on Device A.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceA
                  [DeviceA] user-bind static ip-address 10.0.0.1 mac-address 00e0-fc12-3456
                  [DeviceA] user-bind static ip-address 10.0.0.11 mac-address 00e0-fc12-3478

         Step 2 Enable IPSG.
                  # Enable IPSG on 10GE1/0/1, and 10GE1/0/2 of DeviceA.
                  [DeviceA] interface 10GE 1/0/1
                  [DeviceA-10GE1/0/1] ipv4 source check user-bind enable
                  [DeviceA-10GE1/0/1] quit
                  [DeviceA] interface 10GE 1/0/2
                  [DeviceA-10GE1/0/2] ipv4 source check user-bind enable
                  [DeviceA-10GE1/0/2] quit

                  ----End

Verifying the Configuration
                  # Display static binding entries.
                  [DeviceA] display ip source check user-bind status
                  User-bind table on slot 1:
                  ----------------------------------------------------------------------------------------------------------
                  IP Address        Prefix        Vlan(O/I)        Interface Binding
                               MAC Address          Type          Status
                  -----------------------------------------------------------------------------------------------------------
                  10.0.0.1          -           - /-          -        DHCP
                               00e0-fc12-3456 Static               IPv4/-
                  10.0.0.11          -           - /-          -       DHCP
                               00e0-fc12-3478 Static               IPv4/-
                  -----------------------------------------------------------------------------------------------------------
                  Total count:           2


Configuration Files
                  DeviceA
                  #
                  sysname DeviceA


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                     73
Security Configuration
Security Configuration                                                                         5 IPSG Configuration

                  #
                  user-bind static ip-address 10.0.0.1 mac-address 00e0-fc12-3456
                  user-bind static ip-address 10.0.0.11 mac-address 00e0-fc12-3478
                  #
                  interface 10GE1/0/1
                   ipv4 source check user-bind enable
                  #
                  interface 10GE1/0/2
                   ipv4 source check user-bind enable
                  #
                  return


5.5.3 Example for Configuring IPSG Based on a Static Binding
Table in a VLAN
Networking Requirements
                  In Figure 5-5, PC1 and PC2 access the network through DeviceA, and they both
                  use static IP addresses. The Gateway functions as the enterprise egress gateway.
                  The administrator wants the PCs to use fixed IP addresses to access the Internet
                  through fixed interfaces. For security purposes, the administrator does not allow
                  external hosts to access the intranet without permission.

                  Figure 5-5 Network diagram of configuring IPSG based on a static binding table
                  in a VLAN
                          NOTE

                         In this example, interface1, interface2, interface3, and interface4 represent 10GE1/0/1,
                         10GE1/0/2, 10GE1/0/3, and 10GE1/0/4, respectively.




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                        74
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
                  [DeviceA-10GE1/0/3] port link-type access
                  [DeviceA-10GE1/0/3] port default vlan 10
                  [DeviceA-10GE1/0/3] quit
                  [DeviceA] interface 10GE 1/0/4
                  [DeviceA-10GE1/0/4] port link-type trunk
                  [DeviceA-10GE1/0/4] port trunk allow-pass vlan 10
                  [DeviceA-10GE1/0/4] quit

