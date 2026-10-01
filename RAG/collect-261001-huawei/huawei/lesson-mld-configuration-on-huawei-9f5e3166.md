---
id: collect-261001-huawei/huawei/lesson-mld-configuration-on-huawei-9f5e3166
title: "lesson-mld-configuration-on-huawei-9f5e3166"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/lesson-mld-configuration-on-huawei-9f5e3166.md
source_anchor: ""
source_lines: [1, 27]
sha256: 81d173a34e2360ee7f18bba3b7d5119dae780fad9ebf919c3855582a3aaaae3f
---

# lesson-mld-configuration-on-huawei-9f5e3166

Table of Contents
In this lesson we will give an example topology consist of Huawei switches and multicast hosts and then we will configure MLD on these devices.
For our example, we will use the below topology:
Let’s configure MLD on the Switch 1.
Firstly, we will enalble IPv6 Multicast globally on the switch.
system-view [Switch1] multicast ipv6 routing-enable 
We will use PIM-SM with MLD. So, we will enable IPv6 PIM-SM.
[Switch1] pim ipv6 sm
We will enable MLD under the interface that we will use MLD with “mld enable” command.Before this, we will enalbe Layer 3 on switch with “undo portswitch”. We will do this configuration on each interface that we will use MLD.
[Switch1] interface 10GE 1/1/1
[Switch1-10GE1/1/1] undo portswitch
[Switch1-10GE1/1/1] mld enable
[Switch1-10GE1/1/1] quit
[Switch1] interface 10GE 1/1/2
[Switch1-10GE1/1/2] undo portswitch
[Switch1-10GE1/1/2] mld enable
[Switch1-10GE1/1/2] quit
[Switch1] interface 10GE 1/1/3
[Switch1-10GE1/1/3] undo portswitch
[Switch1-10GE1/1/3] mld enable
As we discussed before, MLD has twı versions; MLDv1 and MLDv2. We will set the version MLDv2 here. We will do this configuration on each interface that run MLD.
[Switch1-10GE1/1/1] mld version 2
[Switch1-10GE1/1/1] quit
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
