---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-zl0227-vpn-teleport-wifiman-wireguard-so-confused-bb260919
title: "r-ubiquiti-comments-zl0227-vpn-teleport-wifiman-wireguard-so-confused-bb260919"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-zl0227-vpn-teleport-wifiman-wireguard-so-confused-bb260919.md
source_anchor: ""
source_lines: [1, 31]
sha256: c139c25bad9a3ce316dec85705b020a98abd909a11a145939678e63feecaf880
---

# r-ubiquiti-comments-zl0227-vpn-teleport-wifiman-wireguard-so-confused-bb260919

VPN - Teleport? WiFiman? WireGuard? So confused 
        
        
        
    
    
    One of the reasons I went with the UDM-SE is because I liked the idea of locking down Internet access on my home network and yet still being able to VPN in to control or view things, like cameras. I would want the access to be the same as if I were there at home on my local network.
I just started diving into the VPN settings and re-watching videos that cover VPN and have found myself very confused. Mostly because there seems to be several ways to use VPN. A lot of the Youtube videos are also a little out of date now since the release of the latest software. The Mactelecom video from 8 months ago shows setting up a VPN server, manually entering all the server and client settings. The Crosstalk video from 3 months ago shows setting up Teleport and generating an http link that gets opened in a client's WiFiman app. The video from Lawrence Systems from 3 months ago seems to go over 3 different ways of setting up VPN, but I don't understand which one is the best or fits my use case.
So to give a concrete example, I'm looking to disable Internet access to my Reolink camera system. The only Internet access I would give is to open the firewall to allow push notifications from the NVR. When I'm away from the house, I would VPN into my network, (as if I were actually home on my network), and view or control the cameras. Since I have the Reolink client on both my Mac and iPhone, I'd like to be able to VPN using either device.
Knowing that, can someone steer me in the right direction of which VPN method might be best or which video best fits my use case?
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
I use Wireguard, IPSec and Teleport into my UDM SE. I found that the Teleport is very easy to setup and it uses Wireguard underneath. Straight Wireguard is faster and IPSec is old school but just works too. Watch u/mactelecomnetworks latest Wireguard setup video here:
https://youtu.be/zGwZGZyAKNs
Thanks for the mention :)
u/mactelecomnetworks thanks for all the great videos. Not sure why I haven't seen this latest one, but I'm watching now!
Thank you so much for linking to the video. I'm subscribed and watch all of Cody's videos, but I don't know why this one hasn't shown up in my search. It's the most recent!
They did implement Wireguard.
If you can use the native wireguard server, that’s what I’d recommend.
Teleport doesn’t work on T-Mobile and hasn’t since it came out. So that’s something. Works fine on WiFi… but I really wish it worked on my cell connection too.
i cant comment prior to a month ago, but have you tried it recently i use Teleport on T-mobile. if it matters to u, i use it on 5G as that's what i get around my area. Hopefully its working for you now!
Oooh maybe that’s it! I’m still using an iPhone 11 Pro so no 5G yet and teleport still isn’t working for me. I’m gonna get a new iPhone this fall so hopefully it works after that.
Commentaire supprimé par le membre
I have an M1 iMac and was able to import the WiFiMan app to that desktop.
yes regular WireGuard as a server is implemented. Unlike some other WireGuard implementations you don’t get a QR code. But easy peasy to copy and paste the config into your client.
But, they did implement regular wireguard.
This.
