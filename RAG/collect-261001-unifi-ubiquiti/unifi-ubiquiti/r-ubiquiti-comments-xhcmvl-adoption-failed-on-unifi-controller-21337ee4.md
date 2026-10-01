---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-xhcmvl-adoption-failed-on-unifi-controller-21337ee4
title: "r-ubiquiti-comments-xhcmvl-adoption-failed-on-unifi-controller-21337ee4"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-xhcmvl-adoption-failed-on-unifi-controller-21337ee4.md
source_anchor: ""
source_lines: [1, 52]
sha256: f8eff4349eb04c5e654ca3bafff4f8a65b13823eab42087dcb5b4b1e7e4dfb86
---

# r-ubiquiti-comments-xhcmvl-adoption-failed-on-unifi-controller-21337ee4

Adoption failed on unifi controller 
        
        
        
    
    
    Hi guys
I am using unfii ap ac LR and controller 7.1.68
As seen in the photo, I encountered an adoption failed error after updating the controller.
Sometimes the problem is solved by physically restarting the access point, but after a few hours, the problem occurs again.
This is while the access point is available and ping is established.
Access point and controller on the same network
Someone has a solution?
Thanks
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
It's been happening to me too on the latest controller with U6-LR's... They suddenly start adopting again after being fully functional for like a day or two and i sometimes have to restart then for it to be comnected again
I had an issue this morning where my APs (which are all wired) decided to connect wirelessly to 1 specific AP via meshing. This resulted in 1 of the 3 AP's going offline.
I turned off "use global settings", meshing, and increased tx power to high. All is well now.
ALL NEW TO THIS I just bought U6 Lite and 6 mesh Both are in Adoption failed loop Firmware and software up to date. I have done both the factory reset (hold till change color) And the long reset ( hold till turns blue then hold another minute) de-energized and re-energized everything, tried different Cat cable. Verified Mac and IP addresses. Even tried with my Android to connect. Did advanced adoption using ubnt/ubnt. Turning the firewall off also, still no luck
24PoE port 100/1000Mbps switch
HELP !
I've been having issues with U6 devices as well. Seems it's a battle to adopts, often have to force the switch port to a specific network, and reset 4-6 times then they finally adopt, and a week later they switch to "Managed by another console" suddenly.
Right now I have 2 U6-Mesh that refuse to adopt no matter what, and both had been previously adopted and worked for a couple weeks. Then "Assigned to another console", remove, reset, reboot, SSH restore, nothing seems to work now.
Also have a brand new 48 port Enterprise switch that worked for 2 weeks with no devices connected followed by a week in production, and is now "Managed by another console", it's in a live environment so I can't just reset it and re-adopt, it's stuck in limbo because it functions, but I have no control over it currently.
I've found that most of my adoption failures go away if I move the device to a switch directly connected to the controller. I think the latency of the extra hops contributes to some of the issues with adoption.
And I'm NOT claiming this is a universal fix.
But after a while you see a lot of us who do this pile all the APs on the desk and adopt them BEFORE putting them in the ceiling or wherever.
Check the DHCP server on your network. Make sure it is enabled and properly assigning IPs and not configured with the wrong inform host
Manually update to the latest firmware if the device has never been adopted before
For previously adopted devices, forget device within unifi, factory reset, power cycle and the adopt again.
When this happened to me I had to do the following.
On the AP press and hold the reset button until the AP starts blinking.
Continue to hold the rest button and power off the AP. DO NOT LET GO OF THE RESET BUTTON.
Wait for 30 seconds. While holding the reset button on the powered off AP.
Apply power to the AP while still holding the rest button until the AP starts blinking.
You should now be able to adopt the AP.
Forget the AP via cloud key, 2. Factory reset AP (paper clip, 6 seconds) 3. Adopt
Spent 6 hours doing that on mine. The only thing that worked was what I posted above.
Rarely do I have to do a firmware update on the AP via SSH when adoption fails, then after the update it adopts fine.
Ok, this is quite an edge case I'm sure, but it may help someone...
This happened to me last night after I installed pihole (that's a spoiler right there!)
So I reset one of the devices (a PoE switch) and ssh'ed into it. The logs were a give away:
Oct 31 21:03:43 US-8-150W daemon.err mcad: mcad[896]: ace_reporter.reporter_fail(): initial contact failed #41, url=http://unifi:8080/inform, rc=6Oct 31 21:04:13 US-8-150W daemon.warn mcad: mcad[896]: url_resolv.evdns_resolv6_cb(): dns6 host unifi resolv failed, result=70/no records in the reply, type=0, count=0, ttl=0
Because I'd whisked DNS away from the unifi application and pushed it to pihole it now couldn't resolve the device. Adding "unifi" and the ip address as a route in pihole fixed it and everything adopted within a minute or two.
Again, I know this won't help everyone, and I've certainly had my fair share of the issues and "just juggle it" solutions presented here.... thought this edge case may be useful to someone else though
You just use UNIFI as the FQDN? No .local or anything like this?
Basically, I was using USG as the DNS and that was the problem. Now I run USG as the local DHCP, but, that has recently been a PITA as it randomly turns DHCP off 🤦
