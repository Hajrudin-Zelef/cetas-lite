---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-14m83an-new-to-unifi-setting-up-vlan-tags-on-trunk-ports-bd9ac12c
title: "r-ubiquiti-comments-14m83an-new-to-unifi-setting-up-vlan-tags-on-trunk-ports-bd9ac12c"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-14m83an-new-to-unifi-setting-up-vlan-tags-on-trunk-ports-bd9ac12c.md
source_anchor: ""
source_lines: [1, 31]
sha256: aba993666deaade5e0db1464698a3ea0fc20b57fc50035660ea8ae5017943596
---

# r-ubiquiti-comments-14m83an-new-to-unifi-setting-up-vlan-tags-on-trunk-ports-bd9ac12c

New to Unifi - setting up VLAN tags on trunk ports? 
        
        
        
    
    
    So im new to Unifi, im testing out a Unifi switch to see if it would potentially work for us as an edge/small office solution. We have always used Cisco small business so thats where my background is.
Im very confused on how Unifi handles VLANs. We have been using ubiquiti APs for quite a while now, and its always worked pretty simply with our cisco switches. Our firewall hosts the VLAN and DHCP for that vlan, we tag our trunk ports with those vlans, and then we create the VLAN-only networkf for the wifi on unifi. Simple.
So Im taking a unifi switch (USW Pro 24 POE) and ive created VLAN only networks, one for each wifi (guest and corp) and the VLAN for our voice (grandstream IP phones)
Under the profile for ALL, my voice vlan is checked, but my IP phones are still pulling DHCP from my DC, rather than the firewall voice VLAN where it would if it was plugged into a cisco switch with tagged trunk ports.
What am I missing here? The way UB handles vlans is confusing.
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Ok, so now that I’m am at my PC and re-read you post:
Did you create a new port profile
set your corp as default (or untagged)
Under advanced select your voice network
Under advanced make sure LLDP-MED is enabled
After saving apply this port profile to the switch port you are testing with
Now that I have thought some more, the broken part is UniFi talk phones not honoring lldp-med voice VLANs, not that the process is broken for all phones
Okay, so I created a profile, called it default+voice, and tagged the voice vlan 2222 in it, as well as the default vlan. I also tagged the uplink port with this same profile. Still the same deal, the phone is pulling DHCP from the local lan.
Your uplink likely needs to be set to all
Is 2222 the VLAN your voice is using up stream?
UniFi does not use any “smart“ process that would allow the phone to know what vlan is currently set up for voip. You have to tell the phone what vlan to connect to
Interesting, so you are saying I would essentially have to dedicate that port to the phone/voice vlan. But then how would computers that are plugged into that phone know to look to the native VLAN for their DHCP? Is this something the unifi switch is not capable of?
I am mobile right now and can’t look up what I’m remembering to share, But on a Cisco switch where you set your voice VLAN and the phone will automatically pick up that VLAN, that process has been broken on UniFi switches for a while. At least it was last time I checked.
It’s supposed to work exactly as you’re thinking, it just hasn’t for a long time
