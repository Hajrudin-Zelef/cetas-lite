---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-18hls67-multisite-options-091c76a8
title: "r-ubiquiti-comments-18hls67-multisite-options-091c76a8"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-18hls67-multisite-options-091c76a8.md
source_anchor: ""
source_lines: [1, 22]
sha256: 8486e87e435e9eaa036db8d25ed0e672f92fdad6f90135e6c15b50b62c78961c
---

# r-ubiquiti-comments-18hls67-multisite-options-091c76a8

Multi-site options 
        
        
        
    
    
    Howdy. I am looking for some up-to-date experience/recommendations please.
We have multiple sites globally, and we are gradually converging our network, access and video onto Unifi where possible. We have a UCK G2 Plus at HQ in the UK. This runs our HQ Wi-Fi, and our HQ door access solution from Unifi.
We also recently enabled multi-site mode and added our USA Unifi APs to the HQ controller via a S2S VPN link and set-inform.
I am now looking at our new middle east site, where I will be deploying APs, Unifi door access and unifi CCTV. I have gone into our HQ controller, and while there's a multi-site system in the Network app, it does not appear to exist in the Access app or the Video/Protect app. Am I doing something stupid?
Is the recommended approach really to deploy a separate physical controller into every "layer two site"? With separate user databases to maintain/sync, separate swipe cards, etc.? If so, then does our chairman etc. needs to carry a card from every site when he travels?
If I have to deploy a new controller, I will need to deploy an NVR anyway for the CCTV - is there a controller and NVR combined that would make more sense?
Looking for experience/recommendations/options please, very much appreciated.
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
It would be way better if unifi let you adopt the cloud key protects as a child device for other sites as protect hubs. That way you could have one master cloud key as a hub and then the others as spokes on the hierarchy.
Personally, I plan to deploy a UCK-G2+ at each physical location for talk, connect, and access. Then a UNVR, UNVR-Pro or stack of either for video. All networking is either a UXG-Pro or UXG-Lite managed by the unifi cloud console. Magic VPN will handle any site to site needs.
This assumes you are NOT looking into UniFi Identity Enterprise, which could solve your multi site needs, especially for door access
