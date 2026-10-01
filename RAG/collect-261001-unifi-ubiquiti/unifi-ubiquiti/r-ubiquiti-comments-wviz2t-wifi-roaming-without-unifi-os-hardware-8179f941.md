---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-wviz2t-wifi-roaming-without-unifi-os-hardware-8179f941
title: "r-ubiquiti-comments-wviz2t-wifi-roaming-without-unifi-os-hardware-8179f941"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-wviz2t-wifi-roaming-without-unifi-os-hardware-8179f941.md
source_anchor: ""
source_lines: [1, 41]
sha256: 743ae2c48179bc61346c189779db962e07eab499b47ac1eeb82b77e20426585c
---

# r-ubiquiti-comments-wviz2t-wifi-roaming-without-unifi-os-hardware-8179f941

WiFi Roaming without Unifi OS hardware? 
        
        
        
    
    
    I'm looking to put a few U6 LR throughout the house, and buy a PoE switch to connect them together (in addition to other ethernet connected devices).
Assuming I don't need user management (a single password for the family is enough), any remote management feature, etc... I just want the internet to "work", for a single-family home.
Do I still need additional Unifi hardware, such as the "consoles" they sell?
I'm trying to keep is both simple and affordable.
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
You can self host the controller to keep the cost down, you don’t have to do that as you can setup the AP without a controller, but having the controller in place gives you much more control
What control are we talking about? I mean, like I said it’s for a single family home. I just want to set it up once and forget about it. If I have to “control” things this sounds overkill. The only thing I can think of is to make sure roaming between access points is seamless.
Even for things like setting the SSID. You don’t need a controller but configuring without one via ssh is a PITA and each individual AP needs to be configured each time a change is made. You don’t need to be running the controller constantly but can instead run it when necessary from your pc or Mac after ‘adopting’ the devices.
Things like changing the bandwidth on the AP,
You may be able to do this if you use the app to individually set up the AP. I’m just not sure what options are Available. You could always buy one AP, set it up using the app. See how it goes
I realize this is the Unifi sub, but Unifi is not the only game in town for APs, especially if you are not planning on diving into their cameras/doorbells/etc.
Aruba has an instant on line that uses a cloud-based controller.
https://www.arubainstanton.com/products/access-points/
Grandstream makes APs that have a built-in controller, accessed through a local browser pointed at the master AP's IP, so you have a local controller, but you don't need a dedicated device to host it.
https://www.grandstream.com/products/networking-solutions/wifi-access-points
Thank you I will look into that!
Go pimp for them on their sub.
My bad. All hail Ubiquiti, the one and only true maker of wireless APs fit for home use. :)
Lol...
I run my house on a raspi pi that is already being used as a Pi-hole. I leave the controller running because I like to get on and look at stats from time to time. I have one indoor ap and on outdoor. They work just fine without the controller though.
Am I reading this and the Unifi site correctly, that to to have 3 U6 Lite APs with seamless roaming/mesh a CloudKey or similar box is required?
I also see mention of controller software on local PC/Mac as an option but do not see where to download it, does it still exist?
While you can setup the APs in standalone mode with the Unifi app, if you want any of the advanced capabilities like seamless roaming or Mesh (wireless uplink), or even basic statistics, you will need to use the Unifi controller, either one of the hardware based ones or the free to download software. And the controller must be running for the previous mentioned capabilites. If you just use the controller software to configure the APs and then shut it down, the APs will still work but they won't roam or Mesh. This is true for most ecosystems.
If you don't want to mess with controllers or software controllers, then maybe look into some used Ruckus APs. They don't require controllers to function. You setup the first one and then each one you add to the Mesh will automatically just work. Just keep in mind that the new equipment will be pricey for a home. So go with the older used equipment that you can get or eBay if you want to go with Ruckus.
Commentaire supprimé par le membre
IMHO, yours is the most useful answer. As long as all the APs are configured to use the same SSID and password, the clients will handle the roaming on their own. I used a collection of Apple Airport Extremes and Expresses for years at home until graduating to a more professional setup and everything roamed just fine.
The biggest advantage of setting up the equipment in the controller would be the ease of making sure channels don't overlap, but if OP only has a couple APs, then even that won't be too bad to do without.
Commentaire supprimé par le membre
A switch from Unifi is different than a switch from any company ?
Commentaire supprimé par le membre
