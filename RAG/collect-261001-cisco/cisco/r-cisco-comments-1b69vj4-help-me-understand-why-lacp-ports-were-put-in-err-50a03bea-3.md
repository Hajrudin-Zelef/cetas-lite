---
id: collect-261001-cisco/cisco/r-cisco-comments-1b69vj4-help-me-understand-why-lacp-ports-were-put-in-err-50a03bea-3
title: "r-cisco-comments-1b69vj4-help-me-understand-why-lacp-ports-were-put-in-err-50a03bea"
domain: cisco
role: reference
task: reference
actors: []
dates: ["3400-33-32"]
keywords: []
source: docs/RAG/collect-261001-cisco/r-cisco-comments-1b69vj4-help-me-understand-why-lacp-ports-were-put-in-err-50a03bea.md
source_anchor: ""
source_lines: [104, 143]
sha256: 507cc3e2203a904791a74647e79ce073b19b79c4c1dbb713a7b254e85816d274
---

# r-cisco-comments-1b69vj4-help-me-understand-why-lacp-ports-were-put-in-err-50a03bea

You also have to look at the other side of the port channel. An LACP channel will shut down if the other side doesn’t negotiate, by design. So it could easily be something on the other switch (or VM host, or whatever it is).
Is a static port channel an option? 'channel 17 mode on'
With some backup appliances I've where the lacp pdus are sourced from a member interface Mac vs the ether channel vmac. In those cases, if they do maintenance or that particular member port drops, then it will source from another member and then my side err-disables
It's an idea, I haven't used a static channel in 20 years...
Check your spanning tree config on each side too.
Unfortunately I cannot easily check the other side since it's in the hands of a contracted company - I can only make suggestions.
Config problem on the other side. Software defect on the other side causing a behavior that trips up the channel detect. Or fundamental difference in LACP interpretation between the two platforms.
I remember a pair of Cisco 4506s with a channel between them that had a misconfig on one of the four ports. Ran that way for months, perhaps two years. Occasionally would drop packets. Never threw a misconfig error. (May have been mode on though…it’s been a decade and I don’t remember all the details.) The fact that the packet drops were only like once a month was weird though. In theory it should have been all the time (at least for half of the hash).
Yeah that's exactly what I think... the problem is they're trying to throw me under the bug because if enforced gasp segmentation between clients, servers and PLCs...
More than likely it’s either the spanning tree or etherchannel config is mismatched between the two vendors. The problem is most likely being triggered when spanning-tree reconverges on the Hirschmann end due to them switching HA state, or reversing directions on an industrial redundancy ring protocol (MRP).
a) Can you post the Cisco and Hirschmann switchport configs, and also which spanning tree the Hirschmann are running? They need to match/support each vendors end.
b) Confirm the Hirschmann are running the latest recommended stable software version. This is important in my experience for bug fixes on their end.
c) Can you post “show spanning-tree summary” from your Cisco?
d) Try adding “no spanning-tree etherchannel guard misconfig” for a period of time, and see if the problem disappears. If so, then it’s a mismatched configuration issue between the two vendors.
e) Why do you even have a Etherchannel trunk configured? Your post above shows that you only have vlan31 enabled, so using a normal switchport configuration may easily solve the problem. The following works for all our Cisco<>Hirchmann interface links: switchport
switchport access vlan 31
no shutdown
(both vendors are running rapid spanning tree)
f) Running “debug spanning-tree etherchannel” and/or “debug etherchannel event/error/detail” and capturing the problem occurring would be the way to t’shoot the root cause. You could try getting the Hirschmann admin to invoke an HA failover on their end, or invoking a ring protocol reversal to trigger the problem while your Cisco debug is enabled.
g) Hirschmanns typically run rapid spanning tree. You should have no problem participating in the Hirschmann spanning tree instance from the Cisco end, as long as you match their config. You can configure specific per-vlan spanning tree only on those vlans you’re peering on the Hirschmann end (vlan 31). (Also make sure your spanning tree priority is lower than the Hirschmann so you don’t become their root).
Finally, these issues are why you should always opt for converged OT/IT switch infrastructure. Next time have your integrators use your Cisco switches (or separate ones) instead of Hirschmann. Cisco IE 31/32/33/3400 series support all of the industrial protocols really well.
Hey thanks for the great answer.
a) as far as the Ciscos go, it's currently changed to this: interface Port-Channel17 description AUTOMATION switchport access vlan 31 switchport mode access speed 1000 duplex full spanning-tree bpdufilter enable
Unfortunately I don't have direct access to the hirschmanns
b) I can confirm they are definitely NOT running the latest firmware. Industrial equipment types and their refusal to update... they have made a change request to do it, hope it'll help (and that they'll manage before a few months)
c) I'll add it as a reply due to the 1000 character limit. Mostly default (open to recommendations actually - I'm a sysadmin required to also do networks, not really a network engineer)
d) will do
e) we have two stacked switches, it's to provide redundancy on our side. They're checking of that switch can be replaced with stacked switches on theirs as well (don't know if Hirschmann can do that). Rapid PVST on both sides is enabled.
f) I was not aware of what I could do to troubleshoot and so I simply shut / no shut the ports to recover communication. If it ever happens again, now I know, thanks
g) they do and that was the thought behind the initial configuration - but now I figure it would be best to avoid it completely and remove the headache, they use a single vlan and the likelyhood of a loop is near zero
I agree entirely on your end comment, I wish I set up this infrastructure myself, but I inherited it, and the OT side is entirely managed by an external company. I have zero visibility and it's a pain.
Thanks again for your thorough response, I appreciate it very much.
Make sure your trunks are the same vlans, make sure your speed and duplex are the same, make sure port channel configuration has the same speed/settings, and lacp configuration is the same
They're access ports
You'll want to figure out the root cause, and you could also set up error recovery if this is a rare issue. e.g.,
errdisable recovery cause channel-misconfig
errdisable recovery interval 60
yeah setting recovery is a partial solution but I'd really like to solve the actual problem...
I’d recommend setting both the port channel and switchports modes explicitly to access (switchport mode access). By default they’re dynamic auto and will attempt to go to a trunk and with the other end outside of your control (plus the unknown of how the Hirschmann handles this) it may be triggering a mismatch.
Quite possible actually, I'll do that thanks.
