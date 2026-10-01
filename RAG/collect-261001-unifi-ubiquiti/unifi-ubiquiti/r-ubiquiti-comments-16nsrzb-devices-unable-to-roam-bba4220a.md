---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-16nsrzb-devices-unable-to-roam-bba4220a
title: "r-ubiquiti-comments-16nsrzb-devices-unable-to-roam-bba4220a"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple", "Intel"]
dates: []
keywords: ["intel"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-16nsrzb-devices-unable-to-roam-bba4220a.md
source_anchor: ""
source_lines: [1, 39]
sha256: 13e1ebd265bd13c4f765356a423b9d451aa60c7a006cb83fc86f387e31c23f15
---

# r-ubiquiti-comments-16nsrzb-devices-unable-to-roam-bba4220a

Devices unable to Roam 
        
        
        
    
    
    I do technology for a school district. We have 200 Unifi UAP-AC-PRO (WiFi 5) devices across 6 sites (1 controller).
At one of our school sites, users are having trouble roaming. Yes, we have fast roaming enabled. Our other 5 sites do not have any of these issues.
Users move from one area (with WiFi from 1 UAP) to another area (with WiFi from another UAP) in the same site. These are Macbook Airs, one an Intel processor, one with an M2 processor. Moving between UAPs, it can take 5 minutes or longer to get back on the Internet.
I have recently replaced a number of these problematic UAPs. Brand-new UAPs still have the same problem (not providing an Internet connection quickly). The computers appear on the Unifi controller connected to the UAP but often have a self-assigned IP - they are without Internet. From the Unifi controller, I have used the "reconnect" command, but the devices reconnected with self-assigned IPs.
      We are buying some U6-plus devices to see if they perform better.
We have replaced multiple computers (no computer changes have made a difference). One user has had three different devices (completely different computer models) and all behave the same.
    
Ideas for what the issue might be? Suggestions for moving forward?
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Fast roaming can be problematic with some clients. Turn that off and see if things improve.
Some SSIDs have it off, some on. The SSID we have had trouble with (the Staff SSID) has Fast roaming turned off. Our Student SSID has it on - not sure why the difference, I've been here a short time and inherited things as they are.
Ok, in that case, we're back to setting appropriate channels, Tx levels, min RSSI, etc.
My bet is that there is something weird with channel overlap and the device thinks its connected to one AP, but the APs think it is connected to a different one.
Also, check to see that the SSIDs are all configured to the APs you want them... it could be that someone has done a manual override and the AP you think you should be roaming to isn't even servicing that SSID.
AND... as a final effort, take a Saturday and wipe all the AP configs and start fresh. There's nothing worse than inheriting someone else's bad config.
We reached out via Reddit Chat requesting any existing support ticket numbers so we can escalate and review. Thank you.
Created a support ticket: 3852676
Clients generally control the roaming. But a few things come to mind.
How is power and channel controlled? Do you manually set them or do you let Unifi auto 'optimize' them?
If you don't have Min RSSI set, you can try that. That way once the devices reach a certain signal threshold, they are kicked off the AP, then the clients 'should' attach to one with a stronger signal.
I have Unifi set to optimize power and channel.
I don't have min RSSI set, I will look into that (it sounds promising).
Self assigned IP addresses sounds like a DHCP issue to me. Is your lease space large enough?
Disable minRSSI and fast roaming. Set min data rate to two values above the default. Test. Adjust and retest.
Also, nothing beats a site survey to determine as-installed coverage and required adjustments to AP tx power levels.
Negotiated rate takes into account received signal level, bit error rate, etc, so it’s a better canary to trigger client roaming than individual knobs.
can you see if you have big enough IP for DHCP?
Apple computers are a pain in the ass ever since 802.11ac was released apple removed roaming decisions from the ap and controller and left it to make its own decisions. If you don’t have 802.11r enabled at the minimum the clients won’t roam until their rssi hits -70. I had problems a couple weeks ago and disabled uapsd and just left 802.11r enabled and the roaming issues I was having went away. These days apple recommends r/k/v be enabled but not all wifi devices support.
