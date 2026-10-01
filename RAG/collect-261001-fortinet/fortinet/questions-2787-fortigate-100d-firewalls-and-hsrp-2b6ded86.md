---
id: collect-261001-fortinet/fortinet/questions-2787-fortigate-100d-firewalls-and-hsrp-2b6ded86
title: "questions-2787-fortigate-100d-firewalls-and-hsrp-2b6ded86"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "research"]
source: docs/RAG/collect-261001-fortinet/questions-2787-fortigate-100d-firewalls-and-hsrp-2b6ded86.md
source_anchor: ""
source_lines: [1, 21]
sha256: 7e9f78cb0dbeb51c4553eca3514d2a6b3bebcacce390ca057dd089a766bd884d
---

# questions-2787-fortigate-100d-firewalls-and-hsrp-2b6ded86

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
5
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
First, please see the attached network diagram to see what I am trying to do.
On the 2 distribution switches, we run HSRP and have one Fortigate going to SW1 and the other going to SW2.
Then we have vmware esx hosts with connections to each of the fortigates. The fortigates are configured in an active-active ha config as are the vmware vswitches.
Now for some reason, when we disconnect the cable from the monitored wan1 port on either of the fortigates, the ip assigned the a vm on either of the esx hosts fails to be reachable however the ip assigned to the fortigates is reachable just fine. For some reason, traffic isn't being passed through the fortigates.
However, if you look at this diagram, the vm ip is reachable just fine when the devices failover.
I've tried setting different interface tracking options in hsrp and tried using "ip sla" as well to no avail.
Am I missing something somewhere either on the fortigates or on the distro switches?
Let me preface this by saying that I have not used Fortigate, but speaking generally.
You should have two links from each firewall (FW ), one to each switch (SW), just as you have a link from each FW to each server (SRV).
This is what I suspect is happening. Assuming SW1 is the HSRP active interface, it initially receives traffic from SRV2 on the link to SW2 and creates an entry in the MAC and ARP tables. When the link between SW2 and FW2 goes down, SW2 removes the entries for that interface from it's tables, but SW1 doesn't know the link is down and maintains it's entries.
When traffic comes in to SW1 for SRV2, it looks up the ARP/MAC information and sends the traffic to SW2. SW2 doesn't have an entry for SRV2 anymore, and floods it out all ports except the one it received the traffic on (normal switch operations). This results in the traffic never reaching SRV2 as none of the other links on SW2 provide a path to SRV2.
With the second link the traffic between FWs and SWs, the flooded traffic would then be received at FW1 and be able to get to SRV2.
If you have outbound traffic from SRV2 after the link goes down, or you clear the MAC entries on SW1, I suspect that this would work as well.
I haven't set it up before but I think you'll need to either configure 'full-mesh' HA between the 100D's and your two switches. We use stacked 3750's for redundancy with 2x 100D's in HA instead, seems to work well enough.
I would assume on the the FG HA is working fine, but your 'Internal' switch isn't getting the message that the Primary 'internal' FG link is no longer the one it should be talking to.
