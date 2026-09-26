---
id: collect-260926-mikrotik/mikrotik/questions-9327-can-a-mikrotik-750gl-have-the-same-ip-as-the-dsl-router-connected-afc70c00
title: "questions-9327-can-a-mikrotik-750gl-have-the-same-ip-as-the-dsl-router-connected-afc70c00"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-9327-can-a-mikrotik-750gl-have-the-same-ip-as-the-dsl-router-connected-afc70c00.md
source_anchor: ""
source_lines: [1, 5]
sha256: 463adea7a4391c2f8d1a735ae367d3c81b3620192a12affd7da426c304a95e2e
---

# questions-9327-can-a-mikrotik-750gl-have-the-same-ip-as-the-dsl-router-connected-afc70c00

On a simple network of about 10 computers setup with static IPs in the 192.168.1.x range, the ISP's DSL router occupies 192.168.1.1. All computers and the router physically connect to a simple network switch.
What I want to achieve is to put a mikrotik in the middle of the DSL and the computers, that is to plug the DSL router in one port of a mikrotik's 750GL and the switch to another port and have mikrotik bridge the traffic while at the same time appear as 192.168.1.1 for the computers in the network. It has to be done without changing the IP of the DSL router or any of the computers.
I know how to do it if the router and the computers are in a different network, but I don't know if it is possible in the same network and IP.
My final goal is to use the mikrotik as a transparent proxy to block access to certain sites. But this will be in a remote location where I won't be able to physically set it up, only prepare and sent the mikrotik. Additionally if something goes wrong at any point, I would like to be able to instruct someone to just take the plugs out to remove the mikrotik and connect them back to the switch in order to return to the previous working configuration.
I'd like to add that I tried to research this but I couldn't come up with the proper keywords or no such answer exists :)
