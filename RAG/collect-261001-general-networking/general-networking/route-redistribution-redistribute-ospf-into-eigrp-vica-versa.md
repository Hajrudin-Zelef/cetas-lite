---
id: collect-261001-general-networking/general-networking/route-redistribution-redistribute-ospf-into-eigrp-vica-versa
title: "route-redistribution-redistribute-ospf-into-eigrp-vica-versa"
domain: general-networking
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/route-redistribution-redistribute-ospf-into-eigrp-vica-versa.md
source_anchor: ""
source_lines: [1, 116]
sha256: bb034c91f0ffca0de627901fc0884fda0a7a8fd6cb95479f093aba8d74fd91b9
---

# route-redistribution-redistribute-ospf-into-eigrp-vica-versa

Different **routing protocols** can be used in different networks. In other words, sometimes a routing protocol is used in one part of a network and another one is used in the other part of the network. In such cases, we use **Route Redistribution**. In this example, we will use a topology that is configured with both OSPF and EIGRP. We will **redistribute** **OSPF into EIGRP** and we will **redistribute EIGRP into OSPF**. In other words, we will configure **Cisco** **OSPF into EIGRP Redistribution** and **Cisco** **EIGRP into OSPF Redistribution**.


**To learn the details of OSPF and EIGRP, you can visit OSPF Overview Lesson and EIGRP Overview Lesson.** 


Here, our **Cisco route redistribution** topology will be like below:



Now, let’s go to our configuration and start with **Redistribute OSPF into EIGRP.**



Table of Contents

To **Redistribute OSPF into EIGRP,** we will do the below configuration on the border router. Firstly, we will enter eigrp configuration mode with eigrp process number 1.

Here, firstly, we will redistribute the static routes with “**redistribute static**” command and directly connected networks with “**redistribute connected**” command.  The first command to **Redistribute OSPF into EIGRP** is “**redistribute ospf 1 metric 1500 10 255 1 1500**” command. Here, 1 is OSPF process number. The numbers that we use in this command is orderly, **bandwidth, delay, reliability, load** and **MTU.** 


Router B(config)# **router eigrp 1**

Router B(config-router)# **redistribute static**

Router B(config-router)# **redistribute connected**

Router B(config-router)# **redistribute ospf 1 metric 1500 10 255 1 1500**

Router B(config-router)# **end**

Router B# **write** 



After Redistribute OSPF into EIGRP configuration, it is time to configure **Redistribute EIGRP into OSPF**.



After the process of **Redistribute OSPF into EIGRP,** now, it is time to **Redistribute EIGRP into OSPF.** Here, firstly, we will enter ospf configuration with the process id 1. Then, we will enter Cisco redistribution commands to **Redistribute EIGRP into OSPF.** 

Here, firstly, we will redistribute the static routes again but this time we will use “**redistribute static metric 200 subnets**” command. We will set **200** as metric for static routes that will be redistributed, and we will also add **subnets** keyword. Again, we will use “**redistribute connected metric 100**” command to redistribute directly connected networks.


The first redistribution command will be “**redistribute eigrp 1 metric 100 subnets**” for **EIGRP into OSPF redistribution.** This command will set 100 to all the routes. Here, **subnets** keyword is important. Because if we do not use this command, only classful network will be advertised. To overcome this bug, we use subnets keyword at the end of “**redistribute eigrp 1 metric 100 subnets**” command. Here, 1 is the process id of EIGRP.


Router B(config)# **router ospf 1**

Router B(config-router)# **redistribute static metric 200 subnets**

Router B(config-router)# **redistribute connected metric 100**

Router B(config-router)# **redistribute eigrp 1 metric 100 subnets**

Router B(config-router)# **end**

Router B# **write** 




**Seed metric** or **default metric** is an important term used during route redistribution. Then, what is seed metric?

As you know all the routing protocols use different **metrics** for routes. To redistribute from one routing protocol to another, these metrics also **need to be** **translated**. Because, one routing protocol uses different parameter as a metric and the other uses different one. If we do not translate this value to other, there will be a problem.


Let’s give an example to understand this more. Think about **RIP**, Routing Information Protocol. RIP uses Hop Count as metric. And **OSPF**, Open Shortest Path First uses cost as metric. When we redistribute RIP into OSPF, we need this translation.


At this point **seed metric** or **default metric** arises. Here, **seed metric** is basically the metric that is used for the redistributed route.


There are some **default seed metrics**. These default seed metrics are given below:


**RIP** – **Infinity** for routing protocols, **1** for Static and Connected

**EIGRP** – **Infinity** for routing protocols, **Internal metric** for Static and Connected

**OSPF** – **1** for BGP, **20** for all the others

**BGP** – Set to **IGP metric value** for routing protocols, **0** for Static and Connected

**ISIS** – **0** for all


You can find the **default seed metrics** table below:



In this Route Redistribution example, we have done **OSPF into EIGRP Redistribution** and **EIGRP into OSPF Redistribution**.


**Below, you can find important router & switch configuration lessons and configuration lab pages:**

**DOWNLOAD Cisco Packet Tracer Labs:** https://ipcisco.com/cisco-packet-tracer-configuration-examples/

**DOWNLOAD Cisco GNS3 Labs:** https://ipcisco.com/cisco-gns3-configurations/

**Join Cisco Packet Tracer Labs Course:** https://ipcisco.com/course/cisco-packet-tracer-course/

**Join Cisco CCNA Labs Couse:** https://ipcisco.com/course/ccna-200-301-labs/

**Join CiscoCCNP ENCOR Labs Course:** https://ipcisco.com/course/ccnp-350-401-encor-labs/


Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
