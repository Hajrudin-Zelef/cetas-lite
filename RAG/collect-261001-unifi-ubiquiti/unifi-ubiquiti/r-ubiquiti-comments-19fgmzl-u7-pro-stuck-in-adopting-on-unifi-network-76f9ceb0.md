---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-19fgmzl-u7-pro-stuck-in-adopting-on-unifi-network-76f9ceb0
title: "r-ubiquiti-comments-19fgmzl-u7-pro-stuck-in-adopting-on-unifi-network-76f9ceb0"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-19fgmzl-u7-pro-stuck-in-adopting-on-unifi-network-76f9ceb0.md
source_anchor: ""
source_lines: [1, 31]
sha256: a978554144b03cb4919a61888795da68fbda387662ed381980415a0b9fc938dd
---

# r-ubiquiti-comments-19fgmzl-u7-pro-stuck-in-adopting-on-unifi-network-76f9ceb0

U7 Pro stuck in "Adopting" on Unifi Network 
        
        
        
    
    
    UPDATE: SOLVED
Hi
I've just got two U7 pros today. Hooked them up a PoE switch, they power up.
I downloaded the Unifi network controller and mongo DB and i'm running them as a docker container, the controller seems ok, I can access it on the LAN. I power up the APs and they just sit there in "adopting" status. Sometimes the controller says no APs, sometimes 1 AP, sometimes 2 AP but always "adopting". I read through the Ubiquiti documentation and it said sometimes they need to be updated first. I see no way to do this while it's in adopting state.
It's been 3 hours now and none of them will add to the controller.
Anyone got any ideas please?
Edit: Additional Info,
I have manually updated the APs to "UniFi firmware 7.0.35 for U7-Pro 23 Jan 2024 " the latest available. The problem remains.
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
If you have the IP of the AP, SSH into it and run 'set-inform http://controllerIP:8080/inform' and then try adopting again. Default login would be ubnt/ubnt.
omg thank you so much. it has now changed to "Getting Ready" and then "Up to date" and I can see channels and SSIDs and things now!!!!
how strange it did discover it but wouldn't adopt without doing that.
you saved my bacon. thanks a lot.
No problem. We have found that the auto adopt has only gotten worse over the years and pretty much just do SSH adoptions anymore.
thanks -- this comment put me on the right path, as well. U7-Pro lives in Location A, where i was running a self-hosted unifi-controller container (haven't gotten around to setting up the new version yet). recently moved my Proxmox node from Location B into Location A, replacing it. didn't remember that A's inform port was set to the default 8080, whereas B's port had to be set to 8181 to resolve a docker port conflict.
when i realized this, i still couldn't get the U7-Pro out of the "Adopting" loop. the `set-inform` command straightened me out.
Fing hero. First time got my hands on ubiquiti APs and it drove me nuts for the past 3 hours troubleshooting it! Tyvm :)
Glad you got it working, but to add some context to why it was an issue. By default the network controller try’s to set the set-inform URL to the local IP address of the network controller. When running on bare os (not docker), this is not an issue as that’s the actual IP of the device hosting the controller. In docker, you have the internal docker networking play, so the controller only see’s the IP address of the container, which other devices (like your AP’s) can’t route to. By forcing the device to inform to the host devices ip, the connection works. The long term work around is to use the “override inform url” option and set it to your hosts static IP. Adoption should work just fine that way.
Thank you. I thought it could possibly be docker related. Even though I have the container in host networking mode. But it just puzzled me how it was detected without me doing anything, but refused to adopt it. But fortunately I've got them added now.
Discovery uses a whole different system than adoption, so many times one works without the other working
