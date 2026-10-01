---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1c5e4xi-ipsec-up-but-no-traffic-passing-927c9e54
title: "r-fortinet-comments-1c5e4xi-ipsec-up-but-no-traffic-passing-927c9e54"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1c5e4xi-ipsec-up-but-no-traffic-passing-927c9e54.md
source_anchor: ""
source_lines: [1, 32]
sha256: 49005851a865363623f317940f0d3afb8067d857579f3357310f06744cc96b1b
---

# r-fortinet-comments-1c5e4xi-ipsec-up-but-no-traffic-passing-927c9e54

IPsec up but no traffic passing 
        
        
        
    
    
    Hi all,
We have a setup with a Fortigate 60F (7.2.8) with a fortiextender in WAN port. Normal internet connection is working fine. But we have some trouble with IPsec VPN.
The tunnels is up both Phase 1 and Phase 2.
When i try to ping from Local lan to remote lan i can see in dianostics that the packets leave the firewall, but it is not received on the other end. Same happens when i try the other way arround.
I have built 100's of tunnels, but this is the first setup with Fortiextender. Am i missing something? :)
Thanks!
Section des commentaires
On your IPsec tunnel network settings, do you have nat traversal enabled? Generally it’s disabled but in this case with the extender in the middle you may need to enable it. If phase1 and phase2 came up but traffic is not getting to the other site in my experience nat traversal is the culprit.
You are a hero, i forgot everything about NAT traversal. Everything working now.
That single setting has caused me hours of troubleshooting before. Glad to hear it resolved the issue for you.
FortiExtender doesn't matter. The basics of IPsec troubleshooting apply:
Is the traffic allowed?
Is the traffic routed correctly?
Is the traffic allowed in the phase 2?
Hi HappyVlane,
Yes to all 3, i think i would see it under dianostics if one of those was the problem? :)
It can't be yes to all three, because you say the other side doesn't receive anything.
Got the same problem recently, raised the ticket with Fortinet and was told that it's a bug. To solve the issue is to disable npu offloading under phase 1.
I have had this exact bug with an 1800F running 6.2.x firmware but have been informed it was fixed in 6.4.11 so may not be the same thing but worth a shot if all else fails
Do a debug flow on both sides to be sure.
We had this in the past a few times between two Fortis (tunnel was up, but no traffic passing through). After checking the SAs we sometimes simply had to flush / reset the tunnel, after that the problem was gone. https://community.fortinet.com/t5/FortiGate/Technical-Tip-How-to-flush-a-VPN-tunnel/ta-p/196631
This is from my notes and always helped me during a troubleshooting session in regards to that:
Maybe that helps.
Also, for some reason I think I've read something about a known issue in 7.2.8 - you would have to look that up.
Never mind regarding the "known issue". I just double-checked and it was this:
IPsec phase 2 negotiation fails with failed to create dialup instance, error 22 error message.
