---
id: collect-260926-mikrotik/mikrotik/questions-1301537-how-to-create-daily-data-download-limit-using-mikrotik-router-9d6265f4
title: "questions-1301537-how-to-create-daily-data-download-limit-using-mikrotik-router-9d6265f4"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-1301537-how-to-create-daily-data-download-limit-using-mikrotik-router-9d6265f4.md
source_anchor: ""
source_lines: [1, 34]
sha256: 68437e0a81c9da1a9cbe3338d0d01bb4b14de0ee5efc39b903c4c7cf8f600609
---

# questions-1301537-how-to-create-daily-data-download-limit-using-mikrotik-router-9d6265f4

I have one Mikrotik Router RB951G. I have two hotspots set up on these two networks:
- 192.168.1.0/24
- 192.168.2.0/24
I also have created two user profiles:
- Guest
- Employee
I have configured bandwidth limitation on the guest profile. On the other hand, I want to configure daily data download limitation on the employee profile.
I have run these two scripts to do so, but it seems not working.
- script for limiting data #Set your dowload limit in MegaBYTES!
:local downloadlimitmb "250"
### You will not need to edit anything below this line ###
:local downloadlimit  [($downloadlimitmb  * 1048576)]
:local counter
:local datadown
:local username
:local macaddress
:foreach counter in=[/ip hotspot active find where user~"^[T][-].{17}"] do={
:set $datadown [/ip hotspot active get $counter bytes-out]
:if ($datadown>$downloadlimit) do={
:set $username [/ip hotspot active get $counter user]
:set $macaddress [/ip hotspot active get $counter mac-address]
/ip hotspot user remove [/ip hotspot user find where name=$username profile=Employee]
/ip hotspot user add name=$username limit-bytes-out=$downloadlimit mac-address=$macaddress profile=Employee
/ip hotspot active remove $counter
:log info "Force logout on user: $username - Reached download quota"
}}
- script for clearing user counters :log info "Checking Users"
:local counter
:foreach counter in=[/ip hotspot user find profile="Employee" ] do={
/ip hotspot user remove $counter
}
Finally I have added two schedules:
- to trigger datalimit and runs every minute.
- to trigger clearcounter and runs every 24 hours. one to run every 24 hours
