---
id: collect-261001-cisco/cisco/questions-146010-cisco-ios-qos-prioritize-ssh-but-not-scp-21f13575
title: "questions-146010-cisco-ios-qos-prioritize-ssh-but-not-scp-21f13575"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-cisco/questions-146010-cisco-ios-qos-prioritize-ssh-but-not-scp-21f13575.md
source_anchor: ""
source_lines: [1, 18]
sha256: a97c14bef7e1f442e872a67091160743db672b694b8af8b6b88da4d52851a301
---

# questions-146010-cisco-ios-qos-prioritize-ssh-but-not-scp-21f13575

How can I prioritize ssh traffic to ensure low-latency but throttle SCP file transfers? I'm looking for a solution that is not host specific so I don't have to add lists of IP addresses. Thanks!
4 Answers 4
You can use packet lengths to distinguish one from the other, but there's a risk that these will be chopped and mixed by the SSH transport.
class-map match-all ssh-interactive
 match access-group name ssh
 match packet length max 600
class-map match-all ssh-files
 match access-group name ssh
 match packet length min 600 
ip access-list extended ssh
 permit tcp any any eq 22
 permit tcp any eq 22 any
The cut-off point is arbitrarily picked out of this air, you'd have to tune and see what works for you.
Most (all? OpenSSH at least) SSH implementations set different IP TOS flags depending on if the session is interactive or bulk data. It sets the lowdelay TOS flag for interactive sessions.
You can then match on the TOS flags in your ACL.
This really is not possible - both operate on the same port, and both encapsulate their traffic inside an encrytped (and optionally compressed) stream. There may be a way for a DPI system to make an guess at what stream is a "standard" terminal ssh session and which is an SCP session, but that would be a wild guess at best.
Well I agree with ErikA about the technical details, I (partly) disagree that it is not possible. You can run multiple ssh daemons on different ports and the prioritize based on ports. Here is a blog post on how to do that.
From there you are going to need different credentials for each ssh daemon (not sure how to do this, but I bet you can), unless you trust your users to choose. Another option would be to only allow the ssh session from certain IPs.
