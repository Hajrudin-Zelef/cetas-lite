---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-8
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [1635, 1960]
sha256: e738fc6d5b030daa6f40d7e82c8bdf1ced12dedf56a0c49e7f709f1544a97dd8
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

----[CIST][Port10(GigabitEthernet0/0/10)][DISCARDING]----
 Port Protocol          :Enabled
 Port Role              :Alternate Port
 Port Priority         :128
 Port Cost(Dot1T )      :Config=auto / Active=20000
 Designated Bridge/Port          :4096.4c1f-cc45-aacc / 128.10
 Port Edged              :Config=default / Active=disabled
 Point-to-point         :Config=auto / Active=true
 Transit Limit         :147 packets/hello-time
 Protection Type         :None
 Port STP Mode            :STP
 Port Protocol Type     :Config=auto / Active=dot1s
 BPDU Encapsulation :Config=stp / Active=stp
 PortTimes               :Hello 2s MaxAge 20s FwDly 15s RemHop 0
 TC or TCN send          :3
 TC or TCN received :90
 BPDU Sent                :4
             TCN: 3, Config: 1, RST: 0, MST: 0
 BPDU Received            :637
             TCN: 0, Config: 637, RST: 0, MST: 0


The greyed line in the preceding information indicates that G0/0/9 and G0/0/10 cost
is 20000 by default.
Change the cost of G0/0/9 to 200000 on S1.
[S1]interface GigabitEthernet 0/0/9
[S1-GigabitEthernet0/0/9]stp cost 200000


Run the display stp brief and display stp interface command on S1 to view the
roles of interfaces.
<S1>display stp interface GigabitEthernet 0/0/9
----[CIST][Port9(GigabitEthernet0/0/9)][DISCARDING]----
 Port Protocol          :Enabled
 Port Role              :Alternate Port
 Port Priority         :128
 Port Cost(Dot1T )      :Config=200000 / Active=200000
 Designated Bridge/Port          :4096.4c1f-cc45-aacc / 128.9
 Port Edged              :Config=default / Active=disabled
 Point-to-point         :Config=auto / Active=true
 Transit Limit         :147 packets/hello-time
 Protection Type         :None




                                                      HUAWEI TECHNOLOGIES   Page38
 Port STP Mode           :STP
 Port Protocol Type     :Config=auto / Active=dot1s
 BPDU Encapsulation :Config=stp / Active=stp
 PortTimes              :Hello 2s MaxAge 20s FwDly 15s RemHop 0
 TC or TCN send          :4
 TC or TCN received :108
 BPDU Sent               :5
          TCN: 4, Config: 1, RST: 0, MST: 0
 BPDU Received           :818
          TCN: 0, Config: 818, RST: 0, MST: 0


<S1>display stp brief
    MSTID       Port                    Role STP State Protection
    0           GigabitEthernet0/0/9    ALTE     DISCARDING NONE
    0           GigabitEthernet0/0/10 ROOT       FORWARDINGNONE


The highlighted lines in the preceding information indicates that G0/0/10 has
become the root port.


Final Configuration

<S1>display current-configuration
#
!Software Version V200R008C00SPC500
 sysname S1
#
 stp mode stp
 stp instance 0 priority 8192
#
interface GigabitEthernet0/0/1
 shutdown
#
interface GigabitEthernet0/0/2
 shutdown
#
interface GigabitEthernet0/0/3
 shutdown
#
interface GigabitEthernet0/0/9
 stp instance 0 cost 200000
#
interface GigabitEthernet0/0/10


                                                  HUAWEI TECHNOLOGIES   Page39
#
interface GigabitEthernet0/0/13
 shutdown
#
interface GigabitEthernet0/0/14
 shutdown
#
user-interface con 0
user-interface vty 0 4
#
return


<S2>display current-configuration
#
!Software Version V200R008C00SPC500
 sysname S2
#
 stp mode stp
 stp instance 0 priority 4096
#
interface GigabitEthernet0/0/1
 shutdown
#
interface GigabitEthernet0/0/2
 shutdown
#
interface GigabitEthernet0/0/3
 shutdown
#
interface GigabitEthernet0/0/6
 shutdown
#
interface GigabitEthernet0/0/7
 shutdown
#
interface GigabitEthernet0/0/9
#
interface GigabitEthernet0/0/10
#
user-interface con 0
user-interface vty 0 4
#



                                      HUAWEI TECHNOLOGIES   Page40
return


<S3>display current-configuration
#
!Software Version V100R006C05
 sysname S3
#
interface GigabitEthernet0/0/1
 shutdown
#
interface GigabitEthernet0/0/13
 shutdown
#
interface GigabitEthernet0/0/7
 shutdown
#
user-interface con 0
user-interface vty 0 4
#
return


<S4>display current-configuration
#
!Software Version V100R006C05
 sysname S4
#
interface GigabitEthernet0/0/14
 shutdown
#
interface Gigabit
Ethernet0/0/1
 shutdown
#
interface GigabitEthernet0/0/6
 shutdown
#
user-interface con 0
user-interface vty 0 4
#
return




                                    HUAWEI TECHNOLOGIES   Page41
Lab 3-2 Configuring RSTP


Learning Objectives

As a result of this lab section, you should achieve the following tasks:

       Enable and disable RSTP .

       Configuration of an edge port.

       Configuration of RSTP BPDU protection.

       Configuration of RSTP loop protection


Topology




                                Figure 3.2 RSTP topology



Scenario

Assume that you are a network administrator of a company. The company network
consists of two layers: core layer and access layer. The network uses a redundancy
design. RSTP will be used to prevent loops. You can configure features to speed up
RSTP route convergence at the edge network and configure RSTP protection
function.


Tasks


Step 1 Preparing the environment

If you are starting this section with a non-configured device, begin here and then


                                      HUAWEI TECHNOLOGIES                  Page42
move to step 3. For those continuing from previous labs, begin at step 2.

Irrelevant interfaces must be disabled to ensure test result accuracy.
Shut down port interfacesGigabitEthernet 0/0/1 on S3,GigabitEthernet 0/0/13 and
Ethernet 0/0/7 on S3; GigabitEthernet 0/0/1, GigabitEthernet 0/0/2, GigabitEthernet
0/0/3, GigabitEthernet 0/0/13, GigabitEthernet 0/0/14 on S1; GigabitEthernet 0/0/1,
GigabitEthernet 0/0/2, GigabitEthernet 0/0/3, GigabitEthernet 0/0/6, GigabitEthernet
0/0/7 on S2;       as well asGigabitEthernet 0/0/1,GigabitEthernet 0/0/14 and
GigabitEthernet 0/0/6 on S4 before starting STP configuration. Ensure that the
devices start without any configuration files. If STP is disabled, run the stp enable
command to enable STP.


<Quidway>system-view
Enter system view, return user view with Ctrl+Z.
[Quidway]sysname S1
[S1]interface GigabitEthernet 0/0/1
[S1-GigabitEthernet0/0/1]shutdown
[S1-GigabitEthernet0/0/1]quit
[S1]interface GigabitEthernet 0/0/2
[S1-GigabitEthernet0/0/2]shutdown
[S1-GigabitEthernet0/0/2]quit
[S1]interface GigabitEthernet 0/0/3
[S1-GigabitEthernet0/0/3]shutdown
[S1-GigabitEthernet0/0/3]quit
[S1]interface GigabitEthernet 0/0/13
[S1-GigabitEthernet0/0/13]shutdown
[S1-GigabitEthernet0/0/13]quit
[S1]interface GigabitEthernet 0/0/14
[S1-GigabitEthernet0/0/14]shutdown
[S1-GigabitEthernet0/0/14]quit


<Quidway>system-view
Enter system view, return user view with Ctrl+Z.
[Quidway]sysname S2
[S2]interface GigabitEthernet 0/0/1
[S2-GigabitEthernet0/0/1]shutdown
[S2-GigabitEthernet0/0/1]quit
[S2]interface GigabitEthernet 0/0/2
[S2-GigabitEthernet0/0/2]shutdown
[S2-GigabitEthernet0/0/2]quit
[S2]interface GigabitEthernet 0/0/3



                                                   HUAWEI TECHNOLOGIES      Page43
[S2-GigabitEthernet0/0/3]shutdown
[S2-GigabitEthernet0/0/3]quit
[S2]interface GigabitEthernet 0/0/6
[S2-GigabitEthernet0/0/6]shutdown
[S2-GigabitEthernet0/0/6]quit
[S2]interface GigabitEthernet 0/0/7
[S2-GigabitEthernet0/0/7]shutdown
[S2-GigabitEthernet0/0/7]quit


<Quidway>system-view
Enter system view, return user view with Ctrl+Z.
[Quidway]sysname S3
[S3]interface GigabitEthernet 0/0/1
[S3-GigabitEthernet 0/0/1]shutdown
[S3-GigabitEthernet 0/0/1]quit
[S3]interface GigabitEthernet 0/0/13
[S3-GigabitEthernet 0/0/13]shutdown
[S3-GigabitEthernet 0/0/13]quit
[S3]interface GigabitEthernet 0/0/7
[S3-GigabitEthernet0/0/7]shutdown


<Quidway>system-view
Enter system view, return user view with Ctrl+Z.
[Quidway]sysname S4
[S4]interface GigabitEthernet 0/0/1
[S4-GigabitEthernet 0/0/1]shutdown
[S4-GigabitEthernet 0/0/1]quit
[S4]interface GigabitEthernet 0/0/14
[S4-GigabitEthernet 0/0/14]shutdown
[S4-GigabitEthernet 0/0/14]quit
[S4]interface GigabitEthernet 0/0/6
[S4-GigabitEthernet0/0/6]shutdown




Step 2 Clean up the previous configuration


