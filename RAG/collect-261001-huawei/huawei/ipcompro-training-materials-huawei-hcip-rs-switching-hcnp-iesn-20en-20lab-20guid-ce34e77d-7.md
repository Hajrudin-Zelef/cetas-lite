---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-7
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [1240, 1504]
sha256: 2f845f2c8494e3df189867014c106e7cd2712d0d2b3b7abf8c8567f228c9e31f
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

Add the G0/0/1 interface on R1 that  connects to S1 and the G0/0/2 
interface on R2 that connects to S1 to VLAN 10 and enable the MUX VLAN 
function. 
[S1]interface GigabitEthernet 0/0/1 
[S1-GigabitEthernet0/0/1]port link-type access 
[S1-GigabitEthernet0/0/1]port default vlan 10 
[S1-GigabitEthernet0/0/1]port mux-vlan enable 
[S1-GigabitEthernet0/0/1]interface GigabitEthernet 0/0/2 
[S1-GigabitEthernet0/0/2]port link-type access 
[S1-GigabitEthernet0/0/2]port default vlan 10 
[S1-GigabitEthernet0/0/2]port mux-vlan enable 
 
Add the G0/0/3 interface on R3 that  connects to S1 and the G0/0/4 
interface on R4 that connects to S2 to VLAN 20 and enable the MUX VLAN 
function. 
[S1]interface GigabitEthernet 0/0/3 
[S1-GigabitEthernet0/0/3]port link-type access 
[S1-GigabitEthernet0/0/3]port default vlan 20 
[S1-GigabitEthernet0/0/3]port mux-vlan enable 
 
[S2]interface GigabitEthernet 0/0/4  
[S2-GigabitEthernet0/0/4]port link-type access  
[S2-GigabitEthernet0/0/4]port default vlan 20  
[S2-GigabitEthernet0/0/4]port mux-vlan enable 
Run the display mux-vlan command to view information about all MUX 
VLANs. 
[S1]display mux-vlan 
Principal Subordinate Type         Interface 
---------------------------------------------------------------------------- 
100       -           principal 
100       20          separate     GigabitEthernet0/0/3 
100       10          group        GigabitEthernet0/0/1 GigabitEthernet0/0/2 
 
[S2]display mux-vlan 
Principal Subordinate Type         Interface 
----------------------------------------------------------------------------

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page26 HUAWEI TECHNOLOGIES HC Series 
 
100       -           principal    GigabitEthernet0/0/5 
100       20          separate     GigabitEthernet0/0/4 
100       10          group 
---------------------------------------------------------------------------- 
 
Run the ping command to test whether the routes from R1 to R2, R3, R4, 
and R5 are reachable. 
[R1]ping -c 1 10.0.10.2 
  PING 10.0.10.2: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.10.2: bytes=56 Sequence=1 ttl=255 time=3 ms 
 
  --- 10.0.10.2 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
    round-trip min/avg/max = 3/3/3 ms 
 
[R1]ping -c 1 10.0.10.3 
  PING 10.0.10.3: 56  data bytes, press CTRL_C to break 
    Request time out 
 
  --- 10.0.10.3 ping statistics --- 
    1 packet(s) transmitted 
    0 packet(s) received 
    100.00% packet loss 
 
[R1]ping -c 1 10.0.10.4 
  PING 10.0.10.4: 56  data bytes, press CTRL_C to break 
    Request time out 
 
  --- 10.0.10.4 ping statistics --- 
    1 packet(s) transmitted 
    0 packet(s) received 
    100.00% packet loss 
 
[R1]ping -c 1 10.0.10.5 
  PING 10.0.10.5: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.10.5: bytes=56 Sequence=1 ttl=255 time=3 ms 
 
  --- 10.0.10.5 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page27 
 
    0.00% packet loss 
    round-trip min/avg/max = 3/3/3 ms 
 
Run the ping command to test whether the routes from R3 to R2, R4, and 
R5 are reachable. 
[R3]ping -c 1 10.0.10.2 
  PING 10.0.10.2: 56  data bytes, press CTRL_C to break 
    Request time out 
 
  --- 10.0.10.2 ping statistics --- 
    1 packet(s) transmitted 
    0 packet(s) received 
    100.00% packet loss 
 
[R3]ping -c 1 10.0.10.4 
  PING 10.0.10.4: 56  data bytes, press CTRL_C to break 
    Request time out 
 
  --- 10.0.10.4 ping statistics --- 
    1 packet(s) transmitted 
    0 packet(s) received 
    100.00% packet loss 
 
[R3]ping -c 1 10.0.10.5 
  PING 10.0.10.5: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.10.5: bytes=56 Sequence=1 ttl=255 time=3 ms 
 
  --- 10.0.10.5 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
    round-trip min/avg/max = 3/3/3 ms 
 
The preceding information shows that: R1 and R2 in VLAN 10 can 
communicate with each other, as well as with R5; R3 and R4 in VLAN 10 can 
communicate only with R5. 
Step 4 Configure voice VLANs. 
To meet service development requirements, voice VLANs are configured 
on S2 and related configuration is completed on the G0/0/24 interface.

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page28 HUAWEI TECHNOLOGIES HC Series 
 
Create VLAN 30 and VLAN 200 on S2. VLAN 200 functions as a voice 
VLAN. 
[S2]vlan batch 30 200 
Info: This operation may take a few seconds. Please wait for a moment...done. 
 
Configure the interface type and default VLAN for the G0/0/24 interface of 
S2. Assume that the voice device belongs to VLAN 30. 
[S2]interface GigabitEthernet 0/0/24 
[S2-GigabitEthernet0/0/24]port hybrid pvid vlan 30 
[S2-GigabitEthernet0/0/24]port hybrid untagged vlan 30 
 
Configure an OUI address on S2. Assume  that the media access control 
(MAC) address of the voice device is 0011-2200-0000 and the subnet mask is 
ffff-ff00-0000. 
[S2]voice-vlan mac-address 0011-2200-0000 mask ffff-ff00-0000 
 
Enable the voice VLAN function and configure an automatic voice VLAN 
on the G0/0/24 interface of S2. 
[S2]interface GigabitEthernet 0/0/24 
[S2-GigabitEthernet0/0/24]voice-vlan 200 enable 
[S2-GigabitEthernet0/0/24]voice-vlan mode auto 
[S2-GigabitEthernet0/0/24]voice-vlan security enable 
 
Run the display voice-vlan oui command to view the OUI address of the 
voice VLAN. 
[S2]display voice-vlan oui 
--------------------------------------------------- 
OuiAddress       Mask             Description 
--------------------------------------------------- 
0011-2200-0000  ffff-ff00-0000 
 
Run the display voice-vlan 200 status  command to view the 
configurations of the voice VLAN. 
[S2]display voice-vlan 200 status 
Voice VLAN Configurations: 
----------------------------------------------------------- 
Voice VLAN ID            : 200 
Voice VLAN status        : Enable 
Voice VLAN aging time    : 1440(minutes)

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page29 
 
Voice VLAN 8021p remark  : 6 
Voice VLAN dscp remark   : 46 
----------------------------------------------------------- 
Port Information: 
----------------------------------------------------------- 
Port                     Add-Mode  Security-Mode  Legacy 
----------------------------------------------------------- 
GigabitEthernet0/0/24    Auto      Security       Disable 
Additional Exercises: Analyzing and Verifying 
Figure out whether computers in different MUX VLANs can communicate 
with each other. 
Final Configurations 
[S1]display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S1 
# 
 vlan batch 10 20 30 100 
# 
 gvrp 
# 
vlan 100 
 mux-vlan 
 subordinate separate 20 
 subordinate group 10 
# 
interface GigabitEthernet0/0/1 
 port link-type access 
 port default vlan 10 
 port mux-vlan enable 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
interface GigabitEthernet0/0/2 
 port link-type access 
 port default vlan 10 
 port mux-vlan enable

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page30 HUAWEI TECHNOLOGIES HC Series 
 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
interface GigabitEthernet0/0/3 
 port link-type access 
 port default vlan 20 
 port mux-vlan enable 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
interface GigabitEthernet0/0/9 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 4094 
 undo ntdp enable 
 undo ndp enable 
 gvrp 
# 
interface GigabitEthernet0/0/10 
 shutdown 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
Return 
 
[S2]display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S2 
# 
 voice-vlan mac-address 0011-2200-0000 mask ffff-ff00-0000 
# 
 vlan batch 10 20 30 100 200 
# 
 gvrp 
# 
vlan 100 
 mux-vlan 
 subordinate separate 20 
 subordinate group 10 
#

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
