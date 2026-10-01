---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-5
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["cost", "ethernet", "memory"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [932, 1197]
sha256: 0c5a82583c159ecf732be81602edd89043df3db6cac64867e458499403d52914
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

<R3>reset saved-configuration
This will delete the configuration in the flash memory.
The device configurations will be erased to reconfigure.
Are you sure? (y/n)[n]:y
 Clear the configuration in the device successfully.



Step 11        Device restart procedure

Use the reboot command to restart the router.
<R1>reboot
Info: The system is now comparing the configuration, please wait.
Warning: All the configuration will be saved to the next startup configuration. Continue ? [y/n]:n
System will reboot! Continue ? [y/n]:y
Info: system is rebooting ,please wait...


<R3>reboot
Info: The system is now comparing the configuration, please wait.
Warning: All the configuration will be saved to the next startup configuration. Continue ? [y/n]:n
System will reboot! Continue ? [y/n]:y


The system asks to save the current configuration. It is necessary to determine
whether the current configuration should be saved based on the requirements for
the lab. If unsure as to whether the current configuration should be saved, do not
save.


Final Configuration

[R1]display current-configuration



                                                   HUAWEI TECHNOLOGIES                               Page23
[V200R007C00SPC600]
#
 sysname R1
 header shell information "Welcome to Huawei certification lab"
#
interface GigabitEthernet0/0/0
 description This interface connects to R3-G0/0/0
 ip address 10.0.13.1 255.255.255.0
#
user-interface con 0
 authentication-mode password
 set authentication password cipher %$%$4D0K*-E"t/I7[{HD~kgW,%dgkQQ!&|;XTDq9SFQJ.27M%dj,%$%$
 idle-timeout 20 0
#
return


[R3]display current-configuration
[V200R007C00SPC600]
#
 sysname R3
#
interface GigabitEthernet0/0/0
 description This interface connect to R1-G0/0/0
 ip address 10.0.13.3 255.255.255.0
#
user-interface con 0
 authentication-mode password
 set authentication password cipher %$%$M8\HO3:72:ERQ8JLoHU8,%t+lE:$9=a7"8%yMoARB]$B%t.,%$%$
user-interface vty 0 4
#
return




                                                   HUAWEI TECHNOLOGIES             Page24
                          Module 3 STP and RSTP

Lab 3-1 Configuring STP


Learning Objectives

As a result of this lab section, you should achieve the following tasks:

      Enable and disable STP

      Change the STP mode that is used by a switch

      Change the bridge priority to control root bridge election

      Change the port priority to control election of the root port and designated

       port

      Change the port cost to control election of the root port and designated port

      Configure an edge port


Topology




                                Figure 3.1 STP topology



Scenario

Assume that you are a network administrator of a company. The company network
consists of two layers: core layer and access layer. The network uses a design that
supports network redundancy. STP will be used to prevent loops. The STP network


                                      HUAWEI TECHNOLOGIES                  Page25
should include setting the bridge priority to control STP root bridge election, and
configuration of features to speed up STP route convergence.


Tasks


    Step 1 Configure STP and verify the STP configuration.

Irrelevant interfaces must be disabled to ensure test result accuracy.
Shut down port interfacesGigabitEthernet 0/0/1 on S3,GigabitEthernet 0/0/13 and
Ethernet 0/0/7 on S3; GigabitEthernet 0/0/1, GigabitEthernet 0/0/2, GigabitEthernet
0/0/3, GigabitEthernet 0/0/13, GigabitEthernet 0/0/14 on S1; GigabitEthernet 0/0/1,
GigabitEthernet 0/0/2, GigabitEthernet 0/0/3, GigabitEthernet 0/0/6, GigabitEthernet
0/0/7 on S2; as well as GigabitEthernet 0/0/1, GigabitEthernet 0/0/14 and
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




                                                   HUAWEI TECHNOLOGIES    Page26
[S2-GigabitEthernet0/0/1]shutdown
[S2-GigabitEthernet0/0/1]quit
[S2]interface GigabitEthernet 0/0/2
[S2-GigabitEthernet0/0/2]shutdown
[S2-GigabitEthernet0/0/2]quit
[S2]interface GigabitEthernet 0/0/3
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
[S3-GigabitEthernet0/0/1]shutdown
[S3-GigabitEthernet0/0/1]quit
[S3]interface GigabitEthernet 0/0/13
[S3-GigabitEthernet0/0/13]shutdown
[S3-GigabitEthernet0/0/13]quit
[S3]interface GigabitEthernet 0/0/7
[S3-GigabitEthernet0/0/7]shutdown


<Quidway>system-view
Enter system view, return user view with Ctrl+Z.
[Quidway]sysname S4
[S4]inter GigabitEthernet 0/0/1
[S4-GigabitEthernet 0/0/1]shutdown
[S4-GigabitEthernet 0/0/1]quit
[S4]inter GigabitEthernet 0/0/14
[S4-GigabitEthernet 0/0/14]shutdown
[S4-GigabitEthernet 0/0/14]quit
[S4]interface GigabitEthernet 0/0/6
[S4-GigabitEthernet0/0/6]shutdown


In the lab, S1 and S2 are connected through two links, and STP is used. Enable STP
on S1 and S2 and set S1 as the root.
[S1]stp mode stp


                                                   HUAWEI TECHNOLOGIES   Page27
Info: This operation may take a few seconds. Please wait for a moment...done.
[S1]stp root primary


[S2]stp mode stp
Info: This operation may take a few seconds. Please wait for a moment...done.
[S2]stp root secondary


Run the display stp brief command to view brief information about STP.
<S1>display stp brief
MSTID         Port                             Role   STP State Protection
    0         GigabitEthernet0/0/9      DESI      FORWARDING        NONE
    0         GigabitEthernet0/0/10            DESI   FORWARDING         NONE


<S2>display stp brief
MSTID    Port                           Role      STP State Protection
     0        GigabitEthernet0/0/9      ROOT FORWARDING            NONE
     0        GigabitEthernet0/0/10 ALTE          DISCARDING       NONE



Run the display stp interface command to view the STP status of a port.



<S1>display stp interface GigabitEthernet 0/0/10
-------[CIST Global Info][Mode STP]-------
CIST Bridge            :0     .d0d0-4ba6-aab0
Config Times           :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
Active Times           :Hello 2s MaxAge 20s FwDly 15s MaxHop 20
CIST Root/ERPC          :0     .d0d0-4ba6-aab0 / 0 (This bridge is the root)
CIST RegRoot/IRPC       :0      .d0d0-4ba6-aab0 / 0
CIST RootPortId        :0.0
BPDU-Protection         :Disabled
CIST Root Type          :Primary root
TC or TCN received :11
TC count per hello :0
STP Converge Mode           :Normal
Share region-configuration :Enabled
Time since last TC :0 days 1h:43m:55s
Number of TC            :29
Last TC occurred       :GigabitEthernet0/0/9
----[Port10(GigabitEthernet0/0/10)][FORWARDING]----
 Port Protocol          :Enabled
 Port Role              :Designated Port


                                                       HUAWEI TECHNOLOGIES      Page28

