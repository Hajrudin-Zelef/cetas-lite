---
id: collect-261001-huawei/huawei/implementing-link-aggregation-on-huawei-routers-11173a25-1
title: "Implementing Link Aggregation on Huawei routers."
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-02-19"]
keywords: ["parameters", "preemption"]
source: docs/RAG/collect-261001-huawei/implementing-link-aggregation-on-huawei-routers-11173a25.md
source_anchor: ""
source_lines: [1, 163]
sha256: b273056c7c15a6f8ba50b9301004409fea6c55162a3c8079b97d7880da7a7039
---

# Implementing Link Aggregation on Huawei routers.

(Source archive : https://lyfeytech.co.ke/implementing-link-aggregation-on-huawei-routers/ — capture 20260313224548)

[ ![Lyfey Technologies](https://lyfeytech.co.ke/wp-content/uploads/2024/01/lyfey.png) ](https://lyfeytech.co.ke/)

[ ![Lyfey Technologies](https://lyfeytech.co.ke/wp-content/uploads/2024/01/lyfey.png) ](https://lyfeytech.co.ke/)

 

  * [Blog](https://lyfeytech.co.ke/blog-2/)
  * [About Us](https://lyfeytech.co.ke/about-us/)
  * [Privacy Policy](https://lyfeytech.co.ke/privacy-policy/)
  * [News](https://lyfeytech.co.ke/news/)



[Home](https://lyfeytech.co.ke "Go to Lyfey Technologies.") __[Blog](https://lyfeytech.co.ke/blog-2/ "Go to Blog.") __[Networking](https://lyfeytech.co.ke/category/networking/ "Go to the Networking category archives.") __[Huawei](https://lyfeytech.co.ke/category/networking/huawei/ "Go to the Huawei category archives.") __Implementing Link Aggregation on Huawei routers.

#  Implementing Link Aggregation on Huawei routers. 

# Implementing Link Aggregation on Huawei routers.

  * __February 19, 2024
  * Posted by: Lyfey Technologies
  * Categories: Huawei, Networking



[__383 Comments](https://lyfeytech.co.ke/implementing-link-aggregation-on-huawei-routers/#comments)

![](https://lyfeytech.co.ke/wp-content/uploads/2024/02/LAGS.png)

**What is Link Aggregation**?

Link aggregation is a common networking technique of bundling multiple physical links to a logical link to increase link bandwidth. Link aggregation provides link backup mechanisms, greatly improving link reliability. The aggregated link is referred to as an **Eth-trunk** in Huawei terminologies.

The interfaces that constitute an Eth-Trunk are referred to as **member interfaces**. The link corresponding to a member interface is a **member link**. Member interfaces can be classified into active interfaces, which forward data, and inactive interfaces, which do not forward data, they are backup links in a LAG (Link Aggregation Group). Link Aggregation can operate in two modes Manual and Link Aggregation Control Protocol (LACP) 

**Benefits of Link aggregation**

  * Increased bandwidth: The total bandwidth of the link aggregation interface is the sum of the bandwidth of member interfaces. If you bundle 4 10G links together, you have a 40G link.
  * Higher reliability: When the physical link of a member interface fails, the traffic on the member link is switched to another member link, ensuring uninterrupted service on the trunk link.
  * Load balancing: The traffic is load-balanced among active member interfaces of the LAG.



**Lab Setup.**

![](https://lyfeytech.co.ke/wp-content/uploads/2024/02/image-33.png)

**Step 1: Configure LAGs on the two routers**
    
    
    ****************************************R1
    interface Eth-Trunk1
     description TO_R2_ETH_TRUNK1
     ip address 10.251.251.0 255.255.255.254
     mode lacp-static
    commit
    #
    
    
    ****************************************R2
    interface Eth-Trunk1
     description TO_R1_ETH_TRUNK1
     ip address 10.251.251.1 255.255.255.254
     mode lacp-static
    commit
    #

**Step 2: Configure physical interfaces and add them to the LAG.**

> In production network where routers with multiple line cards are deployed, its a good practice to distribute ports in a lag to different line cards so that a failure of one line card will not have impact on the LAG. The other ports on other line cards will continue to carry traffic avoiding down time.
    
    
    *************************************R1
    interface GigabitEthernet0/0/0
     description TO_R2_GE0/0/0
     eth-trunk 1
    #
    interface GigabitEthernet0/0/1
     description TO_R2_GE0/0/1
     eth-trunk 1
    #
    interface GigabitEthernet0/0/2
     description TO_R2_GE0/0/2
     eth-trunk 1
    #
    interface GigabitEthernet0/0/3
     description TO_R2_GE0/0/3
     eth-trunk 1
    #
    commit
    
    
    interface Eth-Trunk1
     description TO_R1_ETH_TRUNK1
     ip address 10.251.251.1 255.255.255.254
     mode lacp-static
    commit
    #
    
    *********************************R2
    interface GigabitEthernet0/0/0
    description TO_R1_GE0/0/0
    eth-trunk 1
    #
    interface GigabitEthernet0/0/1
    description TO_R1_GE0/0/1
    eth-trunk 1
    #
    interface GigabitEthernet0/0/2
    description TO_R1_GE0/0/2
    eth-trunk 1
    commit
    #
    interface GigabitEthernet0/0/3
    description TO_R1_GE0/0/3
    eth-trunk 1
    commit
    #

**Verification:**

![](https://lyfeytech.co.ke/wp-content/uploads/2024/02/Screenshot-2024-02-19-160801-1024x763.jpg) ![](https://lyfeytech.co.ke/wp-content/uploads/2024/02/Screenshot-2024-02-19-161028-1024x566.jpg)

Please note, that this is just a simple configuration of a LAG. You may be required to configure other parameters like interface priorities, minimum or maximum number of links that should be up in an eth-trunk, preemption, load balancing parameters, etc.

**Latest Posts**

  * [Step by step guide on how to implement different networking protocols on Juniper MX routers](https://lyfeytech.co.ke/step-by-step-guide-on-how-to-implement-different-networking-protocols-on-juniper-mx-routers/)
  * [L2 EVPN Implementation on Huawei Routers.](https://lyfeytech.co.ke/l2-evpn-implementation-on-huawei-routers/)
  * [VRRP Monitoring of the Uplink Interface status on Huawei routers.](https://lyfeytech.co.ke/vrrp-monitoring-of-the-uplink-interface-status-on-huawei-routers/)
  * [Association between VRRP and BFD Implementation on Huawei routers.](https://lyfeytech.co.ke/association-between-vrrp-and-bfd-implementation-on-huawei-routers/)
  * [Association between VRRP and STP Implementation on Huawei routers.](https://lyfeytech.co.ke/association-between-vrrp-and-stp-implementation-on-huawei-routers/)



  
  


[Huawei](https://lyfeytech.co.ke/tag/huawei/) [LAG](https://lyfeytech.co.ke/tag/lag/) [Reliability](https://lyfeytech.co.ke/tag/reliability/)

### Leave a Reply [Cancel reply](/implementing-link-aggregation-on-huawei-routers/#respond)

You must be [logged in](https://lyfeytech.co.ke/wp-login.php?redirect_to=https%3A%2F%2Flyfeytech.co.ke%2Fimplementing-link-aggregation-on-huawei-routers%2F) to post a comment.

__

##### Recent Comments

  * [Crypto Predictions](https://cryptopredic.com) on [Implementing OSPF LSA filtering & Inter-area route filtering on Huawei Routers](https://lyfeytech.co.ke/implementing-ospf-lsa-filtering-inter-area-route-filtering-on-huawei-routers/#comment-1948)
  * [Lyfey Technologies](https://lyfeytech.co.ke) on [Implementation of Pseudo wire on Huawei Routers.](https://lyfeytech.co.ke/implementation-of-pseudo-wire-on-huawei-routers/#comment-45)
  * Imran on [Implementing HoVPN IPRAN Architecture on Huawei Routers](https://lyfeytech.co.ke/implementing-hovpn-ipran-architecture-on-huawei-routers/#comment-44)
  * Imran on [Implementing Aggregate VLAN on Huawei Switches](https://lyfeytech.co.ke/implementing-aggregate-vlan-on-huawei-switches/#comment-43)
  * [Majani James](https://lyfeytech.co.ke/) on [Implementing Aggregate VLAN on Huawei Switches](https://lyfeytech.co.ke/implementing-aggregate-vlan-on-huawei-switches/#comment-42)



##### Archives

