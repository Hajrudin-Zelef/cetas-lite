---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-63607-ubiquiti-aps-on-manged-or-unmanaged-switch-dc2925a7
title: "questions-63607-ubiquiti-aps-on-manged-or-unmanaged-switch-dc2925a7"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-63607-ubiquiti-aps-on-manged-or-unmanaged-switch-dc2925a7.md
source_anchor: ""
source_lines: [1, 10]
sha256: 5629bbb7a4a9222f19fbe4d829ef9b013d5e9b3d74ae27a0de624e1fc0899e3b
---

# questions-63607-ubiquiti-aps-on-manged-or-unmanaged-switch-dc2925a7

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have three APs, one UAP-AC-IW and two UAP-AC-LITE. I have a Netgear Smart Managed Plus Switch, mainly for VLANs with plenty of open ports. I need to either add an unmanaged POE switch or a POE injector for the APs. Is there any benefit to having each AP plugged into the managed switch through the POE injector over each AP plugged into the unmanaged switch? Are there any downsides? I'm leaning toward the POE injector as that gives more flexibility. I'm assuming the injector route would also maximize bandwidth speed per AP?
The APs broadcast 4 different SSIDs, all on different VLANs. AP administration is handled on a physical connection (not over wireless).
If you use an unmanaged switch to a WAP that uses VLANs, you have no idea what will happen until you do. There is no standard for what an unmanaged switch will do with a VLAN tagged frame. Some unmanaged switches will drop the frame as damaged, some will strip off the VLAN tag and forward the frame, and some will simply pass the frame as is. Two out of the three options are bad, and you will not know until you try it.
A PoE injector will not look at the frames, and it will simply pass on all the layer-1 signals it receives.
