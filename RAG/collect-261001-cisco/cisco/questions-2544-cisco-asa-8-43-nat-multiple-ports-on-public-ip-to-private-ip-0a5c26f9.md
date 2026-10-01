---
id: collect-261001-cisco/cisco/questions-2544-cisco-asa-8-43-nat-multiple-ports-on-public-ip-to-private-ip-0a5c26f9
title: "questions-2544-cisco-asa-8-43-nat-multiple-ports-on-public-ip-to-private-ip-0a5c26f9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-2544-cisco-asa-8-43-nat-multiple-ports-on-public-ip-to-private-ip-0a5c26f9.md
source_anchor: ""
source_lines: [1, 9]
sha256: 4eecb00ed0413b856dbca80543aac2e34e406f7b40333727ce228b7125605221
---

# questions-2544-cisco-asa-8-43-nat-multiple-ports-on-public-ip-to-private-ip-0a5c26f9

I have the following configuration on an ASA running 8.4 already;
object network SERVER-01
 host 192.168.0.1
!
object network SERVER-01
 nat (Inside,Outside) static interface service tcp 3389 5001
I have this config repeated several times to allow RDP to multiple machines behind the ASA, on different ports; 5001 forwards to one server, 5002 to another, and so on.
For this one server in the above config though, I want to forward port 5555 publicly to port 5555 at the same private IP, 192.168.0.1.
If I try to enter the above config but with different port numbers, it is overwritten, I can't seem to add an additional NAT entry, only replace that one. This indicates I can only have one port forwarded to each internal IP at a time. Is that correct, or can I forward various public ports to an internal host, but using some other method? If so, how?
