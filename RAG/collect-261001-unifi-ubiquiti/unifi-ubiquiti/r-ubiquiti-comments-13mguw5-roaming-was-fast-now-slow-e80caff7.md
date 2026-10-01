---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-13mguw5-roaming-was-fast-now-slow-e80caff7
title: "r-ubiquiti-comments-13mguw5-roaming-was-fast-now-slow-e80caff7"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-13mguw5-roaming-was-fast-now-slow-e80caff7.md
source_anchor: ""
source_lines: [1, 19]
sha256: 6c206ff03f9b06f901b0804f9f2fd661ff9780642dd2fa0dc6aff9e45b00afb4
---

# r-ubiquiti-comments-13mguw5-roaming-was-fast-now-slow-e80caff7

Roaming was fast, now slow 
        
        
        
    
    
    I have 2 x AC Lites, 1 was connected via mesh and therefore they were on the same channel. I've since moved the meshed one a few metres into another room but hardwired. As I did this I changed the channels (to both DFS, but different) and now roaming seems super slow...like the quality of the connection can get almost non existent before they change to the next AP. I've since changed the transmission power on both to medium (from high) but that hasn't changed anything. As I'm writing this, I'm wondering if it's to do with the channels, before they were both on CH 60... Any suggestions?
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Roaming is handled by the client so you need to look at the devices specs and requirements for WiFi and adjust your settings to help it but devices decide on what ap to talk to.
But roaming requires a controller, so that's not completely true?
I don't worry about transmission power. If you need more coverage, leave it high. You need BSS Transition and Fast Roaming on. They work together to roam the clients. I've always had issues with those off.
This. You can also adjust Minimum RSSI at the APs settings so the client will be "rejected" if the signal is too weak.
Yeah, I try to stay away from Minimum RSSI because I don't want to kick devices off on the edges of coverage where there's no AP to jump to. I have many APs on high power close together with no issues with any devices roaming and no Minimum RSSI set.
Try putting them on different channels Turn off band steering Disable min rssi Enable fast roaming and bss Do an RF scan you probably don't need them on a dfs channel. 20MHz band width for 5ghz is typically more reliable but a little bit slower. But still fast enough for 99% of things
