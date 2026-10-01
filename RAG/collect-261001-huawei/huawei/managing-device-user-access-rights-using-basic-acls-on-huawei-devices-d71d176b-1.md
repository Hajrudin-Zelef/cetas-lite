---
id: collect-261001-huawei/huawei/managing-device-user-access-rights-using-basic-acls-on-huawei-devices-d71d176b-1
title: "Managing Device/User Access rights using Basic ACLs on Huawei Devices"
domain: huawei
role: reference
task: reference
actors: ["Huawei", "Meta"]
dates: ["2024-01", "2024-02", "2024-03", "2024-04", "2024-05", "2024-06", "2024-07", "2024-08", "2024-09", "2024-10", "2024-11", "2024-12", "2025-01", "2025-02", "2025-09"]
keywords: []
source: docs/RAG/collect-261001-huawei/managing-device-user-access-rights-using-basic-acls-on-huawei-devices-d71d176b.md
source_anchor: ""
source_lines: [1, 204]
sha256: 42da56ea78b42b056bd91888277c0d4cba3794ffbec235a07c20b66529f64fbe
---

# Managing Device/User Access rights using Basic ACLs on Huawei Devices

(Source archive : https://lyfeytech.co.ke/managing-device-user-access-rights-using-basic-acls-on-huawei-devices/ — capture 20251206002350)

[ ![Lyfey Technologies](https://lyfeytech.co.ke/wp-content/uploads/2024/01/lyfey.png) ](https://lyfeytech.co.ke/)

[ ![Lyfey Technologies](https://lyfeytech.co.ke/wp-content/uploads/2024/01/lyfey.png) ](https://lyfeytech.co.ke/)

 

  * [Blog](https://lyfeytech.co.ke/blog-2/)
  * [About Us](https://lyfeytech.co.ke/about-us/)
  * [Privacy Policy](https://lyfeytech.co.ke/privacy-policy/)
  * [News](https://lyfeytech.co.ke/news/)



[Home](https://lyfeytech.co.ke "Go to Lyfey Technologies.") __[Blog-Test](https://lyfeytech.co.ke/blog-test/ "Go to Blog-Test.") __[Networking](https://lyfeytech.co.ke/category/networking/ "Go to the Networking category archives.") __[Huawei](https://lyfeytech.co.ke/category/networking/huawei/ "Go to the Huawei category archives.") __Managing Device/User Access rights using Basic ACLs on Huawei Devices

#  Managing Device/User Access rights using Basic ACLs on Huawei Devices 

# Managing Device/User Access rights using Basic ACLs on Huawei Devices

  * __May 30, 2024
  * Posted by: Lyfey Technologies
  * Categories: Huawei, Networking



[__No Comments](https://lyfeytech.co.ke/managing-device-user-access-rights-using-basic-acls-on-huawei-devices/#respond)

![](https://lyfeytech.co.ke/wp-content/uploads/2024/05/basic_ACLs.png)

This lab demonstrates how to configure basic ACLs on Huawei devices.

**Configurations Steps**

Step 1: Configure IP addresses and hostnames on CE01 and CE02.
    
    
    ******************************CE01
    sys
    #
    sysname CE01
    #
    interface GigabitEthernet0/0/0
     ip address 10.10.10.1 255.255.255.0
    #
    commit
    
    
    ******************************CE02
    sys
    #
    sysname CE02
    #
    interface GigabitEthernet0/0/0
     ip address 20.20.20.1 255.255.255.0
    #
    commit

Step 2: Configure the hostname on PE01. Create two VPN instances and bind them on interfaces connecting CE routers. Assign IP addresses to interfaces connecting CEs.
    
    
    *****************************PE01
    #
    sysname PE02
    #
    ip vpn-instance VPN-A
     ipv4-family
      route-distinguisher 100:1
      vpn-target 100:100 export-extcommunity
      vpn-target 100:100 import-extcommunity
    #
    ip vpn-instance VPN-B
     ipv4-family
      route-distinguisher 200:1
      vpn-target 200:200 export-extcommunity
      vpn-target 200:200 import-extcommunity
    #
    interface GigabitEthernet0/0/0
     ip binding vpn-instance VPN-A
     ip address 10.10.10.254 255.255.255.0
    #
    interface GigabitEthernet0/0/1
     ip binding vpn-instance VPN-B
     ip address 20.20.20.254 255.255.255.0
    #

Step 3: Configure ACLs on PE01 to allow users in VPN-A and deny users in VPN-B.
    
    
    **************************PE01
    #
    acl name CONTROL_ACCESS number 2000
     rule 5 permit vpn-instance VPN-A
     rule 10 deny vpn-instance VPN-B
    #

Step 4: Apply the ACL in Telnet services on the PE.
    
    
    *********************PE01
    #
    user-interface con 0
    user-interface vty 0 4
     acl 2000 inbound
     set authentication password cipher Huawei_123
    user-interface vty 16 20
    #

Step 5: Verify the configuration by trying to use Telnet from CE01 and CE02 to PE01.

![](https://lyfeytech.co.ke/wp-content/uploads/2024/05/image-28.png) ![](https://lyfeytech.co.ke/wp-content/uploads/2024/05/image-29.png)

> As shown above, we can telnet from CE01 to PE01 but we cannot telnet from CE02 to PE01 because our ACL allows access from users in VPN-A and denies access from users in VPN-B.

Thank you for reading our articles. Leave your comments below and subscribe to our Youtube channel for more content on networking: [Lyfey Technologies Channel](https://youtu.be/sPTbKHzY_7s?si=_SH2FI1NOTsfv6S1 "Lyfey Technologies Channel")

**Latest Posts**

  * [Step by step guide on how to implement different networking protocols on Juniper MX routers](https://lyfeytech.co.ke/step-by-step-guide-on-how-to-implement-different-networking-protocols-on-juniper-mx-routers/)
  * [L2 EVPN Implementation on Huawei Routers.](https://lyfeytech.co.ke/l2-evpn-implementation-on-huawei-routers/)
  * [VRRP Monitoring of the Uplink Interface status on Huawei routers.](https://lyfeytech.co.ke/vrrp-monitoring-of-the-uplink-interface-status-on-huawei-routers/)
  * [Association between VRRP and BFD Implementation on Huawei routers.](https://lyfeytech.co.ke/association-between-vrrp-and-bfd-implementation-on-huawei-routers/)
  * [Association between VRRP and STP Implementation on Huawei routers.](https://lyfeytech.co.ke/association-between-vrrp-and-stp-implementation-on-huawei-routers/)



  
  


[Huawei](https://lyfeytech.co.ke/tag/huawei/) [Troubleshooting](https://lyfeytech.co.ke/tag/troubleshooting/)

### Leave a Reply [Cancel reply](/managing-device-user-access-rights-using-basic-acls-on-huawei-devices/#respond)

You must be [logged in](https://lyfeytech.co.ke/wp-login.php?redirect_to=https%3A%2F%2Flyfeytech.co.ke%2Fmanaging-device-user-access-rights-using-basic-acls-on-huawei-devices%2F) to post a comment.

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

  * [Cisco](https://lyfeytech.co.ke/category/networking/cisco/)
  * [Firewalls](https://lyfeytech.co.ke/category/networking/firewall/)
  * [Huawei](https://lyfeytech.co.ke/category/networking/huawei/)
  * [Juniper](https://lyfeytech.co.ke/category/networking/juniper-networking/)
  * [Networking](https://lyfeytech.co.ke/category/networking/)
  * [Nokia](https://lyfeytech.co.ke/category/networking/nokia/)



##### Meta

  * [Log in](https://lyfeytech.co.ke/wp-login.php)
  * [Entries feed](https://lyfeytech.co.ke/feed/)
  * [Comments feed](https://lyfeytech.co.ke/comments/feed/)
  * [WordPress.org](https://wordpress.org/)



  * [ __](https://www.linkedin.com/in/martin-indeche-668402115/)
  * [ __](https://twitter.com/stylemix_themes/)
  * [ __](https://www.youtube.com/@LyfeyTechnologies)



#### Contact info

Phone: +254 739106372  
Email: [[email protected]](/cdn-cgi/l/email-protection)

#### Quick links

