---
id: collect-261001-cisco/cisco/r-cisco-comments-an40lk-asa-not-using-acl-b4a5fc2f
title: "r-cisco-comments-an40lk-asa-not-using-acl-b4a5fc2f"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/r-cisco-comments-an40lk-asa-not-using-acl-b4a5fc2f.md
source_anchor: ""
source_lines: [1, 65]
sha256: dbc3ccd495676c369b1424ee977ded674c1194be68b15d48aa2b255762e88025
---

# r-cisco-comments-an40lk-asa-not-using-acl-b4a5fc2f

ASA not using ACL? 
        
        
        
    
    
    I'm clearly not setting something correctly.
Outbound traffic works perfectly fine. Inbound traffic is hitting a wall. When I run a packet-trace for udp dns it sits at:
      packet-trace in outside udp 4.4.4.4 12345 my.public.ip.address domain detail
Phase: 2
Type: ACCESS-LIST
Subtype:
Result: DROP
Config:
Implicit Rule
Additional Information:
 Forward Flow based lookup yields rule:
 in  id=0x2aaad73a9cb0, priority=111, domain=permit, deny=true
        hits=129783, user_data=0x0, cs_id=0x0, flags=0x4000, protocol=0
        src ip/id=0.0.0.0, mask=0.0.0.0, port=0, tag=any
        dst ip/id=0.0.0.0, mask=0.0.0.0, port=0, tag=any, dscp=0x0
        input_ifc=outside, output_ifc=outside
Result:
output-interface: outside
output-status: up
output-line-status: up
Action: drop
This is weird, because after it didn't initially work I added this ACL entry:
! the access-group was unchanged, including it for the sake of completeness
access-group outside-In in interface outside
!
access-list outside-In extended permit udp any any eq domain
!
!this object-group isn't being used here, it was originally used in the original ACL entry
object-group network DNS1-int-ip
network-object host 10.0.0.1
exit
!
object-group service DNS1-ext-ports-udp udp
port-object eq 53
exit
!
! the 'interface' address is the same as the publicly accessible IP address.
object network nat-DNS1-53-udp
host 10.0.0.1
nat (inside,outside) static interface service udp domain domain
exit
!
It's weird because I've got the ACL set to allow from any to any, the packet trace is set up to test on UDP as per the rules (I do have TCP rules in there too but same result when I run packet-trace with tcp instead of udp). It hits the implicit deny rule, even in that case. It's probably something stupid but I can't see what I'm missing.
Section des commentaires
how is the ASA routing to 10.0.0.1?
Worth noting that it says gateway of last resort isn't set because it's not presently hooked up - our current firewall is. When set, gateway of last resort reads to the IP address of the next hop out (ISP's router at the demarcation point), which is how our current firewall is set up. It is within our /29 subnet of assigned addresses.
10.0.1.254 is a L3 device with correct routing set up, and is currently routing all traffic. The only thing I do there is change its default route to that of the ASA instead of the existing firewall.
was the packet-tracer you ran that gave the output above run when the device was on the Internet?
everything looks fine from what I can tell. your outside and inside interfaces are separate security levels, right?
Something isn't set up correctly. What version of the code are you running? I have just tested it in lab environment and it works fine.
Your phase 2 should show both access-group and access-list you are hitting:
Is phase 1 showing that your NAT rule was hit correctly? You should see UN-NAT.
When you are performing packet-tracer are you using IP address assigned to the outside interface?
Phase 2 definitely didn't show anything about un-natting. Unfortunately I can't get that back without hooking it up again, so I'll do that at lunch to re-run the packet trace. It'll be an outage event but relatively minor for us.
Off the top of your head can you think of anything that I could have set wrong with the overall config that would cause my dynamic nat rules to not apply properly?
It is not the dynamic NAT that is the issue. It is the static NAT that you need to hit:
Your phase 1 should be UN-NAT of transleting your public IP to 10.0.0.1 when UDP port 53 is hit on the outside interface.
My guess is that you are either using a code with a bug (unlikely) or that you are not using correct outside IP address in your packet-tracer command.
Also, is your 10.0.0.1 behind inside interface? That could also be the problem if 10.0.0.1 is behind a different interface name.
