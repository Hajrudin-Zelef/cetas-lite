---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-4
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2016-03-10", "2016-03-11"]
keywords: ["license", "memory", "voice"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [714, 931]
sha256: dabaec781d549f00f19ec88db79bd35896479d7df2bf9cabf4428d0595828880
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

Log out of the system and log back in, using the password set. It should be noted
that this password is required to be set when the router is first initialized.
[R1-ui-console0]return
    <R1>quit




                                               HUAWEI TECHNOLOGIES                             Page18
      Configuration console exit, please press any key to log on

    Login authentication

    Password:

    Welcome to Huawei certification lab
    <R1>



Step 8 Configure interface IP addresses and descriptions

Configure an IP address for the GigabitEthernet 0/0/0 interface of R1. The subnet
mask can be configured using a dotted decimal format (255.255.255.0), or based on
the subnet mask prefix length.
[R1]interface GigabitEthernet 0/0/0
[R1-GigabitEthernet0/0/0]ip address 10.0.13.1 24
[R1-GigabitEthernet0/0/0]description This interface connects to R3-G0/0/0


Run the display this command to check the configuration results at the current
interface view.
[R1-GigabitEthernet0/0/0]display this
[V200R007C00SPC600]
#
interface GigabitEthernet0/0/0
 description This interface connects to R3-G0/0/0
 ip address 10.0.13.1 255.255.255.0
#
return


Run the display interface command to view the interface description.
[R1]display interface GigabitEthernet0/0/0
GigabitEthernet0/0/0 current state : UP
Line protocol current state : UP
Last line protocol up time : 2016-03-11 04:13:09
Description:This interface connects to R3-G0/0/0
Route Port,The Maximum Transmit Unit is 1500
Internet Address is 10.0.13.1/24
IP Sending Frames' Format is PKTFMT_ETHNT_2, Hardware address is 5489-9876-830b
Last physical up time      :   2016-03-10 03:24:01
Last physical down time :      2016-03-10 03:25:29
Current system time: 2016-03-11 04:15:30



                                                   HUAWEI TECHNOLOGIES            Page19
Port Mode: FORCE COPPER
Speed : 100, Loopback: NONE
Duplex: FULL, Negotiation: ENABLE
Mdi      : AUTO, Clock    :-
Last 300 seconds input rate 2296 bits/sec, 1 packets/sec
Last 300 seconds output rate 88 bits/sec, 0 packets/sec
Input peak rate 7392 bits/sec,Record time: 2016-03-10 04:08:41
Output peak rate 1120 bits/sec,Record time: 2016-03-10 03:27:56
Input: 3192 packets, 895019 bytes
  Unicast:                     0,        Multicast:              1592
  Broadcast:             1600,      Jumbo:                       0
  Discard:                     0,        Total Error:            0
  CRC:                         0,        Giants:                     0
  Jabbers:                     0,        Throttles:              0
  Runts:                       0,        Symbols:                0
  Ignoreds:                    0,        Frames:                     0
Output: 181 packets, 63244 bytes
  Unicast:                     0,        Multicast:              0
  Broadcast:             181, Jumbo:                         0
  Discard:                     0,        Total Error:            0
  Collisions:            0,         ExcessiveCollisions: 0
  Late Collisions:       0,         Deferreds:                   0
      Input bandwidth utilization threshold : 100.00%
      Output bandwidth utilization threshold: 100.00%
      Input bandwidth utilization : 0.01%
Output bandwidth utilization :      0%


The command output shows that the physical status and protocol status of the
interface are UP, and the corresponding physical layer and data link layer are
functional.
Once the status has been verified, configure the IP address and description for the
interface of R3.
[R3]interface GigabitEthernet 0/0/0
[R3-GigabitEthernet0/0/0]ip address 10.0.13.3 255.255.255.0 [R3-GigabitEthernet0/0/0]description This interface
connects to R1-G0/0/0

After completing the configuration, run the ping command to test the connection
between R1 and R3.
<R1>ping 10.0.13.3
  PING 10.0.13.3: 56 data bytes, press CTRL_C to break
      Reply from 10.0.13.3: bytes=56 Sequence=1 ttl=255 time=35 ms



                                                      HUAWEI TECHNOLOGIES                        Page20
    Reply from 10.0.13.3: bytes=56 Sequence=2 ttl=255 time=32 ms
    Reply from 10.0.13.3: bytes=56 Sequence=3 ttl=255 time=32 ms
    Reply from 10.0.13.3: bytes=56 Sequence=4 ttl=255 time=32 ms
    Reply from 10.0.13.3: bytes=56 Sequence=5 ttl=255 time=32 ms
  --- 10.0.13.3 ping statistics ---
    5 packet(s) transmitted
    5 packet(s) received
    0.00% packet loss
round-trip min/avg/max = 32/32/35 ms

Step 9 View the file list stored on the current device

Run the dir command in the user view to display the list of files in the current
directory.
<R1>dir
Directory of flash:/
  Idx Attr     Size(Byte)       Date     Time(LMT)      FileName
    0 -rw-       1,738,816      Mar 10 2016 11:50:24 web.zip
1 -rw- 68,288,896          Mar 10 2016 14:17:5          ar2220E-v200r007c00spc600.cc
2 -rw-                 739      Mar 10 2016 16:01:17         vrpcfg.zip
1,927,476 KB total (1,856,548 KB free)


<R3>dir
Directory of flash:/
  Idx Attr     Size(Byte)       Date         Time(LMT)        FileName
    0 -rw-     1,738,816 Mar 10 2016 11:50:58 web.zip
1 -rw-     68,288,896      Mar 10 2016 14:19:0          ar2220E-v200r007c00spc600.cc
2 -rw-                 739 Mar 10 2016 16:03:04        vrpcfg.zip
1,927,476 KB total (1,855,076 KB free)



Step 10        Manage device configuration files

Attempt to display the saved-configuration file.
<R1>display saved-configuration
     There is no correct configuration file in FLASH



Since no save-configuration file exists, save the current configuration file.
<R1>save
  The current configuration will be written to the device.



                                                  HUAWEI TECHNOLOGIES                  Page21
    Are you sure to continue? (y/n)[n]:y
    It will take several minutes to save configuration file, please wait............
    Configuration file had been saved successfully
    Note: The configuration file will take effect after being activated




Run the following command again to view the saved configuration information:
<R1>display saved-configuration
[V200R007C00SPC600]
#
sysname R1
header shell information "Welcome to Huawei certification lab"
#
board add 0/1 1SA
board add 0/2 1SA
……output omit……

Run the following command to view the current configuration information:
<R1>display current-configuration
[V200R007C00SPC600]
#
sysname R1
header shell information "Welcome to Huawei certification lab"
#
board add 0/1 1SA
board add 0/2 1SA
board add 0/3 2FE
……output omit……

A router can store multiple configuration files. Run the following command to view
the configuration file to currently be used after the next startup:
<R3>display startup
MainBoard:
    Startup system software:                              flash:/ar2220E-V200R007C00SPC600.cc
    Next startup system software:                         flash:/ar2220E-V200R007C00SPC600.cc
    Backup system software for next startup: null
    Startup saved-configuration file:                     null
    Next startup saved-configuration file:         flash:/vrpcfg.zip
    Startup license file:                                 null
    Next startup license file:                     null
    Startup patch package:                                null



                                                          HUAWEI TECHNOLOGIES                   Page22
  Next startup patch package:                     null
  Startup voice-files:                            null
  Next startup voice-files:                       null



Delete configuration files from the flash memory.
<R1>reset saved-configuration
This will delete the configuration in the flash memory.
The device configurations will be erased to reconfigure.
Are you sure? (y/n)[n]:y
 Clear the configuration in the device successfully.


