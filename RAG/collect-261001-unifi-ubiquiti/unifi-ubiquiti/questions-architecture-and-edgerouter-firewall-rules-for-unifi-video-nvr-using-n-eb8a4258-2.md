---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-architecture-and-edgerouter-firewall-rules-for-unifi-video-nvr-using-n-eb8a4258-2
title: "questions-architecture-and-edgerouter-firewall-rules-for-unifi-video-nvr-using-n-eb8a4258"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-architecture-and-edgerouter-firewall-rules-for-unifi-video-nvr-using-n-eb8a4258.md
source_anchor: ""
source_lines: [137, 199]
sha256: c78c19cf3eaeba53f75287474aff4add52743fcb2492b0649a01504aacf1259c
---

# questions-architecture-and-edgerouter-firewall-rules-for-unifi-video-nvr-using-n-eb8a4258

                 name SURVEILLANCE_IN
             }
             local {
                 name SURVEILLANCE_LOCAL
             }
         }
         mtu 1500
     }
     vif 40 {
         address 192.168.40.1/24
         description IoT
         firewall {
             in {
                 name IOT_IN
             }
             local {
                 name IOT_LOCAL
             }
         }
         mtu 1500
     }
 }
 ethernet eth1 {
     address dhcp
     description WAN
     duplex auto
     firewall {
         in {
             name WAN_IN
         }
         local {
             name WAN_LOCAL
         }
     }
     speed auto
 }
 ethernet eth2 {
     address dhcp
     description "WAN 2"
     duplex auto
     firewall {
         in {
             name WAN_IN
         }
         local {
             name WAN_LOCAL
         }
     }
     speed auto
 }
 loopback lo {
 }
[edit]
admin@EdgeRouter-Lite-3-Port#
exit
Are there any holes that jump out at ya?
--Thanks,
--Jim
P.S. After further research on the side issue, I'm 100% convinced that the "Space To Keep Free" setting on UNiFi-Video (and surely on UniFi-Protect as well) can only screw you (i.e. risk write fails due to insufficient space) when using a NAS with a shared folder quota. The best answer was to remove the quota and use time-based purging as a very rough proxy for keeping the security videos from using too much of the NAS storage space. There is probably a custom solution to be had, but that is for another thread and another day.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
