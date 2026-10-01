---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-tt2itx-easy-way-to-identify-which-policy-is-blocking-67d0f554
title: "r-fortinet-comments-tt2itx-easy-way-to-identify-which-policy-is-blocking-67d0f554"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-tt2itx-easy-way-to-identify-which-policy-is-blocking-67d0f554.md
source_anchor: ""
source_lines: [1, 14]
sha256: e0f23f613cb37c3090a13e6e5c1ef5a7d125ef96fa4129b1a12368373a6d0ad1
---

# r-fortinet-comments-tt2itx-easy-way-to-identify-which-policy-is-blocking-67d0f554

Easy way to identify which policy is blocking? 
        
    I've inherited a mess of a firewall. There are multiple policy rules setup (some without names) and I'm trying to identify which policy is causing traffic not to route between our SSL VPN IP pool and one of our internal VLANs. Is there an easy way to setup a process so I can try to see which policy is causing the block? I'm thinking if I log into the SSL VPN and ping the internal VLAN, is there a debug command or log that will show me which policy is preventing traffic to route between the two? Thank you for your help!
Section des commentaires
Go to your policy set and enable logging on all rules.
Then go to the Forward Traffic Logs and apply filters as needed.
Depending on how much traffic you receive, you might not want to log everything though if you don't have a FortiAnalyzer.
Thank you!
There is a "policy lookup" feature on the firewall policies screen that lets you put in some details like src/dst ip and the zones and it will tell you what policy it will hit. It can be tricky if you have other security profiles and you need to know a little about the design ... like the traffic flow and what zones it's hitting.
Thank you!
Run a debug flow..
https://community.fortinet.com/t5/FortiGate/Troubleshooting-Tip-First-steps-to-troubleshoot-connectivity/ta-p/192560
Use policy ID or enable it so you can see it in yourpolicy set. In logs you'll see what policy id being blocked. If its unnamed or not doesn't matter, policy id will be unique and you should be capable of finding the policy
Check the hit count for the policys
