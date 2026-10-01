---
id: collect-261001-cisco/cisco/r-cisco-comments-jwdock-bgp-troubleshooting-questions-e095c7ae
title: "r-cisco-comments-jwdock-bgp-troubleshooting-questions-e095c7ae"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/r-cisco-comments-jwdock-bgp-troubleshooting-questions-e095c7ae.md
source_anchor: ""
source_lines: [1, 49]
sha256: 40303e474187e8f61c47bf92eaf0dc0e842bd0d4489515e3ce538a96cffbfee1
---

# r-cisco-comments-jwdock-bgp-troubleshooting-questions-e095c7ae

BGP troubleshooting Questions 
        
        
        
    
    
    I have 2 internet connections served by 2 routers on 2 different providers. They share one AS#, and that routes a /22 we own, which runs on a subint that flips between the 2 routers via HSRP. Normally, if one router goes down, the traffic flips to the other router because the HSRP gateway moves with it. I've verified this works by a sho standby, and by pinging the vip to verify it's moving. I've also tried to just power off the active router, and verified that the vip is still reachable. This worked as of 11/16, when I ran the site on the backup for 24 hours.
Last night the main line went down, internally all my traffic was still hitting the vip on the backup router, but nothing was moving. It's almost like the internet did not know my traffic moved from one provider to the other.
Some traffic always routes out the backup, because I'm using prepending to make the main provider the primary, but traffic local to the backup vendor will always route through them regardless. When I force a failover by shutting the primary int, everything all of a sudden comes to a stop.
All the usuall sho ip bgp commands seem to look ok.
No changes were made on the routers between when it worked on Monday and now, but now it's broken. Any ideas on how I troubleshoot this?
Section des commentaires
How is it going to know? You need to draw up a high level diagram firstly both ways. You only have FHRP available. What mechanism are you using to ensure traffic is preferred or split coming into you on your PI Space? You probably ended up blackholing your return traffic somehow. Is there a link between both routers?
It knows where the gateway is, because my firewall cluster routes directly to the vip, and the routers vip/gateway sits on a L2 VLAN provided by a set of switches between the fw layer and the router layer. The external IP of the firewall cluster sits on the public /22 also with the router vip.
No matter which router owns the vip, it's reachable by the firewall cluster. Even now, if I power off R1, I can still reach the vip which has failed over to R2. And as I said, this worked on Monday, now not so much.
I need to test from external resources some more, but as of now, I can't bring down the network to do this. I won't be able to test this probably until my maintenance window opens on Sunday.
Look on a looking glass to see where your pi space is being advertised to also just to make sure nothing has broken provider side.
Am I assuming this correctly?
R1/R2 both have a different ISP link and address on a physical interface
R1/R2 both advertise your public subnet via eBGP over their respective ISP interfaces
R1/R2 both have an "inward-facing" interface which is configured with an address in that public subnet
The inward-facing interfaces also share an HSRP address.
Outgoing internal traffic is routed to the HSRP address over an L2 fabric
If so:
How does the HSRP status know to respond to changes in the ISP interface status?
When the "main line went down", what exactly happened? Did the BGP session break? Did traffic stop moving? Did the link status change?
Why would it? Your internal gateway moving between HSRP devices has nothing to do with how BGP is advertised. Assuming a healthy iBGP relationship, either router should be able to route traffic out either ISP.
The one failure mode we've seen with this type of setup is if the ISP interface drops BGP or stops moving traffic without the interface actually going down. You need a trigger in both cases to feed back and cause a traffic move event. Incoming traffic doesn't
Correct assumptions.
Q1: It doesn't normally.
Q2: The interface to the first provider1 never went down, traffic just stopped. Link state was 'up'.
In this situation, I have an SLA running on R1 that triggers if an upstream IP in Provider1 fails to respond after a certain time frame, it then deletes the BGP nei on R1, once all those routes disappear the routes on R2 take over.
I also attempted to manually shut gi0/0/0 which is connected to the provider. No effect.
On R2, I could ping to opendns from the provider2 interface, but not from the /22 we own. It almost seems like P2 isn't taking our advertisements or something. I have a ticket open with them.
If the R1 link stopped moving traffic but kept the BGP session / link alive, then it's likely they didn't stop advertising your space either. This would be my first guess and is something we've seen happen. I'm not aware of any real solutions here--we failover to a different DC.
Either way, sounds like an upstream issue.
You have mentioned that you're doing path prepend. Have you checked if you're missing an empty route-map? I had a customer who has the same issue and he was just missing an empty route-map. Or you can have your ISP add a static route on their end but I wouldn't recommend that.
Also, make sure you've next-hop-self for your iBGP peers. These are the two I can think of on top of my head.
Do you run iBGP between your 2 routers? Do you have default timers on BGP sessions?
BGP is slow to converge on the internet, so some outage conditions could lead to up to 2-3 minutes of perceived downtime.
iBGP between your routers should guarantee more stable and reliable internet facing convergence/routing.
A diagram could help. From what I understand, your setup should/could be highly tolerant to any fault.
I do.
I have a loopback (lo0) configured with an ip on both routers. From Either router, a show ip bgp sum shows the other router as a Neighbor. Each router has the other's lo0 int configured in it's router bgp profile.
I want to post a diagram, I know that would be helpful, but we block all file sharing and imaging sites. Reddit is pretty much text only for me.
Do you also run next-hop-self between those iBGP peers ?
I would pair your FHRP with IP SLA or similar to influence reachability against the HSRP primary.
If 1.1.1.1 not reachable, increase AS prepend over link and failover.
It's difficult to circumvent ISP routing issues that don't update when you advertise updates, other than advertising more specific routes out of your other ISP.
