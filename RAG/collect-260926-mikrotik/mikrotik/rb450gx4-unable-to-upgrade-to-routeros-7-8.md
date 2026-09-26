---
id: collect-260926-mikrotik/mikrotik/rb450gx4-unable-to-upgrade-to-routeros-7-8
title: "rb450gx4-unable-to-upgrade-to-routeros-7-8"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "nand"]
source: docs/RAG/lot-mikrotik/RouterOS/rb450gx4-unable-to-upgrade-to-routeros-7-8.md
source_anchor: ""
source_lines: [1, 88]
sha256: 78c86574ec304b34218b94d4c4b062555915f42cb85792e5f95a23d12df2a269
---

# rb450gx4-unable-to-upgrade-to-routeros-7-8

Hi All,

I have two RB450Gx4 which were on RouterOS 6.48.6. When attempting to update to RouterOS 7.8 via WinBox both devices began to repeatedly reboot.

Using NetInstall I installed routeros-7.8-arm.npk however both devices continued to reboot.

Then attempted to install routeros-6.49.7-arm.npk via NetInstall and device successfully boots and starts up. At this stage I also upgraded the routerboard.

Then attempted to install 7.8 again, and device went back to reboot loop. Installing 7.9 also resulted in the reboot loop.

Logging into RouterBOOT shows the following

```
RouterBOOT booter 6.49.7
RB450Gx4
CPU Frequency: 716MHz
Memory size: 1024MiB
NAND size: 512MiB
Press any key within 2 seconds to enter setup..
loading kernel... OK
setting up elf ... OK
jumping to kernel code
jumping to kernel
opendir: No such file or directory
opendir: No such file or directory
ERROR: no system package found!
[   6.994607] Kernel panic - not syncing: Attempted to kill init! exitcode=0x00000000
[   6.994653] CPU: 1 PID: 1 Comm: init Not tainted 5.6.3#2
```

Over my multiple attempts, it appears that Router OS 6.48 and 6.49 loads successfully on the device however 7.8 and 7.9 do not and provide the above error.

We have also recently configured 30+ of the RB450Gx4 onto RouterOS 7.8 with only these two devices having an issue.

Any advice on how I can proceed or could it be a NAND fault?

Thanks

Scott

             
            
           
          
            
            
              You can try to upgrade routerboot to version 7.8 while running ROS 6.48.6, it might make some difference.

The procedure is roughly this:

1. if you don’t have the npk file lying around, download main package file for correct architecture from download archive
2. verify actual firmware type used on offending devices
 use command*/system routerboard print* and have a look at contents of property*firmware-type* (FWtype)
3. open routeros main package NPK file using 7z file archiver. Browse through to */etc/* and extract file named*-.fwf* . The fwf file name has to exactly match the FWtype, determined in previous step, … don’t let yourself get distracted by some single-character difference (e.g. there’s ipq4000-7.8.fwf and ipq4000**L** -7.8.fwf).
4. upload fwf file to device into whatever is root of accessible storage (just like one would do with npk files)
5. go into */system routerboard* … executing*print* should show you “upgrade-firmware” to be available - with version matching the uploaded fwf file. If this is so, execute*upgrade* and reboot device

If you have downloaded main ROS package of correct version and architecture before, then you can skip the step #1 and proceed starting from step #2.

After verifying that device indeed bots with new routerboot, try to netinstall v7 again. Hopefully the above procedure will make a difference.

             
            
           
          
            
            
              Thanks for the suggestions. I followed your instructions and successfully updated to RouterBOOT 7.8.

Then attempted to upgrade to RouterOS 7.8 but again begins to loop with the same above error (except shows RouterBOOT booter 7.8 )

             
            
           
          
            
            
              I am up to ver7.10 beta5, so its possible, but dont remember how I got there… I always do routerboard update after regular firmware update.

             
            
           
          
            
            
              Regular firmware upgrade is preferred method. However, with large intended ROS version jump (e.g. from 6.40 or sonething to 7.x) it doesn’t work. IIRC most devices got routerboot upgrade somewhere around 6.47 which actually allows booting ROS 7 kernel. Netinstall doesn’t touch routerboot (if I’m not mistaken) and both netinstall nor normal upgrade via ROS commands can end up in boot loop. I don’t know if there’s any relevant difference between routerboots 6.49.7 and 7.8, so my suggestion was a “fishing in the mud” idea and sadly it didn’t work.

However, having recent routerboot (7.8 ) still barfing on booting ROS v7.8 indicates possible problem with flash storage.
