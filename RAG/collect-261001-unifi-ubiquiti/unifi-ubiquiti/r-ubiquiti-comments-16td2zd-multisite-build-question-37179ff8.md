---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-16td2zd-multisite-build-question-37179ff8
title: "r-ubiquiti-comments-16td2zd-multisite-build-question-37179ff8"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-16td2zd-multisite-build-question-37179ff8.md
source_anchor: ""
source_lines: [1, 27]
sha256: 373bcd413b31bc6d3da1eccb9db1d642bb41affb1ca9d185ff1edecc71b9fcf2
---

# r-ubiquiti-comments-16td2zd-multisite-build-question-37179ff8

MultiSite build question 
        
        
        
    
    
    Good afternoon people!
I am rather inexperienced with how Ubiquiti does things so I hopefully have a quick question here.
I have a client who wants to switch from Meraki to Ubiquiti to cut ongoing costs with licensing.
They have a total of three sites with a fourth on its way.
I would like to setup the following a Dream machine Pro since the router onsite is dying with 4 or so access points for that office.
From there I was hoping to use a cloud key for each site with two access points per site.
Lastly, I hope to be able to manage the lot with which I assume I'll need to set up an UniFi Controller or can the Dream Machine handle multi-site with two cloud keys?
Or is there a better way?
Hardware details:
Unifi wi-fi 6 lite dual-band x7-10
Dream machine pro x1
UniFi Cloud key gen2 Plus x 2
Some Poe Injectors for sites without poe. x4
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
The UDM Pro/SE runs its own controller. The CloudKey gen2+ runs its own controller. In simple terms if setup the way you describe, you can create 1 Ubiquiti SSO account and connect/control all sites with the 1 account.
I would suggest a UDM Pro or UDM SE, to eliminate the need for PoE injectors, at each site. A little more costly but more functionality and more future proof.
