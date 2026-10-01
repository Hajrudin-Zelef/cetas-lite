---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-2
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [290, 463]
sha256: 5840659821aed61fccbb3e8991fce42152041513166870212ea7420a818001e5
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

Step 2 Configure link aggregation with Eth-trunk. 
The Eth-trunk technology bundles two or more links into a single link to 
increase the bandwidth and improve reliability. Add the G0/0/9 and G0/0/10 
interfaces of S1 and S2 to an Eth-trunk. 
Create an Eth-trunk. 
[S1]interface Eth-Trunk 1 
[S1-Eth-Trunk1] 
 
[S2]interface Eth-Trunk 1 
[S2-Eth-Trunk1] 
 
Configure the Eth-trunk to work in static Link Aggregation Control Protocol 
(LACP) mode. 
[S1-Eth-Trunk1]bpdu enable 
[S1-Eth-Trunk1]mode lacp-static 
 
[S2-Eth-Trunk1]bpdu enable 
[S2-Eth-Trunk1]mode lacp-static 
 
Add the G0/0/9 and G0/0/10 interfaces of S1 and S2 to the Eth-trunk. 
[S1]interface GigabitEthernet 0/0/9 
[S1-GigabitEthernet0/0/9]eth-trunk 1 
[S1-GigabitEthernet0/0/9]quit 
[S1]interface GigabitEthernet 0/0/10 
[S1-GigabitEthernet0/0/10]eth-trunk 1 
 
[S2]interface GigabitEthernet 0/0/9 
[S2-GigabitEthernet0/0/9]eth-trunk 1 
[S2-GigabitEthernet0/0/9]quit 
[S2]interface GigabitEthernet 0/0/10 
[S2-GigabitEthernet0/0/10]eth-trunk 1 
 
Run the display eth-trunk command to view the configuration result. 
[S1]display eth-trunk 1 
Eth-Trunk1's state information is: 
Local: 
LAG ID: 1                   WorkingMode: STATIC

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page4 HUAWEI TECHNOLOGIES HC Series 
 
Preempt Delay: Disabled     Hash arithmetic: According to SA-XOR-DA            
System Priority: 32768      System ID: 4c1f-cc45-aace                          
Least Active-linknumber: 1  Max Active-linknumber: 8                           
Operate status: up        Number Of Up Port In Trunk: 2                     
---------------------------------------------------------------------------- 
ActorPortName          Status   PortType PortPri PortNo PortKey PortState Weight 
GigabitEthernet0/0/9   Selected 1GE      32768   9      305     10100010  1      
GigabitEthernet0/0/10  Selected 1GE      32768   10     305     10100010  1      
 
Partner: 
---------------------------------------------------------------------------- 
ActorPortName          SysPri    SystemID  PortPri PortNo  PortKey   PortState 
GigabitEthernet0/0/9   32768      4c1f-cc45-aace  32768   9      305     10100010 
GigabitEthernet0/0/10  32768      4c1f-cc45-aace 32768   10     305     10100010 
 
The preceding information shows that the Eth-trunk link works in static 
LACP mode and the maximum number of active interfaces is 8. In addition, the 
G0/0/9 and G0/0/10 interfaces are active. 
Change the maximum number of active interfaces. 
[S1-Eth-Trunk1]max active-linknumber 1 
 
[S2-Eth-Trunk1]max active-linknumber 1 
 
View the configurations of the Eth-trunk link. 
[S1]display eth-trunk 1 
Eth-Trunk1's state information is: 
Local: 
LAG ID: 1                   WorkingMode: STATIC 
Preempt Delay: Disabled     Hash arithmetic: According to SA-XOR-DA 
System Priority: 32768      System ID: 4c1f-cc45-aace 
Least Active-linknumber: 1  Max Active-linknumber: 1 
Operate status: up          Number Of Up Port In Trunk: 1 
---------------------------------------------------------------------------- 
ActorPortName          Status   PortType PortPri PortNo PortKey PortState Weight 
GigabitEthernet0/0/9   Selected 1GE      32768   9      305     10111100  1 
GigabitEthernet0/0/10  Unselect 1GE      32768   10     305     10100000  1 
 
Partner: 
---------------------------------------------------------------------------- 
ActorPortName          SysPri    SystemID  PortPri PortNo  PortKey   PortState

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page5 
 
GigabitEthernet0/0/9   32768  4c1f-cc45-aacc  32768  9     305       10111100 
GigabitEthernet0/0/10  32768  4c1f-cc45-aacc  32768  10    305       10100000 
 
The preceding information shows that the G0/0/10 interface changes to the 
Unselect state. In this way, one link in the Eth-trunk link group transmits data 
and the other is a backup, improving network reliability. 
Disable the G0/0/9 interface of S1 to verify the link backup function. 
[S1]interface GigabitEthernet 0/0/9 
[S1-GigabitEthernet0/0/9]shutdown 
 
View the information about the Eth-trunk link. 
[S1]display eth-trunk 1 
Eth-Trunk1's state information is: 
Local: 
LAG ID: 1                   WorkingMode: STATIC 
Preempt Delay: Disabled     Hash arithmetic: According to SA-XOR-DA 
System Priority: 32768      System ID: 4c1f-cc45-aace 
Least Active-linknumber: 1  Max Active-linknumber: 1 
Operate status: up          Number Of Up Port In Trunk: 1 
---------------------------------------------------------------------------- 
ActorPortName          Status   PortType PortPri PortNo PortKey PortState Weight 
GigabitEthernet0/0/9   Unselect 1GE      32768   9      305     10100010  1 
GigabitEthernet0/0/10  Selected 1GE      32768   10     305     10111100  1 
 
Partner: 
---------------------------------------------------------------------------- 
ActorPortName          SysPri    SystemID  PortPri PortNo  PortKey   PortState 
GigabitEthernet0/0/9   0      0000-0000-0000  0      0     0         10100011 
GigabitEthernet0/0/10  32768  4c1f-cc45-aacc  32768  10    305       10111100 
 
The preceding information shows that the G0/0/9 interface in the Eth-trunk 
link changes to the Unselect state, and G0/0/10 changes from the Unselect 
state to the Selected state to transmit data. This proves the link backup 
function. 
Step 3 Configure VLANs. 
Create two VLANs on S1 and S2 to isolate R1 from R2 and R3. Configure 
the Eth-trunk link between S1 and S2 to work in Trunk mode so that the routers 
in the two VLANs can communicate with each other.

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page6 HUAWEI TECHNOLOGIES HC Series 
 
Create VLAN 10 and VLAN 20 on S1 and S2. 
[S1]vlan batch 10 20 
 
[S2]vlan batch 10 20 
 
Configure the G0/0/1 and G0/0/2 interfaces of S1 to work in Access mode, 
and add the two interfaces to VLAN 10. 
[S1]interface GigabitEthernet 0/0/1 
[S1-GigabitEthernet0/0/1]port link-type access 
[S1-GigabitEthernet0/0/1]port default vlan 10 
[S1]interface GigabitEthernet 0/0/2 
[S1-GigabitEthernet0/0/2]port link-type access 
[S1-GigabitEthernet0/0/2]port default vlan 20 
 
Configure the G0/0/3 interface of S2 to work in Access mode, and add this 
interface to VLAN 20. 
[S2]interface GigabitEthernet 0/0/3 
[S2-GigabitEthernet0/0/3]port link-type access 
[S2-GigabitEthernet0/0/3]port default vlan 20 
 
Configure the Eth-trunk link between S1 and S2 to work in Trunk mode. By 
default, the trunk interface allows only packets from VLAN 1 to pass through. 
Therefore, configure the trunk interf ace to allow packets from VLAN 10 and 
VLAN 20 to pass through. 
[S1]interface Eth-Trunk 1 
[S1-Eth-Trunk1]port link-type trunk 
[S1-Eth-Trunk1]port trunk allow-pass vlan 10 20 
 
[S2]inter Eth-Trunk 1 
[S2-Eth-Trunk1]port link-type trunk 
[S2-Eth-Trunk1]port trunk allow-pass vlan 10 20 
 
Run the display port vlan  command to view the status of the trunk 
interface. 
[S2]display port vlan  
Port                    Link Type    PVID  Trunk VLAN List 
---------------------------------------------------------------------------- 
Eth-Trunk1              trunk        1     1 10 20 
GigabitEthernet0/0/1    hybrid       1     -

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page7 
 
