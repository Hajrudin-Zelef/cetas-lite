---
id: collect-261001-huawei/huawei/lesson-nat-configuration-on-huawei-routers-932971f0
title: "lesson-nat-configuration-on-huawei-routers-932971f0"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/lesson-nat-configuration-on-huawei-routers-932971f0.md
source_anchor: ""
source_lines: [1, 51]
sha256: c78ff9adf62d838d82359e1a1bd738bae6f7d231078bdc0d7b7286bbaae3d898
---

# lesson-nat-configuration-on-huawei-routers-932971f0

Table of Contents
NAT (Network Address Translation) is a concept used to translate Private block IP addresses to the Public IP Addresses. By doing this, it provides internet connection to the devices that has Private Blcok IP Addresses. In this lesson we will learn Huawei NAT Configuration.
Here, NAT is a general used name. There are different types of NAT. These types are :
Static NAT maps One-to-One, one Private IP to a one Public IP. This is sued to protect the device inside the network generally. But it is not an effective IP usage.
Dynamic NAT maps Many-to-Many. Devices in the network has Private IPs and these IPs can map to the Public IPs in a predefined IP Pool. Here, the Public IPs are used if they are needed, so it is an efficient solution.
PAT provide mapping with different Port Numbers. Devices in the network can be mapped to one Public IP, with their Private IP addresses and a different Port number. This is a very efficient way.
Now, let’s see the Huawei configuration of these NAT types one by one.
You can download this configuration on Huawei eNSP Labs Page.
In this first NAT configuration example, we will focus on Static NAT Configuration on Huawei Routers. As we have explained above, Static NAT is configured as one to one. We will translate one Private IP Address to one Public IP Address.
For our Huawei Static NAT Configuration Example, we will use the below topology.
Let’s start to configure.
Firtstly we wil configure the interface IP addresses on the Router. Then, under the wan interface of the router, we will ad the nat commadn, “nat static global Public-X indise Private-Y”.
system-view
[Huawei-Router] interface GigabitEthernet0/0/0
[Huawei-Router-GigabitEthernet0/0/0] ip address 10.10.10.1 24
[Huawei-Router-GigabitEthernet0/0/0] quit
[Huawei-Router] interface GigabitEthernet1/1/1
[Huawei-Router-GigabitEthernet1/1/1] ip address 200.222.222.1 24
[Huawei-Router-GigabitEthernet1/1/1] nat static global 200.222.222.5 inside 10.10.10.5
[Huawei-Router-GigabitEthernet1/1/1] quit
To verify our Static NAT Configuration, we will use “display nat statistics” command.
The second NAT Configuration example is Dynamic NAT Configuration. AS we exmplained before, Dynamic NAT uses a predefined IP Pool and provide IP map according to these IP Pool if there is a need.
For our Dynamic NAT Configuration Example, we will use the below topology.
Firstly we wil configure the interface IP addresses on the Router. After that, we will create an access list. In this Acess List, we will add the LAN IPs. Then, we will configure the IP Pool with “nat address-group” command. And then, we will configure NAT under the WAN interface of the router with the defined IP Pool and ACL. This ACL will allow our LAN IPs on the router.
Firstly, IP address configurations will be done.
system-view
[Huawei-Router] interface GigabitEthernet0/0/0
[Huawei-Router-GigabitEthernet0/0/0] ip address 10.10.10.1 24
[Huawei-Router-GigabitEthernet0/0/0] quit
[Huawei-Router] interface GigabitEthernet1/1/1
[Huawei-Router-GigabitEthernet1/1/1] ip address 200.200.200.1 24
[Huawei-Router-GigabitEthernet1/1/1] quit
Now, it is time to create Access List.
[Huawei-Router] acl number 2222
[Huawei-Router-acl-basic-2222] rule 5 permit source 10.10.10.0 0.0.0.255
[Huawei-Router-acl-basic-2222] quit
After ACL definition, let’s configure IP Address Pool globally.
[Huawei-Router] nat address-group 1 200.200.200.5 200.200.200 10
Lastly, we will configure NAT on the WAN interface of the router. We will add IP Pool and ACL also. With this configuration, we will say that, “use these IP Pool for mapping the IPs in the ACL (LAN IPs)”.
[Huawei-Router] interface GigabitEthernet1/1/1
[Huawei-Router-GigabitEthernet1/1/1] nat outbound 2222 address-group 1 no-pat
[Huawei-Router-GigabitEthernet1/1/1] quit
To verify our Dynamic NAT Configuration, we will use “display nat address-group” command and “display nat outbound” command. You can also check the results of these commands below.
By the way, in this article we will not show the configuration of PAT (Port Address Translation) Because only a key word effect this. If you do not use “no-pat” command at the above configuration, you can configure PAT in the same configuration.
In Cisco we use, “overload” command for PAT. If you do not use this keyword then, it is becoming pure Dynamic NAT. This is reverse in Huawei.
In this lesson, we have focused on Huawei NAT Configuration. It will be good for you to try these configurations on your own PC with Huawei eNSP. Because with more practice, you will have a good hands on experience on these configurations. And these configurations are very important on Huawei HCIA Exam.
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
