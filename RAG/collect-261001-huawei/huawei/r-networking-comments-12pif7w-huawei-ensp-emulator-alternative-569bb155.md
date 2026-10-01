---
id: collect-261001-huawei/huawei/r-networking-comments-12pif7w-huawei-ensp-emulator-alternative-569bb155
title: "r-networking-comments-12pif7w-huawei-ensp-emulator-alternative-569bb155"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/r-networking-comments-12pif7w-huawei-ensp-emulator-alternative-569bb155.md
source_anchor: ""
source_lines: [1, 20]
sha256: 10463d1910a9e9c3f20caf7bb300a70c69e3904ce784e21ae220e5b7f6972fbc
---

# r-networking-comments-12pif7w-huawei-ensp-emulator-alternative-569bb155

Huawei ENSP emulator alternative 
        
        
        
    
    
    Anyone has success using the ensp emulator for huawei? I saw that support for it ended a while ago but i needed something to do emulation. I found that it's buggy sometimes. E.g sometimes if i have subinterfaces configured i get no arp entries or unable to ping across the network.I have to reboot the nodes and they might work. If i use the bare interfaces they work fine. I tried the same images in eve but it seems to be the same. Anyone with a working alternative or any ideas?
Section des commentaires
You can get eNSP working on Windows 10, but it is janky.
https://youtu.be/7kpARd4NNl4
If you're not getting arp requests and responses on a router's sub-interfaces, you need to run the command "arp broadcast enable" on them.
Thank you for this helpful comment!
Yes. I used ensp on Windows 7 and works fine. 3 years ago stopped working on Windows 10/11 .
Don't take most up to date virtualbox. 5 something if I remember.
I heard that Huawei works with something cloud to practice but didn't find additional info about that.
Commentaire supprimé par un membre de l’équipe de modération
Thanks for your interest in posting to this subreddit. To combat spam, new accounts can't post or comment within 24 hours of account creation.
Please DO NOT message the mods requesting your post be approved.
You are welcome to resubmit your thread or comment in ~24 hrs or so.
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
