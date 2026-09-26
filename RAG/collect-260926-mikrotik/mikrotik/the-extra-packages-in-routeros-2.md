---
id: collect-260926-mikrotik/mikrotik/the-extra-packages-in-routeros-2
title: "NAME                                                                        VERSION          SCHEDULED"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/the-extra-packages-in-routeros.md
source_anchor: ""
source_lines: [202, 350]
sha256: 3f4eaddb7a33531f6ae9c6be620682b7fcaf0b84646b74625c9303f820f0f2f4
---

# NAME                                                                        VERSION          SCHEDULED

And only if the packages are the correct arch (obviously).

             
            
           
          
            
            
              
I tried both variants, but neither worked 

I think that even the standard downgrade function isn’t working. I have currently this version installed:

/system package update print

channel: stable

installed-version: 6.46.6

To which version will it downgrade if one calls the normal downgrade function (ie. w/o supplying manually an .npk)?

I’ll now try this, but I think I already tried, and it didn’t function.

Update:

Like I said: even the standard downgrade function does not work, as the version is the same as before, after the reboot.

             
            
           
          
            
            
              Discerning (astute) device, it does not like you!    Has Mikrotik developed  networking equipment similar to brooms in Harry Potter?  

             
            
           
          
            
            
              
Then you are doing something wrong.

Read the fine manual again and retry until you succeed.

Only with exercise you will learn how to properly use ROS.

             
            
           
          
            
            
              
[admin2@CRS125] /ip service> /system package update print

channel: stable

installed-version: 6.46.6

[admin2@CRS125] /ip service> /system package downgrade; /system reboot;

Router will be rebooted. Continue? [y/N]:

y

system will reboot shortly

Reboot, yes? [y/N]:

Received disconnect from 192.168.128.253: 11: shutdown/reboot

And after reboot and new login:

[admin2@CRS125] > /system package update print

channel: stable

installed-version: 6.46.6



 Didn’t work as can be seen.

Man, it seems another rocket science to downgrade this device 

IMO there is simply a bug in the downgrade function of this software. Logic & the countless tries with negative outcome today tells me  

             
            
           
          
            
            
              Perhaps operator error--------- bug incurred during conception??

             
            
           
          
            
            
              
You post too much, you read too little. You have half my posts in 2 months and I’ve been here for 15 years.

RTFM and then come back with proper results.

             
            
           
          
            
            
              Does anybody else have a constructive contribution to make (ie. help, tips, hints) for this problem of downgrading to an older RouterOS version?

             
            
           
          
            
            
              You mean constructive, like telling you exactly how to downgrade and you ignoring it?

             
            
           
          
            
            
              It’s childs play. And there is no bug. Processes that are very old and are used daily usually don’t contain bugs anymore.

As programmer you see bugs everywhere , like someone with a hammer sees nails everywhere.

Upload the needed npk to the files directory. Either in RAM (above /flash, the preferred place) or in a top directory in /flash.

**Make sure it is the correct npk (hardware variant) and it is the only instance of that npk in the files directory.**

The hardware variant selection can be confusing.  For CRS125 its the MIPSBE variant. If you load the ARM variant of the CRS325 it will just be deleted.

Check the size of the uploaded npk to verify it is complete. (incomplete files will not install, if there are multiple (incomplete) files it will try one and possibly delete another. So it tries the bad one, fails, and then deletes the good one. Next upload same story.)http://forum.mikrotik.com/t/issues-installing-the-dude/139569/1

If it is a higher version number : take upgrade. If it is a lower version number : take downgrade.  (The channel does not matter in this. E.G. I like the 6.45.6 from the stable channel)

The only, very classical, hurdle in upgrading is that sometimes you need to do an intermediate upgrade step (to be able to upgrade the bootloader firmware).

(I think it was from 6.44 towards 7.0beta … http://forum.mikrotik.com/t/special-instructions-for-installing-routeros-7-beta2/133553/1)

I have never seen a case that needs an intermediate downgrade step.
