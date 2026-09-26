---
id: collect-260926-mikrotik/mikrotik/backup-restore-2
title: "backup-restore-2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters", "reasoning"]
source: docs/RAG/lot-mikrotik/forum/misc/backup-restore-2.md
source_anchor: ""
source_lines: [1, 139]
sha256: 43cbad81437f6554d9579e09a6d2ea39e0b11e68489782d11c28092936619a26
---

# backup-restore-2

Hello everybody!

Id like to share my experience and ask you for solution.

Ive learned few years ago, if you do a backup from the terminal, this way: /export file=filename , than i can restore my config on an another hardware (by that i mean another mikrotik router), while if you do “files → backup” with or without enryption, than you can only restore this backup file to the same hardware where it was made.



Now i wanted to restore a “*.rsc" file (the terminal export version), it says file not found. When i do it from the console, /system backup → load → write in the files name → i get the same error, file not found (6). While the file is there, its valid, etc. After that i tried to restore my "*.backup” file and it worked, wow, i could restore the *.backup file to another hardware.

My question → what are the “rsc” files intended for NOW and what are the backup files (seems to be obvious, but i will ask now ) intended for NOW?

Did this change or did i know this wrong the whole time? I do both backup versions before a fw updgrade and i think i will too, in the future, but id like to clear this misunderstanding in my head.

THANK YOU IN ADVANCE!

             
            
           
          
            
            
              For rsc file, use /import instead of /system backup. Nothing changed in terms of backup and export usage, you should not use backup to restore it on another machine, even if it works.

             
            
           
          
            
            
              
Hi!

Thanks for you answer.

I made on an another router (another hw) a backup using /export file=202005    → this created a file named “202005.rsc”. (which i can open and check settings via notepad)

If i upload this to another router, and type “/import 202005” → error message: there is no file, if i write “/import 202005.rsc” it says ->line 181 column 5 - end of command expected. This is a working and fresh exported config. I deleted the line 181, same error. Now i backed up another mikrotik router using “/export file=filename” and i got now a message (when trying to restore on an another hw) → there is already another interface with this name. I renamed the only one custom named interface, not working. Now i did a factory reset with no def. config, its a plain mikrotik os now, nothing configured, (after uploading my freshly exported RSC) i type in terminal /import 202005.rsc, now it says no such item. Okay, lets try again, on the second run it says failure: already have interface with such name.





So, what is the working OFFICIAL version, where i can back up all the settings and restore it to another hardware?

             
            
           
          
            
            
              
The “official” version is to run *system reset-configuration run-after-reset=path/to/the.exported.file.rsc* (possibly with *no-defaults* and *keep-users=yes*). The file must not disappear after reboot (so on some models, it must be placed in the *flash/* directory), as the old configuration is first deleted, then the router is rebooted, and then the file is read in.

However, there are still things to address:

- if the .rsc file refers to interfaces or other hardware which doesn’t exist on the machine where you import it, the imported script will stop on that error
- if the .rsc file is generated on another RouterOS version than where it is imported, some commands/parameters may be incompatible, the script will also stop with an error
- before placing the file to the router, add *delay 1m* as its first line, because it is typically processed too early, before the interfaces are recognized and added, so the script fails even on interfaces which are actually available on the router

So unless you import at the same model running the same RouterOS version, it needs some handwork in any case.

             
            
           
          
            
            
              
Hi! Thanks for the quick anser!



This sounds painfully complicated, but hey, its okay as long it works, but i cannot get it to work.  Could you please specify the exact command i have to run? From GUI (reset, run after.blabla, no defconf,keepuser) does not work. When using terminal, what you wrote does not work. “system reset-configuration run-after-reset=202005.rsc” (where my rsc file is called 202005.rsc) does not work, if i go to LOG, i cannot see anything realted to the “restoration”. If i try “system reset-configuration run-after-reset=202005.rsc keep-users=yes no-defaults=yes” it does not work. Is there a way to see what the problem is while restoring?

I try to get a config from a “hap ac 2” to a “hap ac lite”, both have 5 physical LAN ports, NO SFP and two (disabled) physical wireless cards (2.4 and 5ghz), so the interfaces are totally the same in the physical way, but when i reset a config, there should be no “sofware” related problem, cuz its resetted to default.

Its a really simple config, there should be no problem i think, there is one WAN (dhcp client - eth1) and the rest of the ports is in a  bridge and there isnt even a dhcp server.

             
            
           
          
            
            
              hAP ac lite is exactly one of the models where the script must be stored as *flash/202005.rsc* and you have to refer to it like this (*run-after-reset=flash/202005.rsc*). But as you seemingly have the router on the table, so you can access it even after using *no-defaults=yes*, you may then import the file manually and see on what it fails.

             
            
           
          
            
            
              
Would you please help me?

It says “expected end of command” (line 181 column 5)

So i would say there is no problem, but if there would be a syntax error, it should begin much more faster, adding to address list with the same syntax begins in my config at line 46 already.

             
            
           
          
            
            
              I agree with your reasoning that if the problem was the format of the line, it would fail already on the first line with that same format. So either the import is too fast, or the line numbering is interpreted differently. I’d suggest to import the file manually by pasting it to the command line window block by block to see what is the real problem.

             
            
           
          
            
            
              
/port

set 0 name=serial0



this is where the whole restoration process breaks. what could this be? is there a serial port on a hap ac2? i dont remember, but i didnt take down the cover.

             
            
           
          
            
            
              You may have connected an USB-to-serial converter in the past so the configuration has been added. Just remove this block from the configuration and try to import the file without it.

             
            
           
          
            
            
              
Thanks! USB-to-Serial? I dont think so, but yes, after finding this line and removing it, with /import flash/“filename.rsc” it worked! Thanks for your time! Now if recovery is not working, first thing i will check is serial port xd
