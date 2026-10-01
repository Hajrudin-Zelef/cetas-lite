---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-6
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [1198, 1423]
sha256: 69cf704a13ee54f2a3036d84901fbfea49ad17fd50d95edced77b5557b7cd0f9
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

 Port Priority         :128
 Port Cost(Dot1T )      :Config=auto / Active=20000
 Designated Bridge/Port          :0.d0d0-4ba6-aab0 / 128.10
 Port Edged              :Config=default / Active=disabled
 Point-to-point         :Config=auto / Active=true
 Transit Limit         :6 packets/s
 Protection Type         :None
 Port STP Mode            :STP
 Port Protocol Type     :Config=auto / Active=dot1s
 BPDU Encapsulation :Config=stp / Active=stp
 PortTimes               :Hello 2s MaxAge 20s FwDly 15s RemHop 20
 TC or TCN send          :52
 TC or TCN received :0
 BPDU Sent                :3189
             TCN: 0, Config: 3189, RST: 0, MST: 0
 BPDU Received            :5
             TCN: 0, Config: 5, RST: 0, MST: 0
 Last forwarding time: 2016/11/21 14:55:11 UTC




<S2>display stp interface GigabitEthernet 0/0/10
-------[CIST Global Info][Mode STP]-------
CIST Bridge            :4096 .d0d0-4ba6-ac20
Config Times            :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
Active Times           :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
CIST Root/ERPC           :0      .d0d0-4ba6-aab0 / 20000
CIST RegRoot/IRPC        :4096 .d0d0-4ba6-ac20 / 0
CIST RootPortId         :128.9 (GigabitEthernet0/0/9)
BPDU-Protection          :Disabled
CIST Root Type          :Secondary root
TC or TCN received :122
TC count per hello :0
STP Converge Mode         :Normal
Share region-configuration :Enabled
Time since last TC :0 days 1h:50m:0s
Number of TC             :17
Last TC occurred       :GigabitEthernet0/0/9
----[Port10(GigabitEthernet0/0/10)][DISCARDING]----
 Port Protocol          :Enabled
 Port Role              :Alternate Port
 Port Priority         :128



                                                     HUAWEI TECHNOLOGIES   Page29
 Port Cost(Dot1T )    :Config=auto / Active=20000
 Designated Bridge/Port           :0.d0d0-4ba6-aab0 / 128.10
 Port Edged             :Config=default / Active=disabled
 Point-to-point        :Config=auto / Active=true
 Transit Limit        :6 packets/s
 Protection Type       :None
 Port STP Mode             :STP
 Port Protocol Type    :Config=auto / Active=dot1s
 BPDU Encapsulation :Config=stp / Active=stp
 PortTimes              :Hello 2s MaxAge 20s FwDly 15s RemHop 0
 TC or TCN send            :0
 TC or TCN received :18
 BPDU Sent                 :2
          TCN: 0, Config: 2, RST: 0, MST: 0
 BPDU Received             :3317
          TCN: 0, Config: 3317, RST: 0, MST: 0




Step 2 Control root bridge election.

Run the display stp command to view information about the root bridge.
<S1>display stp
-------[CIST Global Info][Mode STP]-------
CIST Bridge           :0        .d0d0-4ba6-aab0
Config Times          :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
Active Times          :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
CIST Root/ERPC         :0         .d0d0-4ba6-aab0 / 0 (This bridge is the root)
CIST RegRoot/IRPC       :0        .d0d0-4ba6-aab0 / 0
CIST RootPortId       :0.0
BPDU-Protection        :Disabled
CIST Root Type        :Primary root
TC or TCN received :11
TC count per hello :0
STP Converge Mode          :Normal
Share region-configuration :Enabled
Time since last TC :0 days 2h:32m:25s
……output omit……




<S2>display stp
-------[CIST Global Info][Mode STP]-------


                                                        HUAWEI TECHNOLOGIES       Page30
CIST Bridge             :4096 .d0d0-4ba6-ac20
Config Times            :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
Active Times            :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
CIST Root/ERPC           :0   .d0d0-4ba6-aab0 / 20000
CIST RegRoot/IRPC        :4096 .d0d0-4ba6-ac20 / 0
CIST RootPortId         :128.9 (GigabitEthernet0/0/9)
BPDU-Protection          :Disabled
CIST Root Type          :Secondary root
TC or TCN received :122
TC count per hello :0
STP Converge Mode         :Normal
Share region-configuration :Enabled
Time since last TC :0 days 2h:35m:57s
……output omit……


Configure S2 as the root bridge and S1 as the backup root bridge using priority
values. The device with the same value for the CIST Bridge and CIST Root/ERPC is
the root bridge. A smaller bridge priority value indicates a higher bridge priority.
Change the priorities of S1 and S2 to 8192 and 4096 respectively so that S2 becomes
the root bridge.
[S1]undo stp root
[S1]stp priority 8192


[S2]undo stp root
[S2]stp priority 4096


Run the display stp command to view information about the new root bridge.
<S1>display stp
-------[CIST Global Info][Mode STP]-------
CIST Bridge             :8192 .d0d0-4ba6-aab0
Config Times            :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
Active Times            :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
CIST Root/ERPC           :4096 .d0d0-4ba6-ac20 / 20000
CIST RegRoot/IRPC        :8192 .d0d0-4ba6-aab0 / 0
CIST RootPortId         :128.9 (GigabitEthernet0/0/9)
BPDU-Protection          :Disabled
TC or TCN received :47
TC count per hello :0
STP Converge Mode         :Normal
Share region-configuration :Enabled
Time since last TC :0 days 0h:6m:55s



                                                     HUAWEI TECHNOLOGIES   Page31
……output omit……


<S2>display stp
-------[CIST Global Info][Mode STP]-------
CIST Bridge          :4096 .d0d0-4ba6-ac20
Config Times         :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
Active Times         :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
CIST Root/ERPC        :4096 .d0d0-4ba6-ac20 / 0 (This bridge is the root)
CIST RegRoot/IRPC       :4096 .d0d0-4ba6-ac20 / 0
CIST RootPortId      :0.0
BPDU-Protection       :Disabled
TC or TCN received :135
TC count per hello :0
STP Converge Mode        :Normal
Share region-configuration :Enabled
Time since last TC :0 days 0h:8m:4s
……output omit……


The highlighted lines in the preceding information indicate that S2 has become the
new root bridge.
Shut down interfaces Gigabit Ethernet 0/0/9 and GigabitGigabitEthernet 0/0/10 on
S2 to isolate S2.
[S2]interface GigabitEthernet 0/0/9
[S2-GigabitEthernet0/0/9]shutdown
[S2-GigabitEthernet0/0/9]quit
[S2]interface GigabitEthernet 0/0/10
[S2-GigabitEthernet0/0/10]shutdown


<S1>display stp
-------[CIST Global Info][Mode STP]-------
CIST Bridge          :8192 .d0d0-4ba6-aab0
Config Times         :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
Active Times         :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
CIST Root/ERPC        :8192 .d0d0-4ba6-aab0 / 0 (This bridge is the root)
CIST RegRoot/IRPC       :8192 .d0d0-4ba6-aab0 / 0
CIST RootPortId      :0.0
BPDU-Protection       :Disabled
TC or TCN received :174
TC count per hello :0
STP Converge Mode        :Normal
Share region-configuration :Enabled



                                                    HUAWEI TECHNOLOGIES     Page32
Time since last TC :0 days 0h:12m:51s
……output omit……


The highlighted lines in the preceding information indicate that S1 becomes the root
bridge when S2 is faulty.
Re-enable the interfaces that have been disabled on S2.
[S2]interface GigabitEthernet 0/0/9
[S2-GigabitEthernet0/0/9]undo shutdown
[S2-GigabitEthernet0/0/9]quit
[S2]interface GigabitEthernet 0/0/10
[S2-GigabitEthernet0/0/10]undo shutdown


<S1>display stp
-------[CIST Global Info][Mode STP]-------
CIST Bridge          :8192 .d0d0-4ba6-aab0
Config Times         :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
Active Times         :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
CIST Root/ERPC        :4096 .d0d0-4ba6-ac20 / 20000
CIST RegRoot/IRPC       :8192 .d0d0-4ba6-aab0 / 0
CIST RootPortId      :128.9 (GigabitEthernet0/0/9)
BPDU-Protection       :Disabled
TC or TCN received :47
TC count per hello :0
STP Converge Mode        :Normal
Share region-configuration :Enabled
Time since last TC :0 days 0h:6m:55s
……output omit……


