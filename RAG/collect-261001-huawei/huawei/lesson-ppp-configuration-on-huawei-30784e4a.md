---
id: collect-261001-huawei/huawei/lesson-ppp-configuration-on-huawei-30784e4a
title: "lesson-ppp-configuration-on-huawei-30784e4a"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/lesson-ppp-configuration-on-huawei-30784e4a.md
source_anchor: ""
source_lines: [1, 38]
sha256: f8a9916d0b79c41b691968464c1c0e106b50553ce3151e8fd6aafa8cfeeee5d4
---

# lesson-ppp-configuration-on-huawei-30784e4a

In this lesson, we will talk about PPP configuration on Huawei Routers. For our PPP Configuration example, we will use the below basic topology:
Orderly, we will do the below configuration for our PPP Configuration Example:
1) Enabling PPP
2) PAP and CHAP Authentication
3) Verifying PPP and PPP Authentication
Let’s start to configure PPP for this topology.
You can download this configuration on Huawei eNSP Labs Page.
Table of Contents
In this basic PPP configuration example, firstly we will configure the serial interfaces to use PPP. After that, we will assign the ip address of this interface with its subnet mask. We will do this configuration on Router 1 firstly.
system-view
[Huawei-Router-1] interface Serial 0/0/0
[Huawei-Router-1-Serial0/0/0] link-protocol ppp
Warning : The encapsulation protocol of the link will be changed
Continue? [Y/N] : y
[Huawei-Router-1-Serial0/0/0] ip address 192.168.1.1 30
[Huawei-Router-1-Serial0/0/0] quit
Now, let’s configure Router 2 like above.
system-view
[Huawei-Router-2] interface Serial 0/0/0
[Huawei-Router-2-Serial0/0/0] link-protocol ppp
Warning : The encapsulation protocol of the link will be changed
Continue? [Y/N] : y
[Huawei-Router-2-Serial0/0/0] ip address 192.168.1.2 30
[Huawei-Router-2-Serial0/0/0] quit
We can configure each of these Authentication Protocols on Huawei Routers. Let’s start with PAP.
system-view
[Huawei-Router-1] aaa
[Huawei-Router-1-aaa] local-user abc password cipher abc123
[Huawei-Router-1-aaa] local-user abc service-type ppp
[Huawei-Router-1] interface Serial 0/0/0
[Huawei-Router-1-Serial0/0/0] link-protocol ppp
[Huawei-Router-1-Serial0/0/0] ppp authentication-mode pap
[Huawei-Router-1-Serial0/0/0] quit
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
