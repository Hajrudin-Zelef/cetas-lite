---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-21
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [4328, 4591]
sha256: c7f778b813946871aff16a230ca03ff3bd5827e44eea6ae7ad0eccb44e8bcef9
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

   1    GigabitEthernet0/0/23       DESI  LEARNING        NONE 
   1    GigabitEthernet0/0/24       DESI  LEARNING        NONE 
   2    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   2    GigabitEthernet0/0/23       DESI  LEARNING        NONE 
   2    GigabitEthernet0/0/24       DESI  LEARNING        NONE 
 
[S4]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               DESI  DISCARDING      ROOT 
   0    Ethernet0/0/24              DESI  DISCARDING      ROOT 
 
The interfaces of S4 enter the DISCARDING state and do not forward 
packets. This indicates that the roles of these interfaces do not change and S4 
is still the root switch. 
Delete the priority configurations of instance 0 from S2. 
[S2]undo stp instance 0 priority 
 
View STP information about S2 and S4. 
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
 
[S4]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/1               DESI  FORWARDING      ROOT 
   0    Ethernet0/0/24              DESI  FORWARDING      ROOT 
 
The interfaces restore to the original state (FORWARDING) if they do not 
receive configuration messages with higher priorities within the period 
specified by Max Age. The default period is 20 seconds.

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page99 
 
Step 6 Configure edge interface protection. 
Enable the G0/0/9 interface of S2. 
[S2]interface GigabitEthernet 0/0/9 
[S2-GigabitEthernet0/0/9]undo shutdown 
 
Configure the G0/0/9 interface of S1 as an edge interface and enable edge 
interface protection in global mode. 
[S1]interface GigabitEthernet 0/0/9 
[S1-GigabitEthernet0/0/9]undo shutdown 
[S1-GigabitEthernet0/0/9]stp edged-port enable 
[S1-GigabitEthernet0/0/9]quit 
[S1]stp bpdu-protection 
 
View STP information about S1. 
[S1]display stp interface GigabitEthernet 0/0/9 brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/9        DESI  FORWARDING      BPDU 
 
Enable the G0/0/9 interface of S1 so that it can receive BPDU packets and 
simulate an attack to S1. 
[S1]interface GigabitEthernet 0/0/9 
[S1-GigabitEthernet0/0/9]undo shutdown 
 
Observe S1. 
Dec 21 2011 08:39:51-05:13 S1 %%01IFNET/4/IF_STATE(l)[3]:Interface 
GigabitEthernet0/0/9 has turned into UP state. 
Dec 21 2011 08:39:51-05:13 S1 %%01MSTP/4/BPDU_PROTECTION(l)[4]:This edged-port 
GigabitEthernet0/0/9 that enabled BPDU-Protection will be shutdown, because it 
received BPDU packet! 
Dec 21 2011 08:39:52-05:13 S1 %%01IFNET/4/IF_STATE(l)[5]:Interface 
GigabitEthernet0/0/9 has turned into DOWN state. 
 
After edge interface protection is configured, the edge interface receives 
BPDU packets once being enabled, and then this interface is disabled 
automatically.

HCDP-IESN  Chapter 2 STP and SEP 
 
Page100 HUAWEI TECHNOLOGIES HC Series 
 
Step 7 Configure loop protection. 
Configure loop protection on the E0/0/23 interface of S3. 
[S3]interface Ethernet0/0/23 
[S3-Ethernet0/0/23]stp loop-protection 
 
View STP information about the E0/0/23 interface of S3. 
[S3]display stp interface Ethernet 0/0/23 brief 
 MSTID  Port                        Role  STP State     Protection 
   0    Ethernet0/0/23              ROOT  FORWARDING      LOOP 
   1    Ethernet0/0/23              ALTE  DISCARDING      LOOP 
   2    Ethernet0/0/23              ALTE  DISCARDING      LOOP 
Step 8 Configure TC-BPDU protection. 
Enable the protection function for TC-BPDU packets. 
[S1]stp tc-protection 
 
Additional Exercises: Analyzing and Verifying 
Figure out the impact if the MSTP region names are the same but the 
revision levels are different on switches. 
In step 4, figure out the status changes  of the interfaces of the four 
switches if the priority is changed to 0 for S3 in instance 1. 
Final Configurations 
<S1>display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S1 
# 
 vlan batch 3 to 8 
# 
 stp bpdu-protection 
 stp tc-protection 
#

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     
Page101 
 
 stp region-configuration 
  region-name RG1 
  revision-level 1 
  instance 1 vlan 3 to 5 
  instance 2 vlan 6 to 8 
  active region-configuration 
# 
interface GigabitEthernet0/0/9 
 shutdown 
stp edged-port enable 
# 
interface GigabitEthernet0/0/10 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 4094 
# 
interface GigabitEthernet0/0/13 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 4094 
# 
return 
<S2>display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S2 
# 
 vlan batch 3 to 8 
# 
 stp region-configuration 
  region-name RG1 
  revision-level 1 
  instance 1 vlan 3 to 5 
  instance 2 vlan 6 to 8 
  active region-configuration 
# 
interface GigabitEthernet0/0/9 
# 
interface GigabitEthernet0/0/10 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 4094 
# 
interface GigabitEthernet0/0/23

HCDP-IESN  Chapter 2 STP and SEP 
 
Page102 HUAWEI TECHNOLOGIES HC Series 
 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 4094 
# 
interface GigabitEthernet0/0/24 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 4094 
# 
return 
<S3>display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S3 
# 
 vlan batch 3 to 8 
# 
 stp region-configuration 
  region-name RG1 
  revision-level 1 
  instance 1 vlan 3 to 5 
  instance 2 vlan 6 to 8 
  active region-configuration 
# 
interface Ethernet0/0/1 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 4094 
# 
interface Ethernet0/0/13 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 4094 
# 
interface Ethernet0/0/23 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 4094 
 stp loop-protection 
# 
return 
<S4>display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S4 
#

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     
Page103 
 
 vlan batch 3 to 8 30 
# 
 stp mode stp 
 stp instance 0 priority 4096 
# 
interface Vlanif30 
 ip address 100.100.100.8 255.255.255.0 
# 
interface Ethernet0/0/1 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 4094 
 stp root-protection 
 undo ntdp enable 
 undo ndp enable 
# 
interface Ethernet0/0/14 
 shutdown 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
interface Ethernet0/0/23 
 port link-type access 
 port default vlan 30 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
interface Ethernet0/0/24 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 4094 
 stp root-protection 
 undo ntdp enable 
 undo ndp enable 
# 
Return

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
Page104 HUAWEI TECHNOLOGIES HC Series 
 
Chapter 3 Implementing MPLS technologies 
Lab 3-1 MPLS LDP Configuration 
Learning Objectives 
The objectives of this lab are to learn and understand: 
/g120 Methods used to enable and disable MPLS 
/g120 MPLS LDP Configuration 
/g120 Methods used to configure LSP sessions using MPLS LDP 
/g120 Methods used to modify the LDP LSP trigger policy on each LSR 
Topology 
 
