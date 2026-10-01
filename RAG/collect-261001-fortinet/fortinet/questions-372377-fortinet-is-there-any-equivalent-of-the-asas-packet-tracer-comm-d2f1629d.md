---
id: collect-261001-fortinet/fortinet/questions-372377-fortinet-is-there-any-equivalent-of-the-asas-packet-tracer-comm-d2f1629d
title: "questions-372377-fortinet-is-there-any-equivalent-of-the-asas-packet-tracer-comm-d2f1629d"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-372377-fortinet-is-there-any-equivalent-of-the-asas-packet-tracer-comm-d2f1629d.md
source_anchor: ""
source_lines: [1, 96]
sha256: 308fde0cd057e93707ae99396c710f4d359e15b1ee13e07781fa7a02e51042a2
---

# questions-372377-fortinet-is-there-any-equivalent-of-the-asas-packet-tracer-comm-d2f1629d

I would like to know if there is not Fortigates an equivalent of the packet-tracer command that we can find on the ASA.
Here is an example of execution for those who don't know it:
NAT and pass :
lev5505# packet-tracer input inside tcp 192.168.3.20 9876 8.8.8.8 80
Phase: 1
Type: ACCESS-LIST
Subtype: 
Result: ALLOW
Config:
Implicit Rule
Additional Information:
MAC Access list
Phase: 2
Type: ROUTE-LOOKUP
Subtype: input
Result: ALLOW
Config:
Additional Information:
in   0.0.0.0         0.0.0.0         outside
Phase: 3
Type: ACCESS-LIST
Subtype: log
Result: ALLOW
Config:
access-group inside-in in interface inside
access-list inside-in extended permit tcp any any eq www 
access-list inside-in remark Allows DNS
Additional Information:
Phase: 4
Type: IP-OPTIONS
Subtype: 
Result: ALLOW
Config:
Additional Information:
Phase: 5
Type: VPN
Subtype: ipsec-tunnel-flow
Result: ALLOW
Config:
Additional Information:
Phase: 6
Type: NAT
Subtype: 
Result: ALLOW
Config:
object network inside-network
 nat (inside,outside) dynamic interface
Additional Information:
Dynamic translate 192.168.3.20/9876 to 81.56.15.183/9876
Phase: 7
Type: IP-OPTIONS
Subtype: 
Result: ALLOW
Config:
Additional Information:
Phase: 8
Type: FLOW-CREATION
Subtype: 
Result: ALLOW
Config:
Additional Information:
New flow created with id 94755, packet dispatched to next module
Result:
input-interface: inside
input-status: up
input-line-status: up
output-interface: outside
output-status: up
output-line-status: up
Action: allow
Blocked by ACL:
lev5505# packet-tracer input inside tcp 192.168.3.20 9876 8.8.8.8 81
Phase: 1
Type: ROUTE-LOOKUP
Subtype: input
Result: ALLOW
Config:
Additional Information:
in   0.0.0.0         0.0.0.0         outside
Phase: 2
Type: ACCESS-LIST
Subtype: 
Result: DROP
Config:
Implicit Rule
Additional Information:
Result:
input-interface: inside
input-status: up
input-line-status: up
output-interface: outside
output-status: up
output-line-status: up
Action: drop
Drop-reason: (acl-drop) Flow is denied by configured rule
Is there any equivalent on the Fortigates ?
