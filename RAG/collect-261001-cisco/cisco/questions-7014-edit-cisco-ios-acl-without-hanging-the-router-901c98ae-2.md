---
id: collect-261001-cisco/cisco/questions-7014-edit-cisco-ios-acl-without-hanging-the-router-901c98ae-2
title: "questions-7014-edit-cisco-ios-acl-without-hanging-the-router-901c98ae"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-7014-edit-cisco-ios-acl-without-hanging-the-router-901c98ae.md
source_anchor: ""
source_lines: [177, 197]
sha256: 56b8fc3aa64dbafa91eec0a10652ae5d4eddec9739c7be5e07f70a26641642de
---

# questions-7014-edit-cisco-ios-acl-without-hanging-the-router-901c98ae

You should really consider using two different ACLs for Gigabit0/1.1 and GigabitEthernet0/1.2... this is a guess at what you're trying to do, but it's unclear that I'm interpreting things correctly...
access-list 111 permit ip 192.168.1.0 0.0.0.255 host 192.168.2.44
access-list 111 permit ip host 192.168.1.18 192.168.2.0 0.0.0.255
access-list 111 permit ip host 192.168.1.120 192.168.2.0 0.0.0.255
access-list 111 permit ip host 192.168.1.222 192.168.2.0 0.0.0.255
access-list 111 deny ip 192.168.1.0 0.0.0.255 192.168.2.0 0.0.0.255
access-list 111 permit udp any any
access-list 111 permit ip any any
!
interface GigabitEthernet0/1.1
no ip access-group 110 in
ip access-group 111 in
!
interface GigabitEthernet0/1.2
no ip access-group 110 in
It likely hangs because upon pasting "the first statement is put into effect, and the implicit deny statement that follows could cause you immediate access problems."
Helpful Hints for Creating IP Access Lists ¹
•Create the access list before applying it to an interface. An interface with an empty access list applied to it permits all traffic.
•Another reason to configure an access list before applying it is because if you applied a nonexistent access list to an interface and then proceed to configure the access list, the first statement is put into effect, and the implicit deny statement that follows could cause you immediate access problems.
ACLs should be tuned in after implementation to save processor cycles. The router has to process each line of an ACL until it gets a match on one of the conditions, or otherwise will always match the implicit 'deny any' at the end of all ACLS. Use a show command to see the hits on each ACL statement. Reorder the statements so that the most hit lines are higher in the ACL.
Use notepad to rewrite your ACL.
