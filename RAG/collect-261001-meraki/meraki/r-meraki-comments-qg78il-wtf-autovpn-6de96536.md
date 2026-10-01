---
id: collect-261001-meraki/meraki/r-meraki-comments-qg78il-wtf-autovpn-6de96536
title: "r-meraki-comments-qg78il-wtf-autovpn-6de96536"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-qg78il-wtf-autovpn-6de96536.md
source_anchor: ""
source_lines: [1, 21]
sha256: 456a3233184c3d316495d5a9735d7189a0cd03b10f4831fa3269fdfdb0d58ac0
---

# r-meraki-comments-qg78il-wtf-autovpn-6de96536

WTF auto-vpn?! 
        
    So I've got around 100 various MX firewalls that I have set up as spokes to my one hub (MX100 pair) in my DC. All of a sudden, a whole bunch of spokes decided to not reach my hub at the same time. I have probably around 15 ~ish spokes like this. I replaced one, and now that box can at least see the internet but it still won't connect its auto-vpn tunnel to the hub. All internet connections at all sites have been verified, so the issue has to be somewhere with the Meraki device(s) and not the provider level (as there are multiple providers going on).
Has anybody ever seen anything like this, and if so, where in the hell do I even start troubleshooting this mess?!
Section des commentaires
Did this by chance happen yesterday at about 10 or 11 AM central time?
https://www.reddit.com/r/networking/comments/qfm60p/meraki_down/
I wish it were that easy. These have been down since the 21st, or the 23rd at the latest.
I avoid Meraki most of the times. But I found two ISP were blocking UDP 500 and 4500 for all of the customer after an ISP big security change. I discovered it while capturing traffic for both ports in both sides and one HQ was negotiating phase 1 but NOT phase 2, because packages weren’t arriving the interface.
I'm trying to avoid engaging the providers since these internet connections technically belong to the end users, but I can see providers doing something like this. My concern right now is that most of these devices that aren't connecting to our DC are also not connecting to the portal at all, though they were working just fine until late last week.
This is common for old MX versions, please consider upgrade, I did it while reseating the device to default. Connecting it via DHCP and it will sync to the dashboard. Also, it happened to me that the MX would connect to the dashboard after 24 hours, so, I don’t know what’s going on there
Take a look at these documents. Once you have a good understanding for how the VPNs are setup (via UDP hole punching), it's pretty easy to troubleshoot or find the root cause. Generally, I've had issues with ISP modems before that for whatever reason start blocking the UDP ports.
https://documentation.meraki.com/MX/Site-to-site_VPN/Meraki_Auto_VPN
https://documentation.meraki.com/MX/Site-to-site_VPN/VPN_Status_Page
Start by taking packet captures on your hub and a single spoke and filter for just the hub and spoke's UDP ports. See if there's bidirectional traffic. If there isn't, then something is blocking the traffic upstream.
So wouldn't ya know, after tons of packet captures and beating my head over the desk, I found the problem to be a defunct traffic inspection policy on my edge firewall (hell no my MX firewall isn't my edge!) that'd caused a TON of stale connections. Remove bad policy, clear all connections, magically the tunnels came right back up!
You are about to learn the hell of meraki VPN troubleshooting. I know this isn't going to be helpful but you're going to have to contact their tac, cuz the equipment is so awful at logging any real information, and the fact you don't have any real routing information to deal with either. They're going to tell you destroy your entire organization and rebuild the networking for the lulz. Then you're going to start seriously thinking to yourself about ripping out all that meraki crap and putting in real networking hardware instead. Short version you're going to have to contact meraki there's no other way or anything you're going to be able to do at this point
Oh this job has schooled me about Meraki...the first thing I learned is that Meraki might be good for branches, but it's worse than terrible to put it in a data center. I've already yanked 2 switches out of my DC, and have a project in to mgmt to pull another 4. I like the auto-vpn functionality, but it looks like those days might be numbered, too.
Yeah they had some great sales people that really pulled the wool over some people's eyes for a while there.
Commentaire supprimé par le membre
Many aren't even connecting to the dashboard, but I replaced one, it's connected to the dashboard, but the auto-vpn just won't come up.
