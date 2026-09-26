---
id: collect-260926-mikrotik/mikrotik/questions-1667857-block-video-calls-with-mikrotik-05f1c710
title: "questions-1667857-block-video-calls-with-mikrotik-05f1c710"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-1667857-block-video-calls-with-mikrotik-05f1c710.md
source_anchor: ""
source_lines: [1, 10]
sha256: 346360dbe63e7d9ffbd1739b0f898723e8bc904b26c9b674cedfb3c83eebc164
---

# questions-1667857-block-video-calls-with-mikrotik-05f1c710

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have created a simple firewall filter rule on Mikrotik router to disable internet access on certain ip addresses at a certain time. Everything works fine, but if the user is currently on a video call such as Skype or MS Teams and the rule kicks in, the user is not disconnected and is able to proceed with the video call. I would like the rule to disable anything to do with internet which also includes terminating any active video calls etc. The rule I used is the following:
add action=drop chain=forward comment="Disable Internet by time" src-address-list="!Allow Lan" time=23h-5h,sun,mon,tue,wed,thu,fri,sat
The connection is already in the state table, and therefore it's not checked against the firewall rule list. That's what allows the existing connections to continue.
What you would have to do is drop connections from the state table, or reset all of them. I'm not sure how this works on Mikrotik, a quick google yielded the following:
