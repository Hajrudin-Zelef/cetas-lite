---
id: collect-261001-cisco/cisco/r-networking-comments-hjkgij-catalyst-9600-to-stackwise-virtual-or-not-140bd1d3
title: "r-networking-comments-hjkgij-catalyst-9600-to-stackwise-virtual-or-not-140bd1d3"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/r-networking-comments-hjkgij-catalyst-9600-to-stackwise-virtual-or-not-140bd1d3.md
source_anchor: ""
source_lines: [1, 47]
sha256: a2e0234690f7e9dc0e36f0d273387a3057428b522f0647e402a5d97c73db785b
---

# r-networking-comments-hjkgij-catalyst-9600-to-stackwise-virtual-or-not-140bd1d3

Catalyst 9600 - To Stackwise Virtual or Not 
        
    If you were building out a new core with Catalyst 9600s, could ECMP everything back to the core except for active/standby firewalls (HSRP used), would you implement Stackwise Virtual? I've read some old posts about about Stackwise Virtual, and things have changed since then, so I wanted to see if opinions are any different, or if you have any thoughts that may sway me either direction.
Here's my list of pros, cons, and questions I've come up with for Stackwise Virtual:
Pros
- 
      Single management plane
- 
      Better link utilization, loop-free topology for L2 connectivity
- 
      No FHRP
Cons
- 
      Single control plane = single point of failure, most likely from config mistakes.
- 
      Need to be aware how VSL, DAD, linecard, and sup failures affect the virtual switch. L3 failure scenarios appear straightforward.
Lingering questions
- 
      Is convergence any better in Stackwise Virtual vs L3, when thinking about link, linecard, or sup failures?
- 
      Can ISSU be trusted?
Welcome all thoughts pertaining to Stackwise Virtual on 9600s!
Section des commentaires
ECMP everything with HSRP for the firewall VLAN. VSS works fairly well, but when it breaks it is a huge pain. It isnt worth the failure scenarios, especially if you arent using it to remove spanning tree.
ISSU typically works for small upgrades, but anything major is hit or miss. Then when things break, it is a pain to troubleshoot. You are safer in almost all scenarios by having two redundant boxes with L3 connectivity. I have roughly 20 VSS or stackwise virtual pairs in production (most installed before I was with the company) and anything new, or even replacements, I steer away from VSS/stackwise virtual.
Any specifics you can go into regarding ISSU issues you ran into with major upgrades?
One issue is major release jumps you sometimes have to do two or three ISSU reloads between different codes. If you had two independent devices, you can just reload one with the correct code and not worry about it.
One issue I recall was I did an upgrade, and when the first chassis came back, the VSL links were in some weird state and the switches were both active, not in VSS mode. Keep in mind, I always configure them with at least two VSL and two peer keepalive. After rebooting the second box, when it came back everything was happy.
Another occasion the secondary box just kept rebooting. I think I ended up going to a different code which resolved that. I think there is a bug I found for that one. There have also been multiple other issues people have had that I work with.
There really is no reason to go with stackwise virtual if you are doing all L3. If you are terminating redundant L2 links, it makes more sense to remove spanning tree, but I would steer you towards VPC on Nexus platform, or better yet EVPN/VXLAN.
Not sure what you mean by convergence vs L3. Are you talking about IGP convergence because of a single control plane? If so it's all the same because of NSF.
I've already done 2 ISSU upgrades 16.12.1 -> 16.12.2, 16.12.2 -> 16.12.3a no issues.
Just an FYI. Quad SUPs are still not supported in the 16.12 train running SVL. The other SUP in each chassis sits there like a brick
Yep, that's what I was after regarding convergence!
Yeah, not sold on quad sups yet.
I was doing some failover testing between 2 directly connected stackwise-virtual systems this week that were doing BGP with NSF for both ipv4 and ipv6. During the switchover from active to standby I did not observe a single ping drop on the data plane and only observed 2 ping drops to the control plane of the chassis undergoing the switchover. These were consistent results also.
Thanks for sharing your results! Which platforms were connected?
My setup was host ---- ISR4K ---- 9500-32C ---- 9606R ---- host
I tested constant pings between the 2 hosts while doing the failover on the 9500
We are moving away from VSS in the core and moving to ECMP everywhere. VSS has been ok, but I just like having completely separate devices and control planes. I know that even if an upgrade bricks a core switch, we will be ok as it won't affect the other switch (aside from any weird routing glitches).
I would potentially look at moving away from HSRP and just bring those firewalls into your IGP as well. L3 only in the core and ideally even access nowadays.
Thanks for the input!
We'll have to do an IGP to the firewall soon for another project so there's our opportunity to move away from HSRP.
There are some that don't like running routing protocols on firewalls, but I've never had an issue personally. I think in some cases it depends on which teams manage the firewalls as you might not want security engineers potentially causing a routing problem.
How many Supervisors do you have? 1 or 2 per chassis?
Haven't purchased yet, but was leaning towards 1 sup per chassis. I had to draw out several failure scenarios to understand if there is any benefit for having multiple sups or line cards for the SVL link. After that and reading some more documentation, I'm not convinced quad sups is worth it.
Commentaire supprimé par le membre
