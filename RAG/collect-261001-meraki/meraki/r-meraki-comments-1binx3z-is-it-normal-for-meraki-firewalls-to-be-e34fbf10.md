---
id: collect-261001-meraki/meraki/r-meraki-comments-1binx3z-is-it-normal-for-meraki-firewalls-to-be-e34fbf10
title: "r-meraki-comments-1binx3z-is-it-normal-for-meraki-firewalls-to-be-e34fbf10"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["training"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-1binx3z-is-it-normal-for-meraki-firewalls-to-be-e34fbf10.md
source_anchor: ""
source_lines: [1, 31]
sha256: e924822f1a1e411097c7f986fbede698a5921ba835e844514a5321e90ae1fdd7
---

# r-meraki-comments-1binx3z-is-it-normal-for-meraki-firewalls-to-be-e34fbf10

Is it normal for Meraki firewalls to be configured with an explicit Allow Any/Any? The folks with support that I talked to said this was normal and I still can't wrap my head around it. How are your MX's configured? 
        
    I stepped into this position after the dude that configured the network left.
When I started reviewing the firewall, it started with a long list of obvious allows. There were no blocks in place. Then I got to the end of the MX outbound rules list.. and see a default Allow Any protocol from Any source via Any port to Any destination.
I'm entirely confused. What is the prototypical Meraki MX config look like? Are there a compensating controls that I just haven't discovered yet?
Section des commentaires
It’s a stateful firewall that blocks all incoming traffic and allows all outbound traffic by default. You can create rules as needed, including a deny all directly above it if you want to granularly control a whitelist
I think this might be the nugget I've been digging for. I was not aware of this. Thanks!
The way I think about it is if the traffic originates FROM the network, out to the outside world, by default Meraki will allow it out. If the traffic originated from OUTSIDE the network coming in to the network, the Meraki will block it unless explicitly allowed (Port Forward, 1:1 NAT, etc.).
If you don't mind, I'd love to ask you a couple of questions. I've encountered a couple of people in your place, and I haven't understood why they felt the same way you did, "why is this allowing anything to talk to everything?" I'm not trying to be a jerk or anything, I want to understand your perspective better and thus understand folks that see it this way better, as well.
https://www.dropbox.com/scl/fi/vvtsm0aomw5mnoqalgwdy/chrome_2024-03-19_20-48-47.png?rlkey=buhubpi35erwcc7hh9ckra7pj&dl=0
That's the default rules. Is it just that you're used to a different layout and didn't see that the top is inbound and default DENY or that the ALLOY any/any is under outbound? Or is there something in the way they write it that makes it look like the opposite?
Truly I'm NOT trying to be a jerk, I've just been using this for ten years and so I don't even know what it looks like to a new person anymore! I really just want to hear your perspective, why you felt it was allowing anything to start. I want to be able to explain it better to folks in the future. thanks in advance, and glad you found ForWorkin's explanation useful!
It ends with an explicit permit any/any yeah.
Why you can't change it to a deny without adding a rule in front of it is beyond me, but that's how it is.
I agree this is dumb, part of the philosophy of Meraki is to make things simple and plug and play, baking that in makes it harder for nontechnical users to fuck up their network. As a result individuals with moderate to professional level of technical knowledge get annoyed
You can just add a deny all above it
That is literally what I said, yes
I agree with people's comments on the basis of stateful firewall allowing return and why would you worry about inside going out ...but are the days of vlan isolation gone too, as this is also absent given the any/any permit.
I make it a habit of putting rules to stop every vlan from access internal IP address classes to effectively create vlan isolation. It's a pain but easy enough. Is this not other people's concern too?
Are you worried the outbound firewall rule has an any/any?
I’ve seen this where a vendor says “you must allow this traffic”, and a less experienced tech pops in those allows “as requested”.
Training moment.
This was probably the most confusing part, like why tf are there 47 individual allow rules, only for the 48th rule to be the allow any/any. Wouldn't that make the other rules redundant?
Yep. Probably not a big deal with modern processing, but ACL’s are checked with every packet too until they hit a match. So for efficiency alone you wouldn’t want redundant rules.
I believe it's order of operations, so no. If it failed at a rule above the 48th, it would stop there.
Outbound or inbound? Outbound, no worries. Inbound, lots of worries.
On outbound yes that’s default, on inbound it’s default deny any any
This is normal for MX Appliances. You create your rules over top of that one.
Most of the work I do on Meraki is blocking traffic between vlans, setting up content filters, security filtering, umbrella connection, qos policies, site to site vpn, SD wan settings, setting up static ip's through the dhcp. But all of that is typically done in less than an hour and before I even plug in the device the dashboard is typically already setup to accept it so that at most I might need to input the static ip from the carrier on site. Pretty much full network online inside of an hour.
Outbound only bro
