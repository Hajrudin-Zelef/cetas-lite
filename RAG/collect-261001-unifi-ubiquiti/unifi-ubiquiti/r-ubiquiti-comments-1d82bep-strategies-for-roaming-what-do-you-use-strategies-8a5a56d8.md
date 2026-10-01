---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-1d82bep-strategies-for-roaming-what-do-you-use-strategies-8a5a56d8
title: "r-ubiquiti-comments-1d82bep-strategies-for-roaming-what-do-you-use-strategies-8a5a56d8"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-1d82bep-strategies-for-roaming-what-do-you-use-strategies-8a5a56d8.md
source_anchor: ""
source_lines: [1, 30]
sha256: 64e8deac4cbaa9693d2331115cacdc436f22d0e7a566598f652916863fbcb971
---

# r-ubiquiti-comments-1d82bep-strategies-for-roaming-what-do-you-use-strategies-8a5a56d8

Strategies for roaming - what do you use? Strategies for roaming - what do you use?  
        
        
        
    
    
    Not the worlds biggest problem but it's my big problem as the family are starting to revolt....
So I live rural and have a large property (grounds and house) and it's pretty spread out. For example, the "office" is an outbuilding around 90ish meters from the house.
As an example of my problem - Within the office I have one AP (nanoHD), in the courtyard I have a UAP AC M and in the entrance hall of the house I have another nanoHD.
If I walk from the office to the house for my 100th coffee, my phone will cling on to the office AP. Nothing works but it will hold on to it. If I turn wifi off and on, it catches the one in the house and all is well.
Moving around in the house and garden has the same issue.
Is minimum RSSI the correct tool here? Should I just keep increasing it until life is better?
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Sounds like a great application for minimum rssi. Also fast roaming and bss transition
Except for devices that don't play nicely with fast roaming. Bur min rssi - yes
Thanks. Will have a play. The only possible non-compliant devices I might have don't move around. It's really just SWAMBO and the kids I need to keep happy...
Fast roaming is a standard from 2013 or so... Nearly every device since 2014 supports it
Minimum RSSI is a terrible strategy here. That will boot clients before they can make a roaming decision.
And unless you’re doing enterprise auth, .11r isn’t going to help you.
Whoh this is news to me: 11r doesn't work without enterprise auth?
I’ve turned on bss transition and fast roaming, set the APs to different channels and tuned the power levels such that they’re not set to max dBM, but rather to a level that encourages roaming.
Roaming is a client decision and the client can decide not to roam even if they AP sends a packet to begin the transition. I’ve seen in the logs of some routers that the client rejected the BSS roaming request.
There is only one “strategy” for roaming in WiFi: the client decides when and where.
Commentaire supprimé par le membre
Not sure of which version I currently have but all devices are up to date. Last update I pushed through last week and, oddly, a couple today.
