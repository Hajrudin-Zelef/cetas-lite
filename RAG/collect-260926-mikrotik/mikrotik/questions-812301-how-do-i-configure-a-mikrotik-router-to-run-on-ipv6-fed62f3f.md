---
id: collect-260926-mikrotik/mikrotik/questions-812301-how-do-i-configure-a-mikrotik-router-to-run-on-ipv6-fed62f3f
title: "questions-812301-how-do-i-configure-a-mikrotik-router-to-run-on-ipv6-fed62f3f"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-812301-how-do-i-configure-a-mikrotik-router-to-run-on-ipv6-fed62f3f.md
source_anchor: ""
source_lines: [1, 11]
sha256: 62a52160437b0fcd9c290f1a81834ce688ab96b3b0c7518324110199f346f45d
---

# questions-812301-how-do-i-configure-a-mikrotik-router-to-run-on-ipv6-fed62f3f

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am experimenting with ipv6 and would like to connect to my Mikrotik router via ipv6 on a local level. This would include connecting to it by an ip6 address, and interacting via icmp6. I have installed the ipv6 package and made sure it was enabled through the command line.
Since it seems many people on here only like to be insulting instead of answering a simple question, please be assured I have actually spent quite a bit of time trying to find the answer to this question. How can I connect to the router and interact with it ONLY by ipv6? (no IP4)
I figured it out! I connected to the router by using winbox and the default ipv4 address. I can also connect to it by typing that ipv4 address into a browser window.
From within winbox, I typed "system package print" to reveal all packages available for the router. I saw that ipv6 was on the list, but was disabled, so I typed "system package enable ipv6" to enable the package, then rebooted the router. Ipv6 is now enabled, and I typed "ipv6 address print" to find a list of available ipv6 addresses.
I opened a browser window, and pasted the IPv6 address I got from my list. NOTE: IPv6 addresses cannot be used in quite the same was as IP4. They must be surrounded in square brackets, like so: [xx:xxx::xx etc]. That worked! I am now connected by IPv6.
