---
id: collect-260926-mikrotik/mikrotik/questions-6120-mikrotik-vs-cisco-for-wifi-bridge-application-a502e188
title: "questions-6120-mikrotik-vs-cisco-for-wifi-bridge-application-a502e188"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2009-05-21", "2009-05-23"]
keywords: ["throughput"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-6120-mikrotik-vs-cisco-for-wifi-bridge-application-a502e188.md
source_anchor: ""
source_lines: [1, 26]
sha256: 5b5ce83217f40ec32e60bc6ccbe463d3aa448e8f3eef39887deb844cb65fd105
---

# questions-6120-mikrotik-vs-cisco-for-wifi-bridge-application-a502e188

A local vendor is proposing a wireless bridge solution using MikroTik. I had never heard of MikroTik before but it seems they have a following. My usual first choice is Cisco, however the price difference between the two is significant > 8X. Am I crazy to be considering MikroTik? Anyone ever heard of it? How well do they compare? MikroTik RB433 with wireless card vs Aironet 1410
- 
        Out of curiosity why is youre first choice usually Cisco?sclarson– sclarson2009-05-21 20:00:46 +00:00Commented May 21, 2009 at 20:00
- 
        Familiarity, good quality, good technical resources...James Moore– James Moore2009-05-23 03:53:34 +00:00Commented May 23, 2009 at 3:53
8 Answers 8
I had the same issue trying to decide what to use, as was with you cisco was my first choice but then I saw the licensing and infrastruture costs so I went with a combination of Mikrotik and Ubiquity equipment, I have had great success runnning both, they are nice and customizable (although I will admit the mikrotik cli is a bit annoying at times), the Ubiquity kit is a web driven config but it is very featurefull without any of the fluff and stuff you dont really need, runs on AirOS (linux based as with the mikrotik gear) and can successfully run a 11g link of over 50Km LOS, though I have heard tell of Mikrotik gear going up to 70Km LOS with a parabolic antennae. All in all I'd say your pretty dang safe with either Mikrotik or Ubiquity gear.
You should really also consider looking at vendors like Aruba and Motorola when looking into wireless/wireless bridges.
I'm the head network admin of a small company (3000~ users) in Australia that has been using Mikrotik routers (in place of cisco gear) for about 4 years now.
The mikrotik stuff is fully 802.11x compatible but they do (similarly to any major player) provide proprietary extentions that are only compatible between mikrotik devices.
One of the best examples of this is the nstreme protocol they use which introduces a wireless polling mechanism to allow best throughput with p2mp sites where sites are at varying distances and bundles packets to allow better throughput.
One of my 5.8ghz p2p links running nstreme in Turbo-A mode. img37[dot]imageshack[dot]us/img37/5276/9544426.jpg
As mentioned above, whenever you use such a system you need to ensure you have a good way to monitor and check any issues with it, mikrotik does provide such an inteface in the form of a mangement application named winbox. (or alternatively you can use the dude for en-masse management)
Old question but I'll add my 2c worth for people wondering about Microtik in the future.
I used to work for an ISP that rolled out Microtik access points across a city with both Ubiquiti and Mikrotik hana clients. The performance was fantastic. No, it's not a silly idea to go with Microtik over cisco. In fact, It may be better than cisco. You're buying a name brand with cisco, although they generally are very good. Cisco's licencing policy has never been particular user friendly.
Originally we used another brand of gear (forgot which) but they were very expensive. Performance was a huge problem with them and we couldn't get good support from the company supplying them to the market or the manufacturer. Not have those problems with Mikrotik.
I can highly recommend those two brands to you.
There are many lesser-known vendors who make long-distance wireless bridge solutions, and many of them have been around for a long time. Because one vendor is providing both ends of the bridge, interoperability isn't as much of a concern as it is for a traditional 802.11 wireless LAN.
I am not sure it's still the case, but many vendors used to do things that aren't necessarily Wi-Fi- or 802.11-compliant to make long-distance links work better (like extending certain timeout thresholds or not acknowledging certain packets). The main downside I can see to this is that it'll be more difficult to use standard 802.11 packet analyzers to troubleshoot problems. But it looks like this MikroTik router uses vanilla 802.11abg anyway.
If you are already using Cisco wireless devices throughout your network, and you have a network management system that supports them, then maybe there's some benefit to sticking with what you know. The bridge devices either need to have a good management interface built-in or you have to have good external monitoring software so that you can tell what's going on when the link misbehaves or goes down.
And how could you not like a company whose software is called "The Dude"?
Mikrotik is not bad, it's very popular in Europe. The only one critical thing it miss, is a 802.3af compatibility.
Mikrotik RouterOS (The nix running inside the routerboard) is highly malleable. Easy to do L7 filters, Mangling, Per Connection Classification and that kind of stuff. Next, the forums are basically useful. Basic networking skills required though!
vincent@itworx.co.ke
Kenya.
I haven't kept up with them, but I was pretty happy with the MikroTik products I used for a large wireless deployment a few years ago. Their support was also good and inexpensive (like $200/month, I think).
