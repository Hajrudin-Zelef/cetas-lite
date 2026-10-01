---
id: collect-261001-cisco/cisco/r-cisco-comments-ox3r3b-differences-between-etherchannel-on-lacp-and-pagp-feb44d77
title: "r-cisco-comments-ox3r3b-differences-between-etherchannel-on-lacp-and-pagp-feb44d77"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/r-cisco-comments-ox3r3b-differences-between-etherchannel-on-lacp-and-pagp-feb44d77.md
source_anchor: ""
source_lines: [1, 40]
sha256: 6431cac618a2e8b83ca14311be2298b406becb3d1a16e899a64a8f7baeae9bf0
---

# r-cisco-comments-ox3r3b-differences-between-etherchannel-on-lacp-and-pagp-feb44d77

Differences between etherchannel "on", LACP, and PAGP? 
        
    About a week ago I was creating an etherchannel on a couple of ports on a C9500-48y4c using "channel-group 1 mode on" command on a couple of 10Gig trunk ports and doing the same to the Catalyst 3850 on the other end of the connection. For whatever reason I had a really hard time getting the etherchannel to work. Either the portchannel interface would just err-disable itself or I would get a message to the effect of "cannot put channel-group on this interface" on the ports themselves. Eventually I just used PAGP Auto/Desirable which is working now but I wonder what is the major difference between the 3 methods of enabling an etherchannel?
It's odd to me that I had issues with these two ports on the 9500 (Ports 1/0/47&48) while on ports 1/0/1 & 2 I had configured those for a different etherchannel using "mode on" and didn't have an issue with them.
Section des commentaires
"On" is static; no negotiation occurs between devices.
You explained this so much better in 3 sentences than what I've read on several full page explanations I've been seeing. TheCCNA cert guide 
Cisco with CCNA likes to drone on about the theory of things. The entire CCNA is some weird whimsical experience about the thought process behind everything.
Also, once the ether channel is formed there is no difference. It’s all about negotiating whether to form an ether channel.
Commentaire supprimé par le membre
That gives me warm fuzzies. We have 9 etherchannels with 7 of them being set to "on"...
Ya, I did this once. LACP always ever since.
have you even gone as far as to even go look more alike?
I always do this:
Then do everything else via the Port-Channel. Works every time.
This is basically what I had to do with my troublesome ports. Had to just default both of them before they would let me configure what I needed to.
Be careful with mode "on". If the link is up on one side of the channel and the other side is down (due to a patch panel or something in the path) then the UP side could hash a flow over its UP link even though the remote side of that link is down, causing any traffic on that link to blackhole.
IMO you should never use mode "on" and always use LACP active/active on both sides.
Oh good to know, I think I'll start changing these as I work later evenings.
What's the cleanest way I can switch from the etherchannel configuration to LACP? My thought so far is to shut off one pair of ports, correct their configuration and then no shut them and then do the other pair or will that not jive well?
LACP = L Actice C Passive (Industry Standard)
PAgP = Cisco Auto Desires
Different protocols, same goal.
Some older firmware builds of the 37xx and 38xx series had serious problems with LACP (link flapping etc.) but not in that way that you describe.
On is necessary if you want to use a port channel with a VMWare server, and you don't have the licence that allows you to use LACP.
I feel like if you're trying to use a port channel to a VM host, you're likely doing your VM host networking wrong in the first place.
I'm sure there is exceptions to this (otherwise such a thing wouldn't be possible), but, YOLO.
Cisco does suggest to use 'active/passive' if going the LACP config route. I've never actually tried to use 'on' in a port channel config. Good to know.
On: Static etherchannel, no negotiation or real protocol involved.
LACP: Open standard protocol, has two sub-options:
active: This device will attempt to negotiate a protocol, and will respond to requests made to it.
passive: This device will not attempt to negotiate a protocol, but will respond to requests made to it.
PAGP: Proprietary Protocol, just don't. There's very few reasons to do so (I believe the biggest being that it supports more standby interfaces than LACP, but I could be mistaken -- it's been a long time).
For obvious reasons, for any given pair of interfaces between two devices, at least one interface in the pair must be active. Typically people spray "active" anywhere. There are a couple small concerns with doing so but it's typical practice.
You can have:
And it'll work.
This won't work:
But this will:
Commentaire supprimé par le membre
Don’t forget WLC
