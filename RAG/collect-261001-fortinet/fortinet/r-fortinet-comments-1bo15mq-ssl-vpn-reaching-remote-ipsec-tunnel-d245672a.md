---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1bo15mq-ssl-vpn-reaching-remote-ipsec-tunnel-d245672a
title: "r-fortinet-comments-1bo15mq-ssl-vpn-reaching-remote-ipsec-tunnel-d245672a"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1bo15mq-ssl-vpn-reaching-remote-ipsec-tunnel-d245672a.md
source_anchor: ""
source_lines: [1, 23]
sha256: cced2a377adabdebbfe584fffa102f49c55adf3025de0079e59a1621657652fe
---

# r-fortinet-comments-1bo15mq-ssl-vpn-reaching-remote-ipsec-tunnel-d245672a

SSL VPN reaching remote IPSec Tunnel. 
        
    I have been banging my head against the wall for several days now trying to make this work.
I have a Fortinet with a subnet of 192.168.1.0/24 and a remote sonicwall with a subnet of 192.168.2.0.
They have an ipsec tunnel between them and it works great.
I created an ssl vpn / policy on the fortigate and that works fine, users who connect to the SSLVPN fortinet client, are able to reach all the resources on subnet 192.168.1.0/24.
When I start trying to include the subnet 192.168.2.0/24 i the picture it just doesn't work. I have followed several guides including Fortinets on guide on this topic but no luck.
Please help or ask me more questions if that is not enough info.
Section des commentaires
This has been discussed countless times, both in this subreddit and elsewhere. Have you looked?
The bullet-points are:
firewall policies (have firewall policies that allow SSL-VPN -> IPsec direction; both on the FortiGate and the remote end, if applicable)
routes (ensure all routers on the path know that the SSL-VPN subnet is reachable via the FortiGate, through the IPsec tunnel)
IPsec phase2 traffic selectors (if they're restricted, make sure you also have selectors allowing SSL-VPN-subnet -> <whatever required> IPs)
SSL-VPN config (e.g. if you have split-routing enabled, don't forget to add routes to the remote Sonicwall subnets)
This Pretty much sums it all up.
Can you ping the SonicWall remote subnet directly from the fortigate?
yes, it has an ipsec tunnel. I tried everyting includin the steps outlined in the first post, but never succeeded. I am wondering if it is a compatiblity issue between fortigate and sonicwall
You need a firewall rule from SSL VPN interface to IPSec interface. Maybe a clone reverse.
On the Sonicwall, you need a Static Route of the SSL VPN Subnet pointing to the IPSec tunnel.
You'll need a Sonicwall firewall rule from IPSec to LAN ensuring you include the SSL VPN Subnet in source.
Willing to bet you either didn’t set the policies OR you didn’t share the SSL subnet over the IPsec tunnel.
Sonicwall wouldn’t know how to get to the SSL VPN subnet and vice versa if the routes aren’t being set
