---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-125hgcw-udm-stuck-on-updating-unifi-os-console-updating-a4ef09a2
title: "r-ubiquiti-comments-125hgcw-udm-stuck-on-updating-unifi-os-console-updating-a4ef09a2"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-125hgcw-udm-stuck-on-updating-unifi-os-console-updating-a4ef09a2.md
source_anchor: ""
source_lines: [1, 43]
sha256: 990da4e0e8836077f00228d2e8893bc42a7c3799ae3b412a188862f8941a7e8e
---

# r-ubiquiti-comments-125hgcw-udm-stuck-on-updating-unifi-os-console-updating-a4ef09a2

UDM stuck on updating UniFi OS - Console Updating 
        
        
        
    
    
    Was checking my UniFi stack this morning and noticed I had an update for my UniFi OS on my UDM [v1.xx to v2.xx].
Manually did a config backup and started the procedure, getting a notification that it could take up to 20 minutes.
Now almost 1½ hours later, the tab still shows `Console Updating`, if I open a new tab and go to the local IP I get a `Connection Refused` message and if I try the UniFi SSO login I get a `Console Offline` message.
SSH is disabled because I normally only turn it on when required, so I can't access the machine anymore. Internet still works, so do all the connected UniFi devices, but I'm simply unable to access my UDM.
What to do?
**Update**
      Spoke too soon about the 'all connected UniFi devices still work` bit...
It seems wifi devices can no longer connect 🥲
    
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
After 2½ hours I thought it was ok to reboot the UDM.
Luckily this didn't brick the machine.
After this, it turns out the update had not gone through and it was still on v1.28.38 and an update to v2.4.27 was still available.
With fingers crossed I had tried to update again, this time the update went through in a matter of minutes.
After which the system reported an update to v2.5.17 was also available for me, which again only took a few minutes.
All seems to be well again, UniFi access locally and through SSO working again and everything is up to date.
One thing that struck me as odd:
I have Automatically update UniFi OS turned on, but apparently it hadn't updated at all for quite some time [2 months based on release notes].
I think most on here would highly suggest you NOT set updates to auto. That is a problem waiting to occur. Why do you want that? It's better if you not be the first one out of the gate on updates, and you should schedule them for when your have the time to work out any issues (even if it's just a home situation). That fact it had not been doing the auto updates "for quite some time" should prove it's not a hands free procedure. Tempting fate with auto seems needless. $.02
I have found a different method that seems to be a lot more successful for dealing with this but requires physical access.
From the UDM Pro's LED Menu, go to About. Then, Update from there. I have found that this causes the update to instantly start and truly finished within 20 minutes (~18 minutes for me in this case).
Prior to this, I tried updating twice from the web GUI via local intranet and failed both times.
* First Time: Attempted via Web GUI. Waited 3 hours. Then, I forced a restart from the LED Menu.
* Second Time: Attempted via Web GUI. Waited overnight (roughly 12 hours). Then, I forced a restart from the LED Menu.
Both previous times, the system recovered but failed to update.
* Third Time: Attempted via LED Menu. Instantly started (could see UDM Pro's LED Menu change to Updating...). 5-10 minutes in, it changed from "Updating" to "Critical Updates" mid-process and cautioned to not shut off the system but to wait. And, ultimately, after a total of ~18 minutes, it was updated.
I am currently 30 mins into an estimated ~5 min update from 2.x to 3.x on UDM Pro. I’m going to give it a few more hours…
First time I’ve ever run into any kind of failure like this with UniFi
Yeah, Prior to attempted auto update to 2.x.x, auto update was on but it wasn't working. I only found out as I do a login and check every so often and noticed it wasn't at the latest. Now mine is stuck at updating which I manually started, not sure how I am going to reboot it, going to have to just unplug it, been sitting there updating for the past few hours.
Update: Had to factory reset (hold the little pin button for 10sec), it updated during the setup process, currently on 2.4.27. Had to re-configure my subnets again, but have done this so many times, didn't take long
Dang that sounds rough, I just did my UDMP OS firmware update from 1.x.x to 2.x.x this past Friday, came back this morning to find it still updating (so it was trying to update all weekend). I found that I could get it to reboot from the LED screen menus on the front without pulling power or factory reset. Just told it to update again, crossing my fingers that it goes through this time.
I just rebooted mine from the LED. Did a second update work for you?
