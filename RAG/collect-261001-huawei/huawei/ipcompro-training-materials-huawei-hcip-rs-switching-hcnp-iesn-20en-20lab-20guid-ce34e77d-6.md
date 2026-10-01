---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-6
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [1072, 1239]
sha256: 8f00ad962886a65fafd766d66809d2f1b32295e4d61d8c6478ccd123d6c2d6bb
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

Run the display vlan command to view VLAN information about S1 in the 
current state. 
[S1]display vlan 
The total number of vlans is : 1 
---------------------------------------------------------------------------- 
U: Up;         D: Down;         TG: Tagged;         UT: Untagged; 
MP: Vlan-mapping;               ST: Vlan-stacking; 
#: ProtocolTransparent-vlan;    *: Management-vlan; 
---------------------------------------------------------------------------- 
VID  Type    Ports 
---------------------------------------------------------------------------- 
1    common  UT:GE0/0/1(U)      GE0/0/2(U)      GE0/0/3(U)      GE0/0/4(U) 
                GE0/0/5(U)      GE0/0/6(D)      GE0/0/7(D)      GE0/0/8(D) 
                GE0/0/9(U)      GE0/0/10(D)     GE0/0/11(D)     GE0/0/12(D) 
                GE0/0/13(D)     GE0/0/14(D)     GE0/0/15(D)     GE0/0/16(D) 
                GE0/0/17(D)     GE0/0/18(D)     GE0/0/19(D)     GE0/0/20(D) 
                GE0/0/21(D)     GE0/0/22(D)     GE0/0/23(D)     GE0/0/24(D) 
 
VID  Status  Property      MAC-LRN Statistics Description 
---------------------------------------------------------------------------- 
1    enable  default       enable  disable    VLAN 0001 
 
Create VLAN 10, VLAN 20, and VLAN 100 on S2.

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page22 HUAWEI TECHNOLOGIES HC Series 
 
[S2]vlan batch 10 20 100 
Info: This operation may take a few seconds. Please wait for a moment...done. 
 
View VLAN information about S1 again. 
[S1]display vlan 
The total number of vlans is : 4 
---------------------------------------------------------------------------- 
U: Up;         D: Down;         TG: Tagged;         UT: Untagged; 
MP: Vlan-mapping;               ST: Vlan-stacking; 
#: ProtocolTransparent-vlan;    *: Management-vlan; 
---------------------------------------------------------------------------- 
VID  Type    Ports 
---------------------------------------------------------------------------- 
1    common  UT:GE0/0/1(U)      GE0/0/2(U)      GE0/0/3(U)      GE0/0/4(U) 
                GE0/0/5(U)      GE0/0/6(D)      GE0/0/7(D)      GE0/0/8(D) 
                GE0/0/9(U)      GE0/0/10(D)     GE0/0/11(D)     GE0/0/12(D) 
                GE0/0/13(D)     GE0/0/14(D)     GE0/0/15(D)     GE0/0/16(D) 
                GE0/0/17(D)     GE0/0/18(D)     GE0/0/19(D)     GE0/0/20(D) 
                GE0/0/21(D)     GE0/0/22(D)     GE0/0/23(D)     GE0/0/24(D) 
10   dynamic TG:GE0/0/9(U) 
20   dynamic TG:GE0/0/9(U) 
100  dynamic TG:GE0/0/9(U) 
 
VID  Status  Property      MAC-LRN Statistics Description 
---------------------------------------------------------------------------- 
1    enable  default       enable  disable    VLAN 0001 
10   enable  default       enable  disable    VLAN 0010 
20   enable  default       enable  disable    VLAN 0020 
100  enable  default       enable  disable    VLAN 0100 
 
The preceding information shows that S1 learns the VLAN information 
about S2. 
Create VLAN 30 on S1 and observe the changes of the VLAN information 
about S1 and S2. 
[S1]vlan 30 
[S1-vlan30]quit 
[S1]display vlan 
The total number of vlans is : 5 
---------------------------------------------------------------------------- 
U: Up;         D: Down;         TG: Tagged;         UT: Untagged; 
MP: Vlan-mapping;               ST: Vlan-stacking;

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page23 
 
#: ProtocolTransparent-vlan;    *: Management-vlan; 
---------------------------------------------------------------------------- 
 
VID  Type    Ports 
---------------------------------------------------------------------------- 
1    common  UT:GE0/0/1(U)      GE0/0/2(U)      GE0/0/3(U)      GE0/0/4(U) 
                GE0/0/5(U)      GE0/0/6(D)      GE0/0/7(D)      GE0/0/8(D) 
                GE0/0/9(U)      GE0/0/10(D)     GE0/0/11(D)     GE0/0/12(D) 
                GE0/0/13(D)     GE0/0/14(D)     GE0/0/15(D)     GE0/0/16(D) 
                GE0/0/17(D)     GE0/0/18(D)     GE0/0/19(D)     GE0/0/20(D) 
                GE0/0/21(D)     GE0/0/22(D)     GE0/0/23(D)     GE0/0/24(D) 
10   dynamic TG:GE0/0/9(U) 
20   dynamic TG:GE0/0/9(U) 
30   common  TG:GE0/0/9(U) 
100  dynamic TG:GE0/0/9(U) 
 
VID  Status  Property      MAC-LRN Statistics Description 
---------------------------------------------------------------------------- 
1    enable  default       enable  disable    VLAN 0001 
10   enable  default       enable  disable    VLAN 0010 
20   enable  default       enable  disable    VLAN 0020 
30   enable  default       enable  disable    VLAN 0030 
100  enable  default       enable  disable    VLAN 0100 
 
[S2]display vlan 
The total number of vlans is : 4 
---------------------------------------------------------------------------- 
U: Up;         D: Down;         TG: Tagged;         UT: Untagged; 
MP: Vlan-mapping;               ST: Vlan-stacking; 
#: ProtocolTransparent-vlan;    *: Management-vlan; 
---------------------------------------------------------------------------- 
VID  Type    Ports 
---------------------------------------------------------------------------- 
1    common  UT:GE0/0/1(U)      GE0/0/2(U)      GE0/0/3(U)      GE0/0/4(U) 
                GE0/0/5(U)      GE0/0/6(D)      GE0/0/7(D)      GE0/0/8(D) 
                GE0/0/9(U)      GE0/0/10(D)     GE0/0/11(D)     GE0/0/12(D) 
                GE0/0/13(D)     GE0/0/14(D)     GE0/0/15(D)     GE0/0/16(D) 
                GE0/0/17(D)     GE0/0/18(D)     GE0/0/19(D)     GE0/0/20(D) 
                GE0/0/21(D)     GE0/0/22(D)     GE0/0/23(D)     GE0/0/24(D) 
10   common  TG:GE0/0/9(U) 
20   common  TG:GE0/0/9(U) 
100  common  TG:GE0/0/9(U)

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page24 HUAWEI TECHNOLOGIES HC Series 
 
VID  Status  Property      MAC-LRN Statistics Description 
---------------------------------------------------------------------------- 
1    enable  default       enable  disable    VLAN 0001 
10   enable  default       enable  disable    VLAN 0010 
20   enable  default       enable  disable    VLAN 0020 
100  enable  default       enable  disable    VLAN 0100 
 
The preceding information shows that S2 does not learn the VLAN 
information about S1. 
Step 3 Configure MUX VLANs. 
After devices on the same network segment are added to different VLANs, 
Layer 2 communication is isolated, that is, routers in different VLANs cannot 
communicate with each other. However, the MUX VLAN can allow these 
routers to communicate with a specified VLAN. In addition, the MUX VLAN can 
restrict devices in the same VLAN from communicating with each other. 
Configure VLAN 100 as the primary VLAN of the MUX VLAN, and VLAN 
10 and VLAN 20 as the secondary VLANs. 
Configure the types of the interfaces on the computers that connect to the 
switches to ensure that: all comput ers can communicate with R4; R3 and R4 
cannot communicate with each other or with routers in other VLANs. 
Configure VLAN 100 as a secondary VLAN and add related configurations. 
[S1]vlan 10 
[S1-vlan10]quit 
[S1]vlan 20 
[S1-vlan20]quit 
[S1]vlan 100 
[S1-vlan100]mux-vlan 
[S1-vlan100]subordinate group 10 
[S1-vlan100]subordinate separate 20 
 
[S2]vlan 100 
[S2-vlan100]mux-vlan 
[S2-vlan100]subordinate group 10 
[S2-vlan100]subordinate separate 20 
 
Add the G0/0/5 interface on R5 that connects to S2 to VLAN 100 and 
enable the MUX VLAN function.

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page25 
 
[S2]interface GigabitEthernet 0/0/5 
[S2-GigabitEthernet0/0/5]port link-type access 
[S2-GigabitEthernet0/0/5]port default vlan 100 
[S2-GigabitEthernet0/0/5]port mux-vlan enable 
 
