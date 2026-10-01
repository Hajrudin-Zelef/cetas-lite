---
id: collect-261001-cisco/cisco/questions-578-why-does-my-cisco-6509-bgp-table-use-two-entries-in-my-tcam-966531df
title: "questions-578-why-does-my-cisco-6509-bgp-table-use-two-entries-in-my-tcam-966531df"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "research"]
source: docs/RAG/collect-261001-cisco/questions-578-why-does-my-cisco-6509-bgp-table-use-two-entries-in-my-tcam-966531df.md
source_anchor: ""
source_lines: [1, 31]
sha256: 7981dee0179bbd5664015a6ca116bb614555bd342323c8a64a1e17b87d549a05
---

# questions-578-why-does-my-cisco-6509-bgp-table-use-two-entries-in-my-tcam-966531df

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
10
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a problem on my Cisco 6509, each entry in my BGP table occupies two entries in the TCAM.
If I show capacity forwarding, I see MPLS entries in L3 forwarding resources.
But, I do not use MPLS on my chassis !
#show run | i mpls
mls cef maximum-routes mpls 508
no mpls ldp advertise-labels
no mpls ip
It seems that the 6500 generates MPLS labels for every route if BGP is run in VRF. The fact that your IPv4 and MPLS TCAM usage is almost identical seems to indicate this as well. Can you try this command:
show bgp vpnv4 uni all labels
There seems to be a hidden command that makes IOS allocate labels per VRF instead of per prefix.
Oh the 6500. I run a small service provider network and run the 6500 as a PE router. Worst decision of my life. (That was an embellished statement, but you get my point.)
I run full BGP routes in a VRF and have experienced a lot of problem surrounding this.
You're example is not very surprising. As Daniel said in his post there is an LFIB entry for each VRF prefix as well as a VPNv4 entry. This can be changed by adding the command mpls label mode vrf Internet protocol all-afs per-vrf as was stated; however, this does not get you out of the woods. If you change to per VRF prefixes it does remove the LFIB entry (yay!) but adds an entry for every single prefix into the Adjacency table (wait, what?!). Since the 6500 forwarding hardware is shared between L2 and L3 forwarding this doesn't change your hardware memory usage at all. If anything it makes the problem harder to find.
If you look at your usage once you've changed to per VRF usage (using show platform hardware cef resource-level) it looks as though you've fixed the issue. However if you use the command show platform hardware cef adjacencies resource-level it reveals the problem has just moved to a different location.
Below are the outputs from one of my 6500's resource-level and adjacency usage. Outlining what I'm talking about.
Watermarks apply to regions available for allocation and not pre-reserved
Stats region size for alloc: 444160
Non-stats region size for alloc: 376832
Adjacency Mgr watermarks:
Type Red_WM(%) Green_WM(%) Current usage(%)
---- --------- ---------- ----------------
Stats_WM 95% 80% 97%
Non-Stats_WM 95% 80% 14%
Ivan's post on this was base on my findings here. I am currently working with Cisco to attempt to fix this issue, but unfortunately right now there is no way to fix this.
Your milage may vary since you've got no MPLS adjacencies. Would be interested to see your adjacency usage now that you've made the change.
