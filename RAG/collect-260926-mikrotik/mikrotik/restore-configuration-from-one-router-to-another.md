---
id: collect-260926-mikrotik/mikrotik/restore-configuration-from-one-router-to-another
title: "restore-configuration-from-one-router-to-another"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/restore-configuration-from-one-router-to-another.md
source_anchor: ""
source_lines: [1, 15]
sha256: 28c97bd33057a0253ee2c960802958d5e807e86f86590f1827bb57ee5b33aded
---

# restore-configuration-from-one-router-to-another

Hi,

I have one question . Can I restore a backup from one mikrotik route to another  if they are not the same models?

             
            
           
          
            
            
              The best practice for transferring configuration between models is to use the “/export” CLI command (you can use the file= option to export to a file on the router that can be downloaded), and then either paste or “/import” the setup into the new router.

In general, just taking a backup & restoring it is likely to have problems if the two routers are not the same model AND the same RouterOS version.  Using /export and /import will work as long as the two routers are running the same RouterOS version and have similar interface counts (if the source has 10 Ethernet and 1 wireless, but the target has 5 Ethernet and 0 wireless, the /export and /import will not work either).

In the case of different RouterOS versions, some CLI commands may have changed (especially if the source is v5 and the target is v6, or if the source is newer than the target).  If the source is newer than the target, you should upgrade the target router before importing.  If the target is newer, you may need to use the “verbose” option within your /import command, so that it will tell you which line it failed on and what the problem was.  After you fix the error (you can edit the /export file with any text editor) and re-upload the file, you can then re-run /import using the “from-line” option to continue from where it left off.
