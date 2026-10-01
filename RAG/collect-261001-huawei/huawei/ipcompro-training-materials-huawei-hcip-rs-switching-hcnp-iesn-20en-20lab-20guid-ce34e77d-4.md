---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-4
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [615, 773]
sha256: 2d10404e6108bd18ba664baca863f636658658a3757dcb8833b77a8a45d433db
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

[S2]display port vlan 
Port                    Link Type    PVID  Trunk VLAN List 
---------------------------------------------------------------------------- 
Eth-Trunk1                trunk        1     1 10 20 
GigabitEthernet0/0/1    hybrid       1     -                                    
GigabitEthernet0/0/2    hybrid       1     -                                    
GigabitEthernet0/0/3    hybrid       20   20                                    
GigabitEthernet0/0/4    hybrid       1     -                                    
GigabitEthernet0/0/5    hybrid       1     -                                    
GigabitEthernet0/0/6    hybrid       1     -                                    
GigabitEthernet0/0/7    hybrid       1     -                                    
GigabitEthernet0/0/8    hybrid       1     -                                    
GigabitEthernet0/0/9    hybrid       0     -                                    
GigabitEthernet0/0/10   hybrid       0     -                                    
GigabitEthernet0/0/11   hybrid       1     -                                    
GigabitEthernet0/0/12   hybrid       1     -                                    
GigabitEthernet0/0/13   hybrid       1     -                                    
GigabitEthernet0/0/14   hybrid       1     -                                    
GigabitEthernet0/0/15   hybrid       1     -                                    
GigabitEthernet0/0/16   hybrid       1     -                                    
GigabitEthernet0/0/17   hybrid       1     -                                    
GigabitEthernet0/0/18   hybrid       1     -                                    
GigabitEthernet0/0/19   hybrid       1     -                                    
GigabitEthernet0/0/20   hybrid       1     -                                    
GigabitEthernet0/0/21   hybrid       1     -                                    
GigabitEthernet0/0/22   hybrid       1     -                                    
GigabitEthernet0/0/23   hybrid       1     -

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page11 
 
GigabitEthernet0/0/24   hybrid       1     -                                    
 
The preceding information shows that th e G0/0/1 interface of S1 forwards 
packets from VLAN 20 in untagged mode. The port VLAN ID (PVID) of G0/0/1 
belongs to VLAN 10. The G0/0/3 interface of S2 forwards packets from VLAN 
10 in untagged mode. The PVID of G0/0/3 belongs to VLAN 20. 
 
Run the ping command to test whether the routes from R1 to R2 and R3 
are reachable. 
[R1]ping -c 1 10.0.10.2 
  PING 10.0.10.2: 56  data bytes, press CTRL_C to break 
    Request time out 
 
  --- 10.0.10.2 ping statistics --- 
    1 packet(s) transmitted 
    0 packet(s) received 
    100.00% packet loss 
 
[R1]ping -c 1 10.0.10.3 
  PING 10.0.10.3: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.10.3: bytes=56 Sequence=1 ttl=255 time=13 ms 
 
  --- 10.0.10.3 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
round-trip min/avg/max = 13/13/13 ms 
 
Configuration 2: The link between S1 and S2 works in Hybrid mode. 
Configure the Eth-trunk interfaces on S1 and S2 to work in Hybrid mode. 
Run Untagged VLAN 10 on the Eth-trunk interface of S1 and Untagged 
VLAN 20 on the Eth-trunk interface of S2. 
[S1]interface Eth-Trunk 1 
[S1-Eth-Trunk1]undo port trunk allow-pass vlan 10 20 
[S1-Eth-Trunk1]port link-type hybrid 
[S1-Eth-Trunk1]port hybrid untagged vlan 10 
 
[S2]interface Eth-Trunk 1 
[S2-Eth-Trunk1]undo port trunk allow-pass vlan 10 20 
[S2-Eth-Trunk1]port link-type hybrid

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page12 HUAWEI TECHNOLOGIES HC Series 
 
[S2-Eth-Trunk1]port hybrid untagged vlan 20 
 
View the configurations of the Eth-trunk 1 interface. 
[S1]display port vlan  
Port                    Link Type    PVID  Trunk VLAN List 
---------------------------------------------------------------------------- 
Eth-Trunk1              hybrid       1      - 
GigabitEthernet0/0/1    hybrid       10   - 
GigabitEthernet0/0/2    access       20    -                                    
GigabitEthernet0/0/3    hybrid       1     -                                    
GigabitEthernet0/0/4    hybrid       1     -                                    
GigabitEthernet0/0/5    hybrid       1     -                                    
GigabitEthernet0/0/6    hybrid       1     -                                    
GigabitEthernet0/0/7    hybrid       1     -                                    
GigabitEthernet0/0/8    hybrid       1     -                                    
GigabitEthernet0/0/9    hybrid       0     -                                    
GigabitEthernet0/0/10   hybrid       0     -                                    
GigabitEthernet0/0/11   hybrid       1     -                                    
GigabitEthernet0/0/12   hybrid       1     -                                    
GigabitEthernet0/0/13   hybrid       1     -                                    
GigabitEthernet0/0/14   hybrid       1     -                                    
GigabitEthernet0/0/15   hybrid       1     -                                    
GigabitEthernet0/0/16   hybrid       1     -                                    
GigabitEthernet0/0/17   hybrid       1     -                                    
GigabitEthernet0/0/18   hybrid       1     -                                    
GigabitEthernet0/0/19   hybrid       1     -                                    
GigabitEthernet0/0/20   hybrid       1     -                                    
GigabitEthernet0/0/21   hybrid       1     -                                    
GigabitEthernet0/0/22   hybrid       1     -                                    
GigabitEthernet0/0/23   hybrid       1     -                                    
GigabitEthernet0/0/24   hybrid       1     -                                    
 
[S1]display vlan 10 
---------------------------------------------------------------------------- 
U: Up;         D: Down;         TG: Tagged;         UT: Untagged; 
MP: Vlan-mapping;               ST: Vlan-stacking; 
#: ProtocolTransparent-vlan;    *: Management-vlan; 
---------------------------------------------------------------------------- 
VID  Type    Ports                                                           
---------------------------------------------------------------------------- 
10   common  UT:GE0/0/1(U)      Eth-Trunk1(U)

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page13 
 
VID  Status  Property      MAC-LRN Statistics Description       
---------------------------------------------------------------------------- 
10   enable  default       enable  disable    VLAN 0010 
The preceding information shows that the Eth-trunk 1 interface of S1 works 
in Hybrid mode and forwards packets from VLAN 10 in untagged mode. The 
Eth-trunk 1 interface of S2 works in Hybrid mode and forwards packets from 
VLAN 20 in untagged mode. 
Run the ping command to test whether the routes from R1 to R2 and R3 
are reachable. 
[R1]ping -c 1 10.0.10.2 
  PING 10.0.10.2: 56  data bytes, press CTRL_C to break 
    Request time out 
 
  --- 10.0.10.2 ping statistics --- 
    1 packet(s) transmitted 
    0 packet(s) received 
    100.00% packet loss 
 
[R1]ping -c 1 10.0.10.3 
  PING 10.0.10.3: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.10.3: bytes=56 Sequence=1 ttl=255 time=4 ms 
 
  --- 10.0.10.3 ping statistics --- 
    1 packet(s) transmitted 
    0 packet(s) received 
0% packet loss 
R1 can communicate with R3. 
Additional Exercises: Analyzing and Verifying 
In step 4, figure out the differences in data transmission when the 
Eth-trunk interfaces on S1 and S2 work in Trunk mode and Hybrid mode. 
Final Configurations 
[S1]display current-configuration 
# 
!Software Version V100R005C01SPC100 
 sysname S1

