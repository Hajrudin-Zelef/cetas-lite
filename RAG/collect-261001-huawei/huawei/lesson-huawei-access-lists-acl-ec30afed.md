---
id: collect-261001-huawei/huawei/lesson-huawei-access-lists-acl-ec30afed
title: "lesson-huawei-access-lists-acl-ec30afed"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/lesson-huawei-access-lists-acl-ec30afed.md
source_anchor: ""
source_lines: [1, 39]
sha256: c03b45b08e2b8f002b3922b1f6f0b2bbad2674f8c0ec6c0f2afea0129425e82b
---

# lesson-huawei-access-lists-acl-ec30afed

Table of Contents
In this Huawei HCIA configuration example, we will learn Huawei Access List (ACL) Configuration on Huawei Routers. Configuring Access Lists has some basic steps and we will cover all these steps in this article.
First of all, let’s remember the Access List types. There are three types Huawei Access Lists (ACLs) :
You can check each ACL type and their number ranges below. You can also find the parameters of each ACL type.
You can also check Huawei Configuration Course for Hands on Experience on Huawei Routers and Swithces.
In the first example, we will focus on Basic Access List Configuration. With this Basic Access List Configuration, we will prevent the users on the subnet 10.1.1.0/24 and 10.2.2.0/24 to access the router. And we will allow the users, in the subnet 172.16.0.0/16 to access the router.
system-view
[Huawei-Router] acl 2222
[Huawei-Router-acl-basic-2222] rule deny source 10.1.1.0 0.0.0.255 
[Huawei-Router-acl-basic-2222] rule deny source 10.2.2.0 0.0.0.255
[Huawei-Router-acl-basic-2222] rule permit source 172.16.0.0 0.0.255.255
[Huawei-Router-acl-basic-2222] quit
And lastly, we will add this ACL to the related interface through outbound. Because, we are trying to prevent or allow the traffic going out to this router.
[Huawei-Router] interface GigabitEthernet1/1/1
[Huawei-Router-GigabitEthernet1/1/1] traffic-filter outbound acl 2222
[Huawei-Router-GigabitEthernet1/1/1] quit
To check the ACL, we can use “display acl 2222” command. As you can see below, you can find the three rules in this Access List (ACL).
In this second example, we will focus on Advanced Access List Configuration. With this Advanced Access List Configuration, we can prevent the traffic from a specific network to a specific network. We can allow a specific host to access to a specific server with a specific port. This can be reproducible.
For our example, we will configure an Advanced Access List in the right router. And we will apply this to the inbound direction of GigabitEthernet1/1/1.
We will allow R&D to FTP to the Server, but we will prevent Sales to access the Server. And we will allow Sales to communicate with Marketing and we will prevent R&D to access the Marketing.
Firstly, let’s create the Advanced ACL and write the rules under this ACL. We can use the numbers between 3000 and 3999. Here, in our ACL we will use 3500.
system-view
[Huawei-Router] acl 3500
[Huawei-Router-acl-basic-3500] rule deny source 10.10.5.0 0.0.0.255 destination 172.16.1.0 0.0.0.255
[Huawei-Router-acl-basic-3500] rule deny tcp source 10.10.6.0 0.0.0.255 destination 172.16.2.1 0.0.0.0 destination-port eq 21
[Huawei-Router-acl-basic-3500] rule permit tcp source 10.10.5.0 0.0.0.255 destination 172.16.2.1 0.0.0.0 destination-port eq 21
[Huawei-Router-acl-basic-3500] rule permit source 10.10.6.0 0.0.0.255 destination 172.16.1.0 0.0.0.255
[Huawei-Router-acl-basic-3500] quit
Now, let’s bind this Advanced ACL to the GigabitEthernet1/1/1 interface of the right router. We will add this ACL through inbound direction.
[Huawei-Router] interface GigabitEthernet1/1/1
[Huawei-Router-acl-basic-3500] traffic-filter inbound acl 3500
[Huawei-Router-acl-basic-3500] quit
Lastly, we will check the configured Advanced Access List with “display acl 3500” command.
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
How can access IPCISCO crew
You can have a membership and then you are in Nirati :)
