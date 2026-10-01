---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-59403-question-about-ubiquiti-edgemax-switch-syntax-and-vlans-2963c908
title: "questions-59403-question-about-ubiquiti-edgemax-switch-syntax-and-vlans-2963c908"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-59403-question-about-ubiquiti-edgemax-switch-syntax-and-vlans-2963c908.md
source_anchor: ""
source_lines: [1, 2]
sha256: c4bfd013f58b51431b3f20235198dea10bf5671fae3bb3c55e32a35017ac5c33
---

# questions-59403-question-about-ubiquiti-edgemax-switch-syntax-and-vlans-2963c908

I've been beating myself over the head with a keyboard trying to figure this out. I'm used to switchport syntax on Adtran and Cisco switches, but I've got this situation with a Ubiquiti EdgeSwitch 48 and a Mikrotik router that confuses me. What I want to do is have 4 VLANs, 4 ports are on the router, so for port config, just have whatever ports be access and have one trunk port for the uplink to the router for the internet. Basically, it's like dividing the ports on the switch into 4 parts and having separate networks for each part of the switch; literally the definition of a VLAN.
However, these (horrible, I might add) switches do not have trunk ports or switchport syntax. They have tagging and untagged ports. I read elsewhere that a tagged port is a trunk port in switchport syntax, but it doesn't seem to be working that way in my testing as I can still ping between 2 VLANs. Can someone explain what I can do if I want port 2 to be an uplink and port 3 to be a device on that VLAN?
