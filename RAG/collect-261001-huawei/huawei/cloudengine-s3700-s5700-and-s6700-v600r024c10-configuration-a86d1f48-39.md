---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-39
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [4905, 5023]
sha256: 0e8b813bd3e32c28df250d08ddcb0823c1cfbf6003eaf53dfd8326368c84c5e9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                         You need to enable IPSG on the interface or in the VLAN. The differences are as follows:
                         ● Enabling IPSG on an interface: IPSG checks all packets received by the interface against
                           binding entries. Use this method if you want to perform an IPSG check on specified
                           interfaces and trust other interfaces. This method is suitable for scenarios in which an
                           interface belongs to multiple VLANs, because it eliminates the need to enable IPSG in
                           each VLAN.
                         ● Enabling IPSG in a VLAN: IPSG checks the packets received by all interfaces in the VLAN
                           against binding entries. Use this method if you want to perform an IPSG check in
                           specified VLANs and trust other VLANs. This method is suitable for scenarios in which
                           multiple interfaces belong to the same VLAN, because it eliminates the need to enable
                           IPSG on each interface.
                         IPSG takes effect only on the interface or in the VLAN where it is enabled, and an IPSG
                         check will not be performed on the IPSG-disabled interfaces or in the IPSG-disabled VLANs.
                         As such, if IPSG does not take effect on an interface or in a VLAN, it may not be enabled on
                         this interface or in this VLAN.

                  ----End

5.10.2 IP Packets of Authorized Hosts Are Discarded Because
Static Binding Entries Are Incorrect
Fault Symptom
                  A static binding table has been created and IPSG has been enabled; however, the
                  IP packets of an authorized host are discarded.

Possible Causes
                  The binding entries are improperly configured.

Procedure
         Step 1 Check whether the binding entries are correct.
                  display ip source check user-bind status static [ { interface interface-type interface-number | ip-address
                  ip-address | ipv6-address ipv6-address [ ipv6-prefix ipv6-prefix ]| mac-address mac-address | vlan vlan-
                  id } * ] [ valid | invalid ] [ slot slot-id ]

         Step 2 If the binding entry of the host is not in the binding table, add the host's binding
                entry to the binding table. The device forwards the packets from a host only when
                the host's binding entry exists in the binding table.
                  user-bind static { ip-address { start-ip [ to end-ip ] } &<1-10> | ipv6-address { start-ipv6-address [ to end-
                  ipv6-address ] } &<1-10> | ipv6-prefix prefix | mac-address mac-address } * [ interface { interface-type
                  interface-number | interface-name }] [ vlan vlan-id [ ce-vlan ce-vlan-id ] ]

         Step 3 If the host's binding entry exists in the binding table, check whether the MAC
                address in the entry is the same as the host's MAC address. If the NIC of the host
                has been replaced but the MAC address in the entry has not been updated, delete
                the original entry and configure a new one for the host.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                  86
Security Configuration
Security Configuration                                                                              5 IPSG Configuration

                  undo user-bind static [ ip-address { start-ip [ to end-ip ] } &<1-10> | ipv6-address { start-ipv6-address
                  [ to end-ipv6-address ] } &<1-10> | ipv6-prefix prefix | mac-address mac-address | interface { interface-
                  type interface-number | interface-name } | vlan vlan-id [ ce-vlan ce-vlan-id ] ] *

         Step 4 Check whether the entry contains VLAN information. If so, check whether the
                interface connected to the host has been added to the VLAN. The device forwards
                the packets from the host only when the interface is added to the VLAN.
                  display vlan

                  ----End

5.10.3 IP Packets Are Discarded Because the Upstream
Interface Is Not Trusted

Fault Symptom
                  After IPSG is enabled in a VLAN, none of the hosts in the VLAN can access the
                  Internet, and IP packets from these hosts are discarded.

                  In Figure 5-8, PC1 and PC2 belong to VLAN 10, and Interface1 allows packets
                  from VLAN 10. The static binding entries of PC1 and PC2 have been configured on
                  DeviceA, and IPSG has been enabled in VLAN 10. The PCs can communicate with
                  each other, but cannot access the Internet. The following uses PC1 as an example.
                  ●      PC1 sends a packet to the Internet. When the packet reaches Interface1 on
                         DeviceA, DeviceA detects that the packet matches a binding entry, and
                         forwards it.
                  ●      A packet is sent from the Internet to PC1. When the packet reaches Interface3
                         (which belongs to VLAN 10) on DeviceA, DeviceA detects that the packet does
                         not match any binding entry, and discards the packet.


                  Figure 5-8 Network diagram for scenarios where IP packets are discarded because
                  the upstream interface is not trusted




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                   87
Security Configuration
Security Configuration                                                                              5 IPSG Configuration


Possible Causes
                  The upstream interface is not configured as a trusted interface in the IPSG-
                  enabled VLAN.

Procedure
         Step 1 Check whether the upstream interface belongs to the IPSG-enabled VLAN.
                  display ip source check user-bind status static [ { interface interface-type interface-number | ip-address
                  ip-address | ipv6-address ipv6-address [ ipv6-prefix ipv6-prefix ]| mac-address mac-address | vlan vlan-
                  id } * ] [ valid | invalid ] [ slot slot-id ]

         Step 2 If the upstream interface belongs to the VLAN, configure the interface as a trusted
                interface; otherwise, the return packets will be discarded because they do not
                match the binding entries.
                  1.     Enter the system view.
                         system-view

                  2.     Enable DHCP globally.
                         dhcp enable

                  3.     Enable DHCP snooping globally.
                         dhcp snooping enable

                  4.     Configure Interface3 as a trusted interface in the interface view.
                         dhcp snooping trusted

                  ----End

