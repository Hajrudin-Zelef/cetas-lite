---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-ercpzb-mikrotik-routeros-automatic-backup-and-update-14983190
title: "Mikrotik RouterOS automatic backup and update"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/r-mikrotik-comments-ercpzb-mikrotik-routeros-automatic-backup-and-update-14983190.md
source_anchor: ""
source_lines: [1, 76]
sha256: 0220550da07835f65fbce9f3ebf9d59e7e019193da856b81e24a480a0458fed2
---

# Mikrotik RouterOS automatic backup and update

A community-contributed subreddit for all things Mikrotik. General ISP and network discussion also permitted. Please ensure if you're asking a question you have checked the Wiki First: https://help.mikrotik.com

# Mikrotik RouterOS automatic backup and update

Hello guys, it's been a while since I posted here the first version of my script. I was quite surprised by how many people implemented it on an everyday basis. So now I would like to introduce the second version.

I have entirely rewritten it in order to provide more flexibility and to make it even more useful.

Here is a small list of features:

- 
    Ability to choose script operating mode according to your needs. *(Read below)*
- 
    If automatic updates are enabled, you can set script to install only patch versions of RouterOS updates. *This means if the current RouterOS version is v6.43.6, the script will automatically install v6.43.7 (new patch version) but not v6.44.0 (new minor version), for example.*
- 
    Script includes primary information about the device into the email message. So you can easily find the backup you need among multiple devices.
- 
    For safety purposes, an automatic update process will not be started if script could not send backups to email.

Script operating modes:

    **Backups only** - script creates system and config backups and sends them to specified email as an attachment. Using email account as storage for your backups.
  

    **Backups and notifications about new RouterOS release** - Except backups, script also checks for new RouterOS firmware release and provides this information in the email.
  

    **Backups and automatic RouterOS upgrade** - Script makes a backup, then checks for new RouterOS version, and if new firmware released, script will initiate upgrade process. By the end, you receive two emails. The first one contains system backups of the previous RouterOS version, the second message will be sent when the upgrade process is done.
  

You can find the script and installation instruction here: https://github.com/beeyev/Mikrotik-RouterOS-automatic-backup-and-update

In conclusion, I would like to say that I will be glad to any feedback and if someone could correct my English grammar inside the code and readme files 😳

Cool, I'll check it out. Thanks

That is a nice script! Good job!

I’ll Check this out and thanks.

Awesome, will check it out.

Have been running the old version which works very well indeed. Saved my bacon once when an upgrade trashed my config. Just found the email with the saved config and restored it. Back up and running in <5 min.

Works great. Even updates routerboard firmware as part of the update process.

Thanks man!

Will check it

Excellent script! Thanks for sharing!

This may be a stupid question, but what's the point of not updating to a new minor version as opposed to just updating to the new patch version?


When the bridge filtering change occurred last year, and broke a lot of devices, was it a new minor version or a new patch version?

Patch versions are supposed to fix the bugs and security leaks, so such updates has to be safe for automatic updates.

I don't remember actually.

So you'd recommend using the "osupdate" on production devices? Along with the long-term update channel? Have you deployed this script in production environments with no issue?

Furthermore, when routerOSv7 gets released, what would your script do?

It's completely up to end user to decide when to use this option. I use this feature at home and in several non critic routers.


I have not tested it on routerOSv7.

Seems to be backing up every day rather than just sending a backup when installing a new version. Can it be set to only email backups on new version updates?

    *>  Can it be set to only email backups on new version updates?*
  

Not yet, but it's a good idea to implement. Thanks.
