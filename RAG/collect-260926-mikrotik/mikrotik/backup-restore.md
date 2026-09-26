---
id: collect-260926-mikrotik/mikrotik/backup-restore
title: "backup-restore"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/misc/backup-restore.md
source_anchor: ""
source_lines: [1, 151]
sha256: c7e10d9fcf4e57e3aaca6217bb7d6bb83c347b12f187a0d07d5de5241feb15d9
---

# backup-restore

Hi,

I’m using Mikrotik 3.28V with level-5 licence on IBM server which has 4 ethernet ports, i have taken backup and restored it in Mikrotik 450G 4.1V router board which is a level-4 licence and has 5 ethernet ports. After connecting to 450G router i’m getting access from internet but DHCP customers who are connected in lan are not able to login and some of them are not getting IP address. Is this compatability issue or is there any procedure to restore the backup in 450g router and use it in live with the same configuration.

Initially when i restored the backup in 450g router Vlan’s were not taking the correct ethernet port so i manually selected the ehternet port which was configured in IBM server.

Is this problem because of version mismatch or licence mismatch or beacuse of port mismatch as IBM is having 4 ethernet ports and 450g is having 5 ethernet ports. If any solution is there for this issue please let me know.

Thanks,

Srinivas

             
            
           
          
            
            
              Hi,

you can’t do this..

You need to make in terminal:

export file=backup.rsc

then in files you’ll have all the configuration in txt format.. remove the default lines, the MAC adresses and make the necesary changes. Then copy paste to your RB4xx.



I hope this helps.

             
            
           
          
            
            
              Hi,

Can you explain it more precisely.

Thanks,

Srinivas

             
            
           
          
            
            
              It is strongly recommended to load backup only to the same model router it was taken from. You can follow Martín’ s advice and edit export file. Or you can simply export only configuration that you need, for example, **/ip firewall export**; **/ip dns export**; **/ip address export** etc. Then merge it into one file and import it to the new router.

             
            
           
          
            
            
              Hi,

Thanks for the information, when i’m trying to export file backup.rsc that particular file size is much lesser and with that file i’m not able to restore. I think i’m missing something, if you don’t mind can you give me the detailed procedure.

Thanks,

Srinivas

             
            
           
          
            
            
              Don’t use backup and restore. Use import and export instead. On the source computer

/export file=mysetup.rsc

on the destination computer

/import mysetup.rsc

             
            
           
          
            
            
              With export, you obtain a .RSC file, this is the script commands to configure your routerboard.

You need to enter via terminal and copy/paste the text in the RSC file. You can open the RSC file with notepad.

             
            
           
          
            
            
              
REMEMBER TO EDIT mysetup.rsc AFTER doing the import command. DELETE al MACs and all lines with default configurations or with different interfaces in the new routerboard.

             
            
           
          
            
            
              
Should be after doing export, not import.

Router A:

Router B:

- Import edited export file from Router A

 
            
           
          
            
            
              Hi,

I have tried exporting each fie DHCP, Firewall, Routes, Hotspot and merged them and imported in the new router. After importing i’m getting the below mentioned error message.

“Script file loaded successfullyinput does not match any value of address-pool

interrupted”

I don’t know where i’m going wrong, can you please give me an example.

Thanks,

Srinivas

             
            
           
          
            
            
              sri,

When you export a configuration, the output is generated as if you typed the commands at the terminal.  In your exported config, there may be dependencies (like IP → Pool used in DHCP, or interface names used in Firewall) that have to be met before the export can successfully complete.

So, for DHCP Server, here are the dependencies (as far I know):  Interface and IP Pool.  This means, you’ll have to import your /interface and /ip pool settings before importing /ip dhcp-server settings.

Also, as mentioned previously, you can edit the exported config to suite your needs (ex. change pool name, interface name, etc…)

Hope this clears things up a bit.
