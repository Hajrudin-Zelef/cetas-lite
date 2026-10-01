---
id: collect-261001-general-networking/general-networking/static-route-redistribution-connected-networks-redistribution
title: "static-route-redistribution-connected-networks-redistribution"
domain: general-networking
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/static-route-redistribution-connected-networks-redistribution.md
source_anchor: ""
source_lines: [1, 87]
sha256: 5ec90330fc8a4e7bb8e0a00293fd333f8c5f298a09bea419a6484b5597de96cb
---

# static-route-redistribution-connected-networks-redistribution

There are static routes in a network and directly connected networks. To redistribute static routes and directly connected networks, we use different **Cisco redistribution commands**. Here, we will give an example for both **Static Route Redistribution** and **Connected Network Redistribution**.


As our redistribution topology, we can use the below topology.



You can also visit **OSPF into EIGRP Redistribution** lesson.



Table of Contents

Static Routes are one of the most used routes in a network generally. To **redistribute static routes** into a routing protocol, we use “**redistribute static**” command. We can use this command both without a metric and with a metric.


If we do not use a metric value, we can use only “**redistribute static**” command. Let’s talk about redistribution of static routes into EIGRP. To do this, firstly, we will entetr EIGRP configuration mode with a process number.


Router B(config)# **router eigrp 1**

Router B(config-router)# **redistribute static**




If you want to use a specific metric for static redistributed routes, you can use metric keyword with a metric value.


And a keyword, “**subnets**”, is also used to redistribute static routes into EIGRP to redistribute also the subnets. If we do not use this keyword, this command works for classful network.


Router B(config)# **router eigrp 1**

Router B(config-router)# **redistribute static metric 200 subnets**




Like static routes, directly connected routes are also redistributed with a simple command. To redistribute connected networks into a routing protocol, we use “**redistribute connected**” command.


Again, we can use “**redistribute connected**” command without any metric value or with a metric value. Let’s think about that, we w.ill redistribute directly connected networks into EIGRP. Firstly, let’s do not use any metric value.


Router B(config)# **router eigrp 1**

Router B(config-router)# **redistribute connected**



Now, let’s configure this redistribution with a metric value, 100.


Router B(config)# **router eigrp 1**

Router B(config-router)# **redistribute connected metric 100**




As you can see here, **redistributing connected networks** and **redistribution static routes** are very easy in Cisco world. You will also use these commands while you redistribute from one routing protocol to another routing protocol. Because, in a network there are always such routes.


**Below, you can find important configuration lessons and lab pages:**

**DOWNLOAD Cisco Packet Tracer Labs:** https://ipcisco.com/cisco-packet-tracer-configuration-examples/

**DOWNLOAD Cisco GNS3 Labs:** https://ipcisco.com/cisco-gns3-configurations/

**Join Cisco Packet Tracer Labs Course:** https://ipcisco.com/course/cisco-packet-tracer-course/

**Join Cisco CCNA Labs Couse:** https://ipcisco.com/course/ccna-200-301-labs/

**Join Cisco CCNP ENCOR Labs Course:** https://ipcisco.com/course/ccnp-350-401-encor-labs/



Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
