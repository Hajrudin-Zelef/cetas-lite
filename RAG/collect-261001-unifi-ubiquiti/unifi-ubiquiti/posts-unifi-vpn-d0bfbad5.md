---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/posts-unifi-vpn-d0bfbad5
title: "Device: REDACTED"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["SpaceX"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/posts-unifi-vpn-d0bfbad5.md
source_anchor: ""
source_lines: [1, 40]
sha256: 671736caf9fc937e568e65a04828a2232f253779136dbacb7364d79c0c3e5511
---

# Device: REDACTED

Starlink: installation, first impressions, and running it through UniFi
This weekend’s side quest: installing Starlink as my second internet provider. …
This is a quick post showing how to set up a Wireguard VPN in an UniFi Dream Machine.
So, as you may have heard, 𝕏 is supposed to be blocked in Brazil in the coming hours, and using VPNs to bypass that is supposed to generate a 50k BRL (~8k USD) fine to citizens.
This post just shows how to set up a VPN inside your network, and how to route specific sites through it. It is meant to show you how to do it, so you don’t do it by accident, so you don’t get fined.
For this example, we’ll use Mullvad as the VPN server.
You can create an account there, pay, and then download the Wireguard configuration file for a specific location. It would look like this:
[Interface]
# Device: REDACTED
PrivateKey = REDACTED
Address = 1.2.3.4/32,aaaa:aaaa:aaaa:aaaa::1:8829/128
DNS = 1.2.3.4
[Peer]
PublicKey = REDACTED
AllowedIPs = 0.0.0.0/0,::0/0
Endpoint = 1.2.3.4:51820
In your UDM console, go to VPN, then VPN Client, and then Create New.
There, you can import the configuration file exported by Mullvad.
In my case, it complained about the format of the Address property, so I
removed the IPv6 address from it, and imported it again.
You can then save it, and it should connect:
From there, you’ll probably want to route specific traffic through it, let say, some website called x.com.
You can do that by going to Routing, and then Create Entry.
There, you can give it a name, choose the source devices (or choose all), choose the destination domain (or IP, region), and the interface to use. Here’s an example:
With that done, you should be able to access the websites through the VPN (but don’t do it, its illegal)!
Note that if you were to commit this crime, you would also have to add other
X-owned domains, like pbs.twimg.com, twitter.com, etc.
UniFi has its “Teleport” feature, which allows to easily connect to the network. But, those won’t follow the routing rules we set before, so it won’t work.
But it seems that static routes do work!
You can go to Routing, then Static Routes, and then Create Entry.
There, you give it a name, set distance 1, set destination network to
104.244.40.0/21 (in our example of routing 𝕏), and next hop to the VPN IP
(it will be the “Tunnel IP” in the VPN Client screen):
That /21 CIDR is more or less what is covered
here. If you want to be more precise, you can
create more and more specific routes, but this seems to work well enough.
If you’re not in Brazil, and is not following the news, 𝕏 is blocked in Brazil since this weekend.
Some providers seem to have just removed it from their DNS, some did nothing, Starlink announced that the judge can kiss their ass, but some providers did properly block the 𝕏’s IPs.
Nevertheless, most of the people with 2 or more healthy brain cells looked at all this and we were like:
Most people are using VPNs, and/or saying their “friends”/representatives are posting from another country, or that they are traveling (I’m in Paraguay, btw), and that includes news outlets, politicians, senators, and so on - and that’s just who is being loud about it, I guess a lot of people are in read-only mode for now.
