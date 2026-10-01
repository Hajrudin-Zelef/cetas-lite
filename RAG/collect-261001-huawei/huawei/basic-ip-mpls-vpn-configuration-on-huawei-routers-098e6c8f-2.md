---
id: collect-261001-huawei/huawei/basic-ip-mpls-vpn-configuration-on-huawei-routers-098e6c8f-2
title: "Basic IP MPLS VPN configuration on Huawei Routers"
domain: huawei
role: reference
task: reference
actors: ["Huawei", "Meta"]
dates: ["2024-01", "2024-02", "2024-03", "2024-04", "2024-05", "2024-06", "2024-07", "2024-08", "2024-09", "2024-10", "2024-11", "2024-12", "2025-01", "2025-02", "2025-09"]
keywords: ["training"]
source: docs/RAG/collect-261001-huawei/basic-ip-mpls-vpn-configuration-on-huawei-routers-098e6c8f.md
source_anchor: ""
source_lines: [310, 446]
sha256: 52086335cffdcd1542cee83b812a3e120b399febba3128dcf0cbd02b11e3be80
---

# Basic IP MPLS VPN configuration on Huawei Routers

    ******************************CE04
    sys
    sysname CE04
    #
    interface GigabitEthernet0/0/0
     ip address 10.10.10.9 255.255.255.254
    #
    #
    interface LoopBack0
     ip address 5.5.5.5 255.255.255.255
    #
    bgp 400
     peer 10.10.10.8 as-number 100
     #
     ipv4-family unicast
      undo synchronization
      network 5.5.5.5 255.255.255.255
      peer 10.10.10.8 enable
    #
    commit

Step 8: Verify BGP peering between CEs and PEs is up and routes are propagated between CEs in the same VPN Instance.

![](https://lyfeytech.co.ke/wp-content/uploads/2024/03/image-145-1024x514.png) ![](https://lyfeytech.co.ke/wp-content/uploads/2024/03/image-146-1024x465.png) ![](https://lyfeytech.co.ke/wp-content/uploads/2024/03/image-147.png)

Thank You for reading our blogs, please your comments in the comments section and check out other related posts on our blog.

**Related Posts**

  * [Step by step guide on how to implement different networking protocols on Juniper MX routers](https://lyfeytech.co.ke/step-by-step-guide-on-how-to-implement-different-networking-protocols-on-juniper-mx-routers/)
  * [L2 EVPN Implementation on Huawei Routers.](https://lyfeytech.co.ke/l2-evpn-implementation-on-huawei-routers/)
  * [VRRP Monitoring of the Uplink Interface status on Huawei routers.](https://lyfeytech.co.ke/vrrp-monitoring-of-the-uplink-interface-status-on-huawei-routers/)
  * [Association between VRRP and BFD Implementation on Huawei routers.](https://lyfeytech.co.ke/association-between-vrrp-and-bfd-implementation-on-huawei-routers/)
  * [Association between VRRP and STP Implementation on Huawei routers.](https://lyfeytech.co.ke/association-between-vrrp-and-stp-implementation-on-huawei-routers/)



  
  


[BGP](https://lyfeytech.co.ke/tag/bgp/) [Huawei](https://lyfeytech.co.ke/tag/huawei/) [OSPF](https://lyfeytech.co.ke/tag/ospf/) [Troubleshooting](https://lyfeytech.co.ke/tag/troubleshooting/) [VPN](https://lyfeytech.co.ke/tag/vpn/)

### Leave a Reply [Cancel reply](/basic-ip-mpls-vpn-configuration-on-huawei-routers/#respond)

You must be [logged in](https://lyfeytech.co.ke/wp-login.php?redirect_to=https%3A%2F%2Flyfeytech.co.ke%2Fbasic-ip-mpls-vpn-configuration-on-huawei-routers%2F) to post a comment.

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

  * [Blog](https://lyfeytech.co.ke/blog-2/)
  * [About Us](https://lyfeytech.co.ke/about-us/)
  * [Privacy Policy](https://lyfeytech.co.ke/privacy-policy/)
  * [News](https://lyfeytech.co.ke/news/)



#### Our Service

  * [Company overview](https://lyfeytech.co.ke/company-overview/)
  * [Careers](https://lyfeytech.co.ke/company-overview/careers/)
  * [Company history](https://lyfeytech.co.ke/company-overview/company-history/)
  * [Our approach](https://lyfeytech.co.ke/company-overview/our-approach/)
  * [Partners](https://lyfeytech.co.ke/company-overview/our-partners/)
  * [Our team grid](https://lyfeytech.co.ke/company-overview/our-team-grid/)



© 2026 Lyfey Technologies - Networking Training & Consultancy 

This website uses cookies and asks your personal data to enhance your browsing experience.

Ok, I agree [Privacy Policy](https://lyfeytech.co.ke/privacy-policy/)

[Privacy Policy](https://lyfeytech.co.ke/wpautoterms/privacy-policy/)
