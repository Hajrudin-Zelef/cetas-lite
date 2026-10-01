---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-osl3ud-how-to-configure-a-trunk-port-on-a-ubiquiti-switch-06f20009
title: "r-ubiquiti-comments-osl3ud-how-to-configure-a-trunk-port-on-a-ubiquiti-switch-06f20009"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-osl3ud-how-to-configure-a-trunk-port-on-a-ubiquiti-switch-06f20009.md
source_anchor: ""
source_lines: [1, 20]
sha256: a87bd2911aa1beaf3d12601882b6e5fa46bc7ad1e848ecc0b3d2fd8ac23a2d07
---

# r-ubiquiti-comments-osl3ud-how-to-configure-a-trunk-port-on-a-ubiquiti-switch-06f20009

How to configure a trunk port on a Ubiquiti Switch 
        
        
        
    
    
    So our main networking guy is on his holidays and whilst I know my way around Ubiquiti kit quite well we have had a request to change on the port on a switch to a Trunk port and I have literally no idea how that would work.
Essentially there are 2 VLANS - VOIP and DATA
The telecoms company has a device that we can see attached to the switch and showing up on the correct VLAN, however, nothing on that VLAN can talk to it. Nor can I ping it from the switch management console.
So the telecoms company have asked me to make that port a Trunk Port.
Can anyone give me some sort of step-by-step guide or point me in the right direction?
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
If I'm not mistaken, go to the switch in the console, click on the desired port and select port profile 'All'.
That's exactly right and that's the default for all switch ports. Maybe you need to ask your telvo which VLAN specifically they expect data on.
"all" in Unifi port profile parlance is a trunk port but it means "all CONFIGURED VLANs" so anything not configured won't go through. Thus the need to ask which VLAN they need.
