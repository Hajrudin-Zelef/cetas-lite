---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-9
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [1741, 1946]
sha256: 11e70842a300392a532f744544b854a7c5381b0126b66125004e3c542ebabf37
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page37 
 
[S2-GigabitEthernet0/0/4]shutdown 
 
Add the G0/0/9 interfaces of S1 and S2 to VLAN 3. 
[S2]interface GigabitEthernet 0/0/9 
[S2-GigabitEthernet0/0/9]undo shutdown 
[S2-GigabitEthernet0/0/9]port link-type access 
[S2-GigabitEthernet0/0/9]port default vlan 3 
 
[S1]interface GigabitEthernet 0/0/9 
[S1-GigabitEthernet0/0/9]undo shutdown 
[S1-GigabitEthernet0/0/9]port link-type access 
[S1-GigabitEthernet0/0/9]port default vlan 3 
 
Configure the G0/0/4 interface of S1 to work in Trunk mode and allow the 
access from VLAN 2 and VLAN 3. 
[S1]interface GigabitEthernet 0/0/4 
[S1-GigabitEthernet0/0/4]port default vlan 1 
[S1-GigabitEthernet0/0/4]port link-type trunk 
[S1-GigabitEthernet0/0/4]port trunk allow-pass vlan 2 3 
 
Create two subinterfaces for the G0/0 /1 interface of R4. Configure IP 
addresses for the two subinterfaces and configure them as Dot1q termination 
Ethernet subinterfaces. 
[R4]interface GigabitEthernet 0/0/1.2 
[R4-GigabitEthernet0/0/1.2]control-vid 20 dot1q-termination 
[R4-GigabitEthernet0/0/1.2]dot1q termination vid 2 
[R4-GigabitEthernet0/0/1.2]arp broadcast enable 
[R4-GigabitEthernet0/0/1.2]ip address 10.0.20.1 24 
[R4-GigabitEthernet0/0/1.2]interface GigabitEthernet 0/0/1.3 
[R4-GigabitEthernet0/0/1.3]control-vid 30 dot1q-termination 
[R4-GigabitEthernet0/0/1.3]dot1q termination vid 3 
[R4-GigabitEthernet0/0/1.3]arp broadcast enable 
[R4-GigabitEthernet0/0/1.3]ip address 10.0.30.1 24 
Run the display ip interface brief command to view subinterface 
configurations on R4. 
[R4]display ip interface brief 
*down: administratively down 
(l): loopback 
(s): spoofing 
 
The number of interface that is UP in Physical is 6

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page38 HUAWEI TECHNOLOGIES HC Series 
 
The number of interface that is DOWN in Physical is 4 
The number of interface that is UP in Protocol is 3 
The number of interface that is DOWN in Protocol is 7 
 
Interface                         IP Address/Mask      Physical   Protocol 
Cellular0/0/0                     unassigned           down       down 
Cellular0/0/1                     unassigned           down       down 
Ethernet2/0/0                     unassigned           down       down 
Ethernet2/0/1                     10.0.3.1/24          down       down 
GigabitEthernet0/0/0              unassigned           up         down 
GigabitEthernet0/0/1              10.0.2.1/24          up         up 
GigabitEthernet0/0/1.2            10.0.20.1/24        up         down 
GigabitEthernet0/0/1.3            10.0.30.1/24        up         down 
NULL0                             unassigned           up         up(s) 
Serial1/0/0                       unassigned           up         up 
 
Change the IP addresses and gateways of R1 and R2. 
[R1]inter GigabitEthernet 0/0/1 
[R1-GigabitEthernet0/0/1]ip address 10.0.20.2 24 
[R1-GigabitEthernet0/0/1]quit 
[R1]undo ip route-static 0.0.0.0 0 10.0.2.1 
[R1]ip route-static 0.0.0.0 0 10.0.20.1 
 
[R2]interface GigabitEthernet 0/0/2 
[R2-GigabitEthernet0/0/2]ip address 10.0.30.2 24 
[R2-GigabitEthernet0/0/2]quit 
[R2]undo ip route-static 0.0.0.0 0 10.0.3.1 
[R2]ip route-static 0.0.0.0 0 10.0.30.1 
 
Test whether the route from R1 to R2 is reachable. 
 
[R1]ping -c 1 10.0.30.2 
  PING 10.0.30.2: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.30.2: bytes=56 Sequence=1 ttl=254 time=3 ms 
 
  --- 10.0.30.2 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
round-trip min/avg/max = 3/3/3 ms

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page39 
 
The preceding information shows that  computers in VLAN 2 and VLAN 3 
can communicate with each other. 
Compared with the multi-arm scheme, this scheme reduces the costs in 
purchasing router interfaces. 
In single-arm routing, all data is transmitted over a single interface. W hen 
the number of VLANs increases, the bandwidth of this link may be insufficient. 
Moreover, this link is likely to caus e a single-point failure and then the entire 
network fails. 
Step 4 Configure Layer 3 switching. 
In Layer 3 switching, no router is required and each VLAN has a VLANIF 
interface for communicating with other VLANs. 
Disable the G0/0/4 interface of S1. 
[S1]interface GigabitEthernet 0/0/4 
[S1-GigabitEthernet0/0/4]shutdown 
 
Configure the G0/0/9 interfaces of S1 and S2 to work in Trunk mode and 
allow the access from VLAN 2 and VLAN 3. 
[S1]interface GigabitEthernet 0/0/9 
[S1-GigabitEthernet0/0/9]port default vlan 1 
[S1-GigabitEthernet0/0/9]port link-type trunk 
[S1-GigabitEthernet0/0/9]port trunk allow-pass vlan 2 3 
 
[S2]interface GigabitEthernet 0/0/9 
[S2-GigabitEthernet0/0/9]port default vlan 1 
[S2-GigabitEthernet0/0/9]port link-type trunk 
[S2-GigabitEthernet0/0/9]port trunk allow-pass vlan 2 3 
 
Create the VLANIF 2 and VLANIF 3 interfaces on S1 and configure IP 
addresses for the interfaces. 
[S1]interface Vlanif 2 
[S1-Vlanif2]ip address 10.0.20.1 24 
[S1]interface Vlanif 3 
[S1-Vlanif3]ip address 10.0.30.1 24 
 
Test whether the route from R1 to R2 is reachable. 
[R1]ping -c 1 10.0.30.2 
  PING 10.0.30.2: 56  data bytes, press CTRL_C to break

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page40 HUAWEI TECHNOLOGIES HC Series 
 
    Reply from 10.0.30.2: bytes=56 Sequence=1 ttl=254 time=2 ms 
 
  --- 10.0.30.2 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
    round-trip min/avg/max = 2/2/2 ms 
 
The preceding information shows that Layer 3 communication between 
computers in VLAN 2 and those in VLAN 3 is implemented using the two 
VLANIF interfaces of S1. 
Layer 3 switching delivers higher scalability than single-arm routing and 
can withstand the addition of VLANs. 
Layer 3 switching can effectively se rve a network where most traffic is 
generated by inter-VLAN communication. 
Step 5 Configure VLAN aggregation. 
Similar to Layer 3 switching, VLAN aggregation implements inter-VLAN 
communication on a switch. In VLAN aggr egation, all VLANs are on the same 
network segment to reduce the number of required IP network segments and 
unify gateway configuration. 
Create VLAN 10, VLAN 20, and VLAN 100 on S1 and S2. 
[S1]vlan batch 10 20 100 
Info: This operation may take a few seconds. Please wait for a moment...done. 
 
[S2]vlan batch 10 20 100 
Info: This operation may take a few seconds. Please wait for a moment...done. 
 
Configure the G0/0/9 interfaces of S1 and S2 to allow the access from 
VLAN 10 and VLAN 20. 
[S1]interface GigabitEthernet 0/0/9 
[S1-GigabitEthernet0/0/9]port trunk allow-pass vlan 10 20 
 
[S2]interface GigabitEthernet 0/0/9 
[S2-GigabitEthernet0/0/9]port trunk allow-pass vlan 10 20 
 
Add the G0/0/1 interface of S1 to VL AN 10 and the G0/0/2 interface of S2

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page41 
 
to VLAN 20. 
[S1]interface GigabitEthernet 0/0/1 
[S1-GigabitEthernet0/0/1]port default vlan 10 
 
[S2]interface GigabitEthernet 0/0/2 
[S2-GigabitEthernet0/0/1]port default vlan 20 
 
Configure VLAN 100 as a Super-VLAN, and add VLAN 10 and VLAN 20 to 
the Super-VLAN. Then VLAN 10 and VLAN 20 are Sub-VLANs of this 
Super-VLAN. 
[S1]vlan 100 
[S1-vlan100]aggregate-vlan 
[S1-vlan100]access-vlan 10 20 
 
Configure a VLANIF interface for VLAN 100 and enable the ARP proxy 
function. 
[S1]interface Vlanif 100 
[S1-Vlanif100]ip address 10.0.100.1 24 
[S1-Vlanif100]arp-proxy inter-sub-vlan-proxy enable 
 
Change the IP addresses of R1 and R2 to  ensure that they are on the 
same network segment as the VLANIF 100 interface. Configure the IP address 
of the VLANIF 100 interface as the gateway IP address. 
[R1]interface GigabitEthernet 0/0/1 
[R1-GigabitEthernet0/0/1]ip address 10.0.100.2 24 
[R1-GigabitEthernet0/0/1]quit 
[R1]ip route-static 0.0.0.0 0 10.0.100.1 
 
