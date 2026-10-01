---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-13f8d5j-adoption-failed-dropped-adoption-everytime-1130cc4d-2
title: "r-ubiquiti-comments-13f8d5j-adoption-failed-dropped-adoption-everytime-1130cc4d"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-13f8d5j-adoption-failed-dropped-adoption-everytime-1130cc4d.md
source_anchor: ""
source_lines: [46, 85]
sha256: f2a4ee6384e62351226a5f00c3419109de6eb58a0ad85f397edec3a7505ea706
---

# r-ubiquiti-comments-13f8d5j-adoption-failed-dropped-adoption-everytime-1130cc4d

Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Have you tried resetting ? Is it wired? Disable auto meshing
Did all of that, i figure since I’m running wire i will disable auto meshing
And yes in fact i reset one of them for like 2-4 times lol So frustrating
I had similar issues on 3 of my ap but I can’t recall what fixed the issue. But I did power cycle the ports on use pro. Restated both udm and usw reset the AP put new cables and put the AP on a newer ports but I do remember resetting multiple times standing on the ladder lol
Same thing here. Tried to adopt a new U6 Mesh and it seemed to crash the AP upstream.
When I say crashed, I mean the device upstream immediately errored with failed adoption even though the devices have been running fine for months.
Moved it to a different location that was wired to the USW Lite switch and it crashed the switch.
POS product
added an update to my initial post
Here’s the picture of adoption failed
https://imgur.com/a/OzbLUFh
Ubiquiti software is among the worst I've ever encountered in my whole 20ish years in IT and networks. Absolute fkn trash.
Wasted DAYS worth of time in the past with this horseshit adoption process. Randomly it decides to work and like a house of cards you try desperately to never touch it or even look at it again for fear it might break.
Well.... 6 months later.
I literally just updated one of my APs using the "update" button in unifi. Sure enough it auto rebooted the AP and now adoption has failed and no matter wtf I do I cant get it back online so half my network is out cuz its the AP for that side.
Will NEVER EVER EVER EVER EVER EVER even look at an ubiquiti product again, let alone buy one. You couldnt fkn pay ME to use this shit.
thumb fade public light dinosaurs soft spark price threatening cautious
This post was mass deleted and anonymized with Redact
Yah, outside of the APs I currently have I won't be adding anymore of theirs.
I've had some success sshing into the device and sending the callback command to the Unifi controller ip and port.
Generally have to do that multiple times too tho, just cuz.
I had a similar problem recently. The fix for me was to completely remove the AP from the controller so there was no entry for it at all. Then, I reset the AP and logged in via SSH and used the set-inform command to manually set the IP of the controller. From then on it worked just fine.
Where can i get the id and password dor the aps?
If it’s reset then it should just be Ubnt/ubnt
I tried this earlier, work for some time.. like 2-3 seconds, and then failed to work again
It goes to its fallback ip, 192.168.2.191 I want to switch my default so i can discover the ubnt, too bad that ip range is used for teleport
Hi u/jhanbali We apologize for the frustrations. Please share more info and any related support tickets here so we can properly escalate and assist: http://community.ui.com/social-feedback
I havent get a support ticket, where do i submit the question again?
updated the problem information with more diagnostic that I just did.
I am having the same issue! Except that my AP IP address kept falling back to 192.168.10.xxx. Extremely frustration.
Hi, Dries from UniHosted, please make sure that the inform ip is set correctly in your system.properties. Also if you can, try to access the db and make sure that the failing devices are completely removed from the "devices" collection.
Commentaire supprimé par le membre
Can you elaborate?
refuses to fucking elaborate
