---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-7
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [1424, 1634]
sha256: 1936ad72414f66aa2aaa19f855879d6f99041807e22566d2f3f26b8e51f4edab
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

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



                                                    HUAWEI TECHNOLOGIES     Page33
Time since last TC :0 days 0h:8m:4s
……output omit……


The highlighted lines in the preceding information indicate that S2 has been
restored and has become the root bridge once again.

Step 3 Control root port election.

Run the display stp brief command on S1 to view the roles of the interfaces.
<S1>display stp brief
    MSTID      Port                     Role STP State Protection
    0          GigabitEthernet0/0/9     ROOT      FORWARDINGNONE
    0          GigabitEthernet0/0/10 ALTE         DISCARDING NONE


The preceding information shows that G0/0/9 is the root port and G0/0/10 is the
alternate port. You can change port priorities so that port interface G0/0/10 will
become the root port and G0/0/9 will become the alternate port.
Change priorities of G0/0/9 and G0/0/10 on S2.
The default port priority is 128. A larger port priority value indicates a lower priority.
The priorities of G0/0/9 and G0/0/10 on S2 are set to 32 and 16; therefore, G0/0/10
on S1 becomes the root port.
[S2]interface GigabitEthernet 0/0/9
[S2-GigabitEthernet0/0/9]stp port priority 32
[S2-GigabitEthernet0/0/9]quit
[S2]interface GigabitEthernet 0/0/10
[S2-GigabitEthernet0/0/10]stp port priority 16


Note that the port priorities are changed on S2, not S1.
<S2>display stp interface GigabitEthernet 0/0/9
-------[CIST Global Info][Mode STP]-------
CIST Bridge           :4096 .d0d0-4ba6-ac20
Config Times          :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
Active Times          :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
CIST Root/ERPC          :4096 .d0d0-4ba6-ac20 / 0 (This bridge is the root)
CIST RegRoot/IRPC       :4096 .d0d0-4ba6-ac20 / 0
CIST RootPortId       :0.0
BPDU-Protection         :Disabled
TC or TCN received :147
TC count per hello :0



                                                    HUAWEI TECHNOLOGIES        Page34
STP Converge Mode            :Normal
Share region-configuration :Enabled
Time since last TC :0 days 0h:7m:35s
Number of TC             :41
Last TC occurred       :GigabitEthernet0/0/10
----[Port34(GigabitEthernet0/0/9)][FORWARDING]----
 Port Protocol          :Enabled
 Port Role              :Designated Port
 Port Priority         :32
 Port Cost(Dot1T )      :Config=auto / Active=20000
 Designated Bridge/Port             :4096.d0d0-4ba6-ac20 / 32.34
 Port Edged              :Config=default / Active=disabled
 Point-to-point         :Config=auto / Active=true
 Transit Limit         :6 packets/s
 Protection Type         :None
 Port STP Mode               :STP
 Port Protocol Type     :Config=auto / Active=dot1s
 BPDU Encapsulation :Config=stp / Active=stp
 PortTimes               :Hello 2s MaxAge 20s FwDly 15s RemHop 20
 TC or TCN send           :35
 TC or TCN received :2
 BPDU Sent                   :1013
             TCN: 0, Config: 1013, RST: 0, MST: 0
 BPDU Received               :2
             TCN: 2, Config: 0, RST: 0, MST: 0
 Last forwarding time: 2016/11/22 10:00:00 UTC


<S2>display stp interface GigabitEthernet 0/0/10
-------[CIST Global Info][Mode STP]-------
CIST Bridge            :4096 .d0d0-4ba6-ac20
Config Times            :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
Active Times           :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
CIST Root/ERPC           :4096 .d0d0-4ba6-ac20 / 0 (This bridge is the root)
CIST RegRootIRPC        :4096 .d0d0-4ba6-ac20 / 0
CIST RootPortId         :0.0
BPDU-Protection          :Disabled
TC or TCN received :147
TC count per hello :0
STP Converge Mode            :Normal
Share region-configuration :Enabled
Time since last TC :0 days 0h:8m:19s
Number of TC             :41



                                                        HUAWEI TECHNOLOGIES    Page35
Last TC occurred        :GigabitEthernet0/0/10
----[Port35(GigabitEthernet0/0/10)][FORWARDING]----
 Port Protocol           :Enabled
 Port Role               :Designated Port
 Port Priority          :16
 Port Cost(Dot1T )       :Config=auto / Active=20000
 Designated Bridge/Port              :4096.d0d0-4ba6-ac20 / 16.35
 Port Edged               :Config=default / Active=disabled
 Point-to-point          :Config=auto / Active=true
 Transit Limit          :6 packets/s
 Protection Type         :None
 Port STP Mode                :STP
 Port Protocol Type      :Config=auto / Active=dot1s
 BPDU Encapsulation :Config=stp / Active=stp
 PortTimes                :Hello 2s MaxAge 20s FwDly 15s RemHop 20
 TC or TCN send           :35
 TC or TCN received :1
 BPDU Sent                    :1032
             TCN: 0, Config: 1032, RST: 0, MST: 0
 BPDU Received                :2
             TCN: 1, Config: 1, RST: 0, MST: 0
 Last forwarding time: 2016/11/22 10:00:11 UTC


Run the display stp brief command on S1 to view the role of the interfaces.
<S1>display stp brief
     MSTID       Port                          Role STP State Protection
     0           GigabitEthernet0/0/9          ALTE     DISCARDING NONE
     0           GigabitEthernet0/0/10 ROOT             FORWARDINGNONE


The highlighted lines in the preceding information indicate that G0/0/10 on S1 has
become the root port and G0/0/9 has become the alternate port.
Shut down G0/0/10 on S1 and view the port roles.
[S1]interface GigabitEthernet 0/0/10
[S1-GigabitEthernet0/0/10]shutdown
<S1>display stp brief
     MSTID       Port                          Role STP State Protection
     0           GigabitEthernet0/0/9          ROOT     FORWARDINGNONE


The highlighted line in the preceding information indicates that G0/0/9 has become
the root port. Resume the default priorities of G0/0/9 and G0/0/10 on S2 and


                                                         HUAWEI TECHNOLOGIES   Page36
re-enable the shutdown interfaces on S1.
[S2]interface GigabitEthernet 0/0/9
[S2-GigabitEthernet0/0/9]undo stp port priority
[S2-GigabitEthernet0/0/9]quit
[S2]interface GigabitEthernet 0/0/10
[S2-GigabitEthernet0/0/10]undo stp port priority



[S1]interface GigabitEthernet 0/0/10
[S1-GigabitEthernet0/0/10]undo shutdown


Run the display stp brief and display stp interface command on S1 to view the
roles of interfaces.
<S1>display stp brief
     MSTID       Port                       Role STP State Protection
     0           GigabitEthernet0/0/9       ROOT     FORWARDINGNONE
     0           GigabitEthernet0/0/10 ALTE          DISCARDING NONE



[S1]display stp interface GigabitEthernet 0/0/9
----[CIST][Port9(GigabitEthernet0/0/9)][FORWARDING]----
 Port Protocol           :Enabled
 Port Role               :Root Port
 Port Priority          :128
 Port Cost(Dot1T )       :Config=auto / Active=20000
 Designated Bridge/Port          :4096.4c1f-cc45-aacc / 128.9
 Port Edged               :Config=default / Active=disabled
 Point-to-point          :Config=auto / Active=true
 Transit Limit          :147 packets/hello-time
 Protection Type         :None
 Port STP Mode            :STP
 Port Protocol Type      :Config=auto / Active=dot1s
 BPDU Encapsulation :Config=stp / Active=stp
 PortTimes                :Hello 2s MaxAge 20s FwDly 15s RemHop 0
 TC or TCN send           :4
 TC or TCN received :90
 BPDU Sent                :5
             TCN: 4, Config: 1, RST: 0, MST: 0
 BPDU Received             :622
             TCN: 0, Config: 622, RST: 0, MST: 0


[S1]display stp interface GigabitEthernet 0/0/10



                                                      HUAWEI TECHNOLOGIES   Page37

