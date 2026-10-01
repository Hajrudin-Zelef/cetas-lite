---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-17idi6u-how-do-i-set-up-a-wireguard-or-teleport-vpn-in-2594cdb7
title: "r-ubiquiti-comments-17idi6u-how-do-i-set-up-a-wireguard-or-teleport-vpn-in-2594cdb7"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-17idi6u-how-do-i-set-up-a-wireguard-or-teleport-vpn-in-2594cdb7.md
source_anchor: ""
source_lines: [1, 24]
sha256: c7f050ae713b87bc377647bb9a9e05d900560b57342b1c01e833381d8b02db67
---

# r-ubiquiti-comments-17idi6u-how-do-i-set-up-a-wireguard-or-teleport-vpn-in-2594cdb7

How do I set up a Wireguard or Teleport VPN in conjunction with Magic Link point-to-point VPN? 
        
        
        
    
    
    Equipment:
      Site 1: Cable Internet via MB8611 Modem. <---> UXG-Pro <---> US 16 150W <---> rest of LAN
Site 2: FIOS Internet <---> UXG-Pro <---> USW Flex Mini <---> rest of LAN
    
The Magic Link is working perfectly AFAIK - I can access each site from the other site using 192.168.0.x or 192.168.2.x and internet traffic flows as expected.
However, when I try to set up a Wireguard VPN server and create credentials for a client, while it shows as connected with an IP address of 192.168.5.x, it cannot reach anything on the LAN or the internet. This is true for both of the Wireguard servers I have set up (1 at each site). Similar thing happens with a Teleport VPN.
I'm wondering if there's a problem with running both a (Wireguard or Teleport) VPN server and a Magic Link point-to-point VPN simultaneously? Or maybe there's a problem with VPNs in general on my platform? (running EA software at site 1 btw, general releases at site 2)
thanks in advance for any ideas/suggestions!
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Magic Link?
Sorry, s/b Site Magic
i have the same setup, two houses linked with "Site Magic" (with 6 VLANS) and one site running my Wireguard Server for remote access and OpenVPN to overseas VPN server. No issues access both sites and internet. Can you ping your router when connected? Maybe DNS? Under my setup i have manual, set DNS servers to my piHoles.
Following for a solution as well. I think teleport is just LAN access and Server is just WAN access while away from the network.
