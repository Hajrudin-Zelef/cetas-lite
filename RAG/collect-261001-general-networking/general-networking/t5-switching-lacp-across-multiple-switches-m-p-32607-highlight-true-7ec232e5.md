---
id: collect-261001-general-networking/general-networking/t5-switching-lacp-across-multiple-switches-m-p-32607-highlight-true-7ec232e5
title: "t5-switching-lacp-across-multiple-switches-m-p-32607-highlight-true-7ec232e5"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-lacp-across-multiple-switches-m-p-32607-highlight-true-7ec232e5.md
source_anchor: ""
source_lines: [1, 73]
sha256: e597e124fa6fb23ec06e164de9775dbdf0047df11ae170454ae3962c155cec66
---

# t5-switching-lacp-across-multiple-switches-m-p-32607-highlight-true-7ec232e5

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2018 10:46 AM
The fact that I can't configure it, might actually answer my question, but here goes anyway.
Is it possible to aggregate ports across multiple switches? On MS120-8LP?
I have an MS220-48FP connected to two MS120-8LP's. The uplink ports on the MS220 are aggregated, but each port on each MS120, are not.
The two MS120's are interconnected, as well as connected to each of their own MX64.
LinkedIn ::: https://blog.rhbirkelund.dk/
Like what you see? - Mark as helpful ## Did it answer your question? - Mark it as a Solution
All code examples are provided as is. Responsibility for Code execution is solely your own.
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2018 06:04 PM
You can only aggregate ports across switches when those switches are physically stacked together using their stacking port.
MS120's do not have a physical stacking port.
So in your case, you can not using port aggregation.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2018 11:53 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2018 11:59 AM
I wanted to experiment with if LACP would do some loadbalancing between the two switches, or maybe gain some bandwidth on the uplink, as well as not block ports. So far, even LACP is blocking traffic on one of the links.
LinkedIn ::: https://blog.rhbirkelund.dk/
Like what you see? - Mark as helpful ## Did it answer your question? - Mark it as a Solution
All code examples are provided as is. Responsibility for Code execution is solely your own.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2018 12:22 PM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2018 06:04 PM
You can only aggregate ports across switches when those switches are physically stacked together using their stacking port.
MS120's do not have a physical stacking port.
So in your case, you can not using port aggregation.
