---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-lac163-schedule-script-wont-work-3479ac69
title: "r-mikrotik-comments-lac163-schedule-script-wont-work-3479ac69"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/r-mikrotik-comments-lac163-schedule-script-wont-work-3479ac69.md
source_anchor: ""
source_lines: [1, 37]
sha256: 41fcd7c2a9329135409590d117f5aa3cbc0e84db9312a09b588405ef640e1113
---

# r-mikrotik-comments-lac163-schedule-script-wont-work-3479ac69

schedule script wont work 
        
    Hi !
I have two scripts:
      /system script print
    
      0 name="enable_wlan" owner="user" policy=ftp,reboot,read,write,policy,test,password,sniff,sensitive,romon dont-require-permissions=no run-count=4 source=caps-man interface enable cap1
    
      1 name="disable_wlan" owner="user" policy=ftp,reboot,read,write,policy,test,password,sniff,sensitive,romon dont-require-permissions=no run-count=0 source=caps-man interface disable cap1
    
but only one script start but the second wont run with the scheduler: (disable work enable not )
      /system scheduler print
    
      Flags: X - disabled
    
      # NAME START-DATE START-TIME INTERVAL ON-EVENT RUN-COUNT
    
      0 01_EnableWLAN jan/31/2021 20:31:00 0s /system script run enable_wlan
    
      1 02_DisableWLAN jan/31/2021 22:00:00 0s /system script run disable_wlan
    
I dont understand why ?
If i manually start the script it work without a problem
Please help !
Thanks
regards
Section des commentaires
Check the date and time on the router. If you haven't setup the SNTP client chances are the date is currently default boot date.
You've also got the interval set to 0s which means it'll run at the specified time and date once, you want the interval to be 24 hours to run once a day at the specified time.
already set but Thanks
It’s not Jan 31st.
You have it set to start on a specific time and date. With no recurrence.
i really need to set the right start date crazy -> now i see also the next run scheduler :-)
THANKS !!!
Make sure scripts and scheduler entries have the same permissions.
i have already controll this both scripts -> the problem is if i set nextday start time because 7:30 in the morning is the next day :-) Crazy Thanks !
You may want to try formatting source=(/caps-man interface.....). Another thing to try is copy your script in the terminal and look for highlighted errors.
