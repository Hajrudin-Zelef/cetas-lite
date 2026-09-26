---
id: collect-260926-mikrotik/mikrotik/firmware-version-discrepancy
title: "firmware-version-discrepancy"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/troubleshooting/firmware-version-discrepancy.md
source_anchor: ""
source_lines: [1, 76]
sha256: 2d9d7b073ed077c9e29d8fd6de975a89ffd95bea34b2b765e9e5ec6e8703fd8d
---

# firmware-version-discrepancy

Hello,

I downgraded to RouterOS 7.14.3 and wifi-qcom 7.14.3 because I was getting the infamous “SA Query Timeout” error. I marked version 7.15 for downgrade and rebooted. No warning showed up in the logs and the old packages are gone. Everything appears to have gone fine. When I export the config, the expected RouterOS 7.14.3 is displayed, but when I go to System → RouterBoard, the firmware version is still 7.15.

             
            
           
          
            
            
              routerboard firmware is separate from ROS version … similar to BIOS/UEFI being separate from OS (Windows, linux, …).

The difference between ROS and normal PCs is that routerboard firmware ships with ROS itself and it’s possible to set it to be upgraded automatically (property “auto-upgrade” under */system/routerboard/settings*). I guess ROS doesn’t auto-downgrade though, for that you have to run the procedure manually (*/system/routerboard/upgrade* followed by a reboot). Yup, it says ‘upgrade’ even though you’re downgrading … the ‘upgrade’ procedure actually installs whichever version is listed under *upgrade-firmware* over whatever is listed under *current-firmware* in output of command */system/routerboard/print* … and it doesn’t matter which number is bigger and which one is smaller.

BTW, routerboard firmware version is not vital. The firmware itself doesn’t change often … specially not on mature devices. However the version number changes with ROS … which makes it pretty confusing as to when upgrade is necessary. Until a few years ago, routerboot firmware had it’s own numbering making it much more obvious as to when FW upgrade was necessary. But I guess MT changed it to avoid confusion of users reporting routerboot FW version as ROS version (and vice versa) … similar confusion as you have now.

             
            
           
          
            
            
              I think it’s a good idea to set “/system routerboard settings set auto-upgrade=yes” for more automation on upgrade process.

             
            
           
          
            
            
              @mkx You have make this table before, maybe this can bring more light to the frequent confusing case.

And reuse/recycle some computer bytes 

http://forum.mikrotik.com/t/why-do-i-have-to-update-twice-each-time/176623/5

             
            
           
          
            
            
              If I may ask a question, what is (if any) the *needed* correspondence between firmware and RouterOS?

I mean OP ended up with having (newer) 7.15 firmware but running (older) RouterOS 7.14.3.

Is this only “cosmetic” or it worked because the two versions were very “near”?

In the PC world (exception made for possible regression bugs) the latest Bios/UEFI is usually recommended as a new version is normally released to fix some found bugs in booting or assigning resources, and is independent from OS installed, does the same happen for Routerboards?

             
            
           
          
            
            
              I think the RouterBoot version 6,xx,xx something could not boot version 7.xx.xx of RouterOS.

So you need to update to the latest 6 RouterOS, before you could upgrade to version 7 RouterOS.

And of course when you have the latest ver of 6 you press update firmware(RouterBoot), and then you upgrade the RouterOS.

Sorry Maybe this can be more confusing…

And i think also that RouterBoot(firmware as fwf file in the npk file) is backward compatible with ver 6 of RouterOS.

             
            
           
          
            
            
              
More or less yes. As I already mentioned, changes are more frequent on newly released device models, okder models hardly ever receive bug fix in routerboot … version increases never the less 

So for “mature” device models, changes in routerboot are mostly to introduce some new feature, requiring firmware support … such as device mode.
