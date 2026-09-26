---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-tvdv25-guide-how-to-set-up-wireguard-clients-with-vpn-4946b665-4
title: "r-mikrotik-comments-tvdv25-guide-how-to-set-up-wireguard-clients-with-vpn-4946b665"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/forum/wireguard/r-mikrotik-comments-tvdv25-guide-how-to-set-up-wireguard-clients-with-vpn-4946b665.md
source_anchor: ""
source_lines: [132, 176]
sha256: 7fe3dbc0e3889d347317e4a07d0eb475c55a7a8a2e9a02bc82a2a93322d76117
---

# r-mikrotik-comments-tvdv25-guide-how-to-set-up-wireguard-clients-with-vpn-4946b665

Hello Thanks for your answer, yes i want this "Country A -> VPN-A, Country B -> VPN-B,…, specific sites -> No VPN [like connecting remotely to work], everything else -> VPN-Z", i try usign Case D with B. but i failed, because my D sites are block (UK, FR sites).
Thank you for this great tutorial.
When trying to follow it, I noticed one thing that maybe you could clarify:
The file I got, just like yours but from another VPN provider, had as you said:
Address = [IPaddress]/32
However when I tried to do
/ip address add address=[IPaddress – tunnel DE]/32 interface=KeepSolidVPN-Germany network=[IPaddress – tunnel DE]
with my VPN's parameters, I noticed that MikroTik ROS does not accept /32 subnets.
How did you work around this problem?
Please double-check. ROS accepts IP address with /32 subnet but you can't put subnet in network
Line should look like this: /ip address add address=10.200.14.35/32 interface=vpn-interface network=10.200.14.35
10.200.14.35 is bogus address
True, that goes through.
But when I open the GUI Adress List, there is no /32 any more. If I use /31 or /30, then these /31 or /30 stay after the IP address.
That just may be a quirk.
Did you use /32 in your real world usage and does it work? My address list entry is red right now.
Thanks for this! i got a tunnel working to the outside. traffic goes out no problem. using scenario A and my specific VPN provider.
Only issue I'm running into is port forwarding from the VPN provider back the client using scenario A. I'm a novice with MT so any pointers would be great.
I will say only this - I don’t know.
1 - I don’t have MT anymore
2 - as I am thinking how I would do this it would require:
a) that you have your own IP address from VPN provider. Most of them would put you behind NAT, therefore inbound traffic would be restricted. Some providers are allowing you to have open port for you - this is where you would need to funnel your traffic b) opening port in the VPN interface to allow traffic to go it. I assume similarly to opening port to allow inbound WG connections if you have road warrior setup.
Thats my $0.02 :) Best
Sorry for resurrecting this thread, but was curious if I'm reading correctly that wireguard support in RouterOS wasn't added until v7.xx?
Correct
Appreciate the response. I was able to update the router to 7.12.1 from 6.48. I was worried it wouldn't be available, but all went smoothly. :)
this is more insightful than the mikrotik wiki. thanks, I'm sure this will come in handy somewhere in the future.
great job, thanks mate! I was having troubles configuring WireGuard client on my MikroTik using YouTube videos and MikroTik wiki, but your solution finally helped me to sort it out :)
My configuration is:
- private VPS with custom WireGuard setup inside Docker Container
- MikroTik at home
Initially your setup didn't work for me as well, but then I noticed you used /32 netmask, so I changed it to /24 and it worked like a charm.
Thank you so much! This really helped! However, I have run into a small issue. I am able to route all network traffic through my wireguard server on the internet (scenario B), but response times are rather poor and sometimes the connection fails. I suspect that my routes are not configured properly.
IP wg interface: 10.11.10.2
IP main gateway: 10.11.10.1
DHCP address space: 10.110.10.50-100
DNS: set via DHCP to 8.8.8.8
Is it a problem to have the main gateway, the dhcp address range and the wg interface on the same subnet?
Thanks!
Peter
I was able to follow this and do it in winbox. Super quick to do. Thanks for this!
Thank you so much for this guide. I’ve read so many forum posts that are very similar to this but not right. This article was just what I needed to tweak my current Wireguard setup and get it working. Thank you.
Hey, could someone make a tutorial video and upload it to youtube?
That's a great guide, simple and most importantly works like a charm! Thanks a bunch mate.
Cheers for this. Worked great!
