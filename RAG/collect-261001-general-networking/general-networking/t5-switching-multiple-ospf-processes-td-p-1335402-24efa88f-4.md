---
id: collect-261001-general-networking/general-networking/t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f-4
title: "t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f.md
source_anchor: ""
source_lines: [295, 319]
sha256: 07255e9ce8a1c3256461a15fab52ef3a86557b11de0f758ad7167834bb325fc3
---

# t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f

c) each OSPF process will advertise only prefixes from its own data base. The only way to have one process advertise prefixes from the other process is to redistribute.
In a), did you mean "An interface can be active in only one process."? Or, "A process can be active on only one interface."? I think you may have duplicated the word interface unwittingly. Please help me underestand what you meant to say. Thanks.
Great explanation, nonetheless.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-21-2013 08:46 PM
Kerry
Thank you for your attention to detail in reading an old post. You are absolutely correct that I did unwittingly duplicate the word interface. What I intended to say is that an interface can be active in only 1 process. And to expand that thought I would say that while several OSPF processes may have network statements that include a particular interface that the interface can be active in only one process at a time.
HTH
Rick
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-27-2013 07:54 AM
What about area-0 when there are two processes declared? Can you have two different area-0 in the same router (one per process)?
