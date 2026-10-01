---
id: collect-261001-huawei/huawei/lesson-basic-huawei-configuration-d990caeb
title: "lesson-basic-huawei-configuration-d990caeb"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["training"]
source: docs/RAG/collect-261001-huawei/lesson-basic-huawei-configuration-d990caeb.md
source_anchor: ""
source_lines: [1, 62]
sha256: 19b14f728556afa33718e66181735b7ede07cce22e464e6cdf7fe8d86c837521
---

# lesson-basic-huawei-configuration-d990caeb

In this Basic Huawei Configuration lesson, we will focus on basic Huawei router configuration. We will configure two Huawei Routers on Huawei eNSP with the basic Huawei router commands. This lesson is a key lesson of Huawei HCIA Course and it is importan for Huawei HCIA-Datacom Exam. You can also check other Huawei Configurations for HCIA Exam on Huawei Configuration Course.
Table of Contents
In Huawei routers and switches, to enter the configuration mode, we use “system-view” command. This is basically Cisco’s “enable” command. After “system-view” command, we can do any configuration on a Huawei device.
To change a Huawei router name, we use “sysname” command with the new router name. At the beginning, router name is Huawei. Let’s change routers names as “Router1” and “Router2”. To do this, we will use “sysname Huawei-Training” command. Here, while you are giving a new name to your device, be careful not to add any blank to your router name.
[Huawei] sysname Router1
[Router1]
On router 2, we can use the below name.
[Huawei] sysname Router2
[Router2]
Telnet is one of the important protocols that we use in networking. In a router, telnet configuration is also important. Because, with telnet configuration, we define how to access to this device remotely.
By default, there are 5 sessions (0 to 4), to increase it to the maximum value, 15 sessions (0 to 14), we use “user-interface maximum-vty 15” command. After this command, telnet will allow 15 users.
[Router1] user-interface maximum-vty 15
After that, we can go to these sessions and start tom configure them with “user-interface vty 0 14” command. When we enter this, configuration level will be changed to vty. Under this configuration level, firstly, we will enable telnet with “protocol inbound telnet” command. Then, we will set privilege level of the sessions. Here, I will use the lower value, 1.
[Router1] user-interface vty 0 14
[Router1-ui-vty0] protocol inbound telnet
[Router1-ui-vty0] user privilege level 1
You can also watch the configuration of this scenario on our youtube channel.
Now, it is time to configure authentication method of telnet. Our authentication method will be password. To configure it, we will use “authentication-mode password” command firstly. Then, we will define our authentication password with “set authentication password cipher hw123” command. Here, “hw123” is our password. You can sue anything instead of it.
[Router1-ui-vty0] authentication-mode password
[Router1-ui-vty0] set authentication password cipher hw123
We can also do some detailed configurations on telnet sessions. We can change the default timeout value in minutes with “idle-timeout” command. The default timeout value is 10 minutes.
[Router1-ui-vty0] idle-timeout 30
We can also set the number of lines displayed on the terminal. To do this, I will use “screen-length 0” command. The default value is 24. When we set it to 0, this means that, all command output will be displayed on the screen.
[Router1-ui-vty0] screen-length 0
We can set saved command numbers in buffer with “history-command max-size” command.
[Huawei-Training-ui-vty0] history-command max-size 10
Console configuration is also important for us. Because with this configuration, we set how to access a device with console.
To enter console configuration, we use “user-interface console 0” command.
[Router1] user-interface console 0
To determine the user privilege with highest level, we use “user privilege level 15” command. With this command, we set our user level to the highest level.
[Router1-ui-vty0] user privilege level 15
We will determine the authentication mode for console too. To do this, we will use “authentication-mode password” command to set the authentication mode to password. We will set authentication password as hw456 with “set authentication password cipher hw456” command.
[Router1-ui-vty0] authentication-mode password
[Router1-ui-console0] set authentication password cipher hw456
To configure a Huawei router interface, firstly, we go to that interface level. Here, we will configure GigabitEthernet 0/0/0. So, to go to under GigabitEthernet 0/0/0, we will use “interface GigabitEthernet 0/0/0” command. Then, we will enter interface ip address, 192.168.0.1/24 with “ip address 192.168.0.1 255.255.255.0” command. Here, we use the long form of subnet mask. Lastly, to open the interface, we use “undo shutdown” command. Because, by default, router ports are down.
[Router1] interface GigabitEthernet 0/0/0
[Router1-GigabitEthernet0/0/0] ip address 192.168.0.1 255.255.255.0
[Router1-GigabitEthernet0/0/0] undo shutdown
We can also do the ip configuration of the second router like below:
[Router2] interface GigabitEthernet 0/0/0
[Router2-GigabitEthernet0/0/0] ip address 192.168.0.2 255.255.255.0
[Router2-GigabitEthernet0/0/0] undo shutdown
After this configuration, to check the interfaces’ status, we can use “display ip interface brief” command. We can also display the details of this interface with “display ip interface gigabit 0/0/0” command.
To configure a shell message, we use “header shell information” command. After this command we enter our message. It is important to enter a letter or sign firstly and finish with the same letter or sign. Here, our message will be Good luck!
[Router1] header shell information “Good Luck!”
We can do this also like this: We can write our sign and then press enter. After entering our message, we finish with the same letter.
To save Huawei router configuration, we use “set save-configuration” command. This command is important not to loose our configuration after any restart.
[Router1] set save-configuration
To verify router configuration, we can use some display commands on Huawei routers. For example to display all the current configuration on the router, we use “display current-configuration” command.
[Router1] display current-configuration
We can specifically display any part of the configuration with “|” sign. For example, let’s display the configuration line beginning with “user” in the current Huawei router configuration.
[Router1] display current-configuration | include user
To verify interfaces, we mostly use “display ip interface brief” especially during any real operations. This is one of the most used commands by a network engineer.
[Router1] display ip interface brief
If you would like to display the details of any interface, you can use “display ip interface gigabit 0/0/0” command.
[Router1] display ip interface gigabit 0/0/0
In this Basic Huawei Configuration with eNSP lesson, we have learned how to do basic Huawei router configuration. These configurations will be very important especially on your jobs in which you use Huawei devices. Especially on Huawei HCIA-Datacom Exam and for Huawei HCIA Course, this lesson will be very important.
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
