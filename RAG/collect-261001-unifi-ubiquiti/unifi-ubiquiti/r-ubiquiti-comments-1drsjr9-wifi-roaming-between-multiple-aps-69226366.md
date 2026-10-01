---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-1drsjr9-wifi-roaming-between-multiple-aps-69226366
title: "r-ubiquiti-comments-1drsjr9-wifi-roaming-between-multiple-aps-69226366"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Samsung"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-1drsjr9-wifi-roaming-between-multiple-aps-69226366.md
source_anchor: ""
source_lines: [1, 33]
sha256: 3f2d9fadbe41d1759ab9b0d1ed70ef8180e60605349eea932a8b1a8a9561cc68
---

# r-ubiquiti-comments-1drsjr9-wifi-roaming-between-multiple-aps-69226366

Wifi roaming between multiple APS 
        
        
        
    
    
    Hi
New to Ubiquiti. Got 2 APs and want the ability to roam from 1 AP to the other whilst walking around the house. I received a new AP today and during config was given the option to configure using a gateway or standalone. Do I need to buy a ubiquiti gateway to get the ability to roam between AP's? If not then what do I need to buy to enable roaming?
In the past I just configured the APs standalone but want to do it right this time. I have the following:
Asus router U6 pro UAP-HD-IW
I'm sick of manually connecting to the closest AP.
THANKS
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
To configure and use stuff like BSS transition and fast roaming on Unifi APs, you need to have a Unifi controller. The controller doesn't have to be running all the time, but it is nice for it to. The cheapest/easiest is just to install the controller software on a PC and adopt your APs to it. Other options would be to buy a Cloud Key, or replace your Asus router with something like the Unifi UDR or Unifi Express that has the controller built in.
Ive got the controller installed on my Windows PC. However Ive never been able to roam from AP to AP. Not smart enough to understand what RSSI is or where to change it. In summary I have the following:
Asus RT-AC68U (Garage)
Cisco WS-C3560G-24PS Switch (Garage)
3 x Hikvision CCTV cameras
Ubiquiti U6 pro (TV Room)
Ubiquiti UAP-HD-IW (Bedroom)
The ASUS and both Ubiquiti devices have the same SSID/encryption keys. I dont expect the client roaming to work if going from Asus to Ubiquiti but want ubiquit AP to ubiquiti AP to work. Can I get client roaming to work if I manage the APs in standalone mode via WIndows unifi controller? At the moment I need to stop wifi and re-enable it when I change rooms to ensure it connects to closest AP (mobile phone is a Samsung S24)
If better Im happy to replace the Asus with another Ubiquiti device. Dont want to spend more than $300 though and just need to ensure existing devices on my switch continue to work (I heard there are some client limitations on UniFi Express)
Thanks
The roaming features in the wifi standards are meant to make the transition from one AP to another smooth without interruption, but do not guarantee that you will always be connected to the closest AP. If you are connected to one and the connection is good, the device will normally not be looking to roam to another AP. If the device is being a little too "sticky" and not switching even when it has a bad signal, then adjusting the RSSI can help that by having the AP force a disconnect when the signal gets too weak, at which point the client's only option is to move to the stronger signal. I also have to wonder if having the isolated Asus wifi on the same SSID is confusing the client somehow in this scenario. I would recommend starting by changing your Unifi wifi network ssid to something different from the Asus, and see how it behaves before doing anything else. Replacing the Asus router with a Unifi device would be more ideal and give you more options. UDR is certainly in your price range at $199...if you can catch it in stock. The Express does have strict limits, but is well suited for a simple home network with a couple APs and <50 client devices.
Your computer needs to run 24/7 to enable the roaming.
There must also be overlap between AP coverage areas.
Roaming is not really anything special. Whether or not a client “roams” is up to the client.
If your SSID and encryption keys are the same the client will “roam” to the other AP. You can manipulate this by using RSSI minimums in the AP configuration.
