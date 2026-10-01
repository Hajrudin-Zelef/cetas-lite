---
id: collect-261001-general-networking/general-networking/manual-settingsmenu-html-35cf8d4f-3
title: "manual-settingsmenu-html-35cf8d4f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "memory"]
source: docs/RAG/collect-261001-general-networking/manual-settingsmenu-html-35cf8d4f.md
source_anchor: ""
source_lines: [150, 195]
sha256: bbbaf772415149580a95e7150c3c19091aa1906db9dada244e3745251883b85c
---

# manual-settingsmenu-html-35cf8d4f

| Periodic NetFlow Backup | Periodically backup Netflow state. | 
| Periodic Captive Portal Backup | Periodically backup Captive Portal state. | 
| Power Savings |  | 
| Use PowerD | PowerD allows tweaking power conservation features. The modes are maximum (high performance), minimum (maximum power saving), adaptive (balanced), hiadaptive (balanced, but with higher performance). | 
| On AC Power Mode | Set power mode when on AC (on grid). Default option is: hiadaptive. | 
| On Battery Power Mode | Set power mode when on battery. Default option is: hiadaptive. | 
| On Normal Power Mode | Set power mode the power utility can not determine the power state. Default option is: hiadaptive. | 
| Disk / Memory Settings |  | 
| Swap file | Create a 2 GB swap file. This can increase performance, at the cost of increased wear on storage, especially flash. | 
| /var RAM disk | This can be useful to avoid wearing out flash storage. Everything in /var, including logs will be lost upon reboot. | 
| /tmp RAM disk | See above. | 
| System Sounds |  | 
| Disable the startup/shutdown beep | Disable beeps via the built-in speaker (“PC Speaker”) | 
Logging
Local log settings can be found at , tab “Local”.
The regular log files will use the following standard pattern on disk /var/log/<application>/<application>_[YYYYMMDD].log (one file per day).
Our user interface provides an integrated view stitching all collected files together.  Available settings may change the appearance on disk depending
on space and time constraints for log rotation.
Many plugins have their own logs. In the UI, they are grouped with the settings of that plugin. They mostly log to /var/log/ in text format, so you can view or follow them with tail.
An overview of the local settings:
| Option | Description | 
|---|---|
| Enable local logging | Disable to avoid wearing out flash memory when applicable and set up remote logging instead. | 
| Maximum preserved files | Configures the number of days to keep logs or the number of files if “maximum file size” option is used. | 
| Maximum file size | Limit the file size of the logs instead of keeping one log per day. | 
Tip
When using (very) small file size limits, it is possible to schedule the rotate action more frequently using cron
(). Seek for an action named Rotate log files in the list in that case.
Remote log settings can be found at , tab “Remote”.
Add a new Destination to set up a remote target destination.
| Option | Description | 
|---|---|
| Enabled | Master on/off switch. | 
| Transport | Protocol to use for syslog. | 
| Applications | Select a list of applications to send to remote syslog. Leave empty for all. | 
| Levels | Choose which levels to include, omit to select all. | 
| Facilities | Choose which facilities to include, omit to select all. | 
| Hostname | Hostname or IP address where to send logs to. | 
| Port | Port to use, usually 514. | 
| Certificate | Client certificate to use (when selecting a tls transport type) | 
| Description | Set a description for you own use. | 
Note
When using syslog over TLS, make sure both ends are configured properly (certificates and hostnames), certificate errors are quite common in these type of setups. On OPNsense the general system log usually contains more details. When it comes to tracking syslog-ng messages, this is usually a good resource.
A reconfigure doesn’t always apply the new tls settings instantly, if that’s not the case best stop and start syslog in OPNsense (using the gui).
To activate any changed settings use the “Apply” button below.
To clear all the logs on the system use the “Reset Log Files” button.
