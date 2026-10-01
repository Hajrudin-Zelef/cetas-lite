---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-988168-how-to-edit-delete-ipv4-address-group-from-unifi-gateway-2664564e
title: "questions-988168-how-to-edit-delete-ipv4-address-group-from-unifi-gateway-2664564e"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-988168-how-to-edit-delete-ipv4-address-group-from-unifi-gateway-2664564e.md
source_anchor: ""
source_lines: [1, 27]
sha256: 1748ff89e18107cc8dbc090702383bd22438bdb45cf54b604e575e8815991402
---

# questions-988168-how-to-edit-delete-ipv4-address-group-from-unifi-gateway-2664564e

I'm just starting out with my unifi gateway and I was configuring a firewall rule through the UI. I messed up creating a "IPV4 Address Group". It has the wrong IP address ranges and I'm stuck. I have found no way of editing or deleting a "IPV4 Address Group" in the UI. I'm assuming it might be possible if I were to SSH into the gateway. I'm able to SSH in, but haven't found appropriate documation of how to configure the unifi gateway via command line.
3 Answers 3
Using your Unifi network controller, you can go into the Routing & Firewall of your Settings tab, select Firewall, click on the Groups tab, then just delete the appropriate entry group.
- 
        A screenshot is helpful because it is possible O.P. is using the webserver embedded in the gateway and not a controller. If so, a screenshot of the controller would open a whole new world to O.P.rjt– rjt2019-12-30 01:06:23 +00:00Commented Dec 30, 2019 at 1:06
- 
            
            
- 
        this was the correct solution for me. i completely missed the groups tab. thankspgreen2– pgreen22020-01-04 16:01:47 +00:00Commented Jan 4, 2020 at 16:01
- 
        2For anyone looking for a modern (late 2023 and later) solution Ken's answer below is the correct one for the new user interface.Kean– Kean2023-12-13 00:45:27 +00:00Commented Dec 13, 2023 at 0:45
- 
        12025 Unifi Server v9+ - Settings > Profiles > Network ObjectsJosh Hibschman– Josh Hibschman2025-07-01 18:25:50 +00:00Commented Jul 1, 2025 at 18:25
- 
        Now it's under Settings -> Overview -> Network ListsShannon– Shannon2026-01-26 04:40:23 +00:00Commented Jan 26 at 4:40
I found this thread when trying to figure out how to edit a group.
I have a UDM Pro, a USW-48-PoE, and a U6-LR.
I needed to view the USW-48-Poe in network, then go to settings. While in settings, Instead of going to Firewall & Security, I needed to go to Profiles.
This is where I had a section titled Port/IP Groups which shows the group I had once created.  I can now modify it as needed.
Hope this helps others.
- 
        Thank you, that was indeed a big help. Previously I would switch back and forth between the classic and the new UI just to edit groups. This was on a UDM-Pro running 8.0.7.Kean– Kean2023-12-13 00:44:14 +00:00Commented Dec 13, 2023 at 0:44
- 
        2As of the start of 2025, it's still in Profiles, but the tab is now named "Network Objects". That shows all of the Port and IP Groups and allows modification.Cliff– Cliff2025-02-05 22:43:32 +00:00Commented Feb 5, 2025 at 22:43
I have the same problem. I mistakenly pressed OK but I still need to add some IP ranges to the group - Trying to find out how to edit a group is a bit like playing Zelda, missions, submissions, find the sword, get the wand....
It's as though they go out of their way to make it as hard as possible to find and edit basic settings, so much so that is takes an intra-planetary search (google) to find what is right in front of me. Way to go UX designers!
