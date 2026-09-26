---
id: collect-260926-mikrotik/mikrotik/an-issue-with-scripting-scheduler-for-container-mount-backup
title: "an-issue-with-scripting-scheduler-for-container-mount-backup"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/an-issue-with-scripting-scheduler-for-container-mount-backup.md
source_anchor: ""
source_lines: [1, 50]
sha256: 3c2dbd0ae0bf4152e795ddc308ff21b34bc0f2259c2ca348e4ce435f2b4d43d1
---

# an-issue-with-scripting-scheduler-for-container-mount-backup

Hi,

I have an issue using the scheduler for this script. Is this something that MT should fix? Or perhaps this is an access/permission issue?

This script works when running it manually or via the terminal. However, the scheduler can’t run the script.


**CLI**

```
/system scheduler
add interval=8h name=container-nginx-backup on-event="/system script run container-nginx-backup" policy=ftp,reboot,read,write,policy,test,password,sniff,sensitive,romon start-time=startup
add dont-require-permissions=no name=container-nginx-backup owner=user policy=ftp,reboot,read,write,policy,test,password,sniff,sensitive,romon source=":log info \"Nginx-Backup started\"\r\
    \n\r\
    \n:log info \"DATABASE\"\r\
    \n\r\
    \n/tool fetch address=\"ip\" src-path=\"/pcie1-part1/containers/mounts/nginx-proxy/data/database.sqlite\" user=\"user\" mode=\"ftp\" password=\"pw\" dst-path=\"mounts/nginx-proxy/data/database.sqlite\" upload\
    =yes;\r\
    \n\r\
    \n\r\
    \n:log info \"Nginx-Backup finished\""
```

             
            
           
          
            
            
              You stop the use of the file before backup it?

             
            
           
          
            
            
              Hi,

No, it’s a running container. Additionally, It did work when I ran the script manually without the need to stop the container.

**update**

I stopped the container first and tested the scheduler. It shows the same behavior as before.



**workaround**

Add another script to run the first script and then the scheduler to run the second script.
