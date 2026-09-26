---
id: collect-260926-mikrotik/mikrotik/questions-10872-routeros-creating-user-only-with-access-to-tools-759cb31b
title: "questions-10872-routeros-creating-user-only-with-access-to-tools-759cb31b"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-10872-routeros-creating-user-only-with-access-to-tools-759cb31b.md
source_anchor: ""
source_lines: [1, 5]
sha256: 00778b3165e4a14aa5014c6218b0e45b0e1059ab044ed52f59ecdc77ecbe9166
---

# questions-10872-routeros-creating-user-only-with-access-to-tools-759cb31b

I have 3 different networks and I'd like users from network A to be able to wake up pc network B (ofc with certain password) so I guess it could be good idea to allow them to use WOL capabilities built in router.
- 
        how about setting up a raspberry pi for this task, that one could as well log SNMP data from the router and print nice graphs.Daniel F– Daniel F2014-09-12 19:28:38 +00:00Commented Sep 12, 2014 at 19:28
1 Answer 1
You can use Webfig to create an interface for a certain user where he can only access certain areas of the router management (tools in your scenario).
