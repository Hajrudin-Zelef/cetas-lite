---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-1dcdsc3-site-to-site-vpn-with-one-end-behind-cgnat-434a11c8
title: "r-ubiquiti-comments-1dcdsc3-site-to-site-vpn-with-one-end-behind-cgnat-434a11c8"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-1dcdsc3-site-to-site-vpn-with-one-end-behind-cgnat-434a11c8.md
source_anchor: ""
source_lines: [1, 29]
sha256: 8cb795be5f7176ad9d61ceca1f727323570b57a1f283e878d50f9ff4c88cbb59
---

# r-ubiquiti-comments-1dcdsc3-site-to-site-vpn-with-one-end-behind-cgnat-434a11c8

Site to Site VPN with one end behind CGNAT 
        
        
        
    
    
    I have a Ubiquiti Cloud Gateway Ultra which is on a location(A) that has a Public IP acting as a VPN server. I have another location(B) which is using LTE connection to access the Internet. On that location, the ISP provides an IP address which is behind a CGNAT.
I have setup an openVPN connecting as client towards the Cloud Gateway Ultra(server), but I can't set it up, so that biderectional traffic is allowed. At the moment, only LAN clients of the LocationB can access the LAN subnet of LocationA, but not the other way around.
If both locations had Public IP, I could have setup a parallel openVPN connection on the opposite direction and establish bidirectional traffic.
Any ideas how I could setup the openVPN to allow bidirectional traffic through a single tunnel?
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
I think Unifi site magic is supposed to work automatically even with one site behind a CGNAT so that may be worth trying out
Yes, this is my backup plan, I just would like to explore If I can achieve bidirectional traffic with my current setup.
Is Site Magic SD-WAN an option?
https://help.ui.com/hc/en-us/articles/16750417515159-Introducing-Site-Magic-SD-WAN
Use this for a few clients. Seriously is like magic. A few clicks and it just works.
thanks, I will definately consider this option if I can't make it work with my current Asus router.
Unsure if you can get yours to work, I’m not super familiar with CGNat so good luck. Site magic will work for you though if you can’t get it up like others have said. You just click a few buttons and it all works flawlessly
I e used openvpn over CGNat before and it should work fine. It sounds like a config or routing issue to me.
Either magic sd wan, or use site A as a wireguard server, and other site(s) as wireguard clients. If you use “manual” config when creating the client, you can specify remote subnets for the client side, creating a S2S style vpn.
Bonus: performance will be much faster than OpenVPN
You need a rule on the remote site behind CGNAT permitting Internet In with the source being the remote network and the destination being the local network. UI sucks. I came here looking for a solution and then not finding one other than site magic started poking at it with Wireshark. This fixed it for me. It's lame... but hey that's UI.
Thank you for doing the leg work on this. This helped me a lot. Added the rule on the remote site (behind cgnat) and went from connecting to connected.
Glad this helped someone.
