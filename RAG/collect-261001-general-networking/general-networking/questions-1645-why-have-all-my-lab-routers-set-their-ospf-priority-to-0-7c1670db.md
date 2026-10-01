---
id: collect-261001-general-networking/general-networking/questions-1645-why-have-all-my-lab-routers-set-their-ospf-priority-to-0-7c1670db
title: "questions-1645-why-have-all-my-lab-routers-set-their-ospf-priority-to-0-7c1670db"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1645-why-have-all-my-lab-routers-set-their-ospf-priority-to-0-7c1670db.md
source_anchor: ""
source_lines: [1, 5]
sha256: f00856e43c34dbe34e96b7b01908f205d39590a56993313f9df4665b6683e61b
---

# questions-1645-why-have-all-my-lab-routers-set-their-ospf-priority-to-0-7c1670db

I have 4 routers in my lab. FR1 is the frame relay switch (2621), R1(2621), R2(2651) & R3(2651) are all connected in a mesh using sub interfaces - R1 has pvc to R2&R3 and so on. Without OSPF I can ping all points successfully. With OSPF I can also ping all the loopbacks built onto the routers. OSPF seems to work but when I do:
#show ip ospf neighbor
the neighbours show as FULL but with 0 priority - and checking the detail each router has 0.0.0.0 as their DR.
I'm about to try changing the priorities, but given the fact I did a write erase before building the lab this evening, what could cause the priorities to move from the default of 1 to 0?
Thanks in advance.
