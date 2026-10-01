---
id: collect-261001-general-networking/general-networking/questions-1605518-how-to-create-a-cronjob-that-persists-after-a-firmware-update-c22938a5
title: "questions-1605518-how-to-create-a-cronjob-that-persists-after-a-firmware-update--c22938a5"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-1605518-how-to-create-a-cronjob-that-persists-after-a-firmware-update--c22938a5.md
source_anchor: ""
source_lines: [1, 22]
sha256: 039631aeb2608f0dcf888cd8c450969b5b52ec1d48e877e2aff641c2920c3096
---

# questions-1605518-how-to-create-a-cronjob-that-persists-after-a-firmware-update--c22938a5

We have tens of Ubiquiti Edgerouter-X's on active on customer locations. They among other things responsible for pinging an IP each minute. This is done by a cronjob (set with crontab -e) and works fine as long as the firmware is not updated.
When the firmware is updated to a newer version, the cronjobs are "reset" to the default value. After some research i noticed that i should use the system task-scheduler instead of a cronjob. So i created an ping.sh file in /config/scripts/, set the permissions to 0755 and did an chmod +x ping.sh. This is the content of the ping.sh file (for test purposes)
#!/bin/bash
echo "Ping Pong!"
echo "Ping Pong!!" >> result.txt    
When i execute the script manually ./ping.sh the script works fine, it echo's Ping Pong! and writes Ping Pong!! to result.txt. When i try to execute the script with the system task-scheduler it does nothing. I tried 2 task scheduler configurations and restarted the Edgerouter after commit; save;:
show system task-scheduler task
 task ping-task {
     crontab-spec "* *  *   *   *"
     executable {
         path /config/scripts/ping.sh
     }
 }
and
show system task-scheduler task
 task ping-task {
     executable {
         path /config/scripts/ping.sh
     }
     interval 1m
 }
What do i wrong? I checked multiple topics about it on https://community.ui.com/ but could not find a solution...
