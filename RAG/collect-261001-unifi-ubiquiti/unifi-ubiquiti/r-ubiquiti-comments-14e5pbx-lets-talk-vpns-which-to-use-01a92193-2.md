---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-14e5pbx-lets-talk-vpns-which-to-use-01a92193-2
title: "r-ubiquiti-comments-14e5pbx-lets-talk-vpns-which-to-use-01a92193"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple", "Intel"]
dates: ["2023-06-20"]
keywords: ["intel", "latency"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-14e5pbx-lets-talk-vpns-which-to-use-01a92193.md
source_anchor: ""
source_lines: [7, 63]
sha256: 543c5572bf274611a9c4fa87d73938d6e0d890abb63c52360514efb34b85e708
---

# r-ubiquiti-comments-14e5pbx-lets-talk-vpns-which-to-use-01a92193

    I have been running Unifi for years and have an L2TP VPN set up that works fine. I use the VPN when I am on the road to check back on the network. I also use it to stream Plex music outside the house (double NAT and I don't like the idea of having Plex ports open...)
About 6 months ago I enabled the Teleport option as a backup and that saved me when I had an L2TP hiccup on one trip. But teleport only works with iOS devices, not Macs.
Now I see that there are Wireguard and Open VPN options in the Unifi controller.
L2TP seems to be fast enough to connect and also the performance is acceptable. I am never doing any large applications or traffic; I have occasionally done screen share over VPN and that was acceptable.
Knowing that L2TP is working fine for me, is there a compelling reason to enable either of these?
My primary concerns are simplicity and security. If something is 20% safer but 80% more complicated to manage, it might not be worth it for me.
Interested in your thoughts.
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
I think L2TP is being deprecated in future Unifi versions.
I recently switched to WireGuard and have had a better user experience so far. The WireGuard client setup allows me to automatically connect to my VPN whenever I am joining unknown WiFi networks. Good stuff!
This is one of my concerns.
I will look into Wireguard when I get back from my trip. It works well with Mac I take it?
Works very well with Mac and iOS devices 👌
I think you’ll find WireGuard to be a much better option. WireGuard tends to be faster, have lower latency, and will definitely run into fewer firewall issues. You can also configure your clients to run on-demand so that they connect anytime they’re away from your home network. I’m running it on all our mobile devices (iPhones, Macs, & iPads) with on-demand and all traffic sent over the VPN and it runs flawlessly without me ever having to think about it.
How did you get WireGuard to connect whenever you are on an unknown WiFi? Is that an iOS exclusive feature? I don't see a setting for this in the android app.
It’s called “on-demand activation” in the connection settings in the iOS WireGuard app.
I’ve been able to use teleport with Apple silicon Macs by using the wifiman app.
Hmmm, need to try that, will work for our newer Macs, but not the Intel ones.
My hand was forced when Windows 11 quit supporting L2TP. I landed with WireGuard in a docker before it was available directly on my UDM-SE. I currently have the docker and the UDM-SE both serving WireGuard from on different ports. Once I get all the clients reconfigured I'll be taking down the docker instance.
An you use teleport on your iOS device and then Hot spot to the iOS device?
Yes, I can do that but that is not what I want to do.
Just use l2tp it works fine (I also run Tailscale on my nas so family can access)
Works fine for now. But support is being dropped by most of the industry so I'd rather make the change now while I am in control than have it get dropped when I am hundreds of miles from home.
I just use it because I can configure it on any device without an app, also that’s why I have Tailscale
Until you have to connect two devices on the same remote network to your home/office. Impossible to do with L2TP.
I personally use OpenVPN although I’ve only heard good things about wire guard in general
I use both - OpenVPN for the Unifi native VPN client connectivity (so that I can route certain remote networks via my home lab) and wireguard for individual clients like phones/tablets/laptops.
OpenVPN is now much easier to set up and works well. Stable and reliable. Wireguard is, however, significantly faster. I’ve been travelling with my phone across Europe for the past three months and had the Wireguard VPN client active the entire time - there is no difference in speeds whatsoever, it’s just great. The setup is kind of odd: you have to add the new user, click save, then go back and download the certificate; do it any other way and the file won’t be correct and the connection will fail (including redownloadin the configuration file for an existing user, that will also fail). This is purely Ubiquiti being Ubiquiti - the wireguard component, once set up, works great.
How you guys configure WireGuard?? I activate the VPN Server in UniFi Network, my clients connects, but they never are able to handshake or route any traffic to the UDMP (that is remote to me). Always the same issue, I tried on severy type of clients (MacOS, Windows 11, Windows 10, etc.). All using the Official Client, all with the config generated in the UniFi UI that I downloaded and imported in the Client. Also, tried changing to allow 0.0.0.0 and a also tried specifying Custom DNS (8.8.8.8). I can't ping anything, even the IP Address I connect to.
2023-06-20 19:26:44.553 [APP] Status update notification timeout for tunnel 'wg_c0_a8_03_02 (2)'. Tunnel status is now 'connected'.2023-06-20 19:26:45.105 [NET] peer(2/8V…3Swc) - Handshake did not complete after 5 seconds, retrying (try 2)2023-06-20 19:26:45.105 [NET] peer(2/8V…3Swc) - Sending handshake initiation2023-06-20 19:26:50.269 [NET] peer(2/8V…3Swc) - Handshake did not complete after 5 seconds, retrying (try 2)2023-06-20 19:26:50.269 [NET] peer(2/8V…3Swc) - Sending handshake initiation2023-06-20 19:26:55.592 [NET] peer(2/8V…3Swc) - Handshake did not complete after 5 seconds, retrying (try 2)2023-06-20 19:26:55.593 [NET] peer(2/8V…3Swc) - Sending handshake initiation
Using wireguard. Pretty good and easy to set up.
100 OpenVPN Access server everywhere
I’ve used openvpn for years but am now using wireguard on the udmse and couldn’t be happier. I couldn’t get openvpn on the udm to authenticate username/password.
So, I recently implemented Wireguard on a UDM Pro SE that was already configured for L2TP. I was curious to see if there was a performance difference between the two.
First I compared Speedtest results from both sites using local systems and then the speed from a remote computer connected by the VPN tunnel.
Speedtest - Local Site
235.50 Mbps / 47.89 Mbps
Speedtest - Remote Site
119.87 Mbps / 19.61 Mbps
Speedtest over Wireguard - Remote Site
18.74 Mbps / 17.96 Mbps
Speedtest over L2TP - Remote Site
18.41 Mbps / 20.00 Mbps
Not a huge difference, but L2TP seemed to max out the upload of the remote site more effectively.
Then I started a file copy operation of 110MB from both VPNs from a Windows desktop system at another site to an SMB share. I used Teracopy for the transfers so I would have accurate transfer duration metrics to go off of.
L2TP - file upload in 75 seconds
Wireguard - file upload in 105 seconds
That's nearly a full 50% difference in transfer performance. L2TP was 40% slower than the theoretical max established by the Speedtest results, but Wireguard was 58% slower, utilizing only 8 Mbps of the 18.7 Mbps max.
Then I transferred the same file back.
L2TP - file download in 54 seconds
Wireguard - file download in 56 seconds
That was much more consistent and near the theoretical max that was probably constricted by the upload performance at the remote site.
