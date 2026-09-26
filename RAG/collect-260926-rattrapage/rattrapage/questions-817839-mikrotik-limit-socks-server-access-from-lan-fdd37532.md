---
id: collect-260926-rattrapage/rattrapage/questions-817839-mikrotik-limit-socks-server-access-from-lan-fdd37532
title: "questions-817839-mikrotik-limit-socks-server-access-from-lan-fdd37532"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-rattrapage/servers-reviews/questions-817839-mikrotik-limit-socks-server-access-from-lan-fdd37532.md
source_anchor: ""
source_lines: [1, 11]
sha256: 6e7e0972457f95677d4aac8461c995fdaaedf80f0df19383c2d99d376cd3e0ef
---

# questions-817839-mikrotik-limit-socks-server-access-from-lan-fdd37532

I have a Mikrotik RB750 router that Socks server is configured on it. It has a interface in my LAN and another one in WAN with public IP address. I want to limit access to it from LAN. It seems that some people has found it and using it form WAN!!!
2 Answers 2
Open WinBox . Go to IP > Socks > Access . The Socks Access window is similar to filter rule window. Default action in Socks Access is accept. This means that if no rule get matched, socks server accept that connection. 
Now add a rule with accept action your LAN IP addresses range as its Src Address. Then add a rule with deny action. Left other field unchanged. This rule guaranty denying of any socks request other than your LAN IP addresses.
- 
        
- 
        Yea. Rule matching is done from low index (below # field) to high index. If any rule is matched, it exit rule matching process.Cierra Clark– Cierra Clark2016-11-30 11:06:54 +00:00Commented Nov 30, 2016 at 11:06
If socks server is running on port 1080 and your LAN IP address range is 192.168.10.0/24 use this in Mikrotik terminal:
/ip firewall filter add action=drop chain=input dst-port=1080 protocol=tcp src-address=!192.168.10.0/24
This command filters traffic by Mikrotik firewall.
