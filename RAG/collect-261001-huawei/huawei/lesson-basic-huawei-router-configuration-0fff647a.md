---
id: collect-261001-huawei/huawei/lesson-basic-huawei-router-configuration-0fff647a
title: "lesson-basic-huawei-router-configuration-0fff647a"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2019-12-04"]
keywords: ["ethernet", "training"]
source: docs/RAG/collect-261001-huawei/lesson-basic-huawei-router-configuration-0fff647a.md
source_anchor: ""
source_lines: [1, 86]
sha256: 42f0a2d95b789be8ab679600971a7ee15be6a84b4cccf22091c86185911e04e3
---

# lesson-basic-huawei-router-configuration-0fff647a

Table of Contents
In this Huawei Basic Router Configuration lesson, we will talk about the Basic Huawei Router Configuration Commands. These commands are basic but very improtant commands for your Huawei hands on experience and for Huawei HCIA Certification. You can also check the other Huawei Configuration Lessons.
For our configuration example, we will use a single router.
On Huawei routers and switches, you can use question mark (?) to help about configuration commands. This help can be both to complete a command or for the available all commands under that hierarchy.
If you use “?” alone, then available all commands will be showed at that level. Here, we list all the commands beginning with “display dhcp”.
<Huawei-Router> display dhcp ?
client      DHCP client
option82    DHCP option82
relay       DHCP relay
server      DHCP server
snooping    Snooping
static      Static bind item
statistics  Statistics
 
If you use “?”, with a letter, the commands that begins with that letter will be showed. Below, after “display dhcp” command, we list the commands beginning with “s”.
<Huawei-Router> display dhcp s?
server                                  snooping
static                                  statistics
On every router, there are different view modes. On Huawei routers there are four different viw modes. These are:
User View is the mode that you will enter after login to the router. With this mode, you can check the system configurations.
After that to configure a Huawei router, you should enter System View Mode. To enter this configuration mode, we use “system-view” command. This is the equivalent of Cisco’s “configure terminal” command.
<Huawei-Router> system-view
Enter system view, return user view with Ctrl+Z.
[Huawei-Router]
The other two modes are Interface view and Protocol view. In these modes, router interfaces and routing protocol configurations are done.
There are various shortcuts on Huawei CLI. These shortcuts is useful during router configuration. Some of the important shortcuts are given below:
CTRL – A : To go to the beginning of the line.
CTRL – C : Stops current function
CTRL – Z : Back to user view
TAB : Completes incomplete commands
On Huawei Routers, to change router name, we use “sysname” command.
[Huawei]sysname Huawei-Training
[Huawei-Training]
On all Routers, there is a clock and timezone settings. We can adjust this with various clock commands on Huawei Routers.
To set the timezoze, we will use “clock timezone” command. Here, with “add” keyword, we determine UTC+ timezone and with “mines”, we determine UTC- timezone. Here, we did this configuration for UTC+3.
<Huawei-Router> clock timezone 1 add 3:00:00
After setting time zone, we can set the current time and date with “clock datetime” command.
<Huawei-Router> clock datetime 15:45:10 2019-12-04
In some countries, daylight savingm time is used. So, summer and winter times change. To adjust this on Huawei routers, we us “clock daylight-saving-time” command.
Lastly, you can check the set time and clock with “display clock” command.
<Huawei-Router> display clock
2019-12-04 15:45:14+3:00
Wednesday
Time Zone(1) : UTC+3:00
There are two CLI Message options on Huawei routers. These CLI messages are :
Header Login message is the message that you will see after the router boots. It is configured with “header login information” command.
[Huawei-Router] header login information “Welcome!”
Header Shell message is the message that you will see after login to the router. It is configured with “header shell information” command.
[Huawei-Router] header shell  information “Be Careful!”
Welcome!
Login Authentication
Password:
Be Careful!
On Routers, there are different user privilege levels. Some of these levels can only view authorization, some of them has configuration authorization and others has different types of authorizations. You can find the levels of Huawei routers below:
We can set user levels with “command-privilege level” command like below. Here, we have set user level to 10.
[Huawei-Router] command-privilege level 10 view user save
To configure telnet sessions setting on Hauwei routers, firstly we should enter the telnet configuration level. By default 5 telnet sessions can be configured from 0 to 4. Here, we can set also the user privilege level and user password. Fort his configuration, we will only give monitoring permission to the thelnet users. And this permission is Level 1. Our password will be “abc123”.
[Huawei-Router] line vty 0 4
[Huawei-Router-ui-vty0-4] user privilege level 1
[Huawei-Router-ui-vty0-4] authentication-mode password
[Huawei-Router-ui-vty0-4] set authentication password cipher abc123
For console configuration, we will do the similar configurations with Telnet configurations. Here, by default there is one Console session configuration. We will set our privilege level as level 15 for console connection. And our password will be “abc123”.
[Huawei-Router] line console 0
[Huawei-Router-ui-console0] user privilege level 15
[Huawei-Router-ui-console0] authentication-mode password
[Huawei-Router-ui-console0] set authentication password cipher abc123
On Huawei routers there are also some CLI setting. By setting these values, CLI screen length, Session Timeout and Command History size are configured.
To set CLI screen size, we use “screen-length” command. We can set the displayed lines after the command excution. Below, we set this value for 5 lines.
[Huawei-Router] screen-length 5
To set session timeout value, we use “idle-timeout” command. Below we set this for 1 hour.
[Huawei-Router] idle- timeout 1 60
For command history size, we use “history-command max-size” command. Below, we set to store last 30 commands.
[Huawei-Router] history-command max-size 30
On Huawei routers, to configıre interfaces, firstly we should go to that interface level. Below, we will configure interface Gigabit Ethernet 1/1/1 with ip address 192.168.0.1 and subnet mask 255.255.255.0. After the ip configuration, we will open the interface.
[Huawei-Router] interface GigabitEthernet 0/0/0
[Huawei-Router-GigabitEthernet0/0/0] ip address 192.168.0.1 255.255.255.0
[Huawei-Router-GigabitEthernet0/0/0] undo shutdown
Here, subnet mask can be also written as “24”.
[Huawei-Router-GigabitEthernet0/0/0] ip address 192.168.0.1 24
In this Basic Huawei Router Configuration lesson we have learned basic commands of Huawei CLI. In other words we have learned basics of huawei router configuration. If you would like to learn more on Huawei configurations and Huawei CLI, you can visit other configuration pages like NAT, DHCP, OSPF configurations.
To have Huawei hands on experience, you should focus on configurations more. After learning how to configure a Huawei router, you should practice these configurations on your own PC. For these practices Huawei eNSP is useful. After practicing on eNSP, you will be ready for HCIA exam. On this way, basic Huawei Router Configuration lesson is the beginning lesson for you.
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
