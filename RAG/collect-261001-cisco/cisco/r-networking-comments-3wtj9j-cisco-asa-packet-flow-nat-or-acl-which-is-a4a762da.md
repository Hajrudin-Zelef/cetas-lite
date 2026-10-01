---
id: collect-261001-cisco/cisco/r-networking-comments-3wtj9j-cisco-asa-packet-flow-nat-or-acl-which-is-a4a762da
title: "r-networking-comments-3wtj9j-cisco-asa-packet-flow-nat-or-acl-which-is-a4a762da"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/r-networking-comments-3wtj9j-cisco-asa-packet-flow-nat-or-acl-which-is-a4a762da.md
source_anchor: ""
source_lines: [1, 19]
sha256: 29e4acade25077388542a110396d147b5cfd7386cee080d1ed32e77ad85738c2
---

# r-networking-comments-3wtj9j-cisco-asa-packet-flow-nat-or-acl-which-is-a4a762da

Cisco ASA packet flow -- NAT or ACL: Which is evaluated first? 
        
    Hey all, My new position is requiring me to do more than ASAs than I have previously, and I was wondering what happens first during packet flow on an ASA, NAT or ACLs? I've seen several diagrams after google searching packet flow that say ACL evaluation happens first and if the traffic matches an entry on the ACL then NAT happens, but I watched a video on youtube that says the opposite. Can someone clear this up for me or even suggest a resource that explains ASA packet flow clearly for me?
Thanks a million.
Section des commentaires
This depends on what version the ASA is at. If the ASA is at 8.3 or newer, then NAT will happen before ACL evaluation. If the ASA is pre-8.3, then ACL evaluation will happen first.
Thanks.
I thought un-NAT happened before ACL and NAT happened after? Could be wrong...
I think you are right about that, actually. Usually the context I am thinking about is outside->inside traffic, which would be un-NAT'd before the ACL. inside->outside traffic would be checked against the ACL first, then NAT'd. This is probably the easiest way to keep things relatively straight.
That said, things could probably get a bit hairy if your dealing with both ingress and egress ACLs, and twice NAT. As others have stated (to OP), using the packet-tracer command is a great way to test these sorts of things. My biggest gripe is that it doesn't work well for traffic coming in over a VPN, but other than that it is handy.
A good way to visualize the flow is to use the packet-tracer command. It will step you through the order, which as u/internet_eq_epic mentions can change between firmware versions.
packet tracer command? You mean the Program cisco packet tracer?
ASAs have a built in command/utility called packet-tracer where you can specify traffic (ingress interface, protocol, source/destination IP, source/destination port, etc.) and the utility will show you in detail how the ASA processes the packet, and the result of each step. It is very powerful for troubleshooting, and even for planning and other such things. I have used it in so many different scenarios, I feel like my whole life would be different without it. If you will be working on ASAs, I would suggest getting intimately familiar with packet-tracer.
No, the ASA has a built in command called packet-tracer.
There's a GUI for it in ASDM and you can use it from the command line. Given an example protocol, input interface, source and destination IPs and ports it will show you step by step what the ASA would do and a final result of would the packet be permitted or not.
packet-tracer input (nameif) (tcp/udp) (source IP) (source port) (destination ip) (destination port)
Use this if you every have any question if the firewall is dropping traffic. It will show you where it gets dropped. Sometimes it can be buggy but for the most part it works pretty well.
http://www.nycnetworkers.com/wp-content/uploads/2013/12/asa_orderofops.bmp
That helped me.
