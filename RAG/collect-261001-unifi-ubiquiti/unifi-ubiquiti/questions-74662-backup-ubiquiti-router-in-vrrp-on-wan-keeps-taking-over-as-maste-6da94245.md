---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-74662-backup-ubiquiti-router-in-vrrp-on-wan-keeps-taking-over-as-maste-6da94245
title: "questions-74662-backup-ubiquiti-router-in-vrrp-on-wan-keeps-taking-over-as-maste-6da94245"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-74662-backup-ubiquiti-router-in-vrrp-on-wan-keeps-taking-over-as-maste-6da94245.md
source_anchor: ""
source_lines: [1, 62]
sha256: b11f49224bbbfbc3dc60a997a44fd40f7862d51c317e68590a685a0b9b538129
---

# questions-74662-backup-ubiquiti-router-in-vrrp-on-wan-keeps-taking-over-as-maste-6da94245

I have two EdgeRouter 4s. I have two /28 blocks of public IP addresses. The EdgeRouters each have an uplink to my ISP, and can ping each other.
I'm trying to setup VRRP, and used this resource as a guide: https://help.ui.com/hc/en-us/articles/204962174-EdgeRouter-Virtual-Router-Redundancy-Protocol-VRRP-
I have removed the two middle octets, but below is my relevant config for both routers. You'll notice that I've assigned a static out of the same /28 to both routers, and then configured the rest of the addresses as virtual.
Both routers can ping each other. 209.x.x.82 can ping 209.x.x.83, and vice-versa.
Primary Router's Config:
set interfaces ethernet eth0 address 209.x.x.82/28
set interfaces ethernet eth0 vrrp vrrp-group 3 advertise-interval 1
set interfaces ethernet eth0 vrrp vrrp-group 3 preempt true
set interfaces ethernet eth0 vrrp vrrp-group 3 priority 200
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.84/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.85/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.86/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.87/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.88/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.89/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.90/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.91/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.92/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.93/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.94/28
Backup Router's Config:
set interfaces ethernet eth0 address 209.x.x.83/28
set interfaces ethernet eth0 vrrp vrrp-group 3 advertise-interval 1
set interfaces ethernet eth0 vrrp vrrp-group 3 preempt true
set interfaces ethernet eth0 vrrp vrrp-group 3 priority 100
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.84/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.85/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.86/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.87/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.88/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.89/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.90/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.91/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.92/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.93/28
set interfaces ethernet eth0 vrrp vrrp-group 3 virtual-address 209.x.x.94/28
Here is the primary router's vrrp summary for eth0
$ show vrrp summary
                VRRP    Addr                    Interface       VRRP            
Interface       Group   Type    Address         State           State           
---------       -----   ----    -------         -----           -----           
eth0            3       vip     209.x.x.84/28up              master          
                        vip     209.x.x.85/28
                        vip     209.x.x.86/28
                        vip     209.x.x.87/28
                        vip     209.x.x.88/28
                        vip     209.x.x.89/28
                        vip     209.x.x.90/28
                        vip     209.x.x.91/28
                        vip     209.x.x.92/28
                        vip     209.x.x.93/28
                        vip     209.x.x.94/28
The problem Every time I put the Backup router's config into place, the backup router makes itself the master - when at the same time, the Primary continues to show it is master. So BOTH routers think they are master.
I have tried modifying the backup router to NOT preempt (set interfaces ethernet eth0 vrrp vrrp-group 3 preempt false). Additionally, as soon as I've committed the changes to the backup router, I've tried to force clear its vrrp state, and force it into backup mode, with preempt set to false:
So on the Backup Router, after applying the above config, I run:
$ clear vrrp master interface eth0 group 3
Forcing vyatta-eth0-3 to BACKUP...
And then I run show vrrp summary and I see that Group 3 is in a VRRP State of Backup.
But then I wait a few seconds, and run the same show vrrp summary command on the Backup router, and it is in a state of Master again.
How can I troubleshoot why my Backup router is continuing to try to make itself Master, even though each router can ping each other, and even when I've tried to set preempt to false?
For context, it should be noted my upstream ISP uses Cisco arp caching, and if a MAC address gets into their arp cache, I have to wait 3 hours for it to clear before the IP will start working again. That's quite annoying for my own troubleshooting purposes.
Are there additional firewall flows I have to setup on either of the routers in order to allow vrrp communication to work (i.e. do I need anything besides ICMP ping)? Is it possible that my ISP is doing some sort of filtering on their switch? Any other suggestions?
