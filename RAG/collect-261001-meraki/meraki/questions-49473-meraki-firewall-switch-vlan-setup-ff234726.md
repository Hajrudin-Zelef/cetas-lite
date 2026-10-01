---
id: collect-261001-meraki/meraki/questions-49473-meraki-firewall-switch-vlan-setup-ff234726
title: "questions-49473-meraki-firewall-switch-vlan-setup-ff234726"
domain: meraki
role: reference
task: reference
actors: []
dates: ["2018-03-28", "2018-04-01"]
keywords: ["ethernet"]
source: docs/RAG/collect-261001-meraki/questions-49473-meraki-firewall-switch-vlan-setup-ff234726.md
source_anchor: ""
source_lines: [1, 19]
sha256: 9597dbd128cc84da3c17f271bb3d93e7bb5edea3256d6196a1bf483acb93e82c
---

# questions-49473-meraki-firewall-switch-vlan-setup-ff234726

I have a Meraki MX84 Firewall and a Meraki MS225-48FP switch. I have two VLAN's one is VLAN 1 (Default) the other is VLAN 50. Vlan 1 (Default) has a scope of 192.168.1.x and VLAN 50 has a scope of 192.168.10.X. I am trying to access clients on VLAN 50 from VLAN 1 the default. Specifically I also have a Dell VRTX and the iDRAC is on VLAN 50, I wanted it that way, so VLAN 50 is my Management VLAN. How am I able to access and ping the iDRAC that is on VLAN 50 from VLAN 1????
- 
        You need to edit your question to include the network device configurations.Ron Maupin– Ron Maupin ♦2018-03-28 00:01:37 +00:00Commented Mar 28, 2018 at 0:01
- 
            
            
- 
        1The switch trunks both vlans to the firewall. The firewall then handles inter-vlan routing.Ricky– Ricky2018-03-28 01:13:30 +00:00Commented Mar 28, 2018 at 1:13
- 
        Did any answer help you? If so, you should accept the answer so that the question doesn't keep popping up forever, looking for an answer. Alternatively, you could provide and accept your own answer.Ron Maupin– Ron Maupin ♦2018-04-01 23:28:55 +00:00Commented Apr 1, 2018 at 23:28
2 Answers 2
VLANs separate traffic into distinct L2 segments or broadcast domains - no L2 (Ethernet) traffic crosses between different VLANs.
For them to communicate, you need a router to enable L3 (IP) traffic. The router needs to be connected to all VLANs, either by separate ports or (more desirable) a VLAN trunk. The VLAN trunk ports need to be configured to tag frames accordingly, exactly the same on both sides.
If the inter-VLAN router isn't the default gateway for a connected VLAN the hosts require a route to each 'remote' VLAN or alternatively, the default gateway is required to have a route to each VLAN.
you need to configure your gateways on the Meraki Firewall ( via dashboard Security Appliance => Addressing & Vlans )
after that you need to configure the interconnexion between the firewall and the switch by adding all the necessary Vlan on both side
at the end you need to configure the access vlan part on the switch
and it will be ok
br
