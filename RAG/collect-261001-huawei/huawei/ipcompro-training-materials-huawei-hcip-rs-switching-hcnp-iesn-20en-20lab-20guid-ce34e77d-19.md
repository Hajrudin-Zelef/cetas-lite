---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-19
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [3970, 4152]
sha256: 9dadb01b9509dcf3997693aee835032bcf4a22f149f414a5b9c41394dda696f5
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

[S3]stp region-configuration 
[S3-mst-region]region-name RG1 
[S3-mst-region]revision-level 1 
[S3-mst-region]instance 1 vlan 3 4 5 
[S3-mst-region]instance 2 vlan 6 7 8 
[S3-mst-region]active region-configuration 
 
Configure S2 and S4 to belong to RG2 and set the revision level to 2. 
Create instance 1 and map it to VLANs 3, 4, and 5. 
Create instance 2 and map it to VLANs 6, 7, and 8. Then activate the 
region configurations. 
[S2]stp region-configuration 
[S2-mst-region]region-name RG2 
[S2-mst-region]revision-level 2 
[S2-mst-region]instance 1 vlan 3 4 5 
[S2-mst-region]instance 2 vlan 6 7 8 
[S2-mst-region]active region-configuration 
 
[S4]stp region-configuration 
[S4-mst-region]region-name RG2 
[S4-mst-region]revision-level 2 
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

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page91 
 
   0    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   0    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/10       MAST  FORWARDING      NONE 
   1    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/10       MAST  FORWARDING      NONE 
 
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
   0    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   0    Ethernet0/0/24              ROOT  FORWARDING      NONE 
   1    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   1    Ethernet0/0/24              ROOT  FORWARDING      NONE 
   2    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   2    Ethernet0/0/24              ROOT  FORWARDING      NONE 
 
S1 is the root switch and the E0/0/1 interface of S4 is the alternate 
interface. 
Set the priority to 0 for S3 in instance 0 so that S3 is the Common and 
Internal Spanning Tree (CIST) root, and to 0 in instance 1 so that S3 is the 
regional root in instance 1. Set the priority to 0 for S4 in instance 1 so that S4 is 
the regional root in instance 1. 
[S3]stp instance 0 priority 0 
[S3]stp instance 1 priority 0 
 
[S4]stp instance 1 priority 0 
 
View MSTP information. 
[S1]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/10       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/13       ROOT  FORWARDING      NONE 
   1    GigabitEthernet0/0/10       DESI  FORWARDING      NONE

HCDP-IESN  Chapter 2 STP and SEP 
 
Page92 HUAWEI TECHNOLOGIES HC Series 
 
   1    GigabitEthernet0/0/13       ROOT  FORWARDING      NONE 
   2    GigabitEthernet0/0/10       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
 
[S2]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   0    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/10       MAST  FORWARDING      NONE 
   1    GigabitEthernet0/0/24       ROOT  FORWARDING      NONE 
   2    GigabitEthernet0/0/10       MAST  FORWARDING      NONE 
   2    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
 
[S3]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               DESI  FORWARDING      NONE 
   0    Ethernet0/0/13              DESI  FORWARDING      NONE 
   1    Ethernet0/0/1               DESI  FORWARDING      NONE 
   1    Ethernet0/0/13              DESI  FORWARDING      NONE 
   2    Ethernet0/0/1               DESI  FORWARDING      NONE 
   2    Ethernet0/0/13              ROOT  FORWARDING      NONE 
 
[S4]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   0    Ethernet0/0/24              ROOT  FORWARDING      NONE 
   1    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   1    Ethernet0/0/24              DESI  FORWARDING      NONE 
   2    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   2    Ethernet0/0/24              ROOT  FORWARDING      NONE 
 
Delete the MSTP configurations from S2 and S4. Configure S2 and S4 to 
belong to RG2 and set the revision level to 2. Create instance 1 and map it to 
VLANs 6, 7, and 8. Create instance 2 a nd map it to VLANs 3, 4 and 5. Then 
activate the region configurations. 
[S2]undo stp region-configuration 
 
[S3]undo stp instance 0 priority 
[S3]undo stp instance 1 priority 
 
[S4]undo stp region-configuration

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page93 
 
[S4]undo stp instance 1 priority 
 
[S2]stp region-configuration 
[S2-mst-region]region-name RG2 
[S2-mst-region]revision-level 2 
[S2-mst-region]instance 1 vlan 6 7 8 
[S2-mst-region]instance 2 vlan 3 4 5 
[S2-mst-region]active region-configuration 
 
[S4]stp region-configuration 
[S4-mst-region]region-name RG2 
[S4-mst-region]revision-level 2 
[S4-mst-region]instance 1 vlan 6 7 8 
[S4-mst-region]instance 2 vlan 3 4 5 
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
   1    GigabitEthernet0/0/10       MAST  FORWARDING      NONE 
   1    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/10       MAST  FORWARDING      NONE 
   2    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
 
[S3]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               DESI  FORWARDING      NONE 
   0    Ethernet0/0/13              ROOT  FORWARDING      NONE 
   1    Ethernet0/0/1               DESI  FORWARDING      NONE 
   1    Ethernet0/0/13              ROOT  FORWARDING      NONE 
   2    Ethernet0/0/1               DESI  FORWARDING      NONE

HCDP-IESN  Chapter 2 STP and SEP 
 
Page94 HUAWEI TECHNOLOGIES HC Series 
 
   2    Ethernet0/0/13              ROOT  FORWARDING      NONE 
 
