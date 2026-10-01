---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-3
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [464, 614]
sha256: 5245a21c03afcf7cecb04f083f8e31adfa4c72f468bbc00b8ec371f1372f296b
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

GigabitEthernet0/0/2    hybrid       1    -                                    
GigabitEthernet0/0/3    access       20     -                                    
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
Eth-Trunk1              trunk        1     1 10 20 
 
The preceding information shows that  the Eth-trunk interface works in 
Trunk mode and allows VLAN 10 to communicate with VLAN 20. 
Run the ping command to test whether the routes from R2 to R1 and R3 are 
reachable. 
[R2]ping -c 1 10.0.10.1 
  PING 10.0.10.1: 56  data bytes, press CTRL_C to break 
    Request time out 
 
  --- 10.0.10.1 ping statistics --- 
    1 packet(s) transmitted 
    0 packet(s) received 
    100.00% packet loss 
 
[R2]ping -c 1 10.0.10.3 
  PING 10.0.10.3: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.10.3: bytes=56 Sequence=1 ttl=255 time=3 ms

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page8 HUAWEI TECHNOLOGIES HC Series 
 
 
  --- 10.0.10.3 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
round-trip min/avg/max = 3/3/3 ms 
 
The preceding information shows that computers in different VLANs on two 
switches that are interconnected cannot communicate with each 
other,computers in the same VLAN can;but computers in different VLANs on 
two switches that are interconnected ov er trunk can communicate with each 
other. 
Step 4 Configure hybrid interfaces. 
The hybrid interface allows R1 to communicate with R3 in Layer 2 mode. 
Compare the hybrid interface with the trunk interface to learn the 
configurations and features of the hybrid interface. 
Configuration 1: The link between S1 and S2 works in Trunk mode. 
Configure the G0/0/1 interface of S1 to work in Hybrid mode, and add this 
interface to VLAN 10. Configure the function of removing VLAN tags of 
packets from VLAN 20. 
[S1]interface GigabitEthernet 0/0/1 
[S1-GigabitEthernet0/0/1]port default vlan 1 
[S1-GigabitEthernet0/0/1]port link-type hybrid 
[S1-GigabitEthernet0/0/1]port hybrid pvid vlan 10 
[S1-GigabitEthernet0/0/1]port hybrid tagged vlan 10 
[S1-GigabitEthernet0/0/1]port hybrid untagged vlan  20 
 
Configure the G0/0/3 interface of S2 to work in Hybrid mode, and add this 
interface to VLAN 20. Enable the G0/0/3 interface to forward packets from 
VLAN 10 in untagged mode. 
[S2]interface GigabitEthernet 0/0/3 
[S2-GigabitEthernet0/0/3]port default vlan 1 
[S2-GigabitEthernet0/0/3]port link-type hybrid 
[S2-GigabitEthernet0/0/3]port hybrid pvid vlan 20 
[S2-GigabitEthernet0/0/3]port hybrid untagged vlan 10  
 
Run the display vlan and display port vlan commands to view interface 
status.

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page9 
 
[S1]display vlan 20 
---------------------------------------------------------------------------- 
U: Up;         D: Down;         TG: Tagged;         UT: Untagged; 
MP: Vlan-mapping;               ST: Vlan-stacking; 
#: ProtocolTransparent-vlan;    *: Management-vlan; 
---------------------------------------------------------------------------- 
VID  Type    Ports                                                           
---------------------------------------------------------------------------- 
20   common  UT:GE0/0/1(U)      GE0/0/2(U)                                       
              TG:Eth-Trunk1(U)                                                    
VID  Status  Property      MAC-LRN Statistics Description       
---------------------------------------------------------------------------- 
20   enable  default       enable  disable    VLAN 0020                   
 
[S1]display port vlan 
Port                    Link Type    PVID  Trunk VLAN List 
---------------------------------------------------------------------------- 
Eth-Trunk1                trunk        1     1 10 20 
GigabitEthernet0/0/1    hybrid       10    10                                   
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

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page10 HUAWEI TECHNOLOGIES HC Series 
 
 
[S2]display vlan 10    
---------------------------------------------------------------------------- 
U: Up;         D: Down;         TG: Tagged;         UT: Untagged; 
MP: Vlan-mapping;               ST: Vlan-stacking; 
#: ProtocolTransparent-vlan;    *: Management-vlan; 
---------------------------------------------------------------------------- 
VID  Type    Ports                                                           
---------------------------------------------------------------------------- 
10   common  UT:GE0/0/3(U)                                                       
              TG:Eth-Trunk1(U)                                                    
VID  Status  Property      MAC-LRN Statistics Description       
---------------------------------------------------------------------------- 
10   enable  default       enable  disable    VLAN 0010   
                  
