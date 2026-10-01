---
id: collect-261001-general-networking/general-networking/t5-switching-eigrp-hsrp-dilemma-m-p-610098-highlight-true-d8692230
title: "t5-switching-eigrp-hsrp-dilemma-m-p-610098-highlight-true-d8692230"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "ethernet"]
source: docs/RAG/collect-261001-general-networking/t5-switching-eigrp-hsrp-dilemma-m-p-610098-highlight-true-d8692230.md
source_anchor: ""
source_lines: [1, 138]
sha256: ddf6b90cfd1aa9afc8907d2233f188dc8c5e4ad7d65484f33eef4d2556fc7d8e
---

# t5-switching-eigrp-hsrp-dilemma-m-p-610098-highlight-true-d8692230

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-30-2007 10:11 PM - edited 03-05-2019 02:04 PM
I have 4 routers on one ethernet media. Two routers works in kind of pair and should be backup for each other. Now I have dilemma. Should I use EIGRP and have 4 EIGRP neighbours or .... can I put two routers in HSRP and then use HSRP virtual address as EIGRP neighbour? What is better?
thanks
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Other Switching
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-31-2007 03:55 AM
Dragan
No you can not. EIGRP will form neighbor relationships on the physical (local) address and EIGRP will NOT use the HSRP virtual address.
I am not clear what your environment is and what you are trying to accomplish. But if you want 2 of the routers to talk to each other and not to the other 2 routers, there is a simple solution. Configure 2 of the routers with one AS number and configure the other 2 with a different AS number. Routers on the same segment running EIGRP with different AS numbers will not become neighbors and will not exchange routing updates.
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
01-30-2007 10:27 PM
Hi,
You should always run the routing protocol between all participating Routers.
Let us know more about your network and what you are trying to achieve
-VJ
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-30-2007 11:36 PM
I am trying to find out if I could benefit from using HSRP virtual address as source for EIGRP.
thanks!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-31-2007 03:55 AM
Dragan
No you can not. EIGRP will form neighbor relationships on the physical (local) address and EIGRP will NOT use the HSRP virtual address.
I am not clear what your environment is and what you are trying to accomplish. But if you want 2 of the routers to talk to each other and not to the other 2 routers, there is a simple solution. Configure 2 of the routers with one AS number and configure the other 2 with a different AS number. Routers on the same segment running EIGRP with different AS numbers will not become neighbors and will not exchange routing updates.
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
01-31-2007 04:05 AM
Hi Rick,
basically with two AS numbers the network would be split into two parts not able to exchange routing info other than through redistribution. And I think you would agree not to opt for such a solution if other options exist because of the likelyness of routing loops.
So the two options I would consider are
1) EIGRP with one AS on all 4 routers
2) HSRP plus static routing
The first option is straight forward, easy to configure and unless other/additional requirements are presented what I would implement.
The second option requires to use static routing to the opposite HSRP next hop from both sides. Should also work, but requires more administrative attention in case new networks are introduced.
Regards, Martin
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-31-2007 07:41 AM
Martin
As I said in my post I think that we need a better understanding of what Dragan is wanting to accomplish. Based on this sentence: "Two routers works in kind of pair and should be backup for each other" I assume that he does not want all 4 routers to share advertisements with each other. And that is why I suggested 2 different AS be used. That would allow each pair of routers to work together and to back each other up but not to share advertisements with the other pair.
But we really do not know what Dragan wants to accomplish. And till we do know there are lots of possibilities to consider.
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
01-31-2007 07:45 AM
True Rick,
sorry for the misunderstanding. Did not want to attack you. [Edit on] Did not want to be unpolite [Edit off].
Just having sometimes problems to get the right dosis of fine tuning into my English, being non native speaker.
Regards, Martin
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-31-2007 09:18 AM
Martin
No Problem. I do not feel attacked or any lack of politeness. And I sympathesize with the difficulty of communicating in a language which is not your native language. And I suspect that language differences are part of the reason that we each picked different parts of the original post as guiding our responses.
Occasional differences of opinion keep the forum interesting and frequently point out that there may be more than one viable solution.
Keep up the good work.
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
01-31-2007 11:12 AM
hi rick and martin,
issue is not language... but most important part is both of you tried to solve the problem... if found both you people on the same disscussion forum since i join this forum... and definetly both of you people had resovle the lots of post... and i do have lots of respect for you as well as all of the Netpro experts and member... this forum keep every one update and also save from the boss attack ...
and even lots of time i learned and improve my mistakes from this forum and still doing the same thing... and CISCO is like open sky... as you try to go in deep its become more and more deeper... so definelty my big respect to CISCO also...
regards
Devang
