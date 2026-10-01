---
id: collect-261001-general-networking/general-networking/questions-5364-is-it-possible-to-connect-two-private-ebgp-peers-with-the-same-as-8fa4fca3
title: "questions-5364-is-it-possible-to-connect-two-private-ebgp-peers-with-the-same-as-8fa4fca3"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-5364-is-it-possible-to-connect-two-private-ebgp-peers-with-the-same-as-8fa4fca3.md
source_anchor: ""
source_lines: [1, 14]
sha256: 6511b072e8b2f7d06b8b9ef12fe8772420ee8318516dc44c7c428003a0669d3d
---

# questions-5364-is-it-possible-to-connect-two-private-ebgp-peers-with-the-same-as-8fa4fca3

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
12
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
The Problem
A third party has implemented a small network and private BGP setup within one of our branches. We need to setup a eBGP peer relationship with this network. However we are both using the same private BGP AS number.
The Solution
One of us changes our BGP AS number - This isn't out of the question, and even if there is another solution but its overly complex we may still do this.
BGP Confederation - In my research it appears that setting up a BGP Confederation between the peers might be a work around. The basic idea would be to sub-divide the shared private AS number into multiple private AS numbers. I'm not entirely sure yet whether this would solve the problem properly.
Other solutions? It is here I'm looking for more advice on whether it is possible to work around this problem is an effective way.
Option 1 is likely your best bet for simplicity's sake, but you can use the confederation method if you don't want to change the ASN. You can also do neighbor x.x.x.x local-as <as> in the BGP config but this prepends an ASN of your choosing onto the path, rather than replace the ASN in the path, so updates from one router to the other would be dropped anyway. Unless there's a really good reason as to why you want to maintain the eBGP session, you can also migrate to an iBGP session vs eBGP.
As @DanielDib has pointed out below, a third option (if you're running a version of IOS that supports it) is to use local-as in conjunction with replace-as and no-prepend to strip your ASN from the AS_PATH and send a different one:
