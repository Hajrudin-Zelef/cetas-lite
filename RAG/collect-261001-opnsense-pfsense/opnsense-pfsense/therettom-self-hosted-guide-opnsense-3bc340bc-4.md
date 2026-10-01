---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/therettom-self-hosted-guide-opnsense-3bc340bc-4
title: "Prune the oldest historical environment to control storage footprint"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: ["2026-03-17", "2026-05-18"]
keywords: ["parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/therettom-self-hosted-guide-opnsense-3bc340bc.md
source_anchor: ""
source_lines: [285, 400]
sha256: 811efe31468262e5cb3416449b330ecb579ca1187f1922ea7eb7f668da535e0c
---

# Prune the oldest historical environment to control storage footprint

- Go to Services: Unbound DNS: DNS over TLS . This is where the magic happens to prevent your ISP from knowing what domains you are requesting. There are a few options, but I'll show the configuration for Quad9 and Cloudflare. I recommend one (or two of Quad9's) or the other. Click the orange+ button to create entries:
| Option | Value | 
|---|---|
| Enabled | ☑️ | 
| Domain |  | 
| Server IP | 9.9.9.9 | 
| Server Port | 853 | 
| Forward first | 🔳 | 
| Verify CN | dns.quad9.net | 
| Description |  | 
| Option | Value | 
|---|---|
| Enabled | ☑️ | 
| Domain |  | 
| Server IP | 149.112.112.112 | 
| Server Port | 853 | 
| Forward first | 🔳 | 
| Verify CN | dns.quad9.net | 
| Description |  | 
| Option | Value | 
|---|---|
| Enabled | ☑️ | 
| Domain |  | 
| Server IP | 1.1.1.1 | 
| Server Port | 853 | 
| Forward first | 🔳 | 
| Verify CN | cloudflare-dns.com | 
| Description |  | 
Make sure you pick either Quad9 or Cloudflare, not both. Unless you want to. I don't care.
Assuming you did all of this isolated from the existing main network, or on its on LAN, you should be ready AFTER configuring AdGuard Home. If you don't want to use AdGuard Home, then go to Services: Unbound DNS: General and change the Listen Port to 53. You should be good to go.
🔰 AdGuard Home: DNS blocker that intercepts tracking and advertisements at the gateway level before they ever reach your devices.
🌐 CrowdSec: Collaborative, behavior-based security engine that analyzes system logs to detect malicious activity and automatically blocks aggressive IP addresses using a globally shared threat intelligence network.
🥷 Zenarmor: Lightweight, next-generation firewall plugin that provides enterprise-grade deep packet inspection (DPI), application filtering, and advanced web security controls.
When updating an OPNsense stack running complex plugins always create a ZFS safety checkpoint.
Do NOT use the GUI Check for updates button. The GUI triggers a global, concurrent update that can cause dependency, API, and database migration loops.
Because OPNsense integrates bectl into its backend framework, you can manage snapshots via the CLI or right from the Web GUI dashboard under System: Snapshots.
From GUI
- 
Go to System: Snapshots and click the orange+ button to create a ZFS backup (snapshot).
- 
The name defaults to a time format. You can change that or add a note. The name can be Pre_Change_Backup .
- 
Click Save. Easy and done.
From CLI
- Whenever you are about to modify a major system configuration, install an experimental plugin, or adjust core networking routing logic, run this terminal command:
bectl create Pre_Change_Backup
This freezes your current operating OS dataset. Because ZFS uses a redirect-on-write model, this operation is instantaneous and takes up exactly 0 bytes of space initially. It only consumes storage as files are modified or updated over time.
- To see your existing recovery points:
bectl list
The output would look something like this:
BE                Active Mountpoint Space Created
Pre_Change_Backup -      -          8K    2026-05-18 17:18
default           NR     /          1.85G 2026-03-17 14:54
N means active Now.
R means active upon Reboot.
- For a security perimeter firewall, a weekly rotation is ideal. You do not want a daily cron snapshot script running indefinitely, as a massive influx of automated snapshots can clutter your ZFS storage pool metadata over time.
You can use the built-in OPNsense cron manager to automate a rolling backup script. Skip if you don't care.
Creating automated ZFS backups.
Here's a quick cheat sheet for VIM, which I don't like (yet?).
| What you want to delete | The Keystroke | 
|---|---|
| A single character | x | 
| A whole word | dw | 
| An entire line | dd | 
| Everything from cursor to end of line | d$ | 
| Delete everything inside the file | ggdG | 
Change something quickly: If you want to delete a word and immediately start typing a replacement, use cw (Change Word). It deletes the word and drops you straight into standard typing mode.
If you accidentally butchered the file while fighting the keys, Hit Esc a few times. Type :q! and hit Enter. This forces Vim to quit immediately without saving any of the changes you just made, letting you open the file fresh. If you did make edits you want to keep, press Esc, type :wq and hit Enter to save and crash out.
- First, we need to save the rolling backup logic into a script file on the firewall. In the root shell, create the script file:
vi /usr/local/bin/zfs_rollup.sh
Paste the following rotation script into the file:
#!/bin/sh
# Prune the oldest historical environment to control storage footprint
bectl destroy -F Weekly_Two_Weeks_Ago >/dev/null 2>&1
# Cycle the existing backup backward in rotation
bectl rename Weekly_Last_Week Weekly_Two_Weeks_Ago >/dev/null 2>&1
# Freeze a fresh, current checkpoint of the operational engine
bectl create Weekly_Last_Week
What exactly is this script?
The script (obviously) creates a ZFS backup every week and handles its own cleanup completely through a rolling rotation. It will never run away with your disk space or pile up an infinite number of old snapshots.
Here is exactly how the script controls its footprint every time it runs:
[Execution Step] [Storage State]
- bectl destroy Weekly_Two_Weeks_Ago --> Frees up space from the oldest backup
- bectl rename Weekly_Last_Week ... --> Rolls last week's backup into the old slot
- bectl create Weekly_Last_Week --> Freezes a fresh snapshot of your system
Because of that order of operations, the script maintains a strict maximum of two historical snapshots at any given time (Weekly_Last_Week and Weekly_Two_Weeks_Ago).
The first two commands end with >/dev/null 2>&1. This is a silent safety valve for cron.
The very first time this cron job runs, Weekly_Two_Weeks_Ago and Weekly_Last_Week won't exist yet. Normally, trying to delete or rename a file that doesn't exist causes the system to scream an error message. That syntax silences the empty errors so the script can smoothly cruise right past them and create the very first fresh backup. By the third week, the rotation engine is fully primed, and it will cleanly drop the oldest snapshot before taking a new one—keeping the ZFS pool tidy without you ever needing to intervene.
Save and exit by pressing Shift + Colon, then wq (:wq), then make the script executable so the system is allowed to run it:
chmod +x /usr/local/bin/zfs_rollup.sh
- To make this script visible to the GUI interface, create a new template action file:
vi /usr/local/opnsense/service/conf/actions.d/actions_zfsbackup.conf
Paste this configuration inside. This maps a GUI name and description to your backend script:
[create]
command:/usr/local/bin/zfs_rollup.sh
parameters:
type:script
message:Creating Weekly ZFS Boot Environment Backup
description:ZFS Weekly Rolling Backup
Save and exit (:wq).
- Tell the OPNsense configuration daemon to reload its actions array and read your new file:
service configd restart
- Now that configd knows about the command, it will appear directly in your browser.
1: Log into your OPNsense Web GUI. Navigate to System ➔ Settings ➔ Cron.
2: Click the + (Add) button in the bottom right corner of the table.
3: Set the Description to something recognizable like: Weekly Rolling ZFS Snapshot.
4: Click the Command dropdown menu and select ZFS Weekly Rolling Backup.
5: Configure the schedule for a quiet time (for example, every Sunday morning at 2:00 AM):
* Minutes: 0
* Hours: `2`
* Days: `*`
* Months: `*`
* Days of week: `0` (`0` is Sunday in cron format)
6: Click Save.
- The automated backup task is now fully integrated into the native OPNsense ecosystem. On Sunday at 2:00 AM, the system will run the script, create Weekly_Last_Week , and you can monitor its execution or spacing anytime by runningbectl list from the terminal.
- Update OPNsense:
