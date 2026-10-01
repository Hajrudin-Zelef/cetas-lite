---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-1dqslxz-ubiquity-please-provide-proper-documentation-for-dc3508d8-2
title: "r-ubiquiti-comments-1dqslxz-ubiquity-please-provide-proper-documentation-for-dc3508d8"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer"]
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-1dqslxz-ubiquity-please-provide-proper-documentation-for-dc3508d8.md
source_anchor: ""
source_lines: [131, 196]
sha256: fe49a5a79885112337adb9704a2d8a42c20b8458294be6c41966ee1d163681e5
---

# r-ubiquiti-comments-1dqslxz-ubiquity-please-provide-proper-documentation-for-dc3508d8

Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
Ok, I'll bite -- what exactly do you find difficult about setting up FW rules?
I've set up a couple of VLANs -- for IoT and cameras -- with rules to isolate them from everything but a single machine on the main VLAN, rules to allow access to the devices from the main VLAN (but not the other way), have some port forwards from the Internet to a couple machines, etc. Took a few minutes to set up all of it coming at it from basically nowhere and it just works.
Thanks for your feedback. We have over 1100 upvotes on our Help Article here: https://help.ui.com/hc/en-us/articles/115003173168-Deep-Dive-into-Advanced-Firewall-Rules but understand we can always improve. We'd like to connect you with our leadership team so we can learn more about your experience. We reached out via Reddit Chat to properly connect. Cheers
I strongly agree with this! I'm a Network Engineer and have worked with both Cisco and Palo Alto firewalls for decades. The firewall management Unifi has come up with makes absolutely no sense and is downright difficult to create firewall rules.
Yeah, I mean, documentation is definitely notUbiquiti's strong suit anyway, but they love to reinvent the wheel in non-standard and unintuitive ways.
roof exultant trees literate shelter like rob future wise pie
This post was mass deleted and anonymized with Redact
100% agree. It's almost as if they tried reinventing firewall rule logic, yet it's the same just harder to make sense of.
It's like taking a regular jigsaw puzzle and repainting the pieces in confusing patterns, making it much harder to figure out how they fit together.
Okay I’m glad it’s not just me, I’ve done PA and Fortigate, I’ve somehow never really messed with Cisco and somehow I still wind up confused in unifi firewalls
same here. making rules on a PA is absolute cake compared to whatever the hell UBNT was thinking with their… system
Documentation? Ha
We are talking about a company which support was a community forum for years (and is for most)
It's confusing, it's basically 3 firewalls types. Internet in is from wan to lan. Internet out is from lan to wan. Internet local is everything going to the CPU of the router/firewall.
Funny part is there is no cpu out. It's just another lan client in that respect. Consider it a management interface as a lan client.
Unifi and edgerouter are really weird. It was originally based on VyttaOS.
I prefer playing with Palo Alto or pfsense, though it's been a while since I last played with pfsense.
I've gone full circle with Ubiquiti after 12 years of using them... Started out as a young tech in my first MSP job just using their WAPs then as my career moved on began installing their switches, routers/"dream machines", etc., but honestly with their lack of standardized documentation and laggy support, clients just get Palos/Sonicwalls/Pfsenses now and I only use them for their WAPs, Air Fiber devices, and cameras. I'd like to see their routing/firewall offerings completely revamped but I'm not holding my breath.
That all said, most rules, port forwarding, etc. is pretty simple to do especially if you let an NGINX box do the heavy lifting. But what perplexes me to this day is how bloated/slow/heavy Unifi's OS is, yet at the same time still lacks a plethora of core features that every other firewall out there ships with by default, it's beyond frustrating. I'd rather deal with the quirks of navigating some portions of PAN-OS than Unifi any day.
The tips are pretty much useless as well
u/ubiquiti-inc You should take notice of this entire post.
I design enterprise networks across Palo, Cisco, Juniper, F5.... But took me fkn days to figure out why basic things weren't working. The interface for firewalling in Unifi is the worst experience I have ever had.
The documentation is terrible.
The logging for unifi is a joke (and wtf is it different between USG and UDM...)
Not being able to mix objects and literals sucks
Not being able to refer to host objects (ie: having to assign static address to each device then creating a hostgroup to refer to it) sucks
Traffic Rules are basically useless and keep changing between versions... in non useful ways.
Lack of custom SSL is a joke (and the instructions don't seem to work with current versions...).
With a very small amount of development effort Ubiquiti could make the experience fantastic, it only requires a small amount of effort to make the automation of iptables rules into SDN, instead they redesigned the UI which added very little value.
Some advice that is hopefully helpful that I personally couldn't find written anywhere.
The key to understand is that UDM is simply orchestrating iptables rules and it doesn't naturally do reciprocal rules; ie: it's functionally a stateless firewall... in 2024.
So if you have two 'isolated' networks (which should be the normal state of affair if you're purchasing prosumer+ level kit such as UDM/UXG pro), then you need two rules to open that traffic.
So for example you have a PC lan and Server lan (offering T/80).
Rule#1
Chain: LAN IN
Prot: TCP
Src: network:pc_lan
SrcPort: any
Dst: host group:server_ip
DstPort: port_group 80
Advanced: No
Rule#2
Chain LAN IN
Prot: TCP
Src: host group:server_ip
SrcPort: 80
Dst: network:pc_lan
Advanced: yes (established, related)
Default: LAN IN deny network:pc_lan -> network:server_lan
Default: LAN IN deny network:server_lan-> network:pc_lan
Default: LAN OUT allow (all networks)
This is why I’ve had a pfSense router and an aging CloudKeyGen2+ controller for 5 years despite the “progress” UI have made with management and firewall hardware. Router configuration is bare-bones and unintuitive.
The big thing to realize is that the in/out stuff is almost irrelevant. You place all your inter-VLAN rules in LAN IN. This really just causes it to be in the list in a certain order. Also, I didn't realize this was necessary until last night, but you really should have the following LAN IN rules which are not there by default: At the top:
Allow related sessions.
Drop invalid sessions.
At the bottom just before predefined rules:
Deny all ... But the deny all is not so straightforward. I created IP groups for all my networks by adding them in CIDR notation. You cannot just use ANY/ANY, as this will block internet traffic. You would have to build other rules to allow internet otherwise.
Also make sure your networks are not set to isolated. The firewall doesn't seem to be able to override this. Oh, and get used to adjusting the advanced firewall view to show ALL. There are two toggles to adjust and it's not intuitive.
Once you have the 3 LAN IN rules built, you can start placing your LAN IN ALLOW rules between them and it will work like any other firewall.
I had been fighting with this for almost 2 years and just realized you needed these entries last night. Now all my rules work fine. Ubiquiti just needs to add a few more rules by default.when you create a network. Perhaps a toggle to do deny all if no matching rule. I can see why a lowly consumer might be confused by that.
