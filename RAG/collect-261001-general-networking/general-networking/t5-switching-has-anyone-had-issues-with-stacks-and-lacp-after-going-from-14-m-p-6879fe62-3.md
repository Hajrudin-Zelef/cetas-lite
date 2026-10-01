---
id: collect-261001-general-networking/general-networking/t5-switching-has-anyone-had-issues-with-stacks-and-lacp-after-going-from-14-m-p-6879fe62-3
title: "t5-switching-has-anyone-had-issues-with-stacks-and-lacp-after-going-from-14-m-p--6879fe62"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-has-anyone-had-issues-with-stacks-and-lacp-after-going-from-14-m-p--6879fe62.md
source_anchor: ""
source_lines: [145, 212]
sha256: e68e32cf0d802115317c50ae65550c0c705a15ab92b8f3a5f63c2a2ad8865b18
---

# t5-switching-has-anyone-had-issues-with-stacks-and-lacp-after-going-from-14-m-p--6879fe62

I ended up nailing down a computer that couldn't ping our Cohesity. I got a persistent ping going then I disabled the cross/stack LACP by disabling one of the ports in the 2 port aggregate link. It immediately fixed that computer and a couple others that were having trouble in the same stack.
Every one of our buildings has at least one LACP link to an IDF over SMF which spans 2 switches at the core, and 2 switches in most IDF's. Right now to keep the lights on I'm working on disabling LACP everywhere.
I'd love to hear from people if LACP into a single switch is working?  Or is LACP bad in this version?  
Thank you,
Yellow_Tang
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-12-2023 11:08 PM
I'm running LACP between access switches and cores in different configurations, cross stack on both sides, sometimes one single access switch with 2 LACP uplinks to 2 cores and sometimes a single switch with one uplink going to one core and a backup link (blocked by RSTP) going via an other switch to an other core.
After things settled after the upgrade I didn't notice any LACP issues.
We also experienced things not being reachable in a random way and it was either the LACP links misbehaving or L3 on the cores, but in the end disabling all redundant ports (not unconfiguring LACP) helped and that is what I would try in your situation.
Since we have all links from access switches spanned across 2 cores I went into one of them and disabled all ports (not the uplink and stacking ofc), after waiting a minute or two things went back to normal, then I started enabling ports and the issues never came back.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-30-2024 08:46 AM
Could you provide some more detail on "disable" all ports. Is this on MS425? Did you disable / re-enable both unused ports and also active LCAP ports? We have major connectivity issues between MS390 v C16.7 up link to core MS425 v16.7.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-31-2024 12:29 PM
It's on the MS425 (my core), I don't have MS390's, I have MS225 downstream.
What I mean is that I would disable all redundant links towards given downstream switches from the MS425 perspective.
If I have a LACP connection to some downstream switch (or downstream stack) consisting of 2 ports, I would manually disable one of them and leave only one running (if I would have 5, I would disable 4 and leave one running).
So basically, what I'm doing is killing all redundancy in my network during upgrades, just in case.
And since I'm doing this only on the core switches, if something is going wrong I can easily turn the ports back on (not disabling these links on both sides).
In a perfect world, lets say you have 2 core switches, both of them having connections to all downstream switches (LACP, 2 ports, each ending up in one of the cores), then it would mean that you just need to disable all ports on one of the core switches (assuming you have them stacked, the stack you leave enabled) - that's what I'm doing more/less.
If you read the release notes for the firmwares you will see that there is a lot going on around switches creating loops during reboots on LACP ports and so on.
Its my way of preventing these since some firmware versions ago.
BTW, maybe this known issue is hitting you?
- Switches move LACP ports to an active forwarding state if configured. This can cause loops when connecting to an MS390 or other Catalyst switches unless the bundles are configured on the MS390/Catalyst switches first. All non-catalyst ports are configured in passive LACP mode so that loops do not occur between Meraki switches (always present)
Sorry for the lengthy post and repeating myself, just ask if you have any doubts
And by any means - that is not what Meraki would suggest, its only my way of working around the issues that I kinda got used to during the years.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-01-2024 12:51 PM
I didn't catch that you were doing this ONLY during upgrades. I disabled one downstream port on every switch so we no longer have redundancy at all and it fixed the issue. I cannot reenable LACP though, or the problem will come back. I'm waiting for an upgrade past 15.21.1 to be stable in hopes that they address cross stack LACP. Of interesting note, I think if one side is Cross Stack, and the other goes to a single switch it's still okay. It's only if the redundant links are split across stack members on both sides.
Thank you!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-02-2024 01:14 AM
Yes, when updating or I think when rebooting the devices to be more specific.
But on a daily basis I'm not having any issues, so what you are experiencing is very weird.
Have you contacted support?
Can you share all settings from the ports on both sides here? Maybe there is something that you are just missing?
