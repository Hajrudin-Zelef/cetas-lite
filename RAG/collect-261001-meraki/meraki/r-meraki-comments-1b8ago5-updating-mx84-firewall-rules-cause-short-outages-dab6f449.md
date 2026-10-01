---
id: collect-261001-meraki/meraki/r-meraki-comments-1b8ago5-updating-mx84-firewall-rules-cause-short-outages-dab6f449
title: "r-meraki-comments-1b8ago5-updating-mx84-firewall-rules-cause-short-outages-dab6f449"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-1b8ago5-updating-mx84-firewall-rules-cause-short-outages-dab6f449.md
source_anchor: ""
source_lines: [1, 14]
sha256: 714c71adc4deb323562150d6ce2bd620bc956f4a081be40afc19da356fb590d9
---

# r-meraki-comments-1b8ago5-updating-mx84-firewall-rules-cause-short-outages-dab6f449

Updating MX84 firewall rules cause short outages 
        
    Just checking if this is normal behavior. we have 2 HA MX 84's FW 18.107.2 that do all routing with about 60 L3 firewall rules. whenever I update, change or add a rule all inner vlan traffic stops for a short time and the phone calls start rolling in. I imagine this is the firewall re loading the rules but I would think this would be designed to be near instant so people can update rules at any time of the day, it is annoying to need to wait for after hours to make firewall changes.
Section des commentaires
From every firmware release notes as a known issue:
'After making some configuration changes on MX84 appliances, a brief period of packet loss may occur. This will affect all MX84 appliances on all MX firmware versions'
Thanks!
I should have RTFM before posting
(Also annoying that this is not being fixed at all by Cisco, I would bet this will never be fixed if it has been in just about every release notes.)
Fwiw mx84 is getting old in the tooth I replaced all mine this past few months with mx95
Yeah, I should do the same soon, the thing chokes on our DVR streams from the camera VLAN, our firewall is at 50% utilization at least most of the time. I think I will put it on 2025 capital plan
Also consider to check out the Meraki MVs, really good cameras without DVR!
where do i see the firewall utilization?
This has been an issue since the beginning. Any change to firewalls/vlans/Vpn settings will drop all traffic
