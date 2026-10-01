---
id: collect-261001-general-networking/general-networking/t5-switching-ms425-flexible-stacking-with-link-aggregate-m-p-26587-23c343fe
title: "t5-switching-ms425-flexible-stacking-with-link-aggregate-m-p-26587-23c343fe"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-ms425-flexible-stacking-with-link-aggregate-m-p-26587-23c343fe.md
source_anchor: ""
source_lines: [1, 59]
sha256: b3f43a661c4655fa0c1d2aefa77ae81d05a33106d64c5bbb1e8df4c92a00e42a
---

# t5-switching-ms425-flexible-stacking-with-link-aggregate-m-p-26587-23c343fe

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-26-2018 02:59 PM
Hi All,
I wonder if anyone can pass on some advice on the follow topology please?
We are looking at a Core switch design as shown below, using MS425’s and MS225’s
The plan now is to use Flexible stacking QSFP between the two MS425’s and also have the MS225 connected as a physical stack with a single 10Gb uplink each of the two MS225’s back the MS425’s.
My question is can/should we create a Link Aggregate as well using the two 10Gb uplinks, knowing that we have the MS425’s Flexible stacked and all the MS225’s configured as a stack as well?
Our original approach was to have the Link Aggregate between MS425’s and MS225’s, but not have the MS425’s Flexible Stacked via QSFP.
Any thoughts on this please. Thanks in advance.
BB
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
08-26-2018 05:44 PM
What you have planned is a very standard deployment of MS425's at the core and MS225's at the edge. I would stack the 425's with QSFP and then create LAG's to the 225 stacks with a link going to each of the 425's.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-26-2018 05:44 PM
What you have planned is a very standard deployment of MS425's at the core and MS225's at the edge. I would stack the 425's with QSFP and then create LAG's to the 225 stacks with a link going to each of the 425's.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-26-2018 11:20 PM
I agree wih @MRCUR because this creates a loop free design.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-27-2018 05:19 AM
