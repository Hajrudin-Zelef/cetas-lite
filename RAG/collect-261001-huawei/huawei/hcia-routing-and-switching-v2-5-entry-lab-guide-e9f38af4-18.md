---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-18
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [3979, 4223]
sha256: 1e71719f6919395d7f5c4bc400398e0be3a4f8f238d98cc1b0aa5717e6ce850a
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

Run the dir command before downloading a file or after uploading a file to view the
detailed information of the file.
[R2-ftp]dir
200 Port command okay.
150 Opening ASCII mode data connection for *.
drwxrwxrwx      1 noone    nogroup            0 May 03 18:03 .
-rwxrwxrwx      1 noone    nogroup 114552448 Jan 19 2012 AR2220E-V200R006C10SPC300.cc
-rwxrwxrwx      1 noone    nogroup      159858 May 03 17:59 mon_file.txt
-rwxrwxrwx      1 noone    nogroup      304700 Mar 03 11:11 sacrule.dat
-rwxrwxrwx      1 noone    nogroup         783 Mar 03 11:12 default_local.cer
-rwxrwxrwx      1 noone    nogroup           0 Dec 20 2015 brdxpon_snmp_cfg.efs
-rwxrwxrwx      1 noone    nogroup         777 May 03 18:03 vrpcfg.zip
drwxrwxrwx      1 noone    nogroup            0 Mar 10 11:14 update
drwxrwxrwx      1 noone    nogroup            0 May 03 18:03 localuser
drwxrwxrwx      1 noone    nogroup            0 Mar 17 10:45 dhcp
-rwxrwxrwx      1 noone    nogroup         460 May 03 18:03 private-data.txt
-rwxrwxrwx      1 noone    nogroup 126352896 Mar 10 11:09 AR2220E-V200R007C00SPC600.cc
drwxrwxrwx      1 noone    nogroup            0 Mar 10 11:15 shelldir
-rwxrwxrwx      1 noone    nogroup       11606 May 03 18:00 mon_lpu_file.txt
drwxrwxrwx      1 noone    nogroup            0 Mar 18 14:45 huawei
-rwxrwxrwx      1 noone    nogroup         120 Mar 18 15:02 text.txt226 Transfer complete.
FTP: 836 byte(s) received in 0.976 second(s) 856.55byte(s)/sec.




                                                  HUAWEI TECHNOLOGIES                        Page87
Set the transfer mode for the files to be transferred.


[R2-ftp]binary
200 Type set to I.



Retrieve a file from the FTP server. Note: If the vrpcfg.zip file is not present in the sd1:

directory of R1, use the save command on R1 to create it.


[R2-ftp]get vrpcfg.zip vrpnew.zip
200 Port command okay.
150 Opening BINARY mode data connection for vrpcfg.zip.
226 Transfer complete.
FTP: 120 byte(s) received in 0.678 second(s) 176.99byte(s)/sec.



After downloading the file from FTP server, use the bye command to close the

connection


[R2-ftp]bye
221 Server closing.


<R2>dir
Directory of flash:/


  Idx Attr       Size(Byte) Date          Time(LMT) FileName
    0 -rw-       114,552,448 Jan 19 2012 15:32:52     AR2220E-V200R006C10SPC300.cc
    1 -rw-             270,176 Apr 30 2016 03:17:08    mon_file.txt
    2 -rw-             304,700 Mar 03 2016 11:11:44    sacrule.dat
    3 -rw-                783 Mar 03 2016 11:12:22     default_local.cer
    4 -rw-                  0 Dec 20 2015 00:06:14     brdxpon_snmp_cfg.efs
    5 -rw-                775 Apr 29 2016 17:51:48     vrpcfg.zip
    6 drw-                   - Mar 10 2016 11:28:46    update
    7 drw-                   - Apr 23 2016 17:33:38   localuser
    8 drw-                   - Mar 21 2016 20:59:46    dhcp
    9 -rw-                394 Apr 29 2016 17:51:50     private-data.txt
   10 -rw-       126,352,896 Mar 10 2016 11:14:40      AR2220E-V200R007C00SPC600.cc
   11 drw-                   - Mar 10 2016 11:29:20    shelldir
   12 -rw-              23,950 Apr 27 2016 16:06:06    mon_lpu_file.txt



                                                  HUAWEI TECHNOLOGIES                 Page88
   13 -rw-                120 Mar 24 2016 11:45:44      huawei.zip
   14 -rw-                777 May 10 2016 14:23:43      vrpnew.zip
A file can be uploaded to the FTP server by using the command put, for which a new
file name can also be assigned.
[R2-ftp]put vrpnew.zip vrpnew2.zip
200 Port command okay.
150 Opening BINARY mode data connection for vrpnew2.zip.
226 Transfer complete.
FTP: 120 byte(s) sent in 0.443 second(s) 270.88byte(s)/sec.


After uploading the file, check for the presence of the file on FTP server.
<R1>dir
Directory of flash:/


  Idx Attr       Size(Byte) Date            Time(LMT) FileName
    0 -rw-             286,620 Mar 14 2016 09:22:20    sacrule.dat
    1 -rw-             512,000 Mar 28 2016 14:39:16    mon_file.txt
    2 -rw-         1,738,816 Mar 17 2016 12:05:36      web.zip
    3 -rw-             48,128 Mar 10 2016 14:16:56     ar2220E_v200r001sph001.pat
    4 -rw-                120 Mar 28 2016 10:09:50     iascfg.zip
    5 -rw-                699 Mar 28 2016 17:52:38     vrpcfg.zip
    6 -rw-        93,871,872 Mar 14 2016 09:13:26      ar2220-V200R007C00SPC600.cc
    7 -rw-             512,000 Mar 28 2016 14:40:20    mon_lpu_file.txt
    8 -rw-                699 Mar 02 2016 15:44:16     vrpnew2.zip



Remove the created vrpnew.zip and vrpnew2.zip files on R1 and R2.
<R1>delete flash:/vrpnew2.zip
Delete flash:/vrpnew2.zip? (y/n)[n]:y
Info: Deleting file flash:/vrpnew2.zip...succeed.


<R2>delete flash:/vrpnew.zip
Delete flash:/vrpnew.zip? (y/n)[n]:y
Info: Deleting file flash:/vrpnew.zip...succeed.


Note: Please take extreme care when deleting the configuration files so to ensure
that the entire flash:/ directory of R1 and R2 is not erased.




                                                    HUAWEI TECHNOLOGIES              Page89
Final Configuration

<R1>display current-configuration
[V200R007C00SPC600]
#
 sysname R1
 ftp server enable
 set default ftp-directory flash:
#
aaa
 authentication-scheme default
 authorization-scheme default
 accounting-scheme default
 domain default
 domain default_admin
 local-user admin password cipher %$%$=i~>Xp&aY+*2cEVcS-A23Uwe%$%$
 local-user admin service-type http
 local-user huawei password cipher %$%$f+~&ZkCn]NUX7m.t;tF9R48s%$%$
 local-user huawei privilege level 15
 local-user huawei ftp-directory flash:/
 local-user huawei service-type ftp
#
interface GigabitEthernet0/0/1
 ip address 10.0.12.1 255.255.255.0
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
 ftp server enable
 set default ftp-directory flash:
#
aaa
 authentication-scheme default



                                           HUAWEI TECHNOLOGIES                       Page90
 authorization-scheme default
 accounting-scheme default
 domain default
 domain default_admin
 local-user admin password cipher %$%$=i~>Xp&aY+*2cEVcS-A23Uwe%$%$
 local-user admin service-type http
 local-user huawei password cipher %$%$<;qM3D/O;ZLqy/"&6wEESdg$%$%$
 local-user huawei privilege level 15
 local-user huawei ftp-directory flash:/
 local-user huawei service-type ftp
#
interface GigabitEthernet0/0/1
 ip address 10.0.12.2 255.255.255.0
#
user-interface con 0
 authentication-mode password
 set authentication password cipher %$%$1=cd%b%/O%Id-8X:by1N,+s}'4wD6TvO<I|/pd# #44C@+s#,%$%$
user-interface vty 0 4
#
return




                                           HUAWEI TECHNOLOGIES                    Page91
Lab 5-2 Implementing DHCP


Learning Objectives

As a result of this lab section, you should achieve the following tasks:

      Configuration of a global DHCP pool.

      Configuration of an interface based DHCP pool.

      Enable DHCP discovery and IP allocation for switch interfaces.

      Method of global address pool configuration.

      Method of interface address pool configuration.


Topology




                               Figure 5.2 DHCP topology



Scenario

As the administrator of an enterprise you have been tasked with implementing
DHCP application services within the network. The gateway router in the company
network is to be configured as a DHCP server. IP addressing from an address pool
are to be offered by the gateway(s) (R1 and R3) to respective access layer devices.



                                      HUAWEI TECHNOLOGIES                  Page92
Tasks


Step 1 Preparing the environment

