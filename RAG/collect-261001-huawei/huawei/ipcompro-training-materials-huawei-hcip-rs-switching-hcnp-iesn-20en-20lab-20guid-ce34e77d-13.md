---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-13
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [2723, 2943]
sha256: 0e2a85671b75cc308b9bf91b479fc35c9b3cf894a1f3ee93b04acb2829b39df1
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

View STP status information. 
[S1]display stp 
-------[CIST Global Info][Mode STP]------- 
CIST Bridge         :32768.4c1f-cc45-aadc 
Bridge Times        :Hello 2s MaxAge 20s FwDly 15s MaxHop 20 
CIST Root/ERPC      :32768.4c1f-cc45-aac1 / 20000 
CIST RegRoot/IRPC   :32768.4c1f-cc45-aadc / 0 
CIST RootPortId     :128.9 
BPDU-Protection     :Disabled 
TC or TCN received  :36 
TC count per hello  :2 
STP Converge Mode   :Normal 
Share region-configuration :Enabled 
Time since last TC  :0 days 0h:0m:1s 
……output omit…… 
 
[S2]display stp 
-------[CIST Global Info][Mode STP]------- 
CIST Bridge         :32768.4c1f-cc45-aac1 
Bridge Times        :Hello 2s MaxAge 20s FwDly 15s MaxHop 20 
CIST Root/ERPC      :32768.4c1f-cc45-aac1 / 0 
CIST RegRoot/IRPC   :32768.4c1f-cc45-aac1 / 0 
CIST RootPortId     :0.0 
BPDU-Protection     :Disabled 
TC or TCN received  :20 
TC count per hello  :0 
STP Converge Mode   :Normal 
Share region-configuration :Enabled 
Time since last TC  :0 days 0h:1m:4s 
……output omit…… 
 
[S1]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/9        ROOT  FORWARDING      NONE 
   0    GigabitEthernet0/0/10       ALTE  DISCARDING      NONE 
   0    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/14       DESI  FORWARDING      NONE 
 
[S2]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/9        DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/10       DESI  FORWARDING      NONE

HCDP-IESN  Chapter 2 STP and SEP 
 
Page62 HUAWEI TECHNOLOGIES HC Series 
 
   0    GigabitEthernet0/0/23       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/24       DESI  FORWARDING      NONE 
 
S2 is the root bridge and all interfaces are designated interfaces. 
The actual results of the experim ent may be different due to the 
uncertainty of the MAC addresses of switches. 
Step 2 Control root bridge election. 
Configure S1 as the primary root br idge and S2 as the secondary root 
bridge. 
[S1]stp root primary 
 
[S2]stp root secondary 
 
View STP configuration information. 
[S1]display stp 
-------[CIST Global Info][Mode STP]------- 
CIST Bridge         :0    .4c1f-cc45-aadc 
Bridge Times        :Hello 2s MaxAge 20s FwDly 15s MaxHop 20 
CIST Root/ERPC      :0    .4c1f-cc45-aadc / 0 
CIST RegRoot/IRPC   :0    .4c1f-cc45-aadc / 0 
CIST RootPortId     :0.0 
BPDU-Protection     :Disabled 
CIST Root Type      :Primary root 
TC or TCN received  :67 
TC count per hello  :0 
STP Converge Mode   :Normal 
Share region-configuration :Enabled 
Time since last TC  :0 days 0h:0m:15s  
……output omit…… 
 
[S2]display stp 
-------[CIST Global Info][Mode STP]------- 
CIST Bridge         :4096 .4c1f-cc45-aac1 
Bridge Times        :Hello 2s MaxAge 20s FwDly 15s MaxHop 20 
CIST Root/ERPC      :0    .4c1f-cc45-aadc / 20000 
CIST RegRoot/IRPC   :4096 .4c1f-cc45-aac1 / 0 
CIST RootPortId     :128.9 
BPDU-Protection     :Disabled

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page63 
 
CIST Root Type      :Secondary root 
TC or TCN received  :26 
TC count per hello  :0 
STP Converge Mode   :Normal 
Share region-configuration :Enabled 
Time since last TC  :0 days 0h:0m:1s 
……output omit…… 
 
Configure S1 as the primary root br idge and S2 as the secondary root 
bridge. 
The smaller the priority value, the higher the priority. Change the priority 
value to 8129 for S1 and to 4096 for S2. 
[S1]undo stp root 
[S1]stp priority 8192 
 
[S2]undo stp root 
[S2]stp priority 4096 
 
View STP information. 
[S1]display stp 
-------[CIST Global Info][Mode STP]------- 
CIST Bridge         :8192 .4c1f-cc45-aadc 
Bridge Times        :Hello 2s MaxAge 20s FwDly 15s MaxHop 20 
CIST Root/ERPC      :4096 .4c1f-cc45-aac1 / 20000 
CIST RegRoot/IRPC   :8192 .4c1f-cc45-aadc / 0 
CIST RootPortId     :128.9 
BPDU-Protection     :Disabled 
TC or TCN received  :79 
TC count per hello  :1 
STP Converge Mode   :Normal  
Share region-configuration :Enabled 
Time since last TC  :0 days 0h:0m:0s 
ĂĂoutput omitĂĂ 
 
[S2]display stp  
-------[CIST Global Info][Mode STP]------- 
CIST Bridge         :4096 .4c1f-cc45-aac1 
Bridge Times        :Hello 2s MaxAge 20s FwDly 15s MaxHop 20 
CIST Root/ERPC      :4096 .4c1f-cc45-aac1 / 0 
CIST RegRoot/IRPC   :4096 .4c1f-cc45-aac1 / 0 
CIST RootPortId     :0.0

HCDP-IESN  Chapter 2 STP and SEP 
 
Page64 HUAWEI TECHNOLOGIES HC Series 
 
BPDU-Protection     :Disabled 
TC or TCN received  :88 
TC count per hello  :0 
STP Converge Mode   :Normal  
Share region-configuration :Enabled 
Time since last TC  :0 days 0h:0m:9s 
ĂĂoutput omitĂĂ 
 
S2 becomes the root bridge as it has a higher priority than S1. 
Step 3 Control root interface election. 
View role information about interfaces of S1. 
[S1]display stp  brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/9        ROOT  FORWARDING      NONE 
   0    GigabitEthernet0/0/10       ALTE  DISCARDING      NONE 
   0    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/14       DESI  FORWARDING      NONE 
 
The GigabitEthernet 0/0/9 interface of S1 is the root interface. 
The default priority is 128. The lowe r the priority value, the larger the 
priority. 
S1 interconnects with S2 over the G0/0/9 and G0/0/10 interfaces. 
Set the priority to 32 for the G0/0/9  interface of S2 and to 16 for the 
G0/0/10 interface. 
[S2]interface GigabitEthernet 0/0/9 
[S2-GigabitEthernet0/0/9]stp port priority 32 
[S2-GigabitEthernet0/0/9]interface GigabitEthernet 0/0/10 
[S2-GigabitEthernet0/0/10]stp port priority 16 
 
Note: The priorities of the interfaces of S2, instead of S1, are changed. 
View role information about interfaces of S1. 
[S1]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/9        ALTE  DISCARDING      NONE 
   0    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   0    GigabitEthernet0/0/13       DESI  FORWARDING      NONE

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page65 
 
   0    GigabitEthernet0/0/14       DESI  FORWARDING      NONE 
 
The GigabitEthernet 0/0/10 interface of S1 becomes the root interface. 
Step 4 Control designated port election. 
View status information about the direct interfaces of S3 and S4. 
 [S3]display stp interface Ethernet 0/0/1 
----[CIST][Port1(Ethernet0/0/1)][DISCARDING]---- 
 Port Protocol       :Enabled 
 Port Role           :Alternate Port 
 Port Priority       :128 
 Port Cost(Dot1T )   :Config=auto / Active=199999 
 Designated Bridge/Port   :32768.5489-98ec-f00a / 128.1 
 Port Edged          :Config=default / Active=disabled 
 Point-to-point      :Config=auto / Active=true 
 Transit Limit       :147 packets/hello-time 
 Protection Type     :None 
 Port STP Mode       :STP 
 Port Protocol Type  :Config=auto / Active=dot1s 
 PortTimes           :Hello 2s MaxAge 20s FwDly 15s RemHop 0 
 TC or TCN send      :17 
 TC or TCN received  :52 
 BPDU Sent           :172 
          TCN: 0, Config: 172, RST: 0, MST: 0 
 BPDU Received       :206 
          TCN: 0, Config: 206, RST: 0, MST: 0 
 
[S4]display stp interface Ethernet 0/0/1  
----[CIST][Port1(Ethernet0/0/1)][FORWARDING]----  
 Port Protocol       :Enabled 
 Port Role           :Designated Port 
 Port Priority       :128 
 Port Cost(Dot1T )   :Config=auto / Active=199999 
 Designated Bridge/Port   :32768.5489-98ec-f00a / 128.1 
 Port Edged          :Config=default / Active=disabled 
 Point-to-point      :Config=auto / Active=true 
 Transit Limit       :147 packets/hello-time 
 Protection Type     :None 
 Port STP Mode       :STP  
 Port Protocol Type  :Config=auto / Active=dot1s

HCDP-IESN  Chapter 2 STP and SEP 
 
Page66 HUAWEI TECHNOLOGIES HC Series 
 
