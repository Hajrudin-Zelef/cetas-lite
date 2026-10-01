---
id: collect-261001-huawei/huawei/ipv6-over-ipv4-tunneling-using-gre-tunnel-configuration-on-huawei-routers-e7544773
title: "IPv6 over IPv4 Tunneling Using GRE Tunnel Configuration on Huawei Routers"
domain: huawei
role: reference
task: reference
actors: ["Huawei", "Meta"]
dates: ["2024-01", "2024-02", "2024-03", "2024-04", "2024-05", "2024-06", "2024-07", "2024-08", "2024-09", "2024-10", "2024-11", "2024-12", "2025-01", "2025-02", "2025-09"]
keywords: ["training"]
source: docs/RAG/collect-261001-huawei/ipv6-over-ipv4-tunneling-using-gre-tunnel-configuration-on-huawei-routers-e7544773.md
source_anchor: ""
source_lines: [1, 195]
sha256: 4f4d93c6f180440c4833fadae4f8420e546e57d361ee8e67968504c6134923eb
---

# IPv6 over IPv4 Tunneling Using GRE Tunnel Configuration on Huawei Routers

(Source archive : https://lyfeytech.co.ke/ipv6-over-ipv4-tunneling-using-gre-tunnel-configuration-on-huawei-routers/ — capture 20251205233959)

[ ![Lyfey Technologies](https://lyfeytech.co.ke/wp-content/uploads/2024/01/lyfey.png) ](https://lyfeytech.co.ke/)

[ ![Lyfey Technologies](https://lyfeytech.co.ke/wp-content/uploads/2024/01/lyfey.png) ](https://lyfeytech.co.ke/)

 

  * [Blog](https://lyfeytech.co.ke/blog-2/)
  * [About Us](https://lyfeytech.co.ke/about-us/)
  * [Privacy Policy](https://lyfeytech.co.ke/privacy-policy/)
  * [News](https://lyfeytech.co.ke/news/)



[Home](https://lyfeytech.co.ke "Go to Lyfey Technologies.") __[Blog-Test](https://lyfeytech.co.ke/blog-test/ "Go to Blog-Test.") __[Networking](https://lyfeytech.co.ke/category/networking/ "Go to the Networking category archives.") __[Huawei](https://lyfeytech.co.ke/category/networking/huawei/ "Go to the Huawei category archives.") __IPv6 over IPv4 Tunneling Using GRE Tunnel Configuration on Huawei Routers

#  IPv6 over IPv4 Tunneling Using GRE Tunnel Configuration on Huawei Routers 

# IPv6 over IPv4 Tunneling Using GRE Tunnel Configuration on Huawei Routers

  * __January 21, 2024
  * Posted by: Lyfey Technologies
  * Category: Huawei



[__No Comments](https://lyfeytech.co.ke/ipv6-over-ipv4-tunneling-using-gre-tunnel-configuration-on-huawei-routers/#respond)

![](https://lyfeytech.co.ke/wp-content/uploads/2024/01/image-13.png)

Many ISPs are transitioning their networks from IPv4-only networks to IPv6 networks.

IPv6 over IPv4 tunneling technology can be used by ISPs for a smooth transition from IPv4 networks to IPv6 networks. In this lab, we will demonstrate how to configure IPv6 over IPv4 tunneling on Huawei routers using manual and GRE tunnels.

![](https://lyfeytech.co.ke/wp-content/uploads/2024/01/image-13.png)Topology Diagram. Loopback 100 simulates IPV6 networks.

Configuration procedure:

**Step 1** : Configure IPv4 addresses on interfaces.

**Step 2:** Configure OSPF instance 10 area 0 in the backbone network on R1, R2, and R3 routers. confirm the peering is up and reachability is OK. 

![](https://lyfeytech.co.ke/wp-content/uploads/2024/01/image-14.png)

**Step 3:** Enable IPv6 on routers R1 and R2, and configure loopback 100 with IPv6 addresses to simulate IPv6 clients.
    
    
    ***************************************R1*******************************
    #
    ipv6
    #
    interface LoopBack100
     ipv6 enable 
     ipv6 address 2001:2::2/64 
    #
    
    
    ***************************************R3*******************************
    #
    ipv6
    #
    interface LoopBack100
     ipv6 enable 
     ipv6 address 2001:3::3/64 
    #

4\. Configure tunnels on R1 and R2, You need to enable IPv6, configure tunnel protocol, specify the source IP as Loopback 0, and configure the tunnel destination as the IP address of loopback 0 of the remote router.
    
    
    **************************R1*************************
    interface Tunnel0/0/0
     ipv6 enable                # enable IPv6
     ipv6 address 2003:3::1/64 
     tunnel-protocol gre       # configure tunnel protocol
     source LoopBack0           # specify tunnel source 
     destination 3.3.3.3        # configure tunnel destination
    #
    ***************************R3*************************
    interface Tunnel0/0/0
     ipv6 enable 
     ipv6 address 2003:3::2/64 
     tunnel-protocol gre
     source LoopBack0
     destination 1.1.1.1
    #

**Verification: Try to ping the IPV6 address of Loopback 100 on R3 from R1.**

![](https://lyfeytech.co.ke/wp-content/uploads/2024/01/image-15.png)

We can do a packet capture and verify that IPV6 packets are carried over the GRE tunnel.

![](https://lyfeytech.co.ke/wp-content/uploads/2024/01/image-16.png)

  
  


### Leave a Reply [Cancel reply](/ipv6-over-ipv4-tunneling-using-gre-tunnel-configuration-on-huawei-routers/#respond)

You must be [logged in](https://lyfeytech.co.ke/wp-login.php?redirect_to=https%3A%2F%2Flyfeytech.co.ke%2Fipv6-over-ipv4-tunneling-using-gre-tunnel-configuration-on-huawei-routers%2F) to post a comment.

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



© 2025 Lyfey Technologies - Networking Training & Consultancy 

This website uses cookies and asks your personal data to enhance your browsing experience.

Ok, I agree [Privacy Policy](https://lyfeytech.co.ke/privacy-policy/)

[Privacy Policy](https://lyfeytech.co.ke/wpautoterms/privacy-policy/)
