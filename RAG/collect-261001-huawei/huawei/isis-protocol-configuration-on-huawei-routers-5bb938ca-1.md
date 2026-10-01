---
id: collect-261001-huawei/huawei/isis-protocol-configuration-on-huawei-routers-5bb938ca-1
title: "Basic configuration of IS-IS Protocol on Huawei routers"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-01", "2024-02", "2024-02-15", "2024-03", "2024-04", "2024-05", "2024-06", "2024-07", "2024-08", "2024-09", "2024-10", "2024-11", "2024-12", "2025-01", "2025-02", "2025-09"]
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/isis-protocol-configuration-on-huawei-routers-5bb938ca.md
source_anchor: ""
source_lines: [1, 203]
sha256: 8d35ce64adb357a28f7e7b445ecbf21f5ccb11fea3329b648200c9b777a23161
---

# Basic configuration of IS-IS Protocol on Huawei routers

(Source archive : https://lyfeytech.co.ke/isis-protocol-configuration-on-huawei-routers/ — capture 20260112225055)

[ ![Lyfey Technologies](https://lyfeytech.co.ke/wp-content/uploads/2024/01/lyfey.png) ](https://lyfeytech.co.ke/)

[ ![Lyfey Technologies](https://lyfeytech.co.ke/wp-content/uploads/2024/01/lyfey.png) ](https://lyfeytech.co.ke/)

 

  * [Blog](https://lyfeytech.co.ke/blog-2/)
  * [About Us](https://lyfeytech.co.ke/about-us/)
  * [Privacy Policy](https://lyfeytech.co.ke/privacy-policy/)
  * [News](https://lyfeytech.co.ke/news/)



[Home](https://lyfeytech.co.ke "Go to Lyfey Technologies.") __[Blog](https://lyfeytech.co.ke/blog-2/ "Go to Blog.") __[Networking](https://lyfeytech.co.ke/category/networking/ "Go to the Networking category archives.") __[Huawei](https://lyfeytech.co.ke/category/networking/huawei/ "Go to the Huawei category archives.") __Basic configuration of IS-IS Protocol on Huawei routers

#  Basic configuration of IS-IS Protocol on Huawei routers 

# Basic configuration of IS-IS Protocol on Huawei routers

  * __January 21, 2024
  * Posted by: Lyfey Technologies
  * Category: Huawei



[__385 Comments](https://lyfeytech.co.ke/isis-protocol-configuration-on-huawei-routers/#comments)

![](https://lyfeytech.co.ke/wp-content/uploads/2024/01/image-7.png)

**What is ISIS protocol?**

ISIS(Intermediate System to Intermediate System) is a link-state dynamic routing protocol. It's an Interior Gateway Protocol (IGP) and is used within an autonomous system (AS). It uses the shortest path first (SPF) algorithm to calculate routes.

In this article, we explain how to do basic configuration of IS-IS on Huawei routers. Below is our simple topology.

![](https://lyfeytech.co.ke/wp-content/uploads/2024/01/image-7.png)

**Configuration Steps:**

Configure IP addresses on interfaces, and add descriptions on your physical interfaces.
    
    
    ************************************R1********************************************************
    
    interface LoopBack0
    ip address 1.1.1.1 255.255.255.255
    #
    interface GigabitEthernet0/0/0
    description TO_ROUTER_2_GE0/0/0
    ip address 10.251.10.0 255.255.255.254
    #
    
    
    ******************************************R2****************************************
    interface LoopBack0
     ip address 2.2.2.2 255.255.255.255
    #
    interface GigabitEthernet0/0/0
     description TO_R1_GE0/0/0
     ip address 10.251.10.1 255.255.255.254
    #
    interface GigabitEthernet0/0/1
     description TO_R3_GE0/0/0
     ip address 10.251.10.2 255.255.255.254
    #
    [R1-GigabitEthernet0/0/1]
    
    
    ********************************************Router 3***********************************
    
    interface LoopBack0
     ip address 3.3.3.3 255.255.255.255
    #
    interface GigabitEthernet0/0/0
     description TO_R2_GE0/0/1
     ip address 10.251.10.3 255.255.255.254
    #

Configure and enable ISIS on interfaces.
    
    
    **********************************R1*****************************
    isis 10
     is-level level-2
     cost-style wide
     network-entity 49.0010.0010.0100.1001.00
     is-name R1
    #
    interface LoopBack0
    isis enable 10
    
    interface GigabitEthernet0/0/0
    isis enable 10
    
    interface GigabitEthernet0/0/1
    isis enable 10
    
    **********************************R2*****************************
    isis 10
     is-level level-2
     cost-style wide
     network-entity 49.0010.0020.0200.2002.00
     is-name R1
    #
    interface LoopBack0
    isis enable 10
    
    interface GigabitEthernet0/0/0
    isis enable 10
    
    interface GigabitEthernet0/0/1
    isis enable 10
    **********************************R3*****************************
    
    isis 10
     is-level level-2
     cost-style wide
     network-entity 49.0010.0030.0300.3003.00
     is-name R1
    #
    interface LoopBack0
    isis enable 10
    
    interface GigabitEthernet0/0/0
    isis enable 10

Verify the status of ISIS and check the ISIS routing table on Router 2:

Run the command **display isis peer** to check the neighbor status

![](https://lyfeytech.co.ke/wp-content/uploads/2024/01/image-9.png)R2 has two peers in an established state.

Run the**display isis peer** verbose command to see more details about the peering.

![](https://lyfeytech.co.ke/wp-content/uploads/2024/01/image-10.png)You can see the uptime and area IDs.

Run the command display ip routing-table protocol isis to check the routing table details.

![](https://lyfeytech.co.ke/wp-content/uploads/2024/01/image-11.png)We have learned the Loopback IPs of R1 and R2 through ISIS.

Test connectivity between R1 and R3. 

![](https://lyfeytech.co.ke/wp-content/uploads/2024/01/image-12.png)We can ping R3 from R1.

  
  


[Huawei](https://lyfeytech.co.ke/tag/huawei/) [IGP](https://lyfeytech.co.ke/tag/igp/) [ISIS](https://lyfeytech.co.ke/tag/isis/)

####  385 Comments 

  * [BGP INTER-AS Option B Configuration and troubleshooting on Huawei routers - Lyfey Technologies](https://lyfeytech.co.ke/bgp-inter-as-option-b-configuration-and-troubleshooting-on-huawei-routers/)

[ February 15, 2024 at 3:58 pm ](https://lyfeytech.co.ke/isis-protocol-configuration-on-huawei-routers/#comment-13) [Log in to Reply](https://lyfeytech.co.ke/wp-login.php?redirect_to=https%3A%2F%2Flyfeytech.co.ke%2Fisis-protocol-configuration-on-huawei-routers%2F)

[…] MPLS, and LDP on all interfaces within the AS. Refer to this article for IS-IS Configurations, Configuring ISIS on Huawei. Below is sample IS-IS and interface configuration for […]




### Leave a Reply [Cancel reply](/isis-protocol-configuration-on-huawei-routers/#respond)

You must be [logged in](https://lyfeytech.co.ke/wp-login.php?redirect_to=https%3A%2F%2Flyfeytech.co.ke%2Fisis-protocol-configuration-on-huawei-routers%2F) to post a comment.

__

##### Recent Comments

  * [Crypto Predictions](https://cryptopredic.com) on [Implementing OSPF LSA filtering & Inter-area route filtering on Huawei Routers](https://lyfeytech.co.ke/implementing-ospf-lsa-filtering-inter-area-route-filtering-on-huawei-routers/#comment-1948)
  * [Lyfey Technologies](https://lyfeytech.co.ke) on [Implementation of Pseudo wire on Huawei Routers.](https://lyfeytech.co.ke/implementation-of-pseudo-wire-on-huawei-routers/#comment-45)
  * Imran on [Implementing HoVPN IPRAN Architecture on Huawei Routers](https://lyfeytech.co.ke/implementing-hovpn-ipran-architecture-on-huawei-routers/#comment-44)
  * Imran on [Implementing Aggregate VLAN on Huawei Switches](https://lyfeytech.co.ke/implementing-aggregate-vlan-on-huawei-switches/#comment-43)
  * [Majani James](https://lyfeytech.co.ke/) on [Implementing Aggregate VLAN on Huawei Switches](https://lyfeytech.co.ke/implementing-aggregate-vlan-on-huawei-switches/#comment-42)



##### Archives

  * [September 2025](https://lyfeytech.co.ke/2025/09/)
  * [February 2025](https://lyfeytech.co.ke/2025/02/)
  * [January 2025](https://lyfeytech.co.ke/2025/01/)
  * [December 2024](https://lyfeytech.co.ke/2024/12/)
  * [November 2024](https://lyfeytech.co.ke/2024/11/)
  * [October 2024](https://lyfeytech.co.ke/2024/10/)
  * [September 2024](https://lyfeytech.co.ke/2024/09/)
  * [August 2024](https://lyfeytech.co.ke/2024/08/)
  * [July 2024](https://lyfeytech.co.ke/2024/07/)
  * [June 2024](https://lyfeytech.co.ke/2024/06/)
  * [May 2024](https://lyfeytech.co.ke/2024/05/)
  * [April 2024](https://lyfeytech.co.ke/2024/04/)
  * [March 2024](https://lyfeytech.co.ke/2024/03/)
  * [February 2024](https://lyfeytech.co.ke/2024/02/)
  * [January 2024](https://lyfeytech.co.ke/2024/01/)



##### Categories

