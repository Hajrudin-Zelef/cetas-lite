---
id: collect-260926-mikrotik/mikrotik/questions-1021136-how-to-isolate-networks-with-a-mikrotik-router-b27a79d0
title: "questions-1021136-how-to-isolate-networks-with-a-mikrotik-router-b27a79d0"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-1021136-how-to-isolate-networks-with-a-mikrotik-router-b27a79d0.md
source_anchor: ""
source_lines: [1, 11]
sha256: 8f2742b02713d2b9dabe6a9a7a2f2f774ee1a0e9d8d82bd351ce778cd14e5ca8
---

# questions-1021136-how-to-isolate-networks-with-a-mikrotik-router-b27a79d0

I recently got a Mikrotik router for my network, and I want to create 3 networks that are isolated from each other but all having internet access:
- The "main" network for PCs, etc.
- A network for home automation devices/appliances. I do not want these hosts to be able to access the other networks, but I want some specific hosts on the main network to be able to access specific hosts on this network.
- A guest network for visitors. I want hosts on this network to only have internet access, and be completely isolated from the other networks.
I've been able to setup these three networks using bridges by following these instructions and also mimicking the default configuration that came with the router.
It sound like I now need to define firewall rules to block the traffic between the bridges, and it's here where I need a little help. My understanding is that the Mikrotik firewall software is based on Linux iptables.
- There seems like there's two places to do this: the main firewall configuration in /ip firewall filter , and a bridge-specific section in/interface bridge filter .  Which one would be best to use?  What are the pros and cons of each?
- I'm experimenting with the bridge filters, but all my rules have a little traffic bar icon next to them, which doesn't look good to me. I can't find any explanation of what the icon means.
- How should I setup the rules? Would it be more manageable create a bunch of separate chains for each bridge? If so, how should the chains be organized?
- It sounds like I need to define forward rules for this.  Are there anyinput oroutput rules that I would need as well?
- I should have the rules match on the bridges/interfaces (i.e. in-bridge, out-bridge, WAN interface, etc.), correct? E.g. to block packets from the main network to the home automation network, I would need a rules that's something like in-bridge=main out-bridge=home_automation action=DROP, correct?
