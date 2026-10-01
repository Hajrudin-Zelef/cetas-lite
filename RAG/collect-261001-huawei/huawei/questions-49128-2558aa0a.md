---
id: collect-261001-huawei/huawei/questions-49128-2558aa0a
title: "questions-49128-2558aa0a"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2018-04-01"]
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/questions-49128-2558aa0a.md
source_anchor: ""
source_lines: [1, 9]
sha256: f40978a64a74453553250820575dd7a043dc21c0e3fa5edd687c031319449cd1
---

# questions-49128-2558aa0a

everyone! I want to buy a Huawei switch with GE ports, does Huawei S6720 model have GE port?
- 
        Did any answer help you? If so, you should accept the answer so that the question doesn't keep popping up forever, looking for an answer. Alternatively, you could provide and accept your own answer.Ron Maupin– Ron Maupin ♦2018-04-01 23:05:14 +00:00Commented Apr 1, 2018 at 23:05
2 Answers 2
There's something like 22 models of Huawei S6720, so it would be better to check with the specific model, but this answer should apply to all.
According to this page from Huawei documentation
  10GE SFP+ port A 10GE SFP+ Ethernet optical port supports auto-sensing to 1000 Mbit/s. It sends and receives service data at 1000 Mbit/s or 10 Gbit/s.
So while those switches have only 40/10G interfaces (except for the management interface) those interfaces can operate at gigabit speed.
Huawei S6720 model doesn't have GE port, it is 10G SFP+ switch (also have 40G ports), but ports on S6720 are auto-sensing from GE to 10GE; if you put GE module, then it will be GE ports.
