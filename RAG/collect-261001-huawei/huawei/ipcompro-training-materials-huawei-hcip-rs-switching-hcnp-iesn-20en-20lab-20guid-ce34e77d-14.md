---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-14
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [2944, 3124]
sha256: cd39423050848ee119022dce37980da4e415a5c86a0404be962941ae99ce0eab
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

 PortTimes           :Hello 2s MaxAge 20s FwDly 15s RemHop 20 
 TC or TCN send      :37 
 TC or TCN received  :17 
 BPDU Sent           :181              
          TCN: 0, Config: 181, RST: 0, MST: 0 
 BPDU Received       :172              
          TCN: 0, Config: 172, RST: 0, MST: 0 
 
The Ethernet 0/0/1 interface of S3 is an alternate interface and that of S4 is 
a designated interface. Change the path cost to 2000000 for the E0/0/24 
interface of S4. 
[S4-Ethernet0/0/24]stp cost 2000000 
 
View role information about interfaces. 
[S3]display stp interface Ethernet 0/0/1 
----[CIST][Port1(Ethernet0/0/1)][FORWARDING]---- 
 Port Protocol       :Enabled 
 Port Role           :Designated Port 
 Port Priority       :128 
 Port Cost(Dot1T )   :Config=auto / Active=199999 
 Designated Bridge/Port   :32768.5489-98ec-f022 / 128.1 
 Port Edged          :Config=default / Active=disabled 
 Point-to-point      :Config=auto / Active=true 
 Transit Limit       :147 packets/hello-time 
 Protection Type     :None 
 Port STP Mode       :STP 
 Port Protocol Type  :Config=auto / Active=dot1s 
 PortTimes           :Hello 2s MaxAge 20s FwDly 15s RemHop 20 
 TC or TCN send      :52 
 TC or TCN received  :52 
 BPDU Sent           :284 
          TCN: 0, Config: 284, RST: 0, MST: 0 
 BPDU Received       :380 
          TCN: 0, Config: 380, RST: 0, MST: 0 
 
[S4]display stp interface Ethernet 0/0/1  
----[CIST][Port1(Ethernet0/0/1)][DISCARDING]---- 
 Port Protocol       :Enabled 
 Port Role           :Alternate Port 
 Port Priority       :128 
 Port Cost(Dot1T )   :Config=auto / Active=199999

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page67 
 
 Designated Bridge/Port   :4096.4c1f-cc45-aac1 / 128.30 
 Port Edged          :Config=default / Active=disabled 
 Point-to-point      :Config=auto / Active=true 
 Transit Limit       :147 packets/hello-time 
 Protection Type     :None 
 Port STP Mode       :STP  
 Port Protocol Type  :Config=auto / Active=dot1s 
 PortTimes           :Hello 2s MaxAge 20s FwDly 15s RemHop 0 
 TC or TCN send      :7 
 TC or TCN received  :162 
 BPDU Sent           :8              
          TCN: 7, Config: 1, RST: 0, MST: 0 
 BPDU Received       :1891              
          TCN: 0, Config: 1891, RST: 0, MST: 0 
 
The Ethernet 0/0/1 interface of S3 becomes a designated interface and 
that of S4 becomes an alternate interface. 
Step 5 Configure RSTP and verify the configuration. 
Configure IP addresses for the VLANIF interfaces of S1 and S2. Test 
whether the route from S1 to S2 is reachable. 
[S1]interface Vlanif 1 
[S1-Vlanif1]ip address 10.0.1.1 24 
 
[S2]interface Vlanif 1 
[S2-Vlanif1]ip address 10.0.1.2 24 
 
[S1]ping 10.0.1.2 
  PING 10.0.1.2: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.1.2: bytes=56 Sequence=1 ttl=255 time=9 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=2 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=3 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=4 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=5 ttl=254 time=1 ms 
 
  --- 10.0.1.2 ping statistics --- 
    5 packet(s) transmitted 
    5 packet(s) received 
    0.00% packet loss 
round-trip min/avg/max = 1/2/9 ms

HCDP-IESN  Chapter 2 STP and SEP 
 
Page68 HUAWEI TECHNOLOGIES HC Series 
 
 
View role information about interfaces of S1. 
[S1]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/9        ALTE  DISCARDING      NONE 
   0    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   0    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/14       DESI  FORWARDING      NONE 
 
The GigabitEthernet 0/0/10 interface of  S1 is the root interface. Run the 
ping command to test whether the route from S1 to S2 is reachable for 20 
times. 
Note: Disable the GigabitEthernet 0/0/10 interface of S2 immediately after 
running the ping command on S1. 
 [S1]ping -c 20 10.0.1.2 
  PING 10.0.1.2: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.1.2: bytes=56 Sequence=1 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=2 ttl=254 time=1 ms 
Dec 21 2011 16:20:44-05:13 S1 %%01IFNET/4/IF_STATE(l)[5]:Interface 
GigabitEthernet0/0/10 has turned into DOWN state. 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Reply from 10.0.1.2: bytes=56 Sequence=18 ttl=255 time=15 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=19 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=20 ttl=254 time=1 ms 
 
  --- 10.0.1.2 ping statistics --- 
    20 packet(s) transmitted

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page69 
 
    5 packet(s) received 
    75.00% packet loss 
    round-trip min/avg/max = 1/3/15 ms 
 
[S2]interface GigabitEthernet 0/0/10 
[S2-GigabitEthernet0/0/10]shutdown 
 
View role information about interfaces of S1. 
[S1]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/9        ROOT  FORWARDING      NONE 
   0    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/14       DESI  FORWARDING      NONE 
 
The GigabitEthernet 0/0/9 interface of  S1 becomes the root interface and 
enters the FORWARDING state. There are 15 expired packets and the 
network convergence takes 30 seconds. 
Enable the GigabitEthernet 0/0/10 interface of S2. 
[S2]interface GigabitEthernet 0/0/10 
[S2-GigabitEthernet0/0/10]undo shutdown 
 
Configure RSTP. 
[S1]stp mode rstp 
 
[S2]stp mode rstp 
 
[S3]stp mode rstp 
 
[S4]stp mode rstp 
 
View role information about interfaces of S1. 
[S1]display stp brief 
 MSTID      Port                      Role  STP State     Protection 
   0    GigabitEthernet0/0/9        ALTE  DISCARDING      NONE 
   0    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   0    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/14       DESI  FORWARDING      NONE 
 
The GigabitEthernet 0/0/10 interface of S1 becomes the root interface. 
Run the ping command to test whether the route from S1 to S2 is reachable

HCDP-IESN  Chapter 2 STP and SEP 
 
Page70 HUAWEI TECHNOLOGIES HC Series 
 
