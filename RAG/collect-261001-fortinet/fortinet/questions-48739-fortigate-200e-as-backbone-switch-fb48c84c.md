---
id: collect-261001-fortinet/fortinet/questions-48739-fortigate-200e-as-backbone-switch-fb48c84c
title: "questions-48739-fortigate-200e-as-backbone-switch-fb48c84c"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2018-03-01"]
keywords: ["throughput"]
source: docs/RAG/collect-261001-fortinet/questions-48739-fortigate-200e-as-backbone-switch-fb48c84c.md
source_anchor: ""
source_lines: [1, 19]
sha256: c6f51e46d970b336dc68c79fe803bc09d4fcd467e943d18b415d9a18ed8f01a5
---

# questions-48739-fortigate-200e-as-backbone-switch-fb48c84c

As part of the reorganization of the network of my company, I intend to apply the fortilink architecture, so I want to know if the fortigate 200e can also play the role of the backbone switch of access switches of fotinet (424d and switches 448d)
2 Answers 2
The cited switches do not only have a throughput of 88 Gbps (424D) and 176 Gbps (448D) but also feature 2/4 10GE ports for uplinks.
In comparison, the maximum throughput of a FGT-200E is rated at 20 Gbps, 9 Gbps for small packets (64b). In order to use a Fortigate as a backbone switch it would need to have 10GE ports; aggregating ports in a LACP trunk will be not as efficient and will exhaust the available ports (14 on a FGT-200E).  
The main reason I advise against this deployment pattern is that the main advantage of having a UTM firewall, namely protection via AV, IPS, Application Control etc., will have to be sacrificed for speed.
The FGT is meant to manage the Fortiswitches in your LAN; as such it's very convenient (e.g., VLAN handling), powerful and you can even extend the security perimeter to your access ports.
Just keep in mind that the whole infrastructure will be as powerful as the weakest part, and that would be the FGT if used as a backbone switch. If you use a Fortiswitch for backbone and manage and monitor all switches from the built-in FGT switch controller, all is fine.
- 
        The Fortigate throughput is for layer 3 forwarding/routing. L2 forwarding in a hardware-switch port group is at wire speed.2018-03-01 11:46:41 +00:00Commented Mar 1, 2018 at 11:46
- 
            
            
- 
        @Zac67: correct. As the FGT features 14 GE ports of it's internal switch, and a maximum of 8 ports in a LACP trunk, this would still impose a throughput limit of 14 or 8 Gbps. Depending on the number of access switches the backbone should have more / much more bandwidth than that.user1016274– user10162742018-03-01 11:53:19 +00:00Commented Mar 1, 2018 at 11:53
- 
        Yes - flows are all limited to gigabit speed when the core switch links aren't any faster. LAG can only (potentially) save you from congestion with multiple flows.2018-03-01 12:06:28 +00:00Commented Mar 1, 2018 at 12:06
Looks good:
  FortiGate units can be used to remotely manage FortiSwitch units, which is also known as using a FortiSwitch in FortiLink mode. FortiLink defines the management interface and the remote management protocol between the FortiGate and FortiSwitch.
EDIT after @user1016274's very reasonable comments: Using a switch (the FGT-200E) with only gigabit ports as core may severely limit the overall throughput of your network. Even aggregating multiple GbE ports won't enable you to run multi-gigabit flows across the switch. You should look into options using the FGT as controller only and connecting the faster switches directly.
