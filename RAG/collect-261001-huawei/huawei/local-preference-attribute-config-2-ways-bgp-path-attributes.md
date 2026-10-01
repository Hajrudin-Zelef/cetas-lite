---
id: collect-261001-huawei/huawei/local-preference-attribute-config-2-ways-bgp-path-attributes
title: "local-preference-attribute-config-2-ways-bgp-path-attributes"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/local-preference-attribute-config-2-ways-bgp-path-attributes.md
source_anchor: ""
source_lines: [1, 55]
sha256: b6f7983e073ed01b520a67ffe55fc9bd25e69c26aeb31dc43b7a4651015492fa
---

# local-preference-attribute-config-2-ways-bgp-path-attributes

Table of Contents

In **BGP**  Path selection algorithm, if your device is a different vendor product (not Cisco), the first **BGP Path Attribute** that is taken into account during **Path Selection** is Local Preference attribute.


**Local Preference** is not a vendor specific **BGP Path Attribute** like all the BGP path attributes except **Weight Attribute**.


**Local Preference** attribute is used to choose external BGP paths. Different edge routers has different Local Preferences to an external destination. These values are sent between IBGP (Interior BGP) neighbours and according to these values, the AS (Autonomous System) exit point is determined. Local Preference is applied to the inbound interface and shows the best outbound interface towards another AS. In other words, here, it shows the exit point of your AS towards another AS.


The **default** value of **Local Preference** is **100**. According to BGP Best Path Algorithm, the router which has the **Highest Local Preference** value is selected as a Best **BGP** path towards outside the Autonomous System.


Think about the previous **Weight Attribute**. Weight attribute was configured as local to that router. But Local Preference attribute is not like this. Different routers has different Local Preference values for that destination. And these Local Preference values are shared withing the AS (Autonomous System).


Local Preference attribute is a Well-known and Discretionary **BGP Path Attribute**. As we mentined before, this means, it is supported by every BGP implementation (well-known) and every **BGP** router recognize it, but optionally present in the **BGP Update packet** (discretionary).

**BGP Local Preference** value can be assigned to a router in **two** ways:


Let’s give an example for each of these **BGP Local Preference** assign ways. As a reference topology, we will use the below basic topology:

RouterA(config)# **router bgp 10**

RouterA(config-router)# **neighbour 10.10.10.1 remote-as 10**

RouterA(config-router)# **neighbour 20.20.20.20 remote-as 20**

RouterA(config-router)# **bgp default local-preference 100**



RouterB(config)# **router bgp 10**

RouterB(config-router)# **neighbour 10.10.10.2 remote-as 10**

RouterB(config-router)# **neighbour 30.30.30.30 remote-as 30**

RouterB(config-router)# **bgp default local-preference 200**



You can also test yourself with **BGP Tests and Questions**.

Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
