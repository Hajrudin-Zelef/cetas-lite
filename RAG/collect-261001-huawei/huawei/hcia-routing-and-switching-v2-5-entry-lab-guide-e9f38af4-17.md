---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-17
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet", "parameters", "preemption"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [3665, 3978]
sha256: 6092bc360113751ca6ad31f305589267f990a7e69bef3547ecab741010f1bd82
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

A DR or BDR is elected in non-preemption mode, by default. After router priorities
are changed, a DR is not re-elected, so you must reset the OSPF neighbor
relationship between R1 and R3.
Shut down and re-enable Gigabit Ethernet 0/0/0 interfaces on R1 and R3 to reset the
OSPF neighbor relationship between R1 and R3.
[R3]interface GigabitEthernet0/0/0
[R3-GigabitEthernet0/0/0]shutdown


[R1]interface GigabitEthernet0/0/0
[R1-GigabitEthernet0/0/0]shutdown




                                                      HUAWEI TECHNOLOGIES   Page80
[R1-GigabitEthernet0/0/0]undo shutdown


[R3-GigabitEthernet0/0/0]undo shutdown


Run the display ospf peer command to view the DR and BDR of R1 and R3.
[R1]display ospf peer 10.0.3.3


          OSPF Process 1 with Router ID 10.0.1.1
                   Neighbors


 Area 0.0.0.0 interface 10.0.13.1(GigabitEthernet0/0/0)'s neighbors
 Router ID: 10.0.3.3            Address: 10.0.13.3
    State: Full Mode:Nbr is Master       Priority: 100
    DR: 10.0.13.1 BDR: 10.0.13.3 MTU: 0
    Dead timer due in 52 sec
    Retrans timer interval: 5
    Neighbor is up for 00:00:25
    Authentication Sequence: [ 0 ]


According to the preceding information, R1's priority is higher than R3's priority, so
R1 becomes DR and R3 becomes the BDR.


Final Configuration

<R1>display current-configuration
[V200R007C00SPC600]
#
 sysname R1
#
interface GigabitEthernet0/0/0
 ip address 10.0.13.1 255.255.255.0
 ospf dr-priority 200
 ospf timer hello 15
#
interface GigabitEthernet0/0/1
 ip address 10.0.12.1 255.255.255.0
#
interface LoopBack0
 ip address 10.0.1.1 255.255.255.0
#
ospf 1 router-id 10.0.1.1



                                                     HUAWEI TECHNOLOGIES    Page81
 area 0.0.0.0
    network 10.0.1.0 0.0.0.255
    network 10.0.12.0 0.0.0.255
    network 10.0.13.0 0.0.0.255
#
user-interface con 0
 authentication-mode password
 set authentication password cipher %$%$+L'YR&IZt'4,)>-*#lH",}%K-oJ_M9+'lOU~bD (\WTqB}%N,%$%$
user-interface vty 0 4
#
return




<R2>display current-configuration
[V200R007C00SPC600]
#
 sysname R2
#
interface GigabitEthernet0/0/1
 ip address 10.0.12.2 255.255.255.0
#
interface LoopBack0
 ip address 10.0.2.2 255.255.255.0
#
ospf 1 router-id 10.0.2.2
 area 0.0.0.0
    network 10.0.2.0 0.0.0.255
    network 10.0.12.0 0.0.0.255
#
user-interface con 0
 authentication-mode password
 set authentication password cipher %$%$1=cd%b%/O%Id-8X:by1N,+s}'4wD6TvO<I|/pd# #44C@+s#,%$%$
user-interface vty 0 4
#
return




<R3>display current-configuration
[V200R007C00SPC600]



                                          HUAWEI TECHNOLOGIES                        Page82
#
 sysname R3
#
interface GigabitEthernet0/0/0
 ip address 10.0.13.3 255.255.255.0
 ospf dr-priority 100
 ospf timer hello 15
#
interface LoopBack0
 ip address 10.0.3.3 255.255.255.0
#
interface LoopBack2
 ip address 172.16.0.1 255.255.255.0
#
ospf 1 router-id 10.0.3.3
 default-route-advertise
 area 0.0.0.0
    network 10.0.3.0 0.0.0.255
    network 10.0.13.0 0.0.0.255
#
ip route-static 0.0.0.0 0.0.0.0 LoopBack2
#
user-interface con 0
 authentication-mode password
 set authentication password cipher %$%$ksXDMg7Ry6yUU:63:DQ),#/sQg"@*S\U#.s.bHW xQ,y%#/v,%$%$
user-interface vty 0 4
#
return




                                            HUAWEI TECHNOLOGIES                    Page83
                         Module 5 FTP and DHCP

Lab 5-1 Configuring FTP Services


Learning Objectives

As a result of this lab section, you should achieve the following tasks:

      Establishment of the FTP service.

      Configuration of FTP server parameters.

      Successful file transfer from an FTP server.


Topology




                                Figure 5.1 FTP topology



Scenario

As a network administrator of a company, you have been tasked with implementing
FTP services on the network. You need to implement the FTP service on a router
assigned to be an FTP server. The router should allow clients to successfully establish
a TCP session to the FTP application and transfer files.




                                      HUAWEI TECHNOLOGIES                   Page84
Tasks


Step 1 Preparing the environment.

If you are starting this section with a non-configured device, begin here and then
move to step 2. For those continuing from previous labs, begin at step 2.


<Huawei>system-view
Enter system view, return user view with Ctrl+Z.
[Huawei]sysname R1
[R1]interface GigabitEthernet 0/0/1
[R1-GigabitEthernet0/0/1]ip address 10.0.12.1 24


<Huawei>system-view
Enter system view, return user view with Ctrl+Z.
[Huawei]sysname R2
[R2]interface GigabitEthernet 0/0/1
[R2-GigabitEthernet0/0/1]ip address 10.0.12.2 24


Verify that R1 can reach R2, and vice versa..
 [R1]ping 10.0.12.2
  PING 10.0.12.2: 56 data bytes, press CTRL_C to break
    Reply from 10.0.12.2: bytes=56 Sequence=1 ttl=255 time=10 ms
    Reply from 10.0.12.2: bytes=56 Sequence=2 ttl=255 time=1 ms
    Reply from 10.0.12.2: bytes=56 Sequence=3 ttl=255 time=1 ms
    Reply from 10.0.12.2: bytes=56 Sequence=4 ttl=255 time=10 ms
    Reply from 10.0.12.2: bytes=56 Sequence=5 ttl=255 time=1 ms


  --- 10.0.12.2 ping statistics ---
    5 packet(s) transmitted
    5 packet(s) received
    0.00% packet loss
    round-trip min/avg/max = 1/4/10 ms



Step 2 Enable the FTP service on the router.

The FTP service is disabled by default on the router. It must be enabled before FTP
can be used. Configure an FTP server using R1 with R2 as the client. The same steps


                                                   HUAWEI TECHNOLOGIES    Page85
can be reversed to enable R2 to also act as an FTP server.
[R1]ftp server enable
Info: Succeeded in starting the FTP server
[R1]set default ftp-directory flash:/


Configure user authorization for FTP users to access the server. Unauthorized users
will not be able to access the FTP server, reducing security risks.
[R1]aaa
[R1-aaa]local-user huawei password cipher         huawei123
Info: Add a new user.

[R1-aaa]local-user huawei service-type ftp

Info: The cipher password has been changed to an irreversible-cipher password.

Warning: The user access modes include Telnet, FTP or HTTP, and so security risks exist.

Info: After you change the rights (including the password, access type, FTP directory, and level) of a local user, the
rights of users already online do not change. The change takes effect to users who go online after the change.

[R1-aaa]local-user huawei privilege level 15

Info: After you change the rights (including the password, access type, FTP directory, and level) of a local user, the
rights of users already online do not change. The change takes effect to users who go online after the change.

[R1-aaa]local-user huawei ftp-directory flash:

Info: After you change the rights (including the password, access type, FTP directory, and level) of a local user, the
rights of users already online do not change. The change takes effect to users who go online after the change.




[R1]display ftp-server

   FTP server is running

   Max user number                           5

   User count                            0

   Timeout value(in minute)             30

   Listening port                       21

   Acl number                             0

   FTP server's source address          0.0.0.0



The FTP server is running on R1 and listens on TCP port 21 by default.


                                                    HUAWEI TECHNOLOGIES                                Page86
Step 3 Establish an FTP client connection

Establish a connection to the FTP Server from R2.
<R2>ftp 10.0.12.1
Trying 10.0.12.1 ...
Press CTRL+K to abort
Connected to 10.0.12.1.
220 FTP service ready.
User(10.0.12.1:(none)):huawei
331 Password required for huawei.
Enter password:
230 User logged in.


[R2-ftp]


Following entry of the correct user name and password, the FTP server can be
successfully logged into.

