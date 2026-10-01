---
id: collect-261001-general-networking/general-networking/r-networking-comments-11l9eyq-hsrp-issue-6caf880c
title: "r-networking-comments-11l9eyq-hsrp-issue-6caf880c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution", "ethernet"]
source: docs/RAG/collect-261001-general-networking/r-networking-comments-11l9eyq-hsrp-issue-6caf880c.md
source_anchor: ""
source_lines: [1, 49]
sha256: f09510137ab9ef01218c16485b6a5416d64a550b533ee97d682db39d66d4de60
---

# r-networking-comments-11l9eyq-hsrp-issue-6caf880c

HSRP issue 
        
        
        
    
    
    Hello,
Recently I've configured HSRP on our core switches. Because of the physical limitations of our fiber connections, we only have one trunk interface as the heartbeat interface between the two core switches. I set Core Sw1 as primary for all VLANs and Core Sw2 as the secondary for all VLANs. The configurations were successful and when I pull the ethernet cable going to our ASA, the traffic still flows through Core Sw2. However, when I unplug Core Sw1, we lose connection on the entire network.
Quickly explaining our Core to ASA connection: the connection from Core Sw1 and Core Sw2 is a single cable connection from each of the Core Switch. We only have one ASA running on the network at the moment.
What could be causing this problem? I brought up setting up STP root manually so that the VLANs have a cost to follow to the Distribution then to Core Switches, but my Supervisor is insisting that doesn't matter. If I can get a clear technical explanation as to why our failover isn't happening when Core Sw1 loses power it will be very helpful.
Section des commentaires
A topology drawing would provide a better description. Where does the ASA connect? You say 'when I unplug Core Sw1' - unplug the power? Or unplug it from Sw2?
You'd need to check the state of several things when simulating the failure. Is STP showing Core2 as root correctly? Is Core2 showing as the active gateway for HSRP in all your VLANs? What does the routing table look like pre/post failure?
ASA gi1/5 > Core SW1 ten1/0/32
ASA gi1/7 > Core Sw2 ten1/0/32
-these two interfaces are configured as members in interface Redundant1.
I unplug ASA's gi1/5 and gi1/7 to test the fail over when it comes to interface failure.
Thank you. I will start with that. We have static routes and that may have been the issue? will need to check the downstream switches and their routes to see if they are configured correctly. Thank you.
I see from your past posts you have Cat9500 switches.
If you have models that support it, I’d highly recommend StackWise-Virtual (think VSS but improved) over HSRP. Easier configs and the ISSU upgrades for .3 releases (so eg 17.3 to 17.6 or .9) are a breeze.
I know some say VSS had bugs but those were largely ironed out way back in the early versions of 12. on the older 6500’s.
Yeah after reading I saw that they made the fixes. But the time when I was researching on how to make this redundancy work, I read that there are bugs and decided to drop it. That's why the HSRP. Perhaps after this has settled and other big projects are out of the way I can revisit this and configure it.
They “fixed it” in the release of 12.2(55) around 2011…
VSS and the Stackwise Virtual have been rock solid since then. The one caveat against this is people not reading the update release notes properly and trying to ISSU update to something that isn’t a .3 (or multiple of 3) release or to a new major version 16.12 to 17.something.
This
Does both the core switches have the same routes to the destination you are testing to? And is your ASA to configuered to use the VIP? Is there a also a trunk between the two core switches or how are they configuered?
Yes to all of them. The VIP address is the next hop created on the ASA.
Are the next hops for your PC’s on the switch side also the VIP address?
Probably need much more info to diagnose this one, but my hunch would be spanning-tree. If you kill sw01, traffic should start flowing to sw02 from the downstream switch/device. If not, then it could be because the link is blocked by STP.
That was my initial hunch too because of the way our network is configured at the moment. I tried to bring it up to my supervisor before I implemented this but he was quite adamant about STP at default settings (only having 'spanning-tree mode pvst') should let the network work. I didn't have the technical knowledge at the meeting but I knew I had done this before and remembered it from both labs and my previous work. Just couldn't explain why. Since then I looked into it more and see why setting manual STP root is important. I will bring it up to him in near future.
How are you setting root? Which switch is root in this state? Fwiw, if you don't set root manually switches run an election which sometimes can have wildly unexpected results with switches getting root that have nothing to do with the pair you're working on. I've seen some wildly stupid shit happen with spanning tree.
If you have 2 cores, and your running hsrp on them both, you pretty much imo, should be setting the root/failover manually if those vlans are supposed to reside on those specific switches.
This is where it gets tricky. I suggested to my supervisor that we need to manually set a root to guide the VLANs in case of a failure on a device. Because his idea is that because STP 'does it on it's own' we don't need to manually set the root. But that could make an access switch a root where the traffic doesn't go anywhere. I'm working up a strategy on how to bring this up professionally lol.
So setting the roots is the command:
spanning-tree mode pvst
spanning-tree vlan 5 priority (lower for primary)
spanning-tree vlan 10 secondary (higher for secondary)
would that be the configuration?
What IP does the ASA send traffic to on the switches for the LAN? HSRP virtual IP?
Yes the HSRP virtual IP.
What switches are you using and what code?
If you only have one ASA what is it plugged into?
Do you have your harp priority set and the default gateway for each vlan set as the VIP?
Do the commands:
Sh ip int br
And
Show standby
Give you the results you expect? Show standby should confirm that the vlans are up and active and that the correct HSRP priority is being observed.
It’s been a long time since I used HSRP - VSS (stackwise virtual now) really is the way to go if your 9500’s support it (not all do.)
