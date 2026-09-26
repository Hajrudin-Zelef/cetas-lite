---
id: collect-260926-rattrapage/rattrapage/questions-1293579-mikrotik-rb2011uias-2hnd-in-experiences-wifi-loss-until-wlan-i-dfdd97f2
title: "questions-1293579-mikrotik-rb2011uias-2hnd-in-experiences-wifi-loss-until-wlan-i-dfdd97f2"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["research", "voice"]
source: docs/RAG/lot-rattrapage/servers-reviews/questions-1293579-mikrotik-rb2011uias-2hnd-in-experiences-wifi-loss-until-wlan-i-dfdd97f2.md
source_anchor: ""
source_lines: [1, 14]
sha256: 1f83e5842731a46113e5c33beff3aa5eb8c2d2f6b0d106c5b53c49e4032f0806
---

# questions-1293579-mikrotik-rb2011uias-2hnd-in-experiences-wifi-loss-until-wlan-i-dfdd97f2

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Driven by recently discovered WPA2 security hole (key reinstallation attack) I updated RouterOS in RB2011uias-2hnd-in from v6.39 to v6.40.5.
It didn't improve connection quality, but now I'm experiencing from time to time a complete WiFi loss which could be resolved either by WLAN interface disable/enable or by complete router reboot. WLAN can work for several days or may disappear in few minutes after restart.
Do you experience this too? Did you fix it? How?
Update
Yesterday I updated RouterOS and firmware to v 6.41.2 and now WiFi is always there ... just the access to network disappears with same frequency as WiFi used to disappear. In the recent version they did something to RouterOS bridges, but WLAN restart helps for a while.
The problems you are experiencing with this version are being shared by other users. Make your voice heard, they are asking for anyone with bugs or issues with this release to report them to microtik. You are perfectly safe to roll-back to v6.39 as it is not vunerable to the attack.
https://forum.mikrotik.com/viewtopic.php?t=127485
I'm not sure how that version made it past QA. But deffinately report your problem.
