---
id: collect-261001-huawei/huawei/configuring-bgp-as-path-filter-on-huawei-routers-7195e678-1
title: "Configuring BGP AS-PATH_FILTER on Huawei Routers."
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/configuring-bgp-as-path-filter-on-huawei-routers-7195e678.md
source_anchor: ""
source_lines: [1, 193]
sha256: 2f5e77c0819c60a242e7b39301acb2a53f9f88933e69df20f7c409aa8cfea01b
---

# Configuring BGP AS-PATH_FILTER on Huawei Routers.

(Source archive : https://lyfeytech.co.ke/configuring-bgp-as-path_filter-on-huawei-routers/ — capture 20260113000059)

[ ![Lyfey Technologies](https://lyfeytech.co.ke/wp-content/uploads/2024/01/lyfey.png) ](https://lyfeytech.co.ke/)

[ ![Lyfey Technologies](https://lyfeytech.co.ke/wp-content/uploads/2024/01/lyfey.png) ](https://lyfeytech.co.ke/)

 

  * [Blog](https://lyfeytech.co.ke/blog-2/)
  * [About Us](https://lyfeytech.co.ke/about-us/)
  * [Privacy Policy](https://lyfeytech.co.ke/privacy-policy/)
  * [News](https://lyfeytech.co.ke/news/)



[Home](https://lyfeytech.co.ke "Go to Lyfey Technologies.") __[Blog](https://lyfeytech.co.ke/blog-2/ "Go to Blog.") __[Networking](https://lyfeytech.co.ke/category/networking/ "Go to the Networking category archives.") __[Huawei](https://lyfeytech.co.ke/category/networking/huawei/ "Go to the Huawei category archives.") __Configuring BGP AS-PATH_FILTER on Huawei Routers.

#  Configuring BGP AS-PATH_FILTER on Huawei Routers. 

# Configuring BGP AS-PATH_FILTER on Huawei Routers.

  * __April 13, 2024
  * Posted by: Lyfey Technologies
  * Categories: Huawei, Networking



[__383 Comments](https://lyfeytech.co.ke/configuring-bgp-as-path_filter-on-huawei-routers/#comments)

![](https://lyfeytech.co.ke/wp-content/uploads/2024/04/AS_PATH_FILTER.png)

**Configuration Steps**

Step 1: Configure Interface IPs on R1, R2, and R3.
    
    
    ***************************R1
    sys
    sysname R1
    #
    #
    interface LoopBack0
     ip address 1.1.1.1 255.255.255.255
    #
    interface GigabitEthernet0/0/0
     ip address 10.10.10.4 255.255.255.254
    #
    interface GigabitEthernet0/0/1
     ip address 10.10.10.2 255.255.255.254
    #
    
    
    ***************************R2
    sys
    sysname R2
    #
    #
    interface LoopBack0
     ip address 2.2.2.2 255.255.255.255
    #
    interface GigabitEthernet0/0/0
     ip address 10.10.10.5 255.255.255.254
    #
    interface GigabitEthernet0/0/1
     ip address 10.10.10.1 255.255.255.254
    #
    
    
    ***************************R3
    sys
    sysname R3
    #
    #
    interface LoopBack0
     ip address 3.3.3.3 255.255.255.255
    #
    interface GigabitEthernet0/0/0
     ip address 10.10.10.0 255.255.255.254
    #
    interface GigabitEthernet0/0/1
     ip address 10.10.10.3 255.255.255.254
    #

Step 2: Configure EBGP peering between routers in AS 100, AS 200, and AS 300. Advertise Loopback 0 networks on all routers.
    
    
    *****************************R1
    bgp 100
     router-id 1.1.1.1
     peer 10.10.10.3 as-number 200
     peer 10.10.10.5 as-number 300
     #
     ipv4-family unicast
      network 1.1.1.1 255.255.255.255
      peer 10.10.10.3 enable
      peer 10.10.10.5 enable
    #
    
    
    *****************************R2
    bgp 300
     router-id 2.2.2.2
     peer 10.10.10.0 as-number 200
     peer 10.10.10.4 as-number 100
     #
     ipv4-family unicast
      network 2.2.2.2 255.255.255.255
      peer 10.10.10.0 enable
      peer 10.10.10.4 enable
    #
    
    
    *****************************R3
    bgp 200
     router-id 3.3.3.3
     peer 10.10.10.1 as-number 300
     peer 10.10.10.2 as-number 100
     #
     ipv4-family unicast
      network 3.3.3.3 255.255.255.255
      peer 10.10.10.1 enable
      peer 10.10.10.2 enable
    #

Step 3: Verify BGP peering Status on all routers.

![](https://lyfeytech.co.ke/wp-content/uploads/2024/04/image-56-1024x346.png) ![](https://lyfeytech.co.ke/wp-content/uploads/2024/04/image-57-1024x325.png) ![](https://lyfeytech.co.ke/wp-content/uploads/2024/04/image-58-1024x338.png)

Step 4: Verify the BGP routing table on R1, R2, and R3.

![](https://lyfeytech.co.ke/wp-content/uploads/2024/04/image-62.png) ![](https://lyfeytech.co.ke/wp-content/uploads/2024/04/image-63-1024x419.png) ![](https://lyfeytech.co.ke/wp-content/uploads/2024/04/image-64.png)

R1, R2, and R3 are learning from each other correctly. The objective of this lab is to prevent R1 advertising routes learned from AS 200 to AS 300 and vice versa using AS_PATH filters.

Step 5: Configure AS_PATH filters on R1 to control router advertised from R1 to R2 and R3.
    
    
    *****************************R1
    #
    ip as-path-filter FILTER_AS_200 deny _200_
    ip as-path-filter FILTER_AS_200 permit .*
    ip as-path-filter FILTER_AS_300 deny _300_
    ip as-path-filter FILTER_AS_300 permit .*
    #
    bgp 100
      peer 10.10.10.3 as-path-filter FILTER_AS_300 export
      peer 10.10.10.5 as-path-filter FILTER_AS_200 export
    #

Step 6: Verify that R1 does not advertise routes learned from R2 and R3 to each other.

![](https://lyfeytech.co.ke/wp-content/uploads/2024/04/image-65.png)R1 only advertises its own routes to R2 and R3. Other routes are not advertised. ![](https://lyfeytech.co.ke/wp-content/uploads/2024/04/image-66-1024x373.png) ![](https://lyfeytech.co.ke/wp-content/uploads/2024/04/image-67-1024x392.png)

Thank You for reading our article. Check out other related posts on our website:

**Related Posts**

  * [Step by step guide on how to implement different networking protocols on Juniper MX routers](https://lyfeytech.co.ke/step-by-step-guide-on-how-to-implement-different-networking-protocols-on-juniper-mx-routers/)
  * [L2 EVPN Implementation on Huawei Routers.](https://lyfeytech.co.ke/l2-evpn-implementation-on-huawei-routers/)
  * [VRRP Monitoring of the Uplink Interface status on Huawei routers.](https://lyfeytech.co.ke/vrrp-monitoring-of-the-uplink-interface-status-on-huawei-routers/)
  * [Association between VRRP and BFD Implementation on Huawei routers.](https://lyfeytech.co.ke/association-between-vrrp-and-bfd-implementation-on-huawei-routers/)
  * [Association between VRRP and STP Implementation on Huawei routers.](https://lyfeytech.co.ke/association-between-vrrp-and-stp-implementation-on-huawei-routers/)



Check out our YouTube channel for more content. [Lyfey Technologies Channel](https://www.youtube.com/channel/UCyAw6CxRu2bjhdC6f-U88jQ "Lyfey Technologies Channel")

  
  


[BGP](https://lyfeytech.co.ke/tag/bgp/) [Huawei](https://lyfeytech.co.ke/tag/huawei/)

### Leave a Reply [Cancel reply](/configuring-bgp-as-path_filter-on-huawei-routers/#respond)

You must be [logged in](https://lyfeytech.co.ke/wp-login.php?redirect_to=https%3A%2F%2Flyfeytech.co.ke%2Fconfiguring-bgp-as-path_filter-on-huawei-routers%2F) to post a comment.

__

##### Recent Comments

  * [Crypto Predictions](https://cryptopredic.com) on [Implementing OSPF LSA filtering & Inter-area route filtering on Huawei Routers](https://lyfeytech.co.ke/implementing-ospf-lsa-filtering-inter-area-route-filtering-on-huawei-routers/#comment-1948)
  * [Lyfey Technologies](https://lyfeytech.co.ke) on [Implementation of Pseudo wire on Huawei Routers.](https://lyfeytech.co.ke/implementation-of-pseudo-wire-on-huawei-routers/#comment-45)
  * Imran on [Implementing HoVPN IPRAN Architecture on Huawei Routers](https://lyfeytech.co.ke/implementing-hovpn-ipran-architecture-on-huawei-routers/#comment-44)
  * Imran on [Implementing Aggregate VLAN on Huawei Switches](https://lyfeytech.co.ke/implementing-aggregate-vlan-on-huawei-switches/#comment-43)
  * [Majani James](https://lyfeytech.co.ke/) on [Implementing Aggregate VLAN on Huawei Switches](https://lyfeytech.co.ke/implementing-aggregate-vlan-on-huawei-switches/#comment-42)



##### Archives

