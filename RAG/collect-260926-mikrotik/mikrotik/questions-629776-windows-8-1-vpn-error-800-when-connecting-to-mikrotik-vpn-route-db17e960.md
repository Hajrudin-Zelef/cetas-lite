---
id: collect-260926-mikrotik/mikrotik/questions-629776-windows-8-1-vpn-error-800-when-connecting-to-mikrotik-vpn-route-db17e960
title: "questions-629776-windows-8-1-vpn-error-800-when-connecting-to-mikrotik-vpn-route-db17e960"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vpn/questions-629776-windows-8-1-vpn-error-800-when-connecting-to-mikrotik-vpn-route-db17e960.md
source_anchor: ""
source_lines: [1, 17]
sha256: 72acda017d5053e05e6b53acb73dbecf0329b9cc8997110399f444f86cc74b9b
---

# questions-629776-windows-8-1-vpn-error-800-when-connecting-to-mikrotik-vpn-route-db17e960

After creating a new VPN connection on a Windows 8.1 machine, to a VPN server (using Mikrotik router), the connection times out and an error no 800 is displayed.
- 
        Please post your answer as an answer, not as part of your question.EEAA– EEAA2014-09-19 12:08:37 +00:00Commented Sep 19, 2014 at 12:08
2 Answers 2
Mladen B. answered his own question:
  Digging on the Mikrotik forum I found a post which resolves this
  issue. Shortly, open the registry and navigate to the entry:
  HKEY_LOCAL_MACHINE\System\CurrentControlSet\Control\SecurityProviders\SCHANNEL
There, create a new DWORD key named SendExtraRecord with a Hex value
  of 2.After that retry the VPN connection (no need for a restart).
I just wanted to share this finding, since it is a very common and the page where the solution can be found is not very well optimized for search engines, so I'll just leave it here for people to easier solve it, if they encounter it.
I found a tutorial about fixing vpn error 800 in windows 7, windows 8,windows 8.1 on this website I Trick Buzz
How to fix Error 800 Vpn windows 8.1
- Try pinging the server to make sure it is reachable.
- Check VPN IP address and VPN username password.
- Check firewall and router settings.
Source : How to fix VPN Error 800 in windows 8.1
