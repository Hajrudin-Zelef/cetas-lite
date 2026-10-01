---
id: collect-261001-fortinet/fortinet/t5-support-forum-fortigate-firmware-upgrade-via-fortimanager-script-td-p-404340-ddd00077
title: "t5-support-forum-fortigate-firmware-upgrade-via-fortimanager-script-td-p-404340-ddd00077"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2025-08-06"]
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-support-forum-fortigate-firmware-upgrade-via-fortimanager-script-td-p-404340-ddd00077.md
source_anchor: ""
source_lines: [1, 8]
sha256: 5a1f401611fdf49b10af24d275a4bcfd8a00d4e0dd382414a653f088a63e23b5
---

# t5-support-forum-fortigate-firmware-upgrade-via-fortimanager-script-td-p-404340-ddd00077

I need to upgrade a group of not licenced fortigates, I tried to run a script that execute a restore image tftp but after downloading and validate the image from the tftp server the script stops with errors and doesnt do anything.
I would like to know if this upgrade process is posible and if there are alternatives to do this upgrade on not licenced fortigates. Maybe a script on an external server that execute a ssh on every fortigate.
I found a way to do the upgrade. Executing the image restore from FortiManager cause an error or some kind of incompatibility I dont know why. So the solution I found was create an automation with the image restore as the action.
config system automation-trigger edit "upgrade" set trigger-type scheduled set trigger-frequency once set trigger-datetime 2025-08-06 00:00:00 next end
config system automation-action edit "cmd" set action-type cli-script set script "execute restore image tftp <image> <IP>" set accprofile "super_admin" next end
config system automation-stitch edit "upgrade" set trigger "upgrade" config actions edit 1 set action "cmd" set required enable next end next end
Thanks to all!
Did this topic help you find an answer to your question?
