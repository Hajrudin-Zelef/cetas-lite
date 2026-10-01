---
id: collect-261001-cisco/cisco/r-networking-comments-1b78su-nexus-vpc-config-sync-69870b5f-2
title: "r-networking-comments-1b78su-nexus-vpc-config-sync-69870b5f"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "research"]
source: docs/RAG/collect-261001-cisco/r-networking-comments-1b78su-nexus-vpc-config-sync-69870b5f.md
source_anchor: ""
source_lines: [115, 167]
sha256: 59c8c35b3f1476bb9df35266783d6eeabb215aff90ef8b7c15db6e2cc9689525
---

# r-networking-comments-1b78su-nexus-vpc-config-sync-69870b5f

Section des commentaires
I am NOT using config sync. It is not mandatory and making changes does not cause any issues. I have a friend who tried it out, and it said it was unreliable. It would get stuck sometimes and he'd have to manually go in and make sure it replicated, so he stopped using it and manually makes changes on both sides. He had another reason as well, but I can't remember what it was.
Well I'm intrigued.
Are you using virtual port channels? Also I've been using it for 8 months and find it reliable, as long as you understand what can by sync'd and what can't, and mutual exclusions and yada yada.
Yes. To FEXs and servers later on.
Ohh I think that might have been the other stipulation. It doesn't sync every interface command....only certain one that need to be consistent. So you have to make changes to both 5Ks even if you do us config sync. That sounds familiar, so maybe that's what he was complaining about.
I use vPC only and it never creates an issue regardless of what I do. If I add a VLAN to a 5k, it just goes to vPC inconsistent on that VLAN and it doesn't add it to the peer link or downstream links until I add it to the other 5K as well. There's no impact to the other VLANs or anything like that.
Using FEX, there is only problems when you dual attach the FEX to both the 5Ks (which isn't really the smartest FEX topology). The FEX ports need to be identical on both 5Ks for them to function properly. This also isn't an issue unless you make switchport changes to live servers. And I can't see why you'd do that unless that server was a virtual host (or something like that). And if it was, then you shouldn't be using a dual-attached FEX...you should have dual-attached hosts with single attached FEX.
While I'm sure config sync does work fine for many, I'm just passing on the stories I've heard. I've personally had many issues with Nexus gear being FAR buggier than old school IOS, so I prefer to manually make my changes as well.
That's what I was kind of worried about
I found it to be extremely buggy. The commands were missing half the time on the other switch.
Did you go back to manual sync? If so, any issues with that? beside the obvious human factor errors.
I just removed the config sync commands and just entered the commands twice
Exactly. It's not a feature I trust in a production DC. Do the configs manually.
If you're using virtual port channels or FEX it's my understanding you're going to have to use config-sync.
If you don't, your vPCs will come down no matter how hard you try to manually sync future changes. Which I'm guessing would be bad for your production environment (assuming you're not labbing some 5Ks... :)
Config-sync is sort of weird at first, but it's not that bad. You'll get used to it.
Thanks! It is for a per-production environment but you know how it is: LAB CAN'T GO DOWN!!!
Everywhere, right? "THIS IS DEV! WE NEED 99.999% up time!" Hahahaha.
Also, I'm hoping the experience DC guys will chime in on this one. I've only setup a couple of pairs of 5Ks and have used conf-sync since the beginning so I don't have personal experience of what happens if you don't.
Maybe you can try! :D
With graceful consistency checking (where supported), this is not always the case.
Completely untrue. Research vPC (type 1 and type 2) inconsistencies. You have to try pretty hard to bring down your vPCs with config changes.
Maybe it depends on your definition of bringing down a vPC? Because Cisco seems to say a switch is bringing down a vPC utnil the inconsistency is cleared, but you are passing traffic.
At this point one could argue that you're OK because you're passing traffic, while another could argue that you've technically brought down a vPC on one switch, and that to minizmize disruption maybe config-sync is a good idea to be safe.
Maybe I'm not fully understanding you or the technology (and that wouldn't be the firs time? :D)?
Source.
Commentaire supprimé par le membre
A "show switch-profile status" can help with this. I saw a lot of mutual exclusion errors when I first started using sync because I was so used to just manually configuring everything, and then tried to sync stuff.
Hi, I actually just set up an entire sand-box/lab network using 7k's down to 5k's using VPC's
For the sake of understanding VPC's entirely I'd bang it out manually.
I mean sure we all know the theory, and probably can set up VPC's no sweat, but how often to you set this up? Once a year? I mean unless you're working for a consulting firm, etc.
Everytime I do it, I look through my notes and forget a few things, it's like almost learning it over again.
Plus you feel like it's badass in the end. You know, setting up a highly reliable, highly redundant DC backbone with a shit ton of bandwidth.
As far as fex's go. I don't really know the "name" of the setup I use. But essentailly I'll have two FEX's. One fex goes straight to 5k-A, the other 5k-b, then whatever servers need to go there on the fek's split their nics between FEX a and B.
Any servers/chasis capable of supporting VPC's I just link straight into the 5k's.
If you have any questions, I can possibly answer this for you, because you can run into serious shit if you make a dumb change.
Best of luck
So it seems like you got this? and I got that except it's just a single 10gig , not a PO. The second option allows me to take a 5k down w/o breaking all the FEXs that are attached to it.
Negative. The fex on the left would only have uplinks to the 5k on the left.
Then the fex on the right would only have one uplinke to the other.
No VPC's down to the fex level.
Any server using copper, would have NIC 1 on FEX 1 NIC 2 on FEX2.
Then they'd be load balanced in a variety of ways.
little late to the party, but tossing my .02 in here.
Ran Vpc+ in a production datacenter. Turned off config sync due to bugs. Our DC only had an approved IOS Matrix for our tech stack of 5.X, so we just duplicated on both sides. Although, Nexus can take global commands from an interface config. The ? mark is not your buddy on NX-OS.
I caused a type-1 inconsistency which shutdown a 5020 out of a pair. Thankfully we weren't using any 2232 single-leg fex's on that switch pair.
I got reamed for it... only ran nexus for 2 weeks before that.
left the company shortly after they found their corp firewalls were misconfigured.
What was mismatched if you don't mind me asking ?
You definitely DO NOT need to use config-sync if using vPC or EvPC. I tried config-sync in the beginning and it was horrible. I avoid it completely now.
Thanks! This seems to be the theme here.
I've encountered major issues when trying to use config sync after having applied substantial switch configuration prior to turning it on. With a clean config? Not too bad. But I was burned badly enough in the first case, that I'm not a huge fan. Nice for cases with no graceful consistency checking (like Enhanced vPC dual-homed FEX topologies) -- but when you have config-sync spit a bunch of non-ASCII garbage into your terminal (as I've seen in the former case), you can be understandably leery even in the latter case.
