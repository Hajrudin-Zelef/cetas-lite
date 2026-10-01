---
id: collect-261001-huawei/huawei/pat-or-no-pat-source-nat-on-huawei-usg6000-c77974df
title: "PAT or no-PAT – source NAT on Huawei USG6000"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/pat-or-no-pat-source-nat-on-huawei-usg6000-c77974df.md
source_anchor: ""
source_lines: [1, 155]
sha256: ee2e75bf2fd97f9770047d6e2175078ffc69332f62243c092c61ff9f205c9c64
---

# PAT or no-PAT – source NAT on Huawei USG6000

(Source archive : http://labnario.com/pat-or-no-pat-source-nat-on-huawei-usg6000/ — capture 20260617131433)

__

[__](https://www.facebook.com/Huawei-From-Scratch-376204559133899/ "Facebook")[__](https://twitter.com/labnario "Twitter")[__](https://foursquare.com/mo3aser "Foursquare")

Wednesday , June 17 2026

__

[__](https://www.facebook.com/Huawei-From-Scratch-376204559133899/ "Facebook")[__](https://twitter.com/labnario "Twitter")[__](https://foursquare.com/mo3aser "Foursquare")

##  [ ![Labnario](http://labnario.com/wp-content/uploads/2017/06/logo_labnario.png)**Labnario Huawei From Scratch** ](http://labnario.com/ "Labnario")

  * [HOME](http://labnario.com/)
  * [HUAWEI CHEAT SHEETS](http://labnario.com/huawei-cheat-sheets/)
  * [ABOUT AUTHOR](http://labnario.com/about/)



[Home](http://labnario.com) / [IP Services](http://labnario.com/category/ip-services/) / PAT or no-PAT - source NAT on Huawei USG6000

# PAT or no-PAT - source NAT on Huawei USG6000

__21 Mar, 2024 __[IP Services](http://labnario.com/category/ip-services/), [Security](http://labnario.com/category/security/), [Video](http://labnario.com/category/video/) __[Leave a comment](http://labnario.com/pat-or-no-pat-source-nat-on-huawei-usg6000/#respond)

If you, for some reason, cannot use **easy-ip NAT** , you can use source NAT with **NAT address pool**. Depending on how many public IP addresses you have got, you can configure no-PAT option, when only IP address is translated or you can set PAT, in other words NAT with port translation to assure LAN users accessing Internet. Details in the video 😉

## USG firewall configuration script:

`#  
dhcp enable  
#  
interface GigabitEthernet1/0/0  
undo shutdown  
ip address 10.0.0.1 255.255.255.0  
service-manage ping permit  
dhcp select interface  
dhcp server excluded-ip-address 10.0.0.100  
dhcp server static-bind ip-address 10.0.0.200 mac-address 5489-98b4-6a79  
dhcp server dns-list 10.0.0.100  
#  
interface GigabitEthernet1/0/2  
undo shutdown  
ip address 5.0.0.2 255.255.255.252  
#  
firewall zone trust  
set priority 85  
add interface GigabitEthernet1/0/0  
#  
firewall zone untrust  
set priority 5  
add interface GigabitEthernet1/0/2  
#  
ip route-static 0.0.0.0 0.0.0.0 5.0.0.1  
#  
nat address-group SOURCE-NAT 0  
mode pat  
route enable  
section 0 6.6.6.0 6.6.6.0  
OR  
nat address-group SOURCE-NAT 0  
mode no-pat global  
route enable  
section 0 6.6.6.0 6.6.6.1  
#  
security-policy  
rule name ALLOW  
source-zone local  
destination-zone trust  
destination-zone untrust  
action permit  
rule name NAT_EASY  
source-zone trust  
destination-zone untrust  
source-address 10.0.0.0 mask 255.255.255.0  
action permit  
#  
nat-policy  
rule name SOURCE-NAT  
source-zone trust  
destination-zone untrust  
source-address 10.0.0.0 mask 255.255.255.0  
action source-nat address-group SOURCE-NAT`

Share

  * [__ Facebook](http://www.facebook.com/sharer.php?u=http://labnario.com/?p=2014)
  * [__ Twitter](https://twitter.com/intent/tweet?text=PAT+or+no-PAT+%E2%80%93+source+NAT+on+Huawei+USG6000&url=http://labnario.com/?p=2014)
  * [__ LinkedIn](http://www.linkedin.com/shareArticle?mini=true&url=http://labnario.com/?p=2014&title=PAT+or+no-PAT+%E2%80%93+source+NAT+on+Huawei+USG6000)



Tags [Huawei CLI](http://labnario.com/tag/huawei-cli/) [Huawei USG](http://labnario.com/tag/huawei-usg/) [Huawei VRP](http://labnario.com/tag/huawei-vrp/) [NAT on Huawei USG](http://labnario.com/tag/nat-on-huawei-usg/) [PAT or no-PAT](http://labnario.com/tag/pat-or-no-pat/)

[Previous Easy-IP source NAT on Huawei USG firewall](http://labnario.com/easy-ip-source-nat-on-huawei-usg-firewall/)

[Next Destination NAT on Huawei USG firewall](http://labnario.com/destination-nat-on-huawei-usg-firewall/)

### Leave a Reply [Cancel reply](/pat-or-no-pat-source-nat-on-huawei-usg6000/#respond)

Your email address will not be published. Required fields are marked *

Comment *

Name *

Email *

Website

Save my name, email, and website in this browser for the next time I comment.

Δ

[__](https://www.facebook.com/Huawei-From-Scratch-376204559133899/ "Facebook")[__](https://twitter.com/labnario "Twitter")[__](https://foursquare.com/mo3aser "Foursquare")

__

#### Categories

  * [Basic Configuration](http://labnario.com/category/basic-configuration/)
  * [Cheat Sheets](http://labnario.com/category/cheat-sheets/)
  * [Command Line](http://labnario.com/category/command-line/)
  * [Ethernet](http://labnario.com/category/ethernet/)
  * [FAQ](http://labnario.com/category/faq/)
  * [General](http://labnario.com/category/general/)
  * [How To](http://labnario.com/category/how-to/)
  * [IP Routing](http://labnario.com/category/ip-routing/)
  * [IP Services](http://labnario.com/category/ip-services/)
  * [Multicast](http://labnario.com/category/multicast/)
  * [QoS](http://labnario.com/category/qos/)
  * [Reliability](http://labnario.com/category/reliability/)
  * [Security](http://labnario.com/category/security/)
  * [System Management](http://labnario.com/category/system-management/)
  * [Uncategorized](http://labnario.com/category/uncategorized/)
  * [Video](http://labnario.com/category/video/)
  * [VPN](http://labnario.com/category/vpn/)
  * [WAN](http://labnario.com/category/wan/)



#### Recent Posts

  * [HCIA Datacom Course Summer Sale](http://labnario.com/hcia-datacom-course-summer-sale/)
  * [What's new in the New Year?](http://labnario.com/whats-new-in-the-new-year/)
  * [New Year Sale](http://labnario.com/new-year-sale/)
  * [Huawei AR router USB-based deployment (ZTP)](http://labnario.com/huawei-ar-router-usb-based-deployment-ztp/)
  * [Feel invited to the Huawei HCIA course](http://labnario.com/feel-invited-to-the-huawei-hcia-course/)



Design By: [ThemeTen](themeten.com)
