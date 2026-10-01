---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-xmxla4-teleport-vpn-with-wireguard-client-d6c2a125
title: "r-ubiquiti-comments-xmxla4-teleport-vpn-with-wireguard-client-d6c2a125"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-xmxla4-teleport-vpn-with-wireguard-client-d6c2a125.md
source_anchor: ""
source_lines: [1, 19]
sha256: 3a547c8439a6469fe7231120602a1c2cad4683637cddee7735be76db7dd57fde
---

# r-ubiquiti-comments-xmxla4-teleport-vpn-with-wireguard-client-d6c2a125

Teleport VPN with Wireguard Client? 
        
        
        
    
    
    Is there a way to use Teleport VPN (which as I understand is using wireguard on the backend) with a Wireguard client software? Specifically wireguard for iOS?
Reason: The wireguard iOS client is superior to teleport because it is persistent and auto-connects to vpn the moment you leave predefined SSIDs. So you can set your house SSID to disable the vpn and then enable it for all other SSIDs and cell networks. Also the client allows for split tunneling so you can route DNS and network servers through the VPN tunnel but have your internet traffic go through your primary connection which helps for speed.
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
No but you can manually setup wireguard connections from the commandline.
Without installing a new instance of Wireguard?
Correct. Wireguard is already installed. A couple of commands and some firewall rules are all that is necessary to bring a link up. If you want it to survive a reboot a bit more would be required and if you want it to survive an upgrade install boostchickens on boot stuff. See https://www.reddit.com/r/Ubiquiti/comments/xlex16/how_to_use_wireguard_on_udmp_via_ssh_without/
Not at the moment however rumor is wireguard support is coming soon.
I eventually turned off Teleport and settled on using a wireguard server running on a raspberry pi. I'll probably move it to a docker on my NAS but for now this works fine.
