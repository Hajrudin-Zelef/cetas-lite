---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-16
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [3308, 3530]
sha256: 9034607e5cec4d887bcefdad16ab9c1f7e8d3b80e52ab6106aab9aa8730e1325
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

[S4]vlan batch 1 to 20 
Info: This operation may take a few seconds. Please wait for a moment...done. 
[S4]interface Ethernet0/0/1 
[S4-Ethernet0/0/1]port link-type trunk 
[S4-Ethernet0/0/1]port trunk allow-pass vlan 1 TO 20 
[S4-Ethernet0/0/1]interface Ethernet0/0/14 
[S4-Ethernet0/0/14]port link-type trunk 
[S4-Ethernet0/0/14]port trunk allow-pass vlan 1 TO 20 
[S4-Ethernet0/0/14]interface Ethernet0/0/24 
[S4-Ethernet0/0/24]port link-type trunk 
[S4-Ethernet0/0/24]port trunk allow-pass vlan 1 TO 20 
 
Configure MSTP. 
Configure VLANs 110 to belong to instance 1 and VLANs 1120 to belong 
to instance 2. 
[S1]stp mode mstp 
[S1]stp region-configuration 
[S1-mst-region]region-name RG1 
[S1-mst-region]instance 1 vlan 1 TO 10 
[S1-mst-region]instance 2 vlan 11 to 20 
[S1-mst-region]active region-configuration 
Info: This operation may take a few seconds. Please wait for a moment....done. 
 
[S2]stp mode mstp 
[S2]stp region-configuration 
[S2-mst-region]region-name RG1 
[S2-mst-region]instance 1 vlan 1 TO 10 
[S2-mst-region]instance 2 vlan 11 to 20 
[S2-mst-region]active region-configuration 
Info: This operation may take a few seconds. Please wait for a moment....done. 
 
[S3]STP mode mstp 
Info: This operation may take a few seconds. Please wait for a moment.....done. 
[S3]stp region-configuration 
[S3-mst-region]region-name RG1 
[S3-mst-region]instance 1 vlan 1 to 10

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page75 
 
[S3-mst-region]instance 2 vlan 11 to 20 
 
[S4]STP mode mstp 
Info: This operation may take a few seconds. Please wait for a moment.....done. 
[S4]stp region-configuration 
[S4-mst-region]region-name RG1 
[S4-mst-region]instance 1 vlan 1 to 10 
[S4-mst-region]instance 2 vlan 11 to 20 
 
View the mapping between MSTP instances and VLANs. 
[S1]display stp region-configuration 
 Oper configuration 
   Format selector    :0 
   Region name        :RG1 
   Revision level     :0 
   Instance   VLANs Mapped 
      0       21 to 4094 
      1       1 to 10 
      2       11 to 20 
 
Set the priority of S1 to 4096 in instance 1 and to 8192 in instance 2. 
Set the priority of S2 to 4096 in instance 2 and to 8192 in instance 1. 
[S1]stp instance 1 priority 4096 
[S1]stp instance 2 priority 8192 
 
[S2]stp instance 2 priority 4096 
[S2]stp instance 1 priority 8192 
 
View status information about instance 1 and instance 2. 
[S1]display stp instance 1 
-------[MSTI 1 Global Info]------- 
MSTI Bridge ID      :4096.4c1f-cc45-aadc 
MSTI RegRoot/IRPC   :4096.4c1f-cc45-aadc / 0 
MSTI RootPortId     :0.0 
Master Bridge       :4096.4c1f-cc45-aac1 
Cost to Master      :20000 
TC received         :20 
TC count per hello  :0 
ĂĂoutput omitĂĂ 
 
[S2]display stp instance 2

HCDP-IESN  Chapter 2 STP and SEP 
 
Page76 HUAWEI TECHNOLOGIES HC Series 
 
-------[MSTI 2 Global Info]------- 
MSTI Bridge ID      :4096.4c1f-cc45-aac1 
MSTI RegRoot/IRPC   :4096.4c1f-cc45-aac1 / 0 
MSTI RootPortId     :0.0 
Master Bridge       :4096.4c1f-cc45-aac1 
Cost to Master      :0 
TC received         :16 
TC count per hello  :0 
ĂĂoutput omitĂĂ 
 
S1 is the root bridge of instance 1, and S2 is the root bridge of instance 2. 
View the role information about interfaces in instance 1. 
[S1]display stp instance 1 brief 
 MSTID  Port                        Role  STP State     Protection 
   1    GigabitEthernet0/0/9        DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/10       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/14       DESI  FORWARDING      NONE 
 
[S2]display stp instance 1 brief 
 MSTID  Port                        Role  STP State     Protection 
   1    GigabitEthernet0/0/9        ROOT  FORWARDING      NONE 
   1    GigabitEthernet0/0/10       ALTE  DISCARDING      NONE 
   1    GigabitEthernet0/0/23       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
 
[S3]display stp instance 1 brief 
 MSTID  Port                        Role  STP State     Protection 
   1    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   1    Ethernet0/0/13              ROOT  FORWARDING      NONE 
   1    Ethernet0/0/23              ALTE  DISCARDING      NONE 
 
[S4]display stp instance 1 brief 
 MSTID  Port                        Role  STP State     Protection 
   1    Ethernet0/0/1               DESI  FORWARDING      NONE 
   1    Ethernet0/0/14              ROOT  FORWARDING      NONE 
   1    Ethernet0/0/24              ALTE  DISCARDING      NONE 
 
In instance 1, S1 is the root bridge. Users in VLAN 1 to VLAN 10 of S3 
communicate with users in VLAN 1 to VLAN 10 of S1, S2, and S4 over the 
Ethernet 0/0/13 interface.

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page77 
 
View the role information about interfaces in instance 2. 
[S1]display stp instance 2 brief 
 MSTID  Port                        Role  STP State     Protection 
   2    GigabitEthernet0/0/9        ROOT  FORWARDING      NONE 
   2    GigabitEthernet0/0/10       ALTE  DISCARDING      NONE 
   2    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
 
[S2]display stp instance 2 brief 
 MSTID  Port                        Role  STP State     Protection 
   2    GigabitEthernet0/0/9        DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/10       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/23       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
 
[S3]display stp instance 2 brief 
 MSTID  Port                        Role  STP State     Protection 
   2    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   2    Ethernet0/0/13              ALTE  DISCARDING      NONE 
   2    Ethernet0/0/23              ROOT  FORWARDING      NONE 
 
[S4]display stp instance 2 brief 
 MSTID  Port                        Role  STP State     Protection 
   2    Ethernet0/0/1               DESI  FORWARDING      NONE 
   2    Ethernet0/0/14              DESI  FORWARDING      NONE 
   2    Ethernet0/0/24              ROOT  FORWARDING      NONE 
 
In instance 2, S2 is the root bridge. Users in VLAN 11 to VLAN 20 of S3 
communicate with users in VLAN 11 to VLAN 20 of S1, S2, and S4 over the 
Ethernet 0/0/23 interface. 
Additional Exercises: Analyzing and Verifying 
Figure out how MSTP balances data transmission load among VLANs in 
different areas. 
Figure out the reasons why RSTP can forward data quickly. 
Final Configurations 
[S1]display current-configuration 
# 
!Software Version V100R006C00SPC800

HCDP-IESN  Chapter 2 STP and SEP 
 
Page78 HUAWEI TECHNOLOGIES HC Series 
 
 sysname S1 
# 
vlan batch 2 to 20 
# 
 stp instance 0 priority 8192 
 stp instance 1 priority 4096 
stp instance 2 priority 8192 
# 
 stp region-configuration 
  region-name RG1 
  instance 1 vlan 1 to 10 
  instance 2 vlan 11 to 20 
  active region-configuration 
# 
interface Vlanif1 
 ip address 10.0.1.1 255.255.255.0 
# 
interface GigabitEthernet0/0/9 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 20 
# 
interface GigabitEthernet0/0/10 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 20 
# 
interface GigabitEthernet0/0/13 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 20 
# 
Return 
 
[S2]display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S2 
# 
 vlan batch 2 to 20 
# 
stp instance 1 priority 8192 
 stp instance 2 priority 4096 
 stp instance 0 root secondary 
# 
 stp region-configuration

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page79 
 
