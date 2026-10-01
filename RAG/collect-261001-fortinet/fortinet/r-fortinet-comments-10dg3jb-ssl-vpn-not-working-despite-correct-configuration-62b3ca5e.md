---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-10dg3jb-ssl-vpn-not-working-despite-correct-configuration-62b3ca5e
title: "r-fortinet-comments-10dg3jb-ssl-vpn-not-working-despite-correct-configuration-62b3ca5e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-10dg3jb-ssl-vpn-not-working-despite-correct-configuration-62b3ca5e.md
source_anchor: ""
source_lines: [1, 20]
sha256: 39a660d07be38fc020f77877aba98ded342a8056d6cf06314853e504ff3727c8
---

# r-fortinet-comments-10dg3jb-ssl-vpn-not-working-despite-correct-configuration-62b3ca5e

SSL VPN not working despite correct configuration 
        
        
        
    
    
    This problem has been going on for a while now and I'm starting to run out of troubleshooting ideas to try and find what's the issue.
On my company's FortiGate, we have a VPN tunnel established with another company on a certain interface. The VPN tunnel status shows that it's up, and the configuration is correct from both sides. The other side tries pinging one of our whitelisted IP addresses, it goes through the tunnel but it re-routes the ping to our mail server for some reason. When we try to ping the other side, the ping goes through the default gateway instead of going through the VPN tunnel. Could someone please help me figure out this problem? Thanks in advance!
Section des commentaires
If it was the correct configuration it'd work : )
You don't have blackhole routes setup to prevent internal traffic from hitting the "default gateway", which I assume are your WAN lines. Setup blackhole routes first, it's best practice for both routing and security. https://community.fortinet.com/t5/FortiGate/Technical-Note-Use-of-Black-hole-route-in-site-to-site-IPsec-VPN/ta-p/192526
Also, is this policy based VPN or route based?
Then check if they're trying to reach you by IP or DNS record. If DNS check and see why it's pointing to the mail server (do you have a VIP, a dns translation, how are they configured on the far side? Do they have a hostfile setting perhaps instead of using proper dns?). You need to figure out what is causing them to hit the mail server and if I were a betting man I'd go with either VIP or DNS.
sounds like a routing issue.
are the routes from the output of , get router info routing-table all , for source / destination defined as they should?
w/o a topology and some sanitized config output we cannot help you with your issue.
Also OP, this is an IPSec VPN tunnel, not an SSL VPN tunnel. Big difference. You definitely have a routing issue. As others have said, we need your routing table. Sounds like you don’t have a route for the other side of the tunnel.
tunnel
route
policy
