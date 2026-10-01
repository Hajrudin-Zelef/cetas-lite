---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-38
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [4710, 4904]
sha256: a4d3611b2398b3ab7ba21059cbd0b83333ed403e72332ca22d9ef5a3d1252b4b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  #
                  vlan batch 10
                  #
                  dhcp enable
                  #
                  dhcp snooping enable
                  user-bind static ip-address 10.0.0.3 mac-address 00e0-fc12-3489 interface 10GE 1/0/3 vlan 10
                  #
                  vlan 10
                   dhcp snooping enable
                   dhcp snooping trusted interface 10GE 1/0/4
                   ipv4 source check user-bind enable
                  #
                  interface 10GE 1/0/1
                   port link-type access
                   port default vlan 10
                  #
                  interface 10GE 1/0/2
                   port link-type access
                   port default vlan 10
                  #
                  interface 10GE 1/0/3
                   port link-type access
                   port default vlan 10
                  #
                  interface 10GE 1/0/4
                   port link-type trunk
                   port trunk allow-pass vlan 10
                  #
                  return

                  DeviceB
                  #
                  sysname DeviceB
                  #
                  vlan batch 10
                  #
                  dhcp enable
                  #
                  ip pool 10
                   gateway-list 10.1.1.1
                   network 10.1.1.0 mask 255.255.255.0
                  #
                  interface Vlanif10
                   ip address 10.1.1.1 255.255.255.0
                   dhcp select global
                  #
                  interface 10GE 1/0/1
                   port link-type trunk
                   port trunk allow-pass vlan 10
                  #
                  return



5.7 (Optional) Configuring the IP Packet Check Alarm
Function
Context
                  After the IP packet check alarm function is configured, the device generates logs
                  when IP packets are discarded. If the number of discarded packets reaches the
                  threshold, the device will send alarms to the NMS.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          83
Security Configuration
Security Configuration                                                                         5 IPSG Configuration


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Configure the IP packet check alarm function.
                  ●      Enable the IP packet check alarm function in the interface view.
                         interface interface-type interface-number
                         ip source check user-bind alarm enable
                         quit

                  ●      Enable the IP packet check alarm function in the VLAN view.
                         vlan vlan-id
                         ip source check user-bind alarm enable
                         quit

                  By default, the IP packet check alarm function is disabled.

         Step 3 Configure the threshold of the IP packet check alarm.
                  ●      Configure the threshold of the IP packet check alarm in the interface view.
                         interface interface-type interface-number
                         ip source check user-bind alarm threshold threshold
                         quit

                  ●      Configure the threshold of the IP packet check alarm in the VLAN view.
                         vlan vlan-id
                         ip source check user-bind alarm threshold threshold
                         quit

                  By default, the alarm threshold is 100.

                  ----End


5.8 (Optional) Configuring the Function of Discarding
IP Packets with Identical Source and Destination IP
Addresses
Context
                  IP packets typically have identical source and destination IP addresses only in
                  special scenarios, for example, a network administrator may construct such
                  packets for internal tests. By default, the device forwards these packets. However,
                  if you suspect that such packets are caused by a local area network denial (LAND)
                  attack, configure this function to discard the packets.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Configure the function of discarding IP packets with identical source and
                destination IP addresses.
                  ip anti-attack source-ip equals destination-ip drop { all | slot slot-id }

                  ----End

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                    84
Security Configuration
Security Configuration                                                                              5 IPSG Configuration




5.9 Maintaining IPSG
                  You can clear statistics on packets discarded due to IPSG in the user view if
                  required.



                         NOTICE

                  The cleared IPSG-discarded packet statistics cannot be restored. Exercise caution
                  when performing this operation.


                  Table 5-4 Clearing IPSG related statistics

                   Operation                                             Command

                   Clear statistics on packets discarded                 reset ip source check user-bind
                   due to IPSG.                                          statistics [ vlan vlan-id | interface
                                                                         interface-type | interface-number ]




5.10 Troubleshooting IPSG

5.10.1 IPSG Does Not Take Effect Because It Is Not Enabled on
an Interface or in a VLAN

Fault Symptom
                  Binding entries have been generated, but IPSG does not take effect.

Possible Causes
                  IPSG is not enabled on the specified interface or in the specified VLAN.

Procedure
         Step 1 Check whether IPSG is enabled on the user-side interface.
                  display ip source check user-bind status static [ { interface interface-type interface-number | ip-address
                  ip-address | ipv6-address ipv6-address [ ipv6-prefix ipv6-prefix ]| mac-address mac-address | vlan vlan-
                  id } * ] [ valid | invalid ] [ slot slot-id ]

         Step 2 If IPSG is not enabled on the interface, check whether IPSG is enabled on the user-
                side VLAN in the VLAN view.
                  display this

         Step 3 If IPSG is not enabled on the interface or VLAN, that is, ipv4 source check user-
                bind enable or ipv6 source check user-bind enable is not displayed in the
                command output, enable IPSG in the interface view or VLAN view.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                               85
Security Configuration
Security Configuration                                                                                5 IPSG Configuration


                  ●      Enable the IPv4 packet check function.
                  ipv4 source check user-bind enable

                  ●      Enable the IPv6 packet check function.
                  ipv6 source check user-bind enable

                          NOTE

