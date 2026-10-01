---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-20
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["preemption"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [4153, 4327]
sha256: 36533b6c25a1f0f5d65af5177f7790da6862e4fe31f37972727eaabb361ce54b
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

[S4]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   0    Ethernet0/0/24              ROOT  FORWARDING      NONE 
   1    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   1    Ethernet0/0/24              ROOT  FORWARDING      NONE 
   2    Ethernet0/0/1               ALTE  DISCARDING      NONE 
 
Except instance 0, the MSTP instances in other regions calculate spanning 
trees independently, regardless of whether the instances contain or map the 
same VLAN. That is, the spanning tree processes in the same region run 
independent of each other. 
Step 4 Configure MSTP and STP to be compatible. 
Configure S1, S2, and S3 to belong to the same MSTP region. Enable 
STP on S4. 
Delete the MSTP configurations from S2 and then create RG1 on S2. 
Create instance 1 and map it to VLANs 3, 4, and 5. Create instance 2 and map 
it to VLANs 6, 7, and 8. Then activate the region configurations. 
[S2]undo stp region-configuration 
[S2]stp region-configuration 
[S2-mst-region]region-name RG1 
[S2-mst-region]revision-level 1 
[S2-mst-region]instance 1 vlan 3 4 5 
[S2-mst-region]instance 2 vlan  6 7 8 
[S2-mst-region]active region-configuration 
 
Enable the S0/0/23 interfaces of S2 and S3. 
Configure the direct link between S2 and S3 to work in Trunk mode, 
receive bridge protocol data unit (BPDU) packets, and allow the access from 
all VLANs. 
[S2]int GigabitEthernet 0/0/23 
[S2-GigabitEthernet0/0/23]undo shutdown 
[S2-GigabitEthernet0/0/23]port link-type trunk 
[S2-GigabitEthernet0/0/23]port trunk all vlan all 
[S2-GigabitEthernet0/0/23]bpdu enable

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page95 
 
[S3]int Ethernet0/0/23 
[S3-Ethernet0/0/23]undo shutdown 
[S3-Ethernet0/0/23]port link-type trunk 
[S3-Ethernet0/0/23]port trunk allow-pass vlan all 
[S3-Ethernet0/0/23]bpdu enable 
 
Delete the MSTP configurations from S4 and enable STP on S4. 
[S4]undo stp region-configuration 
[S4]stp mode stp 
 
View STP information. 
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
   0    GigabitEthernet0/0/23       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   1    GigabitEthernet0/0/23       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   2    GigabitEthernet0/0/23       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
 
[S3]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               DESI  FORWARDING      NONE 
   0    Ethernet0/0/13              ROOT  FORWARDING      NONE 
   0    Ethernet0/0/23              ALTE  DISCARDING      NONE 
   1    Ethernet0/0/1               DESI  FORWARDING      NONE 
   1    Ethernet0/0/13              ROOT  FORWARDING      NONE 
   1    Ethernet0/0/23              ALTE  DISCARDING      NONE 
   2    Ethernet0/0/1               DESI  FORWARDING      NONE 
   2    Ethernet0/0/13              ROOT  FORWARDING      NONE

HCDP-IESN  Chapter 2 STP and SEP 
 
Page96 HUAWEI TECHNOLOGIES HC Series 
 
   2    Ethernet0/0/23              ALTE  DISCARDING      NONE 
 
[S4]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               ROOT  FORWARDING      NONE 
   0    Ethernet0/0/24              ALTE  DISCARDING      NONE 
 
STP-enabled S4 and instance 0 in MSTP-enabled S1, S2, and S3 
calculate the CIST together. S1 is the CIST root. 
Set the priority to 4096 for S4 so that it becomes the CIST root. 
[S4]stp priority 4096 
 
View STP information. 
[S1]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   0    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/10       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/10       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
 
[S2]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/10       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/23       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/24       ROOT  FORWARDING      NONE 
   1    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   1    GigabitEthernet0/0/23       DESI  FORWARDING      NONE 
   1    GigabitEthernet0/0/24       MAST  FORWARDING      NONE 
   2    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   2    GigabitEthernet0/0/23       DESI  FORWARDING      NONE 
   2    GigabitEthernet0/0/24       MAST  FORWARDING      NONE 
 
[S3]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   0    Ethernet0/0/13              ALTE  DISCARDING      NONE 
   0    Ethernet0/0/23              ROOT  FORWARDING      NONE 
   1    Ethernet0/0/1               ALTE  DISCARDING      NONE

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page97 
 
   1    Ethernet0/0/13              ROOT  FORWARDING      NONE 
   1    Ethernet0/0/23              ALTE  DISCARDING      NONE 
   2    Ethernet0/0/1               ALTE  DISCARDING      NONE 
   2    Ethernet0/0/13              ROOT  FORWARDING      NONE 
   2    Ethernet0/0/23              ALTE  DISCARDING      NONE 
 
[S4]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               DESI  FORWARDING      NONE 
   0    Ethernet0/0/24              DESI  FORWARDING      NONE 
 
S4 becomes the CIST root and all interfaces on S4 become designated 
interfaces. 
Step 5 Configure designated interface protection. 
Configure designated interface protection for the E0/0/1 and E0/0/24 
interfaces of S4. 
[S4]int Ethernet0/0/1 
[S4-Ethernet0/0/1]stp root-protection 
[S4-Ethernet0/0/1]int Ethernet0/0/24 
[S4-Ethernet0/0/24]stp root-protection 
 
View STP information about S4. 
[S4]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               DESI  FORWARDING      ROOT 
   0    Ethernet0/0/24              DESI  DISCARDING      ROOT 
 
Set the priority to 0 for S2 in instance 0 to simulate preemption of the CIST 
root. 
[S2]stp instance 0 priority 0 
 
View STP information about S2 and S4. 
[S2]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/10       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/23       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/24       DESI  LEARNING        NONE 
   1    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE

HCDP-IESN  Chapter 2 STP and SEP 
 
Page98 HUAWEI TECHNOLOGIES HC Series 
 
