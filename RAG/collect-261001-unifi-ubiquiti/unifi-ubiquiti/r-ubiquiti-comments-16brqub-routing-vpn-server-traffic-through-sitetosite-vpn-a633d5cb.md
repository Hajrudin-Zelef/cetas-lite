---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-16brqub-routing-vpn-server-traffic-through-sitetosite-vpn-a633d5cb
title: "r-ubiquiti-comments-16brqub-routing-vpn-server-traffic-through-sitetosite-vpn-a633d5cb"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-16brqub-routing-vpn-server-traffic-through-sitetosite-vpn-a633d5cb.md
source_anchor: ""
source_lines: [1, 27]
sha256: b878c128b2407eaf91078e84b9958283b99b6103e11bc7111a126abf77fa791a
---

# r-ubiquiti-comments-16brqub-routing-vpn-server-traffic-through-sitetosite-vpn-a633d5cb

Routing VPN Server Traffic through Site-to-Site VPN 
        
        
        
    
    
    Hi,
I have a UDM Pro running UniFi OS UDM Pro 3.0.20.
I have a site to site vpn configured that passes traffic from our DHCP clients ( 192.168.2.0/24) to a remote site utilizing 192.168.50.0/24. Clients are able to ping and remote desktop into machines from 192.168.2.0/24 with absolutely zero problem.
However the client needs to be able to connect to a WireGuard VPN i've configured (192.168.3.0/24) to get to a server located in the other location on subnet 192.168.50.0/24. Please bare in mind that 192.168.50.0/24 is an external vendor and I have asked them to configure their side of the tunnel to also allow traffic from our 192.168.3.0/24 subnet.
When connected to the WireGuard VPN server we've created in Unifi, we can access devices located on 192.168.2.0/24 absolutely flawlessly. However we cannot pass any traffic over to 192.168.3.0/24.
For example, there's a citrix server located at 192.168.50.96. When attempting to run a Tracert, it shows it completes one hop to 192.168.3.1 and then times out.
I have configured a static route for this as well.
What am I missing here? Am I just stupid?
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
The route looks correct, is the UDM the gateway at 3.1? if so the trace hits the UDM but doesnt go out to the 50 vlan?
in the trace above are you tracing from the 3.x network? if so its hitting 3.1 and doesnt know where to go OR:
Is there a firewall rule in place to discard/drop between vlans that needs to be updated to allow the 3. network?
When running the trace, i am connected via the WireGuard VPN Client, so I'm on the 3.x network.
As far as the firewall, I don't see anything obvious unless I am blind?
https://gyazo.com/1e915a09a95294d0d649ae79d2bc28b4
I dont see a lan in for wireguard, I see Internet Local, Local in unifi is the firewall, not the local network. you may need a lan in for .50.x from wireguard. Im only about 75% confident on that one, you may have something in the rules.
