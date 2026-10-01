---
id: collect-261001-huawei/huawei/lesson-huawei-ftp-configuration-94d40bf5
title: "lesson-huawei-ftp-configuration-94d40bf5"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/lesson-huawei-ftp-configuration-94d40bf5.md
source_anchor: ""
source_lines: [1, 31]
sha256: 6c800ac5a54b4087dd9caf288d2033036d3b8f9a5fc5c7f2eb9a80f45177b00e
---

# lesson-huawei-ftp-configuration-94d40bf5

In this Huawei FTP Configuration example, we will configure a router as FTP Server and then we will try to connect this FTP Server router from another router. Here, this another router will be our FTP Client. By the way FTP is the abbreviation for File Transfer Protocol and it uses TCP 20 and TCP 21 Ports by default.
For our Huawei FTP Configuration example, we will use the below topology.
Let’s start to the configuration.
Firstly, we will configure IP addresses of each interfaces of these two routers.
<Huawei-Router1> system-view
[Huawei-Router1] interface GigabitEthernet0/0/0
[Huawei-Router1-GigabitEthernet0/0/0] ip address 10.1.1.1 24
[Huawei-Router1-GigabitEthernet0/0/0] undo shutdown
[Huawei-Router1-GigabitEthernet0/0/0] quit
<Huawei-Router2> system-view
[Huawei-Router2] interface GigabitEthernet0/0/0
[Huawei-Router2-GigabitEthernet0/0/0] ip address 10.1.1.2 24
[Huawei-Router2-GigabitEthernet0/0/0] undo shutdown
[Huawei-Router2-GigabitEthernet0/0/0] quit
After interface configuration, let’s configure Router 2 as FTP Server.
FTP Server configuration starts firstly, by enabling “FTP Server Mode”. After enabling FTP Server Mode, default FTP Directory is set.
[Huawei-Router2] ftp server enable
[Huawei-Router2] set default ftp-directory sdl:
Now, let’s configure FTP user related parameters of FTP Server.
We will enter the “aaa”, and there, we will create FTP user with its username and password.
[Huawei-Router] aaa
[Huawei-Router-aaa] local-user gokhan password cipher abc123
[Huawei-Router-aaa] local-user gokhan service-type ftp
Then, we will also define the FTP Directory location (here we will select flash:/).
[Huawei-Router-aaa] local-user gokhan ftp-diretory flash:/
And we will define, “access-limit” to this user with “timeout” values also.
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
