---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-18
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [3761, 3969]
sha256: 44ad50e691b7dd56d0b7cee6b5fd4b4ca04ba2054f0ad2528f259fd3be12fde3
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

[S2]interface GigabitEthernet 0/0/24 
[S2-GigabitEthernet0/0/24]port link-type trunk 
[S2-GigabitEthernet0/0/24]port trunk allow-pass vlan all 
[S2-GigabitEthernet0/0/24]bpdu enable 
[S2-GigabitEthernet0/0/24]interface GigabitEthernet 0/0/10 
[S2-GigabitEthernet0/0/10]port link-type trunk 
[S2-GigabitEthernet0/0/10]port trunk allow-pass vlan all 
[S2-GigabitEthernet0/0/10]bpdu enable 
 
[S3]interface  Ethernet0/0/1 
[S3-Ethernet0/0/1]port link-type trunk 
[S3-Ethernet0/0/1]port trunk allow-pass vlan all 
[S3-Ethernet0/0/1]bpdu enable 
[S3-Ethernet0/0/1]interface Ethernet0/0/13 
[S3-Ethernet0/0/13]port link-type trunk 
[S3-Ethernet0/0/13]port trunk allow-pass vlan all 
[S3-Ethernet0/0/13]bpdu enable 
 
[S4]interface  Ethernet0/0/1 
[S4-Ethernet0/0/1]port link-type trunk 
[S4-Ethernet0/0/1]port trunk allow-pass vlan all 
[S4-Ethernet0/0/1]bpdu enable

HCDP-IESN  Chapter 2 STP and SEP 
 
Page86 HUAWEI TECHNOLOGIES HC Series 
 
[S4-Ethernet0/0/1]interface Ethernet0/0/24 
[S4-Ethernet0/0/24]port link-type trunk 
[S4-Ethernet0/0/24]port trunk allow-pass vlan all 
[S4-Ethernet0/0/24]bpdu enable 
 
Step 2 Configure multi-instance MSTP . 
Enable MSTP in the system view. 
[S1]stp enable 
[S1]stp mode mstp 
 
[S2]stp enable 
[S2]stp mode mstp 
 
[S3]stp enable 
[S3]stp mode mstp 
 
[S4]stp enable 
[S4]stp mode mstp 
 
Configure all switches to belong to RG1 and set the revision level to 1. 
Create instance 1 and map it to VLANs 3, 4 and 5. Create instance 2 and map 
it to VLANs 6, 7, and 8. Then activate the region configurations. 
[S1]stp region-configuration 
[S1-mst-region]region-name RG1 
[S1-mst-region]revision-level 1 
[S1-mst-region]instance 1 vlan 3 4 5 
[S1-mst-region]instance 2 vlan 6 7 8 
[S1-mst-region]active region-configuration 
 
[S2]stp  region-configuration 
[S2-mst-region]region-name RG1 
[S2-mst-region]revision-level 1 
[S2-mst-region]instance 1 vlan 3 4 5 
[S2-mst-region]instance 2 vlan 6 7 8 
[S2-mst-region]active region-configuration 
 
[S3]stp  region-configuration 
[S3-mst-region]region-name RG1 
[S3-mst-region]revision-level 1

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page87 
 
[S3-mst-region]instance 1 vlan 3 4 5 
[S3-mst-region]instance 2 vlan 6 7 8 
[S3-mst-region]active region-configuration 
 
[S4]stp  region-configuration 
[S4-mst-region]region-name RG1 
[S4-mst-region]revision-level 1 
[S4-mst-region]instance 1 vlan 3 4 5 
[S4-mst-region]instance 2 vlan 6 7 8 
[S4-mst-region]active region-configuration 
 
View MSTP information. 
[S1]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/10       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/10       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/10       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
 
[S2]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   0    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   1    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   2    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
 
[S3]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               DESI  FORWARDING      NONE 
   0    Ethernet0/0/13              ROOT  FORWARDING      NONE 
   1    Ethernet0/0/1               DESI  FORWARDING      NONE 
   1    Ethernet0/0/13              ROOT  FORWARDING      NONE 
   2    Ethernet0/0/1               DESI  FORWARDING      NONE 
   2    Ethernet0/0/13              ROOT  FORWARDING      NONE 
 
[S4]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               ROOT  FORWARDING      NONE

HCDP-IESN  Chapter 2 STP and SEP 
 
Page88 HUAWEI TECHNOLOGIES HC Series 
 
   0    Ethernet0/0/24              ALTE  DISCARDING      NONE 
   1    Ethernet0/0/1               ROOT  FORWARDING      NONE 
   1    Ethernet0/0/24              ALTE  DISCARDING      NONE 
   2    Ethernet0/0/1               ROOT  FORWARDING      NONE 
   2    Ethernet0/0/24              ALTE  DISCARDING      NONE 
 
S1 is the root switch and the E0/0/24 interface of S4 is the alternate 
interface for all MSTP processes. 
In instance 2, set the priority to 0 for S2, to 4096 for S1, and to 8192 for S4 
so that S2 becomes the root switch in instance 2. 
[S2]stp instance 2 priority 0 
 
[S1]stp instance 2 priority 4096 
 
[S4]stp instance 2 priority 8192 
 
View MSTP information. 
[S1]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/10       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/10       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   2    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
 
[S2]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   0    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   1    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/10       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
 
[S3]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               DESI  FORWARDING      NONE 
   0    Ethernet0/0/13              ROOT  FORWARDING      NONE 
   1    Ethernet0/0/1               DESI  FORWARDING      NONE 
   1    Ethernet0/0/13              ROOT  FORWARDING      NONE

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page89 
 
   2    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   2    Ethernet0/0/13              ROOT  FORWARDING      NONE 
 
[S4]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   0    Ethernet0/0/24              ROOT  FORWARDING      NONE 
   1    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   1    Ethernet0/0/24              ROOT  FORWARDING      NONE 
   2    Ethernet0/0/1               DESI  FORWARDING      NONE 
   2    Ethernet0/0/24              ROOT  FORWARDING      NONE 
 
S2 becomes the root switch in instance 2 and the E0/0/1 interface of S3 
becomes the alternate interface. The status of switches does not change in 
instance 1. This proves that MSTP instances calculate spanning tree instances 
independently. 
Step 3 Configure multi-region MSTP. 
Delete the MSTP region and priority configurations made in step 2 for all 
switches. 
[S1]undo stp region-configuration 
[S1]undo stp instance 2 priority 
 
[S2]undo stp region-configuration 
[S2]undo stp instance 2 priority 
 
[S3]undo stp region-configuration 
 
[S4]undo stp region-configuration 
[S4]undo stp instance 2 priority 
 
Configure S1 and S3 to belong to RG1 and set the revision level to 1. 
Create instance 1 and map it to VLANs 3, 4, and 5. 
Create instance 2 and map it to VLANs 6, 7, and 8. 
[S1]stp region-configuration 
[S1-mst-region]region-name RG1 
[S1-mst-region]revision-level 1 
[S1-mst-region]instance 1 vlan 3 4 5

HCDP-IESN  Chapter 2 STP and SEP 
 
Page90 HUAWEI TECHNOLOGIES HC Series 
 
[S1-mst-region]instance 2 vlan 6 7 8 
[S1-mst-region]active region-configuration 
 
