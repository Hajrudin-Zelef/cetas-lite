---
id: collect-261001-huawei/huawei/using-bgp-med-attribute-to-control-route-selection-on-huawei-devices-0ddd867d-1
title: "Using BGP MED Attribute to Control Route Selection on Huawei Devices"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-01", "2024-02", "2024-03", "2024-04", "2024-05", "2024-06", "2024-07", "2024-08", "2024-09", "2024-10", "2024-11", "2024-12", "2025-01", "2025-02", "2025-09"]
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/using-bgp-med-attribute-to-control-route-selection-on-huawei-devices-0ddd867d.md
source_anchor: ""
source_lines: [1, 211]
sha256: 7df124450831a1f6fcdd0a776dffcfeafdbd2cd00975700e96e5875680f098fd
---

# Using BGP MED Attribute to Control Route Selection on Huawei Devices

(Source archive : https://lyfeytech.co.ke/using-bgp-med-attribute-to-control-route-selection-on-huawei-devices/ — capture 20251206001913)

[ ![Lyfey Technologies](https://lyfeytech.co.ke/wp-content/uploads/2024/01/lyfey.png) ](https://lyfeytech.co.ke/)

[ ![Lyfey Technologies](https://lyfeytech.co.ke/wp-content/uploads/2024/01/lyfey.png) ](https://lyfeytech.co.ke/)

 

  * [Blog](https://lyfeytech.co.ke/blog-2/)
  * [About Us](https://lyfeytech.co.ke/about-us/)
  * [Privacy Policy](https://lyfeytech.co.ke/privacy-policy/)
  * [News](https://lyfeytech.co.ke/news/)



[Home](https://lyfeytech.co.ke "Go to Lyfey Technologies.") __[Blog-Test](https://lyfeytech.co.ke/blog-test/ "Go to Blog-Test.") __[Networking](https://lyfeytech.co.ke/category/networking/ "Go to the Networking category archives.") __[Huawei](https://lyfeytech.co.ke/category/networking/huawei/ "Go to the Huawei category archives.") __Using BGP MED Attribute to Control Route Selection on Huawei Devices

#  Using BGP MED Attribute to Control Route Selection on Huawei Devices 

# Using BGP MED Attribute to Control Route Selection on Huawei Devices

  * __May 16, 2024
  * Posted by: Lyfey Technologies
  * Categories: Huawei, Networking



[__390 Comments](https://lyfeytech.co.ke/using-bgp-med-attribute-to-control-route-selection-on-huawei-devices/#comments)

![](https://lyfeytech.co.ke/wp-content/uploads/2024/05/MED_HUAWEI.png)

**Configuration Steps**

Step 1: Configure Hostnames and IP addresses on interfaces for all routers.
    
    
    ***************************R1
    sys
    sysname R1
    #
    interface GigabitEthernet0/0/0
     ip address 10.10.10.0 255.255.255.254
    #
    interface GigabitEthernet0/0/1
     ip address 10.10.10.2 255.255.255.254
    #
    interface LoopBack0
     ip address 1.1.1.1 255.255.255.255
    #
    
    
    ***************************R2
    sys
    sysname R2
    #
    interface GigabitEthernet0/0/0
     ip address 10.10.10.4 255.255.255.254
    #
    interface GigabitEthernet0/0/1
     ip address 10.10.10.3 255.255.255.254
    #
    interface LoopBack0
     ip address 2.2.2.2 255.255.255.255
    #
    
    
    ***************************R3
    sys
    sysname R3
    #
    interface GigabitEthernet0/0/0
     ip address 10.10.10.5 255.255.255.254
    #
    interface GigabitEthernet0/0/1
     ip address 10.10.10.1 255.255.255.254
    #
    interface LoopBack0
     ip address 3.3.3.3 255.255.255.255
    #

Step 2: Configure EBGP peering sessions between R1, R2, and R3. Configure IBGP session between R2 and R3.
    
    
    ***************************R1
    #
    bgp 100
     router-id 1.1.1.1
     peer 10.10.10.1 as-number 200
     peer 10.10.10.3 as-number 200
     #
     ipv4-family unicast
      peer 10.10.10.1 enable
      peer 10.10.10.3 enable
    #
    
    
    *******************************R2
    bgp 200
     router-id 2.2.2.2
     peer 10.10.10.2 as-number 100
     peer 10.10.10.5 as-number 200
     #
     ipv4-family unicast
      peer 10.10.10.2 enable
      peer 10.10.10.5 enable
    #
    
    
    ************************R3
    #
    bgp 200
     router-id 3.3.3.3
     peer 10.10.10.0 as-number 100
     peer 10.10.10.4 as-number 200
     #
     ipv4-family unicast
      peer 10.10.10.0 enable
      peer 10.10.10.4 enable
    #

Step 3: Advertise network 10.10.10.4/31 on R2 and R3 and verify the routing on R1.
    
    
    ***********************R1 and R2
    bgp 200
     network 10.10.10.4 255.255.255.254
    #

> On R1, we observe that there are two valid routes to 10.10.10.4/31 and the route via R2 is preffered due to lower router ID of R2(2.2.2.2) compared to R3(3.3.3.3).

![](https://lyfeytech.co.ke/wp-content/uploads/2024/05/image-13.png)

> We can use MED attribute to control the route selection so that the route learned from R3 will be preffered. MED is like metric for IGPs and is used to control traffic that enters an AS. The route with the smallest MED value is alway preffered, hence we need to configure higher MED value for routes learned from R2 on R1. 

Step 4: Configure high MED for routes from R2.
    
    
    ****************************R2
    #
    route-policy CHANGE_MED permit node 10
     apply cost 100
    #
    bgp 100
     ipv4-family unicast
      peer 10.10.10.3 route-policy CHANGE_MED import
    #

Step 5: Verify the routing table of R1.

![](https://lyfeytech.co.ke/wp-content/uploads/2024/05/image-14-1024x635.png)

> The route from R2 now has MED of 100 while the route from R3 has default MED value of 0 hence its preffered as the best route by R1.

Thank you for reading. Please leave your comments below and check our blog for more interesting content on networking.

**Related Posts**

  * [Step by step guide on how to implement different networking protocols on Juniper MX routers](https://lyfeytech.co.ke/step-by-step-guide-on-how-to-implement-different-networking-protocols-on-juniper-mx-routers/)
  * [L2 EVPN Implementation on Huawei Routers.](https://lyfeytech.co.ke/l2-evpn-implementation-on-huawei-routers/)
  * [VRRP Monitoring of the Uplink Interface status on Huawei routers.](https://lyfeytech.co.ke/vrrp-monitoring-of-the-uplink-interface-status-on-huawei-routers/)
  * [Association between VRRP and BFD Implementation on Huawei routers.](https://lyfeytech.co.ke/association-between-vrrp-and-bfd-implementation-on-huawei-routers/)
  * [Association between VRRP and STP Implementation on Huawei routers.](https://lyfeytech.co.ke/association-between-vrrp-and-stp-implementation-on-huawei-routers/)



  
  


[BGP](https://lyfeytech.co.ke/tag/bgp/) [Huawei](https://lyfeytech.co.ke/tag/huawei/) [Troubleshooting](https://lyfeytech.co.ke/tag/troubleshooting/)

### Leave a Reply [Cancel reply](/using-bgp-med-attribute-to-control-route-selection-on-huawei-devices/#respond)

You must be [logged in](https://lyfeytech.co.ke/wp-login.php?redirect_to=https%3A%2F%2Flyfeytech.co.ke%2Fusing-bgp-med-attribute-to-control-route-selection-on-huawei-devices%2F) to post a comment.

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

