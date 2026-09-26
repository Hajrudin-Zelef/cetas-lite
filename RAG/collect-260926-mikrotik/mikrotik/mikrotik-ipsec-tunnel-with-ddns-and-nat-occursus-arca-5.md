---
id: collect-260926-mikrotik/mikrotik/mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca-5
title: "mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2019-07-17", "2019-07-18", "2019-08-16", "2019-09-26", "2019-09-29", "2019-10-04", "2019-11-10", "2019-11-24", "2019-12-08", "2020-01-05", "2020-01-07", "2020-01-08", "2020-01-29", "2020-08-04", "2020-08-16", "2020-08-23", "2020-08-24", "2020-12-24"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca.md
source_anchor: ""
source_lines: [510, 640]
sha256: 842357e5da8df02f520e545680aa463e71c1861d1630568d2e9b160a0c9be0f7
---

# mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca

## Manuel · July 17, 2019 at 17:18

Hi, and thanks for the fast reply.

I did all you said and also double check de configuration it still doesn’t replace the dummy ip but it resolve the host with the ping i notice in the log that there is nothing related to “changed scheduled script setting”, but when i run the script manually it should change the ip. i have a couple of cuetsions that i’m not sure

should i see the ip changed in the SA Dst. Address field of ipsec/policy tab and the address in ipsec/peers tab when the script runs? or the dummy ip stay there no matter what even when the ip really changes

the schedulers needs to be disabled when are created?

in ip/cloud the update-time must be unchecked?

thanks for everything, i really need this to work and is killing me

## Pessoft · July 18, 2019 at 00:01

Hi Manuel,

IP Cloud update-time decides whether time of your Mikrotik will be updated/synced from Mikrotik Cloud time servers. It is not related to configuration of IPSec tunnel and I have it unchecked as I sync the time using NTP from other sources. Schedulers when created are disabled, because netwatch controls later on when they are enabled and disabled again based on whether the remote Mikrotik is reachable. Dummy IP is there only at the beginning, as soon as the netwatch detects that remote network is not reachable, it enables the scheduler, scheduler executes the script at regular interval, script updates the IP address of IPsec configuration ( peer and SA ) with real IP addresses of remote Mikrotik. If manual start of the script updates the IP addresses in IPSec configuration for you correctly, then there is probably just an issue with your netwatch configuration. Can you check that host parameter of netwatch configuration on each Mikrotik is configured to the IP address of remote Mikrotik ( similar like mentioned in the example within article )?

## Agustin · August 16, 2019 at 17:20

Your post was really helpfull. It helped me to connect 5 different Nat network with each other. Thank you so much!

## Emiel · September 26, 2019 at 13:35

Thanks for this really helpfull post! Just one question, any idea why I’m not able to connect to the remote Mikrotik router using Winbox/via browser?

## Pessoft · September 29, 2019 at 23:55

Hi Emiel, run

`/ip service export verbose`and check for settings ofwwwandwinboxservices. They should have`disabled="no"`( means service is enabled ),`address=""`( means that connection is not restricted by IP range ) and`port="SOME_NUMBER"`( should be 80 for www and 8291 for winbox by default ). If that is ok, check if you have ports 80 for www and 8291 for winbox allowed in firewall input chain from selected IPs/ranges. Be careful that you don’t grant access to these management ports from Internet or other public network.
## Emiel · October 4, 2019 at 12:53

Thanks Pessoft, all settings are correct (by default). Strangely enough I can ping from local router to remote router. But I’m not able to ping from client (connected to locla router) to the remote router. In this case I’m able to connect to the (i.e.) NAS on the remote site and vice versa.

## Joak2602 · November 10, 2019 at 19:51

Hi Pessoft,

First of all, thank you very much for sharing a great tutorial. I have followed your configuration guide and it is working good and I am able to ping the remote route and network from both sides. I want to access internet through the IPSEC tunnel. For example, a device with IP 10.10.20.150 must connect to internet as it will be in 10.10.10.0/24 network. How can I get it working as expected? Thanks again

## Pessoft · December 8, 2019 at 23:29

Hi Joak, thank you for your feedback. Unfortunately, I didn’t yet configure Internet connection via the IPSec tunnel, so I cannot give you exact solution, but you can find a lot of hints on this, including configuration examples, on Mikrotik Forums.

## Kostas · November 24, 2019 at 21:42

Great tutorial, Thank you for sharing!!

Nice to add:

open VPN

## Pessoft · January 8, 2020 at 01:39

Thank you for your feedback. OpenVPN is great VPN technology, but its implementation in RouterOS is missing some security features and therefore I will not consider writing tutorial for it just now. Hopefully it will get improved in near future.

## Harshan · January 5, 2020 at 06:50

Hi,

Can I use a dynamic routing protocol (like OSPF) between the sites? I have more LAN segments in each sites. Usually we add the tunnel IP in OSPF. What is the tunnel IP here?

## Pessoft · January 7, 2020 at 20:41

Hi, this configuration does not use any IPs for tunnel configuration, so I don’t think that it is possible to have OSPF used – although I never tried it. Maybe creating GRE tunnel over IPSec could provide OSPF capability.

## Harshan · January 8, 2020 at 17:25

I have done it with GRE tunnel. But issue is with update script. Could you please help me with that. I have little or no knowledge on scripting side. Could you please email your mail ID?

## Daniel · January 29, 2020 at 17:36

Sir,

What’s your opinion on IPSec IKEv1 with PSK and XAUTH?

I need to connect three remote office branches to a central office, but all the Mikrotik routers will probably have to be NATted behind the fiber router installed by the ISP (I’d like to remove those crappy routers, but their policy is very strict)

IKEv1 with PSK and XAUTH seems to be the easiest and most viable option. Does it have any disadvantages compared to your method?

## Pessoft · January 29, 2020 at 22:33

Hi, method described in my article primarily addresses topics of dynamic IP assignment and ISP router on both sides. Using IKEv1 with PSK and XAUTH should work, assuming IPSec traffic ( or eventually NAT-T ) will not get blocked between your branch and central office. I’d suggest to use IKEv2 as it is more secure and has NAT-T embedded in it. Although IKEv2 does not support XAUTH, so you need to choose a different authentication.

## Rafael · August 4, 2020 at 20:36

Hello, thanks for your settings, they work!

But I have a problem, router1 doesn’t have access to router2’s computers, but router2 accesses router1’s computers, example:

Does not work

10.0.0.84 -> 192.168.1.173:80

It works

192.168.1.122 -> 10.0.0.22:80

Route1 ip 10.0.0.0/24

Route2 ip 192.168.1.0/24

Can you help me?

## Pessoft · August 16, 2020 at 21:35

Hello Rafael, check if VPN traffic reaches the NAT firewall rule or if some other rule is catching the traffic first in NAT table ( in such case reorder NAT rules accordingly ). Also check if you have firewall rule that accepts forwarded traffic of VPN. Both rules are mentioned in my guide, but your results might vary depending on your configuration. Also, just to be sure, check if other service on some other host is also unreachable ( maybe there is local firewall on the machine, … ).

## ismail · August 23, 2020 at 22:32

hi there thanks for all the info.

I would like to ask if you could maybe send me an updated script.

the script to update the peer address works well

but I would like it if the script could also maybe update the local IP/cloud address as well as the peer IP address so that it would be possible to access the VPN from my phone and table.

I hope that you understand what I mean and that you could maybe help me.

## Pessoft · August 24, 2020 at 21:51

Hi ismail, I’m not exactly sure what you need, but VPN tunnel should be accessible by devices connected to the Mikrotik as long as firewall is not blocking them. Also it is quite simple to add to the script setting of additional variables – Mikrotik Wiki can help greatly in this. Otherwise I suppose your request is a bit out of the topic of the article, so feel free to contact me with more details if you need more help.

## Emiel · December 24, 2020 at 18:14

Hi There,

