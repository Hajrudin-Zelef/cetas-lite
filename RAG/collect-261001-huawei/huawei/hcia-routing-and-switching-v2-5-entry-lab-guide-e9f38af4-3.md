---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-3
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2016-03-11"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [473, 713]
sha256: 2b640b0793975f6e5af29928e1565a1a408c0928d65429db08ee402eb612fa94
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

The Wireshark application contains many functions for management of the packet
capture process. One of the more common functions includes the filter function to
isolate the packet capture display to a select group of packets or protocols. This can
be achieved using the filter field below the menu bar. The simplest filter method
involves entering the protocol name (in lower case) and pressing Enter. In the given
example packets for two protocols have been captured, entering either icmp, or arp
into the filter window will result in only the protocol entered in the filter field being
displayed in the output.
The packet capture tool consists of three panels, to show the list of packets, a
breakdown of the content of each packet and finally display the equivalent data
format of the packet. The breakdown is invaluable for understanding the format of
protocol packets and displays the details for protocols as referenced at each layer of
the OSI reference model.



                                       HUAWEI TECHNOLOGIES                    Page13
       Module 2 Basic Device Navigation and Configuration

Lab 2-1 Basic Device Navigation and Configuration


Learning Objectives

As a result of this lab section, you should achieve the following tasks:

       Configure device system parameters including device name, the system time,

        and the system time zone.

       Configure the console port idle timeout duration.

       Configure the login information.

       Configure the login password.

       Save configuration files.

       Configure IP addresses for router interfaces.

       Test the connectivity between two directly connected routers.

       Restart a device using VRP.


Topology




               Figure 2.1 Lab topology for basic VRP navigation and operation.




                                        HUAWEI TECHNOLOGIES                      Page14
Scenario

A company has purchased two AR G3 routers that require commissioning before
they can be used in the enterprise network. Items to be commissioned include
setting device names, the system time, and password management.




Tasks


Step 1 View the system information

Run the display version command to view the software version and hardware
information for the system.
<Huawei>display version
Huawei Versatile Routing Platform Software
VRP (R) software, Version 5.160 (AR2200 V200R007C00SPC600)
Copyright (C) 2011-2013 HUAWEI TECH CO., LTD
Huawei AR2220E Router uptime is 0 week, 3 days, 21 hours, 43 minutes
BKP 0 version information:
......output omitted......


The command output includes the VRP operating system version, device model, and
startup time.

Step 2 Change the system time parameter

The system automatically saves the time. If the time is incorrect, run the clock
timezone and clock datetime commands in the user view to change the system
time.
<Huawei>clock timezone Local add 08:00:00
<Huawei>clock datetime 12:00:00 2016-03-11



The keyword Local can be exchanged with the current regional timezone name, and

add replaced with minus where the timezone is west of UTC+0.


Run the display clock command to check that the new system time has taken effect.


                                               HUAWEI TECHNOLOGIES       Page15
<Huawei>display clock
  2016-03-11 12:00:10
  Friday
  Time Zone(Local) : UTC+08:00




Step 3 Help features & Auto-completion functions

The question mark (?) is a wildcard, and the Tab is used as a shortcut to enter
commands.
<Huawei>display ?
    Cellular               Cellular interface
    aaa                          AAA
    access-user                  User access
    accounting-scheme            Accounting scheme
    acl                          <Group> acl command group
    actual                       Current actual
    adp-ipv4               Ipv4 information
    adp-mpls               Adp-mpls module
    alarm                        Alarm
    antenna                      Current antenna that outputting radio
    anti-attack                  Specify anti-attack configurations
    ap                           <Group> ap command group
    ap-auth-mode           Display AP authentication mode
 ......output omit......


To display all the commands that start with a specific letter or string of letters, enter
the desired letters and the question mark (?). The system displays all the commands
that start with the letters entered. For example, if the string dis? is entered, the
system displays all the commands that start with dis.

If a space exists between the character string and the question mark (?), the system
will identify the commands corresponding to the string and display the parameters
of the command. For example, if the string dis ? is entered and only the display
command matches the dis string, the system displays the parameters of the display
command. If multiple commands start with dis, the system displays an error.
The Tab key can also be pressed to complete a command. For example, if dis is
entered followed by Tab, the system completes the display command. If multiple
commands start with dis, the appropriate command can be selected.
If there are no other commands starting with the same letters, dis or disp can be


                                                  HUAWEI TECHNOLOGIES         Page16
entered to indicate display, and int or inter to indicate interface.



Step 4 Access the system view

Run the system-view command to access the system view to configure interfaces
and protocols.
<Huawei>system-view
Enter system view, return user view with Ctrl+Z.
[Huawei]



Step 5 Change device names

To more easily identify devices, set device names during the device configuration.
Change device names based on the lab topology, as shown below:
Change the name of the R1 router to R1.
[Huawei]sysname R1
[R1]


Change the name of the R3 router to R3.

[Huawei]sysname R3
[R3]



Step 6 Configure the login information

Configure the login information to indicate the login result.
[R1]header shell information "Welcome to the Huawei certification lab."



Run the preceding command to configure the login information. To check whether
the login information has been changed, exit from the router command line
interface, and log back in to view the login information.
[R1]quit
<R1>quit


  Configuration console exit, please press any key to log on




                                                   HUAWEI TECHNOLOGIES    Page17
Welcome to the Huawei certification lab.
<R1>



Step 7 Configure console port parameters

The console port by default does not have a login password. Users must configure a
password for the console port before logging in to the device.
The password can be changed in the password authentication mode to huawei in
plain text.
If there is no activity on the console port for the period of time specified by the
timeout interval, the system will automatically log out the user. When this occurs, log
in to the system again using the configured password.
The default timeout interval is set to 10 minutes. If a 10 minutes idle period is not a
reasonable amount of time for the timeout interval, change the timeout interval to a
more suitable duration, here this is set to 20 minutes.
[R1]user-interface console 0
[R1-ui-console0]authentication-mode password
[R1-ui-console0]set authentication password cipher
Warning: The "password" authentication mode is not secure, and it is strongly recommended to use "aaa"
authentication mode.
Enter Password(<8-128>):
Confirm password:
[R1-ui-console0] idle-timeout 20 0


Run the display this command to check the configuration results.
[R1-ui-console0]display this
[V200R007C00SPC600]
#
user-interface con 0
 authentication-mode password
 set authentication password cipher %$%$fIn'6>NZ6*~as(#J:WU%,#72Uy8cVlN^NXkT51E ^RX;>#75,%$%$
 idle-timeout 20 0


