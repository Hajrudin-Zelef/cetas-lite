---
id: collect-261001-general-networking/general-networking/t5-switching-has-anyone-had-issues-with-stacks-and-lacp-after-going-from-14-m-p-6879fe62-2
title: "t5-switching-has-anyone-had-issues-with-stacks-and-lacp-after-going-from-14-m-p--6879fe62"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-has-anyone-had-issues-with-stacks-and-lacp-after-going-from-14-m-p--6879fe62.md
source_anchor: ""
source_lines: [35, 144]
sha256: b36c511e72dedeef4eff26f8356dc5b71239ce3fbd18b13fb4bf9323d78c2d0e
---

# t5-switching-has-anyone-had-issues-with-stacks-and-lacp-after-going-from-14-m-p--6879fe62

			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-30-2023 04:23 PM
I've not upgraded too many switch stacks from 14.x to 15.x but the ones I have done have gone smoothly.
Without digging too deeply into your issue, it's feasible that you hit one or multiple of the known issues around LACP and stacks on upgrade/reboot. To name a few from the release notes:
- Cross-stack LACP bundles experiencing a switch reboot will cause the remaining online port to experience an outage for up to 30 seconds. The same is seen again when the switch comes back online (present since MS 10)
- Loops can be seen when rebooting a stack member containing a cross-stack lag port (always present)
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-31-2023 10:31 AM
Thanks for the reply Brash!
I had considered those known issues, but we were on 14 code before and didn't have this, I would have expected it there as well since one was always present and the other was from code 10. The other thing that is puzzling is that these errors aren't coming up when something is changed or rebooted. The switch is running just fine for a few days then I get a short burst of these errors from across campus, the errors come up and clear in a short time, but it seems to come in waves.
Currently I haven't had a UDLD error since May 24th, but I had it on 6 completely different switch stacks, it occurred and cleared on all of them all within 10 minutes. Then on May 22nd I had another wave of these errors, all happening and clearing within about 25 minutes. and this was in 14 completely independent stacks and buildings.
I'm not sure, I'm hoping it stops happening though. A big part of me was starting to wonder if one of my core switches didn't upgrade correctly.
Yellow Tang
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-01-2023 02:46 AM
I have a somewhat big setup and the upgrade went terribly wrong, nothing I'm not used when it comes to Meraki switches (sad, but true) - this will be a long story
I have 4x MS425-32 switches in a stack + a number of MS225-24/48 in different combinations (some are single, some are 2pcs stacks and some are 3pcs or more which is very important in the Meraki world) - overall ~40 Meraki switches + ~10 WS2960X.
All "remote" switches are connected to 2 of the core switches using LACP, so no SPOF.
I set up a staged upgrade where I first wanted to upgrade the core switches (they are handling all L3 as in your setup), then 2 pure management switches in server rooms and then the rest of the remote switches.
Here is what happened:
- cores upgraded nice, everything green, all ports up, BUT
- as in your story I had something between 10-20 switches behind the cores offline and not handling traffic despite all ports being up and showing no issues
- so the upgrade got stuck and I needed to react
- first I rebooted all cores from the dashboard - didn't change a thing
- then, since I have a long story with the Meraki switches I started disabling all redundant connections towards the remote switches (so basically all ports on core 02 and 04 disabled, only stacking was up)
- after that almost all switches started regaining connectivity and started upgrading
- by saying almost I mean that there is a story if you have more than 2 switches in a stack and a cross stack LACP connection to the cores they tend to "break", some of the switches in a stack will be offline, some online, the best resolution to that is remove all redundant connections (I disabled that from core side) and then reboot one of the switches that is online, if this doesn't help you pick the next one and next one and this will eventually kick in. If you are onsite, then just remove redundant connections and reboot the whole stack. I assume that rebooting the master switch would help right away, but I didn't know that the dashboard shows the info, at least on the latest firmware
- I waited some time and observed the current firmware version on the dashboard (it will show you if its "not current" or MS 15.21.1)
- after everything upgraded I rebooted the cores once again and then I started enabling the redundant ports and everything is working as expected since then (two weeks ago)
So this was a journey which took me over 2 hours to fix remotely, I even didn't bother to raise a case because I knew that playing with ports and rebooting will work and also I planned to do a full reboot anyway on the cores because I have low trust in the firmware upgrade process (bad experience).
If you have any questions - go ahead
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-03-2023 01:03 AM
@jacek.jagielka what release were you coming from? We did have a similar issue with a much earlier version of 15, but not generally from one 15 version to another. The only thing we do differently is we upgrade the L2 edge first, wat a bit and then the L3 core last.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-04-2023 11:59 PM
We were upgrading from MS14.33.1.
Is this something we should be doing? First upgrade all L2 remote switches and then the cores?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-03-2023 12:18 PM
Yes, we've seen similar issues going from 14.33.1 to 15.21.1 with out MS250-48FP switches. I've done troubleshooting with support, full stack reboot, etc., I'm at the point where I'm looking to revert to 14.33.1
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-05-2023 12:30 AM
Are you still experiencing issues? What is happening?
Curious why you are thinking about an rollback.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-04-2023 03:31 PM
Thank you everyone for posting your experiences here! Unfortunately I will miss the 14 day rollback window due to travel, but if you go forward with that South Paw, I'd love to hear how it goes. Meraki support said there is absolutely no way to roll back after 14 days.
I have not removed LACP, and I've noticed that I'm getting almost no UDLD errors now (none since the 24th actually), plus the switches have been fairly stable (Knock on Wood).  I have my fingers crossed that this all works itself out.  I've never heard of a firmware upgrade needing to break in,.. but maybe something still wasn't finished?  I can only guess and hope lol.
Thanks again!  And if anyone else has experience with this please let me know,
Yellow Tang
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-12-2023 08:44 AM
So,.. this got allot worse. I started noticing that some internal IP addresses would randomly not be reachable from certain machines on the network, but other machines connected to the same switch were just fine. I could even unplug a machine from a port, plug a different computer in, and it would start working. It was very intermittent, and the same machine that wasn't working yesterday would suddenly be working today, but another one wouldn't be. It's worth noting that the Cohesity VIP might not be pingable, but half the node addresses on the same subnet would respond, and half wouldn't. Which ones did and didn't was not consistent across different workstations either though. The only thing that was consistent was that all virtual computers always had a perfect connection to everything. I suspect that is because LACP is not setup between that switch and the core.
As you can imagine this was causing some bizarre and difficult helpdesk tickets with glitchy behavior like this, and those poor guys have enough to do.
