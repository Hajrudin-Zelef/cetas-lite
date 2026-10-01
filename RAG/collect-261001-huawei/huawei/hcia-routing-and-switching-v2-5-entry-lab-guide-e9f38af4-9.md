---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-9
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [1961, 2270]
sha256: 70755d9f23ec3745a55eb5b872536c222d45151fb9efa2c44d874a372a06bdf2
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

Remove the configured STP priority from S1 and S2, and assigned cost on S1.

[S1]undo stp priority
[S1]inter GigabitEthernet 0/0/9




                                                   HUAWEI TECHNOLOGIES   Page44
[S1-GigabitEthernet0/0/9]undo stp cost


[S2]undo stp priority

Step 3 Configure RSTP and verify the RSTP configuration.

Configure S1 and S2 to use RSTP as the spanning tree protocol.
 [S1]stp mode rstp


 [S2]stp mode rstp


Run the display stp command to view brief information about RSTP.
[S1]display stp
-------[CIST Global Info][Mode RSTP]-------
CIST Bridge             :32768.d0d0-4ba6-aab0
Config Times            :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
Active Times            :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
CIST Root/ERPC           :32768.d0d0-4ba6-aab0 / 0 (This bridge is the root)
CIST RegRoot/IRPC        :32768.d0d0-4ba6-aab0 / 0
CIST RootPortId         :0.0
BPDU-Protection          :Disabled
TC or TCN received :362
TC count per hello :0
STP Converge Mode         :Normal
Share region-configuration :Enabled
Time since last TC :0 days 0h:0m:45s
……output omit……


[S2]display stp
-------[CIST Global Info][Mode RSTP]-------
CIST Bridge             :32768.d0d0-4ba6-ac20
Config Times            :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
Active Times            :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
CIST Root/ERPC           :32768.d0d0-4ba6-aab0 / 20000
CIST RegRoot/IRPC        :32768.d0d0-4ba6-ac20 / 0
CIST RootPortId         :128.34 (GigabitEthernet0/0/9)
BPDU-Protection          :Disabled
TC or TCN received :186
TC count per hello :0
STP Converge Mode         :Normal
Share region-configuration :Enabled



                                                    HUAWEI TECHNOLOGIES        Page45
Time since last TC :0 days 0h:3m:55s
……output omit……



Step 4 Configure an edge port.

Configure ports connected to the user terminals as edge ports. An edge port can
transition to the forwarding state without participating in the RSTP calculation. In
this example, interface GigabitEthernet 0/0/1 on S1 and S2 connect to a router and
can be configured as edge ports.
[S1]interface GigabitEthernet 0/0/1
[S1-GigabitEthernet0/0/1]undo shutdown
[S1-GigabitEthernet0/0/1]stp edged-port enable


[S2]interface GigabitEthernet 0/0/1
[S2-GigabitEthernet0/0/1]undo shutdown
[S2-GigabitEthernet0/0/1]stp edged-port enable



Step 5 Configure BPDU protection.

Edge ports are directly connected to user terminal and will not receive BPDUs.
Attackers may send pseudo BPDUs to attack the switching device. If the edge ports
receive the BPDUs, the switching device configures the edge ports as non-edge
ports and triggers a new spanning tree calculation. Network flapping then occurs.
BPDU protection can be used to protect switching devices against malicious attacks.
Configure BPDU protection on both S1 and S2.
[S1]stp bpdu-protection


[S2]stp bpdu-protection


Run the display stp brief command to view the port protection.
<S1>display stp brief
 MSTID Port                            Role STP State    Protection
   0    GigabitEthernet0/0/1           DESI FORWARDING        BPDU
   0    GigabitEthernet0/0/9           DESI FORWARDING        NONE
   0    GigabitEthernet0/0/10          DESI FORWARDING        NONE


<S2>display stp brief
 MSTID Port                            Role STP State    Protection



                                                 HUAWEI TECHNOLOGIES     Page46
   0    GigabitEthernet0/0/1           DESI FORWARDING        BPDU
   0    GigabitEthernet0/0/9           ROOT    FORWARDING      NONE
   0    GigabitEthernet0/0/10          ALTE DISCARDING       NONE


After the configuration is complete, interface GigabitGigabitEthernet 0/0/1 on S1
and S2 shows as supporting BPDU protection.



Step 6 Configure Loop protection

On a network running RSTP, a switching device maintains the root port status and
status of alternate ports by receiving BPDUs from an upstream switching device. If
the switching device cannot receive BPDUs from the upstream device because of link
congestion or unidirectional-link failure, the switching device re-selects a root port.
The original root port becomes a designated port and the original discarding ports
change to the Forwarding state. This switching may cause network loops, which can
be mitigated by configuring loop protection.
Configure loop protection on both the root port and the alternate port.
[S2]display stp brief
 MSTID Port                            Role STP State    Protection
   0    GigabitEthernet0/0/1           DESI FORWARDING        BPDU
   0    GigabitEthernet0/0/9           ROOT    FORWARDING      NONE
   0    GigabitEthernet0/0/10          ALTE DISCARDING       NONE


G0/0/9 and G0/0/10 on S2 are now the root port and alternate port. Configure loop
protection on these two ports.
[S2]interface GigabitEthernet 0/0/9
[S2-GigabitEthernet0/0/9]stp loop-protection
[S2-GigabitEthernet0/0/9]quit
[S2]interface GigabitEthernet 0/0/10
[S2-GigabitEthernet0/0/10]stp loop-protection



Run the display stp brief command to view the port protection.
<S2>display stp brief
 MSTID Port                            Role STP State    Protection
   0    GigabitEthernet0/0/1           DESI FORWARDING        BPDU
   0    GigabitEthernet0/0/9           ROOT    FORWARDING      LOOP
   0    GigabitEthernet0/0/10          ALTE DISCARDING       LOOP




                                                 HUAWEI TECHNOLOGIES        Page47
Since S1 is root, all the ports are designated ports and therefore do not need to
configure loop protection. After completing the configuration, you may wish to set
S2 as the root, and configure loop protection on the root port and alternate port of
S1 using the same process as with S2.


Final Configuration

<S1>display current-configuration
#
!Software Version V200R008C00SPC500
 sysname S1
#
 stp mode rstp
 stp bpdu-protection
#
interface GigabitEthernet0/0/1
undo shutdown
 stp edged-port enable
#
interface GigabitEthernet0/0/2
 shutdown
#
interface GigabitEthernet0/0/3
 shutdown
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




                                      HUAWEI TECHNOLOGIES                Page48
#
 stp mode rstp
 stp bpdu-protection
#
interface GigabitEthernet0/0/1
undo shutdown
 stp edged-port enable
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
 stp loop-protection
#
interface GigabitEthernet0/0/10
 stp loop-protection
#
user-interface con 0
user-interface vty 0 4
#
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



                                    HUAWEI TECHNOLOGIES   Page49
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
interface GigabitEthernet0/0/1
 shutdown
#


interface GigabitEthernet0/0/6
 shutdown
#
user-interface con 0
user-interface vty 0 4
#
return




                                    HUAWEI TECHNOLOGIES   Page50
                   Module 4 Routing Configuration

Lab 4-1 Configuring Static Routes and Default Routes


Learning Objectives

As a result of this lab section, you should achieve the following tasks:

      Configuration of a static route using an interface and an IP address as the

       next hop.

      Verification of static route operation.

      Implementation of the interconnection between a local and external network

       using a default route.

