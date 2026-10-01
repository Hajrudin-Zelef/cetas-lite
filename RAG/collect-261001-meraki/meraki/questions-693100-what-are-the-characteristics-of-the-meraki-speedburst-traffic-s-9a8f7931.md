---
id: collect-261001-meraki/meraki/questions-693100-what-are-the-characteristics-of-the-meraki-speedburst-traffic-s-9a8f7931
title: "questions-693100-what-are-the-characteristics-of-the-meraki-speedburst-traffic-s-9a8f7931"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-meraki/questions-693100-what-are-the-characteristics-of-the-meraki-speedburst-traffic-s-9a8f7931.md
source_anchor: ""
source_lines: [1, 14]
sha256: 25751fed49cd5938de8ad34856e2c35faedd370b3bb29a19855856475c267bbc
---

# questions-693100-what-are-the-characteristics-of-the-meraki-speedburst-traffic-s-9a8f7931

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
4
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I've been using Cisco Meraki wireless access points to provide guest and BYOD access at some of my customer sites. One of the interesting bandwidth management features is the SpeedBurst traffic shaping option.
This is described simply as a temporary suspension of the bandwidth limit to make access feel "snappier", followed by a throttling down to the fixed limit.
While giving a demo of a BYOD network for a client, I was asked how the SpeedBurst option worked, and didn't really have a good answer.
I'm curious about the specifics of this feature.
Is the algorithm described in detail anywhere?
How long is traffic burstable?
What do repeated requests from the connected client look like and how does that impact overall speed and experience?
Enable SpeedBurst: To provide a better user experience in bandwidth-limited environments, an administrator can enable SpeedBurst by selecting the Enable Speedburst checkbox. SpeedBurst allows users to exceed their assigned limit in a "burst" for a short period of time, providing a more satisfying Internet browsing experience while still preventing any one user from using more than his or her fair share of bandwidth over the longer term. Users are allowed up to four times their allotted bandwidth limit for a period of up to five seconds.
