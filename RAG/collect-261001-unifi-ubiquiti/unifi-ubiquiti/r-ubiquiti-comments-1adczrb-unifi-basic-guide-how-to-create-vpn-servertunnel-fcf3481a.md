---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-1adczrb-unifi-basic-guide-how-to-create-vpn-servertunnel-fcf3481a
title: "r-ubiquiti-comments-1adczrb-unifi-basic-guide-how-to-create-vpn-servertunnel-fcf3481a"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-1adczrb-unifi-basic-guide-how-to-create-vpn-servertunnel-fcf3481a.md
source_anchor: ""
source_lines: [1, 42]
sha256: 853a6debe22e0d80186847aa49e8bdce2ab092035519498802d465e4a5f8a28a
---

# r-ubiquiti-comments-1adczrb-unifi-basic-guide-how-to-create-vpn-servertunnel-fcf3481a

UniFi Basic Guide: How to create VPN Server/Tunnel 
        
        
        
    
    
    Intro
VPN server is to make our UniFi network into VPN service provider. This allows me to remotely connect to my own home network as if I am locally in the network.
UniFi provides two main methods/approaches for this.
- 
      Teleport method
- 
      VPN Server method
Teleport vs. VPN Server
Teleport is a hair touch easier to configure when compared to VPN Server approach because VPN Server method is already fairly easy.
For Teleport approach, client device must be able to run WiFiMan App by Ubiquiti. The primary benefit of VPN server approach is its flexibility. This approach allows VPN client setting on almost on any device while providing several customization options.
For the throughput, since Teleport uses Wireguard protocol, two are identical with ~15% of download throughput reduction when compared to no VPN on my test.
General steps:
Teleport method
- 
      Install WiFiMan on the client device
- 
      Enable Teleport on Network Controller
- 
      Link client to VPN
VPN server method (w/ Wireguard)
- 
      Install Wireguard client on the client device
- 
      Create VPN Server entry on Network Controller
- 
      Create client pass on VPN Server
- 
      Register the client pass on the client device
Creating high performance VPN server with UniFI system is super easy. This is my preferred approach of using home network attached services like Plex. I hope someone finds this helpful.
Original article: https://gameandtechfocus.com/unifi-keep-it-simple-unifi-basic-guide-how-to-create-vpn-server-tunnel/
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
